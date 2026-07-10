# 2k — ExternalRepoC push 纠偏后：slice 关闭与用户手验 runtime 缺口（2026-07-10，證據 only）

> **專案證據邊界**：inner commit、class 名、DB job id、merge SHA、branch 名、表名留于 `<PROJECT_ROOT>` sub-plan 05 §执行纪录与 `preflight-feedback-log.md`；Ai-skill 只保留 generalized dogfood metrics 与契约回饋。

## Run 摘要

- **任務**：接续 **2j 负向证据** — stakeholder 纠偏后完成 push **R1 Verifier + S5 UI + S6 外层验收**并标 `slice_compliant_closed`；随后用户手验与 orchestrator 运维暴露 **create 表单 / 远程同步 / pull 入库** 缺口。
- **Transport**：Cursor orchestrator；纠偏阶段有 Executor + Verifier loop；**关闭后**多次 `SURGICAL_BYPASS` 直改 `<INNER_REPO>/<APP_MODULE>` + orchestrator merge/push **无新 Verifier**。
- **Repo**：`<PROJECT_ROOT>` 外层 + `<INNER_REPO>/<APP_MODULE>`（`feature/<slice-slug>` → default branch @ consumer merge）。
- **触发**：2j 后用户反馈「创建推送选不到模版」「商户不存在」「远程同步任务进行中」；orchestrator API/DB 复现与补测。

## 相对 2j 的纠偏（正向对照）

| 2j 缺口 | 2k 纠偏后表现 |
|---|---|
| 0 Verifier | **R1 Verifier** + S6 outer L1–L3 + push-sync integration suite green |
| 无 backfill / `enabled: false` | brief + backfill + `delegation.enabled: true`（纠偏 session） |
| 无外层验收 | Gherkin + BDD + api-surface + Java runtime（list/sync-remote 面） |
| consumer 不挡「不验」 | `verifier-after-executor` gate 已落地（`<PROJECT_ROOT>` consumer） |

**结论**：2j 流程缺口在纠偏轮 **部分闭合**；**关闭声明早于用户可见路径验证** → 本 run 为 **post-close acceptance gap** 证据。

## 失败 / 不如预期（按发现顺序）

| # | 现象 | 根因分类 | 应有证据层 |
|---|---|---|---|
| F1 | 消息模版下拉为空 | **runtime stale JVM** — template-options API 新码未 reload；重启前 API 对 `{}` 返回 0 条 | V5：**restart-aware** + 测 API 契约（非仅 list） |
| F2 | 创建推送「商户不存在」 | **Spec/映射**：多数 preset App `merchantId=null`；create 只从 App 取商户 | V5-U：**UI 关键路径**（unassigned App + template）；acceptance 应含该组合 |
| F3 | 远程同步「任务进行中」 | **拓扑**：sync job queue 表 pending 无 Worker 消费；非「未部署 Admin」 | V5-W：**sync-remote 202 + worker 消费**；dev README 须写 worker |
| F4 | Worker 执行 pull 失败 | **映射**：remote row 缺 required URL 字段 → DB NOT NULL | V5：**pull upsert 样本行**（非 mock-only） |
| F5 | `slice_compliant_closed` 后多次 surgical bypass | **process**：关 slice 后的 hotfix 不走 Verifier | 记 **post-close surgical debt**；或开 follow-up micro-slice |
| F6 | orchestrator merge/push default branch | **release leg** 在 loop 外（延续 2d′） | 与 implementation 验收分离；merge 前可选 smoke |
| F7 | 事后补 create-flow integration gate test | **2j F4 延后修复** — 关 slice 时无 create 路径测试 | backfill 应含 **user-visible form API 链** |

## 仲裁纪要（orchestrator 事后，2026-07-10）

| finding | 处置 | 理由 | 后续 |
|---|---|---|---|
| F1 stale JVM | **fix**（ops） | observation — 本地 dev 纪律 | V5 清单写 restart 步骤 |
| F2 商户解析 | **fix**（code） | acceptance-violation — 手验路径未覆盖 | create-flow merchant resolver + integration test |
| F3 pending job | **fix**（ops） | 环境拓扑非单点 API | 启动 worker；可选 pending TTL/失败可重入队 |
| F4 required URL null | **fix**（code） | pull 映射缺口 | remote sync pull mapper 字段 fallback |
| F5 surgical 累积 | **defer** | 真实但属 hotfix 模式 | consumer 记次数；重大 hotfix 开 micro-slice + Verifier |
| F6 release leg | **defer** | beyond-loop 已共识 | 维持 2d′ 分离；merge 前 checklist |
| F7 补测 | **fix**（test） | 闭合 2j F4 债务 | 纳入 push-sync integration suite |

## 量測欄

| 指標 | 值 |
|---|---|
| 相对 2j Verifier spawn | **≥1**（纠偏轮） |
| 用户手验失败报告 | **≥3**（模版 / 商户 / 远程同步） |
| post-close surgical bypass 次 | **≥3**（模版、商户、测试） |
| 新增 integration test（事后） | **1** — create-flow gate |
| Worker 本地启动 | **1** — 消费 pending job |
| merge/push（orchestrator，无 Verifier） | **1** — `<INNER_REPO>` default branch |
| acceptance-violation（关后暴露） | **≥2** — create 路径、sync 拓扑 |

## 契约回饋（写回 canonical / consumer overlay）

1. **`v5-ui-critical-path`** — combined / user-visible slice 的 V5 须覆盖 **admin UI 关键表单所调 API 链**（从 UI module api surface 静态提取），不能仅 list/sync-remote 注册性 smoke。
2. **`sync-remote-requires-worker`** — sync-remote 返回 202 只证明入队；**验收须 Worker 消费**或 integration 测 job 状态迁移；dev 文档须 **Admin + Worker 双进程**。
3. **`pending-job-stale-guard`** — active-job dedup lookup 对永久 pending 阻塞重试；consumer 候选：**TTL 标 failed** 或 orchestrator 运维脚本（doc-only 先记）。
4. **`post-close-surgical-debt`** — `slice_compliant_closed` 后 bypass 修用户可见 bug = **技术债**；累计 N 次或触及 acceptance 应开 **follow-up micro-slice** + Verifier，禁止无限 bypass。
5. **`2j-correction-partial`** — 2j 纠偏可闭合 **流程与外层三角**；**不自动闭合**手验/runtime 拓扑 — 须 V5-W / V5-U 或明确 defer 到 ops checklist。
6. **Q7 正向** — backfill 应映射 **「用户会点的按钮」** → API → tier；2k 证明仅映射 sync/list 不够。

## 相对 §2j 的增量

| 主题 | 2j | **2k 新增** |
|---|---|---|
| Loop | 0 Verifier | 纠偏后有 Verifier；**关后 bypass** |
| 关闭语义 | inner-only / 未声明 | `slice_compliant_closed` **早于**手验 |
| V5 范围 | 无 / 片面 | list + sync-remote **≠** create 表单 + worker |
| 教训类型 | 跳过 Verifier | **关后 acceptance gap** + 拓扑依赖 |

## 四责任闭环实例（本 run — 断裂点）

```text
Specification — 05 标 slice_compliant_closed
  ↓
Production — surgical bypass hotfixes + merge
  ↓
Evidence — 用户手验 + API 复现（非 Verifier leg）
  ↓
Decision — 仲裁于本文 + 补测
  ↓
Specification — 2k 契约回饋 + consumer V5-W/U 候选
```

**断裂点**：关闭后 Production 继续、Evidence 靠用户反馈而非独立 Verifier — ERA「Producer ≠ Closure Authority」再次实例化。
