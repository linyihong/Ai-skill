# Phase 2 v2 — 精確優先、欄位排序與有界輸出協議

本文件於 v2 replay 前固定方法。v1 候選已拒收，不再要求為結清而審完 191 rows；原 queue 保持未完成。最後準備採用的版本仍須完成全部新增候選覆核。

## 代表案例診斷

[逐欄匹配資料](phase2-v2-representative-matches.json) 使用 v1 NEW index，對 R01 / R03 / R04 / R11 / O01 / O02 / O04 / O06 的 v1 QE 前五 source，逐詞以 SQLite FTS5 欄位限制查詢。它是欄位歸因的索引檢查，並非重建實際 binary 的排名；v1 逐題 report 才是實際排名證據。詞項取 Unicode word runs，沒有來源黑名單。

- R04 全片語可命中 title；R11 原查詢可命中 title。先保護完整 phrase，不讓 substring 等單詞補充結果擠入。
- O01 host-pipeline confuser 命中 host / middleware；正向 lesson 命中 JSON / code / 1002 / 1。相關詞與背景詞會並存。
- O02 confuser 的 API / SQL 在 title / summary 命中，正向來源目前只命中 summary 的 viewport；僅提高 title 權重可能仍失敗，必須實測。
- R03 正向來源只有 path / ID / validation_signal 命中 baseline。完全刪掉這些欄位會損失精確術語入口，不直接移除。
- O06 含「已排除 ECONNREFUSED」，FTS 仍可能命中該字；v2 不假裝具備否定、原因或適用性理解。

## 固定 v2 候選

1. 使用完整 escaped phrase 查詢；有結果時只組裝 phrase 結果，不以鬆散 OR 補滿空位。
2. 完整 phrase 無命中才使用 v1 Unicode 詞項 OR（最多 64 個，無同義詞、停用詞表、CJK 分詞、否定解析或 query-specific source 排除）。
3. FTS5 bm25 欄位權重：id / source_path 各 0.25、title 8、summary 5、tags / when_to_read / validation_signal 各 0。欄位仍能召回，但 boilerplate 不增加詞頻權重。延續既有 priority / confidence / context_cost / id 同分排序；這些既有 metadata 不表示適用性或權威判定。
4. 每段最多取得 64 atoms，依排名按 canonical source 去重。v2 保守輸出最多 1 個來源（並遵守更小的 --limit）；這是固定 shortlist 實驗，不宣稱 top1 等於相關，也不依 confuser 或 label 動態挑選。
5. consumer payload 定義延續原工程代理：實際回傳 results array 的 compact JSON，包含全部 QueryResult metadata / source / reason / summary，不計 CLI diagnostics envelope。上限 2000 UTF-8 runes，仍非模型 tokens。
6. 固定 reason 簡述 match stage，不重複整段 query。先保留完整來源身分，只縮短 summary，以 Unicode rune 邊界加省略號；身分欄位仍無法容納則不輸出該來源。不得剪斷 JSON 或回報未實際輸出的來源命中。

## 評估與 gate

沿用不變的 baseline-cases / overlap-controls-v1、兩份固定 index manifests、B0 / E 舊 binary，Q / QE 換成 v2 pinned binary；重跑八份報告與兩份 matrix validator。Generator / corpus / binary 身分規則不變。查詢實作與 extraction 源碼仍在 experiment 分支；main 只收證據直到品質與審閱通過。

v1 與 v2 都用各自實際 binary 的最終 results 計算 hits / Recall / MRR / confuser / preview；不以預算前候選代替最終品質。另保留 v2 排序後、去重 / top1 / summary budget 前的診斷資料，與 final binary 輸出比對，揭露哪些來源被 shortlist 或預算移除。診斷呼叫同 revision 的 production candidate function，標明不是額外獨立的 CLI 輸出。成本限制造成的召回損失必須呈現。

原品質 gate 不變：diagnostic ≥ 8/14，固定比較不退步，人工負面 0/2，overlap confuser 0，最終 preview ≤ 2000 runes，p95 ≤ 500ms；準備採用版本完整審閱。若任一明確失敗，保留失敗證據與代表案例即可，不為拒收版本結清全部 queue；不能外推一般 precision 或泛化。這些資料已曝光，維持固定比較 / 診斷資料身分。未新增向量或 reranking 基礎設施。
