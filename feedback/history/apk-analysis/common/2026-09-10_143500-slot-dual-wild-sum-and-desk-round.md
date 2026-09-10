# Lesson: dual mult-wild sum + desk ROUND without symbols

Refs: `enforcement/README.md`, `feedback/feedback-lessons.md`

## Goal

When reconciling slot payline estimates against live captures on a shared
table, separate (a) own-spin geometry math from (b) feature payouts and
(c) other-player desk notifies that lack a symbol grid.

## Action

1. **Dual outer multiplier wilds on one payline:** if two participating
   factors appear on the run (typically reel ends), treat the combined
   factor as the **sum**, not product or max. Validate by checking that
   only the sum closes the server total when other single-factor lines
   are held fixed.
2. **Flash / progressive banks:** diamonds on the configured middle reels
   at/above the trigger count can add a bank amount **on top of** paylines.
   Bank chip totals are progressive — record observed remainders as
   evidence, not as a fixed paytable row.
3. **Shared-desk `ROUND` / RESULT with chips only:** a notify that carries
   `PAYOUT` + chip reward but **no symbols** is often another seat’s
   outcome. Do not attach a stale local grid or treat payline delta as a
   missing-line bug. Prefer filtering by own player id before reconcile.
4. **Capture hygiene:** if the parser reads *all* frames after bet but the
   DB stores only the first frame, jackpot / symbol XML may be lost while
   the grid still looks populated — keep full frame sets for audits.

## Validation

- Dual-factor hypothesis matrix: product / max / sum vs server total with
  other lines unchanged; only one mode should match.
- Flash remainder: same diamond-count tier, same line bet, payline sum
  subtracted — remainder should cluster if the bank is slow-moving.
- Desk ROUND: decrypt stored body; absence of symbol tokens + foreign
  player id ⇒ exclude from payline exactness stats.

## Applies / does not apply

- Applies: fixed-line slots with end-reel mult wilds; multi-seat desk feeds.
- Does not apply: all-ways games; single-player isolated captures with full
  RESULT XML already stored.

## Status

`validated` (dual-wild sum + desk ROUND triage). Flash remainder amounts
remain `experimental` (progressive).
