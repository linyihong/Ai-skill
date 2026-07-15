# UI Interaction Knowledge — reusable rule seeds（Ai-skill）

**Research Cycle 2** · RC2-P1/P2/P3 ✅ **Closed** · Interaction Knowledge 🟢 **Stable**

Interaction Knowledge describes the valid temporal lifecycle of UI state **after Pattern selection and Composition have been validated** — not another Composition layer, and **not application business workflow**.

> **Canonical boundary**: Interaction Knowledge describes **UI interaction semantics**, not application business workflow.  
> Example: `dialog_open` → `user_confirm` → `dialog_close` — **not** `order_paid` / payment-order pending.

| Layer | Owns |
| --- | --- |
| Pattern Knowledge（frozen） | What pattern is appropriate |
| Composition Knowledge（frozen） | Spatial / constraint edges between patterns |
| **Interaction Knowledge** | State ownership, transition trigger, invalidation, recovery |
| **Interaction Composition** | Constraints between interaction lifecycles on a Screen / flow |

## Maturity（RC2）

| RC2 | 狀態 |
| --- | --- |
| P1 Representability | ✅ Closed |
| P2 Inferability | ✅ Closed |
| P3 Composition | ✅ Closed |
| **Interaction Knowledge** | 🟢 **Stable** |

## Layout

```text
ui-interaction-knowledge/
  README.md
  interaction_composition_rules.yaml
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

**四欄 vocabulary 已於 RC2-P1 驗證** — post-P3 exit review 前不擴 schema。

## Plan evidence

- P1：[`rc2-p1-interaction-representability-closure.md`](../../../plans/active/2026-07-14-0856-ui-pattern-knowledge-workflow/evidence/rc2-p1-interaction-representability-closure.md)
- P2：[`rc2-p2-interaction-inferability-closure.md`](../../../plans/active/2026-07-14-0856-ui-pattern-knowledge-workflow/evidence/rc2-p2-interaction-inferability-closure.md)
- P3：[`rc2-p3-interaction-composition-closure.md`](../../../plans/active/2026-07-14-0856-ui-pattern-knowledge-workflow/evidence/rc2-p3-interaction-composition-closure.md)

## Post-research maintenance

**Stable Maintenance Dogfood** — [`../maintenance-governance.md`](../maintenance-governance.md)

New incidents: Layer First → Existing Entry/Constraint → PASS → Archive. **No** default vocabulary expansion.
