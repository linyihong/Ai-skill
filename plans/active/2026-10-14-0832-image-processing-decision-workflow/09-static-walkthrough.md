# 09 — Phase 1 Static Walkthrough（無模型）

← [`_plan.md`](_plan.md) · workflow [`workflow/image-processing/`](../../../workflow/image-processing/README.md)

## 目的

不依賴任何分割模型／供應商名，判定：**契約是否完整、約束如何裁可行集、驗收如何擋假成功**。

## Case A — 契約不完整

| 輸入 | 判定 |
| --- | --- |
| 有圖、intent=去背，但未指定 `background` | **停止／ask**；不得執行 |
| `fidelity` preset=`strict` 且未定 `max_rim_px` | **ask**；不得默認全域去污染 |
| 多主體、無 `subject_hint`、analysis confidence=low | **ask／human_review** |

契約欄位 SoT：[`image-task-contract.yaml`](../../../workflow/image-processing/contracts/image-task-contract.yaml)。

## Case B — Constraints 裁可行集

契約：`task_class=A`，`background=transparent`，`immutable=[teeth, gem_holes]`，`fidelity.subject_structure=preserve`，`edge_treatment=rim_decontaminate`，`max_rim_px=4`。

| Candidate | Feasible? | 理由 |
| --- | --- | --- |
| `seg.generic` | yes | 未排除 |
| `post.decontaminate` | yes | edge_treatment 授權 |
| `gen.edit` | **no** | subject preserve／非 allow_local_edit |
| interior binary opening | **no** | immutable + structure preserve |

Selection 只能在 yes 集合內選（[`candidates.yaml`](../../../workflow/image-processing/registry/candidates.yaml)）。

## Case C — Execute ≠ Verify

| 事實 | 允許 finality？ |
| --- | --- |
| Adapter 寫出 RGBA PNG，alpha 全 255 | **否** → `alpha_invalid`；對齊 `neg.alpha-opaque` |
| PNG + 棋盤預覽「看起來透明」 | **否** — 預覽不是證據 |
| 邊乾淨但牙齒被 opening 咬穿 | **否** → `interior_damage`（residue 可 pass） |
| 硬 check 全 pass，軟邊緣 `review_required` | `accepted_with_review` 僅可；非生產完成 |

## Case D — Consumer profile 不污染核心

`consumer_profile=slot-symbol`（328²、face-axis）附加為 **constraints**，不改變通用 A lifecycle 步驟名。
缺少 Sprite 參考 → 對齊相關 check = `inconclusive`／ask，不得硬猜。

## Gate 自檢（本 walkthrough）

- [x] 無模型名出現在「必須步驟」敘述中（僅 candidate id）
- [x] 可指出誰負責 Constraint vs Selection vs Verify
- [x] 假成功（opaque alpha／interior damage）路徑可擋

**Verdict**：Phase 1 Gate 靜態條件滿足；執行器實作屬 Phase 2。
