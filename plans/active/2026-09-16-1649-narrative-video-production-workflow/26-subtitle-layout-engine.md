# Candidate: Subtitle layout engine

Companion to [`01-captions-and-locales.md`](01-captions-and-locales.md)。**候選 solver，不是 workflow schema，也不是 LLM 排版。**  
觀察：[`evidence/2026-09-22-subtitle-layout-engine.md`](evidence/2026-09-22-subtitle-layout-engine.md)。

`content_gate`／`timing_gate` 之後，`layout_gate` 應是 **Need → Constraints → feasible layouts → selection.policy → selected**。`max_lines` 是上限：[`30-max-lines-is-bound.md`](30-max-lines-is-bound.md)、workflow [`subtitle-layout.md`](../../../workflow/narrative-video-production/subtitle-layout.md)。LLM 最多出 `semantic_break_candidates`。字級是 range＋scale＋fit，且 **min／max 機械硬閘**：[`27-typography-layout-profile.md`](27-typography-layout-profile.md)、[`31-font-size-hard-bounds.md`](31-font-size-hard-bounds.md)。Wrap lossless、protected spans、cue-uniform typography：[`32-semantic-safe-wrap-uniform-typography.md`](32-semantic-safe-wrap-uniform-typography.md)。Caption Composition（語意先於 fit、scored 斷點、禁止 equal-length）：[`33-caption-composition-semantic-breaks.md`](33-caption-composition-semantic-breaks.md)。
