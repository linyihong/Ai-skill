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

**Status**: in-progress — Phase 0、Phase 1 已完成；Phase 2a（Ai-skill 自身 dogfood）S1 量測、S2 V2 驗證與修補、S3 PreToolUse 增量化評估完成；Phase 2 外部證據等待中。

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
| Change Snapshot（明確 tree / index / working-tree 身分） | **採納為詞彙，抽象為版控中立** | 契約能力層；行程內快取不算身分。2026-11-25 使用者指出契約不應綁定 Git：身分改寫為比較基準 / 已記錄版本 / 準備提交 / 工作區四項，Git 只作對應範例 |
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
| Ai-skill 自身 hooks | **dogfood 進行中**（見 §Phase 2a）：基線量測完成；V2 已套用（路徑式測試選擇改為保守全跑） | 已有路徑式測試選擇（V2 相關）與 Go 內建測試快取（V3 相關） |

- [ ] 外部 consumer 修復 E2E 漏跑後，唯讀檢查其是否套用 V1–V5 的任何一條，記錄於本計畫 `evidence/`
- [ ] 若 consumer 執行 shadow mode，記錄其預先登記的採用門檻與第一批 audit 結果
- [ ] 尋找第 2、第 3 個有結果快取或增量選擇的 consumer

## Phase 2a：Ai-skill 自身 dogfood（進行中）

2026-11-25 使用者決定以 Ai-skill 自身 hooks 作為第一個實際套用案例，同時加速系統。依契約順序：先量測，再決策記錄，再重用，最後才做 shadow 之後的跳過。

- [x] S1 耗時基線：[`evidence/s1-ai-skill-hook-timing-baseline.md`](evidence/s1-ai-skill-hook-timing-baseline.md)
  - git hooks 空跑都 < 0.2 s → 不做快取
  - `go test ./...` 未快取約 25 s，Go 內建快取約 2 s → **不自建結果快取**（會與 Go 快取重疊）
  - PreToolUse 每次工具呼叫約 0.32 s（3.1 MB transcript），隨 transcript 變大 → 最頻繁的成本
  - 路徑式測試選擇（pre-push 與 CI）不涵蓋測試會讀取的 `runtime/`、`knowledge/`、`plans/` → V2 缺口（observed）
- [x] V2 驗證：**verified**。[`evidence/s2-v2-path-selection-falsification.md`](evidence/s2-v2-path-selection-falsification.md)：只改 `metadata/` 或 `plans/` 的探針讓測試失敗，pre-push 仍放行、不跑 Go 測試；CI paths 過濾同樣不涵蓋（讀設定得出）
- [x] 依 V2 結果修補（2026-11-25 使用者裁決：方案 1 保守全跑）：pre-push 只要有任何變更就跑 `go test ./...`（`hasPrePushChange`）；CI 移除 `paths` 過濾。第一個實際套用的不變式：V2
- [x] PreToolUse 增量化評估：[`evidence/s3-pretooluse-incremental-parse-evaluation.md`](evidence/s3-pretooluse-incremental-parse-evaluation.md)。5.1 MB 時每次 0.52 s；主因是同一次呼叫內全量解析兩次（各 134 ms）。shadow 採用門檻已先登記
- [x] 步驟 0：每次呼叫只解析一次（`loadPreToolUseRouting`）。5.3 MB 時中位數 0.62 → 0.40 s；純重構，輸出與舊版相同
- [ ] 步驟 1：跨呼叫增量解析，**只做 shadow**（2026-11-25 使用者裁決進入；收益估算見 S3：每個 session 中位數省 39 s，最大省 323 s）。執行邊界如下：

  | 階段 | 工作 | 要求 |
  | --- | --- | --- |
  | 1a | 建立增量解析器 | 不影響正式解析結果 |
  | 1b | Shadow 比對 | 現行完整解析仍是唯一權威結果 |
  | 1c | 累積證據 | ≥ 200 次比對、≥ 5 個 session（涵蓋條件見下） |
  | 1d | 一致性裁決 | 0 次不一致才可**提出**切換；切換另需使用者裁決 |
  | 1e | 獨立驗證 | 由獨立執行者確認失效、重置、截斷等邊界條件 |

  **失效一律退回完整解析**（V2）。下列任一情況，增量解析器都必須丟棄狀態、從頭重建：
  - 沒有快取、快取無法解碼，或快取版本不符；
  - 快取記錄的 transcript 路徑或 session 與本次不同；
  - 檔案比已處理的位移量短（被截斷）；
  - 檔案開頭的雜湊不符（被覆寫）；
  - 位移量前的邊界雜湊不符（已處理的區段被修改）；
  - 其他任何無法確認游標仍然有效的情況。

  **Shadow 必須獨立**：
  - 增量解析器自己開檔讀取，自己保存狀態，不使用完整解析的結果或中間狀態。
  - 兩者共用的只有逐行抽取文字的純函式（`extractEntryText`）。因此 shadow 驗證的是游標、邊界與失效機制，**不驗證逐行抽取邏輯本身**；後者的正確性仍由既有測試負責。
  - 比對時兩者處理同一段位元組：完整解析前後檔案大小不同（transcript 在比對期間變長）時，這次比對記為「不計入」，不算一致也不算不一致。

  **證據涵蓋條件**：200 次、0 不一致只是進入 1d 的必要條件，不是永久正確的證明。證據還要涵蓋：
  - 不同 transcript 大小：< 1 MB、1–5 MB、> 5 MB 都要有；
  - 至少 1 個經過 compaction / resume 的 session，至少 1 個有平行工具呼叫；
  - 至少 1 次各種失效退回的實際觸發；
  - 異常輸入：無法解碼的行、超長行、最後一行未寫完。

  **禁止事項**（直到 1d 經使用者裁決之前）：
  - 增量結果不得影響正式 routing、gate 或 advisory；
  - 不得移除或繞過完整解析；
  - 不得宣稱 V3 或 V5 已驗證；
  - Shadow 紀錄只記計數、大小、耗時與是否一致，不記 transcript 文字。
- [ ] （選擇性）`internal/app` 最慢的 runtime compiler 測試切分或共用 fixture；屬測試效能，不屬快取

**計數**：Ai-skill 自身最多算 1 個實際套用案例；仍需 2 個外部案例才進入 Phase 3。

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
