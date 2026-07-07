# Architecture Evolution Protocol

> **Protocol Status**: **Experimental** (1 validated instance; Instance 2 pending)
> **Not** an ADR appendix. **ADR-013 is complete**; **this protocol has just begun**.
> **Sibling docs**: [`decision-promotion-pipeline.md`](decision-promotion-pipeline.md) · [`governance-pattern-template.md`](governance-pattern-template.md)

## Positioning — three things that were once mixed

These are **not the same thing**. They can each stand alone.

| Layer | Output | Standalone? | Status (2026-07-07) |
|---|---|---|---|
| **Architecture Decision** | ADR-013, ADR-014 | ✅ | ADR-013 **Completed** · ADR-014 Proposed |
| **Evolution Methodology** | **This protocol** | ✅ | **Experimental** — promoted, Instance 2 not yet validated |
| **Governance Pattern** | Three-lock · Debt classes · Second-instance gate | ✅ | Proven in Instance 1; reusable per [`governance-pattern-template.md`](governance-pattern-template.md) |

ADR-013 solved **one** architecture problem (Review / `context.stance`). This protocol solves **how the system evolves architecture problems going forward**. The archived ADR-013 plan is **Instance 1 evidence**, not the owner of this file.

---

## The protocol — three invariants

The protocol is **not** the seven integration steps. The steps are one **chapter** (see [§Integration Phase](#integration-phase-seven-steps)) — they may become 6, 7, or 8 as practice refines. What should not change are these three invariants:

### ① Separation of Concerns

**Each layer may depend only on the layer below for contracts, and may never redefine them.**

```text
ADR (Architecture Decision)
  ↓ must align
Governance Contract (owner policy)
  ↓ must align
Runtime Contract (executable fields)
  ↓ navigation describes, never redefines
Navigation (execution-flow, README, taxonomy, routing index)
```

| Layer | May | Must not |
|---|---|---|
| **Navigation** | Describe, link, invoke, thin-surface the Runtime Contract | Redefine capability context, workflow boundaries, or consumer ownership |
| **Runtime** | Validate invoke envelopes, registry fields, projections | Contradict Accepted ADR |
| **Governance** | Own contract prose and companion YAML | Become a second runtime source of truth |
| **ADR** | Record irreversible decision; **Completed** only after Mechanical Closure | Stay open while integration incomplete |

Capability-context evolutions add the **execution corollary**:

```text
Workflow invokes → Capability declares requires_context → Runtime validates → Consumer loads bodies
```

Workflow **must not branch** on contract fields (e.g. `context.stance`).

### ② Mechanical Closure

**Architecture evolution is not complete until Mechanical Closure holds.**

An ADR is **not** done because the decision was accepted or code was merged. It is done only when **all three** are green:

| Closure component | What it proves |
|---|---|
| **Regression** | The runtime contract is test-protected (scenarios + tests) |
| **Drift Lock** | Navigation canon cannot silently re-define the contract |
| **Debt Payoff** | Legacy paths (ownership / path / semantic / artifact drift) are cleared |

This invariant matters **more than step count**. Skip any component → evolution incomplete → ADR must not close.

Debt classes (Governance Pattern — reusable):

| Class | Question |
|---|---|
| **A — Ownership Drift** | Does any doc claim artifacts owned by the capability/contract? |
| **B — Path Drift** | Do active sources point at stub or pre-migration paths? |
| **C — Semantic Drift** | Does prose re-introduce rejected models? |
| **D — Artifact Drift** | Do scenarios/templates match capability output shapes? |

Three-lock pattern (Governance Pattern — Instance 1 reference):

| Lock | Protects |
|---|---|
| **Contract lock** | Runtime schema + registry + invoke semantics |
| **Navigation drift lock** | Canon does not redefine the contract |
| **Ownership drift lock** | Debt classes A–D in active tree |

Each lock follows [`governance-pattern-template.md`](governance-pattern-template.md). This protocol defines **when**; the template defines **how**.

### ③ Generalization

**A framework is not the first case. A framework is the second case still holding without modifying the protocol.**

```text
Instance 1 (ADR-013) → abstract (this file) → Instance 2 (ADR-014) → no protocol changes → emerging
  → Instance 3+ → no protocol changes → stable
```

**Primary maturity signal:** Did the instance complete using the **existing core invariants and Evolution Contract** without amending this file's invariant sections?

**Secondary signal:** Validated instance count and domain diversity (see [§Protocol Status](#protocol-status)).

Allowed growth **without** a protocol amendment: new rows in §Validated Instances, §Protocol Evidence, and per-ADR §Dogfood Evidence only.

---

## Protocol Status

Status tracks **maturity**, not ADR completion. **Do not** call the protocol mature when only Instance 1 exists.

| Status | Primary condition | Reference (instance count) |
|---|---|---|
| **Experimental** | Protocol promoted; ≤1 instance with Protocol Evidence | 0–1 instances |
| **Emerging** | Instance 2 completed; **no protocol core edits** | 2 instances |
| **Stable** | ≥3 instances; **no protocol core edits** across all | ≥3 instances |
| **Mature** | Multiple **domains**; instances complete without core edits; only editorial updates to this file | ≥3 instances, ≥2 domains |

**Current:** **Experimental** — Instance 1 (ADR-013) only. Instance 2 candidate: ADR-014.

**Stability rule (normative):** Case count is a **reference**, not the proof. The proof is: *new evolution completes under existing invariants; this file gains evidence rows only — not new invariant sections or amended Evolution Contract.*

If Instance 2 requires a protocol core amendment → gap found → amend **with evidence**, bump status back or hold at Experimental until re-validated.

---

## Evidence — Dogfood vs Protocol

Two evidence types. Do not conflate.

| Evidence type | Proves | Lives in |
|---|---|---|
| **Dogfood Evidence** | **This ADR succeeded** — decision → implementation → locks → tests | ADR §Dogfood Evidence (or linked plan archive) |
| **Protocol Evidence** | **The protocol did not need to change** for this instance | This file §Validated Instances + §Protocol Evidence |

### Dogfood Evidence (per ADR)

Every **Completed** architecture ADR that introduced a Runtime Contract **must** carry §Dogfood Evidence.

| Field | Content |
|---|---|
| **Instance** | ADR id + archived plan link |
| **Runtime tests** | Go tests + `runtime validate` checks |
| **Regression** | Validation scenario YAML paths |
| **Drift locks** | Validator names / source files |
| **Repository paths** | Owner contract, registry, consumer, navigation canon |
| **Commits** | Primary integration commits |

Template — copy into ADR:

```markdown
## Dogfood Evidence

| Field | Record |
|---|---|
| Instance | ADR-NNN · [plan archived link] |
| Runtime tests | `path/to/*_test.go` · `runtime validate` check name |
| Regression | `validation/scenarios/...` |
| Drift locks | validator · `scripts/.../validator.go` |
| Repository paths | owner · registry · consumer · navigation canon |
| Commits | `abc1234` … |
```

### Protocol Evidence (per instance)

Recorded **here** when an instance closes. Proves generalization.

| Field | Record |
|---|---|
| **Instance** | N · ADR id |
| **Integration complete** | Yes / No |
| **Protocol core modified** | **Yes / No** — if Yes, link amendment commit + reason |
| **Only evidence rows added** | Yes / No |
| **Mechanical Closure** | Regression ✓ · Drift Lock ✓ · Debt Payoff ✓ |

---

## Validated Instances & Protocol Evidence

| Instance | ADR | Status | Protocol core modified? | Mechanical Closure |
|---|---|---|---|---|
| **1** | [ADR-013](../../constitution/ADR-013-cognitive-role-primitive-gate.md) | ✅ Completed | N/A (source instance) | ✅ |
| **2** | [ADR-014](../../constitution/ADR-014-cognitive-stance-capability-context.md) | Proposed — **Instance 2 candidate** | *pending* | *pending* |
| **3** | — | — | — | — |

**Instance 2 pass criterion:** Complete Stages A–F + Integration Phase; **Protocol Evidence: No protocol core modified**; only this table and ADR-014 §Dogfood Evidence grow.

**Stable criterion:** Instance 3+ with same **No protocol core modified** column — then Protocol Status may advance to **Stable**.

---

## Contract stack

Four contract types. Do not conflate.

| Contract | Owns | Example |
|---|---|---|
| **Evolution Contract** | This protocol — no skip, violation shapes | §Evolution Contract below |
| **Step Contract** | Entry / exit per integration step | §Integration Phase |
| **Runtime Contract** | Fields `runtime validate` must understand | `requires_context.stance` |
| **Governance Contract** | Owner-layer policy + companion YAML | `governance/cognitive-stance.md` |

**Evolution Contract > Step Contract > Runtime Contract > Governance Contract** for *process*. Runtime Contract is SoT for *semantics* once established.

### Evolution Contract (no skip)

For any evolution that introduces or materially changes a Runtime Contract: **no step may be silently skipped.**

| Violation | Why it breaks the protocol |
|---|---|
| Discussion → patch Runtime → backfill ADR | Decision without evidence |
| Implementation → **no** Regression | Mechanical Closure broken |
| Implementation → **no** Drift Lock (when navigation canon exists) | Separation broken |
| Implementation → **no** Debt Payoff | Legacy bypass persists |
| Close ADR before Mechanical Closure | Incomplete evolution |
| Extend closed ADR instead of new ADR/plan | Generalization polluted |

Waivers require documented ADR/plan amendment + stakeholder sign-off. Mechanical enforcement deferred; reviewers treat violations as **blocking**.

---

## Integration Phase (seven steps)

> **This section is a workflow chapter**, not the protocol definition. Step count may change; invariants ①–③ may not.

When introducing or changing a Runtime Contract, the **current** validated sequence:

**Conceptual:**

```text
Architecture Decision → Runtime Contract → Navigation Alignment
  → Regression → Drift Lock → Debt Payoff → Close ADR
```

**Implementation order** (ADR-013 validated):

```text
Runtime consistency → Navigation + Drift Lock → Debt Payoff → Regression (may parallel runtime) → Close ADR
```

### Full lifecycle map (Stages A–F)

| Stage | Purpose |
|---|---|
| **A — Problem & evidence** | Symptom vs architecture; generalization evidence |
| **B — Taxonomy / naming** | Vocabulary convergence before Accept |
| **C — Architecture Decision** | ADR **Accepted** |
| **D — Contract establishment** | Owner + registry + enforcement + consumer + dogfood |
| **E — Integration** | Seven steps above → Mechanical Closure |
| **F — Post-close** | New ADR/plan only — no closed-ADR scope creep |

### Step contracts (current version)

| Step | Exit criteria | Lock |
|---|---|---|
| **1. Architecture Decision** | ADR **Accepted** | — |
| **2. Runtime Contract** | `runtime validate` understands fields | Contract lock |
| **3. Navigation Alignment** | Thin flow; fat README; consumer in taxonomy | Drift lock when canon exists |
| **4. Regression** | Scenarios + Go tests | — |
| **5. Drift Lock** | Canon scan active | Navigation drift lock |
| **6. Debt Payoff** | Classes A–D cleared | Ownership drift lock |
| **7. Close ADR** | Mechanical Closure green; plan archived | All locks |

---

## When to use

| Use this protocol | Do not use |
|---|---|
| New or changed **runtime-readable** contract | Doc fix with no contract semantics |
| ADR-accepted boundary → routing, navigation, validators | README-only polish |
| Cross-layer drift risk | Lock-only work → [`governance-pattern-template.md`](governance-pattern-template.md) |

---

## Lineage

ADR-013 and this protocol share origin but **different lifecycles**:

```text
Review three-layer mismatch (2026-07)
  → ADR-013 Completed (architecture decision — done)
  → Pattern abstracted
  → architecture-evolution-protocol.md (methodology — beginning)
  → ADR-014 Instance 2 candidate (protocol validation — pending)
```

| Artifact | Role |
|---|---|
| [ADR-013 plan archive](../../plans/archived/2026-07-06-review-architecture-adr/_plan.md) | Instance 1 evidence |
| [ADR-013](../../constitution/ADR-013-cognitive-role-primitive-gate.md) | Decision + §Dogfood Evidence |
| [ADR-014](../../constitution/ADR-014-cognitive-stance-capability-context.md) | Instance 2 candidate |

---

## Instance 1 — reference (ADR-013 Dogfood)

Full record in [ADR-013 §Dogfood Evidence](../../constitution/ADR-013-cognitive-role-primitive-gate.md). Summary:

| Field | Record |
|---|---|
| **Runtime tests** | `capability_context_test.go` · `capability_registry` |
| **Regression** | `capability-stance-*-v1.yaml` (4 scenarios) |
| **Drift locks** | `review_architecture_doc_drift` · `canonical_ownership_drift` |
| **Commits** | `948f2f1` · `5c732fa` · `fb9cdc1` |

---

## Plan checklist (agent-facing)

- [ ] Three output layers understood: ADR ≠ Protocol ≠ Governance Pattern
- [ ] Invariant ① Separation: no layer redefines below
- [ ] Invariant ② Mechanical Closure: Regression + Drift Lock + Debt Payoff before ADR close
- [ ] Invariant ③ Generalization: instance completes without protocol core edit
- [ ] ADR §Dogfood Evidence filled
- [ ] §Validated Instances + §Protocol Evidence updated on close
- [ ] Post-close scope → new ADR/plan only

---

## Related

- [`decision-promotion-pipeline.md`](decision-promotion-pipeline.md) — where decisions land
- [`governance-pattern-template.md`](governance-pattern-template.md) — how locks are built
- [`system-upgrade-governance.md`](system-upgrade-governance.md) — plan archive checklist
- [`../../architecture/ai-native-cognitive-execution-system.md`](../../architecture/ai-native-cognitive-execution-system.md)
- [`../../constitution/ADR-013-cognitive-role-primitive-gate.md`](../../constitution/ADR-013-cognitive-role-primitive-gate.md) — Instance 1 · **Completed**
- [`../../constitution/ADR-014-cognitive-stance-capability-context.md`](../../constitution/ADR-014-cognitive-stance-capability-context.md) — Instance 2 candidate
