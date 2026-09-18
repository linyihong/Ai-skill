# Candidate: Material fact extraction before semantic inference

Companion to [`_plan.md`](_plan.md)。**候選原則，不是 workflow schema。** Phase 3 **不**把下列分析器寫進 `narrative-video-production`。  
觀察：[`evidence/2026-09-17-material-fact-extraction.md`](evidence/2026-09-17-material-fact-extraction.md)。

原則：劇情分析應先最大化可機械取得的素材證據，再進行語義推理。

```text
layer.observable   shot / keyframe / visual_text+geometry / ASR+transcript / voice_track / speaker_id / face_track / audio
layer.linking      ASR↔OCR↔Voice↔Face↔Shot（共現，非等同）
layer.canonical    identity / dialogue / scene / naming（仲裁後）
layer.narrative    script / template / matching / selection / EDR
```

Stage A（事實層）禁止 LLM 做主判斷。Stage B（敘事理解）才做誰是誰；身份可 unnamed：[`08-identity-precedes-naming.md`](08-identity-precedes-naming.md)。  
OCR 在此層是取得 **visual text evidence** 的一種方法（必帶 normalized box；不分類字幕／浮水印）：[`09-visual-text-evidence.md`](09-visual-text-evidence.md)。  
Face Track 屬 observable；linking 層掛 ASR／OCR／Shot，**不做** Recognition→角色：[`11-face-as-candidate-evidence.md`](11-face-as-candidate-evidence.md)。  
Voice／Speaker 與 ASR transcript 分開掛，**不做** diarization→角色：[`14-voice-speaker-evidence.md`](14-voice-speaker-evidence.md)。  
採集之後的仲裁／審查／政策學習：[`12-evidence-refinement.md`](12-evidence-refinement.md)（獨立於 parser）。  
掃區是 Mechanical Probe，不是 LLM crop：[`13-mechanical-visual-text-probe.md`](13-mechanical-visual-text-probe.md)。  
外部工具產出當 **evidence candidates**；真實片子再數哪些真的被 bible／catalog／matching／EDR 消費。Shot／scene 關係見 [`10-editorial-vs-narrative-transition.md`](10-editorial-vs-narrative-transition.md)。
