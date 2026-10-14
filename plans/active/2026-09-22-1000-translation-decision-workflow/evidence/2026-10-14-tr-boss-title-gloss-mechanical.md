# TR workplace boss-title mechanical repair

## 決策

1. **Script／locale admission ≠ social-address pass** when TR MT reads ZH
   surname+总 as office／gender titles（`Genel Müdür`／`Bayan`／`Şirket Başkanı`）
   or leaves a bare `Zong` syllable.
2. Prefer **expression-type mechanical rewrite** keyed by ZH `X总`（not `总统`）
   to `{Pinyin} Bey`；negation vocative `不是X总` → `Hayır, {Pinyin} Bey`.
3. Scope to **tr only**；FIL／TL keep `Ginoong X` sibling. Do not invent series
   cases — expand only from dogfood bad surfaces.
4. Short exact cues（`陈总`／`郑总`）may replace the whole destination；longer
   leading-negation lines rewrite the office-title span and keep trailing
   time-ish clauses when present.
5. Content Finality stays **REVIEW**；no publish PASS from unit-locked table hits.

## 觀測（sanitized）

- Postfix TR caches showed `不是 陆总` → `Hayır, Genel Müdür Bayan`，
  `郑总` → `Zong`，`陈总` → `Cen` while script gates were green.
- Product：`mechanical_tr_boss_title_gloss` after FIL boss-title on the
  validate／repair path；unit tests lock negation，leading negation，short
  cues，总统 skip，and non-tr noop.
- Offline apply on three TR episode caches；ASS Dialogue rewritten in
  lockstep with cue counts.

## 未完成

- Other Latin locales’ president／CEO glosses for `X总`／`X总裁` remain REVIEW.
- Publish／layout Finality still open.

## 連動

- Sibling：`2026-10-14-fil-phrase-polarity-mechanical.md`（FIL boss-title
  follow-up section）
- Product episode tables stay under `<PROJECT_ROOT>` analysis only.
