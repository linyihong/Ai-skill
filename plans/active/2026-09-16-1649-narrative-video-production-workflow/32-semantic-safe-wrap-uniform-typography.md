# Candidate: Semantic-safe wrap and uniform cue typography

Companion to [`30-max-lines-is-bound.md`](30-max-lines-is-bound.md)、
[`31-font-size-hard-bounds.md`](31-font-size-hard-bounds.md)。**workflow 已補硬閘**
（[`subtitle-layout.md`](../../../workflow/narrative-video-production/subtitle-layout.md)）。  
觀察：[`evidence/2026-09-25-semantic-safe-wrap-uniform-typography.md`](evidence/2026-09-25-semantic-safe-wrap-uniform-typography.md)。

根因不是 max_lines 或 min font 本身，而是把 **wrap** 與 **segmentation** 混用：
layout 為 fit 截字，造成「談｜話」。Wrap 必須 lossless；break 不得進
`protected_spans`。若無合法 break，回 Speech Unit 重切並重做 timing。

另一個硬閘：同一 cue 只有一套 typography；上下行不得各自 auto-fit。任何一行
overflow，整個 cue 用同一候選字級重算，否則 resegment／reject。
