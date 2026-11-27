# Probability Calibration Workflow

`workflow/probability-calibration/` 是 **機率模型校準與驗證**（Probability Model Calibration & Validation）的方法與契約：
Contract → Reference Evidence → Metric Parity → Diagnose → Tune → Freeze → Independent Verify → Status → Gap Ledger → Feedback。

> **狀態**：Phase 1 doc-only pilot（plan [`2026-11-27-1500`](../../plans/active/2026-11-27-1500-probability-calibration-workflow.md)）。
> **沒有** `route.workflow.probability-calibration`；所有 YAML `runtime_projection.enabled: false`。手動啟用本目錄。
> 第一個 consumer profile：[`profiles/slot-rtp.md`](profiles/slot-rtp.md)。中立 core 目前只有**紙上驗證**（N=1 真實 consumer）。

## Scope

處理讓**機率模型或隨機產品**的行為符合目標分布或規格，並以 Monte Carlo 方式驗證其統計行為。涵蓋：

| 類型 | 例 |
| --- | --- |
| 機率估計 | 從樣本估事件率、命中率、期望值 |
| 模型校準 | 調參使產生器符合目標分布（RTP、掉落率、牌型頻率） |
| Monte Carlo validation | 驗證抽樣模型的統計行為與獨立驗收 |

傳統「預測機率 calibration」（預測機率 vs 實際頻率一致性）只是其中一種特例。**不**處理：投資語境的機率表述（[`investment/`](../investment/README.md)）、官方行為逆向取證（[`apk-analysis/`](../apk-analysis/README.md)）。

## 一句話責任邊界

| 元件 | 問句 |
| --- | --- |
| Calibration Contract | 目標是什麼、誰定的、指標怎麼定義、要多準？ |
| Diagnose | 偏差是定義、採集，還是產品真的偏？結構做得到嗎？ |
| Tune | 候選參數怎麼產生？（非驗收） |
| Verification | 獨立驗收是否成立？達到哪幾維獨立性？ |
| Calibration Record | 這輪改了什麼、量到什麼、結果狀態與理由？ |
| Gap Ledger | 還剩哪些已知缺口？ |

## Mechanical invariants（摘要）

| ID | 規則 |
| --- | --- |
| PC1 | 沒有凍結的 calibration contract 不得調參 |
| PC2 | 診斷順序：定義 → 採集（僅具 reference_observation）→ 產品；只有最後一項成立才改參數 |
| PC3 | Execute ≠ Verify；executor（含優化器）不得自證 |
| PC4 | final holdout 不參與搜尋；用過即失去 final 身分 |
| PC5 | 指標分 hard_constraint／primary_objective／diagnostic；hard 不可被抵銷 |
| PC6 | 五個 status 互斥；`partially_calibrated`、`inconclusive` 是合法收工 |
| PC7 | 已知缺口必入 Gap Ledger；鏈式驗證一關失敗＝整體失敗 |
| PC8 | 個案資料與探針程式碼留在 consumer；本目錄只放方法與契約 |

跨域證據形狀參考：[`traceable-evidence`](../cross-cutting/traceable-evidence/README.md)。Execute≠Verify 精神與 [`image-processing`](../image-processing/README.md) 相同，但契約各自獨立。

## 何時讀哪個檔

| 認知階段 | 檔案 |
| --- | --- |
| Lifecycle | [`execution-flow.md`](execution-flow.md) |
| 目標／指標／門檻 | [`contracts/calibration-contract.yaml`](contracts/calibration-contract.yaml) |
| 獨立驗收／holdout | [`contracts/verification.yaml`](contracts/verification.yaml) |
| 每輪紀錄／status | [`contracts/calibration-record.yaml`](contracts/calibration-record.yaml) |
| 缺口／結構天花板 | [`contracts/gap-ledger.yaml`](contracts/gap-ledger.yaml) |
| Failure classes | [`registry/failure-classes.yaml`](registry/failure-classes.yaml) |
| Consumer profile | [`profiles/slot-rtp.md`](profiles/slot-rtp.md) |
| 候選判斷知識（N=1） | [`intelligence/engineering/probability-calibration/`](../../intelligence/engineering/probability-calibration/README.md) |

## 明確不做

- 註冊 `route.workflow.*`／discovery signal／runtime 投影
- 自動優化器、機率引擎程式碼、探針腳本
- 把單一遊戲或 consumer 的個案結論寫成通用規則（N<3 一律 `candidate-intelligence`）
- 預建尚無真實 consumer 的 profile 空殼
