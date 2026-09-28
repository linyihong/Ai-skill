# Caption composition contract

先保語意，再排版。`max_lines` 是 **上限**，不是目標。Constraint ≠ Selection（invariant 5）。TTS／ffmpeg 不進本檔。

欄位：[`records/caption-locale-pack.yaml`](records/caption-locale-pack.yaml) cue `layout`。時軸：[`speech-unit-and-timing.md`](speech-unit-and-timing.md)。Solver 觀察：plan [`26-subtitle-layout-engine.md`](../../plans/active/2026-09-16-1649-narrative-video-production-workflow/26-subtitle-layout-engine.md)。語意斷點：[`33-caption-composition-semantic-breaks.md`](../../plans/active/2026-09-16-1649-narrative-video-production-workflow/33-caption-composition-semantic-breaks.md)。

## 四層（優先序固定）

```text
1. Content Integrity     不可破壞語意單位（atomic / protected spans）
2. Caption Segmentation  要不要拆成兩個 cue（Speech Unit；時間軸）
3. Line Composition      同一 cue 一行或兩行（wrap；lossless）
4. Typography Constraint 三層字級；同 cue 統一；視覺尺度≠同一 px
```

禁止：先塞畫面 → 塞不下就切字 → 還塞不下就縮到看不見。

**Speech Segmentation ≠ Line Breaking。** 前者改 cue／時軸；後者只插入顯示換行。Layout engine **不得**自創任意斷點，只能從 `semantic_break_candidates` 裡選。

## 決策順序

```text
Script → Speech Unit Planning → Semantic Segmentation → Caption Candidate
  → Semantic Break Candidates（AI 評分；mechanical 裁決）
  → Try 1-line → fit? YES accept
  → Try 2-line（只在 scored 合法斷點）→ fit? YES accept
  → Reduce font size（整 cue 同一字級；只在 profile＋max_delta 內）→ retry 1-line／2-line
  → 低於 profile min？ → Re-segment Speech Unit（完整語意邊界）→ reject overcrowded cue
```

```text
line_policy:
  max_lines: 2
  minimize_lines: true
  target_lines: null
  equal_line_length: forbidden   # 禁止字數／2

typography:
  absolute:                 # 安全底線；不可突破。不是單句可用範圍
    min: 36
    max: 56
  profile:                  # 由 layout_script 選，不是 locale。px 留給 dogfood
    preferred: <profile>
    min: <inside absolute>
    max: <inside absolute>
  adjustment:
    max_delta: 4            # 相對 preferred；超出視為另一視覺尺度 → 重切
  visual_target:
    mode: glyph_box         # 跨語系對齊觀看尺寸，不是同一 font_size
  consistency:
    scope: scene
    max_delta: 4

overflow_policy.order:
  - try_one_line
  - try_two_lines
  - reduce_font_size_within_profile_delta
  - resegment_speech_unit
  - reject
```

三層同時成立：`absolute.min ≤ profile.min ≤ preferred − max_delta ≤ actual ≤ preferred + max_delta ≤ profile.max ≤ absolute.max`。到 **profile min** 仍放不下 → 重切 Speech Unit，禁止繼續縮到 absolute min。Profile 的 preferred／min／max 與 `visual_scale` **不在本檔凍死**；locale 只選 `layout_script` profile。

跨 script 用 glyph bounding box 的視覺高度補償，不用「某語言固定 +N px」表。AI 只能出方向與建議區間；mechanical 在 profile＋delta 內選字級。Scene 先定 baseline，cue 優先靠近它。

```text
Need → Constraints → Feasible layouts[]（glyph 實測）
  → selection.policy（minimize_lines → semantic_break_score → readability → larger_font）
  → selected
```

兩行候選 **不得**用平均長度當目標。優先：語意單位、詞組、標點、動詞結構、人名／專名。例如「徐經理」「談話」是 atomic unit。

## Break candidate system

Phase 1 schema：[`records/break-candidate.yaml`](records/break-candidate.yaml)。

`BreakPolicy` 只提供 mechanical constraints／features。`generate_break_candidates` 找出可能斷點，不做最終決策。`hard_violation`（`lexical_unit_split`｜`function_word_dangling`）才把候選趕出可行集。`phrase_integrity` 與 Phase 1 的 `semantic_boundary: unknown` 只是特徵，不能單獨排除。`balance_score` 是左右平衡（越大越平衡）；`score` 是機械成本（越小越好）。沒有 Selection Actor 時，`best_cut` = 去掉 violation 後取最小 `score`。`window` 是搜尋範圍，不是語意範圍；不得只在 `prefer_at` 旁邊幾個字裡找。

Phase 2 才讓 LLM 在可行集內選擇並寫 `break_evidence`。Phase 1 沒有 `source`。單集反例不加成特例。

## Wrap ≠ segmentation；換行必須 lossless

- Wrap 只插入顯示換行。移除換行後必須等於 canonical cue text。
- Cue／Speech Unit 切分只能在上游語意邊界。無合法 break → `no_semantic_break`，再走字級或 resegment。
- `line_break_offsets` 必須 ∈ feasible set（`hard_violation: none`），且 ∉ `protected_spans`。

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
7. Line breaks MUST be chosen from the mechanical feasible set. LLM is the semantic selector inside that set and is not the fit authority.
8. Font size MUST satisfy absolute bound **and** profile bound **and** `max_delta` from preferred. Hitting profile min without fit → Speech Unit resegment, never slide to the absolute floor.
9. Every rendered line in one cue MUST use the same typography.
10. Wrapping MUST preserve every source character and MUST NOT break a protected span.
11. Cross-script consistency targets glyph visual size, not an identical `font_size` number. Script compensation numbers stay in profile／dogfood.
12. Only `hard_violation` removes a Phase 1 candidate from the feasible set. `phrase_integrity` is advisory. `best_cut` is the min-score fallback, not a second selection engine.

無法 fit 時 rollback：`speech_author`（重切 unit）或 `locale_author`（policy／profile）。
