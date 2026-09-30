# Candidate: Bilingual OCR regions & subtitle groups

Companion to [`40-language-role-before-text-resolution.md`](40-language-role-before-text-resolution.md)、
[`24-ocr-role-projection.md`](24-ocr-role-projection.md)、
[`25-text-group-preserve-variants.md`](25-text-group-preserve-variants.md)、
[`21-spoken-vs-subtitle-reconstruction.md`](21-spoken-vs-subtitle-reconstruction.md)。
**不是新 Phase、不是新 Agent、不推翻 Phase 1/2。**
觀察：[`evidence/2026-09-29-bilingual-hardsub-regions.md`](evidence/2026-09-29-bilingual-hardsub-regions.md)。
Workflow：[`text-evidence-regions.md`](../../../workflow/narrative-video-production/text-evidence-regions.md)；
欄位：[`records/text-evidence.yaml`](../../../workflow/narrative-video-production/records/text-evidence.yaml)。

## 問題

同一時間窗內可同時出現 **中英（或多語）硬字幕**（上下疊字）。若 OCR 合成單一字串
`"Don't touch me. 你不要碰我。"`，後面 Language Relation／Text Resolution 只能猜，
並容易把翻譯列當 spoken、或把雙語當 conflict。

這是 **OCR evidence 粒度** 缺口：需要 `text_region`＋機械 grouping，不是雙語 Agent。

## 吸收策略（機械優先）

優先既有：`visual_text_evidence`／`normalized_box`／`text_span_role`／`text_group`／
`spoken_text`／`subtitle_text`／`text_relation`（見 40）。升級 OCR 為多 region：

| 物件 | 用途 |
| --- | --- |
| `text_region` | 單一 OCR box：text、language_candidate、box／normalized_box、timestamp、persistence、visual_style |
| `subtitle_group` | 同窗、空間可疊的 dialogue regions（含 bilingual pair） |
| `language_relation` | same_language／bilingual_pair／translation_pair／unrelated_text／unresolved |

禁止：LLM 單獨猜「哪一行是哪種語言」；機械 script-family／detector 只產 **candidate**。

## 前置管線（插在 Language Relation Gate 之前）

```text
Frame
  → OCR Detection → text_region (+ box, language_candidate)
  → Region Classification (dialogue vs non-dialogue candidate)
  → Subtitle Grouping (spatial/temporal → bilingual_pair?)
  → Language Relation Gate (per region ↔ ASR)
  → Text Resolution
```

## Bilingual Subtitle Pair（機械）

疊字候選信號（同時滿足多數即可）：

- same timestamp／高 temporal_overlap
- similar horizontal span；不同 vertical position（stacked）
- different `language_candidate`（如 en + zh）
- 合理 bottom-band／subtitle-like geometry

產出：

```text
subtitle_group:
  type: bilingual
  regions: [ocr_en, ocr_zh]
  alignment: { temporal_overlap, spatial_relation: stacked, language_relation: en+zh }
```

## OCR ↔ ASR（多 region）

例：OCR EN + OCR ZH + ASR ZH →

| relation | pair |
| --- | --- |
| `cross_language_translation` | ocr_en ↔ asr_zh |
| `same_language_match`（或 variant） | ocr_zh ↔ asr_zh |

劇情／spoken primary：**與 ASR 同語言的 dialogue subtitle region**（＋ASR）。
異語言 dialogue region = **translation evidence**（可進 subtitle.observed／locale），
**不得**當 spoken_text。

「有字幕 → OCR 優先」修正為：有 **已確認為 dialogue subtitle 的 OCR region** →
該語言字幕作 dialogue evidence；非 dialogue（SALE／招牌／UI）不得進劇情文字。

## Phase 3 分類

| 標籤 | 判定 |
| --- | --- |
| `bilingual_collapsed_to_single_string` | **contract_gap candidate**；用 region／group 吸收；不開新 Phase／Agent |

產品 adapter 另改；不得反向定義本契約。翻譯字幕只作 semantic evidence，不可直接當 spoken。

## Projection 契約（2026-09-30）

Layer-0 OCR schema **不必重做**。缺口在 **OCR → dialogue cue → locale pack** 的投影：

- `spoken_timed` = spoken／口播語 burn／phrase 投影，**不**承載 subtitle 原始結構。
- Dialogue Cue Group 必須保留：`speech` + `subtitle_group.regions[]` + `alignment{speech_id, subtitle_group_id}`。
- Export／reload 不得把 bilingual cue 壓成 `{src,dst,start,end}` 後丟掉 regions。
- 目標語 burn：優先使用已存在的同語 OCR region；缺口才 MT。

觀察：[`evidence/2026-09-30-dialogue-cue-projection-retains-subtitle-group.md`](evidence/2026-09-30-dialogue-cue-projection-retains-subtitle-group.md)。
