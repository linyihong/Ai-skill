# FIL／TL short-phrase polarity mechanical repair

## 決策

1. **Script／locale admission ≠ polarity pass** for short FIL／TL discourse and
   state phrases that MT systematically confuses（beauty gloss for okay-filler；
   sick for dirty；recover for obey）.
2. Prefer **expression-type mechanical rewrites keyed by ZH cue text**（exact or
   contains）over stuffing series cases into Selection prompts.
3. Scope repairs to **fil／tl only**；leave other Latin locales untouched.
4. Short cues may replace the whole destination string；longer cues rewrite the
   matched bad span once. Expand the repair table from dogfood, not episode IDs.
5. Content Finality stays **REVIEW** until declared assertions／spot rebuild cover
   table-external polarity；no publish PASS from unit-locked table hits alone.

## 觀測（sanitized）

- Postfix matrix REVIEW packs showed FIL cues where ZH okay-filler／dirty／obey
  mapped to beauty／sick／recover glosses while script gates were green.
- Product added `mechanical_fil_phrase_polarity` after latin name consistency on
  the validate／repair path；unit tests lock the three seed confusions and
  id-skip／tl-alias behavior.
- Same session class as Arabic divorce／marry invert：antonym-prone or
  near-antonym MT confusions need polarity handling beyond “looks like target
  script”.

## 未完成

- Full locale pack rebuild／spot REVIEW for previously poisoned FIL cues.
- Broader declared semantic assertions beyond the seed table.
- Publish／layout Finality still open.

## 連動

- Lesson：
  `feedback/history/narrative-video-production/common/2026-10-14_123700-fil-phrase-polarity-not-script-gate.md`
- Sibling：`2026-10-08-divorce-marry-polarity-invert.md`
- Prior same-day：`2026-10-14-ass-empty-unresolved-skew-and-semantic-review.md`
  （FIL filler semantics marked open → narrowed by this mechanical table）
- Product episode tables stay under `<PROJECT_ROOT>` analysis only.

## Follow-up: boss title gloss

- ZH surname+总 (workplace address) was MT-read as Pangulo / Hindi siya.
- Product: mechanical_fil_boss_title_gloss (fil/tl); 总统 untouched via 总(?!统).
- Dogfood: 不是陆总 → Hindi, Ginoong Lu; 不陆总我现在 drops Pangulo ng Kontinente.

## Follow-up: exact 「好」

- Exact ZH「好」was still MT-glossed as Maganda after the 好不好 rule（exact match
  does not cover the single-character filler）.
- Product table: exact「好」+ Maganda → Okay（fil／tl）.
- Dogfood: two overnight FIL cues rewritten offline；broader Maganda＝beautiful
  remains out of scope.
