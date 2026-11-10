### 2026-11-10 — 驗證重播輸入的起點與連續性

Status: candidate

#### One-line Summary

事件摘要有終點時，仍需以原始起始事件、狀態交接和獨立計數器證明重播輸入完整。

#### Trigger

已保存的事件摘要可結算，但以摘要重建狀態時出現額外變更。

#### Evidence

- Tool: 唯讀資料查詢、原始事件解碼、逐步狀態回算。
- Sanitized excerpt: 把原始起始事件接回後，原先未解釋的變更能由已記錄的動作還原。另有一筆前後狀態位置相同，但剩餘次數一次少兩次，揭露漏段。末段同一份摘要出現兩次，按同狀態及動作序列去重後可回算。
- Reuse Evidence: 單一專案、數百個連續對照；待獨立專案重用。

#### Generalized Lesson

先以來源識別碼配對原始起始事件，再查相鄰狀態的交接及獨立的序號或剩餘次數。狀態可以回到同一位置，單用位置連續會漏掉循環。以狀態、計數器與事件序列共同辨認重複摘要，保留真正重複發生的動作。

#### Goal / Action / Validation

- Goal: 判斷模型的殘差能否作為規則反例。
- Action: 將完整序列、可見缺段與重複摘要分組，以原始事件配對起點。
- Validation: 用已保存的絕對狀態逐項比對預測，分別保存樣本分母與缺段原因。

#### Transfer Contract

- Inputs: 起始事件、摘要動作、前後狀態、獨立序號或計數器、絕對讀回。
- Mechanism: 摘要可能漏掉整段或重複收錄，循環會遮蔽位置斷點。
- Decision: 只在已核對連續的區間裁決規則。缺段保留為未知輸入。
- Verification: 對完整序列與缺段序列分開跑同一預測。
- Counterexample: 若計數器會回補或重設，須先確認更新契約，才能用它判斷缺段。
- Transfer case: 多階段交易或工作流程的事件重播。

#### Promotion Target

- analysis/reverse-engineering/collection/capture-review.md

#### Required Linked Updates

已將起點配對、交接及獨立計數器核對寫入既有抓包審查流程。Step 6 不適用，這是審查方法。Step 7 由本條及審查步驟承接，將漏段與模型反例分開。

#### Abstraction Check

- Swap-project test: PASS。
- Project evidence kept in: <PROJECT_ROOT>/docs/。
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS

#### Agent Action

用原始起始識別碼分組，依序查狀態交接、計數器與重複摘要，再比較保存的絕對狀態。

#### Applies When

可以取得起始或交接事件，且至少有一個獨立序號或計數器。

#### Does Not Apply When

只有終點快照可用，或計數器更新契約未知。

#### Validation

保存連續組和缺段組的分母及殘差，並標記各組的輸入邊界。

#### Intelligence Review

- Decision: method-only
- Reason: 目前證據支持審查步驟，尚無跨目標重用足以獨立成知識單元。
- Target: analysis/reverse-engineering/collection/capture-review.md
