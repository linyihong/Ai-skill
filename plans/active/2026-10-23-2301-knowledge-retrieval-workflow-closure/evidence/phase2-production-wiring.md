# Production lookup wiring and frozen E2E replay

## Implementation status

User accepted the previous engineering slice and requested the actual production candidate-to-selector path. The CLI now exposes `runtime lookup index`, `prepare` and `complete`, using source-level query/freshness, neutral request generation, the accepted byte-identical selector prompt, the production watcher receipt and direct-only P0 packing. Historical fixed-pool choices are not production inputs.

Scope and acceptance were frozen before implementation in [production protocol](phase2-production-protocol.md). ATX heading extraction is the existing v1 extraction fix, promoted without lifecycle changes; Unicode-term OR is explicit to lookup candidate generation, not a default change to exact-phrase query. No weights, selector prompt or selection method are tuned.

`TestLookupPrepareIsProductionCandidateConsumer` first failed with unsupported runtime command. Targeted integration tests now cover new production candidates, actual watcher transport, completion, full summaries/unknown counts, altered bundle/request/raw, failed watcher and source changes. Packer tests preserve ordering, tail deferral and complete-message budget. The selector prompt is checked byte-for-byte against the accepted frozen prompt.

Release binaries, real CLI replay, fresh attempts and full frozen acceptance results are pending this source/binary commit sequence. Phase 2 remains incomplete; no old fixed-pool score is reused as a pass. Real applicability, R11/O03 project evidence, required_read/blocker and adoption remain Phase 3.
