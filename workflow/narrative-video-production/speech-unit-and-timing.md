# Speech unit and timing authority

Generated 旁白／破題／自製口播的 **時間軸契約**。TTS／SSML／聲音 provider／ffmpeg **不進**本檔。

欄位 SoT：[`records/speech-unit.yaml`](records/speech-unit.yaml)。字幕驗收仍見 [`captions-and-locales.md`](captions-and-locales.md)。Layout solver 見 plan [`26-subtitle-layout-engine.md`](../../plans/active/2026-09-16-1649-narrative-video-production-workflow/26-subtitle-layout-engine.md)。

## Timing authority

| 來源 | cue `start`/`end` 的 primary evidence |
| --- | --- |
| 源片對白／硬字幕 | ASR（或 OCR persistence）timing |
| 自製口播／TTS 產出 | **Speech artifact duration**（adapter 回寫） |

禁止用 AI 猜「這句大概 3 秒」當 generated-speech cue 軸。`timing_gate` 讀的是 evidence-backed 時長，不是 layout 漂不漂亮。

## 順序

Speech Unit 是語意／語音 SoT。Caption 是它的視覺投影，不得為了放得下而改 unit 原文。

```text
Script
  → Speech segmentation（標點候選；太碎才合併。不看字級／行數／安全區）
  → Speech Unit SoT
  → Speech generation（adapter）→ actual duration
  → Caption cue timing（start/end = 該 unit 的語音時間）
  → Line break／typography（只在這個 unit 內；見 [`subtitle-layout.md`](subtitle-layout.md)）
```

三件事分開：segmentation 決定怎麼說、cue 決定哪一段時間顯示哪句、line break 決定這句在畫面上怎麼排。

標點建立 boundary candidate，不是 100% 強制切開。`他说：“等等，我还没说完。”` 可以合併過短的逗號 unit。合併或重切屬於 speech loop，而且必須重做語音與 timing。

Layout 無可行換行時回 `layout_blocked` 給 `speech_author`。它不得把「谈话」切成「谈｜话」，也不得把一行硬拆成兩行。

## 兩個 loop

- **Speech loop**：切 unit → 生成 → timing QC → 重切或調語速  
- **Layout loop**：cue 吃真實時軸 → layout（1 行先於 2 行）→ visual review（[`28-layout-review-loop.md`](../../plans/active/2026-09-16-1649-narrative-video-production-workflow/28-layout-review-loop.md)）。無可行 layout → 回到 Speech Unit Planning，不是擠兩行。

失敗 rollback：`speech_author`（timing QC）／`locale_author`（layout）。
