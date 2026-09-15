# Wire REWARDS: parse Collapse from richest block only

> References: `enforcement/README.md`, `feedback/feedback-lessons.md`

## Problem

Desk S2C often embeds multiple `<REWARDS>` copies (seat chrome, stages, duplicates). Scanning the whole body for `<CHIPS>` doubles chip totals and breaks cascade desk-match. Variable-height cabinets also send `<COLLAPSE POSITIONS LEVEL><NEW SYMBOLS="id:reel:row,…"/></COLLAPSE>` that a chips-only parser drops — UI then cannot show explode cells or refill, and may treat the final glass as the pre-explode board.

## Rule

1. Prefer a single richest `<REWARDS>…</REWARDS>` block (score by CHIPS + COLLAPSE + JACKPOT counts); do not concatenate every block.
2. Parse leveled CHIPS (`AMOUNT`, `LEVEL`, `POSITIONS`) and COLLAPSE (`POSITIONS`, `NEW SYMBOLS`) into structured rewards.
3. When protocol glass is the **final** stop board, reverse-tumble with COLLAPSE NewSymbols + remove positions to recover board-before each Way, then forward-sim for pays; always render a labeled final glass.
4. On **multi-level** reverse, after each peeled level run chip-position fill (unique paytable symbol, else wild). Do not only fill the outermost Way — inner explode cells otherwise leak as `"?"` onto earlier boards and post-tumble `gridAfter`.

## Validation

- Paid multi-level spin: Σ chips == SPIN payout; each Way exposes remove-cell highlights; final heights match glass; **no `"?"` cells** on any Way / post-tumble board (ambiguous tiers may show wild).
- Zero-pay spin: no Way boards required; final glass still shown.
- Fixed payline cabinets still parse single CHIPS/JACKPOT without COLLAPSE.

## Applicability

Multiplayer desk S2C parsers and cascade / variable-lines bet-record UI. Not a substitute for Frida deskMatch fixtures that already supply initial boards.
