# ADR-014: Cognitive Stance in Capability Context

## Status

**Proposed**

## Framework Generation

- **世代分類**：Gen 3 — Cognitive Execution System 子系統邊界擴充（capability context layer，非 runtime primitive）
- **當前世代文件**：[`architecture/ai-native-cognitive-execution-system.md`](../architecture/ai-native-cognitive-execution-system.md)；[`constitution/ADR-013-cognitive-role-primitive-gate.md`](ADR-013-cognitive-role-primitive-gate.md)（D2 accept — 引入 `context.stance` 欄位）
- **適用狀態**：本 ADR **不**重開 ADR-013 的 D1/D2 問題。Accept 前僅定義 stance **合法值**、命名邊界、與 ADR-008 `cognitive_mode` 的正交關係。

## Date

2026-07-06

## Source Plan

- [`plans/active/2026-07-06-review-architecture-adr/_plan.md`](../plans/active/2026-07-06-review-architecture-adr/_plan.md)
- [`plans/active/2026-07-06-review-architecture-adr/02-phase0b5-perspective-taxonomy-validation.md`](../plans/active/2026-07-06-review-architecture-adr/02-phase0b5-perspective-taxonomy-validation.md) — Phase 0b.5 / 0d evidence

## Context

### Why ADR-014 exists (separate from ADR-013)

| ADR | Primary question |
|---|---|
| **ADR-013** | Does Cognitive System need **`cognitive_role` runtime primitive**? → **No (D2 accepted)** |
| **ADR-014** | What **bounded values** may `context.stance` carry? |

ADR-013 **Accepted** establishes that capability invoke context **may** carry a cognitive stance field. It **does not** finalize the full stance taxonomy.

### Naming: `stance`, not `perspective` or `role`

Phase 0b.5 used `perspective` as a working label. Phase 0d stakeholder review rejected it:

| Term | Problem |
|---|---|
| `reviewer`, `debugger` | **Actor** labels — recreates Role at context layer |
| `perspective` | Implies viewpoint (Developer / Customer / Operator) — wrong abstraction |
| **`stance`** | **Epistemic / reasoning stance** — e.g. seek counter-evidence vs default forward work |

Alternatives considered: `reasoning_stance`, `cognitive_stance`. **Canonical field name: `stance`** (short, bounded; full term in glossary: *cognitive stance*).

Scientific-method analogy:

```text
Hypothesis → try to prove (default forward path)
          → try to falsify (fault_finding)
```

`fault_finding` is **reasoning direction**, not a persona.

### Orthogonality with ADR-008

| Layer | Question | Example |
|---|---|---|
| `cognitive_mode` (ADR-008) | **How** to execute? | FORENSIC — lineage before write |
| `context.stance` (ADR-014) | **What reasoning stance**? | `fault_finding` — seek disconfirming evidence |
| `capability` | **Which procedure**? | `code-review` vs `debug-trace` |

Composable, not redundant.

---

## Primary Question

> **What are the valid, bounded values for `context.stance` on capability invoke?**

---

## Phase 0d stakeholder decisions (input)

| Item | Decision |
|---|---|
| Reject `perspective: reviewer` | **Accepted** |
| `fault_finding` as stance value | **Accepted** |
| `constructive_build` as enum | **Rejected** — evidence insufficient; premature "everything else" bucket |
| ADR-014 separate from ADR-013 | **Strongly recommended — accepted** |

---

## Standardized values (Phase 0d — conservative)

### Accepted now (ADR-013 + this ADR draft)

| Value | Meaning |
|---|---|
| `fault_finding` | Suspend forward constructive work; prioritize counter-evidence, flaws, root causes, vulnerabilities |
| `default` | Implicit forward path when omitted or explicitly set; **not** a proven shared stance family |

```yaml
invoke:
  capability: code-review
  context:
    stance: fault_finding
    caller_slice: sd-implementation
```

**Anti-explosion rule:** Consumer labels (`reviewer`, `debugger`, `auditor`, `investigator`) live in **capability id / documentation**, not in `context.stance`.

### Explicitly not standardized (deferred)

| Candidate | Why deferred |
|---|---|
| `constructive_build` | Refactoring, documentation, implementation share only "not fault_finding" — insufficient evidence of one stance |
| `incident_investigation` | Partial overlap with `fault_finding`; timeline/authority owned by `sd-incident-observation` slice + FORENSIC mode |
| Actor-derived labels | Would recreate D1 Role at context layer |

New stance values require **evidence of ≥2 activities sharing irreducible reasoning semantics** + ADR/plan gate (same bar as ADR-013 primitive promotion, applied to enum growth).

---

## Activities mapped to `fault_finding` (evidence)

| Activity | Capability | Stance | Artifact |
|---|---|---|---|
| Code / architecture review | `code-review`, `architecture-review` | `fault_finding` | review report |
| Debug / root cause | `debug-trace` | `fault_finding` | root-cause trace |
| Security audit | `security-audit` | `fault_finding` | finding list |
| UI incident observe | incident card workflow | `fault_finding` | incident card |

Counter-examples (use `default` + slice / `objective` / `execution_mode`): refactoring, documentation, planning, implementation.

Full matrix: [`02-phase0b5-perspective-taxonomy-validation.md`](../plans/active/2026-07-06-review-architecture-adr/02-phase0b5-perspective-taxonomy-validation.md).

---

## Open Questions (ADR-014 accept gate)

1. Glossary term: `cognitive_stance` vs `stance` in prose — field name fixed as `stance`?
2. Peer vs self review — same capability + context flag, or separate capability id?
3. Mechanical enforcement of `stance: fault_finding` — advisory vs transition-block (Phase 1)
4. Future second stance family — promotion criteria + evidence bar
5. D2 vs B — is formal `context.stance` envelope required for all review capabilities, or optional?

---

## Recommendation (draft)

> **Accept conservative enum:** `context.stance ∈ { fault_finding, default }`.
>
> **Do not** standardize `constructive_build` until a distinct shared reasoning family is evidenced.
>
> **Governance:** new stance values via ADR amendment or ADR-014 revision — not ad hoc in capability docs.

---

## Decision

**TBD** — Proposed until Phase 1 governance doc + stakeholder sign-off.

---

## Non-goals (enum discipline)

- **Do not** reserve placeholder stance values (`creative`, `planning`, `optimization`, …) to "fill the field"
- New stance values only via **ADR-014 amendment** + evidence of a second mature family
- `stance` is **not** review-directory-local — capabilities declare `requires_context.stance`

---

## Capability contract pattern (Phase 1)

Stance is declared on **any** capability that needs it — not only under `cross-cutting/review/`:

```yaml
id: security-audit
requires_context:
  stance:
    - fault_finding
invoke:
  capability: security-audit
  context:
    stance: fault_finding
    caller_slice: sd-contracts
```

Review capabilities (`code-review`, `architecture-review`, …) use the **same** contract shape as `debug-trace`, `security-audit`, `incident-analysis`.

---

## Consequences (if accepted)

- Governance doc (Phase 1) documents **`requires_context.stance`** on capability metadata — not review-README-local
- [`governance/cognitive-stance.md`](../governance/cognitive-stance.md) — contract owner (Phase 1.1)
- [`knowledge/runtime/capability-registry.yaml`](../knowledge/runtime/capability-registry.yaml) — capability metadata
- Glossary entry: `cognitive_stance` (distinct from `cognitive_mode`, `cognitive_role`)
- Phase 0b.5 working label `perspective` **deprecated** in new docs — historical plans may retain for traceability

---

## Related

- [`constitution/ADR-013-cognitive-role-primitive-gate.md`](ADR-013-cognitive-role-primitive-gate.md) — D1 reject, D2 accept
- [`constitution/ADR-008-runtime-cognitive-modes.md`](ADR-008-runtime-cognitive-modes.md)
- [`workflow/cross-cutting/README.md`](../workflow/cross-cutting/README.md)
