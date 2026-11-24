> 遵守 [共用規則索引](../../../../enforcement/README.md) 與 [回饋規則](../../../feedback-lessons.md)；原始證據留在本機 session transcript 與 hook log。

### 2026-11-25 - 回合中輸出的 Bootstrap Receipt 不一定看得到；120 秒寬限過後，PreToolUse 會擋住整個回合

Status: candidate

#### Reuse Evidence

單一來源：Ai-skill 本 repo 一個 Claude Code（桌面版 Code 分頁）session。hook log 較早的段落也有一串 `deny (no receipt)`，但沒有逐一查證是否同一原因，不計入重用證據。

#### One-line Summary

`gate.bootstrap.receipt_present` 只認 transcript 裡 `type: assistant` 文字區塊中的 `Bootstrap:`。這個 client 在回合進行中，大多不會把夾在工具呼叫之間的文字寫進 transcript。SessionStart 的 120 秒寬限一過，Write / Edit / Bash 就全被擋。agent 重複輸出 Receipt 也沒用，直到回合結束、最後回覆寫入 transcript 後才解除。

#### Human Explanation

1. SessionStart 寫入 flag。PreToolUse 在 120 秒內放行所有工具，但**不**標記 gate 已通過。這是刻意的：先前把 flag 當成永久通過，會讓 gate 整個 session 失效。
2. 寬限期過後，gate 用 `extractAssistantTexts` 掃 transcript，找 `Bootstrap:`。
3. 本次 session 中，agent 在回合中輸出了 4 次 Receipt，transcript 裡一次都沒有對應的 assistant 文字項目。同一回合有少數過場文字有被寫入，所以是「不可靠」，不是「從不寫入」。
4. 結果：8 次 flag 放行後，連續 8 次 `deny (no receipt)`。補讀 3 份 required reads、再輸出 Receipt、改成「Receipt 單獨一則訊息」都沒用。回合結束後，最後回覆（開頭就是 Receipt）寫進 transcript，下一回合第一次呼叫工具就是 `receipt found + required reads verified`。
5. 錯誤訊息寫的是「Receipt missing，請先讀檔再輸出」，會把 agent 導向重讀檔案、重貼 Receipt 的無效迴圈。實際上可行的只有「結束回合」。

`ai-tools/agent/claude.md` 原本把這個「同一回合先輸出 Receipt 文字、再呼叫工具」的問題記為「已修」（靠 flag）。但 flag 已改成只有 120 秒寬限，所以超過 120 秒的長回合仍會遇到。

#### Trigger

- 新 session 第一個回合很長，沒有先結束回合就持續呼叫工具，超過 120 秒。
- PreToolUse 回 `Bootstrap Receipt missing`，但本回合已經輸出過 Receipt。
- 同一個 deny 在補讀檔、重貼 Receipt 後仍然重複出現。

#### Evidence

- Tool: session transcript JSONL、`ai-skill-bootstrap-hook.log`（系統暫存目錄）、`scripts/ai-skill-cli/internal/app/hooks.go`（PreToolUse gate、`extractAssistantTexts`）。
- Sanitized excerpt: SessionStart 後約 106 秒的 Write 被放行，約 132 秒後的 Write 被擋。log 依序為 8 行 `exit_code: 0 (sessionstart flag, not cached)`、8 行 `deny (no receipt)`，接著下一回合的 `exit_code: 0 (receipt found + required reads verified)`。transcript 中含 `Bootstrap: rules=` 的 assistant 文字項目，第一回合只有最後回覆那一筆。
- Evidence path: 本機 transcript 與 hook log；不提交。

#### Generalized Lesson

1. **機械 gate 依賴的證據，要確認 client 何時把它寫進來源。** 「agent 有輸出」不等於「gate 讀得到」；寫入時機屬於 client 實作，不能假設是即時的。
2. **寬限期不是解法，只是延後。** 寬限期一過，所有超過寬限長度的長回合都會重現同一個問題。
3. **阻擋訊息要指出能解除的動作。** 訊息只說「請輸出 X」，但 X 已輸出卻看不到時，agent 會陷入重試迴圈。應說明 gate 讀的是哪個來源，以及「結束回合」可以解除。
4. **重複被擋時，先查 gate 讀什麼，不要重複同一個修法。** 第二次出現相同 deny，就該去讀 gate 的實作或 log。

#### Agent Action

- 新 session 的第一個回合：先讀 `CORE_BOOTSTRAP.md`、`runtime/core-bootstrap.yaml`、3 份 required reads，執行 `ai-skill runtime receipt`，然後**結束回合**，讓含 Receipt 的回覆寫進 transcript，再開始做事。工作量大時尤其要這樣做。
- 如果已在長回合中被擋：不要重讀檔案、重貼 Receipt。直接回報目前進度與尚未寫入的變更，結束回合，請使用者回覆「繼續」。

#### Goal / Action / Validation

- Goal: 長回合不因 Receipt 寫入時機而卡住，也不因此跳過 bootstrap。
- Action: 第一回合完成 bootstrap 後結束回合；被擋時結束回合，而不是重試。
- Validation or reference source: hook log 在第二回合出現 `receipt found + required reads verified`，且 required reads 都在 transcript 中。

#### Applies When

- Claude Code client 在回合中不會即時把 assistant 文字寫進 transcript JSONL（本次觀察為桌面版 Code 分頁）。
- PreToolUse gate 以 transcript 文字判斷 Receipt 是否存在。

#### Does Not Apply When

- 第一回合在 120 秒內就結束。
- 被擋的原因是 `receipt without required reads`（真的缺讀檔），那要補讀檔。

#### Validation

- 新 session 第一回合結束後，下一回合第一次呼叫非 Read 工具時，hook log 為 `receipt found`。
- 被擋時，回報中說明 gate 讀的來源，而不是只說「已重新輸出 Receipt」。

#### Promotion Target

- `enforcement/failure-patterns/bootstrap-receipt-gate-blind-to-mid-turn-text.md`（agent 面行為與 gate 面缺口）
- 程式修補：deny 訊息加上「結束回合可解除」**已完成**（2026-11-25，使用者裁決）。讓 gate 也接受 transcript 中已成功執行的 `ai-skill runtime receipt` 結果（該指令已被 gate 放行）配合 required reads 作為證據，**未做**。

#### Promotion Record

尚未 promotion；保留為 candidate history。

#### Required Linked Updates

- `enforcement/failure-patterns/README.md` 加索引列。
- `feedback/history/development-guidance/README.md` 分類數量與 Recent 表；`common/README.md` 索引列。
- `ai-tools/agent/claude.md` §已知限制：「雙重 bootstrap 已修」改為「只在 120 秒寬限內緩解」，並指向本 lesson。

#### Transfer Contract

- Inputs: 一個以 transcript / log 內容判斷合規的機械 gate，加上非即時寫入的 client。
- Mechanism: 被檢查的事實已發生，但還沒寫進 gate 讀的來源，gate 判定為不存在。
- Decision: 被擋且確定已做時，先確認來源的寫入時機，再決定是等待寫入（結束回合）還是修 gate。
- Verification: 寫入後 gate 解除；log 看得到狀態轉換。
- Counterexample: 真的沒做（例如缺 required reads），gate 擋得正確，就要補做。
- Transfer case: 任何讀 transcript、event log、檔案系統快照的 PreToolUse / Stop / commit gate。

#### Intelligence Review

- Decision: deferred
- Reason: 單一 session、單一 client 的觀察；client 寫入 transcript 的行為沒有公開契約，可能隨版本改變。
- Target: none
- Revisit: 另一個 session 或另一個 client 再次出現「Receipt 已輸出但 gate 看不到」時重新評估。

#### Abstraction Check

- Swap-project test: PASS — 把 Receipt 換成任何「agent 已輸出、gate 讀 transcript 判斷」的證據，結論仍成立。
- Generalized Lesson has no product/cabinet/sample IDs：已確認。
- Project evidence kept in: 本機 transcript 與 hook log（不提交）。

#### Linked Rules

- [`enforcement/failure-patterns/bootstrap-receipt-gate-blind-to-mid-turn-text.md`](../../../../enforcement/failure-patterns/bootstrap-receipt-gate-blind-to-mid-turn-text.md)
- [`enforcement/failure-patterns/pretooluse-block-wrong-exit-code.md`](../../../../enforcement/failure-patterns/pretooluse-block-wrong-exit-code.md)（同屬 gate 名實不符）
- [`enforcement/failure-patterns/bootstrap-bypass-on-resume.md`](../../../../enforcement/failure-patterns/bootstrap-bypass-on-resume.md)
