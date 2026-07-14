# 3h4 — Independence Stress Run（Episode Detail）

**Plan**: [`../../2026-07-14-0856-ui-pattern-knowledge-workflow.md`](../../2026-07-14-0856-ui-pattern-knowledge-workflow.md)  
**Date**: 2026-07-14  
**Mode**: experimental execution（停架構設計）  
**Success criterion（本輪鎖定）**: ≥1 real case that **requires new `composition_rules.yaml` constraint** and **zero `entries/*.yaml` edits**.

**Invariant pointer**: Composition evidence MUST NOT modify validated entries.  
**Workflow gate line** (separate research): [`feedback/history/2026-07-14-workflow-primary-source-read-cursor-evidence.md`](../../../../feedback/history/2026-07-14-workflow-primary-source-read-cursor-evidence.md) — pointer only；不併入本 plan 正文。

---

## Setup

| Item | Value |
| --- | --- |
| Screen | `episode_detail` |
| Frozen entries | `bottom_sheet` / `modal_dialog` / `toast` / `scrim` / `drawer`（not in tree） |
| Pre-stress Rule Count | 4 |
| Pre-stress Deferred Nodes | 2 |
| Pre-stress Entry Modifications | 0 |

---

## Case A — Overlay 互換 / 並存（**首輪故意失敗點**）

### Stress

1. Intent fixed: `share_or_export`（本來選 Bottom Sheet）  
2. Tree 允許把 slot **改寫成 Modal Dialog**（錯誤互換）  
3. 另：`bottom_sheet` 與 `modal_dialog` 在 Episode 皆 optional，**無**「不可同時 open」約束

### Observe

| Question | Result |
| --- | --- |
| Selection（share）是否應改變？ | **是** — 應保持 `bottom_sheet`；換成 Dialog 是選型錯，不是「Entry 規則變了」 |
| Sheet / Dialog entry `selection_rules` 被改了嗎？ | **否**（Entry Modifications = 0） |
| 既有 Constraints 能否擋住「兩者同時 open」？ | **否** — gap（首輪 **FAIL**） |

### Failure (valuable)

壓力開出第一個缺口：Episode 同時掛 Sheet + Dialog 為 optional，卻沒有 **concurrent open** Composition Constraint。  
若回頭改 `bottom_sheet.yaml` / `modal_dialog.yaml` 來「禁止彼此」→ 違反 Invariant。  
正確處置 → **新增 Constraint**。

### Repair（Constraint only）

Added:

```text
overlay.no_concurrent_temporary_overlays
  cannot_coexist_open: [bottom_sheet, modal_dialog]
```

### Re-check

| Check | Result |
| --- | --- |
| New rule exists | ✅ |
| Entry diff | ✅ empty |
| Share → Sheet 邊界未改寫（Independence） | ✅ |

**Case A verdict**: stress-FAIL → Constraint-add → PASS · **Invariant held under pressure**（本輪高價值證據）

---

## Case B — Feedback 插入 / 移除 Toast

### Stress

Episode + Bottom Sheet → insert Toast → remove Toast。

### Observe

| Question | Result |
| --- | --- |
| Overlay Selection Boundary（Sheet）變了嗎？ | **否** — `selection_rules` / neighbors 未改 |
| Toast family 是否進出 Overlay Decision？ | **否** — `feedback.toast_cannot_trigger_selection` 已覆蓋 |
| Metrics | Deferred 不變；Rules 本輪 +1（來自 Case A）；Entry Mods = 0 |

**Case B verdict**: PASS（既有 Constraint 足夠；Independence 穩）

---

## Case C — Scrim 消失（Dialog 無 Scrim）

### Stress

Modal Dialog open，Tree / 狀態刻意去掉 Scrim。

### Observe

| Question | Result |
| --- | --- |
| Constraint 能偵測？ | **是** — `overlay.dialog_requires_scrim` |
| 需要改 `modal_dialog.yaml`？ | **否**（Invariant） |
| 需要新 Constraint？ | **否**（既有規則已定位） |

**Case C verdict**: PASS · Invariant exercised（檢測靠 Constraint，不靠回改 Entry）

---

## Case D — Deferred Node（Floating Hint）

### Stress

加入:

```yaml
pattern: floating_hint
knowledge: deferred
reason: uncovered_pattern
```

### Observe

| Question | Result |
| --- | --- |
| H5 Completeness（零 Unknown）？ | **仍成立** — deferred 明確 |
| Overlay Independence？ | **不受 Deferred 污染** — 未改任何 overlay entry |
| Traceability 到 Recipe？ | **故意不足** — deferred 無 Recipe；記給 **H6**（非本輪補 Entry） |

**Case D verdict**: H4 PASS；Deferred Nodes 2→3；H6 debt recorded（不開新 Entry）

---

## Metrics after H4 stress

| Metric | Before | After |
| --- | --- | --- |
| Deferred Nodes | 2 | **3**（+`floating_hint`） |
| Composition Rule Count | 4 | **5**（+`overlay.no_concurrent_temporary_overlays`） |
| Entry Modifications | 0 | **0** |

---

## Gate vs Success criterion

| Criterion | Result |
| --- | --- |
| ≥1 case needing **new** composition rule, **no** entry edit | ✅ **Case A** |
| H4 Independence stress-validated（非「沒碰到」） | ✅ |
| Wish that ≥1 case fails once | ✅ Case A initial gap |

**H4 mini-cycle status**: **PASS（stress-validated）** — proceed to H5 without new design concepts.
