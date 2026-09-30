> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-01 - OCR join must use token seam not cumulative script

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。來自雙語硬字幕 dogfood：`parts[]` 已切開，derived `text` 因 cumulative mixed script 把後續 Latin 詞黏回。

#### One-line Summary

`parts[]` 是 L0 evidence、`text` 是 derived；join 只看相鄰 token 接縫（Latin|Latin 必空格），禁止用累積 mixed／CJK 否決空格；derived 不得覆蓋 parts。

#### Human Explanation

OCR 已把英文拆成 I'm／from／a／pet／store，但 join 看「整段已是 mixed」就不插空，得到 I'm fromapetstore。這是 projection 損壞 evidence，不是「不會切詞」。單框仍黏的 appointmentfortoday 才走 geometry／lexical；與 join 分開。

#### Trigger

- `parts` 有多個 Latin token，整段 `text` 缺空格（fromapetstore／You'refinally／Needamaster）
- CJK 前綴 + 多個 Latin parts 後段被黏
- 用黏掉的 text 覆蓋已切開的 parts

#### Evidence

- Tool: product dogfood OCR cache；parts 正确、join 坏
- Sanitized：token-seam join vs cumulative script deny
- Paths／titles：`<PROJECT_ROOT>` only

#### Generalized Lesson

1. **parts 永不因 join 遺失**；text = rebuildable derived（`method: script_aware_join`）。
2. **接縫判斷**：Latin|Latin→空格；勿用 cumulative mixed／CJK 否決。
3. **Intra-box 黏串 ≠ join bug**：geometry→lexical→必要時 LLM；分開修。
4. **雙語先 region**：zh／en 各自 normalize，再 timeline。
5. **adapter_only**：不開新 Phase／runtime／agent。

#### Agent Action

修 join 門檻；回寫 derived text；固定 CJK+Latin+Latin… regression；重算 cues／timeline。

#### Validation

- Fixture：I'm/from/a/pet/store → I'm from a pet store
- Door/to/door/cat…、You're/finally/here、I/feel/so/bad…、Need/a/master、Hurry/up/and…
- parts 仍在；derived 可重跑

#### Goal / Action / Validation

- Goal: 所有片共用機械契約——token-seam join + parts 保留。
- Action: workflow invariant + product `_needs_space_between`／rebuild。
- Validation: regression + dogfood OCR／cues 對照。

#### Applies When

- 多 part OCR（尤其中英硬字幕）合成 display／dialogue text

#### Does Not Apply When

- 純 CJK 無 Latin parts；或已正確 spaced 的多 box Latin

#### Promotion Target

- `workflow/narrative-video-production/text-evidence-ocr-boundary.md`
- `workflow/narrative-video-production/records/text-evidence.yaml`
- Plan companion `42-ocr-boundary-and-script-aware-normalization.md`

#### Required Linked Updates

- Evidence：`…/evidence/2026-10-01-ocr-join-token-seam-not-cumulative-script.md`
- `feedback/history/development-guidance/common/README.md` 索引
