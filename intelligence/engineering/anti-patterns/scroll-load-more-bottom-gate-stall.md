# Scroll Load-More Bottom Gate Stall

**Status**: `candidate-intelligence`

**Source**: Vidoe-Test home-ranking scroll load-more incident — phase 2 (2026-07-02 → 2026-07-03). Phase 1 (merge/ref-shadow) in [`pagination-accumulator-ref-shadow.md`](pagination-accumulator-ref-shadow.md). Promote after second independent infinite-scroll trigger stall with same gate sequence.

## Definition

Scroll load-more bottom gate stall: a **scroll-axis bottom trigger** treats one bottom reach as consumed (`bottomReachConsumed`) or ignores re-entry (`wasAtBottom` already true) while the user **remains at or returns to the bottom** before the in-flight fetch completes — so the next page never fires even though `hasMore` / `nextCursor` remain valid.

API and merge can be healthy; the failure is **trigger state machine authority**, not pagination data.

## Pattern

```text
user reaches bottom
↓
trigger fires once; bottomReachConsumed = true
↓
fetch in-flight (user still at bottom OR fast re-scroll to bottom)
↓
wasAtBottom stays true → no scroll-enter-bottom
OR duplicate/empty page → result false but consumed not released
OR fetch completes; pin release without scroll event sync
↓
user at bottom with hasMore but no new trigger
↓
must scroll up and back down (bad UX) or appears "stuck"
```

## Common variants

| Variant | Signal | Real failure |
| --- | --- | --- |
| `duplicate_page_stall` | API 200, cursor advances, `appended: 0` | treated as failed reach; `bottomReachConsumed` blocks retry at bottom |
| `in_flight_bottom_ignored` | fast scroll to bottom during fetch | `wasAtBottom` already true; no `enteredBottom`; pending not queued |
| `in_flight_dwell_no_pending` | user dwells at bottom through entire fetch | no leave/re-enter scroll events → pending never set until dwell-in-flight fix |
| `empty_page_cursor` | `videos.length === 0` but `nextCursor` set | returns `false` instead of retryable skip → gate stall |
| `false_without_rearm` | load-more returns `false` with `hasMore` | `bottomReachConsumed` never released on complete |
| `pin_release_no_scroll_event` | success + programmatic scrollTop nudge | DOM still at bottom band; no synthetic state sync |

## Rule

### Result taxonomy (load-more callback)

| Result | Meaning | Trigger follow-up |
| --- | --- | --- |
| `true` | appended rows | pin release; rearm; sync scroll state |
| `"duplicate"` | cursor advanced, dedup skipped all rows | auto-retry at bottom after cooldown |
| `false` / void | hard stop or gate skip | **release** `bottomReachConsumed`; do not leave consumed latched when `hasMore` |

Empty API page with valid `nextCursor` → **`"duplicate"`**, not `false`.

### In-flight bottom intent

- When `triggerIfAllowed` blocked because fetch in-flight **and** scroll metrics at bottom → set `pendingBottomReachWhileInFlight`.
- When user **dwells at bottom** during in-flight (scroll events without leaving band) → also set pending.
- After fetch completes → schedule `retry-after-pending-bottom-reach` after cooldown (skip pin release when pending).

### Post-fetch state sync

- Always run scroll state reconciliation after fetch complete (e.g. call scroll handler after pin release).
- Pin release alone is insufficient if the browser does not emit scroll events.

### Forbidden shortcuts

- Closing on «Network shows cursor requests» without checking **poster count growth** or explicit duplicate retry logs.
- Testing scroll **request count only** without in-flight bottom re-entry and duplicate-page paths.
- Using `ref.current` as merge source (see ref-shadow doc) — orthogonal but often co-occurring.

## Observe checklist (trigger stall incidents)

Before blaming API or backend cursor:

1. Is user **still at bottom** when stall happens?
2. Does debug show `trigger:blocked:bottom-consumed` or `scroll:leave-bottom-ignored-in-flight`?
3. Did last page return `appended: 0` with cursor advance (`duplicate-page`)?
4. Was fetch in-flight when user reached bottom again?
5. Did merge layer pass (poster count) on prior pages?

## Validation

Minimum evidence chain (four layers for infinite scroll):

```text
trigger (one reach per deliberate bottom entry; in-flight pending retry)
→ merge (functional setState; featured-only dedup)
→ result taxonomy (true | duplicate | false+rearm)
→ DOM readback (poster count grows or explicit duplicate retry)
```

Unit tests should cover:

- duplicate at bottom auto-retry without scroll up
- bottom re-entry during in-flight fetch → pending → post-complete retry
- dwell at bottom during in-flight (scroll events, no leave)
- generic `false` does **not** auto-chain while dwelling (cooldown + no duplicate schedule)
- success does **not** chain while dwelling without pin release / re-entry

Debug contract: `?scrollDebug=1` → log `defer-bottom-reach-in-flight`, `scroll:bottom-dwell-in-flight`, `retry-after-pending-bottom-reach`, `loadMore:duplicate-page`, `loadMore:empty-page-advance-cursor`.

## Related

- [`pagination-accumulator-ref-shadow.md`](pagination-accumulator-ref-shadow.md) — merge / dedup / ref authority (phase 1)
- [`stale-derived-state.md`](stale-derived-state.md) — `wasAtBottom` / `bottomReachConsumed` as stale derived scroll intent
- [`validation-proxy-trap.md`](validation-proxy-trap.md) — API 200 ≠ UI append ≠ trigger rearm
- Project: Vidoe-Test `.ai-skill/project/feedback/scroll-load-more-bottom-gate-stall.md`
