# Run — ep8 zh→ja static walkthrough (workflow dogfood)

**Run ID**: `2026-09-29-ep8-ja-walkthrough`  
**Status**: walkthrough complete → **needs human review**（finality 多數 `needs_review`）  
**Target**: `ja-JP` · content.type=`subtitle`  
**Taxonomy bound**: `cross_locale_core` + **JA-F01–F10** via [`locale-failure-taxonomy.yaml`](../../../workflow/translation/registry/locale-failure-taxonomy.yaml)  
**TDR pack**: [`examples/ep8-ja-walkthrough.yaml`](../../../workflow/translation/examples/ep8-ja-walkthrough.yaml)

## Context binding

```text
translation_context:
  source: { language: zh, locale: zh-CN }
  target: { language: ja, locale: ja-JP, script: Jpan }
  content: { type: subtitle, decision_mode: faithful_translation }
reference_resolution (discourse):
  姐 / 我姐 → elder_sister (speaker's)
  姐夫 → sister's_husband (addressee or referent)
  你姐 → addressee's_wife / elder_sister depending on speaker
```

## Segment decisions（selected = illustrative Selection under guards）

| # | Source | Selected (ja) | Finality | Guards hit / notes |
| --- | --- | --- | --- | --- |
| 1 | 这是巴掌印 | これは平手打ちの跡だ | needs_review | literal OK；確認「巴掌印」語感 |
| 2 | 难道是我姐打的 | まさか、姉さんが殴ったの？ | needs_review | JA-F10 难道 modality 保留 |
| 3 | 你姐她也不是故意的 | お姉さんもわざとじゃないよ | needs_review | 你姐→お姉さん／妻（依說話者）；discourse |
| 4 | 我姐怎么可以这样 | 姉さんがこんなことするなんて | needs_review | 怎么可以 = 非難 speech-act |
| 5 | 她怎么可以打你 | 彼女があなたを殴るなんて | needs_review | 同上 |
| 6 | 我不怪她 | 彼女のことは責めない | needs_review | — |
| 7 | 我去帮你拿点药 | 薬を取ってきてあげる | needs_review | — |
| 8 | 我帮你 | 手伝うわ／僕がやる | needs_review | 性別／角色 register 需 cast |
| 9 | 姐夫好看吗 | 義兄さん、似合ってる？ | needs_review | **JA-F05**：禁止留「姐夫」 |
| 10 | 我姐脾气暴躁 | 姉さんは気が短い | needs_review | — |
| 11 | 我替我姐向 | — | **blocked** | source **truncated**；不可補全 invented |
| 12 | 这丫头竟然发育的这么好 | この娘、まさかこんなに発育がいいとは | needs_review | 竟然 modality；register／敏感 |
| 13 | 糟了 | まずい | pass／review | — |
| 14 | 药酒的劲好像发作 | 薬酒の効き目が出てきたみたい | needs_review | 好像 modality；「发作」≠病發硬翻 |
| 15 | 姐夫下面好像有东西 | 義兄さん、下の方が当たってるみたい | needs_review | JA-F05 + euphemism；禁止 residue |
| 16 | 顶到我了 | 当たっちゃった | needs_review | — |
| 17 | 你先回房休息 | 先に部屋で休んでて | needs_review | — |
| 18 | 我去替你收拾房间 | 部屋、片付けておくね | needs_review | — |
| 19 | 姐夫 | 義兄さん | needs_review | **JA-F05** 必過 |
| 20 | 我姐是女强人 | 姉さんはバリバリのキャリアウーマンだ | needs_review | JA-F09 女强人 ≠ 女の強人 kanji |
| 21 | 一心都在工作上 | 仕事一筋で | needs_review | — |
| 22 | 你平时肯定没少受委屈 | 普段、ずいぶん我慢してるんでしょ | needs_review | 肯定／没少 = 推測＋程度 |
| 23 | 受点委屈倒是没什么 | 少し我慢するくらいならいいけど | needs_review | — |
| 24 | 我只怕你姐 | — | **blocked** | truncated；等下一句 |
| 25 | 怕我姐什么 | 姉さんの何が怖いの？ | needs_review | — |
| 26 | 怕你姐给我戴绿帽子 | お姉さんに浮気されるのが怖い | needs_review | **JA-F08**：戴绿帽子 ≠ 緑の帽子 |

## Anti-patterns explicitly rejected

| Bad dst | Why |
| --- | --- |
| 姐夫（未譯） | JA-F05 fail |
| 緑の帽子をかぶせる | JA-F08 idiom literalization |
| 女の強人 | JA-F01／F09 false／pragmatic |
| 補全「我替我姐向你道歉」無 evidence | I18 semantic_expansion／invented |

## Human review checklist

- [ ] 義兄さん vs お姉さん的な呼び／キャスト既定名是否一致  
- [ ] 說話者性別（我帮你／似合ってる）是否對角色  
- [ ] 敏感句（15–16、12）是否符合作品尺度／平台政策  
- [ ] Truncated #11／#24 是否應與相鄰 cue merge（NVP speech-unit）再決策  

## Acceptance

- [x] Locale taxonomy bound for ja-JP in walkthrough  
- [x] TDR pack landed under `workflow/translation/examples/`  
- [ ] Human reviewer sign-off（pending）
