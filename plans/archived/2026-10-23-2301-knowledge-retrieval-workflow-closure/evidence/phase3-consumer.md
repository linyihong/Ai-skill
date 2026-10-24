# Phase 3 consumer：第一段接線與受控讀回

## 結論

首批 sd-validation 已接上 workflow 層的查找／重查／略過、來源回讀與決策紀錄；未改 Bridge、selector、CLI、activation 或 authority。Phase 3 維持未完成。

[事前範圍](phase3-consumer-protocol.md) 沿用主計畫 gates，沒有新增驗收。新增操作正文以繁體中文撰寫，獨立子文件避免 validation 主檔再膨脹。CLI 語意沿用既有 contract，不新建 YAML gate、trigger engine 或自動採用器。

## 實際 trial

本次真實工作是「為這次 workflow 接線選擇足夠的驗證，避免文件存在被宣稱為採用效果」。現有 release binary 對本 repo 建 cache，查 intelligence，11 個來源中 current-agent 選2個 direct、拒絕1個不適用身份特化，其餘8個保留 unassessed。正式 watcher／complete 成功，未使用舊 fixed-pool choices。

這是已讀來源的 current-agent trial，並非 blind selector 或獨立任務泛化實驗。來源 evidence-chain-validation 與 evidence-model 的全文已讀回，candidate/readback/pre-adoption SHA 一致；採用為本次方法參考，具體用途是分開「找到→讀回→適用→採用→驗證結果」並保留未覆蓋 claim。兩者仍為 candidate-intelligence；metadata 的 scope/replaces 為 unknown，不推定全域有效或無取代關係。

[決策與版本紀錄](phase3-consumer-trial.json)、[原 request](phase3-consumer-request.json)、[原 current-agent 回應](phase3-consumer-raw.json)、[watcher](phase3-consumer-watcher.json)、[complete](phase3-consumer-complete.json) 可覆核。沒有取得實際 runner model/tokenizer，tokens=null；未以 runes 替代。故只算受控讀回，不能當正式 token admission 或兩個獨立任務 gate 之一。

## 安全邊界走讀

以下為 root 對 workflow 契約的人工走讀，不是 runtime integration 成績：

| 情境 | 本次處置與限制 |
| --- | --- |
| route 已選定仍查經驗 | 本 trial 實際執行；未以 detector miss 為條件 |
| required_read 不在候選前列 | bootstrap、validation 與 workflow 入口先讀；候選沒有控制必讀順序；完整低排名負面重播待驗 |
| 未解 blocker | tokenizer／專案證據缺口採用前後都保留，未因找到來源結清 |
| 相似但不適用 | identity-coupled candidate 本次不選；不是本次文件驗證的前提 |
| A→B | isolated copy 增加格式註解、SHA 改變；比對正文重驗 B，沒有把 checksum 變化當結論失效；真實來源變動採用仍待驗 |
| 重查／合理略過／不當略過 | 契約要求前提、checksum 與先前 trace；尚未作真實多事件序列，未標通過 |
| supersession | 現有 maintenance-governance 明列 research-era rule 被 Stable Maintenance 取代；舊專案檔不在本 repo，僅確認取代宣告，沒有冒稱完整歷史／當前採用重播 |
| 缺索引／stale／timeout／feature off | 操作正文保留直接來源 fallback 與原 blocker；Phase 2 CLI 負例不冒稱 Phase 3 consumer 負例已完成 |

## 驗證與剩餘工作

runtime refresh / validate 成功，更新 generated runtime-index.sqlite 的既有 source checksums；runtime.db owner contract 未變，不手改 DB 或新增 mirror。既有 runtime.db age warning 未隱藏。新增 Markdown 連結、入口與 template index 已反查；本段沒有 CLI code 或 binary 變更。

接續完成兩個獨立真實任務與完整負面 consumer 回放；取得 R11/O03 專案版本／指紋證據；固定實際 runner model/tokenizer 與完整 message 成本。未取得的項目維持 not evaluated。全文補讀成本亦未知，不宣稱任務總成本受控。原始暫存資料最多保留7日，repo 只保留最小去敏紀錄。
