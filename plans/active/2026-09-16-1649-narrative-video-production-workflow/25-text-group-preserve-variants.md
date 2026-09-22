# Candidate: Text grouping is not deletion

Companion to [`19-text-resolution-and-narrative-assembly.md`](19-text-resolution-and-narrative-assembly.md)、[`24-ocr-role-projection.md`](24-ocr-role-projection.md)。**同一類：機械層 grouping ≠ discard。未凍結、未進 workflow gate。**  
觀察：[`evidence/2026-09-22-text-group-preserve-variants.md`](evidence/2026-09-22-text-group-preserve-variants.md)。

同 timestamp ＋高相似 **不是**同一 utterance。timestamp 是 alignment signal，不是 identity key。Merge 只對 **exact duplicate**；variant／truncated／conflict 進 `text_group`，raw 全留。LLM 收 candidate set，不得收已合成的單筆。**dst／翻譯不得決定 source merge。** Source resolution → canonical spoken → 才 translation。調 similarity threshold 不算修。Phase 3 **不**改 workflow。
