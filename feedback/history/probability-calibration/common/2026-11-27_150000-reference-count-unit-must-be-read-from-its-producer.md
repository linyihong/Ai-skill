> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-11-27 - A reference count's unit must be read from the tool that produced it, not from the prose that quotes it

Status: candidate

#### Reuse Evidence

N=1；尚無獨立重用證據。

#### One-line Summary

參照數字常在文件裡被轉述成「N 條」「N 次」，單位與產生它的工具不同。把轉述的數字當目標去調參，等於對錯誤的指標做最佳化。診斷步驟 a（指標定義不一致）要做的不是「看一眼定義文字」，而是讀產生參照的工具原始碼或用自己的量測在同一份原始資料上重算該數字，確認單位、分母與去重方式一致。

#### Human Explanation

本輪的參照數字在轉述文件裡寫成「條」（每條連線），產生它的工具實際是「每局最長連線落在該長度的局數」（每局只記一次）。我方探針逐線計數，一局有兩條長連線就記兩次。兩邊單位不同，算出「多 1.4–1.9 倍」的缺口，並據此做了多輪參數搜尋與最佳化；對齊單位後缺口只剩約 3–5%。

這不是粗心的算術。轉述文件、探針的標籤與我方的指標名稱三處都用了同一個字，每一處單看都合理。發現它是因為後來為了別的目的去讀參照工具的原始碼，才看到它用「每局取最長」去重。

#### Trigger

- 目標是從參照樣本取得的計數，而且該計數是在文件或對話裡被轉述過
- 我方指標的標籤和參照的標籤一樣，但兩邊是由不同工具各自產生
- 第一次比較就出現「整數倍」的差距，而差距方向恆定
- 調參很費力：所有候選都在兩個無關指標之間拉扯，卻沒有任何一個靠近目標

#### Evidence

- Tool: 讀參照工具的原始碼確認計數與去重方式；在我方探針加上同口徑的直方圖，與參照用同一組分母重算
- Sanitized excerpt: 逐線計數對每局最長連線計數，差 1.4–1.9 倍；同口徑重測後，7+ 連占有翻倍局比例為 0.110–0.112 對 0.107，兩個種子一致
- Evidence path: `<PROJECT_ROOT>/docs/` 的校準紀錄（§更正）；探針與查詢腳本留在 consumer

#### Generalized Lesson

1. 參照值進入契約前，在契約的指標定義裡寫明**計數單位、分母、去重方式**（每線／每局／每局取最大）。只寫指標名稱不夠。
2. 執行診斷步驟 a 時，要有一個**可重現的對照**：用自己的量測工具，在與參照工具相同的口徑下重算該參照數字。不是用文件裡的轉述。
3. 轉述過的數字（README、報告、對話）預設「單位未驗證」，直到對照過產生它的工具。
4. 調參前先做一次無參數變更的口徑對照；若比值落在整數倍附近，先懷疑單位，不是先懷疑產品。
5. 發現口徑錯誤後，所有建立在錯誤指標上的實驗結果只算診斷線索，不是校準證據；紀錄裡要明寫哪些節作廢，並一併修正含有錯誤比較的下游文件。
6. 工作流程的 a 步驟標記為「一律」，不得因為「這個數字看起來很明確」而略過。

#### Agent Action

- 先做：讀參照工具或在同口徑下重算；把單位、分母、去重寫進契約的指標定義；做一次不改參數的口徑對照。
- 不要做：用轉述文件裡的數字當目標；在兩個工具用同一個詞時假設它們是同一件事；在口徑未驗證前投入搜尋或最佳化。

#### Goal / Action / Validation

- Goal: 讓調參的目標與參照在同一單位上。
- Action: 讀產生者→同口徑重算→寫入指標定義→再調參。
- Validation or reference source: 同口徑的對照表（我方 vs 參照），差距在量測波動內或有解釋。

#### Applies When

- 參照值來自另一個工具、另一個人或另一份文件
- 計數可以有多種合理單位（每線、每局、每場、每人）

#### Does Not Apply When

- 參照與我方指標由同一個程式、同一個定義產生
- 指標是單位已固定的比例或金額（但仍要確認分母）

#### Validation

在我方工具上同時輸出兩種口徑（例如逐線與逐局），確認參照數字與哪一個接近；兩個獨立種子上結論一致。

#### Promotion Target

- `workflow/probability-calibration/execution-flow.md`（診斷步驟 a：同口徑重算與單位寫入契約）
- `workflow/probability-calibration/contracts/` 的指標定義形狀（加 `count_unit`、`denominator`、`dedup`）

#### Promotion Record

尚未 promotion。

#### Required Linked Updates

- [x] `feedback/history/probability-calibration/common/README.md`
- [x] `feedback/history/probability-calibration/README.md`（domain 已列）
- [ ] 診斷步驟 a 與契約指標定義補充（待 plan 擁有者）

#### Transfer Contract

- Inputs: 轉述過的參照計數，與我方獨立產生的同名指標
- Mechanism: 兩邊單位或去重不同，同一個詞指向不同的量
- Decision: 讀產生者或同口徑重算，確認單位後才設為目標
- Verification: 我方在同口徑下的值與參照接近，或差距有解釋
- Counterexample: 指標由同一程式產生，口徑天然一致
- Transfer case: 日誌分析的「錯誤次數」（次／使用者／工作階段）、流量統計的「訪問數」、測試的「失敗數」（案例／斷言）

#### Intelligence Review

- Decision: deferred
- Reason: 單一 consumer（N=1）
- Target: none
- Revisit: plan Phase 3 或第二個 consumer

#### Abstraction Check

- Swap-project test: PASS — 任何與外部參照比較計數的校準都適用
- Project evidence kept in: `<PROJECT_ROOT>/docs/`
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS
