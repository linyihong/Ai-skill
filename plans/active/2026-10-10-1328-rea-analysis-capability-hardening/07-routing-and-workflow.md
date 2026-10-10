---
id: 2026-10-10-1328-rea-07-routing-and-workflow
plan_kind: sub
status: completed
owner: larrylin/cursor-session
created: 2026-10-10
parent: 2026-10-10-1328-rea-analysis-capability-hardening
required_for_completion: true
sub_plan_reason: >
  Runtime integration requires registry routes with named consumers plus
  workflow/reverse-engineering execution entry and artifact gates.
---

# Sub-plan 07 — routing + workflow

Parent: [`_plan.md`](_plan.md) Phases 7–8.

## Acceptance

- [x] `route.analysis.*` / `route.workflow.reverse-engineering` 有 discovery consumer
- [x] `workflow/reverse-engineering/` execution-flow + gates
- [x] 授權第一步；不削弱 apk-analysis 動態主線
