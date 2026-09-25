Status: candidate

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
