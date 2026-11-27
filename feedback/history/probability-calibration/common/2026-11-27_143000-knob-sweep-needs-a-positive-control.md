> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-11-27 - A knob sweep needs a positive control, or "no effect" cannot be told from "override never applied"

Status: candidate

#### Reuse Evidence

N=1；尚無獨立重用證據。

#### One-line Summary

校準實驗用外部覆寫（命令列、設定表）去動一個旋鈕時，若覆寫的鍵名不被消費端認得，系統常常靜默忽略，輸出與基準逐位相同。這個結果和「旋鈕真的沒有效果」長得一模一樣。每個旋鈕族在下結論前，先跑一個已知會動的對照組，確認覆寫真的進得去。

#### Human Explanation

本輪對一個新旋鈕族做了五組變體，五組的所有指標都和基準相同到最後一位。如果當時就寫「這個旋鈕族無效」，紀錄會是錯的。真正原因是覆寫用了常數名，而消費端讀的是字串鍵，覆寫被丟掉了。能發現，是因為「五組完全一樣」本身太整齊，不像真實的隨機波動。

這個失敗模式的危險在於它產生的是「負面發現」。負面發現不會被人再查（沒有東西壞掉），會直接進入結論，甚至進入 Gap Ledger 當成「已排除的旋鈕」。

#### Trigger

- 用字串鍵、環境變數或命令列參數覆寫設定，且消費端對未知鍵不報錯
- 一組變體的指標與基準完全相同（含最後一位），而變體間的設定確實不同
- 覆寫的鍵名由人手寫，而不是由消費端匯出

#### Evidence

- Tool: 比較每個變體與基準的全部輸出指標；再用一個預期會讓指標明顯移動的極端值重跑
- Sanitized excerpt: 第一次五個變體對基準的差為 0（所有指標逐位相同）；改用消費端實際讀的鍵名後，同樣的五個變體指標各自移動，且方向與預期一致
- Evidence path: `<PROJECT_ROOT>/docs/` 的校準紀錄（§實驗紀錄）；探針留在 consumer

#### Generalized Lesson

1. 對每個旋鈕族，實驗序列的第一步是**正對照**：一個明知會移動指標的極端值。對照組沒動，就是覆寫沒進去，不是旋鈕無效。
2. 變體與基準逐位相同是警訊，不是結論。真實的隨機波動不會讓多個不同設定都給出完全相同的輸出。
3. 覆寫的鍵名從消費端的來源匯出或驗證（讀取時記錄「已套用的鍵」），不要靠人手寫。能讓消費端對未知鍵報錯更好。
4. 「旋鈕族無效」要寫進 Gap Ledger 或實驗紀錄前，必須附正對照的結果；沒有正對照的負面發現，狀態是 `inconclusive`，不是「已排除」。
5. 反向也成立：旋鈕有效的結論同樣需要確認，是「覆寫進去了」造成的，而不是別的差異（種子、建置）。

#### Agent Action

- 先做：每個旋鈕族先跑極端值正對照；變體與基準逐位比對；紀錄覆寫是否被套用。
- 不要做：看到「全部相同」就寫「無效」；用手寫的鍵名而不驗證它被消費端認得。

#### Goal / Action / Validation

- Goal: 讓「旋鈕無效」只在旋鈕真的無效時出現。
- Action: 正對照→逐位比對→記錄已套用的覆寫→才下結論。
- Validation or reference source: 對照組指標明顯移動；變體彼此不同。

#### Applies When

- 任何以外部覆寫驅動的參數掃描、敏感度分析、最佳化迴圈
- 消費端對未知鍵靜默忽略

#### Does Not Apply When

- 消費端對未知鍵會報錯或拒絕啟動
- 參數經由型別檢查過的程式介面傳入（鍵名打錯在編譯期就失敗）

#### Validation

故意傳一個不存在的鍵，確認流程要嘛報錯、要嘛在紀錄裡標示「未套用」；正對照的極端值讓指標移動的幅度大於重複量測的波動。

#### Promotion Target

- `workflow/probability-calibration/execution-flow.md`（tune 步驟：旋鈕族的正對照要求）
- `workflow/probability-calibration/contracts/` 的實驗紀錄形狀（C8：加「覆寫已套用的證據」欄）

#### Promotion Record

尚未 promotion。

#### Required Linked Updates

- [x] `feedback/history/probability-calibration/common/README.md`
- [x] `feedback/history/probability-calibration/README.md`（domain 已列）
- [ ] execution-flow 與 C8 實驗紀錄形狀補充（待 plan 擁有者）

#### Transfer Contract

- Inputs: 由外部覆寫驅動的參數掃描
- Mechanism: 覆寫在消費端被靜默丟棄，輸出等於基準
- Decision: 先跑正對照，再解讀負面結果
- Verification: 對照組指標移動；變體彼此不同
- Counterexample: 型別檢查過的程式介面，鍵名錯誤會編譯失敗
- Transfer case: 效能調校的旗標掃描、機器學習超參數網格、功能旗標的 A/B 實驗

#### Intelligence Review

- Decision: deferred
- Reason: 單一 consumer（N=1）
- Target: none
- Revisit: plan Phase 3 或第二個 consumer

#### Abstraction Check

- Swap-project test: PASS — 任何覆寫驅動的實驗都適用
- Project evidence kept in: `<PROJECT_ROOT>/docs/`
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS
