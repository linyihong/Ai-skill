> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-14 - Slots Rewards chips are per-way; Collapse means cascade

Status: candidate

#### One-line Summary

`SpinInfo.Rewards` 的 `SlotsRewardChips` 是 **每一 way 一筆**（Amount=該長度 paytable chip）；`Σ Amount == Payout`。同袋出現 `SlotsRewardCollapse`（含 NewSymbols）代表付完後會 tumble，不能只靠單一 Symbols 快照解釋所有局的 Payout。

#### Human Explanation

Variable-lines ways 公式（最長 L→R × tableMult × TB）在單階段局與 Payout 一致。Rewards 字典（`Dictionary<int,SlotsReward[]>`）用 valuetype Entry 陣列走訪後可見：chips 筆數等於 ways，每筆 Amount 等於該長度的 chip。Collapse 獎勵帶下一輪填入符號，證明 cascade。若 hook 只 `once` 抓第一個 ProcessResult，多段 cascade 的累計 Payout 可能大於該快照的 ways。

#### Trigger

- Desk-match 多數 exact，少數 Payout > ways(Symbols)。
- Need to explain Rewards vs ways without rewriting paytable.

#### Evidence

- Tool: Frida RESULT Rewards walk (Entry stride 0x18) + ways recompute.
- Sanitized observation: one spin had N chips each equal to length-3 chip and N==ways; another had Collapse.NewSymbols count > 0 after chips; Σ chips matched Payout.
- Evidence path: `<PROJECT_ROOT>/<App>/docs/slots/<cabinet>/win-calc.md` Rewards shape section.

#### Generalized Lesson

1. Walk `Rewards` as valuetype Dictionary entries (not Object[]).
2. Treat `SlotsRewardChips` as per-way; verify `count(chips)==ways` and `Σ Amount==Payout`.
3. `SlotsRewardCollapse` ⇒ plan multi-RESULT / stage capture; do not force single-grid ways to equal multi-stage Payout.
4. Keep longest-ways formula; extend capture, not invent phantom pays.

#### Agent Action

Dump Rewards+Collapse detail on RESULT; compare chip sum and ways; if Collapse present and Payout later diverges, capture the full ProcessResult chain.

#### Goal / Action / Validation

- Goal: reconcile ways, Rewards, and cascade with Payout.
- Action: fix dict walk; document chips-per-way + Collapse.
- Validation: ≥1 spin with chipSum==Payout==waysAmount and Collapse observed.

#### When to Use / When Not to Use

- Use: IL2CPP slots with `SpinInfo.Rewards` dictionary.
- Not: cabinets without Rewards / without tumble.

#### Related

- `2026-09-14_150000-slot-variable-lines-ways-desk-match.md`
- `2026-09-14_160500-slot-ways-payout-gap-dump-state-rewards-first.md`

#### Applies / Does Not Apply

- Applies: 本條既有 Trigger、前提與 Evidence 相符時。
- Does not apply: 跨專案、前提不明或缺少原始證據時；先重新驗證。

#### Validation

依本條既有 Evidence 重做對應觀察或檢查；結果與原始結論一致才可採用。

#### Promotion Target

- 尚未 promotion；保留為 candidate history，需有獨立 reuse evidence 才可提升。

#### Required Linked Updates

- Not applicable: 這是歷史 closure 修復；未新增可安全提升的 canonical guidance，category index 已存在。

#### Applies / Does Not Apply

- Applies: 本條既有 Trigger、前提與 Evidence 相符時。
- Does not apply: 跨專案、前提不明或缺少原始證據時；先重新驗證。

#### Validation

依本條既有 Evidence 重做對應觀察或檢查；結果與原始結論一致才可採用。

#### Promotion Target

- 尚未 promotion；保留為 candidate history，需有獨立 reuse evidence 才可提升。

#### Required Linked Updates

- Not applicable: 這是歷史 closure 修復；未新增可安全提升的 canonical guidance，category index 已存在。
