# Plan Evidence Index — `evidence-record-contract`

本目錄存放本 plan 的收斂審計全文。`_plan.md` 只留摘要與連結。

## 引用規則（避免行號漂移）

| 做法 | 說明 |
|---|---|
| **用檔案路徑** | `evidence/<file>.md` 或相對連結；**不要**寫行號 |
| **用標題錨點** | 定位寫 `evidence/foo.md` 內的 `## 標題` |
| **新 run** | 新增 `evidence/<run-id>-<slug>.md` + 在同一 commit 更新本表 |

Canonical 規則：[`governance/lifecycle/plan-evidence.md`](../../../../governance/lifecycle/plan-evidence.md)

## Run 索引

| Run ID | 檔案 | 狀態 | 摘要 |
|---|---|---|---|
| phase0 | [`phase0-convergence-audit.md`](phase0-convergence-audit.md) | done（2026-10-04） | 5 個 domain、4 條 invariant + lifecycle／epistemic 檢查點。I1=3、I2=4、I4=5 → Promote；I3=2+variant → Watch；I5 候選=3 → 待 Phase 1 決定 |
