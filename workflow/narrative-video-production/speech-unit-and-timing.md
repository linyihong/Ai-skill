# Speech unit and timing authority

Generated 旁白／破題／自製口播的 **時間軸契約**。TTS／SSML／聲音 provider／ffmpeg **不進**本檔。

欄位 SoT：[`records/speech-unit.yaml`](records/speech-unit.yaml)。字幕驗收仍見 [`captions-and-locales.md`](captions-and-locales.md)。Layout solver 見 plan [`26-subtitle-layout-engine.md`](../../plans/active/2026-09-16-1649-narrative-video-production-workflow/26-subtitle-layout-engine.md)。

## Timing authority

| 來源 | cue `start`/`end` 的 primary evidence |
| --- | --- |
| 源片對白／硬字幕 | ASR（或 OCR persistence）timing |
| 自製口播／TTS 產出 | **Speech artifact duration**（adapter 回寫） |

禁止用 AI 猜「這句大概 3 秒」當 generated-speech cue 軸。`timing_gate` 讀的是 evidence-backed 時長，不是 layout 漂不漂亮。

## 順序（不是整段先 TTS 再切字幕）

```text
Script → Speech Unit Planning（語意邊界＋screen-fit＋max_lines／width＋語系）
  → Speech Generation（adapter）→ Actual duration
  → Speech Timing QC（CPS／語速／pause；過慢是 speech 問題）
  → Caption cue timing（primary = speech timing）
  → Caption layout（**minimize_lines**；max_lines 是上限；見 [`subtitle-layout.md`](subtitle-layout.md)）
```

Speech timing ≠ caption layout。太長要分責：文案切 unit／TTS 語速／layout 擁擠。兩個 loop 不得合成「字幕不好看請重做」。

## 兩個 loop

- **Speech loop**：切 unit → 生成 → timing QC → 重切或調語速  
- **Layout loop**：cue 吃真實時軸 → layout（1 行先於 2 行）→ visual review（[`28-layout-review-loop.md`](../../plans/active/2026-09-16-1649-narrative-video-production-workflow/28-layout-review-loop.md)）。無可行 layout → 回到 Speech Unit Planning，不是擠兩行。

失敗 rollback：`speech_author`（timing QC）／`locale_author`（layout）。
