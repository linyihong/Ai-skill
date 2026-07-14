# Phase 3 Composition Metrics

**Screen**: `episode_detail`  
**Invariant**: Entry Modifications **must stay 0**.

| Metric | Baseline (start) | After H4 stress (2026-07-14) | Expected at Closure |
| --- | --- | --- | --- |
| Deferred Nodes | 2 | **3**（+`floating_hint`） | ≥ honest count |
| Composition Rule Count | 4 | **5**（+`overlay.no_concurrent_temporary_overlays`） | ≥ honest count |
| Entry Modifications | 0 | **0** | **0** |

## Recount

```bash
grep -c 'knowledge: deferred' workflow/software-delivery/ui-pattern-knowledge/compositions/episode_detail.yaml
grep -c '^  - id:' workflow/software-delivery/ui-pattern-knowledge/composition_rules.yaml
git status -- workflow/software-delivery/ui-pattern-knowledge/entries/
```

## Notes

- Case A 增加 Rule Count 是 **正向**（Invariant 施壓後成立），不是 KPI 膨脹。
- Deferred 增加來自 Case D 誠實記帳，不是 Coverage 推進。
- Evidence：[`3h4-independence-stress.md`](3h4-independence-stress.md)
