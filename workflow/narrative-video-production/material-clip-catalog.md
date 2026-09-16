# Material clip catalog

同一源作品一份 catalog，可增量。回答：實際有哪些**可定位**的剪輯窗。
**不是** bible。Invariant 2。

填 [`records/material-clip-catalog.yaml`](records/material-clip-catalog.yaml)。

## 每條 clip 必填

`clip_id`、`episode_id`、`source_in`、`source_out`、`duration_s`、`visual_description`、`tags[]`。
可選：`dialogue_excerpt`、`duration_band`（v0 名：`hook`／`beat`／`hold`；**秒數 dogfood Q10**）、`ingest.origin`。

禁止：沒有入出點的印象標籤；沒有 `clip_id` 的散文；把整部片當一條 clip。

## retrieval_contract（Invariant 3）

Canonical 查找面 = 本 catalog 的文字欄（id、aliases、visual_description、tags、duration）。

- 查詢結果必須是 **catalog 裡已存在的 `clip_id` 列表**。
- 向量／BM25／LLM retrieval 只是 adapter；仍必須回傳既有 `clip_id`。
- 不得回傳「再去片裡找找」或庫外虛構片段。

受控 tags：能連 entity 的用 `character_id`／`place_id`；其餘用詞表，禁止每次發明同義新 tag（Q9）。

推進：本片將引用的每個 `selected_clip_id` 都在 catalog。
失敗 rollback：`catalog_owner`。
