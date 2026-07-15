# 2u — ExternalRepoC P12-R1 删除 merchant-product-mapping：`implementation_done` + V5-A defer（2026-07-15，正负对照）

> **专案证据边界**：inner commit、migration 文件名、merge SHA 留于 `<PROJECT_ROOT>` sub-plan `12-push-product-id-fixed-no-merchant-bind.md` §执行纪录；Ai-skill 只保留 generalized dogfood metrics 与契约回馈。

## Run 摘要

- **任务**：Stakeholder 定案删除 Admin「商户 Product 映射」功能面；push create 固定 `product_id=1`（plan 12 P12-R1，`slice_kind: combined`）。
- **结果**：**Executor + Verifier 独立 Task** 完成；`close_kind: implementation_done`（**非** `slice_compliant_closed`）— V5-A authenticated HTTP defer。
- **Transport**：Cursor orchestrator；brief commit → Executor `e8f51f9b` → Verifier `c9c1907a`；merge 前 rebase；post-merge surgical TS fix + push。
- **触发**：2j 教训后 Execute 意图 mandatory loop；产品 pivot（固定 product_id、不绑商户）。

## 相对 2j / 2n / 2m 的对照

| 指标 | 2j | 2n | 2m | **2u（本 run）** |
|---|---|---|---|---|
| Verifier spawn | 0 | 6/6 | 有 | **1**（独立 Task） |
| Executor spawn | 1（包办） | 6/6 | 有 | **1**（独立 Task） |
| `delegation.enabled` | false 误豁免 | true | true | **true** |
| close_kind | 无 | `slice_compliant_closed` | mixed | **`implementation_done`** |
| V5-M（migration apply） | — | defer 常见 | stale JVM 复发 | **linked** |
| V5-A（admin_test HTTP） | — | defer（D2/D3） | defer | **defer** — captcha/Redis |
| orchestrator manageCode 写码 | 有 | 0 | 0 | **0**（merge 后 1× surgical bypass） |
| post-merge release leg | — | 外层 only | 有 | **rebase + build gate fix** |

**结论**：2j 流程缺口在本 run **闭合**（E→V 分离、backfill、四栏 findings）；**2n 式全 linked 未达成** — 与 2n D2/D3、2m V5-A 同类 **runtime auth 环境债** 复发，但已显式 `defer` 而非假绿关闭。

## Slice 量测（P12-R1 单 slice）

| 层级 | 结果 | 证据 |
|---|---|---|
| L1 Gherkin | linked | `push-delivery.feature` — mapping 404/410 + fixed product_id |
| L2 BDD | linked | `push-delivery-smoke.test.mjs` 7/7 |
| L3 IT | linked | `AdminPushDeliveryScenarioTest`；runtime IT present |
| inner | linked | mvn push outbound/gateway 10 tests |
| V5-M | linked | `run-migrations.sh`；`ninej2_merchant_product_mapping` dropped |
| V5-A | **defer** | `adminLogin` fail；captcha 不在 peekable Redis；`AdminPushDeliveryRuntimeTest` 4 skipped |
| deliverable | linked | `9j2-merchant-mapping.md` REMOVED；migration + manifest |

## 失败 / defer（显式登记）

| # | 现象 | 根因分类 | 处置 |
|---|---|---|---|
| F1 | V5-A：`admin_test` login 失败 | **runtime env** — captcha plaintext 不在 `127.0.0.1:6380` db0–3；Admin JVM 可能非 feature 构建 | **defer** — restart Admin + `/tmp/redisvenv` + db.env |
| F2 | 未认证 POST mapping/list → 401 非 404 | **test precondition** — 无法在无 token 下验 removal | 须 V5-A 认证路径 |
| F3 | pre-push build gate TS6133 | **deliverable-omission** — `effectiveMerchantId` 未使用 | **fix** — post-merge surgical on master |
| F4 | ff-only merge 失败 | **process** — origin/master 超前 3 commit | rebase feature → retry merge script |
| F5 | `implementation_done` vs stakeholder「删功能」 | **close semantics** — 代码/DB 已删；authenticated API 未验 | 文档区分；手验或 V5-A follow-up |

## 仲裁纪要（orchestrator，2026-07-15）

| finding | 处置 | 理由 |
|---|---|---|
| V1–V4 + V5-M | **accept** | 机械 gate 全绿 |
| V5-A captcha/Redis | **defer** | 环境非 acceptance-violation；对照 2n D2 |
| F3 TS6133 | **fix** | 阻塞 release push |
| F4 rebase | **fix**（process） | merge 脚本纪律 |
| slice 关闭类型 | **`implementation_done`** | C6–C7：runtime tier 未全 linked |

## 量测栏

| 指标 | 值 |
|---|---|
| Executor Task | **1** |
| Verifier Task | **1** |
| 合规 loop（E+V） | **1/1 slice** |
| `slice_compliant_closed` | **0** |
| `implementation_done` | **1** |
| V5-M linked | **1** |
| V5-A defer | **1** |
| post-merge surgical bypass | **1**（TS6133） |
| merge + push master | **1**（`6687c10` + `f81da8b`） |
| orchestrator manageCode diff（loop 内） | **0** |

## 契约回馈（写回 canonical / consumer overlay）

1. **`implementation-done-runtime-defer`** — combined slice 在 V5-M linked、V5-A defer 时，关闭类型应为 **`implementation_done`** + plan backfill `deferred`，**禁止**标 `slice_compliant_closed`（强化 2n D2/D3 判读）。
2. **`v5-a-captcha-redis-matrix`** — V5-A 须 checklist：**Redis 隧道** + **python redis venv**（PEP 668）+ **Admin 进程加载 db.env** + **captcha key 可 peek**；缺任一项 → `runtime-omission` 非 code fail。
3. **`merge-rebase-before-ff`** — consumer `merge-feature-to-master.sh` 遇 origin 超前须 rebase；记入 release leg checklist（延续 2d′/2k F6）。
4. **`post-merge-build-gate`** — pre-push client-admin `vue-tsc` 可拦 merge push；Executor 自验应含 **frontend build** 或 Verifier V5-U build（2u F3）。
5. **`stakeholder-delete-vs-close`** — 「功能删除」用户语言 ≠ loop `slice_compliant_closed`；须口头/文档区分 **代码已交付** vs **authenticated runtime 待验**。
6. **Q7 正向** — 2j 后首次 push 域 **完整 E+V** 样本；证明 overlay + hook 可机械 enforce（对比 2j 负向）。

## 四责任闭环实例（本 run — 断裂点）

```text
Specification — plan 12 acceptance + backfill tier=runtime
Implementation — Executor 删除 mapping + migration（linked）
Verification — Verifier PARTIAL；V5-A runtime-omission
Arbitration — implementation_done；V5-A defer；merge + surgical fix
Release — master pushed；feature branch deleted
```

**断裂点**：Verification 层 V5-A 未 linked → 关闭类型降级为 `implementation_done`（符合 execution-loop C6–C7）。

## 相对 §2j 的增量

| 主题 | 2j | **2u 新增** |
|---|---|---|
| Loop | 0 Verifier | **完整 E→V** |
| 关闭 | 假完成风险 | **显式 implementation_done** |
| 产品 | — | **删除功能面** pivot（非仅弱化 mapping） |
| Release | — | rebase + build gate + surgical post-merge |
