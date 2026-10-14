# Latin MT／OCR letter-space glue

## 決策

1. **Target-script admission ≠ word integrity** for Latin locales when MT／OCR
   inserts a space inside a common function word（`a nd`／`A fter`／`to gether`）.
2. Prefer a **whitelist fragment rejoin**（left＋right → joined）over dictionary
   guessing that would smash legitimate bigrams（`I was`／`a lot`）.
3. Preserve left-fragment case：single-letter `A` → title case（`After`），not
   ALL CAPS（`AFTER`）；multi-letter ALL CAPS left may stay upper.
4. Apply on Latin-script targets only；surface repair shares the mechanical_*
   API with siblings（ZH unused，kept for parity）.
5. Content Finality stays **REVIEW**；glue clears dogfood surface junk，not
   semantic／publish PASS.

## 觀測（sanitized）

- Postfix EN caches on a multi-series overnight pack showed intra-word spaces
  on whitelist fragments；romanization Liu hits that remained were correct 柳→Liu，
  not new 陆 confusions.
- Product：`mechanical_latin_letter_space_glue` after latin name consistency on
  the validate／repair path；unit tests lock seed splits，title-case，legit
  bigram skip，and non-Latin skip.
- Offline apply on two EN episode caches cleared whitelist residual to 0；ASS
  Dialogue texts rewritten in lockstep with cue counts.

## 未完成

- Expand whitelist only from new dogfood fragments（do not invent broad glue）.
- TR boss-title gender／office gloss（e.g. Bayan for 陆总）still open REVIEW.
- Publish／layout Finality still open.

## 連動

- Sibling：`2026-10-14-fil-phrase-polarity-mechanical.md`（mechanical expression
  repair class）
- Product episode tables stay under `<PROJECT_ROOT>` analysis only.
