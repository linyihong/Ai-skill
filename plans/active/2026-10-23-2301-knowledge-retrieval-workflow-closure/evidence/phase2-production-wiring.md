# Production lookup wiring and frozen E2E replay

## Result and boundary

The actual production candidate → fresh selector → watcher → complete path passes the frozen 26-case diagnostic acceptance. Phase 2 is complete within its explicitly narrowed diagnostic / engineering scope. This is exposed-data gold-blind execution, not independent generalization, actual model-token validation or workflow adoption. Phase 3 remains open, including R11/O03 project evidence, applicability, required_read/blocker preservation and real consumer dogfood.

No historical fixed-pool response or decision is a production input. The [protocol](phase2-production-protocol.md) was frozen before wiring. Selector prompt, disposition policy and P0 full-summary packing are unchanged. ATX extraction is the previously measured fix; bounded term-OR is explicit lookup candidate generation, while ordinary exact-phrase query remains unchanged. Source implementation `62bf1ae2`, all five rebuilt binaries `6f1c1bf0`; actual executable SHA and fixed corpus/index/dataset identities are in [results](phase2-production-results.json) and [manifest](phase2-production-manifest.json).

## Actual final-message results

| Measure | Observed |
| --- | --- |
| Specified-source hit | 23/23 positive queries |
| Source MRR | 0.9710144928: 22 rank-one cases plus R09 rank three, divided by 23 |
| Retrieved / direct / final confuser cases | 6/6 → 0/6 → 0/6 |
| Artificial negative suggestions | 0/2; not a natural no-answer error-rate estimate |
| Whole compact Go-JSON retrieval message | maximum 1622 runes; ceiling 2000 |
| Summary truncation / budget deferral | 0 / 0; all 28 direct occurrences emitted |
| Pool / unassessed / explicit unknown | 835 source occurrences / 797 / 3; total unknown 800 |
| Fresh quality attempts / retry | 26 / 0 |
| Warm production prepare p95 | 158.4ms over 520 warm calls; ceiling 500ms |
| Longest watcher observation | 46763ms; per-attempt deadline 300s |
| Continuous batch observation / maximum gap | 586684ms / 1.238695s; limits 3000s / 30s |
| Model snapshot and token usage | unavailable for all 26; never treated as zero |

The 26 exact request and raw files are SHA-bound in results; `phase2-prod-<case>.payload.json` contains the exact compact Go bytes, including JSON escaping, counted in the budget. Outer tool-only receipts/traces are separate. No necessary unknown state is hidden to pass the budget. R09's three sources retain the selector order; the historical lesson is rank three. These costs are retrieval-message runes and observed orchestration elapsed, not actual tokens or model inference latency.

R19 has an empty pool. R20 has a nonempty pool because source-level filtering retains missing status as unknown; its vague artificial query causes abstention. Its 0 is not empty-result precision. O06 excludes the stale-peer solution and leaves auth/session applicability unknown; no guessed revocation solution is emitted. Empty results never establish knowledge absence. Unknown/unassessed counts and trace remain visible.

## Source readback and handoff

[Handoff audit](phase2-production-handoff.json) reviews all 28 emitted occurrences, 23 unique source paths. Source identity and freshness pass actual complete checks; all summaries are preserved. Twenty-one positive queries have body-supported next-read purposes. R11/O03 correctly remain conditional on project evidence: the general lesson cannot supply actual package/model versions or failure fingerprints. O03's binary README is conditional evidence-navigation, not proof the current platform uses ELF. R08's runtime.db is a derived compiler/projection surface, not independent historical authority. R09 includes the summary, canonical governance and historical lesson without merging identities.

The handoff is useful when it points to a source and specifies what to verify; summaries need not become complete project answers. Applicability and adoption are not evaluated here. No new reviewer or selector exploration was added.

## Failure and identity controls

[Failure evidence](phase2-production-failures.json) is excluded from quality scores. Actual CLI watcher timeout at 501ms was followed by parent interruption and interrupted-state readback. Deliberately late valid synthetic JSON was rejected by complete because its watcher receipt failed. Invalid synthetic output was rejected by watcher and complete. In an isolated corpus copy, completion passed before a selected-source edit and failed afterward, emitting no payload. Synthetic fixtures are labeled and are not model attempts. The frozen quality corpus was unchanged.

Integration tests cover altered bundle/request/raw, missing or mismatched successful receipts, stale sources, full-summary order, unknown visibility and budget tail deferral. Existing lifecycle fixtures cover build/refresh/edit/add/delete/move/stale/missing metadata/index/deterministic rebuild and failed-write preservation. No old-pool fallback, retry, output repair or gold-based reordering was used.

## Verification and artifacts

`TestLookupPrepareIsProductionCandidateConsumer` first failed by absence. Full Go suite and targeted race tests passed after release rebuild. New test-only `TestLookupReplayExportPayloads` saves exact production message bytes; it does not change release behavior. The opt-in prepare/monitor/complete tests execute the release binary. [Production traces](phase2-production-traces.json), [continuous monitor](phase2-production-monitor.jsonl), [monitor events](phase2-production-monitor-events.json), [results](phase2-production-results.json), [manifest](phase2-production-manifest.json), [handoff](phase2-production-handoff.json) and [failures](phase2-production-failures.json) preserve provenance. Normalized outer traces replace local transport paths; raw request/response and payload bytes are preserved separately with SHA references. Two requests are stored as base64 (decoded SHA unchanged) because escaped JSON examples falsely match the shared-layer UNC-path scanner; normalized outer JSON uses semantically identical Unicode escapes. No model input or raw response was edited.

Release CI [40102700251](https://github.com/linyihong/Ai-skill/actions/runs/40102700251) passed Linux/macOS/Windows tests and artifact build. Final evidence/test-only commit CI readback is recorded separately when available. No claim is made about an unobserved run.
