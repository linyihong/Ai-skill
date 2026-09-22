---
id: 2026-09-22-1000-translation-decision-workflow
plan_kind: main
status: draft
owner: workflow
owner_layer: workflow
created: 2026-09-22
parent: null
---

# Translation Decision Workflow（`workflow/translation/`）

**Status**: draft — 計畫已落盤；Phase 0 待確認後進 Phase 1（contracts + registry，doc-only）。

**Glossary Impact**: yes — 候選詞 `translation_decision_record`（TDR）、`expression_analysis`、`expression_type_registry`、`translation_context_contract`、`constraint_responsibility`、`selection_responsibility`（後兩者若與 ERA plan 重複則 Phase 5 只 cross-link，不 duplicate 定義）。Phase 5 前不登記 glossary。

## 一句話目標

建立 **governed translation decision** 總綱：先 Expression Analysis，再 Constraints → Feasible Candidates → Selection Policy → Translation Decision Record → Independent Review → Mechanical Gate → Finality；LLM 只當 Selection Actor，workflow／runtime 管限制、驗證與收斂。

## Decision Rationale

### Problem & Why Now

「翻譯正確」不是文字對文字，而是 Meaning → Intent → Register → Cultural Expression → Target Language Expression。若做成 `原文 → LLM 翻譯 → 完成`，典型失敗包括：流行語直譯、成語字面化、諺語失文化、方言被「普通話化」、角色語氣消失、日文察する類語意硬翻、中文「小可愛」被 ASR／OCR 或模型音近誤解。

**Real-data（印尼語）**：`陈小姐` 目標 `id-ID` 時，若 workflow 未把 target locale 帶進每段 decision，模型常走「中文→英文」稱謂（`Miss Chen`／`Ms. Chen`），而非在 `Nona Chen` 等 **locale-aware** 候選中選擇。根因是 **語言判定、名稱解析、稱謂翻譯** 三責任混線——見 [`04-dogfood-case-address-title-id-ID.md`](04-dogfood-case-address-title-id-ID.md) 與 example [`05-example-address-title-chen-xiaojie-id.yaml`](05-example-address-title-chen-xiaojie-id.yaml)。

Ai-skill 已有 Loop-first／Governance-first 與 ERA v2（Evidence constrains Decision Space；Constraint ≠ Selection）。`narrative-video-production` 已有 locale 三閘（content ≠ timing ≠ layout）與 `dialogue.semantic_context`，但 **「這句怎麼翻、為什麼這樣翻」** 仍缺一等公民契約。

### Decision

新增 **`workflow/translation/`** 為 **cross-cutting capability workflow**（非 NVP 子目錄）。第一版是 **Translation Decision Workflow**，不是 Translation Pipeline。

凍結意圖：

- Expression Analysis 為必經中間層；禁止 `source_text → translation` 跳步。
- **Constraint Responsibility**（不能錯什麼）可大量機械化；**Selection Responsibility**（哪個表達最好）由明示 policy + Selection Actor，LLM 不能當 Governance。
- 品質用 **多維 validation + finality**，禁止單一 `translation_quality: 0.92`。
- **translation-core** + **adapters**（document／subtitle／UI；dubbing 後續）；不把四套互不相干的流程寫死。
- Locale 模型：`language ≠ locale ≠ dialect ≠ register`；禁止隱式 `方言 → 標準語 → 目標語` flattening。
- Phase 1 **SoT**：**Translation Context Contract**（P0，與 decision／registry／validation 同批）+ Translation Decision、Expression Type Registry、Validation Contract（含 **Locale Validation**）。不先寫模型 prompt。
- 所有 translation actor **只吃 `TranslationContext`**，禁止 decision 路徑僅有 `{ src, dst }`。
- **Locale Resolution** 必須在 Expression Analysis **之前**（見 04 案例）。

架構與 ERA 對照見 [`01-architecture-and-era.md`](01-architecture-and-era.md)。目錄與 SoT 見 [`02-sot-contracts-and-layout.md`](02-sot-contracts-and-layout.md)。NVP 接線見 [`03-nvp-and-adapters.md`](03-nvp-and-adapters.md)。印尼語稱謂 dogfood 見 [`04-dogfood-case-address-title-id-ID.md`](04-dogfood-case-address-title-id-ID.md)。

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
| Q3 缺 context | 無 speaker／intent／scene 等必要 context 時 **禁止 `finality: accepted`**；須 `needs_review` 或 `unresolved` + `blocking_reasons` |
| Q4 TDR 存放 | v0 只定 **schema**；專案 artifact 路徑由 consumer profile 定 |

## Phases

### Phase 0 — Plan freeze（本 commit）

- [x] 計畫落盤 `plans/active/2026-09-22-1000-translation-decision-workflow/`
- [x] 第一個 real-data case：陈小姐 → id-ID（[`04`](04-dogfood-case-address-title-id-ID.md) + [`05-example-address-title-chen-xiaojie-id.yaml`](05-example-address-title-chen-xiaojie-id.yaml)）
- [ ] Stakeholder 確認 status：`draft` → `in-progress`
- [ ] 確認不做：runtime 投影、route 註冊、模型 prompt、工具實作（Phase 0 邊界）

### Phase 1 — Contracts + Registry（doc-only）

- [ ] 建立 `workflow/translation/` 目錄骨架
- [ ] 落地 SoT YAML（**translation-context**、decision、expression-types、validation + locale_consistency）
- [ ] README + `execution-flow.md`（Locale Resolution → 六層）
- [ ] `examples/`：自 plan `05-example-*` 複製 **address-title-chen-xiaojie-id** + slang、proverb、dialect 各一
- [ ] `registry/`：`name_with_address_title`、`address_title`（locale_aware）
- [ ] `locale_consistency` regression 對照 example `regression_fixtures`

### Phase 2 — Subtitle adapter + NVP link

- [ ] `adapters/subtitle.yaml` ↔ `caption-locale-pack` `content_gate`
- [ ] NVP [`captions-and-locales.md`](../../workflow/narrative-video-production/captions-and-locales.md) inbound pointer（linked update）
- [ ] 明文化 `semantic_context` 為 Translation Context 上游

### Phase 3 — Knowledge skeleton + dogfood

- [ ] `knowledge/translation/` 骨架（含 `locale/title-mapping` 種子）
- [ ] episode 14 案例正式 dogfood run + `evidence/`（可引用既有 segment 紀錄）

### Phase 4 — 可選（不擋 v0）

- [ ] document／UI／dubbing adapter
- [ ] `route.workflow.translation` + runtime 投影（需 Phase 3 證據）

## Out of Scope（v0）

- 翻譯 API／prompt／模型選型寫進 canonical
- 單一 quality score
- 合併 NVP timing／layout 進 translation content gate
- 完整方言／成語知識庫
- 自動 orchestrator

## Success Criteria（v0 完成）

- 三 SoT 可獨立閱讀，不依賴任何模型名
- ≥3 expression examples 填滿 TDR 欄位（含 **陈小姐 id-ID** + locale regression fixtures）
- README 一句話說清 Constraint／Selection／Finality
- subtitle adapter 說明餵 NVP content_gate，不碰 timing／layout
- 未註冊 route、未 runtime 投影

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
