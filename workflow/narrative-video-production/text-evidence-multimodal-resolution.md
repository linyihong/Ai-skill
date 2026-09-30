# Text evidence — Multimodal Text Evidence Resolution

何時讀：OCR＋ASR（或硬字幕＋口播）已過 Language Relation Gate，要做 spoken／subtitle
reconstruction，且可能出現 ASR 字面怪、跨語言字幕語義清晰、或近音錯字。

Plan companion：[`43-multimodal-text-evidence-resolution.md`](../../plans/active/2026-09-16-1649-narrative-video-production-workflow/43-multimodal-text-evidence-resolution.md)。
前置：[`text-evidence-language-relation.md`](text-evidence-language-relation.md)、
[`text-evidence-regions.md`](text-evidence-regions.md)。

## 在管線中的位置

插在 Language Relation Gate **之後**、narrative consumer **之前**（不是新 publish stage）：

```text
… → Language Relation Gate
  → Multimodal Text Evidence Resolution
       Evidence Normalization
       → Phonetic Evidence | Semantic Evidence
       → Candidate Generation
       → Context Resolution
       → Independent Verification
  → Resolved spoken_text / subtitle_text
```

## 核心契約

1. **ASR observed ≠ spoken meaning SoT**；保存 phonetic／timing／quality。
2. **OCR observed ≠ 自動 spoken**；跨語言時另產 `semantic_candidate`（provenance）。
3. **LLM 只選 candidate**，必須寫 `resolution.reason`＋`sources`。
4. **異常只觸發重建**，不機械替代表定案。
5. **權重看 evidence_policy**，禁止全域 OCR>ASR。

## 最小欄位

| 物件 | 必填／建議 |
| --- | --- |
| `asr` | `observed_text`, `language`; 建議 `phonetic`, `timing`, `quality.lexical_confidence` |
| `ocr` | `text`, `language`, `text_role`; 建議 `region_refs` |
| `semantic_candidate` | `text`, `from: ocr_translation\|normalize`, `source_ref` |
| `candidates[]` | `text`, `evidence[]`, `status` |
| `resolution` | `status`, `text`（若 resolved）, `reason[]`, `sources[]`, `confidence.type` |

`confidence.type`：`evidence_supported`｜`needs_review`｜`unresolved`（與 Finality 語意對齊即可，不強制同一 enum）。

## evidence_policy（摘要）

| 情境 | lexical／OCR semantic | ASR |
| --- | --- | --- |
| 有 dialogue subtitle | high | supporting |
| 無字幕 | — | primary |
| 跨語言／雙語字幕 | subtitle_semantics high | supporting |
| ASR lexical anomaly | ocr_semantics high；開 phonetic reconstruction | suspicious |

## 案例族（同一框架）

| 模式 | 典型訊號 | 走向 |
| --- | --- | --- |
| ASR 錯音 | phonetic 近、字面怪 | phonetic candidates＋context |
| ASR 怪／字幕對 | OCR semantic 清晰 | semantic_candidate＋phonetic support |
| 跨語言對齊 | EN OCR＋ZH ASR | `cross_language_translation`；非 conflict |
| 同語言 match | 字面一致 | `same_language_match` |

## 禁止

- 翻譯字幕直接寫入 `spoken_text` 且無 ASR／reason
- 未過 Language Relation Gate 就跑 sanitization／和諧覆蓋
- 用 OCR 顯示詞做同音展開（仍遵守 22）
- 無 evidence 時 LLM 自由改寫

## 產品落點

Adapter 應在 fuse／spoken resolution 產出上述 candidates 與 reason；locale content
消費 **resolved pair**，不消費 raw ASR 字面當唯一 SoT。Record：
[`records/text-evidence.yaml`](records/text-evidence.yaml)。
