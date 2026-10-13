# 2q — ExternalRepoC Integration 默认切流：Verifier 仅 inner 绿 vs 运行态/整合缺口（2026-07-13）

> **专案证据边界**：09 tip、mvn 过滤、密钥空值留于 `<PROJECT_ROOT>` `09-integration-default-cutover.md`；本档只保留 generalized metrics 与契约回馈。

## Run 摘要

- **触发**：Stakeholder 质问「切到 Integration 后功能还通吗？」→ 答：仅 mock/unit，无 live HMAC；再质问为何 Verifier 未跑整合/真实路径，是否与既有 V5 / L1–L3 设定不一致。
- **对照**：2p 曾记为「正向 6/6 E+V」；本 run 暴露 **证据类型缺口**：loop 合规（E+V 都 spawn）≠ **Runtime / Integration constraint** 已满足。
- **结果**：字面可标 `implementation_done` + live defer；**叙事上**易被读成「功能已验通」——与 bookmark 案同构的 **绿灯假象（missing constraint）**。

## 缺口矩阵

| 证据层 | 规则要求（既有） | 2p/INT-D 实际 | 判读 |
|---|---|---|---|
| inner JUnit / mock cutover | 可作 implementation 辅助 | **107**/0 | 有 |
| V5-M migration apply | transport seed 变更时 mandatory | D4 做过 | 部分有 |
| V5-A `sync-remote`（双闸=Integration） | 运营 sync / 用户可观察路径 | **未跑** | **缺** |
| 外层 L1 `docs/features` + L3 IT | combined / 用户可见 | 多片标 `implementation` 跳过 | **缺提醒** |
| Partner live HMAC | 可 defer（无密钥） | defer 正确 | 允许，但须 follow-up |

**根因（非「少 spawn Verifier」）**：Orchestrator 把 transport cutover 标成纯 `implementation` + Verifier 只重跑 `*BrowserBackend*`，**没有任何 acceptance 行强制「Integration 出站路径可观察」** —— Feasible Set 仍含「默认 integration 但从未打通过」。

## 量测栏

| 指标 | 值 |
|---|---|
| Verifier spawn / slice | 6/6（流程合规） |
| Verifier 跑外层 L3 / V5-A Integration | **≈0** |
| Partner secrets present | **no** |
| Stakeholder 纠偏 | **1**（要求写入流程 + dogfood） |
| 与 2j 同构？ | **否**（有 Verifier）；与 **bookmark / 绿灯假象** 同构 **是** |

## 契约回馈（写回）

1. **`transport-cutover-not-inner-only`** — `module_archetype: transport_migration` / Integration 默认切流：**禁止**仅用 inner mock 关闭「路径已通」叙事；至少一条 `tier=runtime`（V5-A 含 sync-remote 或等价）+ 无密钥时 **defer + 未关闭功能通** follow-up。
2. **`features-plus-integration-pair`** — 用户可见 sync/transport：`docs/features/**/*.feature`（L1）须与外层 **integration/L3**（`tests/backend-java` 或 `tests/integration`）互链；Verifier 机械核对存在性（consumer gate）。
3. **`loop-green ≠ path-proven`** — dogfood 正向（2p）须区分：**三角色合规** vs **Runtime Constraint 满足**；缺后者不得在 stakeholder 回复写「功能通」。
4. **不单独 promote schema** — 仍 doc + consumer gate；Q5 open。

## 关联

- 正向流程样本：2p  
- 绿灯假象：bookmark V5 / delegated-execution §8  
- Consumer：`<PROJECT_ROOT>` verifier-checklist + `gate.plan_transport_runtime_evidence`
