# Observation — OCR join glued Latin after correct parts

**Run ID**：2026-10-01-ocr-join-token-seam-not-cumulative-script

**Kind**：Phase 3 dogfood → adapter／implementation defect（非新 Phase）

**Plan**：[`../42-ocr-boundary-and-script-aware-normalization.md`](../42-ocr-boundary-and-script-aware-normalization.md)

**Workflow**：[`../../../workflow/narrative-video-production/text-evidence-ocr-boundary.md`](../../../workflow/narrative-video-production/text-evidence-ocr-boundary.md)

## 去敏情境

雙語硬字幕。L0 `parts[]` 已切開英文詞；derived 整段 `text` 出現 `I'm fromapetstore`、`You'refinally here`、`Needamaster` 等。

## 機械觀察

| 層 | 結果 |
| --- | --- |
| parts | I'm／from／a／pet／store 等已分 token |
| join（cumulative script） | 前綴 CJK 使 out=mixed → 後續 Latin\|Latin 不插空 |
| derived text | 黏字串 |
| 分類 | projection／join defect；非「OCR 不會切詞」；非新 Phase |

## Invariant

1. parts 保留；text = `script_aware_join` derived，可重建。
2. Join 看 token 接縫，不看 cumulative mixed／CJK。
3. Intra-box 單 token 黏串另案（geometry／lexical）。

## Regression（固定）

I'm/from/a/pet/store；Door/to/door/cat…；You're/finally/here；I/feel/so/bad/all/over；Need/a/master；Hurry/up/and…；以及 **CJK + Latin + Latin + Latin** 觸發型。
