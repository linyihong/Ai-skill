# Adapters — image-processing

Adapters bind **abstract candidate_ids** to concrete tools.
Canonical lifecycle stays in [`../execution-flow.md`](../execution-flow.md)；本目錄可提工具名。

| Adapter | Owns | Does not own |
| --- | --- | --- |
| [`bg-removal-vertical-slice.md`](bg-removal-vertical-slice.md) | A 類最小閉環步驟、本機 Python 基準候選綁定 | 通用契約欄位、finality 規則 |
| [`license-checklist.md`](license-checklist.md) | Q5 商用／開源授權核對 | 法律意見本身 |
| [`../tools/`](../tools/README.md) | Independent verify + fixture 生成（機械） | Selection 決策、人審語意 |

**Independent verify 必須是獨立進程／腳本**，不得由執行腳本結尾 `print("OK")` 代替。
