# Candidate: OCR role is projection, not deletion

Companion to [`09-visual-text-evidence.md`](09-visual-text-evidence.md)、[`13-mechanical-visual-text-probe.md`](13-mechanical-visual-text-probe.md)。**候選 invariant，未凍結、未進 workflow gate。**  
觀察：[`evidence/2026-09-22-watermark-exclusion-is-projection.md`](evidence/2026-09-22-watermark-exclusion-is-projection.md)。

Watermark role 可以機械判定；**不得**因此刪 observable visual-text、crop 該區、或 `contains(watermark)` 整條 discard。採集／角色／過濾／下游使用必須拆開。下游用 role-aware projection（narrative／dialogue／scene_text／raw retained）。需要 `text_span_role`：合併 box 內 watermark span 與 scene／subtitle span 分開。watermark STOP 進 name／entity learning。第二、三個同型 episode 後才升 contract。Phase 3 **不**改 workflow。
