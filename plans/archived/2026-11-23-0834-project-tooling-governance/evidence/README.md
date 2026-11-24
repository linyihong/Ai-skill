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
| f2 | [`f2-repair-slice.md`](f2-repair-slice.md) | complete（2026-11-24） | F2 修復切片：宣告式命令表＋雙向狀態比對取代 diff pattern 鎖（T3）；修正 9 個未入 contract 的子命令；`warnCLIDocSync` deprecated；觀察 S1–S3；Consumer B T-B1–T-B5 與 Ai-skill 對照（Phase 4 候選） |
| f6 | [`f6-repair-slice.md`](f6-repair-slice.md) | complete（2026-11-24） | F6 修復切片：T1 加上 shebang 內容偵測（`core.hooksPath` 豁免、沒有內容時降級）；pre-push replay 改用同一條規則（原本只比對 `.sh`）；X1–X3 突變；觀察 V1–V2 |
| s2 | [`s2-portable-kge-survey.md`](s2-portable-kge-survey.md) | complete（2026-11-24） | S2 唯讀盤點：外部專案接觸 `DefaultRules()` 的三條路徑（複製單元、Go import、執行 binary）；本機 88 個 repo 都沒有使用，`CLIDocSyncRule` 和 V1 在盤點範圍內沒有外部使用者；限制 K1–K3（observed，非 verified）；觀察 W1 |
| consumer-b-cli-pilot | [`consumer-b-cli-pilot.md`](consumer-b-cli-pilot.md) | complete（2026-11-23；2026-11-24 查證已合併） | Consumer B 自有的單一 CLI 入口試點：7 閘邏輯去重＋固定樣本快取（快取熱 −75%）；推翻「行程啟動是主要成本」；共用 I/O 實測（每檔平均被讀 7.5 次，單行程共用讀取 59 s → 13～14 s）；觀察 T-B1–T-B5，不改契約 |
| consumer-a-temp-registry | [`consumer-a-temp-registry.md`](consumer-a-temp-registry.md) | complete（2026-11-24；唯讀） | Consumer A 自己已合併的暫存檔登記框架：Scratch 位置與責任符合；G4 機械清理邊界 2／3（沒有先隔離、可恢復）；Role 和 concept-first 是正例；契約候選 C1–C2；R1–R4 重新量測仍未處理；Consumer A 的 agent 依 brief 修改後重新檢查通過（§7；外部依契約修改的案例，限制：鎖不在版控內、R1 沒在真機驗證）；契約候選 C3–C7 |
| phase4 | [`phase4-decision-package.md`](phase4-decision-package.md) | complete（2026-11-24） | Phase 4 裁決資料包：案例 K1–K6、invariant × 案例矩陣、Q11 選項（裁決：只算實際套用，3／3）、契約修改候選（T-B4、T-B5+C5、C6 接受）、Q5／Q7／runtime 接入 deferred、遺留項目移到後續計畫 |
