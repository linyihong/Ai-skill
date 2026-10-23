> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-23 - Wire RESULT GLASS must match ProcessResult strip window

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。來源為同一類「AES game-TCP session-proxy rewrite + Unity ProcessResult strip」lab；具體柜机／样本留在專案 lab 筆記。

#### One-line Summary

用 session-proxy／replay 強制 RESULT 時，GLASS 符號數必須對齊該柜 `ProcessResult` 的 `returnedRowRange`；只發可見窗（例如 5×3＝15）常會讓客戶端卡在停輪態，後續 BET 發不出。

#### Human Explanation

許多柜机的 live RESULT 不是「可見 5×3」而是帶上下緩衝列的 strip（常見 Row −2..4 → 35 cells）。Lab fixture 常只存 `visibleGrid` 方便 desk-match／對獎，若 builder 原樣編成 GLASS，wire 形狀就與 runtime `ProcessResult` 契約不一致。

症狀容易被誤判成「缺 NOTIFY」「packet-id stamp 失敗」「presence-only PAYOUT=0 不夠」。先比對：live／capture 的 strip cell 數、fixture 實際發出的 `SYMBOLS V=` cell 數、以及 `plainLen`。

即便 strip 對齊，部分柜（巨像／特殊 Spine 路徑）wire rewrite 仍可能卡轉——動畫猎取應改 apply-shim。但 **thin GLASS 是可先排除的硬錯誤**。

#### Trigger

- session-proxy `--rewrite-bets` 後第一手 `fixture_result` 成功，但 CTA 卡「停止」／不再發 `SLOTS_BET`
- Lab fixture 宣告 `returnedRowRange` 為全高 strip，卻只有 `visibleGrid`
- 同柜 apply-shim（記憶體內寫 Symbols）能 settle，wire RESULT 不能

#### Evidence

- Tool: session-proxy wire log（`plainLen`／`fixtureId`）、fixture RESULT builder、ProcessResult strip capture
- Sanitized excerpt: visible-only GLASS ≈15 cells／較短 plainLen；live strip ≈35 cells（Row −2..4）時 plainLen 明顯更長
- Evidence path: `<PROJECT_ROOT>/docs/lab/session-proxy.md`；柜机 lab 停轮 A/B 筆記（project-local）

#### Generalized Lesson

1. **先讀柜机 strip 契約**：`capture-config`／`ProcessResult` dump 的 `returnedRowRange` 與 symbol count。
2. **Builder 規則**：`symbols[]` 優先；僅有 `visibleGrid` 時，若契約是全高 strip，必須 pad／展開到同一 Row 窗，或拒絕發出 thin RESULT。
3. **不要用錯柜的 pad 策略**：可見窗就是 `[0,2]` 的柜（例：部分 giant／sparse GLASS）禁止套用 −2..4 pad。
4. **验收不只看第一手 rewrite**：連續 BET 的 `betCount≥2` 且 CTA 回到可再旋轉，才算 wire settle 成功。
5. **NOTIFY 與 stamp 分開排障**：`NOTIFY="desk"` 可能跳過 C2S `-xN`→S2C `ID=` stamp；先對齊 strip，再 A/B NOTIFY。
6. **Wire 仍卡轉 → apply-shim**：上游真 RESULT 停輪後再 patch Symbols／Rewards，專攻動畫帧。

#### Agent Action

- 做 rewrite lab 前：比對 fixture GLASS cell 數與柜机 `returnedRowRange`。
- 發現卡「停止」：先查 `plainLen`／cell 數是否 thin，再查 NOTIFY／stamp。
- 更新 builder 後：離線 assert cell 數 + 裝置上驗 `betCount` 連續遞增。
- 不把單一柜的 pad 窗寫成全庫預設。

#### Goal / Action / Validation

- Goal: wire 強制盤面時客戶端能停輪並連續下注
- Action: GLASS 對齊 ProcessResult strip；必要時 pad visible→strip
- Validation or reference source: wire `fixture_result` 的 cell 數／plainLen；`betCount` 隨 spin 增加；CTA 非永久「停止」

#### Applies When

- AES／自研 game-TCP 上對 `SLOTS_BET`（或同類）回放 RESULT
- Unity／il2cpp 柜机以 ProcessResult strip 套符號

#### Does Not Apply When

- 柜机契約本來就是可見窗 only（無 off-window rows）
- 僅用 Frida apply-shim、不走 wire RESULT
- 需要 sparse／巨像空洞編碼的特殊 GLASS（pad 會破壞 Spine）

#### Validation

- 同一 fixture：pad 前後 GLASS cell 數與 live strip 一致
- 連續兩次 spin：`betCount` +2 且第二次仍為 `fixture_result`
- 對照柜若禁止 pad，builder 測試必須覆蓋「不可誤 pad」

#### Promotion Target

- `feedback/history/apk-analysis/local-proxy/`（本條）
- 專案 `docs/lab/session-proxy.md` 櫃机差異表（已／應同步）
- 若多柜复现：可升到 analysis／workflow 的 replay-lab checklist

#### Promotion Record

尚未 promotion；保留為 candidate history。

#### Required Linked Updates

- 專案 session-proxy 文件註明：strip 柜 visibleGrid 必須 pad 或帶 `symbols[]`
- 新柜注册 rewrite builder 時核對 `returnedRowRange`

#### Abstraction Check

- Swap-project test: PASS — any Unity ProcessResult strip cabinet with session-proxy-style RESULT rewrite needs GLASS cell count to match returnedRowRange; visible-only emission is a hard error regardless of product name
- Project evidence kept in: `<PROJECT_ROOT>/docs/lab/` stop-reel / continuous-BET notes（project-local A/B tables）
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS
