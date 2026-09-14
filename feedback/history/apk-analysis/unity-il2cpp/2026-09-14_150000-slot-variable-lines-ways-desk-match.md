> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-14 - Variable-lines desk-match uses Bet×BetMultiplier and longest ways

Status: validated

#### One-line Summary

`LinesCount=0` 的 variable-lines cabinet：`totalBet = SpinInfo.Bet × BetMultiplier`；一般中獎用「最長連續 L→R ways × tableMult × totalBet」對 `SpinInfo.Payout`，不要用固定線 chipScale。

#### Human Explanation

固定線常用 `chipScale = BetMultiplier / LinesCount`。動態高度線櫃 `LinesCount` 常為 0，HUD 总下注仍可由 `Bet × BetMultiplier` 還原。賠付不是靜態 LineData：對單一符號取從左起最長連續匹配長度 L（≥3），ways 為各軸匹配個數乘積，再乘 paytable 相對总下注倍率。同一條 ways 的較短前綴不要再加一次。

#### Trigger

- Catalog / live `LinesCount=0` 且 HUD 顯示動態「条赔付线」。
- Paytable 金額標註相對某一 totalBet。
- 需要把 RESULT `Payout` 對回 client 賠表。

#### Evidence

- Tool: Frida RESULT SpinInfo Bet/Payout + per-reel symbol counts.
- Sanitized observation: one sample had Bet×BetMultiplier equal HUD total bet; four ways of a length-4 mid-tier symbol matched Payout exactly when using paytable ×TB and omitting shorter prefixes.
- Evidence path: `<PROJECT_ROOT>/<App>/docs/slots/<cabinet>/win-calc.md` ways desk-match section.

#### Generalized Lesson

1. Variable-lines：先證明 `totalBet = Bet × BetMultiplier`（對 HUD），再談 ways。
2. Ways：最長 L→R；`ways = Π counts`；`payout = ways × tableMult(L) × totalBet`。
3. 不要套用 `BetMultiplier / LinesCount`。
4. 第二樣本再驗證 WILD 是否計入 count；多符號同時中獎要分開加總驗證。

#### Agent Action

Hook RESULT 時同時 dump Bet、Payout、jagged symbols；先算 TB 與 ways，再寫 win-calc desk-match。

#### Goal / Action / Validation

- Goal: 可重算一般中獎而不誤用固定線公式。
- Action: document formula + one exact sample in cabinet win-calc.
- Validation: recompute == Payout on at least one non-feature spin.

#### Applies When

- 授權分析 in-app slots；結果 DTO 含 Bet/Payout/Symbols。

#### Does Not Apply When

- 已證明固定 LineSet + chipScale。
- 只做 art map、不做賠付對帳。

#### Validation

- [ ] win-calc 寫明 TB 與 ways 公式
- [ ] 至少一筆 recompute == Payout

#### Related Failure Patterns

- 無

#### Promotion Target

- N/A this turn（cabinet docs already updated）

#### Required Linked Updates

- N/A
