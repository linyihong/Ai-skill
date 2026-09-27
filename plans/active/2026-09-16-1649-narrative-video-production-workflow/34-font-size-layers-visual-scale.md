# Candidate: Three-layer font bounds and visual scale

Companion to [`31-font-size-hard-bounds.md`](31-font-size-hard-bounds.md)、
[`27-typography-layout-profile.md`](27-typography-layout-profile.md)。**workflow 已補**
（[`subtitle-layout.md`](../../../workflow/narrative-video-production/subtitle-layout.md)）。  
觀察：[`evidence/2026-09-28-font-size-layers-visual-scale.md`](evidence/2026-09-28-font-size-layers-visual-scale.md)。

單層 `[36, 56]` 是安全底線，不是單句可用範圍。操作字級必須同時落在：

1. **Absolute** — 不可突破的安全區間。
2. **Profile** — 由 `layout_script` 選擇，不是 `locale`。preferred／min／max **不在本檔凍死**。
3. **Adjustment** — 相對 preferred 與 scene baseline 的窄 `max_delta`。超出就重切 Speech Unit。

跨語系對齊的是 glyph bounding box 的視覺高度，不是同一個 `font_size`。補償係數留在 profile／dogfood。AI 只出方向與建議區間；mechanical 在三層內選字級。
