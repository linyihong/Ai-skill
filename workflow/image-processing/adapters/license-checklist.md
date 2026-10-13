# License checklist（Q5）

← [`README.md`](README.md) · freeze stub [`08`](../../../plans/active/2026-10-14-0832-image-processing-decision-workflow/08-phase-0-freeze-invariants.md)

在將任一 candidate 用於 **生產 dogfood／對外交付** 前，逐項核對並記入 IDR `notes` 或專案授權紀錄。
本清單滿足 Phase 2 Gate「Q5 授權檢查清單已存在」；**打勾是專案／場次責任**，不是 Ai-skill 自動通過。

## 分割／matting 權重

| 檢查 | 狀態 | 備註 |
| --- | --- | --- |
| 權重名稱／版本已記錄 | [ ] | 例：u2net 變體；寫入 IDR |
| 授權條款已讀（可下載 ≠ 可商用） | [ ] | |
| 商用／再分發限制可接受 | [ ] | 否則排除該 candidate |
| 訓練資料爭議已評估或標 `review_required` | [ ] | 無法確認 → 人審 |

## 執行程式庫

| 套件 | 典型授權（請當場核對上游） | 狀態 |
| --- | --- | --- |
| Pillow | HPND-ish／Pillow license | [ ] |
| NumPy | BSD | [ ] |
| SciPy | BSD | [ ] |
| rembg（若使用） | 專案 license + 依賴權重 | [ ] |
| OpenCV（若使用） | Apache-2.0 | [ ] |

傳遞義務（NOTICE／歸因）已處理：[ ]

## 雲端／API（可選）

| 檢查 | 狀態 |
| --- | --- |
| 未使用雲端 → 標 N/A | [ ] |
| 若使用：ToS、資料出境、刪除政策可接受 | [ ] |
| 原圖不上傳未授權第三方 | [ ] |

## Fixture／dogfood 圖

| 檢查 | 狀態 |
| --- | --- |
| Canonical fixtures 僅合成或已去敏可重用圖 | [ ] |
| 業務未授權原圖不進 Ai-skill | [ ] |
| 每圖有來源說明或 `synthetic:` 標記 | [ ] |

## Phase 2 基準候選（本庫）

| Candidate | 本庫綁定 | 授權備註 |
| --- | --- | --- |
| `seg.generic` | 可選綁 rembg；**Gate 不強制安裝** | 使用前完成本表權重列 |
| `post.*` / verify tools | Pillow（tools/） | 合成負向 fixture 僅 Pillow |

**Gate 說明**：負向 fixture 與 independent verify 僅需 Pillow；不把 rembg 裝進 CI 當 Phase 2 完成條件。
