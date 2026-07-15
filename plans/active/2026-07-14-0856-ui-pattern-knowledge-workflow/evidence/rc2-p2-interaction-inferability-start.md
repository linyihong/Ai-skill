# RC2-P2 Start Lock — Interaction Inferability

**Plan**: [`../_plan.md`](../_plan.md)  
**Date**: 2026-07-15  
**Status**: RC2-P2 **Active**（stakeholder kickoff complete — 對齊 RC1-P2 Active 慣例）  
**Prerequisite**: RC2-P1 ✅ Closed — [`rc2-p1-interaction-representability-closure.md`](rc2-p1-interaction-representability-closure.md)

---

## Framing — Dual evidence chains

RC2-P2 同時產出兩條證據鏈（stakeholder 2026-07-15）：

| Chain | Validates |
| --- | --- |
| **Domain Knowledge** | Interaction Inferability |
| **Method Knowledge** | Knowledge Evolution Method **Inferability independent replication** |

- P2 **成功** → 方法在 Representability 之外再次跨 Layer 成立  
- P2 **失敗** → 可分辨 Interaction-layer gap vs Method 適用邊界  

Meta record（**不在 RC2 evidence 目錄**）：[`Architecture Evolution Protocol` §Method Validation Log](../../../../governance/lifecycle/architecture-evolution-protocol.md#method-validation-log)

**Consumer writeback**（mandatory）：每輪 `<AI_SKILL_DOGFOOD_EVIDENCE>` intake / dogfood 結束 → 回寫本 plan `evidence/`。Consumer rule：`<PROJECT_ROOT>/.ai-skill/project/rules/rc2-consumer-evidence-writeback.md` · workflow：`docs/framework-development-workflow.yaml` §`rc2_p2_consumer_evidence_writeback`

---

## Methodology lock（與 RC1-P2 平行）

RC1-P2 驗的是：**Scenario → 推回 Pattern Entry**（不是 Entry → Scenario）。

RC2-P2 驗的是：

```text
Incident / Scenario  →  推回 Interaction Entry
```

**方向不可反。**

**Canonical ordering sentence**

> Inferability must classify the **correct knowledge layer** before identifying the **correct interaction entry**.

先回答：**哪一層？** 再回答：**哪一個 Entry？**

---

## RC2 研究風險（P2 必擋）

Interaction 容易吸收並誤判為本層的鄰居：

| 鄰居 | 常見誤吸 |
| --- | --- |
| Continuation | capture→restore scroll/URL |
| Navigation | entry-return / back-stack |
| Hazard Review | failure taxonomy only |
| Runtime | experience-runtime state YAML |

**P2 失敗模式**：Flow 問題被誤判成 Composition 或 Pattern — 見 **Boundary Misclassification**（Primary Metric）。

---

## 1. Success Definition — Hypotheses

### IH1 — Inferability

給一段 **Incident 敘述**（scenario prompt），能否**唯一**推回：

- ✅ `payment_leave_transition`
- ❌ `preview_gate_transition`（decoy — 同層不同 entry）

**不是**證明「preview 案例也會推理」— P1 已覆蓋表示法；P2 要 **第二獨立案例** 的泛化。

### IH2 — Boundary Inferability（比 IH1 更重要）

能否推斷：**這不是 Composition，也不是 Pattern**？

| Symptom shape | 應推出 | 不應推出 |
| --- | --- | --- |
| Listener 綁錯 video / poll 在 preload | **Interaction** | `modal_dialog` Pattern 選型錯 |
| Accordion unmount 使 pending 失效 | **Interaction** | `overlay.dialog_requires_scrim` Composition 錯 |
| 選錯 bottom_sheet vs modal | **Pattern** | Interaction entry |

若 IH2 推不出層級 → Interaction **Boundary 尚未成熟**。

### IH3 — Repair Localization（推薦）

Interaction 的價值：**知道應改哪一層**。

| Incident | 應推出 Repair chain |
| --- | --- |
| `preview_limit_reached` 未在主舞台發火 | Layer: **Interaction** → Entry: `preview_gate_transition` → **不要**改 `entries/modal_dialog.yaml` |
| pending panel unmount 未經 confirm | Layer: **Interaction** → Entry: `payment_leave_transition` → **不要**改 `composition_rules.yaml` |

---

## Evidence Intake（stakeholder 2026-07-15 — RC1 紀律）

收集目標：**可證偽 RC2-P2 假說**的 Interaction Incident — 不是泛泛 Flow。

```text
Incident → Layer First? → Interaction Entry? → Repair Layer? → Evidence
```

**禁止**：`Incident → 寫新 Entry`

**Entry 節奏**（與先前 kickoff 修正）：

1. **現在**：僅 `preview_gate_transition` 已 landing
2. 用 ~10 個 Incident 嘗試映射到**既有** entry
3. 當 **Layer 判對** 且 **無法合理映射** `preview_gate_transition` → 才建 `payment_leave_transition.yaml`

Consumer intake（Evidence Producer）：  
`<PROJECT_ROOT>/.ai-skill/project/evidence/rc2-p2-interaction-incident-intake.md`

---

## 2. Entry landing 節奏（修正 — preview first）

| 階段 | Entry | 狀態 |
| --- | --- | --- |
| P1 | `preview_gate_transition` | ✅ landed |
| P2 intake | 用既有 entry 分類 Incident | ▶ Active |
| P2 trigger | `payment_leave_transition` | ✅ landed（I-05 blind + rule-trace） |

**Scope lock**（第二 entry，landing 時）：

```text
dialog_open → user_confirm_stay_or_leave → dialog_close
```

**不是** `order_paid` / `payment_pending` business states。

**Consumer anchor**：Readiness C2 · [`r1-consumer-dogfood-2026-07-15.md`](r1-consumer-dogfood-2026-07-15.md) §C2 · intake **I-05**

**P2 dogfood 前置產物**：

1. Consumer incident intake ≥10（layer-first 分類完成） — ✅
2. [`interaction-inferability-scenarios.yaml`](interaction-inferability-scenarios.yaml) — ✅ 8 scenarios（I-01–I-05 + decoys）
4. [`rc2-p2-inferability-run.md`](rc2-p2-inferability-run.md) — rule-trace 8/8 + blind round 2（6/8 layer）
5. `payment_leave_transition.yaml` — ✅ landed（I-05）

---

## 3. RC2-P2 Metrics（non-KPI）

| Metric | Role | P2 目標 |
| --- | --- | --- |
| **Layer Classification Accuracy** | **Primary** | 第一層判對 |
| **Boundary Misclassification** | **Primary** | 誤判到哪一層（非 accuracy KPI） |
| **Existing Entry Reuse Rate** | Supporting | Interaction 層能否重用既有 entry |
| **New Entry Required** | Supporting | 是否真的需要第三/第四 entry |
| Scenario Accuracy | Supporting | IH1：正確 entry id |
| Frozen Layer Mods | Blocking | **0** always |

**不是 KPI**：scenario 數量、entry 數量、覆蓋率。

### P2 Report — Table 1（primary；before accuracy）

**第一張表不是 Accuracy**，而是 layer boundary 分類：

| Incident | Initial Guess | Final Layer |
| --- | --- | --- |
| Payment Leave | Composition | Interaction |
| Projection Break | Runtime | Interaction |
| Listener bound to preload video | Pattern (`modal_dialog`) | Interaction |

RC1 最強證據不是 10/10，而是「這不是 Bottom Sheet」。RC2 對稱：「這不是 Composition / Runtime」。

### Knowledge Layer Confusion Matrix（emerging）

若 Hazard / Runtime / Continuation 持續被誤判，記錄 **Actual × Predicted** 矩陣；矩陣越乾淨 → Knowledge Boundary 越成熟 — 研究價值高於 scenario accuracy 單一數字。

| Actual | Predicted | Notes |
| --- | --- | --- |
| Interaction | Composition | Primary failure mode |
| Interaction | Runtime | P2 watch |
| Interaction | Continuation | P2 watch |
| Composition | Interaction | Inverse mislabel |

---

## 4. Invariants（延續 P1）

```text
Interaction MUST NOT redefine Pattern selection or Composition constraints.
Vocabulary: 仍僅四欄，除非 P2 exit review 批准擴充。
```

---

## 5. Exit Gate（RC2-P2）

| 通過 | 不通過 |
| --- | --- |
| IH1：C2-derived scenarios 推回 `payment_leave_transition` | 只能靠 preview 案例 |
| IH2：layer 分類正確（Interaction vs Pattern vs Composition） | Boundary Misclassification > 0 |
| IH3：repair 定位在 Interaction 層 | 建議改 `entries/*.yaml` 或 `composition_rules` |
| Frozen Layer Mods = 0 | 任何 frozen path 編輯 |
| Schema / Vocabulary Extensions = 0 | mid-P2 schema 膨脹 |

**不是** Exit：Interaction Knowledge 🟢 Stable（仍待 RC2-P3）。

---

## 6. Kickoff 五問

| # | 問題 | 答案 |
| --- | --- | --- |
| 1 | 驗什麼能力？ | Scenario → Interaction Entry + Layer + Repair |
| 2 | 案例？ | `payment_leave_transition`（非 preview） |
| 3 | Primary metric？ | Boundary Misclassification = 0 |
| 4 | 順序？ | Layer first，Entry second |
| 5 | 鄰居風險？ | 不吸 Continuation / Navigation / Hazard / Runtime |
| 6 | 驗方法還是驗 Domain？ | **Both** — dual evidence chains；Method Validation Log 在 Protocol |

---

## Explicit non-goals（RC2-P2）

- [ ] ~~用 preview_gate 作 P2 唯一案例~~
- [ ] ~~Interaction Composition（RC2-P3）~~ — 🔒 Locked
- [ ] ~~升格 Interaction Knowledge 🟢 Stable~~
- [ ] ~~擴 vocabulary 無 Review~~

---

## RC2 maturity snapshot（kickoff）

| RC2 | 狀態 |
| --- | --- |
| P1 Representability | 🟢 Stable |
| P2 Inferability | ▶ Active |
| P3 Composition | ⏸ Locked |
| Interaction Knowledge | 🟡 Research Justified |
