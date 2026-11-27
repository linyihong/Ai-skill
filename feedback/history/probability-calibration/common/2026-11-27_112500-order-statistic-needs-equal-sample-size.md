> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-11-27 - Order statistics need equal sample sizes before a gap is claimed

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate（N=1）。本輪把一個「最大值差距」從缺口清單移除；consumer 側紀錄留在專案文件。

#### One-line Summary

「最大值／最小值／極端分位」是**順序統計量**，會隨樣本數單調漂移；兩邊樣本數不同時，直接比大小不能當缺口，必須先把產品側重抽樣到與參考側相同的單位數。

#### Human Explanation

參考觀察只有幾百個單位（場次），產品側用數十萬單位模擬，後者的最大值幾乎必然更大，這不代表產品尾部偏重。把兩個最大值並排寫進缺口清單，等於把「樣本數不同」誤讀成「模型不同」。正確做法是在產品側以**參考側的樣本數**重複抽樣許多次，得到「同樣樣本數下最大值的分布」，再看參考值落在該分布的哪個位置。固定高分位（如 p99）對樣本數敏感度低，可作為輔助交叉驗證。

#### Trigger

- 缺口清單出現「產品最大值 ≫ 參考最大值」
- 兩邊 `n` 相差一個數量級以上
- 指標是 max／min／極端分位／最長連續段

#### Evidence

- Tool: 以參考樣本數為單位的 bootstrap 重抽（數千次）得最大值分布；比對 p10／p50／p90 與固定高分位
- Sanitized excerpt: 參考側 ~380 單位、產品側 ~10 萬單位；直接比最大值差 2 倍以上；等量重抽後參考值落在產品分布約 p15，p99 兩邊接近
- Evidence path: `<PROJECT_ROOT>/docs/` 的校準缺口紀錄（個案數值不進本檔）

#### Generalized Lesson

1. 順序統計量宣稱缺口前，先確認兩邊 `n`；不等就等量重抽。
2. 報「參考值在產品分布的位置」與信賴區間，不報兩個點值的比值。
3. 用不隨 `n` 漂移的指標（固定分位、超越機率）交叉驗證。
4. 已寫進缺口清單的此類項目要回頭重驗；過期結論依 `stale_claim_reused` 處理。

#### Agent Action

- 先做：比對 `n`；產品側以參考 `n` 重抽；寫入位置與區間。
- 不要做：把大樣本最大值與小樣本最大值直接相除當缺口；為了「壓低最大值」去改尾部參數。

#### Goal / Action / Validation

- Goal: 缺口清單只含真缺口。
- Action: 等量重抽後再判定。
- Validation or reference source: 重抽分布覆蓋參考值（或明確在外）；固定分位交叉一致。

#### Applies When

- 參考樣本遠小於產品模擬樣本
- 指標為 max／min／極端分位

#### Does Not Apply When

- 兩邊 `n` 相同
- 指標為平均、比率等隨 `n` 收斂的估計量（用既有 SE 比較即可）
- 重抽單位與真實獨立單位不一致（如有狀態依賴）；需先處理樣本單位問題

#### Validation

換一個無關指標（如最長連續段）重做等量重抽，確認直接比大小與等量比較給出不同結論。

#### Promotion Target

- `workflow/probability-calibration/execution-flow.md`（Metric Parity：order statistic 需等量比較）
- `workflow/probability-calibration/contracts/calibration-contract.yaml`（estimator enum 與 `tail_event_count` 之外的 order-statistic 處理）
- `workflow/probability-calibration/registry/failure-classes.yaml`（候選 class：`unequal_sample_size_order_statistic`；N≥3 才升級）

#### Promotion Record

尚未 promotion；交由 plan 擁有者於 Phase 3 評估（對應其「分布比較型 estimator」observation）。

#### Required Linked Updates

- [x] `feedback/history/probability-calibration/common/README.md`
- [x] `feedback/history/probability-calibration/README.md`
- [x] `feedback/history/README.md` domain 表
- [ ] workflow 契約／registry（待 plan 擁有者決定，不由本 lesson 直接改）

#### Transfer Contract

- Inputs: 兩個樣本數不同的樣本與一個順序統計量指標
- Mechanism: 順序統計量期望值隨 n 單調成長
- Decision: 產品側以參考 n 重抽，比位置不比點值
- Verification: 重抽分布與固定分位交叉
- Counterexample: 平均／比率類指標
- Transfer case: 掉落稀有度、最大連勝、延遲 p99 等任何「小樣本觀察 vs 大樣本模擬」比較

#### Intelligence Review

- Decision: deferred
- Reason: 單一 consumer（N=1）；承接位置（workflow 契約或 Phase 2 atoms）由 plan 擁有者決定，本 lesson 不直接改 workflow
- Target: none
- Revisit: plan Phase 2／Phase 3 評估這條或第二個 consumer 出現同型情境時

#### Abstraction Check

- Swap-project test: PASS — 任何「小樣本觀察 vs 大樣本模擬」的 max／min 比較都適用
- Project evidence kept in: `<PROJECT_ROOT>/docs/`
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS
