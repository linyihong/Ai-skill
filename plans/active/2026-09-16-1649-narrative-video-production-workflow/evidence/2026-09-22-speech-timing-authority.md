# Observation — speech artifact is cue timing SoT

**Run ID**：2026-09-22-speech-timing-authority  
**Kind**：representation／contract improvement（使用者授權補 workflow 契約；**不**把 TTS／ffmpeg 寫進 execution-flow 新 stage）  
**Extends**：[`01-captions-and-locales.md`](../01-captions-and-locales.md)、[`2026-09-22-subtitle-layout-engine.md`](2026-09-22-subtitle-layout-engine.md)

## 缺口

既有 CPS／cue 窗／折行是驗收。缺少 **誰是時間軸 authority**。AI 猜「大概 3 秒」不得當 generated-speech cue。

## 契約

源片：ASR timing。自製口播：speech artifact duration。Caption layout 另 loop。TTS duration 過慢／pause 異常 = speech generation，不是 layout。

與源片 ASR word timing 同一思想：來源不同，都是 evidence-backed timing。
