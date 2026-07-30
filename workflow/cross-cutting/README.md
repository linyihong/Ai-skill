# Cross-Cutting Workflow Concerns

Concerns that span multiple delivery slices without becoming a new `software-delivery` slice.

## What belongs here

- Runtime models shared by journey, validation, UI contracts, and client code
- Pilot templates consumed by project overlays — **not** executable slice contracts
- Boundary tables that prevent scope drift into `experience-validation` or journey execution

## What does not belong here

- New `route.workflow.*` registry entries (until slice promotion criteria met)
- Runnable integration tests (project repos)
- BDD / Gherkin (project `docs/features/`)

## Current concerns

| Concern | Path | Status |
| --- | --- | --- |
| Experience runtime | [`experience-runtime/README.md`](experience-runtime/README.md) | pilot — player template only |
| Review (fault_finding consumer) | [`review/README.md`](review/README.md) | active — invoke + checklist; not stance contract owner |
| Decision Support Stage | [`decision-support/README.md`](decision-support/README.md) | pilot — generic contract + 1／3 converged case (`workflow/legal/strategy/`) |

## Slice promotion policy

Do **not** register `sd-experience-runtime` until at least **three** converged cases exist (plan: player + editor + onboarding). Until then, projects consume cross-cutting YAML as overlay alignment only.

The same three-case bar applies to **Decision Support Stage**: it stays a cross-cutting pattern that domains instantiate and link back to. Until three domains have converged instantiations (each supplying decision-point inventory, playbooks, verification source and a depth gate), it is **not** a registered stage in any workflow's required order and must not be described as a repo-wide capability. Current: 1／3 (`legal`).
