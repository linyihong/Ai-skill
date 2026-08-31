# 3D Character Production — Execution Flow

Canonical lifecycle。各 stage 填哪個 record、能否推進：見
[`artifact-gates.md`](artifact-gates.md) 與 [`records/`](records/README.md)。
**不要**在本檔寫 mutation class → re-review 對照表。

## Lifecycle

```text
0. Frame                 → 確認可驅動 3D 角色資產（非 ML、非純圖／片）
1. Intake / Lock         → Character Specification + Reference Set
2. Candidate             → Candidate Record（accepted 與 rejected 都留）
3. Identity Acceptance   → 填 records/identity-acceptance.yaml
4. Mesh QA               → 僅當 identity.decision==accepted 且 validity==current
5. Rig / Deformation     → 僅當 mesh_qa.decision==pass 且 identity.validity==current
6. Face / Outfit / Motion→ 僅當 deformation.decision==pass
7. Export / consumption  → 填 runtime-ready pack；export_ok 不得單獨 completion
8. Validate + Close      → artifact-gates；fresh reviewer 才能 runtime-ready
```

下游改資產後：**寫 `mutation_event`，套用 identity-acceptance.yaml 的 rules，更新 `validity`。**
Gate 只讀更新後的 `validity`／`decision`。

## Maturity（Phase 0.2 凍結）

| 級 | 意義 |
| --- | --- |
| exploration | 可無完整 gate；不得宣稱產品可用 |
| prototype | 可開啟、可基本動作；允許未分類缺陷；不得 completion |
| runtime-ready | 全部 blocking gates PASS + runtime contract PASS + known defects 已分類 |
| runtime-ready with known defects | 僅 non-blocking；blocking 未過仍是 prototype |

`exploration`／`prototype` 可省略部分 gate；**completion claim 必須完整 blocking set**。

## Stage 明細

| Stage | 填／讀 | 推進條件（欄位，非 heuristic） | 失敗 rollback |
| --- | --- | --- | --- |
| 0 Frame | 本檔 | 任務是可驅動 3D 角色資產 | — |
| 1 Lock | [`intake.md`](intake.md) | spec + reference lock 完整 | intake_or_spec_author |
| 2 Candidate | [`candidate-generation.md`](candidate-generation.md) | decision+reason 齊才能 promotion | candidate_generation |
| 3 Identity | [`records/identity-acceptance.yaml`](records/identity-acceptance.yaml) | `decision==accepted` 且 `blocking` 空 且 `validity==current` | candidate_generation |
| 4 Mesh | [`mesh-quality.md`](mesh-quality.md) | `mesh_qa.decision==pass` 且 identity `validity==current` | mesh_or_candidate |
| 5 Deform | [`rigging-and-deformation.md`](rigging-and-deformation.md) | `deformation.decision==pass` | rig_weights |
| 6 Face／Outfit | 對應 md | deformation pass；outfit 不重寫 identity rules | 見 gates |
| 7 Export | [`export-and-runtime-validation.md`](export-and-runtime-validation.md) | consumption readback 非 fail；fresh_reviewer 才 completion | failing_readback_stage |
| 8 Close | [`artifact-gates.md`](artifact-gates.md) | 不得用自驗當 runtime-ready | — |

## 禁止

- 在 workflow 寫 `if hair_reconstructed: re_review`（應寫 mutation_event → contract → validity）。
- 歷史 `decision: accepted` 且 `validity != current` 仍推進。
- 無觀察卻填 PASS。
- 裸「AI 建模」當本 domain 充分條件（Phase 5 路由；執行時仍應先 Frame）。
