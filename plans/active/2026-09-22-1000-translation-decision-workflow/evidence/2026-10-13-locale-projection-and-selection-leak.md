# Locale projection gates vs Selection scaffolding leak

## 決策

1. **Target hardsub OCR ≠ default realization** when spoken／pivot is source-language
   dialogue：mono foreign hardsub residue must not override MT for that locale.
2. **Script residue repair order**：peel／normalize residual source glyphs（e.g. CJK
   name leftover in Thai）**before** the admit／reject gate；reject-before-strip
   causes `missing_target_realization` and **projection_loss**（target cue count
   ＜ source cue count → index skew for consumers that zip by order）.
3. **I21 incomplete_source** must cover OCR／layout crumbs（digit-only、bare
   punctuation such as `:0`／`(1`），not only ellipsis／comma cuts — otherwise
   Selection may emit **prompt scaffolding headers**（e.g. “Guards / candidate
   space:”）as destination text.
4. **Script／honorific reject ≠ silent drop**：latin leftover and English honorific
   rejects still need a disposition path that preserves timeline alignment or a
   bounded name／honorific repair；mechanical fail must not imply shifted cues.

## 觀測（sanitized）

- Multi-episode postfix caption matrix（REVIEW only）：acquisition unaccounted=0 on
  completed episodes；locale packs still showed projection_loss on several
  locales when realization was rejected.
- EN locale：mono hardsub region preference overwrote good MT until product gate
  limited OCR override to bilingual／non-source-spoken cases.
- TH：single CJK surname glyph in an otherwise Thai candidate was rejected as
  residual_cjk before strip；after strip-before-admit，cue count matched pivot.
- One episode’s ZH pivot contained digit／punct crumbs；EN dst copied Selection
  adapter section title text — treated as incomplete_source BLOCK going forward.

## 未完成

- Honorific Latin (`Mr.`／`Miss`) and latin_leftover name forms still drop cues
  on some locales；1:1 timed slot policy not yet universal.
- Semantic antonym／filler mistranslation（separate from scaffolding）still needs
  declared assertions／human REVIEW — not closed by these mechanical gates.
- Publish／burn／layout Finality remain open（content≠timing≠layout）.

## 連動

- Lessons：
  - `feedback/history/narrative-video-production/common/2026-10-13_133000-selection-scaffolding-must-not-become-dst.md`
  - `feedback/history/narrative-video-production/common/2026-10-13_133100-residue-admit-order-and-projection-alignment.md`
- Product paths／episode tables stay under `<PROJECT_ROOT>` analysis only.
