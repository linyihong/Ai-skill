# UI Interaction Knowledge — reusable rule seeds（Ai-skill）

**Research Cycle 2** · RC2-P1 ✅ **Closed** · RC2-P2 ▶ **Active**

Interaction Knowledge describes the valid temporal lifecycle of UI state **after Pattern selection and Composition have been validated** — not another Composition layer, and **not application business workflow**.

> **Canonical boundary**: Interaction Knowledge describes **UI interaction semantics**, not application business workflow.  
> Example: `dialog_open` → `user_confirm` → `dialog_close` — **not** `order_paid` / payment-order pending.

| Layer | Owns |
| --- | --- |
| Pattern Knowledge（frozen） | What pattern is appropriate |
| Composition Knowledge（frozen） | Spatial / constraint edges between patterns |
| **Interaction Knowledge** | State ownership, transition trigger, invalidation, recovery |

## Maturity（RC2）

| RC2 | 狀態 |
| --- | --- |
| P1 Representability | ✅ Closed |
| P2 Inferability | ▶ Active |
| P3 Composition | ⏸ Locked |
| Interaction Knowledge | 🟡 Research Justified |

## Layout

```text
ui-interaction-knowledge/
  README.md
  validation/
    interaction-entry-schema.yaml   # Vocabulary freeze: four core fields only
  entries/
    preview_gate_transition.yaml    # P1 entry
    payment_leave_transition.yaml   # P2 second entry（dogfood 前落地）
```

## RC2 Invariant

```text
Interaction evidence MUST NOT redefine Pattern selection or Composition constraints.
Interaction MUST NOT edit ui-pattern-knowledge/entries/* or composition_rules.yaml.
```

## Vocabulary Freeze

**四欄 vocabulary 已於 RC2-P1 驗證**：`state_owner`, `transition_trigger`, `invalidation_event`, `recovery_boundary`

- **Forbidden** mid-P2 schema adds without exit review
- Dogfood gaps → evidence note → P2 exit review

## RC2 Metrics（non-KPI）

| Metric | Target |
| --- | --- |
| **Schema Extensions** | **0** until exit review |
| **Frozen Layer Mods** | **0** always |
| **Boundary Misclassification**（P2 primary） | **0** |

## RC2-P2 Inferability（layer-first）

> Inferability must classify the **correct knowledge layer** before identifying the **correct interaction entry**.

**Direction**：Incident / Scenario → Interaction Entry（不是 Entry → Scenario）

**P2 case**：`payment_leave_transition`（非 `preview_gate_transition` — 泛化證據）

## Hazard Review boundary

> Interaction Knowledge describes how interaction **should** behave; Interaction Hazard Review evaluates how interaction **can fail**.

Shared vocabulary · **not** shared responsibility.

## Plan evidence

- P1 closure：[`plans/active/2026-07-14-0856-ui-pattern-knowledge-workflow/evidence/rc2-p1-interaction-representability-closure.md`](../../../plans/active/2026-07-14-0856-ui-pattern-knowledge-workflow/evidence/rc2-p1-interaction-representability-closure.md)
- P2 kickoff：[`plans/active/2026-07-14-0856-ui-pattern-knowledge-workflow/evidence/rc2-p2-interaction-inferability-start.md`](../../../plans/active/2026-07-14-0856-ui-pattern-knowledge-workflow/evidence/rc2-p2-interaction-inferability-start.md)

**禁止**：Runtime projection；完整 Player / Payment business state machines in Interaction entries；P2 用 preview 作唯一案例.
