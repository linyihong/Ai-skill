# Candidate: Natural boundary before length balance

Companion to [`36-break-candidate-schema.md`](36-break-candidate-schema.md)。**workflow 已補**
（[`break-candidate.yaml`](../../../workflow/narrative-video-production/records/break-candidate.yaml)）。  
觀察：[`evidence/2026-09-28-natural-boundary-before-length.md`](evidence/2026-09-28-natural-boundary-before-length.md)。

句段切分與字幕換行分開。`。`／`，` 先建立 clause candidate；太短可合併，不是看到標點就強制切。
某個 clause 放不下，才在 unit 內產生 line-break candidates。`lexical_unit_split` 在 LLM 之前排除。

`natural_boundary_first`：strong > medium > lexical／phrase > semantic > weak。
`balance_score`、`prefer_at`、行長只在同一 tier 比較。`best_cut` 不得用全局最小 `score` 選中「电脑｜里」。
