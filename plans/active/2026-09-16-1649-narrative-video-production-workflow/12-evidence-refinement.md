# Candidate: Evidence refinement loop

Companion to [`_plan.md`](_plan.md)。**候選獨立契約，不是 ASR／OCR parser，也不是 Phase 2 workflow 檔。** Parser 只採集；Refinement 負責連結、仲裁、驗證、學習。Phase 3 **不**建 `workflow/.../evidence-refinement.md`。  
觀察：[`evidence/2026-09-18-evidence-refinement.md`](evidence/2026-09-18-evidence-refinement.md)。

原則：不是 OCR／ASR 二選一，也不是讓 LLM 一次猜對。多種 evidence 經評估、仲裁、獨立審查與回饋，讓資料室逐步接近正確值。對齊 invariant 5：Need → Constraints → Feasible set → **Selection policy** → Selection。第一版用作品級 policy，**不**寫死全域權重。

```text
layer.observable   （凍結再加 detector）shot / keyframe / visual_text+geometry / ASR+transcript / voice / speaker / face
layer.linking      text resolution／relations／evidence_unit（共現 ≠ 等同 ≠ 劇情）
layer.event        narrative window／event assembly／candidate relations + relevance
layer.story        story state / story evidence
layer.narrative    script / template / matching / EDR / 成片
outer loop         observe → resolve → produce → verify → correct → learn policy
```

禁止：

- 只存 OCR `text`、不存 pixel／normalized box 與 persistence（幾何與持續性是一級 metadata）
- 把 `role.candidate` 寫成已判定事實（清楚時只允許 `resolver: mechanical`）
- 高 persistence 邊角文字因像人名就進 ASR 人名仲裁
- `OCR weight = 1.0` 當全域真理（字幕延遲、誤讀、LOGO 都可能）
- LLM `confidence` 當 Decision SoT（只當另一條 evidence）
- Script 反過來改寫 evidence（script 是 consumer）
- LLM 決定 OCR crop，或把單片 `likely_subtitle_region` 直接寫進全局 probe（見 [`13-mechanical-visual-text-probe.md`](13-mechanical-visual-text-probe.md)）
- `speaker_id`／`voice_track` 直接等於 `character_id`（見 [`14-voice-speaker-evidence.md`](14-voice-speaker-evidence.md)）
- 整段 ASR／「有意思的對話」直接當劇情摘要進 Script／EDR（見 [`15-story-evidence-vs-dialogue.md`](15-story-evidence-vs-dialogue.md)）
- 在 observable 層繼續加 detector，或從 links 直接跳劇情（見 [`16-evidence-unit.md`](16-evidence-unit.md)）
- 定義了 OCR／ASR selection policy，卻讓 narrative consumer 繼續讀 raw ASR；或
  把 `dialogue_cluster` 當 resolved event（見
  [`19-text-resolution-and-narrative-assembly.md`](19-text-resolution-and-narrative-assembly.md)）
- 全域 OCR 優先／ASR fallback；或 LLM 原地改 ASR；或把字幕 sanitization 當成 spoken truth（見
  [`21-spoken-vs-subtitle-reconstruction.md`](21-spoken-vs-subtitle-reconstruction.md)）
- 只存 ASR grapheme、丟掉 phonetic／syllable；或 ASR 直接進 LLM 修句；或開局寫死字幕替換表（見
  [`22-phonetic-text-reconstruction.md`](22-phonetic-text-reconstruction.md)）
- 顯示詞自動改成 spoken 猜測，或把 OCR 和諧詞拿去同音展開，或只做單次掃描、不做整集 Final Text Audit（見
  [`23-sanitization-anomaly-audit.md`](23-sanitization-anomaly-audit.md)）
- 把 watermark role 當刪除／區域 crop／contains 整條丟棄，或讓廠標變體進 name learning（見
  [`24-ocr-role-projection.md`](24-ocr-role-projection.md)）
- 同 timestamp／高相似就 destructive merge；或用翻譯 dst 反推哪個 source 相同（見
  [`25-text-group-preserve-variants.md`](25-text-group-preserve-variants.md)）
- 用字數切行、寫死 bottom％、或讓 LLM 做最終斷行／選位（見
  [`26-subtitle-layout-engine.md`](26-subtitle-layout-engine.md)）
- 全片固定 `font_size` px、或 `字數 × 字級` 估寬（見
  [`27-typography-layout-profile.md`](27-typography-layout-profile.md)）
- 為了 fit 把字縮到 min 以下，或短句放到 max 以上（見
  [`31-font-size-hard-bounds.md`](31-font-size-hard-bounds.md)）
- 讓 AI 的「舒服」直接寫進 profile，或無 iteration budget 調字級（見
  [`28-layout-review-loop.md`](28-layout-review-loop.md)）
- 為 generated 口播猜 cue 秒數，或把 speech／layout 合成一次重做（見
  [`29-speech-timing-authority.md`](29-speech-timing-authority.md)）
- 把 `max_lines` 當填滿行數的目標，或一行放得下仍拆行（見
  [`30-max-lines-is-bound.md`](30-max-lines-is-bound.md)）
- 為 fit 在詞／專名中間截斷，或同一 cue 上下行使用不同字級（見
  [`32-semantic-safe-wrap-uniform-typography.md`](32-semantic-safe-wrap-uniform-typography.md)）
- 用字數均分換行，或 layout 自創未評分斷點（見
  [`33-caption-composition-semantic-breaks.md`](33-caption-composition-semantic-breaks.md)）
