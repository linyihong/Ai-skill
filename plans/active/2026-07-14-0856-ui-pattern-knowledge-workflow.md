---
id: 2026-07-14-0856-ui-pattern-knowledge-workflow
plan_kind: main
status: in-progress
owner: linyihong
created: 2026-07-14
last_updated: 2026-07-14
parent: null
priority: P2
revision:
  - date: 2026-07-14
    note: "Stakeholder review：從 Vocabulary 升格為 UI Pattern Knowledge"
  - date: 2026-07-14
    note: "Rename file+id → ui-pattern-knowledge-workflow"
  - date: 2026-07-14
    note: "Q7–Q11 決議；Core/Extended；禁 Intent DB；transient checklist；ARIA→headless→project；evidence_level"
  - date: 2026-07-14
    note: "Sanitization：dogfood 改 <PROJECT_ROOT>"
  - date: 2026-07-14
    note: "Round-2 Verifier：澄清 D9 trigger（T1–T5 任一成立；T1∨T2 強調句與 ≥2 consumers 不再語意打架）"
  - date: 2026-07-14
    note: "Phase 0 complete → status in-progress；D9 T1 時距自今日起算；Phase 1 未開始"
  - date: 2026-07-14
    note: "Phase 1 landed"
  - date: 2026-07-14
    note: "Phase 1 Verifier 仲裁：fix artifact-gates 全鏈、seed project:null、Lock when_to_load≠runtime；defer Contract Stack 圖"
  - date: 2026-07-14
    note: "Stakeholder Phase1 freeze：Done Definition 鎖產物；pattern-index（非 runtime-index）；validation/entry-schema；Phase2 僅五件 overlay 驗形狀不堆名詞"
  - date: 2026-07-14
    note: "Phase 2 kickoff：目標=Pattern Family 可推理性；成功=5 entries + 10 Selection Scenarios；順序 Scrim→Dialog→Sheet→Drawer→Toast；停止 Phase1 打磨"
  - date: 2026-07-14
    note: "Phase 2 Completed（H1 Selection / H2 Family / H3 Near Neighbor）；closure phase2-summary；Phase 3=Pattern Composition（Episode Page）；entry freeze→composition_rules"
  - date: 2026-07-14
    note: "Research ladder：P1 Representability→P2 Inferability→P3 Composability；Phase3 凍結 H4 Independence / H5 Completeness / H6 Traceability；composition_rules=Constraint not Rule Library"
  - date: 2026-07-14
    note: "Phase 3 formal start：Entry Freeze→Invariant（anti back-propagation）；H4–H6 分型证据；exit=Composition Closure（非完成 Screen）；P4 Orchestrability 僅觀察不入 plan"
  - date: 2026-07-14
    note: "Phase 3 review=constraints；H4→H5→H6 mini-cycles；Composition Metrics；Layer Growth Rhythm → Architecture Evolution Protocol"
---

# UI Pattern Knowledge — Workflow 強化計畫

**Status**: `in-progress` — **Phase 0–2 complete**；**Phase 3 started**（Pattern Composition · Episode Page）  

**Owner**: linyihong  
**建立日期**: 2026-07-14  
**Source**: 2026-07-14 對話 — NameThatUI 對照 + stakeholder 回饋（Knowledge layer，非 glossary；非 NameThatUI clone）。  
**Glossary Impact**: yes — candidates（未註冊）：`ui_pattern_knowledge`、`pattern_selection_rules`、`pattern_composition`、`pattern_family`、`implementation_recipe`、`pattern_prompt_expansion`、`pattern_knowledge_core|extended`、`pattern_evidence_level`（verified / observed / …）。舊詞 `ui_pattern_vocabulary` = 子能力。graduate 後才註冊。
**D9 clock**: `in-progress` since **2026-07-14**（T1 三月時距起算點）

> **Watch-Out List citation**（[`architecture/ai-native-cognitive-ecosystem-system.md`](../../architecture/ai-native-cognitive-ecosystem-system.md) §Watch-Out List）：
> - **process bloat** — 不新建 lifecycle；不建 Intent Ontology / Design-System 百科；Prompt checklist **不**每次 commit。
> - **premature abstraction** — platform_map 只留三層；Material/iOS/Fluent 若需要另開 `platform-adapter`，不混進 Pattern Knowledge。
> - **autonomous optimizer** — Expansion = 從 Knowledge 展開 transient checklist，不是自動改設計。

## Executive summary

定位：**可被 workflow 消費的 UI Pattern Knowledge Layer**（載入 → 推理選型 → 展開 → 驗證），不是人查 glossary，也不是模仿 NameThatUI。

### Research ladder（驗證單位升級 — 耐久）

質變點不在「多了 Pattern Composition 這個名詞」，而在 **unit of validation** 升級：

| Phase | 驗證對象 | 驗證單位 | 核心能力 | 核心問題 |
| --- | --- | --- | --- | --- |
| Phase 1 | Schema | Entry | **Representability**（可表示） | Knowledge 能不能被表示？ |
| Phase 2 | Inferability | Scenario | **Inferability**（可推理） | Knowledge 能不能被推理？ |
| Phase 3 | Composition | Screen | **Composability**（可組合） | 多個 Knowledge 能不能共同工作？ |

```text
Phase 3 chain（驗證單位 = Screen）
Screen → Pattern Tree → Selection → Recipe
```

Entry / Scenario 仍是下層產物；Phase 3 **不以 Entry 數量、也不以新 Screen 堆疊**衡量進度。  
（跡象：Entry → Composition Constraints → Pattern Tree 開始像 **Knowledge Graph** 的邊，而非「很多 YAML」——可另開第四條研究線，本 plan 暫不升格。）

### Authority Boundary

| Owner | Owns（一句） |
| --- | --- |
| **Pattern Knowledge** | *What pattern is appropriate*（選哪個、為何、近鄰互斥） |
| **Design Contract** | *How the chosen pattern is built*（token / primitive / 視覺權威） |
| **UI Governance** | *Whether implementation complies*（合規分類與收口） |
| **Evidence** | *Whether claims are verified*（證明，非選型） |

看到 Bottom Sheet 時：選型進 Pattern Knowledge；**不**先去改 token，**不**先改 governance domain 定義。

```text
Product Requirement
  → Intent（任務側推導；不進 Intent DB）
  → Pattern Selection（Core）
  → Name / Family / Neighbors（Core）
  → Pattern Composition（獨立 artifact）
  → Design Contract（token / primitive）
  → Implementation Recipe（Extended；可 unknown）
  → UI Governance + Evidence
```

| Layer | 我們怎麼做 |
| --- | --- |
| Intent | 僅 `intent_examples` 掛在 entry；**禁止**獨立 Intent DB |
| Selection + Vocabulary + Relationship | **Core 必填** |
| Composition | 獨立 artifact |
| Recipe / Anti-pattern / Platform | **Extended**；recipe 用 maturity status |
| Prompt expansion | Knowledge（canonical）→ transient checklist（任務報告 / plan evidence，**非**每次入主知識庫） |
| Maturity | `evidence_level`（與 Evidence Governance 對齊） |

---

## Decision Rationale

### Problem & Why Now

缺的不只是名字，而是 **可推理、可展開、帶成熟度的知識**。否則 agent 停在「叫 Bottom Sheet」，仍選錯 destructive confirm、漏 focus trap、也不知道哪條規則只是 observed。

### Decision

#### D1 — UI Pattern Knowledge（非 Vocabulary 產品）

```text
UI Contracts → Pattern Knowledge Lock → Design Contract → Implement → UI Governance
```

- Owner：`workflow/software-delivery/`  
- 規則種子：Ai-skill；採用名 / composition / project alias：consumer 專案  
- **掛點（Q1 已決議）**：Pattern Knowledge Lock 寫入 [`ui-contracts.md`](../../workflow/software-delivery/ui-contracts.md)。Screen Mapping 仍不承載 framework pattern 細節；Lock 是 **選型閘門 + template 指標**，與 Screen Mapping 章節分開（新增小節，不塞進 mapping 表）。  
- 不開六道 slice / 不開 Intent repository

#### D2 — Schema = Core + Extended（Q8 決議）

**Core（缺一不可算完整 entry）**

| 欄位 | 作用 |
| --- | --- |
| `canonical_pattern` | 正式名 |
| `plain_name` | 人話 |
| `intent_examples` | 適用意圖示例（非產品 Intent DB） |
| `selection_rules`（含 when / not_when） | 何時選／不選 |
| `family` | 家族（如 overlay） |
| `near_neighbors` | 易混 + 互斥一句 |

**Extended（可後補）**

| 欄位 | 作用 |
| --- | --- |
| `implementation_recipe` | 實作能力清單；`status: unknown \| partial \| complete` |
| `platform_map` | 僅三層（見 D6） |
| `anti_patterns` | 反模式 |
| `project_component` / project alias | 專案指針 |

```yaml
id: bottom_sheet
canonical_pattern: bottom_sheet
plain_name: "從底部抽出的臨時面板，通常帶 scrim"
family: overlay
variant_of: overlay
near_neighbors:
  - id: modal_dialog
    rule: "中央阻斷／破壞性確認用 dialog；底部多選用 sheet"
  - id: drawer
    rule: "持久導航用 drawer/rail；臨時上下文用 sheet"
intent_examples:
  - choose_one_of_many_actions
  - share_or_export
  - contextual_options
selection_rules:
  when: [multiple_actions, temporary, mobile_or_h5, contextual]
  not_when: [destructive_confirm_primary, long_multi_step_workflow, persistent_nav]
  evidence_level: verified | observed | hypothesized
anti_patterns:
  - more_than_10_actions
  - full_form_or_wizard_inside
  - nested_bottom_sheets
platform_map:
  aria: dialog
  headless: Dialog
  project: ProjectShareSheet
implementation_recipe:
  status: unknown | partial | complete
  evidence_level: verified | observed | hypothesized
  required: [portal, scrim, focus_trap, escape_close, body_scroll_lock]
  recommended: [swipe_dismiss, safe_area_inset, enter_exit_motion]
entry_status: declared | legacy_alias | deferred
```

**完整 entry 定義** = Core 齊；Extended 可 `unknown` / 缺省。  
**禁止**：只有名字、沒有 selection / neighbors / family。

#### D3 — Composition 獨立 artifact

```yaml
screen: content_detail
contains:
  - pattern: app_bar
  - pattern: media_player
  - pattern: bottom_sheet
    optional: true
  - pattern: modal_dialog
    optional: true
  - pattern: toast
```

Mapping = 行為追溯；Composition = pattern 結構；二者不互代。

#### D4 — Prompt Expansion：Knowledge canonical，Checklist transient（Q9）

```text
Knowledge (repo)  →  Prompt Expansion (runtime)  →  Execution Checklist (transient)
```

| 產物 | 進哪 | 原因 |
| --- | --- | --- |
| Pattern entry / composition / recipe 知識 | **Repo**（Ai-skill 規則 + 專案 overlay） | canonical Knowledge |
| 某次任務展開的 checklist | 任務報告；可選 plan `evidence/` **摘要** | Runtime，非 Knowledge；**禁止**每次 commit 進主知識庫 |

Expansion 展開時應標出各條的 `evidence_level`（verified vs observed），避免把 observed 當硬約束。

#### D5 — 閘門強度

| 階 | 行為 |
| --- | --- |
| L0 | template + Core schema + seed |
| L1 | pattern/overlay **選型／對齊** claim 須 Core lock（或 `deferred`）；不要求 recipe complete；不挡純資料修補 |
| L1.5 | 改 screen 結構 → 更新 composition（behavioral） |
| L2 | route + scenario（條件） |
| L3 | mechanical — **不承諾** |

#### D6 — platform_map 極限縮（Q10）

```yaml
platform_map:
  aria: <role or concept>
  headless: <headless primitive name>
  project: <project component alias>   # optional
```

**禁止** Material / iOS / Fluent / Bootstrap / Ant / Chakra 百科。跨 DS → 另開 `platform-adapter`。

#### D7 — evidence_level（Q11）

| 值 | 含義 |
| --- | --- |
| `verified` | 有可覆核證據 |
| `observed` | 實務最佳實踐，尚未硬證 |
| `hypothesized` | 暫定；展開時不得當硬 gate |

L1 / Expansion：**verified** 可硬要求；**observed** 建議；**hypothesized** 僅提示。

#### D8 — Intent：不做 Intent DB（Q7）

僅 entry 內 `intent_examples`。`Task → Intent（任務內）→ Pattern`。禁止 Intent Ontology。

#### D9 — promote-or-sunset（事件驅動，非固定日曆）

取代「固定截止日期」。**任一 trigger 成立即啟動 review**（不必等齊全部條件）：

| # | Trigger | 說明 |
| --- | --- | --- |
| T1 | 時距 | 自本 plan `in-progress` 起 **滿 3 個月** |
| T2 | 跨專案採用 | **第一次** cross-project 採用 Core schema |
| T3 | 消費者數 | **≥2** consumers 採用 Core schema（T2 可計入其中之一） |
| T4 | 深度 | **≥5** completed Core pattern entries（Ai-skill 種子 + consumer 合計） |
| T5 | 放棄 | dogfood **abandoned**（書面） |

Stakeholder 強調的壓縮讀法：**T1 OR T2 以先發生者啟動**；T3–T5 同屬「任一成立」集合，不互相覆蓋。

**Outcome（三選一，書面）**：`Promote`（含可選 L2 wiring）/ `Continue`（延續 doc-only + 寫下一次 review 條件）/ `Sunset`。

**禁止**：用單一固定日（如 `2026-09-30`）代替本表。

### Alternatives Considered

- **A. 僅 Vocabulary**：reject。  
- **B. Intent DB**：reject。  
- **C. platform_map 各大 DS**：reject。  
- **D. 每次 Expansion commit**：reject。  
- **E. 強制 recipe 齊才完整**：reject。  
- **F. 固定畢業日（如 2026-09-30）**：reject — 改 D9 事件驅動。  
- **G. Core/Extended + evidence_level + transient checklist + Authority Boundary（accept）**。

### Why Not an ADR Yet

未 dogfood；枚舉仍可微調。Promotion 需消費者採用 + expansion evidence。

### ADR Promotion Criteria

- [ ] foundational + cross-session + cross-project + expensive-to-reverse + explains-why  
- [ ] ≥1 外專案用 Core schema  
- [ ] ≥1 composition；≥1 expansion **evidence**  
- [ ] Open Questions 全解  
- [ ] 無更輕 target 仍夠用  
- [ ] D9 review 曾執行且 outcome 不是 Sunset-without-replacement 的倉促放棄

### Consequences

#### 正面
- 選型／合規／證明分權清楚（Authority Boundary）。  
- AI 可依 intent 選 pattern 並帶成熟度。  

#### 風險
- Core 維護成本 → 首波 overlay family。  
- `evidence_level` 亂標 → Phase 2 升級須掛證據指針。

**Glossary Impact**: yes — 見檔頭；完成前不註冊。

---

## Runtime Execution Path

**Doc-only 語義（明示）**：

| 允許（Phase 1–3） | 禁止自稱 |
| --- | --- |
| 改 workflow markdown / templates / indexes | 「已完成 runtime integration」 |
| 專案 overlay entries / composition | 新增 `route.*` 卻無 consumer |
| plan `evidence/` **摘要**（非每次 checklist 入 Knowledge） | 把 transient checklist commit 進 canonical Knowledge |

**不接入 runtime**（不新增 `runtime/*.yaml` projection、不接 commit-msg 檔名攔截），直到 Phase 5 **且** D9 outcome = Promote 含 L2。  
**Deferred Runtime Projection**：不新增 runtime YAML。  
**Per-surface consumer 表**：N/A until Phase 5。  
**Graduation**：見 D9（事件驅動），非單一日曆日。

---

## Open Questions

| ID | 狀態 | 決議 |
| --- | --- | --- |
| Q1 | **已決議** | Knowledge Lock **掛 `ui-contracts.md`**（與 Screen Mapping 分開小節；見 D1） |
| Q2 | **已決議** | 規則 Ai-skill；採用名／composition 專案 |
| Q3 | **已決議** | 首波 overlay family + toast / scrim / empty_state |
| Q4 | **已決議** | `legacy_alias`；不强制 rename |
| Q5 | **已決議** | 不要圖鑑縮圖 |
| Q6 | **已決議** | L1 只挡 pattern/overlay 選型與對齊 claim |
| Q7 | **已決議** | 不做 Intent DB；僅 `intent_examples` |
| Q8 | **已決議** | Core 必填；Recipe = Extended + status |
| Q9 | **已決議** | Knowledge 進 repo；checklist = transient |
| Q10 | **已決議** | platform_map 僅 aria → headless → project |
| Q11 | **已決議** | `evidence_level`：verified / observed / hypothesized |

---

## Phase 0 — Pre-Build Interrogation

### Phase 0.0 — Open Questions 核對（公版，必填）

逐條核對本 plan §Open Questions，標記處置並回寫：

- [x] 已讀本 plan §Open Questions 全部條目
- [x] 對每條標記 `resolved` / `still-open` / `deferred`
- [x] `resolved` 的條目已同步於 §Open Questions（Q1–Q11 均已決議）
- [x] 若盤點新發現問題，已加入 §Open Questions（本輪無新增 still-open）

| Open Question | 處置 | 證據 / 原因 |
| --- | --- | --- |
| Q1 掛點 | resolved | Stakeholder + Round-1 Verifier：掛 `ui-contracts`；D1 |
| Q2–Q11 | resolved | Stakeholder 決議；見 §Open Questions 表 |

### Phase 0.1 — 其餘盤點

- [x] Q1–Q11 決議入檔  
- [x] Authority Boundary 入 Executive Summary  
- [x] Inner-docs 邊界：knowledge 放 outer `docs/frontend-contracts/ui-pattern-knowledge/`  
- [x] 混名樣本表（Appendix A）≥5  
- [x] 正式簽署 → `in-progress`（stakeholder 2026-07-14：先跑 Phase 0；R2 wording-only 已清）

### Phase 0.2 — Pre-build Interrogation + Architecture Compatibility Preflight

| 欄位 | 內容 |
| --- | --- |
| Trigger | 完成 Phase 0 → 允許之後進 Phase 1（本輪**只關 Phase 0**，不開始 Phase 1 實作） |
| Goal | 補 UI Pattern Knowledge Layer（選型／組成／展開），掛在 contracts→design contract 之間 |
| Scope | Ai-skill：templates + `ui-contracts` Lock + indexes + overlay Core 種子；consumer overlay 屬 Phase 2 |
| Non-goals | Intent DB；DS 百科 platform_map；機械檔名攔截；新 lifecycle / 獨立 slice（先）；自動設計 agent |
| Acceptance（Phase 0） | Q 全 resolved；Appendix A；Authority Boundary；preflight Decision=proceed；status=`in-progress` |
| Checked sources | `ui-contracts.md`、`ui-governance.md`、`artifact-gates.md`、`plans/README.md` Phase 0.0／preflight、本 plan D1–D9 |
| Layer | workflow（gate+template）；專案 overlay 屬 consumer docs；不搶 Design Contract / Governance |
| Compiler / runtime | Phase 1–3 **doc-only**（見 Runtime Execution Path）；無 `runtime/*.yaml`；不宣稱 runtime integration |
| Duplication risk | 低：Lock 獨立小節，不塞 Screen Mapping；不建第二套 Intent/平台百科 |
| Conflicts | 無 blocking；Screen Mapping「不承載 framework pattern」與 Lock 分開小節已寫死 |
| Linked updates（Phase 1 時） | `ui-contracts.md`、`artifact-gates.md` / software-delivery `README` / `execution-flow` 索引；**本 Phase 0 不改那些檔** |
| Open Questions 核對 | Q1–Q11 全 `resolved`（見 0.0） |
| Decision | **proceed**（Phase 0 close；Phase 1 另開，建議 Executor Task） |
| Validation | checklist + R1/R2 Verifier evidence（delegation `evidence/2s-…`） |

- [x] Pre-build Interrogation 記錄齊  
- [x] Architecture Compatibility Preflight Decision = proceed  

**Phase 0 完成條件**：同意項簽署；Appendix A 就緒；preflight proceed — **全部滿足（2026-07-14）**。

---

## Phase 1 — Ai-skill：Core schema + Selection

**Phase 1 Done Definition（鎖定，不再加產物）**：

```text
workflow/software-delivery/
  ui-contracts.md                 § Pattern Knowledge Lock
  templates/ui-pattern-knowledge.* （entry / composition / prompt-expansion）
  ui-pattern-knowledge/
    entries/*.yaml                # 五件 overlay，一 pattern 一檔（非 mega-glossary）
    pattern-index.yaml            # Discovery index（禁止叫 runtime-index）
    validation/entry-schema.yaml  # 結構驗證（非 Runtime）
    README.md
```

- [x] `ui-pattern-knowledge.entry.template.yaml`（Core **FROZEN v1.0**）  
- [x] `ui-pattern-knowledge.composition.template.yaml`  
- [x] `ui-pattern-prompt-expansion.template.md`  
- [x] `ui-contracts.md`：**Pattern Knowledge Lock**  
- [x] indexes 一行  
- [x] overlay Core 種子 ×5（bottom_sheet / modal_dialog / drawer / toast / scrim）  
- [x] `pattern-index.yaml` + `validation/entry-schema.yaml`（stakeholder 補強；非 runtime）

**完成條件**：上表產物齊；Core 凍結；**不再**為此 Phase 加更多 pattern／slice／runtime surface。

---

## Phase 2 — Pattern Family 可推理性驗證（**不是「寫完五份 YAML」**）

> 一句目標：**證明五個 Entry 能形成一個可工作的 Pattern Family**（Knowledge 可被推理），不是累積名詞、也不是再打磨 Phase 1 產物。

**範圍鎖死**：仍只有這五件（禁 popover / tooltip / fab / …）。  
**主證據** = Selection Scenario dogfood；consumer project alias / outer↔inner 雙鏈 = **次要**（可並行，不擋主成功條件）。

### 驗證順序（刻意，非隨意）

| 序 | Pattern | 本輪主要驗什麼 |
| --- | --- | --- |
| 1 | **Scrim** | family / recipe / relationship；**幾乎無「主表面選型」**——寫壞則後面全亂 |
| 2 | **Modal Dialog** | Selection + Relationship + Recipe（例：destructive confirm → Dialog，不是 Sheet） |
| 3 | **Bottom Sheet** | Near Neighbor：Sheet vs Dialog vs Drawer |
| 4 | **Drawer** | persistent vs temporary；Navigation vs Task |
| 5 | **Toast** | **另一個 Family（feedback）**——不屬 Overlay Decision；驗 Family 是否真有用 |

每完成一個 entry 強化回合，就做一輪 **Selection Test**（給意圖 → 應選出正確 pattern）。

### 成功條件（取代「五個 Entry」 alone）

```text
五個 Entry（同波次）+ 十個 Selection Scenario 全部通過
  Scrim ×2 · Modal Dialog ×2 · Bottom Sheet ×2 · Drawer ×2 · Toast ×2
```

- [x] 證據檔：[`evidence/selection-scenarios.yaml`](./2026-07-14-0856-ui-pattern-knowledge-workflow/evidence/selection-scenarios.yaml)（十案定義）  
- [x] 證據檔：[`evidence/2a-family-inferability-run.md`](./2026-07-14-0856-ui-pattern-knowledge-workflow/evidence/2a-family-inferability-run.md)（逐案 PASS/FAIL + rule trace）— **10/10 PASS**（rule-trace，2026-07-14）  
- [x] toast.`family` = `feedback`（與 overlay 決策分開，作為 Family 邊界 dogfood）  
- [x] 依序強化五件 selection_rules / neighbors，使十案可推出正確答案  
- [x] 本 Phase **未**新增其他 pattern entry（遵守）  
- [x] Closure：[`evidence/phase2-summary.md`](./2026-07-14-0856-ui-pattern-knowledge-workflow/evidence/phase2-summary.md)  
- [ ]（次要，defer）`<PROJECT_ROOT>` project alias / recipe partial — 不計入本 Phase gate

**Dogfood loop pointer（#4 defer）**：三角色可選；非完成條件。

### Phase 2 正式結論

> Phase 2 驗證的是 Pattern Knowledge 的 **「可推理性（inferability）」**，而不是 Pattern **Coverage（涵蓋率）**。因此 Phase 2 的完成條件以推理能力為準，而非 Pattern 數量。

| Gate | 判定 | 理由 |
| --- | --- | --- |
| Phase 2 | ✅ **Completed** | H1 Selection、H2 Family、H3 Near Neighbor 皆有可重現證據（rule-trace ≡ blind LLM 10/10） |
| Phase 3 | ✅ **Start** | 目標升到「多 Pattern 可組成 Screen 並保持可推理」 |

**關閉日**：2026-07-14（stakeholder gate decision）。

---

## Phase 3 — Pattern Composition / Composability（**正式開始**）

> 名稱用 **Pattern Composition**。驗證單位 = **Screen**。核心能力 = **Composability**。  
> **審查對象轉變**：不再驗「知識本體」，而驗 **知識之間的約束（constraints）**。Phase 2 審 Selection；Phase 3 審 Constraint。  
> Start lock：[`evidence/phase3-start.md`](./2026-07-14-0856-ui-pattern-knowledge-workflow/evidence/phase3-start.md)  
> Cross-plan rhythm：[`Architecture Evolution Protocol` §Layer Growth Rhythm](../../governance/lifecycle/architecture-evolution-protocol.md#appendix--layer-growth-rhythmoptional-governance-pattern)（*Grow one layer, freeze the previous*）

**一句目標**：證明多個 Pattern Knowledge 能共同工作，且彼此不污染。

### Phase 3 Invariant（升格 — 高於 checklist）

```text
Composition evidence MUST NOT modify validated Pattern Knowledge entries.

New knowledge discovered during composition
  → composition_rules.yaml   (= Composition Constraints)
NOT → entries/*.yaml
```

| 風險 | 說明 |
| --- | --- |
| **Back-propagation（回滲）** | Phase 3 最大風險。每 Screen 回改 Entry → 無法分辨 Entry 正確還是 Episode 在修 Entry |

違反 Invariant = Phase 3 證據作廢（revert entry + 改寫 Constraints）。

### Mini-cycles（H4 → H5 → H6 獨立跑；不要一次跑完整 Episode）

| Cycle | 問什麼 | 刻意不做 | Deliverable |
| --- | --- | --- | --- |
| **H4 Independence（先）** | 施壓共現／置換後 Decision Boundary 變了沒？ | 不「完成畫面」 | Independence evidence（A±B / Dialog↔Sheet） |
| **H5 Completeness（次）** | 每個 Node = Pattern **或** deferred？零 Unknown？ | 不管 Selection；不補 Entry | Completeness matrix；gap → `knowledge: deferred` + `reason: uncovered_pattern` |
| **H6 Traceability（末）** | 每個 Node 能否追？ | 不以 YAML 齊全當成功 | **Trace Graph**（鏈圖） |

#### H4 施壓動作（示例）

```text
Episode Detail + Bottom Sheet + Toast + Dialog
  · 加 Toast / 移掉 Toast
  · Dialog ↔ Sheet 互換
觀察：Bottom Sheet Selection Boundary 是否改變？
  不變 → H4 PASS
  需要改 Entry → 違反 Invariant → 寫 composition_rules.yaml
```

#### H5 deferred 形狀

```yaml
knowledge: deferred
reason: uncovered_pattern   # Completeness ≠ Coverage
```

#### H6 Trace Graph（真正 deliverable）

```text
Episode Detail
└── Bottom Sheet
      ├── selection_rule
      └── implementation_recipe
```

### Composition Metrics（非 KPI；Closure 必填）

| Metric | 意義 | Expected |
| --- | --- | --- |
| **Deferred Nodes** | Tree 上 explicit deferred 數 | ≥0（誠實） |
| **Composition Rule Count** | `composition_rules.yaml` constraints 數 | ≥0（隨發現成長） |
| **Entry Modifications** | Phase 3 期間 `entries/*` 變更數 | **0**（Invariant） |

`Entry Modifications ≠ 0` ⇒ Phase 2 Freeze 未被真正守住；Closure 失敗。

基線／進度見 [`evidence/3-metrics.md`](./2026-07-14-0856-ui-pattern-knowledge-workflow/evidence/3-metrics.md)。  
H4 stress：[`evidence/3h4-independence-stress.md`](./2026-07-14-0856-ui-pattern-knowledge-workflow/evidence/3h4-independence-stress.md)（Case A 首輪 FAIL → 新 Constraint；Entry Mods=0）。

### Scope 鎖死

只驗 **Episode Detail**（不要一開始很多 Screen）。App Bar / Player = deferred；**禁止**為 Composition 新建 Entry。

### Artifacts

- [x] Pattern Tree：[`compositions/episode_detail.yaml`](../../workflow/software-delivery/ui-pattern-knowledge/compositions/episode_detail.yaml)  
- [x] Constraints：[`composition_rules.yaml`](../../workflow/software-delivery/ui-pattern-knowledge/composition_rules.yaml)  
- [x] Start lock：[`evidence/phase3-start.md`](./2026-07-14-0856-ui-pattern-knowledge-workflow/evidence/phase3-start.md)  
- [x] Metrics baseline：[`evidence/3-metrics.md`](./2026-07-14-0856-ui-pattern-knowledge-workflow/evidence/3-metrics.md)  
- [x] **H4 mini-cycle** evidence（Independence stress）— [`3h4-independence-stress.md`](./2026-07-14-0856-ui-pattern-knowledge-workflow/evidence/3h4-independence-stress.md)  
- [ ] **H5 mini-cycle** evidence（Completeness only）  
- [ ] **H6 mini-cycle** evidence（Trace Graph）  
- [ ] Metrics 終值記入 Closure  
- [x] **Invariant**：Entry Modifications = 0（H4 後仍 0）

### Exit Gate = **Composition Closure**

| 不要 | 要 |
| --- | --- |
| Episode Detail 完成 | Episode Detail **Pattern Tree Validated** |
| 一次跑完整 Screen | 三個 mini-cycle 分檔通過 + Metrics |

**Composition Closure** = H4∧H5∧H6 mini-cycles PASS + Entry Modifications = 0。

（觀察：Flow·Orchestrability 可能自然露出——**不**預寫 plan Phase 4。）

## Phase 4 —（可選）Extended 加厚

- [ ] anti_patterns、recipe→partial/complete、evidence_level 升級規則  
- [ ] 仍不引入 DS 百科；governance 一行承接 anti-pattern warn

---

## Phase 5 — Optional L2（僅當 D9 → Promote）

- [ ] `load_when` + scenario + consumer 表；不做檔名機械攔截

---

## Deferred Notes（#4 / #7）

| ID | 處置 | 短註 |
| --- | --- | --- |
| #4 三角色入口 | defer | Phase 2 pointer → delegation kit；非本 plan scope |
| #7 outer↔inner 雙鏈 | defer | dogfood 導航邊界；不升格為雙寫 Knowledge |

---

## 完成條件（整 plan）

- [ ] Phase 0–2（✓）；Phase 3 有 **Pattern Composition** + expansion **evidence**  
- [ ] Core/Extended 邊界寫進 template  
- [ ] 無 Intent DB；platform_map 無 DS 百科  
- [ ] Q 全關；glossary 註冊或明示不註冊  
- [ ] 無 runtime 則不宣稱 runtime integration  
- [ ] D9 review 至少執行一次（Promote / Continue / Sunset 書面）

---

## Stakeholder 同意項目

- [x] 升格 UI Pattern Knowledge  
- [x] Q1–Q11 決議（含 Q1 掛 ui-contracts）  
- [x] Round-1 仲裁：#1/#2/#3/#5/#6/#8 fix；#4/#7 defer；#9–11 reject  
- [x] D9 事件驅動 promote-or-sunset（非固定日）  
- [x] Authority Boundary 四句  
- [x] Round-2 Verifier 通過 → `in-progress`（2026-07-14 Phase 0 close）  
- [x] 簽署開始 Phase 1 實作（stakeholder 2026-07-14：「可以繼續」）  
- [x] Phase 2 Completed / Phase 3 Start（stakeholder 2026-07-14：三假說通過；inferability ≠ coverage）  
- [x] Research ladder + Phase 3 H4–H6 凍結；composition_rules = Composition Constraint（非 Entry／非 Rule Library）  
- [x] Phase 3 formal start：Entry Freeze **升格 Invariant**（anti back-propagation）；Exit = Composition Closure；P4 Orchestrability 不入 plan
- [x] Phase 3 mini-cycles（H4→H5→H6）+ Composition Metrics；Layer Growth Rhythm recorded in Architecture Evolution Protocol

---

## 與其他 plans / surfaces 的關係

| 關係 | 說明 |
| --- | --- |
| `ui-contracts.md` | Knowledge Lock **已決議掛點** |
| `ui-governance.md` | 合規；不擁有選型 |
| Design Contract | 怎麼建成（token/primitive） |
| Evidence Governance | claims 是否被證明；`evidence_level` |
| [`2026-07-08-0825-delegation-…`](../2026-07-08-0825-delegation-verification-arbitration-loop/_plan.md) | dogfood 可選三角色 SOP；本 plan 不擁有該 loop |
| `platform-adapter`（未來） | 跨 DS；不進本 Knowledge |
| NameThatUI | 靈感；我方 = AI-consumable Knowledge Layer |

---

## 已決議摘要（Q1–Q11）

| 問題 | 決議 |
| --- | --- |
| Q1 | Lock 掛 `ui-contracts.md`（獨立小節） |
| Q2 | 規則 Ai-skill；採用名專案 |
| Q3 | 首波 overlay + toast/scrim/empty |
| Q4 | legacy_alias |
| Q5 | 無縮圖 |
| Q6 | L1 窄挡 |
| Q7 | 無 Intent DB |
| Q8 | Core / Extended；recipe status |
| Q9 | checklist transient |
| Q10 | aria → headless → project |
| Q11 | evidence_level 三階 |

---

## Appendix A — 混名樣本（Phase 0 前置證據，≥5）

Sanitized 示例（真實路徑留 `<PROJECT_ROOT>`）；說明「為何需要 Pattern Knowledge」：

| # | 觀察到的符號／說法 | 像哪個 pattern | 易混成 | 本 plan 預期 canonical |
| --- | --- | --- | --- | --- |
| 1 | 底部分享面板（Dialog primitive + sheet 動效） | bottom_sheet | modal_dialog | `bottom_sheet` + `legacy_alias` 可指向現名 |
| 2 | 置中離開確認（title + 雙按鈕） | alert / modal_dialog | bottom_sheet | `modal_dialog`（destructive／阻斷） |
| 3 | 全屏播放層覆蓋 Tab | immersive overlay | bottom_sheet / modal | `overlay` family 變體（首波可 `deferred` 細分） |
| 4 | 全域系統錯誤提示 | toast | modal_dialog | `toast` |
| 5 | 工程文檔「Drawer／Modal」並用 | drawer vs modal | sheet | 依 selection_rules；持久導航≠臨時面板 |
| 6 | 半透明壓暗背後 | scrim | 頁面背景色 | `scrim`（常為 overlay recipe 必備） |

**用途**：證明 Phase 0 非空話；Dogfood Phase 2 用此表回填 `<PROJECT_ROOT>` 真實 alias，不把私有名寫進 Ai-skill 正文。
