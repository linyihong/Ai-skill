# Review Invocation Points

> Capability ids and `context.stance` come from [`knowledge/runtime/capability-registry.yaml`](../../../knowledge/runtime/capability-registry.yaml). This file maps **when** callers invoke — not stance taxonomy.

## Review type × caller × invoke

| Review type | Caller slice / path | Capability | Invoke context |
|---|---|---|---|
| Contract review | `sd-contracts` | `contract-review`¹ | `stance: fault_finding` |
| Architecture review | `architecture/` | `architecture-review` | `stance: fault_finding` |
| Security audit | `sd-contracts`, `sd-implementation` | `security-audit` | `stance: fault_finding` |
| Code review | `sd-implementation` | `code-review` | `stance: fault_finding` |
| Performance review | `sd-test-strategy`, perf-risk-gate | `performance-review`¹ | `stance: fault_finding` |
| UI / compliance review | `sd-ui-governance` | `ui-review`¹ | `stance: fault_finding` |
| Release review | `sd-validation`, `sd-closure` | `release-review`¹ | `stance: fault_finding` |
| Incident analysis | `sd-incident-observation` | `incident-analysis` | `stance: fault_finding` |
| Debug / root cause | `sd-implementation` | `debug-trace` | `stance: fault_finding` |

¹ Capability registered as **draft** in registry until consumer surface ships — invoke shape is the same.

## Example invokes

```yaml
# Security audit (contracts path)
invoke:
  capability: security-audit
  context:
    stance: fault_finding
    caller_slice: sd-contracts

# Incident analysis
invoke:
  capability: incident-analysis
  context:
    stance: fault_finding
    caller_slice: sd-incident-observation
```

## Consumer docs

| Review type | Consumer surface |
|---|---|
| Code (self-review) | [`self-review.md`](self-review.md) |
| Checklist bodies | [`checklist.md`](checklist.md) |
| Index | [`README.md`](README.md) |

## Related

- [`ADR-013`](../../../constitution/ADR-013-cognitive-role-primitive-gate.md) §Review Type × Invocation Point
