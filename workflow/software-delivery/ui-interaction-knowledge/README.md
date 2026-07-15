# UI Interaction Knowledge — reusable rule seeds（Ai-skill）

**Research Cycle 2** · RC2-P1 ✅ **Closed** · RC2-P2 ✅ **Closed** · RC2-P3 ▶ **Active**

Interaction Knowledge describes the valid temporal lifecycle of UI state **after Pattern selection and Composition have been validated** — not another Composition layer, and **not application business workflow**.

> **Canonical boundary**: Interaction Knowledge describes **UI interaction semantics**, not application business workflow.  
> Example: `dialog_open` → `user_confirm` → `dialog_close` — **not** `order_paid` / payment-order pending.

| Layer | Owns |
| --- | --- |
| Pattern Knowledge（frozen） | What pattern is appropriate |
| Composition Knowledge（frozen） | Spatial / constraint edges between patterns |
| **Interaction Knowledge** | State ownership, transition trigger, invalidation, recovery |
| **Interaction Composition**（RC2-P3） | Constraints between interaction lifecycles on a Screen / flow |

## Maturity（RC2）

| RC2 | 狀態 |
| --- | --- |
| P1 Representability | ✅ Closed |
| P2 Inferability | ✅ Closed |
| P3 Composition | ▶ Active |
| Interaction Knowledge | 🟡 Research Justified（P3 exit → 🟢 Stable） |

## Layout

```text
ui-interaction-knowledge/
  README.md
  interaction_composition_rules.yaml   # RC2-P3 constraints（NOT entries）
  validation/
    interaction-entry-schema.yaml
  entries/
    preview_gate_transition.yaml
    payment_leave_transition.yaml
```

## RC2 Invariant

```text
Interaction evidence MUST NOT redefine Pattern selection or Composition constraints.
Interaction MUST NOT edit ui-pattern-knowledge/entries/* or composition_rules.yaml.
```

## RC2-P3 Invariant

```text
Interaction Composition evidence MUST NOT modify validated Interaction entries.
New composition knowledge flows into interaction_composition_rules.yaml ONLY.
```

## Vocabulary Freeze

**四欄 vocabulary 已於 RC2-P1 驗證**：`state_owner`, `transition_trigger`, `invalidation_event`, `recovery_boundary`

## RC2 Metrics（non-KPI）

| Metric | Target |
| --- | --- |
| **Interaction Entry Modifications**（P3 Primary） | **0** |
| **Frozen Layer Mods** | **0** always |
| **Schema Extensions** | **0** until exit review |

## Plan evidence

- P1 closure：[`rc2-p1-interaction-representability-closure.md`](../../../plans/active/2026-07-14-0856-ui-pattern-knowledge-workflow/evidence/rc2-p1-interaction-representability-closure.md)
- P2 closure：[`rc2-p2-interaction-inferability-closure.md`](../../../plans/active/2026-07-14-0856-ui-pattern-knowledge-workflow/evidence/rc2-p2-interaction-inferability-closure.md)
- P3 kickoff：[`rc2-p3-interaction-composition-start.md`](../../../plans/active/2026-07-14-0856-ui-pattern-knowledge-workflow/evidence/rc2-p3-interaction-composition-start.md)

**禁止**：Runtime projection；完整 Player / Payment business state machines；P3 kickoff 預填 constraints.
