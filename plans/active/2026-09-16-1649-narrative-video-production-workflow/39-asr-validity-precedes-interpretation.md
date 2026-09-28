# Candidate: ASR validity precedes sanitization interpretation

Companion to [`23-sanitization-anomaly-audit.md`](23-sanitization-anomaly-audit.md)。**選擇順序契約。本 repo 沒有 Fusion 程式。**  
觀察：[`evidence/2026-09-28-asr-validity-precedes-interpretation.md`](evidence/2026-09-28-asr-validity-precedes-interpretation.md)。

Evidence validity precedes evidence interpretation。

ASR 先過 validity，才進 feasible set。`possible_subtitle_sanitization` 是待驗證假設，不得提高無效 ASR 的優先級，也不得蓋過「OCR valid + ASR invalid → 選 OCR」。

`asr_validity` 分開記：`temporal_alignment`、`speech_rate`、`repetition`、`lexical_coherence`，再匯總 `validity`。語速異常只是一項 evidence，不是 hallucination 的充分條件。`alignment.status: invalid` 是時間錯位；時間有效但用字可疑才走 phonetic／context resolution。連續重複短句（例：同一短語 ≥ 3）是 `repetitive_hallucination`，該段 ASR 離開可行集。
