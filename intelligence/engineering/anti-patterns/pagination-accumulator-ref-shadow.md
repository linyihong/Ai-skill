# Pagination Accumulator Ref Shadow

**Status**: `candidate-intelligence`

**Source**: Vidoe-Test home-ranking scroll load-more incident (2026-07-02). First independent hit — promote after second pagination append failure with same authority sequence.

## Definition

Pagination accumulator ref shadow: a **list accumulation** uses a React ref (or external cache) as the merge source of truth, but the ref is updated **after render** (e.g. `useEffect`) while the next page fetch can run **before** that sync — so each merge reads a **stale empty or partial list** and overwrites instead of accumulating.

## Pattern

```text
page N fetch succeeds
↓
setState(merge(ref.current, incoming))   ← ref still stale
↓
useEffect syncs ref ← too late
↓
page N+1 fetch starts
↓
merge(stale ref, incoming) → overwrite / duplicate / silent zero-append
↓
API looks healthy; UI list unchanged or wrong tail
```

## Common variants

| Variant | Signal | Real failure |
| --- | --- | --- |
| `ref_shadow_empty` | Multiple cursor requests; only last page visible | ref never saw page N state |
| `dedup_scope_mismatch` | API returns new ids; `appended: 0` | dedup set includes hidden SSR tail or session batch |
| `cursor_stale_pair` | Same bottom items after fetch | session videos without matching `nextCursor` |
| `api_success_proxy` | Network 200 + JSON array | no check on list length / poster count |

## Rule

- **Merge source of truth** = functional updater `setItems(prev => merge(prev, incoming))`, not `ref.current` read before async gap.
- If a ref exists for in-flight guards, sync it **inside** the setState updater when merge succeeds.
- **Dedup base** must equal **visible ids only**, not full server prefetch batch.
- Persist **cursor + accumulated list** as a pair; never restore list without cursor.

## Observe checklist (pagination incidents)

Before blaming API or scroll trigger:

1. Did poster / row **count** increase?
2. Does debug log show `appended > 0` or explicit `duplicate-page`?
3. Is merge tested independently of scroll?
4. Is dedup base aligned with rendered slice?

## Validation

Minimum evidence chain:

```text
trigger (≤1 request per bottom reach)
→ merge unit test (sequential pages accumulate)
→ DOM readback (count or last-row id changes)
```

Proxy-only evidence (request count, API JSON) is insufficient — see [`validation-proxy-trap.md`](validation-proxy-trap.md), [`../execution/validation-reasoning/state-visibility-gap.md`](../execution/validation-reasoning/state-visibility-gap.md).

## Related

- [`scroll-load-more-bottom-gate-stall.md`](scroll-load-more-bottom-gate-stall.md) — phase 2 trigger gate / in-flight pending / duplicate retry (same incident, later symptom)
- [`stale-derived-state.md`](stale-derived-state.md) — variant `stale_list_accumulator`
- [`session-scoped-implicit-state.md`](session-scoped-implicit-state.md) — cursor/videos pair restore
- Project: Vidoe-Test `.ai-skill/project/feedback/scroll-load-more-api-success-without-ui-append.md`
