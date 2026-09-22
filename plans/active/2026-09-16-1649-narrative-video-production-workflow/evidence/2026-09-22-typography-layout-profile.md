# Observation — font size is a fit constraint

**Run ID**：2026-09-22-typography-layout-profile  
**Kind**：Phase 3 observation（**不是**立刻 `contract_gap`；**不**散落寫死 px）  
**Extends**：[`2026-09-22-subtitle-layout-engine.md`](2026-09-22-subtitle-layout-engine.md)

## 不是固定樣式

`preferred` 是起點。實際字級由 safe area、max width／height、max_lines、line_height、forbidden regions、**glyph metrics** 共同決定。同字元數寬度可差很大；禁止 `字數 × 48`。多語尤其必須量測。

## 相對畫布再 clamp

`font_size ≈ preferred × (video_height / reference_height)`，再 `min ≤ size ≤ max`。720 高不用 48；4K 不得無上限放大。`scaling.mode: canvas_height` 可配 min_scale／max_scale。

## Overflow 順序（policy）

預設精神：1 行 fit → 否則 2 行（仍在 max_lines 內）→ 再降 font → 仍無可行則重切 Speech Unit。實際順序由 `selection.policy` 決定。禁止「max=2 所以先做成 2 行」。見 [`30-max-lines-is-bound.md`](../30-max-lines-is-bound.md)。

## 同一 profile

font_family、font_size range／scale、line_height、stroke／outline、safe area、max width、max lines。短劇覺得 44 比 48 舒服 → 改 profile／policy，不改 workflow。
