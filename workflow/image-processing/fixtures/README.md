# Fixtures — image-processing

Canonical home for reusable desensitized / synthetic fixtures (Phase 1+).

## Current staging

Phase 0 freeze specs live under the active plan:

[`plans/active/2026-10-14-0832-image-processing-decision-workflow/fixtures/`](../../plans/active/2026-10-14-0832-image-processing-decision-workflow/fixtures/README.md)

Migrate case specs + binaries here when inputs are desensitized and hashed.
Consumer-private assets never land in this directory.

## Layout (target)

```text
fixtures/
  README.md
  cases/          # optional copies or stubs linking to plan until moved
  inputs/         # png + sidecars; git-lfs or small synthetics only
```

## Negative fixtures (Phase 2 Gate)

| Fixture | Path | Purpose |
| --- | --- | --- |
| `neg.alpha-opaque` | [`inputs/neg.alpha-opaque.png`](inputs/neg.alpha-opaque.png) | Alpha present but fully opaque |
| `neg.residue-halo` | [`inputs/neg.residue-halo.png`](inputs/neg.residue-halo.png) | Navy rim residue |

Generate / re-hash：`python3 workflow/image-processing/tools/generate_negative_fixtures.py`
Gate：`python3 workflow/image-processing/tools/run_phase2_gate.py` → [`reports/phase2-gate-latest.json`](reports/phase2-gate-latest.json)
