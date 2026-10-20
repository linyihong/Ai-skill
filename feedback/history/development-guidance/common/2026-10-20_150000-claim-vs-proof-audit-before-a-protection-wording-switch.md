> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-20 - Audit every protective sentence against a proof before switching a security wording

Status: candidate

#### One-line Summary

Before a document changes from "trusted" to "protected", list every "cannot/refuses/prevents" sentence and map each to a test or live transcript; the independent reviewer should attack the wording, not only the code, and ask whether one credential alone can still defeat the claim.

#### Human Explanation

An acceptance run passed its live proofs and rewrote the docs to say that no Batch, releaser or API credential could alter or delete an archived month. The independent review used only the releaser credential, with values it could read from the seal, and dropped a confirmed month without any frozen copy; the daily reconcile and purge then also removed the other copy. The text was stronger than the executor's own capability table. Statement counts were also inflated by repeated rows. Two plan conditions (a skipped proof, and review before the switch) had not been met.

#### Trigger

- A phase ends by changing a public or contract statement about what an attacker cannot do.
- Evidence tables aggregate counts across classes.

#### Generalized Lesson

1. Build a claim-vs-proof table for the changed wording; unsupported sentences are corrected before the switch.
2. Test each credential class alone, not only the combination a deployment is expected to use.
3. Count what was actually run per class, and de-duplicate classes that share an identity.
4. A control that only detects (a verify command) needs an enforced ordering where destructive steps wait for it, or the wording must say the destructive step does not check.
5. Run the independent review before the wording change, as the plan says.

#### Evidence

- Tool: PostgreSQL role-separated archive tier, live acceptance with real processes
- Sanitized excerpt: single-credential release dropped the only cold copy; wording said it could not
- Evidence path: project-local acceptance evidence and verifier report

#### Agent Action

- Do not switch a protection claim until the claim-vs-proof table has no unsupported rows.
- Do not accept aggregate statement counts without per-class breakdown.

#### Goal / Action / Validation

- Goal: Public security wording never exceeds demonstrated protection.
- Action: claim-vs-proof audit plus single-credential attack before the switch.
- Validation: every protective sentence cites a test or transcript; the review attempted each credential alone.

#### Applies When

- Contract, runbook or UI text asserts that some actor cannot do something.

#### Does Not Apply When

- Internal design notes that state intent rather than guarantees.

#### Validation

The table exists, is reviewed independently, and corrections land before the wording switch is pushed.

#### Promotion Target

- Security-claim review checklist in the delivery workflow, if repeated.

#### Required Linked Updates

- Indexed in `feedback/history/development-guidance/common/README.md`.
- No workflow or enforcement change yet; candidate stays local until repeated.
