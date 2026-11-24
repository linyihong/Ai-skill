# Reverse-Engineering Workflow

跨目標 shipped-artifact 分析的**執行入口**（授權 → 分流 → 取證 → 產出／handoff）。深度方法在 `analysis/`；本層只編排順序與 gates。

APK **動態**端到端仍以 [`../apk-analysis/`](../apk-analysis/README.md) 為主。本 workflow 處理跨族分流，並在 Static APK 時可切入 apk-analysis 的靜態分支。

| 文件 | 用途 |
| --- | --- |
| [`execution-flow.md`](execution-flow.md) | 執行順序 |
| [`artifact-gates.md`](artifact-gates.md) | 產出與 Evidence 最低門檻 |

工具治理：分析過程中要新增或修改腳本、探針、查詢或檢查之前，先讀 [`../cross-cutting/project-tooling/`](../cross-cutting/project-tooling/README.md)（先找再建、Entry／Logic／Parameters、放置與生命週期、鎖分類）。這是 pilot 參考契約，不是 gate。

計劃：[`../../plans/active/2026-10-10-1328-rea-analysis-capability-hardening/_plan.md`](../../plans/active/2026-10-10-1328-rea-analysis-capability-hardening/_plan.md)

← [workflow/](../README.md)
