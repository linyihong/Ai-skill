# Evidence — Product locale-layered failure registry + JA／ID mechanical dogfood

Status: observed  
Date: 2026-09-29  
Linked plan: [`14-locale-layered-failure-registry.md`](../14-locale-layered-failure-registry.md)、ep8 Finality ([`13`](../13-ep8-contract-revision.md))

## Scope

Sanitized product dogfood（no live LLM、no host paths、no episode dump）showing:

1. Product failure registry split toward Ai-skill plan 14 shape: **core ∪ locale** + binding.
2. Mechanical Finality／Selection notes for JA truncated／kinship and ID honorific cases.

## Product registry shape（aligned）

```text
failure_registry_binding.json
     ├── translation_failure_patterns.json          # cross_locale_core
     ├── locale/ja-JP/failure_patterns.json         # JA-F05 → source_language_residue
     └── locale/id-ID/failure_patterns.json         # ID-F01 → foreign_honorific_leak
```

Loader: `_load_patterns()` = core only；`_load_patterns_for_locale(lang|locale)` = core ∪ first matching locale registry（aliases `product_detector` id for existing detectors）.

## Dogfood cases（mechanical）

| Case | Target | Result |
| --- | --- | --- |
| truncated `我替我姐向` | ja | Finality `blocked` / `source_truncated`（I21） |
| kinship residue `姐夫` left in JA dst | ja | Finality `needs_review`；pattern `kinship_target_residue`（JA-F05 manifestation） |
| kinship realized `義兄さん` | ja | Finality `accepted`；no kinship residue hit |
| structured `陈小姐` → `Nona Chen` | id | Finality `accepted` |
| `Ms. Chen` on id | id | Finality `needs_review`；`foreign_honorific_leak`（ID-F01） |

Registry merge smoke：ja load includes `JA-F05-kinship_target_residue`；id load includes `ID-F01-foreign_honorific_leak`.

## Plan impact

- Confirms plan 14 binding model is implementable outside Ai-skill docs（product JSON mirror）.
- Does **not** close Phase 4 runtime route；does **not** replace human review of ep8 walkthrough.
- Open：broader JA-F01–F11／ID expansion still evidence-driven；full episode LLM dogfood still pending.

## Sanitization

No Windows paths、SSH hosts、API keys、or raw episode transcripts in this file.
