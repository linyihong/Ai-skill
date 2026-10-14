# Digit-only pivot crumbs are not publishable ZH

## 決策

1. **Arabic digits and short pure Chinese numerals**（七／八／7／８）as
   standalone OCR cues are **layout／slate crumbs**, not dialogue pivots.
2. Extend I21／`incomplete_source` so Selection cannot complete or translate
   these crumbs into locale glosses（e.g. 「八」→ Delapan）.
3. At acquisition fuse：drop digit-only OCR unless ASR text strongly matches the
   same numeral（someone actually spoke the number）.
4. OCR-hardsub accept（ASR-missed lower-band）must **not** promote pure numeral
   tokens.
5. Existing episode caches stay dirty until force re-acquire；REVIEW only — no
   publish PASS from unit locks alone.

## 觀測（sanitized）

- Postfix matrix ZH pivot held single-token numeral cues with `asr_char_count=0`.
- Arabic digit cues already hit I21 → unresolved empty locale slots；CN numerals
  did not，so MT invented number words and ASS length skewed vs locales.
- Product landed shared `is_digit_only_layout_crumb` + fuse／hardsub gates with
  unit locks for crumb I21 and fuse drop／keep-when-ASR-matches.

## 未完成

- Force re-acquire of affected episodes to shrink publishable cue lists.
- Broader numeral phrases（第三课）stay dialogue when non-numeral CJK present.
- Publish／layout Finality still open.

## 連動

- Lesson：
  `feedback/history/narrative-video-production/common/2026-10-14_134500-digit-only-pivot-not-publishable.md`
- Prior I21 crumb note：product `b7bc1a9`（Arabic／punct）；this extends CN numerals
  + acquisition admission.
- Sibling：`2026-10-14-ass-empty-unresolved-skew-and-semantic-review.md`
- Product episode tables stay under `<PROJECT_ROOT>` analysis only.
