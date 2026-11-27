# Optimization Probe Hygiene（校準探針與優化的操作衛生）

**Status**: `candidate-intelligence`
**Evidence**: N=1（單一 slot consumer 的優化迭代）
**Source**: plan Inventory I9；executor 邊界見 [`workflow/probability-calibration/contracts/verification.yaml`](../../../workflow/probability-calibration/contracts/verification.yaml) §`roles`、§`holdout_policy`

## 原則

優化器與探針是 executor 的工具，不是驗收者。用它們時：

1. **解析掃描優先於黑盒優化**：先對單一參數做有序掃描，看清指標隨參數變化的形狀，再決定是否需要黑盒搜尋。
2. **進度可觀察**：長時間探針應串流輸出，避免緩衝導致看不到進度而誤判卡住或重跑。
3. **探針要夠快才會被常用**：先降低單次探針成本（例如快取編譯或執行計畫），再增加探針次數。
4. **seed sweep**：同一組參數跨多個 seed 量測，區分參數效果與隨機波動。
5. **偵測退化解**：目標函數接近完美時先懷疑退化局部最小（例如目標指標趨近 0 讓損失變小），檢查所有指標而非只看損失。

## 判斷訊號

- 優化器回報極佳結果，但只在它自己用過的 seed 上
- 損失值突然大幅下降，某個指標同時坍縮到 0 或極端值
- 同一探針重跑很多次，原因是「看不到它在跑」

## 單一案例觀察（N=1，非通用規則）

- 該案在加快探針與改用串流輸出後，迭代速度明顯提升；也曾出現目標指標≈0 的退化解。具體倍數與數值留在 consumer，不在此引用。

## 適用條件

- 校準需要多次量測或搜尋
- 有 executor 與 verifier 的角色分離（優化結果不得直接成為驗收）

## 反例／不適用

- 參數空間極小、可窮舉：直接全掃，不需要優化器
- 本 atom 不是要求建自動優化器；Ai-skill 不提供優化器實作

## 與既有知識的關係

- 優化器在自己的 seed 上「通過」：failure class 與 scenario `probability-calibration-holdout-required-v1`
- 多變數同時亂調：[`shotgun-debugging`](../anti-patterns/shotgun-debugging.md)

## 升級條件

≥3 個獨立 consumer 的校準紀錄顯示同一衛生項目避免了錯誤結論（例如 seed sweep 推翻單 seed 結果），且至少一例不是 slot。
