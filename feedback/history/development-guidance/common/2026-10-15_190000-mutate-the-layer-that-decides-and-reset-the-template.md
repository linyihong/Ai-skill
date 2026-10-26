> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-15 - A mutant that survives may be in the wrong layer; a shared template must be rebuilt after every mutant

Status: candidate

#### One-line Summary

When mutation-testing a layered database design, apply each mutant in the layer that finally decides (the last statement that runs), and rebuild any shared migrated template database after each mutant; otherwise a "surviving" mutant is a no-op and a "red" one may be red for the previous mutant's reason.

#### Human Explanation

Two things looked like results and were not. A grant appended at the end of one setup step survived every live test because a later step resets all function privileges: the mutant never reached the database. And a batch of four mutants all showed the same seven red tests because the migrated template database still carried the first mutant's routine text, not because their own changes did anything.

#### Generalized Lesson

1. Before running a mutant, state which statement or file decides the behaviour last, and put the change there (or in the earlier layer *and* check that nothing later undoes it).
2. If tests copy a migrated database as a template, rebuild it after each mutant that changes migration text, and run the unmutated suite once to prove the template is clean before the next mutant.
3. Read a mutant's red set against what it changes: a set that is identical for unrelated mutants is the template or the harness, not the mutants.

#### Agent Action

- Do record, for each mutant, the layer it was applied in and whether the live lane went red; a 0-red integration result with a red text pin is a finding about the mutant's layer, not a pass.
- Do not report mutation results from a run that did not rebuild the template since the previous mutant.

#### Goal / Action / Validation

- Goal: mutation tables that say what each guard really protects.
- Action: apply at the deciding layer, rebuild templates between mutants, compare red sets.
- Validation: a clean run of the whole class after the last restore, and `git diff` empty for the mutated files.

#### Evidence

- Tool: PostgreSQL migrations, setup pack applied after a routine file, TEMPLATE-copied test databases
- Sanitized excerpt: grant appended to a routine file had no effect because the ACL step ran after it; four mutants shared the same seven failures until the template was rebuilt
- Evidence path: project-local Phase 7 evidence file

#### Applies When

- Mutation checks of privileges or routines that are installed by several ordered steps, with tests that clone a prebuilt database.

#### Does Not Apply When

- Single-step designs where the mutated file is the only place the behaviour is decided and tests build their own database.

#### Validation

The mutation table lists, per mutant, the layer, the red set of the live lane and of the text pins, and a note when the live lane stayed green and why.

#### Promotion Target

- Test-strategy guidance for mutation checks of layered database objects, if repeated.

#### Required Linked Updates

- Indexed in `feedback/history/development-guidance/common/README.md`.
- 已承接為既有 evidence-chain 的候選補充；尚無獨立重用證據，維持 candidate。

#### Trigger

- 突變結果可能被後續覆寫或共享衍生狀態污染。

#### Transfer Contract

- Inputs: 多層安裝步驟與共用衍生範本，突變測試宣稱守衛有效
- Mechanism: 後續步驟覆寫突變，或範本殘留前一突變，會切斷原因與測試結果的關聯
- Decision: 先證實最後決定層已載入當次突變，再以隔離範本判讀失敗集合
- Verification: 讀回當次有效物件並比對乾淨基線；不同突變固定相同失敗集合時先排查測試隔離
- Counterexample: 單一步驟且每次重建全部狀態的純函式測試不需要額外範本重建
- Transfer case: 生成設定被啟動參數覆蓋時，先確認實際有效設定再比較故障注入結果

#### Intelligence Review

- Decision: existing
- Reason: 增量是宣稱範圍與證據鏈覆蓋的判斷，沿用既有 atom，避免按事件建立重複智慧。
- Target: intelligence/engineering/execution/validation-reasoning/evidence-chain-validation.md

#### Abstraction Check

- Swap-project test: PASS — 生成設定被啟動參數覆蓋時，先確認實際有效設定再比較故障注入結果；為可攜推演，未宣稱獨立實測。
- Project evidence kept in: `<PROJECT_ROOT>/` 的受控驗證紀錄。
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS

#### 承接修訂 2026-10-26

已與 [evidence-chain 候選判斷](../../../../intelligence/engineering/execution/validation-reasoning/evidence-chain-validation.md) 雙向接線；分類／domain 索引與 RAG 已同步檢查，維持 candidate，不把抽取當 promotion。
