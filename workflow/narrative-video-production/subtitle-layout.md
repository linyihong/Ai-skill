# Subtitle layout contract

`max_lines` 是 **上限**，不是目標。Constraint ≠ Selection（invariant 5）。TTS／ffmpeg 不進本檔。

欄位：[`records/caption-locale-pack.yaml`](records/caption-locale-pack.yaml) cue `layout`。時軸：[`speech-unit-and-timing.md`](speech-unit-and-timing.md)。Solver 觀察：plan [`26-subtitle-layout-engine.md`](../../plans/active/2026-09-16-1649-narrative-video-production-workflow/26-subtitle-layout-engine.md)。

## 語意

```text
line_policy:
  max_lines: 2
  minimize_lines: true
  target_lines: null

typography.font_size:
  preferred: 48
  min: 36          # 硬下限；到此仍 fit 失敗 → 重切 Speech Unit
  max: 56          # 硬上限；短句不得因畫面空而放大
  step: 2          # 離散字級；禁止每 cue 連續 ±1 造成跳變

overflow_policy.order:
  - try_one_line
  - try_two_lines
  - reduce_font_size_within_bounds
  - resegment_speech_unit
  - reject
```

`min ≤ actual_font_size ≤ max` 由 **mechanical engine 強制**。AI 可建議語意切點，**不得**把字縮到 min 以下「塞進去」，也不得把短句放到 max 以上。

Scene 已有主字級時，超過 `stability.max_delta_from_scene` 的縮字視為應重切 unit，不是再縮。

```text
Need（這句要可讀、可放）
  → Constraints（max_lines、max_width、font [min,max]、step、safe area）
  → Feasible layouts[]（glyph 實測，非字數）
  → selection.policy（例 minimize_lines → readability → larger_font）
  → selected
```

LLM 可出 `semantic_break_candidates`；**不得**當最終斷行／行數／字級裁決。

## Wrap ≠ segmentation；換行必須 lossless

- **Line wrap** 只插入顯示換行，不得刪字、補字、換字或重排。移除換行後的
  rendered text 必須等於 canonical cue text。
- **Cue／Speech Unit segmentation** 會建立新的時間單位，只能在上游的有效語意
  邊界發生；layout engine 不得為了 fit 自行截字。
- Break offset 不得落在 `protected_spans` 內。至少保護：詞彙／專名、數字＋單位、
  成語／固定短語與人工鎖定 span。例如「談話」不可變成 cue A「談」＋cue B「話」。
- 若沒有合法 break，回傳 `no_semantic_break`，再走字級 bounds 或 Speech Unit
  resegment；禁止用字元索引硬切。

Mechanical integrity：

```text
remove_only_inserted_linebreaks(rendered_lines) == cue.text
all(line_break_offsets ∉ protected_spans)
source_span_coverage == complete_and_non_overlapping
```

任何缺字／重複／順序改變都是 `layout_gate: fail`。

## 同一 cue 的 typography 必須一致

一個 cue 只選一次 `font_family`、`font_size`、`line_height`、stroke／outline。
兩行字幕不得各自 auto-fit，禁止上行大、下行小。若任一行 overflow，整個 cue
共用同一候選字級重新求解；仍無可行解就 resegment／reject。跨 cue／scene 的穩定性
另由 profile 與 temporal review 管理。

## 七條＋字級硬閘

1. `max_lines` is an upper bound, never a target.  
2. Prefer the **minimum** line count that satisfies all layout constraints.  
3. A one-line layout that fits MUST NOT be split because `max_lines > 1`.  
4. Line count uses rendered glyph dimensions, not character count.  
5. LLM may propose break points; mechanical engine is final fit authority.  
6. If one line fails width, evaluate two-line candidates **before** reducing font, per selection policy.  
7. Font size MUST stay in `[min, max]` on `step`／`allowed` values. Hitting min without fit → **Speech Unit resegment**, never shrink further. Short cues MUST NOT exceed max.  
8. If no feasible layout exists inside those bounds, **return to Speech Unit segmentation** — do not emit an overcrowded or unreadably small cue.
9. Wrapping MUST preserve every source character and MUST NOT break a protected span.
10. Every rendered line in one cue MUST use the same selected typography.

無法 fit 時 rollback：`speech_author`（重切 unit）或 `locale_author`（policy／profile）。
