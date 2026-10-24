# Phase 2 remaining engineering protocol

Frozen before implementation. Selector method exploration is closed for the exposed 26 cases. No new reviewer or semantic tuning is planned.

## Scope and acceptance

- Seven Phase 0 validation-reasoning sources only; `phase2-intelligence-cases.json` pins queries, expected canonical paths and source SHA-256. Compare old/new body coverage at source level with explicit query executable SHA. Exact phrase queries remain unchanged; this diagnostic does not estimate general relevance or independent generalization.
- Source-level query deduplicates before limit, reports canonical source checksum and status separately from atom compatibility status. Missing scope/domain/supersession is unknown; no inferred authority or applicability. Readback remains required.
- Verify one-hop graph limit, edge type and deterministic order; depth is explicit, cycles never cause recursive expansion.
- Verify initial build, edit/add/delete/move, stale detection, missing metadata/index and deterministic rebuild. Build replacement is atomic so failed rebuild cannot destroy a usable cache. Canonical sources remain authoritative.
- Formal file-response watcher validates fixed requests and strict response schema, one attempt, explicit deadline, no fallback. Timeout/failure exposes an interrupt-required handoff to the owning agent adapter; it must never claim that recording a timeout stopped a child. Usage absent is unavailable, never zero. Test the handoff and failure branches without rerunning selector quality cases.

## Test-first and boundaries

Behavior tests precede implementation and must fail for the missing behavior. Full Go suite, binary parity and GitHub CLI workflow are release checks. Source changes require rebuilding all five binaries. Actual workflow applicability/adoption, project evidence R11/O03 and actual-token budgets remain Phase 3. Phase 2 is not declared complete by these engineering checks alone.
