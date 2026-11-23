# Plan Evidence Index — `2026-11-23-0834-project-tooling-governance`

本目錄存放本 plan 的量測、試判與 dogfood 全文。`_plan.md` 與 [`01-design-contracts.md`](../01-design-contracts.md) 只留摘要與連結。

> **主計畫吸收**：使用 `evidence/` 時，主紀錄在資料夾內的 `_plan.md`。

## 引用規則（避免行號漂移）

| 做法 | 說明 |
|---|---|
| **用檔案路徑** | `evidence/<file>.md` 或相對連結；不寫行號 |
| **用標題錨點** | 定位寫 `evidence/foo.md` 內 `## <標題>` 或表格 `#` 欄 |
| **專案細節** | repo 名、檔名、裝置序號、SHA、可重跑指令 → 各 consumer 的 project-local、非 git `.ai-skill/dogfood/2026-11-23-0834-project-tooling-governance.md`；本目錄只留 generalized metrics（Consumer A = ExternalRepoD、Consumer B = ExternalRepoE） |
| **新 run** | 新增 `evidence/<run-id>-<slug>.md` + 更新本表同一 commit |

Canonical 規則：[`governance/lifecycle/plan-evidence.md`](../../../../governance/lifecycle/plan-evidence.md)

## Run 索引

| Run ID | 檔案 | 狀態 | 摘要 |
|---|---|---|---|
| phase0 | [`phase0-baseline.md`](phase0-baseline.md) | complete（2026-11-23） | 兩個 consumer 的 baseline 量測、A1–A5／B1–B4 verify、concept 試判、placement × state 試標、framework discovery 與 preflight |
| phase2 | [`phase2-readonly-evaluation.md`](phase2-readonly-evaluation.md) | complete（2026-11-23） | 兩個 consumer 唯讀評估：Role 違規句型、concept / discovery、placement × state × ownership、lock review、契約缺口 G1–G6 與歧義 A-1／A-2、建議改動清單（不執行） |
| phase3-0 | [`phase3-0-scope-inventory.md`](phase3-0-scope-inventory.md) | complete（2026-11-23） | Ai-skill 工具樹唯讀盤點：規模、16+2 能力、46 hook 鎖 / 44 registry entry、Go-first 確認、觀察 O1–O6、切片裁決 S1 + 五點邊界 |
| phase3 | [`phase3-application.md`](phase3-application.md) | complete（2026-11-23） | S1 套用：bin 流程預查、18 列 catalog、hook→executor→registry 對帳（L1 scope 缺口）、Role（無違規）、7 把工具治理鎖審查、產品不變式鎖未審清單、O4 skip 標記、後續 F1–F7 |
| f1 | [`f1-repair-slice.md`](f1-repair-slice.md) | complete（2026-11-23） | F1 修復切片：hook-call coverage 修正 L1；17 個 hook callee 逐一分類（enforcement / advisory / helper）並登記；新增 `enforcement_self_governance`；T4 executor 方向樣本通過；F4 policy 同步；觀察 R1–R2 |
