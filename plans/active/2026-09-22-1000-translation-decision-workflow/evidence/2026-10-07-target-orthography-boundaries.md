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

## Native-loader isolation

Reusable candidate: [known-case lookup and revalidation boundary](../../../../feedback/history/development-guidance/common/2026-10-07_125417-native-model-load-known-case-first.md).

Capture native failures in an owned child with parent-retained exit status and
stack logs: an in-process Python exception handler cannot guarantee evidence
retention after a native fault. A controlled read-method comparison can isolate
weight materialization from device initialization or translation. Complete
loading after that change still does not establish generation or target quality.
Keep compatibility adapters bounded to the intended weights, restore temporary
integration hooks, and test unrelated-file preservation. A whole-file eager read
is not equivalent to bounded lazy reads when memory headroom is limited.

## Realization versus semantic acceptance

A bounded local-only run subsequently produced derived target text, a timed pack,
captions and a burned artifact. Mechanical orthography, caption text retention
and immutable source checks passed. This closes only that adapter execution
slice: source-subtitle compression versus acoustic evidence still requires
semantic review, and overlaying a new locale on existing hard captions does not
prove production layout. Do not promote ASR to automatic truth, repair one line
with a hard-coded replacement, or mark full-source content/timing/render PASS.

## Provider enforcement and source-authority verification

Local-only selection must constrain empty-output retry, quality fallback and
validation repair, even when cloud credentials are configured. Positive tests
for explicitly authorized automatic fallback complement forbidden-provider
negative tests; clearing a key in a diagnostic is not a production guarantee.

Distinguish OCR-only translation diagnostics from resolved spoken-pivot identity
delivery. Compare the actual translator input and provenance before attributing
acoustic detail loss to translation. An independent identity regression must
retain resolved content and observed evidence without translating raw ASR or
compressed OCR again. Incompatible frozen evidence must remain rejected; use
a compatible rebuild rather than rebranding cache versions. A generic semantic
prompt guard that leaves the observed failure unchanged is not a repair and
must not be deployed or counted as semantic acceptance. Full gates remain open.
