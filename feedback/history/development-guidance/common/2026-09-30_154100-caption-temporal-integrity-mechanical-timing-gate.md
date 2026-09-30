> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-30 - Caption temporal integrity is a mechanical timing gate

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。來自雙語硬字幕成片 dogfood：可見兩套台詞時間互壓；Caption Pack 無 overlap，burn ASS same-track 有小數秒互壓。

#### One-line Summary

字幕互壓是 Rendering／Temporal Integrity，不是 Evidence Resolution；`timing_gate` 必須機械檢查 same-track overlap（未明示則禁），雙語 same-group 允許，並在 burn 前／artifact 驗 EDR↔rendered。

#### Human Explanation

內容對錯走 Text Resolution；時間軸上兩條字幕是否撞車，可用 `next.start < current.end` 無歧義判定。LLM 不該看片判 overlap。雙語同組 region 同窗是合法的，不可當成 same-track fail。若 Pack 乾淨但成片互壓，屬 render adapter（倍速／ASS merge／未 scrub 硬字幕）缺陷，不是新 Phase。

#### Trigger

- 成片／ASS 上相鄰 cue 時間重疊
- 源片硬字幕仍可見又燒了新 caption（兩套台詞）
- Caption Pack 乾淨但 burn 後互壓

#### Evidence

- Tool: product dogfood；timed cues overlap=0，ASS Default track 有鄰接 overlap
- Sanitized：same-track cue_overlap vs bilingual_same_group allowed
- Paths／titles：`<PROJECT_ROOT>` only

#### Generalized Lesson

1. **分層**：content ≠ timing ≠ layout ≠ artifact。
2. **same_track 預設 forbidden**：`next.start < current.end` → `cue_overlap` fail。
3. **bilingual_same_group allowed**：同 `subtitle_group` 多 locale region 不算 fail。
4. **group 之間互壓**：`subtitle_group_overlap` fail。
5. **Pack vs render**：Pack 互壓＝producer／timing fail；Pack 乾淨成片壞＝adapter defect。
6. **publish-ready** 需 content∧timing∧layout∧artifact 各自 PASS。

#### Agent Action

擴 `timing_gate`／burn 前機械 validator；發現 overlap 先分類再修 adapter 或擋 publish；不要開 LLM QC stage。

#### Validation

- Validator 對 known overlapping fixture 回 `cue_overlap`。
- Bilingual same-group fixture 回 PASS。
- Burn 前跑 gate；嚴重 overlap 不得 silent pass。

#### Goal / Action / Validation

- Goal: 未明示的 same-track caption overlap 在 render 前被機械擋住或修復並留證。
- Action: 擴 timing_gate temporal integrity；wire burn-precheck；區分 pack vs adapter。
- Validation or reference source: fixture + dogfood ASS／pack 對照。

#### Applies When

- Caption／硬字幕 burn、locale pack QC、publish-ready
- 多 cue 同視覺字幕層

#### Does Not Apply When

- 純音訊、無字幕交付
- Explicit transition／karaoke policy 已標允許

#### Promotion Target

- `workflow/narrative-video-production/captions-and-locales.md`
- `workflow/narrative-video-production/artifact-gates.md`
- Plan companion `01-captions-and-locales.md`

#### Required Linked Updates

- Evidence：`…/evidence/2026-09-30-caption-temporal-integrity-overlap.md`
- `feedback/history/development-guidance/common/README.md` 索引
