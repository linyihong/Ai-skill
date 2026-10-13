# Image Processing — Execution Flow

Canonical lifecycle。欄位 SoT 在 [`contracts/`](contracts/)。**不要**寫 provider／CLI 逐步碼／固定權重名。

## Lifecycle

```text
0. Bind Image Task Contract     → intent + background + fidelity 三維 + immutable
1. Input Validation             → decode、EXIF、content hash；缺背景意圖 → ask
2. Image Task Analysis          → subject_count / edge_type / risk_flags / confidence
3. Apply Constraints            → contract + consumer_profile + license → Feasible Candidates
4. Selection Policy             → 在可行集內選 candidate 組合；寫入 IDR
5. Execute (adapter)            → mask/matte/assemble；產出輸出 + 預覽（預覽≠證據）
6. Independent Verify           → check_id / status / failure_class / hardness
7. Disposition                  → repair | human_review | abort | close
8. Repair Loop (bounded)        → max_repair_attempts；必須換 candidate 或升級人審
9. Finality + IDR Closeout      → accepted | accepted_with_review | rejected | aborted
```

分析結果是 **路由依據**，不是不可推翻真相。低信心 → ask／human_review。

## 責任分離

| 階段 | 負責 | 禁止 |
| --- | --- | --- |
| Constraints | 契約／profile／授權裁可行集 | Selection Actor 自訂驗收 |
| Selection | 在可行集內選 | 選可行集外工具 |
| Execute | Adapter 產檔 | 寫 `status: pass` |
| Verify | 獨立程序／人審 | 複用執行器自評當唯一依據 |

## Finality（對齊 [`contracts/finality.yaml`](contracts/finality.yaml)）

| Finality | 條件（摘要） |
| --- | --- |
| `accepted` | 硬 check 全 pass；無 `review_required` |
| `accepted_with_review` | 硬 pass；軟 check 需人簽（生產完成前不得當 done） |
| `rejected` | 硬 fail 且 repair 用盡／無候選 |
| `aborted` | 契約衝突不可修，或 `fidelity_violation` 需停生成式路徑 |

`human_review`／`repair` 是 **disposition**，不是 finality 枚舉值。

## 禁止

- 棋盤格預覽代替 alpha 機械檢查
- 無真值卻用遮罩面積單獨證明 PASS
- 對 `immutable` 內部做 morphology opening「清乾淨」
- 無限重跑同一 `selected` candidate
- 在本檔寫死單一分割供應商為唯一路徑

## Fixtures / dogfood

案例規格與負向矩陣：plan [`fixtures/`](../../plans/active/2026-10-14-0832-image-processing-decision-workflow/fixtures/README.md)；本目錄 [`fixtures/`](fixtures/README.md) 為 Phase 1+ canonical 目標位置。
