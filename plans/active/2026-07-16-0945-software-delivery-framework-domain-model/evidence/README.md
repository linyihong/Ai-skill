# Evidence — Software Delivery Framework Domain Model

Plan: [`../_plan.md`](../_plan.md)

## 引用規則

- 正文只留摘要 + 本目錄路徑連結；禁止行號引用（用 `evidence/<file>.md` 路徑）。
- Run 索引見下表。本機 consumer 綁定（專案名、絕對路徑）寫 repo 根 `local/plan-evidence/`（gitignored），不進索引。

## Run 索引

| Run ID | 檔案 | 說明 |
| --- | --- | --- |
| phase-0-2026-07-16 | [`phase-0-classification-matrix.md`](phase-0-classification-matrix.md) | Primary Model 驗證：完整分類矩陣、N 判定、否證記錄、Domain 圖 |
| phase-2-2026-07-16 | [`phase-2-dogfood.md`](phase-2-dogfood.md) | 3 個新 artifact 問答 + placement dual-write grep |
| phase-3-external-greenfield-consumer | [`phase-3-external-greenfield-consumer.md`](phase-3-external-greenfield-consumer.md) | 外部 greenfield consumer 分類 + ERA 接點（去敏；本機邊界見 `local/plan-evidence/`） |
| phase-3-external-greenfield-consumer-execute | [`phase-3-external-greenfield-consumer-execute.md`](phase-3-external-greenfield-consumer-execute.md) | classify-before-create + artifact audit（preflight，未 E+V） |
