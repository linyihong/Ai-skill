# Attribution precision and acquisition integrity

Status: contract confirmation / adapter validation gaps; not full content, timing or render acceptance.

## Reusable findings

1. Preserving Latin observations repairs an evidence-retention defect, but does not prove subtitle precision. Pair real dialogue with clothing, logos and device UI negatives in the same source window.
2. A changing dialogue band can accidentally lend credibility to nearby scene text. Band proximity is supporting evidence, not independent region-role attribution. Keep ambiguous observations uncertain rather than publish them or erase them.
3. Mixed scripts and vertically separated boxes do not prove bilingual subtitles: one box may be a clothing label. Region classification precedes bilingual grouping, and downstream fusion must preserve that identity.
4. Decoder/resource failure can truncate acquisition while later resolution still completes. A successful process or nonzero cue count cannot certify an exhaustive scan. Record declared source extent, observed scan extent and termination reason; incomplete acquisition remains REVIEW or BLOCK according to the consumer's requirements.
5. Cached frame labels, derived frame-rate times and independently decoded source presentation times can disagree. Validate media identity and source timing before certifying a short dialogue window or applying an offset. Do not turn an observed discrepancy into a global timing correction.
6. Resolve the exact caption artifact from the run record, not an ambiguous directory glob. Metadata sidecars are not caption documents; missing cue structure must fail artifact integrity before exclusion-only checks, rather than default to an empty caption list and falsely pass.

## Validation boundaries

- Offline retention and geometry fixtures protect specific contracts, not full corpus quality.
- Frame-reviewed negative pairs and an isolated mechanical attribution replay revealed a band-borrowing failure. The failure remains an executable regression; neighboring true dialogue remains a required positive assertion.
- A resource-truncated acquisition is excluded from complete baselines. GPU-heavy integration runs should be serialized unless whole-pipeline resource budgets and concurrency have been verified.
- The adapter's source-window timing discrepancy is unresolved; retained raw text is not independently verified timing evidence.
- An artifact-kind negative was demonstrated before tightening the acceptance checker; metadata cannot serve as caption exclusion evidence. This verifies the test infrastructure, not subtitle quality.
- Content recovery, bilingual propagation, temporal integrity and rendered-output checks remain open. No full-pack or publish-ready checkbox is promoted.

## Linked updates

Maps to [resolution companion](../43-multimodal-text-evidence-resolution.md) and [candidate precision companion](../47-subtitle-candidate-detector.md). Existing workflow contracts already require joint evidence and independent timebase validation; no new route, agent, global sampling policy or canonical numeric threshold is introduced. Concrete titles, source timestamps, hashes, class names and run outputs remain in consumer project documentation.
