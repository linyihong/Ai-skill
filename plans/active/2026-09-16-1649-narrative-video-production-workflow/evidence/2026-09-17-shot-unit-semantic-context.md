# Observation — shot unit semantic context

**Run ID**：2026-09-17-shot-unit-semantic-context  
**Kind**：Phase 3 observation（**不是** `contract_gap`；**不**擴 Phase 2／3 workflow contract）  
**Supersedes framing**：不是「台詞 + 查找語」；也不是只給 dialogue 加 search keyword。

## 抽象

每個需要 AI 理解／匹配的分鏡單元：

```text
Unit (dialogue | action | visual)
  text              → 給觀眾／編輯讀的自然敘事（可省略）
  semantic_context  → 已知可解析的機器語義（含 summary）
```

| 讀者 | 吃什麼 |
| --- | --- |
| Audience | `text` |
| Editor | `text` + `semantic_context` |
| Matching / AI | `semantic_context`（尤其 `summary` + entity 欄）→ Need → Constraints |

`summary` 保留語義，不是 `search_text` 關鍵字串。正式名仍是 **`semantic_context`**（不用 `retrieval_context`）。

Matching 仍：Need → Constraints → 可行集 → Selection Policy。Context 只進約束。

## 真實片子要數

- 哪些 shot 是 dialogue／action／visual  
- 各類有多少必須靠 `semantic_context` 才能對上 catalog  
- 穩定需要的欄位（speaker vs actor vs entities／location）是否收斂  

若幾乎所有分鏡描述都需要，再升成 `script-and-shot-list` 正式抽象。本 round 不改 workflow。
