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
- Force re-acq showed fuse drops numerals，but **resolution_loss re-probe**
  reintroduced them onto publishable cues → product demote-on-write
  （OCR-only digit pivots → rejected）verified offline：cues 40→38，digitish empty.
- Sibling ASS empty-slot rebuild on another episode：locale ASS length matched
  pivot including one timed empty unresolved slot（43＝43）.

## 未完成

- Optional full locale re-projection of demoted episodes.
- Broader numeral phrases（第三课）stay dialogue when non-numeral CJK present.
- Publish／layout Finality still open.

## 連動

- Lesson：
  `feedback/history/narrative-video-production/common/2026-10-14_134500-digit-only-pivot-not-publishable.md`
- Prior I21 crumb note：product `b7bc1a9`（Arabic／punct）；this extends CN numerals
  + acquisition admission.
- Sibling：`2026-10-14-ass-empty-unresolved-skew-and-semantic-review.md`
- Product episode tables stay under `<PROJECT_ROOT>` analysis only.

## Follow-up (same day)

- Offline locale reproject: drop orphan numeral locales, pad missing ZH starts
  as unresolved empties; ASS Dialogue residual vs ZH cleared on dogfood pack.
- Product helper `align_timed_cues_to_zh_pivot` + matrix rewrite-before-burn.
- Dirty surname pack: 陆* `Liu`→`Lu` reapplied (table already present).
