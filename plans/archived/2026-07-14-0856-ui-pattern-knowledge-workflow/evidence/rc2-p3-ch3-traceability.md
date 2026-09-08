# RC2-P3 CH3 — Traceability（可追溯 + 可停止）

**Plan**: [`../_plan.md`](../_plan.md)  
**Kickoff**: [`rc2-p3-interaction-composition-start.md`](rc2-p3-interaction-composition-start.md)  
**Prerequisites**: CH1 ✅ · CH2 ✅  
**Date**: 2026-07-15  
**Mode**: Interaction Composition Trace Graph

## Traceability redefined

不是「每個 Screen 節點都有獨立 Interaction entry」。  
而是：**每個節點都必須有合法且可解釋的終點** — `complete` 或 `waived`。

```text
entry: null（無 disposition）     → 非法
trace_status: waived + reason     → 合法終點
trace_status: complete            → entry 和／或 composition constraint 鏈完整
```

---

## Exit Gate

| ID | Rule | Result |
| --- | --- | --- |
| **CT1** Continuous Trace | 非 waived：Screen → Interaction entry →（可選）composition constraint | ✅ |
| **CT2** Governed Stop | CH2 disposition 節點必有 `waive_reason` + `owner` | ✅ |
| **CT3** No Broken Edge | 終點只能是 **complete** 或 **waived**；無 `?????` | ✅ |

---

## Path A — 完整 Trace（`player_immersive_episode` primary）

### A1 — Preview gate entry

```text
player_immersive_episode
    │
    ▼
preview_gate_transition
  (entries/preview_gate_transition.yaml — four fields)
    │
    ▼
recovery_boundary
  consumer: episode-coupon-redeem-journey · player-preview-gate integration
```

`trace_status`: **complete**

### A2 — Composition edges（CH1）

```text
player_immersive_episode
    │
    ├─ preview_gate_transition (active)
    │     └─ interaction.preview_gate_preserves_confirm_semantic_path
    │     └─ interaction.preview_gate_owner_stable_during_nested_confirm
    │
    └─ confirm_intent channel (nested)
          → composition_only terminal via constraints（非新 entry）
```

`trace_status`: **complete**（composition constraint 鏈）

### A3 — Membership payment（C3）

```text
membership_payment_flow
    │
    ▼
payment_leave_transition
  (entries/payment_leave_transition.yaml)
    │
    ▼
recovery_boundary
  consumer: payment-leave-confirm integration
```

`trace_status`: **complete**（`single_entry_sufficient` — 無第二 composition rule 缺口）

---

## Path B — Waived Trace（CH2 disposition 終點）

| Node | Disposition | `waive_reason` | `owner` | `trace_status` |
| --- | --- | --- | --- | --- |
| `coupon_redeem_pending` API chain | implementation_only | domain_api_and_screen_mapping_not_interaction_composition | implementation | waived |
| Full player state machine | out_of_scope | p3_unit_is_composition_edges_not_state_machine | rc2-p3 | waived |
| Episode sheet hit-test (I-06) | out_of_scope | composition_layer_decoy | composition | waived |
| `guard_condition` vocabulary | out_of_scope | vocabulary_freeze_until_p3_exit_review | rc2 | waived |
| Home tab / player-return (I-07) | out_of_scope | continuation_layer_not_interaction_composition | continuation | waived |

CH2 已 disposition → CH3 補 **waive_reason + owner** → 非黑洞。

---

## Full Screen Trace Graph（終點表）

| Node | Path | `trace_status` | Terminal |
| --- | --- | --- | --- |
| `preview_gate_transition` on episode screen | A1 | complete | Entry + recovery evidence |
| `confirm_intent` under preview | A2 | complete | Composition constraints（2） |
| `payment_leave_transition` on membership flow | A3 | complete | Single entry sufficient |
| `coupon_redeem_pending` | B | waived | Governed waiver |
| Full player state machine | B | waived | Governed waiver |
| Episode sheet hit-test | B | waived | Governed waiver |
| `guard_condition` defer | B | waived | Governed waiver |
| Player-return continuation | B | waived | Governed waiver |

**Broken edges**: **0**  
**Interaction Entry Modifications**: **0**  
**No Unknown / ?????**: ✅

---

## Relation to CH2

CH2 證明 Composition 能 **拒絕新增 Interaction entry**（`confirm_intent` → `composition_only`）。  
CH3 證明該拒絕不是黑洞：**waived** 與 **complete** 皆為一等公民終點。

---

## Hypothesis result

| Hypothesis | Result |
| --- | --- |
| CH3 Traceability | ✅ **PASS** — CT1∧CT2∧CT3 |
| CH1∧CH2∧CH3 | ✅ **PASS** |
| Interaction Entry Modifications | **0** |

**Interaction Composition Closure**：可宣告（三 mini-cycles + Entry Mods=0）。

---

## Explicit non-actions

- [x] 未新增 Interaction entry
- [x] 未修改 frozen Pattern / Composition paths
- [x] RC2-P3 Closure — [`rc2-p3-interaction-composition-closure.md`](rc2-p3-interaction-composition-closure.md)
