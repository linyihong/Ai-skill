# Candidate: Language & text-role before Text Resolution

Companion to [`19-text-resolution-and-narrative-assembly.md`](19-text-resolution-and-narrative-assembly.md)、
[`21-spoken-vs-subtitle-reconstruction.md`](21-spoken-vs-subtitle-reconstruction.md)、
[`23-sanitization-anomaly-audit.md`](23-sanitization-anomaly-audit.md)。
**不是新 Phase、不是新 Agent、不推翻 Phase 1/2。**
觀察：[`evidence/2026-09-29-cross-language-subtitle-spoken-alignment.md`](evidence/2026-09-29-cross-language-subtitle-spoken-alignment.md)。
Workflow 落點：[`workflow/narrative-video-production/text-evidence-language-relation.md`](../../../workflow/narrative-video-production/text-evidence-language-relation.md)。

## 問題

OCR 英文硬字幕與 ASR 中文口播**可以同時正確**。若直接做「字面 conflict／OCR 優先／和諧嫌疑」，會把正常的跨語言字幕對齊誤判成 anomaly，並把英文硬字幕寫進 spoken canonical。

這是 **locale／language relation** 缺口，不是單純 OCR vs ASR 誰優先。

## 吸收策略（優先既有欄位）

優先用既有候選：`spoken_text`、`subtitle_text`、`text_relation`、`text_span_role`、
`visual_text_evidence.role`、locale 三閘（content ≠ timing ≠ layout）。
每個 evidence 補齊：

| 通道 | 必記 |
| --- | --- |
| OCR → observed_text | `language`、`locale`（若可知）、`text_role`（subtitle／signage／watermark／…） |
| ASR → spoken_text | `language`、`locale`（若可知）、`speaker`（若有） |

`source_locale`（片／口播來源）≠ `display_locale`（字幕／目標語系 pack）。

## 前置管線（插在 Text Resolution 前）

```text
OCR / ASR
  → Language & Text-Role Detection   (mechanical first)
  → Evidence Alignment               (time / window)
  → Language Relation Gate
  → Text Resolution                 (selection / reconstruction)
  → Narrative Analysis
```

LLM 不得單獨宣告「OCR 英文所以是翻譯」。機械先建：

- `ocr.language = en`、`asr.language = zh`
- `candidate_relation = cross_language`

再由 LLM 只做 **semantic alignment support / unsupported**（Selection／Interpretation Actor）。

## Language Relation Gate

```text
Same language?
  ├── YES → normal text-resolution
  │         (match / variant / sanitization_candidate / unresolved)
  └── NO  → Cross-language alignment
            ├── translation/subtitle relation supported → 不進 sanitization
            └── unsupported → cross_language_unresolved（anomaly 另議）
```

## `text_relation.type`（建議四＋一）

| type | 例 |
| --- | --- |
| `same_language_match` | OCR／ASR 同語言且對齊 |
| `same_language_variant` | 同語言近音／錯字 → phonetic reconstruction |
| `cross_language_translation` | EN subtitle ↔ ZH spoken 且語義對齊 |
| `sanitization_candidate` | **同語言**語意衝突（顯示詞 vs 口播） |
| `unresolved` | 對齊或語言不足以判定 |

禁止：跨語言對齊直接標成 `OCR_ASR_CONFLICT` 或 `subtitle_sanitization`。

## Sanitization 前置條件

Sanitization Detection **必須**先過 Language Relation Gate。
「Little sweetheart」vs「你這個禽獸」：若 OCR=`en`／ASR=`zh` 且為 translation relation → **不**判和諧；只有 same-language semantic conflict 才進 sanitization hypothesis。

## OCR `text_role`（Evidence，不是刪除）

`dialogue_caption`／`translated_subtitle`／`hardcoded_subtitle`／`signage`／`UI_text`／`watermark`／`title`。
例：畫面 “Tokyo Airport” + ASR「東京機場」→ 可能是 signage／place，不得自動當 dialogue subtitle translation。role 只 candidate；watermark 仍走 projection 不刪 evidence（見 [`24-ocr-role-projection.md`](24-ocr-role-projection.md)）。

## 與 Translation Decision／locale 三閘

對齊 [`2026-09-22-1000-translation-decision-workflow`](../2026-09-22-1000-translation-decision-workflow/_plan.md)：Evidence 層先分 language／role，再談怎麼翻。Locale pack 的 content gate 消費 **resolved spoken／subtitle 對**，不得直接吃 raw OCR 拉丁硬字幕當中文口播源。

## Phase 3 分類

| 分類 | 含義 |
| --- | --- |
| `cross_language_collapsed_to_conflict` | 跨語言對齊被当成 conflict／OCR 勝出 | **contract_gap candidate**；優先既有欄位吸收（本檔＋workflow 落點）；不開 Q12/Q13、不加 Phase |

本 round：記缺口＋workflow／companion 吸收；產品 adapter 另改，不得反向定義本契約。
