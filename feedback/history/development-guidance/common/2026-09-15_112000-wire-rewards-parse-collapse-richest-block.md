Status: candidate

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
