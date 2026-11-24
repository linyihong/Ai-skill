# Bootstrap Receipt Gate Blind To Mid-turn Text（Receipt gate 看不到回合中的文字）

Status: candidate
Class: `enforcement-self-leak` / `validation-gap`

## Trigger

- 新 session 的第一個回合持續呼叫工具超過 SessionStart 的 120 秒寬限。
- PreToolUse 回 `Bootstrap Receipt missing`，但本回合已輸出 Receipt。
- 補讀檔、重貼 Receipt 後，同一個 deny 仍重複出現。

## Failure Mode

`gate.bootstrap.receipt_present` 只認 transcript 中 `type: assistant` 文字區塊裡的 `Bootstrap:`。有些 client 在回合進行中不會即時寫入夾在工具呼叫之間的文字，於是 gate 判定 Receipt 不存在。deny 訊息要求「讀檔並輸出 Receipt」，agent 照做也無法解除，陷入重試迴圈；只有結束回合、讓最後回覆寫入 transcript 後才會解除。

## Risk

- 長回合中途卡住，變更只寫了一半，留下 dirty tree。
- agent 為了解除阻擋反覆重讀檔案、重貼 Receipt，浪費回合。
- 若有人為了解除而改成 flag 永久放行，gate 會整個 session 失效（先前已發生過並修正）。

## Required Agent Action

1. **第一回合完成 bootstrap 後就結束回合**，讓含 Receipt 的回覆寫進 transcript，再開始做事。
2. **被擋時先分辨 deny 類型。** `receipt without required reads` → 補讀檔；`no receipt` 但本回合已輸出 → 不重試，回報進度與尚未寫入的變更，結束回合。
3. **同一個 deny 第二次出現時，去讀 gate 實作或 hook log**，不要重複同一個修法。

## Prevention Gate

- **文件層（已存在）：** 本 pattern、[`feedback/history/development-guidance/common/2026-11-25_084854-bootstrap-receipt-gate-blind-to-mid-turn-text.md`](../../feedback/history/development-guidance/common/2026-11-25_084854-bootstrap-receipt-gate-blind-to-mid-turn-text.md)、[`ai-tools/agent/claude.md`](../../ai-tools/agent/claude.md) §已知限制。
- **Gate 層（待定）：** deny 訊息加上「結束回合可解除」；或讓 gate 也接受 transcript 中已成功執行的 `ai-skill runtime receipt` 結果加上 required reads 作為證據。尚未實作，需使用者裁決。
- **Promotion 條件：** 另一個 session 或 client 再次觀察到時，才考慮提升為 `validated`。

## 驗證

- 第一回合結束後，下一回合 hook log 出現 `receipt found + required reads verified`。
- 被擋時的回報說明 gate 讀的來源與 deny 類型。

## Linked Rules

- [`pretooluse-block-wrong-exit-code.md`](pretooluse-block-wrong-exit-code.md)（同屬 gate 名實不符）
- [`bootstrap-bypass-on-resume.md`](bootstrap-bypass-on-resume.md)
- [`failure-learning-system.md`](../failure-learning-system.md)
