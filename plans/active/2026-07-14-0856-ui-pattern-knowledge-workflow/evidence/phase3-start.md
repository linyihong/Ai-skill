# Phase 3 Start Lock — Composability

**Plan**: [`../../2026-07-14-0856-ui-pattern-knowledge-workflow.md`](../../2026-07-14-0856-ui-pattern-knowledge-workflow.md)  
**Date**: 2026-07-14  
**Status**: Phase 3 **formally started**

---

## What Phase 3 reviews

Phase 2 審 **Selection（知識可否推理）**。  
Phase 3 審 **Constraints（知識之間的約束）**。

驗證單位 = Screen；審查方式 ≠ 一次跑完整 Episode。

Cross-method：**Grow one layer, freeze the previous** — [`Architecture Evolution Protocol` §Layer Growth Rhythm](../../../../governance/lifecycle/architecture-evolution-protocol.md#appendix--layer-growth-rhythmoptional-governance-pattern)。

---

## Phase 3 Invariant（升格 — 不可違反）

```text
Composition evidence MUST NOT modify validated Pattern Knowledge entries.

New knowledge discovered during composition flows into:
  composition_rules.yaml   (= Composition Constraints)
NOT
  entries/*.yaml
```

**Risk**: **Back-propagation（回滲）**。違反 Invariant ⇒ 證據作廢。

---

## Mini-cycles（分開跑）

| Order | Cycle | Evidence shape |
| --- | --- | --- |
| 1 | **H4 Independence** | 施壓：±Toast、Dialog↔Sheet；Decision Boundary 不變 |
| 2 | **H5 Completeness** | 每 Node → Pattern \| deferred(`reason: uncovered_pattern`)；零 Unknown（≠ Coverage） |
| 3 | **H6 Traceability** | Deliverable = **Trace Graph**（非 YAML 齊全） |

H6 鏈形：

```text
Episode Detail
└── Bottom Sheet
      ├── selection_rule
      └── implementation_recipe
```

---

## Composition Metrics

| Metric | Expected |
| --- | --- |
| Deferred Nodes | ≥0（誠實） |
| Composition Rule Count | ≥0 |
| **Entry Modifications** | **0** |

見 [`3-metrics.md`](3-metrics.md)。

---

## Exit Gate = Composition Closure

**Pattern Tree Validated** = H4∧H5∧H6 mini-cycles PASS + Entry Modifications = 0。  
不是「Episode Detail 完成」。

## Explicit non-promotion

Flow / Orchestrability 若自然露出，**不**預寫進本 plan Phase 4。
