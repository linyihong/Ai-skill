# Translation examples & regression fixtures

Phase 1 = **doc-only**. No test runner.

| File | Role |
| --- | --- |
| [`address-title-chen-xiaojie-id.yaml`](address-title-chen-xiaojie-id.yaml) | **P0** — id-ID title + `target_locale_residue` |
| [`address-title-chen-xiaojie-ja.yaml`](address-title-chen-xiaojie-ja.yaml) | **P0** — ja-JP **name realization**（Chenさん → review） |
| [`title-kongjie-yiriqianli.yaml`](title-kongjie-yiriqianli.yaml) | **P0** — short-drama **title**；idiom + invented_information（I12） |
| [`social-address-laozhang.yaml`](social-address-laozhang.yaml) | **P0** — 老張；F5 social_address_misinterpretation（≠ オヤジ） |
| [`literal.yaml`](literal.yaml) | literal |
| [`idiom.yaml`](idiom.yaml) | idiom |
| [`slang.yaml`](slang.yaml) | slang |
| [`proverb.yaml`](proverb.yaml) | proverb |
| [`dialect.yaml`](dialect.yaml) | dialect |
| [`wordplay.yaml`](wordplay.yaml) | wordplay |

## Static walkthrough A（陈小姐 → id-ID）

1. context.target = `id-ID`（authoritative）  
2. Analysis splits 陈／小姐  
3. Candidate Space titles（Nona, Miss, Ms.）  
4. Feasible：`Nona Chen` true；`Miss`/`Ms.` false  
5. Policy + decision_basis  
6. `Miss Chen` → locale_consistency=review → not accepted without waiver  
7. `Nona Chen` → accepted if I9  

## Static walkthrough B（陈小姐 → ja-JP）

1. context.target = `ja-JP` + realization_profile  
2. Analysis：proper_name → `name_realization`；title → locale_aware  
3. Name Candidate Space：チェン／陳／Chen；Title：さん／ミス  
4. Composed：チェンさん／陳さん／Chenさん  
5. `Chenさん` → **name_realization=review**（title OK；≠ locale_consistency residue）  
6. `チェンさん` → preferred pass → accepted if I9  
7. **禁止**非片假一律 FAIL  

A+B **PASS**（2026-09-22）：[`08-static-walkthrough-pass.md`](../../../plans/active/2026-09-22-1000-translation-decision-workflow/08-static-walkthrough-pass.md)。

## Static walkthrough C（空姐被一日千里-上）

1. `content.type=title` + `subtype=short_drama_title`  
2. Title Structure：idiom + part_marker + nonstandard  
3. 一日千里 → Candidate Space（急速な進展／目まぐるしい展開…），非唯一 Rapid Changes  
4. `#ギリギリのセクシーなドラマ` → `invented_information=fail`（I12）  
5. Observed model dst → **not** `finality.accepted`  

見 [`title-kongjie-yiriqianli.yaml`](title-kongjie-yiriqianli.yaml)、[`09-title-content-type.md`](../../../plans/active/2026-09-22-1000-translation-decision-workflow/09-title-content-type.md)。

## Static walkthrough D（老張）

1. Analysis：`老` = social_address_marker；`張` = proper_name  
2. Candidate Space：張さん／張／チョウさん；**オヤジ張さん** infeasible（F5）  
3. I14–I16：preserve social relation；naturalness 不得改 identity  

見 [`social-address-laozhang.yaml`](social-address-laozhang.yaml)、[`11`](../../../plans/active/2026-09-22-1000-translation-decision-workflow/11-semantic-syntactic-realization.md)。
