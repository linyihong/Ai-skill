# Observation — three-layer font bounds and visual scale

**Run ID**：2026-09-28-font-size-layers-visual-scale  
**Kind**：contract（製片觀察；已寫進 workflow）  
**Extends**：[`2026-09-24-font-size-hard-bounds.md`](2026-09-24-font-size-hard-bounds.md)

## 反例

- 同一部片在 absolute 36–56 內出現 56、48、38，畫面忽大忽小。
- 中文、英文、泰文用同一個 px，視覺重量並不相同。

## Contract

字級同時受 absolute、`layout_script` profile、以及 preferred／scene `max_delta`。
到 profile min 仍放不下 → 重切 Speech Unit，不得再縮向 absolute floor。
跨 script 的目標是 glyph 視覺高度。Profile 的 px 與 visual_scale **不凍死**，留給 dogfood。
