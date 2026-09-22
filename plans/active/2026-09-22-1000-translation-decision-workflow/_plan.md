---
id: 2026-09-22-1000-translation-decision-workflow
plan_kind: main
status: in-progress
owner: workflow
owner_layer: workflow
created: 2026-09-22
parent: null
---

# Translation Decision Workflow（`workflow/translation/`）

**Status**: in-progress — Phase 0–3 + title／I12 + **Failure Pattern（I13）**（[`10`](10-failure-pattern-learning.md)）。Phase 4 optional。證據 [`evidence/`](evidence/README.md)。


**Glossary Impact**: yes — 候選詞 `translation_decision_record`（TDR）、`expression_analysis`、`expression_type_registry`、`translation_context_contract`、`candidate_space`、`decision_basis`、`constraint_responsibility`、`selection_responsibility`（後兩者若與 ERA plan 重複則 Phase 5 只 cross-link）。Phase 5 前不登記 glossary。

## 一句話目標

建立 **governed translation decision** 總綱：`TranslationContext` → Locale Resolution → Expression Analysis（artifact）→ Candidate Space → Constraints → Feasible Candidates → Selection Policy → Selection Actor → TDR → Independent Review + Mechanical Gate → Finality。LLM 只在可行集內做 Selection；workflow 管限制、驗證與收斂。

## Decision Rationale

### Problem & Why Now

「翻譯正確」不是文字對文字，而是 Meaning → Intent → Register → Cultural Expression → Target Language Expression。若做成 `原文 → LLM 翻譯 → 完成`，典型失敗包括：流行語直譯、成語字面化、諺語失文化、方言被「普通話化」、角色語氣消失、日文察する類語意硬翻、中文「小可愛」被 ASR／OCR 或模型音近誤解。

**Real-data（印尼語）**：`陈小姐` 目標 `id-ID` 時，若 workflow 未把 target locale 帶進每段 decision，模型常走「中文→英文」稱謂（`Miss Chen`／`Ms. Chen`），而非在 `Nona Chen` 等 **locale-aware** 候選中選擇。根因是 **語言判定、名稱解析、稱謂翻譯** 三責任混線——見 [`04-dogfood-case-address-title-id-ID.md`](04-dogfood-case-address-title-id-ID.md) 與 example [`05-example-address-title-chen-xiaojie-id.yaml`](05-example-address-title-chen-xiaojie-id.yaml)。

Ai-skill 已有 Loop-first／Governance-first 與 ERA v2（Evidence constrains Decision Space；Constraint ≠ Selection）。`narrative-video-production` 已有 locale 三閘（content ≠ timing ≠ layout）與 `dialogue.semantic_context`，但 **「這句怎麼翻、為什麼這樣翻」** 仍缺一等公民契約。

### Decision

新增 **`workflow/translation/`** 為 **cross-cutting capability workflow**（非 NVP 子目錄）。第一版是 **Translation Decision Workflow**，不是 Translation Pipeline。

凍結意圖（Phase 0 freeze；細節 [`06`](06-phase-0-freeze-invariants.md)）：

- Expression Analysis 為必經 **artifact**；禁止 `source_text → translation` 跳步；**≠** 「必須 LLM 分析」。
- **Constraint ≠ Selection**；LLM 不能當 Governance。
- **Locale Resolution ≠ Language Detection**；`target_locale` 是 **authoritative Constraint 輸入**，不是推論出的 translation decision。
- **Candidate Space ≠ Feasible Candidates ≠ Selected**；`title_mapping` 只種子 Candidate Space。
- 品質用多維 validation + finality；禁止單一 quality／confidence score。
- 所有 actor 只吃 `TranslationContext`；禁止僅 `{ src, dst }`。
- Phase 1 **只釘資料契約**（Context／Analysis／Decision／Registry／Validation／examples／P0 regression）；**不加** memory／glossary engine／prompt／routing／auto-correct／runtime route。

架構 [`01`](01-architecture-and-era.md) · SoT [`02`](02-sot-contracts-and-layout.md) · NVP [`03`](03-nvp-and-adapters.md) · dogfood [`04`](04-dogfood-case-address-title-id-ID.md) · freeze [`06`](06-phase-0-freeze-invariants.md) · realization [`07`](07-target-locale-realization.md) · walkthrough [`08`](08-static-walkthrough-pass.md) · title [`09`](09-title-content-type.md) · failure learning [`10`](10-failure-pattern-learning.md)。

### Domain Boundary

| 任務 | 本 workflow | 其他 |
| --- | --- | --- |
| Expression analysis、strategy、candidates、selection、TDR、content-oriented validation | Primary | — |
| 字幕 timing／layout／CPS／burn | Adapter 消費方定義 | `narrative-video-production` locale 三閘 |
| ASR／OCR／phonetic 噪音 | Evidence 輸入 | NVP sanitization／phonetic companions |
| 翻譯模型、API、prompt 全文 | Adapter only | `software-delivery`（若實作工具） |
| 術語、方言、文化、角色語氣知識 | `knowledge/translation/` | 不塞進 execution-flow 逐步碼 |

### Alternatives Considered

- **A. 把翻譯步驟寫進 NVP execution-flow**：reject。翻譯跨文件／UI／影片；NVP 應 consume adapter，不是 owns 決策契約。
- **B. 只做 prompt／模型選型文件**：reject。無 TDR、無 registry，無法追溯「為什麼這樣翻」。
- **C. 單一 pipeline + quality score**：reject。與 governance 方向衝突。
- **D. Translation Decision Workflow + 三 SoT 先行（accept）**。

### Why Not an ADR Yet

尚無第二個 domain consumer（除 NVP subtitle 規劃外）與真實 dogfood TDR。先 plan + workflow contracts；schema 穩定且被 ≥2 adapter 消費後再評估 ADR。

### ADR Promotion Criteria（completed 時評估）

- [ ] 三 SoT 穩定 ≥1 次 schema revision 仍向後相容或明確 migration。
- [ ] ≥1 真實 dogfood（影片或文件）含完整 TDR + finality。
- [ ] subtitle adapter 與 NVP `content_gate` 對接已文件化且至少一例通過 content validation。
- [ ] 未把特定翻譯供應商寫進 canonical 步驟。
- [ ] Open Questions 全解或有 deferred owner。

### Consequences

#### 正面

- 翻譯爭議可沿 Source → Analysis → Strategy → Candidates → Selection → Review → Final 追溯。
- 換模型／供應商只換 Selection Actor，不重做 workflow。
- 與 delegation／ERA dogfood 同一套 Constraint／Selection 語言。

#### 負面

- 每 segment 多一份 analysis／TDR 成本（可 batch／可 defer review）。

#### 風險

- **Workflow inflation**：把 ffmpeg／API 步驟寫進 execution-flow。緩解：只到紀錄與閘；工具留 adapter。
- **Registry 過早膨脹**：緩解：v0 最小 expression type 集 + examples，知識庫骨架 Phase 3。

## 已裁決 Open Questions

| # | 決策 |
| --- | --- |
| Q1 計畫形態 | **獨立 main plan**（NVP 為 consumer，非 sub-plan） |
| Q2 目標語 | **契約語系無關**；examples 用 ja／zh-TW／**zh-CN→id-ID**（稱謂案例） |
| Q3 缺 context | 無必要 context 時 **禁止 `accepted`**；見 I9 Finality closure |
| Q4 TDR 存放 | v0 只定 **schema**；專案路徑由 consumer profile 定 |
| Q5 Phase 0 freeze | **2026-09-22 accept** — 進 Phase 1；I1–I10 釘死（[`06`](06-phase-0-freeze-invariants.md)） |

## Phases

### Phase 0 — Plan freeze

- [x] 計畫落盤
- [x] 第一個 real-data case：陈小姐 → id-ID（[`04`](04-dogfood-case-address-title-id-ID.md) + [`05`](05-example-address-title-chen-xiaojie-id.yaml) = **P0 regression fixture**）
- [x] Stakeholder review：Phase 0 freeze → `in-progress`；邊界 [`06`](06-phase-0-freeze-invariants.md)
- [x] 確認不做：runtime／route／prompt／工具／memory／glossary engine／score／auto-correct（Phase 0–1 邊界）

### Phase 1 — Contracts + Registry（doc-only）

- [x] 建立 `workflow/translation/` 目錄骨架
- [x] 落地 SoT：context／analysis／decision／validation／finality + registry（含 realization）
- [x] README + `execution-flow.md`
- [x] `examples/`：P0 id+ja + literal／idiom／slang／proverb／dialect／wordplay
- [x] Registry invariant：Candidate Space ≠ Final Answer；I11 Realization
- [x] 寫作檢查 I1–I3；靜態走讀 A+B **PASS**（[`08`](08-static-walkthrough-pass.md)）

### Phase 2 — Subtitle adapter + NVP link

- [x] `adapters/subtitle.yaml` ↔ `caption-locale-pack` `content_gate`
- [x] NVP [`captions-and-locales.md`](../../workflow/narrative-video-production/captions-and-locales.md) inbound pointer
- [x] 明文化 `semantic_context` 為 Translation Context 上游（`dialogue-semantic-context.yaml` inbound）

### Phase 3 — Knowledge skeleton + dogfood

- [x] `knowledge/translation/` 骨架（ja-JP name-realization + [`title-mapping`](../../knowledge/translation/locale/title-mapping.yaml) 最小種子）
- [x] Qwen ep12–14 dogfood run 入 [`evidence/`](evidence/README.md)（[`2026-09-22-qwen-ep12-14-locale-realization.md`](evidence/2026-09-22-qwen-ep12-14-locale-realization.md)）
- [x] Title／content_type 補強（I12）：[`09`](09-title-content-type.md) + fixture [`title-kongjie-yiriqianli`](../../workflow/translation/examples/title-kongjie-yiriqianli.yaml) + evidence [`2026-09-22-title-yiriqianli`](evidence/2026-09-22-title-yiriqianli.md)（**不成** title-translation-workflow）
- [x] Failure Pattern learning（I13）：[`10`](10-failure-pattern-learning.md) + [`failure-pattern`](../../workflow/translation/contracts/failure-pattern.yaml)／[`failure-patterns`](../../workflow/translation/registry/failure-patterns.yaml) — dogfood → abstract guard，**不**堆 Selection prompt

### Phase 4 — 可選（不擋 v0）

- [ ] document／UI／dubbing adapter
- [ ] `route.workflow.translation` + runtime 投影（需 Phase 3 證據）

## Out of Scope（v0／Phase 1）

- 翻譯 API／prompt／模型選型／model routing 寫進 canonical
- translation memory、glossary engine、automatic terminology extraction
- 單一 quality／confidence score、automatic correction
- 合併 NVP timing／layout 進 translation content gate
- 完整方言／成語知識庫當 dictionary truth
- 自動 orchestrator、runtime route、test runner（regression 先 doc fixture）
- Selection adapter 內 `if locale: += 例句` 膨脹（改走 failure_pattern → knowledge）
- LLM 自行把 learning candidate 提升為 active governance（I13）

## Success Criteria（v0 完成）

- [x] SoT 可獨立閱讀；I1–I13 寫進 contracts（含 content.type／title／failure_pattern）
- [x] ≥3 expression examples + P0 陈小姐 **id + ja** + **空姐被一日千里** title fixtures
- [x] README 說清 Context／Content-Type／Analysis／Realization／Registry／Guards／Constraints／Candidates／Policy／Actor／Finality／Failure Learning
- [x] subtitle adapter 餵 NVP content_gate，不碰 timing／layout
- [x] Phase 3 dogfood evidence 入庫（含 title invented_information）
- [x] 未註冊 route、未 runtime 投影、未另開 title-translation-workflow、未用 locale if-case 膨脹 Selection prompt

## Linked Updates（Phase 1+ 觸發）

- [`workflow/README.md`](../../workflow/README.md) — 新增 translation 入口（Phase 1）
- [`workflow/workflow-routing.md`](../../workflow/workflow-routing.md) — 選路表（Phase 4 或 Phase 1 若先註冊 draft route）
- [`knowledge/runtime/routing-registry.yaml`](../../knowledge/runtime/routing-registry.yaml) — Phase 4 only
- NVP plan [`2026-09-16-1649-narrative-video-production-workflow/_plan.md`](../2026-09-16-1649-narrative-video-production-workflow/_plan.md) — consumer 關係一行（Phase 2）

## References

- ERA / Constraint≠Selection：[`plans/active/2026-07-08-0825-delegation-verification-arbitration-loop/_plan.md`](../2026-07-08-0825-delegation-verification-arbitration-loop/_plan.md)
- NVP invariants：[`plans/active/2026-09-16-1649-narrative-video-production-workflow/03-architecture-invariants.md`](../2026-09-16-1649-narrative-video-production-workflow/03-architecture-invariants.md)
- Captions：[`workflow/narrative-video-production/captions-and-locales.md`](../../workflow/narrative-video-production/captions-and-locales.md)
- Semantic context：[`workflow/narrative-video-production/records/dialogue-semantic-context.yaml`](../../workflow/narrative-video-production/records/dialogue-semantic-context.yaml)
