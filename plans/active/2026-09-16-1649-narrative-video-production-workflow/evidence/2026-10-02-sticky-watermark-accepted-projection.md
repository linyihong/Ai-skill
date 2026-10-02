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
┌─ 投影成功且像對白 → accepted；spoken= projected 或更乾淨的 ASR
├─ 僅 sticky／brand 殘片 → rejected（evidence 不成立）
└─ 投影後仍與 ASR 衝突 → uncertain（保留 observed＋projected＋asr）
```

## 規則

1. **Projection not deletion**：`subtitle.observed` 保留原 OCR；只改 `spoken.selected`／publishable 投影。
2. **Sticky affix／brand** 先 peel，再 decide status。
3. ASR 更乾淨且與投影對白可對齊 → spoken 優先 ASR（observed 仍 OCR）。
4. 失敗 → `uncertain`，reason `sticky_watermark_unresolved`／`brand_mix`；**禁止** coerce reject-unless 整段皆非對白。
5. 不開新 OCR／新 Phase。

## Validation

- [ ] unit：黏串→投影；純「哥外」→rejected；ASR 乾淨→spoken=ASR
- [ ] ep16 re-audit：accepted sticky 比例下降；uncertain／rejected 可審計
- [x] 本 evidence + README 索引

## 產品落點

`sticky_ocr_projection` helper → `resolve_fused_to_cues` 出口前 post-pass。
