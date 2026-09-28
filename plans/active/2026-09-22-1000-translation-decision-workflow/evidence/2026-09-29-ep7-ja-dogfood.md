# Run — ep7 zh→ja dialogue dogfood

**Run ID**: `2026-09-29-ep7-ja-dogfood`  
**Status**: observed → Phase 1 contract reinforcement  
**Related**: [`../12-ja-pragmatic-lexical.md`](../12-ja-pragmatic-lexical.md)、[`failure-patterns`](../../../workflow/translation/registry/failure-patterns.yaml)

## Verdict

Not “a few typos”. ≥7–9 structural failure modes. Direction of workflow OK; nail **Semantic Analysis → Target Realization → Pragmatic／Register／Modality validation**. Do **not** append prompt cases.

## Segment map（sanitized）

| Source | Observed dst | Gate | Pattern |
| --- | --- | --- | --- |
| 今天晚上定让你跪地求饶 | 今夜は間違いなく、あなたを膝をつかせます | review | JA-F03／pragmatic incomplete |
| 老婆今晚给你开开荒 | お婆さんが今夜、開拓をしてくれますよ | **fail** | JA-F01 false cognate + literalization |
| 你别碰我 | 触らないで | pass | — |
| 又装矜持 | また、矜持を装う | **fail** | JA-F08／JA-F02 |
| 良田就得多耕 | 良田はよく耕すべし | review | JA-F07 register_drift |
| 容易长草 | すぐに飽きてしまう | review | JA-F09 ambiguity |
| 我说了别碰我 | 触らないでと言ったのに | review | JA-F10／speech_act |
| 赵伟是谁 | 趙偉とは誰ですか | review | JA-F06 name realization |
| 赵伟是我同事 | 趙偉は私の同僚です | review | JA-F06 + register |
| 今天我们公司聚餐 | 今日は会社の飲み会です | pass／review | — |
| 你早点休息吧 | 早めに休んでね | pass | — |
| 我老婆也太反常了 | 妻の振る舞いが異常すぎる | **fail** | JA-F04 semantic_expansion |
| 她该不会是出轨了吧 | 彼女、浮気してない？ | review | JA-F10 modality_loss |
| 姐夫 | 姐夫 | **fail** | JA-F05 target residue |
| 刚才听到我姐的声音 | さっき、姉さんの声を聞いた | pass | — |
| 你们该不会吵架了吧 | 喧嘩したの？ | review | JA-F03＋JA-F10 |
| 你姐要出去应酬 | お姉さんが外で接待することになりました | **fail** | JA-F09 应酬≠接待 |
| 我说了她两句 | 彼女に少し言及しました | **fail** | JA-F02 construction |
| 姐夫你的脸怎么了 | 姐夫、あなたの顔は… | **fail** | JA-F05 + unnatural |
| 没事我不小心撞的 | 気にしないで、… | review | naturalness |

## Contract takeaways

1. **Chinese ≠ Japanese hanzi friends** — 老婆／矜持 need `lexical_handling.mode=false_cognate_risk`／lexicalized.
2. **Pragmatic constructions** — 说了她两句／装矜持／该不会…吧 are not dictionary glosses.
3. **Semantic expansion** — 振る舞い is invented content (I18), not “more natural Japanese”.
4. **Residue** — 姐夫 untranslated → mechanical／locale fail.
5. **Name** — 赵伟 → 趙偉 needs identity + series cast + script policy, not auto-kanji.
6. **Register** — べし elevates colloquial proverb → register fail/review.

## Acceptance

- [x] Mapped to JA-F01–F10 without prompt if-blocks  
- [x] Phase 1 contracts／registry updated（[`12`](../12-ja-pragmatic-lexical.md)）  
- [ ] Consumer re-run with guards（out of band）
