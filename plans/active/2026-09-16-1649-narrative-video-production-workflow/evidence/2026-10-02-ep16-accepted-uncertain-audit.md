# Observation — ep16 accepted quality & uncertain recoverability (offline audit)

**Run ID**：2026-10-02-ep16-accepted-uncertain-audit
**Kind**：Phase 3 observation（接續 [`2026-10-02-destructive-finalization-vs-preservation`](2026-10-02-destructive-finalization-vs-preservation.md)；**不加 OCR、不開新 workflow**）
**方法**：cached OCR／ASR → fuse → `resolve_fused_to_cues`（preserve path）；機械 spot-check tags（非人工 ground truth）

## Counts

| bucket | n |
| --- | --- |
| accepted | 30 |
| uncertain | 8 |
| evidence_retained | 38 |
| publishable（≈accepted） | 30 |

## Uncertain（8）— 原因分佈

| 模式 | n | 例（sanitized） | 應有 status |
| --- | --- | --- | --- |
| `semantic_mismatch` + sticky watermark／UI 字黏在 OCR | ~5 | OCR 含「哥外／外卖」類殘片，ASR 為短稱謂／不同句 | **uncertain**（evidence 成立但 interpretation 不清；**禁止** delete／強行選邊） |
| `semantic_mismatch` + brand／logo 混入 | ~1 | OCR 品牌行＋人名，ASR 另一人名 | uncertain |
| ASR 為 OCR 子集／部分對齊 | ~1–2 | ASR「算了」vs OCR 更長黏串 | uncertain／後續可再解 |

**判定**：uncertain preservation **仍成立** — 這 8 條沒被 hard-delete；多數是「對白價值＋黏連雜訊／跨模態衝突」，適合留待 voice／再採樣／region projection，不是立刻 reject。

## Accepted（30）— 品質 spot-check

| quality_tag（機械） | 約 n（accepted 內） | 解讀 |
| --- | --- | --- |
| `asr_subset_of_ocr`／`plausible_align` | ~15 | OCR↔ASR 大致同句；**較像可用對白** |
| `likely_watermark_or_sticky_ui` | ~11 | OCR 對白後黏「哥外／外卖」等；ASR 常已是乾淨句 — **accepted 過度樂觀** |
| `brand_or_logo_mix` | ~4 | PUMA／PUT 等與對白同框 — **spoken／subtitle 不該整串進 publishable** |

**結論（Accepted quality）**：由「待驗證」→ **partial fail／contract gap**。
Preserve 修好了「歸零」；但 **accepted ≠ 乾淨對白**。問題常是 **sticky watermark／logo 被併進同一 OCR 字串後仍 `resolved`**，不是 uncertain 太多。

對齊既有 watermark 方向：[`2026-09-22-watermark-exclusion-is-projection`](2026-09-22-watermark-exclusion-is-projection.md) — 應 **role-aware projection**（去掉水印區、保留對白），**禁止**整條 delete，也 **禁止**把黏串當 clean spoken。

## Phase 3 checklist（更新）

| 項 | 狀態 |
| --- | --- |
| Evidence preservation | **PASS** |
| Uncertain preservation | **PASS**（8 條保留；以 semantic_mismatch／黏連為主） |
| Accepted quality | **partial fail** — ~半數 accepted 帶 sticky watermark／brand；~半數對白可用 |
| Uncertain recoverability | **仍待驗證**（需 region／voice／再解析；本輪只分類原因） |

## 下一步（仍不加功能爆炸）

1. **不**為 cues=0 或 OCR recall 再開 Phase。
2. 優先：accepted 路徑上的 **sticky watermark／logo projection**（從 OCR observed 投影出 clean spoken／subtitle），失敗則降為 `uncertain`，不得維持「整串 accepted」。
3. uncertain 維持一級；semantic_mismatch 繼續累積 corpus。
4. 產品／人工抽樣：對 ~15 條 plausible accepted 做成片級對白抽查（本輪未做影像對看）。

## 連動

- Workflow：[`text-evidence-multimodal-resolution.md`](../../../workflow/narrative-video-production/text-evidence-multimodal-resolution.md)（accepted 不得吞 sticky non-dialogue）
- Watermark：[`2026-09-22-watermark-exclusion-is-projection`](2026-09-22-watermark-exclusion-is-projection.md)
- Parent：[`2026-10-02-destructive-finalization-vs-preservation`](2026-10-02-destructive-finalization-vs-preservation.md)

## Validation

- [x] Full 38-row audit JSON／MD 於 `<PROJECT_ROOT>/docs/analysis/`（本機／SoT analysis only）
- [x] uncertain 原因已分類
- [x] accepted sticky／plausible 比例已標出
- [ ] role-aware sticky projection adapter
- [ ] uncertain recoverability dogfood
- [ ] 人工／成片抽查 plausible accepted
