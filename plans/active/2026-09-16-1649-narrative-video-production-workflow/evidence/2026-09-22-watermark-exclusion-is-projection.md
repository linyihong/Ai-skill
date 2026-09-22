# Observation — watermark exclusion is a projection

**Run ID**：2026-09-22-watermark-exclusion-is-projection  
**Kind**：Phase 3 `contract_gap` **candidate**（watermark 判定合理；錯在判定後的處理。**不**改 workflow；第二、三個同型 episode 後才升 invariant／gate）  
**Extends**：[`2026-09-17-visual-text-evidence.md`](2026-09-17-visual-text-evidence.md)、[`2026-09-18-mechanical-visual-text-probe.md`](2026-09-18-mechanical-visual-text-probe.md)、[`2026-09-17-identity-precedes-naming.md`](2026-09-17-identity-precedes-naming.md)

## 案例（去敏）

左上 persistent 廠標被機械判 `role.candidate: watermark`（合理）。同一畫面另有 scene／title 字（例航空公司＋人名）。因區域／字串過濾，有效文字被切掉前綴（廠標「天…」與「天美…」重疊），下游變成殘缺專名。廠標變體同時進了 `ocr_name_mentions`／`preferred_names`。

## 要拆的混管線

採集 ≠ 角色判定 ≠ 過濾 ≠ 下游使用。Role Detection ≠ Filter。

禁止：`deleted: true`、`ignore_region` crop、整條 OCR item 因 contains(廠標字) discard。允許：該位置 watermark **prior**，同區其他 token 仍可進 candidate。合併字串先 span segmentation，再逐 span 給 `text_span_role`。

## 可用性投影

raw 全留。projection：`narrative_text`／`spoken_text_support`／`scene_text`／`watermark_excluded`。`role: watermark` → STOP identity／name learning。候選 invariant：**Watermark exclusion is a projection rule, not an evidence deletion rule.**
