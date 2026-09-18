# Real-run review — narrative representation gap

**Run ID**：2026-09-18-narrative-representation-gap
**Kind**：去敏真實 source-analysis 複核（不是完整 EDR run）
**Conclusion**：Evidence Extraction 可用；Dialogue → Story Event 的直接路徑
不足。瓶頸是中間的 Narrative Representation，不是 detector 或模型數量。

原始媒體、逐字內容、劇名、人物與路徑留在外部專案。本檔只保留可重用的
失效類型。

## Human comparison

人工合看 ASR 與字幕型 visual text 後，可辨識出關係訊號、互動升級、情境中斷、
後續延續與揭示 setup。既有輸出只保留極少 story candidates，而且主要候選仍
使用錯誤 ASR 語義。

關鍵 visual text 已被 observable 層抽取，故缺口不是「沒有資料」。既有
selection policy 沒有穿透到 narrative consumer；下游仍直接使用 raw ASR。

## Findings

### F11 — selected text not consumed downstream

ASR 與 subtitle candidate 衝突時，event semantics 仍採 raw ASR。需要
role-qualified、span-level `text_resolution`，保留 alternatives 與來源；
不得簡化為全域 OCR 優先。

### F12 — event representation is dialogue-centric

Event candidates 主要是 dialogue grouping，未表達 interaction、relationship、
action、situation、revelation 或 transition。這是 segmentation 的重新命名，
不是已解析事件。

### F13 — local relevance cannot assemble context

多個單句各自為 medium，但合看前後 dialogue、visual text、shot、participant
與 action 後形成明顯 escalation／interruption。逐句 relevance 無法取代
contextual assembly。

### F14 — event relations are absent

平行 event rows 無法表達 continuation、escalation、interruption、consequence、
revelation 或 setup。Relation 必須有 evidence 與 resolution status；相鄰不等於
因果。

### F15 — story-state vocabulary is under-specified

空 `before`／`after` 不可升格，但不是每個候選都應被迫填完整世界狀態。先觀察
typed claims：relationship、interaction、situation。只有宣稱 state change 時，
才要求具體 subject／change／可驗證前後或新增事實。

## Phase 3 assessment

- Observable extraction、OCR geometry、Evidence Unit、traceability：可繼續。
- ASR ↔ subtitle text resolution：partial；未穿透 narrative。
- Event Candidate：fail；過度 dialogue-centric。
- Narrative Window／Assembly、Narrative Relation：missing。
- Story Event／Story State：fail；上游 representation 不足。

這不是「Story Event 數量太少」或 threshold tuning。下一步只 dogfood
Text Resolution → Narrative Window → Event Assembly → Narrative Relation →
Story State，不新增 Agent、不擴 detector、不立即改 workflow。
