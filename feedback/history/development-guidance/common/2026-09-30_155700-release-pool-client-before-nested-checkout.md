> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-30 - Release pool client before nested same-pool checkout

Status: candidate

#### One-line Summary

Never hold a DB pool client across an await that also checks out from the same pool; saturation turns into an uninterruptible hang that soft reconcile cannot clear.

#### Human Explanation

Long-running job loops often keep one pool client for the account lifetime, then call a replacement helper that peeks/leases via additional `pool.connect()` calls. When several jobs do this under a small `pool.max`, every holder waits for a free client while still occupying one. `shared.stop` / stale-wait interrupt cannot unblock a hung `connect()`, and soft requeue skips in-process `activeJobs`, so the durable row stays `running` with zero leases until process restart—even when the idle funded pool is large.

#### Trigger

- Logs stop right after “release and replace” / insufficient-funds with no “replace account → …” line.
- Job status remains `running`, open lease count is 0, idle funded accounts remain pickable.
- Stale-wait interrupt logs (if any) do not resume the job; only worker restart recovers.

#### Evidence

- Tool: Node `pg.Pool` job worker + reconcile soft requeue / stale-wait interrupt
- Sanitized excerpt: after insufficient-funds release, replace hung on nested `pool.connect()` while outer loop still held its client; restart requeued orphaned running jobs and they leased immediately
- Evidence path: project-local worker account-loop + replace-lease path (keep host/account labels out of Ai-skill)

#### Generalized Lesson

1. Before any helper that checks out from the same pool, release (or hand off) the outer client; disconnect session sockets first if needed.
2. Size `pool.max` for concurrent jobs **plus** peek/lease/reconcile checkouts; add `connectionTimeoutMillis` so a starved connect fails instead of hanging forever.
3. Do not rely on in-process stop flags to recover a hung pool wait; design so reconcile paths themselves are not starved of pool clients.
4. Treat “running + 0 leases + quiet logs + pickable inventory” as pool/deadlock suspicion, not farm shortage.

#### Agent Action

- Do release the loop client before nested replace/mint/wait helpers that use the same pool.
- Do not assume soft reconcile will heal in-process hangs blocked on `pool.connect()`.
- Do not raise farm/mint capacity as the first diagnosis when pickable inventory is already non-empty.

#### Goal / Action / Validation

- Goal: Account replacement never deadlocks the worker pool.
- Action: Early client release before replace; pool timeout + adequate max; restart only as recovery for already-hung processes.
- Validation: Force insufficient-funds replace under concurrent jobs → see “replace account →” without worker restart; inject pool pressure → connect times out instead of permanent zombie.

#### Applies When

- One Nesting unit of work holds a pool client and calls another function that also checks out from that pool.
- Soft reconcile skips in-process active work and interrupt only flips a flag.

#### Does Not Apply When

- Replacement reuses the same already-checked-out client for peek/lease (no nested checkout).
- Each unit of work uses a dedicated connection outside a shared pool.

#### Validation

Under concurrent account loops, force an insufficient-funds replace path; confirm a “replace → next account” log appears without worker restart, and that a saturated pool surfaces a connect timeout instead of a permanent `running` + 0-lease zombie.

#### Promotion Target

- `workflow/development-guidance/` (optional later — connection-pool nesting hazard)
- Project worker README / job-runner comments (project docs, not Ai-skill)

#### Required Linked Updates

- N/A for Ai-skill workflow docs this turn — lesson is candidate pattern only; project evidence stays under `<PROJECT_ROOT>`.

#### Related

- Promotion: N/A (candidate; project fix already deployed)
- Linked updates: N/A
