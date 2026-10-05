# Candidate: Multimodal Text Evidence Resolution

Companion to [`19-text-resolution-and-narrative-assembly.md`](19-text-resolution-and-narrative-assembly.md)、
[`21-spoken-vs-subtitle-reconstruction.md`](21-spoken-vs-subtitle-reconstruction.md)、
[`22-phonetic-text-reconstruction.md`](22-phonetic-text-reconstruction.md)、
[`40-language-role-before-text-resolution.md`](40-language-role-before-text-resolution.md)。
**不是新 Phase、不是新 Agent、不改凍結的 Narrative Video Workflow 架構。**
觀察：[`evidence/2026-09-30-asr-anomaly-ocr-semantic-reconstruction.md`](evidence/2026-09-30-asr-anomaly-ocr-semantic-reconstruction.md)、
[`evidence/2026-09-30-semantic-anchor-last-lot-auction.md`](evidence/2026-09-30-semantic-anchor-last-lot-auction.md)、
[`evidence/2026-10-05-evidence-resolution-loss-monitor.md`](evidence/2026-10-05-evidence-resolution-loss-monitor.md)。
Workflow 落點：[`text-evidence-multimodal-resolution.md`](../../../workflow/narrative-video-production/text-evidence-multimodal-resolution.md)。

## 問題

ASR 字面結果不是 spoken meaning 的 SoT。OCR／ASR／語言模型／場景應共同做
**Evidence Resolution**，不是固定「OCR > ASR」或「ASR 字面即台詞」。

典型：ASR「不偿不偿我」（phonetic）＋ OCR `Compensate compensate me`（lexical／semantic）
→ 合理 spoken reconstruction「補償補償我」。與「秦舍→禽獸／小可愛」同屬更高層問題。

## 升格後的管線（插在既有 Language Relation Gate 之後）

```text
Video → OCR | ASR | Timing | Speaker | Scene/Event
         ↓
Evidence Normalization（language / text_role / region / validity）
         ↓
Phonetic Evidence          Semantic Evidence
（ASR observed + pinyin／syllable）  （OCR observed + translation candidate）
         ↓
Candidate Generation（不得用 OCR 顯示詞做同音展開；semantic 保留 provenance）
         ↓
Context Resolution（local / narrative window / scene）
         ↓
Independent Verification + resolution_reason
         ↓
Resolved Text（spoken_text / subtitle_text 分家）
```

這把既有「OCR↔ASR Text Resolution」升成 **Multimodal Text Evidence Resolution**。
不新增 publish stage；仍是 acquisition／assemble 的 evidence 能力。

## 證據分工（三角）

| 極 | 來源 | 提供 |
| --- | --- | --- |
| Acoustic | ASR observed＋phonetic | 該時窗有人說了接近某串音的話 |
| Lexical | OCR observed（字幕／硬字幕） | 畫面上實際寫的詞 |
| Semantic | OCR→translation candidate、scene／speaker | 這句在語境裡應表達什麼 |

三者交集才升格 resolved；缺支撐 → `unresolved`，不得發明台詞。

## ASR / OCR 必須拆欄位

ASR 至少：`observed_text`、`phonetic`（pinyin／syllable）、`timing`、`speaker`（若有）、
`quality.lexical_confidence`。`observed_text` 是 observation，不是真理。

OCR 至少：`text`、`language`、`text_role`。跨語言字幕再出 **semantic_candidate**
（translation／normalization），**不得**覆寫 OCR observed；只作 candidate＋provenance。

## semantic_candidate（本輪新增抽象）

```yaml
text_resolution:
  observed:
    asr: "<asr_observed>"
    ocr: "<ocr_observed>"
  candidates:
    - text: "<spoken_reconstruction>"
      evidence: [ocr_semantic, asr_phonetic, context]
      status: candidate
  resolution:
    text: "<selected>"
    reason: [ocr_exact_semantic_match, asr_phonetic_support, contextual_fit]
    sources: [...]
    confidence: { type: evidence_supported }  # 或 unresolved
```

LLM 只在 candidate set 上選擇並寫 `resolution_reason`，不得從 raw ASR 直接生成台詞。

## 語意異常（觸發重建，不直接改字）

機械可先標：`low_lexical_naturalness`、`unusual_phrase`（例「不偿不偿我」）。
只產生 **ASR suspicious → Candidate Reconstruction**，禁止機械替代表寫死答案。

## evidence_policy（動態權重，非固定 OCR>ASR）

| 情境 | 傾向 |
| --- | --- |
| subtitle_dialogue | lexical high；ASR supporting |
| no_subtitle | ASR primary |
| bilingual_subtitle／cross_language | subtitle_semantics high；ASR supporting |
| suspicious_asr | ocr_semantics high；phonetic_reconstruction enabled |

統一案例族：錯音重建、字幕正確／ASR 怪、跨語言對齊、同語言 match。詳見 workflow 落點。

## 禁止

- 把 OCR 翻譯結果直接當 `spoken_text` 且丟棄 ASR／provenance
- 固定全域 OCR>ASR 或 ASR 字面即 canonical
- 無 phonetic／semantic／context 支撐時用 LLM 自由改寫
- 改變凍結 Phase／新增 Story Agent

## semantic_reconstruction（同能力，非 typo correction）

正式區分 **observed / candidate / resolved**。OCR 領域術語可作 `semantic_anchor`
（例 `Last Lot`→auction→「最後一件拍賣品」），與 ASR lexical anomaly（「拍皮」）
收斂後寫入 candidate；resolved 才進敘事。OCR 非必要，但提高 evidence strength。

## 下一輪只驗

1. 產品是否產出 `semantic_candidate`＋`resolution_reason`（至少 dogfood 一例）；
2. ASR lexical anomaly 是否觸發 reconstruction 而非靜默採用字面；
3. 跨語言路徑是否仍先過 Language Relation Gate（40／41）；
4. `Last Lot`＋「拍皮」類是否走 `semantic_anchor`／`semantic_reconstruction` 且保留 ASR observed。
5. resolution funnel 大幅縮量是否保留 uncertain／rejected／merged trace，並由定向 re-probe 處理；derived cache 是否先通過 integrity validation。

Adapter progress：定向重探與 before／after、失敗保留、box／parts 失配及 nested
decision eligibility 的 fixture 已驗證；live accepted 品質仍須獨立核對，不以 cue
數增加當 recovery PASS。見上述 resolution-loss evidence。

Cross-source regression gate 仍 open：新舊 OCR ownership 格式、watermark
precision 與 true-dialogue retention 必須分開驗證；主字幕帶頻率不能單獨
否定低頻的另一字幕帶。新採集實跑與舊快取 replay 分開記錄，pipeline
完成不等於 content／timing／render gate 通過。
