# Candidate: Sanitization anomaly detection and final text audit

Companion to [`22-phonetic-text-reconstruction.md`](22-phonetic-text-reconstruction.md)。**Text Reconstruction 內的警覺／複核能力，不是新 Agent，也不是 workflow schema。**  
觀察：[`evidence/2026-09-21-sanitization-anomaly-audit.md`](evidence/2026-09-21-sanitization-anomaly-audit.md)。

禁止「看到顯示詞就自動改成 spoken 猜測」。學的是：看似正常甚至語意衝突的 OCR 詞，**可能**是和諧／替換，因此要提高警覺。`text_alert` 只 candidate。

兩次掃描都要：讀劇時 suspicion（alignment 後：和諧詞候選、異常語境、音近異常、歷史模式）才進 reconstruction；整集敘事建成後 **Final Text Audit** 再掃未解決 anomaly。詞彙／語境／多證據三級 risk；涵蓋面是 **Sanitization / Substitution Anomaly**（審核弱化、同音、諧音、OCR／ASR 誤識），不是暴力詞對照表。產出進 Learning Inbox，多集後才 verified。Phase 3 dogfood 驗 scope，**不**改 workflow。
