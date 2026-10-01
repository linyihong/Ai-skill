# Text evidence — OCR boundary recovery & script-aware normalize

何時讀：Latin／多語硬字幕 OCR 出現無空格長串、單框整句、或 `ocr_lang=ch` 卻讀到英文時；
在 Subtitle Grouping／Language Relation 之前。
欄位：[`records/text-evidence.yaml`](records/text-evidence.yaml)。
Plan：[`42-ocr-boundary-and-script-aware-normalization.md`](../../plans/active/2026-09-16-1649-narrative-video-production-workflow/42-ocr-boundary-and-script-aware-normalization.md)。
銜接：[`text-evidence-regions.md`](text-evidence-regions.md)。

> **執行契約，不重新解釋契約。** 不新增 LLM 斷詞；raw 不可被 derived 覆蓋。

## 在 lifecycle 的位置

```text
OCR Detection
  → Raw OCR Evidence
  → Text Segmentation / Boundary Recovery
       (multi-box space | intra-box word-gap | script-aware normalize)
  → Normalized Text Evidence
  → Subtitle Grouping → Language Relation → Text Resolution
```

對應 execution-flow **6b** 最前段（region／group 之前）。

## `recognition_language` ≠ `observed_script`

| 欄位 | 用途 |
| --- | --- |
| `recognition_language` | 引擎模型語言（adapter 配置） |
| `observed_script` / `language_candidate` | 觀測字串的 script（機械） |

禁止：用 recognition_language 決定「是否刪空格」。Latin 觀測一律保留詞界策略；CJK 可去空白。

## Boundary

| status | 含義 |
| --- | --- |
| `ok` | 多 box 已有合理空格、CJK 無需詞空格、或整段是合法單字（`Unexpectedly`） |
| `suspicious` | Latin 且機械判定可能缺詞界（非「凡 Latin 都拆」） |
| `recovered` | 已產出 derived candidate（須標 `method`） |
| `unresolved` | 無法可靠恢復 |

### Suspicious 閘門（先於 recovery）

**不要** `Latin → split`。至少一項成立才 `suspicious`：

- 單框／單 part 長串 + `space_count=0`，且 **非整段 dictionary exact-match**
- mid-capital／字母數字交界（`WhenI`）
- 字元間距異常（有 char／ink evidence 時）
- lexical 可全覆蓋拆成 ≥2 詞（僅作 **觸發**，不是真理）
- 與 ASR／context 明顯衝突（可選）

合法單字（`Unexpectedly`／`uncomfortable`／`volunteer`）→ `ok`，不進拆詞。

### Layered recovery（座標優先於 closed-class）

```text
OCR raw evidence
  → Script detection
  → Latin + boundary_check.suspicious?
       │
       ├─ 1. Existing OCR boxes（word／part／char boxes 若有）
       ├─ 2. Intra-box geometry（ink／connected-components／gap÷median_char_width）
       ├─ 3. Lexical boundary candidates（closed-class + content lexicon；candidate only）
       ├─ 4. ASR phonetic + subtitle context（candidate）
       ↓
  Candidate set → resolver／QC（必要時 LLM）
```

| 層 | `method`（derived） | 角色 |
| --- | --- | --- |
| multi-box | `multi_box_space` | 已有多 box／part 空格 |
| geometry | `geometry_word_gap`／`ink_projection` | **第一刀 recovery**；畫面 evidence |
| lexical | `lexical_candidate`／`closed_class_candidate`／`orthographic_mid_capital` | **候選產生器**，非 truth generator |
| ASR／context | （下游 multimodal） | 加候選／重排；不直接覆寫 raw |

**整句 bbox 只有一個 box 時，「有座標」≠「座標足以拆詞」** → 必須做 **intra_box_boundary_recovery**（對 bbox 再取字元／墨水間距），不是跳過幾何直接 closed-class。

Closed-class（`a`／`to`／`for`／`is`／`her`…）保留，但只當 lexical candidate generator：
`aghostis`→`a ghost is`、`tomasturbate`→`to masturbate` 是好候選，不是第一刀真理。

Regression corpus（驗證機械能力，**禁止** `if raw == "…"` 特例）：
[`records/latin-boundary-regression.yaml`](records/latin-boundary-regression.yaml)。

## Raw / derived

必保留 `raw_text`。`normalized_text` 或 `derived.candidates[]` 帶 `method`＋`status=candidate`。
下游 Text Resolution／locale 優先消費 normalized／recovered，但 audit 可回看 raw。

**`parts[]` 是較底層 evidence；整段 `text` 是 derived projection。**
禁止用黏掉的 derived `text` 回頭覆蓋／刪除已切開的 `parts`。join 演算法可改 → **只重跑 projection**，不必重 OCR。

```text
L0 OCR parts
  → script-aware join（token 接縫）
  → derived.text  (method: script_aware_join)
```

## Script-aware join（token boundary，非 accumulated script）

判斷**相鄰兩個 token 接縫**，禁止用「目前累積字串是 mixed／CJK」決定後面全部怎麼接：

| 接縫 | 預設 |
| --- | --- |
| Latin \| Latin | 插入 `" "` |
| Latin \| CJK | policy（硬字幕常插空或依 profile） |
| CJK \| Latin | policy（同上） |
| CJK \| CJK | `""` |

**禁止：** `observed_script(cumulative_out) == mixed` → 否決後續 Latin\|Latin 空格
（典型壞例：`parts=[I'm,from,a,pet,store]` → `I'm fromapetstore`）。

單 part 內仍黏（如 `appointmentfortoday`）屬 **intra-box boundary recovery**
（boxes → geometry → lexical candidate → ASR／context → resolver），與 join 分開；
不得一開始就 LLM 斷詞，也不得跳過幾何直接 closed-class。

雙語疊字：先 `subtitle_group` 分 zh／en region，各自 script-aware normalize，再進 timeline——勿把中英 parts 先併成一大串再拆。

分類：parts 已切、join 黏壞 = **adapter／implementation defect**（非新 Phase）。

## Quality（建議）

`frame_width`／`frame_height`／`text_height_px`／`boundary_confidence` —
低解析＋單框 Latin 應標 low，避免下游當高信心字面。

## 推進條件

| 條件 | 失敗 |
| --- | --- |
| Latin 長串無空格時有 boundary 標記或 recovery 嘗試紀錄 | 靜默接受黏字串當唯一 SoT |
| suspicious 閘門：合法單字 exact-match 不拆 | `Unexpectedly` 被強制拆詞 |
| recovery 順序 geometry → lexical；closed-class 僅 candidate | 跳過幾何／詞表當第一刀真理 |
| derived 不覆蓋 raw／parts | raw 或 parts 被改寫／刪除 |
| normalize 依 observed_script | `ocr_lang=ch` 刪光 Latin 空格 |
| join 依 token 接縫 | 用 cumulative mixed／CJK 否決 Latin\|Latin 空格 |
| CJK+Latin+Latin+… 可機械重建空格 | `I'm fromapetstore` 類 derived 無標記仍當 SoT |
| dogfood 黏串進 regression corpus | 用 `if raw==…` 特例代替能力驗證 |
