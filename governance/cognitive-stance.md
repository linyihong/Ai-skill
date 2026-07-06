# Cognitive Stance — Capability Context Contract

> **Status**: `active` (Phase 1.1)  
> **ADR**: [`ADR-013`](../constitution/ADR-013-cognitive-role-primitive-gate.md) (Accepted) · [`ADR-014`](../constitution/ADR-014-cognitive-stance-capability-context.md) (Proposed — value taxonomy)  
> **Machine registry**: [`knowledge/runtime/capability-registry.yaml`](../knowledge/runtime/capability-registry.yaml)  
> **Schema**: [`metadata/capability-context-schema.md`](../metadata/capability-context-schema.md)

## Purpose

Define how **capabilities** declare required **runtime context** — specifically `context.stance` — without promoting `cognitive_role` to a runtime primitive.

This document is the **contract owner** for stance. Workflow consumers (e.g. `cross-cutting/review/`) are **not** contract owners.

## Five orthogonal dimensions

| Dimension | Question | Layer |
|---|---|---|
| **Workflow** | When? | Caller slice / phase |
| **Capability** | What procedure? | `capability` id |
| **Cognitive Mode** | How (depth, gates)? | ADR-008 `cognitive_mode` |
| **Stance** | What reasoning stance? | `context.stance` |
| **Artifact** | What output? | review report, trace, incident card, … |

Stance is **epistemic / reasoning direction** (e.g. seek counter-evidence), not an actor label (`reviewer`) or viewpoint (`Developer` / `Customer`).

## Runtime pattern

```text
Capability
  → declares requires_context (metadata)
Workflow
  → invoke capability + context
Consumer
  → loads capability body (never redefines stance contract)
Runtime
  → validates invoke against requires_context (Phase 1.2+)
```

## Invoke envelope

```yaml
invoke:
  capability: code-review
  context:
    stance: fault_finding
    caller_slice: sd-implementation
```

Omitted `stance` or `stance: default` = forward / constructive path (implicit). Capabilities that **require** `fault_finding` must declare it in registry metadata.

## Capability metadata (registry)

Each capability that needs non-default stance **must** declare:

```yaml
id: code-review
requires_context:
  stance:
    - fault_finding
```

Canonical registry: [`knowledge/runtime/capability-registry.yaml`](../knowledge/runtime/capability-registry.yaml).

## Stance enum (Phase 1 — conservative)

| Value | Status |
|---|---|
| `fault_finding` | **Standardized** — only non-default value in Phase 1 |
| `default` | Implicit when omitted; do not require in `requires_context` |

**Do not** reserve placeholder values (`creative`, `planning`, `optimization`, `constructive_build`, …). New values only via **ADR-014** amendment + evidence.

## Governance invariant

> **Capabilities declare context requirements. Consumers never own runtime context contracts.**

| Role | May define `stance` contract? | Example |
|---|---|---|
| **Governance** (`cognitive-stance.md`) | Yes — field semantics, enum policy | This file |
| **Capability registry** | Yes — per-capability `requires_context` | `capability-registry.yaml` |
| **Workflow consumer** | **No** — invoke only | `cross-cutting/review/README.md` |
| **Slice hook** | **No** — references capability + passes context | `review_invocation` in slice |

**Forbidden:**

```text
workflow/cross-cutting/review/README.md
  stance: fault_finding   # ❌ consumer must not privately define contract
```

**Required:**

```text
knowledge/runtime/capability-registry.yaml
  code-review.requires_context.stance: [fault_finding]
workflow/cross-cutting/review/
  → invoke code-review with context.stance: fault_finding
```

## Phase 1 sub-phases

| Phase | Scope | Status |
|---|---|---|
| **1.1** | Contract + registry + schema — **no** review directory | **active** |
| **1.2** | Runtime enforcement — **Warning** on missing/mismatched stance (not hard block) | planned |
| **1.3** | `workflow/cross-cutting/review/` — first consumer only | blocked until 1.2 |
| **1.4** | Dogfood: `code-review`, `security-audit`, `incident-analysis` share contract without review-specific runtime logic | blocked until 1.3 |

### Phase 1.2 enforcement (planned)

When `requires_context.stance` includes `fault_finding` but invoke omits or mismatches `context.stance`:

| Behavior | Phase 1 |
|---|---|
| Auto-fill `fault_finding` | **No** |
| Warning | **Yes** |
| Hard block | **No** (defer to later phase) |
| Fallback `default` | **No** when capability requires `fault_finding` |

## Done definition (Phase 1 complete)

1. Runtime understands `requires_context.stance` (registry + validation path).
2. `fault_finding` is the only standardized non-default stance.
3. At least **three** capabilities (`code-review`, `security-audit`, `incident-analysis`) share the same contract **without** review-specific runtime logic.
4. `workflow/cross-cutting/review/` is **first consumer only** — does not own or redefine stance.

## Related

- [`constitution/ADR-013-cognitive-role-primitive-gate.md`](../constitution/ADR-013-cognitive-role-primitive-gate.md)
- [`constitution/ADR-014-cognitive-stance-capability-context.md`](../constitution/ADR-014-cognitive-stance-capability-context.md)
- [`plans/active/2026-07-06-review-architecture-adr/_plan.md`](../plans/active/2026-07-06-review-architecture-adr/_plan.md)
- [`constitution/ADR-008-runtime-cognitive-modes.md`](../constitution/ADR-008-runtime-cognitive-modes.md) — orthogonal `cognitive_mode`
