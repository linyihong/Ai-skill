# Observation — material fact extraction (observable layer)

**Run ID**：2026-09-17-material-fact-extraction  
**Kind**：Phase 3 observation（**不是** `contract_gap`；**不**擴 workflow）  
**原則**：Analysis should maximize deterministic evidence before semantic inference.

不是「讓 AI 分析影片」。是先用 deterministic／classical CV／音訊處理，把片子變成可查詢素材庫；AI 只在後面理解、歸納、決策。對齊現有 bible → catalog → matching → EDR。

## Stage A — Material Fact Extraction（LLM 不做主判斷）

外部分析可提供的 **evidence candidates**（本 phase 全部 optional，不進 contract）：

| 候選 | 機械產物（例） | 尚未等於 |
| --- | --- | --- |
| 時間結構 | duration／FPS／scene／shot 切點／黑幀／freeze／近重複／keyframes | 敘事 beat |
| ASR | `transcript.text` + timing；**不**內嵌角色名 | `character_id` |
| Voice／Speaker | `speaker_id`／`voice_track` 掛在 segment 上；embedding 可空 | `character_id`（見 [`2026-09-18-voice-speaker-evidence.md`](2026-09-18-voice-speaker-evidence.md)） |
| Face | detect → track → `face_track_id` + keyframes；cluster／embedding 可空 | 角色名／`character_id` |
| OCR | 產出 `visual_text_evidence`：text + pixel／normalized box + timestamp + persistence 可導出 | 字幕／浮水印已判定；`character_id` |
| Visual mechanical | keyframe、亮度／模糊／histogram | 「悲傷」標籤 |
| Visual embedding | retrieval 表示 | LLM 場景判決 |
| Audio | loudness／silence／speech vs music | 音樂情緒 |
| SFX | 專門 classifier（門／鈴聲）可選 | LLM 音效理解 |

匯聚進 catalog 時可以只有 `face_track`／`speaker_id`／ASR／OCR，**還沒有** `character_chen`。Face cluster 可空。連結層見 [`2026-09-17-face-as-candidate-evidence.md`](2026-09-17-face-as-candidate-evidence.md)。

## Stage B — Narrative Understanding（才允許語義）

誰是誰、關係、事件、`semantic_context`、template、selection。Entity resolution：`speaker_02` → 可能是某 `character_id`。

## 真實片子要數

哪些 observable 欄位實際被 source-bible／clip-catalog／matching／EDR **消費**。沒被消費的不進 schema。禁止為完整而把十種分析一次寫進 workflow。
