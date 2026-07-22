# Dogfood Evidence — external player project (Change Intent Lock pilot)

**Plan**: [`_plan.md`](_plan.md) Phase 4 (second evidence path)
**Date**: 2026-06-25 (pilot design) · **observed / partial-verified** 2026-06-29
**Project**: an external landscape video-player project (project-layer; **not** Ai-skill).
Concrete repo, plan-doc paths, component names, and hook wiring are recorded
**project-locally** in that project's own `.ai-skill/` overlay (sanitized out of the
shared layer per policy).
**Evidence class**: **partial-verified** (happy path / structure-transition only)

Complements [`01-dogfood-evidence.md`](01-dogfood-evidence.md) (**verified** — `force_exit` path).
This file records **cross-repo** evidence that `preparatory_refactoring` + Change Intent
Lock can constrain real feature work without ai-skill commit-msg runtime projection.

> **Not** a success-story completion claim. Phase 1 feature work remained in progress;
> this evidence validates **execution contract classification** (structure → transition →
> feature started), not full feature delivery.

---

## Evidence maturity (this file)

```text
Observed → Partial Verified → Verified (behavior proven) → Promoted (independently auditable)
```

| Level | Meaning | This file |
|-------|---------|-----------|
| Observed | Narrative + artifacts named | yes |
| Partial Verified | Structure transition + guard observed; canonical `exit_when` mapped | **current** |
| Verified | Observable equivalence / behavior proven (Gate A) | **not yet** |
| Promoted | Independently auditable pointer + ai-skill validator/scenario wiring | **explicitly no** (observation period) |

**Upgrade gates**: Gate A → **Verified**; pointer/SHA → **Promoted** (not blocking collection or Verified).

---

## Intake routing (mapped to dual-axis)

```yaml
change_kind: feature
blocked_by_structure: true
# Rationale: the target player frame is preview-gated, swipe-mutex, tab-shell z-index,
# and streaming-metadata coupled — the feature cannot land safely without contracts first.
execution_mode: preparatory_refactoring   # project analog; native plan uses Phase 0/1 wording
```

**Not** `replacement` / `migration` — no parity inventory; orthogonal per `intake.md`.

---

## Implementation plan (retrospective mapping)

> **Retro-map declaration**: the YAML below is **retrospective mapping for evidence
> review**; it is **not** the project's native plan schema. The authoritative project
> plan is prose Phase 0 / Phase 1 in the project's own plan doc.

```yaml
execution_mode: preparatory_refactoring   # implied; plan labels Phase 0 / Phase 1
steps:
  - id: phase-0-contracts
    intent: structure
    behavior_change:
      allowed: false
    action: screen-mapping + player-spec §Landscape + hazard notes + BDD trace refs
    checkpoint:
      observable_equivalence:
        required: true
        # Intent (not yet proven): portrait + vertical-snap + preview-gate unchanged

  - id: phase-1-mvp
    intent: feature
    behavior_change:
      allowed: true
    action: player-frame landscape mode + page style variants
    validation: landscape-mode integration test + deploy smoke

  - id: phase-3-closure
    intent: feature
    behavior_change:
      allowed: true
    action: integration green + claim registry + plan status completed
```

**Intent transition**: `structure → feature` allowed after Phase 0 artifacts exist
(project file_exists + human gate). **Not** the same as `observable_equivalence_passed`
until regression evidence is attached (see Gate A).

---

## Path classification

| Segment | Classification | Status (2026-06-29) |
|---------|----------------|---------------------|
| Phase 0 contracts | structure intent | **observed** — artifacts exist |
| Structure → feature transition | transition gate | **partial-verified** — guard + mapping/spec gate |
| Phase 1 implementation | feature | **in progress** — WIP; not independently auditable from Ai-skill |
| force_exit | N/A | Not triggered — contracts reduced local feature cost |

**Valid evidence path**: **Happy path (partial)** — structure phase + transition observed →
feature work started. Differs from [`01-dogfood-evidence.md`](01-dogfood-evidence.md)
(`force_exit` teaching case).

---

## Gate A — observable_equivalence（blocking for Verified）

Per [`execution-modes.md`](../../../workflow/software-delivery/implementation/execution-modes.md): **checkpoint ≠ observable equivalence**.

| Claim level | What we have | Status |
|-------------|--------------|--------|
| `checkpoint_exists` | mapping file + spec §Landscape + integration scaffold | **yes** (2026-06-29) |
| `checkpoint_valid` | portrait / snap / preview **regression executed and green** | **no** — not recorded in this evidence file |

**Do not conflate** project guard `file_exists` checks with `observable_equivalence_passed`.

**Verified 2026-06-29 (artifact gate only)**: the landscape screen-mapping doc and the
player spec §Landscape both exist. **Pending for Verified upgrade**: command + date +
pass/fail for the portrait-player regression owned by the Phase 0 checkpoint intent.

**Failure lens**: aligns with §8 `fake equivalence` if we claimed Valid without regression proof.

> Superseded as the happy-path Verified source by
> [`03-external-spa-return-qty-verified.md`](03-external-spa-return-qty-verified.md),
> which closes Gate A with an isolated observable-equivalence run.

---

## Gate B — exit_when（recorded at Partial Verified)

Canonical vocabulary from [`execution-modes.md`](../../../workflow/software-delivery/implementation/execution-modes.md) §4. Satisfied for **partial-verified**; does not alone satisfy **Verified** (see Gate A).

| `exit_when` candidate | Applies? | Evidence |
|-----------------------|----------|----------|
| `target_change_becomes_local` | **partial** | Phase 0 contracts + guard unblocked **starting** the player-frame work; feature not closed |
| `target_test_becomes_expressible` | **yes (primary)** | landscape integration scaffold + fixture exist — landscape acceptance can be expressed before/alongside implementation |
| `new_abstraction_created` | no | No new seam/abstraction retained for feature to consume |

**Recorded exit_when (partial-verified)**: **`target_test_becomes_expressible`** — the
integration scaffold + fixture make landscape behavior testable; structure phase exited
into feature intent on that basis **plus** artifact gate, **not** on equivalence proof
(Gate A gap). `force_exit_when`: **none** (happy partial path).

---

## Mechanical enforcement (project overlay)

The project carries its own commit-time guard (a plan-phase registry + a pre-commit hook
wired through its editor tooling + a BDD test, with an opt-out env var). **Guard rule**:
staging the player-frame component or its page styles **denies commit** until the mapping
doc + spec §Landscape exist. Concrete file paths are in the project's local overlay.

**Ai-skill posture**: project overlay only; aligns with the observation-period
**no ai-skill commit-msg validator** decision.

---

## Phase 0 artifacts (structure phase)

The structure phase produced (in the project repo): a main plan doc, a landscape screen
mapping, a visual spec §Landscape, a landscape integration scaffold, and a player fixture.
Exact paths live in the project's local `.ai-skill/` overlay.

---

## External pointer（blocking for Promoted only)

| Field | Value |
|-------|-------|
| Repo | external (not vendored in Ai-skill workspace) |
| Commit SHA | *TBD — attach in the project-local overlay when Phase 0 artifacts land on remote* |
| Verifier | *TBD — name + date + command for guard simulation* |
| Reproducibility | **not independently reproducible from Ai-skill repo alone** (disclosed) |

**Not** blocking Phase 4 collected. **Not** blocking **Verified**. **Blocks Promoted** until independently auditable.

---

## Failure-mode review lens

| failure | Project observation |
|---------|---------------------|
| intent oscillation | Plan orders Phase 0 → 1 → 3; guard blocks implementation-first commits |
| structure inflation | Phase 0 scoped to mapping + spec regions, not broad player rewrite |
| fake equivalence | **Gate A open** — artifact gate ≠ portrait regression proof |
| abstraction orphan | N/A — no unused parser/seam (contrast `01` prep-02) |
| compatibility collapse | Native plan uses Phase 0/1; no forced `execution_mode` on legacy plans |
| replacement parity misuse | Plan explicitly frontend-only; no replacement inventory |
| illegal transition | Not wired — would need native `steps[]` YAML + planvalidate advisory |

---

## What this evidence supports in `_plan.md`

| Plan section | Supported? | Notes |
|--------------|------------|-------|
| Phase 4 evidence path collected | **yes** | Happy **partial** complements `01` force_exit **verified**; superseded for Verified by `03` |
| Phase 3 routing cross-links | **partial** | Concept via project plan wording, not ai-skill `loading_surfaces` |
| Phase 5 glossary | **no** | Project uses Phase 0/1 labels |
| Validator hook / enforcement | **explicitly no** | Project overlay ≠ ai-skill promotion |

---

## Recorded fields (Phase 4 checklist)

| Field | Value |
|-------|-------|
| change_kind | feature |
| execution_mode | preparatory_refactoring (project analog) |
| evidence_class | partial-verified |
| intent sequence | structure (Phase 0) → feature (Phase 1+, in progress) |
| transitions | structure → feature after artifact gate; equivalence **not** verified |
| exit_when | `target_test_becomes_expressible` (primary); `target_change_becomes_local` (partial) |
| force_exit trigger | none |
| checkpoint | `checkpoint_exists` yes; `checkpoint_valid` **pending** |
| blocking | project editor hook; ai-skill planvalidate `Blocking=false` |
