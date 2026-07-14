# Plan Evidence Index — `<plan-slug>`

本目錄存放本 plan 的 **dogfood / 量測 / 契約回饋** 全文。`_plan.md` 與 companion 檔（如 `01-dogfood-prompt-kit.md`）只留摘要與連結。

> **主計畫吸收**：使用 `evidence/` 時，主紀錄必須在資料夾內的 `_plan.md`；禁止頂層仍留 `<slug>.md`。

## 引用規則（避免行號漂移）

| 做法 | 說明 |
|---|---|
| **用檔案路徑** | `evidence/<file>.md` 或相對連結；**不要**寫「kit L449」類行號 |
| **用標題錨點** | 定位寫 `evidence/foo.md` 內 `## 量測欄` 或表格 `#` 欄 |
| **專案細節** | inner commit、class 名、live 環境 → consumer `<PROJECT_ROOT>` plan；本目錄只留 generalized metrics（[`enforcement/sanitization.md`](../../../enforcement/sanitization.md)） |
| **新 run** | 新增 `evidence/<run-id>-<slug>.md` + 更新本表同一 commit |

Canonical 規則：[`governance/lifecycle/plan-evidence.md`](../../../governance/lifecycle/plan-evidence.md)

## Run 索引

| Run ID | 檔案 | 狀態 | 摘要 |
|---|---|---|---|
| _example_ | [`_example-record.md`](_example-record.md) | 模板 | 刪除或替換為第一筆真實 run |
