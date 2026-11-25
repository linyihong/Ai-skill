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
| S2 | [`s2-v2-path-selection-falsification.md`](s2-v2-path-selection-falsification.md) | 完成 | V2 缺口 verified：只改 `metadata/` 或 `plans/` 的兩個探針都讓測試失敗，pre-push 仍放行且不跑 Go 測試；CI paths 過濾同樣不涵蓋。已修補：pre-push 任何變更都跑 Go 測試，CI 移除 paths 過濾。旁支：pre-commit 在 git worktree 中無法執行 |
| S3 | [`s3-pretooluse-incremental-parse-evaluation.md`](s3-pretooluse-incremental-parse-evaluation.md) | 完成（評估；步驟 0 已實作） | PreToolUse 在 5.1 MB 時每次 0.52 s，與 transcript 大小成正比；主因是同一次呼叫內全量解析兩次。步驟 0（只解析一次）已實作，中位數 0.62 → 0.40 s；之後以 shadow 做跨呼叫增量（步驟 1），採用門檻已登記 |
