# Legacy Git hook path

Canonical hook adapters live in [`.githooks/`](../../.githooks/) at the repository root.

After clone, run:

```bash
ai-skill hooks install
```

Files here forward to `.githooks/` for clones that still use `core.hooksPath=scripts/git-hooks`.
