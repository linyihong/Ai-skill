# Observation — material fact extraction (L0)

**Run ID**：2026-09-17-material-fact-extraction  
**Kind**：Phase 3 observation（**不是** `contract_gap`；**不**擴 workflow）  
**原則**：Analysis should maximize deterministic evidence before semantic inference.

不是「讓 AI 分析影片」。是先用 deterministic／classical CV／音訊處理，把片子變成可查詢素材庫；AI 只在後面理解、歸納、決策。對齊現有 bible → catalog → matching → EDR。

## Stage A — Material Fact Extraction（LLM 不做主判斷）

外部分析可提供的 **evidence candidates**（本 phase 全部 optional，不進 contract）：

| 候選 | 機械產物（例） | 尚未等於 |
| --- | --- | --- |
| 時間結構 | duration／FPS／scene／shot 切點／黑幀／freeze／近重複／keyframes | 敘事 beat |
| ASR | `text` + start/end；可選 word 時間軸 | `character_id` |
| Diarization | `speaker_id` + 聲線 profile | `character_id` |
| Face | detect → track → cluster → `person_cluster`；出現窗 | 角色名 |
| OCR | text + bbox + timestamp | 劇情解釋 |
| Visual L1 | keyframe、亮度／模糊／histogram | 「悲傷」標籤 |
| Visual L2 | embedding 當 retrieval 表示 | LLM 場景判決 |
| Audio | loudness／silence／speech vs music | 音樂情緒 |
| SFX | 專門 classifier（門／鈴聲）可選 | LLM 音效理解 |

匯聚進 catalog 時可以只有 `person_cluster`／`speaker_id`／ASR／OCR，**還沒有** `character_chen`。

## Stage B — Narrative Understanding（才允許語義）

誰是誰、關係、事件、`semantic_context`、template、selection。Entity resolution：`speaker_02` → 可能是某 `character_id`。

## 真實片子要數

哪些 L0 欄位實際被 source-bible／clip-catalog／matching／EDR **消費**。沒被消費的不進 schema。禁止為完整而把十種分析一次寫進 workflow。
