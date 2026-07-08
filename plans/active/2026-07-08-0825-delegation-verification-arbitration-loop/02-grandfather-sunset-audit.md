# Grandfather Sunset 前置調查 — pre_2026_05_28_doc_only_completion

> **產出脈絡**：本檔為 grandfather flag（`governance/lifecycle/system-upgrade-governance.yaml`
> §`pre_2026_05_28_doc_only_completion`，L300–354）2026-08-31 sunset 的前置事實調查，
> 由 delegation-verification-arbitration-loop 的隔離調查者產出。
> **處置決定保留給 maintainer** — 本報告只提供現況事實、判定與建議。
> 調查日：2026-07-08。所有行號以本 worktree HEAD（`bcfdda7`）為準。

## 核心發現（TL;DR）

**Flag 條款所描述的 orphan 狀態已全部過時。** Grandfather flag 於 2026-05-28 由 commit
`4c0b514` 寫入；**同日稍晚** commit `f222fdd`（gen3 plan Phase 4）即已把全部 5 條
covered surfaces 補 wire 完成。今日 `ai-skill runtime audit` 確認 4 個 covered plans
的每條 surface 都分類為 auto-detected / consumed / intentionally-manual —
**post_sunset_evaluation_rule 的選項 (a) 已對全部 4 個 plan 滿足**，無任何 surface
需要降 orphan。剩餘工作是「行政收尾」：更新 plans/README.md 的 4 個 ⚠️ 標籤與
flag 條款自身的 sunset 處理。

### Audit 分類總覽（`ai-skill runtime audit --json`，2026-07-08 執行）

| Surface | Covered plan | Audit 分類 | Audit evidence 欄 |
|---|---|---|---|
| `route.governance.cognitive-state-evidence` | plan 1 | **auto-detected** | discovery signal description references route |
| `enforcement.evidence_hierarchy.contract` | plan 1 | **consumed** | Go source references target_key |
| `route.memory.retrieval-activation` | plan 2 | **auto-detected** | discovery signal description references route |
| `route.models.model-aware-routing` | plan 3 | **auto-detected** | discovery signal description references route |
| `route.runtime.cognitive-modes` | plan 4 | **intentionally-manual** | manual_activation reason: validators_consume_by_file_path_not_route_lookup |

---

## Plan 1 — `plans/archived/2026-05-20-1501-cognitive-state-evidence-governance.md`

### Surface 1a: `route.governance.cognitive-state-evidence`

**現況事實**：

- Route 仍存在於 `knowledge/runtime/routing-registry.yaml` **L786**
  （`- id: route.governance.cognitive-state-evidence`）。
- Consumer（discovery signal）已存在：`runtime/cognitive-modes-discovery.yaml`
  **L212–220**，signal `user_keyword_evidence_claim`（signal_type: user_keyword，
  pattern `完成|completed|✅|done|證據|evidence|claim|confidence|inflated`，
  governance_mode: STRICT），其 **L220** description 明確引用
  `route.governance.cognitive-state-evidence + enforcement.evidence_hierarchy.contract`。
  該區塊上方 L211 註解即標明 `Phase 4 high-priority orphan wires (gen3-runtime-trigger-audit)`。
- `ai-skill runtime audit` 分類：**auto-detected**。

**判定**：(a) **已 wired** — consumer 為 discovery signal `user_keyword_evidence_claim`
（`runtime/cognitive-modes-discovery.yaml` L212–220）。

### Surface 1b: `enforcement.evidence_hierarchy.contract`

**現況事實**：

- Generated surface 仍存在（source: `enforcement/evidence-hierarchy.yaml`；audit
  surfaces 表列出該 target_key）。
- Consumer（commit-msg validator）已存在：`scripts/ai-skill-cli/internal/app/hooks.go`
  - **L3402–3418**：`validateEvidenceHierarchy` 函式，註解明載 "consumes the executable
    contract `enforcement.evidence_hierarchy.contract` (source: enforcement/evidence-hierarchy.yaml)"。
  - **L4005–4007**：註冊於 validator registry map（`obligation.commit.evidence_hierarchy`）。
  - **L4061**：列於 `defaultCommitMsgDispatchOrder`（Phase 6 dispatcher 的 fallback 順序，
    hooks.go L2489–2490 為 dispatcher 本體）。
  - 對應 obligation 已註冊於 `runtime/core-bootstrap.yaml` **L377–382**
    （`obligation.commit.evidence_hierarchy`，severity: block，opt-out `[skip-evidence-hierarchy]`）。
  - Fixture tests 存在：`scripts/ai-skill-cli/internal/app/evidence_hierarchy_test.go`。
- `ai-skill runtime audit` 分類：**consumed**。
- 這正是 flag `wire_plan` 提示（"add validateEvidenceHierarchy commit-msg validator"）
  所指的動作 — 已由 gen3 plan Phase 4（commit `f222fdd`，2026-05-28）落地。

**判定**：(a) **已 wired** — consumer 為 `validateEvidenceHierarchy`
（hooks.go L3402–3462 + dispatch L4005/L4061）。

### 處置建議（plan 1）

**升 auto-detected / consumed（維持現狀，行政收尾）**。兩條 surface 均已由聲明的
wire_plan 路徑滿足：route 由 `user_keyword_evidence_claim` signal 滿足 auto-detected；
contract 由 `validateEvidenceHierarchy` 滿足 consumed。無需補 wire、無需降 orphan。
建議收尾動作：更新 `plans/README.md` **L295** 的 ⚠️ 標籤（該行仍寫「無 discovery
signal / commit-msg validator 消費」，與現況不符）。

---

## Plan 2 — `plans/archived/2026-05-20-1745-memory-retrieval-activation-governance.md`

### Surface: `route.memory.retrieval-activation`

**現況事實**：

- Route 仍存在於 `knowledge/runtime/routing-registry.yaml` **L1095**。
- Consumer（discovery signal）已存在：`runtime/cognitive-modes-discovery.yaml`
  **L222–230**，signal `file_diff_memory_layer`（signal_type: file_diff_scope，
  pattern `^enforcement/memory-|^enforcement/conversation-goal-ledger|^memory/`，
  memory_mode: SELECTIVE_REPLAY），其 **L230** description 明確引用
  `route.memory.retrieval-activation`。
- `ai-skill runtime audit` 分類：**auto-detected**。
- Flag `wire_plan` 提示為 "add discovery signal in runtime/cognitive-modes-discovery.yaml
  OR Go consumer" — 已由前者（discovery signal 路徑）滿足。

**判定**：(a) **已 wired** — consumer 為 discovery signal `file_diff_memory_layer`
（`runtime/cognitive-modes-discovery.yaml` L222–230）。

### 處置建議（plan 2）

**升 auto-detected（維持現狀，行政收尾）**。wire_plan 的 discovery-signal 選項已落地。
收尾動作：更新 `plans/README.md` **L296** ⚠️ 標籤（該行仍寫「audit 仍判 orphan」，
與現況不符）。

---

## Plan 3 — `plans/archived/2026-05-20-1802-model-aware-execution-routing.md`

### Surface: `route.models.model-aware-routing`

**現況事實**：

- Route 仍存在於 `knowledge/runtime/routing-registry.yaml` **L1059**。
- Consumer（discovery signal）已存在：`runtime/cognitive-modes-discovery.yaml`
  **L232–240**，signal `file_diff_model_selection`（signal_type: file_diff_scope，
  pattern `^models/|model-aware|execution-strategy`，context_mode: SOURCE_BACKED），
  其 **L240** description 明確引用 `route.models.model-aware-routing`。
- `ai-skill runtime audit` 分類：**auto-detected**。
- Flag `wire_plan` 提示為 "add commit-msg validator referencing model selection contract
  OR discovery signal" — 已由後者（discovery signal 路徑）滿足。

**判定**：(a) **已 wired** — consumer 為 discovery signal `file_diff_model_selection`
（`runtime/cognitive-modes-discovery.yaml` L232–240）。

### 處置建議（plan 3）

**升 auto-detected（維持現狀，行政收尾）**。收尾動作：更新 `plans/README.md`
**L297** ⚠️ 標籤（該行仍寫「無 commit-msg validator 引用」— 字面上仍為真，但
audit 判準已由 discovery signal 滿足，標籤的 doc-only 定性已過時）。

---

## Plan 4 — `plans/archived/2026-05-22-1629-runtime-cognitive-modes-system.md`

### Surface: `route.runtime.cognitive-modes`

**現況事實**：

- Route 仍存在於 `knowledge/runtime/routing-registry.yaml` **L2574**，且 **L2577–2581**
  已帶 `manual_activation` 註記：
  - reason: `validators_consume_by_file_path_not_route_lookup`
  - note 明載：cognitive modes 由 commit-msg validators（`validateExecutionModeFloors`、
    `validateGovernanceModeConsistency` 等）按檔案路徑 heuristic 消費，route 本身供人類
    導覽；並明文寫入 "Grandfather-flagged per pre_2026_05_28_doc_only_completion until
    2026-08-31 + conditional extension to 2026-11-30, **then this annotation makes it
    intentionally-manual permanently**"。
- 該 note 所稱的實際 consumers 可證：`validateGovernanceModeConsistency` 等 validators
  存在於 hooks.go（plan 4 自身 §Phase 3 表格 L581 亦記錄 `validateGovernanceModeConsistency`
  in hooks.go 完成），且 `runtime/core-bootstrap.yaml` L281–316 列出 cognitive-mode
  家族的 per-commit validators（contract_source 指向 runtime/cognitive-modes-*.yaml）。
- `ai-skill runtime audit` 分類：**intentionally-manual**
  （evidence: `manual_activation reason: validators_consume_by_file_path_not_route_lookup`）。
- Flag `wire_plan` 提示（"candidate for manual_activation annotation"）— 已落地，
  且正是 audit 判 intentionally-manual 的依據。

**判定**：(a) **已 wired（intentionally-manual 路線）** — `manual_activation` 註記
已存在於 routing-registry.yaml L2577–2581，audit 承認該分類。

### 處置建議（plan 4）

**升 intentionally-manual（維持現狀，行政收尾）**。wire_plan 的 manual_activation
註記路線已完成，且註記自述 sunset 後「permanently intentionally-manual」— sunset 到期時
此 surface 自動滿足規則 (a) 的 intentionally-manual 分支，maintainer 只需（可選）把
註記 note 中的 grandfather 引用改寫為歷史陳述。收尾動作：更新 `plans/README.md`
**L277** ⚠️ 標籤（該行仍寫「Phase 4 將決定升 manual_activation 或補 signal」— 該決定
已做出並落地）。

---

## 延展條件核對（conditional_extension_trigger @ 2026-08-15 checkpoint）

規則：2026-08-15 時**任一**條件成立 → primary_sunset 2026-08-31 自動延展至 2026-11-30。

### 條件 1: `audit_tool_age_lt_60_days`

- **事實**：`ai-skill runtime audit` 由 commit `0f53e91`（2026-05-28，"feat(audit): add
  ai-skill runtime audit subcommand with 4-way classification"）引入。
  `git log --follow -- scripts/ai-skill-cli/internal/app/runtime_audit.go` 顯示該檔
  **僅此一個 commit**（無後續 rewrite 可能重置 "age" 的解讀空間）。
- **推算**：2026-05-28 → 2026-08-15 = **79 天** ≥ 60。
- **判定**：條件**不成立**（audit tool age ≥ 60 days）。

### 條件 2: `phase_4_high_priority_wires_incomplete`

- **事實**：Phase 4 所指為 gen3 plan（現已 archived：
  `plans/archived/2026-05-28-1200-gen3-runtime-trigger-audit-and-completion.md`）。
  其 §Phase 4（L315–341）全部 checkbox 為 `[x]`，含「5 wires 明細」表（L325–333）
  逐條列出本 flag 5 條 surface 的 before/after 分類與 wire 方式；§Phase 4 完成條件
  （L335–341）亦全 `[x]`（≥5 wired 達標 5/5、audit 升級確認、8 個 fixture tests 綠）。
  Wire commit `f222fdd`（2026-05-28，"feat(audit): wire 5 high-priority orphans
  (Phase 4 complete)"）可查。本次 audit 重跑（2026-07-08）確認 5 條 surface 分類
  與該表 after 欄一致。
- **判定**：條件**不成立**（Phase 4 high-priority wires 已 complete）。

### 結論

兩條件均不成立 → **2026-08-15 檢查點不會觸發延展**，primary_sunset **2026-08-31 生效**。
由於全部 covered surfaces 已滿足 post_sunset_evaluation_rule 選項 (a)，sunset 生效時
無需任何降 orphan / registry 移除動作；flag 本身依規則於該日 sunset。

---

## 給 maintainer 的整體建議（決定權在 maintainer）

1. **無需任何補 wire 或降 orphan 動作** — 5/5 surfaces 已分類達標（見上表）。
2. **行政收尾清單**（可在 sunset 日或之前一次完成）：
   - `plans/README.md` L277 / L295 / L296 / L297：4 個 ⚠️ 標籤改回 completed
     （或改註「wired 2026-05-28, grandfather flag sunset 2026-08-31」），因為這些行
     描述的 orphan 狀態已與 audit 事實不符。
   - `governance/lifecycle/system-upgrade-governance.yaml` §pre_2026_05_28_doc_only_completion：
     sunset 日將 `status: active` 改為 sunsetted/resolved（或依 governance 慣例移除
     並留歷史指標）。條款自述 "the flag itself is sunsetted on this date regardless"。
   - `knowledge/runtime/routing-registry.yaml` L2577–2581 note：（可選）把 grandfather
     引用改為歷史陳述，明確 intentionally-manual 為永久狀態。
   - 相關 scenario `validation/scenarios/failure-derived/pre-2026-grandfather-coverage-v1.yaml`
     （flag §related_scenario）屆時需同步檢視是否更新 — 本調查未展開該檔內容（見 unverified）。
3. **時序備註**：flag 條款（`4c0b514`）與 wire 動作（`f222fdd`）同為 2026-05-28 落地，
   flag 寫入在先、wire 在後，條款文字從未回頭更新 — 這是「宣告 vs 事實」落差的來源，
   不是 regression。

## Unverified / 範圍外

- `validation/scenarios/failure-derived/pre-2026-grandfather-coverage-v1.yaml` 的內容
  與其在 sunset 後的預期行為：**unverified**（未讀該檔；不影響 4 個 plan 的 surface 判定，
  但屬 flag 收尾時應檢視的關聯物）。
- post_sunset_evaluation_rule 文字提及 "route / surface / **scenario**"，但 covered_plans
  只列 orphan_surfaces（不含 scenario 清單）；audit 顯示全 repo 仍有 205 orphan scenarios。
  covered plans 是否有 scenario 層級的殘留義務：**unverified**（flag 條款未列，本調查
  依條款列出的 surfaces 為準）。
- 「audit tool age」的起算定義（首次引入 vs 最近重大改版）條款未明文；本報告採「首次
  引入」且該檔無後續 commit，兩種解讀結論相同，故不影響判定。

## 證據索引（可覆核指令）

```
# Surface 分類
./scripts/ai-skill-cli/bin/ai-skill-darwin-arm64 runtime audit --json
# Registry 行號
grep -n "route.governance.cognitive-state-evidence\|route.memory.retrieval-activation\|route.models.model-aware-routing\|route.runtime.cognitive-modes" knowledge/runtime/routing-registry.yaml
# Discovery signals
sed -n '200,250p' runtime/cognitive-modes-discovery.yaml
# Go validator
grep -n "validateEvidenceHierarchy" scripts/ai-skill-cli/internal/app/hooks.go
# 時序
git log --format='%h %ad %s' --date=short -1 0f53e91   # audit tool 2026-05-28
git log --format='%h %ad %s' --date=short -1 4c0b514   # flag 條款 2026-05-28
git log --format='%h %ad %s' --date=short -1 f222fdd   # 5 wires 2026-05-28
# gen3 Phase 4 checkbox
sed -n '315,341p' plans/archived/2026-05-28-1200-gen3-runtime-trigger-audit-and-completion.md
```
