> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-08 - Honorific address surface is not an established person name

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。來自多語字幕 dogfood：姓＋職稱呼語有
OCR／ASR 證據，但無真名；獨立覆核將稱謂譯名候選全部 defer，個人名譯名才
accept。

#### One-line Summary

姓＋職稱／呼語可以當 address surface 證據，但不能在身份未決時寫進
established／accepted 譯名；人名、職稱、普通名詞要分開。

#### Human Explanation

 ass常見捷徑是：畫面上反覆出現「某总／某老师」，就當成角色真名並寫進
各語固定譯詞。正確鏈是 observed surface →（可選）identity cluster →
canonical name（需 evidence_refs）→ 再做 locale 譯名的**獨立** accept。
稱謂層只證明「有人被這樣叫」，不證明 canonical identity。

#### Trigger

- 字幕／對白翻譯要把職稱呼語做成 glossary 或 established name
- cast／role cache 以出場權重 `confirmed` 宣稱身份已解析
- 想把 Boss／Mr.＋姓 的 MT 輸出直接當目標語驗收名稱

#### Evidence

- Tool: per-series read-only reference consumer + independent name review on frozen local outputs
- Sanitized: honorific surname+title surfaces deferred; personal-name locales accepted only with independent reviewer binding; conflicting MT and Latin residue rejected
- Evidence path: Translation plan `evidence/2026-10-08-independent-name-review-binding.md`；NVP plan `evidence/2026-10-08-identity-before-honorific-established-name.md`

#### Generalized Lesson

1. **Surface ≠ identity**：address／honorific 可留 unresolved。
2. **Identity ≠ target acceptance**：來源 resolved 不自動關閉 locale Finality。
3. **Independent binding**：accepted + independent reviewer_role + evidence_refs。
4. **衝突全留**：禁止 first-match；拉丁殘留≠目標語 established。
5. **勿合成單閘**：Reference 進展 ≠ timing／layout／burn PASS。

#### Agent Action

遇到「把某总譯名寫進詞表」：先查是否有非稱謂真名證據；沒有則 defer。
只把獨立覆核 accept 的個人名寫入 established；衝突與殘留標 reject／keep
candidates。不重跑相同 MT 碰運氣。

#### Goal / Action / Validation

- Goal: 避免未決稱謂污染固定譯名與角色表。
- Action: 分 source identity／target-name review／content Finality 三欄記錄。
- Validation: 稱謂案例 established 不增加；個人名 accept 數＝獨立 accept 數。

#### Applies When

- 多語字幕／配音的人物與稱謂 realization
- per-series cast／identity review export
- Translation Decision dogfood Reference／Acceptance

#### Does Not Apply When

- 純排版／burn 幾何
- 已有獨立真名證據且 review binding 完整的角色

#### Validation

- Adapter／evidence 明文禁止稱謂在 identity unresolved 時 accepted
- 凍結重放：honorific defer；personal-name accept 可計數核對

#### Promotion Target

- `workflow/translation/adapters/dogfood-acceptance.md`（已同步補強）
- 候選：`workflow/narrative-video-production/source-bible.md`（升格時）
- NVP companions `05-series-cast-canonicalization`／`08-identity-precedes-naming`

#### Promotion Record

本輪已寫入 dogfood-acceptance 稱謂／MT 邊界；尚未升格 source-bible schema。

#### Required Linked Updates

- Translation／NVP plan evidence 索引
- common README 索引
