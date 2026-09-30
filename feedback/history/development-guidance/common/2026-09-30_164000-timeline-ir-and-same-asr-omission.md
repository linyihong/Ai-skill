> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-30 - Timeline IR + forbid unrelated same-ASR caption collapse

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。來自 Phase 3 雙語硬字幕 dogfood：L0 OCR 有短 cue，成片／dialogue cues 卻缺句。

#### One-line Summary

「Evidence 有、成片沒有」是 projection／trace 問題；先查 Timeline IR＋coverage，禁止只因共用同一長 ASR latch 就把 text-unrelated OCR cue destructive merge。

#### Human Explanation

硬字幕每一屏常是獨立對白。長 ASR（或錯聽跨多秒）可能時間上同時蓋住多條 OCR。若 resolve 後「同 ASR → 合併成一句並加寬時窗、只留分數最高者」，就會把中間那句（例如短中文硬字幕）無聲刪掉。這不是 content 譯錯，是 artifact omission。正確做法：Timeline IR 一等交付；coverage 標 `suspicious_omission`；同 ASR 合併只允許 text 為 duplicate／truncated／variant。

#### Trigger

- OCR／visual text 有句，dialogue cues／ASS／成片沒有
- 多條 cue 共用同一 `asr_observation` 且時間相鄰後只剩一句
- 只能靠看 MP4 才發現缺字幕

#### Evidence

- Tool: product dogfood；OCR 保留、fuse／TextAlignment 保留、resolve 後 same-ASR collapse 丟句
- Sanitized：`unrelated_forced_merge` vs allowed same-utterance fragment merge
- Paths／titles：`<PROJECT_ROOT>` only

#### Generalized Lesson

1. **MP4 不是唯一檢查面**：EDR → Timeline IR → Mechanical QC → Render。
2. **selected 必須可 trace**：無 downstream → fail／`suspicious_omission`，不可 silent。
3. **subtitle_group 整組投影**：雙語 region 不得拆丟。
4. **same-ASR merge 有文本門檻**：僅 `duplicate`／`truncated_variant_of`／`variant_of`；`unrelated` 禁止。
5. **loser 進 candidates／coverage**：即使合併同一口播碎片，也要留下 not_used reason。

#### Agent Action

擴 captions／artifact-gates；修 collapse 門檻；bump cues 並重算；出 coverage／timeline 中間產物。

#### Validation

- Fixture：兩條 unrelated OCR + 同一長 ASR → 仍為兩 cue。
- Fixture：Tonight／Last Lot 同 ASR 且語意碎片 → 仍可合併勝出者。
- 重算後目標短句出現在 cues／ASS；coverage 無未解釋 omission。

#### Goal / Action / Validation

- Goal: 高品質硬字幕 evidence 不得因同 ASR latch 無聲消失；缺句可機械定位層級。
- Action: Timeline／coverage 契約 + collapse 文本門檻 + product regenerate。
- Validation or reference source: fixture + dogfood cues／ASS 對照。

#### Applies When

- OCR／硬字幕 → cues → burn；雙語 subtitle_group；Phase 3 dogfood

#### Does Not Apply When

- 純音訊無字幕；明確只交付 ASR spoken 且無硬字幕责任

#### Promotion Target

- `workflow/narrative-video-production/captions-and-locales.md`
- `workflow/narrative-video-production/artifact-gates.md`
- `workflow/narrative-video-production/assemble-and-qc.md`
- Plan companion `01-captions-and-locales.md`

#### Required Linked Updates

- Evidence：`…/evidence/2026-09-30-ocr-caption-omission-same-asr-collapse.md`
- `feedback/history/development-guidance/common/README.md` 索引
