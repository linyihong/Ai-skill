# 3D Character Production Workflow

`workflow/3d-character-production/` 負責**可驅動 3D 角色資產**的生產順序與 gate。
AI 生成／手工／外包是 acquisition strategy，不是 domain 名稱。

> **狀態**：Phase 3 workflow core。**沒有** `route.workflow.3d-character-production`（Phase 5）。
> YAML **不**投影（`runtime_projection.enabled: false`）。Detector 應維持 no_match。

> **執行契約，不重新解釋契約。** Heuristic 在 Phase 1 scenarios；欄位在
> [`records/`](records/README.md)。Invalidation **唯一 SoT**：
> [`records/identity-acceptance.yaml`](records/identity-acceptance.yaml)。

## 何時讀哪個檔

| 認知階段 | 檔案 | load_when |
| --- | --- | --- |
| Lifecycle | [`execution-flow.md`](execution-flow.md) | 任何本 domain 任務 |
| Intake / Lock | [`intake.md`](intake.md) | Specification／Reference 未鎖 |
| Spec 欄位 | [`character-specification.md`](character-specification.md) | 填 Lock |
| Reference 欄位 | [`reference-governance.md`](reference-governance.md) | 填 Lock |
| 候選 | [`candidate-generation.md`](candidate-generation.md) | 生成／比較／promotion |
| Identity | [`records/identity-acceptance.yaml`](records/identity-acceptance.yaml) | 候選後、任何下游 mutation 後 |
| Mesh | [`mesh-quality.md`](mesh-quality.md) | Identity 可進 mesh_qa 之後 |
| Rig／變形 | [`rigging-and-deformation.md`](rigging-and-deformation.md) | mesh_qa pass 之後 |
| 表情 | [`facial-expression.md`](facial-expression.md) | body deformation pass 之後 |
| 服裝／動作 | [`outfit-and-animation.md`](outfit-and-animation.md) | facial-expression acceptance pass 之後 |
| 匯出／消費 | [`export-and-runtime-validation.md`](export-and-runtime-validation.md) | 宣稱完成前 |
| Eligibility | [`artifact-gates.md`](artifact-gates.md) | 每一 stage 推進與 completion |
| VRM profile | [`profiles/README.md`](profiles/README.md) | Phase 4 才寫 `vrm.md` |

## 核心原則

1. 填 [`records/`](records/) 欄位；markdown 不複製 `invalidation_rules`。
2. 下游 eligibility 讀 `identity_acceptance.validity == current` **且** `decision == accepted`（hold／rejected 不得當 pass）。
3. 無 evidence → 不得推導 PASS（absence = fail／hold）。
4. 可匯出 ≠ `runtime-ready`。
5. 工具名只出現在 consumer evidence 或未來 profile；本 core 不綁 DCC／生成服務。
6. 程式實作走 `software-delivery`；本 workflow 只答資產能否被正確消費。

## 與 software-delivery

資產 acceptance 在此；Blender Python／exporter／viewer **程式**在 `software-delivery`。
Inventory UI、選角 API、權限、同步不在本 domain。
