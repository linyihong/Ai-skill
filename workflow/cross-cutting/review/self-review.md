# Self Review (Post-Implementation)

> **Consumer** for capability `code-review` from caller slice `sd-implementation`.  
> Stance contract: [`governance/cognitive-stance.md`](../../../governance/cognitive-stance.md) — **not defined here**.

## When to invoke

After implementation changes are complete and before validation / closure claims:

- Feature or bug-fix code paths touched
- AI-generated code in the diff
- Governance expects a review report artifact (see [`software-delivery-governance.md`](../../../governance/ai-runtime-governance/software-delivery-governance.md))

Do **not** use this path for forward implementation — use [`implementation/`](../software-delivery/implementation/README.md) default path.

## Invoke envelope

```yaml
invoke:
  capability: code-review
  context:
    stance: fault_finding
    caller_slice: sd-implementation
```

Registry: [`code-review`](../../../knowledge/runtime/capability-registry.yaml) — `requires_context.stance: [fault_finding]`.

Pre-flight (advisory):

```bash
ai-skill runtime capability-invoke --capability code-review --stance fault_finding
```

## Agent behavior (findings-only)

1. **Stop** feature coding — no new implementation in this pass
2. Run checklist: [`checklist.md`](checklist.md) (sections relevant to the change)
3. Produce **findings-only** output — blockers, suggestions, evidence gaps
4. Emit **review report** artifact (project template or governance-required format)

## Artifact

| Field | Required |
|---|---|
| Findings list | yes — severity + location + evidence |
| Blocker vs advisory | yes |
| Checklist sections covered | yes |

## Related

- [`invocation-points.md`](invocation-points.md) — all review types
- [`implementation/README.md`](../software-delivery/implementation/README.md) — `sd-implementation` slice
