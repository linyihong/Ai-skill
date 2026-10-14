# Adapter notes — task-class B／C／D fragments

← [`README.md`](README.md)

Fragments live under [`../contracts/`](../contracts/). They **do not** own a second lifecycle.

| Class | Fragment | Typical candidates | Fixtures |
| --- | --- | --- | --- |
| B | [`task-class-b.yaml`](../contracts/task-class-b.yaml) | `seg.portrait_matting`, `post.matte_refine` | soft-fringe proxy；`neg.over-erode` |
| C | [`task-class-c.yaml`](../contracts/task-class-c.yaml) | `gen.edit`（需 edit_region） | pattern via `neg.interior-damage` |
| D | [`task-class-d.yaml`](../contracts/task-class-d.yaml) | `post.deterministic`／`post.canvas_place` | `neg.consumer-layout`／`pos.d-canvas-328` |

## Load rule

1. Bind core [`image-task-contract.yaml`](../contracts/image-task-contract.yaml).
2. If `primary` ∈ {B,C,D}（或 secondary 含 D），merge fragment constraints.
3. Feasible set = registry `task_class_load[primary]` ∩ license ∩ fidelity.
4. Continue shared steps 1–9 in [`execution-flow.md`](../execution-flow.md).
