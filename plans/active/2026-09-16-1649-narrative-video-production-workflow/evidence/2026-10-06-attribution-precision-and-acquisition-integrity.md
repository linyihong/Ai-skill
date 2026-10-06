# Attribution precision and acquisition integrity

Status: contract confirmation / adapter validation gaps; not full content, timing or render acceptance.

## Reusable findings

1. Preserving Latin observations repairs an evidence-retention defect, but does not prove subtitle precision. Pair real dialogue with clothing, logos and device UI negatives in the same source window.
2. A changing dialogue band can accidentally lend credibility to nearby scene text. Band proximity is supporting evidence, not independent region-role attribution. Keep ambiguous observations uncertain rather than publish them or erase them.
3. Mixed scripts and vertically separated boxes do not prove bilingual subtitles: one box may be a clothing label. Region classification precedes bilingual grouping, and downstream fusion must preserve that identity.
4. Decoder/resource failure can truncate acquisition while later resolution still completes. A successful process or nonzero cue count cannot certify an exhaustive scan. Record declared source extent, observed scan extent and termination reason; incomplete acquisition remains REVIEW or BLOCK according to the consumer's requirements.
5. Cached frame labels, derived frame-rate times and independently decoded source presentation times can disagree. Validate media identity and source timing before certifying a short dialogue window or applying an offset. Do not turn an observed discrepancy into a global timing correction.
6. Resolve the exact caption artifact from the run record, not an ambiguous directory glob. Metadata sidecars are not caption documents; missing cue structure must fail artifact integrity before exclusion-only checks, rather than default to an empty caption list and falsely pass.
7. Complete raw-observation accounting does not prove identity-preserving projection. Independently owned bilingual regions can be acquired correctly but collapse into a single language with missing boxes during fusion or resolution. Verify both language regions, stable distinct identities, owned geometry and the final group relation. A cross-language resolution label alone proves neither bilingual retention nor semantic correspondence.

## Validation boundaries

- Offline retention and geometry fixtures protect specific contracts, not full corpus quality.
- Frame-reviewed negative pairs and an isolated mechanical attribution replay revealed a band-borrowing failure. A joint scene/style adapter repair now passes that regression while retaining neighboring dialogue and explicitly qualified subtitle positives. This is scoped evidence, not exhaustive precision acceptance; profile thresholds are not promoted into canonical policy.
- A resource-truncated acquisition is excluded from complete baselines. GPU-heavy integration runs should be serialized unless whole-pipeline resource budgets and concurrency have been verified.
- The adapter's source-window timing discrepancy is unresolved; retained raw text is not independently verified timing evidence.
- An artifact-kind negative was demonstrated before tightening the acceptance checker; metadata cannot serve as caption exclusion evidence. This verifies the test infrastructure, not subtitle quality.
- Independent frame review and raw-region readback confirmed a bilingual pair, while its final cue originally lost one language and owned geometry. Test-first projection repair and a frozen-input downstream replay now pass the specified owned-region and accepted-artifact checks. The first repaired replay restored both regions but exposed a final gate that compared a flattened bilingual string against one box; that failure was retained and regression-tested before the gate repair. Do not compensate for projection loss by increasing global acquisition frequency.
- Content recovery, bilingual propagation, temporal integrity and rendered-output checks remain open. No full-pack or publish-ready checkbox is promoted.

## Adapter refinement confirmed in scoped replay

- Match downstream ownership with each region's source identity **and its own observed text**, within the relevant time window. Identity-only matching or flattened display-text matching does not establish attribution.
- Subtitle qualification precedes pairing. Shared acquisition group, frame/source ownership and overlapping region windows can support a pair; equal timestamps, mixed scripts or nearby boxes alone cannot.
- Namespace targeted-probe identities separately from whole-source identities, so unrelated observations cannot alias during trace or projection.
- Account for owned region texts in candidate-to-cue tracing. Restoring boxes without correcting downstream gates and accounting can still leave valid evidence uncertain or apparently missing.
- Offline fixtures and a specified frozen-input projection check do not independently certify spoken-audio truth, semantic equivalence, fresh acquisition completeness, locale propagation, timing or rendered output.

## Linked updates

Maps to [resolution companion](../43-multimodal-text-evidence-resolution.md) and [candidate precision companion](../47-subtitle-candidate-detector.md). Existing workflow contracts already require joint evidence and independent timebase validation; no new route, agent, global sampling policy or canonical numeric threshold is introduced. Concrete titles, source timestamps, hashes, class names and run outputs remain in consumer project documentation.
