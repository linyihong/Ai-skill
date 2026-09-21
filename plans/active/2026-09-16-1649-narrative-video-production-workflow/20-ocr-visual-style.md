# Candidate: OCR visual style as evidence feature

Companion to [`_plan.md`](_plan.md) 與 [`09-visual-text-evidence.md`](09-visual-text-evidence.md)。**候選欄位，不是 workflow schema，也不是新 detector。**  
觀察：[`evidence/2026-09-21-ocr-visual-style.md`](evidence/2026-09-21-ocr-visual-style.md)。

顏色／對比／描邊屬 **observable visual text 的 feature**，不是「OCR 判斷這是字幕／這是誰說的」。正式名：`visual_style`（不要 `subtitle_color`）。機械層算好 palette／ratio／outline／shadow／contrast；後面 AI 不從截圖猜色。

`yellow ≠ character`。style 只進 `subtitle_style_candidate`／watermark 多證據，再進 text resolution 與 narrative window。Phase 3 **不**改 workflow。
