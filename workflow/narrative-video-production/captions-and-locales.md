# Captions and locales

一部片子可有多個 locale pack（同一 `film_id`）。這是內容交付契約，不是 TTS 產品。

欄位 SoT：[`records/caption-locale-pack.yaml`](records/caption-locale-pack.yaml)。

## 三閘分開（Invariant 7）

| Gate | 裁決 | 不是 |
| --- | --- | --- |
| `content_gate` | 語意、專名、source residue | 排版漂不漂亮 |
| `timing_gate` | 讀得完（CPS／cue 窗；軸秒 vs 觀眾秒要寫明） | 有沒有擋住臉 |
| `layout_gate` | 放得下、安全區、不遮擋 | 譯文對不對 |

**放得下 ≠ 讀得完 ≠ 沒擋住臉。** 禁止合成單一「字幕 PASS」。Publish QC 三閘都要有各自 `decision`。

`text_origin`：`script`／`asr`／`ocr`／`human`／`translated` 分源，不得混成一條無標記對白。
`layout_script`（cjk／latin／…）≠ `locale`。
`burn_mode`：`sidecar`／`burned`／`none`。

供應商聲音、翻譯模型、ffmpeg 濾鏡 = adapter。首輪要幾種語 = dogfood Q6。

## Content 決策（Translation Decision）

語意／專名／稱謂／locale／name realization 的 **content** 決策走  
[`workflow/translation/`](../translation/README.md)（TDR + Finality）。  
Subtitle 接線：[`workflow/translation/adapters/subtitle.yaml`](../translation/adapters/subtitle.yaml)。

`content_gate` 可引用 `cue.translation_decision_ref`；**不得**因此省略或合併 `timing_gate`／`layout_gate`。

## Speech timing authority（generated 口播）

自製破題／旁白／口播：先切 **Speech Unit**（語意＋screen-fit），再 adapter 生成語音，用 **實際 duration** 當 cue 時軸。禁止整段先 TTS 再切字幕；禁止猜秒數。TTS 過慢是 `speech_timing_gate`，不是 `layout_gate`。源片對白仍用 ASR timing。契約：[`speech-unit-and-timing.md`](speech-unit-and-timing.md)。

