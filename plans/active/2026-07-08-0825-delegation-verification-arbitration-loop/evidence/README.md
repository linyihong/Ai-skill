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
| **2h** | [`2h-externalrepoc-common-url-verification-gaps.md`](2h-externalrepoc-common-url-verification-gaps.md) | 證據 only | ExternalRepoC 03 common-url Execute：RBAC 三连漏网、V5 仅 list、combined defer L1–L3、api-surface gate 回饋 |
| **2i** | [`2i-externalrepoc-user-feedback-pull-execute.md`](2i-externalrepoc-user-feedback-pull-execute.md) | 證據 only | ExternalRepoC 04 user-feedback S0–S4 Execute：2h 教训迁移、Stop/resume、inventory gate、sync_jobs 分表、mapping defer |
| **2j** | [`2j-externalrepoc-push-execute-skip-verifier-loop.md`](2j-externalrepoc-push-execute-skip-verifier-loop.md) | 负向证据 | ExternalRepoC 05 push Execute：**0 Verifier**、单 Task 包办、`delegation.enabled:false` 误豁免；consumer verifier-after-executor gate 回饋 |
| **2k** | [`2k-externalrepoc-push-post-close-runtime-gaps.md`](2k-externalrepoc-push-post-close-runtime-gaps.md) | 證據 only | ExternalRepoC 05 push **2j 纠偏后**：slice 关闭 vs 用户手验（模版/商户/远程同步）；Worker 拓扑、pull 映射、post-close surgical debt |
| **2l** | [`2l-externalrepoc-common-url-s2-mirror-skip-loop.md`](2l-externalrepoc-common-url-s2-mirror-skip-loop.md) | 负向证据 | ExternalRepoC 03 S2′ mirror：**0 Executor/Verifier**、surgical bypass 滥用、Shell 绕过 preToolUse；2j/2k 教训未内化 |
| **2m** | [`2m-externalrepoc-phase-g-mirror-batch-retrofit.md`](2m-externalrepoc-phase-g-mirror-batch-retrofit.md) | 正负对照 | ExternalRepoC **Phase G-mirror** 批量 retrofit：V-m1–V-m5 + 登记总表；02/01 合规 loop vs 03/2l；stale JVM V5-A 复发 |
| **2n** | [`2n-externalrepoc-push-delivery-s1-s6-compliant-loop.md`](2n-externalrepoc-push-delivery-s1-s6-compliant-loop.md) | 正向证据 | ExternalRepoC **07 push delivery** DEL-S1–S6：6/6 E+V loop、sub-plan `completed`、零 post-close bypass；对照 2j/2k/2l |
| **2o** | [`2o-consumer-tab-scroll-single-vs-delegation.md`](2o-consumer-tab-scroll-single-vs-delegation.md) | 正负对照 | `<PROJECT_ROOT>` `/h5` tab-scroll：**单 session 部分关** vs **三角色全 linked**；authority model 桥接 State Trust；deploy smoke≠L3 |
| 2d | —（inline） | 證據 only | 見 [`01-dogfood-prompt-kit.md`](../01-dogfood-prompt-kit.md) → `### 2d` |
| 2c | —（inline） | 證據 only | 見 kit → `### 2c` |
| 2g | —（inline） | 證據 only | 見 kit → `### 2g` |
| 2e | [`2e-grandfather-sunset-audit.md`](2e-grandfather-sunset-audit.md) | 完成 | Research 域 grandfather sunset；Q6/Q7/Q8 跨域观察 |
| 2f | —（inline，預註冊） | 進行中 | 見 kit → `### 2f`；產出候選 [`03-repo-naming-candidates.md`](../03-repo-naming-candidates.md) |
| 2b | —（inline） | 完成 | 見 kit → `### 2b` |
| 2a / 2a-external | —（inline） | 完成 | 見 kit → `### 2a` / `### 2a-external` |

> **漸進遷移**：2026-07-09 起新證據進本目錄；舊 run 仍 inline 於 kit，後續可按需拆檔，拆後更新本表「檔案」欄。
