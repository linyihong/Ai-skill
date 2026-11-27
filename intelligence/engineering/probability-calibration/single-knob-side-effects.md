# Single-Knob Side Effects（調一個旋鈕，量全部指標）

**Status**: `candidate-intelligence`
**Evidence**: N=1（單一 slot consumer 的多目標聯合校準）
**Source contract**: [`workflow/probability-calibration/execution-flow.md`](../../../workflow/probability-calibration/execution-flow.md) 步驟 4（每調一旋鈕量全部指標）；實驗紀錄形狀見 [`calibration-record.yaml`](../../../workflow/probability-calibration/contracts/calibration-record.yaml) §`experiment_ledger`（C8）

## 原則

在多指標校準中，一個參數幾乎從不只影響一個指標。每次只改一個旋鈕，並在改後重量**全部** hard constraint 與 primary objective，把副作用記入 experiment ledger。若兩個效果總是綁在同一個旋鈕上，應考慮把它拆成兩個正交旋鈕，而不是在兩個效果之間來回妥協。

## 判斷訊號

- 某指標改善時，另一個看似無關的指標同步移動
- 同一旋鈕被反覆往兩個方向調
- 實驗紀錄只寫了目標指標的前後值，沒有其他指標

## 單一案例觀察（N=1，非通用規則）

- 增加同類聚集以提升跨軸相關，會同步抬高總回報率（同類組合變多）；該案把「類別密度聚集」與「同類聚集」拆成兩個旋鈕後才能分別控制。

## 適用條件

- 有兩個以上指標同時受契約約束
- 產生器參數可以單獨變動並重新量測

## 反例／不適用

- 單一指標校準：沒有副作用需要量
- 指標彼此在數學上恆等相關（例如一個是另一個的固定倍數）：拆旋鈕沒有意義
- 量測成本極高、無法每次重量全部指標：至少量 hard constraint，並在 Gap Ledger 記 `process_gap`

## 與既有知識的關係

- 一次改多個變數而無法歸因：[`shotgun-debugging`](../anti-patterns/shotgun-debugging.md)。本 atom 是其正面做法在校準中的實例：單變數變更 + 全指標 after 對照
- 為何不能用加總分數掩蓋副作用：C2 objective tiers（契約）

## 升級條件

≥3 個獨立 consumer 的 experiment ledger 顯示「單旋鈕跨指標副作用 → 拆正交旋鈕」模式，且至少一例為非 slot。
