---
case_id: neg.residue-halo
kind: reproducible_case
status: spec_frozen
input_hash: 5c3d55e616ca94d8e4ff3e51da662ed257aab4a2294d403fff1c26373319239c
---

# Case: `neg.residue-halo`（負向）

對齊 [`05`](../../05-independent-verification.md) fixture `neg.residue-halo`。

## 契約摘要

| 欄位 | 值 |
| --- | --- |
| `task_class` | A |
| `background` | transparent |
| `fidelity` | `strict` + `edge_treatment=rim_decontaminate` |
| `acceptance` | 無明顯 rim 色暈 |

## 輸入／合成

- 路徑（預定）：`fixtures/inputs/neg.residue-halo.png`（或「錯誤輸出」固定樣本）
- 特徵：主體外仍有 navy／綠／暗 rim（模擬 rembg 未後處理）
- `input_hash`：`5c3d55e616ca94d8e4ff3e51da662ed257aab4a2294d403fff1c26373319239c`

## 期望

| 項 | 期望 |
| --- | --- |
| Verify | `residue.rim`（或背景殘留族）→ `fail` |
| `failure_class` | `edge_halo` 或 `background_residue` |
| Finality | 不得 `accepted`；可 `repair` → `post.decontaminate` |

## 重跑步驟

1. 載入錯誤輸出樣本（或執行 baseline seg 不加 decontaminate）。
2. Independent verify。
3. 斷言擋住；IDR 記 `repair_candidate`。
