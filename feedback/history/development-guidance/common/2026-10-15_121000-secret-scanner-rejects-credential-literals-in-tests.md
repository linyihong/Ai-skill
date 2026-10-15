> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-15 - Secret scanners reject credential-looking literals in tests

Status: candidate

#### One-line Summary

Repository secret scanners flag variable names such as `password = "..."` and `Password=` literals in test code; build throwaway-database connection strings from a builder and neutral names instead of fighting the gate at commit time.

#### Human Explanation

Live tests that create throwaway logins need connection strings. Literal `Password=...` fragments and `password` variables tripped the pre-commit scanner for several unrelated slices, each costing a fix round. The gate is correct to be blunt; the fix is the way the test is written.

#### Generalized Lesson

1. Generate the secret at runtime (random) and pass it through a connection-string builder property rather than a literal fragment.
2. Use neutral variable names (e.g. `credential`, `PSW=` keyword form where the driver allows) in test helpers, and keep one shared helper so the pattern is not repeated.
3. Never whitelist the scanner for test directories to make this pass.

#### Agent Action

- Do check how existing tests in the repo obtain throwaway credentials before writing new ones.
- Do not commit literal passwords, even for local containers.

#### Goal / Action / Validation

- Goal: Test credentials never look like committed secrets.
- Action: runtime-generated values via a builder and a shared helper.
- Validation: pre-commit scanner passes without exceptions.

#### Evidence

- Tool: repository pre-commit secret scanner
- Sanitized excerpt: commits rejected for a variable named like a password assigned a literal, and for a `Password=` fragment in a test connection string
- Evidence path: project-local executor reports for several test slices

#### Applies When

- Tests that create throwaway logins or databases and need connection strings.

#### Does Not Apply When

- Pure configuration documents that carry no credentials.

#### Validation

The pre-commit scanner passes without allow-list changes and the test still authenticates with the runtime value.

#### Promotion Target

- Test-strategy guidance for throwaway database fixtures, if repeated.

#### Required Linked Updates

- Indexed in `feedback/history/development-guidance/common/README.md`.
- No workflow or enforcement change yet; candidate stays local until repeated.
