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
7. Assemble vs EDR       → 時間線對 shot_id／selected_clip_id
7b. Locale packs         → content／timing／layout 三閘分開
8. Publish QC            → 平台規格；publish-ready 需 fresh verification
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
| 7 Assemble | [`assemble-and-qc.md`](assemble-and-qc.md) | 成片軸對 `shot_id` | editor |
| 7b Locale | [`captions-and-locales.md`](captions-and-locales.md) | 三閘分別有 decision | locale_author |
| 8 Publish | [`artifact-gates.md`](artifact-gates.md) | `fresh_reviewer` + blocking 空 | independent_verifier |
| 9 Outcome | [`publish-outcome.md`](publish-outcome.md) | 窗口欄位；不足樣 → `insufficient_sample` | outcome_author |

## 禁止

- 時長最近自動當選且不寫 `selection.policy`。
- 從 mp4 反推「當初為什麼這樣剪」。
- 無觀察填 PASS；合成單一「字幕 PASS」。
- 裸「AI 影片」當已註冊 route（尚未註冊）。
