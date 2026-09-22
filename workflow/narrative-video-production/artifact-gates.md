# Artifact Gates

Eligibility **只讀 record 欄位**。對照表：[`records/artifact-gates.yaml`](records/artifact-gates.yaml)。

## Blocking

| Gate | 讀取 | 未過 |
| --- | --- | --- |
| brief lock | intake 必填 | 不得 matching |
| bible + catalog | `bible_id`；引用的 `clip_id` 都存在 | 不得 selected |
| template | 主 id ∈ catalog + `template_kind` | 不得開 EDR 為 cut-ready |
| matching | 每 shot 可行集非空（或明確 blocked）；`selection.policy`；selected ∈ 可行集 | 不得 assemble |
| EDR SoT | 結構化 EDR 存在 | 不得宣稱任何發布成熟度 |
| assemble | 成片軸 ↔ `shot_id`／`selected_clip_id` | 不得 cut-ready |
| locale 三閘 | 每個發布語三閘各自 decision；generated 口播 cue 時軸來自 speech artifact | 不得 publish-ready |
| completion | blocking 空 + `fresh_reviewer` 已填且獨立 | 不得 publish-ready |

作者自驗可推進到 cut-ready。**completion authority must be independent from the producer when claiming publish-ready.**
