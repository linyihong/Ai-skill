# Bison Bash payline geometry — live GetLineSet dump

## Lesson
Do not invent remaining payline paths from paytable OCR alone when Frida can call `SlotsInfo.GetLineSet("")` on a live `SlotsBuffaloRush` instance. Struct array `SlotsLineData[]` is inline (stride 16): `string ID` + `SlotsLinePoint[] Points`. Numeric line IDs are array order 1..25.

## Chip scale
Live `BetMultiplier=50`, `LinesCount=25` → `chipScale=2`.  
`lineWin = tableMult × lineBet × wildFactor × chipScale` matches server `@PAYOUT` on regular payline spins.

## Corrections
Early spot-crop map swapped lines 1/2 vs live order and mislabeled `[1,0,0,0,1]` as line 21 (live id 6). Path `[1,0,1,0,1]` is not in the live set.

## Tooling
`pojiepokerist/scripts/frida/bison_dump_lineset.js`
