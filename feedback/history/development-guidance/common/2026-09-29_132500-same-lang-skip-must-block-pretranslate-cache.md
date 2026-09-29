> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-29 - Same-lang skip must block pretranslate cache (OCR Latin junk)

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。來自中文源→中文目標成片 dogfood：硬字幕 OCR 把短詞認成 Latin token 後被「譯」成臆造目標語。

#### One-line Summary

`skip_translate` / 源語≈目標語時必須短路整集預譯與 `translations/{lang}` 寫入；不可只靠 per-line `detect_text_lang`，否則 OCR 拉丁垃圾會被當成 en→zh 翻譯。

#### Human Explanation

中文成片做中文時，語系決策可標 `skip_translate`，但若 orchestrator 仍呼叫整集預譯，短純拉丁 OCR（如 2–12 字母 token）會被腳本偵測成 `en`，再送進翻譯通道寫入 `translations/zh`，燒錄成錯誤人名／部門名。正確行為：語系相同 → 不建／不寫同語 cache；拉丁 OCR 噪點保持原樣或丟棄，禁止臆造。

#### Trigger

- 目標語與片源語相同（尤其 zh→zh），卻出現 `translations/{target}` 新檔
- 字幕出現源片沒有的「譯名」，且 cue src 是短純拉丁串
- `dialogue_skip_translate=true` 只寫入 cache_extra，未閘住 `ensure_episode_translations`

#### Evidence

- Tool: product dogfood，中文-only reference_dub；OCR/dialogue_cues 短拉丁 token → 錯誤 zh dst
- Sanitized：same-lang skip 未閘預譯；per-line script detect 把 OCR junk 當 foreign
- Evidence path / episode titles：留在 `<PROJECT_ROOT>` analysis only

#### Generalized Lesson

1. **Decision-level skip 必須閘預譯**：`skip_translate` 為真時跳過 `ensure_episode_translations` 與翻譯寫 cache，改 identity pass-through。
2. **Per-line detect 不能當唯一閘**：整集以漢字為主時，短純拉丁 token 應視為 OCR junk，不得發明目標語專有名詞。
3. **同語 cache 空才是健康態**：源≈目標時 `translations/{lang}` 不應因噪點而誕生。
4. **Finality／Validation 可補一刀**：publish 前若 target=source 且 dst≠src 且 src 為短拉丁，標 needs_review／blocked。

#### Agent Action

改翻譯／配音管線時：先查 skip 是否短路整集預譯；dogfood 同語 run 後確認無妄生的 `translations/{lang}`；發現 Latin OCR→臆造譯名時優先修閘門而非加 prompt 例句。

#### Goal / Action / Validation

- Goal: 源語≈目標語（含 `skip_translate`）時不得因 OCR 拉丁噪點寫入 `translations/{lang}` 或臆造譯名。
- Action: decision-level skip 短路 `ensure_episode_translations`；同語走 identity pass-through；短純拉丁 token 當 OCR junk，不送外語翻譯通道。
- Validation or reference source: same-lang dogfood 後無妄生 cache；若 target=source 且 dst≠src 且 src 為短拉丁 → needs_review／blocked。

#### Applies When

- 字幕／配音／OCR hardsub 管線在源≈目標語時仍呼叫整集預譯
- per-line `detect_text_lang` 會把短純拉丁 OCR 標成外語的產品路徑

#### Does Not Apply When

- 真跨語翻譯（源≠目標）且拉丁串為合法外語內容
- 純人工譯稿、無預譯 cache／burn-in 的離線編輯

#### Validation

Same-lang run：無新 `translations/{target}`；短拉丁 OCR cue 保持原樣或丟棄，不出現源片沒有的「譯名」；`dialogue_skip_translate` 真正閘住預譯而非只寫 cache_extra。

#### Promotion Target

- `workflow/translation/`（Selection／cache／same-lang skip 邊界；待產品閘門穩定後再輕量回寫）
- Narrative Video Production burn-in／OCR 路徑（產品側；不在本 repo 強制改檔）

#### Promotion Record

本輪僅沉澱 feedback lesson + domain／common index；workflow／adapter 契約尚未改。

#### Required Linked Updates

- `feedback/history/development-guidance/common/README.md` index（已更新）
- `feedback/history/development-guidance/README.md` recent／count（已更新）
- workflow／adapter 正文：not applicable this commit（lesson 仍 candidate；產品閘門未回寫前不改契約）

#### Linked Plans

- Translation Decision：Selection／cache 邊界與 same-lang skip
- Narrative Video Production：burn-in 文案不得吞 OCR 臆造譯
