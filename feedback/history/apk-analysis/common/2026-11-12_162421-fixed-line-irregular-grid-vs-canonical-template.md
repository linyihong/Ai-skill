> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-11-12 - Fixed-line irregular page vs canonical panel-grid template

Status: deprecated

Successor: [2026-11-12_171754-canonical-grid-fail-recount-cells-not-invent-layout.md](2026-11-12_171754-canonical-grid-fail-recount-cells-not-invent-layout.md)

#### Reuse Evidence

Deprecated. Do not reuse. Superseding lesson records the corrected mid-course path.

#### One-line Summary

Deprecated: this file wrongly treated printed-line tallies as an irregular panel layout; use the superseding recount-`--cells` lesson instead.

#### Human Explanation

Historical writeup encoded a mistaken rematch branch (uneven columns from printed ids). Kept only as a deprecation pointer so agents do not keep following it.

#### Trigger

- N/A (deprecated)

#### Evidence

- Correction path: superseding lesson + project payline-coordinates anti-pattern table under `<PROJECT_ROOT>/…/docs/slots/`

#### Generalized Lesson

1. Do not follow this file.
2. Follow the superseding lesson: canonical preset miss → re-eye-count rectangular `COLS×ROWS` → `--cells`; do not invent irregular layout modes from printed lineId tallies.

#### Agent Action

- Do: open the superseding lesson.
- Don’t: implement irregular column-row extract flags based on this file.

#### Goal / Action / Validation

- Goal: Prevent continued use of a wrong candidate lesson.
- Action: Status deprecated + pointer.
- Validation or reference source: category README marks deprecated; superseding lesson is candidate.

#### Applies When

- Never (deprecated)

#### Does Not Apply When

- Always — use superseding lesson

#### Validation

- Category index lists this row as deprecated
- Superseding file exists and carries Abstraction Check PASS

#### Promotion Target

- none (deprecated)

#### Promotion Record

Deprecated 2026-11-12; replaced by `2026-11-12_171754-canonical-grid-fail-recount-cells-not-invent-layout.md`.

#### Required Linked Updates

- Category README deprecated row
- Superseding lesson Required Linked Updates

#### Transfer Contract

- Inputs: this deprecated history file path + superseding lesson path from Status/Successor fields
- Mechanism: Status=deprecated + Successor pointer replace the old Generalized Lesson so agents cannot treat uneven printed-line tallies as a legal panel layout mode
- Decision: if Status is deprecated → open Successor only; do not implement irregular column-row extract flags from this file
- Verification: category README marks deprecated; Successor file exists with Abstraction Check PASS; no active Agent Action remains here beyond redirect
- Counterexample: following this file’s historical irregular-column wording after deprecation (would reintroduce the mid-course wrong path)
- Transfer case: any agent that still lands on this filename via old chat/search — must bounce to Successor before rematch tooling changes

#### Intelligence Review

- Decision: deferred
- Reason: deprecated stub; no new intelligence extraction
- Target: none
- Revisit: only if Successor is renamed/deleted or Status is mistakenly set back to candidate — then restore Successor link or re-deprecate
- Discovery: not applicable

#### Abstraction Check

- Swap-project test: PASS — deprecation stub only redirects; no project-bound rule remains active here
- Project evidence kept in: `<PROJECT_ROOT>/…/docs/slots/`
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS
