# Subtitle adapter（NVP）

Doc index for [`subtitle.yaml`](subtitle.yaml).

**Owns**: Translation Decision → `content_gate` / cue text.  
**Does not own**: `timing_gate`、`layout_gate`、burn／TTS／ffmpeg.

Static walkthrough A+B PASS：[`08-static-walkthrough-pass.md`](../../../plans/active/2026-09-22-1000-translation-decision-workflow/08-static-walkthrough-pass.md)。

Dogfood 驗收分欄／Reference／Semantic／ja-JP lexical 責任：
[`dogfood-acceptance.md`](dogfood-acceptance.md)。接產品驗收摘要時讀取；
不新增 workflow、不讓機械 pass 代替 semantic Finality。

## Product Selection Actor wiring（integration note）

When a product dub／subtitle pipeline uses an LLM provider:

1. Inject **Selection adapter notes only**（constraints + Candidate Space seeds）— not per-locale case lists.
2. Run **Independent Validation** after Selection（locale／honorific／kinship residue／numeric／temporal guards）.
3. Map to **Finality** ternary before publish：`blocked`（e.g. I21 truncated source）and hard residue must not write cue／dub text；`needs_review` is explicit handoff；`accepted` only when gates pass.
4. Kinship／address seeds expand Candidate Space（I23）；they are **not** sole fixed glosses.

Reusable lesson：[`2026-09-29_105651-translation-decision-selection-adapter-finality-gates.md`](../../../feedback/history/development-guidance/common/2026-09-29_105651-translation-decision-selection-adapter-finality-gates.md)。
