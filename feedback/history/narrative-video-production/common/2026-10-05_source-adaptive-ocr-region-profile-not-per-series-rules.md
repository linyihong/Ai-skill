> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-05 - Source-adaptive OCR region profile, not per-series hand-tuning

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。來自硬字幕＋中置 sticky 免責水印黏進對白的 dogfood：字幕未歸零，但 accepted 品質被污染；手調單片規則無法規模化。

#### One-line Summary

用機械 probe 從影片自學 region behavior／source_ocr_profile（含 sticky），再 targeted OCR＋projection；禁止把單片修正寫成全域 watermark 座標或字典。

#### Human Explanation

固定 bottom OCR 或「這部片調一次水印」會陷入：修 A 怕壞 B、黏字串字典刪不穩、probe inconclusive 被誤讀成沒字幕。正確鏈是 boxes→region clustering→behavior（persistence／repetition／stability）→versioned source profile→scan policy；LLM 只校準 ambiguous regions。Sticky layer 應先空間分離，文字層 projection 是 fallback。

#### Trigger

- 討論為單部片調 OCR band／watermark_y／刪字字典
- 中置或非常規位置 sticky 水印黏進對白
- 問「這次調整會不會搞壞上一部」卻沒有 profile impact／fixture regression

#### Evidence

- Tool: product OCR／dialogue probe dogfood on hardsub short-drama with sticky disclaimer overlays
- Sanitized: center sticky high-repetition overlay glued into dialogue lines; windows still yield multiple cues; probe inconclusive ≠ no subtitle
- Evidence path: plan evidence `2026-10-05-source-adaptive-ocr-region-profile`

#### Generalized Lesson

1. **Source profile is observation**：`source_ocr_profile` 綁 source／version，不是 canonical global rule。
2. **Region before text**：空間／時序分層優先於 OCR 後字典刪字。
3. **Sticky = temporal behavior**：長時＋穩 bbox＋高文本重複 → watermark_candidate／sticky；無需 LLM。
4. **Probe wide, OCR narrow**：probe 學 layout；正式 OCR 跟 profile。
5. **LLM once per source**：只標 ambiguous roles；其後機械。
6. **Impact + tiny fixtures**：profile v1→v2 必報 candidates／accepted／uncertain／rejected 前後差；難例 clip corpus 防回歸。
7. **不開新 Phase**：接 OCR discovery／regions／sticky projection／acquisition-loop。

#### Agent Action

遇到「這部片水印怎麼濾」：先寫／更新 source-adaptive profile 契約與 evidence；產品用本集自學 sticky spans／region roles，禁止全域單片字典；驗證用污染率＋小 fixture，不靠感覺測上一部。

#### Goal / Action / Validation

- Goal: 新 sticky 片源無需手調全域規則即可降低黏串污染，且可審計 impact。
- Action: 回寫 NVP ocr-discovery／evidence／lesson；產品 episode-local sticky mining＋projection。
- Validation: 窗內 cues 非空；黏串污染下降；無新增全域單片名 watermark 字串。

#### Applies When

- 硬字幕 OCR／layout discovery／sticky projection
- 多區域文字（字幕＋水印＋UI／聊天）並存

#### Does Not Apply When

- 純 burn／排版 typography
- 真的 acquisition miss（無 OCR box、無片源）

#### Validation

- 契約寫明 `source_ocr_profile` 為 observation、可版本化；禁止單片座標／字典升全域
- ocr-discovery 含 region behavior／sticky／profile impact 欄位
- 產品：episode-local sticky mining＋projection；單元測試無新增單片名全域 watermark 字串
- 窗內 cues 非空時不以 probe inconclusive 當「沒字幕」

#### Promotion Target

- `workflow/narrative-video-production/text-evidence-ocr-discovery.md`
- 交叉：[`text-evidence-regions.md`](../../../../workflow/narrative-video-production/text-evidence-regions.md)、sticky-watermark projection evidence

#### Promotion Record

尚未 full promotion；本輪已同步強化 ocr-discovery＋plan evidence。

#### Required Linked Updates

- 更新 common README 索引
- 強化 ocr-discovery：region discovery → behavior → source profile → scan policy；sticky 空間優先；profile impact
- 新增 plan evidence `2026-10-05-source-adaptive-ocr-region-profile` 並掛 evidence README
- 產品：`mine_sticky_spans`／episode-local peel（seed 可留，禁止單片字典全域化）
- 專案幀／片名細節留 `<PROJECT_ROOT>`；本 lesson 不含本機路徑／host

#### Links

- Evidence: [`2026-10-05-source-adaptive-ocr-region-profile`](../../../../plans/active/2026-09-16-1649-narrative-video-production-workflow/evidence/2026-10-05-source-adaptive-ocr-region-profile.md)
- Workflow: [`text-evidence-ocr-discovery.md`](../../../../workflow/narrative-video-production/text-evidence-ocr-discovery.md)、[`text-evidence-regions.md`](../../../../workflow/narrative-video-production/text-evidence-regions.md)
- Related: sticky projection evidence、ocr-layout-typography lesson
