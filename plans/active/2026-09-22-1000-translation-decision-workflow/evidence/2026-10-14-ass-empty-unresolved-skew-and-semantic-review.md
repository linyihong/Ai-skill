# ASS empty unresolved slots vs semantic REVIEW

## 決策

1. **Cache 1:1 timed slots ≠ deliverable alignment** until ASS／burn（and any dense
   cue list export）preserve empty／`text_source=unresolved` events. Skipping blank
   `Dialogue` lines re-creates **index skew** for consumers that zip by order.
2. **Join by `source_index` or start time**, never by dense ASS line index, when
   comparing pivot ↔ locale for REVIEW or repair.
3. **Latin honorific soften**（allow `Mr.`／`Miss` inside a real Latin-script
   sentence; still reject bare honorific+name crumbs）reduces false
   `en_honorific` drops without reopening non-Latin residue hard reject.
4. **Semantic／romanization REVIEW** remains open：filler polarity（e.g. “好不好”
   → beauty gloss）、dirty≠sick、and inconsistent surname romanization（same
   person `Lu` vs `Liu`）are content Finality issues — not closed by cue-count
   gates.

## 觀測（sanitized）

- Postfix multi-episode caption matrix finished REVIEW：acquisition
  unaccounted=0；locale workers 0 failed after re-locale of one kinship
  NameError race episode.
- One locale pack：translation cache length matched pivot and held an empty
  unresolved slot；exported ASS omitted that event → subsequent lines looked
  like “mistranslation” under index zip；time keys showed cascading skew.
- Another pack：digit-only／numeral pivot crumbs still produced a shorter
  translated ASS list and skew.
- FIL／ID spot REVIEW：known filler／adjective confusions and surname
  romanization inconsistency persist on otherwise mechanically green packs.
- Mid-run product soften for Latin `en_honorific` stopped mixed-sentence drops
  on later samples；bare honorific+name still rejected in tests.

## 未完成

- ~~ASS／burn must emit timed placeholders for unresolved／empty text~~ → product
  landed（empty／unresolved keep timed Dialogue；wrong-script → empty slot not
  delete）. Re-verify on next matrix／rebuild；REVIEW join still prefer
  `source_index`／start.
- Declared semantic assertions／glossary binding for high-churn surnames and
  FIL fillers.
- Publish／layout Finality still open；no publish PASS from this dogfood.

## 連動

- Lesson：
  `feedback/history/narrative-video-production/common/2026-10-14_084000-empty-unresolved-must-survive-ass-export.md`
- Prior evidence：`2026-10-13-locale-projection-and-selection-leak.md`（1:1 cache
  policy；honorific drop still open → partially narrowed by Latin soften）
- Product episode tables stay under `<PROJECT_ROOT>` analysis only.
