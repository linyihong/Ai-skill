# Observation — dialogue cue projection must retain subtitle_group

**Run ID**：2026-09-30-dialogue-cue-projection-retains-subtitle-group

**Kind**：Phase 3 dogfood observation → **`contract_gap` confirmation**（adapter projection）

**Plan**：[`../41-bilingual-ocr-regions-and-subtitle-groups.md`](../41-bilingual-ocr-regions-and-subtitle-groups.md)

**Extends**：[`2026-09-29-bilingual-hardsub-regions.md`](2026-09-29-bilingual-hardsub-regions.md)、[`2026-09-29-cross-language-subtitle-spoken-alignment.md`](2026-09-29-cross-language-subtitle-spoken-alignment.md)

## 去敏情境

參考成片硬燒 **中英雙語字幕**；口播為中文。Layer-0 OCR／dialogue cue 已能產出 bilingual `subtitle_group`＋per-region OCR。下游 locale pack／burn 卻把 cue **壓成** `spoken_timed{src,dst,start,end}`，雙語 region 與 speech alignment 丟失。

## 機械觀察

| 層 | 狀態 | 備註 |
| --- | --- | --- |
| Layer-0 OCR（hardsub） | 保留 glued text + `parts[]`／boxes | **不需重做 schema** |
| Dialogue cue | 已有 `subtitle_group`（bilingual）+ `ocr_regions`（zh／en） | region 已分語 |
| Locale pack export | 只寫 `{start,end,text}` + `kind: spoken_timed` | **projection 丟結構** |
| Target-lang burn | 從 spoken 再翻譯 | 忽略已存在的 OCR 異語 region |

## 契約澄清

`spoken_timed` = **spoken／口播語** 的 burn／phrase 投影，**不得**充當 subtitle 原始結構。

正確 Dialogue Cue Group 投影：

```text
dialogue_cue_group:
  speech: { start, end, text, language }          # ASR／spoken
  subtitle_group: { id, start, end, type, regions[] }
  alignment: { speech_id, subtitle_group_id }
```

`regions[]` 各帶 `text`／`language`／`script`（translation evidence ≠ spoken）。

## 分類

| 標籤 | 判定 |
| --- | --- |
| `bilingual_collapsed_to_single_string` | 上游部分已吸收；**downstream projection 仍塌縮** → adapter **contract_gap** |
| 不是 | 重做 Layer-0 OCR schema／新雙語 Agent |

## Adapter 驗收

1. Export／load timed pack 保留 `subtitle_group`＋`ocr_regions`＋`speech`／alignment
2. `spoken_timed` entries 僅 spoken 側；不得刪除 cue 上的 subtitle structure
3. 目標語 burn：若 `subtitle_group` 已有該語 region，優先作 subtitle evidence（可再做 boundary recovery），而非只從 spoken 重譯並丟 region
