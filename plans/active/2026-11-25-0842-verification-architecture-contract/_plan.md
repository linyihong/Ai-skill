---
id: 2026-11-25-0842-verification-architecture-contract
plan_kind: main
status: in-progress
owner: linyihong
owner_layer: workflow
created: 2026-11-25
parent: null
---

# Verification Architecture Contract（驗證器架構契約）

**Status**: in-progress — Phase 0、Phase 1 已完成；Phase 2 起等外部證據。

## 一句話目標

把「驗證器如何決定要不要跑、結果能不能重用」的正確性不變式，補進 project-tooling 的驗證腳本架構，並在累積到足夠實際套用案例前維持 `candidate`。

## 背景

[`project-tooling`](../../../workflow/cross-cutting/project-tooling/README.md) 契約已經定義了檢查怎麼**登記**（[`lock-taxonomy.md`](../../../workflow/cross-cutting/project-tooling/lock-taxonomy.md) §Governance check architecture：single entry / check registry / scope registry / result contract），但沒有定義檢查怎麼**選擇與重用**。外部 consumer 的驗證效能計畫（W0–W8：共享讀檔、Git blob 指紋 attestation、input closure、HIT / MISS / SKIP / NOT RECORDED、E2E 模組化並行、突變測試與獨立驗證）收到一份外部建議，提出六層 Incremental Verification Engine 與 Selection Accuracy Audit。本計畫評估那份建議，只把**跨專案成立的部分**寫進契約。

外部 consumer 的 repo 只做唯讀評估，不在其 repo 內修改；名稱與模組名在本計畫中去識別化。

## Phase 0：建議評估（完成）

裁決原則：Ai-skill 只收**正確性不變式與詞彙**；效能實作（daemon、遠端快取、預熱、資源排程）屬各專案 Selection。

| 建議項目 | 裁決 | 落點 / 理由 |
| --- | --- | --- |
| 必要集合不因效能機制縮小（排程只改順序與並行度） | **採納** | 契約 V1 |
| 依賴不確定不得跳過（Unknown Dependency Fallback） | **採納** | 契約 V2（fail-closed） |
| 每次重用要能證明為什麼安全 | **採納** | 契約 V3 |
| Verification Planner：決策與執行分離、plan-only | **採納** | 契約 V4；與 role-contract §入口編排一致 |
| Selection Accuracy Audit + Shadow Mode | **採納** | 契約 V5 + §Selection Accuracy Audit；補上「門檻要先登記」 |
| Change Snapshot（明確 tree / index / working-tree 身分） | **採納為詞彙** | 契約能力層；行程內快取不算身分 |
| 多層快取 + Cache Layer Isolation | **只採 isolation** | 層數與層名屬實作；「一層命中不推導另一層」是不變式 |
| 決策 enum（execute / reuse / skip / blocked） | **採納，並收緊** | `skip` 只用於政策不適用；判斷不了是 `blocked`，與 lock-taxonomy「無法檢查 ≠ 違規」對齊 |
| Verification Profile | **只採約束** | profile 必須有涵蓋清單；不得讓執行器自行刪檢查 |
| Persistent Worker / daemon、Remote Artifact Cache、Cache Warming、Resource Scheduler | **不收入契約** | 實作選擇；只受 V1–V3 約束 |
| 導入 Nx 等外部工具取代既有驗證器 | **不收入契約** | 屬專案 Selection；建議本身也不推薦 |
| 外部 consumer 的三階段 roadmap（先修 E2E 漏跑） | **不收入契約，記為給 consumer 的建議** | 見 §給外部 consumer 的建議 |

## Phase 1：寫入契約（完成）

- [x] 新增 [`verification-architecture.md`](../../../workflow/cross-cutting/project-tooling/verification-architecture.md)（`candidate`）
- [x] project-tooling [`README.md`](../../../workflow/cross-cutting/project-tooling/README.md) 契約索引加一列
- [x] [`lock-taxonomy.md`](../../../workflow/cross-cutting/project-tooling/lock-taxonomy.md) §Governance check architecture 加一行指向新契約

不加 validator、不加 route、不加 runner、不加 runtime 接入。

## Phase 2：收集實際套用證據（等待中）

計數規則沿用 project-tooling Q11：只算**實際套用**；同一 consumer 只算一次；唯讀評估不算。

| 案例 | 狀態 | 涵蓋的不變式 |
| --- | --- | --- |
| 外部 consumer（驗證效能計畫） | 來源案例，**尚未計入**：建議尚未被採用，E2E 漏跑（其 W3-0）尚未修復 | 現有 attestation / input closure 部分對應 V3；V1–V2、V4–V5 未套用 |
| Ai-skill 自身 hooks | **不適用**：pre-push 目前沒有結果快取或增量選擇 | — |

- [ ] 外部 consumer 修復 E2E 漏跑後，唯讀檢查其是否套用 V1–V5 的任何一條，記錄於本計畫 `evidence/`
- [ ] 若 consumer 執行 shadow mode，記錄其預先登記的採用門檻與第一批 audit 結果
- [ ] 尋找第 2、第 3 個有結果快取或增量選擇的 consumer

## Phase 3：升級裁決（未開始）

- [ ] 達到 3 個實際套用案例後，決定 `candidate` → 併入 project-tooling pilot，或修改 / 撤回不變式
- [ ] 若某條不變式在案例中被證明無法遵守或不需要，在契約中標記 refuted 並附證據，不默默刪除
- [ ] Plan Completion Closure

**重啟 / 撤回條件**：2027-05-31 前仍只有 1 個來源案例 → 重新評估契約是否撤回為純參考。

## 給外部 consumer 的建議（不寫入其 repo）

1. **P0 先修 E2E 漏跑**；在必要集合被證明完整前，不擴大快取與增量範圍（V1）。
2. 現有 W0–W8 維持「執行最佳化」範圍，不加入新層。
3. 修好後另開 linked plan（建議名稱 `incremental-verification-intelligence`），範圍限定：Change Snapshot、Dependency / Impact、Planner 決策記錄、Selection Accuracy Audit（shadow mode）。
4. daemon、遠端快取、預熱、資源排程留到 shadow 證據出來之後。

## 完成條件

- [x] 建議已逐項裁決（Phase 0）
- [x] 契約已寫入並連入索引（Phase 1）
- [ ] Phase 2 至少 3 個實際套用案例，或到期撤回裁決
- [ ] Phase 3 升級裁決已記錄

**Glossary Impact**: 契約內詞彙（Change Snapshot、Selection Accuracy Audit、decision 四值）為 project-tooling 契約局部用語，不進 framework glossary；注意 `Evidence` 層與 Ai-skill plan `evidence/` 目錄同名不同義（terminology collision，觀察中）。
