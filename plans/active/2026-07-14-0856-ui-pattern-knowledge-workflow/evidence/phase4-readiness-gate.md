# Readiness Gate — Interaction Knowledge（原 Phase 4 Readiness）

**Plan**: [`../_plan.md`](../_plan.md)  
**Status**: ✅ **Closed**（R1∧R2∧R3 PASS — stakeholder 2026-07-15）  
**Research Cycle 2**: ▶ **Started** · Interaction Knowledge = 🟡 **Research Justified**  
**Date opened**: 2026-07-14 · **Date closed**: 2026-07-15

> **檔名保留** `phase4-readiness-gate.md` 以免連結斷裂；語意上本檔為 **Readiness Gate**，不是「Phase 4」。  
> **不再使用「Phase 4」** 指 Interaction — 見 [`../_plan.md`](../_plan.md) §Research Cycle 2。

---

## Research rhythm（結案版）

```text
Research Cycle 1
        │
        ▼
Knowledge Evolution Method (Emerging)
        │
        ▼
Readiness Gate (R1 → R2 → R3)
        │
        ├── R1 FAIL → 留在 Phase 3 / 繼續搜證
        │
        └── R1∧R2∧R3 PASS
               │
               ▼
        Research Cycle 2 — Interaction Knowledge
        （新 Cycle；非 Cycle 1 的 Phase 4）
```

Method：**不是下一層看起來合理就建下一層；是上一層第一次解釋不了真實現象，才建下一層。**

---

## Research Cycle 2 開啟條件（已滿足）

> 發現一個**經過驗證（validated）**、且**無法由既有 Knowledge Layer 表達（cannot be represented）**的 Flow 問題。

**Evidence**: [`r1-consumer-dogfood-2026-07-15.md`](r1-consumer-dogfood-2026-07-15.md) — C1（preview gate projection break）+ C2（payment leave confirm counterfactual）；兩個獨立 consumer。

---

## Gates = 必要且有序（結案）

```text
R1  ✅ PASS
      │
      ▼
R2  ✅ PASS
      │
      ▼
R3  ✅ PASS
      │
      ▼
Research Cycle 2 Start
```

### R1 — Pattern ✅ + Composition ✅，Flow ❌ — ✅ PASS

| Case | Pattern | Composition | Flow |
| --- | --- | --- | --- |
| **C1** preview gate projection break | modal_dialog ✅ | episode_detail + scrim + constraints ✅ | listener owner / projection break ❌ |
| **C2** payment leave confirm (C.5) | modal_dialog ✅ | dialog + scrim + accordion ✅ | pending invalidation ❌ |

**不是**單一 incident：**兩個獨立 consumer** 滿足同一 gate 形狀。

### R2 — 不能靠 composition_rules +1 修復 — ✅ PASS

已試且**不足以**修復的 constraint 類型（spatial）：

- `overlay.dialog_requires_scrim`（已有）
- `overlay.no_concurrent_temporary_overlays`（已有）
- `player.cannot_coexist_with: adjacent_preload`（假設 — 與 vertical snap 衝突）

**證偽的是 Constraint 類型**，不是「覺得 rule 不夠」。需要的是：

- State Owner
- Transition Trigger
- Recovery Boundary

→ **非 Spatial Constraint** → R2 PASS。

### R3 — Interaction Knowledge 一句定義 — ✅ PASS

**English（canonical）**

> Interaction Knowledge describes the valid temporal lifecycle of UI state **after Pattern selection and Composition have been validated**, including state ownership, transition triggers, invalidation events, and recovery boundaries.

**中文**

> Interaction Knowledge 描述在 Pattern 選型與 Composition 驗證完成後，UI 狀態於時間軸上的合法生命週期，包括狀態擁有者、轉移觸發、失效事件與恢復邊界。

**鎖定句**：`after Pattern selection and Composition have been validated` — Interaction 是**下一層**，不是另一種 Composition。

---

## Stakeholder judgment（2026-07-15）

| 項目 | 判定 |
| --- | --- |
| R1 | ✅ PASS（2 independent consumer cases） |
| R2 | ✅ PASS（Constraint 類型已證偽） |
| R3 | ✅ PASS（Interaction 邊界可一句話定義） |
| Interaction Knowledge | 🟡 **Research Justified**（≠ Stable；≠ Observation） |
| Research Cycle 2 | ✅ **可以啟動** |
| ~~Phase 4 Start~~ | **不簽** — 改用 Research Cycle 2 命名 |

---

## Domain vs Method 成果

| 層 | 成果 |
| --- | --- |
| **Domain（UI）** | Pattern Knowledge · Composition Knowledge |
| **Domain（下一層）** | Interaction Knowledge — 🟡 Research Justified |
| **Method（Architecture）** | Layer Growth Rhythm · Constraint Accumulation · Governed Trace Termination · **Readiness-before-New-Layer** |

Knowledge Evolution Method 仍 🟡 **Replicated once**（RC2-P1 independent replication）— RC2-P2 為 Method Inferability replication。Canonical log：[`Architecture Evolution Protocol` §Method Validation Log](../../../../governance/lifecycle/architecture-evolution-protocol.md#method-validation-log)

---

## Final maturity（Readiness 結案）

| 對象 | 狀態 | 下一步 |
| --- | --- | --- |
| Pattern Knowledge | 🟢 Stable | 維護與擴充 Pattern |
| Composition Knowledge | 🟢 Stable | 新 Screen 持續 dogfood |
| Knowledge Evolution Method | 🟡 Replicated once | RC2-P1 independent replication；見 Protocol §Method Validation Log |
| Readiness Gate | ✅ Closed | — |
| Interaction Knowledge | 🟡 **Research Justified** | RC2-P1/P2 ✅ · RC2-P3 ▶ Active |
| Research Cycle 2 | ▶ Started | 見 [`../_plan.md`](../_plan.md) §Research Cycle 2 |

🟡 **Research Justified** = 值得開始研究；**不是**已知道怎麼做（≠ Stable）。

---

## Log（主動搜證）

| Date | Candidate | R1 | R2 | R3 | Disposition |
| --- | --- | --- | --- | --- | --- |
| 2026-07-15 | C1 player preview gate projection break | PASS | PASS | PASS | RC2 justified；見 [`r1-consumer-dogfood-2026-07-15.md`](r1-consumer-dogfood-2026-07-15.md) |
| 2026-07-15 | C2 payment leave confirm (C.5) | PASS | PASS | PASS | 同上（獨立 consumer #2） |

---

## Explicit non-design（Readiness 期間 — 仍有效）

- [x] Readiness 設計完整 → Gate closed  
- [ ] ~~Interaction Knowledge schema~~ — **RC2 Phase 1 可啟動**（仍非 Stable claim）  
- [ ] ~~Runtime Projection for interaction~~ — 仍不做，直到 RC2 promote 條件
