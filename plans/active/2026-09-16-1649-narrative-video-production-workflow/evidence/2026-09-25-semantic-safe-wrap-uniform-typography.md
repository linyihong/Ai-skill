# Observation — semantic-safe wrap and uniform typography

**Run ID**：2026-09-25-semantic-safe-wrap-uniform-typography  
**Kind**：contract（實際製片反例；已寫進 workflow）  
**Extends**：[`2026-09-22-max-lines-is-bound.md`](2026-09-22-max-lines-is-bound.md)、
[`2026-09-24-font-size-hard-bounds.md`](2026-09-24-font-size-hard-bounds.md)

## 反例

- 為 fit 在詞內切 cue：句尾「談」／下一 cue「話」，語意與 timing 都被破壞。
- 同一個兩行 cue 各自 fit，造成上行大、下行小。

## Contract

Wrap 只加 line break；`concat(lines)` 必須還原 cue text。Break offset 不得在詞彙、
專名、數字＋單位、成語／固定短語或人工鎖定 span 內。Segmentation 是上游
Speech Unit 決策；generated speech 重切後必須重做 artifact／timing。

同一 cue 的 font family／size／line height／stroke 一致。不存在 per-line font
selection。無 uniform feasible layout → resegment／reject。
