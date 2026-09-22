# Observation — subtitle layout as constraint solver

**Run ID**：2026-09-22-subtitle-layout-engine  
**Kind**：Phase 3 observation（`layout_gate` 已有；缺 parameterized engine。**不**改 workflow；**不**讓 LLM 斷行／選位）  
**Extends**：[`01-captions-and-locales.md`](../01-captions-and-locales.md)、[`2026-09-22-watermark-exclusion-is-projection.md`](2026-09-22-watermark-exclusion-is-projection.md)

## Content ≠ Layout

spoken／subtitle **寫什麼** 與 **怎麼放** 分開。Translation 只到 content；layout 在 NVP。

## 機械 solver

Constraints：dynamic safe area（後接 obstruction map）、`max_lines` policy、safe width／height ratio、glyph measurement（非 `len(text)`）、break_policy、font preferred／min／max、preferred_position＋fallback、stability（cue 窗內 hold，禁止逐幀跳位）。Overflow → 再斷行 → 降字號 → split cue／reject；禁止硬縮到不可讀。

高相似／禁區：watermark／UI／face 是 obstruction **evidence**，不是刪 OCR。

## Layout profile

9:16／16:9／1:1 換 profile 數字，不換 workflow。對齊 invariant 5。

## 分期

- V0：畫布、letterbox、safe area、max_lines／width、glyph break、font 區間、preferred＋fallback、render overflow QC  
- V1：OCR／face／watermark／UI obstruction、position stability  
- V2：LLM semantic break hints、語境節奏、locale typography  

本 round 只驗 V0 範圍是否足以表達「最大兩行、動態可用區」，不做完整引擎。
