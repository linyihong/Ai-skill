# Software Delivery Implementation（`sd-implementation`）

Implementation stage 在 requirements / architecture / validation target 足夠明確後開始。本目錄是 **`sd-implementation`** 的 canonical surface（retained slice，不另開 top-level lifecycle stage）。

## 載入指引

| 任務 | 讀什麼 |
|------|--------|
| 選 execution mode、Change Intent Lock、structure preparation、stop condition | [`execution-modes.md`](execution-modes.md) |
| **多倉庫處理模式**（orchestrator + sibling repos、隊友 merge 後接 sync / adapter 切片） | [`multi-repository-workspace-mode.md`](multi-repository-workspace-mode.md) |
| SDK 缺陷閉環、同工作階段閉環（code + 持久文件） | [`execution-flow.md`](../execution-flow.md) §3–§4 |
| 外科手術式 diff 紀律（feature / direct_change） | [`surgical-changes.md`](../surgical-changes.md) |
| intent → validation 分流 | [`test-strategy.md`](../test-strategy.md) |
| Post-implementation self review | invoke `code-review` capability — 見 [`README.md`](../README.md) §Review invoke |

## 快速選型

```text
change_kind: feature + blocked_by_structure?
  yes → execution_mode: preparatory_refactoring → 讀 execution-modes.md
  no  → execution_mode: direct_change（預設）
```

Plan：[`plans/active/2026-06-29-1430-preparatory-refactoring-workflow/_plan.md`](../../../plans/active/2026-06-29-1430-preparatory-refactoring-workflow/_plan.md)
