> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-30 - Dialogue cue projection must not flatten subtitle_group

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。來自雙語硬字幕 dogfood：dialogue cue 已有 bilingual `subtitle_group`／`ocr_regions`，locale pack export 卻壓成 `spoken_timed{src,dst,start,end}`。

#### One-line Summary

`spoken_timed` 只投影 spoken／phrase；Dialogue Cue Group 的 `speech`＋`subtitle_group.regions[]`＋`alignment` 必須在 export／reload 後仍可讀，burn 優先同語 OCR region。

#### Human Explanation

上游已把中英疊字硬字幕拆成 region，並掛在 dialogue cue 上。若下游為了「燒字幕方便」只保留一句 text，異語 region 與 speech alignment 會消失，目標語只能從 spoken 再翻譯，等於在 projection 層重現 `bilingual_collapsed_to_single_string`。正確做法：entries 可薄（spoken_timed），cues 必須保留結構；Layer-0 OCR schema 不需為此重做。

#### Trigger

- dialogue cue 有 `subtitle_group.type=bilingual`／多語 `ocr_regions`
- `translations/{lang}/epXX.json` 的 cues 只剩 `{start,end,text}`
- 目標語 burn 文案明顯不如已存在的同語 hardsub region

#### Evidence

- Tool: product dogfood；dialogue_cues 含 regions，locale pack 壓扁
- Sanitized：projection flatten spoken_timed vs retained subtitle_group
- Episode titles／paths：`<PROJECT_ROOT>` only

#### Generalized Lesson

1. **Layer-0 與 projection 分開**：Layer-0 visual_text／OCR schema 正確 ≠ export 正確。
2. **`spoken_timed` ≠ subtitle SoT**：只承載口播語 burn／phrase。
3. **Cue group 必帶**：`speech`＋`subtitle_group.regions[]`＋`alignment`。
4. **Burn 優先同語 region**：缺口才 MT；不得為簡化丟 region。

#### Agent Action

改 locale pack／timed export／burn 時：確認 reload 後 cues 仍有 ≥2 regions（雙語窗）；目標語 text_source 可追溯 ocr_region 或 mt；不重寫 Layer-0 OCR schema 來修 projection。

#### Validation

- Reload locale pack cues after export: bilingual windows still have ≥2 `ocr_regions`.
- Target-lang timed cue `text_source` can be `ocr_region` when hardsub region exists.
- Layer-0 OCR files unchanged by the projection-only patch.

#### Goal / Action / Validation

- Goal: 雙語 hardsub 的 region／alignment 在 locale pack 與 burn 路徑可追溯。
- Action: export／load 保留 Cue Group 欄位；target burn 優先同語 region。
- Validation or reference source: 重跑後 zh／target cues 含 `ocr_regions`＋`subtitle_group`；EN（或他語）burn 可引用 region 而非僅 spoken→MT。

#### Applies When

- 硬字幕多語疊字且上游已產 `subtitle_group`／regions
- locale pack／burn adapter 從 dialogue cues 投影

#### Does Not Apply When

- 純單語字幕且無 region 結構
- 尚未做 OCR region split（仍屬上游 contract gap）

#### Promotion Target

- Workflow：`workflow/narrative-video-production/text-evidence-regions.md`
- Plan：`plans/active/2026-09-16-1649-narrative-video-production-workflow/41-bilingual-ocr-regions-and-subtitle-groups.md`

#### Required Linked Updates

- Plan evidence：`…/evidence/2026-09-30-dialogue-cue-projection-retains-subtitle-group.md`
- `feedback/history/development-guidance/common/README.md` 索引列
