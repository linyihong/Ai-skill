<!-- archival-provenance -->
<!-- Fixed-version historical URLs and corpus paths below retain their original active location. -->

# Phase 1 基準協議

## 資料與適用範圍

[baseline-cases.json](baseline-cases.json) 共 20 題：18 題由去敏歷史 lesson 的 Trigger / Evidence 改寫的重播問題、2 題人工負面控制。不是原始使用者提問逐字紀錄，也不能宣稱 20 個獨立真實任務。

每題保留 id、group、split、query intent、task context、來源與 SHA-256、expected source paths、相關理由、適用性待驗狀態。來源可能仍為 candidate；expected 只代表相關，不代表可信或可採用。18 條來源已核對 Trigger / Evidence / 不適用條件。

14 題 diagnostic，6 題 holdout（4 題歷史重播、2 題人工負面）。同一問題家族不得跨組；兩種 time-source 問題同留 holdout，相關的 baseline / restore 問題都在 diagnostic。診斷案例可能來自同一較大 incident，不能視為獨立事件數。

有 3 題精確名稱、1 題歷史理由追溯，其餘為當前情境查找；包含中文與中英混合。歷史較舊或 candidate 不作自動排除條件。不存在 status 的控制只驗 metadata 空集合；nonce 控制只驗不存在識別字，不代表自然語言無答案能力。自然語意無答案、真實 superseded / stale、不適用採用案例留待 Phase 3 scenario，不能用這兩個控制冒充。

## Preflight 與生成路徑

Trigger：接續 Phase 1。Checked sources：Phase 0 列出的 generator / query / metadata / runtime ownership / sd-validation，另讀 CLI fixture / BDD docs 與 lesson 來源。Decision：只新增 test-only harness，不改 production search、Bridge、workflow 或 metadata。

Runner 用現有 `buildNativeRuntimeSQLiteIndex` 建 temporary index，再用 committed host binary 的 `runtime query` 查詢；不複用可能 stale 的本機 cache，不重寫全庫 reports。這是正式 generator 的針對性重建，非完整 runtime refresh。驗證來源 checksum，並固定 corpus revision、dataset / binary SHA、generator records / edges / sources 的 logical SHA、OS / arch、Go version。SQLite 檔案時間與 temporary path 不作 corpus identity。

## 量尺與先行凍結的驗收

- 取既有 top 5 atoms，再按 canonical source 去重；不從第 6 名補位。重複 atoms 占掉名額是既有查詢行為，不能靠評分器隱藏。
- Recall@5 = hit expected sources / expected sources；本版每正例一個 expected source。MRR@5 = 正例首個相關 canonical source reciprocal rank 平均；未命中為 0。負面控制不混入正例分母。
- 無答案誤推 = 有任何候選的負面題數 / 負面題數；不得外推成一般未知問題錯誤率。
- 每題第一個 CLI invocation 另記，之後 20 次重跑供 warm latency；包含 process / CLI / DB query，不包含人工回讀。每次檢查候選一致。OS cache 未清理，第一筆不是 cold-start proof。
- p95 為 nearest-rank；效能容許同機波動，候選內容與排序須一致。
- Preview token 為 result JSON 的 UTF-8 rune count 工程估算，不是模型 tokenizer，不宣稱已機械執行 token cap。
- Phase 2 固定 acceptance（已在 baseline-cases.json 寫入才開始量測）：diagnostic 至少 8/14 題命中；holdout Recall@5 / MRR@5 不低於本 baseline；負面誤推 0/2；warm p95 ≤ 500ms、preview estimate ≤ 2000。不可看到結果後改分母或 query。
- 若 preview 估算不通過，需另訂有版本的實際 tokenizer 協議，不可靜默更換量尺。
- 採用錯誤、source-read、適用性分布、required-read / blocker 保持均為 `not evaluated`；route discovery 也 `not evaluated`。本次只量 workflow knowledge lookup，不宣稱 Bridge 改善。

## 重跑方式

下列是原始基準的資料 / 報告參數；新版 harness 另強制要求顯式 query binary 路徑與預期 SHA、clean corpus。B0 / Q / E / QE 必須再提供預建 index / manifest / reference，完整環境與 snapshot 步驟見 [phase2-execution-contract.md](phase2-execution-contract.md)，不能只設定兩個變數就執行新版實驗。

從 repo root 設定兩個環境變數為絕對路徑（需 export 或單次命令環境傳入），再於 `scripts/ai-skill-cli` 執行：

```text
AI_SKILL_RETRIEVAL_CASES=<repo>/plans/active/2026-10-23-2301-knowledge-retrieval-workflow-closure/evidence/baseline-cases.json
AI_SKILL_RETRIEVAL_REPORT=<absolute-output-path>.json
go test ./internal/app -run '^TestRetrievalBaselineReplay$' -count=1 -v
```

`AI_SKILL_RETRIEVAL_CORPUS` 可指向固定 corpus revision 的獨立 checkout；runner 本身從目前含 harness 的 checkout 執行，避免舊 revision 尚無測試檔而無法重播。未指定時用目前 repo。精確重現 baseline 時應使用 baseline-results.json 的 corpus_revision 建獨立 checkout，並使用記錄的 binary / dataset SHA；不同 revision 的結果要另存並比較，不能覆寫 baseline。

變數缺席時 runner skip，CI 不會把 live-corpus 品質當單元測試通過證據。獨立 metrics tests 正常執行，覆蓋來源去重、MRR 與 p95。

## Holdout 使用紀律

初始 baseline 存 aggregate 與原始輸出以可重現；後續方法設計只參考 diagnostic，避免讀 holdout 個案錯誤作診斷。資料標記者已看過題目，故非密封盲測，只稱「不參與調參的預留測試組」。Phase 2 評估一次後若用於診斷 / 方法選擇 / prompt / 調參，須標已消耗，Phase 4 另收留出資料。Baseline 執行前已固定 query、split、expected sources 與 acceptance，不能因結果不佳重寫。
