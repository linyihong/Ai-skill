> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-11-12 - Fixed-line irregular page vs canonical panel-grid template

> **Deprecated 2026-11-12.** This lesson incorrectly blessed printed-line tallies as an irregular panel layout. Successor: [2026-11-12_171754-canonical-grid-fail-recount-cells-not-invent-layout.md](2026-11-12_171754-canonical-grid-fail-recount-cells-not-invent-layout.md). Do not follow the Generalized Lesson below.

Status: deprecated

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。

#### One-line Summary

When a fixed-line paytable page is an irregular column/row panel layout, do not force the canonical dual-label template for that line count; eye-map bijection can ship under a non-rematch geometry status until tooling accepts an explicit `cells-*` stamp for that count.

#### Human Explanation

Mechanical rematch gates often bind a popular line count (e.g. 25) to one dual-label grid preset. Some in-game help books instead print one label per board in uneven columns (different row counts per column). Forcing the preset yields high divider-content scores and cannot `apply`. Agents then either invent OCR paths or leave print lineIds dump-ordered. The recoverable path is: keep GetLineSet for geometry authority, eye-read the irregular page into an eye-map with a full dump bijection, publish lock + client formulas, and leave geometry status at client-proven (or equivalent) until check-grid/apply accept an atypical `cells-*` template for that line count.

#### Trigger

- Fixed payline count whose canonical rematch template is a dual-label NxM grid
- Line page clearly has a different panel rectangle count / uneven column row counts
- `max_divider_content_frac` stays ≫ threshold after offset tuning on the canonical template
- `apply` / pre-commit refuse atypical `cells-*` because line count expects the canonical stamp

#### Evidence

- Tool: panel-grid extract + eye-map bijection against live GetLineSet dump
- Sanitized excerpt: 25-line book with irregular 4-column uneven row counts could not pass canonical dual-label extract; column eye-map still achieved 25/25 dump bijection and shipped under client-proven
- Evidence path: `<PROJECT_ROOT>/pojiepokerist/docs/slots/` cabinet package `assets/paytable-line-audit/` + rematch check-grid docs

#### Generalized Lesson

1. Count **panel rectangles** on the line page first; do not assume lineCount → canonical dual-label template.
2. If the page is irregular, do **not** force the canonical template to “make apply green.”
3. Still produce a print lineId ↔ GetLineSet bijection via eye-map; write package lock + client win tables from that map.
4. Keep rematch-claim geometry status off until check-grid/apply accept an explicit atypical `cells-*` (or per-column edges) for that line count; document the tooling gap in the cabinet next-window.

#### Agent Action

- Do: eye-count panels; extract only when gutters pass; otherwise column/board eye-map → bijection → lock/win sync under non-claim status.
- Don’t: force canonical dual-label cuts, switch to banned OCR-label layout, or claim `paytable-panel-eye` when check-grid cannot pass.
- Don’t: leave failed high-frac `capture/paytable-line-debug/**/crop-boxes.json` on disk — pre-commit may scan the filesystem and fail even if those files are gitignored.

#### Goal / Action / Validation

- Goal: Ship correct print lineId + stroke colors without lying about panel-grid success.
- Action: Irregular-page eye-map + GetLineSet bijection; client-proven (or equivalent) until tooling allows atypical cells for that count.
- Validation or reference source: eye-map maps every dump index exactly once; package analysis gate accepts lock + non-provisional note; check-grid not claimed.

#### Applies When

- Fixed-line cabinets with mechanical rematch gates tied to line-count templates
- Line art uses single-label boards in uneven columns/rows

#### Does Not Apply When

- Page really is the canonical dual-label grid (then use the preset and finish `apply`)
- Win mode is not fixed paylines
- Dump geometry set is incomplete or non-bijective with print art

#### Validation

- Eye-map check reports N/N unique dump coverage
- Canonical-template extract still shows high divider content (proves irregular)
- Commit hooks pass without a rematch claim, or with atypical `cells-*` once tooling allows it

#### Promotion Target

- Project rematch / panel-grid calibration docs under `<PROJECT_ROOT>/…/docs/slots/`
- Optional later: rematch check-grid allow-list for 25-line irregular `cells-*` (tooling change)

#### Promotion Record

尚未 promotion；保留為 candidate history。

#### Required Linked Updates

- Project cabinet docs already record irregular page + client-proven status
- Ai-skill: this lesson only; no enforcement change until a second cabinet hits the same gate
- Category index: update `feedback/history/apk-analysis/common/README.md` if present

#### Transfer Contract

- Inputs: fixed paylineCount, line-page screenshot, GetLineSet dump, rematch check-grid rules
- Mechanism: lineCount→template coupling rejects irregular panel geometry; eye-map can still bijection dump shapes to print ids
- Decision: if panel count ≠ canonical dual-label cells → irregular path (eye-map + non-claim status); else canonical extract→apply
- Verification: bijection N/N; divider frac on forced canonical template remains high; hooks pass without false rematch claim
- Counterexample: true dual-label 25-line books that pass regular-5x4 with frac≈0
- Transfer case: any fixed-line help book with uneven column stacks (e.g. 8+6+6+5) under a lineCount that tooling still maps to a dual-label preset

#### Intelligence Review

- Decision: deferred
- Reason: needs a second irregular fixed-line cabinet or an explicit rematch tooling change before promoting into workflow
- Target: none
- Revisit: when rematch check-grid gains an allow path for atypical `cells-*` at popular line counts, or a second cabinet hits the same refusal
- Discovery: not applicable (deferred)

#### Abstraction Check

- Swap-project test: PASS — any fixed-line rematch stack that couples lineCount to a dual-label grid preset will block irregular single-label pages the same way
- Project evidence kept in: `<PROJECT_ROOT>/pojiepokerist/docs/slots/`
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS
