# Phase 2 — 26 題最終輸出與資訊充分性診斷

結論：本輪三項工作已執行：明確縮限審閱對象、凍結模型 retrieval payload、完成全 26 題選擇及最終摘要的充分性評估。**Phase 2 未完成；沒有正式 workflow consumer 接線。**

這是當前 agent 已看過 gold／confuser 的知情診斷，並非獨立方法驗證。全部命中不能當成模型泛化或 production 改善；選擇方法、輸入與 payload schema 也不同於前輪六個 fresh subagents，不作單變項增益主張。

## 交付範圍的明確決定

依使用者對 e73d1f3d 的審閱，主計畫原「全池新增候選逐筆分類／unknown 未結清不得通過」已明確修訂為：**準備輸出的來源逐筆審閱；其餘保留數量、處置與可追查入口。** 未選 unknown 不必清零，explicit unknown 可為本次終態。這不減少全 26 題最終輸出的品質與資訊充分性責任，也不取消 intelligence、索引生命週期或 runner gates。

[協議](phase2-final-output-protocol.md) 先提交 `e2d68c63`；當前 agent 讀完固定池 324 個 atom cards（311 個 distinct source paths，含同來源的不同摘要），再逐題選擇。沒有池外注入、gold blacklist 或全池強制分類。新 [decisions](phase2-final-output-decisions.json) 記錄 26 題順序、顯式 relation、理由、需要的資訊與 unknown 處置；舊原始輸出與評分不改。

[readback](phase2-final-output-readback.json) 記錄 21 份不同準備輸出來源的全文 Git blob／checksum、當前版本與 scope 審閱；21/21 身分與當前 checksum 相符。來源對問題的關係經知情審閱，但真實任務的前提、證據與採用仍待驗。所選共 26 筆 source occurrences，所有實際輸出均有逐筆 source review。

## 凍結並實測模型 retrieval payload

Test-only exporter 與輸入決策／readback 先提交 `5f420095`，再以 `AI_SKILL_RETRIEVAL_FINAL_REPORT=<new path> go test ./internal/app -run '^TestRetrievalFinalOutputExport$' -count=1` 產出新 [results](phase2-final-output-results.json)。scorer revision／SHA、protocol SHA、全部輸入 SHA 均保留。

每題 `model_payload` 是**本輪唯一預定送入模型的 retrieval message**，不是外層 audit report，也不包含既有任務對話／system prompt。模型使用者原始 query/task context 已在任務中；這輪沒有真的向另一個模型送出這些 payload，只有組裝、閱讀及評分。不要把這個離線原型當 production consumer integration。

模型可見：guard、state、pool、unassessed、unknown（含未評估）、deferred、trace，以及每份來源的 source／kind／status／完整 indexed summary。explicit unknown 數可由 unknown 減 unassessed 得到；逐筆理由與覆核紀錄由 trace 追查。候選的原 status 保持不變，direct 不賦予權威。R08 的 YAML 索引 status 是 `new-layer-promoted`，canonical 正文為 active；兩種 surface 不混為相似度或可信度。

工具／紀錄可見但不送入本 message：完整候選池、逐筆 assessments／scope review、source revision／checksum、request／artifact hash、BM25、priority／confidence／context_cost、評分與充分性判斷。trace 精確對應本 results 的 observation `trace=final-output/<case-id>`，可由外層取回 source audit。未來若將這些附加內容載入模型，須另計成本並重新驗收。

整個 `model_payload` 的 Go compact JSON 都計入 ≤2000 runes，包括欄位、guard、unknown 狀態與 trace。**26/26 在預算內，最大 1322（R08）**；全部 P0 完整摘要、0 截短、0 budget deferred。此前 1922 results-array 與 2208 prototype envelope 結果仍成立；本輪換成凍結的模型欄位投影，並非刪除歷史成本或把必要狀態移走。實際 tokens、總任務 context 成本與模型使用額度仍 not evaluated。

## 最終輸出結果

| 指標 | 本輪知情診斷結果 | 邊界 |
| --- | --- | --- |
| 歷史正例 source hit | 18/18 | 原 diagnostic／comparison labels 不改；全屬曝光資料 |
| overlap 正向 source hit | 5/5 | O06 沒有完整 gold，不補入分母 |
| 指定來源 MRR（23 正例） | 1.0 | 人工已知 gold 的選擇順序，不宣稱模型排序增益 |
| retrieved overlap confuser | 6/6 | 候選池原值保持 |
| selected-output confuser | 0/6 | 只代表預標 confuser，非一般無關率 |
| 人工負面控制輸出 | 0/2 | nonce／不存在 status，不代表自然 no-answer |
| 有輸出／空輸出 | 23/26 ／ 3/26 | 空為 R19、R20、O06，各有不同理由 |

770 source occurrences 中僅明列 41 筆 relation：direct 26、background 8、not-applicable 5、explicit unknown 2；其餘 729 未標記。unknown 合計 731 是本輪 coverage 現況，**不是 731 個缺陷或清零待辦**。這與前輪 sparse assessments 是不同方法，不拿數量增減評模型品質。

O06：peer lesson 明確排除 token invalidation；App session／guest bootstrap 兩份線索的服務／裝置前提不明，標 explicit unknown，本次不輸出。處置已結清為「不拿它們當解法；待具體 provider／服務前提再查正式流程」，不是要求強行分類整個池。空 payload 保留 unknown／trace，不能讀成「沒有相關知識」或「已解決重新授權」。

## 資訊充分性：評的是實際留下的摘要

逐題完整判斷在 [sufficiency](phase2-final-output-sufficiency.json)，綁定 results 檔案 SHA。每題都有 information_needed、actual payload 提供的訊號與缺失／下一步。全文 readback 不被假裝算入模型已收到的內容。

| 分組 | sufficient-for-next-step | partial | insufficient |
| --- | --- | --- | --- |
| 23 正例 | 17/23 | 6/23 | 0/23 |
| 2 人工空集合控制 | 2/2（正確留空） | 0/2 | 0/2 |
| O06（無完整 gold） | 0/1 | 0/1 | 1/1 |

`sufficient-for-next-step` 只代表有具體下一個查證方向，不是完整答案、實際適用或任務解決。所有來源都仍須真實 workflow 讀回並核對前提；19/26 不能稱任務成功率。

| partial 題目 | 最終摘要缺的關鍵資訊 |
| --- | --- |
| R06 | blob 改變 base 與用原 manifest URL resolve 的明確解釋在正文 |
| R08 | 歷史 prose 跳步的理由在正文，規範摘要不能代替歷史因果 |
| R11／O03 | 實際 fingerprint／stack／platform／version 比對欄位與隔離方法不在 summary |
| R12 | server-side cancellation 停止、DB readback 與同 key retry 斷言不在 summary |
| R16 | 索引目錄解析與移除連結負例的具體驗法不在 summary |

因此本輪沒有把「全命中且成本符合」當成資訊充分性全部通過，也未偷偷擴 summary 或引入 P1。這六題的下一步是明確 source-read／任務證據補讀，必要時再另版測試資訊組裝；不需要把未選池分類完才能處理。

## 驗證與剩餘 gates

- `go test ./...`、正式 opt-in export 通過。新增負面 controls 驗 budget 不足時保留 guard／unknown／deferred、不截完整摘要、不暗中填其他來源；必要狀態本身超額即回錯誤。
- 獨立 Python 重算 26 題 hits／MRR／confuser、source membership、原 summary／kind／status、不重複來源、所有輸入 hashes；獨立 Go 重算全部 26 個完整 model payload 的 runes。
- 21 份準備輸出來源均核對固定 blob 與當前 checksum，沒有將 identity equality 當成採用通過。
- 前輪 [CI 40056269659](https://github.com/linyihong/Ai-skill/actions/runs/40056269659) 成功。本輪 [CI 40057004671](https://github.com/linyihong/Ai-skill/actions/runs/40057004671) 撰寫時仍執行中，不先當通過。

後續順序：先依六個 partial 與 O06 未知前提處理實際輸出資訊缺口；再依主計畫補 intelligence 比較資料、canonical metadata／bounded graph、索引建立／refresh／修改刪移／stale／rebuild／ordering；另外用新凍結契約驗證 runner 的連續監測、timeout 與可取得 usage。這輪 current-agent 的對話與工具時間不冒充受控 inference latency，未修復舊時間缺口。Phase 3 真實 consumer／來源版本 race／required_read／blocker／supersession／適用與採用鏈仍 not evaluated。
