# Phase 2 source lookup and runner engineering

## Scope

The exposed 26-case selector method exploration is closed per user review. This change adds production source-level lookup and a formal file-response watcher; it does not wire the workflow consumer or adopt v1/v2 ranking.

Protocol: [engineering scope](phase2-engineering-protocol.md). The seven intelligence probes were frozen before implementation in [cases](phase2-intelligence-cases.json); selected source bodies/checksums remain unchanged. [Old executable baseline](phase2-intelligence-before.json) is 0/7, MRR 0. [New-source preflight](phase2-intelligence-after.json) found 7/7 at rank 1. Release source is `f9570120`; all five binaries were rebuilt from that commit. [Release replay](phase2-intelligence-production.json) is 7/7, source MRR@5 1.0; [manifest](phase2-engineering-manifest.json) binds binary SHA, generator version/source inputs, dataset SHA and source checksums. No selected source body changed. Warm CLI p95 is 123.763 ms over 21 samples (three warm repetitions per probe; first invocation excluded), below 500 ms in this diagnostic. This is a local observation, not an inference-time or long-term production p95 claim. Artifact transport paths are normalized placeholders; canonical source identities, scores and executable/index hashes are unchanged. These are source-specific exposed body-coverage diagnostics, not a natural-language or generalization benchmark.

## Implemented behavior and verification

- Selected validation-reasoning body FTS is separate from preview summary. Exact phrase search stays intact. One atom per canonical source; opt-in `--source-level` deduplicates before limit (maximum 64).
- Canonical checksum/status/domains/scope/replaces/readback are separate from atom compatibility status and relevance rank. Missing metadata stays unknown; status filtering excludes declared mismatches and retains unknown for readback. No scope or trust is inferred from a folder name.
- Existing graph query stays single-hop; output now includes depth=1. Cycle, edge-type filter, limit and deterministic-order fixture checks pass. Relation does not establish support or authority.
- Input manifest includes generator version, effective atom data, source contents and graph/glossary membership. Source-level queries check before and after lookup. Initial build, edit/add/move/delete, missing metadata/index, canonical-vs-route status, source dedup, deterministic rebuild and failed-build cache preservation are covered. `TestRuntimeFailedIndexPopulationPreservesCache` forces a duplicate-primary-key failure during SQL population, verifies byte-identical preservation of the old DB and removal of the temporary DB; this exercises replacement failure after writes begin, separately from metadata parse failure. Old atom lookup keeps legacy schema compatibility and does not gain the freshness guarantee.
- Build writes a same-directory temporary SQLite DB and replaces the cache only after successful population and stable-input checks. Generated indexes remain ignored/rebuildable.
- `runtime selection-watch` is a real CLI consumer of the previously test-only strict response parser. Fixed request, atomic response publication, preexisting-output rejection, deadline, invalid output, request mutation and usage failure have explicit states. Selected must exactly equal direct assessments; no repair/retry/fallback. Supplied usage 0 is available; missing usage is unavailable/null.

Test-first evidence: `TestRuntimeIntelligenceBodyCoverage` failed with `body source absent: []` before implementation; watcher behavior tests failed with `unsupported runtime command: selection-watch`. Targeted behavior suite and `go test ./...` passed after implementation. Targeted `go test -race` passed; full `go test ./...` passed again after the all-platform rebuild, including binary/source parity. GitHub workflow readback remains pending until push.

## Actual timeout / interrupt handoff

[Raw timeout and adapter readback](phase2-runner-timeout.json) records one deliberately nonresponsive existing child, `/root/blind_26`, without rerunning selector quality. Watcher returned timeout at 501 ms, usage unavailable/null, interrupt_required=true and interrupt_confirmed=false. The parent then called `collaboration.interrupt_agent`, which returned previous_status=running; `collaboration.list_agents` readback returned interrupted. No response file was created.

[Release-binary cancellation](phase2-runner-cancel.json) additionally records actual macOS SIGINT → canceled with interrupt_required=true and exit 30; it does not claim a child was stopped by that signal.

This validates the timeout-to-host-interrupt handoff. The CLI cannot invoke host tools or claim it stopped a child by itself. Production workflow orchestration and process/inference accounting remain separate; no inference-duration or dollar-cost estimate is derived from elapsed time.

## Remaining gates

Phase 2 remains incomplete. The historical lesson query/extraction improvements and selector still need a reusable production candidate-to-selection path and the full frozen acceptance replay; this change does not silently adopt the failed v1/v2 strategies. Full query intent coverage/generalization, whole-consumer context cost and the formal production consumer boundary are not established by seven probes or structural watcher tests. R11/O03 project evidence and actual applicability/adoption remain Phase 3. Actual tokens and model snapshot/usage not exposed by the host remain not evaluated/unavailable.

## Reproduction

From `scripts/ai-skill-cli`, build the cache with `AI_SKILL_ENGINEERING_INDEX=<new-index> go test ./internal/app -run TestRetrievalEngineeringExportIndex -count=1`. Then set `AI_SKILL_ENGINEERING_QUERY_BINARY` to an explicitly pinned binary, `AI_SKILL_ENGINEERING_QUERY_INDEX` to that cache, `AI_SKILL_ENGINEERING_OUTPUT` to a new report path and `AI_SKILL_ENGINEERING_SOURCE_LEVEL=1`; run `go test ./internal/app -run TestRetrievalEngineeringSourceComparison -count=1`. Independent artifact readback recomputed both score sets, dataset/source/binary/input fingerprints and confirmed 0/7 → 7/7. The harness verifies frozen source SHA and binary SHA, records generator file checksums, deduplicates before source@5 scoring, and preserves normalized query output. It refuses to overwrite a report. Old baseline used the old generator/executable with legacy atom output normalized to source@5; it does not use the new source-level flag.
