# Phase 3 Start Lock — Composability

**Plan**: [`../../2026-07-14-0856-ui-pattern-knowledge-workflow.md`](../../2026-07-14-0856-ui-pattern-knowledge-workflow.md)  
**Date**: 2026-07-14  
**Status**: Phase 3 **formally started**

---

## Phase 3 Invariant（升格 — 不可違反）

```text
Composition evidence MUST NOT modify validated Pattern Knowledge entries.

New knowledge discovered during composition flows into:
  composition_rules.yaml   (= Composition Constraints)
NOT
  entries/*.yaml
```

**Why**: Phase 2 證明的是 Entry 具有 inferability。若每遇到一個 Screen 就回改 Entry，無法分辨「Entry 正確」還是「Episode 在修 Entry」——破壞研究可驗證性與 Phase 之間假說獨立性。

**Risk name**: **Back-propagation（回滲）** — Phase 3 最大風險不是 Composition 本身，而是回滲污染已驗證的 Pattern Knowledge。

---

## H4–H6 = 三種不同 Evidence（勿混寫）

### H4 — Independence

```text
Pattern A + Pattern B  →  Decision Boundary(A) unchanged
```

例：Bottom Sheet 加入 Toast → Selection Rules（sheet）不變。

### H5 — Completeness

```text
Screen UI nodes  →  each has Pattern | explicit deferred
                 →  zero Unknown
```

不是 100% Coverage；是 **沒有 Unknown**。

### H6 — Traceability（鏈，非散文）

```text
Episode Detail
    │
    ▼
Bottom Sheet
    │
    ▼
Selection Rule
    │
    ▼
Implementation Recipe
```

任何 Screen node 都必須可畫成這條鏈；禁止「不知道為什麼存在」。

---

## Exit Gate = Composition Closure（不是「完成 Screen」）

| 不要 | 要 |
| --- | --- |
| Episode Detail 完成（產品語意） | Episode Detail **Pattern Tree Validated** |
| 完成了多少 UI | Knowledge 能否組成一個 Screen |

**Composition Closure** = Pattern Tree 在 H4–H6 分型證據下通過，且整輪 **零** entries/*.yaml 變更（Invariant 抽查）。

---

## Explicit non-promotion

若 Phase 3 暴露 Interaction / State Transition 問題，可能自然導向下一層（Flow / Orchestrability）。**現在不寫進 plan Phase 4**——尚無證據；僅作觀察備註。
