> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-14 - Ways≠Payout triage: dump StateInfo + Rewards before revising formula

Status: validated

#### One-line Summary

Variable-lines longest-ways 對上多數 `SpinInfo.Payout` 後，若出現穩定 under-count，先 dump `SlotsStateInfo`（series/activity/extra/mult）與 `SpinInfo.Rewards`，不要立刻改 ways 公式或 paytable 倍率。

#### Human Explanation

同一 RESULT 物件上的 jagged symbols + Bet/Payout 已足夠證明 baseline ways。少數局 Payout 明顯大於 ways 重算時，差額可能來自活動／系列／額外籌碼或 `Rewards` 明細，而不是 L→R ways 定義錯了。先把 StateInfo 與 Rewards chips amount 記進 fixture，再決定是否擴公式。

#### Trigger

- `ways × tableMult × (Bet×BetMultiplier)` 對多數樣本 exact，但少數樣本 Payout 更高。
- Paytable chip 表已對過仍對不上。

#### Evidence

- Tool: Frida on RESULT — SpinInfo Bet/Payout/Rewards；player StateInfo series/activity/extra/multiplier.
- Sanitized observation: several exact matches coexisted with two under-counts whose gaps were not explained by empty StateInfo on later zero-win dumps; Rewards walk still incomplete.
- Evidence path: `<PROJECT_ROOT>/<App>/docs/slots/<cabinet>/win-calc.md` open-mismatch table.

#### Generalized Lesson

1. Keep longest-ways as baseline once ≥2 exact samples exist (with and without WILD).
2. On under-count: record StateInfo + Rewards before changing tableMult or adjacency rules.
3. Do not invent phantom symbol pays from the gap arithmetic alone.
4. Mark open mismatches in package next-windows; do not silently “fix” by discarding fixtures.

#### Agent Action

Extend RESULT hook to emit StateInfo + Rewards alongside Bet/Payout; persist under fixture `deskMatch`; recompute ways and compare gap to state extras.

#### Goal / Action / Validation

- Goal: triage Payout gaps without corrupting a proven ways formula.
- Action: dump StateInfo/Rewards; document open mismatches.
- Validation: exact samples remain exact; gaps either equal state extras or stay listed open.

#### When to Use / When Not to Use

- Use: variable-lines (or ways) desk-match after a working baseline.
- Not: first-ever desk-match before TB/ways proof.

#### Related

- `2026-09-14_150000-slot-variable-lines-ways-desk-match.md`
