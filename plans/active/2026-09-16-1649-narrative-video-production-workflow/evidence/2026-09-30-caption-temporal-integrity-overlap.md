# Observation — caption temporal integrity (same-track overlap)

**Run ID**：2026-09-30-caption-temporal-integrity-overlap

**Kind**：Phase 3 dogfood observation → timing_gate / render adapter defect classification

**Plan**：[`../01-captions-and-locales.md`](../01-captions-and-locales.md)

**Workflow**：[`../../../workflow/narrative-video-production/captions-and-locales.md`](../../../workflow/narrative-video-production/captions-and-locales.md)

## 去敏情境

雙語硬字幕參考成片再 burn 目標語 caption。觀測：成片上兩套台詞時間互相壓到。內容對不對屬 Evidence Resolution；**時間軸是否互壓屬 Rendering / Temporal Integrity**，必須機械 gate，禁止 LLM 判斷。

## 機械觀察（去敏）

| 層 | 結果 |
| --- | --- |
| Locale timed cues（Caption Pack） | same-track overlap = 0 |
| Burn ASS `Default` track | 鄰接 cue `next.start < current.end`（小數秒級） |
| Bilingual same-group regions | 同窗雙語是允許的；不可當 same-track fail |
| Source hardsub still in pixels | 若未 scrub／遮蓋，與新 burn 形成第二套可見台詞（artifact／adapter） |

## 分類

| 情況 | 標籤 |
| --- | --- |
| Caption Pack 本身 A∩B | producer／timing_gate **fail**（`cue_overlap`） |
| Pack 乾淨，ASS／burn 後互壓 | **render adapter／implementation defect** |
| Pack 無法表達 track／group | **contract_gap**（需 subtitle_group／track） |
| 不是 | 新 Phase／用 LLM 看畫面判 overlap |

## timing_gate 擴張（temporal integrity）

在既有 CPS／cue 窗外，機械必檢：

1. `cue.start < cue.end`
2. same-track：`next.start < current.end` → `fail: cue_overlap`（未明示 policy 則禁止）
3. bilingual `subtitle_group` 同組多 region → **PASS**（same_group_overlap allowed）
4. `subtitle_group` 之間互壓 → `fail: subtitle_group_overlap`
5. duplicate identical window／text → `fail: duplicate_cue`
6. impossible duration／CPS 窗
7. （layout／artifact）temporal∩spatial collision；EDR↔rendered burn-time consistency

`overlap_policy`：same_track 預設 forbidden；bilingual_same_group allowed；transition／karaoke 僅 explicit。

## Publish

```text
publish-ready =
  content PASS ∧ timing PASS ∧ layout PASS ∧ artifact PASS
  ∧ independent verifier PASS
```

Artifact gate：確認 speedup／concat／ASS merge／frame rounding 未把乾淨 pack 改壞。
