# 2m — ExternalRepoC Phase G-mirror 批量 retrofit（2026-07-10，正负对照证据）

> **专案证据边界**：inner SHA、class 名、BDD 路径留于 `<PROJECT_ROOT>` main `_plan` §Platform mirror + sub-plan 01/02/03 §执行纪录；Ai-skill 只保留 generalized dogfood metrics 与契约回馈。

## Run 摘要

- **任务**：`<PROJECT_ROOT>` **Phase G-mirror** — 对已交付 sync 模块补 `mirror_status` 远端稽核（03 S2′、02 S-mirror、01 S-mirror）；main plan §Platform mirror 登记总表 + 共用验收 V-m1–V-m5。
- **结果**：**5/6 模块 done**（04/05 原生带 mirror；01–03 retrofit）；Phase G **closed**；publish/crash 未开 plan。
- **Transport 对照**：
  - **03 S2′** — 负向，见 **2l**（0 Executor/Verifier、surgical bypass）
  - **02 S-mirror** — **正向**：Orchestrator → Executor Task → Verifier Task → fix leg → `implementation_done` @ `8b430c3`
  - **01 S-mirror** — **正向**：同型 loop @ `aa243eb`；Verifier `cb84e43a`

## 相对 2l 的纠偏（同 Phase 内）

| 指标 | 2l（03 S2′） | 2m 正向 slice（02/01） |
|---|---|---|
| Executor Task spawn | 0 | **≥1** |
| Verifier Task spawn | 0 | **≥1** |
| SURGICAL_BYPASS | 整 slice | **未使用** |
| close_kind | 无 / 事后 R1 | **`implementation_done`** + 仲裁表 |
| verification_backfill | 事后补 | **Execute 前** V-m1–V-m5 |

**结论**：同一 orchestrator、同一 Phase G 内，**纪律可执行**；2l 为 regression，2m 正向 run 证明 **retrofit brief + V-m 模板** 可复用（clone `AdminCommonUrlRemoteSyncService` / migration 范式）。

## 量测栏（Phase G 批次）

| 指标 | 值 |
|---|---|
| retrofit slice 数 | **3**（03/02/01） |
| 合规 loop（E+V） | **2/3** |
| 负向 bypass | **1/3**（→ 2l） |
| 共用验收模板 V-m1–V-m5 使用 | **3/3** |
| Verifier 抓 manifest 遗漏（F1 同型） | **≥2**（bookmark、common-url） |
| stale admin JVM → V5-A defer | **≥2**（02/01 arbitration **defer**） |
| 外层 BDD mirror gate | **≥2**（7/7 级） |
| main plan 登记总表更新 | **是** |

## 失败 / 不如预期（跨 slice 模式）

| # | 现象 | 根因分类 | 仲裁 / 处置 |
|---|---|---|---|
| F1 | migration manifest 遗漏（02/03） | **deliverable-omission** | **fix** — Verifier 抓；Executor 补 manifest |
| F2 | stale JVM → list API 无 `mirrorStatus` | **runtime-env** | **defer** — V5-A restart；非 acceptance 假绿 |
| F3 | 03 整 slice bypass（2l） | **process-omission** | **fix** — retroactive R1 Verifier |
| F4 | Phase checkbox「5/6 done」含 1 负向 slice | **phase-vs-slice granularity** | **observation** — 登记总表须 per-slice close_kind |

## 契约回馈（Platform mirror → delegation loop）

1. **`retrofit-v-m-template`** — main plan **V-m1–V-m5** 作为 mirror retrofit 的 **mandatory verification_backfill** 有效；第二、三模块 clone 第一模块 RemoteSync + migration 形状，Verifier L2 可对照 reference SHA（`83fe413`）— **promote 为 consumer overlay 默认**。
2. **`platform-mirror-registry-table`** — §Platform mirror **登记总表**（模块 × migration × RemoteSync × UI × slice）让 orchestrator **不丢 backlog**；建议写入 `9j2-sync-module-alignment.md` §9 或 consumer project overlay 模板。
3. **`stale-jvm-v5-a-checklist`** — **recurring defer**（push 05、bookmark/app-url S-mirror）；Verifier `verifier_only` 应含「restart admin + V5-A list 字段」或标 **blocking** 若 slice 声明 UI 列必达。
4. **`phase-close-mixed-compliance`** — Phase G **closed** 但含 2l 负向 → **phase 关闭 ≠ 全 slice `slice_compliant_closed`**；kit 应区分 **phase milestone** vs **slice close_kind**（对齐 sd-delegated-execution C1–C7）。
5. **`positive-pair-after-2l`** — 2m 与 2l **同 Phase 对照**可作为 dogfood kit **「纠偏后正向 run」** 范例：stakeholder 质问 loop 后，02/01 强制 Executor+Verifier。
6. **`manifest-f1-recurrence`** — manifest 遗漏在 03 S2′、02 S-mirror **重复** → Verifier L3 **`verifier_only`：migration 文件 ∈ manifest** 应进入 mirror retrofit brief 固定项。

## 四责任闭环实例（02 S-mirror — 正向）

```text
Specification — V-m1–V-m5 backfill + brief（clone S2′ @ 83fe413）
Production    — Executor Task f089d898 → inner 8b430c3
Evidence      — Verifier Task 56539e88（F1 manifest）；fix leg 24e6cbd0
Decision      — Orchestrator 仲裁表：F1 fix / F3 defer stale JVM
```

## 相对 §2k / §2l 的增量

| 主题 | 2k / 2l | **2m 新增** |
|---|---|---|
| 范围 | 单模块 push / common-url | **三模块 batch retrofit** + main 登记总表 |
|  polarity | 负向为主 | **正负对照**（2l + 02/01 合规） |
| 验收 | 单 slice | **V-m1–V-m5 共用模板** 跨模块复用 |
| 关闭 | slice / post-close | **Phase G milestone** vs per-slice close_kind |
