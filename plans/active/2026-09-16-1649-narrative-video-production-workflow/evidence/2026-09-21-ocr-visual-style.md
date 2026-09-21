# Observation — OCR visual style features

**Run ID**：2026-09-21-ocr-visual-style  
**Kind**：Phase 3 observation（**不是** `contract_gap`；**不是**新 observable detector；**不**擴 workflow）  
**Extends**：[`2026-09-17-visual-text-evidence.md`](2026-09-17-visual-text-evidence.md)

## 定位

在已有 `text`／timestamp／bbox／normalized_box／frame_ref／persistence／region 上，機械層再存：

```yaml
visual_style:
  text_color:
    dominant_rgb: [245, 245, 245]
    dominant_hsv: [0, 0, 0.96]
    color_ratio: 0.89
    palette:
      - { rgb: [245, 245, 245], ratio: 0.89 }
      - { rgb: [40, 40, 40], ratio: 0.07 }
  background:
    dominant_rgb: [16, 16, 16]
  outline: { detected: true }
  shadow: { detected: true }
  contrast: { score: 0.91 }
```

不要只存單一 `"#FFFFFF"`：抗鋸齒、描邊、陰影、壓縮、半透明會污染單色。禁止 LLM 從截圖猜色。

屬 **layer.observable**。下一層才產生 `subtitle_candidate`／`watermark_candidate`／`subtitle_style_candidate`（多證據：位置＋面積＋持續＋文字是否固定＋style）。再下一層才決定是否構成 narrative evidence。

## 四個用途（皆 candidate evidence）

1. **字幕 vs 浮水印**：固定位置＋長 persistence＋固定文＋固定 style → watermark candidate；底部＋短窗＋文變＋style 一致 → subtitle candidate。不是單靠顏色。
2. **角色字幕樣式**：白／黃／藍等只當 `subtitle_style_candidate`，再與 speaker／face 共現。`yellow ≠ character_lin`。
3. **ASR／OCR 衝突**：style 連續＋位置連續＋同一 dialogue window，與 ASR 一併進 [`19-text-resolution-and-narrative-assembly.md`](../19-text-resolution-and-narrative-assembly.md)。禁止寫死 OCR 永遠贏。
4. **Narrative window**：同一 style＋時間連續 → 較易組成 interaction window，而不是三條孤立 dialogue。連 [`15-story-evidence-vs-dialogue.md`](../15-story-evidence-vs-dialogue.md)。

鏈：geometry + visual_style + ASR + face／speaker + temporal continuity → linking → text resolution → narrative window → event assembly。

## 真實片子要數

style 是否穩定可聚類；是否真能分開 watermark／subtitle；衝突時是否被 text_resolution 消費。沒被消費就不進 schema。
