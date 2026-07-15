# UI Interaction Knowledge — reusable rule seeds（Ai-skill）

**Research Cycle 2** · RC2-P1 **Interaction Representability** in progress.

Interaction Knowledge describes the valid temporal lifecycle of UI state **after Pattern selection and Composition have been validated** — not another Composition layer, and **not application business workflow**.

> **Canonical boundary**: Interaction Knowledge describes **UI interaction semantics**, not application business workflow.  
> Example: `dialog_open` → `user_confirm` → `dialog_close` — **not** `order_paid` / payment-order pending.

| Layer | Owns |
| --- | --- |
| Pattern Knowledge（frozen） | What pattern is appropriate |
| Composition Knowledge（frozen） | Spatial / constraint edges between patterns |
| **Interaction Knowledge** | State ownership, transition trigger, invalidation, recovery |

## Layout（RC2-P1 — do not bloat）

```text
ui-interaction-knowledge/
  README.md
  validation/
    interaction-entry-schema.yaml   # Vocabulary freeze: four core fields only
  entries/
    preview_gate_transition.yaml    # First entry — RC2-P1 dogfood only
```

**Sequence**（locked）：

```text
interaction-entry-schema
        │
        ▼
preview_gate_transition.yaml
        │
        ▼
Dogfood（representability evidence）
        │
        ▼
Frozen Layer Mods = 0
```

## RC2 Invariant

```text
Interaction evidence MUST NOT redefine Pattern selection or Composition constraints.
Interaction MUST NOT edit ui-pattern-knowledge/entries/* or composition_rules.yaml.
```

## Vocabulary Freeze（RC2-P1）

Until first entry **representability** exit review:

- **Allowed** interaction primitives: `state_owner`, `transition_trigger`, `invalidation_event`, `recovery_boundary`
- **Forbidden** mid-dogfood schema adds（e.g. `guard_condition`, `rollback`, `checkpoint`, `priority`）
- If dogfood surfaces a gap → **Evidence note only** → dogfood end → Review → decide schema extension

Symmetric discipline to RC1 **Entry Freeze**.

## RC2 Metrics（non-KPI）

| Metric | RC2-P1 target |
| --- | --- |
| **Schema Extensions** | **0** until exit review |
| **Interaction Entry Mods** | **0** after first entry lands（one entry only） |
| **Frozen Layer Mods** | **0** always |

## Hazard Review boundary

> Interaction Knowledge describes how interaction **should** behave; Interaction Hazard Review evaluates how interaction **can fail**.

Shared vocabulary · **not** shared responsibility.

## Plan evidence

- Kickoff：[`plans/active/2026-07-14-0856-ui-pattern-knowledge-workflow/evidence/rc2-p1-interaction-representability-start.md`](../../../plans/active/2026-07-14-0856-ui-pattern-knowledge-workflow/evidence/rc2-p1-interaction-representability-start.md)
- Dogfood：[`plans/active/2026-07-14-0856-ui-pattern-knowledge-workflow/evidence/rc2-p1-preview-gate-representability-run.md`](../../../plans/active/2026-07-14-0856-ui-pattern-knowledge-workflow/evidence/rc2-p1-preview-gate-representability-run.md)

**禁止**：Runtime projection；多 entry 堆疊；完整 Player / Payment business state machines in Interaction entries.
