# RC2-P3 CH2 — Completeness via Deferred Disposition

**Plan**: [`../_plan.md`](../_plan.md)  
**Kickoff**: [`rc2-p3-interaction-composition-start.md`](rc2-p3-interaction-composition-start.md)  
**Prerequisite**: [`rc2-p3-ch1-independence-stress.md`](rc2-p3-ch1-independence-stress.md) ✅  
**Date**: 2026-07-15  
**Mode**: rule-trace + consumer screen inventory（停 entry 編輯）

## Success criterion（鎖定）

**不是** Deferred → 0。  
**是**：Every deferred **Interaction composition** node has an **explicit disposition**.

Allowed dispositions（Interaction-adapted from RC1 H5）：

| disposition | Meaning |
| --- | --- |
| `uncovered_interaction_entry` | 候選新 Interaction entry；本波次 **不** landing |
| `composition_only` | 既有 entry(ies) + `interaction_composition_rules` 邊界；**不**開新 entry |
| `single_entry_sufficient` | 單一 entry 已足夠；無 composition edge 缺口 |
| `implementation_only` | 實作／DOM 細節；不進 Interaction Knowledge |
| `out_of_scope` | 錯層或本波次故意不收 |

Research question：許多 Interaction 失敗 **不是** New Entry，而是 **Existing Entry + Composition Constraint**（CH1 已證）。

---

## Inventory（`player_immersive_episode` + defer C3）

| Node | Before CH2 | Disposition | Verdict |
| --- | --- | --- | --- |
| `confirm_intent` under active preview gate | deferred / I-04 partial entry | **`composition_only`** | CH1 Case A constraint — **非**新 `coupon_confirm_transition` entry |
| `preview_limit` owner vs nested confirm inflight | deferred / hazard | **`composition_only`** | CH1 Case B constraint |
| `coupon_redeem_pending` terminal chain | uncovered lifecycle | **`implementation_only`** | Domain API + screen mapping；composition 邊界已由 CH1 覆蓋 |
| Full player state machine graph | gap anxiety | **`out_of_scope`** | P3 unit = screen composition edges，非 state machine |
| Episode sheet + preview hit-test (C4 / I-06) | mis-route | **`out_of_scope`** | Composition layer — decoy |
| `membership_payment_flow` accordion unmount (C3) | defer | **`single_entry_sufficient`** | `payment_leave_transition` 單 entry 足夠；無第二 composition rule 缺口 |
| `guard_condition` vocabulary (P1 defer) | defer | **`out_of_scope`** | Vocabulary freeze — P3 exit review，非 composition |
| Home tab / continuation on player-return | adjacent incident | **`out_of_scope`** | Continuation layer（I-07）— 非 Interaction composition |

**Zero Unknown.** Deferred nodes dispositioned: **8**（不追求數量下降）。

---

## Probe — `confirm_intent` under preview（CH2 價值點）

問：這是新 Interaction entry，還是 Composition？

```text
confirm_intent under preview_gate
    ?= new entry (coupon_confirm_transition)
    |  → NO (would inflate entry count under Completeness pressure)
    ?= preview_gate_transition + composition constraints
    |  → YES (composition_only)
```

| Check | Result |
| --- | --- |
| New Interaction entry created? | **No** |
| `preview_gate_transition.yaml` modified? | **No**（Invariant） |
| Disposition explicit? | **Yes** — `composition_only` |
| CH1 constraints reference node? | **Yes** — 2 rules |
| Completeness（no Unknown）? | **Yes** |

---

## C3 probe — `membership_payment_flow`

| Question | Result |
| --- | --- |
| 需要第二 composition rule？ | **No** — accordion unmount 已在 `payment_leave_transition` invalidation_event |
| 需要改 entry？ | **No** |
| Disposition | **`single_entry_sufficient`** |

本波次 **不** 為 C3 新增 `interaction_composition_rules` — 正確 defer，非缺口。

---

## Metrics（CH1 → CH2）

| Metric | Role | CH1 | CH2 |
| --- | --- | --- | --- |
| Interaction Entry Modifications | Primary | **0** | **0** |
| Interaction Composition Rule Count | Supporting | 2 | **2**（無 CH2 增量 — 正確） |
| Deferred nodes with disposition | Supporting | 2 unknown | **8 / 8** |
| Frozen Layer Mods | Invariant | 0 | **0** |

---

## Hypothesis result

| Hypothesis | Result |
| --- | --- |
| CH2 Completeness | ✅ **PASS** — disposition coverage；非 Deferred 消滅 |

---

## Structural meaning（對稱 RC1 H5）

```text
Deferred Interaction composition node
        │
        ├─ composition_only      → CH1 constraints（不開 entry）
        ├─ single_entry_sufficient → payment_leave alone
        └─ out_of_scope / implementation_only → 錯層或實作
```

**Completeness = 每個 defer 有名字**，不是把 defer 清成零。

---

## Explicit non-actions

- [x] 未新增 Interaction entry
- [x] 未新增 CH2 composition rule（C3 正確判定 single_entry_sufficient）
- [ ] CH3 Traceability — next mini-cycle
