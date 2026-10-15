> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-15 - now() is the transaction start; age and window gates need clock_timestamp()

Status: candidate

#### One-line Summary

In PostgreSQL `now()` / `current_timestamp` is frozen at transaction start; a holder of an old open transaction can back-date any timestamp the routine writes with it, so minimum-age and reset-window gates must stamp with `clock_timestamp()`.

#### Human Explanation

A tamper-protection routine recorded the time a month was marked archived and refused release until a minimum age had passed. Written with `now()`, a caller could open a transaction, idle, do the work in another session, and finish the marking inside the old transaction: the stored time was an hour old and the age gate passed immediately. Tests passed because every test used a fresh transaction.

#### Trigger

- A SQL routine, trigger or default stamps a time that is later compared against a window, minimum age, expiry or ordering.
- The caller is less trusted than the database owner.

#### Generalized Lesson

1. Stamp gate-relevant times with `clock_timestamp()` (statement wall clock), including event-trail rows and column defaults.
2. Test with a transaction held open for a few seconds before the marking statement, and assert stored time ≥ the earlier committed time.
3. Add an ordering invariant (staged ≤ committed ≤ archived) as a verifier rule when a CHECK would break existing rows.

#### Agent Action

- Do grep routines for `now()` when reviewing any age/window gate.
- Do not accept a gate test that never holds a transaction open.

#### Goal / Action / Validation

- Goal: Age and window gates cannot be back-dated by transaction age.
- Action: clock_timestamp() plus a held-open-transaction test.
- Validation: the test is red against `now()` and green against `clock_timestamp()`.

#### Evidence

- Tool: PostgreSQL 16 SECURITY DEFINER routines with a minimum-age gate
- Sanitized excerpt: a transaction held open 1 h, marking done inside it stored the time one hour earlier; the age gate passed 0.2 s after the commit
- Evidence path: project-local plan evidence for a cold-tier protection slice

#### Applies When

- Routines, triggers or column defaults whose timestamps gate an age, window, expiry or ordering decision.

#### Does Not Apply When

- Purely descriptive timestamps no decision reads.

#### Validation

A test holds a transaction open before the marking statement; it is red with `now()` and green with `clock_timestamp()`.

#### Promotion Target

- `workflow/software-delivery/` database-routine review checklist, if repeated.

#### Required Linked Updates

- Indexed in `feedback/history/development-guidance/common/README.md`.
- No workflow or enforcement change yet; candidate stays local until repeated.
