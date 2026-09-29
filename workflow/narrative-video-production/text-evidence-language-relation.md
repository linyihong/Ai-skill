# Text evidence — language, role, and relation gate

何時讀：源片／參考成片有 OCR＋ASR（或硬字幕＋口播）要進 Text Resolution、和諧偵測、或 locale content 源之前。
欄位契約：[`records/text-evidence.yaml`](records/text-evidence.yaml)。
Plan companion：[`40-language-role-before-text-resolution.md`](../../plans/active/2026-09-16-1649-narrative-video-production-workflow/40-language-role-before-text-resolution.md)。

> **執行契約，不重新解釋契約。** 不寫模型名／ffmpeg。LLM 只做 alignment／ranking，不是 language constraint 的來源。

## 在 lifecycle 的位置

插在「可觀測 OCR／ASR」與「Text Resolution／locale content 源」之間（acquisition／assemble 準備證據時即可跑；**不是**新的 publish stage）：

```text
OCR → text_region / subtitle_group   (見 text-evidence-regions.md)
  → Language & Text-Role Detection
  → Evidence Alignment (per region ↔ ASR)
  → Language Relation Gate
  → Text Resolution
  → (spoken／subtitle 對) → locale content／narrative
```

## 每個 evidence 必填

| 物件 | 必填 |
| --- | --- |
| `observed_text`（OCR） | `text`、`language`、`text_role`（candidate）；可選 `locale` |
| `spoken_text`（ASR） | `text`、`language`；可選 `locale`、`speaker_id` |

`source_locale`（片／口播）≠ `display_locale`（目標字幕 pack）。
`spoken_text` ≠ `subtitle_text`：兩者可同時正確。

## Language Relation Gate（機械優先）

```text
language(ocr) == language(asr)?
  ├── YES → same-language path
  │         text_relation: same_language_match | same_language_variant
  │         | sanitization_candidate | unresolved
  └── NO  → cross-language path
            candidate_relation = cross_language
            LLM 只判 semantic alignment support?
              ├── YES → text_relation: cross_language_translation
              │         （不進 sanitization）
              └── NO  → cross_language_unresolved
```

禁止：

- 把跨語言對齊標成 `OCR_ASR_CONFLICT` 或直接 OCR 勝出當 spoken canonical；
- 未過本閘就跑 sanitization／和諧假設；
- 讓 LLM 單獨斷言「英文 OCR＝翻譯」而不先有 mechanical language tags。

## `text_role`（OCR）

candidate only：`translated_subtitle`／`hardcoded_subtitle`／`dialogue_caption`／`signage`／`UI_text`／`watermark`／`title`。
signage／UI／watermark **不得**僅因 ASR 同義就當 subtitle translation。watermark 排除仍是 projection，不刪 raw（見 visual-text／role projection companions）。

## 推進條件

| 條件 | 失敗 |
| --- | --- |
| 每條對齊窗有 ocr／asr 的 `language`（或明示 unknown） | 不得進 sanitization／不得宣告 spoken canonical |
| cross-language 時有 `text_relation` ∈ {cross_language_translation, unresolved, …} | 不得用 same-language conflict 規則 |
| locale content 源使用 resolved spoken／subtitle **對** | 直接餵 raw OCR 拉丁硬字幕當中文口播源 = 閘失敗 |

## 與三閘／Translation Decision

- Locale **content** gate 消費本檔產出的 resolved 對；timing／layout 不變。
- Translation Decision 的「怎麼翻」在 language／role 清楚之後才談。
