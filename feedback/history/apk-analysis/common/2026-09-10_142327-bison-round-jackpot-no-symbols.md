# Bison Bash ROUND RESULT may be jackpot-only (no Symbols)

## Lesson
Desk `EVENT TYPE="ROUND"` RESULT packets can carry `SPIN/@PAYOUT` with
`<REWARDS><JACKPOT POSITIONS="reel:row,…" AMOUNT="…"/>` and **no**
`SlotsSymbolData` grid. Ops must not treat empty payline analysis on a
stale/associated visible_grid as “missing paylines.”

## Evidence
Spin #48: PAYOUT=18000, JACKPOT LEVEL=1 with 7 POSITIONS; RESULT has no
symbol grid. Heuristic `roundType` may show `ROUND` from EVENT/@TYPE.

## Ops UI
Show JACKPOT reward / positions separately from payline strokes when
present; flag records whose RESULT lacks Symbols.
