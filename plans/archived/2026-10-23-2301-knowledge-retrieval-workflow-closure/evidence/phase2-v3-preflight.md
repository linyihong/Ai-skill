# Phase 2 v3 preflight — 固定候選與離線分離比較

本輪完成 [v3 協議](phase2-v3-protocol.md)、候選凍結、離線比較及語意 request 準備。**語意重排序尚未執行；Phase 2 品質 gate 維持未完成，沒有 production 採用。** 不把此次準備稱為 v3 語意實驗完成。

## 固定產物與來源

方法先固定在 `47ac4f19`，test-only harness 為 `4ea4130fce452ef7bb61a64dc1364370daf96832` 的 `scripts/ai-skill-cli/internal/app/retrieval_fixed_pool_test.go`。沒有 production Go source 或 binary 變更，不需新增向量 DB 或 runtime surface。

[fixed-pools.json](phase2-v3-fixed-pools.json) 保存完整身分、每題 membership 和 324 個不同 atom 的 source metadata / title / summary；它是可重建索引快照，不取代 canonical source。匯出來自固定 v1 binary 的實際 `--limit 64`，非手動用 gold 補池。

- Corpus 保持 `5f12fc3906051e40228959e9e6bfaed714e0658e`；NEW index / generator 身分沿用 v1 manifest，SHA `1221f2402084ec4454e506f05c77404d7773bc683737010cf890bb556af24cfc`。
- Binary 從 `df95535b` 已提交的 darwin/arm64 產物取出，SHA `d59cc24c93a7b0f448b1fba1f3e61375610226edca63ffbe32ae30ad1b19cf20`；version readback 隨 pool 保存。
- 26 題累計 812 個 atom occurrences，每題至多 64；各 selector 的 membership 不變。所有 18 個歷史正例及 5 個 overlap 正向錨點均已入池（各題一個 gold，不外推全庫 coverage）。
- O03 pool 有 49 atoms，正向在 v1 rank 3；同一 pool 的 v2 權重使其排到 rank 33。這次完整保留，沒有先取五筆再要求模型救回。
- S1 / S2 的離線 reference query 在全部 26 題上與固定 v2 trace 的前五 source／score 一致；它只驗證 reference 方程與舊證據對齊，不宣稱使用 v2 CLI 重播了不同輸出政策。

公開 pool 與本機映射 manifest 的 JSON 字串內 hyphen 使用標準 Unicode escape；這避免既有公開檔名 `ondisk-encrypted-…` 被原始文字 scanner 誤判為 API key。解碼前後資料物件一致，負數 rank 不變，file SHA 與 manifest 關係已重算；重播比較應按 JSON 解碼物件與 membership SHA，而非要求 pretty-print bytes 相同。

每題 membership SHA 與 dataset SHA 隨各比較 row 保存。全文／索引覆蓋未改；R08 的 YAML 與歷史 lesson 依各自來源保留，不改 gold 或預設 YAML 不相關。

## 分離排序與 packer 的結果

[offline-comparison.json](phase2-v3-offline-comparison.json) 共 156 rows：26 題 × S0 / S1 / S2 × P0 / P1。S0 為 v1 排序；S1 只改 v2 bm25 權重；S2 再加 exact-phrase 優先 tier。P0 完整摘要，P1 摘要限 128 runes；其他 QueryResult 欄位與短 reason 相同。各 packer 依順序填入，下一筆超預算即停止，沒有硬設 top1 或固定筆數。

| 選擇 | 組裝 | Diagnostic hit / 14 | Diagnostic MRR | 固定比較 / 4 | Overlap 正向 / 5 | Confuser / 6 | 最大 runes | 截短 occurrences |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| S0 v1 | P0 full | 14/14 | 0.880952 | 4/4 | 5/5 | 6/6 | 1982 | 0 |
| S0 v1 | P1 compact | 14/14 | 0.880952 | 4/4 | 5/5 | 6/6 | 1802 | 47 |
| S1 weighted | P0 full | 13/14 | 0.892857 | 4/4 | 4/5 | 6/6 | 1967 | 0 |
| S1 weighted | P1 compact | 13/14 | 0.892857 | 4/4 | 4/5 | 6/6 | 1802 | 51 |
| S2 phrase tier | P0 full | 13/14 | 0.892857 | 4/4 | 4/5 | 6/6 | 1967 | 0 |
| S2 phrase tier | P1 compact | 13/14 | 0.892857 | 4/4 | 4/5 | 6/6 | 1802 | 51 |

所有組合的人工負面仍 0/2。這是固定池內、來源去重後的最終 payload 指標；selection 前五 source 與原基準「前五 atoms 後去重、不補位」分別保留，不宣稱兩種協議完全相同。各組先報 pool coverage，無模型也沒有新增正例。

本資料最多可放 3 個來源。P0 可放 1–3 筆（兩個空結果控制為 0）；P1 多數非空結果為 3 筆，原生精確案例可只有 1 筆。R08 及 O02 在多筆輸出下保留正向來源；O02 的 confuser 同時仍在，不能把找回正向當成已選對。

可支持的判斷：

- 同一 S0 排序下，完整摘要和多筆輸出已能在 proxy 預算內找回指定來源，無需硬限 top1。成本計算共用短 reason、去重與結果 JSON；不是直接把原 v1 五筆未改 payload 從 3503 降到 1982。
- P1 這次確實發生摘要截短，但在此資料未增加 gold hits，也未降低 confuser；不能把截短當成資訊充分性驗證。
- 固定相同 pool / packer 後，S1 相對 S0 丟失 R03 / O03 指定來源；MRR 可升而 recall 降，不能單選 MRR 宣告改善。S2 對本次最終 payload 無額外增益。
- 三種排序均留 6/6 confuser。有限的詞彙排序與組裝實驗沒有解決焦點／否定／適用範圍問題，語意排序 pilot 仍有可測假說。

每筆 report 的 `information_sufficiency` 都是 `not evaluated`。需要審閱 trigger、適用／不適用條件、機制與回讀入口；即使 P0 保留完整 indexed summary，也不是完整來源內容。版本／checksum 在 pool sidecar，未混入 results 代理成本；真正 workflow context、tokenizer 成本與採用安全仍待後續。

## 語意 pilot：已有可送審輸入，尚無模型結果

[semantic manifest](phase2-v3-semantic-manifest.json) 綁定六份 request 的 file SHA、固定 pool SHA、opaque-to-source 本機映射。模型 request 本體只有 system、query、opaque ID / title / summary，不含 case ID、source path metadata、gold / excluded labels 或 lexical rank／score；映射不應傳給模型。

| 案例 | 固定候選來源數 | 輸入檔 runes（只作診斷，非 tokens） |
| --- | --- | --- |
| R03 | 5 | 1868 |
| R08 | 55 | 21336 |
| O02 | 64 | 21654 |
| O03 | 49 | 15939 |
| O04 | 57 | 19084 |
| O06 | 23 | 8307 |

[成本契約](phase2-v3-semantic-contract.md) 提案為 6 次總 attempts、每次 input ≤16384 / output ≤4096 實際 tokens、每題 timeout 30 秒、總 model wall time ≤180 秒、費用上限提案 USD 2。**這些不是已核准支出**；runner／固定模型／tokenizer／價格及資料目的地尚未確定。任一輸入若超實際 context cap，整題停下，不能為放得下而刪掉 O03 等指定來源。

目前無外部模型呼叫，usage / latency / semantic relevance 均 `not evaluated`。已透過問題請使用者指定既有 runner／模型與總費用上限；缺少這些資訊只影響語意執行，不把離線工作混稱 blocked 或全計畫完成。不能以目前 agent 已看過 gold 的人工重排冒充獨立語意實驗。

## 驗證與重播

完整 `go test ./...` 及 pre-push preflight 通過；新 harness 正常 CI 只跑 packing controls，真實 pool export 必須顯式 env 啟動。Packing controls 驗證多筆對照、Unicode / JSON escaping、identity/status 保留、來源去重與超額拒絕。另以 Go 標準 JSON encoder 重算 156 份保存的最終輸出，rune 數逐題相符且均 ≤2000；獨立核對 pool membership、source unique、hits/confuser、reference case 覆蓋與六份 request hashes。

```text
AI_SKILL_RETRIEVAL_V3_OUT=<new absolute output directory>
AI_SKILL_RETRIEVAL_CORPUS=<clean frozen corpus checkout>
AI_SKILL_RETRIEVAL_INDEX=<v1 NEW extraction snapshot>
AI_SKILL_RETRIEVAL_INDEX_MANIFEST=<matching v1 manifest>
AI_SKILL_RETRIEVAL_QUERY_BINARY=<v1 df95535b darwin/arm64 binary>
AI_SKILL_RETRIEVAL_QUERY_SHA256=d59cc24c93a7b0f448b1fba1f3e61375610226edca63ffbe32ae30ad1b19cf20
go test ./internal/app -run '^TestRetrievalV3FixedPool$' -count=1 -v
```

變數須 export，命令在 CLI module 執行；輸出既有檔案拒絕覆寫。模型 request 的 title / summary 來自固定 pool，不是新查詢服務。若更換模型輸入、集合深度、prompt 或 packing 政策，另版留下變更，不能悄悄覆寫本輪。

v2 CI [40033981207](https://github.com/linyihong/Ai-skill/actions/runs/40033981207) 的三平台 tests / artifacts 已全數成功；新 test-only harness 的 CI 為另一次執行，不能沿用 v2 結果宣稱已全綠。Phase 2、production 搜尋改善、資訊充分性及語意增益均尚未通過。
