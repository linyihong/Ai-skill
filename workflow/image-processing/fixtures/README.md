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

## Negative fixtures

| Fixture | Path | Purpose |
| --- | --- | --- |
| `neg.alpha-opaque` | [`inputs/neg.alpha-opaque.png`](inputs/neg.alpha-opaque.png) | Alpha present but fully opaque |
| `neg.residue-halo` | [`inputs/neg.residue-halo.png`](inputs/neg.residue-halo.png) | Navy rim residue |
| `neg.interior-damage` | [`inputs/neg.interior-damage.png`](inputs/neg.interior-damage.png) + [`inputs/refs/`](inputs/refs/) | Interior punched |
| `neg.over-erode` | [`inputs/neg.over-erode.png`](inputs/neg.over-erode.png) + ref | Subject over-eroded |
| `neg.consumer-layout` | [`inputs/neg.consumer-layout.png`](inputs/neg.consumer-layout.png) | Wrong canvas vs profile |
| `pos.d-canvas-328` | [`inputs/pos.d-canvas-328.png`](inputs/pos.d-canvas-328.png) | D positive：328² + true alpha |

## Dogfood（synthetic proxies）

| Path | Purpose |
| --- | --- |
| [`dogfood/sources/`](dogfood/sources/) | Category proxies（object／glow／ui／multi／soft） |
| [`dogfood/outputs/`](dogfood/outputs/) | rembg outputs |
| [`dogfood/reports/`](dogfood/reports/) | IDR + verify JSON |

Generate：`generate_negative_fixtures.py` / `generate_dogfood_sources.py`
Gates：`run_phase2_gate.py` · `run_phase3_gate.py` → [`reports/`](reports/)
