# Final Core Exit Verification — Knowledge Governance Engine

**Date**: 2026-07-31  
**Plan**: `2026-07-30-0950-knowledge-governance-runtime`  
**Purpose**: Archive gate — prove Core Exit surfaces work; document intentional Later gaps.

## Verdict

| Question | Answer |
| --- | --- |
| Core Exit met (北極星：Knowledge becomes executable in Ai-skill)? | **YES** |
| Ready to archive this plan? | **YES** after this evidence + checkbox cleanup |
| Remaining Later work blocks archive? | **NO** — Plans pack / IDE host / external dogfood / Unified Execution are follow-up plans |

## Inventory (as of verification)

| Surface | Count / Location | Result |
| --- | --- | --- |
| Portable `DefaultRules` | **29** rules in `portable/kge/presentation.go` | OK |
| `obligation.commit.*` (runtime obligations list) | **29** lines | Parity signal OK |
| CLI | `kge check` / `validate [--advisory]` / `diagnose` | OK (below) |
| Git adapters | commit-msg count path; pre-push `kge check` | Wired in adapter |
| CI | `.github/workflows/ai-skill-cli.yml` step `kge validate --advisory` | Present |
| Glossary | `knowledge_governance_engine`, `capability_registry`, `runtime_adapter` | Registered |
| Portable boundary | `scripts/ai-skill-cli/portable/kge/README.md` | Documented |

## Test matrix (executed 2026-07-31)

| # | Check | Command / method | Result |
| --- | --- | --- | --- |
| V1 | Portable unit tests | `go test ./portable/kge/ -count=1` | **PASS** |
| V2 | Adapter / hook package tests | `go test ./internal/app/ -count=1` | **PASS** |
| V3 | D9 validate + full advisory | `ai-skill kge validate --advisory --root <repo>` | **PASS** — `Validation ok.` / `No advisory findings.` (clean tree) |
| V4 | D9 push watershed | `ai-skill kge check --root <repo>` | **PASS** — `Ready to push (validation ok).` |
| V5 | D9 IDE JSON surface | `ai-skill kge diagnose --root <repo>` | **PASS** — `[]` (no findings) |
| V6 | CI workflow contains advisory step | grep workflow | **PASS** — line present |
| V7 | Rule IDs include archival + registry transition | DefaultRules list | **PASS** — both present |

## Gaps (intentional Later — not Core blockers)

| Gap | Why not Core | Follow-up |
| --- | --- | --- |
| External project copy dogfood | Phase 3 optional | Separate dogfood when needed |
| IDE host (Cursor task → `kge diagnose`) | CLI surface exists; host wiring is tooling | Optional IDE task |
| Plans validation as KGE domain pack | Q4 deferred from day one | Later — Plans pack |
| Unified KGE + Delegation Execution Model | §D7 / Q10 explore | Separate explore plan |
| Expand more `behavioral_only` → advisory | Phase 5 partial (document_sizing only) | Incremental later |
| ADR promotion | Criteria partially met; external dogfood + registry↔capability formal mapping still open | When ADR criteria all green |

## ADR Promotion Criteria disposition

| Criterion | Disposition |
| --- | --- |
| KGE naming vs `runtime/` distinguishable | **met** — glossary + plan naming |
| Cap registry + ≥2 rules (projection optional) | **met** — `Cap*` contracts; most rules need no Projection stage |
| ≥1 non-plan pack enforced in-repo | **met** — Ai-skill first-party commit-msg / CLI pack |
| ≥1 external pack + thin adapter, reversible | **deferred** — Phase 3 Later |
| Enforcement Registry expresses rule_class → KGE capability without second coverage system | **deferred** — adapters map today; formal registry field = ADR follow-up |

## Archive recommendation

1. Mark plan `status: completed`.
2. Convert ADR criteria unchecked task-list into the disposition table above (no leftover `- [ ]`).
3. Move `plans/active/…-knowledge-governance-runtime/` → `plans/archived/…`.
4. Add row to `plans/README.md` §目前狀態 as completed.
5. Commit with body noting **deferred** Later items (external dogfood / Plans pack / IDE / Unified Execution / ADR remainder).
