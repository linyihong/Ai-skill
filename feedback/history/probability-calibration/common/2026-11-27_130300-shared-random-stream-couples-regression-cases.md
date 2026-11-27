> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-11-27 - A shared random stream couples unrelated regression cases and masks rare failures

Status: candidate

#### Reuse Evidence

N=1；尚無獨立重用證據。

#### One-line Summary

迭代多個案例共用同一條亂數流的回歸測試，會讓前面案例的取樣次數改變後面案例看到的隨機值；無關的參數變更因此讓某個案例翻紅，也可能讓真實的低機率缺陷長期沒被抽到。改成每案例自己的種子，再用多種子量出缺陷率。

#### Human Explanation

共用一條流時，案例 i 看到的值取決於案例 1..i−1 消耗了多少隨機數。任何會改變前面案例消耗量的變更（包含參數檔變更）都會移動後面所有案例。症狀是「改了 A，B 的測試壞了」，但 B 與 A 無關。把測試改為以案例識別碼導出種子後，測試穩定，但這同時代表先前的通過可能只是剛好沒抽到壞情況；必須另外以多種子量測，判斷是不是真的有低機率缺陷。

#### Trigger

- 測試迴圈對多個案例共用一個隨機數產生器
- 無關變更造成某案例失敗，單獨跑又通過
- 失敗訊息指向看似與變更無關的元件

#### Evidence

- Tool: 對疑似案例以數千個獨立種子重複執行，統計失敗次數與失敗類型
- Sanitized excerpt: 共用流版本在一次參數變更後 3/3 失敗；每案例種子版本通過；3000 個獨立種子量出該案例約 0.07%（2/3000）的守衛例外，其他案例 0/3000
- Evidence path: `<PROJECT_ROOT>/docs/` 的校準紀錄；探針與量測腳本留在 consumer

#### Generalized Lesson

1. 測試的隨機輸入由案例識別碼導出種子，不共用流。
2. 修成穩定測試之後，不要把「現在通過」當成「沒有缺陷」；用多種子量出失敗率。
3. 量到非零率時，記為潛在缺陷（Gap Ledger `known_model_limitation` 或 consumer 的缺陷清單），不要藏進測試的種子選擇。
4. 釘值（regression lock）與校準是兩回事（`pinned_value_claimed_as_calibration`）；測試穩定性修復也不得被描述成校準證據。

#### Agent Action

- 先做：每案例種子；多種子量失敗率；非零率入帳。
- 不要做：挑一個會讓測試通過的種子然後結案；把共用流造成的失敗歸咎於當次變更。

#### Goal / Action / Validation

- Goal: 讓回歸測試只因相關變更而失敗，並讓低機率缺陷可見。
- Action: 案例獨立種子→多種子量測→入帳。
- Validation or reference source: 多種子失敗率表。

#### Applies When

- 迴圈式測試共用亂數產生器
- 測試對象本身含隨機行為

#### Does Not Apply When

- 測試輸入完全固定、無隨機
- 測試本來就是以單一固定種子作為釘值且不跨案例

#### Validation

把案例順序打亂或在前面案例插入額外取樣，確認結果不變；多種子失敗率在另一組種子上同量級。

#### Promotion Target

- `workflow/probability-calibration/contracts/verification.yaml`（釘值分層：記錄測試種子獨立性）
- `intelligence/engineering/probability-calibration/sampler-strategy-as-data.md`（零 RNG 漂移護欄補充：案例獨立種子）

#### Promotion Record

尚未 promotion。

#### Required Linked Updates

- [x] `feedback/history/probability-calibration/common/README.md`
- [x] `feedback/history/probability-calibration/README.md`（domain 已列）
- [ ] verification 契約或 I8 atom 補充（待 plan 擁有者）

#### Transfer Contract

- Inputs: 共用亂數流的多案例回歸測試
- Mechanism: 前案例的消耗量決定後案例的輸入
- Decision: 每案例獨立種子並以多種子量缺陷率
- Verification: 打亂順序結果不變；多種子失敗率穩定
- Counterexample: 無隨機輸入的測試
- Transfer case: 任何以單一 RNG 迭代多個設定的模擬型測試、基準測試

#### Intelligence Review

- Decision: deferred
- Reason: 單一 consumer（N=1）
- Target: none
- Revisit: plan Phase 3 或第二個 consumer

#### Abstraction Check

- Swap-project test: PASS — 與領域無關的測試設計問題
- Project evidence kept in: `<PROJECT_ROOT>/docs/`
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS
