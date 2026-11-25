# Plan Evidence Index — `2026-11-25-0842-verification-architecture-contract`

本目錄存放本 plan 的量測與 dogfood 全文。`_plan.md` 只留摘要與連結。

## 引用規則（避免行號漂移）

| 做法 | 說明 |
|---|---|
| **用檔案路徑** | `evidence/<file>.md` 或相對連結；不要寫行號 |
| **用標題錨點** | 定位寫 `evidence/foo.md` 內的章節標題 |
| **專案細節** | 外部 consumer 的細節不放這裡；本目錄只放 Ai-skill 自身量測與去識別化摘要 |
| **新 run** | 新增 `evidence/<run-id>-<slug>.md` + 同一 commit 更新本表 |

Canonical 規則：[`governance/lifecycle/plan-evidence.md`](../../../../governance/lifecycle/plan-evidence.md)

## Run 索引

| Run ID | 檔案 | 狀態 | 摘要 |
|---|---|---|---|
| S1 | [`s1-ai-skill-hook-timing-baseline.md`](s1-ai-skill-hook-timing-baseline.md) | 完成 | Ai-skill 自身 hooks 耗時基線：git hooks 空跑 < 0.2 s；`go test` 未快取約 25 s、Go 內建快取約 2 s；PreToolUse 每次工具呼叫約 0.32 s 且隨 transcript 變大；測試選擇的路徑範圍有 V2 缺口（observed） |
