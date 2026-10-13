---
case_id: illust-glow-halo
kind: reproducible_case
status: spec_frozen
input_hash: PENDING_BINARY
---

# Case: `illust-glow-halo`

## 契約摘要

| 欄位 | 值 |
| --- | --- |
| `task_class` | A（primary） |
| `background` | transparent |
| `immutable` | 主體臉／配件；禁止 gen.edit |
| `fidelity` | preset `strict` → structure/color `preserve`；`edge_treatment=rim_decontaminate`；`max_rim_px` ≤ 4（未定則 ask——本案例固定 4） |
| `consumer_profile` | none（通用 A） |

## 輸入

- 路徑（預定）：`fixtures/inputs/illust-glow-halo.png`
- 特徵：插畫主體 + 明顯深色／彩色外光暈底
- `input_hash`：`PENDING_BINARY`（放入去敏或合成圖後改 sha256）

## 期望（比較基準）

| 項 | 期望 |
| --- | --- |
| 若僅跑 `seg.generic` 無 rim 處理 | 硬 check `residue.rim` → `fail`；`failure_class=edge_halo` |
| 完整閉環（授權 rim decontaminate） | 可達硬 check pass；軟邊緣可 `review_required` |
| 不可 | 用棋盤預覽當 alpha 證據；對內部 opening |

## 重跑步驟

1. 載入本契約 + 輸入圖（hash 核對）。
2. 執行選定 candidate（IDR 記錄）。
3. Independent verify（非執行器）跑 `alpha.*` + `residue.rim`。
4. 比對上表期望；寫 IDR finality。
