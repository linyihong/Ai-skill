# Execution SoT for record fields

Phase 2 freeze（歷史、勿改 heuristic 分叉）：
`plans/active/2026-08-31-1032-3d-character-production-workflow/contracts/`

本目錄是 **Phase 3 起唯一 execution source of truth**。Phase 2 contracts 是歷史快照，
不得作為 scenario 或 stage eligibility 的 runtime input。`invalidation_rules` **只**存在於
[`identity-acceptance.yaml`](identity-acceptance.yaml)。Workflow markdown 與
[`../artifact-gates.yaml`](../artifact-gates.yaml) 只讀欄位（尤其 `validity`）。

`runtime_projection.enabled: false`。Phase 5 才考慮 executable-contract 投影。

跨域 pattern：[`traceable-evidence`](../../cross-cutting/traceable-evidence/README.md)。`candidate-record.yaml` 實作 TE1（`rejected_must_retain`）、TE2（`provenance`）；`artifact-gates.yaml` promotion 實作 TE4；`identity-acceptance.yaml` 的 `invalidation_rules` 是 TE3（watch）的依賴變更觸發實例。
