# Plan evidence — narrative-video-production

## 引用規則

引用寫 `evidence/<file>.md` 或 markdown 連結。禁止用檔案內絕對行號定位。去敏：專案名、絕對路徑、原始媒體檔名、host、金鑰留在 `<PROJECT_ROOT>`。

Product dogfood that touches captions／locale pack／font／layout／timing／content≠timing≠layout **must** write back here（or workflow）per the Mac project overlay rule `aiskill-plan-feedback-loop`（also linked to Translation Decision plan）. Do not leave NVP-relevant findings only in the product repo.

本輪使用者確認的修復順序與 regression matrix：
[48-evidence-retention-and-regression-matrix.md](../48-evidence-retention-and-regression-matrix.md)（驗收規劃，非實測 PASS）。

## Run 索引

| Run ID | 檔案 | 狀態 | 摘要 |
| --- | --- | --- | --- |
| 2026-10-06-attribution-precision-and-acquisition-integrity | [2026-10-06-attribution-precision-and-acquisition-integrity.md](2026-10-06-attribution-precision-and-acquisition-integrity.md) | adapter／partial validation | band／雙語／fuzzy scrub／same-speech temporal groups 局部修復；uncertain≠recovery；採集／PTS／consumer gates 仍 open |
| 2026-10-05-evidence-resolution-loss-monitor | [2026-10-05-evidence-resolution-loss-monitor.md](2026-10-05-evidence-resolution-loss-monitor.md) | contract_gap／adapter | evidence funnel 縮量監測；uncertain 保留；定向 re-probe；artifact integrity 與 versioned cache |
| 2026-10-05-source-adaptive-ocr-region-profile | [2026-10-05-source-adaptive-ocr-region-profile.md](2026-10-05-source-adaptive-ocr-region-profile.md) | contract_gap／adapter | 勿每片手調；region behavior→source_ocr_profile；sticky 空間優先；impact+fixture |
| 2026-10-05-en-dogfood-source-lang-vs-qwen | [2026-10-05-en-dogfood-source-lang-vs-qwen.md](2026-10-05-en-dogfood-source-lang-vs-qwen.md) | validation／PASS clip | gemini 避 Qwen OOM；源口播勿用 en script gate；ep16 EN 成片 1 |
| 2026-10-02-ep16-verify-duipian-uncertain-en | [2026-10-02-ep16-verify-duipian-uncertain-en.md](2026-10-02-ep16-verify-duipian-uncertain-en.md) | validation／partial | 对片抽检 PASS+ASR notes；uncertain LIKELY；EN kick |
| 2026-10-02-sticky-watermark-accepted-projection | [2026-10-02-sticky-watermark-accepted-projection.md](2026-10-02-sticky-watermark-accepted-projection.md) | adapter／contract | accepted 黏串→projection；純殘片 rejected；衝突 uncertain；非 delete |
| 2026-10-02-ep16-accepted-uncertain-audit | [2026-10-02-ep16-accepted-uncertain-audit.md](2026-10-02-ep16-accepted-uncertain-audit.md) | observation／partial fail accepted | uncertain 8 保留；accepted~半數 sticky watermark／brand；下一步 projection 非再開 OCR |
| 2026-10-02-destructive-finalization-vs-preservation | [2026-10-02-destructive-finalization-vs-preservation.md](2026-10-02-destructive-finalization-vs-preservation.md) | contract_gap／PASS preserve | cues=0＝finalization failure；uncertain 一級；Evidence≠Production；metric 具名 |
| 2026-10-02-evidence-scope-source-vs-window | [2026-10-02-evidence-scope-source-vs-window.md](2026-10-02-evidence-scope-source-vs-window.md) | contract_gap／adapter | source hardsub／cache hit ≠ window cues；window-local OCR fallback |
| 2026-10-01-subtitle-candidate-detector | [2026-10-01-subtitle-candidate-detector.md](2026-10-01-subtitle-candidate-detector.md) | contract_gap／adapter partial | scene text≠subtitle；classifier 第一刀已落地；完整 pack dogfood 仍 open |
| 2026-10-01-evidence-acquisition-loop | [2026-10-01-evidence-acquisition-loop.md](2026-10-01-evidence-acquisition-loop.md) | contract_gap | Monitor＋escalation；OCR×ASR 窗密度缺口不得當最終無對白 |
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
| 2026-09-28-font-size-layers-visual-scale | [2026-09-28-font-size-layers-visual-scale.md](2026-09-28-font-size-layers-visual-scale.md) | contract | 三層字級；跨 script 對齊 glyph 視覺高度；profile px 不凍死 |
| 2026-09-28-break-candidate-system | [2026-09-28-break-candidate-system.md](2026-09-28-break-candidate-system.md) | contract | heuristic 是候選特徵；LLM 在可行集內選擇；升格前不改 policy |
| 2026-09-28-break-candidate-schema | [2026-09-28-break-candidate-schema.md](2026-09-28-break-candidate-schema.md) | contract | Phase 1：balance_score、hard_violation、無 source；best_cut 是 fallback |
| 2026-09-28-natural-boundary-before-length | [2026-09-28-natural-boundary-before-length.md](2026-09-28-natural-boundary-before-length.md) | contract | 標點／lexical 先於 prefer_at；字數只在同一 tier 比較 |
| 2026-09-28-speech-unit-before-caption | [2026-09-28-speech-unit-before-caption.md](2026-09-28-speech-unit-before-caption.md) | contract | Speech Unit 是語音 SoT；caption 只投影；換行不改 unit |
| 2026-09-28-asr-validity-precedes-interpretation | [2026-09-28-asr-validity-precedes-interpretation.md](2026-09-28-asr-validity-precedes-interpretation.md) | contract | 先判 ASR validity；和諧是假設，不能蓋過無效 ASR |
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
| 2026-09-29-windows-console-encoding-nonlatin-captions | [2026-09-29-windows-console-encoding-nonlatin-captions.md](2026-09-29-windows-console-encoding-nonlatin-captions.md) | observed | GBK console print of Thai／non-Latin cue text must not abort UTF-8 caption jobs |
| 2026-09-29-cross-language-subtitle-spoken-alignment | [2026-09-29-cross-language-subtitle-spoken-alignment.md](2026-09-29-cross-language-subtitle-spoken-alignment.md) | contract_gap candidate | EN hardsub＋ZH spoken 被当成 conflict／OCR 勝出；Language Relation Gate 先於 Text Resolution／sanitization |
| 2026-09-29-bilingual-hardsub-regions | [2026-09-29-bilingual-hardsub-regions.md](2026-09-29-bilingual-hardsub-regions.md) | contract_gap candidate | 中英疊字硬字幕被黏成單一字串；需 text_region＋subtitle_group 再對齊 ASR |
| 2026-09-30-ocr-latin-word-boundary-loss | [2026-09-30-ocr-latin-word-boundary-loss.md](2026-09-30-ocr-latin-word-boundary-loss.md) | contract_gap candidate | Latin 單框無空格／ch 路徑去空白；需 raw＋boundary recovery，非 LLM 斷詞 |
| 2026-10-01-ocr-join-token-seam-not-cumulative-script | [2026-10-01-ocr-join-token-seam-not-cumulative-script.md](2026-10-01-ocr-join-token-seam-not-cumulative-script.md) | adapter／join | parts 已切、derived 黏：token-seam join；CJK+Latin+Latin regression |
| 2026-10-01-ocr-latin-boundary-geometry-before-lexical | [2026-10-01-ocr-latin-boundary-geometry-before-lexical.md](2026-10-01-ocr-latin-boundary-geometry-before-lexical.md) | adapter／boundary | Latin 單 part 黏串：geometry→lexical candidates；closed-class 非第一刀 |
| 2026-09-30-asr-anomaly-ocr-semantic-reconstruction | [2026-09-30-asr-anomaly-ocr-semantic-reconstruction.md](2026-09-30-asr-anomaly-ocr-semantic-reconstruction.md) | contract_gap candidate | ASR 字面怪＋OCR 跨語言語義清晰 → 需 phonetic＋semantic_candidate＋resolution_reason；非固定 OCR>ASR |
| 2026-09-30-semantic-anchor-last-lot-auction | [2026-09-30-semantic-anchor-last-lot-auction.md](2026-09-30-semantic-anchor-last-lot-auction.md) | contract extension | OCR `Last Lot` 作 semantic_anchor＋ASR「拍皮」異常 → semantic_reconstruction「拍卖品」；observed／candidate／resolved 三層 |
| 2026-09-30-dialogue-cue-projection-retains-subtitle-group | [2026-09-30-dialogue-cue-projection-retains-subtitle-group.md](2026-09-30-dialogue-cue-projection-retains-subtitle-group.md) | contract_gap confirmation | Layer-0／cue 已有 bilingual regions；locale pack 壓成 spoken_timed 丟 subtitle_group — projection 必須保留 speech＋regions＋alignment |
| 2026-09-30-caption-temporal-integrity-overlap | [2026-09-30-caption-temporal-integrity-overlap.md](2026-09-30-caption-temporal-integrity-overlap.md) | timing_gate／adapter | 成片兩套台詞互壓：Pack 乾淨但 ASS same-track 有 overlap；擴 temporal integrity；雙語 same-group 允許 |
| 2026-09-30-ocr-caption-omission-same-asr-collapse | [2026-09-30-ocr-caption-omission-same-asr-collapse.md](2026-09-30-ocr-caption-omission-same-asr-collapse.md) | projection／omission | OCR 有句、cues／成片無：同長 ASR latch + unrelated merge；要 Timeline IR＋coverage＋merge 文本門檻 |
| 2026-10-01-source-publish-timebase | [2026-10-01-source-publish-timebase.md](2026-10-01-source-publish-timebase.md) | contract refinement | 1× vs sped 成片同秒互比是假錯位；analysis=source、speed=publish transform；CPS 在 publish 軸 |
| 2026-10-01-ocr-probe-discovery-not-exclusion | [2026-10-01-ocr-probe-discovery-not-exclusion.md](2026-10-01-ocr-probe-discovery-not-exclusion.md) | contract_gap candidate | bottom／`no_subtitle_like` miss≠無字幕；probe=discovery；需 recovery＋分層 coverage |
