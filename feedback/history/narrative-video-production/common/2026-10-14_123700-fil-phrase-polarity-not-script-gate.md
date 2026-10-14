> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-14 - FIL filler／state polarity is not a script gate

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。來自 FIL／TL dogfood：來源「好不好／脏了／听话」
在 MT 產出 beauty／sick／recover 類 gloss，而 locale script admission 仍綠。

#### One-line Summary

目標語看起來像 Filipino／Tagalog 不能關閉短語極性；okay-filler≠beauty、dirty≠sick、
obey≠recover 要用 ZH-keyed mechanical repair 或 declared assertions。

#### Human Explanation

短 filler 與狀態句常被模型「說得通」但極性錯。正確鏈是：locale／Latin
gates → expression-type polarity repair（或 declared semantic check）→ Finality。
不要把系列句子塞進 Selection prompt；用可擴充的 ZH→bad-pattern→good 表。

#### Trigger

- fil／tl repair 只報告 script／Latin／empty flags
- 來源含詢問語氣 filler、髒／病、聽話／康復等易混對
- unit 或 spot REVIEW 看到 Maganda／Nagkasakit／gagaling 類固定混淆

#### Evidence

- Tool: product mechanical_fil_phrase_polarity + unit locks；plan evidence
  `plans/active/2026-09-22-1000-translation-decision-workflow/evidence/2026-10-14-fil-phrase-polarity-mechanical.md`
- Sanitized: three seed confusions rewritten for fil／tl；other locales skipped
- Product episode paths stay under `<PROJECT_ROOT>` analysis

#### Generalized Lesson

1. **Script pass ≠ polarity pass** for short discourse／state phrases.
2. Key repairs by **source cue text + bad destination patterns**, not series IDs.
3. Keep locale scope narrow（fil／tl）until other locales show the same confusion class.
4. Table hits do not alone authorize publish PASS；REVIEW／declared assertions remain.

#### Agent Action

對 fil／tl：在 validate／repair 路徑套用 expression-type polarity rewrite；表外混淆
標 REVIEW，勿為單一 episode 擴 prompt。

#### Goal / Action / Validation

- Goal: 阻止 FIL／TL 固定極性混淆被 script gate 放行。
- Action: ZH-keyed repair table + unit locks；dogfood 擴表。
- Validation: seed cues rewrite to expected polarity；non-fil locales unchanged.

#### Applies When

- 翻譯／配音 locale repair、FIL／TL semantic dogfood、polarity REVIEW

#### Does Not Apply When

- 非 fil／tl 目標語、或與 beauty／sick／recover 無關的開放語意爭議

#### Validation

- Unit locks：好不好／脏了／听话 seed cues rewrite to expected polarity
- Non-fil／tl locales unchanged on the same bad destination strings
- Linked from Translation Decision evidence README

#### Promotion Target

- `workflow/translation/` locale failure／dogfood notes（when promoting）
- Translation Decision failure-registry（FIL／TL manifestation）

#### Promotion Record

本輪僅 candidate lesson + plan evidence。

#### Required Linked Updates

- `feedback/history/narrative-video-production/common/README.md` 索引
- `plans/active/2026-09-22-1000-translation-decision-workflow/evidence/README.md`
