# 3h5 — Completeness via Deferred Disposition

**Plan**: [`../../2026-07-14-0856-ui-pattern-knowledge-workflow.md`](../../2026-07-14-0856-ui-pattern-knowledge-workflow.md)  
**Date**: 2026-07-14  
**Mode**: experimental（停補 Pattern；從 Deferred 開始）

## Success criterion（鎖定）

**不是** Deferred → 0。  
**是**：Every Deferred has an **explicit disposition**.

Allowed dispositions:

| disposition | Meaning |
| --- | --- |
| `uncovered_pattern` | 候選新 Pattern；尚未建 Entry（Coverage ≠ Completeness） |
| `composition_only` | Existing Pattern + Placement／關係；**不**開新 Entry |
| `implementation_only` | 純實作細節，不進 Pattern Knowledge |
| `out_of_scope` | 本波次故意不收 |

Research question：很多 UI 不是 New Pattern，而是 **Existing Pattern + Placement**.

---

## Inventory（Episode Detail）

| Node | Before | Disposition | Verdict |
| --- | --- | --- | --- |
| `app_bar` | deferred / uncovered | **`uncovered_pattern`** | 結構族候選；非 sheet/toast placement |
| `player` | deferred / uncovered | **`uncovered_pattern`** | 媒體主表面；非 overlay composition |
| `floating_hint` | deferred / uncovered | **`composition_only`** of `toast` + `near_player_controls` | **非新 Pattern** — H5 核心發現 |

Zero Unknown. Deferred count stays **3**（不追求下降）。

---

## Floating Hint probe（H5 價值點）

問：這是新 Pattern，還是 Composition？

```text
floating_hint
    ?= new entry
    |  → NO (would inflate Coverage under Completeness pressure)
    ?= toast + placement
    |  → YES (composition_only)
```

| Check | Result |
| --- | --- |
| New Entry created? | **No** |
| `toast.yaml` modified? | **No**（Invariant） |
| Disposition explicit? | **Yes** — `composition_only` |
| Completeness（no Unknown）? | **Yes** |

---

## Metrics direction

| Metric | Role | Δ H4→H5 |
| --- | --- | --- |
| Entry Modifications | Primary | **0** |
| Rule Count | Supporting | 0 |
| Deferred Nodes | Supporting | 0（仍 3；已全部 disposition） |

---

## Verdict

**H5 mini-cycle：PASS** — Completeness = disposition coverage，不是 Deferred 消滅。  
Composition Knowledge remains Emerging；Pattern Knowledge untouched（Stable）。  
Next：H6 Trace Graph（含 deferred 如何標「waive Recipe」而不補 Entry）。
