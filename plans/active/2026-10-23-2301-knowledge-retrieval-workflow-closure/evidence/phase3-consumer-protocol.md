# Phase 3：sd-validation consumer 執行範圍

> 2026-10-25 更新：token 執行前提已由使用者明確暫緩，見 [最新處置](phase3-token-disposition.md)；以下保留事前範圍歷程，不再以缺 tokenizer 阻擋目前 agent。

## 啟動與相容性

使用者已接受 `f173083a` 的縮限 Phase 2，允許進入 Phase 3，沒有新增驗收項目。本輪沿用主計畫 Phase 3 gate，不開第二輪 selector 探索。

| 項目 | 決定 |
| --- | --- |
| 目標 | 讓 sd-validation 在形成驗證方案前找經驗、讀回並留下用途與採用／拒絕／unknown 理由 |
| Canonical owner | workflow/software-delivery/validation.md；操作子文件與專案 receipt 模板，不建立通用 trigger framework |
| 重用 | runtime lookup prepare/watch/complete、現有 required_read 與 blocker；不改搜尋、selector、路由、權威順序 |
| Bridge | Phase 0 無重構必要證據；保留 miss advisory，workflow route hit 仍可查經驗 |
| 連動 | validation、workflow README、template index、plan/evidence；索引若已有 source row 則 refresh，不手改 DB |
| 範圍外 | 新模型／API／vector、全池 unknown 清零、相似度提升來源權威 |
| 語言 | 目標 workflow 為繁體中文混必要欄位；新操作正文採繁體中文，既有英文 template index 延續原語言 |
| 測試策略 | 文件行為以情境走讀及真實 consumer 紀錄驗證；CLI failure coverage 沿用 Phase 2，不寫鏡像文件字串測試 |

主計畫 Q1/Q3/Q4/Q5 已由 Phase 0/2 回答；Q2 初版限制已量測，Q6 語意擴充仍屬條件式 Phase 4。沒有新的架構 blocker。R11/O03 原始專案證據及實際 runner tokenizer 身分仍未取得：不推測、不以通用 lesson 代替，對應 gate 保持未驗收。

## 驗收與紀錄

先執行現有 required_read，記當前 blocker，再決定 lookup/requery/skip。輸出來源逐筆確認 domain、project、前提、證據、supersession、lifecycle 與 checksum；lookup complete 的來源新鮮只證明那一時點的版本，不取代採用前重驗。

沿用既定情境：route hit/miss、required_read 低排名、blocker 保持、相似不適用、歷史追溯與當前採用、A→B 變更、lookup/requery/skip、stale/unavailable/timeout/無答案、feature off。人工情境只算情境證據；至少兩個獨立真實任務的完整鏈才算主 gate 通過。文件存在、helper pass 或重播舊決策均不足。

實際模型 message 的 tokenizer/model/version 與納入範圍必須可辨識，≤2000 實際 tokens 才能宣稱 token gate 通過。缺資訊時不以 runes 冒充；可做受控 readback trial 或回到既有直接來源查找，但不宣稱已完成正式 consumer admission。來源正文補讀另列成本；不把需要送入模型的說明藏到 tool metadata。

本輪 raw transport 保存在本機暫存目錄，保留最多7日；可覆核去敏決策紀錄放 plan evidence，長期隨 repo 保存。未採用候選只保留數量與 trace，不保存全量對話。外部專案證據留該專案，這裡只保留去敏引用與處置。
