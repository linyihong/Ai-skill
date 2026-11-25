# S1 — Ai-skill 自身 hooks 耗時基線

- 日期：2026-11-25
- 環境：macOS，Apple M1 Pro（10 核），Go 1.26.2，repo-local binary `ai-skill-darwin-arm64`（`7516c7fe`）
- 方法：暫存區的計時腳本，每項重複 3–5 次取中位數；stdout / stderr 丟棄，只計 wall time。計時腳本不進 repo。

## 量測結果

### Git hooks（每次 commit / push）

| 項目 | 次數 | 中位數 | 範圍 |
| --- | --- | --- | --- |
| `hooks run pre-commit`（無 staged） | 5 | 0.087 s | 0.080–0.093 |
| `hooks run commit-msg`（上一個 commit 的訊息） | 5 | 0.149 s | 0.141–0.186 |
| `hooks run pre-push`（無待推 commit；含 kge check） | 5 | 0.172 s | 0.167–0.212 |

結論：不值得做快取。

### `go test ./...`（pre-push 在 CLI 路徑有變更時執行）

| 項目 | 次數 | 中位數 | 範圍 |
| --- | --- | --- | --- |
| `-count=1`（不用 Go 測試快取） | 3 | 25.4 s | 23.3–35.1 |
| 預設（Go 測試快取命中） | 3 | 2.0 s | 2.0–2.2 |

- 各 package：`internal/app` 27.4 s，其餘每個 ≤ 1.4 s。
- `internal/app`：590 個 top-level 測試，各自耗時加總 90 s（並行後約 27 s）。最慢三個都是 runtime compiler 測試，合計約 21 s。
- **Go 內建測試快取已經是一種結果重用**：以 build 輸入，加上測試實際開啟的檔案與讀取的環境變數作為 key。pre-push 執行的是 `go test ./...`（沒有 `-count=1`），所以會用到這個快取。真正會跑滿的情況是 CLI 原始碼有改動的時候，這時 `internal/app` 本來就必須重跑。
- **結論：不要自己做 `go test` 結果快取**，那會和 Go 的快取重疊。要降低耗時，該處理的是 `internal/app` 的切分，或最慢的 runtime compiler 測試。

### Claude Code hooks（每次工具呼叫 / 每次提問）

| 項目 | transcript | 次數 | 中位數 | 範圍 |
| --- | --- | --- | --- | --- |
| `pre-tool-use` | 3.1 MB / 1163 行（本 session） | 5 | 0.318 s | 0.255–1.200 |
| `pre-tool-use` | 1.5 MB / 前 300 行 | 5 | 0.114 s | 0.074–0.241 |
| `pre-tool-use` | 0.76 MB / 前 60 行 | 5 | 0.039 s | 0.027–0.068 |
| `post-tool-use` | 3.1 MB | 5 | 0.027 s | 0.014–0.050 |
| `user-prompt-submit` | 3.1 MB | 5 | 0.143 s | 0.111–0.320 |
| `stop` | 3.1 MB | 5 | 0.019 s | 0.011–0.030 |

- PreToolUse 每次工具呼叫都會執行，耗時**隨 transcript 變大而增加**。gate 已通過時（`.done` 快取命中）仍需約 0.32 s，所以主要成本不在 receipt 檢查，而在後續每次從 transcript 重建狀態（`RuntimeContext is rebuilt from the transcript each call`）。整個 session 的總成本約略隨長度平方成長。
- 比較限制：截短的 transcript 沒有 `.done` 快取，走的分支不同，只能看出趨勢，不能精確比較。

## V2 觀察：測試選擇的路徑範圍（observed，未 verified）

- pre-push 的 `hasCLICIPreflightChange` 只在 `scripts/ai-skill-cli/`、`scripts/git-hooks/`、`.githooks/`、CLI workflow 有變更時才跑 `go test`。GitHub CI（`.github/workflows/ai-skill-cli.yml`）同樣以 `scripts/**` 與 workflow 檔過濾。
- `internal/app` 測試有 37 處呼叫讀取真實 repo 的 helper。`filepath.Join` 常見的目標包含 `runtime/`、`knowledge/`（如 `routing-registry.yaml`）、`plans/`、`runtime.db`。
- 因此，**只改 `runtime/`、`knowledge/` 或 `plans/` 的 push，本機與 CI 都不會跑任何 Go 測試**，即使有測試讀取這些路徑。依契約 V2，依賴範圍無法證明時應該執行。
- 尚未驗證的部分：是否真的存在「只改這些路徑就會讓某個測試失敗卻沒被發現」的情況。其他 gate（kge check、runtime index freshness）可能已涵蓋部分路徑。可證偽的驗證方式：在暫存 worktree 只改一個被測試讀取的 runtime YAML，讓測試失敗，再確認 pre-push 是否放行。

## 對契約的回饋

- V3：Go 測試快取是一個現成、以實際讀取的檔案為 key 的重用機制，可作為「重用要能說出理由」的參考實作。
- V1 / V2：這次的缺口不是快取造成的，而是**路徑式選擇**。這說明 V2 不只適用於快取，也適用於「依路徑決定要不要跑」。
- 成本主要集中在**最頻繁的那一層**（每次工具呼叫的 PreToolUse），不在最慢的那一步（`go test`）。
