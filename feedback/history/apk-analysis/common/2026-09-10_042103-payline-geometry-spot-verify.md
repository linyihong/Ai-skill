# Lesson: Payline geometry — verify paths, do not invent full maps

Policy refs: [`enforcement/README.md`](../../../../enforcement/README.md), [`feedback/feedback-lessons.md`](../../../feedback-lessons.md).

## Goal

When documenting fixed-payline slot rules from a client paytable screenshot or HUD art, record the **win formula** and only those line paths that are **spot-verified**. Do not publish a complete N-line coordinate table from a single multimodal “likely paths” pass.

## Action

1. Separate **algorithm** (L→R, wild substitution, `tableMult × lineBet × wildFactors`, server `Payout` authority) from **geometry** (row index per reel per line id).
2. For geometry: crop or OCR each numbered diagram; store `rows[reel] = visibleRow` only when the label and lit cells are clear.
3. Mark status as `partial-…` until every line id is verified; keep an explicit `unverifiedLineIds` list.
4. Client-visible symbol multipliers may be transcribed as **formula inputs** with a server-authority caveat; do not treat them as RTP reconstruction or store progressive jackpot bank amounts.

## Validation

- Spot-check at least two non-horizontal lines against labeled crops.
- Reject any auto-generated full map that disagrees with a spot-check.
- Schema / package status must not claim `client-proven` geometry until all paths are verified.

## Applies / does not apply

- Applies: payline cabinets documented from UI art or screenshots.
- Does not apply: all-ways games; live `LineSet` XML dumps that already list every path.

## Status

`stable`
