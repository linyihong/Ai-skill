# Caption composition contract

先保語意，再排版。`max_lines` 是 **上限**，不是目標。Constraint ≠ Selection（invariant 5）。TTS／ffmpeg 不進本檔。

欄位：[`records/caption-locale-pack.yaml`](records/caption-locale-pack.yaml) cue `layout`。時軸：[`speech-unit-and-timing.md`](speech-unit-and-timing.md)。Solver 觀察：plan [`26-subtitle-layout-engine.md`](../../plans/active/2026-09-16-1649-narrative-video-production-workflow/26-subtitle-layout-engine.md)。語意斷點：[`33-caption-composition-semantic-breaks.md`](../../plans/active/2026-09-16-1649-narrative-video-production-workflow/33-caption-composition-semantic-breaks.md)。

## 四層（優先序固定）

```text
1. Content Integrity     不可破壞語意單位（atomic / protected spans）
2. Caption Segmentation  要不要拆成兩個 cue（Speech Unit；時間軸）
3. Line Composition      同一 cue 一行或兩行（wrap；lossless）
4. Typography Constraint min／max／step；同 cue 統一字級
```

禁止：先塞畫面 → 塞不下就切字 → 還塞不下就縮到看不見。

**Speech Segmentation ≠ Line Breaking。** 前者改 cue／時軸；後者只插入顯示換行。Layout engine **不得**自創任意斷點，只能從 `semantic_break_candidates` 裡選。

## 決策順序

```text
Script → Speech Unit Planning → Semantic Segmentation → Caption Candidate
  → Semantic Break Candidates（AI 評分；mechanical 裁決）
  → Try 1-line → fit? YES accept
  → Try 2-line（只在 scored 合法斷點）→ fit? YES accept
  → Reduce font size（整 cue 同一字級，∈ [min,max] step）→ retry 1-line／2-line
  → font < min? → Re-segment Speech Unit（完整語意邊界）→ reject overcrowded cue
```

```text
line_policy:
  max_lines: 2
  minimize_lines: true
  target_lines: null
  equal_line_length: forbidden   # 禁止字數／2

typography.font_size:
  preferred: 48
  min: 36
  max: 56
  step: 2

overflow_policy.order:
  - try_one_line
  - try_two_lines
  - reduce_font_size_within_bounds
  - resegment_speech_unit
  - reject
```

`min ≤ actual_font_size ≤ max` 由 **mechanical engine 強制**。Scene 已有主字級時，超過 `stability.max_delta_from_scene` 的縮字視為應重切 unit。

```text
Need → Constraints → Feasible layouts[]（glyph 實測）
  → selection.policy（minimize_lines → semantic_break_score → readability → larger_font）
  → selected
```

兩行候選 **不得**用平均長度當目標。優先：語意單位、詞組、標點、動詞結構、人名／專名。例如「徐經理」「談話」是 atomic unit。

## AI vs mechanical

| AI | Mechanical |
| --- | --- |
| Semantic segmentation 建議 | Glyph 寬、max lines、safe area |
| `semantic_break_candidates` + score | 只從候選選 offset；forbidden 落點直接淘汰 |
| Break quality／meaning preservation | Font [min,max] step；cue-uniform typography |
| Visual review（adjustment candidate） | Overflow；forbidden region |

AI 不得把「談話」拆成「談／話」。Mechanical 不接受 unscored／protected 內的 break。

`break_protection` 來自 locale 的 linguistic units（詞／專名／數字＋單位／成語），**不是**每次讓 AI 寫一長串 `forbidden_boundaries`。AI 可標 `preferred_boundaries`；最終仍走 scored candidates。

## Wrap ≠ segmentation；換行必須 lossless

- Wrap 只插入顯示換行。移除換行後必須等於 canonical cue text。
- Cue／Speech Unit 切分只能在上游語意邊界。無合法 break → `no_semantic_break`，再走字級或 resegment。
- `line_break_offsets` 必須 ∈ scored candidates，且 ∉ `protected_spans`。

```text
remove_only_inserted_linebreaks(rendered_lines) == cue.text
all(line_break_offsets ∈ scored_semantic_break_offsets)
all(line_break_offsets ∉ protected_spans)
source_span_coverage == complete_and_non_overlapping
```

缺字／重複／重排／未評分斷點 → `layout_gate: fail`。

## 同一 cue 的 typography 必須一致

一個 cue 只選一次 `font_family`、`font_size`、`line_height`、stroke。禁止上行大、下行小。任一行 overflow，整個 cue 用同一字級重算。

## 契約條

1. Content integrity precedes layout fit.
2. `max_lines` is an upper bound, never a target.
3. Prefer the **minimum** line count that fits.
4. A one-line layout that fits MUST NOT be split because `max_lines > 1`.
5. Two-line wrap MUST NOT optimize for equal character/pixel length.
6. Line count uses rendered glyph dimensions, not character count.
7. Line breaks MUST be chosen from scored `semantic_break_candidates`; LLM is not final fit authority.
8. Font size MUST stay in `[min, max]` on `step`／`allowed`. Hitting min without fit → Speech Unit resegment.
9. Every rendered line in one cue MUST use the same typography.
10. Wrapping MUST preserve every source character and MUST NOT break a protected span.

無法 fit 時 rollback：`speech_author`（重切 unit）或 `locale_author`（policy／profile）。
