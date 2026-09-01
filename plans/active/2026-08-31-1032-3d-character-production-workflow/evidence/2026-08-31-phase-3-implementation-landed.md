# Phase 3 implementation landed — 不是驗收完成

Implementation lock commit `5abffb7a`（`6567ae4e` = execution layer files）。

不是：3D character production workflow validated。
不是：Phase 4 VRM graduation。
不是：Phase 5 route／glossary／projection。

Detector `no_match` = **EXPECTED**（registry 刻意未登記）。

## 已凍結的執行不變量

- Identity 放行：`decision == accepted` **且** `validity == current`。
- `accepted + stale` / `accepted + re_review_required` / `hold + current` → stop。
- mutation → `identity-acceptance.yaml` 更新 validity → gates 只讀新 state。
- 無 evidence → 不得 PASS。
- `export_ok` only → prototype，不是 runtime-ready。

## 下一驗證（2026-09-01 gate-graph repair 後；未授權則不做）

Phase 3 **execution dogfood**（真實資料上停／放行），不是再擴寫 workflow 文件：

| # | Case | 預期 |
| --- | --- | --- |
| 1 | accepted ∧ current | allow |
| 2 | accepted ∧ re_review_required | stop |
| 3 | accepted ∧ stale | stop |
| 4 | hold ∧ current | stop |
| 5 | missing evidence | stop / no PASS |
| 6 | export_ok only | prototype |
| 7 | mutation → Identity SoT → validity → gate observes new state | stop 或 allow **只依更新後欄位** |
| 8 | facial mapping／visible evidence = partial | stop；不得進 outfit／animation |
| 9 | UV/material fail 或 not_evaluated + named defer | 只允許 diagnostic；maturity ≤ prototype |
| 10 | helper／overlay visible，但 real face geometry/readback 不完整 | stop；不得當 facial PASS |
| 11 | export readback 其餘 pass，但 identity stale／re_review_required | stop；不得 runtime-ready |

## Authorization gate（未說這句就不跑）

觸發口令：**「跑 execution dogfood」**。

授權後只驗上表十一條既有 semantics；**不**再改 contract、**不**擴 workflow、**不**碰 Phase 4／5。

第 7 條必須證明這條鏈，而不是 execution layer 自己判斷 mutation：

```text
mutation_event
      ↓
Identity SoT (identity-acceptance.yaml)
      ↓
state mutation (validity)
      ↓
artifact-gate (reads state only)
```
