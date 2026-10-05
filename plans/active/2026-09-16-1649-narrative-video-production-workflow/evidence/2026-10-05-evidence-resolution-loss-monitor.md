# Observation — Evidence resolution loss must be observable and recoverable

**Run ID**: 2026-10-05-evidence-resolution-loss-monitor
**Kind**: Phase 3 dogfood observation → contract/adapter gap
**Scope**: text-evidence acquisition, multimodal resolution, and derived cue-cache integrity.

## Generalized observation

An evidence pipeline can acquire a substantial candidate corpus, then produce a
small publishable cue set without a mechanically explainable disposition for
the difference. Treating every non-accepted candidate as absent destroys the
ability to distinguish three different conditions:

1. evidence is invalid (`rejected` with a concrete reason);
2. evidence is valid but unresolved (`uncertain`);
3. evidence was consolidated (`merged` with a target).

The defect is a resolution-loss anomaly, not proof that either OCR or ASR is
globally insufficient. OCR visual text, ASR speech timing/phonetics, and
geometry/region identity are complementary evidence dimensions; none should be
promoted to unconditional truth.

## Contract refinement

Record a named funnel (`raw_ocr`, `usable_candidates`, `fused_candidates`,
`accepted`, `uncertain`, `rejected`, `merged`). A large transition into
`accepted` is suspicious when the remaining candidates have no traceable
disposition. The response is a scope-limited re-probe—region OCR, ASR window,
or frame sample—then re-resolution. It is not a blanket rerun and not a fixed
whole-video coverage threshold.

Coverage is compared against expected evidence: resolved cues over speech
windows and over `subtitle_like` windows. A visually quiet or speech-free
timeline interval is not, by itself, subtitle loss.

## Watermark and cache implications

Sticky text must be classified from spatial persistence, repetition, and visual
features before dialogue resolution. Role-aware projection preserves observed
evidence while preventing watermark/UI material from contaminating the
publishable dialogue text.

Derived cue and coverage artifacts are cache inputs only after schema validation
and atomic publication. Cache provenance must identify source content plus OCR,
ASR, region-profile, fusion-policy, and resolution versions. An invalid or
incompatible artifact triggers a targeted rebuild and cannot be accepted as a
healthy cache hit.

## Linked updates and validation

- Workflow: `text-evidence-multimodal-resolution.md` adds the resolution-loss
  monitor and cache-integrity contract.
- Record: `records/text-evidence.yaml` adds funnel, cache metadata, artifact
  integrity, and corresponding gates.
- Plan companion and this evidence index identify the Phase 3 validation work.

Validation is still pending in the product adapter: produce a fixture with
candidate reduction, verify that all non-accepted candidates remain
traceable, force an invalid derived artifact, and confirm cache reuse is
rejected before a targeted re-probe/rebuild.
