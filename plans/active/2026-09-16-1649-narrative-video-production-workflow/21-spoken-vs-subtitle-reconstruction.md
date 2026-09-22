# Candidate: Spoken vs subtitle text reconstruction

Companion to [`19-text-resolution-and-narrative-assembly.md`](19-text-resolution-and-narrative-assembly.md)。**候選，不是 workflow schema，也不是換 ASR 模型。**  
觀察：[`evidence/2026-09-21-spoken-vs-subtitle-reconstruction.md`](evidence/2026-09-21-spoken-vs-subtitle-reconstruction.md)。

不是「OCR 優先、ASR fallback」。OCR 可能字對但語意經劇組改寫；ASR 可能時間對但字詞錯。從兩個缺陷不同的訊號重建 **spoken text**，同時保留 **subtitle text**。

Raw ASR／OCR 永不覆寫。LLM 產出 reconstruction candidate。`duration`／語速／`char_count` 是 plausibility constraint，不是鎖定字數。優先 word timestamp 與 temporal alignment。`text_relation` 可標 `semantic_mismatch` 或 `subtitle_sanitization`。三層 context：local 1–3 句、narrative window、episode state／詞彙。無 evidence 不得發明台詞；不確定 → unresolved。本 round：timestamp → 對齊 → 分家 → conflict → reconstruction → verify，再進 narrative。Phonetic／syllable 候選與 LLM 最後一層見 [`22-phonetic-text-reconstruction.md`](22-phonetic-text-reconstruction.md)。同窗多筆近形句：group 保留 variants，見 [`25-text-group-preserve-variants.md`](25-text-group-preserve-variants.md)。Phase 3 **不**改 workflow。
