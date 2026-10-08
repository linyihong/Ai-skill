> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-08 - Nearby personal-name tokens do not close honorific identity

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。來自多語字幕 dogfood：稱謂表面旁出現
近形／近音人名與 MT 拉丁名，但無同窗共現把兩者等同；另有廠標水印殘片被
regex 當成「人名」。

#### One-line Summary

同系列裡出現的個人名、OCR↔ASR 近形衝突、或水印廠標，都不能用來關閉
未決的姓＋職稱身份；沒有等同證據就繼續 defer。

#### Human Explanation

找「某总」真名時，常見誤判是：同片出現另一個同姓個人名、或 ASR／OCR 差
一個形近字，就宣告身份已解。另一類是廠標／片頭浮水印被切成短字串，看起來
像人名。正確做法是要求**等同證據**（同句／同窗明確共指、或獨立覆核），否則
稱謂維持 unresolved，個人名另開 identity。

#### Trigger

- 想把同片另一個同姓個人名綁到未決稱謂表面
- OCR 與 ASR 對同一時間窗給出近形姓／名，想自動選一邊當 canonical
- regex／token 掃描把水印／廠標短串當成 cast 名

#### Evidence

- Tool: read-only overnight／bible JSON probe + relation window check（no realization）
- Sanitized: honorific surface stays deferred; nearby personal-name token treated as separate candidate; watermark-like short tokens excluded from identity bind
- Product paths stay under `<PROJECT_ROOT>` analysis only

#### Generalized Lesson

1. **Co-series ≠ co-identity**：同片出現不夠；要同窗／明確共指或獨立覆核。
2. **OCR↔ASR 近形衝突先保留兩側**，不可 silent winner。
3. **Watermark／brand residue ≠ person name**。
4. **Regex 候選要人工／獨立覆核裁假陽性**（稱謂後接動詞等）。
5. 與既有 lesson「honorific surface ≠ established name」並用，不取代。

#### Agent Action

探查稱謂真名時輸出：honorific counts、non-honorific candidates、同窗共現數、
水印嫌疑。共現為 0 則 `defer_no_bind`；有共現也只升到 needs_independent_review，
不自動 established。

#### Goal / Action / Validation

- Goal: 阻止假陽性人名關閉稱謂身份閘門。
- Action: relation／window probe 先於 identity_translations accept。
- Validation: 無共現時 established 不因探查增加；衝突 token 仍分列。

#### Applies When

- 稱謂真名追蹤、cast bind、identity_translations 覆核
- 從 overnight OCR／ASR／MT 挖人名候選

#### Does Not Apply When

- 已有明確共指或官方 credits 的真名綁定
- 純機械 script／empty-output 修復且不碰身份

#### Validation

- Probe 報告含 verdict=`defer_no_bind` 或 `needs_independent_review`
- 無共現時不寫 accepted personal name for the honorific surface

#### Promotion Target

- `workflow/translation/adapters/dogfood-acceptance.md`（升格時補 relation 條款）
- NVP companion `08-identity-precedes-naming`（升格時）

#### Promotion Record

本輪僅 candidate lesson + plan evidence；未改 workflow 正文。

#### Required Linked Updates

- common README 索引
- Translation plan evidence 短記 + evidence README（若有）
