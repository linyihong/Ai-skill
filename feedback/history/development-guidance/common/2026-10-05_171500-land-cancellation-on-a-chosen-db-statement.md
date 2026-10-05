> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-05 - Land a cancellation on a chosen database statement, including COMMIT

Status: candidate

#### One-line Summary

To prove a cancellation-safety or idempotency claim, make the request block on the exact statement you care about (lock it, or delay COMMIT with a deferred trigger), then let a short client timeout cancel it — instead of hoping a random cancel lands there.

#### Human Explanation

Converting request paths to async with a forwarded CancellationToken changes when work can stop: mid-query, mid-allocation, or during COMMIT. Claims like "a cancelled allocation either rolls back or leaves a replayable result" are only believable if each window is actually hit. Random timeouts almost never land inside COMMIT. Holding a lock from a second connection makes a chosen statement block deterministically; a `DEFERRABLE INITIALLY DEFERRED` constraint trigger that sleeps makes COMMIT itself block. A short client timeout then cancels precisely there, and the retry behaviour can be read back from the database.

#### Trigger

- A change forwards a request CancellationToken into code that writes (allocation, sequence numbers, handshakes, outbox).
- A design argument relies on idempotent retry after cancellation ("the device retries the same id").

#### Evidence

- Tool: real API process + PostgreSQL, throwaway runner mode
- Sanitized excerpt: table lock on the read → OCE, no rows, retry yields one allocation; FOR UPDATE row lock on the sequence row → query cancelled (no backends left blocked), no gap; deferred constraint trigger with pg_sleep(4) at commit → full rollback, retry yields one allocation
- Evidence path: project-local plan verification evidence for an async conversion slice

#### Generalized Lesson

1. Per window, pick a blocking mechanism: `LOCK TABLE … IN ACCESS EXCLUSIVE MODE` for a specific read; `SELECT … FOR UPDATE` on the row a write needs; a deferred constraint trigger calling `pg_sleep` for COMMIT.
2. Set the client timeout shorter than the block; confirm the server-side query really stopped (no backend left waiting on the lock), not just the client.
3. Read back after each case: rows written, sequence position, audit rows; then send the same idempotency key again and compare the response byte for byte.
4. Run the same cases at the parent commit to show the behaviour change, not just the end state.

#### Agent Action

- Do cover each cancellation window separately (before write, during write, during COMMIT, after COMMIT).
- Do not accept "cancellation is safe" from code reading alone when the claim is about retries.

#### Goal / Action / Validation

- Goal: Cancellation-safety claims are demonstrated, not argued.
- Action: Block the chosen statement deterministically and cancel there; read back and retry.
- Validation: Each window shows rollback or a replayable committed result; retries produce exactly one result with no gaps.

#### Applies When

- Request-path writes with idempotency keys, sequences, or external side effects, on PostgreSQL (similar techniques exist for other databases).

#### Does Not Apply When

- Read-only paths, or writes with no retry/idempotency contract.

#### Validation

For each window, the database read-back and the retry response match the claimed behaviour, and the same cases run at the parent commit show the intended difference.

#### Promotion Target

- `workflow/software-delivery/` testing guidance for concurrency and cancellation, if repeated in another project.

#### Required Linked Updates

- Indexed in `feedback/history/development-guidance/common/README.md`.
- No workflow or enforcement change yet; candidate stays local until repeated.
