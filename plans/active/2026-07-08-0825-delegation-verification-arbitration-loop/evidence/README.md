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
| **2a** | [`2a-software-delivery-review-invoke.md`](2a-software-delivery-review-invoke.md) | 完成 | SD Review invoke 整合備註；violation 0；transport adaptation（Task subagent） |
| **2b** | [`2b-plans-sop-expansion.md`](2b-plans-sop-expansion.md) | 完成 | plans §Delegation SOP 擴充；fix 1（tool-neutral）；Q1 正向證據 |
| **2a-external** | [`2a-external-sync-adapter-step6.md`](2a-external-sync-adapter-step6.md) | 完成 | 外部 sync-adapter Step 6；6 tests；pre-loop orchestrator 越界 1 |
| **2c** | [`2c-tiered-archive-platform.md`](2c-tiered-archive-platform.md) | 完成 | tiered archive 8 slices；violation 2/8；多 slice loop 證據 |
| **2d** | [`2d-outbound-sync-phase3.md`](2d-outbound-sync-phase3.md) | 完成 | outbound sync Phase 3；consumer overlay + slice_kind + backfill |
| **2d′** | [`2d-prime-externalrepoc-module-alignment.md`](2d-prime-externalrepoc-module-alignment.md) | 證據 only | ExternalRepoC 9j2 模組 01/02 對齊 follow-on；integration gate、remote_absent_delete |
| **2h** | [`2h-externalrepoc-common-url-verification-gaps.md`](2h-externalrepoc-common-url-verification-gaps.md) | 證據 only | ExternalRepoC 03 common-url Execute：RBAC 三连漏网、V5 仅 list、api-surface gate |
| **2e** | [`../02-grandfather-sunset-audit.md`](../02-grandfather-sunset-audit.md) | 完成 | Research 域 grandfather sunset |
| **2f** | —（預註冊） | 進行中 | 產出候選 [`../03-repo-naming-candidates.md`](../03-repo-naming-candidates.md) |
