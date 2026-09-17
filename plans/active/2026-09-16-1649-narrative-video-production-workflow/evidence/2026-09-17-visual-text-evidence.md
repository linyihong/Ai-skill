# Observation — visual text evidence

**Run ID**：2026-09-17-visual-text-evidence  
**Kind**：Phase 3 observation（**不是** `contract_gap`；**不**擴 workflow）  
**current_support**：partial（plan／locale 有 `text_origin: ocr`；源片分析未要求 timestamped visual text）  
**candidate_concept**：`visual_text_evidence`

## 與現況

- 已有：ASR vs OCR vs 腳本，禁止混成一條無標記對白（locale／caption）。
- 沒有：Video → OCR → timestamped text → clip catalog／narrative／identity 的明確一級產物。
- 不把 OCR 改成「必須由 LLM 做」。它屬 Material Fact Extraction 的 evidence source。

## 兩個證據通道

| 通道 | 回答 | 例欄 |
| --- | --- | --- |
| `speech_text` | 有人說了什麼 | ASR + start/end + `speaker_id` |
| `visual_text` | 畫面上出現什麼字 | OCR + start/end + region；不一定是字幕 |

Visual text 可能是硬字幕、對白燒錄、手機訊息、招牌、地名、文件、UI、名片、時間／集數標記。

可與 ASR 互證，也可單獨提供 ASR 沒有的資訊（例如畫面「三年後」）。可成為 [`08-identity-precedes-naming.md`](../08-identity-precedes-naming.md) 的 `name_evidence`，不讓 LLM 第一集猜名。

## 真實片子要數

名稱、地點、時間、對白、手機訊息、劇情提示、場景文字：哪些穩定出現、哪些被 bible／catalog／matching／EDR **消費**。穩定且被消費才考慮升 contract。
