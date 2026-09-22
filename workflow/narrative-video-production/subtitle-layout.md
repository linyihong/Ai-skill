# Subtitle layout contract

`max_lines` 是 **上限**，不是目標。Constraint ≠ Selection（invariant 5）。TTS／ffmpeg 不進本檔。

欄位：[`records/caption-locale-pack.yaml`](records/caption-locale-pack.yaml) cue `layout`。時軸：[`speech-unit-and-timing.md`](speech-unit-and-timing.md)。Solver 觀察：plan [`26-subtitle-layout-engine.md`](../../plans/active/2026-09-16-1649-narrative-video-production-workflow/26-subtitle-layout-engine.md)。

## 語意

```text
line_policy:
  max_lines: 2          # upper bound
  minimize_lines: true  # 能 1 行就不拆 2 行
  target_lines: null    # 禁止把 max 當 target
```

```text
Need（這句要可讀、可放）
  → Constraints（max_lines、max_width、font [min,max]、safe area）
  → Feasible layouts[]（glyph 實測，非字數）
  → selection.policy（例 minimize_lines → readability → larger_font）
  → selected
```

LLM 可出 `semantic_break_candidates`；**不得**當最終斷行／行數裁決。

## 七條

1. `max_lines` is an upper bound, never a target.  
2. Prefer the **minimum** line count that satisfies all layout constraints.  
3. A one-line layout that fits MUST NOT be split because `max_lines > 1`.  
4. Line count uses rendered glyph dimensions, not character count.  
5. LLM may propose break points; mechanical engine is final fit authority.  
6. If one line fails width, the engine MAY evaluate two-line candidates **before** reducing font, per selection policy.  
7. If no feasible layout exists inside font／width／line-count bounds, **return to Speech Unit segmentation** — do not emit an overcrowded cue.

無法 fit 時 rollback：`speech_author`（重切 unit）或 `locale_author`（policy／profile）。
