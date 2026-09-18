# Candidate: Voice / speaker evidence

Companion to [`_plan.md`](_plan.md)。**候選掛點，不是 workflow。** 正式名是 **Voice／Speaker Evidence**，不是「ASR 判斷人物」。Phase 3 **不**改 `narrative-video-production`、不接聲紋辨識產品。  
觀察：[`evidence/2026-09-18-voice-speaker-evidence.md`](evidence/2026-09-18-voice-speaker-evidence.md)。

原則：ASR 負責「說了什麼」；Voice／Speaker 負責「誰在說」（本段音訊的 speaker cluster）。`speaker_id` ≠ `character_id`，與 `face_track` ≠ `character_id` 同一條：[`11-face-as-candidate-evidence.md`](11-face-as-candidate-evidence.md)。

`asr_segment` 掛 `speaker_id`／`voice_evidence_ref`，以及同時段的 `visual_text_refs`／`face_track_refs`。**只記時間重疊，不在 ASR 階段寫 speaker = face = character。** Identity Resolution 才組合。

無對白人名時，不同 `speaker_id` 仍可標出不同說話者，供後續人物分析。變聲、一演員多角、配音、誤切 speaker，都禁止 `voice_04 = character_lin`。Phase 3 **不**改 workflow。
