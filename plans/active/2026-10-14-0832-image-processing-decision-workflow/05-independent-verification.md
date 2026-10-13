# 05 — Independent Verification Contract

← [`_plan.md`](_plan.md)

## 原則

1. **執行成功 ≠ 品質成功**
2. 驗收器不得複用執行器的「自評分數」作為唯一依據
3. 無前景真值時，統計指標最多做 **異常偵測**，不得單獨證明品質良好
4. `inconclusive` 是一等公民狀態

## 驗收維度

| check 族 | 機械檢查 | 視覺／語意檢查 |
| --- | --- | --- |
| 檔案有效性 | 可解碼、尺寸、格式 | 通常不需要 |
| 透明通道 | alpha 存在、直方圖／非全不透明 | 透明區是否符合意圖 |
| 主體完整性 | 遮罩面積、連通域異常 | 缺件／缺髮／缺配件 |
| 背景殘留 | 遮罩外亮度／彩度統計、邊界帶 | 白邊／色邊／殘景 |
| 邊緣品質 | alpha 梯度、硬切比例 | 是否自然／過度削切 |
| 原圖保真 | 未授權區像素差（需對齊幾何） | 主體是否被改臉／改色 |

棋盤格預覽 **不可** 替代 alpha 機械檢查。

## Check result schema

每條檢查：

| 欄位 | 值 |
| --- | --- |
| `check_id` | 穩定 id，如 `alpha.present` |
| `status` | `pass` \| `fail` \| `inconclusive` |
| `evidence` | 路徑、統計摘要、可覆核觀察 |
| `failure_class` | 見 [`06`](06-failure-taxonomy-and-repair.md)；僅 fail／部分 inconclusive |
| `repair_candidate` | 建議下一策略 id 或 `none` |
| `review_required` | bool |

## Finality 規則（草案）

| 條件 | finality |
| --- | --- |
| 全部必選 check = pass，且無 review_required | `accepted` |
| 必選 pass，但有標記需人審的 soft check | `accepted_with_review`（人簽前不得當生產完成） |
| 任一硬 check = fail | 不得 `accepted`；進 repair 或 `rejected` |
| 關鍵 check = inconclusive | 不得偽造成 pass；`human_review` 或收集更多證據 |

## 獨立性要求

| 允許 | 禁止 |
| --- | --- |
| 第二程序讀輸出檔做統計 | 執行腳本結尾 `print("OK")` 當驗收 |
| 人審對照白／黑／棋盤預覽 | 同一 LLM 回合既執行又宣稱「完美」 |
| 合成參考圖（已知前景）回歸 | 無真值卻用面積閾值當 PASS 證明 |

## Gate（Phase 2）

- schema 固定；至少 1 個負向 fixture：工具產出 PNG 但 `alpha`／`residue` fail
- 文件寫明：何種情況必須 `inconclusive`
