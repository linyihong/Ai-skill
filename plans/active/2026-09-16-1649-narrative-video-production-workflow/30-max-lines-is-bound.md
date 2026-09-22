# Candidate: max_lines is a bound not a target

Companion to [`26-subtitle-layout-engine.md`](26-subtitle-layout-engine.md)。**workflow 已補** [`subtitle-layout.md`](../../../workflow/narrative-video-production/subtitle-layout.md)。不是改 prompt。  
觀察：[`evidence/2026-09-22-max-lines-is-bound.md`](evidence/2026-09-22-max-lines-is-bound.md)。

`max_lines: 2` 錯在被讀成「做成兩行」。正確：`minimize_lines: true`；一行放得下禁止為「看起來完整」而換行。可行集含 1 行／2 行 × 字級；`selection.policy` 才選。無可行 layout → 重切 Speech Unit，不是擠進 cue。EDR／locale 記 `layout.lines` 與 `max_lines` 分開。
