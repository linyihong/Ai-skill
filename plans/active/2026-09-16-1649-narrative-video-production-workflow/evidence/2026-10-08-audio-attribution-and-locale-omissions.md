# Audio fallback attribution and locale omission accounting

## Cause and adapter boundary

A fusion adapter can return audio observations when visual evidence is sparse.
If downstream treats every fused text as observed OCR, it manufactures visual
candidates/regions, then a visual-role gate can erase valid audio-only dialogue.
Capture modality provenance at the producer boundary before cleaning copies
objects. Do not infer origin from text equality or invent visual ownership.
Audio validity and resolution still apply; missing visual evidence cannot be
used as the sole rejection reason for an explicitly audio-only input.
Invalid repeated audio remains retained but unaccepted. Genuine OCR remains
subject to visual attribution and scene-text precision controls.

## Locale projection must account for every selected cue

A nonempty translation cache and a completed worker can still yield missing
target captions. Distinguish absent realization from target-script rejection,
semantic rejection and projection failure. Retain source identity/time, candidate
text and reason for every unprojected selected cue; accepted rows need a target
reference. Do not relax all residue guards simply to restore cue counts.
Reprojection with no usable captions must not retain an older timed list as
the apparent new output. Phrase candidates can remain for review/recovery.
Per-cue validation must receive the corresponding source, not a stale loop value.

## Cached rejection must remain recoverable

Nonempty phrase entries are candidates, not permanent admission decisions.
Recheck target-locale script/residue and existing Finality restrictions at both
episode scheduling and individual cache reuse. A rejected hit must reach the
configured realization actor; a valid hit must remain reusable without a model
call. Preserve the previous candidate, rejection reason and whether a repair
was attempted in the selected cue's projection disposition, including after
successful replacement. Failed repair remains an explicit omission.
Short cache lookup responses must not silently remove the remaining sources.

Paired fixtures reproduce the two cache-bypass defects before repair and verify
valid-hit reuse, retry routing, failed-retry retention and lookup cardinality.
Read-only frozen-cache replay verifies scheduling, not improved translation.
Actual realization needs a separate bounded local-only child, isolated copied
cache, native-exit retention and immutable baseline checks. Target-script PASS
does not establish entity identity, event equivalence or full semantic Finality.

Admission reporting also requires nonempty target realization. A permissive
script helper may legitimately return true for an empty string; the consumer
must not count that as a caption restored. Paired result-auditor fixtures keep
empty outputs, rejected nonempty outputs, runtime failures and pending positions
separate. A successful parent/native exit reports execution only. Deterministic
retranslation can reproduce rejected candidates; retry scheduling is not recovery.
Preserve failed realization evidence and inspect actor output versus sanitization
before another attempt, rather than relaxing guards or repeatedly replaying the
same selection request. Target realization/content acceptance remains open.

## Verification and limits

Paired fixtures verify audio-only builder wiring, invalid audio non-promotion,
scene-text exclusion, disposition persistence, empty reprojection and per-cue
source association. Frozen real evidence replay verifies the provenance defect
without mutating baseline evidence or calling a realization provider.
Audio-stream presence/decoded volume is not proof of intelligible dialogue.
Raw observation disposition completeness is not final dialogue recall; generated
caption artifacts are not content/timing/layout acceptance.

Supports scoped consumer adapter repair only. Target realization quality,
bounded reacquisition, independent audio/visual review and rendered acceptance
remain open. Concrete sources, counts, logs and code live in `<PROJECT_ROOT>`.
Existing acquisition/resolution/caption contracts were checked; no new actor,
fixed glossary, workflow stage, lifecycle definition or corpus PASS is introduced.
