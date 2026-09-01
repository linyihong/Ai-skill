# Phase 2 Artifact Contracts（historical snapshot；非 execution SoT）

本索引回答：**要留下什麼資料才能判斷 Phase 1 scenario 描述的行為？**
YAML 在 [`contracts/`](contracts/)（資料檔，不是 plan companion `.md`）。

不是第三份規則。三層分工：

| 層 | 回答 | 現況 |
| --- | --- | --- |
| Phase 1 `validation/scenarios/` | 要證明什麼行為？ | 已改讀 execution records；detector Phase 5 前仍應 no-match |
| **`contracts/*.yaml`** | Phase 2 當時留下哪些欄位？ | 歷史快照；不得作 runtime input |
| Phase 3 `workflow/3d-character-production/records/` | 執行時如何組織 contract？ | 唯一 execution SoT |

2026-09-01 起，新增或修正欄位只改
[`workflow/3d-character-production/records/`](../../../workflow/3d-character-production/records/)。
Identity invalidation 的 execution SoT 是該目錄的 `identity-acceptance.yaml`。

`runtime_projection.enabled: false`。不得標 `executable-contract/v1`。

| Phase 2 Record | 歷史快照 | 原始 scenario（Phase 3 已改讀 execution records） |
| --- | --- | --- |
| Character Specification | [`contracts/character-specification.yaml`](contracts/character-specification.yaml) | specification-lock-blocks-generation-v1 |
| Reference Set | [`contracts/reference-set.yaml`](contracts/reference-set.yaml) | 同上 |
| Identity Acceptance + invalidation | [`contracts/identity-acceptance.yaml`](contracts/identity-acceptance.yaml) | identity-acceptance-blocks-mesh-v1、identity-downstream-mutation-invalidation-v1 |
| Candidate Record | [`contracts/candidate-record.yaml`](contracts/candidate-record.yaml) | provenance-blocks-promotion-v1 |
| Mesh QA Report | [`contracts/mesh-qa-report.yaml`](contracts/mesh-qa-report.yaml) | mesh-gate-blocks-rig-v1 |
| Deformation Acceptance Set | [`contracts/deformation-acceptance-set.yaml`](contracts/deformation-acceptance-set.yaml) | deformation-gate-blocks-animation-v1 |
| Runtime-ready Character Pack | [`contracts/runtime-ready-character-pack.yaml`](contracts/runtime-ready-character-pack.yaml) | export-runtime-readback-v1 |
| Stage eligibility + rollback owner | [`contracts/artifact-gates.yaml`](contracts/artifact-gates.yaml) | 全部 stage-gate；只列「用哪些欄位裁決」 |

Facial Expression Acceptance 是 2026-09-01 Phase 3 correction 新增的 execution record，
不回寫 Phase 2 snapshot；見
[`workflow/3d-character-production/records/facial-expression-acceptance.yaml`](../../../workflow/3d-character-production/records/facial-expression-acceptance.yaml)。
