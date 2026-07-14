# Architecture Evolution Protocol

> **Protocol Status**: **Experimental** (1 reference implementation; first independent validation pending)
> **Not** an ADR appendix. **ADR-013 is complete** (Reference Implementation #1); **this protocol has just begun**.
> **Sibling docs**: [`decision-promotion-pipeline.md`](decision-promotion-pipeline.md) · [`governance-pattern-template.md`](governance-pattern-template.md)

## Positioning — four layers

These are **not the same thing**. Each can stand alone. Stability increases down the stack.

| Layer | Role | Stability | Current (2026-07-07) |
|---|---|---|---|
| **Architecture Decision** | Solve one concrete architecture problem | Project-level | ADR-013 **Completed** · ADR-014 Proposed |
| **Architecture Evolution Protocol** | Govern *how* architecture evolves | Framework-level | **Experimental** |
| **Governance Patterns** | Reusable mechanical shapes (Three-lock, Debt classes, Second-instance gate) | Framework-level | Proven in RI #1 |
| **Reference Implementation** | One **complete** protocol walkthrough — spec + verifiable evidence | Evidence-level | **RI #1 = ADR-013** |

ADR-013 is **not only** an ADR. It is the **first Reference Implementation** — a full, verifiable walkthrough of this protocol. When discussing the protocol, ask: *「ADR-013 是怎麼走完整個 Protocol 的？」* — not *「當初 Review 是怎麼討論的？」*

**Spec and reference implementation validate each other; they must not be conflated.**

### Governance maturity loop

Healthy architecture knowledge **accumulates**; it does not merely accumulate documents:

```text
Architecture Decision
  → Reference Implementation (complete protocol walkthrough + Dogfood Evidence)
  → Protocol Validation (Protocol Evidence: no core modified?)
  → Governance Maturity (Protocol Status advances)
  → (next ADR repeats)
```

Each new ADR solves a problem **and** tests whether the protocol is stable enough to reuse.

---

## Protocol Core (normative)

**Primary maturity signal:** *No protocol core modified* — not instance count alone.

Three instances with three core edits → **not** Stable. Two instances with zero core edits → **Emerging**.

### Core (amending any item = Protocol Core Change → re-evaluate maturity)

| Core element | Location in this file |
|---|---|
| **Three invariants** | §The protocol — ① Separation · ② Mechanical Closure · ③ Generalization |
| **Evolution Contract** (no skip rule + violation shapes) | §Contract stack · §Evolution Contract |
| **Layer dependency invariant** | §① Separation of Concerns |
| **Mechanical Closure requirement** | §② Mechanical Closure |
| **Second-instance / generalization principle** | §③ Generalization · §Protocol Status |
| **Protocol Core definition** | This section |
| **Dogfood vs Protocol Evidence distinction** | §Evidence |
| **Changelog impact rules** | §Protocol Changelog |

### Non-core (editorial / evidence growth — does not reset maturity)

| Non-core element | Examples |
|---|---|
| Integration Phase step count or order | Seven steps → six or eight |
| Step contract tables | Per-step exit criteria wording |
| Plan checklist | Agent-facing bullets |
| Example paths, commits, scenario filenames | ADR-013 dogfood summary |
| **Validated Instances** rows | New instance evidence |
| **Reference Implementation** index rows | RI #2, #3 |
| Protocol Status **current** label | Experimental → Emerging |
| Lineage narrative | Historical discussion |

**Judgment rule:** If the change alters *what must always hold* → core. If it only adds *evidence that it held again* → non-core.

---

## Protocol Changelog

When editing this file, classify each change:

| Change type | Affects maturity? | Action |
|---|---|---|
| **Editorial** | ❌ No | Typo, clarity, translation — no status change |
| **Example update** | ❌ No | New paths, commits, scenario names in RI section |
| **Additional validated instance** | ❌ No | New row in §Validated Instances / §Protocol Evidence |
| **Additional reference implementation** | ❌ No | New RI index row + linked ADR §Dogfood |
| **New governance pattern** | ⚠️ Maybe | If pattern is **optional appendix** → non-core. If it amends invariants or Mechanical Closure → core |
| **Protocol Core Change** | ✅ Yes | Re-evaluate Protocol Status; document in changelog below; may revert to Experimental |

### Changelog record

| Date | Type | Summary |
|---|---|---|
| 2026-07-07 | Protocol Core Change | Initial promote: three invariants, Evolution Contract, Mechanical Closure, Protocol Status |
| 2026-07-07 | Protocol Core Change | Protocol Core definition, four-layer positioning, Reference Implementation, Changelog classification |
| 2026-07-14 | New governance pattern | Optional appendix §Layer Growth Rhythm（Grow one layer, freeze the previous）— non-core；sourced from UI Pattern Knowledge Phase 1–3 rhythm |
| 2026-07-14 | Editorial | Pair Constraint Accumulation sentence into §Layer Growth Rhythm；cite H4 stress evidence (Entry Mods=0) |

---

## Positioning — what was once mixed (summary)

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

**Current:** **Experimental** — Reference Implementation #1 (ADR-013) only.

**First independent validation:** ADR-014 — not merely "the next ADR", but the protocol's **first cross-instance test**. Pass → **Experimental → Emerging** (Instance 2 complete, **no protocol core modified**).

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

## Reference Implementations

A **Reference Implementation (RI)** is one evolution that walked the **full protocol** with Mechanical Closure and recorded Dogfood Evidence. RIs are **evidence** for the protocol — not substitutes for it.

| RI | ADR | Role | Plan / evidence |
|---|---|---|---|
| **#1** | [ADR-013](../../constitution/ADR-013-cognitive-role-primitive-gate.md) | **Source RI** — first complete walkthrough; abstracted into this protocol | [archived plan](../../plans/archived/2026-07-06-review-architecture-adr/_plan.md) |
| **#2** | [ADR-014](../../constitution/ADR-014-cognitive-stance-capability-context.md) | **First independent validation** — protocol crosses single-case threshold | *pending Accept + Integration* |

**How to use RI #1:** When executing a new evolution, map your plan phases to ADR-013's archived plan Phase Map (0a–2.4) and §Dogfood Evidence — not to Review-specific prose.

---

## Validated Instances & Protocol Evidence

| Instance | ADR | RI | Status | Protocol core modified? | Mechanical Closure |
|---|---|---|---|---|---|
| **1** | [ADR-013](../../constitution/ADR-013-cognitive-role-primitive-gate.md) | **#1** | ✅ Completed | N/A (source → abstracted) | ✅ |
| **2** | [ADR-014](../../constitution/ADR-014-cognitive-stance-capability-context.md) | **#2** (candidate) | Proposed — **first independent validation** | *pending* | *pending* |
| **3** | — | — | — | — | — |

**Instance 2 pass criterion:** Stages A–F + Integration + Mechanical Closure; **Protocol Evidence: No protocol core modified**; only non-core rows added (this table, RI index, ADR-014 §Dogfood Evidence).

**Emerging criterion:** Instance 2 pass with **No** in protocol core column → Protocol Status **Experimental → Emerging**.

**Stable criterion:** Instance 3+ same **No** column — instance count is reference; **no core edits** is proof.

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

```text
Review three-layer mismatch (2026-07)
  → ADR-013 Completed + Reference Implementation #1
  → Protocol abstracted from RI #1
  → ADR-014 = first independent protocol validation (pending)
```

| Artifact | Role |
|---|---|
| [ADR-013 plan archive](../../plans/archived/2026-07-06-review-architecture-adr/_plan.md) | RI #1 walkthrough evidence |
| [ADR-013](../../constitution/ADR-013-cognitive-role-primitive-gate.md) | Decision + RI #1 + §Dogfood Evidence |
| [ADR-014](../../constitution/ADR-014-cognitive-stance-capability-context.md) | First independent validation (not "just next ADR") |

---

## RI #1 — ADR-013 summary (Dogfood pointer)

Full record in [ADR-013 §Dogfood Evidence](../../constitution/ADR-013-cognitive-role-primitive-gate.md). Summary:

| Field | Record |
|---|---|
| **Runtime tests** | `capability_context_test.go` · `capability_registry` |
| **Regression** | `capability-stance-*-v1.yaml` (4 scenarios) |
| **Drift locks** | `review_architecture_doc_drift` · `canonical_ownership_drift` |
| **Commits** | `948f2f1` · `5c732fa` · `fb9cdc1` |

---

---

## Appendix — Layer Growth Rhythm（optional Governance Pattern）

> **Classification**: optional appendix / reusable research rhythm — **non-core**（does not amend Protocol Core invariants）.  
> **Source observation**: UI Pattern Knowledge workflow（Phase 1–3）+ prior Evidence / Governance layering practice.

### Pattern

```text
Grow one layer, freeze the previous layer.
A frozen layer may accumulate constraints without reopening its knowledge objects.
```

| Sentence | Role |
| --- | --- |
| Grow one layer, freeze the previous | **Growth** — open the next knowledge surface |
| A frozen layer may accumulate constraints… | **Constraint Accumulation** — new pressure becomes edges/rules on the new layer; frozen objects stay closed |

Pair: Growth ⊕ Constraint Accumulation. Together they prevent back-propagation while still allowing Composition Knowledge to emerge.

| Phase（instance） | New artifact layer | Freeze target（previous） |
| --- | --- | --- |
| 1 | Entry / Schema layer | Schema contract |
| 2 | Scenario / Inferability layer | Entry |
| 3 | Composition Constraints layer | Entry + Scenario |

Each phase **adds** one knowledge surface and **freezes** what the prior phase already validated. New discoveries flow into the **new** layer — not back into frozen layers（anti **back-propagation**）. H4 Episode stress (2026-07-14)：Composition Failure → new Constraint → Entry Modifications = 0（first positive evidence of Constraint Accumulation）.

### Why reusable

This is not UI-Pattern-specific. The same rhythm appears when Evidence, Governance, and Runtime contracts evolve: promote a layer only after the layer below is frozen enough that later screens/instances cannot silently rewrite it.

### When to apply

- Multi-phase research / architecture dogfood where later phases compose earlier artifacts
- Any plan that risks “fix by editing the prior phase’s validated unit”

### When not to apply

- Single-shot bugfix with no layered knowledge claim
- Explicit plan revision that **unfreezes** a prior layer（rare；must be written，not accidental）

---

## Plan checklist (agent-facing)


- [ ] Four layers understood: ADR · Protocol · Governance Pattern · Reference Implementation
- [ ] Protocol Core vs non-core: know if your edit triggers maturity re-eval
- [ ] Changelog type recorded for any protocol edit
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
- [`../../constitution/ADR-013-cognitive-role-primitive-gate.md`](../../constitution/ADR-013-cognitive-role-primitive-gate.md) — **Reference Implementation #1** · Completed
- [`../../constitution/ADR-014-cognitive-stance-capability-context.md`](../../constitution/ADR-014-cognitive-stance-capability-context.md) — first independent protocol validation
