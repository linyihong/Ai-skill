# Candidate: Typography as layout constraint

Companion to [`26-subtitle-layout-engine.md`](26-subtitle-layout-engine.md)。**候選 profile，不是 CSS、也不是 workflow 寫死 px。**  
觀察：[`evidence/2026-09-22-typography-layout-profile.md`](evidence/2026-09-22-typography-layout-profile.md)。

`font_size` 是 **range + canvas scale + fit solver**，不是 `48`。行數：先測 1 行能否放下，再評 2 行，再依 `selection.policy` 降字級；無可行則重切 Speech Unit（[`30-max-lines-is-bound.md`](30-max-lines-is-bound.md)）。縮放相對 reference canvas，再 clamp min／max。字數公式禁止。font／line_height／stroke／safe area／max width／max lines 同屬 **Typography／Layout Profile**。AI 不直接改字級：[`28-layout-review-loop.md`](28-layout-review-loop.md)。
