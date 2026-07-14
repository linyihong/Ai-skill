# UI Pattern Knowledge — reusable rule seeds（Ai-skill）

本目錄存放 **Ai-skill reusable 規則種子**（首波：`overlay` family Core entries）。

| 誰擁有 | 內容 |
| --- | --- |
| **Ai-skill** | Core selection / family / neighbors / intent_examples；optional Extended 雛形 |
| **Consumer 專案** | 採用名、`platform_map.project`、composition、project alias；可覆寫 / 延伸 entries |

使用方式：

1. 選型不明或宣告 pattern 對齊時，先讀 [`../ui-contracts.md`](../ui-contracts.md) **Pattern Knowledge Lock**。
2. 複製 [`../templates/ui-pattern-knowledge.entry.template.yaml`](../templates/ui-pattern-knowledge.entry.template.yaml) 擴充；**完整 entry = Core 齊**。
3. 專案 overlay 放 consumer repo（例如 `docs/frontend-contracts/ui-pattern-knowledge/`），不要把專案私有名寫進本目錄。

**禁止**：`platform_map` 使用 Material / iOS / Fluent 等 DS 百科 key；不做獨立 Intent DB。
