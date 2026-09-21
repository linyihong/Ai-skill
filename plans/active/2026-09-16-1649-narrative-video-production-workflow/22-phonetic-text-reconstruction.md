# Candidate: Phonetic text reconstruction

Companion to [`21-spoken-vs-subtitle-reconstruction.md`](21-spoken-vs-subtitle-reconstruction.md)。**候選層，不是重跑 ASR，也不是 workflow schema。**  
觀察：[`evidence/2026-09-21-phonetic-text-reconstruction.md`](evidence/2026-09-21-phonetic-text-reconstruction.md)。

OCR 優先仍不夠：字幕可完全正確卻 sanitization；ASR 近音錯字（例 grapheme 不同、syllable 接近）。要問的是 ASR／OCR／context／knowledge **各提供什麼 constraint**，能否在 candidate space 裡重建 spoken text。

ASR 除 raw text 外保存 phonetic／syllable／phoneme；再做 **音→字第二次 decoding**（normalization → near-homophone set → context ranking）。`char_count` 弱；syllable 序列較強。LLM **最後**才 ranking／reconstruction，不得從 raw ASR 直接生成台詞。字幕替換當 Learning Inbox 的 `sanitization_mapping` candidate（多集後才 supported），禁止開局寫死對照。無 phonetic／context／OCR／knowledge 支撐 → unresolved。前後二次掃描與 `text_alert` 見 [`23-sanitization-anomaly-audit.md`](23-sanitization-anomaly-audit.md)。Phase 3 **不**改 workflow、**不**換 ASR 模型。
