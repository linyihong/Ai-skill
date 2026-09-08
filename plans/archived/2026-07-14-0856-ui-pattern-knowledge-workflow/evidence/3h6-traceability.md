# 3h6 — Traceability（可追溯 + 可停止）

**Plan**: [`../_plan.md`](../_plan.md)  
**Date**: 2026-07-14  
**Mode**: experimental Trace Graph  
**Deliverable**: Trace Graph with first-class **Waiver**（非 `recipe: null`）

## Traceability redefined

不是「每個節點都有 Recipe」。  
而是：**每個節點都必須有合法且可解釋的終點** — `complete` 或 `waived`。

```text
recipe: null          → 不知道（非法）
trace_status: waived  → 知道，且刻意停止（合法）
```

---

## Exit Gate

| ID | Rule | Result |
| --- | --- | --- |
| **T1** Continuous Trace | 非 Deferred：Screen → Pattern → Selection → Recipe | ✅ |
| **T2** Governed Stop | Deferred 必有 `disposition` + `waive_reason` + `owner` | ✅ |
| **T3** No Broken Edge | 終點只能是 Recipe **或** Waived；無 `?????` | ✅ |

---

## Path A — 完整 Trace（例：Bottom Sheet）

```text
Episode Detail
    │
    ▼
Bottom Sheet
    │
    ▼
Selection Rule
  (entries/bottom_sheet.yaml#selection_rules)
    │
    ▼
Implementation Recipe
  (partial: portal, scrim, focus_trap, escape_close, body_scroll_lock)
```

同形可追：`modal_dialog` · `toast` · `scrim`（皆 `trace_status: complete`）。

---

## Path B — Waived Trace（例：Floating Hint）

```text
Episode Detail
    │
    ▼
Floating Hint
    │
    ▼
Disposition: composition_only
  composition_of: toast
  placement: near_player_controls
    │
    ▼
Waive Recipe
  waive_reason: composition_only_resolves_to_toast_placement
  owner: composition
```

可選延續（非必須）：從 `toast` 再走 Path A（H5 邊界：不開 floating_hint Entry）。

`app_bar` / `player`：

```text
… → Disposition: uncovered_pattern
  → Waive Recipe
    waive_reason: uncovered_pattern_no_recipe_yet
    owner: composition
```

---

## Full Screen Trace Graph（終點表）

| Node | Path | `trace_status` | Terminal |
| --- | --- | --- | --- |
| bottom_sheet | A | complete | Recipe（partial） |
| modal_dialog | A | complete | Recipe（partial） |
| toast | A | complete | Recipe（partial） |
| scrim | A | complete | Recipe（partial） |
| floating_hint | B | waived | Governed waiver |
| app_bar | B | waived | Governed waiver |
| player | B | waived | Governed waiver |

**Broken edges**: 0  
**Entry Modifications**: 0（Invariant）  
**No Unknown / ?????**: ✅

---

## Relation to H5

H5 證明 Composition 能 **拒絕新增 Pattern**（`floating_hint` → composition_only）。  
H6 證明該拒絕不是黑洞：Waiver 是一等公民終點。

---

## Verdict

| Check | Result |
| --- | --- |
| T1 | PASS |
| T2 | PASS |
| T3 | PASS |
| H6 | **PASS** |
| Entry Mods | **0** |

**Composition Knowledge**：H4 Independence + H5 Completeness + H6 Traceability → **Stable Candidate → Stable**（見 plan maturity）。  
**Composition Closure**：可宣告（三假說齊 + Entry Mods=0）。

Protocol sentence exercised：*Every trace must terminate explicitly…*
