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
| locale 三閘 | 每個發布語三閘各自 decision；generated 口播 cue 時軸來自 speech artifact，且文字等於 speech unit；**timing 含 temporal integrity**（same-track 未明示不得 overlap；雙語 same-group 允許）；max_lines 不得當填滿目標；字級同時受 absolute／profile／max_delta；wrap lossless；斷點的 hard_violation 必須是 none，且不得為靠近 prefer_at 而跨過更高 boundary tier；cue typography 統一 | 不得 publish-ready |
| timeline projection | EDR 選定且要求發布的 caption／audio／clip 都在 Timeline IR 有 `source_ref`／`edr_ref` 實例；雙語 `subtitle_group` 完整；coverage 無未解釋的 `suspicious_omission` | 不得 publish-ready |
| caption artifact | burn／speedup／concat／ASS 轉換後，rendered cue 軸與 Caption Pack／Timeline IR 一致；源片硬字幕已 scrub／遮蓋，避免與新 burn 疊成兩套可見台詞 | 不得 publish-ready |
| completion | blocking 空 + `fresh_reviewer` 已填且獨立 | 不得 publish-ready |

作者自驗可推進到 cut-ready。**completion authority must be independent from the producer when claiming publish-ready.**
