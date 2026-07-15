# RC2 — Vocabulary Exit Review（post-P3）

**Plan**: [`../_plan.md`](../_plan.md)  
**Date**: 2026-07-15  
**Trigger**: RC2-P3 Closure — [`rc2-p3-interaction-composition-closure.md`](rc2-p3-interaction-composition-closure.md)  
**Prerequisite**: P1 gap log — [`rc2-p1-preview-gate-representability-run.md`](rc2-p1-preview-gate-representability-run.md) §Vocabulary gap log

---

## Review question

RC2 三階梯完成後，是否應將 P1 dogfood 記錄的候選 primitive **擴充進** `interaction-entry-schema.yaml`？

**Default**：**不擴** — 除非 ≥2 independent entries 證明四欄不足且無法以 composition constraint 表達。

---

## Candidate inventory

| Candidate | First observed | P1–P3 usage | Review decision | Rationale |
| --- | --- | --- | --- | --- |
| **`guard_condition`** | C1 `onEnded` guard during preview | Expressible via `invalidation_event` + `recovery_boundary` on `preview_gate_transition`；CH1/CH2 composition 覆蓋 confirm 衝突 | **reject** | 單案實作 guard ≠ schema primitive；擴欄會誘發 entry 膨脹 |
| **`rollback`** | P1 gap log | No observed need across 2 entries | **reject** | — |
| **`checkpoint`** | P1 gap log | No observed need | **reject** | — |
| **`priority`** | P1 gap log | No observed need | **reject** | — |
| **Fifth field（任意）** | — | Two entries + 2 composition rules sufficient | **reject** | Stable 維護期維持四欄 |

---

## What would trigger re-open

| Signal | Action |
| --- | --- |
| ≥2 entries 無法四欄表示且 composition rules 無法修 | Stakeholder review → schema RFC |
| Layer boundary break（Interaction 吸收 Composition） | 新 RC incident，非 vocabulary patch |
| Consumer dogfood 連續 2 案需同名第五欄 | 記錄於 gap log → 下一 review |

---

## Schema status

| Item | Status |
| --- | --- |
| `interaction-entry-schema.yaml` | **unchanged** — four core fields |
| **Schema Extensions** | **0** for RC2 full cycle |
| Vocabulary freeze | **continues** until re-open signal above |

---

## Relation to CH2/CH3

`guard_condition` 在 CH2 標為 `out_of_scope`（vocabulary freeze）；CH3 `waived` with `vocabulary_freeze_until_p3_exit_review`.  
**本 review 結論**：waive 維持 — **不** 因 exit review 自動升格為 schema field。

---

## Verdict

**Vocabulary Exit Review：PASS** — 四欄 vocabulary **validated for Stable**；無 mid-cycle schema 膨脹。  
記錄歸檔；**不** 修改 `entries/*.yaml` 或 `interaction-entry-schema.yaml`。
