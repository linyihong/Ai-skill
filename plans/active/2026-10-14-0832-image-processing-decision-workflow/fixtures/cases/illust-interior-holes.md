---
case_id: illust-interior-holes
kind: reproducible_case
status: spec_frozen
input_hash: PENDING_BINARY
---

# Case: `illust-interior-holes`

## 契約摘要

| 欄位 | 值 |
| --- | --- |
| `task_class` | A |
| `background` | transparent |
| `immutable` | 牙齒／口腔陰影／配件孔洞 |
| `fidelity` | `subject_structure=preserve`；`subject_color=preserve`；`edge_treatment=rim_decontaminate`（`max_rim_px` ≤ 3） |
| `consumer_profile` | none |

## 輸入

- 路徑（預定）：`fixtures/inputs/illust-interior-holes.png`
- 特徵：主體內應保留的孔洞／牙縫（深色但屬主體）
- `input_hash`：`PENDING_BINARY`

## 期望

| 項 | 期望 |
| --- | --- |
| 對非 rim 區做 binary opening／暗像素填背景 | 硬 check 主體完整性或保真 → `fail`；`failure_class=interior_damage` 或 `fidelity_violation` |
| 正確路徑 | 孔洞 alpha 保持不透明主體；僅 rim 去暈 |
| 與 `neg.over-erode` 正交 | 本案例焦點是 **內部** 損傷，不是外輪廓過度侵蝕 |

## 重跑步驟

1. 載入契約 + 輸入；核對 hash。
2. 故意或回歸：若 pipeline 含 interior opening，verify 必須 fail。
3. 合法 candidate：無 interior morphology；IDR 記錄。
4. Independent verify；對照期望。
