# Phase 2 readiness — 全資料 coverage 與 P0 preview 盤點

結論：兩項審閱要求已固定為離線 preview 契約並有 prototype evidence；全資料相關性審閱、資訊充分性、成本／時間及索引生命週期尚未通過，Phase 2 未完成。沒有接 Phase 3 workflow，沒有新增 agent／API 呼叫，沒有改 production CLI 或既有 raw／結果。

## 身分與範圍

[契約](phase2-readiness-contract.md) 與 test-only exporter 先提交於 `ac79c829`，再執行 `AI_SKILL_RETRIEVAL_READINESS_REPORT=<new output> go test ./internal/app -run '^TestRetrievalPhase2ReadinessExport$' -count=1`。新 [results](phase2-readiness-results.json) 保存 scorer revision／SHA、所有輸入 SHA 及逐題 coverage／preview。這是對固定 26 題的 readiness 盤點，不是 26 題新語意評估。

## 全資料 coverage

| 範圍 | 題數 | 來源出現次數 | 已 assessment | 未 assessment | explicit unknown | unknown 總數 |
| --- | --- | --- | --- | --- | --- | --- |
| 原六題 | 6 | 253 | 33 | 220 | 1 | 221 |
| 其餘二十題 | 20 | 517 | 0 | 517 | 0 | 517 |
| 全部 | 26 | 770 | 33 | 737 | 1 | 738 |

來源在不同 query 重複出現分別計數，不能把某題 relevance 沿用到其他題。812 是 atom occurrences，770 是每題按 source 去重後的來源出現次數，不可混用。

Disposition 維持 direct 9、background 15、not-applicable 8、unknown 738。這只是已有 33 筆 assessment 的範圍，不是 738 筆無關判定。二十題沒有捕捉 semantic response，trace 明列 unavailable；其中 R19／R20 固定池為空，不代表全庫沒有知識。

Consumer preview prototype 逐題提供候選數、已評／未評數、explicit unknown 與 unknown 總數，即使已有 direct 也保留 unknown。空清單分成 `pending-assessment`、`no-direct-in-assessed-pool`、`no-candidates-in-fixed-pool`，皆不宣稱 knowledge absence；direct 仍為 pending verification。完整既有 decision trace 由結果檔路徑與輸入 SHA 綁定。這是離線範例，真實 consumer 是否讀取／處理仍 not evaluated。

## P0 與成本缺口

預設 prototype 僅使用原六題 P0 完整摘要；其餘二十題空 preview 是未評估／池空的顯式狀態，不作無答案品質評分。P1 仍留在原實驗中，未改成預設，也沒有為降低成本重排或丟來源。

| 成本範圍 | 最大 runes | 判讀 |
| --- | --- | --- |
| 原 direct-only P0 results array | 1922（R08） | 原凍結的 ≤2000 results-array gate 保持通過 |
| 本版 consumer-visible JSON envelope | 2208（R08） | 若以整個 envelope 套用 2000，則 1/26 超額；其他題較短不能抵銷 |

新增 envelope 包含 coverage、verification、packer、knowledge_absence 與 results，側錄 audit／hash／完整 decision trace 仍在外層 report。2208 只對應本版 prototype schema；不能冒稱所有未來 consumer metadata 或實際 tokens 都已計入。若整份 trace 也載入模型須另計成本。

這個缺口不回改舊子實驗；也不以切 P1、調 top-k 或靜默裁掉 direct 來掩蓋。後續需先凍結 consumer-visible 整體成本範圍、序列化與門檻，再測組裝方案；目前不可宣稱完整 context budget 通過。

## 本輪驗證與後續順序

- `go test ./...` 通過，新增 controls 區分池空、未評估、explicit unknown、已判不適用、有 direct 但仍未評完；來源數量守恆。
- 正式 opt-in exporter 通過；獨立 Python 重算 26 題 counts、輸入 hashes；獨立 Go 重算全部 26 份 envelope 及 results-array runes，確認只有 R08 envelope 超過 2000。
- 前輪 scorer CI [40055737331](https://github.com/linyihong/Ai-skill/actions/runs/40055737331) 已完成 success（Linux／macOS／Windows tests 及 artifacts）；本輪新 exporter 的 CI 需獨立追蹤，不以舊成功代替。

下一個執行順序依契約：另版完整 assessment 協議（舊每題最多八筆理由不夠）、全資料逐題 relevance／資訊充分性審閱、intelligence 比較資料、metadata／graph／索引 lifecycle 證據、受控新 runner 時間及整體 context 成本。不要用新離線處理時間修復舊模型時間缺口。Phase 2 gates 通過後才安排 Phase 3 source-read／適用性／採用鏈；主計畫未勾選的 gate 均保持原狀。
