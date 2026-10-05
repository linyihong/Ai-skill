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

Adapter validation now covers atomic cue/coverage publication and cache reuse:
the fixture confirms that invalid cue integrity or invalid coverage JSON is
rejected rather than treated as a healthy cache hit. The adapter emits the
named funnel and marks an unexplained contraction as `targeted_required`.

The adapter now runs a bounded window OCR/re-resolution pass with before/after
counts, preserving failed or unresolved observations. Fixtures distinguish
explained duplicate reduction from missing dispositions, unrelated texts with
shared timestamps, quiet timeline gaps, and failed acquisition.

## Attribution and eligibility findings

Lexicon segmentation can change the number of parts while detection boxes
remain unchanged. A positional pairing fallback that copies the aggregate
line into each box gives watermark text a dialogue box's geometry. Box-local
raw text must retain ownership regardless of segmentation count. The region
contract and a mixed watermark/dialogue fixture now capture this requirement.

A second gap is status projection: a nested unresolved or rejected decision
cannot become publishable through an absent/stale outer status. Recovery must
pass the same finalization gate as initial resolution, including episode-level
role projection; an increased cue count alone is not recovery evidence.
An accounted `merged` observation can still point to an uncertain target.
Retry outcomes must inspect the target's final status after eligibility gates;
accounting closure alone must not be reported as recovered dialogue.

Another attribution hazard is conversion through a text/time-only summary or
replaying raw boxes over already-projected evidence. Both can discard region
identity and source-level persistence, making a short retry incorrectly
rehabilitate sticky or scene text. Preserve box-local metadata across retries
and require attributable dialogue support at finalization. Source-local changing
text bands can support mid-frame captions; short retries must not establish
that source profile by themselves. Ambiguous roles remain uncertain.
Bounded retry scheduling should prioritize attributable dialogue losses over
unresolved role noise; preserve the latter for review without starving the
actual omission windows. Source-profile matching must tolerate box jitter.

Validation remains partial: scoped acquisition executes and preserves unresolved
observations, but accepted-content quality and remaining loss must be checked
independently. No complete subtitle coverage or publish-ready claim is made
from fixture passes or increased cue counts.
