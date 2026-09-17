# Observation — dialogue semantic ambiguity

**Run ID**：2026-09-16-dialogue-semantic-context  
**Kind**：Phase 3 observation（審查提煉，**不是**真實片子 dogfood，**不是** `contract_gap`）  
**Status**：optional `dialogue.semantic_context` 已在 workflow；**更廣的 unit 抽象**見後續 observation，本檔不單獨升 contract。

後續收窄：[`2026-09-17-shot-unit-semantic-context.md`](2026-09-17-shot-unit-semantic-context.md)（dialogue／action／visual 都是 text + semantic_context）。

## 觀察名

`dialogue_semantic_ambiguity` / retrieval context gap

短觀眾台詞省略主詞時，機器不能靠加長台詞去猜「你／他」是誰。結構化補充叫 `semantic_context`，不叫查找語。

## 處置（本 round）

- 落點：`script-and-shot-list` → matching（進 Constraints）→ EDR 痕跡
- **optional**：有 `semantic_context` 才必填 `speaker` + `summary`
- 條件欄：`addressee`／`subject_refs`／`object_refs`／`intent`；未知則省略，禁止 `unknown`
- 不破壞 Selection：不代替 `selection.policy`
- 真實片子需計：有對白的 shot 裡，多少必須靠 `semantic_context` 才能對上 catalog

## 升格規則

| 計數 | 下一步 |
| --- | --- |
| 幾乎每段有對白的 shot 都需要 | 考慮升成有對白時的 contract |
| 少數特例 | 維持 optional |
