# 2p — ExternalRepoC Integration 默认切流 INT-D0–D5 一口气多 slice 合规 loop（2026-07-13）

> **专案证据边界**：inner tip、migration 文件名、mvn 过滤命令留于 `<PROJECT_ROOT>` sub-plan `09-integration-default-cutover.md`；本档只保留 generalized dogfood metrics 与契约回馈。

## Run 摘要

- **任务**：`<PROJECT_ROOT>` sub-plan **09-integration-default-cutover** — INT-D0（盘点/例外冻结）→ D5（全量相关单测）；契约内可切模块接线 + DB/example 默认 Integration；common-url / push-create 永久例外。
- **结果**：sub-plan `status: completed`；**6/6** slice 各走 Orchestrator brief commit → Executor Task → Verifier Task；相关 Surefire **107**/0；live HMAC/sync-remote/真推送 **显式 defer**（无 Integration 密钥）。
- **Transport**：Cursor 主 session = Orchestrator；每 slice `Task` spawn Executor / Verifier（**不**指定 `model`）；同一内层 feature branch 连续推进（非每 slice 新 branch）。
- **触发**：Stakeholder 要求「全部默认切 Integration + example/DB + 全测」；相对既有「可切但默认 legacy」能力做默认切流。

## 相对 2n / 2j 对照

| 指标 | 2j（负） | 2n（正，6-slice） | **2p（本 run）** |
|---|---|---|---|
| Verifier Task / slice | 0 | 6/6 | **6/6** |
| Executor Task / slice | 1 包办 | 6/6 | **6/6** |
| `delegation.enabled` | false 误豁免 | true | **true** |
| verification_backfill | 无 | Execute 前每 slice | **Execute 前每 slice** |
| 同一 feature branch 连续 | — | 否（每 slice 多 branch） | **是**（tip 链 D0→D5） |
| orchestrator `server/**` 实作 diff | 有 | 0 | **0**（gate 挡住） |
| orchestrator 写 `Data/docs` | — | 少 | **有**（D0 运输矩阵 — gate 允许 docs，非 server） |
| 用户压力「一口气做完」 | — | 有序 slice | **有**；仍未跳 Verifier |
| sub-plan terminal | — | completed | **completed** |
| live / runtime | — | defer 登记 | **defer 登记**（无密钥） |

**结论**：在「单 session 连续多 slice + 用户催促一口气」压力下，**仍可维持 6/6 E+V**；补强 2n 的正向样本，并新增 **same-branch multi-slice** 与 **docs-only orchestrator 边界** 观察。

## Slice 量测（INT-D0–D5）

| Slice | close_kind | E | V | 备注 |
|---|---|---|---|---|
| D0 盘点+例外 | `implementation_done` | Orchestrator 自写 plan/docs | ✓ | docs-only；无 server |
| D1 Adapter CRUD | `implementation_done` | ✓ | ✓ | inner only |
| D2 Sync 双闸 | `implementation_done` | ✓ | ✓ | cutover 单测 |
| D3 Push 分流 | `implementation_done` | ✓ | ✓ | create→Gateway 永不 records/create |
| D4 seed/example | `implementation_done` | ✓ | ✓ | V5-M 有 DB 密码则 apply |
| D5 全测 | `implementation_done` | ✓（无新 commit） | ✓ | live defer |

## 失败 / defer（显式，非假绿）

| # | 现象 | 处置 |
|---|---|---|
| D-live-hmac | HMAC live / Admin sync-remote via Integration | **defer** — 无 Integration 密钥 |
| D-live-push | Push Gateway true send | **defer** — 无 Push Gateway 密钥 |
| D-live-webhook | Partner webhook 真推 | **defer** — 无 webhook secret / 对方未推 |
| O1 | 契约无 common-url Integration 路径 | **永久例外** legacy（写入矩阵，非 defer 假装可切） |
| O2 | 无 push records create | **永久例外** create→面 B |


## 仲裁纪要（本 run）

| finding 类 | 处置 | 理由 |
|---|---|---|
| Verifier 全 pass / 无 acceptance-violation | **accept close** | 每 slice 独立重跑 mvn 或 docs 核对 |
| Live runtime 缺密钥 | **defer** | 写 plan deferred 清单；不以假绿关 |
| Executor D3 顺带收紧 Admin/Merchant deliver 路径 | **observation** | 仍在 push create→Gateway 验收内；未开 fix 回派 |
| Orchestrator 写 Data/docs（D0） | **accept / 记边界** | 非 `server/**`；与「零 manageCode 实作」不完全同构——见契约回馈 |

## 量测栏

| 指标 | 值 |
|---|---|
| slice 总数 | **6** |
| 合规 loop（E+V）/ slice | **6/6**（D0 E=Orchestrator docs） |
| Task spawn 约计 | Executor×5 + Verifier×6 ≈ **11**（D0 无独立 Executor） |
| 外层 plan commit 约计 | **≥8**（每 slice brief 覆写 + close） |
| orchestrator `server/**` diff | **0** |
| post-close surgical bypass | **0** |
| acceptance-violation（关后暴露） | **0**（本 run 内） |
| 相关单测 | **107**/0（common+admin 过滤） |
| context compaction mid-run | **有**（对话摘要后续跑）；loop 仍闭环 |
| sub-plan terminal status | **`completed`** |

## 契约回馈（写回 canonical / consumer overlay）

1. **`same-branch-multi-slice-ok`** — 连续 tip 链（非每 slice 新 branch）在 Cursor Task transport 下可维持 E+V；与 2n「每 slice branch」可并存为两种合法形态；kit 应注明 **branch 策略由 consumer 约束，非 loop 不变量**。→ **观察保留**（未升格为 invariant）。
2. **`orchestrator-docs-vs-server`** — consumer gate 禁 `server|client|Data/sql` 但允许 `Data/docs` 时，D0 类「盘点冻结」可由 Orchestrator 直写 docs；dogfood 应区分 **implementation diff=0** vs **manageCode 任意路径=0**。→ **观察保留**。
3. **`一口气压力 ≠ 跳过 Verifier`** — **已写回** `plans/README.md` §Delegation、`delegated-execution.md` anti-pattern、kit Cursor 备注、consumer `plan-delegation-orchestrator.md`（2026-07-13）。
4. **`brief-churn-cost`** — **已写回** kit 使用流程 + consumer orchestrator「Brief 写法」：累积表优先于整份覆写（2026-07-13）。
5. **`verifier-report-shape-drift`** — **已写回** `plans/README.md` 四栏强制、kit 模板 B「缺表不合规」、consumer Verifier 清单（2026-07-13）。
6. **`permanent-exception-vs-defer`** — 契约缺口标永久例外优于假切再 defer → **观察**（acceptance 可实现性预检；未另开 schema）。
7. **不关闭 Phase 3** — 本 run 增强 sd 域多 slice 正向信号；**不**单独 promote schema（Q5 仍 open）。

## 关联

- 正向对照：2n（6-slice completed）
- 负向对照：2j（跳 Verifier）
- Consumer plan：`<PROJECT_ROOT>` `09-integration-default-cutover.md`
- 前序能力：同树 `08-integration-api-migration`（可切、默认 legacy）
