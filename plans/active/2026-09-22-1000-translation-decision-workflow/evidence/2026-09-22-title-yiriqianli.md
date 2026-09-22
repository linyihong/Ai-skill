# Run — short-drama title「空姐被一日千里-上」

**Run ID**: `2026-09-22-title-yiriqianli`  
**Status**: observed → contract FAIL／NEEDS_REVIEW mapping  
**Related**: [`../09-title-content-type.md`](../09-title-content-type.md)、[`../../../workflow/translation/examples/title-kongjie-yiriqianli.yaml`](../../../workflow/translation/examples/title-kongjie-yiriqianli.yaml)

## Observed

| Field | Value |
| --- | --- |
| source | 空姐被一日千里-上 |
| model dst | Airhostess Falls Victim to Rapid Changes - Part One #shortdrama #ギリギリのセクシーなドラマ |

## Contract mapping

| Observation | Classification | Gate |
| --- | --- | --- |
| 空姐 → Airhostess | occupation render | OK class |
| 被 → Falls Victim to | over-narrativized passive | review under title policy |
| 一日千里 → Rapid Changes | idiom gloss without Candidate Space | `semantic=review`（idiom_semantic_uncertain） |
| 上 → Part One | part_marker | OK class |
| #shortdrama | platform metadata | consumer policy |
| #ギリギリのセクシーなドラマ | **invented_information** | **fail**（I12） |

## Verdict

Observed dst **must not** `finality.accepted`：

1. `invented_information=fail`（セクシー claim unsupported）  
2. idiom path lacks ambiguity Candidate Space → at best `needs_review`

Root cause for workflow：missing **content.type=title** + Title Structure Analysis — not「日文規則」or「再調 prompt」。

## Acceptance

- [x] Mapped to I12 + title content_type without new workflow  
- [x] Fixture landed under `workflow/translation/examples/`  
- [ ] Consumer re-run with title policy（out of band）
