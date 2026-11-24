# S2 唯讀盤點：外部專案怎麼使用 `portable/kge`

- **日期**：2026-11-24
- **基準**：`c037e906`
- **開放**：使用者 2026-11-24「開 S2 唯讀盤點」。範圍：[`f2-repair-slice.md`](f2-repair-slice.md) §7 的 S2（`CLIDocSyncRule` 要不要在可攜層 deprecated，先查外部使用方式和替代檢查），並包含 [`f6-repair-slice.md`](f6-repair-slice.md) §4 的 V1（`NoNewShellScriptsRule` 行為改變對外部 adapter 的影響）。
- **方式**：全程唯讀。沒有修改任何外部 repo，也沒有寫入它們的 `.ai-skill/dogfood/`。

## 1. 外部專案接觸 `DefaultRules()` 的路徑

先從 Ai-skill 原始碼確認，外部專案可能透過哪些路徑跑到可攜規則：

| # | 路徑 | 會跑到的規則 | 依據 |
| --- | --- | --- | --- |
| P1 | **複製單元**：把 `portable/kge/` 複製到自己的 Go module，自己寫 adapter | adapter 自己選的規則。README 的複製步驟範例寫的就是 `kge.NewEngine(kge.CognitiveCostRule{}, kge.CLIDocSyncRule{})` | `portable/kge/README.md` §複製步驟 |
| P2 | **Go import**：直接 import `github.com/linyihong/Ai-skill/scripts/ai-skill-cli/portable/kge` | 同 P1 | `go.mod` module path |
| P3 | **執行 Ai-skill binary**：在外部 repo 跑 `ai-skill kge check/validate/diagnose`，或讓 git hook 呼叫 `ai-skill hooks run pre-commit/commit-msg/pre-push` | 完整 `DefaultRules()`（含 `CLIDocSyncRule`、`NoNewShellScriptsRule`） | `kge_cmd.go` 預設 `--root` 是目前目錄；呼叫 KGE 的只有 `runKGENoNewShellScripts`（pre-commit）、`runKGECLIDocSync`（commit-msg）、`attachKGEAdvisoryCount`（commit-msg）、`attachKGECheck`（pre-push） |

Claude Code 的 hook（`session-start`、`pre-tool-use` 等）**不會**呼叫任何 KGE 規則，所以只裝了 Ai-skill Claude hook 的專案不屬於 P3。

## 2. 盤點結果

範圍：本機專案根目錄往下 4 層內的 88 個 git repo，其中包含 Consumer A、Consumer B。

| 檢查 | 方法 | 結果 |
| --- | --- | --- |
| P1 複製單元 | 搜尋原始碼裡的 `package kge`、`portable/kge`、`CLIDocSyncRule`、`kge.DefaultRules` | **0 個 repo** |
| P2 Go import | 搜尋 `go.mod` 裡的 `linyihong/Ai-skill` | **0 個 repo** |
| P3 git hook 呼叫 binary | 19 個 repo 有 git hook（設了 `core.hooksPath` 或有非 sample 的 `.git/hooks/*`）；在 hook 檔內容搜尋 `ai-skill` | **0 個 repo 呼叫 binary**。有 9 個 repo 的 hook 檔提到 `.ai-skill/`，但只是路徑字串，用在 sanitization 的 forbidden pattern 或說明文件連結，不是呼叫 |
| P3 文件或設定指示執行 `ai-skill kge` | 搜尋原始碼與 `.claude/settings*.json` 裡的 `ai-skill kge`、`ai-skill hooks run` | **0 個 repo** |
| Consumer A | 同上 | P1–P3 都沒有。git hook 是自己的 pre-commit，不呼叫 Ai-skill |
| Consumer B | 同上 | P1–P3 都沒有。`.githooks` 是自己的 Node pipeline（見 [`consumer-b-cli-pilot.md`](consumer-b-cli-pilot.md)），不呼叫 Ai-skill |

**替代檢查**：沒有任何外部 repo 用到 `CLIDocSyncRule`，所以「外部專案拿掉它之後有沒有替代檢查」這個問題目前沒有對象，判為**不適用**。

## 3. 對 S2 和 V1 的意義（觀察，不是裁決）

- **S2**：在本機可見範圍內，`CLIDocSyncRule` 目前唯一的使用者是 Ai-skill 自己，透過 `DefaultRules()` 和 `warnCLIDocSync`。F2 審閱擔心的是「外部專案可能把它當成唯一的同步檢查」，這個情況在盤點範圍內**沒有發生**。
- **V1**：同理，F6 改了 `NoNewShellScriptsRule` 的行為，在盤點範圍內**沒有外部 adapter 受影響**。
- **仍要決定的事**：`CLIDocSyncRule` 要不要 deprecated、什麼時候 deprecated，仍然由使用者裁決，本盤點不預設結論。如果要 deprecated，至少要一起處理兩個地方：
  - `portable/kge/README.md` 的複製範例本身就用 `CLIDocSyncRule`，會把新的外部使用者引向這條規則；
  - `DefaultRules()` 的內容。

## 4. 盤點的限制

| # | 限制 |
| --- | --- |
| K1 | 只涵蓋這台機器上 clone 的 repo。其他機器、只存在遠端的 repo、或其他人的 fork 都看不到。沒有查 GitHub 上的 public dependents（需要對外查詢） |
| K2 | P1 只靠字串搜尋。如果有人複製後改了 package 名稱和型別名稱，就搜不到 |
| K3 | 狀態是 **observed**（本機掃描），不是 verified（全部 consumer） |

## 5. 範圍外的觀察（observed，不處理）

| # | 觀察 |
| --- | --- |
| W1 | 有 8 個外部 repo 的 hooks 目錄裡有 `commit-sanitization-check.py` 加 `forbidden-patterns.json`，看起來是同一份 sanitization hook 被複製到多處。這不屬於 `portable/kge`，和 S2 無關；是 consumer 自己的工具。只記錄，沒有檢查來源和版本差異 |

← [Evidence index](README.md)
