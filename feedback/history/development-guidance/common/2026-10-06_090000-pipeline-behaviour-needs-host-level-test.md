> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

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
- No workflow or enforcement change yet; candidate stays local until repeated.
