> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-24 - Slot free-spin series may use SPIN + SERIES_FINISH, not BONUS RoundType

Status: candidate

#### One-line Summary

IL2CPP `SlotsRoundType` 枚舉裡即使有 `BONUS`，自然 free-spin **系列**仍可能全程走 `ProcessFreeSpin` + RoundType **`SPIN`**，結束用 **`SERIES_FINISH`**；`EventSink_SlotsBonusRound` 可 hook 卻從不命中。不要把「看到 BONUS／BonusRound」當成 FS 對桌的必要條件。

#### Human Explanation

Agent 常假設「進 feature = RoundType.BONUS 或 BonusRound sink」。實際自然進入後：

1. 觸發當下仍是 `RESULT`（例如 scatter 計數達標）。
2. 局內 free spin 走專用 `ProcessFreeSpin`（或同等 handler），RoundType 名仍是 **`SPIN`**。
3. 系列結束出現 **`SERIES_FINISH`**（及對應 `ProcessSeriesFinish`）。
4. 枚舉值 `BONUS`／`EventSink_SlotsBonusRound` 可能存在但本局從未觀測——視為 **unused or alternate path**，不是「還沒抓到所以缺證」。

對桌／probe 應記錄**實際 handler + RoundType 名**；缺 `BONUS` 時用 `ProcessFreeSpin` + `SERIES_FINISH` 閉環即可。

#### Trigger

- Probe 只等 `BONUS`／BonusRound，自然 FS 整局結束仍標「未確認」。
- 把 base `SPIN`／`RESULT` 與 in-feature `ProcessFreeSpin`+`SPIN` 混成同一狀態機。
- 以為沒有 BonusRound sink 就不能寫 FS wire 文件。

#### Evidence

- Tool: Frida RoundType／handler probe across natural feature enter → in-feature spins → series end.
- Sanitized pattern: enter on RESULT；in-feature hits ProcessFreeSpin+SPIN；end SERIES_FINISH；BONUS enum present, sink never fired.
- Evidence path: `<PROJECT_ROOT>` cabinet `lab/roundtype-live-probe.json`（project-local）。

#### Generalized Lesson

1. **先觀測再命名**：feature wire 以 live handler／RoundType 字串為準，不以枚舉清單臆測。
2. **SPIN 可重入**：同一 RoundType 名可同時服務 base 與 free-spin；用 **handler 名** 區分。
3. **SERIES_FINISH ≠ BONUS**：系列結束常是獨立 RoundType。
4. **未命中的 sink ≠ 未完成**：enum／method 存在但零命中時，標 `unobserved / possibly unused`，改用已觀測路徑閉環。

#### Agent Action

自然進 feature 時同步記 RoundType + Process*／EventSink 名；文件寫「觀測路徑」與「枚舉有但未見」兩欄，勿阻塞在 BONUS。

#### Goal / Action / Validation

- Goal: FS 狀態機可對桌，不誤等 BONUS。
- Action: 自然進 feature → probe → 寫 SPIN／SERIES_FINISH／handler 表。
- Validation: in-feature 至少一次 ProcessFreeSpin+SPIN；結束見 SERIES_FINISH；BONUS 未見則標 unobserved。

#### Applies / Does Not Apply

- Applies: Unity IL2CPP slot cabinets with free-spin **series** handlers.
- Does not apply: cabinets that genuinely enter via BONUS／BonusRound（仍以該柜 live 為準）。

#### Related

- `2026-09-24_085100-slot-ui-harvest-lab-fixture-apply-shim-not-live-payout-forge.md`
- `2026-09-07_172200-slot-s2c-may-apply-via-processspin-gameaction-not-parseresult.md`

#### Validation

- 以自然 feature session 記錄 handler 與 RoundType，確認 `ProcessFreeSpin`／`SERIES_FINISH` 閉環；未命中的 BONUS path 必須標為 unobserved。

#### Promotion Target

- apk-analysis slot capture SOP：RoundType／handler presence table；勿強制 BONUS。

#### Reuse Evidence

- 尚未有獨立 cabinet／專案的重用證據；保留為 candidate，避免把單一 runtime path 視為全域協定規則。

#### Promotion Record

- 尚未 promotion；僅保留為 feedback history candidate。

#### Required Linked Updates

- 已更新 `unity-il2cpp/README.md` category index；尚未改寫 capture SOP，待第二個觀測來源驗證後再評估。
