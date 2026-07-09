# Plan Tree Hierarchy（governance rule）

> **Status**：validated（2026-06-04，plan `2026-06-02-1200-plan-tree-hierarchy-governance` Phase 5 收尾；5 validators 已落地並機械強制）
> **Canonical source**：本檔為 plan-tree governance 的人類可讀 rule。Schema enum
> 與 validator dispatch 已進 `scripts/ai-skill-cli/internal/app/plan_tree.go`，
> 由 commit-msg hook 機械強制。本檔不重複定義 schema 細節，detail 連 plan。

## 目的

定義 plan 的階層治理規則，避免 `plans/active/` 扁平化導致 main / sub plan 關係
不可追蹤。Hierarchy 由 frontmatter `parent` pointer 決定（**單一 source of truth**），
folder + filename 為 UI convention。

## 核心原則（Minimal Governance Model）

| 原則 | 規範 |
|---|---|
| Single source of truth | `id` + `parent` 是 hierarchy 唯一真實來源；不維護 `children:`，runtime scan 推導 |
| UI vs runtime 分離 | folder + `_plan.md` + `NN-` 前綴是 UI convention，validator 只發 warning；hierarchy 由 frontmatter pointer 決定 |
| Lifecycle vs storage 分離 | `status` 是生命週期；`plans/active/` vs `plans/archived/` 是儲存位置。Archive gate 只看 `status`，不混 location |
| 業務語意優先 | `required_for_completion: bool` 描述「是否屬於 parent acceptance」；archive blocker 由此推導，不直接命名為機制 |
| Free-text reason | `sub_plan_reason` 是 free text，validator 只擋空字串；不維護 enum，避免 framework 隨情境膨脹 |
| Warning over restriction | 深度 ≥ 3 發 warning，不立硬上限；讓真實案例決定 |

## Frontmatter 規範（簡述；canonical 見 plan）

主計畫必填：`id` / `plan_kind: main` / `status` / `owner` / `created` / `parent: null`

Sub-plan 必填：上列 + `parent: <main-id>` / `required_for_completion: bool` / `sub_plan_reason: <非空 free text>`

Spike：等同 sub-plan，建議 `required_for_completion: false`。

詳細 schema、fixture 範例見：
- [`plans/active/2026-06-02-1200-plan-tree-hierarchy-governance/_plan.md`](../../plans/active/2026-06-02-1200-plan-tree-hierarchy-governance/_plan.md)
- [`plans/active/2026-06-02-1200-plan-tree-hierarchy-governance/fixtures/`](../../plans/active/2026-06-02-1200-plan-tree-hierarchy-governance/fixtures/)

## 何時開 sub-plan

`sub_plan_reason` 非空為 validator 唯一強制條件。**不審內容**。

下列為**建議**參考（recommended triggers），寫進 reason 時可引用但不強制：

- Independent sign-off — 該支線需獨立 stakeholder 簽核
- Multi-phase with own acceptance — 該支線跨 ≥ 3 個 phase 有獨立 completion criteria
- Independent runtime trigger — 該支線有自己的 routing-registry / generated_surface
- Parallel owners — 兩個 agent / session 需 owner-lock 分隔
- Independent archive (spike) — 工作完成可獨立 archive，parent 仍 in-progress

未來出現第 6 / 7 種情境，**直接寫進 `sub_plan_reason` 即可，不需升 framework**。
重複 ≥ 3 次再評估 promote 為 recommended trigger 範例。

**不該開 sub-plan**：
- 單一 phase 內的 step → checkbox
- < 1 工作 session → inline 寫進主計畫
- 純文件補強 / rename / typo → 直接 commit，不開 plan
- 同 acceptance 下不同 angle → 同 plan 多 phase

## Validator 行為（已落地 `plan_tree.go`）

| Validator | Severity | 規則 |
|---|---|---|
| `validatePlanTreeFrontmatter` | block | sub-plan 缺 `parent` / `sub_plan_reason`（空字串視為缺）/ `required_for_completion` |
| `validatePlanTreeArchiveOrder` | block | 主計畫 archive 時，所有 `parent == <main>` 且 `required_for_completion: true` 的 sub-plan 必須 `status: completed`。**只看 status，不看 location** |
| `validatePlanTreeParentReference` | block | sub-plan `parent` 指向的 id 必須存在於 active + archived 全集；防 orphan node |
| `validatePlanTreeUniqueID` | block | 全 repo plan `id` 必須唯一；防 parent pointer 指錯 |
| `validatePlanTreeFolderConvention` | warning | folder 缺 `_plan.md` / 檔名不符 `NN-` 前綴 / 深度 ≥ 3（**`evidence/` 子目錄豁免**）/ **頂層 flat multi-file cluster** |
| `validatePlanEvidenceConvention` | block | `evidence/` 存在時：`evidence/README.md` 必填、Run 索引覆蓋所有 evidence `.md`、禁行號引用（warning 另發）— 見 [`plan-evidence.md`](plan-evidence.md) |

Sub-plan 之間的依賴（`depends_on` → DAG）目前 **不在治理範圍**；
promotion gate：≥ 3 個自然發生的 C-depends-on-A 案例後再評估。

## Delegation（Ai-skill consumer-layer 擴充，**非 portable invariant**）

> **邊界警告**：`delegation` **不是** plan-tree 不變式，**不屬** portable engine（`planvalidate/engine.go`）。它是 Ai-skill 的 **workflow contract**，只由 consumer-layer `validatePlanTreeFrontmatter` 驗證。採用 canonical plan tree 的外部 repo **不需要**、也不會被強制 delegation。上方「Frontmatter 規範」「Validator 行為」兩表是 portable invariant；本節刻意分開，勿併入。目前零跨 repo 證據 → 維持 Ai-skill-only；未來若有真實外部需求，再依 decision-promotion-pipeline 評估是否 promote 進 engine。

**用途**：讓一個 sub-plan 可交給另一個執行者（人或 agent）獨立完成，不需讀整個 main plan。**optional**：未宣告 `delegation` = 不可委派 / 行為不變（既有 sub-plan 不受影響）。

**三層邊界**：
- **Layer 1 portable engine** — plan-structure invariant（unique id / parent / archive order / required-sub completion / schema compat）。engine 不知道 delegation。
- **Layer 2 Ai-skill consumer** — delegation 住這。回答「能不能交給另一執行者」，不是「是不是合法 plan tree」。
- **Layer 3 tool adapter**（`ai-tools/`）— worktree / Agent / 另一 session / Codex / CI agent 全是 adapter。delegation 不知道它們；工具細節只在 `execution.constraints` + `ai-tools/`。

**Schema**（capability-first：`brief` 是 portable 能力描述，`execution` 是 workflow）：

```yaml
delegation:
  enabled: true             # 開放委派
  brief:                    # capability — tool-neutral
    goal: <一句話目標>
    acceptance:             # 做到什麼算完成（必填，非空）
      - ...
    verification:           # 怎麼驗（必填，非空）
      - ...
    context:                # optional
      required: [ plans/..., docs/... ]
  execution:
    modes: [ human ]        # 至少一個（必填）：human / agent / …（未來 reviewer/scheduler 也放這）
    constraints: [ ... ]    # optional — 恆 optional，可完全沒有 tool
```

**必填集（`enabled: true` 時，validator block）= 恰 4 個**：`brief.goal`、`brief.acceptance`（非空）、`brief.verification`（非空）、`execution.modes`（非空）。`brief.context`、`execution.constraints` 皆 optional。

**雙路徑共用同一份 `brief`**：human = 把 brief 貼給另一 session / 開發；agent = 把 brief 餵給 Agent/Task 工具（可選 worktree，屬 Layer 3）。SOP 見 [`plans/README.md`](../../plans/README.md) §Delegation。

## 計畫切檔決策（同一主題，三種合法形狀）

Agent 或作者在 plan 變長、要拆檔時，先判斷屬於哪一類——**不要**把「相關的兩個 main plan」誤當成「同一 plan 的多檔 companion」。

| 情境 | 判斷訊號 | 儲存形狀 | 範例 |
| --- | --- | --- | --- |
| **A. 同一 plan 的附檔** | 檔名為 `<slug>.md` + `<slug>-<companion>.md`；同一 `id` 語意；companion 無獨立 acceptance | `<slug>/_plan.md` + `NN-<companion>.md` | [`2026-06-29-1430-preparatory-refactoring-workflow/`](../plans/active/2026-06-29-1430-preparatory-refactoring-workflow/_plan.md)（dogfood evidence） |
| **B. 獨立 acceptance 的支線** | 需獨立 sign-off / archive / multi-phase gate；有獨立 `id` | plan tree：`parent` + `plan_kind: sub`；可選 folder + `NN-` | [`2026-06-22-1009-plans-system-portability-and-delivery-integration/`](../plans/archived/2026-06-22-1009-plans-system-portability-and-delivery-integration/_plan.md) |
| **C. 序貫的兩個 main plan** | 不同 `id`、不同 timestamp-slug；後者引用前者 baseline，但各自是 main | **維持**兩個頂層 `.md`；用 frontmatter / prose 連結，**不** folderize | diagnosis [`2026-06-23-1500-adr-004-migration-drift-diagnosis`](../plans/archived/2026-06-23-1500-adr-004-migration-drift-diagnosis.md) → completion [`2026-06-24-1100-adr-004-migration-completion`](../plans/archived/2026-06-24-1100-adr-004-migration-completion.md)（`baseline_ref`；已 closed/archived） |

**機械偵測**：`validatePlanTreeFolderConvention` 只對 **A** 發 flat-cluster warning；`ai-skill plans folderize --dry-run` 預覽遷移。**B** 走 sub-plan frontmatter；**C** 不觸發 folderize（`plans folderize --dry-run` 全 repo 掃描為空）。

## 與其他 governance rule 的關係

- [`directory-structure-governance.md`](directory-structure-governance.md) — plan tree 的 folder convention 不違反目錄治理原則
- [`decision-promotion-pipeline.md`](decision-promotion-pipeline.md) — plan completed 後若符合 promotion criteria，依此 pipeline 升 ADR
- `plans/README.md` — plan 模板必填章節（main plan 適用；sub-plan 模板較簡，見 plan `_plan.md` §Decision Rationale Q3 resolution）

## 變更流程

修改本 rule 前：
1. 先改 plan `_plan.md` 的對應章節（Decision / Schema / Open Questions）
2. 同步本檔（companion）
3. 若涉及 validator 行為改動，連動更新 `scripts/ai-skill-cli/internal/app/plan_tree.go`（Phase 2 後存在）

← [Back to lifecycle governance index](README.md)
