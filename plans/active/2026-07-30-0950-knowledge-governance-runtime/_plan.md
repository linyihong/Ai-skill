---
id: 2026-07-30-0950-knowledge-governance-runtime
plan_kind: main
status: in-progress
owner: linyihong
created: 2026-07-30
last_updated: 2026-07-30
required_for_completion: false
parent: null
glossary_impact: candidate
canonical_name: Knowledge Governance Engine
canonical_abbrev: KGE
legacy_working_title: Knowledge Governance Runtime (KGR)
current_phase: 1
revision:
  - date: 2026-07-30
    note: "初稿落檔 — KGR 定位、兩層 Framework/Plugin、pipeline draft"
  - date: 2026-07-30
    note: "審閱整併 — 定名 KGE；Pipeline 改 Context Builder + optional Projection/Dependency；補 Capability Registry；升級 Architecture Evolution 定位；記錄與 Delegation Loop 收斂觀察"
  - date: 2026-07-30
    note: "進入 Phase 0 — 目標改為驗證 Architecture Hypothesis（非功能盤點）；Phase 0 內含 Mini Spike；首批 mapping 目標 = Commit Message / README Sync / Linked Updates（暫不搬 Plans）"
  - date: 2026-07-30
    note: "Phase 0 首跑 evidence/p0-architecture-mapping.md — 五問暫過；Mini Spike A=cognitive_cost、B=cli_doc_sync 紙上完成；~7 capability 工作集；發現 cli_doc_sync 內嵌 git 需在 Phase 1 剝離"
  - date: 2026-07-30
    note: "使用者簽核 Phase 0 Exit → 進入 Phase 1（Contracts + Thin Engine）；Rule A/B 實作路徑授權"
  - date: 2026-07-30
    note: "Portable copy boundary — engine 移至 scripts/ai-skill-cli/portable/kge/（可整夾複製到外部專案）；Ai-skill 只留 Adapter wiring；刻意多檔不壓成單檔"
  - date: 2026-07-30
    note: "Phase 1 adapter — internal/app/kge_adapter.go 委派 cognitive_cost + cli_doc_sync；git diff 僅在 adapter；Q2 工作假設 first-party Go + 外部 copy pack"
  - date: 2026-07-30
    note: "補 Enforcement→KGE 對照 — taxonomy 已有但缺遷移清單；釐清 validation(block)/advisory(warn)/out-of-scope；document_sizing=advisory 候選；neutral_language=not_mechanizable 不強制進 KGE block"
  - date: 2026-07-30
    note: "D9 Adapter Presentation Policy — Q6 改問法並拍板：Severity×Adapter=Presentation；commit=checkpoint（count+pointer）；push/kge check=watershed summary；IDE=完整 advisory；validate --advisory=完整"
  - date: 2026-07-30
    note: "Phase 2b+4 slice — document_sizing advisory rule；ai-skill kge check/validate --advisory；D9 presentation helpers；RunAvailable"
  - date: 2026-07-30
    note: "D9 git+IDE adapters — commit-msg advisory count-only；pre-push 掛 kge check；ai-skill kge diagnose JSON；Finding.Path；Windows /tmp→os.TempDir 已修"
  - date: 2026-07-30
    note: "Phase 2a batch 2 — glossary_retro_own、runtime_yaml_projects、token_budget 遷入 portable/kge"
  - date: 2026-07-30
    note: "Phase 2a batch 3 — execution_mode_floors、governance_mode_consistency、memory_mode_subdir + shared path_policy"
  - date: "2026-07-30"
    note: "Phase 2a batch 4 — activation_signals、capability_snippet、adaptive_triggers + CapKnownSignals"
  - date: "2026-07-30"
    note: "Phase 2a batch 5 — evidence_hierarchy、plan_status_sync（git-free CapCommitMsg+CapStagedPaths）"
  - date: "2026-07-30"
    note: "Phase 2a batch 6 — plan_checkbox_sync（PathDiffs/CapStagedDiff）、plan_archival_audit（CapStagedContent）"
---

# Knowledge Governance Engine（知識治理引擎）

**Status**: `in-progress`（**Phase 1 Exit ✅** — adapter 已委派 Rule A/B；下一步 Phase 2 擴大遷移）  
**Owner**: framework maintainer (linyihong)  
**建立日期**: 2026-07-30  
**Canonical name**: **Knowledge Governance Engine (KGE)**  
**Folder slug**: `…-knowledge-governance-runtime`（歷史 working title；id 維持穩定，不改 folder）  
**Priority**: **P1**  
**Maturity framing**: Architecture Hypothesis **validated (Phase 0)** → Contracts + Thin Engine

> ## Success Criterion（北極星）
>
> **成功 ≠ 多一堆 hook / script，也 ≠ 又一個 lint framework。**  
> 成功 = Ai-skill 補上 Architecture Evolution 缺的一塊：
> **Knowledge becomes executable** — Knowledge Layer 經 Engine 變成可掛載的 Rule Capability，
> 再由 **Runtime Adapter** 接到 Git / CI / IDE / AI / MCP。
>
> 外部專案與本 repo 只「掛規則（plugin）」，不必各自發明 pre-commit／CI。  
> Agent 忘記 prose 時，`validation` kind 仍擋住；其餘 kinds 分級，不做 false block。

> ## 本 plan 在 Architecture Evolution 的位置
>
> | 既有研究成果 | 治理切面 |
> | --- | --- |
> | ERA | 決策治理 |
> | Delegation Loop | 執行／任務治理（Task Execution） |
> | Executable Contract | Contract 治理 |
> | Runtime / Projection（`runtime.db`） | Session／phase Runtime 治理 |
> | **KGE（本 plan）** | **知識治理（Knowledge Execution）** |
>
> KGE **不是突然冒出的新概念**；它把你們一路在做的模式  
> `Knowledge → (optional) Projection → Executable surface → Adapter`  
> 橫切套到「agent 易忘的 governance / enforcement」上。

---

## 審閱結論摘要（2026-07-30）

| 維度 | 結論 |
| --- | --- |
| 總評 | 方向正確且承接既有原則；約 **8.8–9.2/10**（未成熟，但定位對） |
| 最大優點 | Adapter≠Engine；`Validate(Context)→Findings`；Severity taxonomy；與 Evolution 同構 |
| 必須修 | (1) 避開 `Runtime` 撞名 → **Engine**；(2) Projection 非 core 必經；(3) Dependency 降級 optional；(4) 補 **Capability Registry** |
| 定位升級 | 不要縮成「Validation Engine / Lint」；真定位是 **Knowledge Layer 的可執行邊界** |
| 下一跳觀察 | KGE（Knowledge Execution）與 Delegation Loop（Task Execution）或可收斂到同一更高層模型（共享 Capability / Registry）— **探索項，非 Phase 1** |

---

## 討論脈絡（conversation capture）

來源：2026-07-30 設計討論 + 同日完整審閱。

### 核心共識（v2，審閱後）

1. **不是 Hook / Pre-commit Framework** — adapter 可增長；Engine 不必改。
2. **正式名 = Knowledge Governance Engine (KGE)** — `Runtime` 留給 `runtime/` phase machine 與 **Runtime Adapter**，避免「哪個 Runtime」。
3. **兩層 + Registry**：
   - **Engine（Ai-skill）**：Context Builder、optional Projection / Dependency、Capability Registry、Rule Dispatch、Reporter。
   - **Plugin Pack（專案）**：只宣告 rules／所需 capability；Engine 不懂 README 語意。
4. **Rule 契約**：`Validate(Context) -> Findings`；不知 Git / Commit / Hook。
5. **Pipeline（精簡核心 + 可選能力）**：
   ```text
   Input Snapshot
        │
        ▼
   Context Builder          ← 必經
        │
        ├──► Projection         ← optional（rule 宣告需要才跑）
        ├──► Dependency Provider← optional（rule 宣告需要才跑）
        │
        ▼
   Capability Registry      ← 誰可以跑／綁哪種 implementation
        │
        ▼
   Rule Dispatcher → Reporter → Adapter Exit Policy
   ```
6. **Rule kinds**：Validation / Advisory / Discovery / Evolution / Telemetry — 不全是 Fail/Pass。
7. **定位句**：Knowledge becomes executable — 不是「又一套 validator」。

---

## Problem statement

| 現象 | 根因 |
| --- | --- |
| Agent 反覆漏 linked-updates / README sync / plan checkbox / API 對齊 | Knowledge Layer 有規則；缺橫切 **Knowledge Execution** 引擎 |
| 每條新機械鎖 ≈ `hooks.go` 特規 + YAML obligation | 缺 Rule / Context / Findings / Capability 契約 |
| 外部只能抄 hook 或只用 `plans validate` | 僅 plan 子集外放；通用 knowledge execution 未抽象 |
| 全部像 Fail/Pass | 缺 severity taxonomy → 易變成「所有 rule 都 Blocking」 |
| 敘事若停在 Validation / Lint | Discovery / Evolution / Telemetry 被矮化；Architecture boundary 說不清 |
| `runtime` 詞彙過載 | phase machine、projection、bootstrap、本引擎若都叫 Runtime → 溝通崩壞 |

---

## Decision Rationale

### Problem & Why Now

1. **Behavioral debt**：governance / enforcement 增長；agent forgetfulness 是 registry 已點名的 meta-pattern。
2. **Portability pressure**：plans 外放已 dogfood；下一步應升成通用 Knowledge Execution 能力。
3. **Naming risk**：Pre-commit Toolkit 或亂叫 Runtime 都會鎖死／撞名。
4. **Evolution gap**：ERA / Delegation / Executable Contract / `runtime.db` 各治一塊；**Knowledge → executable** 的橫切層尚未獨立成邊界。

### Decision（審閱後）

#### D1 — 定名

**採用：Knowledge Governance Engine (KGE)。**

| 候選 | 結論 |
| --- | --- |
| Knowledge Governance Runtime (KGR) | Reject 作為正式名 — 與 `runtime/`、phase machine、`runtime.db` 撞名 |
| Knowledge Validation Engine | Reject 作為主稱 — 把定位鎖進 Lint；容不下 Discovery / Evolution / Telemetry |
| **Knowledge Governance Engine (KGE)** | **Accept** — Governance 對齊切面；Engine ≠ Adapter；Runtime 留給 Adapter / 既有 `runtime/` |

副敘事允許：*Incremental knowledge validation* 描述 Phase 1 子集行為，**不是**系統正式名。

#### D2 — Architecture Boundary（定位升級）

```text
Knowledge Layer          （enforcement / governance / domain docs）
        │
        ▼
Knowledge Governance Engine   ← 本 plan 的邊界（Capability + Dispatch）
        │
        ▼
Rule Capability               ← Registry：Rule → Capability → Implementation
        │
        ▼
Runtime Adapter               ← Git / CI / IDE / AI / MCP（可增長）
```

這層 **不是** Hook、不是 Validator 清單、不是 CLI、甚至不是 Git。  
它是 **Knowledge 變成可執行表面** 的邊界。

#### D3 — Framework ↔ Plugin

```text
┌─ KGE（Ai-skill）─────────────────────────────────────────┐
│  Context Builder · optional Projection · optional Dep     │
│  Capability Registry · Rule Dispatcher · Reporter         │
│  Runtime Adapters (transport only)                        │
└──────────────────────────▲───────────────────────────────┘
                           │ loads / binds
┌──────────────────────────┴───────────────────────────────┐
│  Project Plugin Pack                                       │
│  rules + required capabilities + severity kind             │
└────────────────────────────────────────────────────────────┘
```

#### D4 — Core pipeline（必經 vs optional）

**必經**

1. **Input Snapshot** — 由 Adapter 組裝（dirty / PR / buffers / …）
2. **Context Builder** — 正規化成 `Context`
3. **Capability Registry lookup** — rules → capability → implementation
4. **Rule Dispatcher** — `Validate(Context) -> Findings`
5. **Reporter + Adapter Exit Policy**

**Optional（僅當 rule / capability 宣告需要）**

- **Projection** — 例：commit-msg 只需 Context；不得強迫全走 Projection
- **Dependency Provider** — 不是 Graph Runtime；多數 rule 無 dependency。用 *Provider* 而非必經 *Resolver stage*，避免「KGE = Graph Runtime」

#### D5 — Capability Registry（審閱補上的缺口）

```text
Rule  →  Capability  →  Implementation
```

Plugin 變多時，Dispatcher 不能只靠硬編碼。須有：

| Registry 面 | 回答的問題 |
| --- | --- |
| Rule Registry | 有哪些 rule、kind、severity、啟用條件 |
| Capability Registry | 需要什麼能力（`paths+content`、`commit_msg`、`projection:plan_tree`…） |
| Plugin / Implementation binding | 誰提供 implementation；對齊 Enforcement Registry coverage / executor |

**與 Enforcement Registry（Layer 2.5）**：不另造第二套 coverage 哲學；KGE binding 應能被指到 `rule_class → capability/executor`。

#### D6 — Rule kind taxonomy

| Kind | 預設 severity | Phase 1 |
| --- | --- | --- |
| `validation` | block | Yes |
| `advisory` | warning（Finding severity） | Phase 1 taxonomy；**呈現由 Adapter Policy（D9）決定，≠ 一定刷屏 warn** |
| `discovery` | suggestion | Later |
| `evolution` | suggestion（**human promotion required**） | Later |
| `telemetry` | metrics | Later |

> **Severity ≠ UI。** Finding 的 kind/severity 由 Engine 決定；**哪個 Adapter 怎麼呈現** 由 D9 Presentation Policy 決定。

#### D7 — 與 Delegation Loop 的更高層觀察（非 Phase 1）

| | KGE | Delegation Loop |
| --- | --- | --- |
| 執行對象 | Knowledge Execution | Task Execution |
| 典型輸入 | Knowledge / dirty context | Intent / brief |
| 共享候選 | Capability、Registry、Adapter、Findings 形狀 | （探索） |

**Phase 0–3 不合併實作**。記為 architecture follow-up，避免兩套各自長成 orchestration 後再硬併。

#### D8 — Portable copy unit（外部可不引用 Ai-skill）

**Accept（2026-07-30）**：驗證引擎與規則實作放在

`scripts/ai-skill-cli/portable/kge/`

| 原則 | 說明 |
| --- | --- |
| 複製單元 | **整個目錄**（多檔：contracts / engine / rule_* / tests / README），不壓成單一巨大 `.go` |
| 零宿主依賴 | 不 import Ai-skill app、不讀 `runtime.db`、Rule 不呼叫 git |
| 外部用法 | 不引用本 repo 也可：拷目錄進自己的 Go module → 自寫 Adapter 組 `Context` → `Engine.Run` |
| 本 repo 用法 | 只在 `internal/app`（等）寫 Adapter／hook／CI **串流程**；不把流程鎖進 portable 包 |

文件：[`scripts/ai-skill-cli/portable/kge/README.md`](../../../scripts/ai-skill-cli/portable/kge/README.md)

#### D9 — Adapter Presentation Policy（Q6 拍板，2026-07-30）

**Reject 舊問法**：「commit-msg 要不要 warn / silent？」— 那是 hook-framework 思維。

**Accept 新問法**：**同一個 Advisory，在哪個 Adapter 應如何呈現？**

```text
Severity (Engine Finding)
        ×
Adapter
        =
Presentation Policy
```

##### Checkpoint vs watershed

| 動作 | 語意 | KGE 角色 |
| --- | --- | --- |
| **Commit** | Save checkpoint（存一下） | **Validation only**（block on fail）；Advisory **不展開**，最多 count + pointer |
| **Push / `kge check`** | 準備分享出去 | Validation（block）+ Advisory **摘要**（不擋 push） |
| **IDE** | 連續書寫 | Advisory / Discovery **即時完整**（最佳體驗） |
| **`kge validate --advisory`** | 使用者主動來看 | Advisory / Discovery **完整列表** |
| **CI** | 整合審查 | Validation fail + **Full advisory report**（給 review） |

##### Presentation matrix（canonical）

| Adapter | Validation | Advisory | Discovery |
| --- | --- | --- | --- |
| **IDE / MCP diagnostic** | 即時紅線（可選） | ✅ 即時完整提醒 | ✅ Suggestion |
| **`kge validate`（預設）** | ❌ Fail on error | （預設可不跑） | — |
| **`kge validate --advisory`** | ❌ Fail on error | ✅ **完整輸出** | ✅ 完整（若啟用） |
| **`kge check`**（新 adapter） | ❌ Fail → block | ⚠️ **摘要**（例前 3 條）+ count；**不擋** | Optional 一行 |
| **pre-commit** | ❌ Block | ⚠️ 極簡 summary 或省略 | ❌ |
| **commit-msg** | ❌ Block | ⚠️ **≤3 行**：count + `kge validate --advisory`；**不展開內文** | ❌ |
| **pre-push**（建議掛 `kge check`） | ❌ Block | ⚠️ Summary（同 check）；**不擋** | ❌ / Optional |
| **CI** | ❌ Fail | ⚠️ Full report artifact | Optional |
| **AI review** | ❌ Block（可選） | ✅ 詳細 | ✅ 詳細 |

##### commit-msg 範例（上限）

```text
✔ commit accepted

KGE: 2 advisory findings (document sizing, README sync).
Run: kge validate --advisory
```

- **不要**在 commit 洗十幾條 Warning（Commit ≠ Review）。
- **也不要**完全 silent（否則建議在整合前被忘記）。
- **效能**：commit-msg **不必跑完全部 Advisory rules**；Validation 過後可只算 **Advisory count**（或快取／輕量計數），細節留給 `validate --advisory` / IDE。

##### `kge check`（建議新增 CLI adapter）

```text
kge check  →  Validation + Advisory summary (+ optional discovery line)
```

- Validation fail → non-zero（可被 pre-push 用來 block）
- 僅 Advisory → zero exit + “Ready to push / N recommendations”
- git push 前可自動跑；**Advisory 永不單獨擋 push**（除非專案 profile 顯式升級，本 plan 預設不升級）

### Alternatives Considered

| ID | 方案 | 結論 |
| --- | --- | --- |
| A | 繼續只加 hook 特規 | Reject 長期 |
| B | Pre-commit framework | Reject 命名/抽象 |
| C | 正式名 KGR（Runtime） | Reject — 撞名（審閱後） |
| D | 主稱 Validation Engine | Reject — 定位過窄（審閱後） |
| E | **KGE + optional Projection/Dep + Capability Registry** | **Accept** |
| F | 強制全 pipeline 含 Graph Projection | Reject — 易變 orchestration |
| G | Engine 只活在 `internal/`、外部必須引用整庫 | Reject — 違反「可複製組裝驗證器」 |
| H | 壓成單一 `.go` 方便 copy | Reject — 難增刪規則；改用目錄多檔 |
| I | **`portable/kge` 目錄為 copy unit + 宿主 Adapter** | **Accept（D8）** |
| J | commit-msg 展開全部 Advisory warn | Reject — Commit≠Review；噪音 |
| K | commit-msg 對 Advisory 完全 silent | Reject — push 前建議易被忽略 |
| L | **Severity×Adapter Presentation Policy + commit summary / push check** | **Accept（D9 / Q6）** |

### Why Not an ADR Yet

- Capability Registry 與 Enforcement Registry 精確 schema 尚未 spike。
- 與 Delegation 更高層模型仍是觀察，不是決策。
- Plugin 格式 / manifest 路徑未決。

### ADR Promotion Criteria

- [ ] KGE 名詞穩定，且與 `runtime/` phase machine **口語可分辨**
- [ ] Capability Registry 最小 schema + ≥2 rules（一需／一不需 projection）通過
- [ ] ≥1 非-plan domain plugin 在本 repo 以 plugin 形式 enforce
- [ ] ≥1 外部專案 plugin pack + thin adapter，可逆移除
- [ ] Enforcement Registry 能表達 `rule_class → KGE capability`，無第二套 coverage

### Consequences

**正向**：Adapter 增長不改 Engine；severity 防全域 Blocking；對齊 Evolution；撞名風險下降。

**風險**：Engine/Registry 抽太早 → orchestration；Discovery/Evolution 過早 → process bloat；與 Delegation 過早合併 → 範圍爆炸。

**Mitigations**：Phase 0 inventory；Projection/Dep 預設 off；Evolution 不自動寫回（Gen 4）；Delegation 收斂另開探索。

---

## Target architecture（draft v2）

### Framework surface

```text
# Portable copy unit（外部整夾複製；Ai-skill 與外專案共用同一包）
scripts/ai-skill-cli/portable/kge/
  README.md          # copy boundary
  contracts.go       # Context, Finding, Rule, Severity, Capability
  engine.go          # dispatch + capability_missing
  rule_*.go          # pure Validate — 可增刪單檔
  engine_test.go

# Ai-skill-only wiring（不隨 copy 帶走）
scripts/ai-skill-cli/internal/app/
  … adapters: 組 Context、hook/CI/CLI 委派 …

# Candidate A — repo 第一級 validation/ 目錄仍延後
```

### Project plugin surface

```text
<PROJECT_ROOT>/
  <manifest TBD>/
    manifest.yaml
    rules/
      readme-sync.rule.yaml      # may require: paths+content
      plan-tree.rule.yaml        # may require: projection:plan_tree
      commit-msg.rule.yaml       # commit_msg only — no projection
```

### Adapter map + Presentation Policy

> 細節見 **D9**。此表為速查。

| Adapter | InputSnapshot | Validation | Advisory presentation |
| --- | --- | --- | --- |
| IDE / MCP | open/save buffers | 即時紅線（可選） | ✅ 完整即時 |
| `kge validate` | `--root` + paths | Fail on error | 預設不跑；`--advisory` 完整 |
| `kge check` | dirty / staged / range | Fail → block-capable | ⚠️ 摘要（前 N 條）；不擋 |
| pre-commit | staged | Block | 極簡／可省 |
| commit-msg | msg + staged | Block | ⚠️ ≤3 行 count + pointer |
| pre-push | 建議呼叫 `kge check` | Block on validation | ⚠️ Summary；不擋 |
| CI | PR / tree | Fail | Full report |
| AI review | dirty + hints | Block 可選 | ✅ 詳細 |

### Non-goals

- 不取代 `runtime/` phase machine / bootstrap / cognitive modes。
- 不強制全部 `behavioral_only` 機械化。
- Phase 1 不做 Evolution 自動寫回；不做 KGE↔Delegation 合併。
- 不把 KGE 建成 Graph Runtime 或「凡事必 Projection」。
- 不要求外部安裝 `runtime.db` 或全套 governance。
- **不**把 portable engine 壓成單一巨大 `.go`（複製單元 = 目錄；規則可增刪單檔）。
- **不**要求外部專案引用整個 Ai-skill repo 才能用 engine（可 copy `portable/kge`）。

---

## Relationship to existing systems

| 既有系統 | 關係 |
| --- | --- |
| Enforcement Registry | Capability / Implementation binding 的 coverage 面；**KGE kind 必須尊重 6-bucket，不另造哲學** |
| Commit-msg / README / Linked-updates validators | **Phase 0–1 首批**；Phase 2a 繼續遷既有 **mechanical** |
| Enforcement `behavioral_only`（如 document_sizing） | **Phase 2b / 5**：優先做 **advisory（提醒）**，非一律 block；對齊各 class 的 `sunset_decision` |
| Enforcement `not_mechanizable`（如 neutral_language） | **預設不進 KGE block**；最多極窄 heuristic discovery（opt-in），或維持 agent/review 行為 |
| Plans validation engine | **Deferred** — Later pack |
| `per_commit_obligations` | Rule ids 映射；commit-msg = Adapter |
| Delegation Loop | Task Execution 對偶；共享抽象 = **follow-up** |
| Cognitive Execution System | Knowledge Execution 橫切 |

---

## Enforcement → KGE 對照（規劃補齊 — 2026-07-30）

> **使用者提問**：很多 `enforcement/` 現在「沒效果」（behavioral）；像 document-sizing、neutral-language 要不要進 KGE？有的強制、有的提醒？  
> **盤點結論**：Plan **已有** severity taxonomy（D6）與「不全機械化」Non-goal，但 **Phase 2 清單寫得太窄**（只遷既有 commit-msg validators），**沒有**把 enforcement coverage 桶對到 KGE kind。本節補上。

### 現況數字（`ai-skill enforcement coverage`，2026-07-30）

| Coverage | 約略 | 對 KGE 的預設處置 |
| --- | --- | --- |
| `mechanical`（21） | 已有 executor | Phase 2a：遷入 / 對齊 KGE `validation` 或既有 warning |
| `behavioral_only`（13） | prose + sunset | Phase 2b／5：挑 **可測子集** → 多半 `advisory`；達 sunset 才升 `validation` |
| `not_mechanizable`（5） | 永久主觀 | **不**做 block；避免假安全感 |
| `research_required`（2） | 路徑未清 | 不進 KGE 實作佇列 |

### Kind × Adapter（強制 vs 提醒 — 已由 D9 取代「一律 warn」）

| KGE kind | Engine severity | 呈現 |
| --- | --- | --- |
| `validation` | error | 各 Adapter：**可 block/fail**（見 D9 矩陣） |
| `advisory` | warning | **依 Adapter**：IDE 完整；commit-msg 僅 count；push/`check` 摘要；`--advisory`/CI 完整 |
| `discovery` | info | IDE / `--advisory` / AI review；commit-msg **不跑** |
| （不掛 KGE） | — | agent prose / review |

**Q6：✅ Resolved（D9）** — 不再問「commit-msg warn vs silent」；改為 Presentation Policy。commit-msg = **count + pointer**（非 silent、非洗版）。

### 具體 class 處置（首批對照）

| Rule class | Registry coverage | 進 KGE？ | 建議 kind | 備註 |
| --- | --- | --- | --- | --- |
| 既有 commit-msg / pre-commit mechanical | mechanical | **Yes（2a）** | `validation`（少數既有 warning 維持 warn） | 已開始（cognitive_cost、cli_doc_sync） |
| `linked_updates` | mechanical（部分）+ 大量 behavioral 表 | **子集 Yes** | 可機械 proxy=`validation`；全文圖=`advisory` 或仍 behavioral | 勿宣稱「linked-updates.md 全機械化」 |
| `document_sizing` | **behavioral_only** | **Yes 候選（2b）** | **`advisory` 先** | sunset：deterministic split-suggestion lint；行數超標可 warn，**拆哪一節仍是 judgment → 不 block** |
| `dependency_reading` | behavioral_only | 晚 | advisory／discovery | 需 transcript read-log；bootstrap 已有機械子集勿重複發明 |
| `neutral_language` | **not_mechanizable** | **預設 No block** | 最多 opt-in discovery heuristic | Registry 已寫：字面 lint 會被 gaming；**強制進 KGE = 違反 not_mechanizable** |
| `tool_neutral_documentation` | not_mechanizable | No block | — | 同左 |
| `rule_weight` / `decision_efficiency` 等 | behavioral / not_mech | No／極晚 | — | 衝突判斷與效率偏好不適合 Fail/Pass |

### 與 Non-goals 的一致性

- 「不強制全部 `behavioral_only` 機械化」仍然成立。  
- KGE **不是**把 13 條 behavioral 一次改成 block。  
- 正確用法：對 **可客觀測量的切片** 掛 **提醒（advisory）**；達 registry `success_criteria` 再升 validation。

### Phase 切分修正

| Phase | 內容（更新） |
| --- | --- |
| **2a** | 遷更多 **已是 mechanical** 的 validators → `portable/kge`（markdown_yaml_sync、bootstrap thinness…） |
| **2b** | Enforcement advisory pack：**document_sizing**（Finding=`advisory`）；**呈現走 D9**（IDE 完整、commit 僅 count） |
| **4** | Multi-adapter：**IDE diagnostic** + **`kge check`** + pre-push 掛 check；CI full report |
| **5** | 擴大 advisory／discovery；**落實 D9 矩陣**（含 commit count-only 效能路徑）；不強制 not_mechanizable |

---

## Phase 0 — 驗證 Architecture Hypothesis（進行中）

> **Phase 0 要回答的不是「有哪些 validator？」**  
> 而是：**「這個 Architecture 能不能裝得下目前的 validator？」**  
> Inventory 是手段；Architecture Mapping + Mini Spike 才是目的。避免 Paper Design。

### Phase 0 三軌

```text
Phase 0
├── Inventory           （既有 validator → Context / Projection / Capability 映射表）
├── Architecture Mapping（用五問驗證抽象邊界）
└── Mini Spike          （Rule A 無 Projection + Rule B 可選 co-change／非必經 Projection）
```

### 五問（Inventory / Mapping 必須回答）

| # | 問題 | Phase 0 要證明 / 否證 |
| --- | --- | --- |
| A | **Context 是否足夠？** | 有沒有 rule 一直逼 Context Builder 加特例欄位？ |
| B | **Projection 是否真的 Optional？** | 有多少 rule 根本不用 Projection？是否少數？ |
| C | **Capability 是否能收斂？** | Capability 數量是否失控，還是收成有限穩定集合？ |
| D | **Registry 是否合理？** | 多條 rule 能否共享同一 capability？ |
| E | **Adapter 是否真的透明？** | 同一 rule 邏輯能否概念上同時掛 CLI 與 Hook、不改 Rule？ |

### 首批 mapping 目標（Resolved — 暫不碰 Plans）

| 優先 | 候選 | 為什麼適合 Phase 0 |
| --- | --- | --- |
| 1 | **Commit Message** 類 validators | Context 幾乎只要 msg；驗證「無 Projection」路徑 |
| 2 | **README Sync**（或目錄索引同步類） | paths + content；Projection 可有可無 |
| 3 | **Linked Updates**（子集 / advisory 或 block 行為） | forgetfulness 高；Dependency 可能按需出現 |
| later | Plans validation engine | 成熟但依賴最多 — **Phase 0 不做真搬遷** |

### Mini Spike（提前納入 Phase 0）— ✅ 完成

| Rule | Projection | 實際選型 | 要觀察 |
| --- | --- | --- | --- |
| **Rule A** | 不用 | `cognitive_cost` | Context → Capability → Dispatch 是否自然 |
| **Rule B** | **不用**（用 optional `path_cochange`） | `cli_doc_sync` | Adapter 注入 diff 後 Rule 是否仍透明；**不**強迫 Projection |

交付物：[`evidence/p0-architecture-mapping.md`](evidence/p0-architecture-mapping.md)。紙上 walkthrough ✅；Phase 1 做薄實作。

### Phase 0 Exit（進 Phase 1 的門檻）

全部傾向「成立」才進 Phase 1（Contracts + Thin Engine）：

- [x] Context Builder **沒有**在 mapping 中一直長特例（見 evidence §3A — 特例風險在 Adapter 注入 diff，非無限 Context 欄位）
- [x] Projection **確實**只有少數 rule 需要（焦點三域 ≈ 0；Plans/runtime Later）
- [x] Capability Registry 能收斂成 **有限且穩定** 的一組能力（~7-cap 工作集）
- [x] 同一條 Rule 可經不同 Adapter（CLI、Hook）執行的路徑在概念上成立、**不必改 Rule**（條件：Rule B 剝離內嵌 `git`）
- [x] Mini Spike（A+B）顯示 Dispatcher / Capability / Context 自然（紙上 walkthrough 完成）
- [x] 確認 Plans **不**作為本階段 migration target

**Evidence**：[`evidence/README.md`](evidence/README.md) · [`evidence/p0-architecture-mapping.md`](evidence/p0-architecture-mapping.md)

**簽核狀態**：✅ **使用者已簽核 Phase 0 Exit（2026-07-30）** → 進入 Phase 1。

若五問或 spike 否證架構 → **修 Architecture Hypothesis**（修 plan），不硬進 Phase 1 堆 contracts。

---

## Phase 1 — Contracts + Thin Engine（進行中）

### 目標

1. 落地最小契約：`Context` / `Finding` / `Rule` / `Severity` / `Capability`
2. Thin engine：Context 組裝（無 Git）→ Capability lookup → Dispatch → Findings
3. 掛 **Rule A**（`cognitive_cost`）與 **Rule B**（`cli_doc_sync`，diff 由 Adapter 注入）
4. Hook 行為相容（可先雙路徑：舊函式 delegate 到 KGE，或 adapter 包裝）
5. CLI 能對同一 Context 跑同一 Rule（證明 Adapter 透明）

### Placement

**Portable copy unit（外部可整夾複製）**：

`scripts/ai-skill-cli/portable/kge/`

- 多檔模組：`contracts` / `engine` / `rule_*` / tests + `README.md`
- **零** Ai-skill / `runtime.db` / git 依賴
- 宿主專案（含本 repo）只寫 Adapter 組 `Context`、解釋 `Findings`

**Ai-skill wiring（不可隨 engine 複製的流程）**：`scripts/ai-skill-cli/internal/app/`（未來 hook／CLI 委派）

### Phase 1 Exit

- [x] portable `kge` contracts + engine 可編譯／有單測（`go test ./portable/kge/` 綠）
- [x] **Copy boundary 文件化**（`portable/kge/README.md`）— 外部可整夾複製、本 repo 只串 Adapter
- [x] Rule A 經 KGE 路徑與既有行為等價（`validateCognitiveCost` → `runKGECognitiveCost`；app tests 綠）
- [x] Rule B **不**在 Rule 內呼叫 `git`；缺 `staged_diff` → `capability_missing`（單測）；hook 路徑由 adapter 注入 diff
- [x] 至少一個測試入口不經 commit-msg hook 呼叫同一 Rule（package + adapter tests）
- [x] 本 repo hook 無行為倒退 — `go test ./internal/app/` 綠（雙路徑委派）
- [x] Q2 工作假設：Phase 1 = **first-party Go plugins only**；外部 = **copy `portable/kge` pack**（非強制 go get）

**Phase 1 技術 Exit：✅**（尚未擴大到全 obligation 遷移 — 屬 Phase 2）

---

## Phased roadmap

| Phase | 目標 | Exit |
| --- | --- | --- |
| **0 — Architecture Hypothesis validation** | Inventory + Mapping（五問）+ Mini Spike；首批三域；不搬 Plans | 見上方 Exit checklist |
| **1 — Contracts + thin engine** | Context/Finding/Rule/Severity/Capability；實作薄 engine；掛首批 mapping 通過的 rules | hook 行為相容；CLI 同路徑 |
| **2 — Migrate Ai-skill pack** | **2a** 既有 mechanical → KGE；**2b** document_sizing 等 advisory 原型 | coverage 不倒退；advisory 不誤升 block |
| **3 — External plugin dogfood** | 外部少量 rules + thin adapter；可逆 | portability success 子集 |
| **4 — Multi-adapter** | IDE + **`kge check`** + pre-push + CI report | D9 矩陣可觀測；advisory 不誤擋 push |
| **5 — Advisory / Discovery** | 落實 D9 呈現；擴大提醒類；**不**強制 not_mechanizable | commit ≤3 行；`--advisory` 完整 |
| **Later — Plans pack** | 將 plans validation 收斂為 KGE domain pack | 首批三域 + Phase 1–2 穩定後 |
| **Explore** | KGE + Delegation 統一 Execution Model | **獨立 spike/plan** |

---

## Runtime Execution Path（草案宣告）

> 章節名依 plan 模板；**不是**把 KGE 改回叫 Runtime。  
> Phase 0：**不接入** executable runtime wire；只做 mapping + mini spike evidence。  
> Phase 1 起才談 contracts / thin engine consumer。

| 欄位 | 草案 |
| --- | --- |
| Engine owner | `scripts/ai-skill-cli`（KGE host） |
| Trigger flow | Adapter → InputSnapshot → Context Builder → (opt Projection/Dep) → Capability Registry → Dispatch → Findings → Adapter exit |
| Activation | manifest + rules/capabilities；本 repo pack ↔ obligations（Phase ≥1） |
| Generated surface | Phase ≥2 再宣告；Phase 0–1 可不 project |
| Validation scenarios | Phase 0：五問 + spike evidence；Phase 1：hook 相容 + optional-stage 單測 |
| Consumer | Phase 0：無強制 consumer；Phase 1：CLI / git adapter |

Graduation：Phase 0 Exit 通過後進 Phase 1；若長期卡住則修 hypothesis 或 archive / orphan。

---

## Architecture Compatibility Preflight

| Check | 結果 |
| --- | --- |
| Source-of-truth | owner-layer 保留規則正文 |
| `runtime/` 邊界 | phase machine 與 KGE 口語／owner 分離 |
| Optional stages | Projection / Dependency 非必經 |
| Capability Registry | 對齊 Enforcement Registry，不平行 coverage |
| Plans portability | shared-binary + thin adapter + monotonic removal |
| Over-engineering | `portable/kge` 可 copy；第一級目錄延後；Delegation 合併延後 |

---

## Open Questions

| # | 問題 | 狀態 |
| --- | --- | --- |
| Q1 | 正式名？ | **Resolved: KGE** |
| Q2 | Plugin 格式：純 YAML vs YAML+Go？Phase 1 first-party Go only？ | **Resolved (working)**: Phase 1 = first-party Go rules in `portable/kge`；外部 = copy pack（非強制 go get）；YAML plugin 格式延後 |
| Q3 | Candidate B → 第一級目錄門檻？ | open |
| Q4 | plans engine：併入 KGE pack vs subdomain 呼叫？ | **Deferred** — Phase 0 不搬 Plans；Q4 延到 Later—Plans pack |
| Q5 | 第一個非-plan pack 優先序？ | **Resolved: Commit Message → README Sync → Linked Updates**（Plans later） |
| Q6 | Advisory 呈現？→ **改問：Severity×Adapter Presentation？** | **Resolved: D9** — commit=count+pointer；push/`kge check`=summary 不擋；IDE/`--advisory`/CI=完整 |
| Q7 | 外部 manifest 路徑 / branding？ | open |
| Q8 | Discovery/Evolution 永需 human promotion？ | **lean Yes** |
| Q9 | Capability Registry 最小 schema？ | **Phase 0 產出**（對五問 C/D） |
| Q10 | 與 Delegation 統一模型是否另開 plan？ | open（建議 Yes） |

---

## Acceptance

### 進入 Phase 0（已滿足）

- [x] Architecture Boundary 穩定（Knowledge → Engine → Capability → Adapter）
- [x] Core Pipeline 收斂（Projection / Dependency optional）
- [x] Capability Registry 已入架構
- [x] 定名 KGE；定位升為 Architecture Hypothesis
- [x] Phase 0 目標改為驗證架構（含五問 + Mini Spike）；首批三域確定；Plans 暫緩

### 進入 Phase 1（✅ 已簽核）

- [x] 五問 A–E 傾向成立（p0-map §3；**已簽核**）
- [x] Mini Spike A+B evidence 完成（A=`cognitive_cost`；B=`cli_doc_sync` — **共變而非 Projection**）
- [x] Capability 候選集合草案寫出（~7 caps，見 p0-map §1）
- [ ] Q2 至少有工作假設（可改）— Phase 1 進行中補
- [x] **使用者簽核** Phase 0 Exit → 授權 Phase 1

---

## Document TODO

| ID | Item | Status |
| --- | --- | --- |
| T1 | 使用者審閱 Decision / Naming / Non-goals | done |
| T2 | Phase 0：Architecture Mapping（五問） | **done + 簽核** |
| T3 | Phase 0：Mini Spike A/B | **done + 簽核** |
| T4 | Glossary：KGE、Capability Registry、Runtime Adapter（vs `runtime/`） | pending |
| T5 | 交叉引用 Delegation Loop | pending |
| T6 | Unified Knowledge/Task Execution Model 探索 plan | pending |
| T7 | Phase 0 evidence/ | done |
| T8 | 可選：`p0-capability-schema-draft.md` | deferred（caps 已在 p0-map §1） |
| T9 | 使用者簽核 Phase 0 Exit | **done** |
| T10 | Phase 1：portable `kge` + Rule A/B + adapter 委派 | **done**（Exit ✅；擴大遷移 → Phase 2） |
| T11 | Portable copy README / 邊界 | **done** |
| T12 | Phase 2a：更多 mechanical validators → KGE | **partial** — +plan_checkbox_sync, plan_archival_audit |
| T13 | Phase 2b：document_sizing advisory Finding + D9 呈現 | **done**（rule + `kge check`/`validate --advisory`） |
| T14 | 明確不把 neutral_language 做成 KGE block（對齊 not_mechanizable） | **done** |
| T15 | Q6 / Adapter Presentation Policy | **done（D9）** |
| T16 | 實作 `kge check` + commit-msg advisory count path + pre-push 掛載 | **done** — CLI check/validate/diagnose；commit-msg count；pre-push `kge check` |

---

## Next Action

1. Phase 2a：繼續遷 mechanical validators。
2. CI full advisory report（D9 CI 列）。
3. Plans pack 仍 Later。
4. IDE host 接線（Cursor/VS Code task 呼叫 `kge diagnose`）可另開。
