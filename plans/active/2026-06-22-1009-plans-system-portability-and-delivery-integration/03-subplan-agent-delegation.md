---
id: 2026-06-22-1009-subplan-agent-delegation
plan_kind: sub
status: in-progress
owner: linyihong
created: 2026-06-22
parent: 2026-06-22-1009-plans-system-portability-and-delivery-integration
required_for_completion: true
sub_plan_reason: >
  委派 schema 會擴充 plan-tree frontmatter，依賴 01 對 portable core 邊界的
  共識（新欄位算 portable 還是 Ai-skill-only？是否進 plan_profile？），
  因此排在 01 之後。獨立成 sub-plan 以便 delegation schema + 人工/agent 雙路徑
  契約可獨立設計與 sign-off，並可先 reservation（只設計 schema 不實作自動派發）。
---

# Sub-plan Agent Delegation（sub-plan）

**Status**: `in-progress`（Phase 0 done — delegation = Ai-skill-only）
**Owner**: linyihong
**Parent**: [`_plan.md`](_plan.md)

## Source Request
用子計畫系統讓一個 sub-plan 可交給其他 agent 執行。使用者澄清：**人工派發與 agent 派發兩種都要支援，依專案需求選用或並用**。

## Scope
- **In**：sub-plan frontmatter 新增 **nested `delegation` 物件**，`brief`（portable capability：goal / acceptance / verification / optional `context.required`）與 `execution`（workflow：`modes` / optional `constraints`）分層；人工派發與 agent 派發共用同一份 `brief` 的契約。
- **Out**：自動 orchestrator（自動偵測 + 自動 spawn agent + 自動收斂結果）— 保留為 future；本輪採 reservation pattern（schema + 雙路徑契約，不建自動 orchestrator）。
- **Affected**：plan-tree frontmatter schema、`scripts/ai-skill-cli/internal/app/plan_tree.go`（`PlanFrontmatter` struct + 可選 `validatePlanTreeFrontmatter` 擴充）、`governance/lifecycle/plan-tree-hierarchy.md`、`plans/README.md`。

## Decision Rationale（sub 層）
sub-plan 已具獨立 acceptance / archive，是天然的委派單元，缺的是**自足 brief 契約**：執行者（人或 agent）不讀整個 main plan 也能獨立完成。

**Schema 修正（回應 review #5）：避免把「可委派」與「brief 存在」綁死。** 扁平的 `agent_assignable: bool` + `delegation_brief` 會讓「是否可派」「派給誰」「brief」混在一起，未來要支援 manual-only / agent-only / hybrid / forbidden 就得破 schema。改用 **nested `delegation` 物件**：

**Capability-first 分層（回應 2026-07-03 review）**：`brief` 與 `execution` 分開——`brief` 是真正 portable 的能力描述（做什麼 / 何為完成 / 怎麼驗 / 需要什麼 context），`execution` 是 workflow（用哪些路徑執行、有什麼限制）。未來若冒出 `reviewer` / `scheduler` / `external-service` 等新路徑，全部疊進 `execution.modes`，`brief` 不動。

```yaml
delegation:
  enabled: true             # 開放委派（取代扁平 agent_assignable）
  brief:                    # Layer 2 portable capability — 不知道任何工具
    goal: <一句話目標>
    acceptance:             # 做到什麼算完成（必填，非空）
      - ...
    verification:           # 我要怎麼驗（必填，非空）
      - ...
    context:                # optional — 非所有 delegation 都需額外 context
      required:
        - plans/...
        - docs/...
  execution:                # workflow — 未來 reviewer/scheduler/external-service 也放這
    modes:                  # 至少一個（必填）：human / agent / ...
      - human
    constraints:            # optional — 恆 optional，可完全沒有 tool/worktree/sandbox
      - ...
```

**三層邊界（不可混）**：
- **Layer 1 portable engine**（`planvalidate/engine.go`）：只認 plan-structure invariant（unique id / parent resolution / archive ordering / required-sub completion / schema compat）。**engine 不知道 delegation。**
- **Layer 2 Ai-skill consumer**（`validatePlanTreeFrontmatter`）：delegation 住這。回答「這個 sub-plan 能不能交給另一個執行者」，不是「這是不是合法 plan tree」。
- **Layer 3 tool adapter**（`ai-tools/`）：worktree / Agent / 另一 session / Codex / Claude / CI agent 全是 adapter。**delegation 不知道它們**；工具細節只出現在 `execution.constraints` + `ai-tools/`。

雙路徑共用同一份 `brief`（tool-neutral，Q6）：human 路徑把 brief 貼給另一開發 / session；agent 路徑把 brief 餵給 Agent/Task 工具（可選 worktree，屬 Layer 3）。

### Alternatives
- A. 純文件慣例（sub-plan 標「可指派」但無 schema）：partial — 使用者要兩路徑且要可靠 brief，純慣例不足。
- B. 扁平 `agent_assignable` + `delegation_brief`：reject — 隱藏耦合（assignable 綁 brief 存在），不支援 mode 變化，未來必破 schema（review #5）。
- C. 只做 agent 自動 orchestrator：reject — over-engineering，使用者要人工也能用。
- D. nested `delegation` 物件 + reservation（accept）。

## Open Questions（本 sub）
- Q5（delegation schema 最小欄位集；人工/agent 共用同一 brief 是否可行）。
- Q6（agent 派發是否綁工具 vs tool-neutral brief 契約）。

## Phase 0 — Pre-Build Interrogation

### Phase 0.0 — Open Questions 核對（公版，必填）
- [x] 已讀 main + 本 sub §Open Questions 全部條目
- [x] 對每條標記 `resolved` / `still-open` / `deferred`
- [x] `resolved` 條目回寫（Phase 0.1 alignment 已決，見下）
- [x] 新問題已加入 §Open Questions（無新問題）

| Open Question | 處置 | 證據 / 原因 |
|---|---|---|
| Q5 delegation 最小契約 | still-open（形狀已定，待 Phase 1 落地） | 收斂到 **4 必填**：`brief.goal` / `brief.acceptance` / `brief.verification` / `execution.modes`；`context`/`constraints` optional。`brief`(capability) vs `execution`(workflow) 分層 |
| Q6 tool-neutral vs 綁工具 | still-open（方向已定，待 Phase 1 落地） | `brief` 全 tool-neutral；工具細節只在 `execution.constraints` + Layer 3 `ai-tools/` |

### Phase 0.1 — 架構盤點（需與 01 對齊 frontmatter schema；**不依賴外部 repo 能力**）
- [x] 讀 `plan_tree.go` `PlanFrontmatter` struct（`:36`，扁平）+ `validatePlanTreeFrontmatter`（`:274`，consumer-layer validator）。新增 optional 欄位不破壞既有 5 validators（它們只讀 parent/id/reason/required_for_completion）→ 向後相容成立。
- [x] **與 01 對齊決定（2026-07-03）：`delegation` = Ai-skill-only，NOT portable。** schema + 驗證落在 consumer 層 `validatePlanTreeFrontmatter`，**不進** `plan_profile.core`（`planvalidate/engine.go` 的抽出 portable engine）。理由：目前**零外部證據**有 repo 需要 delegation 驗證；promote 進 portable engine 是較重、外部採用後難回退的 rung；delegation 是 Ai-skill workflow dogfood 特性，非 universal plan-structure invariant。依 Q8 / falsification-ladder 紀律 → 留 consumer 層，待真實外部需求再 promote。
- [x] 確認新欄位為 **optional**（未宣告 `delegation` = 不可委派 / 行為不變；既有 sub-plan 不受影響）。

## Phase 1 — Delegation schema + brief 契約設計 ✅（doc-only，2026-07-03）
- [x] 定義 nested `delegation` 物件（optional，未宣告 = 不可委派 / 行為不變），`brief`（capability）與 `execution`（workflow）分層。
- [x] **Q5 必填集 = 恰 4 個**（`enabled: true` 時）：`brief.goal` / `brief.acceptance`（非空）/ `brief.verification`（非空）/ `execution.modes`（非空，至少一路徑）。`brief.context`（`context.required`）與 `execution.constraints` **恆 optional**。
- [x] tool-neutral（Q6）：`brief` 不綁任何工具；工具/隔離細節只出現在 `execution.constraints` + `ai-tools/`（Layer 3）。
- [x] 文件化：`governance/lifecycle/plan-tree-hierarchy.md` §Delegation（**明確標記 consumer-layer 擴充，非 portable invariant，不進既有不變式表**）+ `plans/README.md` §Delegation（雙路徑 SOP）+ glossary `delegation` 詞條。

## Phase 2 — Validator 擴充（optional 欄位）+ 雙路徑說明
- [ ] 擴充 **consumer-layer** `validatePlanTreeFrontmatter`（**不碰 `planvalidate` engine**）：當 `delegation.enabled: true` 時，驗 4 必填（`brief.goal` / `brief.acceptance` 非空 / `brief.verification` 非空 / `execution.modes` 非空）（block）；`context`/`constraints` 不驗必填；未宣告 `delegation` 則不變（向後相容）。
- [ ] 測試：tmp fixture（enabled+4 必填齊 pass / enabled+缺 verification fail / enabled+空 modes fail / enabled 但無 context+constraints pass / 未宣告 pass）。
- [ ] 文件化雙路徑：human 派發 SOP + agent 派發 SOP（後者可選 worktree isolation；保持 tool-neutral，工具細節放 `ai-tools/`）。
- [ ] **若擴充 validator 行為，補 Runtime Execution Path trigger flow**（commit-msg validator 已是既有 dispatch，新增子驗證須宣告）。

## Phase 3 — Dogfood（回應 review #6：兩路徑各一次）
- [ ] 挑一個真實 sub-plan 設 `delegation.enabled: true` + 完整 brief。
- [ ] **human 路徑 evidence**：另一 session / 開發僅憑 brief 獨立完成一次。
- [ ] **agent 路徑 evidence**：以 Agent/Task 工具僅憑 brief 在 worktree 執行一次。
- [ ] 兩次皆記錄 brief 是否足夠自足（缺漏回饋修正 schema）。

## 完成條件
- [ ] nested `delegation` schema（`brief` capability / `execution` workflow 分層）+ brief 契約落地（Q5 resolved：4 必填 + context/constraints optional）
- [ ] 欄位 optional、向後相容（既有 sub-plan 不受影響）
- [ ] **consumer-layer** validator 擴充 + 測試通過（含 4 必填各 violation + no-context/constraints pass），**未碰 `planvalidate` engine**
- [ ] human + agent 雙路徑 SOP 落地（tool-neutral，Q6 resolved）
- [ ] dogfood evidence：human 一次 + agent 一次
- [x] 與 01 `plan_profile` 邊界對齊（`delegation` = **Ai-skill-only / consumer-layer**，非 portable — 見 Phase 0.1，2026-07-03）

## Glossary Impact
Glossary Impact: yes — 新增 `delegation`（nested 委派 schema，`brief` capability / `execution` workflow 分層，Ai-skill consumer-layer，非 portable）；Phase 1 落地時註冊到 `knowledge/glossary/ai-skill.md`。取代早期扁平 `agent_assignable` / `delegation_brief` 提案。

## 與其他 plans 的關係
- 擴充 [`archived/2026-06-02-1200-plan-tree-hierarchy-governance/_plan.md`](../../archived/2026-06-02-1200-plan-tree-hierarchy-governance/_plan.md) 的 frontmatter schema 與 `validatePlanTreeFrontmatter`。
- 依賴 [`01-external-repo-plan-system-shared-binary.md`](01-external-repo-plan-system-shared-binary.md) 的 `plan_profile` 邊界決定新欄位是否 portable。
