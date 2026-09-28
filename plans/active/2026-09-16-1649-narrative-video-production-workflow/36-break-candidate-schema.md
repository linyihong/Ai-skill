# Candidate: Phase 1 break-candidate schema

Companion to [`35-break-candidate-system.md`](35-break-candidate-system.md)。**欄位已凍**
（[`break-candidate.yaml`](../../../workflow/narrative-video-production/records/break-candidate.yaml)）。  
觀察：[`evidence/2026-09-28-break-candidate-schema.md`](evidence/2026-09-28-break-candidate-schema.md)。

Phase 1 只定 schema，不實作 LLM 選擇。四個責任：

1. `BreakPolicy` — 語系的 mechanical constraints／features。
2. `generate_break_candidates` — 找出可能斷點，不做最終決策。
3. `hard_violation` — 唯一能把候選趕出可行集的條件。
4. `best_cut` — 沒有 Selection Actor 時，去掉 violation 後取最小 `score`。

相對初稿的四個修正：`readability` → `balance_score`；`hard_penalty` → `hard_violation`；
Phase 1 沒有 `source`；`semantic_boundary` 是 `unknown | boundary | non_boundary`。
`phrase_integrity` 不得單獨排除候選。`window` 是搜尋範圍，候選不得只圍在 `prefer_at` 旁邊。
