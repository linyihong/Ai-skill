# 02 — Image Task Contract

← [`_plan.md`](_plan.md)

## 目的

在選工具之前，先凍結 **輸入、意圖、輸出、不可變條件**。契約不完整 → 禁止宣告 `accepted`。

## Contract 最小欄位（草案）

| 欄位 | 說明 |
| --- | --- |
| `task_class` | `A` \| `B` \| `C` \| `D`（可多標，須有 primary） |
| `intent` | 使用者意圖摘要（去背／換底／裁切／修復…） |
| `source` | 原圖路徑、hash、格式、尺寸、色彩模式、既有 alpha、EXIF 方向 |
| `output` | 目標格式（預設透明 PNG）、尺寸策略、是否允許裁切／縮放 |
| `background` | `transparent` \| `solid` \| `replace`（+ 色碼／圖） |
| `immutable` | 不可改動清單（臉、字形、配件、品牌色、幾何比例、牙齒／孔洞…）— **清單，不是等級** |
| `subject_hint` | 可選：點／框／「最大主體」／「與參考 silhouette 對齊」 |
| `fidelity` | 見下方 **Fidelity 三維**；可用 preset `strict`／`normal`／`allow_local_edit` 展開 |
| `consumer_profile` | 可選：`slot-symbol`／`product`／`portrait`… 附加約束 |
| `acceptance` | 可機器化的通過條件 + 需人審條件 |

## Fidelity 三維（取代單一「像素級前景」標籤）

一個 `strict` 不夠區分「主體不可動」與「邊緣可修」。契約必須能分開表達：

| 維度 | 含義 | 典型約束 |
| --- | --- | --- |
| `subject_structure` | 主體拓撲、孔洞、薄結構、內部陰影是否保留 | `preserve`（預設）\| `allow_fill_holes`（須明示） |
| `subject_color` | 主體本體配色／材質細節 | `preserve`（預設）\| `allow_recolor` |
| `edge_treatment` | 邊緣帶允許的操作與寬度上限 | `none` \| `rim_decontaminate` \| `matte_refine`（各可帶 `max_rim_px`） |

### Preset 展開（草案；Selection 前展開進契約）

| Preset | `subject_structure` | `subject_color` | `edge_treatment` | 另排除 |
| --- | --- | --- | --- | --- |
| `strict` | `preserve` | `preserve` | 預設 `rim_decontaminate`（`max_rim_px` 由 profile／分析定；未定 → ask） | `gen.edit`；激進 interior morphology |
| `normal` | `preserve` | `preserve` | `matte_refine` 或 `rim_decontaminate` | `gen.edit` 預設不在可行集 |
| `allow_local_edit` | 契約另列可改區 | 可局部 | 依契約 | 僅授權區可 `gen.edit` |

### 責任邊界（防 B1 衝突）

1. **未授權的像素變更** → `fidelity_violation`（見 [`06`](06-failure-taxonomy-and-repair.md)）。
2. **已授權 `edge_treatment` 帶內的變更**不算違反 `subject_structure`／`subject_color`，但仍受 `max_rim_px` 與 `immutable` 清單約束（例如牙齒／孔洞永不在 rim 內）。
3. **`immutable` 是物件／區域清單**；**不可**寫成 `immutable=strict`。等級只用 `fidelity` preset 或三維欄位。
4. 去污染／局部修正（[`04`](04-bg-removal-vertical-slice.md)）只能落在已授權的 `edge_treatment`；超過 rim → fail。

## 輸入檢查（Input Validation）

必做：

1. 解碼成功；記錄 W×H、mode、bits、icc（若有）
2. 正規化 EXIF 方向（記錄是否旋轉）
3. 原圖 content hash（例如 sha256）
4. 確認 `background` 意圖；未指定 → **ask** 或 `inconclusive`，不猜測
5. 明確 `immutable` 清單；未聲明時預設 **禁止改主體外觀／禁止生成式重畫**
6. 展開 `fidelity` preset → 三維；`strict` 且 `edge_treatment` 寬度未定 → ask，不默認「可任意修邊」

## 輸出不變條件（Examples）

### A — 去背／透明化（P0）

- 輸出含有效 alpha；「看起來透明」的棋盤格預覽 **不是** 契約滿足
- `subject_*` = `preserve` 時：不重繪臉／配件／牙齒／孔洞、不改變主體本體配色
- 僅在授權 `edge_treatment` 帶內可去暈／refine matte
- 背景像素 alpha=0 且 RGB 應歸零或依 profile 規定（防預乘污染）

### D — 確定性處理

- 輸出尺寸／格式／色彩空間 = 規格
- 未列操作不得執行

## Consumer profile 例：`slot-symbol`（非核心）

業務約束可附加，**不得**污染通用 A 類核心：

- 輸出固定 canvas（如 328²）
- 可要求與灰階 Sprite silhouette／face-axis 對齊
- UI 殘字／髮絲線必須剔除

這些寫在 consumer profile，由 Selection 階段變成額外 constraints。

## Incomplete contract 行為

| 缺什麼 | 行為 |
| --- | --- |
| 無輸出背景意圖 | 停止或 ask |
| 多主體且無 subject_hint、分析低信心 | ask／human_review |
| 要求「更好看」但 `fidelity` preset=`strict`（或 `subject_*`=`preserve`） | 拒絕生成式路徑 |
| `strict` 但未定 `edge_treatment`／`max_rim_px` | ask；禁止默認全域去污染 |

## Gate（Phase 1）

不依賴任何模型名，也能判斷：契約是否完整、三維保真是否展開、哪些候選被 constraints 排除。
