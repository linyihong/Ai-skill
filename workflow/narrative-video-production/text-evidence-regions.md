# Text evidence — OCR regions, subtitle groups, bilingual pairs

何時讀：硬字幕可能多語疊字、多 box、或 OCR 把多語黏成一字串時；在 Language Relation Gate／Text Resolution 之前。
欄位：[`records/text-evidence.yaml`](records/text-evidence.yaml)。
Plan：[`41-bilingual-ocr-regions-and-subtitle-groups.md`](../../plans/active/2026-09-16-1649-narrative-video-production-workflow/41-bilingual-ocr-regions-and-subtitle-groups.md)。
銜接：[`text-evidence-language-relation.md`](text-evidence-language-relation.md)。

> **執行契約，不重新解釋契約。** LLM 不猜語言行別；機械 region／group 先於 alignment。

## 在 lifecycle 的位置

```text
OCR Detection
  → Raw OCR Evidence
  → Text Segmentation / Boundary Recovery  (見 text-evidence-ocr-boundary.md)
  → text_region (text, box, language_candidate, raw/derived, …)
  → Region Classification (dialogue vs non-dialogue candidate)
  → Subtitle Grouping (incl. bilingual_pair)
  → Language Relation Gate (region ↔ ASR)
  → Text Resolution
```

對應 execution-flow **6b** 的前半；6b 後半仍是 Language Relation Gate。

## `text_region`（必填／可選）

| 必填 | 可選 |
| --- | --- |
| `id`、`text`、`language_candidate` | `box`／`normalized_box`、`box_size`、`timestamp`／frame、`persistence`、`visual_style`、`text_role` |

`language_candidate`：Latin-heavy→en；CJK-heavy→zh／ja／ko；Thai→th；Arabic→ar；未知→unknown。只作 candidate。

## Subtitle Grouping

同時間窗＋空間合理（常見：stacked、相似寬度、bottom band）→ `subtitle_group`。
`type`：`mono`｜`bilingual`｜`multi`｜`unresolved`。
bilingual 時記 `language_relation`（如 en+zh）與 `spatial_relation: stacked`。

## 與 ASR／spoken

| 情況 | spoken／dialogue primary | 異語 region |
| --- | --- | --- |
| bilingual ZH+EN，ASR=zh | zh dialogue region + ASR | EN = translation evidence（subtitle） |
| 僅 EN OCR，ASR=zh | ASR spoken；EN = cross_language subtitle | 見 language-relation |
| OCR 含 SALE／UI | 排除出 dialogue group | 保留 raw；role projection |

禁止：翻譯字幕直接當 `spoken_text`；未分 region 就把雙語黏字串當單一 OCR 證據。

## 推進條件

| 條件 | 失敗 |
| --- | --- |
| 多語疊字窗有 ≥2 `text_region` 或明示 split 失敗原因 | 不得用單一字串進 Language Relation 當唯一 OCR |
| bilingual group 有 per-region language_candidate | 不得 LLM 猜行別 |
| spoken 來自口播語 dialogue evidence | EN-only hardsub 當 zh spoken = 閘失敗 |
