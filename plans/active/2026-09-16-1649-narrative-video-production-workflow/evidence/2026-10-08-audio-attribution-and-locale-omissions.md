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
