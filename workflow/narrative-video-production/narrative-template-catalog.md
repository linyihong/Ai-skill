# Narrative template catalog

模板是一級 artifact。每部成片必有**一個主** `narrative_template_id`；次模板只解釋局部。
禁止執行時發明匿名模板。

填／對照 [`records/narrative-template-catalog.yaml`](records/narrative-template-catalog.yaml)。

## v0 種子（experimental；不拆 taxonomy）

每個 id 必填 `template_kind: structure | mechanism | format`（分類槽，不是三套目錄）。

| id | 建議 kind | 備註 |
| --- | --- | --- |
| `hero_journey` | structure | |
| `mystery_box` | structure | |
| `cold_open_hook` | mechanism | |
| `identity_clash` | mechanism | |
| `rising_conflict` | mechanism | |
| `list_payoff` | mechanism | |
| `recap_explained` | format | |

七個可共存。Phase 3 有真實 EDR 前不拆 catalog。去留 = experimental。

推進：主 id ∈ 本 catalog 且 `template_kind` 已填。
失敗 rollback：`template_author`。
