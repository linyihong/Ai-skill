# 04 — Background Removal Vertical Slice（P0）

← [`_plan.md`](_plan.md)

## 目標

打通 **A 類去背** 最小閉環，用來驗證決策／執行／驗收／修復是否分離。不在此階段「支援所有圖片」。

## 階段拆解

### 1. Input Validation

見 [`02`](02-image-task-contract.md)。額外：若已有 alpha，記錄是否可信（可能是假透明／白底）。

### 2. Image Analysis

至少標：`subject_count`、`edge_type`、`content_type`、`risk_flags`（hair／text_overlay／glow_halo／thin_structure）。

### 3. Candidate Selection

建議起點（**基準候選，可替換**）：

| 情境 | 先試 |
| --- | --- |
| 一般單一主體 | `seg.generic`（如 rembg 通用模型） |
| 人像／細髮 | `seg.portrait_matting` 或 matting 後處理 |
| 主體錯 | `seg.interactive`／要求 subject_hint |
| 邊緣色污染 | `post.decontaminate` |
| 遮罩形狀錯 | `post.mask_fix`（勿只調顏色） |

### 4. Mask / Matte（分開責任）

可組合、**非每張全做**；每步必須對齊契約 [`02`](02-image-task-contract.md) 的保真度三維：

| 步驟 | 需要的契約授權 |
| --- | --- |
| 前景遮罩取得 | A 類核心（必做） |
| 邊緣精細化／matting | `edge_treatment` ∈ {`matte_refine`, `rim_decontaminate`} |
| 色彩去污染 | `edge_treatment=rim_decontaminate`（限 `max_rim_px`） |
| 小雜點／小浮島去除（有面積上限） | 確定性 post；不得侵入 `immutable` 內部 |
| 局部修正 | 僅 rim／授權區；**不得**當 interior inpaint |

**反模式**（來自 [`07`](07-dogfood-baseline-el-dorado-cutout.md)）：

- 對非金色區域做 binary opening → 咬穿牙齒／內部細節（`subject_structure` 違規）
- 全域硬二值化毁掉柔邊
- 用「暗邊一律內塗」掩蓋錯誤分割（超出 rim＝`fidelity_violation`）
- 把 `fidelity=strict` 解讀成「禁止任何邊緣處理」或「可任意修邊」——兩者都錯；必須讀展開後的三維

### 5. Output Assembly

- 原圖前景 RGB × matte → RGBA
- 透明區 RGB 清零（或依 profile）
- 另產 **預覽**：白底／黑底／棋盤格（**僅預覽**）
- 記錄輸出 hash

### 6. Independent Verify

見 [`05`](05-independent-verification.md)。執行 adapter 不得寫 `status: pass`。

### 7. Repair Loop

見 [`06`](06-failure-taxonomy-and-repair.md)。預設 `max_repair_attempts = 2`（可配置）；超過 → `human_review`。

## 工具分層（不在 Phase 0 定死供應商）

| 層 | 例 | 責任 |
| --- | --- | --- |
| 確定性 | Pillow、OpenCV、NumPy | 尺寸、格式、遮罩運算、像素統計 |
| 分割／Matting | rembg 等（候選） | 主體／matte |
| 生成式編輯 | 僅 C 類且契約允許 | 補畫；**不**替代單純去背 |

授權：權重／程式碼／API **分開**核對；可下載 ≠ 可商用。

## 非目標（本切片）

- 自動對齊外部灰階 Sprite（consumer）
- 自動修 UI 殘字的通解（可作 risk_flag + 專用 post，但非 P0 必達）
- 批次雲端服務化

## Gate（Phase 2）

- 每個案例有 IDR + verification report
- 至少一條「執行成功但 verify fail」的負向案例被正確擋住
- 修復不得無界重試
