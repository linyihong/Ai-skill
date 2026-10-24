# Phase 0 — 檢索現況與架構相容性盤點

日期：2026-10-23。程式與契約基準：`86d0cdb0`；lookup cache 為本機既有 cache，未聲稱全庫已 refresh。以下計數與探測只代表此 snapshot，Phase 1 必須重建並凍結 corpus 再正式量測。

## Preflight

| 欄位 | 結果 |
| --- | --- |
| Trigger | 使用者授權分段執行、commit / push；完成計畫修訂後進入 Phase 0 |
| Checked sources | `internal/app/runtime.go`、`discovery.go`、`hooks.go`；`runtime/README.md`、`runtime.db` schema / generated surfaces；`knowledge/README.md`、SQLite README；`metadata/schema.md`；software-delivery README / execution-flow / validation；intelligence README / engineering README / validation-reasoning README；memory retrieval governance README |
| Conflicts | 未見需新搜尋服務的架構理由；cache 與 canonical runtime.db 不可混為一談；metadata schema 與已索引欄位並不相等 |
| Interrogation | 目標為既有經驗被找回並驗證適用；核心 consumer 是 workflow 內查找；不改 activation、不新建 memory replay 路徑 |
| Decision | proceed Phase 1；Phase 2 前用 benchmark 決定最小正文覆蓋，不先全庫 chunking |
| Validation | 讀 generator / query / hook、查 schema / 計數 / checksum、執行 filtered query；沒有 benchmark 或採用鏈通過宣稱 |

## 來源覆蓋與生成路徑（Q1 resolved）

`runtimeIndexRecords` 只合併 summaries、routes、feedback 三種 records；不是遞迴掃描所有 intelligence 正文。`runtimeIndexSources` 再對 records 指向的 source 去重並計算 checksum；source 表不是全庫 inventory。

| Snapshot 欄位 | 計數 / 觀察 |
| --- | --- |
| atoms / sources / fts / edges | 474 / 460 / 474 / 350 |
| `intelligence/**/*.md`（包含 README） | 233 |
| intelligence sources / atoms | 6 / 7 |
| intelligence sources 中非 README 文件 | 1 |
| 6 個 intelligence source 的 checksum | 與當前檔案相符；不代表其他 layer 全部 fresh |
| intelligence atoms 的非空 domains | 0 |

六個 intelligence source 為 agent-architecture README、highest-leverage-analysis-path、apk-analysis README、architectural-fit README、heuristics README、requirements README。未索引例子：`intelligence/engineering/execution/validation-reasoning/evidence-collapse-point.md`；不是拿 README 與正文總數比值冒充 Recall。

Phase 2 的最小候選：先針對 validation-reasoning 文件建立 source-level 正文檢索；標題 / 段落與 source identity 保留，是否切 section 由 Phase 1 失敗案例決定。不預先替 233 份文件全部造 summary。這是範圍決策，不是已完成正文索引。

## 查詢與可重跑入口（Q2 still-open）

從 repo root 呼叫平台 binary；flags 必須在 positional query 前，或統一用 `--keyword`。本輪發現 Go flag parser 遇到第一個 positional 即停止解析，`query evidence --layer intelligence --limit 5 --json` 不會套用後面 filters；正式 baseline 固定下面形式，避免把呼叫問題誤判成檢索能力問題。

```text
scripts/ai-skill-cli/bin/ai-skill-darwin-arm64 runtime query --layer intelligence --limit 5 --keyword evidence --json
scripts/ai-skill-cli/bin/ai-skill-darwin-arm64 runtime query --layer intelligence --limit 5 --keyword "evidence collapse" --json
```

前者命中 4 個 atoms（其中兩個是同一 source），後者 0 筆。`runtimeFTSMatchLiteral` 將整個輸入包成 quoted phrase，並非長句 query expansion。source-level dedup 與長句處理有明確調查理由；中文 / 混合詞完整評估留 Phase 1，這兩次探測不是 20 題基準。

## Metadata 與來源權威（Q3 resolved）

| 意義 | 已有來源 / 欄位 | 本計畫處置 |
| --- | --- | --- |
| domain | metadata `domains`、atoms domains | 重用；summary / route producer 未填值時為 unknown，不據此硬排除 |
| lifecycle | metadata status、文件 Status；route compatibility state 也會進 atoms.status | 保留原值與來源，不把 candidate-intelligence 或 compatibility state 自動當 validated |
| supersession | metadata 選填 `replaces` | 從 canonical source 有宣告才使用；目前 atoms 不含 replaces，不虛構已可 filter |
| project / 前提 | 無通用 project scope 索引欄位；各 source / workflow 可能以正文表達 | 回讀人工核對，缺值 unknown；不在此階段新增 schema |
| 版本 | sources checksum、git corpus revision | 索引 freshness 與實際採用版本分開；checksum 改變觸發重驗 |
| 證據是否失效 | 不由 checksum 單獨決定 | 依 source 的證據與前提確認；unknown 不算通過 |

memory retrieval / replay 有獨立 qualification 與 budget。本次索引候選不得自動進入 memory/working，不接管該 pipeline。`runtime.discovery.config` 的 generated surface 已存在且狀態 synced；不新增 target_key 或 runtime YAML mirror。

## 首批 Consumer 與查找契約（Q4 resolved）

選定 `workflow/software-delivery/validation.md`（sd-validation）作第一批 consumer；既有用途就是判斷測試與證據是否足以支持完成宣稱，對應 intelligence validation-reasoning。當前 CI 排查是可覆核的真實驗證任務入口，但尚未接入檢索採用鏈，不能算 Phase 3 dogfood PASS。

| 事件 / 條件 | 預定動作與可覆核記錄 |
| --- | --- |
| 實作完成、選擇最低足夠證據並形成驗證方案前 | 查 validation-reasoning，記 query / task scope / 候選來源；required reads 仍直接讀 |
| 測試失敗，證據鏈第一個失去一致性的環節不明 | 查 evidence-chain / collapse-point 類經驗，回讀後判斷前提是否一致 |
| 關鍵前提、來源版本、失敗類型或風險範圍改變 | 重查並重新驗適用性；不沿用上次採用判定 |
| 同任務相同前提與 source checksum，且已有可回讀的查找紀錄 | 可略過重複 query；引用先前 query / 版本 / 適用性理由 |
| 純格式改動且既有驗證 checklist 已直接覆蓋 | 可略過經驗搜尋；記 diff 範圍與 checklist 理由，不豁免 required_read |

這些是 Phase 3 待接入 consumer 的具體契約，不是在本輪改動 workflow。Phase 1 必須從現有真實問題選第二個獨立任務，不能只改寫本次 CI 案例湊兩例。

## Bridge 與成本（Q5 resolved；Q6 deferred）

本輪未發現必須改 Bridge 呼叫點才能交付 workflow lookup 的依賴。保留既有 weighted Light ranking、miss eligibility、30ms p95 配置與 advisory cap；不接入新增正文查詢，不抽通用 trigger framework。只有後續證據顯示重複實作或可量測收益才重開 Bridge 改動；該改動不能借 workflow 的寬 budget 放寬既有 Bridge 限制。

Phase 1/2 首批 workflow lookup 的初始工程 budget：本機 warm lookup（含 CLI、排除人工回讀）p95 ≤ 500ms、候選 top 5、候選預覽總計 ≤ 2000 tokens；cold start 另外報告，超預算時回到 source lookup，不擴大 blocker bypass。Phase 1 用至少 20 次重跑量 latency；token 計數方法需凍結，估算值必須標估算。這是待 benchmark 驗證的上限，不是本輪效能結果。

檢索層 Phase 2、採用層 Phase 3 驗收；目前 Recall、MRR、採用錯誤、blocker 保持均為 `not evaluated`。Q6 留 Phase 4，需剩餘語意漏找與獨立 holdout 才開實驗。

## Linked Updates / 後續交付

本次只回寫 plan、Run 索引與 plans 狀態；未改 metadata、workflow、routing、compiler 或 lookup source，故不重建 generated artifacts。Phase 1 需固定重建 corpus、工具版本與 split，再量測；Phase 2 變更索引時才做 refresh / validate，按 affected source 檢查 generated consumers。
