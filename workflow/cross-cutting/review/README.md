# Review (Cross-Cutting Consumer)

Fault-finding capabilities invoked from workflow slices — **not** a `software-delivery` lifecycle slice and **not** the stance contract owner.

## Contract vs consumer

| Role | Owner | Path |
|---|---|---|
| Stance contract | Governance | [`governance/cognitive-stance.md`](../../../governance/cognitive-stance.md) |
| Capability metadata | Runtime registry | [`knowledge/runtime/capability-registry.yaml`](../../../knowledge/runtime/capability-registry.yaml) |
| Enforcement | Runtime | [`runtime/capability-context.yaml`](../../../runtime/capability-context.yaml) |
| **This directory** | **Consumer only** | Invoke + checklist bodies |

**Governance invariant:** Do not define `stance` values or `requires_context` in this README. Capabilities declare requirements; consumers invoke.

## Runtime pattern

```text
Workflow (caller slice)
  → invoke capability + context.stance
  → artifact (review report, findings, …)
```

Validate invoke (Phase 1.2 — warning only):

```bash
ai-skill runtime capability-invoke --capability code-review --stance fault_finding
```

## Consumer surfaces

| Surface | Purpose |
|---|---|
| [`README.md`](README.md) | Consumer index (this file) |
| [`invocation-points.md`](invocation-points.md) | Review type × caller slice × capability |
| [`self-review.md`](self-review.md) | Post-implementation code review consumer |
| [`checklist.md`](checklist.md) | Tool-neutral review checklist bodies |

## Registered capabilities (fault_finding)

See [`capability-registry.yaml`](../../../knowledge/runtime/capability-registry.yaml). Phase 1 dogfood trio:

| Capability | Typical caller | Artifact |
|---|---|---|
| `code-review` | `sd-implementation` | review report |
| `security-audit` | `sd-contracts`, `sd-implementation` | security finding list；付費媒體見 [`media-entitlement-control-plane.md`](../../../analysis/security/media-entitlement-control-plane.md) |
| `incident-analysis` | `sd-incident-observation` | incident card |

Other fault-finding capabilities (`architecture-review`, `debug-trace`, …) share the **same invoke contract** — no review-specific runtime logic.

## Slice hooks (thin)

Caller slices pass `context.stance: fault_finding` when invoking — they do not redefine stance:

| Caller slice | Consumer doc |
|---|---|
| `sd-implementation` | [`self-review.md`](self-review.md) |
| `sd-contracts` / `sd-implementation` | [`invocation-points.md`](invocation-points.md) §security |
| `sd-incident-observation` | [`invocation-points.md`](invocation-points.md) §incident |
| `architecture/` | [`invocation-points.md`](invocation-points.md) §architecture |
| `sd-validation` / `sd-closure` | [`invocation-points.md`](invocation-points.md) §release |

## Related

- [`ADR-013`](../../../constitution/ADR-013-cognitive-role-primitive-gate.md) — D2 accepted
- [`ADR-014`](../../../constitution/ADR-014-cognitive-stance-capability-context.md) — stance taxonomy
- [`../README.md`](../README.md) — cross-cutting policy
