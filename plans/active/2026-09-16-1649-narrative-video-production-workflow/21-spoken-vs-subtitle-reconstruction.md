# Candidate: Spoken vs subtitle text reconstruction

Companion to [`19-text-resolution-and-narrative-assembly.md`](19-text-resolution-and-narrative-assembly.md)。**候選，不是 workflow schema，也不是換 ASR 模型。**  
觀察：[`evidence/2026-09-21-spoken-vs-subtitle-reconstruction.md`](evidence/2026-09-21-spoken-vs-subtitle-reconstruction.md)。

不是「OCR 優先、ASR fallback」。OCR 可能字對但語意經劇組改寫；ASR 可能時間對但字詞錯。從兩個缺陷不同的訊號重建 **spoken text**，同時保留 **subtitle text**。

Raw ASR／OCR 永不覆寫。LLM 產出 reconstruction candidate。`duration`／語速／`char_count` 是 plausibility constraint，不是鎖定字數。優先 word timestamp 與 temporal alignment。

**先** OCR `text_region`／subtitle group（[`41-bilingual-ocr-regions-and-subtitle-groups.md`](41-bilingual-ocr-regions-and-subtitle-groups.md)），再 Language Relation Gate（[`40-language-role-before-text-resolution.md`](40-language-role-before-text-resolution.md)）：同語言才走 match／variant／`sanitization_candidate`；跨語言且語義對齊 → `cross_language_translation`（保留 `subtitle_text`＋`spoken_text`，不得 OCR 英文當 spoken canonical）。雙語疊字時 spoken primary＝與 ASR 同語言的 dialogue region；異語列只作 translation evidence。`text_relation` 擴充見 40／41；舊稱 `semantic_mismatch` 僅在同語言路徑使用。

三層 context：local 1–3 句、narrative window、episode state／詞彙。無 evidence 不得發明台詞；不確定 → unresolved。本 round：timestamp → 對齊 → language／role → 分家 → relation → reconstruction → verify，再進 narrative。Phonetic／syllable 候選與 LLM 最後一層見 [`22-phonetic-text-reconstruction.md`](22-phonetic-text-reconstruction.md)。同窗多筆近形句：group 保留 variants，見 [`25-text-group-preserve-variants.md`](25-text-group-preserve-variants.md)。Workflow 落點已吸收：[`text-evidence-language-relation.md`](../../../workflow/narrative-video-production/text-evidence-language-relation.md)（使用者授權；非新 Phase）。
