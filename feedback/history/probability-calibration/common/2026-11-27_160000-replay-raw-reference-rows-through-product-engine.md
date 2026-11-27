> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-11-27 - Replay the raw reference rows through the product engine before any statistic is compared

Status: candidate

#### Reuse Evidence

N=1；尚無獨立重用證據。

#### One-line Summary

校準的 parity 步驟若只比較統計量（頻率、比值、直方圖），會漏掉產品引擎對單一情境的系統性偏差，因為偏差可能只影響少數局、又被平均掉。把參照的原始樣本（盤面、事件標記、派彩）逐筆餵進產品引擎重放，逐局比對輸出，才是直接的 parity 測試。統計量比較應該放在這之後。

#### Human Explanation

本輪先後比了翻倍率、長連線比例、形狀占比，都在合理範圍內，還做了多輪參數搜尋。直到為了別的問題而重放原始盤面，才發現產品引擎在約 1.7% 的局上少賠，總賠付少 2.3%。這個偏差只在「前導 wild 接可賠付符號，而 wild 自己的階梯更高」的盤面出現，對整體統計量幾乎看不出來，但它改變了 RTP 的基準，使既有的校準建立在錯誤的引擎上。

重放的成本低（數十秒，唯讀），結果是逐局的 yes/no，不需要統計顯著性判斷。

#### Trigger

- 參照有逐筆原始資料（盤面、派彩、事件標記）
- 校準目標是 RTP、賠付分布或事件頻率，而產品有獨立的賠付引擎
- 統計量比較「大致對得上」，但要調的旋鈕與目標之間關係不清楚
- 另有一份獨立的參照重算（例如某個 TS 重算）已與原始資料一致，但產品引擎尚未被直接對到原始資料

#### Evidence

- Tool: 唯讀匯出參照原始列；用產品引擎評估同一盤面與同一事件標記；逐列比對派彩與連線數
- Sanitized excerpt: 102,591 列中，產品引擎與參照派彩不一致 1,720 列（全部少賠）；修正一個前導分支後 0 列不一致
- Evidence path: `<PROJECT_ROOT>/analysis/` 的校準紀錄（§0.2）；重放程式留在 consumer（未持久化）

#### Generalized Lesson

1. 校準的 parity 步驟分兩段：先逐列重放（引擎對參照的逐局一致性），通過後才比統計量。
2. 逐列重放的通過條件是「0 列不一致」或「不一致的列可被解釋並列入 Gap Ledger」。不接受「平均後接近」。
3. 參照的獨立重算（如另一語言的重算）與產品引擎是兩件事：重算一致只代表重算正確，不代表產品引擎正確。兩者都要對到原始資料。
4. 發現引擎偏差時，不要用參數去補。偏差會改變基準，既有校準要整批作廢或重跑，並記入紀錄。
5. 若修正只在副本上測試，紀錄必須寫明「未改 repo」，並列出需要的決策（修引擎、重校準、其他情境是否受影響）。

#### Agent Action

- 先做：取得參照原始列；唯讀重放產品引擎；逐列比對；把不一致列分類。
- 不要做：只比統計量就宣稱 parity；用參數補償引擎偏差；因為偏差「只影響一小部分局」就忽略它。

#### Goal / Action / Validation

- Goal: 確認產品引擎在參照的每一局上與參照一致。
- Action: 原始列→產品重放→逐列比對→分類不一致列→入帳。
- Validation or reference source: 逐列比對結果表（不一致數、方向、受影響的情境）。

#### Applies When

- 有逐筆參照原始資料
- 產品引擎與參照的規則由程式實作，而非只靠統計擬合

#### Does Not Apply When

- 參照只有彙總統計、沒有逐筆資料
- 目標本身就是統計分布，且沒有單局的正確答案（例如純隨機性的隱藏種子）

#### Validation

逐列重放結果：不一致列數與修正後的列數；至少兩種情境（例如有與沒有某個分支）各自重放。

#### Promotion Target

- `workflow/probability-calibration/execution-flow.md`（parity 步驟：先逐列重放，後統計量）
- `workflow/probability-calibration/contracts/`（parity 證據形狀：加逐列重放的一致性欄位）

#### Promotion Record

尚未 promotion。

#### Required Linked Updates

- [x] `feedback/history/probability-calibration/common/README.md`
- [x] `feedback/history/probability-calibration/README.md`（domain 已列）
- [ ] execution-flow parity 步驟補充（待 plan 擁有者）

#### Transfer Contract

- Inputs: 有逐筆原始資料的參照，與一個獨立實作的產品引擎
- Mechanism: 引擎在某個特定分支上與參照規則不同，統計量上不明顯
- Decision: 先逐列重放，不一致列入帳後才談統計量與參數
- Verification: 逐列不一致數為零，或每列都有解釋
- Counterexample: 產品引擎就是參照本身（不需要 parity）
- Transfer case: 任何有線上日誌可回放的遊戲、計費、風控規則引擎

#### Intelligence Review

- Decision: deferred
- Reason: 單一 consumer（N=1）
- Target: none
- Revisit: plan Phase 3 或第二個 consumer

#### Abstraction Check

- Swap-project test: PASS — 任何需要與參照原始資料對齊的規則引擎都適用
- Project evidence kept in: `<PROJECT_ROOT>/analysis/`
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS
