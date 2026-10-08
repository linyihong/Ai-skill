> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-08 - Translation Repair Loop: governance owns correction; Actor only selects／repairs

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。來自產品 dogfood 已落地的 `selection_notes`／`repair_of`／Independent Validation，對齊 Loop Engineer（Governance Bootstrap → Local Autonomous Loop）的抽象。

#### One-line Summary

校正能力屬於 Workflow／Governance（Failure Evidence → Classification → Constraint／Candidate → repair → Independent Validation → Escalation），不屬於 ChatGPT／本地模型；Actor 只在縮小後的 Feasible Candidate Space 做 Selection／Repair，且不得自關 PASS。

#### Human Explanation

產品若把「誰校正」想成「永遠叫強模型仲裁」或「fine-tune 記答案」，會把答案記憶誤當成決策程序。正確切法：Validator 與 Governance 可以非 AI；Selection／Repair 才交給 Actor。`selection_notes` 宣告 HARD／FORBIDDEN／CANDIDATE_SPACE；`repair_of` 保留完整失敗候選。累積的是 failure pattern + violated_constraint + repair_direction，不是單句對錯表。known → 本地修；unknown → escalate。成功率≠學會了——觀測 escalation／known-repair／regression。

#### Trigger

- 每次 Validation FAIL 都呼叫 ChatGPT／Claude 當仲裁者
- 用越來越長的案例提示取代 registry／constraint
- 本地模型輸出「我覺得對」就寫 accepted／改 SoT
- 把 fine-tune「陈小姐→Nona Chen」當優先學習路徑
- 只報 first-pass 成功率，不看 escalation／unknown／constraint reuse

#### Evidence

- Tool: product `selection_notes` + `repair_of` + mechanical／declared Independent Validation；Ai-skill `adapters/repair-loop.md`
- Sanitized: Thai HIT soft-paraphrase、Arabic divorce／marry polarity、identity_translations binding — repair via constraint／candidate refinement，非 cloud 當唯一校正者
- Evidence path: plan evidence `2026-10-08-repair-loop-interfaces.md`；episode／host／keys 留 `<PROJECT_ROOT>` analysis

#### Generalized Lesson

1. **Correction ∈ Workflow**：Governance 擁有 registry／constraints／Finality／escalation；Actor 可換。
2. **Loop Interfaces**：`selection_notes`＝Selection Space Guard；`repair_of`＝完整失敗候選 + validation 宣告。
3. **Accumulate Failure Evidence**，不是答案表：`pattern_id`／manifestation／violated_constraint／repair_direction／known。
4. **Governance Bootstrap → Local Autonomous**：Phase A 強模型／人教 loop；Phase B known 本地自修；unknown 才 escalate。
5. **Local Arbitration Actor 可提案，不可改世界**：arbitration YAML → Mechanical Validator + Governance 決定是否接受。
6. **Telemetry > raw success**：escalation↓、known repair↑、unknown↓、regression recurrence。

#### Agent Action

- 讀 [`workflow/translation/adapters/repair-loop.md`](../../../../workflow/translation/adapters/repair-loop.md) 與 [`execution-flow.md`](../../../../workflow/translation/execution-flow.md) §12b。
- 產品配線：Validation FAIL → structured Failure Evidence → `selection_notes`／`repair_of`；禁止每錯都 escalate。
- 禁止把 dogfood 單句對錯寫進可重用 registry；先抽象 pattern。

#### Goal / Action / Validation

- Goal: 校正閉環屬 Workflow／Governance；known 本地修、unknown escalate；Actor 不得自關 PASS。
- Action: Failure Evidence → Classification → Constraint／Candidate → `selection_notes`＋`repair_of` → Independent Validation；觀測 Loop Telemetry。
- Validation or reference source: product known-script residue → `local_repair`；unknown pattern_id → `arbitration`＋`ESCALATION` notes；Ai-skill adapter／execution-flow §12b／plan evidence。

#### Applies When

- 字幕／配音／caption 把 LLM 當 Selection／Repair Actor 的產品路徑
- dogfood 後要把「校正」沉澱成可重用 loop，而非案例提示或 fine-tune 答案表

#### Does Not Apply When

- 純人工譯稿、無 Independent Validation／Finality 的離線編輯
- timing／layout／burn／TTS／publish packaging（非 translation content Finality）

#### Validation

- Adapter `repair-loop.md` 與 execution-flow §12b 交叉引用存在
- Product：`build_failure_evidence` known→`local_repair`；unknown→`arbitration`；`project_repair_selection_notes` 含 `FAILURE_EVIDENCE`／相容 `VALIDATION_FAILURE`
- 不宣稱 overnight／ASS／burn／publish PASS

#### Promotion Target

- `workflow/translation/adapters/repair-loop.md`
- `workflow/translation/adapters/README.md`
- `workflow/translation/execution-flow.md`
- `workflow/translation/adapters/dogfood-acceptance.md`

#### Promotion Record

本輪已新增 repair-loop adapter，並更新 README／execution-flow／dogfood-acceptance 交叉引用；契約 YAML schema 未改。

#### Required Linked Updates

- `workflow/translation/adapters/repair-loop.md`
- `workflow/translation/adapters/README.md`
- `workflow/translation/execution-flow.md`
- `workflow/translation/adapters/dogfood-acceptance.md`
- `feedback/history/development-guidance/common/README.md` index
- `plans/active/2026-09-22-1000-translation-decision-workflow/evidence/README.md` Run 索引
- `plans/active/2026-09-22-1000-translation-decision-workflow/10-failure-pattern-learning.md`

#### Links

- Adapter：[`repair-loop.md`](../../../../workflow/translation/adapters/repair-loop.md)
- Prior：[`2026-09-29_105651-translation-decision-selection-adapter-finality-gates.md`](2026-09-29_105651-translation-decision-selection-adapter-finality-gates.md)
- Plan：[`10-failure-pattern-learning.md`](../../../../plans/active/2026-09-22-1000-translation-decision-workflow/10-failure-pattern-learning.md)
