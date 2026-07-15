# RC2-P3 CH1 — Independence Stress Run（player_immersive_episode）

**Plan**: [`../_plan.md`](../_plan.md)  
**Kickoff**: [`rc2-p3-interaction-composition-start.md`](rc2-p3-interaction-composition-start.md)  
**Date**: 2026-07-15  
**Mode**: rule-trace + consumer integration anchors（停 entry 編輯）  
**Success criterion**: ≥1 real case that **requires new `interaction_composition_rules.yaml` constraint** and **zero `entries/*.yaml` edits**.

**Invariant pointer**: Interaction Composition evidence MUST NOT modify validated Interaction entries.

---

## Setup

| Item | Value |
| --- | --- |
| Screen / flow | `player_immersive_episode` |
| Frozen Interaction entries | `preview_gate_transition` · `payment_leave_transition` |
| Pattern ref（frozen） | `modal_dialog` · `episode_detail` |
| Pre-stress Rule Count | **0** |
| Pre-stress Entry Modifications | **0** |
| Consumer anchor | `<PROJECT_ROOT>` `episode-coupon-redeem-journey.md` · I-04 |

---

## Case A — Preview gate + confirm_intent pointer channel（C1 · **首輪 FAIL 點**）

### Stress

1. `preview_gate_transition` active — `preview_limit_reached` · mask visible on main `playerStage`  
2. User expands coupon shell on preview overlay and triggers **解锁本集**（confirm_intent on interaction channel）  
3. **無** interaction composition rule 規範 preview lifecycle 與 nested confirm channel 的並存邊界

### Observe

| Question | Result |
| --- | --- |
| `preview_gate_transition` entry 是否已涵蓋 confirm channel？ | **否** — entry scope 為 preview→gated owner binding；I-04 僅 partial |
| 改 `preview_gate_transition.yaml` 加 confirm 欄位能修嗎？ | **能但違反 Invariant** — 會把單 transition entry 膨脹成 screen graph |
| 既有 `interaction_composition_rules` 能擋 pointer-only 失效？ | **否** — scaffold empty（首輪 **FAIL**） |
| Integration 現況 | L2 semantic path + L1 capability poll **PASS**（`episode_coupon_redeem_journey.integration.mjs`）— 證明需 semantic path constraint，非 pointer-only |

### Failure (valuable)

Phase C finding：`event injected ≈ interaction success` 在 preview overlay 下 falsified — **兩條 Interaction lifecycle 的 composition edge 缺失**。  
若回頭改 `preview_gate_transition` → 違反 RC2-P3 Invariant。  
正確處置 → **新增 Interaction Composition Constraint**。

### Repair（Constraint only）

Added:

```text
interaction.preview_gate_preserves_confirm_semantic_path
  when_active: preview_gate_transition
  while_nested_lifecycle: confirm_intent
  requires: semantic_confirm_path (role / keyboard / focus)
  forbids_trust: pointer_only_on_overlay_actions
```

### Re-check

| Check | Result |
| --- | --- |
| New rule exists | ✅ |
| Interaction entry diff | ✅ empty |
| Preview entry scope 未改寫（Independence） | ✅ |

**Case A verdict**: stress-FAIL → Constraint-add → PASS

---

## Case B — Owner invalidation before nested confirm completes（C2）

### Stress

1. `preview_gate_transition` owns `previewLimitReached` on `ImmersivePlayerFrame`  
2. Redeem flow opens confirm (`alertPopup`) while preview gate still active  
3. Early success path can clear mask / owner state **before** confirm reaches terminal state

### Observe

| Question | Result |
| --- | --- |
| Hazard class `owner-invalidation-before-complete` 落在哪層？ | **Interaction composition** — 兩 lifecycle 的 invalidation 順序 |
| `preview_gate_transition.invalidation_event` 擴寫能修嗎？ | **違反 Invariant** — 會把 nested confirm 塞進單 entry |
| Integration `confirm_stable_during_inflight` | **PASS** — 現實修復已存在；知識缺口在 composition rules 未命名 |

### Repair（Constraint only）

Added:

```text
interaction.preview_gate_owner_stable_during_nested_confirm
  when_active: preview_gate_transition
  while_nested: confirm_root_stable_during_inflight
  forbids_invalidation: preview_limit_cleared_before_confirm_terminal
```

### Re-check

| Check | Result |
| --- | --- |
| New rule exists | ✅ |
| Interaction entry diff | ✅ empty |
| Screen mapping invariant 對齊 | ✅ `confirm_root_stable_during_inflight` |

**Case B verdict**: stress-FAIL → Constraint-add → PASS

---

## Case C — Decoy（C4 · wrong layer）

Episode sheet TabBar hit steal（I-06）— **Composition** overlay hit-test，不是 Interaction composition。  
**不**進 `interaction_composition_rules.yaml`。

---

## CH1 Metrics（post-stress）

| Metric | Pre | Post |
| --- | --- | --- |
| Interaction Composition Rule Count | 0 | **2** |
| Interaction Entry Modifications | 0 | **0** |
| Frozen Layer Mods | 0 | **0** |

---

## Hypothesis result

| Hypothesis | Result |
| --- | --- |
| CH1 Independence | ✅ **PASS** — 2 cases · constraint-only repair · Entry Mods = 0 |

---

## Structural meaning（對稱 RC1 H4）

```text
Stress（兩條 Interaction lifecycle 並存）
        │
        ▼
Interaction Composition Constraint
        │
        ├── Interaction entries = 0 modifications
        └── interaction_composition_rules +2
```

第一次證明：**Interaction Composition 可自我吸收新知識，而不污染 Interaction entries 或 frozen Pattern/Composition。**

---

## Explicit non-actions

- [x] 未修改 `preview_gate_transition.yaml` / `payment_leave_transition.yaml`
- [x] 未修改 `ui-pattern-knowledge/composition_rules.yaml`
- [ ] CH2 Completeness — deferred disposition（下一 mini-cycle）
- [ ] CH3 Traceability
