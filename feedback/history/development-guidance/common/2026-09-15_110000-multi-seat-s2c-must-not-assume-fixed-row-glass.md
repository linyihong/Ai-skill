# Multi-seat S2C must not assume fixed-row glass

> References: `enforcement/README.md`, `feedback/feedback-lessons.md`

## Problem

A desk notify carries every seat in one S2C body. A parser that always densifies glass into a fixed visible-row matrix (e.g. rows `0..2` outer) will:

1. Drop taller cells when a cabinet uses variable reel heights.
2. Store outer length = row count, so a UI that expects reel-major jagged glass shows only the first few X columns filled and pads the rest empty.
3. If RESULT+NEXT merge then copies donor `PAYOUT`/`CHIPS` whenever primary payout is `0`, a seat-scoped RESULT can inherit another seat’s desk-wide NEXT money.

UI “fixing” cannot repair this: the stored grid orientation and payout attribution are already wrong.

## Rule

1. Detect jagged / variable-height glass from protocol points (`maxRow > fixedRows` or reel count above the fixed payline width) and densify **reel-major** (outer = reel L→R, inner = top→bottom). Keep fixed row-major only for cabinets that truly use a constant visible row band.
2. When merging RESULT with NEXT, take donor payout only if the donor parse was seat-scoped for the same player (or neither side was scoped). Never borrow desk-wide NEXT chips onto a seat-scoped zero RESULT.
3. After changing orientation, reparse stored spins that hydrated `grid_cells` from the old shape; prefer rebuild from symbols/`visible_grid`, not stale cells.

## Validation

- Fixture with six reels and heights taller than three: outer length = reel count; product of reel lengths matches HUD ways product.
- Multi-seat RESULT `PAYOUT=0` + foreign NEXT `PAYOUT>0` with seat scope: merged payout stays `0` and notes skip foreign NEXT.
- Fixed 3×N payline glass regression still parses row-major and `full` when complete.

## Applicability

Applies to multiplayer desk S2C / notify parsers and any downstream ways/collapse UI that assumes reel-major jagged boards. Does not replace cabinet-specific reward-level Collapse wiring.
