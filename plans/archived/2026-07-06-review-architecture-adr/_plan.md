---
id: 2026-07-06-review-architecture-adr
plan_kind: main
status: completed
owner: linyihong
created: 2026-07-06
priority: P1
required_for_completion: false
parent_discussion: review-workflow-slice-v1-v4
last_revised: 2026-07-07
---

# Review Architecture — Cognitive Role Primitive Gate（ADR-013）

**Status**: `completed` — **ADR-013 closed**（2026-07-07）· **Archived** 2026-07-07

**Owner**: linyihong  
**Scope**: ADR-013 (**Completed**) · 本 plan **closed** · 後續演進 → [`ADR-014`](../../constitution/ADR-014-cognitive-stance-capability-context.md)（**不在本 plan Phase 3**）

---

## Phase Map（完整 — 本 plan 唯一 phase 編號）

> **重要：** 本 plan 的 phase 編號到 **Phase 2.4** 為止。對話或舊 backlog 曾出現的「Phase 3 / Phase 4」**不屬於 ADR-013 計畫** — 見下方 [§Phase 3 從哪來？](#phase-3-從哪來--已取消--不執行)。

| Phase | 名稱 | 目的 | 狀態 | 主要 commit |
|---|---|---|---|---|
| **0a** | ADR draft | 問題框架 + ADR-013 Proposed | ✅ | `84a4605` |
| **0b** | 泛化 evidence | Reject D1 證據 | ✅ | `84a4605` |
| **0c** | Stakeholder input | 確認 Reject D1 / D2 provisional | ✅ | `84a4605` |
| **0b.5** | Taxonomy validation | `fault_finding` 收斂 | ✅ | `4df3a0c` |
| **0d** | ADR accept | `stance` 命名；ADR-013 Accepted | ✅ | `4053008` |
| **1.1** | Runtime Contract | stance contract + registry + schema | ✅ | `030f105` |
| **1.2** | Runtime Enforcement | capability-invoke Warning | ✅ | `5482450` |
| **1.3** | Consumer | `cross-cutting/review/` | ✅ | `195159d` |
| **1.4** | Dogfood | 3 capability 共用 contract | ✅ | `195159d` |
| **2.1** | Runtime 一致性 | routing + refresh + graph | ✅ | `948f2f1` |
| **2.2** | Navigation 對齊 | Documentation Drift Lock | ✅ | `5c732fa` |
| **2.3** | Debt Payoff | Canonical Ownership Lock | ✅ | `fb9cdc1` |
| **2.4** | Contract Regression | 4-case invoke regression | ✅ | `948f2f1` |

**Architecture Evolution Pattern（本 plan 驗證完成）：**

```text
Architecture Decision → Runtime Contract → Navigation Alignment
  → Regression → Drift Lock → Debt Payoff → Close ADR
```

---

## Phase 3 從哪來？ — 已取消 / 不執行

| 來源 | 內容 | 本 plan 處置 |
|---|---|---|
| **舊 backlog（v2 plan 草稿）** | Phase 2 = 對齊 governance/README/routing；Phase 3 = typed review 文件 + validation scenarios；Phase 4 = Brower overlay | **重編為 Phase 2.1–2.4**（2026-07-07 stakeholder） |
| **Agent 對話摘要** | 「Phase 3 typed reviews」 | **從未寫入本 `_plan.md`**；agent 口頭提及 ≠ plan phase |
| **Stakeholder 2026-07-07** | ADR-013 close 後 **不開 Phase 3** | **Accepted** — 本 plan `status: completed` |

**舊「Phase 3」項目去向（Post-close backlog，非本 plan phase）：**

| 原項目 | 新歸屬 | 狀態 |
|---|---|---|
| Typed review docs（architecture/contract/release 獨立 consumer 面） | ADR-014 或獨立 plan | deferred |
| 更多 validation scenarios | ADR-014 / consumer 擴充時 | deferred |
| Brower project overlay sync | optional project overlay | deferred |
| Architecture Evolution Lifecycle 方法論 | enforcement 層文件 | deferred（驗證期後） |
| ADR-014 stance taxonomy acceptance | [`ADR-014`](../../constitution/ADR-014-cognitive-stance-capability-context.md) | **新 ADR，非 Phase 3** |
| Mechanical enforcement hard block（Warning → block） | ADR-014 / runtime phase | deferred |

---

## Executive summary

Software-delivery 存在 Governance / Execution / README **三層契約不一致**。多輪討論結論：問題不是加 link，而是 AI 缺乏 **Reviewer 視角** 的可發現執行模型。

**兩個問題必須分開：**

| 問題 | 狀態 |
|---|---|
| Review 是不是 Workflow Phase？ | **已證明：否**（Phase 0a） |
| Role 是不是 Runtime Primitive？ | **Reject D1（Accepted）** → **Accept D2（Phase 0d）** |

**Phase 0b 產出**：[`01-phase0b-perspective-generalization-evidence.md`](01-phase0b-perspective-generalization-evidence.md)

**Phase 0b.5 / 0d 產出**：[`02-phase0b5-perspective-taxonomy-validation.md`](02-phase0b5-perspective-taxonomy-validation.md)

**ADR-014（Proposed）**：[`ADR-014`](../../constitution/ADR-014-cognitive-stance-capability-context.md) — stance 合法值 taxonomy（**Post-close，非 Phase 3**）

**Outcome：** ADR-013 **Completed**；Review = capability invoke + `context.stance`；無 `sd-review` slice；無 `cognitive_role` primitive。

---

## 完整討論歷程（v1 → Phase 0b）

### v1 — Execution-centric 缺口

**現象**

- [`execution-flow.md`](../../workflow/software-delivery/execution-flow.md) Cognitive Slice 表無 review
- [`README.md`](../../workflow/software-delivery/README.md) step 5：「需要時參考 review-checklist」
- [`execution-flow.yaml`](../../workflow/software-delivery/execution-flow.yaml) 無 review loading surface
- Brower 專案 mirror slice order，零 review-checklist 引用

**診斷**：review = 外部事件（PR / release），非 delivery 內建。AI 路徑：Implement → Test → Finish。

**v1 提案**：在 execution-flow 加 link 或 Review Gate — 被判定為 **symptom fix**。

---

### v2 — Delivery-centric vs Execution-centric

**使用者論點**

- Review 應是 delivery 一部分，非 optional
- 理想流：Implementation → Self Review → Validation → Closure
- execution-flow 應 thin；review-checklist 不應塞進正文

**Plan 調整**

- ADR 先行
- `sd-review` slice + Review Gate（implementation 後）
- 三層契約對齊（governance / execution / routing）

**殘留問題**：仍 **Code Review 綁架** — 單一 post-impl gate 無法覆蓋 Architecture / Contract / Release review 時機。

---

### v3 — Fundamental abstraction + Option D（Cognitive Role）

**升級**

- ADR 問題：「Review 在 Cognitive System 中扮演什麼角色？」
- Evaluation Criteria 六維；Options A/B/C/D
- 四層：Workflow → Role → Capability → Artifact
- thin execution / fat slice 原則保留

**Option D**：Review 不是 slice，是 **Cognitive Role**；capability 是 role 的工具。

---

### v4 — Primitive Promotion Gate

**使用者 challenge**

- Role 不應轻易成 primitive，除非 **獨立語意 + 泛化**
- D 分叉：**D1**（Runtime Primitive）vs **D2**（Capability Context）
- ADR 核心：「Cognitive System 是否需要 cognitive_role primitive？」
- Hybrid (C) 退化風險 → reject
- Recommendation **不預設 D1**
- 五題必答 + 泛化測試
- Primitive 成本：routing / runtime / graphs / overlays

---

### Phase 0a — ADR draft（2026-07-06）

- [`constitution/ADR-013-cognitive-role-primitive-gate.md`](../../constitution/ADR-013-cognitive-role-primitive-gate.md)（Proposed）
- [`plans/archived/2026-07-06-review-architecture-adr/_plan.md`](_plan.md)
- [`architecture/ai-native-cognitive-execution-system.md`](../../architecture/ai-native-cognitive-execution-system.md) ADR 表新增 ADR-013

---

### Phase 0b — 泛化 / 反例 evidence（2026-07-06）

**使用者 Phase 0b 方法論**

- 不找更多 Review case；找 **非 Review** case + **反例**
- 成功標準：**證明 Role 邊界**，非只找支持 Role 的例子
- 主矩陣：Activity ×（Workflow? / Capability 即可? / Role 更自然?）
- 刻意反例：Refactoring、Documentation、Test Authoring
- 若僅 Review (+ 弱 Debugger) → **D2**；若 Planner+Architect+Reviewer+Debugger+Validator 同模式 → **D1**

**Phase 0b 結論**（詳見 evidence 檔）

- Review：**Perspective 候選**（非 Workflow Phase ✓）
- Debugger：**未證偽** — 與 Review 可能共享 `fault_finding` perspective（Phase 0b.5）
- Planning / Architecture / Validation：**slice 已擁有** → 不升 Role
- Refactoring / Documentation / Test authoring：**Capability 反例**
- **≥3 共享 primitive 語意：不成立** → **Reject D1**

### Phase 0c — Stakeholder input（2026-07-06）

| 決策 | 狀態 |
|---|---|
| Reject D1 | **Accepted** — 證據方向正確（泛化不足，非 Review 特殊） |
| D2 working model | **Accepted provisionally** |
| D2 final acceptance | **Deferred** — perspective taxonomy 未完成 |
| `perspective: reviewer` 定案 | **Deferred** — 可能改為 `fault_finding` |

**核心洞察：** Review / Debug / Security Audit / Incident Analysis 可能共享 **Negative Evidence Seeking** — 研究對象應是 **Perspective**，不是 **Role**。

---

## 三層契約不一致（證據表 — Phase 0 起點；Phase 2 已修正）

| 層 | 檔案 | Phase 0 現況 | Phase 2 修正 |
|---|---|---|---|
| Governance | [`software-delivery-governance.md`](../../governance/ai-runtime-governance/software-delivery-governance.md) | Artifact completeness 含 review report | review report → capability artifact（consumer 產出） |
| Execution | [`execution-flow.yaml`](../../workflow/software-delivery/execution-flow.yaml) | 無 review surface | `invoke_review_capability` + `review_invocation` loading surface |
| README | [`workflow/software-delivery/README.md`](../../workflow/software-delivery/README.md) | review-checklist = optional | §Review invoke；capability-centric |
| Routing | [`routing-registry.yaml`](../../knowledge/runtime/routing-registry.yaml) | 無 stance contract routes | `route.runtime.capability-context` + `route.governance.cognitive-stance` |
| Validation | [`validation.md`](../../workflow/software-delivery/validation.md) | review-report-template 错置於 Validate | ownership → `code-review` capability |

---

## 架構選項空間（完整）

| Option | Fundamental abstraction | Placement | Phase 0b 結論 |
|---|---|---|---|
| **A** | Workflow Phase | `sd-review` slice | **Reject** |
| **B** | Capability | `cross-cutting/review/` | Fallback（弱于 D2） |
| **C** | Hybrid | hook + capability | **Reject**（退化） |
| **D1** | Cognitive Role primitive | role → capability | **Reject**（泛化不足） |
| **D2** | Cognitive stance as capability context | capability + `context.stance` | **Accepted (Phase 0d)** |

### D2 invoke envelope（draft）

```yaml
invoke:
  capability: code-review
  context:
    stance: fault_finding    # ADR-013: fault_finding | default
    caller_slice: sd-implementation
    objective: optional      # refactor | plan — not stance
```

### 四層堆疊（D2 版）

```text
Workflow (caller slice)
  → invoke capability
  → context.stance (fault_finding | default)
  → artifact (review report, trace, …)
```

**不引入** `cognitive_role` runtime primitive。Stance 合法值擴充 → **ADR-014**。

---

## Evaluation Criteria（七維）

| Criteria | 說明 |
|---|---|
| Lifecycle Independence | 跨階段 review |
| Cognitive Separation | Reviewer vs Implementer |
| Discoverability | AI 知何時 invoke |
| Extensibility | 新 review 類型成本 |
| Contract Clarity | workflow thin |
| Runtime Cost | routing 維護 |
| Primitive Justification | 泛化 + 不可替代性 |

### Scoring（Phase 0b 更新）

| Criteria | A | B | C | D1 | D2 |
|---|---|---|---|---|---|
| Lifecycle Independence | W | S | W | S | S |
| Cognitive Separation | A | W | A | S | A |
| Discoverability | W | A | W | S | A |
| Extensibility | W | S | W | S | S |
| Contract Clarity | W | S | F | S | S |
| Runtime Cost | A | S | W | W | S |
| Primitive Justification | F | A | F | **F** | S |
| **Overall** | reject | fallback | reject | **reject** | **lead** |

---

## Review Type × Invocation Point（D2）

| Review type | Caller | D2 invoke |
|---|---|---|
| Contract | `sd-contracts` | `contract-review` + `stance: fault_finding` |
| Architecture | `architecture/` | `architecture-review` + `stance: fault_finding` |
| Security | contracts / impl | `security-review` + `stance: fault_finding` |
| Code | `sd-implementation` | `code-review` + `stance: fault_finding` |
| Performance | test-strategy / perf-risk-gate | `performance-review` + `stance: fault_finding` |
| UI compliance | `sd-ui-governance` | `ui-review` + `stance: fault_finding` |
| Release | `sd-validation` / `sd-closure` | `release-review` + `stance: fault_finding` |

---

## 五題必答（Phase 0b 更新）

| # | 問題 | 答案 |
|---|---|---|
| 1 | Review 跨 Workflow？ | **Yes** |
| 2 | Review 需 Persona 切換？ | **Yes** |
| 3 | Persona 需 Runtime Primitive？ | **No（Phase 0b 傾向）** |
| 4 | Capability 足以表達 Reviewer Context？ | **Yes** |
| 5 | Primitive 泛化到其他 Domain？ | **No** |

---

## Phase 計劃

### Phase 0a — ✅

- [x] ADR-013 Proposed
- [x] 本 plan
- [x] 決策樹 + Criteria 框架

### Phase 0b — ✅

- [x] [`01-phase0b-perspective-generalization-evidence.md`](01-phase0b-perspective-generalization-evidence.md)
- [x] 主矩陣 + 反例 + 五題更新
- [x] D2 推薦（待 0c 確認）

### Phase 0c — ✅

- [x] Stakeholder review Phase 0b evidence
- [x] **Reject D1** accepted
- [x] **D2 working model** accepted provisionally
- [x] **D2 final acceptance** — closed in Phase 0d（perspective → stance；ADR-013 Accepted）

### Phase 0b.5 — ✅

- [x] validation matrix；`fault_finding` 收斂
- [x] ADR-014 scope 決定

### Phase 0d — ✅

- [x] Stakeholder sign-off：Reject D1、Accept D2、Reject `reviewer` / `constructive_build`
- [x] 欄位命名：`perspective` → **`stance`**（epistemic stance，非 actor perspective）
- [x] ADR-013 Proposed → **Accepted**
- [x] ADR-014 Proposed（stance taxonomy 分離）
- [x] 保守 enum：`fault_finding | default`

### Phase 1 — ✅ complete

**Done definition satisfied** — Phase 1.1–1.4。

**順序：** 先 Contract，再 Consumer。Review 是第一個 consumer，不是 contract owner。

#### Phase 1.1 — Runtime Contract ✅

| 交付 | 路徑 |
|---|---|
| Stance 契約（contract owner） | [`governance/cognitive-stance.md`](../../governance/cognitive-stance.md) |
| Capability metadata schema | [`metadata/capability-context-schema.md`](../../metadata/capability-context-schema.md) |
| Capability registry（`requires_context.stance`） | [`knowledge/runtime/capability-registry.yaml`](../../knowledge/runtime/capability-registry.yaml) |

#### Phase 1.2 — Runtime Enforcement ✅

| 交付 | 路徑 / 指令 |
|---|---|
| Executable enforcement contract | [`runtime/capability-context.yaml`](../../runtime/capability-context.yaml) |
| Registry validate（blocking） | `ai-skill runtime validate` → check `capability_registry` |
| Invoke validate（warning, exit 0） | `ai-skill runtime capability-invoke --capability <id> [--stance <v>]` |
| Go implementation | `scripts/ai-skill-cli/internal/app/capability_context.go` |

| 行為 | Phase 1 決策 |
|---|---|
| 缺少 `stance` | **Warning**（不 auto-fill、不 hard block） |
| `stance` 與 `requires_context` 不符 | **Warning** |
| registry 結構無效 | **Error**（阻擋 runtime validate） |

#### Phase 1.3 — Consumer ✅

- [x] `workflow/cross-cutting/review/` — README, self-review, invocation-points, checklist
- [x] 遷移 `review-checklist.md` → `cross-cutting/review/checklist.md`（stub 保留舊路徑）
- [x] `sd-implementation` hook → self-review.md
- [x] execution-flow / README 導航更新

#### Phase 1.4 — Dogfood ✅

- [x] validation scenario [`capability-stance-fault-finding-v1.yaml`](../../validation/scenarios/runtime/capability-stance-fault-finding-v1.yaml)
- [x] 執行 scenario smoke + 記錄 evidence（2026-07-06）

#### Done Definition（Phase 1 完成）

1. Runtime 理解 `requires_context.stance`
2. `fault_finding` 是唯一標準化非 default stance
3. ≥3 capability 共用 contract，無 review 專屬 runtime 邏輯
4. `cross-cutting/review/` 僅 consumer，不重新定義 stance

**Governance invariant：** Capability 宣告 Context；Consumer 不得私自定義 Runtime Context。

**五維邊界：** Workflow (When) · Capability (What) · Cognitive Mode (How) · Stance (Reasoning) · Artifact (Output)

---

### Phase 2 — ✅ complete（Integration / Architectural Debt Payoff）

**順序（stakeholder 2026-07-07）：** 2.1 Runtime → 2.2 Navigation → 2.3 Debt Payoff → 2.4 Regression（2.4 可與 2.1 同步落地）

| Sub-phase | 真正目的 | 狀態 |
|---|---|---|
| 2.1 | Runtime Contract 成立 | ✅ |
| 2.2 | Navigation 與 Runtime 對齊 | ✅ |
| 2.3 | 清除舊架構殘留（Architectural Debt） | ✅ |
| 2.4 | Regression 保護 | ✅ |

#### Phase 2.1 — Runtime 一致性 ✅

**驗收：** `runtime validate` 完全理解 `requires_context.stance` 鏈；非「文件更新」。

| 交付 | 路徑 / 行為 |
|---|---|
| Routing registry routes | `route.runtime.capability-context`、`route.governance.cognitive-stance` |
| SD workflow wiring | `required_dependencies` + `loading_surfaces.review-invocation` |
| Refresh policy | `capability_registry` surface in `refresh-policy.yaml` |
| Graph edges | `knowledge/graphs/capability-context.yaml` |
| Governance invariant | Workflow 不得 branch on `stance`；只能 invoke Capability |
| manual_activation | 新 routes 已 annotate（runtime-trigger-wiring） |

**Commit：** `948f2f1`

#### Phase 2.2 — Navigation 對齊（Documentation Drift Lock）✅

**完成條件：** 所有導航層均**僅描述 Runtime Contract**，不得重新定義 Capability Context / Workflow / Consumer 邊界。

| 交付 | 內容 |
|---|---|
| Thin execution-flow | Implementation → invoke review capability → Validation |
| Fat README | §Review invoke — why / what / when / examples |
| Taxonomy §7.6 | `cross-cutting/review` as **consumer** — 不新增 `sd-review` |
| execution-flow.yaml | `invoke_review_capability` step + `review_invocation` loading surface |
| Lock | `review_architecture_doc_drift` — forbid review phase/workflow/slice in navigation canon |

**Canon 文件：** `execution-flow.md`、`README.md`、`cognitive-slice-taxonomy.md`、`routing-registry.yaml`

**Commit：** `5c732fa`

#### Phase 2.3 — Architectural Debt Payoff ✅

**完成定義（Debt Class，非檔案導向）：**

| Debt Class | 描述 | 狀態 |
|---|---|---|
| **A — Ownership Drift** | validation 不擁有 review report；stance 不於 README 重定義 | ✅ |
| **B — Path Drift** | 活躍 canonical source 指向 cross-cutting/review，非 stub path | ✅ |
| **C — Semantic Drift** | 無 review flow/phase/workflow；僅 capability invoke | ✅ |
| **D — Artifact Drift** | scenarios 與 templates 對齊 capability output | ✅ |

**Lock：** `canonical_ownership_drift`（runtime validate + pre-commit）

**主要修正：** `validation.md` ownership · 6 SD scenarios path · intake/perf-risk-gate · glossary/graphs

**Commit：** `fb9cdc1`

#### Phase 2.4 — Contract Regression ✅

**目的：** 保護 Runtime / Routing / Capability Metadata contract — 非 review 行為。

| Case | 預期 | 覆蓋 |
|---|---|---|
| 要求 `fault_finding`，invoke 有提供 | Pass | `capability_context_test.go` |
| 要求 `fault_finding`，invoke 未提供 | Warning, exit 0 | CLI + test |
| 不要求 stance | Pass, 無 warning | synthetic capability test |
| invoke stance 與 registry 不一致 | Warning, exit 0 | mismatch test |

| 交付 | 路徑 |
|---|---|
| Go tests | `scripts/ai-skill-cli/internal/app/capability_context_test.go` |
| Scenario | `validation/scenarios/runtime/capability-stance-contract-regression-v1.yaml` |
| Doc drift scenario | `capability-stance-doc-drift-v1.yaml` |
| Ownership scenario | `capability-stance-ownership-drift-v1.yaml` |
| Dogfood scenario | `capability-stance-fault-finding-v1.yaml` |

**Commit：** `948f2f1`（contract）+ `5c732fa` / `fb9cdc1`（drift locks）

#### Phase 2 close-out

- [x] ADR-013 Status → **Completed**（`constitution/ADR-013-...md`）
- [x] Plan `status: completed`
- [x] 三 lock 全通過：`capability_registry` · `review_architecture_doc_drift` · `canonical_ownership_drift`
- [x] **不開 Phase 3**（stakeholder 2026-07-07 — decision recorded at close）

**Post-close（deferred — 見 §Phase 3 從哪來）：** Architecture Evolution Lifecycle 方法論；ADR-014；Brower overlay

---

## 非目標

- 不新增 `sd-review` lifecycle slice
- 不預留未證據 stance enum（`creative` / `planning` / `optimization` …）
- 不把 `stance` 定義侷限在 `cross-cutting/review/README` 內

---

## 檔案索引

| 檔案 | 角色 |
|---|---|
| [`ADR-013`](../../constitution/ADR-013-cognitive-role-primitive-gate.md) | Canonical ADR — **Completed** |
| [`ADR-014`](../../constitution/ADR-014-cognitive-stance-capability-context.md) | Stance taxonomy — Proposed（**後續演進，非 Phase 3**） |
| [`01-phase0b-...`](01-phase0b-perspective-generalization-evidence.md) | Phase 0b 證據 |
| [`02-phase0b5-...`](02-phase0b5-perspective-taxonomy-validation.md) | Phase 0b.5 perspective taxonomy |
| [`governance/cognitive-stance.md`](../../governance/cognitive-stance.md) | Stance contract owner |
| [`cross-cutting/review/`](../../workflow/cross-cutting/review/README.md) | Consumer（Phase 1.3） |
| [`capability-context.yaml`](../../runtime/capability-context.yaml) | Executable enforcement |
| [`capability-registry.yaml`](../../knowledge/runtime/capability-registry.yaml) | Capability metadata |
| [`documentation_drift.go`](../../scripts/ai-skill-cli/internal/app/documentation_drift.go) | Doc drift lock（2.2） |
| [`canonical_ownership_drift.go`](../../scripts/ai-skill-cli/internal/app/canonical_ownership_drift.go) | Ownership lock（2.3） |
| [`.cursor/plans/review_workflow_slice_12ecc222.plan.md`](../../../.cursor/plans/review_workflow_slice_12ecc222.plan.md) | Cursor 鏡像（historical） |

### Validation scenarios（Phase 1.4 + 2.4）

| Scenario | Phase | 保護對象 |
|---|---|---|
| [`capability-stance-fault-finding-v1.yaml`](../../validation/scenarios/runtime/capability-stance-fault-finding-v1.yaml) | 1.4 | 3 capability dogfood |
| [`capability-stance-contract-regression-v1.yaml`](../../validation/scenarios/runtime/capability-stance-contract-regression-v1.yaml) | 2.4 | 4-case invoke contract |
| [`capability-stance-doc-drift-v1.yaml`](../../validation/scenarios/runtime/capability-stance-doc-drift-v1.yaml) | 2.2 | Navigation drift |
| [`capability-stance-ownership-drift-v1.yaml`](../../validation/scenarios/runtime/capability-stance-ownership-drift-v1.yaml) | 2.3 | Debt class A–D |

---

## Post-close backlog（非 Phase 3 — 獨立追蹤）

> 以下項目**不擴張 ADR-013**。開新 plan 或走 ADR-014 acceptance 流程。

| 項目 | 建議歸屬 | 優先 |
|---|---|---|
| ADR-014 Proposed → Accepted | ADR-014 | P1 |
| Stance enum 新值（需 evidence） | ADR-014 amendment | P2 |
| Invoke hard block（Warning → block） | ADR-014 / runtime | P2 |
| Typed review consumer surfaces（architecture/contract/release 獨立 hook） | 新 plan 或 ADR-014 | P3 |
| `software-delivery-governance.md` artifact completeness 用語對齊 capability | governance patch | P3 |
| Architecture Evolution Lifecycle 方法論 | `governance/` enforcement | P3（驗證期後） |
| Brower overlay `review_invocation` hooks | project overlay | optional |
| Archived plans 內舊 path 引用 | 不阻塞；archived 豁免 drift scan | low |

---

## Brower 專案備註（Post-close optional）

`<PROJECT_ROOT>/docs/browser-manage-development-workflow.yaml`（project overlay 範例）目前 mirror execution slices、零 review。ADR accept D2 後加 `review_invocation` hooks，非新 lifecycle step。
