Status: candidate

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
