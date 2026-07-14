# Phase 3 Composition Metrics

**Screen**: `episode_detail`  
**Primary Metric**: **Entry Modifications**（Invariant）。Supporting：Δ Deferred、Δ Rules。

看 **方向**，不把 3/5/0 當 KPI。

## Snapshots + Δ

| Metric | Role | Start | After H4 | After H5 | Δ（H4→H5） | Signal |
| --- | --- | --- | --- | --- | --- | --- |
| **Entry Modifications** | **Primary** | 0 | **0** | **0** | **0** | Composition Complexity↑ 時仍須 =0；→1 = Invariant 失守 |
| Composition Rule Count | Supporting | 4 | 5 | 5 | 0 | 可 5→8→12；表示 Constraint Accumulation |
| Deferred Nodes | Supporting | 2 | 3 | 3 | 0 | 可不降；須 **every deferred has disposition** |

## Reading

```text
Composition Complexity ↑
        │
        ▼
Entry Stability still = 0   ← 唯一不可破
```

Rules↑ 與 honest Deferred↑ 都可接受。唯 Entry Mods≠0 → Composition Closure FAIL。

## Recount

```bash
grep -c 'knowledge: deferred' workflow/software-delivery/ui-pattern-knowledge/compositions/episode_detail.yaml
grep -c '^  - id:' workflow/software-delivery/ui-pattern-knowledge/composition_rules.yaml
git status -- workflow/software-delivery/ui-pattern-knowledge/entries/
```

Evidence：[`3h4-independence-stress.md`](3h4-independence-stress.md) · [`3h5-completeness-disposition.md`](3h5-completeness-disposition.md)
