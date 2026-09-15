> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-15 - Stale peer refused → refresh discovery; do not fail the job

Status: candidate

#### One-line Summary

When a long-lived TCP peer returns `ECONNREFUSED` / host-unreachable, refresh the official discovery source, persist the new endpoint, and keep the durable job `running` instead of marking it failed.

#### Human Explanation

Farm or session workers often cache a game/media peer from a prior successful connect. That peer can go away while the product’s discovery endpoint still returns a live address. Treating connect-refused as a fatal job error destroys recoverable work and forces operator remints. The recoverable path is: classify unreachable connect errors, force-refresh discovery, rebind, update stored host, and only defer finish (leave status running for reconcile/requeue) if refresh also fails.

#### Trigger

- Worker/client logs show `connect ECONNREFUSED <cached-ip>:<port>` (or `EHOSTUNREACH` / connect timeout) against a previously stored peer.
- Official app or phone still connects; only the automation peer is stale.
- Jobs flip to `failed` on open/login even though credentials are valid.

#### Evidence

- Tool: TCP client + HTTP discovery refresh + job-runner status policy
- Sanitized excerpt: cached peer refused; discovery returned a different live host; job left `running` with wait note instead of `failed`
- Evidence path: `<PROJECT_ROOT>` worker/client host-resolve + job finish policy (project-local)

#### Generalized Lesson

1. Distinguish **peer unreachable** (`ECONNREFUSED`, host/net unreachable, connect timeout) from auth/balance failures.
2. On peer unreachable: clear discovery cache → force resolve → rebind session → persist new host on the account/session row.
3. Never set durable job status to `failed` solely for peer unreachable; use defer-finish / wait note so reconcile can requeue.
4. Ignore known-dead historical peers when an explicit host is supplied so rows can self-heal on next open.

#### Agent Action

- Do implement client-level failover and worker-level non-fatal policy together; client-only still lets outer `catch` fail the job.
- Do not treat phone/manual connect success as proof the automation host cache is current—always have an auto discovery path.
- Do not ask the operator to remint accounts for a refused peer alone.

#### Goal / Action / Validation

- Goal: Keep long-running collect/bet jobs alive across peer rotation.
- Action: Classify unreachable errors; CM4/discovery refresh; persist host; `deferFinish` / leave `running`.
- Validation: Inject refused connect → see discovery refresh + host column update + job status remains `running` (or returns to `pending` via no-lease reconcile), never `failed` for that class alone.

#### Applies When

- Work uses a cached TCP peer plus a separate discovery URL that lists current peers.
- Jobs are durable DB rows that should survive transient network/peer moves.

#### Does Not Apply When

- Auth/token invalidation, account ban, or true insufficient-funds paths (those need their own policies).
- The discovery endpoint itself is down and no alternate resolve exists (then wait/defer is correct; inventing peers is not).

#### Validation

Reproduce with a stale host in storage while discovery returns a live host; confirm one open path recovers and job status is not `failed`.

#### Promotion Target

- `workflow/development-guidance/` (optional later)
- Project worker README host-failover section (project docs, not Ai-skill)

#### Required Linked Updates

- N/A for Ai-skill workflow docs this turn — lesson is candidate pattern only; project evidence stays under `<PROJECT_ROOT>`.
