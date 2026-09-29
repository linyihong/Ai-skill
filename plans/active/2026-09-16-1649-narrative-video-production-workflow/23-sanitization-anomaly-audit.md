# Candidate: Sanitization anomaly detection and final text audit

Companion to [`22-phonetic-text-reconstruction.md`](22-phonetic-text-reconstruction.md)。**Text Reconstruction 內的警覺／複核能力，不是新 Agent，也不是 workflow schema。**  
觀察：[`evidence/2026-09-21-sanitization-anomaly-audit.md`](evidence/2026-09-21-sanitization-anomaly-audit.md)。

顯示詞不是錯字，也**不得**拿去做同音判斷。OCR 是 `subtitle` observed；和諧嫌疑是獨立 `sanitization` alert，只是待驗證假設。

**Language Relation Gate 是前置條件**（[`40-language-role-before-text-resolution.md`](40-language-role-before-text-resolution.md)）：僅 **same-language** semantic conflict 才可標 `sanitization_candidate`。跨語言（例 EN subtitle ↔ ZH spoken）且為 `cross_language_translation` → **禁止**進和諧偵測。

Validity 先於解釋：見 [`39-asr-validity-precedes-interpretation.md`](39-asr-validity-precedes-interpretation.md)。Phonetic candidate 只從**有效** ASR 鏈長出。兩條 chain 最後才由 LLM 串：疑似和諧 → 去找 spoken evidence → reconstruction。OCR＝ASR 時不得因歷史 pattern 強改。Alert 機制，不是替換字典。前警覺＋Final Text Audit。Workflow 閘：[`text-evidence-language-relation.md`](../../../workflow/narrative-video-production/text-evidence-language-relation.md)。
