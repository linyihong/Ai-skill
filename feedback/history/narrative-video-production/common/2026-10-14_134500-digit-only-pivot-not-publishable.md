> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-14 - Digit-only pivot crumbs are not publishable ZH

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。來自 matrix dogfood：standalone OCR「7／八」
進入 ZH pivot；Arabic 已 I21-block，CN 數字仍被翻成 locale 數詞並造成 ASS 長度差。

#### One-line Summary

純數字／短中文數詞 OCR 不是對白 pivot；I21 與 fuse 都必須擋，除非 ASR 同字證明口播。

#### Human Explanation

黑板／角標／頁碼常被 OCR 成單獨「7」或「八」。若當 publishable ZH，翻譯會發明
數詞，locale ASS 與 pivot 對不齊。正確鏈：layout crumb 識別 → 收購丢棄或 I21
BLOCK → 僅當 ASR 文本就是該數詞才保留。

#### Trigger

- ZH cue 為 1–2 字純數詞或 ≤4 位阿拉伯／全形數字
- `asr_char_count=0` 卻仍 accepted
- locale 對數詞 cue 出現 invented number-word 或 empty unresolved skew

#### Evidence

- Tool: product `is_digit_only_layout_crumb` + fuse／hardsub gates；plan evidence
  `plans/active/2026-09-22-1000-translation-decision-workflow/evidence/2026-10-14-digit-only-pivot-not-publishable.md`
- Sanitized: CN numeral joined Arabic／punct I21 crumbs；fuse keeps only ASR text match
- Product episode paths stay under `<PROJECT_ROOT>` analysis

#### Generalized Lesson

1. **Digit-only ≠ dialogue** unless ASR carries the same numeral.
2. I21 incomplete_source must cover **CN numerals**, not only `[0-9]`.
3. Acquisition gates beat post-hoc translation BLOCK for publishable pivot hygiene.
4. Unit locks ≠ publish PASS；force re-acquire dirty caches.

#### Agent Action

對 standalone numeral OCR：fuse drop／I21 block；勿為系列特判。重跑 acquisition
後再比 ASS 長度。

#### Goal / Action / Validation

- Goal: 阻止 digit-only OCR 成為 publishable ZH pivot。
- Action: shared crumb helper + fuse／hardsub／I21。
- Validation: unit locks；「八／7」无 ASR 同字時不進 kept；「第一」不誤殺。

#### Applies When

- 對白證據合流、I21 Finality、matrix REVIEW of cue-length skew

#### Does Not Apply When

- 口播確實是報數且 ASR 同字；或句中含非數詞漢字（第一课／三十岁）

#### Validation

- Unit：I21 crumb + fuse drop／keep-when-ASR-matches
- 「第一／三十岁」不觸發 crumb
- Linked from Translation Decision evidence README

#### Promotion Target

- `workflow/translation/` failure／dogfood notes（when promoting）
- Translation Decision failure-registry（I21 / CORE-F21 manifestation）

#### Promotion Record

本輪僅 candidate lesson + plan evidence。

#### Required Linked Updates

- `feedback/history/narrative-video-production/common/README.md` 索引
- `plans/active/2026-09-22-1000-translation-decision-workflow/evidence/README.md`
