# v3 現有 subagent runner — 執行前紀錄

使用者已選擇現有 agent／subagent runner，不需另購 API。本輪依 [更新成本契約](phase2-v3-semantic-contract.md) 使用六個 fresh subagents，每題一個、`fork_turns=none`，不設定 model override。Root 保有 gold 用於外層評分，ranker 不接收本對話、case ID、來源映射或其他結果。

六份原 request 的 bytes 與 SHA 維持 [manifest](phase2-v3-semantic-manifest.json)；只複製成 /tmp 下依 file SHA 命名的中性輸入，避免以檔名曝露 case ID。子 agent 只能讀 assigned request；不讀 bootstrap／repo／來源全文或使用搜尋。這是使用者指定的 bounded JSON ranking 子任務，不是 repository 修改任務。

固定調度 envelope：讀完全部 request，遵循其中 system；對全部 opaque IDs 回傳 JSON `ranking` 完整 permutation。可另附最多八筆 `assessments`（id、direct/background/not-applicable/unknown、簡短理由），不改 primary ranking-only 評估或 P0/P1 政策。只做一次 attempt，不要求修正、縮集合或重試。

原始 final text、agent ID／狀態、request SHA、dispatch／外層觀測完成時間、SHA 及 parser 結果都保存。精確 model snapshot、usage、價格若工具未回傳，一律 unavailable；不以當前 agent 自述猜型號。模型額度消耗存在，不稱零成本。額外 API 呼叫數為 0，沒有新採購／計費服務。

現有 subagent 的執行環境仍提供系統指令與工具；fresh context 不代表完全無先驗或能證明未受訓練資料曝光。工具 trace 只檢查實際讀取範圍，不把「無對話 fork」說成獨立 holdout。這六題仍是曝光診斷 pilot。

完成後用原 pool／opaque 映射驗證 permutation，再套同一 P0/P1；不得改 prompt、top-k、權重或摘要政策。Parser 失敗／timeout 保留原文且不偷偷 fallback。評分與判讀寫另份 run report；Phase 2 完整 gate、候選審閱與資訊充分性仍另驗。
