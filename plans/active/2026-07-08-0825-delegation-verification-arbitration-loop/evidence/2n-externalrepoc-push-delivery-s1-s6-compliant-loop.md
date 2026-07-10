# 2n — ExternalRepoC push delivery DEL-S1–S6 合规三角色 loop（2026-07-10，正向证据）

> **专案证据边界**：inner commit、class 名、consumer plan §执行纪录留于 `<PROJECT_ROOT>` sub-plan `07-push-delivery-broadcast.md`；Ai-skill 只保留 generalized dogfood metrics 与契约回馈。

## Run 摘要

- **任务**：`<PROJECT_ROOT>` sub-plan **07-push-delivery-broadcast** — DEL-S1（gateway sign）→ S6（outer acceptance closeout）；**6 slice** 全程 Orchestrator brief → Executor Task → Verifier Task（S6 含 verifier 证据回填 commit `450b371`）。
- **结果**：sub-plan `status: completed`；main plan Phase J-4 勾选；`slice_compliant_closed` on S5/S6；**无 post-close surgical bypass**（对比 2k）。
- **Transport**：Cursor orchestrator；每 slice 独立 feature branch；外层 docs commit 与 inner manageCode 分支分离（S6 closeout branch 无新 manageCode 代码）。
- **触发**：2j/2k/2l 负向链后 stakeholder 要求按 slice 顺序执行至 sub-plan 闭环。

## 相对 2j / 2k / 2l 的对照

| 指标 | 2j | 2k | 2l | **2n（本 run）** |
|---|---|---|---|---|
| Verifier Task spawn | 0 | ≥1（纠偏轮） | 0 | **6/6 slice** |
| Executor Task spawn | 1（包办） | 有 | 0 | **6/6 slice** |
| `delegation.enabled` | false 误豁免 | true（纠偏后） | bypass | **true** |
| verification_backfill | 无 | 部分 | 事后 | **Execute 前每 slice** |
| close_kind | 无 | `slice_compliant_closed` 早于手验 | 无 | **`slice_compliant_closed` S5/S6** |
| post-close surgical bypass | — | **≥3** | 滥用 | **0** |
| orchestrator manageCode diff | 有 | 有（hotfix） | 有 | **0**（orchestrator 仅外层 plan/docs） |

**结论**：同一 consumer overlay、同一 push 域，**纪律可稳定执行至 sub-plan completed**；2n 为 2j→2k 纠偏链之后的 **正向闭环样本**，补强 ADR 条件「至少一个真实 SD 任务完整 loop」。

## Slice 量测（DEL-S1–S6）

| Slice | close_kind | Executor | Verifier | 外层链 |
|---|---|---|---|---|
| S1 gateway sign | `implementation_done` | ✓ | ✓ | inner + api.md |
| S2 token ingestion | `implementation_done` | ✓ | ✓ | inner 10/10 |
| S3 admin gateway | `implementation_done` | ✓ | ✓ | L1–L3 + feature |
| S4 mapping carve-out | `implementation_done` | ✓ | ✓ | mapping IT |
| S5 merchant broadcast | `slice_compliant_closed` | ✓ | ✓ | DEL-2 BDD/L3 |
| S6 outer acceptance | `slice_compliant_closed` | ✓ | ✓ | runtime test + plan §5 |

## 失败 / defer（显式登记，非假绿）

| # | 现象 | 处置 | tier |
|---|---|---|---|
| D1 | L4 live 终端收到通知 | **defer** — 无 `PUSH_GATEWAY_*` 凭证 | L4 |
| D2 | V5-A `AdminPushDeliveryRuntimeTest` skip | **defer** — admin JVM 未起 / unreachable | runtime |
| D3 | S2/S3 历史 V5-A stale JVM | **defer** — 重启 feature 分支 admin 后手验 | runtime |

**判读**：defer 在 plan §7 / §13 显式登记；**不**影响 `slice_compliant_closed`（外层 L1–L3 机械 gate 全绿）。

## 仲裁纪要（Phase 3 佐证，非本 run 内逐条）

| finding 类 | 处置 | 理由 |
|---|---|---|
| L4 / V5-A defer | **defer** | 环境/凭证超出 slice scope；已写 plan + deferred ledger |
| 2n 是否关闭 Phase 3 | **defer** | Q5 schema promotion 仍 open；2n 增强信号不足单独 promote schema |
| Q6/Q7/Q8 | **partial** | 2e 已给跨域观察；2n 强化 sd 域 backfill ≥2 次使用 |

## 量测栏

| 指标 | 值 |
|---|---|
| slice 总数 | **6** |
| 合规 loop（E+V）/ slice | **6/6** |
| Verifier 独立重跑矩阵 | **≥4** gates（push-delivery / push-sync / feature-canonical / runtime） |
| orchestrator manageCode implementation diff | **0** |
| post-close surgical bypass | **0** |
| sub-plan terminal status | **`completed`** |
| acceptance-violation（关后暴露） | **0**（对比 2k F2/F3/F7） |
| consumer evidence commits | `<PROJECT_ROOT>` outer plan closeout（去敏，见 consumer §13） |

## 契约回馈（写回 canonical / consumer overlay）

1. **`multi-slice-subplan-closeout`** — 6-slice sub-plan 可逐 slice `implementation_done` → 最终 S6 `outer_acceptance` 一次关闭；每 slice brief 须 pre-backfill。
2. **`runtime-test-skip-tolerant`** — `assumeApiReachable` + skip 合法，但 **Verifier 须在 §13 记 V5-A defer**，禁止 silent pass。
3. **`positive-chain-after-negative`** — 2j→2k→**2n** 三连证明 consumer overlay **可纠偏可闭环**；负向证据不删除，与正向对照共读。
4. **`stale-jvm-v5-a-checklist`** — 延续 2k/2m；DEL-S6 runtime 面仍触发 → checklist 应进 consumer dev-run README（defer 至 project-docs）。

## 关联

- 负向链：2j / 2k / 2l
- 正向对照（同 Phase 内）：2m
- 跨域：2e（Q6/Q7/Q8 已观测）
- Consumer plan：`<PROJECT_ROOT>` `07-push-delivery-broadcast.md`
- Deferred ledger：`feedback/pipeline/deferred/entries/DF-20260710-001.md`（closed）
