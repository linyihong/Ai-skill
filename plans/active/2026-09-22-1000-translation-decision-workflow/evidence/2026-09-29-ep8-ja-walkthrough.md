# Run — ep8 zh→ja static walkthrough (workflow dogfood)

**Run ID**: `2026-09-29-ep8-ja-walkthrough`  
**Status**: retuned PASS／REVIEW／BLOCK（[`13`](../13-ep8-contract-revision.md)）  
**Target**: `ja-JP` · [`ep8-ja-walkthrough.yaml`](../../../workflow/translation/examples/ep8-ja-walkthrough.yaml)  
**Taxonomy**: core + **CORE-F21** + JA-F01–F10

## Disposition summary（after I22）

| Disposition | Count | Segments |
| --- | --- | --- |
| **PASS** `accepted` | 19 | 1–2, 4–7, 9–10, 13–14, 17–23, 25–26 |
| **REVIEW** `needs_review` | 5 | 3（reference）、8（voice）、12／15／16（sensitive） |
| **BLOCK** `blocked` | 2 | 11／24 truncated（I21） |

## Key regressions

| Case | Outcome |
| --- | --- |
| 糟了 → まずい | PASS |
| 一心都在工作上 → 仕事一筋で | PASS（非逐字 ≠ REVIEW） |
| 女强人 → キャリアウーマン | PASS semantic-equivalent（seed ≠ sole） |
| 戴绿帽子 → 浮気 | PASS JA-F08 |
| 姐夫 → 義兄さん via reference | PASS JA-F05＋I23 |
| 你姐 → お姉さん／妻 | REVIEW `reference_ambiguous` |
| 我帮你 | REVIEW `character_voice_unknown` |
| 我替我姐向 | BLOCK source_truncated |

## Human checklist

- [ ] Confirm PASS list against cast（especially 9／19 kinship form）  
- [ ] Resolve #3 reference（お姉さん vs 妻）  
- [ ] Bind #8 speaker gender／voice  
- [ ] Platform gate on #12／15／16  

## Acceptance

- [x] Locale taxonomy bound  
- [x] Finality ternary applied（not blanket REVIEW）  
- [ ] Human sign-off（pending）
