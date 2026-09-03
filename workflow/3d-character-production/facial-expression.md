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
- `artifact_state`：該表情資料是否真的存在於匯出物中。
- `consumer_readback`：使用者實際操作的介面上是否看得到（無互動消費面時填 `not_applicable`）。

任一側只有部分證據時，observation 與 stage decision 都是 `partial`；detached helper、
screen overlay、export success 或只有 morph delta 均不得升成 pass。

## Negative readback

`fail` 是對資產本身的判斷，必須先排除另外兩種成因，否則只能填 `partial`：

- **未產出**：產出步驟含 opt-in 開關時，記錄旗標實際取值，而非假設預設值。空的資料項與
  失敗的資料項在 record 上長得一樣，但修復位置不同。
- **探針失明**：自製解析器須覆蓋容器格式的壓縮與稀疏分支，並先在已知正例讀到訊號
  （`probe_calibration.known_positive_observed == true`），其 0 才可作為證據。

## Consumer-surface readback

資產最終由互動 viewer／runtime 消費時，離線算圖只是資產層證據，不能替代消費層證據。
消費面驗收須驅動真實控制項，逐項斷言可觀測效果與回復基準態，並確認：

- 預設視圖未被除錯輔助物遮蔽，輔助物顯示狀態與其開關一致。
- blocking 由對應 record 欄位逐項推導；硬編碼的全域 block 會在缺陷修復後繼續阻擋。
