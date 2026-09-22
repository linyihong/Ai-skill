# Phase 3 dogfood protocol

Companion to [`_plan.md`](_plan.md)。**不新增架構、不補 Q12/Q13、不接 runtime。**  
Phase 2 落點：`34f778d4`（workflow）、`8757a578`（plan 狀態）。

目標不是「成功做出一部片」，而是：一部真實片子能否把 **素材選擇、敘事結構、EDR 決策、locale、QC、publish 狀態** 留成可驗證決策鏈。

## 核心原則

拿一部真實片子，證明 Phase 2 contract 能承載真實決策。卡住時先分類，禁止直接加欄位或加 phase。

| 分類 | 含義 | 本 round 允許 |
| --- | --- | --- |
| `contract_gap` | Phase 2 欄位／閘無法表達真實決策 | 記缺口；不改 workflow 除非使用者授權 |
| `dialogue_semantic_ambiguity` | 短台詞省略主詞 | optional `dialogue.semantic_context` 已落地；升必填等計數 |
| `shot_unit_semantic_gap` | 分鏡單元（dialogue／action／visual）缺機器可讀語義 | **不是**立刻 `contract_gap`。候選：每個單元 = `text` + `semantic_context`。本 phase **不**擴 workflow。見 [`evidence/2026-09-17-shot-unit-semantic-context.md`](evidence/2026-09-17-shot-unit-semantic-context.md) |
| `identity_naming_gap` | 第一集尚無名仍要建角色 | **不是**立刻 `contract_gap`。身份先於名稱；link 不覆寫歷史。見 [`evidence/2026-09-17-identity-precedes-naming.md`](evidence/2026-09-17-identity-precedes-naming.md) |
| `character_naming_gap` | 角色多名稱／笼统詞／寫稿詞彙不受控 | **不是**立刻 `contract_gap`。候選 **Series Cast Canonicalization**（已解析表）。Phase 3 **不改** workflow。見 [`evidence/2026-09-17-series-cast-canonicalization.md`](evidence/2026-09-17-series-cast-canonicalization.md) |
| `material_fact_gap` | 素材事實層未先於語義推理 | **不是**立刻 `contract_gap`。observable 證據 → identity／context → 決策。分析器當外部 evidence candidates。見 [`evidence/2026-09-17-material-fact-extraction.md`](evidence/2026-09-17-material-fact-extraction.md) |
| `visual_text_gap` | 畫面文字未當一級素材證據，或缺幾何／持續性、parser 直接分類角色 | **不是**立刻 `contract_gap`。候選 `visual_text_evidence`（OCR 是取得方式；`role` 只 candidate）。見 [`evidence/2026-09-17-visual-text-evidence.md`](evidence/2026-09-17-visual-text-evidence.md) |
| `watermark_as_deletion` | watermark role 後刪 evidence／crop 區／contains 整條丟棄，連帶切掉有效專名 | **`contract_gap` candidate**。role ≠ filter；span-level projection；watermark STOP name learning。未改 workflow。見 [`evidence/2026-09-22-watermark-exclusion-is-projection.md`](evidence/2026-09-22-watermark-exclusion-is-projection.md) |
| `similarity_as_identity` | 同 timestamp＋高相似就 merge 成一筆並丟 variant | **`contract_gap` candidate**。group first, resolve later；僅 exact duplicate 可刪；dst 不決定 source merge。見 [`evidence/2026-09-22-text-group-preserve-variants.md`](evidence/2026-09-22-text-group-preserve-variants.md) |
| `layout_as_charcount` | 用字數切行／寫死 bottom％／LLM 排版 | **不是**立刻改 workflow。`layout_gate` 用 constraint solver＋glyph 寬＋profile；LLM 只出 break candidates。見 [`evidence/2026-09-22-subtitle-layout-engine.md`](evidence/2026-09-22-subtitle-layout-engine.md) |
| `font_as_fixed_px` | 全片寫死 `font_size: 48`，或用字數估寬 | **不是**立刻改 workflow。preferred／min／max＋canvas scale＋overflow.order；同屬 layout profile。見 [`evidence/2026-09-22-typography-layout-profile.md`](evidence/2026-09-22-typography-layout-profile.md) |
| `ai_writes_font` | AI review 直接改 profile／無限調字級 | **不是**立刻改 workflow。bounded loop；AI 只出 adjustment candidate；跨集才 promotion。見 [`evidence/2026-09-22-layout-review-loop.md`](evidence/2026-09-22-layout-review-loop.md) |
| `cue_time_guessed` | generated 口播用 AI 猜 cue 秒數，或整段 TTS 再切字幕 | **contract** 已補 Speech Unit／timing authority；TTS 仍 adapter。見 [`evidence/2026-09-22-speech-timing-authority.md`](evidence/2026-09-22-speech-timing-authority.md) |
| `ocr_color_as_label` | 用顏色直接判字幕／說話人／角色 | **不是**立刻 `contract_gap`。`visual_style` 是 observable feature；subtitle／character 只 candidate。見 [`evidence/2026-09-21-ocr-visual-style.md`](evidence/2026-09-21-ocr-visual-style.md) |
| `probe_llm_crop` | 用 LLM 決定 OCR 掃區／改全局 crop | **不是**立刻 `contract_gap`。Mechanical Probe 可 fallback；LLM 只分類歧義角色，觀察不直接 promotion。見 [`evidence/2026-09-18-mechanical-visual-text-probe.md`](evidence/2026-09-18-mechanical-visual-text-probe.md) |
| `face_identity_shortcut` | Face 被當成人名／角色判定器 | **不是**立刻 `contract_gap`。Face Track 是 observable candidate evidence；由 linking 掛 ASR／OCR／Shot。本 phase **不**接辨識模型。見 [`evidence/2026-09-17-face-as-candidate-evidence.md`](evidence/2026-09-17-face-as-candidate-evidence.md) |
| `voice_identity_shortcut` | 把 diarization／聲線當成角色判定（「ASR 判斷人物」） | **不是**立刻 `contract_gap`。Voice／Speaker Evidence 與 transcript 分開；`speaker_id` ≠ `character_id`。見 [`evidence/2026-09-18-voice-speaker-evidence.md`](evidence/2026-09-18-voice-speaker-evidence.md) |
| `dialogue_as_plot` | 有對話／ASR 摘要被當成劇情 | **不是**立刻 `contract_gap`。Relevance → Event → Story State；低相關保留不刪。見 [`evidence/2026-09-18-story-evidence-vs-dialogue.md`](evidence/2026-09-18-story-evidence-vs-dialogue.md) |
| `skip_evidence_unit` | 從 cross-modal links 直接跳劇情，或繼續加 observable detector | **不是**立刻 `contract_gap`。先 evidence_unit；凍結 observable 擴張。見 [`evidence/2026-09-18-evidence-unit.md`](evidence/2026-09-18-evidence-unit.md) |
| `promotion_gate_gap` | upstream unresolved／空 state claim 卻升 `story_evidence: accepted` | Observation-layer gate gap；`traceable` 不等於 valid。至少 resolved upstream + resolved event + valid state claim（若有）+ independent verification。見 [`evidence/2026-09-18-real-run-promotion-gaps.md`](evidence/2026-09-18-real-run-promotion-gaps.md) |
| `text_resolution_not_consumed` | 已有 ASR／subtitle alternatives 與 selection policy，但 narrative 仍讀 raw ASR | **不是**再加模型。先做 role-qualified、span-level resolution，保留 alternatives；不得全域 OCR 優先。見 [`evidence/2026-09-18-narrative-representation-gap.md`](evidence/2026-09-18-narrative-representation-gap.md) |
| `ocr_asr_pickone` | OCR 優先／ASR fallback，或把字幕當 spoken text | **不是**立刻 `contract_gap`。spoken ≠ subtitle；duration／word time 當 constraint；LLM 只出 reconstruction candidate。見 [`evidence/2026-09-21-spoken-vs-subtitle-reconstruction.md`](evidence/2026-09-21-spoken-vs-subtitle-reconstruction.md) |
| `asr_grapheme_only` | 只存 ASR 文字、讓 LLM 直接修句，或把字幕替換當凍結對照 | **不是**立刻 `contract_gap`。phonetic／syllable 第二次 decoding；LLM 最後 ranking；sanitization 走 Learning Inbox。見 [`evidence/2026-09-21-phonetic-text-reconstruction.md`](evidence/2026-09-21-phonetic-text-reconstruction.md) |
| `sanitization_auto_map` | 看到顯示詞就自動改 spoken、把 OCR 詞拿去同音展開，或只掃一次 | **不是**立刻 `contract_gap`。OCR＝subtitle observed；phonetic 另鏈；alert 不是字典。見 [`evidence/2026-09-21-sanitization-anomaly-audit.md`](evidence/2026-09-21-sanitization-anomaly-audit.md) |
| `dialogue_centric_event` | event candidate 只是 `dialogue_cluster`，沒有跨 evidence 組裝或 event relation | **不是**放寬 threshold。候選 Narrative Window → Event Assembly → Candidate Relations；本 phase 不凍 schema。見同上 |
| `mention_as_identity` | vocative／OCR mention 被直接升 canonical identity | `data_insufficient` + identity lifecycle violation。稱呼先解析 addressee；mention ≠ entity；保持 candidate／unresolved。見 [`evidence/2026-09-18-real-run-promotion-gaps.md`](evidence/2026-09-18-real-run-promotion-gaps.md) |
| `episode_as_knowledge` | 本集分析結果直接寫進 identity／knowledge／rules | **不是**立刻 `contract_gap`。Episode Evidence ≠ Learning Candidate ≠ Knowledge Store／Mechanical Registry。見 [`evidence/2026-09-18-episode-vs-knowledge-accumulation.md`](evidence/2026-09-18-episode-vs-knowledge-accumulation.md) |
| `evidence_refinement_gap` | 把 OCR／ASR 當二選一、寫死權重、或讓 script／LLM confidence 當 evidence 權威 | **不是**立刻 `contract_gap`。獨立 refinement loop：作品級 selection policy、歧義才升級 LLM、修正回寫政策。見 [`evidence/2026-09-18-evidence-refinement.md`](evidence/2026-09-18-evidence-refinement.md) |
| `transition_layer_gap` | 畫面切換被當成劇情轉場 | **不是**立刻 `contract_gap`。editorial ≠ narrative；shot ≠ scene。見 [`evidence/2026-09-17-editorial-vs-narrative-transition.md`](evidence/2026-09-17-editorial-vs-narrative-transition.md) |
| `data_insufficient` | bible／catalog／locale 還沒填夠 | 補資料，不改契約 |
| `adapter_only` | ffmpeg／TTS／模型／GUI 問題 | 留在外部工具；canonical 不吸收 |
| `design_error` | invariant 本身擋不住或互相矛盾 | 停手，回 Phase 1 護欄討論 |

## 觀察鏈（依序，不可跳過分類）

```text
real brief → source bible → clip catalog → template
  → matching script → feasible candidates → explicit selection policy
  → EDR → locale packs → QC / independent verification
  → publish-ready → outcome
```

對照檔：[`workflow/narrative-video-production/execution-flow.md`](../../workflow/narrative-video-production/execution-flow.md)。

特別確認：每個 shot 有可行集 + `selection.policy`；查找只回既有 `clip_id`；成片對 EDR 而非反推；三閘分開；`publish-ready` 有獨立 verifier。Q4／Q6／Q10 只觀察，不在本 phase 凍結。另計：semantic unit、身份／名稱、visual text 幾何、機械探針／coverage fallback、Face Track／Voice Speaker／evidence link、evidence_unit、evidence refinement loop、story relevance／state change、editorial vs narrative transition、series_cast、observable 證據消費。**本 phase 不因這些觀察擴 workflow**；**不**再加 observable detector。優先反例：閒聊、關鍵對白、跨 shot 對白、字幕／ASR 衝突、人物切鏡。

首份真實 source-analysis 子鏈見
[`evidence/2026-09-18-real-run-promotion-gaps.md`](evidence/2026-09-18-real-run-promotion-gaps.md)：
observable／evidence unit 可用；promotion、identity、state claim 未通過。
這是有效 dogfood evidence，**不是**整體 Phase 3 PASS。
下一集 identity dogfood 另驗 [`18-episode-vs-knowledge-accumulation.md`](18-episode-vs-knowledge-accumulation.md)：vocative 停在 Episode Evidence；跨集才進 Learning Inbox；無 verifier 不得寫 knowledge。
同一 run 的人工語義複核見
[`evidence/2026-09-18-narrative-representation-gap.md`](evidence/2026-09-18-narrative-representation-gap.md)：
Text Resolution 未被 consumer 使用、Event 過度 dialogue-centric、Window／
Assembly／Relation 缺失。這比 event count 更早，仍不加 Agent／detector。

主鏈每站分類見
[`evidence/2026-09-22-phase-3-chain-station-ledger.md`](evidence/2026-09-22-phase-3-chain-station-ledger.md)：
matching／EDR 仍 `data_insufficient`；不是 `design_error`。

## 本庫 vs 外部專案

- **執行與原始媒體**在 `<PROJECT_ROOT>`（非本庫）。
- **回寫本庫**只允許去敏摘要：進 [`evidence/`](evidence/README.md)。禁止片名／路徑／host／金鑰／未授權肖像。
- 虛構示範 [`sanitized-matching-and-edr.yaml`](../../workflow/narrative-video-production/records/examples/sanitized-matching-and-edr.yaml) **不算** Phase 3。

## Phase 3 PASS 最低條件

1. 外部有一部真實片子走過觀察鏈（outcome 可停在 `insufficient_sample`）。
2. 本庫 `evidence/` 有一份去敏 run：鏈上每站 `pass`／卡住分類。
3. 未註冊 route、未開 runtime projection、未把 provider 寫進 workflow。
