> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-01 - Cap recovery loops; never reset the counter that gates them

Status: candidate

#### One-line Summary

A recovery path that clears the same retry counter it is meant to advance will thrash forever; give failover a separate budget and prefer replace/requeue over `failed` for known-transient lobby errors.

#### Human Explanation

Long-running bet/collect jobs often soft-retry desk enter, then rotate host (CM4) on lobby flakes (`X007*`, “Authorization first”). If host rotation resets the desk-resync counter, every CM4 starts a fresh enter budget → infinite loop, expired leases, zombie `running` with zero spins. A bare `/auth/` matcher can also misclassify “Authorization first” as hard auth death. Cap host-rotation attempts per account, keep the enter counter across lobby CM4, heartbeat the lease during recovery, and on budget exhaust replace the account (or requeue) — do not mark the durable job `failed` for lobby-only errors.

#### Trigger

- Logs alternate `enter resync 1/5` → `lobby transient re-enter` → `CM4 failover` → `reconnected` with no settle inserts.
- Job stays `running` with 0 open leases / 0 recent spins while peers on other cabinets still progress.
- Historical job ends `failed` with only `GetCategoryInfo failed: X007…` / Authorization-first text.

#### Evidence

- Tool: Node worker account loop + host discovery failover
- Sanitized excerpt: lobby CM4 reset deskResyncs; thrash for hours; restart seated and spun; fix caps lobby CM4 and stops counter reset
- Evidence path: project-local worker job-runner + host-resolve helpers

#### Generalized Lesson

1. Separate budgets: soft enter retries ≠ host-rotation attempts ≠ TCP-drop attempts.
2. Never reset the counter that decides “give up enter” inside the recovery that runs after give-up.
3. Heartbeat leases during long recovery so peers cannot steal the account mid-thrash.
4. Classify lobby flakes as transient; do not use substring `auth` alone for hard auth-death.
5. Exhausted lobby budget → replace account / requeue, not durable `failed`.

#### Agent Action

- Do add an explicit host-rotation budget and log `n/max` on each consume.
- Do not reset enter/resync counters on lobby-only host rotate.
- Do not fail keep jobs solely on `X007*` / Authorization-first after soft retries.

#### Goal / Action / Validation

- Goal: Keep jobs either spinning or replacing accounts during lobby storms; never infinite CM4 thrash or lobby-only `failed`.
- Action: Cap lobby CM4; preserve deskResyncs; tighten auth-death regex; heartbeat; replace on exhaust.
- Validation: Inject repeated GetCategoryInfo X007 → see budget logs then replace/requeue without `failed`; confirm deskResyncs does not reset on lobby CM4.

#### Applies When

- Nested soft-retry + host failover on the same error class.
- Durable job rows that operators expect to keep running across lobby flakes.

#### Does Not Apply When

- True auth/token invalidation or account ban (those need invalidate + remint).
- Hard product stop conditions (max spins, operator cancel).

#### Validation

Force repeated lobby GetCategoryInfo failures under a running job; confirm finite CM4 budget, account replace, and status never becomes `failed` for that error class alone.

#### Promotion Target

- `workflow/development-guidance/` (optional later — recovery budget nesting)
- Project worker README lobby section (project docs)

#### Required Linked Updates

- N/A for Ai-skill workflow docs this turn — candidate pattern; project evidence stays under `<PROJECT_ROOT>`.
