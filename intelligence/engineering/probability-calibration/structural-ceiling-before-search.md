# Structural Ceiling Before Search（先證明結構天花板，再擴大搜尋）

**Status**: `candidate-intelligence`
**Evidence**: N=1（單一 slot consumer 的跨軸相關校準）；Phase 1 紙上走查提出「保底機制可能屬結構問題」但**未**驗證
**Source contract**: [`workflow/probability-calibration/contracts/gap-ledger.yaml`](../../../workflow/probability-calibration/contracts/gap-ledger.yaml) §`structural_ceiling`（C7）；判定邊界見 [`calibration-record.yaml`](../../../workflow/probability-calibration/contracts/calibration-record.yaml) §`status_boundary`

## 原則

當調參在合理範圍內持續無法逼近某個目標時，先問「目前的產生器結構能不能表達這個分布」。若結構本身表達不了，任何參數搜尋都只是在浪費算力，正確動作是停止搜尋、提出結構變更需求。

本 atom 只描述**判斷習慣**；判定所需欄位與 status 規則以契約為準，本檔不重新定義。

## 判斷訊號

- 多輪搜尋後，目標指標的最佳值停在同一個水準附近，且擴大範圍也不改善
- 目標描述的是變數之間的**關係**（相關、條件依賴、序列記憶），而產生器每個部分彼此獨立
- 改善某指標必須讓另一個 hard constraint 明顯退化，且無法用正交旋鈕分開（見 [`single-knob-side-effects.md`](single-knob-side-effects.md)）

## 單一案例觀察（N=1，非通用規則）

- 每個軸（或位置）獨立抽樣的結構無法產生跨軸相關；該案在證明天花板後換成耦合式抽樣，而非繼續調權重。

## 適用條件

- 已完成定義對齊與（適用時的）採集偏差排除
- 有搜尋紀錄可說明已嘗試的參數範圍

## 反例／不適用

- 搜尋範圍很窄或只試過一次：不足以宣稱天花板，應判 `inconclusive` 而非 `structural_blocked`
- 指標定義本身不一致：先修定義，那不是結構問題
- 保底、計數器等有狀態機制：**是否**構成天花板尚未驗證，不得套用本 atom 下結論（Phase 3 觀察項）

## 與既有知識的關係

- 判定流程與證據欄位：C7 契約（本檔不重複）
- 「不斷調權重希望會好」的行為面：[`shotgun-debugging`](../anti-patterns/shotgun-debugging.md)（無歸因的多變數嘗試）

## 升級條件

≥3 個獨立 consumer（不同遊戲或玩法類型）各自出現「證明天花板 → 換結構 → 目標達成」的紀錄，且至少一例為非 slot。
