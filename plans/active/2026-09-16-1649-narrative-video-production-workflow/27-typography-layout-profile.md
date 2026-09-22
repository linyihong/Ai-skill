# Candidate: Typography as layout constraint

Companion to [`26-subtitle-layout-engine.md`](26-subtitle-layout-engine.md)。**候選 profile，不是 CSS、也不是 workflow 寫死 px。**  
觀察：[`evidence/2026-09-22-typography-layout-profile.md`](evidence/2026-09-22-typography-layout-profile.md)。

`font_size` 是 **range + canvas scale + fit solver**，不是 `48`。preferred → glyph 量測 → 放不下再降到 min → 仍溢出則 overflow.order（rebreak → reduce_font → split_cue → reject）。縮放相對 reference canvas（例高度 1920），再 clamp min／max。字數公式禁止。font／line_height／stroke／safe area／max width／max lines 同屬 **Typography／Layout Profile**；dogfood 只改 profile，不改 workflow。AI 不直接改字級：[`28-layout-review-loop.md`](28-layout-review-loop.md)。Phase 3 **不**改 workflow。
