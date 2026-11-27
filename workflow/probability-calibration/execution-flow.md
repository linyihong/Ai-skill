# Probability Calibration — Execution Flow

Canonical lifecycle。欄位 SoT 在 [`contracts/`](contracts/)；本檔只定義步驟順序與禁止事項。**不要**寫具體引擎、語言、優化器或 consumer 識別字。

## Lifecycle

```text
0. Bind Calibration Contract    → targets（C1 角色）+ metrics（定義／tier／estimator／precision_target／tolerance）
                                  沒有凍結的契約 → 不得調參
1. Gather Reference Evidence    → 參照來源、樣本範圍、uncertainty；熱／農場樣本只取尺度不變結構
2. Metric Parity                → 兩側指標逐一同定義（分子、分母、基準回合、聚合單位、比較統計量、樣本單位、比較單位數 n）
                                  順序統計量（最大值、極端分位數）只在相同 n 下比較
3. Diagnose Deviation           → 依序排除，見 §診斷順序；只有「產品真的偏」成立才能進 4
   3s. Structural Ceiling exit  → 結構無法表達目標 → 寫 gap-ledger.structural_ceiling，停止調參
4. Tune (executor)              → 只對 implementation target 收斂；training／validation seed；
                                  重要實驗記入 experiment_ledger；每調一旋鈕量全部指標
5. Freeze Candidate             → 參數快照 + hash；之後才揭露 final_holdout_seed_set
6. Independent Verify           → verifier 依 verification.yaml；記錄 independence_achieved
7. Classify Status              → executor 只填 executor_assessment（初判）；verifier／decider 依
                                  calibration-record.evaluation_order 產生 calibration_status（互斥、first match wins）
8. Gap Ledger + Record Closeout → 所有已知缺口入 ledger；status 附 reason 與證據
9. Feedback                     → FeedbackDecision；新失敗進 registry／evidence candidate（N≥3 才升級）
```

**契約變更**：凍結後新增、刪除或修改 Target、必要指標、精度要求、容差或指標定義，都要升 `contract_version` 並重新凍結；舊證據保留，但不是新版本下的驗收證據；重跑範圍依變更影響決定並記錄理由（見 contract §`change_control`）。

## 診斷順序（I3，硬性契約）

旗標出現時依序排除，**不得跳步**：

| 序 | 假設 | 適用條件 | 成立時 |
| --- | --- | --- | --- |
| a | 指標定義不一致 | 一律 | 修定義或對照表，重量；**不改參數** |
| b | 採集或樣本偏差 | **僅當 Target 具 `reference_observation`**（規格型 Target 無此步） | 重新界定參照樣本範圍；**不改參數** |
| c | 產品實際行為偏差 | 一律 | 才可進步驟 4 |

不得只因 live 頻率與產品不同就改賠付或權重表；必須先走完 a（與適用時的 b）。b 被略過時，Record 必須寫明「Target 無 `reference_observation`，b 不適用」。

**契約前已調參**：視為 `diagnosis_skipped_tuned_payout` 與 Gap Ledger `process_gap`；回退該變更或把它凍結為 baseline，再從步驟 0 重新開始，不得把契約前的結果當作本輪證據。

## Structural Ceiling（C7）

調參在合理範圍內持續無法逼近目標時，先問「目前結構能不能表達這個分布」，而不是繼續擴大搜尋。判定須補齊 [`gap-ledger.yaml`](contracts/gap-ledger.yaml) §`structural_ceiling` 全部欄位；欄位不齊時 status 只能是 `inconclusive`，不得判 `structural_blocked`。搜尋得到的前緣只證明「這個搜尋方法找到的最好結果」，只能當診斷證據；`why_not_parameter_not_found` 必須是解析上界、縮減空間窮舉或其他可重現的結構性論證。

## 責任分離

| 階段 | 負責 | 禁止 |
| --- | --- | --- |
| Contract | decision owner（沿用 [`decision-support`](../cross-cutting/decision-support/README.md) provisional→confirmed） | executor 自訂目標或門檻 |
| Tune | executor（含優化器） | 使用 `final_holdout_seed_set`；寫最終 status |
| Verify | verifier | 共用 executor 的自評作為唯一依據 |
| Accept | decider | 由 executor 兼任 |

## 驗證分層

| 層 | 證明 | 不證明 |
| --- | --- | --- |
| 釘值（固定 seed 回歸鎖） | 實作沒有漂移 | 已校準 |
| 長跑量測（附 SE） | 指標落在 precision_target 內 | 對未見 seed 成立 |
| Holdout 驗收 | 本輪校準宣稱 | 下一輪（新 holdout 需重建） |

鏈式驗證中任一機械鎖失敗＝整條鏈失敗，不得口頭沖掉。

## 禁止

- 沒有 calibration contract 就調參
- 以單一加總分數判定成功
- 拿 final holdout 參與搜尋、或驗收失敗後對同一組 holdout 重調
- 把釘值測試當作校準證據
- 單次搜尋失敗就宣稱 `structural_blocked`
- 引用舊文件的頻率或結論而不回查來源與日期
