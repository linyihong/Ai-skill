# Dogfood 證據索引（Delegation Verification & Arbitration Loop）

本目錄存放 **dogfood 量測與契約回饋** 的全文；[`01-dogfood-prompt-kit.md`](../01-dogfood-prompt-kit.md) 保留 **傳輸模板 A/B/C** 與精簡指標。

> **Canonical 規則**：[`governance/lifecycle/plan-evidence.md`](../../../governance/lifecycle/plan-evidence.md)（commit-msg `validatePlanEvidenceConvention` 機械強制）

## 引用規則（避免行號漂移）

| 做法 | 說明 |
|---|---|
| **用檔案路徑** | `evidence/<slug>.md` 或相對連結，**不要**寫「kit L449」類行號 |
| **用標題錨點** | 需要定位時寫 `evidence/foo.md` 內 `## 量測欄` 或表格 `#` 欄，不用絕對行號 |
| **專案細節** | inner commit、class 名、live 環境 → 留 consumer `<PROJECT_ROOT>` plan §執行紀錄；本目錄只留 generalized metrics（[`enforcement/sanitization.md`](../../../../enforcement/sanitization.md)） |
| **新 run** | 一律新增 `evidence/<run-id>-<slug>.md` + 更新本索引；kit 只留一行指標摘要 + 連結 |

## Run 索引

| Run ID | 檔案 | 狀態 | 摘要 |
|---|---|---|---|
| **2d′** | [`2d-prime-externalrepoc-module-alignment.md`](2d-prime-externalrepoc-module-alignment.md) | 證據 only | ExternalRepoC 9j2 模組 01/02 對齊 follow-on；integration gate、remote_absent_delete、live teardown、release-time gate |
| 2d | —（inline） | 證據 only | 見 [`01-dogfood-prompt-kit.md`](../01-dogfood-prompt-kit.md) → `### 2d` |
| 2c | —（inline） | 證據 only | 見 kit → `### 2c` |
| 2g | —（inline） | 證據 only | 見 kit → `### 2g` |
| 2e | [`02-grandfather-sunset-audit.md`](../02-grandfather-sunset-audit.md) | 完成 | Research 域 grandfather sunset；kit 摘要 → `### 2e` |
| 2f | —（inline，預註冊） | 進行中 | 見 kit → `### 2f`；產出候選 [`03-repo-naming-candidates.md`](../03-repo-naming-candidates.md) |
| 2b | —（inline） | 完成 | 見 kit → `### 2b` |
| 2a / 2a-external | —（inline） | 完成 | 見 kit → `### 2a` / `### 2a-external` |

> **漸進遷移**：2026-07-09 起新證據進本目錄；舊 run 仍 inline 於 kit，後續可按需拆檔，拆後更新本表「檔案」欄。
