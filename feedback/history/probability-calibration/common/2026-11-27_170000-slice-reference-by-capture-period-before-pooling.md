> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-11-27 - Slice the reference by capture period before pooling it into one rate

Status: candidate

#### Reuse Evidence

N=1：同一 consumer 的第二次完整校準（補強前一輪留下的缺口）；尚無獨立 consumer 重用證據。

#### One-line Summary

參照事件率若來自跨多個採集時段的語料，先按列順序或日期切段，逐段看是否穩定；混合多種採集狀態的全期比率不是任何一種狀態的真值，用它當目標會把產品調到錯的點。

#### Human Explanation

前一輪的參照事件率（1.45%）用全期語料一次算出。補缺口時按列 id 與日期切開，結果是三段：第一段 3.6% 且沒有結算記錄、中間一段每天 1.22–1.48% 合計 1.40%、最後一段 0.34%。只有中間一段穩定且有完整記錄管道。全期值恰好落在穩定值附近，所以看不出問題；但產品原本就比它高 11%，換成穩定段後是高 15%，且遺失比例的來源也變清楚（約三分之二的遺失集中在沒有結算記錄的那一段，與結果大小無關）。

#### Trigger

- 參照語料跨數週以上，採集器、遊戲版本或帳號狀態可能變過
- 事件率或遺失比例只有一個全期數字
- 另一個端點的記錄在某段時間完全沒有

#### Evidence

- Tool: 按列 id 與日期分段的事件率與遺失率（唯讀聚合）
- Sanitized excerpt: 全期 1.45%；分段後 3.6%／1.40%／0.34%，其中中間段每日值與二項雜訊一致，遺失率從 0%（第一段）到 86% 抓到（中間段）
- Evidence path: `<PROJECT_ROOT>/docs/` 的校準契約與紀錄；查詢腳本留在 consumer 並已刪除

#### Generalized Lesson

1. 設參照值之前，先畫（或表列）事件率與記錄完整率對列順序或日期的分段值。
2. 取穩定且記錄完整的一段當參照，其餘段落以缺口記錄，標明未解釋。
3. 每日值是否與二項雜訊一致是判斷「穩定」的最小檢查。
4. 契約的 `sample_scope` 寫明用了哪些段、排除了哪些段與原因。

#### Agent Action

- 先做：分段表；只在穩定段計參照；排除段入帳。
- 不要做：因全期值看起來合理就跳過分段；在不知道排除段原因時假裝已解釋。

#### Goal / Action / Validation

- Goal: 參照值代表一個明確的採集狀態。
- Action: 分段、取穩定段、記錄排除。
- Validation or reference source: 分段表與每日雜訊檢查。

#### Applies When

- 語料跨多個採集時段，或採集管道可能改過

#### Does Not Apply When

- 單次受控實驗、採集條件固定

#### Validation

用另一批時段重算，穩定段的值落在原區間內。

#### Promotion Target

- `workflow/probability-calibration/contracts/calibration-contract.yaml`（`sample_scope` 與診斷 b：分段檢查）

#### Promotion Record

尚未 promotion；交由 plan 擁有者評估。

#### Required Linked Updates

- [x] `feedback/history/probability-calibration/common/README.md`
- [x] `feedback/history/probability-calibration/README.md`（domain 已列）
- [ ] 契約診斷 b 的分段說明（待 plan 擁有者）

#### Transfer Contract

- Inputs: 跨時段語料與逐列順序
- Mechanism: 採集狀態變動使全期比率混合多個分布
- Decision: 取穩定段、排除段入帳
- Verification: 每日值與二項雜訊一致、另一批重算
- Counterexample: 固定條件的受控實驗
- Transfer case: 日誌取樣、A/B 實驗期間的版本變動、感測器韌體更新

#### Intelligence Review

- Decision: deferred
- Reason: 單一 consumer（N=1）；承接位置由 plan 擁有者決定
- Target: none
- Revisit: plan Phase 3 或第二個 consumer

#### Abstraction Check

- Swap-project test: PASS — 與領域無關的資料穩定性問題
- Project evidence kept in: `<PROJECT_ROOT>/docs/`
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS
