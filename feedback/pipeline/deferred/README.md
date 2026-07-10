# Deferred Feedback Ledger（执行债索引层）

跨 session 可回收的 **NEEDED + DEFERRED|UNAVAILABLE** learning，避免静默流失。

设计来源：[`plans/active/2026-07-10-0955-final-form-feedback-execution-closure.md`](../../plans/active/2026-07-10-0955-final-form-feedback-execution-closure.md) Phase 1。

## 这是什么、不是什么

| 是 | 不是 |
|---|---|
| 已判定 worth capturing 的 **执行债** 索引 | plan authority |
| committed、去敏后的 durable pointer | 对话 transcript 替代品 |
| closure 前须 resurface 的 open 项 | enforcement（Phase 4 前无机械 block） |

## Entry schema

每条 entry 一个文件：`entries/DF-YYYYMMDD-NNN.md`

| 字段 | 必填 | 说明 |
|---|---|---|
| `id` | ✓ | `DF-YYYYMMDD-NNN` |
| `created` | ✓ | ISO date |
| `source_repo_context` | ✓ | `LOCAL` \| `NON_LOCAL` |
| `source_summary` | ✓ | 去敏一句话（无 inner SHA、无凭证） |
| `target` | ✓ | `feedback-history` \| `intelligence` \| `workflow` \| `enforcement` \| `project-docs` |
| `status` | ✓ | `open` \| `closed` \| `refuted` \| `expired` |
| `closure_evidence` | closed 时 | 路径或 commit 指针（指向 **authority 层**，非另一 entry） |
| `related_evidence` | 可选 | Ai-skill / consumer 广义 dogfood 链 |

## Invariants

1. **entry 不可指向 entry** — `closure_evidence` 必须落在 workflow/governance/plan/evidence 等 authority 路径。
2. **ledger 是索引层不是 authority** — 真正 closure 在 writeback 目标层（如 `evidence/2n-*.md`、consumer plan §7）。

## 去敏

遵守 [`enforcement/sanitization.md`](../../enforcement/sanitization.md)：不写 API key、DB 密码、live host 细节；consumer 专有名词用 `<PROJECT_ROOT>` / ExternalRepoC。

## 索引

| ID | 状态 | 摘要 | target |
|---|---|---|---|
| [DF-20260710-001](entries/DF-20260710-001.md) | **closed** | ExternalRepoC push DEL-S6 → delegation 2n writeback | workflow |
| [DF-20260710-002](entries/DF-20260710-002.md) | open | L4 live push receipt 需 gateway 凭证 | project-docs |
| [DF-20260710-003](entries/DF-20260710-003.md) | open | V5-A runtime smoke 需 feature 分支 admin 重启 | workflow |

## 量测栏（人工月结）

| 指标 | 2026-07-10 |
|---|---|
| entries created | 3 |
| closed | 1 |
| open | 2 |
| closure ratio | 33%（首笔，观察期） |
