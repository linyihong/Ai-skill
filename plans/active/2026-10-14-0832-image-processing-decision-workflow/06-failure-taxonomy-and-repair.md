# 06 — Failure Taxonomy & Repair Policy

← [`_plan.md`](_plan.md)

## Failure classes（P0 最小集）

| failure_class | 含義 | 典型證據 |
| --- | --- | --- |
| `subject_missing` | 主體大部消失 | 遮罩面積過低、連通域崩壞 |
| `wrong_subject` | 選到錯物件／背景塊 | 多主體、位置偏移 |
| `background_residue` | 背景殘留 | 遮罩外高能量像素 |
| `edge_halo` | 深色／彩色光暈邊 | rim 統計、視覺色邊 |
| `edge_detail_loss` | 髮絲／薄結構被削 | 與原圖邊緣差、風險區消失 |
| `interior_damage` | 主體內部被誤挖／誤填 | 牙齒洞、配件孔被填死或咬穿 |
| `alpha_invalid` | 無 alpha／假透明 | mode、alpha 直方圖 |
| `fidelity_violation` | 未授權改主體 | 像素差、生成痕跡 |
| `contract_incomplete` | 契約不足仍執行 | 缺欄位 |
| `tool_error` | 執行例外／逾時 | stderr、exit |

可在 Phase 3 擴充；升 registry 前需人覆核。

## Repair policy（對應，非唯一）

| failure_class | 優先 repair_candidate | 禁止 |
| --- | --- | --- |
| `wrong_subject` | `seg.interactive`／ask subject_hint | 只重跑同一 generic seg |
| `background_residue` | 更緊 mask／decontaminate | 無界 erode 主體 |
| `edge_halo` | decontaminate、rim 重上色（限 rim） | 對內部做 opening |
| `edge_detail_loss` | 換 matting／降低二值化 | 更大 erode |
| `interior_damage` | 回退激進 morphology；保護內部 | 再 opening 一次 |
| `subject_missing` | 放寬閾值或換模型；否則 human_review | 盲目 dilate 背景 |
| `alpha_invalid` | 重組裝 RGBA | 用 JPG 當完成 |
| `fidelity_violation` | abort 生成式路徑 | 「再 generate 修臉」 |

## 重試與終止

| 規則 | 值（草案） |
| --- | --- |
| `max_repair_attempts` | 2（同一 case） |
| 同一 `selected` 連續失敗 | 必須換 candidate 或升級 human_review |
| 契約與結果衝突且不可修 | `aborted` + 保留證據 |
| 人審佇列 | 標記 `review_required`；不自動 `accepted` |

## Failure learning → guards

僅當：

1. 真實案例失敗
2. 人覆核確認根因
3. 抽象成可重用 guard（去專案細節／去敏感路徑）

才可寫入 `workflow/image-processing/registry/` 或 `enforcement/failure-patterns/`。

**禁止**：把單次模型口頭判斷寫成永久規則。

## Gate（Phase 2–3）

- 每個 verify fail 必須帶 `failure_class`
- 修復日誌可看出「換了策略」而非「同參重跑」
