# Dogfood Evidence — Happy path (structure → checkpoint → stop → feature), Gate A

**Plan**: [`_plan.md`](_plan.md) Phase 4
**Date**: 2026-07-22
**Evidence class**: **verified** (happy-path; observable equivalence proven — Gate A satisfied)
**Project**: external Angular SPA with a mock domain store — project-layer dogfood.
Concrete repo, commit SHA, and file paths are recorded **project-locally** in that
project's non-git `.ai-skill/dogfood/` overlay (sanitized out of the shared layer).
**Task**: add a mock-store command-validation rule that was gated by a monolithic
command-dispatch `switch`.

Pairs with:
- [`01-dogfood-evidence.md`](01-dogfood-evidence.md) — **verified**, `force_exit` / stop-design path.
- `02` (earlier project-layer dogfood) — **partial-verified**, structure→transition only (equivalence not proven).

This record is what `02` was missing: an **observable-equivalence checkpoint** on the
structure step (Gate A → Verified).

## Intake routing

```yaml
change_kind: feature
blocked_by_structure: true    # moderate — see honesty note
execution_mode: preparatory_refactoring
```

**Not** replacement parity — no old/new capability inventory; the intake parity gate untouched.

## Intent sequence (as executed)

```yaml
steps:
  - id: prep-01
    intent: structure
    behavior_change: { allowed: false }
    action: extract the ~190-line command-dispatch switch into a handler registry
    checkpoint:
      observable_equivalence:
        required: true
        evidence: SPA browser e2e (Playwright) — full suite green, unchanged from baseline
  - id: stop
    exit_when: new_abstraction_created   # per-command seam exists; structure mode ends
  - id: feat-01
    intent: feature
    behavior_change: { allowed: true }
    action: reject a return quantity greater than the line's available qty (422)
    validation: new acceptance — first store-contract unit test on the seam
```

No `feature → structure` re-entry; no illegal transition. Stop fired on the first
`exit_when` (`new_abstraction_created`) — the registry gave the feature its landing seam.

## Observable Equivalence Checkpoint (Gate A)

The structure step is behavior-preserving. Two independent oracles held green:

| Oracle | Before structure | After structure (isolated) | After full change (structure+feature) |
|--------|------------------|----------------------------|----------------------------------------|
| Playwright e2e | 10 passed / 2 skipped | **10 passed / 2 skipped** | **18 passed / 2 skipped** ¹ |
| Karma unit | 30 passed | — | **32 passed** (30 prior + new spec + 1 tree-added) |

¹ The e2e suite grew mid-session (see environment note); the 18-green run carries
both structure + feature but no e2e exercises the new rejection path (the return e2e
is a skipped stub), so it still witnesses equivalence for all pre-existing behavior.
The **isolated** structure-only equivalence (10→10, feature not yet added) is the
clean Gate A proof; the 30→30 unchanged unit tests corroborate at unit level.

`checkpoint_exists` **and** `checkpoint_valid`: the oracle is a real regression suite
that would fail on behavior drift, not a tautological assertion.

## Feature acceptance (new observable behavior)

- Rule: a return quantity must not exceed the line's available (pending-return) qty.
- The guard rejects `qty > available` with a `422` error **before any mutation**, so
  no return-type domain event is written — matching the scenario's
  "command rejected / no Return event written".
- The new unit test is registered in the feature's FeRef; the project's own BDD
  scenario↔FeRef sync gate enforces the linkage (it blocked the commit until linked —
  an independent check that the acceptance is real).

## Maturity ladder

```text
Observed → Partial Verified → Verified (behavior proven) → Promoted (independently auditable)
```

- **Verified** — claimed. Behavior proven: structure equivalence (Gate A) + feature
  acceptance, both via the project's own test gates (karma + Playwright + BDD sync),
  all green in the pre-commit hook.
- **Promoted** — a real SHA pointer exists (in the project repo, recorded in its
  `.ai-skill/dogfood/` overlay), **but the commit is local / not pushed**, so
  independent audit is not yet possible. Promoted remains **pending push + external audit**.

## Honesty notes (do not over-read)

1. **`blocked_by_structure` strength = moderate, not strong.** The command-dispatch
   function was already exported and unit-invocable, so the rule *could* have been
   inlined into the old switch. The genuine structural payoff is (a) per-command
   isolation and (b) it being the first store-contract unit-test target. This is a
   *seam-enabling* case, not a *cannot-proceed-without-refactor* case — recorded as
   such rather than inflated.
2. **Environment drift mid-session.** The project branch advanced during the run, a
   declared dependency was uninstalled (`npm ci` resolved it), and one verify script
   had lost its exec bit (restored in the same commit). None caused by the dogfood;
   noted for reproducibility. Details in the project-local pointer.
3. **Commit granularity.** Structure and feature landed in one commit, not two.
   Change-Intent-Lock commit separation is a byproduct, not the mechanism; the temporal
   isolation of the equivalence proof (structure-only e2e run before the feature edit)
   is what carries Gate A, independent of commit shape.

## Gate status

- [x] Gate A — observable equivalence — **passed** (structure-only e2e 10→10; unit 30→30)
- [x] Feature acceptance — **passed** (karma 32/32, new spec green, FeRef-linked)
- [x] Verified — **claimed**
- [ ] Promoted — pending `git push` of the project commit + independent audit
