> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-08 - MT name forms are candidates until independent accept

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。來自 locale pack／overnight MT 人名形式
與獨立覆核對照：主譯可 accept，衝突譯與拉丁殘留 reject。

#### One-line Summary

自動翻譯產出的人名只是 candidate space；唯有獨立覆核 accept 且綁定
identity／locale／evidence 才能進 established_names。

#### Human Explanation

同一來源人名在目標語常出現多種 MT 寫法，或在非拉丁 locale 殘留拉丁拼寫。
若系統 first-match 或因英文一致而接受，會把錯誤或半成品寫進角色譯名。
正確做法是保留全部衝突候選，由 independent reviewer 逐筆 accept／reject／
defer，並保留 provenance。

#### Trigger

- 要把 phrase cache／locale pack 人名寫進系列詞表
- 目標語輸出混有拉丁拼寫卻被標 pass
- 同一 identity 出現兩個以上譯名時自動選一筆

#### Evidence

- Tool: adapter-owned per-series identity_translations review export + independent decisions on frozen replay
- Sanitized: one locale personal-name form accepted; conflicting form and Latin residue rejected; honorific forms deferred
- Evidence path: `plans/.../translation-decision-workflow/evidence/2026-10-08-independent-name-review-binding.md`

#### Generalized Lesson

1. MT／overnight pack → `needs_review` 候選，預設不是 accepted。
2. 衝突候選全留；reject 必須留原因，不要刪除歷史。
3. 非目標文字系統殘留≠該 locale established name。
4. Selection actor／extract pipeline 的 reviewer_role 不得自關 Acceptance。

#### Agent Action

匯出候選時標明來源為 machine／selection；請獨立覆核後再升格。不把 en
pivot 自動投影到其他 locale。

#### Goal / Action / Validation

- Goal: established_names 只含獨立接受項。
- Action: review export 分 status；consumer 只讀 accepted+independent。
- Validation: accept 後計數與覆核清單一致；reject／defer 不進入 established。

#### Applies When

- Translation Decision Reference／Acceptance dogfood
- 多語字幕人物名 realization

#### Does Not Apply When

- 已有 human glossary 且契約允許作為 seed（仍需 locale binding）

#### Validation

- Consumer 忽略非 independent／非 accepted 項
- 覆核後 established 計數可機械核對

#### Promotion Target

- `workflow/translation/adapters/dogfood-acceptance.md`
- 可選 failure-registry manifestation（locale residue／name conflict）

#### Promotion Record

Adapter 邊界已補強；locale failure manifestation 尚未新增。

#### Required Linked Updates

- Translation evidence README 索引
- 交叉 lesson：honorific-surface-not-established-name
