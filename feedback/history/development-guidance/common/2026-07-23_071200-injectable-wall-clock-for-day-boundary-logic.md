> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md)、[reusable-guidance-boundary](../../../../enforcement/reusable-guidance-boundary.md)、[sanitization](../../../../enforcement/sanitization.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-07-23 — Injectable wall clock for day-boundary logic

Status: candidate

#### One-line Summary

Time-dependent business logic (local day, near-midnight windows, quotas) must take an injectable clock or explicit `Instant`/`LocalDate` parameters; tests freeze that clock instead of changing the OS date.

#### Human Explanation

Reviewers often reject suites that assert “today / near midnight” behavior while production code calls a global “now”. Those tests either flake with the real wall clock or force operators to change the machine date. The reusable fix is a small injectable clock abstraction (or method parameters) plus fixed-clock unit tests that prove day-boundary branches without touching the OS.

#### Trigger

- Business rules depend on local calendar day, remaining time until midnight, or “created today”
- Code or tests call global wall-clock helpers for those decisions
- Reviewers ask for mockable time / clock injection before accepting the test coverage

#### Evidence

- Tool: SDK platform session (sanitized)
- Sanitized excerpt: introduced a shared injectable clock; wired day-resolution / planner paths to it; added frozen-clock unit tests for near-midnight vs early-day branches
- Evidence path: keep target module/class names in `<PROJECT_ROOT>` docs only

#### Generalized Lesson

```
1. Prefer Instant/LocalDate as method parameters when a single event time is enough
2. When a component repeatedly needs "now" / local day, inject a shared clock type
3. Production default = system clock; tests use fixed (or controllable) clock
4. Never validate day-boundary behavior by changing the OS date
5. Audit new Instant.now()/LocalDate.now() call sites in business logic during review
```

Do not claim full coverage until the day-boundary paths under test actually read the injected clock.

#### Agent Action

When adding or reviewing time-dependent SDK / service logic:

1. Search for global wall-clock calls in the path under change
2. Extract an injectable clock or pass explicit time into the decision method
3. Add at least one frozen-clock test for the risky window (e.g. near local midnight)
4. Leave hostnames, package names, and live run results in project docs

#### Validation

- Fixed clock at a known near-midnight Instant yields the late-day branch
- Fixed clock early in the same local day yields a non-late-only branch set
- Tests pass without changing the host OS timezone/date

#### Applicable / Not applicable

- Applicable: schedulers, daily quotas, first-day seeding, cooldown calendars, persona/day rolls
- Not applicable: pure logging timestamps where wall clock is the product requirement and no business branch depends on it

#### Linked updates

- N/A (candidate lesson; no enforcement rule promotion in this pass)

#### Promotion

- candidate → consider `intelligence/engineering/` testing guidance if the pattern repeats across SDKs

#### Goal / Action / Validation

- Goal: 保留既有 lesson 的可驗證結論。
- Action: 依原始 Evidence 與 Trigger 重做相關檢查。
- Validation: 結果與原始結論一致才可重用。

#### Applies / Does Not Apply

- Applies: 本條既有 Trigger、前提與 Evidence 相符時。
- Does not apply: 跨專案、前提不明或缺少原始證據時；先重新驗證。

#### Promotion Target

- 尚未 promotion；保留為 candidate history，需有獨立 reuse evidence 才可提升。

#### Required Linked Updates

- Not applicable: 這是歷史 closure 修復；未新增可安全提升的 canonical guidance，category index 已存在。

#### Goal / Action / Validation

- Goal: 保留既有 lesson 的可驗證結論。
- Action: 依原始 Evidence 與 Trigger 重做相關檢查。
- Validation: 結果與原始結論一致才可重用。

#### Applies / Does Not Apply

- Applies: 本條既有 Trigger、前提與 Evidence 相符時。
- Does not apply: 跨專案、前提不明或缺少原始證據時；先重新驗證。

#### Promotion Target

- 尚未 promotion；保留為 candidate history，需有獨立 reuse evidence 才可提升。

#### Required Linked Updates

- Not applicable: 這是歷史 closure 修復；未新增可安全提升的 canonical guidance，category index 已存在。
