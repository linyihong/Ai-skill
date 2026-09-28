# Phase 3 delivery contract

Companion to [`_plan.md`](_plan.md) and [`04-phase-3-dogfood.md`](04-phase-3-dogfood.md).
This file makes Phase 3 executable without treating a fictional example as a
real run. It does not add a workflow route, runtime projection, provider, or
new narrative-analysis schema.

## Two dogfood levels

| Level | Purpose | Completion authority | Does not require |
| --- | --- | --- | --- |
| Phase 3A — contract dogfood | Prove that real source material can form an auditable decision chain through `cut-ready`. | Producer self-check plus evidence review of the de-identified bundle. | Publishing, outcome window, profile/locale/band questions being frozen. |
| Phase 3B — publication dogfood | Prove locale QC, independent completion authority, and the outcome loop. | Fresh reviewer for `publish-ready`; outcome author for the recorded window. | A positive performance result; `insufficient_sample` is valid. |

Phase 3A is the next blocking milestone. Phase 3B follows it and must not be
silently claimed by a 3A run.

## Phase 3A minimum bundle

One de-identified real run must include all of the following, using stable
opaque ids that allow the artifacts to be cross-checked without revealing the
source work:

1. `external_run_ref`: an opaque, non-reversible run token; its private mapping
   remains in `<PROJECT_ROOT>`.
2. Brief lock: platform, target duration, promise, and forbidden items.
3. Source bible: a `bible_id` and either at least one de-identified entity or
   `no_character_entities: true`.
4. Clip catalog: at least three existing `clip_id` entries with in/out, duration,
   visual description, and tags.
5. Template selection: one primary `narrative_template_id` and `template_kind`.
6. Matching: one to three shots, each with constraints, non-empty feasible set,
   explicit selection policy and rationale, and a selected id in that set.
7. EDR: matching ids and shots align; the recorded timeline has passed producer
   self-check and is `cut-ready`.
8. Locale: one publish locale has separately recorded content, timing, and layout
   decisions. A hold is a valid result when its reason and rollback owner are
   recorded; it is not a `publish-ready` claim.

The evidence file records counts, decision ids, gate verdicts, and failure
classification only. It must not contain media names, raw dialogue, people,
paths, hosts, credentials, or reversible hashes.

## Phase 3B additions

Phase 3B extends the same run or a later comparable run with:

1. An assembled output reconciled to EDR shots.
2. A fresh review record containing reviewer role, reviewed artifact ids,
   blocking-gate verdicts, exceptions, rollback target, and timestamp.
3. Publish evidence and an outcome record. An unopened or insufficient window
   must be recorded as `insufficient_sample`, never inferred as success.

## Gate evidence modes

Before mechanical tooling exists, a gate remains valid but its evidence mode
must be explicit:

| Mode | Meaning |
| --- | --- |
| `structural` | Required ids, membership, and field presence can be checked mechanically. |
| `manual` | A reviewer checks the rendered artifact or semantic claim. |
| `adapter-backed` | An external tool measures it; the tool is not canonical workflow logic. |

Each Phase 3 evidence bundle names the mode used for every reported gate. A
natural-language eligibility rule must not be represented as an already-running
mechanical validator.

## Stop rules

- A missing bundle field is `data_insufficient`, not a reason to add a detector
  or expand the workflow.
- A true Phase 2 contract inability is recorded as `contract_gap` with the
  exact decision that cannot be represented; it needs user approval before a
  workflow change.
- Candidate research may be linked from the evidence index but cannot block 3A
  unless it violates an existing invariant or gate.
