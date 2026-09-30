# Observation — OCR caption omitted by same-ASR collapse

**Run ID**：2026-09-30-ocr-caption-omission-same-asr-collapse

**Kind**：Phase 3 dogfood observation → projection／omission／adapter defect（非新 Phase）

**Plan**：[`../01-captions-and-locales.md`](../01-captions-and-locales.md)

**Workflow**：[`../../../workflow/narrative-video-production/captions-and-locales.md`](../../../workflow/narrative-video-production/captions-and-locales.md)

## 去敏情境

雙語硬字幕短劇。L0 OCR 在相鄰秒有多條獨立對白（含短中文＋Latin region）。成片／dialogue cues 缺其中一句。觀測者以為「字幕生成漏字」；實際是 resolve 後 **same-ASR overlapping collapse** 把 text-unrelated cue 合併丟掉。

## 機械觀察（去敏）

| 層 | 結果 |
| --- | --- |
| OCR／visual text | 短 cue 存在（含 bilingual parts／regions） |
| fuse + TextAlignment | 該 cue 仍獨立（與鄰句 `unrelated`） |
| resolve + same-ASR collapse | 多條 cue 因 latched 同一長 ASR observation → 合併成寬窗，只留 rank 最高文案 |
| dialogue cues／ASS／成片 | 缺原短句 |

## 分類

| 情況 | 標籤 |
| --- | --- |
| OCR 無句 | evidence／OCR miss |
| fuse／alignment 丟 | fusion／TextAlignment defect |
| resolve collapse 因 same ASR + unrelated 文丟 | **projection／omission**（`unrelated_forced_merge`） |
| Timeline／Pack 有、MP4 無 | render adapter defect |
| 不是 | 新 Phase；用 LLM 看片猜「有沒有這句」 |

## 契約補強

1. **Timeline IR** 為 EDR→render 一等中間產物；ASS／NLE XML 是 adapter。
2. **Coverage／omission report**：selected／高品質 subtitle candidate 無 downstream → `suspicious_omission`。
3. **same-ASR merge 門檻**：僅 text relation ∈ {duplicate, truncated_variant_of, variant_of}；`unrelated` 禁止 destructive merge。
4. **subtitle_group** 整組 trace，禁止單 region silent drop。

## Publish

```text
publish-ready 另需：
  timeline projection PASS
  ∧ coverage 無未解釋 suspicious_omission
  ∧ caption artifact ↔ Timeline IR 一致
```
