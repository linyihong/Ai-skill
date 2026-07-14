# Phase 3 Composition Metrics（baseline）

**Date**: 2026-07-14  
**Screen under test**: `episode_detail`  
**Invariant check**: Entry Modifications **must stay 0** through Closure.

| Metric | Baseline | Expected at Closure | Source |
| --- | --- | --- | --- |
| Deferred Nodes | **2** | ≥ baseline（誠實增加可接受） | `app_bar`, `player` in `compositions/episode_detail.yaml` |
| Composition Rule Count | **4** | ≥ baseline | `composition_rules.yaml` `constraints:` |
| Entry Modifications | **0** | **0** | `git diff` on `ui-pattern-knowledge/entries/` during Phase 3 |

## How to recount

```bash
# Deferred Nodes
rg -c 'knowledge: deferred' workflow/software-delivery/ui-pattern-knowledge/compositions/episode_detail.yaml

# Composition Rule Count
rg -c '^- id:' workflow/software-delivery/ui-pattern-knowledge/composition_rules.yaml

# Entry Modifications (Phase 3 window — expect empty)
git log --oneline -- workflow/software-delivery/ui-pattern-knowledge/entries/
```

## Notes

- Metrics 不是 KPI；不優化「減少 Deferred」。
- `Entry Modifications ≠ 0` → Composition Closure **FAIL**（Invariant 失敗）。
