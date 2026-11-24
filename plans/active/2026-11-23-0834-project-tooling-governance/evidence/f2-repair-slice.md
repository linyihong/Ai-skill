# F2 修復切片：命令表改為宣告式，T3 改用狀態比對

- **日期**：2026-11-24
- **基準**：`2e9447da`
- **開放**：使用者 2026-11-24「開 F2 修復切片」。範圍沿用 [`phase3-application.md`](phase3-application.md) §8 的 F2，以及 [`f1-repair-slice.md`](f1-repair-slice.md) R3 的誤報。
- **做法裁決**（使用者 2026-11-24）：採用宣告式命令表。dispatch 讀這份表；Go test 雙向比對命令表與 command-contract；diff pattern 鎖改為 deprecated，replaced_by 指向新檢查。
- **框架修改裁決**（使用者 2026-11-24）：Consumer B dogfood 不直接改契約本文，在 §6 補上 Ai-skill 的對照證據，交給 Phase 4 一次裁決。

## 1. 盤點：為什麼不能「直接從 dispatch 產生」

F2 原本寫的是「由 dispatch 產生 command-contract 命令表」。盤點後發現 dispatch 寫法不一致，沒有可以穩定推導的單一來源：

| 寫法 | 命令群 |
| --- | --- |
| `switch` 字串 case | 頂層 `app.go`、enforcement、plans、image-processing、kge（在旗標解析之後）、router（`proposals` 底下）、goals（在 `buildGoalsResult` 內） |
| `if` 鏈加上 `!=` 白名單 | runtime（`lookup` / `selection-watch` 先 `if` 分流，其餘用 9 個 `!=` 串接） |
| 單一 `!=` | roo、copilot、glossary |
| 跨函式組合 | hooks（`install` / `run` 先判斷，`run <hook>` 再分成 Claude hooks 和 git hooks 兩個 switch） |

所以順序反過來：**先把命令集合變成資料（Parameters），dispatch 和文件都從這份資料檢查**，而不是從程式碼反推資料。

**盤點時發現的實際漂移**

- `goals` 的 8 個子命令和 `runtime capability-invoke` 可以執行，但 command-contract 裡沒有對應的列。舊的 T3 鎖沒有發現，因為它只看 diff 的新增寫法，不看現有狀態。
- `enforcement` 不帶參數時印出的 usage 寫 `<lint|coverage>`，少了 `transition-check`。手寫的 usage 字串本身也是一份會漂移的副本。

## 2. 修正內容

- **命令表**：[`internal/app/commands.go`](../../../../scripts/ai-skill-cli/internal/app/commands.go) 的 `cliCommands` 列出 17 個頂層命令，以及每個命令的子命令（巢狀形式如 `run pre-commit`、`proposals list`）、摘要與 handler。表在 `init()` 填入，避免 handler 讀表時造成初始化循環。
- **load-bearing**：
  - 頂層 dispatch 和 `ai-skill help` 直接讀表。
  - 每個命令群在執行任何邏輯前，都用表判斷子命令是否已宣告。未宣告的子命令在入口就被拒絕，所以只在 switch 加 case、不加進表，這個命令就跑不到。
  - 錯誤訊息和 exit code 維持原樣；goals、router 原本回傳結構化結果，這點也保留。
  - 各群的 usage 字串改由表產生，§1 的 usage 漂移因此消失。
- **狀態比對測試**：[`commands_contract_test.go`](../../../../scripts/ai-skill-cli/internal/app/commands_contract_test.go)
  - `TestCommandContractMatchesCommandTable`：表中每個命令或子命令都有 contract 列；每個 contract 列都對應到表中宣告的命令。
  - `TestContractRowParsing`：檢查列解析的各種形狀（`<file>`、`[--flag]`、`{a,b}` 展開、表格外的敘述不算列），避免解析退化讓主測試空轉通過。
  - `TestCommandTableGatesDispatch`：未宣告的命令或子命令在 11 個入口都回傳 `ExitInvalidUsage`；`help` 列出表中所有命令。
  - `TestStandaloneBinariesInCatalog`：`cmd/` 底下每個獨立 binary 都必須登記在工具樹 capability catalog。這補上 T3 原本 `cmd/` 不在 scope 內的缺口。
- **contract 修正**：`goals` 列改用 `{status,init,...}` 列出 8 個子命令，並新增 `runtime capability-invoke` 列。
- **deprecated**：
  - 原本的 `validateCLIDocSync`（commit-msg、block）改名為 `warnCLIDocSync`，移出 `commitMsgValidatorRegistry`，由 `runCommitMsgHook` 直接輸出 warning，不再擋 commit。
  - registry 的 `cli_doc_sync` class 改由 `TestCommandContractMatchesCommandTable`（pre-push 的 go test、block）執行，`warnCLIDocSync` 標為 deprecated（warn）。
  - core-bootstrap 的 obligation 改為 severity warning，並附 rationale；`cli-modification-policy.yaml` 的 `gate.cli.command_contract_synced` 一併同步。
  - 改名後 F1 的 hook-call lint 立刻要求登記 `warnCLIDocSync`：F1 剛建的鎖在這裡第一次實際攔下未登記的 hook callee。
- **移除條件**：2026-12-24 之後依明確決策移除 `warnCLIDocSync`，前提是 deprecation 期間沒有任何一次它的 warning 抓到狀態測試漏掉的漂移。不會自動移除。

## 3. T3 樣本狀態（沿用 Phase 3 §5 的分類）

| 項目 | Phase 3 | F2 之後 | 證據 |
| --- | --- | --- | --- |
| 新增頂層 `case` 抓不到 | 缺口 | **結構消除** | 頂層 dispatch 只走命令表；表中新增的命令沒有 contract 列就讓 M1 類型的測試失敗 |
| `cmd/` 不在 scope 內 | 缺口 | **通過** | `TestStandaloneBinariesInCatalog` |
| `portable/` 不在 scope 內 | 缺口 | **不適用** | `portable/` 沒有任何 CLI 入口；它的規則透過既有命令（`kge`）執行，命令集合不會因此改變 |
| test fixture 字串誤報（F1 R3） | 缺口 | **結構消除** | 新檢查比對的是命令表與文件的狀態，不讀 diff 文字 |
| 現有狀態漂移（§1：9 個子命令沒有列） | 未偵測 | **通過** | 主測試在修正 contract 前失敗、修正後通過 |

**突變驗證**（每次都還原）：

| # | 突變 | 結果 |
| --- | --- | --- |
| M1 | 只在表中新增 `plans bogus` | 失敗：dispatchable 但沒有 contract 列 |
| M2 | 刪除 contract 的 `hooks run stop` 列 | 失敗：dispatchable 但沒有 contract 列 |
| M3 | 只在 contract 新增 `plans ghost` 列 | 失敗：contract 列沒有對應的子命令 |

`go test ./...` 全部通過；`ai-skill enforcement lint` 為 0 FAIL；`runtime validate` 通過（runtime.db 與 runtime index 已重建）。

## 4. 代價與取捨

- **回饋時點變晚**：原本在 commit-msg 就擋下，現在要到 pre-push 的 go test，或 CI。舊鎖的早期回饋既有漏抓也有誤報，所以這裡用時點換正確性。deprecated advisory 在 deprecation 期間仍會在 commit 時提醒。
- **pre-push 何時跑**：`cli_ci_preflight` 只在 `scripts/ai-skill-cli/`、hook 或 workflow 有變更時才跑 go test。command-contract 位於 `scripts/ai-skill-cli/docs/`，在範圍內。

## 5. 觀察（observed，本切片不處理）

| # | 觀察 |
| --- | --- |
| S1 | core-bootstrap `per_commit_obligations[].severity` 只是宣告：commit-msg dispatcher 把任何非空的 validator 結果都當成 block，不讀 severity。所以這次要降級，必須在程式碼裡把 validator 移出 registry、改走 warning 路徑。`plan_tree_folder_convention` 的 `severity: warning` 能成立，是因為它的 KGE adapter 自己決定回傳值 |
| S2 | `portable/kge` 的 `CLIDocSyncRule` 仍在 `DefaultRules()` 內（KGE Mini Spike 的 Rule B）。這是可攜引擎的預設規則包，外部專案也會用到；要不要在可攜層 deprecated，另行決定 |
| S3 | `runCommitMsgHook` 對同一個 `git diff --cached --name-only` 呼叫了兩次（`stagedForWarn` 和 `staged`）。這是 §6 T-B2「重複讀取」在 Ai-skill 的小型實例，沒有量測 |

## 6. Consumer B dogfood 對照：交給 Phase 4 的候選

依裁決，契約本文不修改。下表把 [`consumer-b-cli-pilot.md`](consumer-b-cli-pilot.md) 的 T-B1–T-B5 和 Ai-skill 自己的證據並列，標出每一條目前有幾個獨立來源，供 Phase 4 裁決。

| # | Consumer B | Ai-skill 對照 | 來源數 | 建議 Phase 4 處置 |
| --- | --- | --- | --- | --- |
| T-B1 single check entry + registry + 結果形狀，可用 consumer 自己的語言實作 | Node 單一入口試點 | `ai-skill hooks run` 是單一入口，`commitMsgValidatorRegistry` 是 check registry，`Check` / `CommandError` 是統一的結果形狀；Go 實作，不依賴共用 runner | 2 | 作為 Q5 維持 deferred 的證據；契約 §Governance check architecture 不需修改 |
| T-B2 鎖的執行成本（重複讀檔） | 每檔平均被讀 7.5 次；共用讀取省約 77% | 只有觀察（S3：同一個 git 呼叫做兩次；KGE workspace context 已經是共用 context 的形狀），**沒有量測** | 1 個量測、1 個觀察 | 維持候選。若要成為第二個來源，需在 Ai-skill 量一次 hook pipeline 的重複 I/O；是否量測由 Phase 4 決定 |
| T-B3 compensating lock（成對檔案）仍存在 | 刻意保留 wrapper 檔名 | 類似的結構：`scripts/git-hooks/*` 和 `.githooks/*` 是兩套入口外殼（Phase 3 O3），已標 deprecated 並寫明移除條件，沒有用 parity 鎖撐住 | 2（處理方式不同） | 屬 consumer 決策；Ai-skill 的做法（deprecated + 移除條件，而不是加 parity 鎖）可以當作對照例 |
| T-B4 失敗訊息要指出修復動作 | 依賴版本落後時，前置檢查以不相關的解析錯誤失敗 | 分兩半：三個 git hook 函式中，9 個**政策鎖**的 block 點都有 `Remediation`；5 個**工具／環境失敗**（staged 清單讀取、runtime compile、runtime.db stage、runtime validate、commit message 讀取）都沒有，和 Consumer B 是同一種情況。commit-msg 的治理檢查透過 `validateCommitMessageGovernance` 回傳，也都有 Remediation | 2 | **有兩個來源**，而且兩邊的缺口都落在「工具本身失敗」，不是「違反政策」。候選：lock-taxonomy 的 result contract 加上「失敗結果必須含修復動作，**包含鎖自身無法執行的情況**」 |
| T-B5 allowlist 靠逐一列舉名稱維護 | 上游新增一項，導致 66 個測試同時失敗 | `internal_helper_allowlist` 已逐一列舉 82 個名稱，F1 又加了 5 個，是同一種形狀；目前沒有因此大量失敗 | 2 | **有兩個來源**。候選：lock-taxonomy 反例加上「allowlist 以列舉名稱維護，不以屬性 / 介面標記」；這是 Review Rule「能否從結構消除」的套用，不是新規則 |
| 新：diff 文字比對型的鎖 | — | T1（副檔名比對漏掉 shebang 腳本）、T3（diff pattern 漏抓頂層 `case`、誤報 fixture 字串）；F2 改為狀態比對後兩個方向都消除 | 1（Ai-skill 內 2 把鎖） | 候選：pattern-based 鎖要宣告比對層級（文字 / AST / 狀態集合），文字比對須列出排除的情境。只有單一來源，Phase 4 先觀察 |

← [Evidence index](README.md)
