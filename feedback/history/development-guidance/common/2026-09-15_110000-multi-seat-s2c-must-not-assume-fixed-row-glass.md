Status: candidate

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

#### One-line Summary

既有 lesson 的結論維持於本檔原始內容；此 closure 補記不新增專案事實。

#### Evidence

既有工具輸出、觀察與專案證據已記於本 lesson 的原始段落；未新增或推論額外證據。

#### Generalized Lesson

將本條的具體情境視為候選通則；未在獨立情境重複驗證前，維持 candidate。

#### Agent Action

重用前先核對本條既有前提、限制與驗證方法；前提不符時重新取證。

#### Goal / Action / Validation

- Goal: 保留既有 lesson 的可驗證結論。
- Action: 依原始 Evidence 與 Trigger 重做相關檢查。
- Validation: 結果與原始結論一致才可重用。

#### Applies / Does Not Apply

- Applies: 本條既有 Trigger、前提與 Evidence 相符時。
- Does not apply: 跨專案、前提不明或缺少原始證據時；先重新驗證。

#### Promotion Target

- 尚未 promotion；保留為 candidate history，需有獨立 reuse evidence 才可提升。

#### Required Linked Updates

- Not applicable: 這是歷史 closure 修復；未新增可安全提升的 canonical guidance，category index 已存在。

#### One-line Summary

既有 lesson 的結論維持於本檔原始內容；此 closure 補記不新增專案事實。

#### Evidence

既有工具輸出、觀察與專案證據已記於本 lesson 的原始段落；未新增或推論額外證據。

#### Generalized Lesson

將本條的具體情境視為候選通則；未在獨立情境重複驗證前，維持 candidate。

#### Agent Action

重用前先核對本條既有前提、限制與驗證方法；前提不符時重新取證。

#### Goal / Action / Validation

- Goal: 保留既有 lesson 的可驗證結論。
- Action: 依原始 Evidence 與 Trigger 重做相關檢查。
- Validation: 結果與原始結論一致才可重用。

#### Applies / Does Not Apply

- Applies: 本條既有 Trigger、前提與 Evidence 相符時。
- Does not apply: 跨專案、前提不明或缺少原始證據時；先重新驗證。

#### Promotion Target

- 尚未 promotion；保留為 candidate history，需有獨立 reuse evidence 才可提升。

#### Required Linked Updates

- Not applicable: 這是歷史 closure 修復；未新增可安全提升的 canonical guidance，category index 已存在。
