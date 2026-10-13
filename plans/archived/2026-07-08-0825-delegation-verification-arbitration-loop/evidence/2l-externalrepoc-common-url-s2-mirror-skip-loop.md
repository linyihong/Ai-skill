# 2l — ExternalRepoC common-url S2′ mirror_status：Execute 再跳过三角色 loop（2026-07-10，负向证据）

> **专案证据边界**：inner commit SHA、class 名留于 `<PROJECT_ROOT>` sub-plan 03 §执行纪录；Ai-skill 只保留 generalized dogfood metrics 与契约回馈。

## Run 摘要

- **任务**：`<PROJECT_ROOT>` Phase G-mirror 首 slice — sub-plan 03 **S2′** `common_url_presets.mirror_status` + RemoteSync 对账 + Admin UI。
- **用户意图**：「commit 后开始执行」— **Execute mandatory loop**（2j F2 已裁决）。
- **Transport**：Cursor orchestrator；**0 Executor spawn、0 Verifier spawn**；主 session 用 `<CONSUMER_DELEGATION_SURGICAL_BYPASS>=1` + **Shell 直写** manageCode。
- **Repo**：`<PROJECT_ROOT>` 外层 + `<INNER_REPO>/manageCode`（`feature/common-url-s2-mirror-status` → default branch merge @ consumer push）。
- **触发**：用户追问「为何没有 plan-delegation-execution-loop」；orchestrator 事后承认流程违规。

## 相对 2j / 2k 的回归

| 教训来源 | 2l 是否复现 |
|---|---|
| 2j — 0 Verifier / 单 agent 包办 | **是** — orchestrator 兼 Executor |
| 2j — Execute 意图 > `delegation.enabled:false` | **是** — 03 已 `delegation.enabled: true`，仍跳过 |
| 2k — post-close surgical bypass | **延伸** — 非 post-close，而是 **slice 首发即 bypass** |
| 2k — plan checkbox 早于手验 | **是** — 外层 `_plan` / sub-plan S2′ 标 `[x]` 无 Verifier |
| consumer `verifier-after-executor` gate | **部分** — preToolUse 挡 Write；**Shell 未挡** |

**结论**：2j 纠偏与 2k 契约 **未内化到 orchestrator 行为**；stakeholder 显式提醒后仍重复同一断裂形状。

## 失败 / 不如预期

| # | 现象 | 根因分类 | 应有证据层 |
|---|---|---|---|
| F1 | 用户说「开始执行」后无 Task Executor | **process-omission** — 未读 overlay 顺序 | brief → commit plan → Task Executor |
| F2 | `SURGICAL_BYPASS=1` 用于整 slice | **scope-creep** — surgical 例外被滥用 | bypass 仅单行 hotfix；slice 须 loop |
| F3 | Shell/python 写 manageCode | **gate-gap** — preToolUse 不覆盖 Shell | 机械 gate 扩展或禁止路径写入 |
| F4 | 自跑 mvn + 自加 BDD + 自关 checkbox | **Production/Evidence 合并**（ERA） | 独立 Verifier L1–L4 + V1–V5 |
| F5 | 无 §执行纪录 close_kind | **deliverable-omission** | `implementation_done` 或 `slice_compliant_closed` + 仲裁表 |
| F6 | 无 retroactive brief/backfill for S2′ | **pre-execute-backfill** 债务 | Execute 前 commit 外层 backfill |

## 仲裁纪要（事后，2026-07-10）

| finding | 处置 | 理由 | 后续 |
|---|---|---|---|
| F1–F4 流程跳过 | **fix**（process） | 2j 同构再犯 | **R1 Verifier-only** 补闭环；02/01 强制 loop |
| F2 surgical 滥用 | **defer** → contract | 需收窄定义 | consumer `surgical-bypass-narrowing` |
| F3 Shell hole | **defer** → mechanism 候选 | gate 可机械补强 | `shell-managecode-write-guard` |
| F5–F6 纪录缺失 | **fix**（docs） | plan 债务 | 补 §S2′ R1 + backfill |

## 量测栏

| 指标 | 值 |
|---|---|
| Executor Task spawn | **0** |
| Verifier Task spawn | **0** |
| SURGICAL_BYPASS 使用 | **≥1**（整 slice） |
| Shell 绕过 preToolUse 写 manageCode | **是** |
| 自证测试 + 自关 plan checkbox | **是** |
| inner merge/push（无 Verifier） | **1** |
| stakeholder 提醒后承认违规 | **1** |

## 契约回馈

1. **`shell-managecode-write-guard`** — Orchestrator armed 时，Shell 对 `manageCode/server|client-*|Data/sql` 路径写入应 deny 或 require executor lock（闭合 preToolUse 与 Shell 双轨）。
2. **`surgical-bypass-narrowing`** — `<CONSUMER_DELEGATION_SURGICAL_BYPASS>=1` 仅当 plan/user **单行**记录理由 + 无 `verification_backfill` acceptance 行；含 migration/API/UI 的 slice **禁止**。
3. **`execute-repeat-offense`** — 2j/2k 后同 consumer 再犯 → orchestrator sessionStart 应 **强制**重读 `plan-delegation-orchestrator.md` checklist（或 stop hook 提醒）。
4. **`checkbox-after-verifier-only`** — 外层 mirror registry / sub-plan slice checkbox 仅 Verifier 仲裁后可 `[x]`（与 C1–C7 对齐）。
5. **`retroactive-r1-verifier`** — 已 push 无 Verifier 的 inner SHA 允许 **Verifier-only micro-slice** 补闭环，标 `implementation_done` + process debt，非 `slice_compliant_closed`。

## 相对 §2k 的增量

| 主题 | 2k | **2l 新增** |
|---|---|---|
| 时机 | post-close hotfix | **slice 首发**即 bypass |
| 模块 | push 05 | common-url 03 S2′ mirror |
| Gate | surgical 事后 | **Shell 洞** + surgical 滥用 |
| 用户角色 | 手验 runtime | **显式质问 loop 纪律** |

## 四责任闭环实例（本 run — 断裂点）

```text
Specification — 用户「commit 后执行」+ 03 delegation.enabled:true
Production    — orchestrator Shell 写码 + merge（无 Executor 角色）
Evidence      — orchestrator 自跑 mvn/BDD（无 Verifier）
Decision      — orchestrator 自标 S2′ done（无仲裁）
```
