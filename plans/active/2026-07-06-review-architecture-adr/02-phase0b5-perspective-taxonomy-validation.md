---
id: 2026-07-06-review-architecture-adr-phase0b5
parent: 2026-07-06-review-architecture-adr
phase: 0b5
status: complete
created: 2026-07-06
purpose: >
  Validate Perspective taxonomy before D2 final acceptance. Test whether Review,
  Debugger, Incident Analysis, and Security Audit share a higher-order perspective
  (e.g. fault_finding / negative_evidence_seeking) rather than separate role labels.
stakeholder_position: >
  Reject D1 (accepted). D2 as working hypothesis (accepted). D2 final acceptance
  deferred until perspective taxonomy is validated.
---

# Phase 0b.5 / 0d — Cognitive Stance Validation

> **Historical note:** Phase 0b.5 used `perspective` as a working label. **Phase 0d final:** field name is **`stance`**; working label deprecated. See [`ADR-014`](../../../constitution/ADR-014-cognitive-stance-capability-context.md).

## Phase 0d stakeholder decision (final)

| Item | Decision |
|---|---|
| Reject D1 | **Accepted** |
| Accept D2 | **Accepted** |
| Reject `perspective: reviewer` | **Accepted** |
| Accept `stance: fault_finding` | **Accepted** |
| Reject `constructive_build` enum | **Accepted** — use `default`; evidence insufficient |
| Field name `stance` (not `perspective`) | **Accepted** — epistemic stance, not actor viewpoint |
| ADR-014 separate taxonomy | **Strongly recommended — accepted** |
| ADR-013 final acceptance | **Accepted** |

### Why `stance`, not `perspective`

| Term | Fits? | Reason |
|---|---|---|
| `perspective` | **No** | Implies Developer / Customer / Operator **viewpoint** |
| `stance` | **Yes** | **Epistemic / reasoning stance** — e.g. falsify vs default forward |
| `reviewer` | **No** | **Actor** label — recreates Role at context layer |

Scientific-method mapping: **try to falsify** = `fault_finding`.

### Why not `constructive_build`

Refactoring, documentation, and implementation share only **「不是 fault_finding」** — not a proven shared reasoning family. Phase 0d: use **`default`** until a second stance family is evidenced via ADR-014.

---

## Stakeholder decision (Phase 0c input — superseded by 0d above)

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

## Validation matrix (Phase 0b.5) — filled

| Activity | Workflow slice / path | Capability | Proposed perspective | Same as Review? | Evidence |
|---|---|---|---|---|---|
| Code review | `sd-implementation` | `code-review` | `fault_finding` | **Yes** | ADR-013 Q2：停止 feature coding；findings-only；review report artifact |
| Architecture review | `architecture/` | `architecture-review` | `fault_finding` | **Yes** | Critique fit / risk，非 build；與 code review 同 stance、不同 capability |
| Debug / root cause | FORENSIC / RECOVERY path | `debug-trace` | `fault_finding` | **Yes** | RECOVERY：`write_without_root_cause` forbidden；FORENSIC：lineage before action: [`cognitive-modes-phase-integration.yaml`](../../../runtime/cognitive-modes-phase-integration.yaml) |
| Security audit | contracts / impl | `security-audit` | `fault_finding` | **Yes** | Threat / vuln seeking；review-checklist 含 security 聚焦；adversarial to design |
| UI incident observe | `sd-incident-observation` | incident card | `fault_finding` | **Partial → Yes** | 共享「證據優先、禁止 implementation-first / root-cause guess」；artifact 是 incident card 非 findings list: [`incident-observation.md`](../../../workflow/software-delivery/incident-observation.md) |
| Refactoring | `sd-implementation` | prep refactor | `constructive_build` | **No** | `execution_mode: preparatory_refactoring` 已足夠 — Phase 0b 反例 |
| Documentation | closure / linked-updates | `doc-sync` | `constructive_build` | **No** | `objective: author` 已足夠 — Phase 0b 反例 |

**Matrix summary:** 4/4 fault-seeking activities collapse to **`fault_finding`**; incident observe is **Partial** only at artifact shape, not cognitive stance.

---

## Orthogonality: `stance` vs `cognitive_mode` (ADR-008)

| Layer | Question | Example |
|---|---|---|
| **`cognitive_mode`** | **How** to execute? | FORENSIC → read full lineage, block writes until analysis complete |
| **`context.stance`** | **What reasoning stance**? | `fault_finding` → seek disconfirming evidence / flaws / causes |
| **Capability** | **Which procedure**? | `code-review` vs `debug-trace` vs `security-audit` |

**Conclusion:** `FORENSIC + fault_finding` is **composable, not redundant**. Mode governs depth and gates; stance governs epistemic intent.

Example stack:

```yaml
invoke:
  capability: debug-trace
  context:
    stance: fault_finding
    caller_slice: sd-implementation
cognitive_mode:
  execution_mode: RECOVERY
  context_mode: CHECKLIST_FIRST
  governance_mode: STRICT
  memory_mode: FAILURE_REPLAY
```

---

## Stance taxonomy (Phase 0d — conservative)

### Standardized (ADR-013 + ADR-014 draft)

| Value | Meaning |
|---|---|
| `fault_finding` | Suspend forward work; seek counter-evidence, flaws, causes, vulnerabilities |
| `default` | Forward path (explicit or omitted); **not** a proven shared family |

**Anti-explosion rule:** Consumer labels (`reviewer`, `debugger`, …) live in **capability id / docs**, not in `context.stance`.

### Rejected / deferred

| Candidate | Status |
|---|---|
| `constructive_build` | **Rejected** — premature "everything else" bucket |
| `perspective: reviewer` | **Rejected** — actor label |
| Future second family | **Deferred** — ADR-014 evidence gate |

---

## Phase 0b.5 preliminary recommendation → Phase 0d final

| Item | Phase 0b.5 | Phase 0d final |
|---|---|---|
| Reject actor enum | Leaning yes | **Accepted** |
| `fault_finding` | Leaning yes | **Accepted** as `stance` value |
| `constructive_build` | Leaning yes | **Rejected** → `default` |
| Field name | `perspective` (working) | **`stance`** |
| ADR-014 | Leaning yes | **Proposed** |
| ADR-013 | Working model | **Accepted** |

### Canonical D2 invoke envelope

```yaml
invoke:
  capability: code-review
  context:
    stance: fault_finding
    caller_slice: sd-implementation
    objective: optional
```

---

## Phase 0b.5 success criteria

- [x] Matrix filled with **Yes / No / Partial** for "same stance as Review"
- [x] **Collapse** to `fault_finding` + `default` — Phase 0d final
- [x] D2 invoke envelope updated in ADR-013 — **Accepted**
- [x] ADR-014 Proposed for stance taxonomy
- [x] Reject `constructive_build` — Phase 0d

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

Phase 1 **unblocked** — ADR-013 Accepted. Implement cross-cutting review capabilities + `context.stance: fault_finding` hooks. Stance enum growth remains **ADR-014** gate.
