# S3 — PreToolUse 增量化評估

- 日期：2026-11-25
- 環境：同 S1（macOS，Apple M1 Pro，Go 1.26.2）；repo-local binary `d9845ad6`
- 對象：本 session 的 transcript，5.1 MB / 1878 行
- 方法：
  - 分段計時：暫存的 Go 測試直接呼叫內部函式，每段跑 5 次取最小值。測試沒有提交。
  - 端到端：以 stdin JSON 執行 `hooks run pre-tool-use` 5 次，取中位數。
- 判定：評估完成。步驟 0 已實作；步驟 1 的 shadow（1a、1b）已上線，正在累積證據（1c），增量結果不影響任何決策。

## 量測結果

### 端到端

| transcript | PreToolUse 中位數 |
| --- | --- |
| 0.76 MB | 0.039 s（S1） |
| 1.5 MB | 0.114 s（S1） |
| 3.1 MB | 0.318 s（S1） |
| 5.1 MB | **0.524 s**（本次） |

每次呼叫的耗時大致與 transcript 大小成正比。一個 session 有 N 次工具呼叫，總成本約略隨 N² 成長。

### 分段（gate 已通過，`.done` 快取命中的路徑）

| 階段 | 耗時 | 每次呼叫執行幾次 |
| --- | --- | --- |
| `readRuntimeRoutingRegistry` | 5.3 ms | 2 |
| `extractTranscriptMessages`（全量 JSON 解析） | **134 ms** | **2** |
| `BuildRuntimeContext` | 5.7 ms | 2 |
| `transcriptHasRequiredBootstrapReads` | 366 ms | 0（本 session 沒有鎖定的 route）；有鎖定 route 時 1 |
| `extractAssistantTexts` | 83 ms | 0（只在 gate 未通過時） |
| `transcriptHasRuntimeReceiptEvidence` | 7 ms | 0（只在 gate 未通過時） |
| Discovery 設定與 proposal 查詢 | < 1 ms | 1 |

- **主要成本是重複的全量解析**。`finishPreToolUse` 先在 `workflowPrimarySourceGate` 解析整份 transcript，`tryDiscoveryAdvisory` 再解析一次，兩次結果相同。
- 只抽出 66 則 user / assistant 文字，卻要 JSON 解碼整個 5.1 MB。依位元組計：attachment 29%、tool_result 21%、tool_use 12%、api-request 8%、thinking 6%，這些對 detector 都沒有文字貢獻。
- 有鎖定 route 時，`transcriptHasRequiredBootstrapReads` 還要再全量掃一次（366 ms），單次呼叫可能接近 0.9 s。

## 修補選項

### 步驟 0：每次呼叫只解析一次（不是快取）

- `finishPreToolUse` 解析一次 transcript 與 registry，傳給 gate 與 Discovery 共用。
- 預估：0.52 s → 約 0.38 s（省掉一次 134 ms 與一次 registry 讀取）。
- 風險：屬於純重構，同一次呼叫內輸入完全相同，不跨呼叫重用，所以不涉及契約 V3；以既有測試驗證即可，不需要 shadow。
- 限制：仍然是 O(transcript)，只把斜率減半。

### 步驟 1：跨呼叫增量解析（是快取，受 V2 / V3 / V5 約束）

- 快取檔：`os.TempDir()` 下以 transcript 路徑雜湊命名，權限 `0600`。內容包括：版本、transcript 路徑、已處理的位移量、位移量之前最後 4 KB 的雜湊，以及累積狀態（detector messages、已讀檔案、receipt 證據）。
- 每次呼叫：
  1. 檢查檔案大小 ≥ 位移量、邊界雜湊相符、版本相符。任一不符就全量重建（V2：不確定就執行）。
  2. 從位移量往後只讀**完整的行**；最後一行若還沒寫完就不推進位移量。
  3. 以 temp 檔加 rename 原子寫回；讀寫失敗一律退回全量。平行工具呼叫會同時觸發 hook，不能假設只有一個寫入者。
- 單次 pass 同時產出 messages、已讀檔案證據與 receipt 證據，取代目前三個各自全量掃描的函式。
- 預估：每次呼叫只處理新增的位元組（一次工具呼叫約數 KB → < 1 ms）加上讀寫快取檔。整體成本不再隨 transcript 長度成長。
- 隱私：快取檔含 transcript 文字的副本，所以放使用者自己的 TempDir、權限 `0600`，不寫進 repo。

### Shadow 方式與採用門檻（先登記，V5）

- Shadow 期間兩條路都跑，**決策仍用全量結果**。比較增量與全量的 messages、`ActiveRoute`、已讀檔案集合、receipt 證據；不一致時寫進 hook log 並累計計數。
- 採用門檻：
  - ≥ 200 次比對，來自 ≥ 5 個 session；
  - 其中至少 1 個經過 compaction / resume，至少 1 個有平行工具呼叫；
  - **0 次不一致**。
- 任何一次不一致：記錄差異，停止採用評估，先修正再重新計數。
- 撤回條件：若 Claude Code 會改寫既有行（不是只附加），邊界雜湊檢查會頻繁觸發全量重建。這時要先量化重建比例；比例過高就放棄步驟 1，只保留步驟 0。

## 步驟 0 結果（2026-11-25）

- 實作：新增 `loadPreToolUseRouting`，`finishPreToolUse` 每次呼叫只讀一次 registry、只解析一次 transcript、只建一次 RuntimeContext，交給 `workflowPrimarySourceGateWith` 與 `tryDiscoveryAdvisoryWith` 共用。`workflowPrimarySourceGate` 保留原簽名（既有測試與其他計畫引用它）。
- 語意差異：原本兩次 `BuildRuntimeContext` 各取一次 `time.Now()`，相差不到 1 秒；現在共用同一個時間點。
- 驗證：`go test ./internal/app` 通過；新舊 binary 對同一份 transcript 的 stdout 完全相同。
- 量測（真實 transcript 5.3 MB，走 `.done` 快取命中路徑；新舊 binary 交錯各跑 15 次）：

| binary | 中位數 | 最小 | 最大 |
| --- | --- | --- | --- |
| 舊（`d9845ad6`） | 0.62 s | 0.41 s | 1.09 s |
| 新（步驟 0） | **0.40 s** | 0.28 s | 1.30 s |

- 中位數約減少 35%，符合「省掉一次全量解析」的預估。成本仍與 transcript 大小成正比；要讓它不再隨 session 變長而增加，需要步驟 1。

## 步驟 1 的實際收益估算（2026-11-25）

外部審查提醒：步驟 0 成功，不代表步驟 1 值得做；要先量化收益，再和狀態管理成本比較。

- 方法：讀取本專案目前留存的 8 份 Claude transcript，記下每次工具呼叫發生時 transcript 的大小。用步驟 0 後的成本模型估算（每次約 0.03 s 固定成本，加上每 MB 約 0.07 s），步驟 1 則假設只剩固定成本。估算腳本在暫存區，沒有提交。
- 限制：只涵蓋本專案、只有 8 個 session；成本模型由單一 transcript 的量測外推。

| 指標 | 中位數 | p90 / 最大 |
| --- | --- | --- |
| transcript 大小 | 5.1 MB | 15.0 MB |
| 每個 session 的工具呼叫數 | 186 | 586 |
| 每個 session 的 PreToolUse 總耗時（步驟 0 後） | 44 s | 341 s |
| 步驟 1 可省下的時間 | 39 s | 323 s |
| 最慢的單次呼叫 | 0.38 s | 1.08 s |

- 收益集中在長 session：一半的 session 超過 5 MB，最大的那個每次工具呼叫要等超過 1 s。
- Compaction 不會縮小 transcript 檔（本 session 經過 compaction，檔案仍持續變大），所以長 session 的成本不會自然歸零。

### 不需狀態的替代方案：已否決

- 試過先只解碼 `type` / `role` 兩個欄位，不是 user / assistant 的行就跳過。
- 結果與現行完全相同，但在兩份最大的 transcript 上反而**慢 15–30%**（186 → 216 ms、76 → 98 ms）。原因是 user / assistant 行仍要完整解碼一次，而 Go 的 JSON 解碼不論目標型別都要掃過整行。
- 結論：不需狀態的方法沒有明顯收益，步驟 1 的狀態管理成本有其必要。

## 步驟 1 shadow 實作（1a、1b，2026-11-25）

- 程式：`scripts/ai-skill-cli/internal/app/transcript_incremental.go`。`finishPreToolUse` 在完整解析之後呼叫 `runIncrementalShadow`，回傳值直接丟棄。
- 狀態檔：`$TMPDIR/ai-skill-transcript-incremental-<路徑雜湊>.json`，權限 `0600`，以 temp 檔加 rename 原子寫入。
- Shadow 紀錄：`$TMPDIR/ai-skill-pretooluse-incremental-shadow.jsonl`。每行記錄 session 雜湊、大小、結果、模式、重建原因、新增位元組、兩邊的訊息數與耗時，不含 transcript 文字。
  - 結果分四種：`equal`、`mismatch`、`not_counted`、`error`。
  - macOS 的 `$TMPDIR` 重開機可能被清空，所以收集證據時要先把紀錄複製出來。
- 關閉方式：環境變數 `AI_SKILL_PRETOOLUSE_INCREMENTAL_SHADOW=0`，這時紀錄為 `not_counted`。
- 重建原因已實作：`no_state`、`state_decode_failed`、`state_version_mismatch`、`transcript_path_mismatch`、`cursor_invalid`、`transcript_truncated`、`head_hash_mismatch`、`boundary_hash_mismatch`、`cursor_unverifiable`。
- 為了能和完整解析比較，增量解析刻意複製了它的兩個現有行為：
  - 最後一行還沒寫完時，照樣解析進本次結果，但游標不推進；
  - 遇到超過 2 MB 的行就停止，之後的行全部不處理（`stalled`），和 `bufio.Scanner` 一樣。
- **已知限制**：雜湊只檢查檔案開頭與游標前各 4 KB。中間區段如果被修改，增量解析不會發現，只有 shadow 比對會抓到。這是 shadow 不能省略的原因之一，切換前要用 shadow 證據說明這種情況是否發生。
- **旁支觀察（observed）**：完整解析遇到超過 2 MB 的行會靜默停止，之後的訊息全部被忽略。這是現行完整解析本身的問題，不在步驟 1 範圍；shadow 紀錄的 `stalled` 欄位會顯示它是否真的發生。

### 上線前驗證（不計入 1c 的 200 次）

- 單元測試涵蓋：
  - 附加後增量續讀；
  - 每一種重建原因，以及重建之後下一次能恢復增量；
  - 最後一行未寫完、無法解碼的行、空行；
  - 超長行停止；
  - 四種紀錄結果，以及紀錄中不含 transcript 文字；
  - 狀態檔權限 `0600`；
  - 被竄改的增量狀態只會表現為 `mismatch`，不會進入正式解析。
- 真實 transcript（5.65 MB）：第一次重建 210 ms；之後增量 2.5–3.4 ms。三次都與完整解析一致；狀態檔 75 KB。
- 離線重播：把 8 份 transcript 從頭逐步寫回，每 7 行比較一次，共 2159 次比較、0 次不一致，8 次重建（都是各檔第一次）。這是離線模擬，不是實際 hook 執行，**不計入** 1c 的證據。

## 建議

1. ~~先做步驟 0~~：已完成，見上。
2. 步驟 1 以 shadow 方式實作，達到門檻再提出切換。執行邊界見 `_plan.md` §Phase 2a 步驟 1。它屬於 Phase 2a 的 dogfood 主體，會首次在 Ai-skill 自身套用 V3 與 V5。
3. Stop / user-prompt-submit hook 也會掃 transcript，但它們每回合只跑一次（S1 顯示 0.02–0.14 s），不在本範圍。

## 對契約的回饋

- 成本來源不只是「沒有快取」，也包括**同一次呼叫內的重複計算**。先移除重複，再談跨呼叫重用，順序比較安全。
- 增量快取的失效條件（檔案縮短、邊界雜湊不符、版本不符）就是 V3 要求的「重用理由」；任一條件無法確認時退回全量，就是 V2。
