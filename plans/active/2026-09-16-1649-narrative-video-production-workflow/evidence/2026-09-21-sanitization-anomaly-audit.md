# Observation — sanitization anomaly and final text audit

**Run ID**：2026-09-21-sanitization-anomaly-audit  
**Kind**：Phase 3 observation（**不是** `contract_gap`；**不是**新 Agent；**不**擴 workflow）  
**Extends**：[`2026-09-21-phonetic-text-reconstruction.md`](2026-09-21-phonetic-text-reconstruction.md)、[`2026-09-18-episode-vs-knowledge-accumulation.md`](2026-09-18-episode-vs-knowledge-accumulation.md)

## 兩條獨立 evidence chain

OCR 顯示詞保持 `role: subtitle`、`status: observed`。若與 spoken 線索衝突，只標 `sanitization.suspicion`（`possible_substitution`、`conflict_with_spoken_evidence`）。**不得**把該顯示詞送進 homophone decoder。

ASR grapheme → phonetic → spoken candidate 是另一條鏈。LLM 最後才解釋：字幕可能替換；ASR 近音＋語境可能指向 spoken candidate。路徑是「疑似和諧 → 找 spoken evidence」，不是「顯示詞 → spoken 對照」。

OCR 與 ASR 同詞時，歷史 sanitization pattern 不得強改。OCR 與另一無關 ASR 詞時亦然。

## 警覺來源（不是字典 lookup）

OCR-only suspicious、OCR／ASR 語意衝突、ASR phonetic anomaly、context anomaly、historical pattern（只加權，不結案）。機械發現異常 → phonetic 產候選 → LLM 語境選擇 → verifier。Learning candidate 可累積「顯示詞可能是和諧」，不得升永久規則。

## 前警覺 + 後複核

Alignment 後 suspicion；整集後 Final Text Audit。三級 risk：詞彙 low、語境 medium、多證據 high。統稱 Sanitization / Substitution Anomaly。
