# Phase 2 — 最終輸出為中心的 26 題診斷協議

在新 26 題選擇前固定。依使用者對 e73d1f3d 的最新指示，**明確取消全池 unknown 清零／每筆候選分類作為 gate**，以準備輸出來源逐筆審閱取代。舊實驗結果與原 corpus／dataset 不改；這是主計畫驗收範圍的具名修訂，不是默默忽略。

## 審閱對象與方法

- 當前 agent 讀固定 pool 的所有 distinct cards（title／完整 indexed summary／source identity）及 26 題 query／task_context；依問題焦點選 direct，沒有 direct 可留空並說明。每題只記真正判斷過的處置，未標記維持 unassessed，不製造全池分類。explicit unknown 可作終態，記下一步或不輸出的理由。
- 這個對話已知 gold／confuser／既有錯誤；新結果只能算 **current-agent、known-gold exposed diagnostic**，不是 fresh／blind runner，不與前輪六個 subagents 混作同方法或單變項因果比較。不得從 gold 直接建選擇清單；選擇理由必須由實際讀到的候選內容支持。
- 固定 broad pool、corpus、原 dataset labels；逐題實際檢視 pool 候選，不加入池外 gold，不用 confuser 黑名單。選擇不必對全池排完整新 permutation。未選候選保留數量、完整 pool／audit 入口。
- 每個準備輸出來源均讀回固定 Git blob、核對 checksum 與當前版本，並審其與 query 的關係／必要前提／摘要限制。身份一致不是適用通過；不符合 scope 可以不輸出或只標待驗證線索。審閱 trace 與未選理由留外層。

## 真正送入模型的 payload

固定 compact JSON object，頂層欄位為 `guard`、`state`、`pool`、`unassessed`、`unknown`、`deferred`、`trace`、`results`。

- guard 固定 `Candidates only; read sources and verify applicability; empty does not establish absence.`，計入預算。
- state 為 `direct-pending-verification`、`pending-assessment`、`no-direct-in-reviewed-candidates` 或 `no-candidates-in-fixed-pool`。有 direct 時仍回報 unknown；全池未評不以空清單推定無知識。
- pool／unassessed／unknown 為數量（unknown 包含未評估），deferred 為因預算未輸出的 direct 數；trace 是可交給工具解析的固定 `final-output/<case-id>` 引用，完整 audit map 在外層同名 results 記錄，模型可要求補讀。
- 每個 result 僅含 `source`、`kind`、`status`、完整 indexed `summary`，不修改來源 status／type，不將 direct 變成可信度。P0 優先，不截摘要、沒有 P1 fallback；BM25、priority、confidence、context_cost、原 atom ID、checksum／revision、hash、評分／逐筆理由只留外層工具 audit。模型使用來源前仍必須回讀；版本由 audit 引用解析核對，不能省略核對。
- **整個物件**（含 key、guard、必要狀態與來源）以 Go JSON.Marshal 計 ≤2000 runes。外層 trace 不送入模型；若之後送入，另計成本。實際模型 tokens 仍 not evaluated。
- 保持 agent 選擇順序，逐筆加入完整 result；下一筆超額即留在 deferred 清單並更新計數，必要時回退已填來源以讓整個物件符合上限。無自動挑更短摘要、無按 gold 補位。品質在最後 payload 上評分，deferred 損失照實記錄。

## 資訊充分性與評分

26 題各列任務需要的資訊、選擇理由、最後摘要實際提供什麼，以及 `sufficient-for-next-step`／`partial`／`insufficient`／`unknown`。sufficient 只表示能引導下一個查證步驟，**不代表能直接解決真實問題或已適用／採用**。若完整答案需來源正文或任務證據，明列缺失；空池人工控制可記正確 abstention，但不當自然無答案率。

分開報 retrieved 原 confuser、selected-output confuser、source hit／MRR、資訊充分性、unassessed／explicit unknown、budget deferred、來源讀回與 scope 未確認數。採用仍 not evaluated；不藉消滅 unknown 過 gate。

這輪沒有新外部 API／subagent attempt，不把整段含工具與審閱的對話時間算成受控推論 latency。成本證據僅為凍結 payload 的工程代理值。全 26 題完成本診斷後，才依主計畫補 intelligence、metadata／graph／索引生命週期、受控 runner 與 Phase 3 採用鏈；Phase 2 仍不因此自動完成。
