# Stable Knowledge — Maintenance Governance

> **Classification**: post-research governance — **not** a Research Cycle plan.  
> **Applies after**: Pattern 🟢 Stable · Composition 🟢 Stable · Interaction 🟢 Stable（RC1 + RC2 closed）.  
> **Does not authorize**: RC3 · new Knowledge Layer · vocabulary expansion by default.

## Core policy

**Stable knowledge evolves through evidence mapping, not vocabulary expansion.**

Research cycles（Representability → Inferability → Composability）**已結束**的 layer，進入本節奏。新 incident **不**自動重開 research。

---

## Stable Maintenance Dogfood（固定流程）

**不叫** RC2 Intake · Research Intake · 新 Entry 收集。

```text
New Incident
      │
      ▼
Layer First                    ← Primary gate（FAIL → stop; do not add entry）
      │
      ▼
Existing Knowledge Mapping     ← Pattern entry · Composition constraint · Interaction entry/constraint
      │
      ▼
PASS
      │
      ▼
Archive                        ← consumer evidence + optional Ai-skill maintenance log
```

### Allowed outcomes

| Outcome | Action |
| --- | --- |
| **PASS** | Map to existing entry / constraint · archive · **no schema change** |
| **Layer First FAIL** | Document misclassification · **do not** add entry · fix layer routing discipline |
| **Boundary Break** | Escalate → **Vocabulary Exit Review** or **Readiness Gate**（新 layer）— **not** silent entry add |

### Forbidden defaults

| Forbidden | Why |
| --- | --- |
| Incident → new Entry | Research-era habit；Stable 期預設 reuse |
| Incident → schema field | Vocabulary expansion 需 exit review |
| Composition pressure → edit `entries/*` | Anti back-propagation invariant |
| Interaction pressure → edit Pattern/Composition frozen paths | RC2 invariant |

---

## Knowledge surfaces（mapping targets）

| Layer | Stable artifacts | Mapping unit |
| --- | --- | --- |
| **Pattern** | `ui-pattern-knowledge/entries/*.yaml` | Pattern entry id |
| **Composition** | `ui-pattern-knowledge/composition_rules.yaml` · `compositions/*.yaml` | Constraint id |
| **Interaction** | `ui-interaction-knowledge/entries/*.yaml` | Interaction entry id |
| **Interaction Composition** | `ui-interaction-knowledge/interaction_composition_rules.yaml` | Interaction constraint id |

**Neighbor layers**（Continuation · Navigation · Pagination_runtime · Hazard · Runtime）— classify correctly · **do not** absorb into Interaction without Readiness evidence.

---

## Reopen signals（僅兩條）

### 1. Vocabulary Exit Review

Trigger when **all** hold:

- Layer First **PASS**（已判對層）
- ≥2 independent incidents **cannot** map with existing four-field Interaction vocabulary **and** composition constraints insufficient
- Documented in consumer evidence + stakeholder review

**Not** triggered by: single incident · decoy miss · blind LLM layer confusion（→ method protocol fix）.

See [`plans/active/2026-07-14-0856-ui-pattern-knowledge-workflow/evidence/rc2-vocabulary-exit-review.md`](../../plans/active/2026-07-14-0856-ui-pattern-knowledge-workflow/evidence/rc2-vocabulary-exit-review.md).

### 2. Readiness Gate（新 Knowledge Layer）

Trigger when **all** hold（RC1 Readiness 對稱）:

- Pattern selection **validated** for case
- Composition constraints **validated** for case
- Flow / lifecycle **still fails** — constraint **type** insufficient（非「規則不夠用」）

Continuation · Navigation · Evidence · Governance layers are **Neighbor** until this gate passes — **not** Research Justified by incident count alone.

---

## Consumer ↔ Ai-skill writeback

| Repo | Role |
| --- | --- |
| `<PROJECT_ROOT>` | Evidence Producer — full incident anchors |
| Ai-skill plan `evidence/` | Closure Authority — generalized maintenance logs only when review needed |

Pilot consumer rule: `<PROJECT_ROOT>/.ai-skill/project/rules/stable-maintenance-dogfood.md`

Research-era writeback rule（archived for reference）: `rc2-consumer-evidence-writeback.md` — superseded by Stable Maintenance for post-RC2 work.

---

## Method maturity（刻意不升格）

| Object | Status | Note |
| --- | --- | --- |
| Knowledge Evolution Method | 🟡 **Replicated once** | RC1 + RC2 = UI Knowledge Family；**不**追求 🟢 Stable until **different family** replicates full ladder |
| Research Cycle 3 | **Not planned** | No automatic layer growth |

Method Validation Log: [`governance/lifecycle/architecture-evolution-protocol.md`](../../governance/lifecycle/architecture-evolution-protocol.md#method-validation-log)

---

## Research closure reference

| Cycle | Status | Retrospective |
| --- | --- | --- |
| RC1 Pattern + Composition | ✅ Closed | [`research-cycle-1.md`](../../plans/active/2026-07-14-0856-ui-pattern-knowledge-workflow/evidence/research-cycle-1.md) |
| RC2 Interaction | ✅ Closed | [`research-cycle-2.md`](../../plans/active/2026-07-14-0856-ui-pattern-knowledge-workflow/evidence/research-cycle-2.md) |

**Research Closure Hygiene**（plan sync · stale summary cleanup · maintenance handoff）— prerequisite for treating RC2 as **truly** ended. See plan `_plan.md` §Research Closure Hygiene.
