# Phase 2 Artifact Contracts（draft；非 workflow）

本索引回答：**要留下什麼資料才能判斷 Phase 1 scenario 描述的行為？**
YAML 在 [`contracts/`](contracts/)（資料檔，不是 plan companion `.md`）。

不是第三份規則。三層分工：

| 層 | 回答 | 現況 |
| --- | --- | --- |
| Phase 1 `validation/scenarios/` | 要證明什麼行為？ | 已寫；detector / workflow 仍應 FAIL |
| **`contracts/*.yaml`** | 留下哪些欄位才能判斷？ | Phase 2 |
| Phase 3 `workflow/3d-character-production/` | 執行時如何組織 contract？ | **未授權**；禁止為了測綠而建立 |

Phase 3 只應 **promote／引用** 這些欄位，不得再抄一套 heuristic。Invalidation 只在
[`contracts/identity-acceptance.yaml`](contracts/identity-acceptance.yaml) 一處。

`runtime_projection.enabled: false`。不得標 `executable-contract/v1`。

| Record | 檔案 | 對應 scenario（只連 id，不複製 then） |
| --- | --- | --- |
| Character Specification | [`contracts/character-specification.yaml`](contracts/character-specification.yaml) | specification-lock-blocks-generation-v1 |
| Reference Set | [`contracts/reference-set.yaml`](contracts/reference-set.yaml) | 同上 |
| Identity Acceptance + invalidation | [`contracts/identity-acceptance.yaml`](contracts/identity-acceptance.yaml) | identity-acceptance-blocks-mesh-v1、identity-downstream-mutation-invalidation-v1 |
| Candidate Record | [`contracts/candidate-record.yaml`](contracts/candidate-record.yaml) | provenance-blocks-promotion-v1 |
| Mesh QA Report | [`contracts/mesh-qa-report.yaml`](contracts/mesh-qa-report.yaml) | mesh-gate-blocks-rig-v1 |
| Deformation Acceptance Set | [`contracts/deformation-acceptance-set.yaml`](contracts/deformation-acceptance-set.yaml) | deformation-gate-blocks-animation-v1 |
| Runtime-ready Character Pack | [`contracts/runtime-ready-character-pack.yaml`](contracts/runtime-ready-character-pack.yaml) | export-runtime-readback-v1 |
| Stage eligibility + rollback owner | [`contracts/artifact-gates.yaml`](contracts/artifact-gates.yaml) | 全部 stage-gate；只列「用哪些欄位裁決」 |
