# Candidate: Material fact extraction before semantic inference

Companion to [`_plan.md`](_plan.md)。**候選原則，不是 workflow schema。** Phase 3 **不**把下列分析器寫進 `narrative-video-production`。  
觀察：[`evidence/2026-09-17-material-fact-extraction.md`](evidence/2026-09-17-material-fact-extraction.md)。

原則：劇情分析應先最大化可機械取得的素材證據，再進行語義推理。

```text
layer.observable   shot / keyframe / ASR / OCR+geometry / speaker_id / face_track / audio
layer.linking      cross-modal evidence_link（ASR↔OCR↔Face↔Shot；非仲裁器）
layer.canonical    identity / dialogue / scene / naming（仲裁後）
layer.narrative    script / template / matching / selection / EDR
```

Stage A（事實層）禁止 LLM 做主判斷。Stage B（敘事理解）才做誰是誰；身份可 unnamed：[`08-identity-precedes-naming.md`](08-identity-precedes-naming.md)。  
OCR 在此層是 **visual text** 候選來源，不是 locale 字幕的同義詞：[`09-visual-text-evidence.md`](09-visual-text-evidence.md)。  
Face Track 屬 observable；linking 層掛 ASR／OCR／Shot，**不做** Recognition→角色：[`11-face-as-candidate-evidence.md`](11-face-as-candidate-evidence.md)。  
採集之後的仲裁／審查／政策學習：[`12-evidence-refinement.md`](12-evidence-refinement.md)（獨立於 parser）。  
外部工具產出當 **evidence candidates**；真實片子再數哪些真的被 bible／catalog／matching／EDR 消費。Shot／scene 關係見 [`10-editorial-vs-narrative-transition.md`](10-editorial-vs-narrative-transition.md)。
