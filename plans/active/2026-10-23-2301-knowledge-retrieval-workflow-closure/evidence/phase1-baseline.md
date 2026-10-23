# Phase 1 — Workflow 知識查找 Baseline

日期：2026-10-24。Corpus revision：`5f12fc3906051e40228959e9e6bfaed714e0658e`。本輪只建立量測與證據，production retrieval、workflow、Bridge 均未變更。**基準已建立不等於 Phase 1 gate 已完整通過**；指標重算、未證實原因與資料曝光紀錄見 [phase1-audit.md](phase1-audit.md)。

## 產物與資料範圍

- [phase1-protocol.md](phase1-protocol.md)：資料來源、評分、成本、holdout 與重跑契約。
- [baseline-cases.json](baseline-cases.json)：固定 20 題；18 題歷史 lesson 重播、2 題人工負面控制，14 / 6 切分。
- [baseline-results.json](baseline-results.json)：由 replay harness 輸出的 baseline；含 dataset / corpus inputs / binary SHA、逐題輸出、延遲及分母。
- Harness：`scripts/ai-skill-cli/internal/app/retrieval_baseline_test.go`；test-only，production CLI 不增加命令，不需發佈新 binary。

這不是 20 個獨立原始使用者查詢；目前 expected sources 都是 feedback lesson，代表第一批「歷史經驗查找」重播，不代表已量測全庫 intelligence 的 recall。intelligence 來源覆蓋缺口沿用 Phase 0 的獨立 inventory，不混進本表分母。後續將 intelligence 來源加入比較集時須另定 dataset 版本，不能回改本 baseline。

## 結果

| 組別 | 正例命中 / 正例數 | Recall@5 | MRR@5 | 負面誤推 / 負面題數 |
| --- | --- | --- | --- | --- |
| Diagnostic | 1 / 14 | 0.071429 | 0.071429 | not evaluated（本組沒有負面題） |
| Holdout | 0 / 4 | 0 | 0 | 0 / 2（僅人工控制） |

固定先取前 5 atoms 再按來源去重，不補位。每正例一個 expected source；小樣本不宣稱統計顯著，不將全部未命中解讀成沒有相關知識。

| 成本 / 可重現項 | 結果 |
| --- | --- |
| 環境 | darwin/arm64，Go 版本與 binary SHA 見 JSON；binary source `b7d8b88aabfc844461452bc8f0e3169c62bd323b` |
| Warm CLI samples | 20 題 × 20 次 = 400 次；另各一筆 first invocation |
| Warm CLI p95 | 11.953 ms（本機，非服務等級保證） |
| 最大預覽估算 | 545 UTF-8 runes；不是模型 token 數，也不是已接入的 cap |
| 第二次整組 replay | corpus / dataset / binary fingerprint、候選、排序與分數一致；p95 10.391 ms；未因重跑改 query / label / acceptance |
| Cold start | not evaluated；有記每題首筆，但未控制 OS cache |
| 適用性分布 / 採用錯誤 / source-read / blocker 保持 | not evaluated；採用鏈尚未接入 |
| Route discovery | not evaluated；Bridge 保留現況，不能用本結果推論 route 推薦品質 |

## 診斷集漏找原因

只對 diagnostic 做方法診斷，不展開 holdout 個案。13 個 missed expected source 都存在於重建索引的 sources，因此本版不是 `source_not_indexed`，也沒有證據可直接分類為語意落差。

目前可證明的分類為 `query_processing_or_indexed_fields`：expected source 在索引，但整句 quoted FTS query 對該 source 沒有命中，尚不能分開計算查詢處理與欄位遺漏各占多少。

- `runtimeFTSMatchLiteral` 把整句當 phrase，不是獨立術語 OR / AND，更沒有 query expansion。
- `runtimeIndexFeedbackRecords` 的 title 先找 `###`，summary 尋找 `#### One-line Summary`。舊 lesson 若使用 `#` 標題與 `## One-line Summary`，會退回檔名且漏掉摘要。R04、R11 是可讀回的格式差異例子，不能把這種缺口歸因為 embeddings 不夠好。
- 人工 FTS control 驗證：`host pipeline` 能命中原序片語、`pipeline host` 不能；預設 FTS 下 `模型` 不能命中整個 token `模型載入失敗`。這只是 tokenizer / phrase 行為探測，不是假扮歷史真實案例。

## 邊界案例覆蓋

`TestRetrievalReplayStaleAndHistoricalControls` 使用臨時生成資料：保留 deprecated source 供歷史查詢、以明確 validated filter 排除不適用狀態、來源改動後 freshness check 拒絕舊 checksum。這些為人工 fixture controls，**不是**真實 supersession、跨版本採用或適用性裁決已成功的證據；相關採用情境仍由 Phase 3 驗證。

資料包含歷史追溯、candidate 來源與中文 / 英文混合問題，但尚無自然語意無答案、實際已取代決策的獨立歷史重播或全庫 relevance 標註。保留這些限制，不把人工控制的 0/2 當成通用無答案表現。

## 下一階段決策

Phase 2 先對 diagnostic 的 query / indexed-field 候選原因做分離與合併對照，再決定修正與較小的 source-level 正文覆蓋；尚未有逐題原因歸屬的 ablation 證據。不因這次 1/14 就導入向量。先前凍結門檻：diagnostic ≥ 8/14、holdout Recall / MRR 不退步、人工負面誤推 0/2、warm p95 ≤ 500ms、preview estimate ≤ 2000。沒有新 reranker、vector provider 或 Bridge 接線承諾。

初始診斷未用預留個案結果設計修法，但資料標記與本次 artifact 重算已曝光題目 / 原始結果；保守視為固定 comparison / regression set，不主張嚴格獨立 holdout。需要獨立泛化驗證時另收資料，Phase 4 也重新核對。

## Validation

- [x] Provenance SHA 與來源存在檢查、group 不跨 split。
- [x] 正式 generator 重建 temporary index、開始與結束的來源 freshness 檢查。
- [x] 兩次 live replay 的候選 / 排序 / 分數與 fingerprints 一致。
- [x] 來源去重 / MRR、nearest-rank p95、FTS phrase / CJK 與 stale / historical controls 通過。
- [x] 本機完整 Go tests 通過（首次 full suite + 新增 controls 後再次確認）。

本輪無 metadata / workflow / routing / compiler source 變更，不重建全庫 generated reports；基準索引以既有 generator 專用 temporary rebuild 產生。
