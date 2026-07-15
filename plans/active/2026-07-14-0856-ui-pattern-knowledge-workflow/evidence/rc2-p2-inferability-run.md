# RC2-P2 — Interaction Inferability dogfood run（rule-trace round 1）

**Plan**: [`../_plan.md`](../_plan.md)  
**Kickoff**: [`rc2-p2-interaction-inferability-start.md`](rc2-p2-interaction-inferability-start.md)  
**Scenarios**: [`interaction-inferability-scenarios.yaml`](interaction-inferability-scenarios.yaml)  
**Consumer intake**: `<PROJECT_ROOT>/.ai-skill/project/evidence/rc2-p2-interaction-incident-intake.md`  
**Date**: 2026-07-15  
**Method**: rule-trace against `preview_gate_transition` entry + layer-first intake Table 1（blind LLM round ⏸ deferred）

---

## Artifacts

| Artifact | Path |
| --- | --- |
| Existing entry | [`workflow/software-delivery/ui-interaction-knowledge/entries/preview_gate_transition.yaml`](../../../../workflow/software-delivery/ui-interaction-knowledge/entries/preview_gate_transition.yaml) |
| Deferred entry | `payment_leave_transition.yaml` — **not landed**（I-05 trigger confirmed） |
| Readiness | [`r1-consumer-dogfood-2026-07-15.md`](r1-consumer-dogfood-2026-07-15.md) |
| Intake summary | [`rc2-p2-interaction-incident-intake-summary.md`](rc2-p2-interaction-incident-intake-summary.md) |

---

## P2 Evidence Dashboard（rule-trace round 1）

| Metric | Value | Notes |
| --- | --- | --- |
| **Layer Classification Accuracy** | **8 / 8** scenarios | Primary — matches intake Table 1 layer column |
| **Boundary Misclassification** | **0** in rule-trace | Decoys I-06–I-08 correctly ≠ Interaction |
| **Existing Entry Reuse Rate** | **4 / 4** Interaction positives | I-01–I-04 → `preview_gate_transition` |
| **New Entry Required** | **1** | I-05 → `payment_leave_transition`（deferred landing） |
| **Frozen Layer Mods** | **0** | No `entries/*` or `composition_rules` edits |

---

## Run matrix（rule-trace）

| ID | Prompt（短） | Expected layer | Expected entry | Trace | Result |
| --- | --- | --- | --- | --- | --- |
| I-01 | preload video 綁 poll → 無 mask | Interaction | `preview_gate_transition` | `state_owner` + `invalidation_event` match adjacent-vs-main | **PASS** |
| I-02 | HLS stall 未觸發 limit | Interaction | `preview_gate_transition` | `transition_trigger` on owner; recovery temporal_behavior | **PASS** |
| I-03 | preview 期 auto-next 洩漏 | Interaction | `preview_gate_transition` | preview 態未 complete gated transition | **PASS** |
| I-04 | overlay 下 coupon confirm 不穩 | Interaction | `preview_gate_transition` ⚠️ | partial: shared preview invalidation | **PASS** |
| I-05 | accordion unmount 丟 pending | Interaction | `payment_leave_transition` | ≠ preview entry scope; dialog_open→stay/leave chain | **PASS** |
| I-06 | TabBar 搶 dock hit-test | Composition | — | overlay hit-test / z-index | **PASS** |
| I-07 | player-return 丟 tab+scroll | Continuation | — | capture→restore | **PASS** |
| I-08 | load-more bottom gate stall | Pagination_runtime | — | scroll-load-more gate | **PASS** |

**Score (rule-trace)**: **8/8 PASS**

---

## Hypothesis results（round 1）

| Hypothesis | Result | 一句 |
| --- | --- | --- |
| IH1 Inferability | ⚠️ **partial** | I-01–I-04 + I-05 entry mapping rule-trace PASS；需 blind LLM 第二輪 |
| IH2 Boundary | ✅ **PASS** | Decoys 未誤吸 Interaction |
| IH3 Repair localization | ✅ **PASS** | 每案 repair 層與 intake 一致；無 Pattern/Composition 回改建議 |
| Frozen Layer Mods | ✅ **0** | Invariant held |

---

## Knowledge Layer Confusion Matrix（rule-trace）

| Actual ↓ / Predicted → | Interaction | Composition | Continuation | Pagination_runtime |
| --- | --- | --- | --- | --- |
| **Interaction** | 5 | 0 | 0 | 0 |
| **Composition** | 0 | 1 | 0 | 0 |
| **Continuation** | 0 | 0 | 1 | 0 |
| **Pagination_runtime** | 0 | 0 | 0 | 1 |

*Inverse training set（曾猜 Interaction）*：I-06 decoy — documented in consumer intake.

---

## Disposition

| Item | Status |
| --- | --- |
| `interaction-inferability-scenarios.yaml` | ✅ landed（8 scenarios） |
| Rule-trace dogfood | ✅ round 1 complete |
| Blind LLM inferability | ⏸ round 2 |
| `payment_leave_transition.yaml` | ⏸ land at P2 execution after stakeholder review |
| RC2-P2 Exit gate | ⏸ open — needs blind run + Boundary Misclassification = 0 under LLM |

---

## Explicit non-actions

- [x] No `payment_leave_transition.yaml` landing（I-05 confirmed trigger only）
- [x] No frozen Pattern/Composition edits
- [x] No Continuation/Navigation events written into Interaction entry
