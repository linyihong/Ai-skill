# Phase 4：方法決策與 deferred 交接

## 決策

使用者已接受 Phase 3 在既定 current-agent 範圍結清，並要求直接完成方法決策，不為階段名稱新增實驗。依目前證據，**沒有足夠證據要求再增加檢索基礎設施**；這不等於沒有任何語意缺口。

| 項目 | 本輪處置 | 依據與邊界 |
| --- | --- | --- |
| 向量搜尋／vector DB | 不啟用 | Query/extraction 有互補因果證據；目前沒有證明需新增 candidate generation 能力。未量測的任務／全庫 recall 不推定已足夠 |
| 既有 SQLite／source-level lookup | 保留 | [production E2E](phase2-production-wiring.md) 使用實際 pool，23/23 指定來源、MRR .971014、selected confuser 0/6、完整 payload ≤1622 runes；限曝光診斷 |
| 現有 agent 語意選擇 | 已納入並有範圍驗證 | [gold-blind execution](phase2-blind-selector.md)、production E2E、[真實 dogfood](phase3-dogfood.md)、[supersession consumer](phase3-supersession-consumer.md)；不外推自主通用適用性或泛化 |
| 額外 reranker／模型服務 | 不新增 | 本輪未證明超過現有 agent 選擇的必要收益；沒有新服務、API、付費模型或新方法實驗 |
| Phase 4 新模型／獨立 holdout／成本契約 | 本輪實驗未啟動，因此不適用 | 舊資料已曝光，不重稱獨立 holdout；未來若新增方法須先另凍結未曝光資料、方法版本及成本／品質門檻，不能拿這次不啟動當通過那些門檻 |

V1／V2 被拒的結果、舊 fixed pool 分數與其限制原樣保留。沒有重跑新 selector、調參、回改排名、覆蓋原 raw 或新增 reviewer。參考 token 計數與實際 usage 分開；Windows 與實際 tokens 的暫緩由使用者明確指定，不冒稱通過。

## Deferred owner／重啟條件

Owner 指責任承接角色，不建立背景工作、排程、服務或新的 active goal；需要重啟時由 larrylin 指定當次執行者。

| 項目 | Owner | 本輪狀態 | 重啟條件與最小交接 |
| --- | --- | --- | --- |
| 可重跑漏找／索引覆蓋（含尚未索引來源與全庫 intelligence） | Ai-skill CLI／knowledge maintainer；由 larrylin 指派 | 七份 intelligence probes 及固定 lesson 診斷之外不宣稱 recall | 真實任務有「應找到但未找到」的既有來源；保存 query／intent、來源 revision、index／binary identity、filters、實際輸出，先分辨未索引、過濾、query 或 consumer 問題 |
| 向量 candidate generation | CLI／retrieval maintainer，larrylin 決定啟動 | 本輪不啟用 | 上述可重跑漏找經診斷仍是 semantic candidate-generation 缺口；先凍結獨立比較資料與成本契約，再決定是否值得新增能力 |
| 選擇品質／成本、額外 reranker | sd-validation／workflow maintainer；larrylin 指派 | 現有 agent 保留，額外服務不新增 | 候選已入池，但真實選擇品質或時間／context 成本不再滿足需求；保留候選集合與最終 consumer 輸出，分開比較方法，不先縮到 top1 掩蓋問題 |
| 獨立泛化／自然無答案誤推率／其他 supersession 形式 | Retrieval evaluation maintainer + consumer task owner | 未由曝光資料與人工控制證明 | 有跨任務泛化或此類品質宣稱的實際需求；收新資料並記使用歷程，不把原26題或本輪遷移 fixture 稱新 holdout |
| 實際模型 token usage | Consumer／context maintainer | deferred-by-user；實際 usage unknown；參考 encoding 計數已另存 | 有對應模型／tokenizer／訊息範圍或 runner usage，且使用者恢復此成本驗收需求時再量測；目前 agent 執行不因此停止 |
| Windows 當前 tuple／native repair | 該產品 maintainer／larrylin | deferred-by-user；適用性 unknown；未套用舊修法 | **只有使用者明確恢復 Windows 調查**才重啟；先核對當前平台、版本、指紋與 scoped repair，必要時 focused tests；不自動重連、重啟、載入模型或擴大產品 PASS |
| 無人值守 runner／其他 host adapter | Workflow／adapter maintainer | 只驗本輪既有 agent／已記錄 adapter 邊界 | 真實 workflow 需要不同 host 或 unattended dispatch；先驗 deadline→parent interrupt→readback，usage 不可得仍明列，不由 watcher alone 宣稱 child 已停止 |
| Bridge Deep Discovery／其他經驗沉澱 roadmap | 原 Discovery Bridge／Feedback Execution Closure 計畫 owner | 不由本計畫接管或宣稱完成 | 依各原計畫的觸發／驗收恢復；本輪不改 Bridge activation／advisory 邊界 |

## 結清範圍

Phase 0–3 的縮限／current-agent 結論與使用者接受範圍不變；Phase 4 是有證據的「本輪不新增」決策。Deferred 已有 owner／條件，故不因未知需全部清零或未導入 vector 阻擋核心計畫結清。新任務出現可重跑缺口時，按上表恢復，而非以本次完成宣稱長期不再有缺口。
