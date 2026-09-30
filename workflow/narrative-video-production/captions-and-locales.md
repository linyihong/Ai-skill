# Captions and locales

一部片子可有多個 locale pack（同一 `film_id`）。這是內容交付契約，不是 TTS 產品。

欄位 SoT：[`records/caption-locale-pack.yaml`](records/caption-locale-pack.yaml)。

## 三閘分開（Invariant 7）

| Gate | 裁決 | 不是 |
| --- | --- | --- |
| `content_gate` | 語意、專名、source residue | 排版漂不漂亮 |
| `timing_gate` | 讀得完（CPS／cue 窗）**＋ temporal integrity**（同 track 互壓／非法窗） | 有沒有擋住臉 |
| `layout_gate` | 放得下、安全區、不遮擋（含空間碰撞） | 譯文對不對 |

**放得下 ≠ 讀得完 ≠ 沒擋住臉。** 禁止合成單一「字幕 PASS」。Publish QC 三閘都要有各自 `decision`。

### `timing_gate` — temporal integrity（機械，非 LLM）

字幕內容對錯屬 Evidence Resolution；**時間軸是否互壓屬 Rendering／Temporal Integrity**。

機械必檢（Caption Pack 與 burn ASS／events）：

| 檢查 | 未過 |
| --- | --- |
| `cue.start < cue.end` | `impossible_window` |
| same-track：`next.start < current.end`（未明示 policy） | `cue_overlap` |
| 同 `subtitle_group` 多 locale region 同窗 | **PASS**（`bilingual_same_group` allowed） |
| 不同 `subtitle_group` 在同一視覺字幕層互壓 | `subtitle_group_overlap` |
| 完全相同窗／文案重複 | `duplicate_cue` |
| CPS／min–max cue 窗 | 既有超窗規則 |

`overlap_policy` 預設：`same_track=forbidden`；`bilingual_same_group=allowed`；`transition`／`karaoke` 僅 `explicit`。

分類：Pack 本身互壓 → producer／timing fail；Pack 乾淨、成片／ASS 互壓 → **render adapter defect**（倍速、ASS merge、硬字幕未 scrub）；無法表達 track／group → contract_gap。見 evidence `2026-09-30-caption-temporal-integrity-overlap`。

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

自製破題／旁白／口播：先建立 **Speech Unit**（標點候選，不看字級與行數），再逐 unit 生成語音，用 **實際 duration** 當 cue 時軸。Caption 只投影該 unit，不得改 unit 原文。禁止整段先 TTS 再切字幕；禁止猜秒數。TTS 過慢是 `speech_timing_gate`。源片對白仍用 ASR timing。契約：[`speech-unit-and-timing.md`](speech-unit-and-timing.md)。

## Layout：Caption Composition（語意先於 fit）

`layout_gate` 用 glyph 可行集＋`selection.policy`（`minimize_lines`，再看同一 tier 的 break）。
一行放得下禁止因 `max_lines>1` 而拆行。兩行 **禁止** 以字數均分為目標。換行只在一個 Speech Unit 內。
選擇是 `natural_boundary_first`。Schema：[`records/break-candidate.yaml`](records/break-candidate.yaml)。字級是三層：absolute 安全底線、`layout_script`
profile、相對 preferred 的窄 `max_delta`。到 **profile min** 仍放不下 → 重切 Speech Unit，
禁止滑到 absolute floor。跨語系對齊 glyph 視覺高度，不是同一 px。Profile 數字不在本檔凍死。
契約：[`subtitle-layout.md`](subtitle-layout.md)。cue 應記
`layout.lines` 與 `layout.max_lines` 分開。

換行是 lossless display transform：去掉換行後必須等於 cue 原文，且不得切進
`protected_spans`。同一 cue 的所有行共用一套 typography；禁止上行大、下行小。

