# Export and Runtime Validation

填 [`records/runtime-ready-character-pack.yaml`](records/runtime-ready-character-pack.yaml)。

- `export_ok: true` 不是 completion。
- 資產消費：viewer_open、humanoid_mapping、expression_trigger、motion_play、outfit_switch。
- 不含 Inventory UI、選角 API、權限、同步。
- `fresh_reviewer: false` 不得標 `runtime-ready`。
- known defects 必須分類；blocking 未過 = prototype。
