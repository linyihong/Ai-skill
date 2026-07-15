# RC2-P1 Start Lock — Interaction Representability

**Plan**: [`../_plan.md`](../_plan.md)  
**Date**: 2026-07-15  
**Status**: RC2-P1 **formally started**（stakeholder kickoff）  
**Readiness**: [`phase4-readiness-gate.md`](phase4-readiness-gate.md) ✅ Closed · Evidence [`r1-consumer-dogfood-2026-07-15.md`](r1-consumer-dogfood-2026-07-15.md)

---

## Methodology lock（與 RC1 平行）

RC2-P1 **不**回答「Interaction 長什麼樣」，先回答 **Interaction 的最小可表示單位（minimum representable unit）是什麼** — 對應 RC1 Phase 1 Representability。

| RC1 | RC2 |
| --- | --- |
| Representability | **Interaction Representability**（本 phase） |
| Inferability | Interaction Inferability（RC2-P2） |
| Composability | Interaction Composition（RC2-P3） |

驗證的是 **Knowledge 是否可表示**，不是 Schema 文件膨脹。

Cross-method：[`Architecture Evolution Protocol` §Layer Growth Rhythm](../../../../governance/lifecycle/architecture-evolution-protocol.md#appendix--layer-growth-rhythmoptional-governance-pattern) — Pattern + Composition **frozen**；Interaction 為新 surface。

**Canonical boundary（比 Schema 更重要）**

> Interaction Knowledge describes **UI interaction semantics**, not application business workflow.

例：`dialog_open` → `user_confirm` → `dialog_close` ✅ · `order_paid` / payment-order `pending` ❌

---

## Vocabulary Freeze（RC2 對稱 RC1 Entry Freeze）

第一個 entry **representability exit** 完成前：

- **不得**新增 interaction primitive（schema 或 entry）
- 現行四欄：`state_owner` · `transition_trigger` · `invalidation_event` · `recovery_boundary`
- Dogfood 若覺得需要 `guard_condition` / `rollback` / `checkpoint` / `priority` → **只寫 evidence** → dogfood 結束 → Review → 再決是否擴 schema

---

## RC2 Metrics（non-KPI）

| Metric | RC2-P1 target |
| --- | --- |
| **Schema Extensions** | **0** until exit review |
| **Interaction Entry Mods** | **0** after first entry lands |
| **Frozen Layer Mods** | **0** always |

---

## `preview_gate_transition` scope lock

**不是**完整 Player State Machine。只表示：

```text
preview → preview_limit_reached → gated
```

目的：證明四欄位足夠 — **不是**證明 Interaction Schema 很完整。

---

## 1. Success Definition

### H1 — Representability

**問題**：Interaction Knowledge **最小**可以表示什麼？

**不是**：能表示所有 Flow。

**足夠的最小形狀**（候選 — RC2-P1 驗證目標，非最終 API）：

```yaml
interaction:
  state_owner:        # 誰擁有這段 UI 狀態
  transition_trigger: # 什麼事件觸發轉移
  invalidation_event: # 什麼事件後狀態不可再信
  recovery_boundary:  # 什麼證據使狀態再次可信
```

**RC2-P1 明確排除**（留後續 phase）：`state_machine`、`timeline`、`event_graph`、`async`、`animation`、`retry`。

### H2 — Layer Boundary（比 Schema 更重要）

**問題**：什麼 **不是** Interaction Knowledge？

| 留在 RC1（frozen） | 升格到 Interaction？ |
| --- | --- |
| Pattern Selection | ❌ |
| Composition Constraint | ❌ |
| Overlay Family | ❌ |
| Spatial Relationship | ❌ |

**Interaction 只負責**：

- State ownership
- Transition trigger
- Invalidation
- Recovery

若 Boundary 不清 → RC2 膨脹為「另一套 Composition」。

### H3 — First Entry Shape（Coverage ≠ 目標）

**只做一個 entry** — stakeholder 鎖定：

| 選擇 | ID | 理由 |
| --- | --- | --- |
| ✅ **選定** | `preview_gate_transition` | C1 R1 最強：真實 incident + integration；`modal_dialog` 已 validated |
| defer | `payment_leave_transition` | C2 為 counterfactual；可作 RC2-P2 候選 |

**目的**：驗證最小形狀 **能不能表示** C1 的 Flow 失敗（listener owner / `preview`→`gated`），不是證明所有 Flow。

**Consumer anchor**（sanitized）：`<PROJECT_ROOT>` player preview gate · Readiness C1 · `r1-consumer-dogfood` §C1。

---

## 2. RC2 Invariant（anti back-propagation）

對稱 RC1 Phase 3：

```text
Interaction evidence MUST NOT redefine Pattern selection or Composition constraints.

Interaction may REFERENCE frozen layers (e.g. modal_dialog entry id).
Interaction MUST NOT edit:
  ui-pattern-knowledge/entries/*.yaml
  ui-pattern-knowledge/composition_rules.yaml
```

**責任鏈**（單向依賴）：

```text
Pattern
    │
    ▼
Composition
    │
    ▼
Interaction   ← 只能建立 state transition ON TOP OF validated pattern/composition
```

**禁止語意**：Interaction 不能說 `dialog_requires_scrim` 錯了；只能建立 `modal_dialog` → `preview`→`gated` transition。

**Violation** ⇒ RC2-P1 evidence invalid until Interaction-layer edits revert frozen files.

---

## 3. Interaction Hazard Review — Boundary（Canonical）

**English**

> Interaction Knowledge describes how interaction **should** behave; Interaction Hazard Review evaluates how interaction **can fail**.

| Interaction Knowledge | Interaction Hazard Review |
| --- | --- |
| Canonical lifecycle | Failure taxonomy |
| State ownership | Hazard identification |
| Transition semantics | Risk evaluation |
| Recovery boundary | Prevention evidence |

**共享 vocabulary**；**不共享責任**。

- Knowledge canonical：`workflow/software-delivery/`（RC2 產物待定目錄）
- Hazard Review pilot：`<PROJECT_ROOT>` `.ai-skill/project/rules/interaction-hazard-review.md` + Ai-skill `plans/active/2026-06-16-1030-interaction-hazard-review-workflow.md`

RC2-P1 **不**合併兩者 artifact；只釘清邊界。

---

## 4. Exit Gate（RC2-P1）

| 通過條件 | 不通過 |
| --- | --- |
| 第一個 Interaction entry（`preview_gate_transition`）**可表示** H1 四欄 | 為了表示而回寫 Pattern/Composition |
| H2 boundary 文件化且無反例升格 | Interaction 吸收 Selection / Spatial constraint |
| Entry Modifications on `entries/*` = **0** | 任何 `composition_rules.yaml` 編輯 |
| Composition Modifications = **0** | — |

**不是** Exit：「所有 Flow 都能表示」。

**是** Exit：**一個** validated Interaction entry 證明 minimum representable unit 成立。

---

## 5. Kickoff 五問（結案 checklist）

| # | 問題 | 本檔答案 |
| --- | --- | --- |
| 1 | Success Definition：最小可表示單位？ | H1 四欄 `interaction.*` |
| 2 | Layer Boundary：Interaction 不負責什麼？ | H2 表；不升格 RC1 四類 |
| 3 | First Entry：一個真實案例驗證表示法 | `preview_gate_transition`（C1） |
| 4 | Invariant：不得重新定義 Pattern/Composition | §RC2 Invariant |
| 5 | Exit Gate：一 entry 可表示，非全 Flow | §Exit Gate |

---

## Explicit non-goals（RC2-P1）

- [ ] ~~Interaction schema 覆蓋所有 async UI~~
- [ ] ~~Runtime projection~~
- [ ] ~~合併 interaction-hazard-review 為同一 artifact~~
- [ ] ~~第二 entry（payment_leave）~~ — defer RC2-P2 前

---

## Next execution

| Step | Status | Artifact |
| --- | --- | --- |
| 1. interaction-entry-schema | ✅ | [`ui-interaction-knowledge/validation/interaction-entry-schema.yaml`](../../../../workflow/software-delivery/ui-interaction-knowledge/validation/interaction-entry-schema.yaml) |
| 2. preview_gate_transition entry | ✅ | [`entries/preview_gate_transition.yaml`](../../../../workflow/software-delivery/ui-interaction-knowledge/entries/preview_gate_transition.yaml) |
| 3. Dogfood | ✅ | [`rc2-p1-preview-gate-representability-run.md`](rc2-p1-preview-gate-representability-run.md) |
| 4. Frozen Layer Mods = 0 | ✅ | Metrics in dogfood run |

## Stakeholder evaluation（RC2-P1）

| 項目 | 狀態 |
| --- | --- |
| Method | 🟢 Ready |
| Scope | 🟢 Well bounded |
| Invariant | 🟢 Defined |
| Entry | 🟢 First validation（dogfood run） |
| Schema | 🟢 Vocabulary validated（extensions = 0） |
