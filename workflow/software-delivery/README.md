# Software Delivery Workflow

`workflow/software-delivery/` 負責「App 開發審查與指引的執行流程」。本目錄保存 agent 在進行 app 開發審查時可照著執行的 planning flow、review flow、handoff flow、review checklists 與 contract-first 開發流程，讓開發與審查過程可重複、可驗證。

## 何時進入此 Workflow

當工作任務需要**開發、實作、修改程式碼、進行 code review / design review / release review** 時，agent 應自行判斷並載入本 workflow。不需要 runtime 觸發——agent 知道什麼時候需要開發。

進入方式：
1. 讀取 [`execution-flow.md`](execution-flow.md) thin index 了解執行流程與 focused loading surfaces
2. 依 task intent 載入需要的 execution surface：[`intake.md`](intake.md)、[`incident-observation.md`](incident-observation.md)、[`ui-incident-governance-workflow.md`](ui-incident-governance-workflow.md)、[`layer-ownership-matrix.md`](layer-ownership-matrix.md)、[`change-retrospective.md`](change-retrospective.md)、[`contracts.md`](contracts.md)、[`ui-contracts.md`](ui-contracts.md)、[`ui-governance.md`](ui-governance.md)、[`test-strategy.md`](test-strategy.md)、[`validation.md`](validation.md)、[`closure.md`](closure.md)、[`surgical-changes.md`](surgical-changes.md)
3. 依流程的 Start From Evidence / Change Intake 開始；新需求、重構、parity、缺失資訊或既有專案回填時載入 [`intake.md`](intake.md)
4. 需要 artifact 規範時參考 [`artifact-gates.md`](artifact-gates.md)
5. Post-implementation review：invoke capability（見下方 §Review invoke）— 不是 lifecycle slice
6. 需要完整開發流程 overview 或 embedded / producer-consumer fallback 時參考 [`development-process.md`](development-process.md)
7. 需要前端、行動、CLI、SDK 或其他 consumer surface 的 Screen Mapping、Consumer / UI Behavior / Screen / ViewModel Contract 或 Screen Traceability 時參考 [`ui-contracts.md`](ui-contracts.md)
8. 需要 UI compliance、design-system enforcement、accessibility evidence、behavior pattern checks、visual baseline review 或 AI visual review scoping 時參考 [`ui-governance.md`](ui-governance.md)；responsive taxonomy / authority / severity 本體引用 [`responsive-ui`](../../intelligence/engineering/governance/responsive-ui/README.md)
9. 需要 pre-build interrogation / product impact alignment / requirements cognition / BDD-lite / acceptance criteria / ambiguity resolution 時參考 [`requirements/`](requirements/README.md)
10. 需要 architecture fit analysis、DDD / CQRS / event sourcing / microservices decision 時參考 [`architecture/`](architecture/README.md)
11. 需要 Simplicity First / Surgical Changes / Think Before Coding 的行為範例時參考 [`examples/EXAMPLES.md`](examples/EXAMPLES.md)；examples 預設 suppress，僅在明確要求範例或 ambiguity 時載入
12. PR 觸動效能敏感路徑或含 AI 生成程式碼時，執行 [`perf-risk-gate.md`](perf-risk-gate.md) 的 5 步檢查（靜態 anti-pattern scan、L1 smoke 或 hot-path micro-benchmark、reviewer perf checklist、pre-deploy observability gate、canary rollout）；L0–L2 delivery model 見 [`perf-governance.md`](perf-governance.md)
13. 當成功訊號可能不同於真實系統狀態時，載入 validation reasoning：[`state-visibility-gap.md`](../../intelligence/engineering/execution/validation-reasoning/state-visibility-gap.md)、[`evidence-model.md`](../../intelligence/engineering/execution/validation-reasoning/evidence-model.md)、[`evidence-chain-validation.md`](../../intelligence/engineering/execution/validation-reasoning/evidence-chain-validation.md)、[`evidence-depth.md`](../../intelligence/engineering/execution/validation-reasoning/evidence-depth.md)
14. UI / consumer incident（Navigation / Continuation / Recovery 或 modification layer 未決）時：先 [`incident-observation.md`](incident-observation.md) → [`ui-incident-governance-workflow.md`](ui-incident-governance-workflow.md) + [`layer-ownership-matrix.md`](layer-ownership-matrix.md) → governance gate [`software-delivery-governance.md`](../../governance/ai-runtime-governance/software-delivery-governance.md) §Incident *；Ship 後 [`change-retrospective.md`](change-retrospective.md)

## Review invoke（ADR-013 — 導航層說明）

**Review 不是 Workflow Phase，也不是 lifecycle slice。** 它是 **capability invoke** — workflow 只決定 *when* 與 *which capability*；stance 需求由 capability registry 宣告；runtime 驗證 invoke envelope。

### 五層責任（Runtime Contract 為 canonical）

```text
Runtime Contract (governance/cognitive-stance.md)
  ↑
Capability Metadata (capability-registry.yaml — requires_context.stance)
  ↑
Execution Flow (thin — invoke boundary only)
  ↑
README (why / what / when / examples — 本節)
  ↑
Taxonomy (cross-cutting consumer — 非 sd-review)
```

### Review 是什麼？

| 問題 | 答案 |
|---|---|
| Review 是什麼？ | **Invoke a capability**（例如 `code-review`） |
| Stance 從哪來？ | `requires_context.stance` in [`capability-registry.yaml`](../../knowledge/runtime/capability-registry.yaml) |
| 標準非 default stance | `fault_finding`（ADR-014） |
| Workflow 能做什麼？ | 決定 **when** invoke、**which** capability id |
| Workflow 不能做什麼？ | `if stance == fault_finding` 分支；私有定義 stance contract |

### Why invoke?

- **Cognitive separation**：實作後切換到 fault-finding 程序，但不引入 `cognitive_role` primitive
- **Reuse**：`code-review`、`security-audit`、`incident-analysis` 共用同一 invoke contract
- **Discoverability**：caller slice 在 Implementation 後 invoke，Validation 前

### What capability?

| Review 類型 | Capability id | Typical caller | Stance（registry） |
|---|---|---|---|
| Code | `code-review` | `sd-implementation` | `fault_finding` |
| Security | `security-audit` | `sd-contracts`, `sd-implementation` | `fault_finding` |
| Incident | `incident-analysis` | `sd-incident-observation` | `fault_finding` |
| Architecture | `architecture-review` | `architecture/` | `fault_finding`（draft） |
| Contract | `contract-review` | `sd-contracts` | `fault_finding`（draft） |
| Release | `release-review` | `sd-validation`, `sd-closure` | `fault_finding`（draft） |

完整對照：[`cross-cutting/review/invocation-points.md`](../../cross-cutting/review/invocation-points.md)

### When should invoke?

| 時機 | Invoke |
|---|---|
| Post-implementation（Validation 前） | `code-review` |
| Contract / security 敏感變更 | `security-audit` 或 `contract-review` |
| 未知 incident observable | `incident-analysis` |
| Pre-release / merge readiness | `release-review` |

Consumer 入口：[`workflow/cross-cutting/review/`](../../cross-cutting/review/README.md)（checklist、self-review hook — **consumer only**，不擁有 stance contract）

### Example invoke envelope

```yaml
invoke:
  capability: code-review
  context:
    stance: fault_finding
    caller_slice: sd-implementation
```

Validate（Phase 1.2 — warning only, exit 0）：

```bash
ai-skill runtime capability-invoke --capability code-review --stance fault_finding
```

### 舊模型（已否決 — 勿 drift）

- ~~`sd-review` lifecycle slice~~ → **Reject**（ADR-013 Option A）
- ~~Review as Workflow Phase~~ → workflow 只 invoke capability
- ~~Consumer 私有定義 stance~~ → capability registry owns contract

Contract owner：[`governance/cognitive-stance.md`](../../governance/cognitive-stance.md) · ADR：[`ADR-013`](../../constitution/ADR-013-cognitive-role-primitive-gate.md)

## Scope

本 workflow 涵蓋以下流程與審查類型：

### 開發流程

- **Requirements Stage**：Pre-build interrogation + product impact alignment + BDD-lite / requirements cognition，包含需求拷問、framework source-of-truth discovery、duplication risk check、Impact Map × Customer Journey Map、behavior-driven discovery、acceptance definition、ambiguity resolution、traceability 與 validation target。
- **Architecture Stage**：domain architecture cognition，包含 DDD fit、bounded context discovery、consistency boundary design、architecture escalation。
- **Contract-First Development Process**：從企劃書到實作的完整開發流程，包含 Default Flow、Required Contracts、Product Brief Validation Gate、Change Intake Gate、Contract Governance Gate、Traceability Gate、BDD Execution Closure、Test Strategy Gate、Embedded/Hardware Flow、Missing Information Gate、Existing Project Documentation Backfill 等。
- **UI / Consumer Contract Process**：在 provider/consumer 平行實作前建立 Screen Mapping、Consumer Contract、UI Behavior Contract、Screen Contract、Frontend ViewModel Contract、Accessibility Contract 與 Screen Traceability；多入口 screen 需盤點 entrypoints、navigation contract 與 back-stack owner，避免 AI agent 只依 API shape 或 route marker 生成語意脫節的前端。
- **UI Governance Process**：當 UI contract 需要 compliance evidence 時，分類 governance domain、render context、collection method、validation mechanism、evidence class、severity 與 project-local design-system policy；browser review 是 evidence acquisition，visual regression / AI review 是 validation mechanisms，不是 governance domains。Responsive governance 的 contract、taxonomy、authority mapping 與 severity policy 由 [`responsive-ui`](../../intelligence/engineering/governance/responsive-ui/README.md) 提供，workflow 只負責在正確階段載入與收口。
- **UI Incident Process（Observe → Classify → Select Layer → Retrospective）**：未知 UI incident 時，**證據驅動變更** 決策鏈（非線性需求→合約→實作）：[`incident-observation.md`](incident-observation.md)、[`ui-incident-governance-workflow.md`](ui-incident-governance-workflow.md)、[`layer-ownership-matrix.md`](layer-ownership-matrix.md)（domain → owner → allowed modifications）、Ship 後 [`change-retrospective.md`](change-retrospective.md)；gate 在 [`software-delivery-governance.md`](../../governance/ai-runtime-governance/software-delivery-governance.md)。
- **Evidence-Oriented Validation**：當 API/adapter/UI 成功訊號不足以證明 persisted、external、identity-specific 或 user-observable state 時，依 engineering validation reasoning 建立 evidence chain、選擇 evidence depth，必要時要求 live system proof 與 independent observation。Critical journey validation 使用 BDD-owned Journey Specification，並在 validation layer 驗證 side-effect chain、expected outcomes 與 observable evidence。
- **Refactor / Replacement Parity**：當新入口、平台遷移、工具改寫或架構重組要取代舊能力時，先建立新舊能力 parity inventory，逐項列出舊入口、現有能力、副作用、外部依賴、新入口、parity 狀態與測試證據。

### 審查類型 → Capability invoke（非 lifecycle slice）

審查不是 workflow phase。下表映射 **review 類型 → capability id**；invoke 細節見 §Review invoke。

| 類型 | Capability | When invoke |
|---|---|---|
| Design / Contract | `contract-review` | 實作前 contract / API 穩定後 |
| Architecture | `architecture-review` | architecture fit 決策後 |
| Code | `code-review` | post-implementation、validation 前 |
| Security | `security-audit` | contract 或 implementation 安全敏感變更 |
| Release | `release-review` | validation / closure 前 |
| Embedded / Firmware | `code-review` + domain checklist | 同 code review invoke + checklist bodies |

Checklist bodies：[`cross-cutting/review/checklist.md`](../../cross-cutting/review/checklist.md) · 舊路徑 [`review-checklist.md`](review-checklist.md) 為 stub

## 核心原則

1. **Review 是預防不是懲罰**。目的是在問題進入 production 前發現，不是追究責任。
2. **Checklist 是輔助不是取代**。Checklist 確保基本項目不被遺漏，但 reviewer 仍需使用工程判斷。
3. **Review 結果必須 actionable**。每個 finding 應包含：問題描述、風險等級、建議修復方式。
4. **Review 記錄應可追溯**。每個 review 的 finding、decision 與 resolution 應可追溯到對應的 commit 或 ticket。
5. **Simplicity First（簡潔優先）**：從最簡單的實作開始。不要預先加入抽象層、Strategy pattern、或 speculative features。當需求證明需要複雜度時再重構。參見 [`examples/EXAMPLES.md`](examples/EXAMPLES.md) §2。
6. **Surgical Changes（外科手術式修改）**：只改解決問題所需的行。匹配既有 code style，不要順便 refactor 不相關的 code。參見 [`surgical-changes.md`](surgical-changes.md)（`sd-surgical-caveats` slice）和 [`examples/EXAMPLES.md`](examples/EXAMPLES.md) §3。
7. **Parity Before Replacement（替換前先對照）**：重構、遷移或 replacement 若會替代既有行為、入口、腳本、API、資料流程或操作能力，先盤點舊能力到新能力的對照與驗證證據，再開始實作。

## 與既有層的關係

- `workflow/software-delivery/` 是 App 開發指引執行流程的主要入口。所有 agent 應優先參考本目錄的內容。
- `analysis/development-guidance/` 提供安全控制、實作模式、平台指引、語言陷阱的 catalog 參考，被本 workflow 引用。
- `analysis/repo/` 可被本 workflow 引用來分析 repository 結構。
- `intelligence/` 可被本 workflow 引用來輔助工程判斷。
- `workflow/software-delivery/requirements/` 提供 requirements cognition 的執行流程；其 source intelligence 來自 `intelligence/engineering/requirements/`。
- `workflow/software-delivery/architecture/` 提供 architecture fit 的執行流程；其 source intelligence 來自 `intelligence/engineering/architecture/architectural-fit/` 與 `intelligence/engineering/architecture/domain-modeling/`。
- `governance/ai-runtime-governance/software-delivery-governance.md` 定義 requirements / behavior / contract delivery gates。
- `governance/ai-runtime-governance/software-delivery-architecture-governance.md` 定義 software-delivery architecture governance gate，但不把 DDD promotion 成 runtime invariant。
- `feedback/history/development-guidance/` 儲存開發指引的具體課程記錄。
- `skills/app-development-guidance/` 是原始 skill 目錄，已刪除。所有內容已遷移至本層。

## 遷移狀態

| 來源 | 目標 | 狀態 |
|------|------|------|
| `skills/app-development-guidance/WORKFLOW.md` | [`execution-flow.md`](execution-flow.md)、[`artifact-gates.md`](artifact-gates.md)、[`analysis/development-guidance/risk-translation.md`](../../analysis/development-guidance/risk-translation.md) | ✅ 已遷移，舊目錄已刪除 |
| `skills/app-development-guidance/process/` | [`development-process.md`](development-process.md) | ✅ 已遷移，舊目錄已刪除 |
| `skills/app-development-guidance/checklists/` | [`review-checklist.md`](review-checklist.md) | ✅ 已遷移，舊目錄已刪除 |

## 已提取內容

| 檔案 | 來源 | 說明 |
|------|------|------|
| [`execution-flow.md`](execution-flow.md) | `WORKFLOW.md` §1, §5-8（已刪除） | Start From Evidence、Change Intake、BDD Closure Loop、SDK Defect Closure、Same-Session Closure、Performance Gate、Backfill Rules、Validate |
| [`execution-flow.yaml`](execution-flow.yaml) | `execution-flow.md` | Software delivery execution executable contract：change intake、requirements、BDD closure、parity、performance、validation gates |
| [`intake.md`](intake.md) | `execution-flow.md` §1/§6 + `development-process.md` intake gates | Focused execution surface：需求接收、Change Intake、Pre-build Interrogation、Requirements Cognition、Parity Gate、Product Brief Validation、Missing Information、Backfill |
| [`contracts.md`](contracts.md) | `development-process.md` contract gates | Focused execution surface：Required Contracts、Contract Governance、Traceability、Contract-First Rules |
| [`ui-contracts.md`](ui-contracts.md) | `development-process.md` frontend / consumer contract gap | Focused execution surface：Screen Mapping、Consumer Contract、UI Behavior Contract、Screen Contract、Frontend ViewModel Contract、Accessibility Contract、Screen Traceability |
| [`ui-governance.md`](ui-governance.md) | `plans/archived/2026-06-08-1408-ui-governance-workflow.md` Phase 1 + [`plans/archived/2026-06-08-1544-evidence-acquisition-layer.md`](../../plans/archived/2026-06-08-1544-evidence-acquisition-layer.md) Phase 1 + [`responsive-ui`](../../intelligence/engineering/governance/responsive-ui/README.md) | Focused execution surface：UI governance classification order, evidence routing, and closure; responsive taxonomy / authority / severity are referenced from intelligence governance |
| [`test-strategy.md`](test-strategy.md) | `execution-flow.md` §2/§4 子節 + `development-process.md` BDD/Test Strategy gates | Focused execution surface：Docs-first BDD closure、Journey Specification、test strategy、mutation testing、test-first ordering |
| [`validation.md`](validation.md) | `execution-flow.md` §5/§7 + [`plans/archived/2026-06-08-1544-evidence-acquisition-layer.md`](../../plans/archived/2026-06-08-1544-evidence-acquisition-layer.md) Phase 1 | Focused execution surface：validation、Journey Validation、evidence acquisition execution、performance gate、old/new behavior proof、completion evidence |
| [`closure.md`](closure.md) | `execution-flow.md` §8 + `development-process.md` DoR/DoD | Focused execution surface：Definition of Ready/Done、handoff、close-loop、reusable lesson feedback |
| [`surgical-changes.md`](surgical-changes.md) | `execution-flow.md` §9 | Focused failure surface：surgical change discipline、diff purity、orphan cleanup boundary |
| [`requirements/pre-build-interrogation.md`](requirements/pre-build-interrogation.md) | mattpocock/skills `/grill-me` pattern + Ai-skill framework failure learning | Plan / implementation 前的需求拷問、framework discovery 與 source-of-truth duplication gate |
| [`artifact-gates.md`](artifact-gates.md) | `DOCUMENTATION.md`（已刪除） | Reusable Note Structure、Content Classification、Guidance Boundary、Linked Update Statement、Good Guidance Criteria |
| [`artifact-gates.yaml`](artifact-gates.yaml) | `artifact-gates.md` | Software delivery artifact executable contract：artifact shape、owner layer、sanitization、linked updates、quality gates |
| [`analysis/development-guidance/risk-translation.md`](../../analysis/development-guidance/risk-translation.md) | `WORKFLOW.md` §2-5（已刪除） | Risk Translation Table、Owner Layer Selection、Control Definition、Guidance Classification、Linked Updates |
| [`review-checklist.md`](review-checklist.md) | `skills/app-development-guidance/checklists/`（已刪除） | 6 種審查 checklist 的 catalog（Mobile Design Review、Mobile PR Review、Mobile Release Review、API Security Review、Contract Governance Review、Embedded Firmware Review） |
| [`development-process.md`](development-process.md) | `skills/app-development-guidance/process/README.md`（已刪除） | Contract-first 開發流程：Default Flow、Required Contracts、Product Brief Validation Gate、Change Intake Gate、Contract Governance Gate、Traceability Gate、BDD Execution Closure、Test Strategy Gate、Embedded/Hardware Flow、Missing Information Gate、Existing Project Documentation Backfill、Contract-First Rules、Definition of Ready/Done |

## 建議 invoke 流程（capability-centric）

### Post-implementation code review

```
1. Implementation 完成（tests green、same-session closure 進行中）。
2. invoke: capability code-review, context.caller_slice sd-implementation.
3. Runtime 依 registry 驗證 requires_context.stance（fault_finding）。
4. Consumer 載入 checklist bodies（cross-cutting/review/checklist.md）。
5. 產出 review report artifact → 進入 Validation。
```

### Pre-release review

```
1. Validation evidence 就緒。
2. invoke: capability release-review, context.caller_slice sd-validation.
3. 確認 build / signing / rollback plan → closure。
```

### Design / contract review（pre-implementation）

```
1. Contract / architecture artifacts 就緒。
2. invoke: capability contract-review 或 architecture-review.
3. Findings actionable → 回到 contracts / implementation。
```

細節 checklist 與 report 模板在 consumer 層 — execution-flow 不載入。

## 產出格式

本 workflow 各階段使用標準化輸出模板，確保產出格式一致、可追溯、可被後續階段自動消費：

| 模板 | 對應階段 | 用途 |
|------|---------|------|
| [`templates/product-impact-alignment-template.md`](templates/product-impact-alignment-template.md) | Product Impact Discovery | 記錄 Impact Map、Customer Journey Map、cross-check decision 與進入 BDD 前的缺口 |
| [`templates/change-brief-template.md`](templates/change-brief-template.md) | Change Intake | 記錄變更類型、證據、範圍、blocker 評估 |
| [`templates/contract-template.md`](templates/contract-template.md) | Contract Governance | 記錄 domain model、架構決策、API / error / consumer / UI 合約 |
| [`templates/bdd-scenario-template.md`](templates/bdd-scenario-template.md) | BDD Closure Loop | 記錄 requirement link、behavior boundary、Given/When/Then、acceptance criteria、validation target、regression scope |
| [`templates/implementation-plan-template.md`](templates/implementation-plan-template.md) | Implementation | 記錄任務拆解、檔案路徑、驗收條件、風險評估 |
| [`templates/review-report-template.md`](templates/review-report-template.md) | Review（6 種類型） | 記錄 finding、decision、reviewed artifacts |

每次 review 應產出：

- **Review 摘要**（≤200 tokens）：審查類型、範圍、verdict。
- **Finding 清單**（每個 finding ≤100 tokens）：問題描述、severity、建議修復方式。
- **Decision 記錄**（≤100 tokens）：最終決定、決定理據、相關連結。
