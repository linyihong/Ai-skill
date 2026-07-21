# Observation — workflow primary_source gate may miss Cursor Read evidence

**Date**: 2026-07-14（updated 2026-07-21）  
**Disposition**: validated → hooks fix  
**Symptom**: Agent Read `workflow/software-delivery/execution-flow.md` repeatedly but PreToolUse still blocked with `gate.workflow.primary_source_read` until a later user turn / full transcript flush.

## Confirmed root cause (2026-07-21, kaizenwms session `3f5268f2`)

1. **Cursor agent-transcript JSONL flush lag（主因）**  
   Within one user turn, completed `Read` tool_use rows are often **not yet on disk** when the next `Write`/`Shell` PreToolUse runs. Truncating the live transcript to the user message (before the flushed Read) reproduces `block=true`; including the Read line clears it.

2. **Occasional wrong `transcript_path`**  
   Hook DIAG showed PreToolUse for the active session sometimes receiving a **sibling** session’s jsonl (e.g. `8606518c` while working in `3f5268f2`).

3. **Secondary**: transcript tool-name allowlist historically omitted `functions.read` (still added for defense-in-depth). Live Cursor transcripts in this incident used `"name":"Read"`, so allowlist alone was not the blocker.

## Fix

- PreToolUse(Read) records the file path into a **project-scoped side-channel** (`/tmp/ai-skill-wf-ps-<projectHash>.seen`) immediately.
- `workflowPrimarySourceGate` accepts that side-channel when transcript evidence is missing (flush race).
- `isTranscriptBootstrapReadTool` / Cursor `preToolUseReadAllowed` include `functions.read`.
- BLOCK_WORKFLOW denials are also appended to `/tmp/ai-skill-bootstrap-hook.log`.

## Action needed

Keep harness coverage in `workflow_gate_test.go` (`TestWorkflowGate_SideChannelClearsFlushRace`). Do not weaken fail-open.

**Not**: this note is not Protocol Core change.
