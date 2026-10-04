# Observation — sticky watermark on accepted path needs projection

**Run ID**：2026-10-02-sticky-watermark-accepted-projection
**Kind**：Phase 3 adapter／contract（接續 accepted-quality audit；**projection 非 deletion**）
**Extends**：[`2026-09-22-watermark-exclusion-is-projection`](2026-09-22-watermark-exclusion-is-projection.md)、[`2026-10-02-ep16-accepted-uncertain-audit`](2026-10-02-ep16-accepted-uncertain-audit.md)

## 問題

preserve 後 `accepted` 仍可能把 **對白 + sticky 水印／brand** 整串當 publishable：

- OCR：「…哥外卖／哥外／PUMA…」
- ASR：常已是乾淨對白

不得：整條 delete。不得：維持整串 `spoken_selected`。

## 目標行為

```text
OCR observed (raw 全留)
        ↓
sticky / brand span projection（機械）
        ↓
dialogue_projected + watermark_spans[]
        ↓
┌─ 投影成功且像對白 → spoken = projected hardsub（中文硬字幕權威）
├─ ASR ≡ projected 或 trivial containment → resolved；spoken = projected
├─ 僅 sticky／brand 殘片 → rejected（evidence 不成立）
└─ ASR 與 projected 語義衝突 → uncertain；spoken 仍 = projected（禁止 loose Jaccard 選 ASR）
```

## 規則

1. **Projection not deletion**：`subtitle.observed` 保留原 OCR；只改 `spoken.selected`／publishable 投影。
2. **Sticky affix／brand** 先 peel，再 decide status。
3. **Hardsub authority after peel**：中文硬字幕片源，投影後對白以 OCR projected 為 spoken；ASR 僅在完全相同或 trivial containment（短子集／≤2 字差）時與 projected 同向 resolved。**禁止** `jaccard≥0.45` 這類 loose 對齊讓 ASR 覆蓋（會把「应酬→诱惑」「私会→死回」 silently prefer）。
4. 語義衝突 → `uncertain`，spoken 仍 projected；保留 observed＋projected＋asr；**禁止** coerce reject-unless 整段皆非對白。
5. 不開新 OCR／新 Phase。

## Validation

- [x] unit：黏串→投影；純「哥外」→rejected；应酬/诱惑 conflict → spoken=应酬 + uncertain
- [x] ep16 re-audit：accepted 30→28；rejected sticky-only 3；**accepted spoken sticky=0**；retained=38
- [x] 本 evidence + README 索引
- [x] prefer-OCR 後 ep16 offline audit：accepted 21 / uncertain 14 / rejected 3；应酬≠诱惑、私会≠死回 → spoken=OCR + uncertain；`sticky_projection: projected=21 uncertain=12 rejected=3`

## Dogfood result（sanitized）

| metric | before | after sticky projection |
| --- | --- | --- |
| accepted / publishable | 30 | 28 |
| uncertain | 8 | 7 |
| rejected | 0 | 3（純「…哥外」殘片） |
| evidence_retained | 38 | 38 |
| accepted spoken 仍含 sticky／brand | 多 | **0** |

`sticky_projection: projected=21 uncertain=5 rejected=3`

## 產品落點

`sticky_ocr_projection` helper → `resolve_fused_to_cues` 出口前 post-pass。
