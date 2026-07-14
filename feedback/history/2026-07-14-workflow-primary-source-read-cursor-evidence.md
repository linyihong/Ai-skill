# Observation — workflow primary_source gate may miss Cursor Read evidence

**Date**: 2026-07-14  
**Disposition**: candidate → enforcement / hooks  
**Symptom**: After conversation summary, agent Read `workflow/software-delivery/execution-flow.md` repeatedly but PreToolUse still blocked with `gate.workflow.primary_source_read` until a fresh user turn cleared/relatched transcript evidence.

**Hypothesis**: Cursor transcript tool-use shape / flush timing / tool name (`functions.Read` vs `read`) not always visible to `transcriptHasRequiredBootstrapReads` / `isTranscriptBootstrapReadTool` when route.software-delivery is locked.

**Action needed**: Reproduce with harness; consider adding `functions.read` to transcript read-tool allowlist and/or waiting for tool_result before evaluating non-Read PreToolUse. Do not weaken fail-open.

**Not**: this note is not Protocol Core change.
