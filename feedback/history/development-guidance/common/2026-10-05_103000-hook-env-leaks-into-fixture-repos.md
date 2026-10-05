> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-05 - Hook environment leaks into fixture repositories

Status: candidate

#### One-line Summary

A gate's self-test that builds a throwaway git repo must strip `GIT_*` and the gate's own scope variables first; inside a hook they redirect or silence the fixture.

#### Human Explanation

Gate scripts often pin their semantics with a fixture test that runs `git init` in a temp directory and asserts on history. Run by hand it passes. Run from pre-commit it inherits the hook's environment: git may export `GIT_DIR` / `GIT_INDEX_FILE` pointing at the real repository, and the verify orchestrator exports its own switches (for example a staged-only scope flag). The first can make fixture `git config` / `commit` write into the real repo; the second silently disables the branch under test, so an assertion that "this must go red" fails only inside the hook — or, worse, a negative assertion passes vacuously.

#### Trigger

- A fixture test passes from the shell and fails only when the commit hook runs it.
- A gate has a mode flag set by the hook runner (staged scope, CI mode) and a fixture that exercises the other mode.

#### Evidence

- Tool: Node fixture test for a git-history freshness gate, run by a versioned pre-commit hook
- Sanitized excerpt: history assertion "code-only commit after measurement is stale" failed only under pre-commit because the hook runner exported the staged-scope flag; real repo checked afterwards and found untouched
- Evidence path: project-local `scripts/lib/test_*` gate fixtures

#### Generalized Lesson

1. Before any fixture git call, delete every `GIT_*` variable and every orchestrator flag the code under test reads.
2. Reproduce hook conditions by hand once: run the fixture with those variables set to hostile values.
3. Keep mode-dependent assertions explicit: set the flag inside the test for the case that needs it, then remove it.

#### Agent Action

- Do scrub `process.env` (or pass a clean `env`) in fixture tests that spawn git.
- Do run the fixture once with `GIT_INDEX_FILE=/nonexistent` and the scope flag set before committing.
- Do not trust "passes locally" for a test that a hook will run.

#### Goal / Action / Validation

- Goal: Gate fixtures behave the same by hand and inside hooks, and never touch the real repository.
- Action: Strip hook-inherited variables at the top of the fixture section.
- Validation: Fixture passes with hostile `GIT_*` / scope variables exported; real repo `git config --local` and `git log` unchanged afterwards.

#### Applies When

- Fixture or unit tests for git-aware gates, hooks, release scripts.

#### Does Not Apply When

- Tests that intentionally operate on the current repository's index (staged-scope gates themselves).

#### Validation

Export `GIT_INDEX_FILE=/nonexistent` and the orchestrator's scope flag, run the fixture, confirm it passes and the real repository is unchanged.

#### Promotion Target

- `workflow/software-delivery/` test-gate guidance, if a second project repeats the failure.

#### Required Linked Updates

- Indexed in `feedback/history/development-guidance/common/README.md`.
- No workflow or enforcement change yet; candidate stays local until repeated.
