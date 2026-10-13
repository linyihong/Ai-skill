# 2j — ExternalRepoC push Execute：跳过三角色 loop（2026-07-10，證據 only）

> **專案證據邊界**：inner commit、class 名留于 `<PROJECT_ROOT>` sub-plan 05 与 `preflight-feedback-log.md`；Ai-skill 只保留 generalized dogfood metrics 与契约回饋。

## Run 摘要

- **任務**：9j2 Operations Sync — 模組 **05 push**（Admin 出站 create + worker pull + `mirror_status`）；用户要求 **commit plan 后 Execute S0–S4**。
- **Transport**：Cursor orchestrator + **单次** `Task` generalPurpose agent（自报 S0–S4 + inner mvn green）。
- **Repo**：`<PROJECT_ROOT>` 外层（plan 仍 `delegation.enabled: false`、无 `verification_backfill`）+ `<INNER_REPO>/manageCode`（`feature/push-ninej2-sync`）。
- **触发**：用户「开始执行」+ stakeholder 确认 OQ 已关；orchestrator **未**走 Executor → Verifier → 仲裁。

## 相对 2i 的退步（反模式对照）

| 2i 纪律（已验证） | 2j 表现 |
|---|---|
| Execute 前 brief + backfill + commit 外层 | **未做** — 05 仍 `delegation.enabled: false`、无 backfill |
| Task **Executor** + **Verifier** 分 spawn | **单 Task** 包办实现 + 自验 mvn |
| Verifier V5 runtime smoke | **无** Verifier session |
| combined / user-visible 外层 L1–L3 | **S5/S6 未做**；无 Gherkin/BDD |
| orchestrator 零 manageCode diff | 可能 bypass（或未启用 hook）；**无独立 Verifier 证据** |

## 失败 / 不如预期（按发现顺序）

| # | 现象 | 根因分类 | 应有证据层 |
|---|---|---|---|
| F1 | 用户说「开始执行」后 orchestrator **直接 spawn 单 agent 写码** | **process-omission** — 跳过 loop 步骤 2–5 | 先 commit brief/backfill → Executor → **Verifier** |
| F2 | 用 **`delegation.enabled: false`** 当不 loop 理由 | **Specification 误读** — Execute 意图 ≠ frontmatter 开关 | stakeholder 裁决：**Execute plan = 三角色 loop 默认** |
| F3 | 无 `verification_backfill` 仍宣称 S0–S4 done | **Evidence-first Execution 缺口**（Q7） | orchestrator 步骤 2 mandatory |
| F4 | 仅 inner JUnit ~15 pass 即回报完成 | **inner-only 关闭** + verifier 降級为自验 | C1b/C2/C4 block；须 V1–V5 + 外层链 |
| F5 | S5 UI + S6 外层验收未做却未标 `implementation_done` | **关闭状态未声明** | 应标 `implementation_done` + follow-up S5/S6 slice |
| F6 | 机械 gate 未阻止「单 agent 包办」 | gate 只挡 orchestrator **写** manageCode，未追踪 **Verifier spawn** | consumer 补强：`verifier_required` Task gate（2j 回饋） |

## 仲裁纪要（orchestrator 事后复盘，2026-07-10）

| finding | 处置 | 理由 | 后续 |
|---|---|---|---|
| F1 跳过 Verifier | **fix**（流程） | acceptance-violation — loop 未关闭 | 补 05 backfill + `delegation.enabled: true` → spawn **Verifier** on inner commit |
| F2 `enabled: false` 借口 | **reject**（误读） | Execute 意图触发 loop；frontmatter 须在 Execute 前改为 true | 更新 consumer overlay + sessionStart 提醒 |
| F3 无 backfill | **fix**（Specification） | Q7 证据：backfill 是 Execute 前硬步骤 | orchestrator 顺序步骤 2 不可跳过 |
| F4 inner-only | **fix** | user-visible push slice 默认 combined | Verifier + S5/S6 follow-up |
| F5 关闭状态混乱 | **fix** | 须书面 `implementation_done` | 05 §执行纪录 |
| F6 gate 缺口 | **fix**（mechanical） | 2c/2d gate 挡写不挡「不验」 | `check-plan-delegation-orchestrator.py` verifier tracking |

## 量測欄

| 指標 | 值 |
|---|---|
| 模組 slice 数（计划） | **7**（S0–S6） |
| Executor spawn 次数 | **1**（与 Verifier **合并** — 反模式） |
| Verifier spawn 次数 | **0** |
| acceptance-violation（流程） | **≥1** — 无独立 Verifier |
| verifier 降級次数 | **1** — executor 自报 mvn = 自证循环 |
| orchestrator 写 manageCode（应经 Executor） | **未知 / 可能 0**（单 Task 代劳） |
| 外层 L1–L3 | **未交付** |
| stakeholder 纠偏（Execute 后） | **1** — 「执行 plan 应走三角色 loop」 |

## 契约回饋（写回 canonical / consumer overlay）

1. **`execute-intent-overrides-delegation-flag`** — 用户 Execute 意图（「开始执行 / 执行 sub-plan / commit 后执行」）→ **mandatory 三角色 loop**，**不得**用 `delegation.enabled: false` 豁免；Execute 前 orchestrator 须改 frontmatter 为 `true` 并补 backfill。
2. **`no-single-agent-combined-slice`** — 禁止单个 Task 同时担任 Executor + Verifier；Task prompt 若含 implement + verify 全包 → 机械 gate **deny** 或拆成两次 spawn。
3. **`verifier-after-executor-gate`** — consumer 层：`executor` subagent 完成后、下一 Task spawn 若仍为 Executor/implement 且本轮 Verifier 未完成 → **deny**；须先 spawn Verifier（fresh）。
4. **`pre-execute-backfill-commit`** — 无 `verification_backfill` + 无 plan commit 锚点 → **不得** spawn 第一个 Executor（orchestrator 步骤 2–3）。
5. **`implementation_done-vs-closed`** — 仅 inner S0–S4 完成时须标 **`implementation_done`** + S5/S6 follow-up，**禁止**暗示 `slice_compliant_closed`。
6. **Q5 仍 doc-only** — 2j 为 **negative evidence**（loop 被跳过）；强化 Phase 3 门檻：role boundary 需 **Verifier spawn tracking** 类 mechanical gate，不只 orchestrator write block。
7. **Q7 正向（反面教材）** — 无 backfill 的 Execute = 「做完再想怎么验」；与 Evidence-first Execution 直接冲突。

## 相对 §2i 的增量

| 主题 | 2i | **2j 新增** |
|---|---|---|
| Loop 完整性 | 5 Executor + 2 Verifier | **0 Verifier** — 回归单 agent |
| frontmatter | 04 已 delegation + backfill | **05 `enabled: false`** 被误当豁免 |
| 机械 gate | 挡 orchestrator 写码 | **未挡** 跳过 Verifier |
| 关闭语义 | slice 分级 + resume | **未声明** implementation_done |
| 教训类型 | transport 脆弱性 | **orchestrator 流程纪律** |

## 四责任闭环实例（本 run — 断裂点）

```text
Specification — 05 sub-plan OQ resolved（缺 backfill / brief）
  ↓
Production — 单 Task agent（应仅 Executor leg）
  ↓
Evidence — ❌ 缺失（无 Verifier leg）
  ↓
Decision — ❌ 未仲裁
  ↓
Specification — 2j 本文 + consumer gate 补强
```

**断裂点**：Production 与 Evidence 合并 → 无法回答「预计与实现的落差」。
