# Captions & locales — source adapt notes

Companion to [`_plan.md`](_plan.md)。只記錄從參考包**抽象出的契約**；不複製 prompt、Azure 聲音 id、模型檔名或主機。

## 為什麼值得進 workflow

參考包把「能燒進畫面」和「人能讀完」拆開：時間軸用 **CPS + 最短／最長 cue 窗**；折行用 **語系 ScriptPattern + BreakPolicy**；多語用 **目標語短碼目錄** 與「源語言殘留 = 失敗」。這些是產片正確性，不是某一家 TTS 的實作細節。

## 建議變成 EDR 的一等公民（locale pack）

一部片子可有多個 **locale pack**（同一 `film_id`）。每個 pack 至少：

| 欄位 | 含義 |
| --- | --- |
| `locale` | BCP-47 短碼（`zh`／`en`／`nan`…），不是供應商聲音名 |
| `role` | `source`（對白／原片）／`commentary`（解說）／`caption`（觀眾看到的字） |
| `text_origin` | `script`／`asr`／`ocr`／`human`／`translated` |
| `cues[]` | `shot_id` 或軸秒 `start`/`end`、顯示文案、可選對齊源文 |
| `content_gate` / `translation_qc` | 語意、專名、source residue（**content correctness**）；譯文決策見 [`workflow/translation/`](../../../workflow/translation/README.md) + [`adapters/subtitle.yaml`](../../../workflow/translation/adapters/subtitle.yaml) |
| `timing_gate` | 讀得完：CPS／cue 窗（**timing correctness**） |
| `layout_gate` | 放得下、安全區、不遮擋（**layout correctness**） |
| `layout_script` | 排版族（cjk／latin／…），≠ locale |
| `burn_mode` | `sidecar`／`burned`／`none` |

主 EDR 填：`source_locale`、`publish_locales[]`、每個發布語是否有通過的 caption pack。

## 從參考包 adapt 的規則（工具中立）

1. **閱讀速度依語系族，不是全球同一個 CPS。** 漢字與拉丁字母不可共用同一「字／秒」。具體數字留給 profile；workflow 只要求「有語系族、有 cue 窗、有超窗處置」。
2. **超窗先拆句再截字。** 強斷點（句末）> 中斷點（逗號）> 弱斷點（空格）。禁止在語系 sticky 位置硬切（細節進 catalog，不進 execution-flow 逐步碼）。
3. **倍速播放要換算體感時長。** 軸秒與觀眾秒不同；QC 必須寫用的是哪一軸。
4. **源文字取得路徑分開記：** 腳本寫定 ≠ ASR ≠ OCR 硬字幕。混用必須在 `text_origin` 標明，不能假裝同一條對白。
5. **多語是「同一 EDR 的多 pack」，不是另開一部片。** 模板 id 通常共用；文案與 cue 不共用。
6. **譯文不得留源語腳本。** 專名規則（音譯／保留）寫在 pack 的 `name_policy`，不要把某模型 system prompt 貼進 repo。
7. **字幕-only 與配音是兩種交付。** 可以只出字幕、只出配音、或兩者；EDR 要能表達，避免「沒聲音卻當配音完成」。
8. **三閘分開，且為 workflow governance rule（不只 companion 備註）。** 放得下 ≠ 讀得完；讀得完 ≠ 沒擋住臉／安全區。Publish QC 三閘都要，禁止合成單一「字幕 PASS」。
9. **供應商聲音表、離線翻譯模型、ffmpeg 濾鏡 = adapter。** Canonical 只留 `locale` + `layout_script` + gates。

## drop

- 具體 Neural 聲音字串、Kie／Gemini／Qwen 呼叫與下載路徑
- 「電視雙行英文字幕規範」若與直式短片衝突：參考包已選直式多行 + cue 窗；我們跟短片，不跟廣播規範當預設
- 把語系目錄寫死成必須支援上述每一種語言才能完成 workflow

## 建議 workflow 檔（Phase 2）

已落地：[`workflow/narrative-video-production/captions-and-locales.md`](../../workflow/narrative-video-production/captions-and-locales.md)、[`records/caption-locale-pack.yaml`](../../workflow/narrative-video-production/records/caption-locale-pack.yaml)。

Phase 3：locale 的 `text_origin: ocr` ≠ 源片 **visual text evidence**。後者見 [`09-visual-text-evidence.md`](09-visual-text-evidence.md)。`layout_gate` 的 solver 候選：[`26-subtitle-layout-engine.md`](26-subtitle-layout-engine.md)。
