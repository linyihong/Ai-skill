# Candidate: Subtitle layout engine

Companion to [`01-captions-and-locales.md`](01-captions-and-locales.md)。**候選 solver，不是 workflow schema，也不是 LLM 排版。**  
觀察：[`evidence/2026-09-22-subtitle-layout-engine.md`](evidence/2026-09-22-subtitle-layout-engine.md)。

`content_gate`／`timing_gate` 之後，`layout_gate` 應是 **Need → Constraints → feasible layouts → selection.policy → selected**，不是字數切兩行或寫死 bottom 10%。LLM 最多出 `semantic_break_candidates`；glyph 寬、safe area、max_lines、font 區間、overflow QC 由 mechanical engine。決策單位是 cue／scene window，不是每幀重定位。Obstruction 消費 OCR／face **projection**（watermark 保留當禁區）。先 V0（尺寸、letterbox、glyph break、min/max font、preferred＋fallback）。Phase 3 **不**改 workflow。
