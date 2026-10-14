# 12 — Phase 4 B／C／D fragments

← [`_plan.md`](_plan.md)

## 目標

在 **不重寫** shared decision loop 的前提下，為 B／C／D 各加契約片段、候選偏置、fixtures／驗收。

## 產物

| Class | Fragment | Acceptance 焦點 |
| --- | --- | --- |
| B | [`task-class-b.yaml`](../../../workflow/image-processing/contracts/task-class-b.yaml) | 柔邊／髮絲；常 `accepted_with_review` |
| C | [`task-class-c.yaml`](../../../workflow/image-processing/contracts/task-class-c.yaml) | 僅授權區可 edit；禁整圖重畫當去背 |
| D | [`task-class-d.yaml`](../../../workflow/image-processing/contracts/task-class-d.yaml) | canvas／規格；`pos.d-canvas-328` vs `neg.consumer-layout` |

Adapter：[`task-class-fragments.md`](../../../workflow/image-processing/adapters/task-class-fragments.md)

## Gate

```text
python3 workflow/image-processing/tools/run_phase4_gate.py --repo <AI_SKILL_REPO>
```

## Acceptance

- [x] 三片段存在且由 registry `fragment_index` 指向
- [x] `execution-flow` 僅加 `0b` lazy-load，主鏈 0–9 保留
- [x] D 正／負 consumer canvas 機械可驗證
- [x] 未新增第二套 pipeline／未註冊 route
