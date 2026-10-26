> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-06 - Pipeline behaviour needs a host-level test, not a component test

Status: candidate

#### One-line Summary

When the claim is about what the running host does (status logged, log level, response for an aborted request), test it through the real composed pipeline; a middleware tested in isolation can be correct while another layer in the real order still produces the old behaviour.

#### Human Explanation

A change mapped client-aborted requests to 499 with an Information log inside a global exception middleware. Its unit tests passed and were mutation-sensitive. In the real host, request logging sat inside that middleware, saw the cancellation first, and logged "responded 500" at Error — also persisted to the operational log. The component was right; the system behaviour the change was meant to deliver did not happen. Order-dependent cross-cutting concerns (request logging, exception handling, auth, compression, CORS) interact only when composed.

#### Trigger

- The acceptance is phrased as host behaviour: "no Error log", "logged as 499", "header present on every response".
- The changed component is one of several middleware/filters registered in a host composition method.

#### Evidence

- Tool: ASP.NET Core host with Serilog request logging and a global exception middleware
- Sanitized excerpt: middleware unit tests green; real host logs `ERR HTTP POST … responded 500` with the cancellation stack for a client abort
- Evidence path: project-local plan verification evidence for an async-conversion slice

#### Generalized Lesson

1. Write at least one test through the real pipeline (e.g. a web application factory with the production composition and real log sinks captured) for any acceptance stated as host behaviour.
2. Assert on what operators see: every logger's events and levels, the recorded status, persisted operational logs — not only the component's return value.
3. Check registration order explicitly when two concerns both react to the same exception.

#### Agent Action

- Do add a host-level assertion when a change targets logging, status mapping or other cross-cutting behaviour.
- Do not close such a change on component tests alone, however mutation-sensitive.

#### Goal / Action / Validation

- Goal: Cross-cutting behaviour is verified as composed in the host.
- Action: Host-level test capturing all log events and the final status.
- Validation: The host-level test fails on the pre-fix composition and passes after.

#### Applies When

- Middleware, filters, interceptors, handlers whose effect depends on composition order.

#### Does Not Apply When

- Pure functions or components whose contract is self-contained and not observed through the pipeline.

#### Validation

A host-level test exists for the acceptance, captures every logger, and fails on the composition that exhibited the bug.

#### Promotion Target

- `workflow/software-delivery/` test-strategy guidance on integration vs component tests, if repeated.

#### Required Linked Updates

- Indexed in `feedback/history/development-guidance/common/README.md`.
- 已承接為既有 evidence-chain 的候選補充；尚無獨立重用證據，維持 candidate。

#### Transfer Contract

- Inputs: 驗收描述最終 host 行為，多個元件會處理同一事件
- Mechanism: 元件順序與共同副作用可能使局部正確仍產生錯誤最終結果
- Decision: 跨元件的宣稱採實際組合入口驗證，觀測所有受影響的輸出與副作用
- Verification: 在舊組合與修正組合執行同一案例，核對最終狀態與所有相關紀錄
- Counterexample: 自足純函式的局部契約不需一律提升為 host 測試
- Transfer case: 訊息處理器的重試與去重順序共同決定最終通知次數

#### Intelligence Review

- Decision: existing
- Reason: 增量是宣稱範圍與證據鏈覆蓋的判斷，沿用既有 atom，避免按事件建立重複智慧。
- Target: intelligence/engineering/execution/validation-reasoning/evidence-chain-validation.md

#### Abstraction Check

- Swap-project test: PASS — 訊息處理器的重試與去重順序共同決定最終通知次數；為可攜推演，未宣稱獨立實測。
- Project evidence kept in: `<PROJECT_ROOT>/` 的受控驗證紀錄。
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS

#### 承接修訂 2026-10-26

已與 [evidence-chain 候選判斷](../../../../intelligence/engineering/execution/validation-reasoning/evidence-chain-validation.md) 雙向接線；分類／domain 索引與 RAG 已同步檢查，維持 candidate，不把抽取當 promotion。
