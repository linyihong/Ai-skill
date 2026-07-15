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
| Blind LLM inferability | ✅ **round 2 complete**（8 subagents · simple prompt · see below） |
| `payment_leave_transition.yaml` | ✅ landed（I-05 blind + rule-trace trigger） |
| RC2-P2 Exit gate | ❌ **not met** — blind Boundary Misclassification **2**（I-07, I-08） |

---

## Blind LLM round 2（2026-07-15）

**Method**: 8 isolated subagents（`gemini-2.5-flash`）· each received **only** one incident prompt + `preview_gate_transition.yaml` · **no** scenarios expected fields · **no** intake.

**Normalizer**: map free-text layer → canonical `{Interaction, Composition, Continuation, Navigation, Pagination_runtime, Pattern, Runtime}`.

### Blind matrix

| ID | Blind layer (raw → canonical) | Blind entry | Expected layer | Expected entry | Layer | Entry |
| --- | --- | --- | --- | --- | --- | --- |
| I-01 | Interaction Knowledge → **Interaction** | `preview_gate_transition` | Interaction | `preview_gate_transition` | ✅ | ✅ |
| I-02 | ui-interaction-knowledge → **Interaction** | `preview_gate_transition` (via invalidation_event) | Interaction | `preview_gate_transition` | ✅ | ✅ |
| I-03 | player interaction logic → **Interaction** | ambiguous literal | Interaction | `preview_gate_transition` | ✅ | ⚠️ |
| I-04 | ui-interaction → **Interaction** | ambiguous literal | Interaction | `preview_gate_transition` | ✅ | ⚠️ |
| I-05 | ui-interaction → **Interaction** | `null`（拒絕 preview） | Interaction | `payment_leave_transition` | ✅ | ⚠️ new entry implied, id 未說出 |
| I-06 | UI Layering → **Composition** | `null` | Composition | — | ✅ | ✅ |
| I-07 | UI_Navigation_State → **Navigation** | `null` | Continuation | — | ❌ | ✅ |
| I-08 | ui-interaction-knowledge → **Interaction** | `new_entry_required` | Pagination_runtime | — | ❌ | ✅ |

**Blind layer accuracy**: **6 / 8**（strict canonical）· **7 / 8**（I-03 entry ambiguous 不計 layer）  
**Boundary Misclassification**: **2**（I-07 Navigation≠Continuation · I-08 Interaction≠Pagination_runtime）  
**IH1 entry（Interaction positives）**: I-01–I-04 → `preview_gate_transition`；I-05 正確拒絕 preview → 觸發第二 entry landing

### P2 Evidence Dashboard（blind round 2）

| Metric | rule-trace | blind LLM |
| --- | --- | --- |
| Layer Classification Accuracy | 8/8 | **6/8** |
| Boundary Misclassification | 0 | **2** |
| Existing Entry Reuse (I-01–I-04) | 4/4 | 2 clear + 2 ambiguous |
| New Entry Required (I-05) | confirmed | confirmed（null entry） |
| Frozen Layer Mods | 0 | 0 |

### Hypothesis results（cumulative）

| Hypothesis | rule-trace | blind LLM | Verdict |
| --- | --- | --- | --- |
| IH1 Inferability | partial | partial | I-05 需命名 `payment_leave_transition` — entry landed post-blind |
| IH2 Boundary | PASS | **FAIL** | I-07/I-08 誤吸 Interaction/Navigation |
| IH3 Repair localization | PASS | PASS | 無 Pattern/Composition 回改建議 |
| Frozen Layer Mods | 0 | 0 | held |

### Knowledge Layer Confusion Matrix（blind — normalized）

| Actual ↓ / Predicted → | Interaction | Composition | Continuation | Navigation | Pagination_runtime |
| --- | --- | --- | --- | --- | --- |
| **Interaction** | 5 | 0 | 0 | 0 | 0 |
| **Composition** | 0 | 1 | 0 | 0 | 0 |
| **Continuation** | 0 | 0 | 0 | **1** | 0 |
| **Pagination_runtime** | **1** | 0 | 0 | 0 | 0 |

---

## Disposition（updated）

---

## Explicit non-actions

- [x] No frozen Pattern/Composition edits
- [x] `payment_leave_transition.yaml` landed（I-05 blind + rule-trace trigger）
- [x] Blind LLM round 2 documented
- [ ] RC2-P2 Exit — blocked on blind Boundary Misclassification = 0
