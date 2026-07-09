# 2h — ExternalRepoC common-url Execute：验证不严与 runtime 漏网（2026-07-09，證據 only）

> **專案證據邊界**：inner commit、SQL patch 名、runtime URL 留于 `<PROJECT_ROOT>` sub-plan 03 §执行纪录与 `preflight-feedback-log.md`；Ai-skill 只保留 generalized dogfood metrics。

## Run 摘要

- **任務**：9j2 Operations Sync — 模組 **03 common-url**（原型 B，`push_on_write`）Execute S1–S4；事后规划 **04 user-feedback**（pull_only，Admin 只读 + memo）。
- **Transport**：Cursor orchestrator + Task Executor/Verifier（combined slice）；流程补丁由 orchestrator 写 `<PROJECT_ROOT>` overlay（RBAC gate、V5 api-surface smoke）。
- **Repo**：`<PROJECT_ROOT>` 外層 + `<INNER_REPO>/manageCode` 内层；common-url 已 merge inner `master`。
- **触发**：使用者手验 UI 报 `No static resource .../sync-remote`；追问「验 UI = 功能做完整」暴露验证链缺口。

## 失败 / 不如预期（按发现顺序）

| # | 现象 | 根因分类 | 应有证据层 |
|---|---|---|---|
| F1 | Admin 侧边栏无「常用 URL」 | migration 缺 **`admin_role_menus`**（仅有 permissions + menus） | L2 读码 + **`verifier_only` RBAC 三连 gate** |
| F2 | `POST .../common-urls/sync-remote` → static-resource 404 | **V5 只验 `list`**，未覆盖 UI 调用的全部 `@PostMapping`；或 **admin JVM 未 restart** | L3 **api-surface** runtime smoke + `verifier_only` frontend↔controller 路径 diff |
| F3 | combined slice 拟以 inner IT + `pnpm build` 关 slice，**defer L1–L3** | 误用「用户要手验 UI」作跳过 Gherkin/BDD 理由 | C1–C5：**外层三角链为 user-visible combined 关闭条件** |
| F4 | stakeholder 手验通过 `list` 后以为可交付 | **手验范围未绑定 brief acceptance 全表**；与 V5 缺口叠加 | backfill 每条 acceptance 映射 tier；手验 checklist ⊆ M 表 |
| F5 | Verifier 曾仅重跑 executor 自验或未跑 L3 负向 | verifier 降級为「复读自验」 | L1 **不充分** 契约再实例化 |
| F6 | 04 user-feedback 初稿含 delete/trash/商户面 | **Execute 前产品 scope 未锁定**；parity 沿用旧表结构假设 | Specification 演化：stakeholder 仲裁进 plan **先于** Execute |

## 仲裁纪要（orchestrator，2026-07-09）

| finding | 处置 | 理由 | 后续 |
|---|---|---|---|
| F1 RBAC 漏 role_menus | **fix** | acceptance-violation（菜单不可见） | patch migration + `test:gate-admin-operations-rbac` |
| F2 sync-remote 未验 / 未部署 | **fix** | acceptance-violation + process | api-surface smoke + AdminCommonUrlPresetsRuntimeTest；restart 写入 V5 |
| F3 defer L1–L3 | **reject**（作为关闭策略） | 违反 combined slice 契约 | overlay + verifier 反模式表更新 |
| F4 手验替代外层验收 | **reject** | 手验 = V5 补充，非替代 L1–L3 | preflight-feedback-log 条目 |
| F5 verifier 降級 | **fix** | 违反 L1–L3 契约 | checklist 补「全 UI API 面」 |
| F6 user-feedback 删除模型 | **fix**（Specification） | stakeholder 澄清：只读镜像 + memo，不删、不开商户 | sub-plan 04 修订 §1 |

## 量測欄

| 指標 | 值 |
|---|---|
| 模組 slice | **1**（03 common-url Execute；04 仅 planning） |
| acceptance-violation（runtime） | **≥2** — role_menus、sync-remote api-surface |
| verifier 降級次数 | **≥1** — 手验/inner-only 拟关 combined slice |
| stakeholder 纠偏次数 | **≥2** — 「验 UI 要完整」、user-feedback 产品面 |
| 新 consumer 机械 gate | **≥2** — RBAC 三连 gate、admin-operations api-surface smoke |
| orchestrator 写 manageCode（应经 Executor） | **0**（patch 后）；早期 workflow 补丁写 `<PROJECT_ROOT>` outer tests |
| 模型自然落位 | **是** — ERA：哪条证据能关哪个状态（inner ≠ user-visible closed） |

## 契约回饋（写回 canonical / consumer overlay）

1. **`V5-api-surface`** — user-visible 模块的 `runtime_smoke` 必须枚举 **client-admin `api/*.ts` 调用的全部路径**（含 `sync-remote`、trash 等），不得仅 `list`；建议 outer `api-surface.test.mjs` 静态 gate + Java runtime IT。
2. **`RBAC-triple`** — Admin 运营 migration 机械检查：`admin_permissions` + `admin_menus` + **`admin_role_menus`** 同批；bookmark 为对照样本。
3. **`combined-no-inner-close`** — `slice_kind: combined` 关闭须 L1–L3 linked；「stakeholder 手验」只作 V5 证据，不可替代 Gherkin/BDD。
4. **`restart-aware-runtime`** — V5 若依赖已运行 JVM，backfill 须写 **restart 步骤** 或探测 stale classpath（static-resource 类失败 = 部署证据缺失）。
5. **`spec-before-execute`** — pull_only / 只读镜像类模块：产品决策（删否、商户面、memo）须在 brief acceptance **Execute 前** 锁定，避免 Part A 既有 CRUD 误导规划。
6. **Q5 仍 doc-only** — 2h 强化 consumer gate 与 backfill 字段，不构成 schema promotion。

## 相对 §2d / §2d′ 的增量

| 主题 | 2d/2d′ | **2h 新增** |
|---|---|---|
| API 覆盖 | integration merge、live delete | **UI 全 API 面** runtime + 静态对齐 |
| RBAC | （未单列） | **role_menus 为侧边栏必要条件** |
| 关闭条件 | slice_compliant vs merge | **手验不能替代 L1–L3** |
| Specification | brief v2 范例 | **只读/无删除** 产品面须在 Execute 前写入 plan |
