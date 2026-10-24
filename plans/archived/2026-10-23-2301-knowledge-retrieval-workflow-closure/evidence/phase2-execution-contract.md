# Phase 2 前置：執行檔、索引與候選品質控制

本文件只定義並驗證實驗控制；尚未修改 production query / extraction，也沒有 Q / E / QE 增益結果。舊 baseline 題目與結果保持原檔。

## 實際產物矩陣

| 組別 | 索引產物 | Query 執行檔 |
| --- | --- | --- |
| B0 | OLD index + manifest | OLD binary SHA |
| Q | 與 B0 完全相同的 index + manifest | NEW binary SHA |
| E | NEW extraction index + manifest | 與 B0 完全相同的 binary SHA |
| QE | 與 E 完全相同的 index + manifest | 與 Q 完全相同的 binary SHA |

每個 index manifest 記錄 corpus revision、generator revision、generator module source SHA、來源路徑 / checksum 集合 SHA、generator inputs SHA 與 SQLite 檔案 SHA。Generator checkout 與 corpus checkout 分開；corpus 必須 clean，generator 的 production source 必須已 commit。測試 harness 可在另一個 checkout，不能因此把當前 generator 冒充舊 extraction。

重播只讀指定的 index，不隱式重建。顯式指定 query binary 路徑與預期 SHA，開始及結束皆核對；同時記錄該 binary 的 `version --json`。不從 corpus checkout 自動選 binary，不因 Go source 已改就假設查詢已改。SHA 不符、corpus 不符、index 被改、Q 使用舊 binary、E 使用舊 inputs 等皆拒跑。

漏找分類移除 host-side `runtimeFTSMatchLiteral` 查詢；只依實際 binary 輸出與 source row 是否存在，記為 `source_not_indexed` 或 `indexed_source_not_returned_cause_not_attributed`。Query / extraction 因果歸屬由四組結果比較，不自行模擬新 binary 的 query parser。

## 執行步驟

所有路徑用絕對路徑傳入、環境變數需 export；命令在含新版 harness 的 `scripts/ai-skill-cli` 執行。舊 extraction snapshot 使用舊 production generator source，新 extraction snapshot 使用已 commit 的新 production generator source；測試 helper 可攜帶至對應 checkout，但不得改 generator production source 再沿用舊 revision。

1. 設 `AI_SKILL_RETRIEVAL_CORPUS` 為固定 clean checkout、`AI_SKILL_RETRIEVAL_INDEX_OUT` 為新的 SQLite 路徑；執行 `go test ./internal/app -run '^TestRetrievalIndexSnapshot$' -count=1 -v`。產出 `<index>.manifest.json`，拒絕覆寫既有 index。
2. 各組設定下表變數，再執行 `go test ./internal/app -run '^TestRetrievalBaselineReplay$' -count=1 -v`。舊 baseline 與新 overlap suite 分開跑、分開存 report。
3. 將同一 dataset 的四份 report 絕對路徑存成 JSON array；設 `AI_SKILL_RETRIEVAL_MATRIX_REPORTS` 為該檔，執行 `go test ./internal/app -run '^TestRetrievalAblationMatrix$' -count=1 -v`。這只核對產物配對，不能替代品質驗收。

| 變數 | 必填內容 |
| --- | --- |
| `AI_SKILL_RETRIEVAL_CASES` / `REPORT` | 固定 dataset 與獨立輸出 JSON |
| `AI_SKILL_RETRIEVAL_CORPUS` | 固定 corpus checkout；本次 B0 使用 `5f12fc3906051e40228959e9e6bfaed714e0658e` |
| `AI_SKILL_RETRIEVAL_QUERY_BINARY` / `QUERY_SHA256` | 實際 binary 路徑與預期完整 SHA；缺席即拒跑 |
| `AI_SKILL_RETRIEVAL_INDEX` / `INDEX_MANIFEST` | 明確的預建 index 及其 manifest |
| `AI_SKILL_RETRIEVAL_GROUP` | B0 / Q / E / QE |
| `AI_SKILL_RETRIEVAL_REFERENCE` | 凍結的 [baseline-results.json](baseline-results.json)，用其舊 query / inputs / corpus 身分檢查單組 |
| `AI_SKILL_RETRIEVAL_REVIEW_ALL` | 設 `1`，對診斷正例的候選也產生 relevance review 待辦 |

四組配對 validator 要求同一 dataset / corpus / 原始來源集合，B0=Q 的完整 index identity、E=QE 的完整 index identity、B0=E 的 binary SHA、Q=QE 的 binary SHA，且新舊 query / extraction 身分確實不同。新增正文 / corpus coverage 另組實驗，不得塞進 E。

## Overlap 補充資料與候選審閱

[overlap-controls-v1.json](overlap-controls-v1.json) 是另版人工自然語言 scope 控制，不回改原 baseline：6 題、5 個正向錨點、6 個事先標註不相關 distractor；最後一題沒有完整 gold answer，不能把所有其他新來源自動判錯。這些控制不是獨立真實提問、不是 holdout，也不是自然情境誤推率估計。部分使用原固定比較組的 pool / clock 題材，因此更不能把兩份資料合稱獨立評估。

例：host / middleware 詞彙出現，但實際問題是自足的 JSON 數字比較；DB pool 與 now() 都出現，但時間戳只作描述日誌、失敗在 nested checkout。標註依 lesson 的明確適用 / 不適用條件與本題主要 intent，不按詞彙相似度裁定。

每個回傳 source 標為 `prelabelled-relevant`、`prelabelled-confuser` 或 `unreviewed`，不能把有候選直接算通過。Phase 2 的新增 gate：

- 原 baseline 的診斷 ≥ 8/14、固定比較集不退步、人工負面 0/2、warm CLI p95 ≤ 500ms、preview ≤ 2000 runes 仍保留；runes 不冒充實際 tokens。
- Overlap suite 的已標 confuser 命中必須為 0；尚未標註的候選逐筆判定 `relevant-to-intent` / `background-only` / `irrelevant` / `unknown`，附來源、版本、判定理由與 reviewer。Owner 為 query / extraction 執行者整理，larrylin 或指定 reviewer 覆核。
- 所有診斷新增候選（Q / E / QE 相對 B0）同樣需審閱；`unknown` 或未審閱項未結清，不宣告 Phase 2 通過。空結果不算 precision=100%。
- 已覆核為不相關、且合併版本仍回傳的候選需修正或記錄具體未達標；不得僅憑 8/14 Recall 宣告更有用。這是有限案例的品質 gate，不外推一般 precision / 誤推率。若 reviewer 推翻預標註，新增 v2 並對全部組別重跑；保留 v1，不只修改對某一組有利的標記。

## 本輪證據

Pinned B0 使用舊 query binary、舊 extraction index、固定 corpus，重現 diagnostic 1/14、固定預留正例 0/4、人工負面 0/2。完整 report 見 [prephase2-b0-results.json](prephase2-b0-results.json)，overlap 的舊版輸出見 [overlap-v1-b0-results.json](overlap-v1-b0-results.json)。這只是控制路徑 smoke，**不是** query / extraction 已改善。

測試覆蓋缺 binary、錯 SHA、index 竄改、dirty corpus / generator、Q/E 分組錯置、四組 index / binary 配對錯置、dataset 或 corpus coverage 偷換。實際 Q / E / QE 與新候選人工審閱仍未執行。
