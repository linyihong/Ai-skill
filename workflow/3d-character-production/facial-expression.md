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

## Texture-baked feature articulation

若眼睛、嘴型或其他應動特徵實際烘在貼圖裡，先確認 representation，再決定修復 owner；不得
因為概念上「應該有幾何」就直接切網格、加 detached helper，或把未裁切的 2D 圖層疊到表面。

可接受的 surface-driven adaptation 依序執行：

1. **建立 neutral parity**：保存原始 consumer 畫面；任何重組後的正中／靜止狀態須與基準
   一致，不得用「看起來差不多」代替 readback。
2. **量測表示與遮擋**：從實際 texture/UV/mesh mapping 找出可動特徵、開口／邊界與局部
   座標；色彩分類須先取樣，不能用未校準的固定亮度閾值。
3. **分離 foreground / background / occlusion**：可動內容、被其蓋住的底圖、以及限制其
   可見範圍的遮罩分開產出。沒有遮罩的平貼移動會越過眼皮、嘴唇或材質邊界。
4. **在局部域內補底圖**：只用確認屬於底圖的像素作 inpaint 邊界；深色抗鋸齒、輪廓線或
   陰影若混入邊界，會把填補區平均成錯誤色。補圖只回填原 foreground 覆蓋區，保留原輪廓。
5. **從 UV 回到共同空間**：成對特徵的 texture island 可能各自旋轉；不得各沿自己的影像
   長軸移動。先由 UV-to-surface mapping 解出共同 world/local direction，再映回各 island。
6. **共享物理幅度、限制日常控制**：成對構造使用共同位移量，取較緊的可見邊界；完整極限
   留給壓力測試，日常按鈕使用較小範圍，避免每次互動都撞到遮擋邊界。
7. **驗證動態與隔離性**：真實控制項需證明雙向端點不同、成對方向一致、區外 readback
   無漂移、可回 neutral，且能與既有表情／頭部或骨架動作組合。

若 surface-driven adaptation 只存在於特定 viewer，這證明的是該 consumer path 可用，不會
自動把底層資產宣稱為具有獨立幾何／骨骼；record 仍須誠實標記 control mapping 的 method
與適用 surface。

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
