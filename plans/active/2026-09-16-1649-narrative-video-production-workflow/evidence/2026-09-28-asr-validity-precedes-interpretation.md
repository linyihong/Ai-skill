# Observation — ASR validity precedes interpretation

**Run ID**：2026-09-28-asr-validity-precedes-interpretation  
**Kind**：contract（選擇順序；本 repo 無 Fusion 程式）  
**Extends**：[`2026-09-21-sanitization-anomaly-audit.md`](2026-09-21-sanitization-anomaly-audit.md)

## 反例

OCR「糟了」有效。ASR 在錯的時間重複「小雅你先回房休息」。Heuristic 把「OCR 短 + ASR 長 + 不重合」標成 `possible_subtitle_sanitization`，並讓這個 relation 擋住「無效 ASR 時選 OCR」。

## Contract

先判 ASR validity，再解釋和諧。和諧是 hypothesis。無效 ASR 不進可行集。時間錯位與用字可疑分開。兩者都 valid 時才進 phonetic／context resolution。
