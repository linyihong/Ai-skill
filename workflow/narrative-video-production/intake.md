# Intake — Brief lock

未鎖不得開匹配腳本、不得宣稱 cut-ready。

填 [`records/edit-decision-record.yaml`](records/edit-decision-record.yaml) 的 brief 段（可先獨立草稿，鎖定後抄入 EDR）：

- `platform`（發布面，不是工具名）
- `target_duration_s`
- `audience`
- `promise`（一句開場承諾）
- `forbidden[]`（不得出現的題材／人物／商標）
- `source_locale`（建議；首輪語數量 = dogfood Q6）
- `profile`（可空；Q4 未凍）

`exploration` 可草擬腳本但不選 clip、不填 EDR `qc` 為 pass。

推進：brief 欄位齊（見 [`records/artifact-gates.yaml`](records/artifact-gates.yaml) `brief_lock`）。
失敗 rollback：`intake_author`。
