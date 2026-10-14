# 05 — Independent Verification Contract

← [`_plan.md`](_plan.md)

## 原則

1. **執行成功 ≠ 品質成功**
2. 驗收器不得複用執行器的「自評分數」作為唯一依據
3. 無前景真值時，統計指標最多做 **異常偵測**，不得單獨證明品質良好
4. `inconclusive` 是一等公民狀態

## 驗收維度

| check 族 | 機械檢查 | 視覺／語意檢查 | 硬／軟（P0 草案） |
| --- | --- | --- | --- |
| 檔案有效性 | 可解碼、尺寸、格式 | 通常不需要 | 硬 |
| 透明通道 | alpha 存在、直方圖／非全不透明 | 透明區是否符合意圖 | 硬（含「假透明」） |
| 主體完整性 | 遮罩面積、連通域異常 | 缺件／缺髮／缺配件 | 硬（異常偵測）；語意為軟／人審 |
| 背景殘留 | 遮罩外亮度／彩度統計、邊界帶 | 白邊／色邊／殘景 | 硬（異常）；視覺為軟 |
| 邊緣品質 | alpha 梯度、硬切比例 | 是否自然／過度削切 | 軟（無真值時）或硬（有合成參考） |
| 原圖保真 | 未授權區像素差（需對齊幾何） | 主體是否被改臉／改色 | 硬（相對原圖／授權 rim） |
| Consumer profile | canvas／置中／縮放規格 | 臉軸／silhouette（若 profile 要求） | 硬（規格欄）；對齊視覺可軟 |

棋盤格預覽 **不可** 替代 alpha 機械檢查。

## Check result schema

每條檢查：

| 欄位 | 值 |
| --- | --- |
| `check_id` | 穩定 id，如 `alpha.present`、`alpha.not_fully_opaque` |
| `status` | `pass` \| `fail` \| `inconclusive` |
| `evidence` | 路徑、統計摘要、可覆核觀察 |
| `failure_class` | 見 [`06`](06-failure-taxonomy-and-repair.md)；僅 fail／部分 inconclusive |
| `repair_candidate` | 建議下一策略 id 或 `none` |
| `review_required` | bool |
| `hardness` | `hard` \| `soft` |

## Finality 與處置轉換（草案）

**Check status** 與 **finality** 分離。`human_review` 是 **disposition（處置）**，不是 finality 枚舉值。

| 條件 | Disposition | Finality（IDR） |
| --- | --- | --- |
| 全部硬 check = `pass`，軟 check 無 `review_required` | `close` | `accepted` |
| 全部硬 check = `pass`，但有軟 check `review_required` | `queue_human_review` | `accepted_with_review`（人簽前不得當生產完成） |
| 任一硬 check = `fail`，且仍有未用盡的不同 `repair_candidate` | `repair` | （尚未 final；IDR 記 intermediate） |
| 任一硬 check = `fail`，repair 用盡或無候選 | `reject_or_abort` | `rejected` 或 `aborted` |
| 關鍵硬 check = `inconclusive` | `human_review` 或 `gather_evidence` | 不得偽造成 `accepted`；人審後才可寫 finality |
| 契約與結果衝突且不可修 | `abort` | `aborted` |

規則摘要：

1. `fail`（硬）→ 不得 `accepted`。
2. `inconclusive`（關鍵）→ 不得當成 `pass`；最多停在 `accepted_with_review` 或保持未 final。
3. 無真值的面積／梯度閾值 → 最多 `inconclusive` 或異常 `fail`，**不得單獨**證明品質 `pass`。

## 負向 fixture 矩陣（Phase 2 Gate 必測）

每類至少 1 條；工具可產出 PNG，但 verify 必須擋住。對應 [`06`](06-failure-taxonomy-and-repair.md)。

| Fixture id | 注入失敗 | 必 fail／inconclusive 的 check | 期望 `failure_class` |
| --- | --- | --- | --- |
| `neg.alpha-opaque` | 有 alpha 通道但實質全不透明（或幾乎全 255） | `alpha.not_fully_opaque` | `alpha_invalid` |
| `neg.residue-halo` | 背景大致清掉仍有 rim 色暈 | `residue.rim`／背景殘留族 | `edge_halo` 或 `background_residue` |
| `neg.interior-damage` | 牙／孔洞／薄結構被咬穿或誤填 | 保真／主體完整性（相對原圖或標註 ROI） | `interior_damage` |
| `neg.over-erode` | 邊「乾淨」但主體過度侵蝕 | 主體完整性／邊緣（與 residue **正交**：residue 可 pass） | `subject_missing` 或 `edge_detail_loss` |
| `neg.consumer-layout` | 去背可接受但 canvas／置中／縮放違規 | consumer profile 規格 check | 契約／profile 違規（可記 `fidelity_violation` 或 profile-specific class） |

Phase 2 最小負向：至少 `neg.alpha-opaque` 與 `neg.residue-halo`。
Phase 3：上表五類皆可執行（組裝 `ai-skill image-processing verify --expect-block`；見 [`tools/README.md`](../../../workflow/image-processing/tools/README.md)）。
## 獨立性要求


| 允許 | 禁止 |
| --- | --- |
| 第二程序讀輸出檔做統計 | 執行腳本結尾 `print("OK")` 當驗收 |
| 人審對照白／黑／棋盤預覽 | 同一 LLM 回合既執行又宣稱「完美」 |
| 合成參考圖（已知前景）回歸 | 無真值卻用面積閾值當 PASS 證明 |

## 必須 `inconclusive` 的情況（文件化）

- 無前景真值，且統計未達「異常」門檻，但視覺風險高（髮絲／玻璃／光暈）
- 多主體且 `subject_hint` 缺失、分析低信心
- Consumer 對齊依賴外部 Sprite，但參考檔缺失或 hash 不匹配
- 機械 check 與人審預覽明顯衝突且無法自動裁決

## Gate（Phase 2）

- [x] schema 固定（含 `hardness`）— [`verification.yaml`](../../../workflow/image-processing/contracts/verification.yaml) + verify tool
- [x] finality／disposition 轉換表可操作 — [`finality.yaml`](../../../workflow/image-processing/contracts/finality.yaml)
- [x] 負向 fixture：`neg.alpha-opaque`、`neg.residue-halo` 可重現擋住 — `ai-skill image-processing verify … --expect-block`
- [x] 文件寫明：何種情況必須 `inconclusive`（本檔上節）
