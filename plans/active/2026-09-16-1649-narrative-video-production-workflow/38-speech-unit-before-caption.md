# Candidate: Speech unit before caption

Companion to [`29-speech-timing-authority.md`](29-speech-timing-authority.md)、
[`37-natural-boundary-before-length.md`](37-natural-boundary-before-length.md)。**workflow 已補**
（[`speech-unit.yaml`](../../../workflow/narrative-video-production/records/speech-unit.yaml)）。  
觀察：[`evidence/2026-09-28-speech-unit-before-caption.md`](evidence/2026-09-28-speech-unit-before-caption.md)。

Speech Unit 是語音／語意 SoT。Caption 是視覺投影。順序：標點候選切 unit（不看字級／行數）→ TTS → 實際 duration → cue 時間 → 只在該 unit 內換行。

標點不是強制切開；過短 unit 由 selection 合併，然後才生成語音。Layout 不得改 unit 原文。BreakCandidate 留在 caption layout。
