---
id: 2026-07-14-0856-ui-pattern-knowledge-workflow
plan_kind: main
status: draft
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
    note: "Q7–Q10 決議入檔：Core/Extended schema；禁止 Intent DB；Knowledge vs transient checklist；platform_map 僅 ARIA→Headless→Project；新增 evidence_level / recipe status maturity"
  - date: 2026-07-14
    note: "Sanitization：dogfood 改 <PROJECT_ROOT>；去掉專案專屬 token，以過 shared-layer pre-commit"
---

# UI Pattern Knowledge — Workflow 強化計畫

**Status**: `draft` — Q7–Q10 與 maturity 欄位已決議；勾選同意項後 → `in-progress`  
**Owner**: linyihong  
**建立日期**: 2026-07-14  
**Source**: 2026-07-14 對話 — NameThatUI 對照 + stakeholder 回饋（Knowledge layer，非 glossary；非 NameThatUI clone）。  
**Glossary Impact**: yes — candidates（未註冊）：`ui_pattern_knowledge`、`pattern_selection_rules`、`pattern_composition`、`pattern_family`、`implementation_recipe`、`pattern_prompt_expansion`、`pattern_knowledge_core|extended`、`pattern_evidence_level`（verified / observed / …）。舊詞 `ui_pattern_vocabulary` = 子能力。graduate 後才註冊。

> **Watch-Out List citation**（[`architecture/ai-native-cognitive-ecosystem-system.md`](../../architecture/ai-native-cognitive-ecosystem-system.md) §Watch-Out List）：
> - **process bloat** — 不新建 lifecycle；不建 Intent Ontology / Design-System 百科；Prompt checklist **不**每次 commit。
> - **premature abstraction** — platform_map 只留三層；Material/iOS/Fluent 若需要另開 `platform-adapter`，不混進 Pattern Knowledge。
> - **autonomous optimizer** — Expansion = 從 Knowledge 展開 transient checklist，不是自動改設計。

## Executive summary

定位：**可被 workflow 消費的 UI Pattern Knowledge Layer**（載入 → 推理選型 → 展開 → 驗證），不是人查 glossary，也不是模仿 NameThatUI。

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

**與既有分工**：Knowledge = 選什麼／為什麼；Design Contract = token/primitive 合規；UI Governance = compliance 分類與收口；Evidence = 證明。互不重疊。

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
- 先掛 `ui-contracts`；不開六道 slice / 不開 Intent repository

#### D2 — Schema = Core + Extended（Q8 決議）

**Core（缺一不可算完整 entry）**

| 欄位 | 作用 |
| --- | --- |
| `canonical_pattern` | 正式名 |
| `plain_name` | 人話 |
| `intent_examples` | 適用意圖示例（非產品 Intent DB） |
| `selection_rules`（含 when / not_when；可與 when_to_use / when_not 同義並存於模板） | 何時選／不選 |
| `family` | 家族（如 overlay） |
| `near_neighbors` | 易混 + 互斥一句 |

**Extended（可後補）**

| 欄位 | 作用 |
| --- | --- |
| `implementation_recipe` | 實作能力清單；**不**要求首日 complete |
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
  evidence_level: verified | observed | hypothesized   # D7
anti_patterns:                    # Extended
  - more_than_10_actions
  - full_form_or_wizard_inside
  - nested_bottom_sheets
platform_map:                     # Extended；僅三層 — D6
  aria: dialog
  headless: Dialog                # e.g. Base UI primitive 名
  project: ProjectShareSheet      # 專案 alias 示例；可空
implementation_recipe:            # Extended — Q8
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

**只允許三層**：

```yaml
platform_map:
  aria: <role or concept>
  headless: <headless primitive name>
  project: <project component alias>   # optional
```

**禁止**塞進 Pattern Knowledge：Material / iOS / Fluent / Bootstrap / Ant / Chakra 等百科。  
若未來需要跨 DS 對照 → 另開 **`platform-adapter`** 計畫／artifact，不混本層。

#### D7 — evidence_level / maturity（stakeholder 追加，接受）

知識條目與選型規則、recipe 可標：

| 值 | 含義（對齊 Evidence Governance 哲學） |
| --- | --- |
| `verified` | 有可覆核證據（dogfood / 實作 / 測試） |
| `observed` | 實務觀察／最佳實踐，尚未硬證 |
| `hypothesized` | 暫定假設，展開時不得當硬 gate |

Prompt Expansion 與 L1 claim：**verified** 可當硬要求；**observed** 預設建議；**hypothesized** 僅提示。

#### D8 — Intent：不做 Intent DB（Q7）

- Intent = Product Domain；Pattern = UI Domain。  
- 只在 entry 內放 `intent_examples`（choose one action / destructive confirmation / …）。  
- Workflow：`Task → Intent（任務內推導）→ Pattern`。  
- **禁止**獨立 Intent Ontology / Intent Knowledge Base（避免產品知識與 UI 知識耦合坑）。

### Alternatives Considered

- **A. 僅 Vocabulary**：reject as final shape。  
- **B. Intent DB + Pattern DB**：reject — Ontology 坑（Q7）。  
- **C. platform_map 含各大 DS**：reject — 維護成本無限（Q10）。  
- **D. 每次 Expansion commit**：reject — repo 爆炸（Q9）。  
- **E. 強制 recipe.required 齊才算完整**：reject — 改 `status: unknown|partial|complete`（Q8）。  
- **F. Core/Extended + evidence_level + transient checklist（accept）**。

### Why Not an ADR Yet

未 dogfood；schema 已定方向但仍可能微調枚舉名。Promotion 需：他專案採用 Core schema + 至少一次 expansion evidence（不强制把 checklist  upstream 成 Knowledge）。

### ADR Promotion Criteria

- [ ] foundational + cross-session + cross-project + expensive-to-reverse + explains-why  
- [ ] ≥1 外專案用 Core schema  
- [ ] ≥1 composition；≥1 expansion **evidence**（非要求 checklist 進 canonical）  
- [ ] Open Questions 全解  
- [ ] 無更輕 target 仍夠用

### Consequences

#### 正面
- AI 知道「為什麼選／不選／跟誰不同」；展開時知道成熟度。  
- 與 Design Contract / Governance / Evidence 自然分工。  
- 相對 NameThatUI：可載入、可推理、可展開、可驗證。

#### 風險
- Core 仍需維護成本 → 首波只 overlay family。  
- `evidence_level` 被亂標 → Phase 2 dogfood 規定升級規則（observed→verified 要掛證據指針）。

**Glossary Impact**: yes — 見檔頭；完成前不註冊。

---

## Runtime Execution Path

**Phase 1–3：doc-only。** Phase 5 條件升 L2 時復用 `route.workflow.software-delivery`；不新增無 consumer 的 surface；不做檔名 commit-msg。

**Deferred Runtime Projection**：不新增 `runtime/*.yaml`。  
**Per-surface consumer 表**：N/A until Phase 5。

---

## Open Questions

| ID | 狀態 | 決議 |
| --- | --- | --- |
| Q1 | 暫定 | 先掛 `ui-contracts` |
| Q2 | 暫定接受 | 規則 Ai-skill；採用名／composition 專案 |
| Q3 | 暫定 | 首波 overlay family + toast/scrim/empty |
| Q4 | 暫定接受 | legacy_alias；不强制 rename |
| Q5 | 暫定接受 | 不要圖鑑縮圖 |
| Q6 | 暫定接受 | L1 只挡 pattern/overlay 選型與對齊 claim |
| Q7 | **已決議** | 不做 Intent DB；僅 `intent_examples` |
| Q8 | **已決議** | Core 必填（含 selection/family/neighbors/intent）；Recipe = Extended + `unknown\|partial\|complete` |
| Q9 | **已決議** | Knowledge 進 repo；任務 checklist = transient（報告／evidence 摘要） |
| Q10 | **已決議** | platform_map 僅 `aria` → `headless` → `project` |
| Q11 | **已決議** | 採納 `evidence_level`（verified / observed / hypothesized） |

---

## Phase 0 — Pre-Build Interrogation

- [x] Q7–Q11 決議入檔  
- [ ] Authority 邊界複核（Knowledge ≠ Design Contract ≠ Governance）  
- [ ] Inner-docs 邊界：knowledge 放 outer `docs/frontend-contracts/ui-pattern-knowledge/`（不寫入受邊界限制的 inner docs 政策區）
- [ ] 混名樣本 ≥5  
- [ ] 同意項勾選 → `in-progress`

**完成條件**：同意項簽署；樣本表就緒。

---

## Phase 1 — Ai-skill：Core schema + Selection

- [ ] `ui-pattern-knowledge.entry.template.yaml`（標 Core vs Extended；`evidence_level`；recipe.status）  
- [ ] `ui-pattern-knowledge.composition.template.yaml`  
- [ ] `ui-pattern-prompt-expansion.template.md`（輸出形狀；註明 **勿 commit 為 Knowledge**）  
- [ ] `ui-contracts.md`：Pattern Knowledge Lock  
- [ ] indexes 一行  
- [ ] overlay family Core 種子（3–5 條）

**完成條件**：Core 被定義為完整性標準；Recipe 允許 `unknown`。

---

## Phase 2 — `<PROJECT_ROOT>` dogfood（Core 齊，Extended 可薄）

- [ ] `docs/frontend-contracts/ui-pattern-knowledge/` + ≥8 entries（**Core 齊**）  
- [ ] recipe 可 `unknown`/`partial`；至少 1–2 條 `partial`+observed 作對照  
- [ ] platform_map 只填三層（能填多少算多少）  
- [ ] Outer knowledge ↔ inner shared-components README 雙向連結  
- [ ]（可選）結構 BDD：Core key 存在

**完成條件**：任務引用 selection（不必等 recipe complete）。

---

## Phase 3 — Composition + Expansion evidence

- [ ] ≥1 composition  
- [ ] 跑一次 Expansion；evidence 留 **摘要**（非把 checklist 升成 entry）  
- [ ] 記錄 verified vs observed 在展開中的差異（若可觀察）

**完成條件**：書面 keep L1 / 開 Phase 4–5。

---

## Phase 4 —（可選）Extended 加厚

- [ ] 補 anti_patterns、recipe→partial/complete、evidence_level 升級規則  
- [ ] **仍不**引入 Material/iOS 百科；governance 一行承接 anti-pattern warn

---

## Phase 5 — Optional L2（條件）

- [ ] `load_when` + scenario + consumer 表；不做檔名機械攔截

---

## 完成條件（整 plan）

- [ ] Phase 0–2；Phase 3 有 composition + expansion **evidence**  
- [ ] Core/Extended 邊界寫進 template  
- [ ] 無 Intent DB；platform_map 無 DS 百科  
- [ ] Q 全關；glossary 註冊或明示不註冊  
- [ ] 無 runtime 則不宣稱 runtime integration

---

## Stakeholder 同意項目

- [x] 升格 UI Pattern Knowledge（非僅 Vocabulary）— 已反映在正文  
- [x] Q7–Q11 決議 — 已入檔  
- [ ] 正式簽署：D1–D8 作為執行依據  
- [ ] 正式簽署：開 Phase 0 剩餘項 → `in-progress`  
- [ ] （可選）立即 commit 本 draft plan

---

## 與其他 plans / surfaces 的關係

| 關係 | 說明 |
| --- | --- |
| `ui-contracts` | Knowledge Lock |
| `ui-governance` | anti-pattern / evidence 收口；不擁有選型 |
| Design Contract | token/primitive |
| Evidence Governance | `evidence_level` 哲學對齊 |
| 未來 `platform-adapter` | 跨 DS 對照（若需要）；**不**進本 Knowledge |
| NameThatUI | 靈感；我方 = AI-consumable Knowledge Layer |

---

## 已決議摘要（Q7–Q11）

| 問題 | 決議 |
| --- | --- |
| Q8 | Core = 名 + plain + intent_examples + selection + family + neighbors；Recipe Extended + `unknown\|partial\|complete` |
| Q7 | 不做 Intent DB |
| Q9 | Knowledge 進 repo；Expansion checklist = transient |
| Q10 | 僅 ARIA → Headless → Project |
| Q11 | `evidence_level`：verified / observed / hypothesized |
