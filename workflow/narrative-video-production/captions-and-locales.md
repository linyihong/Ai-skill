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
