# Capability Context Schema

> **Registry version**: `capability-registry/v1`  
> **Canonical contract**: [`governance/cognitive-stance.md`](../governance/cognitive-stance.md)  
> **Instance registry**: [`knowledge/runtime/capability-registry.yaml`](../knowledge/runtime/capability-registry.yaml)

Schema for capability metadata — especially `requires_context` declarations consumed by workflow invoke and runtime validation (Phase 1.2+).

## Registry root fields

| Field | Type | Required | Purpose |
|---|---|---|---|
| `registry_version` | string | yes | e.g. `capability-registry/v1` |
| `status` | enum | yes | `draft` \| `candidate` \| `active` |
| `owner_layer` | string | yes | `knowledge/runtime` |
| `stance_enum` | object | yes | Policy for legal `context.stance` values |
| `capabilities` | array | yes | Capability metadata records |

## `stance_enum` object

| Field | Type | Purpose |
|---|---|---|
| `standardized` | string[] | Non-default values approved for use (Phase 1: `[fault_finding]` only) |
| `implicit` | string[] | Values that need not appear in `requires_context` (e.g. `default`) |
| `reserved_policy` | string | Human-readable: no placeholder enums without ADR-014 |

## Capability record

| Field | Type | Required | Purpose |
|---|---|---|---|
| `id` | string | yes | Stable kebab-case capability id (matches invoke `capability`) |
| `status` | enum | yes | `draft` \| `active` \| `deprecated` |
| `summary` | string | yes | One-line procedure description |
| `requires_context` | object | no | Context requirements for invoke |
| `requires_context.stance` | string[] | when fault-seeking | Subset of `stance_enum.standardized` |
| `artifact` | string | no | Primary output artifact type |
| `typical_caller_slices` | string[] | no | Workflow hints — not contract |

## `requires_context` rules

1. **Declare** `stance: [fault_finding]` when the capability **requires** counter-evidence seeking (ADR-013 D2).
2. **Omit** `requires_context` when `default` (forward) stance suffices.
3. **Do not** list `default` in `requires_context.stance` — omission means default allowed.
4. Values **must** be subset of `stance_enum.standardized` + `implicit`.

## Invoke validation (Phase 1.2+)

```yaml
# Registry
id: code-review
requires_context:
  stance:
    - fault_finding

# Valid invoke
invoke:
  capability: code-review
  context:
    stance: fault_finding
    caller_slice: sd-implementation

# Phase 1.2: missing stance → Warning (not auto-fill, not hard block)
```

## Anti-patterns

| Anti-pattern | Why |
|---|---|
| Consumer README defines `stance=` | Violates governance invariant — capability registry owns contract |
| Placeholder stance in `stance_enum.standardized` | Premature taxonomy — ADR-014 gate |
| Actor labels in `stance` (`reviewer`, `debugger`) | Recreates rejected `cognitive_role` at context layer |

## Related

- [`governance/cognitive-stance.md`](../governance/cognitive-stance.md)
- [`metadata/schema.md`](schema.md) — knowledge atom schema (different concern)
