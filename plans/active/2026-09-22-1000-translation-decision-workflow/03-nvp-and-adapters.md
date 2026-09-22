# NVP integration & adapters

Companion to [`_plan.md`](_plan.md)。

## 與 narrative-video-production 的關係

| NVP 已有 | Translation workflow 角色 |
| --- | --- |
| `dialogue.semantic_context` | **Translation Context Contract** 上游 evidence（speaker、addressee、intent、subject_refs…） |
| `caption-locale-pack` `content_gate` | Subtitle adapter **輸出**接點（語意、專名、source residue） |
| `timing_gate` / `layout_gate` | **不在** translation-core；subtitle adapter 只 **reference** NVP |
| Invariant 7 三閘分開 | Translation 只擁 **content correctness** 決策鏈；禁止合成「字幕 PASS」 |

`dialogue-semantic-context.yaml` 已列 `subtitle_or_translation_context` 為 use case；Phase 2 明文化欄位映射。

## Translation Context Contract（Phase 1，P0）

`source_locale`／`target_locale` 由 **NVP locale pack／job** 傳入，屬 **authoritative Constraint**（I1：Locale Resolution ≠ Language Detection）。Actor 禁止只吃 `{ src, dst }`；禁止每段猜 target。**dst 不得用來判斷哪些 source 該 merge**；canonical spoken 先於 translation（NVP [`25-text-group-preserve-variants.md`](../2026-09-16-1649-narrative-video-production-workflow/25-text-group-preserve-variants.md)）。

稱謂案例：[`04`](04-dogfood-case-address-title-id-ID.md)。Freeze：[`06`](06-phase-0-freeze-invariants.md)。

從 episode evidence 補足缺主詞等問題後再 Selection。缺必要 context 或 blocking validation 未 resolved → 不得 `accepted`（I9）。

建議必填（dialogue translation）：

- `source_locale` / `target_locale`（來自 locale pack）
- `semantic_context` 或等價 structured context
- `text_origin`（script／asr／ocr／human／translated）

## Subtitle adapter（Phase 2 — landed）

檔案：[`workflow/translation/adapters/subtitle.yaml`](../../workflow/translation/adapters/subtitle.yaml)

職責：

1. 輸入：segment + TDR／decision state + `semantic_context`（若有）
2. 輸出：cue 文案 + TDR id → NVP `content_gate`／`translation_qc`
3. **不**裁決 CPS、安全區、burn_mode

Linked updates：

- [`captions-and-locales.md`](../../workflow/narrative-video-production/captions-and-locales.md)
- [`caption-locale-pack.yaml`](../../workflow/narrative-video-production/records/caption-locale-pack.yaml)
- [`dialogue-semantic-context.yaml`](../../workflow/narrative-video-production/records/dialogue-semantic-context.yaml) — upstream for TranslationContext／Selection
- NVP companion [`01-captions-and-locales.md`](../2026-09-16-1649-narrative-video-production-workflow/01-captions-and-locales.md)

Static gate：[`08-static-walkthrough-pass.md`](08-static-walkthrough-pass.md)。

## 其他 adapters（Phase 4）

| adapter | 用途 |
| --- | --- |
| document | 章節／段落；無 cue timing |
| ui | 字串長度、placeholder、RTL（機械欄位擴 validation-rules） |
| dubbing | lip-sync／duration 約束（後續；可能接 NVP 配音交付） |

## 與 ERA delegation loop（advisory）

高風險或 `needs_review` 堆積時，可走 orchestrator → executor（產 candidates）→ verifier（L1–L3 對 brief acceptance）→ 仲裁；brief 的 acceptance 應對 TDR 欄位與 validation 維度，非「讀起來像中文」。

本 plan **不**強制每 segment 走三角色；dogfood 時可記錄何時需要 independent verification（對齊 NVP invariant 8：publish-ready 需 fresh verification；translation 的 `accepted` 可對齊 content 層的 completion 定義，Phase 1 在 finality contract 寫清）。
