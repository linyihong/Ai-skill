# Facial Expression

Body deformation pass 後，填
[`records/facial-expression-acceptance.yaml`](records/facial-expression-acceptance.yaml)。
只有 `decision == pass` 可進 outfit／motion；`partial` 是未完成證據，不是較弱的 PASS。

VRM blendshape 清單屬 Phase 4 profile，不在本檔展開。

## Preconditions

| Gate | 條件 |
| --- | --- |
| Identity | `decision == accepted` 且 `validity == current`（見 identity-acceptance） |
| Mesh QA | geometry decision pass；`uv_material.result == pass`，或具名 deferral 且 maturity 封頂 prototype |
| Deformation | `deformation_acceptance_set.decision == pass`（body minimal set 已觀察） |

## Control-mapping gate

在驅動 morph／blendshape 前，必須建立可讀回的 control mapping。Mapping 方法可為 landmark、
joint、curve 或 profile 定義的其他方法；core 不保存特定 DCC、generator 或 mesh index 規則。

最低要求：

- mapping 明確綁定 source／target asset revision。
- 機械 readback 與人工可見 evidence 都對同一 revision。
- source 與 runtime mesh 結構不同時，必須有轉換／重建證據；不得假設 index 可直接沿用。
- mapping 缺失或只有單側證據時填 `partial` 或 `fail`，不得填 pass。

## Expression observation

至少觀察一個 extreme facial expression。每項同時記錄：

- real asset geometry readback。
- visible／runtime readback。
- neutral baseline 是否保持。

任一側只有部分證據時，observation 與 stage decision 都是 `partial`；detached helper、
screen overlay、export success 或只有 morph delta 均不得升成 pass。
