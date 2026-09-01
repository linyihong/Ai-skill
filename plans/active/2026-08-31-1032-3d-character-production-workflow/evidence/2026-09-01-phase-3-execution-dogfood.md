# Phase 3 execution dogfood — eleven stop/go cases

**Run ID**：2026-09-01-phase-3-execution-dogfood  
**Authorization**：user「跑一下 Dogfood」（等同口令「跑 execution dogfood」）  
**Scope**：只驗 [`2026-08-31-phase-3-implementation-landed.md`](2026-08-31-phase-3-implementation-landed.md) 十一條。未改 contract、未擴 workflow、未碰 Phase 4／5。

Filled instances 在 consumer `<PROJECT_ROOT>` 的 probe 目錄。本檔不含專案名、絕對路徑、素材檔名。

## 結論

**十一條執行結果均符合 contract。** 這不是 workflow 驗收完成，也不是 Phase 6 full dogfood，也不是 runtime-ready。

Detector `no_match` 仍 **EXPECTED**。

Live pack 仍是 **prototype**：identity `hold`、mesh geometry fail、body deformation fail、facial `partial`、`export_ok` 為 true 但 completion 不是 runtime-ready。

## 方法

Gate 只讀 `decision` / `validity` / 各 record 欄位（`records/artifact-gates.yaml`）。  
Mutation 只經 `records/identity-acceptance.yaml` `invalidation_rules` 得到 `identity_effect` → `validity`；gate 不讀 `mutation_class`。

Case 1–3、11 的 `accepted` 在 **live 填表不存在**（live identity 是 `hold`）。這幾條用 **標註為 overlay 的欄位組合** 驗證閘謂詞，沒有把 live identity 改成 accepted。

## 結果

| # | Case | 資料 | 預期 | 實測 | 判定 |
| --- | --- | --- | --- | --- | --- |
| 1 | accepted ∧ current | overlay | allow | allow（identity 閘） | PASS |
| 2 | accepted ∧ re_review_required | overlay | stop | stop | PASS |
| 3 | accepted ∧ stale | overlay | stop | stop | PASS |
| 4 | hold ∧ current | live identity | stop | stop | PASS |
| 5 | missing evidence | live mesh holes `not_inspected`；body poses `not_observed`；facial 非 pass | stop / no PASS | stop | PASS |
| 6 | export_ok only | live pack `export_ok: true` | prototype | maturity `prototype`；`fresh_reviewer: false` | PASS |
| 7 | mutation → Identity SoT → validity → gate | live `mesh_repair` + `visual_drift_on_fixed_views` | stop | SoT `re_review_required`；event 欄位一致；gate stop | PASS |
| 8 | facial mapping／visible = partial | live facial `decision: partial` | stop；不得進 outfit | stop | PASS |
| 9 | UV/material fail，無對應 named defer | live `uv_material: fail` | diagnostic；maturity ≤ prototype | 同左；pack 無 UV defer | PASS |
| 10 | helper／overlay visible，real face 不完整 | live forbidden `overlay_only` + `geometry_delta_without_visible_readback` | stop；不得 facial PASS | `decision` 非 pass | PASS |
| 11 | export 其餘假設 pass，identity stale | overlay | stop；不得 runtime-ready | identity 閘 stop | PASS |

## Case 7 鏈

```text
mutation_event (mesh_repair, visual_drift_on_fixed_views)
      ↓
identity-acceptance.yaml invalidation_rules
      ↓
identity_effect = re_review_required → validity
      ↓
artifact-gates identity_eligible 只讀 decision + validity → stop
```

未在 gate 寫 `if mesh_repair → re_review`。

## 不是什麼

- 不是把 live 角色升成 `accepted` 或 `runtime-ready`
- 不是 Phase 4 VRM profile
- 不是 Phase 5 route／glossary／projection
- 不是 Phase 6 從 maturity 重新進場的完整 dogfood
