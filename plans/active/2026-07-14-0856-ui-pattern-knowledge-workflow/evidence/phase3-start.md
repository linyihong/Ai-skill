# Phase 3 Start Lock — Composability

**Plan**: [`../_plan.md`](../_plan.md)  
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
| 1 | **H4 Independence** | 施壓；Failure→Constraint；Invariant stress-validated |
| 2 | **H5 Completeness** | Every Deferred has **disposition**（非 Deferred→0） |
| 3 | **H6 Traceability** | complete\|waived 終點；T1/T2/T3；Waiver first-class |

H6 鏈形：

```text
Episode Detail
└── Bottom Sheet
      ├── selection_rule
      └── implementation_recipe
```

---

## Composition Metrics

| Metric | Role | Expected |
| --- | --- | --- |
| **Entry Modifications** | **Primary** | **0** |
| Composition Rule Count | Supporting | Δ can ↑ |
| Deferred Nodes | Supporting | disposition 齊全 > 歸零 |

見 [`3-metrics.md`](3-metrics.md)。

---

## Exit Gate = Composition Closure

**Pattern Tree Validated** = H4∧H5∧H6 mini-cycles PASS + Entry Modifications = 0。  
不是「Episode Detail 完成」。

**Protocol trio**：Growth · Constraint Accumulation · Explicit Termination。

## Explicit non-promotion

Flow / Orchestrability 在 Readiness 已升格為 RC2（2026-07-15）；本句為 Phase 3 當時紀律。
