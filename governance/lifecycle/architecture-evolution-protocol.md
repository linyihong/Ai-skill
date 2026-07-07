# Architecture Evolution Protocol

> **Status**: promoted protocol (normative). **Not** an ADR appendix.
> **Sibling docs**: [`decision-promotion-pipeline.md`](decision-promotion-pipeline.md) (where decisions land) · [`governance-pattern-template.md`](governance-pattern-template.md) (how mechanical locks are built).
>
> **Scope**: ADR-013 solved one architecture problem. **This file solves how the system evolves architecture problems going forward** — how a new **Runtime Contract** enters Ai-skill. The archived ADR-013 plan is **Instance 1 evidence**, not the owner of this protocol.

## Positioning

| Layer | What it answers | Example |
|---|---|---|
| **ADR** | What did we decide, and why? | ADR-013 — Reject D1; Accept D2 (`context.stance`) |
| **This protocol** | How does any accepted decision become a live, verified, non-drifting runtime contract? | Seven-step Integration Phase + Evolution Contract |
| **Governance-pattern template** | How is each mechanical lock shaped? | Observation → Registry → Executor → Validation |

A methodology cannot be proven by one case. The intended maturity path:

```text
Instance 1 (ADR-013)
  → Abstract (this file)
  → Instance 2 (ADR-014)
  → Still holds without editing this file
  → Protocol mature
```

## Contract stack

Four contract types stack. Do not conflate them.

| Contract | Owns | Example |
|---|---|---|
| **Evolution Contract** | **This protocol** — mandatory phase order, no skip, violation shapes | §Evolution Contract below |
| **Step Contract** | Entry / exit criteria per integration step | §Step contracts |
| **Runtime Contract** | Executable fields `runtime validate` must understand | `requires_context.stance`, `capability-context.yaml` |
| **Governance Contract** | Owner-layer policy prose + companion YAML | `governance/cognitive-stance.md` |

**Evolution Contract > Step Contract > Runtime Contract > Governance Contract** for *process*. Runtime Contract is source of truth for *semantics* once established.

## Evolution Contract (meta — mandatory)

The seven integration steps are **not suggestions**. They are the **Architecture Evolution Contract**.

### No skip rule

**Question:** What steps can be skipped?

**Answer:** **None** — for any evolution that introduces or materially changes a Runtime Contract.

Waivers require a **new ADR or plan amendment** that documents why the skip is safe, plus stakeholder sign-off. Silent skip is a violation.

### Violation shapes (anti-patterns)

| Violation | Why it breaks the protocol |
|---|---|
| Discussion → patch Runtime → backfill ADR later | Decision without evidence; constitution drift |
| ADR Accept → implementation → **no** Regression | Contract unprotected; regressions invisible |
| ADR Accept → implementation → **no** Drift Lock when navigation canon exists | Navigation re-defines runtime; three-layer mismatch returns |
| ADR Accept → implementation → **no** Debt Payoff | Legacy paths bypass new contract forever |
| Close ADR before Integration Phase complete | Completed ADR without verified alignment |
| Extend closed ADR scope instead of new ADR/plan | Scope creep; Instance 2 polluted |

Future mechanical enforcement of Evolution Contract violations is **deferred** (no validator yet). Plan reviewers and ADR §Related must treat violations as **blocking** until enforcement exists.

## Layer dependency invariant

**Each layer may depend only on the layer below for contracts, and may never redefine them.**

Dependency direction (authority flows **up**; description flows **down**):

```text
ADR (Architecture Decision)
  ↑ governed by evidence
Governance Contract (owner policy)
  ↑ must match ADR
Runtime Contract (executable fields)
  ↑ navigation describes, never redefines
Navigation (execution-flow, README, taxonomy, routing index)
```

| Layer | May | Must not |
|---|---|---|
| **Navigation** | Describe, link, invoke, thin-surface the Runtime Contract | Redefine capability context, workflow boundaries, or consumer ownership |
| **Runtime** | Validate invoke envelopes, registry fields, projections | Contradict Accepted ADR |
| **Governance** | Own contract prose and companion YAML | Become a second runtime source of truth |
| **ADR** | Record irreversible decision + Completed only after protocol | Stay open while Integration Phase incomplete |

Capability-context evolutions add the **execution invariant**:

```text
Workflow invokes → Capability declares requires_context → Runtime validates → Consumer loads bodies
```

Workflow **must not branch** on contract fields (e.g. `context.stance`).

## Thesis — Integration Phase (seven steps)

When introducing or changing a **Runtime Contract**, execute:

```text
Architecture Decision
  → Runtime Contract
  → Navigation Alignment
  → Regression
  → Drift Lock
  → Debt Payoff
  → Close ADR
```

That sequence is the **Integration Phase** (Stage E). It follows decision evidence and contract establishment (Stage A–D). See [Full lifecycle map](#full-lifecycle-map).

## When to use

| Use this protocol | Do not use |
|---|---|
| New or changed **runtime-readable** contract | Single doc fix with no contract semantics |
| ADR-accepted boundary propagating to routing, navigation, validators | Editorial README-only polish |
| Cross-layer drift risk | Mechanical lock onboarding only → [`governance-pattern-template.md`](governance-pattern-template.md) |
| Decision classified for `constitution/ADR-*` | Session-scoped decision with no durable contract |

## Full lifecycle map

| Stage | Purpose | Typical deliverables |
|---|---|---|
| **A — Problem & evidence** | Separate symptoms from architecture; generalization evidence | Plan §Decision Rationale; companion evidence |
| **B — Taxonomy / naming** | Converge vocabulary before Accept | Matrix; anti-explosion rule |
| **C — Architecture Decision** | ADR **Accepted** | `constitution/ADR-*` |
| **D — Contract establishment** | Owner contract + registry + enforcement + consumer + dogfood | `governance/*`, `runtime/*`, consumer dir |
| **E — Integration (7 steps)** | Runtime + navigation + locks + debt + regression + ADR close | §Integration steps |
| **F — Post-close evolution** | Taxonomy growth, hard blocks — **new ADR/plan only** | e.g. ADR-014 |

## Integration steps

### Conceptual vs implementation order

**Conceptual** (explain the protocol):

```text
Architecture Decision → Runtime Contract → Navigation Alignment
  → Regression → Drift Lock → Debt Payoff → Close ADR
```

**Implementation** (ADR-013 validated order):

```text
Runtime consistency (2.1)
  → Navigation alignment + Drift Lock (2.2)
  → Debt Payoff + Ownership Lock (2.3)
  → Contract regression (2.4; may parallel 2.1)
  → Close ADR
```

### Step contracts

| Step | Entry criteria | Exit criteria | Mechanical lock |
|---|---|---|---|
| **1. Architecture Decision** | Problem framed; alternatives rejected; generalization evidence | ADR **Accepted** | — |
| **2. Runtime Contract** | ADR maps to owner YAML/MD + registry | `runtime validate` understands fields; enforcement path exists | e.g. `capability_registry` |
| **3. Navigation Alignment** | Runtime contract is SoT | Thin execution-flow; fat README; taxonomy → consumer | Required when navigation canon exists |
| **4. Regression** | Contract testable | Scenario YAML + Go tests (pass/warn/mismatch) | `validation/scenarios/` |
| **5. Drift Lock** | Navigation canon enumerated | Validator forbids deprecated shapes in canon | e.g. doc drift validator |
| **6. Debt Payoff** | Debt classes A–D identified | Active canonical tree cleared | e.g. ownership drift validator |
| **7. Close ADR** | Steps 2–6 complete | ADR **Completed**; plan archived; links updated | All locks green |

### Debt classes (step 6)

| Class | Question |
|---|---|
| **A — Ownership Drift** | Does any doc claim artifacts owned by the capability/contract? |
| **B — Path Drift** | Do active sources point at stub or pre-migration paths? |
| **C — Semantic Drift** | Does prose re-introduce rejected models? |
| **D — Artifact Drift** | Do scenarios/templates match capability output shapes? |

## Three-lock pattern

| Lock | Protects |
|---|---|
| **Contract lock** | Runtime schema + registry + invoke semantics |
| **Navigation drift lock** | Canon does not redefine the contract |
| **Ownership drift lock** | Debt classes A–D in active tree |

Each lock follows [`governance-pattern-template.md`](governance-pattern-template.md). This protocol defines **when**; the template defines **how**.

## Dogfood Evidence (required for every important ADR)

Every **Completed** architecture ADR that introduced a Runtime Contract **must** carry a **§Dogfood Evidence** section (in the ADR or linked plan archive). Purpose: **Decision → Evidence** permanently traceable.

### Required fields

| Field | Content |
|---|---|
| **Instance** | ADR id + archived plan link |
| **Runtime tests** | Go tests + `runtime validate` checks |
| **Regression** | Validation scenario YAML paths |
| **Drift locks** | Validator names / source files (if navigation canon exists) |
| **Repository paths** | Owner contract, registry, consumer, navigation canon |
| **Commits** | Primary integration commits |

### Template (copy into ADR §Dogfood Evidence)

```markdown
## Dogfood Evidence

| Field | Record |
|---|---|
| Instance | ADR-NNN · [plan archived link] |
| Runtime tests | `path/to/*_test.go` · `runtime validate` check name |
| Regression | `validation/scenarios/...` |
| Drift locks | validator name · `scripts/.../validator.go` |
| Repository paths | owner · registry · consumer · navigation canon |
| Commits | `abc1234` … |
```

## Validated instances

| Instance | ADR | Plan evidence | Protocol edit required? |
|---|---|---|---|
| **1** | [ADR-013](../../constitution/ADR-013-cognitive-role-primitive-gate.md) | [archived plan](../../plans/archived/2026-07-06-review-architecture-adr/_plan.md) | N/A (source instance) |
| **2** | [ADR-014](../../constitution/ADR-014-cognitive-stance-capability-context.md) (Proposed) | *pending Accept* | **Pass criterion:** complete Integration Phase **without editing this file** |

### Second-instance gate (maturity)

Protocol is **promoted but not mature** until Instance 2 completes:

```text
Instance 1 → abstract → Instance 2 → still holds → mature
```

**Mature** means: ADR-014 (or next Runtime Contract evolution) runs Stages A–F using this file as-is; only **§Validated instances** table and ADR §Dogfood Evidence grow — not the step contracts, Evolution Contract, or invariants.

## Success criteria (protocol-level)

**Not:** “ADR-013 succeeded.”

**Yes:** Any future Runtime Contract evolution — stance taxonomy, new capability context, invocation metadata, runtime governance surface — **completes using this protocol without modifying this file**.

| Signal | Meaning |
|---|---|
| Instance 2 closes with zero protocol edits | Generalization likely |
| Instance 2 requires protocol amendment | Gap found; amend with evidence, not ad-hoc |
| Instance 3+ reuses same seven steps | Protocol is system infrastructure |

This is the long-term value bar: **Architecture Evolution Protocol as Gen 3 infrastructure**, not ADR-013’s summary.

## Lineage — where this protocol came from

### Origin (2026-07)

Software-delivery **three-layer contract mismatch** (governance / execution-flow / README). Symptom fixes rejected.

| Version | Approach | Outcome |
|---|---|---|
| v1 | Add review link | Symptom fix — rejected |
| v2 | ADR-first delivery model | → ADR-013 |
| v3 | `cognitive_role` primitive (D1) | Rejected (Phase 0b) |
| v4 | `context.stance` invoke (D2) | Accepted |

### Abstraction path

```text
Review execution gap
  → ADR-013 (Instance 1)
  → Pattern named in plan Phase Map
  → Stakeholder: value exceeds Review
  → architecture-evolution-protocol.md (this protocol)
  → ADR-014 queued as Instance 2
```

### Evidence chain

| Artifact | Role |
|---|---|
| [`.cursor/plans/review_workflow_slice_12ecc222.plan.md`](../../.cursor/plans/review_workflow_slice_12ecc222.plan.md) | Historical entry |
| [ADR-013 plan archive](../../plans/archived/2026-07-06-review-architecture-adr/_plan.md) | Instance 1 primary evidence |
| [ADR-013](../../constitution/ADR-013-cognitive-role-primitive-gate.md) | Decision + §Dogfood Evidence |
| [ADR-014](../../constitution/ADR-014-cognitive-stance-capability-context.md) | Instance 2 candidate |

## Instance 1 — ADR-013 Dogfood Evidence

| Field | Record |
|---|---|
| **Instance** | ADR-013 · [`2026-07-06-review-architecture-adr`](../../plans/archived/2026-07-06-review-architecture-adr/_plan.md) |
| **Runtime tests** | `capability_context_test.go` · `capability_registry` in `runtime validate` |
| **Regression** | `validation/scenarios/runtime/capability-stance-contract-regression-v1.yaml` · `capability-stance-fault-finding-v1.yaml` |
| **Drift locks** | `review_architecture_doc_drift` · [`documentation_drift.go`](../../scripts/ai-skill-cli/internal/app/documentation_drift.go) · `canonical_ownership_drift` · [`canonical_ownership_drift.go`](../../scripts/ai-skill-cli/internal/app/canonical_ownership_drift.go) |
| **Repository paths** | [`runtime/capability-context.yaml`](../../runtime/capability-context.yaml) · [`capability-registry.yaml`](../../knowledge/runtime/capability-registry.yaml) · [`cross-cutting/review/`](../../workflow/cross-cutting/review/README.md) · [`execution-flow.md`](../../workflow/software-delivery/execution-flow.md) · [`cognitive-stance.md`](../cognitive-stance.md) |
| **Commits** | `948f2f1` · `5c732fa` · `fb9cdc1` |

## Plan checklist (agent-facing)

- [ ] §Decision Rationale: symptom vs architecture separated
- [ ] Phase 0 evidence before ADR Accept ([`decision-promotion-pipeline.md`](decision-promotion-pipeline.md))
- [ ] Stage D: owner + registry + enforcement + consumer + dogfood
- [ ] Stage E: Runtime → Navigation → Debt → Regression — **no skip**
- [ ] Drift Lock when navigation canon exists
- [ ] Debt classes A–D checked or waived with documented reason
- [ ] ADR **Completed** only after locks + regression green
- [ ] ADR §Dogfood Evidence filled
- [ ] Plan archived; protocol linked from ADR §Related
- [ ] Post-close scope → new ADR/plan only

## Related

- [`decision-promotion-pipeline.md`](decision-promotion-pipeline.md)
- [`governance-pattern-template.md`](governance-pattern-template.md)
- [`system-upgrade-governance.md`](system-upgrade-governance.md)
- [`../../architecture/ai-native-cognitive-execution-system.md`](../../architecture/ai-native-cognitive-execution-system.md)
- [`../../constitution/ADR-013-cognitive-role-primitive-gate.md`](../../constitution/ADR-013-cognitive-role-primitive-gate.md) — Instance 1
- [`../../constitution/ADR-014-cognitive-stance-capability-context.md`](../../constitution/ADR-014-cognitive-stance-capability-context.md) — Instance 2 candidate
