# Facial Expression

僅當 `deformation_acceptance_set.decision == pass` **且** mesh QA 未 blocking（`uv_material != fail`，或已 defer 並列 known defect）。
本 stage 不定義 identity 規則。表情證據進入 runtime-ready pack 的
`expressions`／`expression_trigger` 欄（見 records）。

VRM blendshape 清單屬 Phase 4 profile，不在本檔展開。

## Preconditions

| Gate | 條件 |
| --- | --- |
| Identity | `decision == accepted` 且 `validity == current`（見 identity-acceptance） |
| Mesh QA | `uv_material` 不得 silent pass；face/neck smear 須 fail 或 defer+classified |
| Deformation | `deformation_acceptance_set.pass`（minimal set 已觀察） |

## Landmark alignment gate（blocking for expression driving）

在驅動 morph / blendshape **之前**，必須完成 landmark readback。適用所有 generator；以下 failure class 來自 consumer dogfood（2026-08-31）。

### Generator class: `tripo_face_plate`

- 臉部貼圖在 **face plate material**（多 mesh / 多 material），非單一 head shell。
- Blender joined mesh 的 `vert_indices` **不能**直接映射到 runtime GLB split mesh。
- **Landmark record 雙軌**：`center_gltf`（runtime viewer）+ `vert_indices`（DCC morph zone，若存在）。

### 建議流程（工具中立）

1. 定義五點（眼 L/R、鼻、口、眉）的 **screen-space 期望位置**（相對 forward face-plate 螢幕 bbox 的 u/v fraction）。
2. 自期望螢幕座標 **raycast** 到 forward face meshes，寫入 landmark record。
3. 允許 **人工點選微調** 後覆寫 record。
4. **機械 verify**：marker 投影 vs 期望 fraction（不可用整身 bbox 代替 face-plate bbox）。
5. **人工 verify**：截圖確認球落在 **贴图特征** 上；agent 不得單獨宣稱 PASS。
6. 通過後才允許 morph / blendshape 驅動測試與 deformation set 觀察。

### Landmark record（draft shape）

| 欄位 | 用途 |
| --- | --- |
| `center_gltf` | `[x, y, z]` runtime / glTF Y-up |
| `landmark_source` | `screen_pick` \| `blender_auto` |
| `calibrated_at` | ISO8601 |
| `verify_screenshot` | 可選；consumer repo 路徑或 artifact ref |

## Expression observation

### Known defect: `texture_baked_eyes`

- 症狀：mesh morph 有 delta，但 eye **像素**不變（贴图烘焙眼）。
- 判定：`one_extreme_facial_expression` 可記 **partial**（mesh）而非 pass（visible）。
- 修復：Riggle / 手雕眼皮 / 換 mesh；不得用 screen overlay 假 blink。

## Anti-patterns

- 只用 DCC 幾何 z-band 或 shell centroid 當 landmark SoT。
- 將 Blender `vert_indices` 直接索引到 runtime 最大 submesh。
- 無截圖、無人工確認的機械 PASS。
- 將 blink 視覺失敗忽略為 expression pass。

## Related

- Mesh QA failure class：`face_plate_neck_smear`（見 plan contract notes）
- Plan evidence：[`2026-08-31-doubao-facial-landmark-dogfood.md`](../../plans/active/2026-08-31-1032-3d-character-production-workflow/evidence/2026-08-31-doubao-facial-landmark-dogfood.md)
