# Observation — evidence unit before narrative

**Run ID**：2026-09-18-evidence-unit  
**Kind**：Phase 3 observation（**不是** `contract_gap`；**不**擴 workflow；**不**再加 observable detector）  
**原則**：Keep the existing evidence chain. Aggregate same-window facts into an evidence unit before any plot LLM. Observable objects store what they are; relations live on the linking layer.

已驗證方向（保留）：normalized box、spatial／temporal features、機械 watermark／subtitle **candidate**、ASR word timing、`speaker_id`／`voice_evidence_ref`、face track、ASR↔OCR／ASR↔Face links、selection policy + ambiguity escalation。機械把邊角高 persistence 字判成 watermark candidate，正是幾何＋持續性，不是再加模型。

## 缺的是聚合，不是更多採集

`asr_ocr_fuse` 夠辨識「誰／什麼字」。直接跳劇情分析會把閒聊當劇情。先組：

```yaml
evidence_unit:
  id: eu_001
  time_range: { start: 123.4, end: 125.8 }
  shot_refs: [shot_025]
  dialogue: { asr_refs: [asr_00182], visual_text_refs: [ocr_00182] }
  speaker_refs: [speaker_01]
  face_refs: [face_023]
  relations: [temporal_overlap, dialogue_visual_match]
  resolution: { dialogue_source: ocr }
```

這只表示「同一事件窗口」，不是劇情。下一層 `event_candidate` 才帶 `type`／summary，且 `narrative_relevance.status: unresolved`。閒聊仍完整保留 observable + unit，relevance=low、無 state change → 不進主劇情。

## 反向連結不要長在每個採集物件上

現在 ASR 掛 `visual_text_refs`／`face_track_refs` 可用。資料變大後：採集物件只保存自身；**relations 集中在 linking／evidence_unit**。現有結構不必立刻改；固化 schema 時再收斂。

## 真實片子要數（優先於加欄位）

閒聊、關鍵對白、跨 shot 對白、字幕／ASR 衝突、人物切鏡：現有 evidence 能否支撐 evidence_unit → event → relevance，而不是再加 detector。
