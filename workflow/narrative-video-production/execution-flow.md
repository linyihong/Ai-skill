# Narrative Video Production — Execution Flow

Canonical lifecycle。各 stage 填哪個 record、能否推進：見
[`artifact-gates.md`](artifact-gates.md) 與 [`records/`](records/README.md)。
**不要**在本檔寫 ffmpeg／TTS／模型呼叫。

## Lifecycle

```text
0. Frame                 → 敘事影音成品（非 3D 資產、非寫產片器）
1. Intake / Brief lock   → 平台、時長、受眾、一句承諾、禁止項
1b. Source bible         → 源作品共用實體 id
1c. Clip catalog         → 可剪片段入庫；查找只回既有 clip_id
2. Template select       → 主 narrative_template_id + template_kind slot
3. Matching script       → Need → Constraints → feasible set → selection.policy → selected
4. EDR open              → 決策 SoT；shot 對齊腳本／bible／clip_id
5. Continuity            → 需要時鎖角色／場景／風格
6. Acquisition           → 策略可換；結果回寫 EDR（不得用 mp4 當 SoT）
6b. Text evidence        → OCR Layout Discovery／scan_profile（probe≠exclusion）→ targeted OCR raw＋boundary recovery＋text_region／subtitle_group → language／text_role → Language Relation Gate → Text Resolution（跨語言／雙語≠conflict；翻譯列≠spoken；黏字串≠靜默 SoT；bottom miss≠無字幕）
7. Assemble vs EDR       → Timeline IR 對 shot_id／selected_clip_id；機械 coverage；source 軸決策，publish transform 另記
7a. Timeline transform   → constant_speed（或未來 ramp）寫入 EDR；不得對 sped media 重跑 OCR／ASR／matching（見 source-publish-timebase）
7b. Locale packs         → content／timing／layout；timing_gate 在 publish 軸；content 源 = resolved spoken／subtitle 對；generated 口播 cue 時軸 = speech artifact（TTS=adapter）；EDR→Timeline projection 不得 silent drop
8. Publish QC            → 平台規格；temporal_integrity ≠ presentation_comfort；publish-ready 需 fresh verification
9. Outcome window        → evidence_status 回寫模板假設（非 truth）
```

## Maturity

| 級 | 意義 | 誰可簽 |
| --- | --- | --- |
| exploration | 可無完整 gate；不得宣稱可發布 | producer |
| cut-ready | EDR 與成片時間線對得上；locale 可未齊 | producer 自驗可推進 |
| publish-ready | blocking QC + locale 三閘 + **fresh verifier** | 獨立 completion authority |
| outcome-scored | 窗口填完 | 不是流量好；只是 evidence 已記 |

`publish-ready` 不得由 producer／同一 agent 自簽。

## Stage 明細

| Stage | 填／讀 | 推進條件（欄位） | 失敗 rollback |
| --- | --- | --- | --- |
| 0 Frame | 本檔 | 任務是敘事成片 | — |
| 1 Intake | [`intake.md`](intake.md) | brief lock 完整 | intake_author |
| 1b Bible | [`source-bible.md`](source-bible.md) | `bible_id` + 實體表 | bible_owner |
| 1c Catalog | [`material-clip-catalog.md`](material-clip-catalog.md) | 每 clip 有 in/out／`clip_id` | catalog_owner |
| 2 Template | [`narrative-template-catalog.md`](narrative-template-catalog.md) | 主 id ∈ catalog；有 `template_kind` | template_author |
| 3 Matching | [`matching-script.md`](matching-script.md) | 每 shot：可行集 + policy + selected ∈ 可行集 | matching_author |
| 4 EDR | [`edit-decision-record.md`](edit-decision-record.md) | 結構化 EDR 存在；對齊 script | edr_author |
| 6 Acquisition | EDR `shots[]` 回寫 | 實際入出點仍指向 `selected_clip_id` 或記 mutation | acquisition |
| 6b Text evidence | [`text-evidence-ocr-discovery.md`](text-evidence-ocr-discovery.md)、[`text-evidence-ocr-boundary.md`](text-evidence-ocr-boundary.md)、[`text-evidence-regions.md`](text-evidence-regions.md)、[`text-evidence-language-relation.md`](text-evidence-language-relation.md)、[`records/text-evidence.yaml`](records/text-evidence.yaml) | Probe miss＝inconclusive＋recovery（非無字幕）；Latin 黏字串有 boundary／recovery 或標記；raw 保留；雙語疊字有 text_region／group；cross-language／bilingual 不得當 same-language conflict／sanitization；spoken 非翻譯列／raw 拉丁硬字幕充中文 | text_evidence |
| 7 Assemble | [`assemble-and-qc.md`](assemble-and-qc.md)、[`source-publish-timebase.md`](source-publish-timebase.md) | Timeline IR 對 `shot_id`；selected 項可 trace；有 speed 時 EDR 含 `timeline_transform` | editor |
| 7a Transform | [`source-publish-timebase.md`](source-publish-timebase.md) | evidence 仍在 source；publish 僅投影；比較成片用 mapped time | editor |
| 7b Locale | [`captions-and-locales.md`](captions-and-locales.md)、[`speech-unit-and-timing.md`](speech-unit-and-timing.md)、[`subtitle-layout.md`](subtitle-layout.md) | 三閘分別有 decision；`timing_gate.timebase=publish`；content 源用 resolved spoken／subtitle 對；generated 口播 cue 有 speech timing evidence；layout.lines ≤ max_lines 且非把 max 當 target；timeline projection 無未解釋 omission | locale_author |
| 8 Publish | [`artifact-gates.md`](artifact-gates.md) | `fresh_reviewer` + blocking 空 | independent_verifier |
| 9 Outcome | [`publish-outcome.md`](publish-outcome.md) | 窗口欄位；不足樣 → `insufficient_sample` | outcome_author |

## 禁止

- 時長最近自動當選且不寫 `selection.policy`。
- 從 mp4 反推「當初為什麼這樣剪」。
- 無觀察填 PASS；合成單一「字幕 PASS」。
- 裸「AI 影片」當已註冊 route（尚未註冊）。
- 為 generated 口播猜 cue 秒數，或把 TTS／ffmpeg 寫成本 lifecycle 新 stage。
- 把 speech timing QC 與 layout QC 合成一次「字幕重做」。
- 把 `max_lines` 當「做成 N 行」的 target，或一行放得下仍拆行。
- 為了 fit 把 font 縮過 profile min（或滑到 absolute floor），或讓單句偏離 preferred／scene baseline 超過 max_delta。
- 為 fit 用字元索引截斷詞／專名，或讓同一 cue 上下行使用不同 typography。
- 用字數均分換行，或 layout 自創未評分的斷點。
- 為了靠近 prefer_at 而切開 lexical unit，或把標點當成強制切句。
- 讓 layout 改 Speech Unit 原文，或讓字幕切割與 TTS 各自決定停頓。
- 把 script heuristic 當成 NEVER-BREAK 最終答案，或為單集反例直接改 canonical BreakPolicy。
- 把「OCR 英文、ASR 中文」直接當 conflict／OCR 優先 spoken，或未過 Language Relation Gate 就跑和諧偵測。
- 讓 LLM 單獨斷言「英文 OCR＝翻譯」或「哪一行是哪種語言」或「怎麼斷英文詞」而不先有 mechanical `text_region`／`language`／boundary tags。
- 把雙語硬字幕黏成單一字串再進 Text Resolution；或把翻譯列當 spoken_text。
- 用 `recognition_language=ch` 對 Latin 觀測刪光空格，或用 derived 覆蓋／刪除 `raw_text`。
- 把預設 bottom／`no_subtitle_like` probe miss 當成「影片沒字幕」而 STOP／exclusion；缺 global discovery → targeted OCR recovery（見 [`text-evidence-ocr-discovery.md`](text-evidence-ocr-discovery.md)）。
- 讓 LLM 先猜 `subtitle_y` 並寫死唯一 crop，且無 mechanical discovery 兜底。
- 把 ASR 字面當 spoken meaning SoT，或固定 OCR>ASR 覆蓋；跨語言字幕缺 `semantic_candidate`／`resolution_reason` 就定案；領域術語硬字幕應走 `semantic_anchor`→`semantic_reconstruction`（observed／candidate／resolved 三層，見 [`text-evidence-multimodal-resolution.md`](text-evidence-multimodal-resolution.md)）。
- 只因多條 OCR 硬字幕 latched 同一長 ASR observation，就把 **text-unrelated** cue destructive merge 掉（見 captions timeline／omission）。
- 以 MP4 當唯一檢查面、略過 Timeline IR／coverage report。
- 對 sped／publish 媒體重跑 OCR／ASR／story／matching，或把 speed 當成新 evidence source（見 [`source-publish-timebase.md`](source-publish-timebase.md)）。
- 用同一 wall-clock 秒比較不同 rate 的兩條成片，而不換算 mapped source time。
