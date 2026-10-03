# Plan Evidence Index — `2026-10-03-2104-security-audit-capability-hardening`

本目錄存放本 plan 的 **dogfood / 契約回饋** 全文。`_plan.md` 只留摘要與連結。

## 引用規則（避免行號漂移）

| 做法 | 說明 |
|---|---|
| **用檔案路徑** | `evidence/<file>.md` 或相對連結；**不要**寫行號 |
| **用標題錨點** | 定位寫 `evidence/foo.md` 內的 `## 標題` 或表格欄 |
| **專案細節** | finding list、coverage ledger、commit、class 名 → consumer `<PROJECT_ROOT>`；本目錄只留去敏後的結論（[`enforcement/sanitization.md`](../../../../enforcement/sanitization.md)） |
| **新 run** | 新增 `evidence/<date>-dogfood-<n>-<slug>.md` + 更新本表，同一 commit |

Canonical 規則：[`governance/lifecycle/plan-evidence.md`](../../../../governance/lifecycle/plan-evidence.md)

## Run 索引

| Run ID | 檔案 | 狀態 | 摘要 |
|---|---|---|---|
| dogfood-1 | [`2026-10-03-dogfood-1-refresh-family.md`](2026-10-03-dogfood-1-refresh-family.md) | done（`verdict_kind: expected`） | consumer 專案 refresh-token family 內部切片的事後 audit；fresh verifier 抓到 reachability 錯誤與 deferral 冒充 risk acceptance，template / ledger 契約已修 |
| dogfood-2 | — | pending | 同專案下一切片加入第一個使用者時，驗證 coverage invalidation 實際觸發 |
