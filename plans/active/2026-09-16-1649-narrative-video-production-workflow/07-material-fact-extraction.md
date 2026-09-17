# Candidate: Material fact extraction before semantic inference

Companion to [`_plan.md`](_plan.md)。**候選原則，不是 workflow schema。** Phase 3 **不**把下列分析器寫進 `narrative-video-production`。  
觀察：[`evidence/2026-09-17-material-fact-extraction.md`](evidence/2026-09-17-material-fact-extraction.md)。

原則：劇情分析應先最大化可機械取得的素材證據，再進行語義推理。

```text
layer.observable   shot / ASR+timeline / speaker_id / face cluster / OCR / audio segments
layer.identity     series_cast / entity / semantic_context
layer.decision     template / matching / selection / EDR
```

Stage A（事實層）禁止 LLM 做主判斷。Stage B（敘事理解）才做誰是誰、劇情、template。  
`speaker_id` ≠ `character_id`；`person_cluster` ≠ 角色；vision embedding ≠ 「悲傷場景」。  
外部工具產出當 **evidence candidates**；真實片子再數哪些真的被 bible／catalog／matching／EDR 消費。
