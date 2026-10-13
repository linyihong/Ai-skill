---
case_id: neg.alpha-opaque
kind: reproducible_case
status: spec_frozen
input_hash: 79064bb70d987ef9f4fe01fb2460499eb4186d8479d4e5a304e3a9e304d1f7c2
---

# Case: `neg.alpha-opaque`（負向）

對齊 [`05`](../../05-independent-verification.md) fixture `neg.alpha-opaque`。

## 契約摘要

| 欄位 | 值 |
| --- | --- |
| `task_class` | A |
| `background` | transparent |
| `fidelity` | `strict` preset |
| `acceptance` | 必須有 **有效** 透明（非僅 mode=RGBA） |

## 輸入／合成

- 路徑（預定）：`fixtures/inputs/neg.alpha-opaque.png`
- 合成法（二進位可延後）：RGBA 圖、alpha 通道全 255（或不透明主體鋪滿且無透明背景）
- `input_hash`：`79064bb70d987ef9f4fe01fb2460499eb4186d8479d4e5a304e3a9e304d1f7c2`（合成後填 sha256）

## 期望

| 項 | 期望 |
| --- | --- |
| 執行器若「寫出 PNG」 | 仍 **不得** `accepted` |
| 硬 check | `alpha.not_fully_opaque` → `fail` |
| `failure_class` | `alpha_invalid` |
| 棋盤預覽 | 不可當通過證據 |

## 重跑步驟

1. 放置或合成輸入。
2. 跑 Independent verify（可跳過執行器，直接驗「偽成功輸出」）。
3. 斷言：硬 fail + class；finality ≠ `accepted`。
