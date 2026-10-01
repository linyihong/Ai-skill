> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-01 - OCR Latin boundary: geometry before closed-class

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。來自雙語硬字幕 dogfood：join 修好後仍有單 part 黏串；整句 bbox 不足以拆詞。

#### One-line Summary

Latin boundary recovery 必須分層：OCR boxes → intra-box geometry → lexical candidates（closed-class 只是候選產生器）→ ASR／context → resolver；禁止凡 Latin 都拆，禁止跳過幾何把詞表當第一刀真理。

#### Human Explanation

`appointmentfortoday`／`Justlethervolunteer` 的共同點不是漏某個固定詞，而是 engine 只給一整區文字、沒可靠 word boundary。只有整句 box 時要先對 bbox 做墨水／間距幾何；幾何不夠才產生 lexical／closed-class 候選。合法單字如 Unexpectedly 不該進拆詞。

#### Trigger

- 單 part Latin 黏串（parts=1）且僅有整句 bbox
- closed-class／詞典被當成第一刀或直接改 raw
- 合法長單字被標 suspicious 並強制拆開
- 用 `if raw == "…"` 特例代替 regression corpus

#### Evidence

- Tool: product dogfood OCR cache（sticky single-part Latin）
- Sanitized：geometry-before-lexical layered recovery
- Paths／titles：`<PROJECT_ROOT>` only

#### Generalized Lesson

1. **座標優先**：有 char／word boxes 或可做 ink gap 時，幾何先於 lexical。
2. **整句 bbox ≠ 夠用**：必須 intra-box geometry attempt。
3. **Closed-class = candidate generator**，不是 truth generator。
4. **Suspicious 閘門**：dictionary exact-match → ok；非 Latin→split。
5. **Regression corpus** 驗證能力；禁止硬編碼 raw 特例。
6. **adapter_only**：不開新 Phase／runtime／斷詞 Agent。

#### Agent Action

改 OCR boundary 時：確認 recovery 呼叫序；lexical 只 append derived candidates；單字 exact-match 不拆；案例寫入／對齊 `latin-boundary-regression.yaml`。

#### Validation

- Corpus raw→expected：`appointmentfortoday`、`tomasturbate`、`Workhard`、`aghostis`、`alsoromantic`、`Justlethervolunteer`、`commitmentremainsvalid`、`Itcan'tbesuchacoincidence`、`Ihaveahusband's`
- Negative：`Unexpectedly`／`uncomfortable`／`volunteer` 保持 boundary ok（不強制拆）
- Recovery 有 geometry 嘗試或 lexical `status=candidate`；raw 未覆蓋

#### Goal / Action / Validation

- Goal: 單 part Latin 黏串可經分層 recovery 產出可追溯 candidates，且不誤拆合法單字。
- Action: 契約＋corpus writeback；adapter 對齊順序與閘門。
- Validation: regression cases + negative singles；raw 未覆蓋。

#### Applies When

- Latin 硬字幕單 part／單框黏串；intra-box boundary recovery；closed-class／詞典參與斷詞候選

#### Does Not Apply When

- 多 part 已切開、僅 join 接縫問題（見 join token-seam lesson）
- 純 CJK 無需 Latin 詞界

#### Promotion Target

- `workflow/narrative-video-production/text-evidence-ocr-boundary.md`
- `workflow/narrative-video-production/records/latin-boundary-regression.yaml`
- `workflow/narrative-video-production/records/text-evidence.yaml`
- Plan companion `42-ocr-boundary-and-script-aware-normalization.md`

#### Required Linked Updates

- Evidence：`…/evidence/2026-10-01-ocr-latin-boundary-geometry-before-lexical.md`
- `feedback/history/development-guidance/common/README.md` 索引

#### Closure

Dogfood 黏串進 corpus；workflow／plan companion 寫明 geometry-before-lexical。

#### References

- `workflow/narrative-video-production/text-evidence-ocr-boundary.md`
- `workflow/narrative-video-production/records/latin-boundary-regression.yaml`
- Plan companion `42-ocr-boundary-and-script-aware-normalization.md`
- Evidence：`…/evidence/2026-10-01-ocr-latin-boundary-geometry-before-lexical.md`
