# Phase 3：實際 workflow 查找、讀回與處置

## 結論

本輪使用目前 agent，沿正式 prepare → watcher → complete 完成兩項不同的實際調查，並在第二項發生 transport failure 後重查。留下找到、來源讀回、前提核對、方法採用／具體修法拒絕與驗證結果。沒有新 selector 方法、模型 API、計數工具工程或自動採用層。

Phase 3 維持未完成：目前主機修正／版本核對不可達，直接引用的 ownership 遷移已另行讀回，但候選來源 supersession 的完整 consumer 重播仍未結清。兩個有範圍的決策鏈不自動等於完整 gate 通過；未知不填零。

## 任務一：runtime.db 年齡警告是否代表內容失效

真實觸發是本 repo 的 runtime validate 仍回報 compiler-age warning。先讀 sd-validation／既有 bootstrap 契約，列出「warning 本身能證明什麼」的 claim，再查 intelligence。11 個候選中選 evidence-model，回讀全文並核對 candidate／readback／採用前 SHA 一致；scope/replaces 缺值保持 unknown，candidate-intelligence 不升格。

採用其 evidence-scope 方法：回到 nativeCompilerMetadataCheck，確認 age>24h 分支回傳 warning 而非錯誤。實際 runtime validate 成功；針對既有 old-metadata、invalid-JSON、missing-table 三個測試執行驗證，皆通過。這支持「age alone 不是內容失效判定」，不支持整個 DB 在語意上都最新。因此拒絕僅為年齡警告自動重新編譯／修補；沒有改 DB 來把 warning 消掉。

重複檢查同一 claim 時，實際核對來源 SHA／任務前提未變，引用既有 trace 略過第二次 query。required_read 照常保留，沒有讓搜尋排序決定必讀來源。

## 任務二：原生載入已知案例能否直接重用修法

本輪找到 lesson 指向的原始專案診斷、對應產品報告及 focused adapter test 原始檔。三者 SHA 留存，原文與具體平台／套件 tuple／堆疊／產品細節留該專案，不複製進 reusable docs。R11/O03 的「找不到來源專案證據」缺口已補；「當前環境是否匹配」仍是另一件事。

正式 lookup 的64個候選中選既有 native-load lesson，回讀全文及專案證據。其方法是先比對失敗階段、指紋、版本及 scoped repair，並分開載入、生成與產品品質，不能僅看錯誤碼或歷史載入成功。

本機 overlay 明確指定產品 code/runtime 在遠端 Windows，而 Mac 只有操作與分析文件。依既有規則做一次有 timeout 的唯讀 SSH 核對；連線在遠端命令執行前逾時（exit255）。故當前 tuple／修正存在性／focused tests 皆未觀測，沒有把本機舊測試原始檔當成當前測試 PASS。

決策：採用 known-case-first 調查方法；**拒絕直接套用舊 backend 當本次修復**，維持當前適用性 unknown。沒有重新載入模型、重跑已知崩潰 baseline、重啟服務、改遠端配置或切換雲端。未產生新的產品／翻譯／版面 PASS，也不將這次 retrieval readback 說成其他 active plan 已推進。

### 失敗類型改變後的實際重查

當未知從「native 修正適用性」變成「遠端 readback 不可達」，另跑 production query，39 個候選中選 evidence-collapse-point，讀全文確認 collapse point 不等於 root cause。第一個缺證節點是 SSH readback；當前 tuple／修正／載入均 not-observed，不能反推模型失敗。

同時回讀 stale-peer lesson，拒絕套用其 discovery refresh／durable job policy：該方法需要 cached TCP peer、官方 discovery 及 durable job，目前固定 Windows SSH 主機並無這些前提。拒絕理由是適用條件，不是來源黑名單。

## Consumer 安全回放

[負面結果](phase3-consumer-negative-results.json) 使用實際 CLI subprocess；producer 為明列的合成 fixture，不是假稱 fresh model 回應，也不加入兩項真實任務成績。

| 情境 | 實際處置與證據範圍 |
| --- | --- |
| route 已選定仍查經驗 | 兩任務都由 sd-validation consumer 執行 lookup，不要求 detector miss；不宣稱 detector 自動命中這段對話 |
| required_read 排名／候選缺席 | intelligence pool 不含 workflow 必讀契約，但事先讀取且 feature-off 仍直接讀；沒有搜尋前幾名就略過 |
| blocker 保持 | 當前 Windows tuple／修正保持 unknown；候選與方法採用沒有結清它 |
| 相似不適用 | 真實 requery 回讀後拒絕 stale-peer 方法；runtime query 亦拒絕身份特化 |
| 缺索引 | 實際 prepare 失敗，無成功 lookup；consumer 回到直接來源並保留原 blocker |
| 無答案 | nonce 的空 pool、空 selected 正常完成，保留 absence guard；人工控制不估自然無答案率 |
| timeout／晚到 JSON | 真實 watcher timeout 後，晚到合成有效 JSON 仍被 complete 拒絕；沒有舊成功 fallback。current-agent 無另外執行的 child，因此不宣稱測到 host child interrupt |
| 來源 A→B | isolated corpus 的選中來源加格式註解；舊 bundle 被拒絕。fresh index／prepare／readback B 後，確認正文前提不變才保留方法用途；未推論結論失效 |
| 再次變動 C | B 完成後、重用前再改版本；B bundle 再次被拒絕，consumer 停止沿用，待 C 重驗 |
| 合理略過／不當略過 | 實際相同前提／SHA 核對後略過重複 query；受控「已找到 native lesson 就略過核對」提案被拒，跟著執行實際 requery |
| feature off | 受控事件不呼叫 lookup，直接讀 canonical workflow，保留 required_read／unknown；未改任何 global feature flag |
| supersession | 實際讀回 token 限制 A 與使用者批准的 B，只用 A 追溯歷史、B 控制本次執行；這不是完整 source-level replaces 重播，候選來源的 supersession gate 仍未驗收 |

另以實際 workflow-context CLI 做 [受控 probes](phase3-context-controls.json)：nonce／literal implement 為 no-match，canonical user signal SDK 會命中 software-delivery；保留全部結果，沒有為了命中改 routing。這不冒稱本對話由 detector 自動選定。

既有 WorkflowGate、advisory-never-lock、Bridge manual-lock bypass 與 discovery-provider-never-block 測試通過，見 [原測試紀錄](phase3-existing-gates.txt)。這只支持舊 activation/gate regression，不取代本輪 consumer 行為證據。

### 實際 ownership 遷移讀回

本輪 evidence review 沿既有 stable inbound link 讀回舊 review-checklist。歷史 blob 自稱 canonical，當前 stub 明確宣告已遷移；實際追到 cross-cutting owner，讀取 Contract Governance／Reusable Guidance Boundary，依現行來源審核 claim 與去敏邊界，拒絕把歷史正文重新當現行 owner。見 [讀回與 findings](phase3-supersession-readback.json)。這是實際 direct-reference 遷移，不冒稱模型已能自行辨識所有 metadata replaces。

## 產物與驗證

[三次查找決策](phase3-dogfood-decisions.json) 綁定 binary/index/request/raw SHA、候選與回讀版本、scope／replaces unknown、方法用途及處置。`phase3-dogfood-<task>-request.json.b64` 是完整原始 request bytes，解碼後驗 SHA；raw 保留原字節，外層 CLI receipts 只正規化本機 transport 路徑。所有查找為 source-informed current-agent 執行，非 gold-blind 或泛化實驗。

[事件紀錄](phase3-consumer-events.json)、[feature-off](phase3-feature-off.json)、[runtime validate](phase3-runtime-validation.json) 分開存放。原始本機 transport 最多7日，repo 只留必要去敏紀錄與原始通用 request／回應；不建立對話庫。

本輪不改 production code、CLI binary、workflow 契約或元資料。需續做：Windows 可達後核對當前 tuple／repair 與必要 focused validation，以及候選來源 supersession 的完整 consumer 重播。遠端 unknown 是本次拒絕舊修法的有效終態；只有要進一步宣稱當前修正適用／驗證成功時才需主機證據，不把它擴成所有 workflow 的新 blocker。計數依使用者指示暫緩，不是阻擋項。沒有把 Phase 3 標成完成。
