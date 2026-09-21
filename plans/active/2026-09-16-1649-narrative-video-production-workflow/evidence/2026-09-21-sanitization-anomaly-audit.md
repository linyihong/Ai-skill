# Observation — sanitization anomaly and final text audit

**Run ID**：2026-09-21-sanitization-anomaly-audit  
**Kind**：Phase 3 observation（**不是** `contract_gap`；**不是**新 Agent；**不**擴 workflow）  
**Extends**：[`2026-09-21-phonetic-text-reconstruction.md`](2026-09-21-phonetic-text-reconstruction.md)、[`2026-09-18-episode-vs-knowledge-accumulation.md`](2026-09-18-episode-vs-knowledge-accumulation.md)

## 不學硬對照

顯示詞與 spoken 猜測對上一次，仍不得寫成自動替換。要學的是：正面／中性 OCR 詞在衝突語境下可能是 sanitization，因此產生 `text_alert`（`possible_sanitized_subtitle` 等），status 只 candidate。Raw OCR／ASR 不動。

## 前警覺 + 後複核

Alignment 後做 Suspicion Detection：和諧詞候選、異常語境、音近異常、歷史替換模式。訊號疊加（ocr_asr_mismatch、phonetic_anomaly、contextual_mismatch）才提高 risk 並進 reconstruction。

整集 episode analysis 後做 Final Text Audit：同一顯示詞反覆出現在衝突句式時，第一階段「看起來正常」的 span 重新進 reconstruction。未解決 anomaly 不得當 Story Truth。

## 三級警覺

| 級 | 含義 | 例 |
| --- | --- | --- |
| A 詞彙 | 常見可替換詞本身無過 | risk low |
| B 語境 | vocative／action／emotion 衝突 | risk medium |
| C 多證據 | OCR＋ASR phonetic＋語境同時衝突 | risk high → reconstruct |

統稱 **Sanitization / Substitution Anomaly**（辱罵／敏感／審核弱化／同音諧音／OCR／ASR 誤識）。不是暴力詞表。Learning candidate 多集後才 `verified` mapping。
