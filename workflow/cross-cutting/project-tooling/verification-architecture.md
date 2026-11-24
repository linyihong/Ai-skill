# Verification Architecture（驗證器架構契約）

[`lock-taxonomy.md`](lock-taxonomy.md) §Governance check architecture 定義了**檢查怎麼登記**（single entry / check registry / scope registry / result contract）。本檔定義**檢查怎麼決定要不要跑、結果能不能重用**：當專案的驗證器開始做快取、增量選擇或並行時，哪些不變式不能被「加速」削弱。

> **狀態**：`candidate`（1 個來源案例，未計入 project-tooling 的 3／3 收斂）。只定義語意與判定，**不**規定語言、目錄、runner、daemon 或快取後端。
> 來源計畫：[`plans/active/2026-11-25-0842-verification-architecture-contract/`](../../../plans/active/2026-11-25-0842-verification-architecture-contract/_plan.md)。

## 核心不變式（不隨實作改變）

| # | 不變式 | 違反時的樣子 |
| --- | --- | --- |
| V1 | **必要集合不因效能機制縮小。** 快取、增量選擇、並行、排程只能改變「怎麼跑、何時跑、跑幾個同時」，不能改變「哪些必須被證明」。 | 並行數調低後某組測試沒被排進去；快取命中卻說不出涵蓋了哪些檢查 |
| V2 | **不確定 → 執行（fail-closed）。** 依賴關係、影響範圍或快取有效性無法證明時，決策是 execute 或 blocked，不是 skip / reuse。 | 依賴圖沒收錄某個動態載入，於是判定「不受影響」 |
| V3 | **每次重用都附得出理由。** reuse 必須能指出：輸入指紋、先前的證據、涵蓋範圍。 | 只記「上次通過」，不記上次驗了什麼輸入 |
| V4 | **決策與執行分離。** 先產生計畫（每個檢查的決策 + 理由），再交給執行器；執行器不自行決定略過。 | 每個 gate 各自判斷快取、各自決定 SKIP |
| V5 | **增量選擇要先被證明再被採用。** 從「全跑」改成「只跑受影響的」之前，必須有 shadow 期的選擇準確度證據。 | 依賴分析一上線就直接取代完整驗證 |

V4 對應 [`role-contract.md`](role-contract.md) §入口編排：含快取判斷或範圍選擇的編排屬 Logic，不屬 Entry。planner 與 executor 都是 Logic，Entry 只呼叫它們。

## 能力層（詞彙，不是必建清單）

| 層 | 回答的問題 | 最小契約 |
| --- | --- | --- |
| Change Snapshot | 這次驗證的輸入是什麼 | 一份共用快照，有明確身分（比較基準、已記錄的版本、準備提交的內容、尚未準備的工作區變更分開辨識）；行程內快取**不算**身分 |
| Dependency / Impact | 變更影響到誰 | 可由反向依賴查出受影響的檢查；查不到 → V2 |
| Planner | 每個檢查要做什麼 | 產出決策記錄（見下）；支援只出計畫不執行（plan-only / dry-run） |
| Result Reuse | 什麼結果可以重用 | 依 V3；各層快取隔離（見下） |
| Execution | 怎麼跑 | 批次、並行、行程重用；只受 V1 約束 |
| Evidence | 怎麼證明 | 結果證明、失效原因、涵蓋檢查、獨立覆核 |

**版本控制中立**：本契約不綁定特定版控工具。各專案把 Change Snapshot 的四個身分對應到自己的工具即可。例：Git 是 base commit / HEAD / index（staged）/ working tree；沒有 staging 區的工具，「準備提交的內容」就等於提交時送出的變更集。指紋可用版控工具自己的內容雜湊，或自行計算的內容雜湊；只要同樣內容得到同樣指紋即可。

專案不必六層齊備。**先補缺的是決策正確性（Snapshot、Dependency、Planner），不是執行速度。**

### 決策記錄（最小形狀）

```text
check:        檢查 ID
decision:     execute | reuse | skip | blocked
reason:       一句話；skip 與 reuse 不得為空
input:        輸入指紋（對應 Change Snapshot 身分）
evidence:     reuse 時指向先前證據
coverage:     這個決策涵蓋哪些檢查 / 範圍
```

- `skip` 只用於**政策上不適用**（範圍外）；「判斷不了」是 `blocked`，與 [`lock-taxonomy.md`](lock-taxonomy.md) 的「無法檢查 ≠ 違規」同一語意，必須附恢復方式。
- 快取的可觀測結果（例：HIT / MISS / SKIP / NOT RECORDED）應能對回上面四種決策。

### 快取層隔離

每一層快取（檔案索引、依賴圖、計畫、檢查結果、建置產物、測試產物）各自有版本、輸入與失效條件。**一層命中不推導另一層有效**：檔案索引命中不代表依賴圖有效；結果命中不代表產物存在。

## Selection Accuracy Audit（V5 的證據形狀）

增量選擇上線前，以 shadow mode 運作：正式驗證仍跑完整範圍，另外記錄預測。

```text
snapshot:             Change Snapshot 身分
predicted:            增量選擇預測要跑的檢查
actually_required:    完整執行下實際需要（例：結果與上次不同、或失敗）的檢查
false_negative:       該跑卻沒被預測 → 嚴重，任何一筆都阻擋採用
false_positive:       多跑 → 只是效能損失
unknown_dependencies: 無法證明是否受影響 → 依 V2 執行
```

採用門檻由專案自訂，但「false_negative 必須為零的期間 / 樣本數」要先寫下來（pre-registered），不能看了結果再訂。

## 不在本契約範圍

- 常駐 daemon、遠端快取、cache warming、資源排程器：實作選擇，各專案依量測自行決定；只受 V1–V3 約束。
- 驗證強度 profile（local / commit / push / release）：可以有，但每個 profile 必須有明確的涵蓋清單與延後補驗條件；profile 不是讓執行器自行刪檢查（V1）。
- 共用 runner：依 project-tooling Phase 4 裁決 Q5 維持 deferred。
- 是否採用某個外部建置 / monorepo 工具：屬專案 Selection，不是契約。

← [Back to project-tooling index](README.md)
