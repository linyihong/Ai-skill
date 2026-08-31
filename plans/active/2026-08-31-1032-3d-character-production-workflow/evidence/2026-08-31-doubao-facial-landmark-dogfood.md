# Doubao facial landmark dogfood — consumer evidence（去敏）

**Run ID**：2026-08-31-doubao-facial-landmark
**Plan**：`2026-08-31-1032-3d-character-production-workflow`
**Maturity 判定**：**prototype**（非 runtime-ready；Phase 3 execution dogfood 七條 **未**授權）

## 觸及 stage

| Stage | 觀察 | Contract 欄位 |
| --- | --- | --- |
| Mesh QA | Tripo face plate（`root.9` 等）與 shell 分離；脖子/胸口 UV smear | `uv_material` → **fail** |
| Facial expression | Landmark 自動幾何偵測不可靠；screen-raycast 可對齊貼圖 | 新 gate：landmark alignment readback |
| Deformation | `aa` morph 在 mouth 區域有 mesh delta；贴图眼無像素變化 | `one_extreme_facial_expression` → **partial** |
| Export / readback | GLB 11 mesh；viewer 可驅動 morph | partial |

## 核心 lesson（promote 到 workflow）

1. **Generator class `tripo_face_plate`**：臉部貼圖在獨立 material plate，不是單一 head shell。Blender joined-mesh `vert_indices` 與 runtime split-mesh **不可互換**。
2. **Landmark SoT 雙軌**：runtime 用 `center_gltf`；Blender morph zone 用 `vert_indices`。
3. **Verify anti-pattern**：寬鬆 3D bbox 或 agent 宣稱 PASS **不等於**貼圖對齊。必須 **screen-space expected fraction + 截圖 + 人工確認**。
4. **有效流程**：face-plate screen bbox → raycast → 可选手動微調 → 存 landmark record → Playwright verify screenshot。
5. **`texture_baked_eyes`**：morph 有 delta 但 blink 無視覺 → 記 **known defect**，不得判 expression pass。

## Contract fill（本 run）

### mesh_qa_report

- `uv_material.result`: **fail**
- `uv_material.notes`: failure class `face_plate_neck_smear` — plate 材質出現在脖子/胸口且 UV 拉伸

### deformation_acceptance_set

- `one_extreme_facial_expression.result`: **partial**
- `notes`: mesh delta observed on `aa` near mouth landmark；visible eye pixels unchanged（`texture_baked_eyes`）

### landmark alignment（workflow 新增 gate；contract 待 Phase 4 profile）

- `landmark_source`: `screen_pick`（非 `blender_auto` z-band）
- 機械 verify：marker 螢幕座標 vs face-plate bbox 期望 fraction（dist ≤ 22% bbox）
- 人工 verify：**pending stakeholder**（agent 截圖不可單獨作 completion claim）

## Anti-patterns observed

| Anti-pattern | 後果 |
| --- | --- |
| Shell-only z-band landmark | 球落在脖子/胸口 |
| Whole-body screen bbox verify | 5/5 假 PASS |
| Agent PASS without pixel check | 使用者截圖推翻結論 |

## 與 plan phase 的關係

- **不是** Phase 3 execution dogfood（需口令「跑 execution dogfood」）。
- **是** Phase 2.5 延伸 + Phase 6 預演：證明 facial-expression stage 文件缺口與 mesh QA blocking 合理。
- Workflow 更新：[`facial-expression.md`](../../../../workflow/3d-character-production/facial-expression.md)

## Consumer evidence 位置

詳細路徑、腳本、截圖留在 consumer `<PROJECT_ROOT>`（VTuber / Doubao），本檔不含本機絕對路徑。
