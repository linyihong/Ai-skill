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

## Phase 2 — Validator 擴充（consumer-layer）✅（2026-07-03）

**六條 acceptance gate（放行時使用者補全）**：
1. `enabled: true` ⇒ 4 必填（`brief.goal` / `brief.acceptance` / `brief.verification` / `execution.modes`，非空）→ block。
2. `brief.context`、`execution.constraints` 恆 optional。
3. 未宣告 `delegation` ⇒ **100% 舊行為**（byte-identical，非只是 no-error）。
4. **`enabled: false` 語義明確 = 等同未宣告**（不驗任何 brief；code 中與 `nil` 走同一 no-op 分支）。
5. **Consumer Exclusive**：portable engine 對 delegation 完全無感——機械鎖（source grep + `NormalizedPlanModel`/`RawPlan` reflection 雙證），consumer 刪除 engine 不需重編。
6. glossary 保留「future promotion requires cross-repository evidence」，防止功能成熟後有人略過證據門檻直接搬進 engine。

- [x] 擴充 **consumer-layer** `validatePlanTreeFrontmatter`（**未碰 `planvalidate` engine**）：`delegation.enabled: true` 時驗 4 必填（block）；`context`/`constraints` 不驗；未宣告 / `enabled:false` 不變。`plan_tree.go` 用 `yaml.v3` 解析 nested block（unmarshal 失敗 → `Delegation=nil` = 未宣告，保 zero-behavior-change）；`flexStrings` 容忍 scalar 或 list。
- [x] 測試（`plan_tree_delegation_test.go`）：complete+no-optional pass / 缺 verification fail / 空 modes fail / scalar brief pass / 未宣告 baseline pass / **`enabled:false`==未宣告輸出相等** / Consumer Exclusive engine-source-grep + reflection 雙鎖。全 suite 綠、`go vet` 乾淨。
- [x] 雙路徑 SOP 已於 Phase 1 落地（`plans/README.md` §Delegation）。
- [x] **Runtime Execution Path**：未新增 `route.*` / generated_surface / 新 validator dispatch——僅在既有 commit-msg validator `validatePlanTreeFrontmatter` 內新增 enabled-gated 子檢查（既有 dispatch，無新 wiring），故無 `validateRuntimeTriggerWiring` 觸發面。

**Phase 2 review 追加（2026-07-03，兩個反向鎖）**：
- [x] **Reverse-direction lock（consumer 不得懂工具）**：`execution` schema 凍結為恰 `{modes, constraints}`（`TestDelegation_ExecutionSchemaLockedToModesAndConstraints` reflection),機械擋掉未來 `execution.worktree`/`execution.<tool>`;且 validator **從不解讀 `constraints` 內容**（`TestDelegation_ConstraintsContentNeverInterpreted`：任意/工具狀/不存在路徑的 constraints 仍 pass）。方向:engine 不懂 delegation ⊕ consumer 不懂 workflow/tool。
- [x] **Capability Removal Test（可移除性證明，使用者唯一 merge gate）**：`TestDelegation_CapabilityRemoval_OnlyDelegationFindingsChange`——兩份僅差 delegation block 的 plan set,parent/archive-order/unique-id 三 validator byte-identical;frontmatter 移除 delegation 行後 == 無-delegation baseline(以共用的 non-delegation finding 保 header/trailer 兩邊都在,非空對空的 vacuous 比較)。證明:移除 delegation capability 只失去 delegation 驗證,不改任何其他 validator 語意。

### 完成條件（Phase 2）
- [x] enabled=true ⇒ 4 必填（gate 1）
- [x] context / constraints optional（gate 2）
- [x] 未宣告 ⇒ 100% 舊行為（gate 3）
- [x] enabled:false 明確 = 未宣告（gate 4）
- [x] Consumer Exclusive 機械證明（gate 5，grep + reflection）
- [x] glossary cross-repo promotion 門檻說明（gate 6）
- [x] Reverse-direction lock：consumer 不得解讀工具/workflow（review 追加）
- [x] Capability Removal Test：可完整移除、不影響其餘 validator（review 追加，merge gate）

## Phase 3 — Dogfood（回應 review #6：兩路徑各一次）

**焦點紀律（Phase 2 review 定調）**：Phase 3 驗的是 **brief 的獨立性（brief 是否真的形成 capability）**,**不是** Agent / human 的能力。任何「執行者不夠強」都要先反問「brief / `context.required` 是否寫得夠自足」。**不看** token 花費、不看執行者聰不聰明;只看「是否僅憑 brief（+ `context.required`）就能完成」。**保持小**——這已不是 architecture,是 dogfood。

**Brief Independence Score（Phase 3 驗收指標,非 pass/fail）**：
| 分數 | 判準 |
|---|---|
| ★★★★★ | 完全依 `brief` 完成,連 `context.required` 都沒回查 |
| ★★★★☆ | 依 `brief` + `context.required` 完成,未讀其他 |
| ★★★☆☆ | 需讀 `context.required` 以外、但仍在 repo 內的少量檔 |
| ★★☆☆☆ | 需回讀 main plan 才能完成（brief 未成熟） |
| ★☆☆☆☆ | 需人工補充需求（brief 不構成 capability） |
判準:human 路徑「是否一直回查 main plan」、agent 路徑「是否需 read whole repository」= 直接指向 `context.required` 寫得好不好,不是執行者問題。★★★★☆ 以上視為 brief 已形成 capability;★★★☆☆ 以下 → 回饋修正 brief/schema 後重跑。

- [x] 挑一個真實、小、可驗收的 sub-plan/task 設 `delegation.enabled: true` + 完整 brief → [`04-phase3-dogfood-glossary-spike.md`](04-phase3-dogfood-glossary-spike.md)（註冊 glossary `plan_profile`，真實 audit 缺口 ×46）。commit-msg delegation validator 接受此真實 brief（Phase 2 validator 首次真實 pass）。
- [ ] **human 路徑 evidence**：**完全關掉 main plan**,只給 `delegation.brief`,另一 session 能否獨立完成 → 記 Brief Independence Score。（brief 已交付使用者;待回填）
- [x] **agent 路徑 evidence：★★★★☆（2026-07-06）**。乾淨 agent、worktree、僅餵 brief → **只讀 `context.required`,未讀 main plan / 未 read whole repo**。產出正確、已採納 land 進 glossary（關閉真實缺口）。詳見 spike §Dogfood 結果。
- [x] **回饋迴路（非改 Agent）**：agent 揭露 brief 未 pin `introduced-by`/`owner-layer` → 修回 spike brief v2;worktree-not-in-brief 確認為 Layer 3 concern（capability/workflow 分層成立,非缺陷）。human 缺漏待你回填後併入。

## 完成條件
- [x] nested `delegation` schema（`brief` capability / `execution` workflow 分層）+ brief 契約落地（Q5 resolved：4 必填 + context/constraints optional）— Phase 1
- [x] 欄位 optional、向後相容（既有 sub-plan 不受影響）— Phase 2（`Undeclared_IdenticalToBaseline` / `EnabledFalse_EqualsUndeclared` 測試）
- [x] **consumer-layer** validator 擴充 + 測試通過（含 4 必填各 violation + no-context/constraints pass），**未碰 `planvalidate` engine**（Consumer Exclusive 雙鎖）— Phase 2
- [x] human + agent 雙路徑 SOP 落地（tool-neutral，Q6 resolved）— Phase 1（`plans/README.md` §Delegation）
- [ ] dogfood evidence：**agent ✓（★★★★☆，2026-07-06）** + **human 待使用者回填** — 唯一實質未決,gated on human path
- [x] 與 01 `plan_profile` 邊界對齊（`delegation` = **Ai-skill-only / consumer-layer**，非 portable — 見 Phase 0.1，2026-07-03）

## Glossary Impact
Glossary Impact: yes — 新增 `delegation`（nested 委派 schema，`brief` capability / `execution` workflow 分層，Ai-skill consumer-layer，非 portable）；Phase 1 落地時註冊到 `knowledge/glossary/ai-skill.md`。取代早期扁平 `agent_assignable` / `delegation_brief` 提案。

## 與其他 plans 的關係
- 擴充 [`archived/2026-06-02-1200-plan-tree-hierarchy-governance/_plan.md`](../../archived/2026-06-02-1200-plan-tree-hierarchy-governance/_plan.md) 的 frontmatter schema 與 `validatePlanTreeFrontmatter`。
- 依賴 [`01-external-repo-plan-system-shared-binary.md`](01-external-repo-plan-system-shared-binary.md) 的 `plan_profile` 邊界決定新欄位是否 portable。
