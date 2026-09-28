# Phase 1/3 reinforcement — JA pragmatic／lexical dogfood (ep7)

Companion to [`_plan.md`](_plan.md)。Evidence：[`evidence/2026-09-29-ep7-ja-dogfood.md`](evidence/2026-09-29-ep7-ja-dogfood.md)。

## Conclusion

Workflow direction stands. Nail **Semantic Analysis → Target Realization → Contextual／Pragmatic／Register／Modality validation**. Especially zh→ja:

1. Chinese lexicon ≠ Japanese kanji gloss  
2. Character relationship／address  
3. Chinese ellipsis／construction recovery  
4. Idiomatic／pragmatic rewrite  

**Do not** grow `if japanese: 老婆→…` prompt lists.

## New／strengthened contract concepts

| Concept | Role |
| --- | --- |
| `lexical_handling.mode` | literal／lexicalized／idiomatic／collocational／**false_cognate_risk**／culturally_bound |
| `construction`／`pragmatic_expression`／`speech_act_expression` | expression types |
| Reference／Identity Resolution | 姐夫／姐／老婆／赵伟 before surface |
| Modality／uncertainty | 该不会…吧 must not collapse to bare したの？ |
| `semantic_expansion` | naturalness must not add 振る舞い etc.（I18） |
| Register preservation | colloquial ≠ べし elevation（I20） |

## Invariants

- **I17** — Cross-lingual false cognate／lexical decomposition risk MUST be flagged; Selection MUST NOT default to character-wise kanji gloss.
- **I18** — Semantic expansion（unsupported content for fluency）is a distinct fail／review from marketing invented_information; both block accepted without waiver.
- **I19** — Preserve participants、event、modality／uncertainty、polarity（not only predicate）.
- **I20** — Meaning-correct + register-elevated／depressed → register fail or review.

## JA Failure Taxonomy（→ registry, not prompt）

| ID | Pattern | Example |
| --- | --- | --- |
| JA-F01 | chinese_lexical_false_cognate | 老婆 → お婆さん |
| JA-F02 | chinese_pragmatic_construction_literalization | 我说了她两句 → 言及しました |
| JA-F03 | semantic_role_or_modifier_loss | 你们该不会吵架了吧 → 喧嘩したの？ |
| JA-F04 | semantic_expansion | 我老婆也太反常了 → 妻の振る舞い… |
| JA-F05 | target_locale_residue | 姐夫 → 姐夫 |
| JA-F06 | foreign_name_realization | 赵伟 → 趙偉 only |
| JA-F07 | register_drift | 良田… → 耕すべし |
| JA-F08 | idiom_literalization | 又装矜持 → 矜持を装う |
| JA-F09 | pragmatic_ambiguity | 应酬 → 接待 |
| JA-F10 | modality_loss | 该不会…吧 → したの？ |

Disposition：Constraint／Mechanical vs Selection／Knowledge — per pattern `response` in registry.

## Landing

| 產物 | 路徑 |
| --- | --- |
| Evidence | `evidence/2026-09-29-ep7-ja-dogfood.md` |
| Patterns | `registry/failure-patterns.yaml`（JA-F*） |
| Analysis | `lexical_handling` + modality + discourse layers |
| Validation | semantic_expansion／modality／register_drift rules |
| Fixtures | `laopo-false-cognate.yaml`、`shuoleshe-liangju.yaml` |
