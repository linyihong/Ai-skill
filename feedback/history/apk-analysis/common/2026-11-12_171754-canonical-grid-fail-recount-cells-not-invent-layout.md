> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-11-12 - Canonical grid fail → re-count rectangular cells (do not invent layout)

Status: candidate

Supersedes: [2026-11-12_162421-fixed-line-irregular-grid-vs-canonical-template.md](2026-11-12_162421-fixed-line-irregular-grid-vs-canonical-template.md)（該條把「印刷線號→不規則欄」誤寫成合法路徑，已 deprecated）

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。同日使用者糾正：canonical 預設失敗後改眼數矩形 `COLS×ROWS`（多標號同框）即過 divider，無需新 layout 旗標。

#### One-line Summary

When the line-count canonical panel-grid preset fails, stop and re-eye-count a rectangular `COLS×ROWS` of panel frames (multi-label = one cell); do not count printed lineIds into uneven columns or invent a new extract layout mode.

#### Human Explanation

Rematch tooling often maps a popular payline count to one dual-label preset. Agents treat that map as the grid answer. When extract divider scores stay high, the failure feels like “this page is irregular,” especially if counting printed numbers down each column yields uneven stacks. Project anti-patterns already ban that move: printed line count ≠ panel count; multi-label boards share one rectangle. The recoverable branch is always the same toolkit—`--cells COLSxROWS` (plus force when it conflicts with the preset) and rectangle knobs / edges—not a new per-column layout primitive. Mid-course correction must fire **before** writing new CLI flags or encoding “irregular” into feedback/docs.

#### Trigger

- Canonical template for a fixed paylineCount fails `max_divider_content_frac` after ordinary offset/pad tuning
- Impulse to explain the page as uneven column stacks from printed lineId tallies
- Impulse to add a new extract flag / layout mode / stamp family to fit that story
- Existing docs already say: after `COLS×ROWS` is chosen, only tune rectangles

#### Evidence

- Tool: panel-grid extract comparison (canonical preset vs rectangular `--cells`)
- Sanitized excerpt: same line page rejected the lineCount dual-label preset; rectangular multi-label cells cleared divider≈0 and matched eye-read panel crops
- Evidence path: `<PROJECT_ROOT>/…/docs/slots/` payline-coordinates + panel-grid-calibration anti-pattern tables

#### Generalized Lesson

1. `paylineCount → canonical template` is a **default**, not an observation. Observation = eye-count panel **frames**.
2. Printed labels per column (8+6+…) measure line art, not the panel lattice. Multi-label frames collapse several ids into one cell.
3. When the preset fails: **re-count** `COLS×ROWS` → `--cells` (+ force if needed) → tune `--dx/--dy/--gutter-pad/--band-frac`／edges. Do **not** invent irregular column-rows tooling.
4. **Mid-course stops** (must halt before new code/docs):
   - Before adding extract CLI / stamp vocabulary: re-read the project anti-pattern table for “irregular columns / lineId factorization.”
   - After one failed offset-tuning round on a hypothesis: try the next rectangular `COLS×ROWS`, do not deepen the same wrong model.
   - Divider still high on a brand-new layout story: treat the **story** as wrong, not the threshold.
5. Do not write feedback lessons that bless the irregular story until a rectangular `--cells` path has been tried and shown impossible on annotated crops.

#### Agent Action

- Do: on canonical fail → open annotated page → count gutter lattice → `--cells` → check-grid → `panel-line-map` for multi-label cells → apply.
- Don’t: tally printed ids into uneven columns; add `--column-rows`-class flags; promote “irregular page” into Ai-skill while the rectangular path is untested.
- Don’t: use banned OCR-label layout to “see” the grid; clarity comes from annotated gutters + `panels/rXcY.png`.

#### Goal / Action / Validation

- Goal: Keep rematch on one rectangular panel-grid contract; correct wrong hypotheses early.
- Action: Mid-course halt checklist above; prefer `--cells` over new layout modes.
- Validation or reference source: alternate rectangular extract reaches divider ≤ threshold; panel crops show multi-label groups; no new layout flag required.

#### Applies When

- Fixed-line rematch with lineCount-coupled presets
- Help-book pages that pack multiple print ids into shared panel frames

#### Does Not Apply When

- Annotated crops already prove a true non-rectangular board set after rectangular candidates were tried
- Win mode is not fixed paylines

#### Validation

- At least one rectangular `--cells` attempt documented before any “irregular” claim
- Divider pass on the chosen `COLS×ROWS` without new layout primitives
- Feedback/docs do not encode the rejected uneven-column story as policy

#### Promotion Target

- Project `docs/slots/payline-coordinates.md` / `panel-grid-calibration.md` (already state the ban; keep as authority)
- Optional: rematch extract help text — “canonical fail → try `--cells`, do not add layout modes”

#### Promotion Record

尚未 promotion；保留為 candidate history。

#### Required Linked Updates

- Deprecate `2026-11-12_162421-fixed-line-irregular-grid-vs-canonical-template.md`
- Category index README row update
- Project cabinet notes that still say “irregular N-column stacks” should be corrected to rectangular `cells-*` when rematch resumes (project-local)

#### Transfer Contract

- Inputs: paylineCount, line-page image, canonical preset result, anti-pattern docs
- Mechanism: agents confuse lineCount defaults and printed-id tallies with panel lattice geometry, then escalate into new tooling
- Decision: preset fail → re-count rectangular cells → `--cells`; invent layout only after rectangular hypotheses are exhausted on annotated crops
- Verification: divider ≤ threshold on `--cells`; multi-label map covers all print ids; no new layout flag
- Counterexample: genuine non-grid art where every rectangular candidate cuts through boards on annotated crops
- Transfer case: any rematch stack that auto-picks a dual-label preset from line count

#### Intelligence Review

- Decision: deferred
- Reason: single corrected incident; project docs already carry the ban—lesson records the mid-course/tooling-escalation failure
- Target: none
- Revisit: if a second agent repeats “new layout flag” after canonical fail
- Discovery: not applicable (deferred)

#### Abstraction Check

- Swap-project test: PASS
- Portability note: any pipeline that binds an ID count to a grid preset hits the same recount-versus-invent fork after the preset misses gutters
- Project evidence kept in: `<PROJECT_ROOT>/…/docs/slots/`
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS
