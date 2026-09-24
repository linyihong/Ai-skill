# Observation — font min/max are hard constraints

**Run ID**：2026-09-24-font-size-hard-bounds  
**Kind**：contract（使用者授權；機械強制，不是 AI 調到塞進畫面）  
**Extends**：[`2026-09-22-typography-layout-profile.md`](2026-09-22-typography-layout-profile.md)、[`2026-09-22-max-lines-is-bound.md`](2026-09-22-max-lines-is-bound.md)

Technically fit 但不可讀 = 沒有封 `min`。短句放大到 80px = 沒有封 `max`。每 cue ±1px = 沒有 `step`／`allowed`。Scene 主字級 48 卻單句 40，優先重切 unit。Overflow：1 行 → 2 行 → 只在 bounds 內降字級 → resegment → reject。
