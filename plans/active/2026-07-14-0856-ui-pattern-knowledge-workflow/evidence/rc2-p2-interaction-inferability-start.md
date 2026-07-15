# RC2-P2 Start Lock — Interaction Inferability

**Plan**: [`../_plan.md`](../_plan.md)  
**Date**: 2026-07-15  
**Status**: RC2-P2 **Active**（stakeholder kickoff complete — 對齊 RC1-P2 Active 慣例）  
**Prerequisite**: RC2-P1 🟢 Stable — [`rc2-p1-interaction-representability-closure.md`](rc2-p1-interaction-representability-closure.md)

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

## 2. Second entry（P2 前置 — 仍用 P1 vocabulary）

| 選擇 | ID | 理由 |
| --- | --- | --- |
| defer P2 dogfood | `preview_gate_transition` | P1 已用 — **禁止**作為 P2 唯一案例 |
| ✅ **P2 案例** | `payment_leave_transition` | C2 counterfactual；第二獨立 consumer；證明 schema **泛化** |

**Scope lock**（同 P1 紀律）：

```text
dialog_open → user_confirm_stay_or_leave → dialog_close
```

**不是** `order_paid` / `payment_pending` business states。

**Consumer anchor**：Readiness C2 · [`r1-consumer-dogfood-2026-07-15.md`](r1-consumer-dogfood-2026-07-15.md) §C2

**前置產物**（P2 dogfood 前）：

1. [`entries/payment_leave_transition.yaml`](../../../../workflow/software-delivery/ui-interaction-knowledge/entries/payment_leave_transition.yaml) — 同四欄 vocabulary；Schema Extensions = 0
2. [`interaction-inferability-scenarios.yaml`](interaction-inferability-scenarios.yaml) — incident prompts + expected layer + entry + repair

---

## 3. RC2-P2 Metrics（non-KPI）

| Metric | Role | P2 目標 |
| --- | --- | --- |
| **Boundary Misclassification** | **Primary** | 0（Flow→Composition / Flow→Pattern = FAIL） |
| Scenario Accuracy | Supporting | IH1：正確 entry id |
| Frozen Layer Mods | Blocking | **0** always |

**不是 KPI**：scenario 數量、entry 數量、覆蓋率。

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
