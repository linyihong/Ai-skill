# Target orthography belongs to derived output

## Observation and cause

Script-family acceptance cannot distinguish Chinese orthographies. A default
Simplified Chinese locale can therefore accept Traditional Chinese model
output or old cached text. Prompt-only changes do not affect cache-hit or
same-language identity paths; prompts inherited from non-Chinese targets may
even contradict the requested Chinese target.

## Adapter boundary

- Resolve target locale and orthography separately from source language.
- Normalize only derived target display/speech text, never raw OCR/ASR,
  region identity, timing, or source-language evidence.
- Exercise fresh realization, cache hit, identity skip, timed pack reload and
  direct caption generation. Old cache bytes need not be destructively migrated.
- Native/runtime orthography conversion must have an explicit availability
  contract and fail visibly, not silently pass through a wrong target script.
- Orthography normalization is not semantic repair, regional vocabulary
  realization, or proof of complete subtitle content coverage.

## Validation scope

Scoped adapter regression verifies conversion idempotence, supplementary
Unicode, non-target-language preservation, provider-result normalization,
old-cache compatibility, unchanged source evidence and actual caption output.
Concrete modules, platform mechanics and run reports remain in `<PROJECT_ROOT>`.
No full real-source translation, TDR/finality or rendered corpus gate is closed.

This supports locale realization and the subtitle-adapter/content-gate seam;
Phase 3 acceptance remains open. NVP caption locale projection consumes the
same boundary without changing evidence classification or OCR sampling.

## Real-provider verification boundary

When local-only realization is requested, availability of cloud credentials
must not silently authorize fallback. Use explicit per-run provider enforcement
without mutating global settings. A basic device probe or complete weights is
not full model-loader readiness: native loading failure must remain a runtime
failure, not an orthography/content verdict. Real-provider translation and
render gates stay open if no target artifact was generated; detailed diagnostics
and environment-specific evidence remain in `<PROJECT_ROOT>`.
