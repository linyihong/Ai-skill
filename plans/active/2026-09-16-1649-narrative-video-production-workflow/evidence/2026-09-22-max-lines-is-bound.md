# Observation — max_lines is not a fill target

**Run ID**：2026-09-22-max-lines-is-bound  
**Kind**：contract（使用者授權寫進 workflow；根因不是 prompt）  
**Extends**：[`2026-09-22-subtitle-layout-engine.md`](2026-09-22-subtitle-layout-engine.md)、[`2026-09-22-speech-timing-authority.md`](2026-09-22-speech-timing-authority.md)

一行 glyph 放得下卻拆成兩行 = 把 Constraint 當 Selection。應產可行 candidates（1 行 48 寬度 fail、2 行 48 pass、1 行 44 pass），再用 `minimize_lines` 等 policy 選。locale evidence 必須能區分「選了 2 行因為 1 行放不下」vs「max=2 所以用了 2 行」。
