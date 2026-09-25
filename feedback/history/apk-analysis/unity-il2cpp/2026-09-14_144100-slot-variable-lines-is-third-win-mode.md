> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-14 - Slot winMode needs a third class for variable-height lines

Status: candidate

#### One-line Summary

動態「條賠付線」且每軸符號高度會變時，不要塞進固定 `paylines` 或固定格 `all-ways`；用獨立 `variable-lines`（或同等命名），並禁止沿用固定線的 chipScale / LineData 工作流。

#### Human Explanation

Slot registry 常只分 paylines / all-ways / unknown。第三種常見表面是：HUD 顯示很大的「N 条赔付线」，catalog `linesCount=0`，paytable 寫 VARIABLE LINES，停輪後每軸高度不同，線數 ≈ 各軸高度乘積。

常見誤判：

1. 看到「条赔付线」就標 `paylines`，並假設有靜態 `LineData[N]`。
2. 看到 `linesCount=0` 就標 `all-ways`（與「不要把 0 當成 all-ways」的既有規則衝突，但仍會發生）。
3. 直接複製固定線 cabinet 的 capture → chipScale → `PaylineGrid` 管線。

正確做法：先用 client paytable + 一轉 height product 對 HUD 證明模式，再在 registry / play-rules 開第三 pattern；fixture 用 jagged per-reel symbols，不要硬填矩形 grid。

#### Trigger

- 新 cabinet 的 HUD 是「N 条赔付线」但線數隨每轉變化。
- Paytable 出現 VARIABLE LINES / 任意高度算一顆。
- Live `LinesCount=0` 且無法用 `BetMultiplier / LinesCount` 做 chipScale。
- 想把新 cabinet 塞進既有 complete payline 套件模板。

#### Evidence

- Tool: cabinet play-rules / games-registry winMode taxonomy, spin fixture with per-reel heights.
- Sanitized observation: idle HUD line count ≠ fixed path count; one RESULT showed line count equal to product of six reel heights; registry previously labeled the cabinet `paylines` until reclassified to `variable-lines`.
- Evidence path: `<PROJECT_ROOT>/<App>/docs/slots/` games-registry patterns + one variable-height cabinet package.

#### Generalized Lesson

1. Win-mode taxonomy 至少四格：`paylines` | `all-ways` | `variable-lines` | `unknown`。
2. 「条赔付线」≠ 已證明固定幾何；先問線數是否每轉變、高度是否變。
3. `variable-lines`：`paylineCount=null`（或只存 idleHudLineCount）；`lines=[]`；platform UI 不要預設 `PaylineGrid`。
4. 禁止把固定線 chipScale（`BetMultiplier / LinesCount`）套到 `LinesCount=0` 的 variable cabinet。
5. Fixture：authoritative 是 jagged `symbols[]` / per-reel heights；矩形 `visibleGrid` 只在同一語意契約下使用，不可為通過 schema 而捏造。
6. Worker / desk-match 前先確認 pattern；不要 fork 固定線 worker 當 shortcut。

#### Agent Action

Onboard 時若 paytable 或 height product 證明 VARIABLE LINES，立刻把 knowledge `winMode` 設成 `variable-lines`（或同等），更新 registry patterns 表，並在 package README 寫明「不是 bison / 不是 balloon」。

#### Goal / Action / Validation

- Goal: 後續 session 不會把動態高度線機當成已完成的固定線或 all-ways 模板重用。
- Action: registry enum + play-rules `winMode` + empty/pending geometry file + next-windows 標明 third pattern。
- Validation or reference source: registry `winModeHint=variable-lines`; play-rules `paylineCount=null`; one spin where ∏(heights) equals HUD line count.

#### Applies When

- 授權分析 in-app slots，結果由 server DTO 驅動。
- 需要跨 cabinet 選 capture / platform UI / desk-match 工作流。

#### Does Not Apply When

- 已有 client LineSet dump 證明固定 N 條路徑。
- HUD 明確 ALL WAYS 且格數固定。
- 任務不建 registry / play-rules。

#### Validation

- [ ] games-registry（或同等）含 `variable-lines` pattern 條目
- [ ] 該 cabinet `winMode` / `winModeHint` 不是誤標的 `paylines`
- [ ] package 註明不可重用固定線 chipScale

#### Related Failure Patterns

- 無（分類缺口；補強既有 play-rules 比較 lesson，非新 failure-pattern 檔）

#### Promotion Target

- `feedback/history/apk-analysis/unity-il2cpp/2026-09-08_171200-slot-play-rules-compare-cabinets-without-rng.md`（winMode 比較句應含 variable-lines）— optional follow-up
- project slots `games-registry` patterns table — done in target repo this window

#### Required Linked Updates

- N/A for Ai-skill workflow docs this turn（target registry already updated）
