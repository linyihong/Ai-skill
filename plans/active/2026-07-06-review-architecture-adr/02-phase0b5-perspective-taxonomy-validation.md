---
id: 2026-07-06-review-architecture-adr-phase0b5
parent: 2026-07-06-review-architecture-adr
phase: 0b5
status: in-progress
created: 2026-07-06
purpose: >
  Validate Perspective taxonomy before D2 final acceptance. Test whether Review,
  Debugger, Incident Analysis, and Security Audit share a higher-order perspective
  (e.g. fault_finding / negative_evidence_seeking) rather than separate role labels.
stakeholder_position: >
  Reject D1 (accepted). D2 as working hypothesis (accepted). D2 final acceptance
  deferred until perspective taxonomy is validated.
---

# Phase 0b.5 — Perspective Taxonomy Validation

## Stakeholder decision (Phase 0c input)

| Decision | Status |
|---|---|
| **Reject D1** (`cognitive_role` runtime primitive) | **Accepted** — Phase 0b evidence sufficient |
| **D2 as working model** | **Accepted provisionally** — implement design against D2 envelope |
| **D2 final acceptance** | **Deferred** — pending this phase |
| **Accept `perspective: reviewer` as final enum** | **Deferred** — may collapse to higher-order perspective |

**Not the same as "Accept D2".** Working hypothesis allows Phase 1 planning; final ADR acceptance waits on perspective validation.

---

## Why Phase 0b.5 exists

Phase 0b rejected D1 because Role did not generalize as a **runtime primitive**.

Phase 0b also flagged Debugger as a **weak candidate #2** — stakeholder feedback: **that conclusion was premature**.

Debugger may share a **deeper pattern** with Review:

| Activity | Surface goal | Shared cognitive move? |
|---|---|---|
| Review | Find flaws | Negative evidence seeking |
| Debug | Find root cause | Negative evidence seeking |
| Security audit | Find vulnerabilities | Negative evidence seeking |
| Incident analysis | Find evidence / broken authority | Negative evidence seeking |

**Hypothesis:** `reviewer`, `debugger`, `auditor`, `investigator` may be **consumers of one perspective**, not four separate context values.

---

## Research question

> Are Review, Debugger, Incident Analysis, and Security Audit the same **Perspective**, expressed through different capabilities?

If **Yes** → D2 envelope should use a **bounded perspective taxonomy** (e.g. `fault_finding`), not proliferating role-like labels.

If **No** → document why each perspective is **irreducible** and bound the enum explicitly.

---

## Candidate perspective taxonomy (working)

### Option P1 — Activity labels (Phase 0b draft — **under review**)

```yaml
context:
  perspective: reviewer | debugger | author | default
```

**Risk:** recreates role taxonomy at context layer (`reviewer`, `debugger`, …).

### Option P2 — Higher-order perspective (stakeholder hypothesis)

```yaml
context:
  perspective: fault_finding   # or negative_evidence_seeking
  capability: code-review      # consumer-specific procedure
```

| Capability consumer | Same perspective? | Different artifact? |
|---|---|---|
| `code-review` | `fault_finding` | review report |
| `debug-trace` | `fault_finding` | root-cause trace |
| `security-audit` | `fault_finding` | finding list |
| `incident-analysis` | `fault_finding` | incident card / evidence chain |

**Implementer / Planner / Author** perspectives (if ever needed) would be **different family** — e.g. `constructive_build` vs `fault_finding`.

### Option P3 — Split fault_finding vs incident_investigation

If incident response needs **timeline + authority trace** beyond flaw-finding:

```yaml
perspective: fault_finding | incident_investigation | constructive_build | default
```

Phase 0b.5 must **falsify or confirm** whether one or two values suffice.

---

## Validation matrix (Phase 0b.5)

| Activity | Workflow slice / path | Capability | Proposed perspective | Same as Review? | Evidence needed |
|---|---|---|---|---|---|
| Code review | `sd-implementation` | `code-review` | `fault_finding`? | TBD | Agent must stop coding; findings-only |
| Architecture review | `architecture/` | `architecture-review` | `fault_finding`? | TBD | Critique fit, not build |
| Debug / root cause | FORENSIC mode path | `debug-trace` | `fault_finding`? | **Likely yes** | Suspend forward fix; hypothesis elimination |
| Security audit | contracts / impl | `security-audit` | `fault_finding`? | TBD | Threat / vuln seeking |
| UI incident observe | `sd-incident-observation` | incident card | `incident_investigation`? | Partial | Evidence before classify |
| Refactoring | `sd-implementation` | prep refactor | `constructive_build` | **No** | Counter-example — objective not perspective |
| Documentation | closure / linked-updates | `doc-sync` | `constructive_build` | **No** | Counter-example |

---

## Debugger deep-dive (why not "weak candidate" yet)

| Dimension | Coding (default) | Debugging |
|---|---|---|
| Goal | Complete requirement | **Invalidate hypotheses** |
| Stance toward code | Extend / fix forward | **Suspect invariants** |
| Success | Tests pass, feature ships | **Root cause identified** |
| Artifact | Code, tests | Trace, failure classification |
| Overlap with Review | — | Both **adversarial to current belief** |

**Overlap with ADR-008:** `FORENSIC` / `RECOVERY` modes encode **how** (depth, governance). Perspective encodes **why this task feels different** — orthogonal if bounded.

**Open:** Is `FORENSIC mode + fault_finding perspective` redundant, or composable?

---

## Phase 0b.5 success criteria

- [ ] Matrix filled with **Yes / No / Partial** for "same perspective as Review"
- [ ] Either **collapse** to `fault_finding` (+ optional second family) **or** document **irreducible** split with anti-explosion rule
- [ ] D2 invoke envelope updated in ADR-013 (still Proposed)
- [ ] Recommendation text: **D2 working model** vs **D2 accepted**
- [ ] Decision:是否需要 **ADR-014**（Perspective Taxonomy）或 ADR-013 §amendment

**Explicit non-goal:** Re-open D1 unless perspective validation **requires** runtime primitive (unlikely).

---

## Provisional outcome paths

| Outcome | Next step |
|---|---|
| Single `fault_finding` perspective | Amend ADR-013 D2 envelope; **Accept D2** |
| `fault_finding` + `constructive_build` (+ maybe `incident_investigation`) | Bounded enum in `governance/review-capability.md`; **Accept D2** |
| Cannot collapse; many irreducible perspectives | Cap enum; **stay D2 working model**; slow enum growth via plan gate |
| Perspective duplicates slice/mode | **Reject perspective layer**; fall back to B (capability only) |

---

## Relationship to Phase 1

Phase 1 **may begin planning** against D2 working model (cross-cutting review capabilities + invoke hooks) **without** final `perspective` enum.

Implementation must treat perspective value as **configurable / draft** until Phase 0b.5 + 0d close.
