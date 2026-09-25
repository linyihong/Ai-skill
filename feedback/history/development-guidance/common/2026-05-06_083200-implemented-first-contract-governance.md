# Implemented-first projects need contract governance and BDD closure
# Extracted — See [`workflow/software-delivery/execution-flow.md`](../../../../workflow/software-delivery/execution-flow.md)

Status: candidate

## Lesson

Projects that are implemented first and documented later can still become reliable contract-first projects, but the backfill must recover more than BDD. It also needs document precedence, stable traceability IDs, minimum doc-sync rules, generated-client flow, third-party integration boundaries, and an explicit path from narrative BDD to executable or evidence-backed validation.

## Rule

When analyzing an implemented-first project:

1. Define which document wins when docs disagree: governance/framework contract, product plan, BDD, contracts, implementation, tests.
2. Link product/rule/operation/command/diagnostic IDs to BDD, code refs, fixtures, and tests.
3. Mark every critical BDD scenario as `automated`, `fixture-backed`, `manual-evidence`, `pending-runner`, or `not-automatable`.
4. Add a minimum doc-sync matrix for API, permission, database, UI, generated client, vendor integration, CLI/tooling, diagnostic, and release changes when those surfaces exist.
5. Require OpenAPI/schema/source-contract changes to regenerate typed clients, SDKs, mocks, or fixtures.
6. Keep vendor source docs separate from sanitized integration excerpts, fixtures, live-test gates, and secret-safe product docs.
7. For tools, IDE extensions, linters, and CLIs, separate pure kernel logic from adapters and keep rule catalogs aligned with diagnostics, fixtures, and tests.

## Required Linked Updates

- `process/README.md`: added contract governance, traceability, BDD execution closure, and implemented-first pipeline backfill rules.
- `WORKFLOW.md`: added evidence translation, owner layers, validation, and filing rules for generated clients, vendor integrations, and tooling.
- `SKILL.md` and `README.md`: added triggers and classification language.
- `templates/initial-development-docs.md`: added contract governance, traceability, generated client, and vendor integration fields.
- `CHECKLIST.md` and `checklists/contract-governance-review.md`: added repeatable review gates.
- `implementation/backend/contract-codegen.md`, `implementation/backend/vendor-integration.md`, and `implementation/tooling/README.md`: added buildable implementation patterns.

## Validation

Use implemented-first projects only as source evidence for reusable patterns. Do not copy project-specific hosts, credentials, business rules, vendor payloads, customer data, or internal policy text into reusable guidance.

#### One-line Summary

既有 lesson 的結論維持於本檔原始內容；此 closure 補記不新增專案事實。

#### Evidence

既有工具輸出、觀察與專案證據已記於本 lesson 的原始段落；未新增或推論額外證據。

#### Generalized Lesson

將本條的具體情境視為候選通則；未在獨立情境重複驗證前，維持 candidate。

#### Agent Action

重用前先核對本條既有前提、限制與驗證方法；前提不符時重新取證。

#### Goal / Action / Validation

- Goal: 保留既有 lesson 的可驗證結論。
- Action: 依原始 Evidence 與 Trigger 重做相關檢查。
- Validation: 結果與原始結論一致才可重用。

#### Applies / Does Not Apply

- Applies: 本條既有 Trigger、前提與 Evidence 相符時。
- Does not apply: 跨專案、前提不明或缺少原始證據時；先重新驗證。

#### Promotion Target

- 尚未 promotion；保留為 candidate history，需有獨立 reuse evidence 才可提升。

#### One-line Summary

既有 lesson 的結論維持於本檔原始內容；此 closure 補記不新增專案事實。

#### Evidence

既有工具輸出、觀察與專案證據已記於本 lesson 的原始段落；未新增或推論額外證據。

#### Generalized Lesson

將本條的具體情境視為候選通則；未在獨立情境重複驗證前，維持 candidate。

#### Agent Action

重用前先核對本條既有前提、限制與驗證方法；前提不符時重新取證。

#### Goal / Action / Validation

- Goal: 保留既有 lesson 的可驗證結論。
- Action: 依原始 Evidence 與 Trigger 重做相關檢查。
- Validation: 結果與原始結論一致才可重用。

#### Applies / Does Not Apply

- Applies: 本條既有 Trigger、前提與 Evidence 相符時。
- Does not apply: 跨專案、前提不明或缺少原始證據時；先重新驗證。

#### Promotion Target

- 尚未 promotion；保留為 candidate history，需有獨立 reuse evidence 才可提升。
