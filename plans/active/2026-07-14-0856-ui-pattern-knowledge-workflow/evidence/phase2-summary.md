# Phase 2 Closure Evidence — Pattern Knowledge Inferability

**Plan**: [`../../2026-07-14-0856-ui-pattern-knowledge-workflow.md`](../../2026-07-14-0856-ui-pattern-knowledge-workflow.md)  
**Date**: 2026-07-14  
**Gate decision**: Phase 2 **Completed** → Phase 3 **Start**

---

## Original hypothesis

Phase 2 要驗證的不是「寫完五個 Entry」，而是三個假說：

| # | Hypothesis |
| --- | --- |
| H1 | **Selection Rules 可推理** — Pattern Knowledge 能支撐選型，不是靜態詞典 |
| H2 | **Family 有決策意義** — Overlay vs Feedback 真正參與 Selection |
| H3 | **Near Neighbor 有效** — Sheet vs Dialog vs Drawer 形成 decision boundary |

**設計原則（耐久）**：Phase 2 驗證的是 Pattern Knowledge 的 **可推理性（inferability）**，而不是 Pattern **Coverage（涵蓋率）**。完成條件以推理能力為準，而非 Pattern 數量。

---

## Evidence

| Artifact | 角色 |
| --- | --- |
| [`selection-scenarios.yaml`](selection-scenarios.yaml) | 十案定義（每 pattern ×2） |
| [`2a-family-inferability-run.md`](2a-family-inferability-run.md) | rule-trace + blind LLM 矩陣 |
| `entries/*.yaml`（五件）+ toast.`family=feedback` | 推理對象（非「數量 KPI」） |

代表性命中：付款方式 → `bottom_sheet`；確定刪除 → `modal_dialog`；已儲存 → `toast`。

---

## Result

| Hypothesis | Result | 一句 |
| --- | --- | --- |
| H1 Selection | ✅ PASS | rule-trace 10/10 與 blind LLM 10/10 一致 |
| H2 Family | ✅ PASS | toast 為第一個非-overlay family，仍可正確選型 |
| H3 Near Neighbor | ✅ PASS | Sheet / Dialog / Drawer 邊界可穩定選對 |

**Decision**: Pattern Knowledge Layer **validated**（inferability）。

---

## Unexpected findings

1. **Scrim 不是主表面選型** — 它是 compositional capability；必須在 scenario 裡拆「任務面」vs「壓暗層」，否則會把刪除確認錯答成 scrim。
2. **Family 改標立即改變競爭集** — toast 若仍標 `overlay`，會與 modal/sheet 搶同一決策空間；標 `feedback` 後 S9/S10 自然脫離 Overlay Decision。
3. **盲測與 rule-trace 零分歧** — 本輪未出現「文件寫得通、模型推不出」；若未來出現，優先懷疑 near_neighbor 模糊，而不是再開更多 entry。

---

## Decision

| Field | Value |
| --- | --- |
| Pattern Knowledge Layer | **validated**（inferability） |
| Phase 2 | **Completed** |
| Phase 3 | **Start** — Pattern Composition（Episode Page 先行） |
| Entry freeze | Phase 3 **不得**為修問題回頭改 `entries/*.yaml`；新增 `composition_rules` |

## Known limitation

| Limitation | Note |
| --- | --- |
| Only Overlay (+ toast Feedback) tested | Navigation / form / data-display families **not yet verified** |
| Coverage intentionally thin | 五件不是成熟度指標 |
| Consumer project alias | 次要；未擋 Phase 2 gate |
| App Bar / Player | 進 Phase 3 composition 時可列結構節點，**不**等於已驗證 entry |
