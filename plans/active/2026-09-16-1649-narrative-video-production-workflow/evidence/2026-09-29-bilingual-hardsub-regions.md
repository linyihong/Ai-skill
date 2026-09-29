# Observation — bilingual hardsub as stacked OCR regions

**Run ID**：2026-09-29-bilingual-hardsub-regions
**Kind**：Phase 3 dogfood observation → **`contract_gap` candidate**（`bilingual_collapsed_to_single_string`）
**Plan**：[`../41-bilingual-ocr-regions-and-subtitle-groups.md`](../41-bilingual-ocr-regions-and-subtitle-groups.md)
**Extends**：[`2026-09-29-cross-language-subtitle-spoken-alignment.md`](2026-09-29-cross-language-subtitle-spoken-alignment.md)、[`2026-09-17-visual-text-evidence.md`](2026-09-17-visual-text-evidence.md)

## 去敏情境

參考成片硬燒 **中英雙語字幕**（上下疊字）；音軌為中文口播。
片名／路徑／job id 留在 `<PROJECT_ROOT>`；本檔只記可复用 pattern。

## 機械觀察（同一 speech window）

| 通道 | 例（去敏） | language_candidate | 角色 |
| --- | --- | --- | --- |
| OCR region A | Latin hardsub line（上） | en | dialogue subtitle／translation |
| OCR region B | CJK hardsub line（下） | zh | dialogue subtitle |
| ASR | 中文口播 | zh | spoken |

錯誤塌縮：合成單一字串或只留其中一語 → Language Relation／spoken 選錯。

## 正確分流

1. OCR → 多個 `text_region`（各自 box／language_candidate）
2. Subtitle Grouping → `bilingual_pair`（stacked + temporal overlap）
3. OCR↔ASR：zh region ↔ ASR = same_language_*；en region ↔ ASR = cross_language_translation
4. spoken canonical = 口播語（通常 ASR／CJK dialogue region）；EN 僅 translation evidence
5. 非 dialogue OCR（促銷／UI）不進 spoken／narrative

## 分類

| 標籤 | 判定 |
| --- | --- |
| `bilingual_collapsed_to_single_string` | **contract_gap candidate** |
| 不是 | 新「雙語字幕 Agent」／新 Phase |

吸收：[`41-…`](../41-bilingual-ocr-regions-and-subtitle-groups.md)＋[`text-evidence-regions.md`](../../../workflow/narrative-video-production/text-evidence-regions.md)。
