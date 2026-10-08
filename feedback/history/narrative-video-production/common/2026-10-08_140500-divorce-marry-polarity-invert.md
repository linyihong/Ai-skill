> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-08 - Divorce/marry polarity can invert under locale MT

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。來自阿拉伯語 dogfood：來源「離婚」在
local constrained realization 兩次產出「結婚」詞根，直到正面 طلاق 範例＋明確
禁止 تزوج 才命中離婚義。

#### One-line Summary

機械通過（目標文字、無 Latin）不表示語意通過；離婚／結婚這類反義對必須用
declared assertion 或詞根檢查，不能只看 script gate。

#### Human Explanation

字幕修復常把「有阿拉伯字、沒有英文」當成修好。但模型可能把離婚翻成結婚
（或反向）。正確鏈是：mechanical admission → 關鍵謂語／極性檢查（或
declared semantic assertions）→ 再 Finality。約束 prompt 要同時給正例詞根
與反例禁詞，並保留失敗輸出作 repair_of。

#### Trigger

- locale repair 只報告 script／empty／Latin flags
- 來源含離婚／結婚／離職／入職等易反義對
- constrained local MT 「看起來通順」就準備 accept

#### Evidence

- Tool: bounded local-only Arabic realization with constraint escalation
- Sanitized: first outputs used marry root; later output hit divorce root after
  positive طلاق example + explicit تزوج ban; script flags alone were insufficient
- Product episode paths stay under `<PROJECT_ROOT>` analysis

#### Generalized Lesson

1. **Script pass ≠ meaning pass** for antonym-prone predicates.
2. Track **marry_risk / divorce_hit**（或 generic polarity flags）as repair metrics.
3. Constraint escalation：禁反義詞根 + 正例詞根 + 引用上一輪錯誤輸出。
4. Content Finality 仍要 declared assertions／independent review，不靠 prompt 運氣。

#### Agent Action

對離婚／結婚類 cue：在 flags 裡記 polarity；marry_risk 且無 divorce_hit 不得
標 content repair_ready。必要時第二輪 forced repair，保留失敗輸出。

#### Goal / Action / Validation

- Goal: 阻止反義謂語被 script gate 放行。
- Action: 加 polarity flags；失敗則 escalate constraint。
- Validation: divorce cue 最終輸出 divorce_hit=true 且 marry_risk=false。

#### Applies When

- 翻譯／配音 locale repair、semantic dogfood、acceptance binding

#### Does Not Apply When

- 純人名音譯且無謂語
- 已有 independent semantic assertions 覆蓋該 action

#### Validation

- 報告同時含 mechanical flags 與 polarity flags
- 反義失敗輪次被保留且不寫成 accepted

#### Promotion Target

- `workflow/translation/adapters/dogfood-acceptance.md`（升格時補 polarity）
- Translation Decision failure-registry（locale manifestation）

#### Promotion Record

本輪僅 candidate lesson + plan evidence。

#### Required Linked Updates

- common README 索引
- Translation plan evidence 短記 + evidence README
