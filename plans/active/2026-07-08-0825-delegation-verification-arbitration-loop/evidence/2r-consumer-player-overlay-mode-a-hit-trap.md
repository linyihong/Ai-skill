# 2r — <PROJECT_ROOT> player overlay：Mode A 单 session hit-trap 复发（2026-07-13）

> **专案证据边界**：`server_doc` tip、L3 exit code、release stamp 留于 `<PROJECT_ROOT>`；本档只保留 generalized metrics 与契约回馈。

## Run 摘要

- **任务**：H5 player dock「打开选集」被 TabBar 挡住 → 修 soft-nav → 用户 cold `/player` URL 全控件不可点。
- **模式**：**Mode A**（单 session 直连，无 Orchestrator/Executor/Verifier 分角色）——与 **2o** 同 family。
- **结果**：soft-nav hit-test 绿后宣称可验；cold URL 暴露 empty `@player` overlay + `pointer-events:none` cached pane。

## 假绿机制

| 证据 | 为何不够 |
|---|---|
| `querySelector().click()` / 旧 drama-nav L3 | 绕过命中检测 |
| 仅 soft-nav `elementFromPoint` | partial authority — 只证明刚修路径 |
| `Boolean(player)` 判 overlay | empty placeholder 仍 truthy |

## 契约回馈（已写回 consumer）

1. Execute 前 backfill **入口矩阵**（home soft-nav / search soft-nav / cold URL）
2. `verifier_only` + `elementFromPoint`；禁止仅 programmatic click
3. overlay 激活用 `useSelectedLayoutSegment('player')`，非 `Boolean(player)`
4. Mode A/B 对照写入 `plan-delegation-execution-loop.md`；workflow failure_modes 增补

## 量测栏

| 指标 | 值 |
|---|---|
| Verifier Task spawn | **0** |
| 用户纠偏次数 | **2**（soft-nav 测不够 → cold URL；再要求补测） |
| 与 2o 同构 | **是**（partial authority / single-session） |
| 与 2j 同构 | 否（非「0 Verifier 因 enabled:false」；是 surgical Mode A） |

## 关联

- 2o Mode A/B 对照
- consumer feedback：`programmatic-click-and-single-entry-path-miss-hit-trap.md`
- L3：`player-episode-sheet-hit-target.integration.mjs`
