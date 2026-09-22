# Candidate: Speech timing authority

Companion to [`01-captions-and-locales.md`](01-captions-and-locales.md)。**workflow 已補契約**（Speech Unit → Speech Timing → Caption Cue）。TTS 仍是 adapter。  
觀察：[`evidence/2026-09-22-speech-timing-authority.md`](evidence/2026-09-22-speech-timing-authority.md)。

Locale pack 原先偏字幕驗收。自製口播的 cue 時軸 **primary source = speech artifact duration**（源片則 ASR）。先 Speech Unit Planning（語意＋screen-fit），再逐 unit 生成；禁止整段 TTS 再回頭切字幕、禁止猜秒數。Speech QC ≠ layout QC。兩個 loop 分開。見 [`speech-unit-and-timing.md`](../../../workflow/narrative-video-production/speech-unit-and-timing.md)。
