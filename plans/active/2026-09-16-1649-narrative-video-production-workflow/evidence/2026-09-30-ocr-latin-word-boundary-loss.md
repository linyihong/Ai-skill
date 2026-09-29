# Observation — OCR Latin word-boundary loss

**Run ID**：2026-09-30-ocr-latin-word-boundary-loss
**Kind**：Phase 3 dogfood observation → **`contract_gap` candidate**（`ocr_latin_word_boundary_lost`）
**Plan**：[`../42-ocr-boundary-and-script-aware-normalization.md`](../42-ocr-boundary-and-script-aware-normalization.md)
**Extends**：[`2026-09-29-bilingual-hardsub-regions.md`](2026-09-29-bilingual-hardsub-regions.md)、[`2026-09-17-visual-text-evidence.md`](2026-09-17-visual-text-evidence.md)

## 去敏情境

參考成片硬燒 **英文**字幕；畫面可見詞間空白。OCR 快取／dialogue cue 卻成
無空格長串。片名／路徑／job id 留在 `<PROJECT_ROOT>`。

## 機械觀察

| 信號 | 例（去敏） |
| --- | --- |
| 畫面 | Latin words with visible gaps |
| OCR `text` / `parts` | 單一字串無 space |
| `boxes` | **box_count = 1**（整句一框） |
| 解析度 | 短邊偏小（如 ~480 寬） |
| 管線 | `recognition_language=ch` → 全域去空白政策 |

錯誤：把黏字串當最終 subtitle／spoken source，或用 LLM「斷詞」。

## 正確分流

1. 保留 `raw_text`
2. `observed_script=latin` + 單框 + 無空格 → `boundary_status=suspicious`
3. Intra-box geometry／ink-gap → derived candidate（不覆蓋 raw）
4. `recognition_language ≠ observed_script`：script-aware normalize／join
5. Lexical detokenize 僅 candidate

## 分類

| 標籤 | 判定 |
| --- | --- |
| `ocr_latin_word_boundary_lost` | **contract_gap candidate** |
| 不是 | 雙語 Agent／新 Phase／LLM 斷詞 stage |

吸收：[`42-…`](../42-ocr-boundary-and-script-aware-normalization.md)＋[`text-evidence-ocr-boundary.md`](../../../workflow/narrative-video-production/text-evidence-ocr-boundary.md)。
