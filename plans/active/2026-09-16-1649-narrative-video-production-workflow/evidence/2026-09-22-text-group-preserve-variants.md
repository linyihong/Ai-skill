# Observation — similarity merge is not identity

**Run ID**：2026-09-22-text-group-preserve-variants  
**Kind**：Phase 3 `contract_gap` **candidate**（`text_alignment`／`candidate_grouping`；**不是**把 threshold 從 0.9 調到 0.95）  
**Extends**：[`2026-09-22-watermark-exclusion-is-projection.md`](2026-09-22-watermark-exclusion-is-projection.md)、[`2026-09-21-spoken-vs-subtitle-reconstruction.md`](2026-09-21-spoken-vs-subtitle-reconstruction.md)、[`2026-09-21-phonetic-text-reconstruction.md`](2026-09-21-phonetic-text-reconstruction.md)

## 案例（去敏）

同一時間三筆近形：完整句、末 token 衝突、完整句 prefix。目前 similarity＋同 timestamp 只留一筆。應留 canonical＋variants；三條 raw 不消失。Prefix／末字衝突用 span relation（`variant_of`、`truncated_variant_of`），不是刪。

## 四類 alignment

| 類 | 處置 |
| --- | --- |
| duplicate | 字面全同 → evidence 層可 destructive merge |
| variant | 局部差異 → preserve |
| truncated | 短的是長的 prefix → preserve |
| conflict | 差異 span 可能有語意（地名／物件）→ resolution，不得 similarity 吃掉 |

高相似 ≠ duplicate。例：去甲地 vs 去乙地。

## Pipeline

raw → temporal alignment（group）→ canonical selection → variants 保留 → phonetic／context → LLM on **group** → verify → canonical spoken → **然後** translation。禁止用 dst 反推哪個 source 相同（翻譯會把誤識補成完整語意）。

候選 invariant：Only exact duplicates may be destructively merged at the evidence layer. Differing spans that may carry meaning MUST stay candidates.
