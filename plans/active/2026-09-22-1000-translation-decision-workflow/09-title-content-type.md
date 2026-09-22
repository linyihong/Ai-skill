# Phase 1/3 reinforcement — Title / content_type

Companion to [`_plan.md`](_plan.md)。**不改 Phase 0 freeze**；不成新 workflow。

## Dogfood signal

```text
空姐被一日千里-上
→ Airhostess Falls Victim to Rapid Changes - Part One
  #shortdrama #ギリギリのセクシーなドラマ
```

| 問題 | 契約落點 |
| --- | --- |
| 一日千里 → Rapid Changes | idiom ambiguity；需 Candidate Space（I5） |
| #ギリギリのセクシーなドラマ | **invented_information**（I12） |
| 當普通句子翻 | 缺 `content.type=title` + Title Structure Analysis |
| Part One | part_marker normalization — OK class |

## Lifecycle 插入

```text
Locale Resolution
  → Content-Type / Context Resolution   ← NEW (not new workflow)
  → Title Structure Analysis (when title)
  → Expression Analysis
  → …
```

## Modes

| Mode | `content.decision_mode` | 允許 | 禁止 |
| --- | --- | --- | --- |
| A | `faithful_translation` | semantic equivalent | new plot／emotion／marketing |
| B | `title_adaptation` | 重構句型、短劇 hook 風格 | 仍禁止 unsupported marketing |

`allow_marketing_expansion` 預設 false；需 upstream `marketing_tags`。

## Landing

| 產物 | 路徑 |
| --- | --- |
| content.type／decision_mode | [`translation-context.yaml`](../../workflow/translation/contracts/translation-context.yaml) |
| title_structure | [`expression-analysis.yaml`](../../workflow/translation/contracts/expression-analysis.yaml) |
| I12 | [`translation-decision.yaml`](../../workflow/translation/contracts/translation-decision.yaml)、[`validation.yaml`](../../workflow/translation/contracts/validation.yaml) |
| types／strategies | expression-types + `title_semantic_adaptation` |
| P0 fixture | [`title-kongjie-yiriqianli.yaml`](../../workflow/translation/examples/title-kongjie-yiriqianli.yaml) |

證明：`invented_information`、idiom、selection policy、content context 會在 production dogfood 互相碰撞——不是理論。
