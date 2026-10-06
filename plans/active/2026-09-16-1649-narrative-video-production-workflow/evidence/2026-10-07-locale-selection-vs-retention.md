# Locale selection is not metadata retention

Status: scoped adapter regression confirmation; full-chain gates remain open.

## Finding

Preserving all region identities, temporal groups and alignment references on
serialization does not ensure the selected display text covers those identities.
A first-match target-language region selector can replace a complete translated
speech cue with one partial observation. Reordering regions can then change
which sentence reaches the caption artifact without changing the source evidence.

## Bounded adapter response

When distinct target-language readings exist in a flat region collection,
decline the single-region override and retain the existing whole-cue translation
path. This is not fixed OCR precedence or permission to concatenate observations.
No complete translation means no authority to invent or join dialogue. Preserve
the observations and their original temporal identities for later resolution.
Repeated identical readings and an unambiguous region remain control cases.

## Verification boundary

Exercise the real locale projection, serializer, reload and caption writer with
temporary synthetic packs. Check both expected sentence bodies downstream,
unchanged group/region/alignment metadata, no source mutation and order invariance.
Pair the omission case with single-region, repeated-reading and bilingual-language
controls. Cache equality alone is not the final consumer assertion.

The earlier plural-group write/reload loss has scoped repaired regression
evidence; this newly exposed display-selection defect required another paired
regression. Neither repair proves real-source semantic correspondence, temporal
group scheduling, layout comfort, rendered-output quality or Phase 3 acceptance.
Concrete source hashes, test names and host run records remain in the consumer
project. No new workflow route, agent, sampling rate or timing authority is added.

Maps to [resolution companion](../43-multimodal-text-evidence-resolution.md).
