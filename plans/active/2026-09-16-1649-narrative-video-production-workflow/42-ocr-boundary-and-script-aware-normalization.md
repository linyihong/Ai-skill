# Candidate: OCR boundary recovery & script-aware normalization

Companion to [`41-bilingual-ocr-regions-and-subtitle-groups.md`](41-bilingual-ocr-regions-and-subtitle-groups.md)、
[`09-visual-text-evidence.md`](09-visual-text-evidence.md)、
[`24-ocr-role-projection.md`](24-ocr-role-projection.md)。
**不是新 Phase、不是新 Agent、不開英文字幕專用 workflow、不推翻 Phase 1/2。**
觀察：[`evidence/2026-09-30-ocr-latin-word-boundary-loss.md`](evidence/2026-09-30-ocr-latin-word-boundary-loss.md)。
Workflow：[`text-evidence-ocr-boundary.md`](../../../workflow/narrative-video-production/text-evidence-ocr-boundary.md)；
欄位：[`records/text-evidence.yaml`](../../../workflow/narrative-video-production/records/text-evidence.yaml)。

## 問題

OCR 把「字串內容」與「詞邊界／排版結構」綁死。例：畫面有空格的 Latin 硬字幕，
engine 吐出單框 `Compensatecompensateme`，或 `ocr_lang=ch` 路徑用
`re.sub(r"\s+", "", text)` 刪光空白。這是 **OCR Evidence Acquisition / boundary**
缺口，不是雙語、也不是 Text Resolution、更不該交給 LLM 斷詞。

## 核心分離

| 概念 | 含義 |
| --- | --- |
| `recognition_language` | OCR **引擎／模型**語言（如 Paddle `ch`） |
| `observed_script` | **觀測文字**的 script family（latin／cjk／…） |

二者 **不可等同**。用 `ch` 模型讀英文硬字幕，不得因此對 Latin 去空白。

## 前置管線（插在 Subtitle Grouping 之前）

```text
OCR Detection → Raw OCR Evidence (text, bbox, script candidate, …)
  → Text Segmentation / Boundary Recovery
       multi-box spacing | intra-box word-gap | script-aware normalize
  → Normalized Text Evidence (derived; raw retained)
  → Subtitle Grouping → Language Relation → Text Resolution
```

禁止：新增 LLM 斷詞階段；dictionary 直接 overwrite raw。

## Raw vs derived（不可覆蓋）

```text
raw_text: "Compensatecompensateme"
normalized_text / derived.candidates[]:
  text: "Compensate compensate me"
  method: geometry_word_gap | multi_box_space | lexical_candidate
  status: candidate
boundary:
  status: ok | suspicious | recovered | unresolved
```

Derived 必須可重跑；演算法升級不得污染 raw。

## Boundary suspicious（機械）

同時多數滿足即可標 `boundary_status: suspicious` → 觸發 recovery：

- `observed_script = latin`
- `box_count = 1`（或 join 後仍無空白）
- 長串 + `space_count = 0`

優先：**geometry / ink projection / character boxes**；lexical／dictionary 只作 candidate。

## Evidence quality（可選欄位）

`frame_width`／`frame_height`／`text_height_px`／`boundary_confidence` —
低解析 Latin 單框應降低 boundary confidence，供下游勿盲目信任黏字串。

## Phase 3 分類

| 標籤 | 判定 |
| --- | --- |
| `ocr_latin_word_boundary_lost` | **contract_gap candidate**；機械 adapter 優先；schema 不足才回頭擴 contract |

不開 Q12/Q13、不加「英文斷詞 Agent」。
