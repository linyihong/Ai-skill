# Dogfood 2y — KaizenWMS Phase 2 SPA scaffold（external greenfield consumer）

**Date**: 2026-07-21  
**Consumer**: `kaizenwms`（Bitbucket；`plans/active/2026-07-14-1610-wms-material-foundation`）  
**Loop**: Orchestrator → Executor → Verifier → Arbitration  
**Outcome**: `slice_compliant_closed`（`phase-2-spa-scaffold`）

## What ran

| Leg | Notes |
| --- | --- |
| Orchestrator | Brief + verification_backfill in plan YAML；git-anchored before Execute |
| Executor | `KaizenWms.sln` + layered `src/` + Angular 19/PrimeNG + esproj；TFM initially `net9.0`（host SDK max was 9） |
| Verifier | Independent V1–V4；**A1–A5 PASS** |
| Arbitration | MAJOR on **C1b**：A3 was `tier: deliverable` only on a **combined / user-visible** acceptance |

## Feedback to this plan / SD contract

1. **C1b still earns its keep on greenfield scaffolds.** Implementation can look “done” (routes + dualLabel code) while backfill stays deliverable-only. Verifier correctly refused `slice_compliant_closed` until Orchestrator linked integration + acceptance-spec smoke (published deep-link HTTP 200 + bundle zh/en dual-label).
2. **Recommend consumer overlay reminder**: for `slice_type: combined`, any acceptance whose text mentions UI render/navigate must backfill `integration` or `acceptance-spec` *before* Execute — not after Verifier MAJOR.
3. **Consumer verify hygiene (local)**: macOS `/bin/bash` 3.2 + `set -u` broke `[[ -v VAR ]]` and empty `"${arr[@]}"` in project verify scripts during close-out; fixed in consumer. Not an Ai-skill defect, but dogfood friction when agents force `/usr/bin` ahead of Homebrew bash.

## Evidence pointers（consumer repo）

- `plans/.../evidence/2026-07-21-phase-2-verifier.md`
- `plans/.../evidence/2026-07-21-phase-2-a3-integration-smoke.md`
- `plans/.../evidence/2026-07-21-phase-2-arbitration.md`

## Disposition

- No schema change requested this round.
- Optional doc tweak（advisory）：kit / delegated-execution example row for “scaffold shell + routes = user-visible → not deliverable-only”.
