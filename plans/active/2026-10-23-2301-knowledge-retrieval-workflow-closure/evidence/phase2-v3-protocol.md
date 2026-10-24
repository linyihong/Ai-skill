# Phase 2 v3 — 固定候選、選擇與輸出組裝分離

本輪先完成離線比較與語意 pilot 的可覆核準備。v1 / v2 已拒收，原紀錄與未完成審閱清單不回改；Phase 2 gate 維持未完成。此協議在 v3 資料產生前固定，不預設向量 DB，不修改 production query。

## 固定候選集合

- Corpus / NEW extraction index 沿用 v1 manifest；26 題沿用 baseline-cases 與 overlap-controls-v1。以 v1 pinned binary `df95535b` 的實際 `runtime query --limit 64` 匯出每題候選，開始／結束核對 binary / index SHA 與 clean corpus。
- Pool 是前 64 atoms，完整保存 membership 與原序，按 source 去重用於選擇。不得只取舊 report 的前五、不得按 gold 補來源。若正向未入池，記 candidate-generation miss，該題不得歸功或歸罪重排序。
- 每題 pool membership SHA、query、dataset SHA、corpus / index / binary 身分凍結；排序方法只能重排同一組來源。Source metadata / title / summary 來自同一 index；來源內容仍以 corpus checksum 為準。
- 先核對 O03 正向是否在 pool，並記 v1／v2 位置。禁止為修正 O03 單獨擴池；若需改深度，另版全組重跑。

## 分離選擇方法

| 方法 | 改變項 | 固定項 |
| --- | --- | --- |
| S0 | 保留 v1 lexical 排序 | 同一 pool、index、source 去重、packer |
| S1 | 只使用 v2 的 bm25 欄位權重 | 不套 top1、不改候選 membership、不改 packer |
| S2 | S1 排序加 exact-phrase 優先 tier | 不刪 phrase 以外來源；其餘保持 S1 |
| S3 | 語意排序 pilot，尚未執行 | 同一 pool / cards、同一 packer；不加向量或新來源 |

S1 / S2 的離線 SQLite reference 使用 v2 相同權重和既有 tie-break；在 v2 無 phrase 的題目核對其前五 atom rank／score 與 v2 trace，在 phrase 題核對 phrase tier 與 v2 trace。明確標為 reference comparison，不冒充 v2 CLI 的 top1 consumer 輸出。S0 pool 才是 pinned v1 binary 實際輸出。

所有方法按 source 去重（保留排名第一 atom）；S0/S1/S2 pool membership 必須一致。獨立呈現 selection 的前五 source 分數，不能混同原基準「前五 atoms 不補位」；另保留原 atom 指標作對照，避免改 dedup 協議後宣稱直接改善 v1。

## 固定排序後，只改輸出組裝

同一排序各跑兩個 packer，不預設 top1、沒有固定結果筆數：

- P0 full：原 QueryResult 欄位；兩組共用短 match_reason，不重複 query。依順序加入完整摘要，下一筆超過 2000 runes 即停止。
- P1 compact：完全相同的欄位／排序／reason，只將每筆 summary 限 128 Unicode runes（含省略號）。是否截短在 report sidecar 逐筆記錄；不能暗改 source / status / confidence。
- 兩者都以最終 compact JSON results array 的 UTF-8 rune 數限制 2000，保留合法 JSON。空結果不算 precision=100%，大型 metadata 放不下亦記清楚，不只拿摘要文字計成本。
- 比較同一 S 的 P0/P1，才可歸因組裝；比較同一 P 的 S0/S1/S2，才可歸因排序。實驗不改 runtime CLI，輸出是 offline consumer-payload simulation，非已接入 workflow 的 production 成果。
- 回傳多筆／截短成功只證明工程預算。另列資訊充分性審閱：是否仍能辨識 trigger、適用／不適用前提、失敗機制與回讀入口；未審閱標 not evaluated，不宣稱「充分資訊」。

## 指標與 gate

先報 candidate source coverage，再報 selection source Recall@5／MRR@5、confuser、最後報各 packer 的實際 sources、hits、confuser、摘要截短、runes 與移除原因。R08 的相關 YAML 可另審，但指定 lesson 仍按原 gold 記漏找。

歷史基準的 ≥8/14、固定比較不退步、人工控制 0/2、overlap confuser 0、2000 runes 與完整採用版本審閱不變。離線測試不能代替 CLI latency ≤500ms gate；assembly runtime 與 semantic latency 另報，未測不填 0。資料全已曝光，只作診斷，不宣稱獨立泛化。最後採用仍需完整 26 題及後續 intelligence / workflow / tokens gates。

語意部分限 6 個已曝光診斷案例 R03 / R08 / O02 / O03 / O04 / O06，以同一 pool 比較問題焦點／否定與詞彙排序；不計作全資料 Phase 2 acceptance。執行前需完成 [semantic cost contract](phase2-v3-semantic-contract.md) 的 runner / model / tokenizer / 價格 / 資料傳送確認，尚未完成不得外呼或宣稱語意結果。
