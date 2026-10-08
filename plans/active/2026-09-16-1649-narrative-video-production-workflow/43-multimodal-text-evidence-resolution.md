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
另需 watermark lexicon extension 的 region-boundary regression：重複
首字不等於水印末字，擴詞不可跨入另一 dialogue box。

分階段 attribution 修復的機械重播已累積，但 cross-source gate 仍 open：
英文短字幕不得僅以 Latin 長度判 junk；雙語 script split 不可在缺 box
時聲稱已確認 stacked geometry。這些缺口須保留 failing regression，
並補獨立 frame／region／final-cue 驗證。既有污染 derived lexicon 也需
重驗與隔離；完整 resolver 的 uncertain 保留不等於 accepted 恢復。

後續 precision／acquisition-integrity 核對仍有缺口：附近 scene text
不能借用 changing dialogue band 升格；decoder 提前中止不能因後續
resolver 完成就當全片採集完成；cached frame 與 source presentation time
不一致時，先核對來源與時軸。見
[`evidence/2026-10-06-attribution-precision-and-acquisition-integrity.md`](evidence/2026-10-06-attribution-precision-and-acquisition-integrity.md)。

獨立雙語 region 已採集但 final cue 丟失語言與 owned geometry 的實跑反例，
已由 test-first adapter 修正與指定 frozen-input 重播驗證局部恢復。
下游 gate 必須核對各 region 的來源 identity＋自身文字，不能用合併雙語
字串去比單一 box；配對仍須 acquisition／frame／source 與時間交集。
raw disposition 全數有去向與 cross-language resolved label 均不能代替
identity／semantic 驗收。完整 cross-source／locale／timing／render gate 仍 open。

後投影 fuzzy watermark scrub 吃掉完整 box-local dialogue 的反例已由
test-first adapter 修正；shared affix 不具 ownership authority，legacy
region discovery 也不得覆蓋已分類的 non-dialogue role。凍結 diagnostic
確認完整候選保留，但其自身 region 尚未進 broad accepted speech cue；
uncertain retention、cleaning recovery 與 identity propagation 分開驗收。
此 diagnostic 僅停用部分 reconstruction，不標純機械或 normal-provider
baseline；詳見同份 attribution precision evidence。全鏈 gate 不升格。

Same-speech collapse 的 winner-only projection 已定位並以 adapter regression
修正：同一 speech 可對齊多個獨立 temporal subtitle group，不能因 selected
speech 字串相同就丟掉第二個 region。保留 verbatim geometry／identity／time
與複數 group alignment；singular group 僅相容主視圖，不將 sequential groups
當 bilingual stack。凍結 resolver 驗證局部恢復，排程／成片 gate 仍 open。

Target-locale consumer 的 synthetic write／reload regression 已重現複數
temporal groups 被 allowlist 丟掉、而 singular view／regions／alignment
仍存在的 adapter gap。Spoken-side projection PASS 不可替代 locale round-trip
驗證；修正後仍須獨立驗 translation projection、caption scheduling 與成片。
見同份 attribution precision evidence；consumer／全鏈 gate 保持 open。

後續 cache round-trip 的 plural-group retention regression 已局部通過，
但實際 display selector 又暴露 first-region partial override：metadata 全保留
仍可漏掉另一句。保守 adapter 修正拒絕歧義的單 region override，沿用
whole-cue translation，並以真實 locale／cache／caption consumer 的合成
regression 驗 content coverage 與 order invariance；不拼 raw region、不改時間。
見 [`evidence/2026-10-07-locale-selection-vs-retention.md`](evidence/2026-10-07-locale-selection-vs-retention.md)。
Scoped adapter evidence 不升格 full-source／semantic／timing／render gate。

中文 target orthography 必須在 derived locale／caption boundaries 明確實現，
不能把 script-family accepted 當簡體驗收，亦不能改寫 raw OCR／ASR。
見 Translation Decision [boundary evidence](../2026-09-22-1000-translation-decision-workflow/evidence/2026-10-07-target-orthography-boundaries.md)。
cache／identity／caption regression 僅局部驗 adapter；content／render gate 仍 open。

Local-only realization 的 native-loader 問題需獨立程序保留堆疊與退出狀態，
以單變量驗證讀取相容性；完整載入不能替代 generation／caption／render 驗收。
細節與驗收邊界見同份 Translation Decision boundary evidence。

OCR-only translation diagnostic 與 resolved spoken-pivot identity delivery
需分開驗 source authority；不可把前者的 acoustic detail 缺失當後者的
成片結論，也不可將舊版 frozen cache 改標後通過 integrity gate。
Local-only policy 必須涵蓋 repair／fallback，不只清空測試 key。
一般 semantic prompt guard 無效時不得算 repair；同份 boundary evidence
已記錄此 adapter 驗證方法，full semantic／timing／render gate 仍 open。

Source-analysis candidate 的 schema／span anchoring 不等於 meaning／reference
正確；invalid model generation 也必須保留。見 Translation Decision
[analysis boundary](../2026-09-22-1000-translation-decision-workflow/evidence/2026-10-07-source-analysis-and-provider-selection.md)。不新增 actor 或升格全鏈 gate。

Seeded cross-source background verification 需先固定 source／locale manifest、
baseline 與 code provenance，隔離 evidence／translation cache，保留每階段
native exit／timeout／NOT_RUN。Scheduling fixture 與 transport-independent
dispatch 僅驗 harness，不能當 corpus content／timing／layout PASS；bounded
caption burn 不是 full publish。方法見同份 Translation Decision boundary
evidence；實跑仍須逐階段驗收，cross-source gate 仍 open。

背景產物檢查暴露 audio fallback 被誤標為視覺觀測，以及 target projection
無聲 omission。見 [audio attribution／locale omission](evidence/2026-10-08-audio-attribution-and-locale-omissions.md)：
producer modality 先於 visual gate；每條 selected cue 都要 downstream disposition。
凍結 replay／paired fixtures 僅支持 adapter 修復，不關閉 corpus 三閘。
