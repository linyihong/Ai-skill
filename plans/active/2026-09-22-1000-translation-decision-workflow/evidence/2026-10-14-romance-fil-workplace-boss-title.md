# Romance／FIL workplace boss-title mechanical repair

## 決策

1. **Workplace `X总`／`X大总` ≠ head-of-state** when MT emits
   `presidente`／`président`／`Panggawa` on Latin／FIL surfaces.
2. Keep **`总裁`／`总统` untouched** — those may legitimately gloss as CEO／president.
3. Shared surname extraction prefers `X大总` then plain `X总`，skips modifier
   chars（大／老／副）so `陈大总` yields 陈 not 大.
4. es／fr／pt → `jefe`／`Monsieur`／`Sr.` span rewrite；FIL extends existing
   boss gloss to `Panggawa` → `Ginoong X`.
5. Content Finality stays **REVIEW**；no publish PASS.

## 觀測（sanitized）

- Dogfood residual after TR knife：three workplace rows（es／fr presidente on
  `陈大总`；FIL `Si Panggawa` on `郑总`）. `总裁` CEO rows intentionally left.
- Product：`_workplace_boss_surnames` + `mechanical_romance_boss_title_gloss`；
  FIL gate／rewrite includes Panggawa；unit tests lock surname extract and
  总裁 skip.
- Offline apply cleared workplace presidente／Panggawa residual to 0.

## 未完成

- Broader semantic REVIEW（including intentional `总裁`→CEO）and burn／publish.
- Expand only from new dogfood bad surfaces.

## 連動

- Sibling：`2026-10-14-tr-boss-title-gloss-mechanical.md`、
  `2026-10-14-fil-phrase-polarity-mechanical.md`
- Product episode tables stay under `<PROJECT_ROOT>` analysis only.
