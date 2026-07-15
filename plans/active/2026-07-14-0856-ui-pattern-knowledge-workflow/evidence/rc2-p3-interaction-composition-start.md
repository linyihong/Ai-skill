# RC2-P3 Start Lock — Interaction Composition

**Plan**: [`../_plan.md`](../_plan.md)  
**Date**: 2026-07-15  
**Status**: RC2-P3 **Active**（stakeholder kickoff — 對齊 RC1 Phase 3 / RC2-P1/P2 慣例）  
**Prerequisite**: RC2-P2 ✅ Closed — [`rc2-p2-interaction-inferability-closure.md`](rc2-p2-interaction-inferability-closure.md)

---

## Framing — Dual evidence chains

RC2-P3 同時產出兩條證據鏈（延續 P2）：

| Chain | Validates |
| --- | --- |
| **Domain Knowledge** | Interaction Composability（多 transition 在同一 Screen / flow 的約束） |
| **Method Knowledge** | Knowledge Evolution Method **Composability** independent replication |

- P3 **成功** → Representability → Inferability → **Composability** 三階梯在 Interaction 層完整  
- P3 **失敗** → 可分辨 Interaction composition gap vs Method 適用邊界  

Meta record：[`Architecture Evolution Protocol` §Method Validation Log](../../../../governance/lifecycle/architecture-evolution-protocol.md#method-validation-log)

**Consumer writeback**（mandatory）：每輪 `<AI_SKILL_DOGFOOD_EVIDENCE>` composition dogfood 結束 → 回寫本 plan `evidence/`。Rule：`<PROJECT_ROOT>/.ai-skill/project/rules/rc2-consumer-evidence-writeback.md`

---

## Methodology lock（與 RC1 Phase 3 平行）

RC1-P3 審 **Pattern 之間的 Composition Constraints**（單位 = Screen）。  
RC2-P3 審 **Interaction transition 之間的 Composition Constraints**（單位 = Screen / flow）。

| RC1 Phase 3 | RC2-P3 |
| --- | --- |
| H4 Independence | **CH1 Independence** |
| H5 Completeness | **CH2 Completeness** |
| H6 Traceability | **CH3 Traceability** |
| `composition_rules.yaml` | [`interaction_composition_rules.yaml`](../../../../workflow/software-delivery/ui-interaction-knowledge/interaction_composition_rules.yaml) |
| Entry Modifications = 0 | **Interaction Entry Modifications = 0** |

驗證單位 = **Screen / flow**；審查方式 ≠ 一次跑完整 Player 或 Payment 業務流程。

Cross-method：**Grow one layer, freeze the previous** — Pattern + Composition + Interaction **entries** frozen。

---

## RC2-P3 Invariant（升格 — 不可違反）

```text
Interaction Composition evidence MUST NOT modify validated Interaction entries.

New knowledge discovered during composition flows into:
  interaction_composition_rules.yaml   (= Interaction Composition Constraints)
NOT
  entries/*.yaml

Interaction Composition MUST NOT edit:
  ui-pattern-knowledge/entries/*
  ui-pattern-knowledge/composition_rules.yaml
```

**Risk**: **Back-propagation（回滲）**。違反 Invariant ⇒ 證據作廢。

---

## 1. Success Definition — Hypotheses

### CH1 — Independence（首要 mini-cycle）

**問題**：同一 Screen 上兩條 Interaction lifecycle **並存或交錯**時，能否用 **Constraint only** 修復，而不改 `entries/*.yaml`？

**首輪 stress screen**（consumer）：`player_immersive_episode`

| 元件 | Interaction entry | Composition ref |
| --- | --- | --- |
| Preview gate | `preview_gate_transition` | `episode_detail` |
| Coupon confirm under preview overlay | ⚠️ partial → `preview_gate_transition` 或新 constraint | `modal_dialog` |

**Stress shape**（來自 P2 intake I-04）：`preview_limit_reached` 態下 `confirm_intent` 在 interaction channel 被攔截 — **兩條 lifecycle 的 invalidation 邊界衝突**。

**成功** = ≥1 real case 需要新 `interaction_composition_rules.yaml` constraint 且 **Interaction Entry Modifications = 0**。

### CH2 — Completeness

**問題**：Pattern Tree 上每個 **deferred interaction composition node** 是否有 **disposition**（complete \| waived \| constraint_added）？

**不是** Deferred → 0 強行歸零。

### CH3 — Traceability

**問題**：Screen / flow → Interaction entries + composition constraints 能否追到 **complete \| waived** 終點？

**鏈形（候選）**：

```text
player_immersive_episode
├── preview_gate_transition
│     ├── pattern: modal_dialog
│     └── composition: episode_detail
└── [CH1 stress] coupon confirm interaction channel
      └── interaction_composition_rules constraint (TBD)
```

---

## 2. Primary screen / flow lock

| 優先 | Screen / flow ID | 理由 | Consumer anchor |
| --- | --- | --- | --- |
| ✅ **P3 primary** | `player_immersive_episode` | 兩 entry 已存在 + I-04 多 lifecycle 衝突 | I-04 · `episode-coupon-redeem-journey.md` |
| defer | `membership_payment_flow` | `payment_leave_transition` 單鏈；可作 CH2/CH3 第二輪 | I-05 · C.5 trial |

**不得**用 kickoff 預建 constraint — 只鎖 screen 與 stress 形狀。

---

## 3. RC2-P3 Metrics（non-KPI）

| Metric | RC2-P3 target |
| --- | --- |
| **Interaction Entry Modifications** | **0**（Primary） |
| Interaction Composition Rule Count | Supporting · Δ can ↑ |
| Deferred Interaction Nodes | Supporting · disposition 齊全 > 歸零 |
| Frozen Layer Mods（Pattern + Composition + Interaction entries） | **0** |
| Schema Extensions | **0** until P3 exit review |

見 kickoff 後落地 [`rc2-p3-composition-metrics.md`](rc2-p3-composition-metrics.md)（待 CH1 首輪後）。

---

## 4. Mini-cycles（分開跑）

| Order | Cycle | Evidence shape |
| --- | --- | --- |
| 1 | **CH1 Independence** | stress → constraint-only repair |
| 2 | **CH2 Completeness** | deferred disposition table |
| 3 | **CH3 Traceability** | complete \| waived 終點 |

**禁止** 一檔混跑三 cycle — 對稱 RC1 `3h4` / `3h5` / `3h6` 分檔慣例。

---

## 5. Exit Gate（RC2-P3）

| 通過 | 不通過 |
| --- | --- |
| CH1∧CH2∧CH3 mini-cycles PASS | 為修 composition 回寫 Interaction entries |
| Interaction Entry Modifications = **0** | 回寫 Pattern / Composition frozen paths |
| ≥1 validated interaction composition constraint | mid-P3 schema 膨脹無 Review |
| Frozen Layer Mods = 0 | — |

**是** Exit：**Interaction Composition Closure** = CH1∧CH2∧CH3 + Entry Mods = 0。

**也是** Exit（整層）：Interaction Knowledge → 🟢 **Stable**（RC2 三階梯完成）。

**不是** Exit：「所有 Screen 都有 interaction composition rules」。

---

## 6. Kickoff 五問

| # | 問題 | 答案 |
| --- | --- | --- |
| 1 | 驗什麼能力？ | Screen → Interaction entries + Composition Constraints |
| 2 | 案例？ | `player_immersive_episode` + I-04 stress |
| 3 | Primary metric？ | Interaction Entry Modifications = 0 |
| 4 | 順序？ | CH1 Independence → CH2 → CH3 |
| 5 | 鄰居風險？ | 不吸 Pattern selection / Composition spatial / Hazard taxonomy |
| 6 | 驗方法還是驗 Domain？ | **Both** — Composability replication |

---

## Explicit non-goals（RC2-P3）

- [ ] ~~用單一 entry 代表整個 Player state machine~~
- [ ] ~~在 kickoff 預填 interaction composition constraints~~
- [ ] ~~合併 Interaction Hazard Review artifact~~
- [ ] ~~升格前未跑 CH1–CH3~~

---

## RC2 maturity snapshot（kickoff）

| RC2 | 狀態 |
| --- | --- |
| P1 Representability | 🟢 Stable |
| P2 Inferability | ✅ Closed |
| P3 Composition | ▶ **Active** |
| Interaction Knowledge | 🟡 Research Justified（P3 exit → 🟢 Stable） |

---

## Next execution

| Step | Status | Artifact |
| --- | --- | --- |
| 1. `interaction_composition_rules.yaml` scaffold | ✅ | [`interaction_composition_rules.yaml`](../../../../workflow/software-delivery/ui-interaction-knowledge/interaction_composition_rules.yaml) |
| 2. Consumer P3 intake | ▶ | `<PROJECT_ROOT>/.ai-skill/project/evidence/rc2-p3-interaction-composition-intake.md` |
| 3. CH1 Independence stress | ✅ | [`rc2-p3-ch1-independence-stress.md`](rc2-p3-ch1-independence-stress.md) |
| 4. CH2 Completeness | ⏸ | `rc2-p3-ch2-completeness-disposition.md` |
| 5. CH3 Traceability | ⏸ | `rc2-p3-ch3-traceability.md` |
