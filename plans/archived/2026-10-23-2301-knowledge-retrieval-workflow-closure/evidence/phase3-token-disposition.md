# Phase 3：token 計數暫緩

使用者 2026-10-25 明確澄清：「可以先不用記 token 數，用這個 agent 去做沒問題」，並詢問可自行計算的工具。這取代本計畫先前將 tokenizer 身分當正式 consumer 執行前提的要求。

決定：使用目前 agent 繼續 Phase 3，token 計數為 `deferred-by-user`，不再阻擋。保留既有完整 retrieval payload ≤2000 runes，required_read、適用性、來源版本、supersession 與 blocker 檢查照常。舊 trial 的 raw 與歷史結果不回改，原本缺計數不會因此變成實際 token gate 通過；本輪執行不再因同一缺口被降為只能試跑。

可選工具為 [tiktoken](https://github.com/openai/tiktoken)：固定 encoding 與版本可做可重現的本機文字計數。對應模型 encoding 未知時只能標 reference count；不等於 agent 的實際輸入／輸出用量，也不包含未取得的系統／工具／訊息封裝。當時僅提供工具建議；後續使用者已授權安裝，實際安裝與參考量測見 [token counter](phase3-token-counter.md)，未增加模型 API 呼叫。

主 plan、workflow 操作與 template 已同步。舊 phase3-consumer-protocol 與 trial 的 token 限制屬當時狀態，現在以此修訂為準。後續聚焦真實任務採用與負面 consumer 回放；R11/O03 專案證據仍待取得。
