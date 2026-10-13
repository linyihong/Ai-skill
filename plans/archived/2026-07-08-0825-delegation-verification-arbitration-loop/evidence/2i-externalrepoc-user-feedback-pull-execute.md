# 2i — ExternalRepoC user-feedback pull Execute：多 slice 闭环与 Stop 恢复（2026-07-09，證據 only）

> **專案證據邊界**：inner commit hash、class 名、migration 文件名留于 `<PROJECT_ROOT>` sub-plan 04 §10–§15 与 `preflight-feedback-log.md`；Ai-skill 只保留 generalized dogfood metrics。

## Run 摘要

- **任務**：9j2 Operations Sync — 模組 **04 user-feedback**（`pull_only`，Admin 只读镜像 + `admin_memo` + worker 队列）；**S0→S4 完整 Execute**（schema → ModuleSync → RemoteSync/worker/API → 收敛/UI → 外层 L1–L3）。
- **Transport**：Cursor orchestrator + **Task** Executor / Verifier（多 slice 瀑布式 spawn）；用户授权「先 commit 再执行」。
- **Repo**：`<PROJECT_ROOT>` 外层（plan + S4 tests）+ `<INNER_REPO>/manageCode` 内层（`feature/user-feedback-pull`，4 inner commits）。
- **触发**：规划阶段 stakeholder 多轮产品决策（D1–D14）后 Execute；**Verifier 被用户 Stop 中断** → orchestrator 恢复并完成 S0 验证 + S1–S4。

## 相对 2h 的改进（已应用教训）

| 2h 缺口 | 2i 表现 |
|---|---|
| combined defer L1–L3 | **S4 同轮交付** Gherkin + BDD + L3 Java + api-surface smoke；`test:user-feedback-pull` 8/8 |
| RBAC 漏 `admin_role_menus` | S3 migration 含三连；`test:gate-admin-operations-rbac` **2/2 pass** |
| V5 仅 list | api-surface test 枚举 list / sync-remote / update-memo + 负向路径 |
| spec-before-execute 弱 | **Execute 前** 锁定：只读、无删除、双系统 RBAC、mapping defer（`merchant_id` NULL）、`ninej2_sync_jobs` 独立表 |

## 失败 / 不如预期（按发现顺序）

| # | 现象 | 根因分类 | 应有证据层 |
|---|---|---|---|
| F1 | 用户 **Stop** 中断 Verifier（S0） | transport 脆弱性：subagent 无 checkpoint | orchestrator **resume 协议**：重 spawn Verifier 或 shell 补验；plan §执行纪录标 pending |
| F2 | S2 Verifier spawn **`resource_exhausted`** | 平台配额 / 并发限制 | orchestrator **fallback**：shell 跑 `brief.verification`；记入「verifier 降級」量测 |
| F3 | S0 inventory：**mapper 仍查 `member_feedbacks`**，非 `merchant_user_feedbacks` | plan 假设 Part A 与代码偏差 | **S0 inventory 硬门禁**（§2.2 P1–P10）→ Strategy A+收敛；非 executor 失误 |
| F4 | `ninej2_sync_failure` 与 pull 入队语义混淆 | platform 表职责未在 brief 写清 | **Specification**：Execute 前决议独立 `ninej2_sync_jobs`（`job_kind` + `module_key`） |
| F5 | S4 runtime：legacy delete 路径 **403** 非 404/410 | **stale admin JVM**（2h F2 同类） | V5 backfill 写 **restart 步骤**；区分 403 envelope vs 410 |
| F6 | Stop hook **多次** close-out repair（Bootstrap / Cognitive / Git Report） | orchestrator 长 session 漏 per-turn obligations | 与 delegation 正交；记为 **transport friction** observation |
| F7 | 单 session 连跑 S0–S4（5 Executor + 2 Verifier） | 协调成本高但 **scope 边界清晰**（每 slice Do NOT 列表） | slice 瀑布可行；大 combined 仍建议 S3a/S3b 拆分 |

## 仲裁纪要（orchestrator，2026-07-09）

| finding | 处置 | 理由 | 后续 |
|---|---|---|---|
| F1 Stop 中断 Verifier | **fix**（流程） | acceptance-violation 若未验就标 closed | resume：重跑 V5-M；S0 → `slice_compliant_closed` |
| F2 resource_exhausted | **fix**（fallback） | 不能无证据关 S2 | shell `mvn test` 15/15；标 shell-verifier |
| F3 mapper 表名偏差 | **defer** → S3a | 真实但 S0 范围外 | Executor 收敛 slice 切换 mapper |
| F4 sync_failure vs sync_jobs | **fix**（Specification） | 避免 worker 接错队列 | D13 + OQ-UF12 写入 plan 后 Execute |
| F5 stale JVM 403 | **defer** | runtime 待用户重启 admin | M2/M8 pending；写入 §执行纪录 |
| F6 close-out hook 摩擦 | **defer** | beyond delegation loop | observation → Cursor transport |
| F7 多 slice 单 session | **observation** | 成本↑但闭环完整 | 保留瀑布；不强制拆 session |

## 量測欄

| 指標 | 值 |
|---|---|
| 模組 slice 数 | **5**（S0 deliverable + S1 inner + S2 inner/runtime + S3a/S3b + S4 L1–L3） |
| Executor spawn 次数 | **5** |
| Verifier spawn 次数 | **2**（1× Stop 中断 + 1× S1 pass；S2 shell fallback） |
| acceptance-violation（runtime，修前） | **0**（inner/tests 全绿）；**1 待验**（stale JVM delete 路径） |
| verifier 降級次数 | **1** — S2 `resource_exhausted` → orchestrator shell 重跑 tests |
| stakeholder 纠偏次数（Execute 前） | **≥4** — 只读/memo、worker 队列、双系统 RBAC、mapping defer |
| orchestrator 写 manageCode（应经 Executor） | **0** |
| 外层 L1–L3 defer？ | **否**（相对 2h 改进） |
| 新 consumer 机械 gate 沿用 | RBAC 三连 gate、api-surface smoke（来自 2h，2i 验证有效） |
| 模型自然落位 | **是** — ERA：S0 schema 证据 ≠ S4 user-visible closed；inventory 决定收敛策略 |

## 契约回饋（写回 canonical / consumer overlay）

1. **`resume-after-stop`** — Task Verifier 被 Stop/abort 时，orchestrator 须有 **显式 resume 步骤**（重 spawn 或 shell 补跑 `brief.verification`），不得在 §执行纪录 留 `pending Verifier` 却标 slice closed。
2. **`verifier-fallback`** — `resource_exhausted` 等 spawn 失败时，orchestrator 可 **降级** 跑 verification 命令，但须记入量测栏「verifier 降級」且不得省略 V5-M 语义（migration + psql）。
3. **`s0-inventory-gate`** — `pull_only` 且 plan 写「Part A 已有」时，**S0 必须**跑 inventory（表/mapper/API/UI/RBAC）并选 Strategy A/B/C；避免 S2 写对表、S3 才发现 mapper 指错表。
4. **`platform-queue-typing`** — 多表队列（`sync_failure` vs `sync_jobs`）须在 brief **Execute 前** 写清 `job_kind` / 禁止复用语义，否则 Executor 易接错 worker。
5. **`mapping-defer-pattern`** — `merchant_id` NULL + 配置 `product_id` 列表可 **解除** mapping 表硬依赖；backfill 标 `M0′ defer` 即可 Execute — 适合 9j2 单站未确认多租户时。
6. **`2h-lessons-transfer`** — 同一 consumer overlay 上，**RBAC gate + api-surface + combined L1–L3** 在 2i 一次通过，支持 2h 契约回饋可迁移性。
7. **Q5 仍 doc-only** — 2i 为第 **3** 个 ExternalRepoC Execute run（03 common-url、04 user-feedback）；仍不足以单独 promote schema，但 **强化** Q7（backfill 映射 tier）与 resume/fallback 契约。

## 相对 §2h 的增量

| 主题 | 2h | **2i 新增** |
|---|---|---|
| 模块类型 | push_on_write common-url | **pull_only** + worker enqueue + `mirror_status` |
| 关闭策略 | 曾拟 inner-only 关 combined | **S4 同轮** 外层三角链 |
| 规划深度 | 事后修 04 产品面 | **Execute 前** D1–D14 + OQ 全关 |
| Transport | — | **Stop 中断** + orchestrator resume + verifier fallback |
| Platform | RBAC 漏项 | **`ninej2_sync_jobs`** 与 `sync_failure` 分表决议 |
| Inventory | — | **mapper 指错表** 由 S0 发现、S3a 收敛 |

## 四责任闭环实例（本 run）

```text
Specification — sub-plan 04 + S0–S4 brief + verification_backfill（stakeholder D1–D14）
  ↓
Production — Executor ×5（`<INNER_REPO>` + `<PROJECT_ROOT>` outer tests）
  ↓
Evidence — Verifier ×2 + shell fallback；tests 15+8+10
  ↓
Decision — orchestrator 写 §执行纪录 close_kind；mapping defer = defer 非 fix
  ↓
Specification — 2i 本文 + consumer preflight 链回
```
