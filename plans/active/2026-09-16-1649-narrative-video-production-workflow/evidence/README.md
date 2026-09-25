# Plan evidence — narrative-video-production

## 引用規則

引用寫 `evidence/<file>.md` 或 markdown 連結。禁止用檔案內絕對行號定位。去敏：專案名、絕對路徑、原始媒體檔名、host、金鑰留在 `<PROJECT_ROOT>`。

## Run 索引

| Run ID | 檔案 | 狀態 | 摘要 |
| --- | --- | --- | --- |
| 2026-09-16-dialogue-semantic-context | [2026-09-16-dialogue-semantic-context.md](2026-09-16-dialogue-semantic-context.md) | observation | 台詞≠查找；dialogue optional 已落地 |
| 2026-09-17-shot-unit-semantic-context | [2026-09-17-shot-unit-semantic-context.md](2026-09-17-shot-unit-semantic-context.md) | observation | 單元 = text + semantic_context；含 action／visual；不擴 contract |
| 2026-09-17-series-cast-canonicalization | [2026-09-17-series-cast-canonicalization.md](2026-09-17-series-cast-canonicalization.md) | observation | 角色命名 SoT 候選；不改 Phase 2 workflow |
| 2026-09-17-identity-precedes-naming | [2026-09-17-identity-precedes-naming.md](2026-09-17-identity-precedes-naming.md) | observation | 身份先於名稱；link 不覆寫歷史；series_cast 是已解析表 |
| 2026-09-17-material-fact-extraction | [2026-09-17-material-fact-extraction.md](2026-09-17-material-fact-extraction.md) | observation | observable 機械證據先於語義；分析器當 evidence candidates |
| 2026-09-17-visual-text-evidence | [2026-09-17-visual-text-evidence.md](2026-09-17-visual-text-evidence.md) | observation | visual_text_evidence SoT；normalized box + persistence；role 只 candidate |
| 2026-09-22-watermark-exclusion-is-projection | [2026-09-22-watermark-exclusion-is-projection.md](2026-09-22-watermark-exclusion-is-projection.md) | contract_gap candidate | watermark 判定對；刪除／crop／contains 切掉有效字；role-aware projection |
| 2026-09-22-text-group-preserve-variants | [2026-09-22-text-group-preserve-variants.md](2026-09-22-text-group-preserve-variants.md) | contract_gap candidate | 高相似≠同一段；group 保留 variants；禁止 dst 反推 source |
| 2026-09-22-subtitle-layout-engine | [2026-09-22-subtitle-layout-engine.md](2026-09-22-subtitle-layout-engine.md) | observation | content≠layout；glyph／safe area solver；V0 先於 face obstruction |
| 2026-09-22-typography-layout-profile | [2026-09-22-typography-layout-profile.md](2026-09-22-typography-layout-profile.md) | observation | font 是 range＋canvas scale＋overflow.order；禁止字數估寬 |
| 2026-09-22-layout-review-loop | [2026-09-22-layout-review-loop.md](2026-09-22-layout-review-loop.md) | observation | AI 只出 adjustment；engine 執行；budget；跨集才升 profile |
| 2026-09-22-speech-timing-authority | [2026-09-22-speech-timing-authority.md](2026-09-22-speech-timing-authority.md) | contract | 自製口播 cue 時軸 = speech artifact；TTS adapter；兩 loop 分開 |
| 2026-09-22-max-lines-is-bound | [2026-09-22-max-lines-is-bound.md](2026-09-22-max-lines-is-bound.md) | contract | max_lines 上限不是 target；minimize_lines；記 selected lines |
| 2026-09-24-font-size-hard-bounds | [2026-09-24-font-size-hard-bounds.md](2026-09-24-font-size-hard-bounds.md) | contract | font min／max／step 硬閘；禁止縮過下限塞進畫面 |
| 2026-09-25-semantic-safe-wrap-uniform-typography | [2026-09-25-semantic-safe-wrap-uniform-typography.md](2026-09-25-semantic-safe-wrap-uniform-typography.md) | contract | wrap 不截詞／不丟字；同 cue 上下行 typography 統一 |
| 2026-09-25-caption-composition-semantic-breaks | [2026-09-25-caption-composition-semantic-breaks.md](2026-09-25-caption-composition-semantic-breaks.md) | contract | 語意先於 fit；只從 scored 斷點換行；禁止字數均分 |
| 2026-09-21-ocr-visual-style | [2026-09-21-ocr-visual-style.md](2026-09-21-ocr-visual-style.md) | observation | visual_style 機械量測；非 subtitle_color；非新 detector |
| 2026-09-17-editorial-vs-narrative-transition | [2026-09-17-editorial-vs-narrative-transition.md](2026-09-17-editorial-vs-narrative-transition.md) | observation | 畫面切了 ≠ 劇情轉場；shot ≠ scene |
| 2026-09-17-face-as-candidate-evidence | [2026-09-17-face-as-candidate-evidence.md](2026-09-17-face-as-candidate-evidence.md) | observation | Face Track 可被 evidence link 引用；非身份判定器；不接辨識 |
| 2026-09-18-evidence-refinement | [2026-09-18-evidence-refinement.md](2026-09-18-evidence-refinement.md) | observation | 獨立 refinement loop；OCR 幾何；作品級 policy；script 是 consumer |
| 2026-09-18-mechanical-visual-text-probe | [2026-09-18-mechanical-visual-text-probe.md](2026-09-18-mechanical-visual-text-probe.md) | observation | 探針可 fallback；LLM 不改 crop；清楚角色用機械 candidate |
| 2026-09-18-voice-speaker-evidence | [2026-09-18-voice-speaker-evidence.md](2026-09-18-voice-speaker-evidence.md) | observation | Voice 與 ASR 文字分開；speaker ≠ character；共現不是等同 |
| 2026-09-18-story-evidence-vs-dialogue | [2026-09-18-story-evidence-vs-dialogue.md](2026-09-18-story-evidence-vs-dialogue.md) | observation | 劇情=state change；有對話≠劇情；低相關 archive 不刪 |
| 2026-09-18-evidence-unit | [2026-09-18-evidence-unit.md](2026-09-18-evidence-unit.md) | observation | 凍結 observable 擴張；先 evidence_unit 再 event；不改 schema |
| 2026-09-18-real-run-promotion-gaps | [2026-09-18-real-run-promotion-gaps.md](2026-09-18-real-run-promotion-gaps.md) | real-partial | 首份真實 source-analysis；observable／unit 可用，promotion／identity／state fail |
| 2026-09-18-episode-vs-knowledge-accumulation | [2026-09-18-episode-vs-knowledge-accumulation.md](2026-09-18-episode-vs-knowledge-accumulation.md) | observation | 本集觀察 ≠ 知識寫入；Learning Inbox 先於 Knowledge／Mechanical Registry |
| 2026-09-18-narrative-representation-gap | [2026-09-18-narrative-representation-gap.md](2026-09-18-narrative-representation-gap.md) | real-review | resolved text 未被 narrative 消費；event 過度 dialogue-centric；window／relation 缺失 |
| 2026-09-21-spoken-vs-subtitle-reconstruction | [2026-09-21-spoken-vs-subtitle-reconstruction.md](2026-09-21-spoken-vs-subtitle-reconstruction.md) | observation | spoken ≠ subtitle；禁止 OCR 優先；LLM 不改 raw ASR；duration／word time 當 constraint |
| 2026-09-21-phonetic-text-reconstruction | [2026-09-21-phonetic-text-reconstruction.md](2026-09-21-phonetic-text-reconstruction.md) | observation | 音→字第二次 decoding；syllable 優於字數；LLM 最後；sanitization 走 inbox |
| 2026-09-21-sanitization-anomaly-audit | [2026-09-21-sanitization-anomaly-audit.md](2026-09-21-sanitization-anomaly-audit.md) | observation | OCR 不進同音；兩條 chain；alert 非字典；前警覺＋整集 audit |
| — | — | EDR chain pending | 尚未走完 matching → EDR → locale → QC → publish／outcome |
| 2026-09-22-phase-3-chain-station-ledger | [2026-09-22-phase-3-chain-station-ledger.md](2026-09-22-phase-3-chain-station-ledger.md) | station ledger | 主鏈卡在 bible／catalog／matching／EDR（`data_insufficient`）；子鏈 promotion／watermark 已分類 |
