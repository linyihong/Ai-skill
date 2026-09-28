# Observation — speech unit before caption

**Run ID**：2026-09-28-speech-unit-before-caption  
**Kind**：contract  
**Extends**：[`2026-09-22-speech-timing-authority.md`](2026-09-22-speech-timing-authority.md)、[`2026-09-28-natural-boundary-before-length.md`](2026-09-28-natural-boundary-before-length.md)

## 反例

同一段文案同時交給字幕切割與 TTS。兩邊各自決定停頓、換 cue、換行、縮字與秒數，最後不同步，或把「谈话」「电脑里」切壞。

## Contract

Speech Unit 先決定怎麼說。標點是候選，不是強制切。TTS 回實際 duration。Cue 時間跟 unit。換行只在 unit 內，且不得改 unit 原文。
