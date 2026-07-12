# 2o — <PROJECT_ROOT> tab-scroll `/h5`：单 session vs 三角色 loop 对照（2026-07-10）

> **专案证据边界**：inner commit、release_id、L3 exit code 留于 `<PROJECT_ROOT>` `docs/plans/active/2026-06-16-state-trust-transition-pilot.md` §Delegation 2026-07-10b；本档只保留 generalized metrics 与契约回馈。

## Run 摘要

- **任务**：`<PROJECT_ROOT>` continuation slice — `/h5` tab-shell scroll keep-alive（player overlay return + tab-bar round-trip）。
- **对照设计**：同一 incident、同一 acceptance 集（TAB-005/005b/001/002*），先后以 **Mode A（单 session 直连）** 与 **Mode B（Orchestrator → Executor Task → Verifier Task）** 执行。
- **结果**：Mode A 部分关闭（TAB-005 only）；Mode B 全 linked @ release `6561e6f`；根因深度与修复面显著不同。

## Mode A — 单 session 直连（负向/部分证据）

| 指标 | 值 |
|---|---|
| 角色 | 1 agent session（规划+实现+自验+deploy） |
| `delegation.enabled` | false（隐式） |
| verification_backfill | 无（事后补 plan） |
| Verifier 独立 session | **0** |
| inner unit 依赖 | 高（`scroll-load-more.test.ts` 绿即倾向关 slice） |
| deploy smoke | pass（不含 `tab-scroll-keepalive`） |
| L3 @ test `/h5` release `46759c7` | TAB-005/005b **pass**；TAB-001/002* **fail** |
| 宣称关闭 | 倾向「deploy smoke 绿 = 可交付」→ **假绿**（authority gap） |

**Mode A 漏抓根因（Executor 后补全的 4 项）**：

1. `tab-shell-last-url.ts` 未传 `basePath` → `/h5` 下 entry URL / `?tab=` 永不 persist
2. 跨 tab 返回后 `updateWhenActive` 仍刷新 cached pane → `viewport-missing`
3. 每 tab 页各挂 TabBar → 隐藏 pane 拦截点击
4. `router.replace('/')` 在 infra 逃逸到 gateway `/` → 需 prefixed `/h5` display href

**对称盲点（symmetric blind spot）**：单 session 自跑 L3 时仍可能只盯「刚修的路径」（player return），未把 tab-bar round-trip 当同等 authority case。

## Mode B — 三角色 loop（正向证据）

| 指标 | 值 |
|---|---|
| Orchestrator | 主 session — plan/backfill/commit 外层 → spawn Task |
| Executor Task | 1 — 5 inner commits → deploy `6561e6f` |
| Verifier Task | 1 — fresh session，只跑 L3（readonly） |
| `delegation.enabled` | **true**（§Delegation 2026-07-10b brief） |
| verification_backfill | Execute **前** 填表（7 行 tier/owner） |
| L3 @ test `/h5` `6561e6f` | TAB-005/005b/001/002/002b/002c **pass** |
| orchestrator implementation diff | **0**（server_doc 由 Executor） |
| close_kind | `tab_bar_continuation_closed`（slice 级，非全 plan） |

**Verifier 增量发现（非 Executor 宣称）**：

| finding | classification |
|---|---|
| 并行跑两条 L3 命令 → TAB-005 首次 file-level fail | `observation` — 须顺序跑同一 host |
| deploy stamp API 404，release SHA 未 HTTP 复核 | `observation` — 以 L3 行为为准 |

## 对照表（核心）

| 维度 | Mode A 单 session | Mode B 三角色 loop |
|---|---|---|
| 验收量尺 | 实现者自选（inner + 部分 L3） | brief `acceptance` + backfill tier |
| 关闭权威 | 同一 session 自判 | Verifier 独立 L3 → Orchestrator 仲裁 |
| 根因个数（本 slice） | 2（player path + base path 局部） | **4**（含 TabBar hoist、URL prefix） |
| deploy smoke vs L3 | smoke 绿、journey 红 | 显式登记 smoke gap → defer |
| 假绿风险 | **高**（smoke 不含 TAB-001） | **低**（backfill 强制 linked 才关） |
| session 成本 | 1× | 3×（+2 Task） |
| 适合 | surgical 单路径 hotfix | user-visible combined slice |

## 与 Phase C authority model 的桥接（consumer 回馈）

同一 falsification family，不同 domain：

| Phase C（interaction） | 本 run（continuation） |
|---|---|
| `event injected ≈ interaction success` | `inner unit green ≈ scroll capability holds` |
| L1 semantic + capability close-out | L3 browser journey on running H5 |
| pointer vs keyboard authority | deploy smoke vs TAB-scroll integration authority |

**不替代 C.5**（无 Predictive Interception Trial）；**强化** State Trust pilot §forming authority model。

## 契约回馈（写回 canonical / consumer overlay）

1. **`deploy-smoke-not-l3-authority`** — deploy smoke pass 不得关闭 user-visible continuation slice；须 backfill 列 L3 script（consumer: add `tab-scroll-keepalive` to deploy.sh — defer）。
2. **`l3-sequential-same-host`** — 同一 `H5_TEST_BASE_URL` 勿并行跑多条 browser integration（Verifier flake 登记）。
3. **`single-session-partial-close`** — 单 session 在 combined slice 上易产出 **partial authority**（一条 journey 绿即关）；三角色 loop 强制全 acceptance linked。
4. **`evidence-responsibility-inner-tier`** — 延续 2g（server_doc placement）：inner = producer assist only；closure authority = outer L3（Q8 信号增强）。
5. **`positive-after-partial`** — 2o 与 2j/2n 不同：非「0 Verifier」负向，而是 **同 incident 前后模式对照**；建议 consumer overlay 保留「模式对照」dogfood 模板。

## 量测栏

| 指标 | Mode A | Mode B |
|---|---|---|
| Verifier Task spawn | 0 | **1** |
| Executor Task spawn | 0（主 session 包办） | **1** |
| acceptance linked（TAB-001–005b） | 2/6 | **6/6**（002b/c 含于 pattern） |
| inner-only 关 slice 尝试 | **是** | **否** |
| orchestrator server_doc diff | 有 | **0** |
| post-close bypass | 0 | 0 |

## 关联

- Consumer：`2026-06-16-state-trust-transition-pilot.md` §Dogfood + §Delegation 2026-07-10b
- Consumer overlay：`plan-delegation-execution-loop.md` §反模式
- Consumer feedback：`inner-unit-test-alone-cannot-close-user-visible-slice.md`
- 负向对照（同 family）：2j（0 Verifier）
- 正向对照（同 family）：2n（6/6 E+V）
- Trial declaration：`2026-07-07-trial-declaration-enforcement-pilot` — 本 slice **未触发** detector（非 Dialog/Overlay 路径信号）
