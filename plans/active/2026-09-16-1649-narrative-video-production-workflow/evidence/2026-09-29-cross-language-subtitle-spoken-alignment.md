# Observation — cross-language subtitle ↔ spoken alignment

**Run ID**：2026-09-29-cross-language-subtitle-spoken-alignment
**Kind**：Phase 3 dogfood observation → **`contract_gap` candidate**（`cross_language_collapsed_to_conflict`）
**Plan**：[`../40-language-role-before-text-resolution.md`](../40-language-role-before-text-resolution.md)
**Extends**：[`2026-09-21-spoken-vs-subtitle-reconstruction.md`](2026-09-21-spoken-vs-subtitle-reconstruction.md)、[`2026-09-21-sanitization-anomaly-audit.md`](2026-09-21-sanitization-anomaly-audit.md)、[`2026-09-28-asr-validity-precedes-interpretation.md`](2026-09-28-asr-validity-precedes-interpretation.md)

## 去敏情境

短劇來源：畫面硬燒 **英文**字幕；音軌為 **中文**口播（站點語系標記為 zh）。
`external_run_ref` 與媒體路徑留在 `<PROJECT_ROOT>`；本檔只記可复用 pattern。

## 機械觀察（同一 speech window）

| 通道 | 例（去敏） | language | role |
| --- | --- | --- | --- |
| OCR observed | Latin hardsub span（常無空白黏字） | `en` | hardcoded／translated subtitle candidate |
| ASR spoken | 中文口播 span | `zh` | spoken |

兩邊時間可對齊，語義常為互譯。ASR 中文候選存在且語速／時長往往更合理；最終 canonical `text` 卻選了 OCR 英文，並標成同窗 `semantic_mismatch`／由 LLM 重建「選 OCR」。

## 錯誤抽象

把「OCR=en、ASR=zh」直接當：

- OCR／ASR conflict，或
- OCR 優先 spoken，或
- 和諧／anomaly

都錯。這是 **`cross_language_translation` 候選**，必須先過 Language Relation Gate。

## 正確分流

1. Mechanical：`language(ocr) ≠ language(asr)` → `candidate_relation = cross_language`
2. LLM（僅 alignment）：語義是否支持 translation／subtitle relation
3. Supported → 保留 `subtitle_text`（en）與 `spoken_text`（zh）；**spoken canonical = ASR／reconstruction 中文**
4. Sanitization **不**觸發（非 same-language semantic conflict）

## 分類

| 標籤 | 判定 |
| --- | --- |
| `cross_language_collapsed_to_conflict` | **contract_gap candidate** |
| 不是 | 立刻換 ASR 模型／新 Phase／新 Agent |

吸收：既有 `spoken_text`／`subtitle_text`／`text_relation`／`text_span_role`＋本 plan companion；workflow 見 [`text-evidence-language-relation.md`](../../../workflow/narrative-video-production/text-evidence-language-relation.md)。產品 fuse 適配另開，不得反向定義契約。
