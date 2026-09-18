---
id: 2026-09-16-1649-narrative-video-production-workflow
plan_kind: main
status: in-progress
owner: workflow
owner_layer: workflow
created: 2026-09-16
parent: null
---

# Narrative Video Production Workflow（`workflow/narrative-video-production/`）

**Status**: in-progress — Phase 1 PASS、Phase 2 PASS（`34f778d4` + `8757a578`）。Phase 3 dogfood **進行中**：真實片子驗證既有契約，不改架構。**不**做工具、**不**接 runtime。

**Glossary Impact**: yes — 另增候選 `selection_policy`、`feasible_set`、`semantic_context`、`series_cast`／`call_name`、`evidence_link`／`face_track`、`evidence_policy`／`visual_text_evidence`／`normalized_box`／`mechanical_probe`、`voice_evidence`／`speaker_id`、`story_state`／`narrative_role`。字幕 [`01-captions-and-locales.md`](01-captions-and-locales.md)；素材 [`02-source-bible.md`](02-source-bible.md)；invariant [`03-architecture-invariants.md`](03-architecture-invariants.md)；dogfood [`04-phase-3-dogfood.md`](04-phase-3-dogfood.md)；cast [`05-series-cast-canonicalization.md`](05-series-cast-canonicalization.md)；單元 [`06-shot-unit-semantic-context.md`](06-shot-unit-semantic-context.md)；事實層 [`07-material-fact-extraction.md`](07-material-fact-extraction.md)；身份 [`08-identity-precedes-naming.md`](08-identity-precedes-naming.md)；畫面文字 [`09-visual-text-evidence.md`](09-visual-text-evidence.md)；轉場 [`10-editorial-vs-narrative-transition.md`](10-editorial-vs-narrative-transition.md)；人臉證據 [`11-face-as-candidate-evidence.md`](11-face-as-candidate-evidence.md)；證據收斂 [`12-evidence-refinement.md`](12-evidence-refinement.md)；機械探針 [`13-mechanical-visual-text-probe.md`](13-mechanical-visual-text-probe.md)；聲線 [`14-voice-speaker-evidence.md`](14-voice-speaker-evidence.md)；劇情證據 [`15-story-evidence-vs-dialogue.md`](15-story-evidence-vs-dialogue.md)。Phase 5 前不登記 glossary。

## 一句話目標

為「AI 產片」建立 governed creative execution 總綱：共享 bible／catalog、明示 Constraint 與 Selection、EDR 當決策 SoT、獨立驗證後才 publish-ready、outcome 當 evidence——而不是先做工具。

## Decision Rationale

### Problem & Why Now

外部參考包（AI-Editor-Colleague code-only，2026-09-16）是一套可跑的產片程式（短劇分屏量產、電影解說爆款、參考成片配音）。它證明管線存在，但缺這些我們要的東西：

1. **成片沒有獨立、可審查的剪輯紀錄**：任務狀態與 ffmpeg 輸出不能代替「這部片為什麼這樣剪」。
2. **腳本模板沒有當成一等公民**：模型提示裡有 hook／arc_role，但沒有可盤點的模板目錄，無法回答「這部片是破題法還是英雄旅程」。
3. **沒有發布後回饋**：有「已處理檔名跳過」的作業庫，不是用觀看／完播來驗證腳本。
4. **字幕與多語卻已有可抽象的正確性規則**（CPS／cue 窗、語系斷句、源語殘留檢查）。這部分**值得進總綱**，但不能把參考包的 TTS／翻譯實作成為 workflow。
5. **沒有可查找的素材層**：人物／集數散在長片裡；更缺「抽出的片段 + 畫面 tags + 時長」，匹配腳本無法用同一套文字庫找回材料，也無法依長短挑最適合的 clip。

使用者本輪只要 workflow 總綱，作為日後若做工具時的契約；不是實作 ShortDramaTool 的 fork。

### Decision

新增 `workflow/narrative-video-production/`，domain = **敘事向影音從 brief 到發布結果的生產與驗收**。

鎖定：

- **不做工具**。Launcher、ffmpeg 編排、TTS、模型 gateway、管理後台、資料庫部署全部 out of scope。未來工具必須 **consume** 本 workflow 的紀錄契約，不得反向定義流程。
- 每部片子必有一份 **Edit Decision Record（EDR）**：同一份資料，YAML（或同等結構化）給機器，Markdown 投影給人看。沒有 EDR 不得宣稱可發布。
- 每部片子必填 **`narrative_template_id`**（可多標、必有主標）。模板目錄是 workflow 的 catalog，不是某家模型的 prompt 全文。
- **字幕與多語是 EDR 的 locale pack，不是後期美工。** 對白來源（腳本／ASR／OCR）、目標語、燒錄或 sidecar、閱讀速度與折行族都要能審查。規則摘要見 [`01-captions-and-locales.md`](01-captions-and-locales.md)。供應商聲音與翻譯模型只當未來 adapter。
- **Source bible + clip catalog 適合進 workflow。** 人物／集數穩定 id；擷取片段寫入畫面描述與 tags。查找必須回傳既有 `clip_id`。**Constraint 產生可行候選集；Selection 是明示 policy，時長接近度不得默認等於「最好」。** 見 [`02-source-bible.md`](02-source-bible.md)、[`03-architecture-invariants.md`](03-architecture-invariants.md)。
- **EDR 是決策 SoT；mp4 是產出。** Assemble 必須對得上 EDR；`publish-ready` 需要獨立於 producer 的 fresh verification。
- 發布後走 **publish-outcome loop**：outcome 是 **evidence status**（目前證據是否支持模板假設），不是「已證明模板為真」。
- 字幕三閘獨立：content ≠ timing ≠ layout。見 locale companion。
- 參考包三種 product mode 降為 **profile 候選**（Q4 dogfood，不擋 Phase 1）。
- 首輪不建 `analysis/`、`intelligence/` 空殼；不註冊 route；YAML **不**投影。至少一次真實 EDR dogfood 後才有資格當 runtime candidate。

### Domain Boundary

| 任務 | 本 workflow | 其他 |
| --- | --- | --- |
| Brief、源片 bible、clip catalog（畫面 tags／時長）、匹配腳本、模板、EDR、字幕語系、流量回寫 | Primary | — |
| 可驅動 3D 角色／VRM 資產 | 不適用 | `3d-character-production`（該 plan 已把純影片生成劃出） |
| 產片器、剪輯 GUI、ASR/TTS/翻譯模型、帳號與金鑰 | 只定 locale pack 欄位與閘 | `software-delivery`（若真的要寫程式） |
| 訓練／fine-tune 影片模型 | 不適用 | ML／software |
| 把某課程或某 zip 當 canonical | 禁止 | 只允許去敏後的 retain/adapt/drop |

### Alternatives Considered

- **A. 把參考包當 workflow**：reject。程式、金鑰、主機、提示詞會鎖死工具與版權。
- **B. 擴充 `3d-character-production`**：reject。那邊驗的是可驅動角色資產；影片敘事與流量回饋會污染 identity gate。
- **C. 先做產片工具再補流程**：reject。使用者明確只要總綱；沒有 EDR／模板／outcome 契約，工具只會再產無法比較的 mp4。
- **D. `narrative-video-production` workflow 總綱（accept）**。

### Why Not an ADR Yet

尚未有第二個真實成片＋流量窗口。先 plan + workflow docs；模板 taxonomy 與 EDR schema 穩定且被工具消費後，再評估 ADR。

### ADR Promotion Criteria（completed 時評估）

- [ ] 至少 5 部片子填完整 EDR（含主模板）。
- [ ] 至少 2 個模板各有 ≥ 2 部 outcome 窗口，能比較 **evidence status**（非宣稱模板已證成）。
- [ ] 沒有把特定產片軟體寫進 canonical 步驟。
- [ ] Open Questions 全解或有 deferred owner。
- [ ] 沒有更輕的 documentation-only 頁面就足夠。

### Consequences

#### 正面

- 人工與未來工具讀同一份剪輯真相。
- 腳本改進有模板粒度，而不是「再讓模型寫一次」。
- 人物／集數可查找；抽出片段有畫面 tags 與時長，匹配腳本用同一文字庫選材。
- 流量用來驗假設，不是拿來當虛榮數字。

#### 負面

- 每片多一份 EDR／腳本成本；bible 應跨片共用以免複製。
- 模板目錄初期會吵（分類重疊）。

#### 風險

- **Workflow inflation**（Gen 4 Watch-Out #2）：把 ffmpeg 步驟寫進 execution-flow。緩解：流程只到紀錄與閘；渲染細節留給未來 adapter。
- **Telemetry explosion**（Watch-Out #4）：接所有平台 API。緩解：outcome 只收最小欄位；人工貼匯出也算。
- **量化幻覺／把 Selection 機械化**：時長最近自動當選。緩解：feasible set + 明示 `selection.policy`（invariant 5）。
- **參考包污染**：把私有路徑、金鑰、提示詞貼進 repo。緩解：zip 留在 `.agent-goals/`，本庫只留去敏盤點。
- **Bible 變字幕全集／clip 變印象標籤**：全文當庫或沒有入出點。緩解：濃縮實體 + 有時長的 clip 列；全文另存。

## 參考包盤點（retain / adapt / drop）

來源：code-only 包目錄盤點（不 commit zip；不引用私有主機／.env）。

| 來源概念 | 處置 | 進 workflow 的抽象 |
| --- | --- | --- |
| `GroupPlan`（title、hook、segments、arc_role、prev_bridge、next_tease） | **adapt** | EDR 的鏡頭／鉤子／橋段欄位 |
| `DramaBible` + `ReelSpec` | **adapt** | bible 身份 id；另建 **clip catalog**（擷取窗、畫面描述、tags、時長） |
| heuristic reel roles（cold open hook、rising conflict、identity clash） | **adapt** | 模板／beat 種子，不是唯一分類 |
| 電影模式：merge → ASR → 規劃多條爆款 → TTS → title/hook/commentary → finalize | **adapt 為階段名** | 只保留階段與產物，不保留實作 |
| 短劇左右分屏 1.75× 量產 | **profile 候選** | 非預設敘事路徑 |
| 參考成片配音／翻譯 | **profile 候選** | 與「原創敘事」分開 |
| `TITLE` 解說文案庫、Gemini prompt 全文 | **drop 原文** | 只承認「解說／recap」是一種模板族 |
| `JobState.rebuild_flags`（recut / rewrite_script / redo_bible） | **adapt** | EDR 的 mutation／重做原因 |
| 字幕 CPS／min-max cue、語系 `BreakPolicy`／`ScriptPattern`、多語目錄、源語殘留檢查、`text_zh`+目標語並行 | **adapt** | locale caption pack；見 companion |
| ASR vs OCR 硬字幕 vs 腳本對白 | **adapt** | `text_origin` 分源，禁止混成一條無標記對白 |
| Azure／Edge 聲音 id、離線 Qwen 下載與 GPU 細節 | **drop** | adapter；canonical 不用供應商字串 |
| 已處理檔名跳過庫 | **drop as outcome** | 那是作業去重，不是流量驗證 |
| launcher、ps1、exe、site-packages、金鑰、部署 | **drop** | 工具層 |

## Proposed Workflow Shape

```text
workflow/narrative-video-production/
  README.md
  execution-flow.md
  intake.md
  source-bible.md                 # 集數／人物等查找目錄（Phase 2）
  material-clip-catalog.md        # 擷取片段、畫面 tags、時長 bands（Phase 2）
  narrative-template-catalog.md   # 破題／英雄旅程／解說爆款… id 與適用條件
  matching-script.md              # AI 匹配腳本：必須引用 bible id
  script-and-shot-list.md
  edit-decision-record.md         # EDR 人機雙讀契約
  captions-and-locales.md         # 字幕時間窗、折行族、多語 pack（Phase 2）
  assemble-and-qc.md              # 對 EDR 的符合性，不是教 ffmpeg
  publish-outcome.md              # 流量窗口 → 腳本／模板結論
  artifact-gates.md
  evidence-refinement.md          # 候選獨立契約；parser 不吸收；Phase 3 不建檔
  records/                        # YAML SoT（EDR schema、template enum）
  profiles/README.md              # 首輪空；recap / original / split-promo 後補
```

本輪 plan **不**建立 `workflow/narrative-video-production/` 正文，直到 Phase 1 invariant 凍結（本檔 + 03）被接受且進入 Phase 2。

Phase 2 已寫入：[`workflow/narrative-video-production/`](../../workflow/narrative-video-production/README.md)。

## 建議總綱（agent 執行順序）

```text
0. Frame                 敘事影音成品（非 3D 資產、非寫產片器）
1. Intake / Brief lock   平台、時長、受眾、一句话承諾、禁止項
1b. Source bible         分析源片：集數／人物等穩定 id（多成片共用）
1c. Clip catalog         擷取片段；畫面描述 + tags 寫入統一文字庫；標 duration
2. Template select       主 narrative_template_id + 可選次模板
3. Matching script       Constraints → 可行候選集 → 明示 selection policy → selected（留下 rationale）
4. EDR open              決策紀錄 SoT；shot 對齊腳本／bible／clip_id
5. Continuity            角色／場景／風格 lock（需要時）
6. Acquisition           生成或實拍或剪輯——策略可換；結果回寫 EDR
7. Assemble vs EDR       成片必須對得上 EDR；mp4 不得反過來當決策真相
7b. Locale packs         content／timing／layout 三閘分開
8. Publish QC            平台規格；**publish-ready 需 fresh verification**
9. Outcome window        evidence status 回寫模板假設（非 truth）
```

成熟度：`exploration`／`cut-ready`（EDR 與成片對得上，可自驗推進）／`publish-ready`（**獨立 verifier** + QC）／`outcome-scored`（窗口填完，**不是**流量好）。Completion claim 的 `publish-ready` 不得由 producer 自簽。

### EDR 最小欄位（契約草案，未凍結）

| 欄位 | 用途 |
| --- | --- |
| `film_id` / `title` / `profile` / `bible_id` | 身份；成片連到共用素材聖經 |
| `matching_script_id` | 本片所依據的格式化匹配腳本 |
| `narrative_template_id`（主）+ `secondary_template_ids` | 這部片靠哪個模板產生 |
| `template_beats[]` | 模板段落 ↔ `shot_id` |
| `shots[]` | 入出點、`selected_clip_id`、feasible candidates、`selection.policy`／rationale |
| `hook` | 開場承諾是否在前 N 秒兌現 |
| `mutations[]` | recut / rewrite_script / redo_bible 與原因 |
| `qc` | 對 EDR 的符合性；producer 自驗 ≠ completion |
| `publish` | 平台、識別碼、時間、`publish_locales[]` |
| `locale_packs[]` | 每語 + **content／timing／layout** 三閘 |
| `outcome` | 窗口與最小欄位；`evidence_status`：`supports`／`contradicts`／`insufficient_sample` |

禁止：無觀察的綜合品質分；把平台原始 dump 整包提交進本知識庫。

### 模板目錄 v0（experimental；kind slot 已凍結）

七個 id 可留作種子，**不假設同一 taxonomy level**。每個預留 `template_kind: structure | mechanism | format`（例如 `hero_journey`≈structure，`cold_open_hook`≈mechanism，`recap_explained`≈format）。Phase 3 前不拆 catalog。選定：一個主模板；次模板只解釋局部。

### 流量如何當 evidence

- 單位是 **（模板 × 窗口）**。
- 同模板至少 2 部、同一平台窗口，才允許 `supports`／`contradicts`（**目前證據是否支持該假設**，不是已證明為真）。
- 單部爆款 → `insufficient_sample`。
- 去敏後回寫模板 evidence 索引。

## Runtime Execution Path

本 plan **不接入 runtime**。Narrative domain correctness 必須先經過至少一次真實 EDR dogfood，才有資格成為 runtime candidate。未 wire 前不得聲稱已 runtime integration。

預定接入（未開始）：`route.workflow.narrative-video-production`；consumer = `DetectWorkflows` + `workflowPrimarySourceGate`。未 wire 前不得聲稱已 runtime integration。

### Per-surface consumer 表

本輪無 generated surface。`manual_activation` 不適用（尚未建 route）。

### Deferred Runtime Projection

不建立 `runtime/*.yaml`。不 project。預定：workflow YAML 契約在第一次 dogfood 後再評估 `runtime_projection.enabled`。

## Pre-build Interrogation

- **Goal**：可重跑的敘事產片總綱 + EDR + 模板 + 流量回寫。
- **Scope**：workflow 文件與 records schema。
- **Non-goals**：產片工具、平台 API 整合、複製參考包、commit zip。
- **Acceptance**：Phase 1 = 凍結 03 的十條 invariant；欄位細節允許 dogfood 修正。
- **Duplication**：與 3D workflow、software-delivery、未來 ML 路徑已分界。
- **Assumption**：流量可由人工貼最小欄位；不依賴即時爬蟲。

## Open Questions

- [x] **Q0** domain → `narrative-video-production`
- [x] **Q-tools** → 不做工具
- [x] **Q1** 模板語意 → **resolved（架構）**：模板是一級 artifact；v0 七 id 可留；必有 `template_kind` slot；不拆 taxonomy、不發明匿名模板。七個 id 的去留 = experimental。
- [x] **Q2** 成熟度 → **resolved（架構）**：自驗可推進到 `cut-ready`；宣稱 `publish-ready` 必須 fresh verification。Outcome 不是 completion 必要條件。
- [x] **Q3** EDR SoT → **resolved（架構）**：結構化檔（YAML 或同等）是 EDR canonical；Markdown 是人讀投影。
- [ ] **Q4** first profile → **dogfood（C）**：不擋 Phase 1。
- [x] **Q5** outcome 欄位 → **v0（B）**：先用 impressions／views／avg_watch_pct／engagements／窗口天數；dogfood 可修，非永久 schema。
- [x] **Q-captions** → locale pack；content／timing／layout 三閘。
- [ ] **Q6** 首輪語 → **dogfood（C）**：建議源語+一目標語；不擋 Phase 1。
- [x] **Q-bible** → bible + clip catalog + 查找必須既有 clip_id。
- [x] **Q7** 所有權 → **resolved（架構）**：一部源作品一份 bible／一份 catalog；clip 增量進同一 catalog。
- [x] **Q8** 匹配腳本檔形 → **resolved（架構）**：與 Q3 相同（結構化 SoT + 人讀投影）。
- [x] **Q9** 受控 tags → **resolved（架構）**：受控詞表 + 能連 entity_id；禁止每次發明同義詞當新 tag。
- [ ] **Q10** duration band 秒數 → **dogfood（C）**：v0 可用 hook／beat／hold 分檔；數字可修。Band 存在本身屬約束，秒數不凍死。
- [x] **Q11** Selection Responsibility → **resolved（架構）**：Constraints → feasible set → 明示 selection policy → selected。時長接近度可作 criterion，不得默認等於唯一「最好」。見 03。

## Phase 0 — Architecture Compatibility Preflight

### Phase 0.0 — Open Questions 核對（公版，必填）

- [x] 已讀本 plan §Open Questions 全部條目
- [x] 對每條標記 `resolved` / `still-open` / `deferred`
- [x] `resolved` 的條目已同步勾選 / 附註於 §Open Questions
- [x] 盤點新發現已加入（Q11 Selection、invariant companion 03、Q 分級）

| Open Question | 處置 | 證據 / 原因 |
|---|---|---|
| Q0 / Q-tools / Q-captions / Q-bible | resolved | 使用者先前凍結 |
| Q1 模板語意 + kind | resolved | review：一級 artifact；不拆 taxonomy |
| Q2 成熟度 | resolved | cut-ready 可自驗；publish-ready 要 fresh verifier |
| Q3／Q8 SoT 形狀 | resolved | 結構化 canonical + 人讀投影 |
| Q4 profile | still-open / dogfood | 非架構 invariant |
| Q5 outcome 欄位 | v0 | 五欄可修 |
| Q6 語系 | still-open / dogfood | 非 completion 核心 |
| Q7 所有權 | resolved | 每源作品一份 bible+catalog |
| Q9 受控 tags | resolved | 受控詞表 + entity_id |
| Q10 band 秒數 | still-open / dogfood | 分檔要有；數字可修 |
| Q11 Selection | resolved | feasible set + 明示 policy |

### Phase 0.1 — Candidate / layer

| 項 | 結果 |
| --- | --- |
| Candidate | `workflow/narrative-video-production/**`；本 `_plan.md`；`workflow/README.md`、`workflow-routing.md`（**未**改 routing-registry） |
| Source of truth | workflow markdown + `records/` YAML；不是參考包程式 |
| Layer | workflow；非 runtime、非 ai-tools |
| Compiler | 不適用（本 round 無 projection） |
| Linked updates now | `plans/README.md`；`workflow/README.md`；`workflow-routing.md` 手動入口。registry **刻意不改**（Invariant 10） |
| Conflicts | 無雙 SoT。注意勿把 zip／.env 加進 git |
| Decision | **Phase 2**：寫 workflow contract；十條 invariant 落在各檔 gate |

## Phase 1 — 凍結 domain invariants（非凍死全部 schema）

目標：**Freeze the domain invariants, not every domain detail.** Architecture = ready；schema = experimental。

完成條件：

- [x] [`03-architecture-invariants.md`](03-architecture-invariants.md) 十條 invariant 寫入
- [x] Selection：Constraints → feasible set → 明示 policy（Q11）
- [x] EDR SoT／fresh verification／outcome=evidence／locale 三閘／catalog 只回 clip_id
- [x] 使用者確認 03 可作為 Phase 2 寫 workflow 的契約（2026-09-16：Phase 1 PASS → Phase 2 proceed）

不擋 Phase 1 的：Q4、Q6、Q10。Q5 以 v0 帶進 Phase 2。

## Phase 2 — 寫 workflow 總綱（仍無工具、無 route）

**PASS**（2026-09-16）。落點 commit：`34f778d4`、`8757a578`。

完成條件：Proposed Shape 檔存在；artifact-gates 寫明 publish-ready 的 completion authority 獨立於 producer；EDR／matching script／catalog 示範可人工填（去敏）；`workflow/README.md` 列入並註明未註冊 route。

- [x] Proposed Shape 檔存在於 `workflow/narrative-video-production/`
- [x] `artifact-gates.md`：publish-ready 獨立於 producer
- [x] 去敏示範 [`records/examples/sanitized-matching-and-edr.yaml`](../../workflow/narrative-video-production/records/examples/sanitized-matching-and-edr.yaml)
- [x] `workflow/README.md` 列入；route 未註冊
- [x] 十條 invariant 在 README 落點表 + 各 contract／gate（不只引用 03）
- [x] 使用者判定 Phase 2 PASS → 進 Phase 3（不補 Q12/Q13、不回頭改 runtime）

## Phase 3 — 一份真實 EDR dogfood

協議：[`04-phase-3-dogfood.md`](04-phase-3-dogfood.md)。證據索引：[`evidence/README.md`](evidence/README.md)。

- [ ] 外部專案一部真實片子走完觀察鏈（outcome 可 `insufficient_sample`）
- [ ] 本庫 `evidence/` 去敏 run：每站 pass 或卡住分類（`contract_gap`／`data_insufficient`／`adapter_only`／`design_error`）
- [ ] 卡住不自動加欄位／加 phase（dialogue optional 維持；identity／series_cast／observable 分析器／Face Recognition／聲紋產品／權重模型 **不**寫進本 phase workflow；Face／Voice 只留 track 掛點；refinement 只留觀察契約）
- [ ] 虛構 YAML 示範不算本 phase
- [ ] 仍無 route、無 runtime projection

完成條件：外部專案產出一部片子的 EDR；本庫只收去敏 scenario。成功 = 決策鏈可驗證；失敗 = 真實 contract gap（都算有價值）。

## Phase 4 — 視需要才考慮 route

Entry：Phase 2+3 完成且 activation 反例寫好（裸「AI 影片」不得鎖路）。

## Stakeholder 同意項目

- [x] Domain 名稱 `narrative-video-production`
- [x] 本輪不做工具
- [x] 每片要有可解析、人可讀的剪輯紀錄
- [x] 每片要能追溯敘事模板
- [x] 流量用來驗證腳本／模板，不是作業去重
- [x] 字幕與多語進總綱（locale pack；非翻譯產品）
- [x] 源片 bible + clip catalog + 查找既有 clip_id
- [x] Constraint ≠ Selection；明示 selection policy
- [ ] `dialogue.semantic_context`：optional 已落地；**shot unit（dialogue／action／visual）** 升格等真實計數
- [ ] Series Cast Canonicalization：觀察中；**已解析表**，不是發現層；升格前不改 Phase 2 workflow
- [ ] Identity precedes naming：觀察中；entity 可 unnamed；link 不覆寫歷史
- [ ] Face as candidate evidence：觀察中；`face_track` 可被 evidence link 引用；Face ≠ identity 判定器；升格前不改 workflow、不接辨識模型
- [ ] Voice / speaker evidence：觀察中；ASR 只管 transcript；`speaker_id` ≠ `character_id`；共現連結不是身份等同；升格前不改 workflow
- [ ] Story evidence vs dialogue dump：觀察中；Relevance／Event／Story State；低相關 archive 不刪；候選 invariant 11 未凍結、未進 workflow gate
- [ ] Evidence refinement loop：觀察中；獨立於 parser；作品級 `evidence_policy`；Selection policy 先於 weight model；script 是 consumer；升格前不建 workflow 檔
- [ ] Visual text evidence：觀察中；SoT 是「何時何地出現什麼字」；pixel + normalized box + persistence 為一級 metadata；`role` 只 candidate；升格前不改 workflow
- [ ] Mechanical visual-text probe：觀察中；LLM 不決定掃區；coverage 不足才 expand；改 probe 須 Observation→Validation→Promotion
- [ ] Editorial vs narrative transition：觀察中；shot ≠ scene；關係不塞進 catalog 本體
- [ ] Material fact extraction（observable layer）：觀察中；外部 evidence candidates；升格前不塞十種分析器進 workflow
- [x] publish-ready 需 fresh verification；outcome 是 evidence
- [x] Locale content／timing／layout 分閘
- [x] 確認 03 invariant 後開 Phase 2 寫 workflow
- [ ] Q4／Q6／Q10 留待 dogfood

## 完成條件

- [x] Phase 1 凍結
- [x] Phase 2 workflow 文件
- [ ] Phase 3 至少一份去敏 EDR 示範或外部 dogfood 指標
- [ ] 未把參考包、金鑰、主機寫進 reusable docs
- [ ] 未聲稱 runtime integration
- [ ] Plan Completion Closure（若宣告 completed）

## 與其他 plans 的關係

- Watch-Out：[`architecture/ai-native-cognitive-ecosystem-system.md`](../../architecture/ai-native-cognitive-ecosystem-system.md) §Watch-Out List wall 2（workflow inflation）、wall 4（telemetry）。
- [`2026-08-31-1032-3d-character-production-workflow`](../2026-08-31-1032-3d-character-production-workflow/_plan.md)：3D 資產域；本 plan 填補 media/creative 空位。Fresh verification 對齊該 plan 的 self-check ≠ completion。
- software-delivery：僅當未來真的寫產片器。
