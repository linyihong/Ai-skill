# Observation — Phase 3 execution-chain station ledger

**Run ID**：2026-09-22-phase-3-chain-station-ledger  
**Kind**：去敏 station 分類（**不是**完整 EDR；虛構 YAML **不算**）  
**Extends**：[`2026-09-18-real-run-promotion-gaps.md`](2026-09-18-real-run-promotion-gaps.md)、[`2026-09-22-watermark-exclusion-is-projection.md`](2026-09-22-watermark-exclusion-is-projection.md)

對照 [`execution-flow.md`](../../../workflow/narrative-video-production/execution-flow.md)。本庫沒有外部媒體／matching 原始檔；只記每站能不能宣稱 pass。

## Lifecycle 主鏈

| Stage | 本庫狀態 | 分類 |
| --- | --- | --- |
| 0 Frame | 真實敘事短片 source-analysis 已發生 | pass（子鏈） |
| 1 Intake | 未回寫 brief lock | `data_insufficient` |
| 1b Bible | 無去敏 bible_id／實體表 | `data_insufficient` |
| 1c Catalog | 無既有 `clip_id` 入庫 | `data_insufficient` |
| 2 Template | 無主 `narrative_template_id` | `data_insufficient` |
| 3 Matching | 無可行集＋`selection.policy` | `data_insufficient` |
| 4 EDR | 無真實 EDR | `data_insufficient` |
| 7 Assemble | 無成片對 shot | 未開始 |
| 7b Locale | 無三閘 decision | `data_insufficient`（Q6 仍 open） |
| 8 Publish | 無獨立 verifier | 未開始 |
| 9 Outcome | 可停 `insufficient_sample`，但尚未開窗 | 未開始 |

因此 **Phase 3 未 PASS**。卡住點是「還沒走 matching／EDR」，不是 invariant 互斥（非 `design_error`）。

## Source-analysis 子鏈（已跑）

| 站 | 狀態 | 分類 |
| --- | --- | --- |
| observable OCR／ASR／face／unit | 可繼續 | pass |
| watermark role | 判定合理 | — |
| watermark 當刪除／crop／contains | 切掉有效專名 | `watermark_as_deletion`（contract_gap candidate） |
| spoken ≠ subtitle／phonetic／alert | 已記觀察 | 尚未在 consumer 落地 |
| text resolution 進 narrative | 未消費 | `text_resolution_not_consumed` |
| story promotion | fail | `promotion_gate_gap` |
| vocative／OCR mention → identity | fail | `mention_as_identity` |

## 下一輪外部專案只帶四件事回來（去敏）

1. 至少一 shot：Need → constraints → feasible `clip_id[]` → `selection.policy` → selected ∈ 可行集  
2. 對應 EDR shot 列（可 cut-ready，不必 publish-ready）  
3. 同一 shot 的 OCR span：watermark **retained** + scene／subtitle **eligible**（驗證不再切前綴專名）  
4. 同一 shot 的 spoken／subtitle 分欄；watermark STOP name learning  

仍不加 detector、不改 Phase 2 workflow、不註冊 route。
