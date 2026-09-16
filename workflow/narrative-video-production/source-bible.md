# Source bible

一部**源作品**一份 bible。回答：這個世界裡有誰、哪一集、什麼地方、哪些 canonical entity。
**不是**可剪素材表（見 [`material-clip-catalog.md`](material-clip-catalog.md)）。Invariant 2。

填 [`records/source-bible.yaml`](records/source-bible.yaml)。

## 必填

| 欄位 | 規則 |
| --- | --- |
| `bible_id` | 穩定；多部成片共用 |
| `episodes[]` | `episode_id` |
| `characters[]` | `character_id` + 顯示名 + aliases |
| `places[]` | 可空；有則 `place_id` |

禁止：把整集 ASR／字幕全文當 bible；沒有 id 的散文人物。

推進：至少一個 `episode_id` 與一個 `character_id`（或書面 `no_character_entities: true`）。
失敗 rollback：`bible_owner`。
