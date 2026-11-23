# Ai-skill CLI Runtime

本目錄是跨平台 `ai-skill` CLI / runtime toolchain 的開發根目錄。目標是把現有 `scripts/` 中依賴 shell 或本機環境假設的流程，逐步升級成 Windows、macOS、Linux 可執行的單一 Go binary；runtime Ruby/Python surfaces 已移除。

## 目錄分層

| 路徑 | 用途 |
| --- | --- |
| [`docs/`](docs/README.md) | 文件先行產物：命令契約、舊腳本 parity 盤點、支援矩陣、BDD-lite 場景、fixture 計畫 |
| `cmd/ai-skill/` | Go CLI 入口；Phase 1 已建立 `doctor` skeleton |
| `internal/` | Go internal packages；目前包含 command dispatch、`doctor` checks、path normalization、JSON / plain output 與 exit code |
| `testdata/` | 未來測試 fixtures / golden 輸出（尚未建立） |

## Capability catalog

Discovery 的搜尋起點，也是 owner 的唯一記錄點（契約：[`lifecycle-contract.md`](../../workflow/cross-cutting/project-tooling/lifecycle-contract.md) §Capability catalog）。新增能力前先查這張表；同一概念的變化加參數或子命令，不另開入口。

| capability | entry | placement | state | owner | ttl |
| --- | --- | --- | --- | --- | --- |
| 建置版本資訊 | `ai-skill version` | Domain | maintained | CLI toolchain | — |
| 本機 runtime 與 repo 就緒檢查 | `ai-skill doctor` | Domain | maintained | CLI toolchain | — |
| 新專案 AI 工具 bootstrap | `ai-skill init-project` | Domain | maintained | tool bootstrap | — |
| 對話目標帳本 | `ai-skill goals` | Domain | maintained | conversation goal ledger | — |
| close-loop 收尾（dry-run / commit / push） | `ai-skill close-loop` | Domain | maintained | repo close-loop | — |
| Git hook 安裝與執行（全部鎖的執行入口） | `ai-skill hooks`；adapter `.githooks/*` | Domain | maintained | enforcement | — |
| Cursor bundle mirror 規劃（write mode 未實作） | `ai-skill sync-cursor-bundle` | Domain | maintained | tool bootstrap | — |
| runtime compile / refresh / validate / query | `ai-skill runtime` | Domain | maintained | runtime | — |
| Roo Code 設定守門 | `ai-skill roo` | Domain | maintained | tool bootstrap | — |
| Copilot guided bootstrap prompt | `ai-skill copilot` | Domain | maintained | tool bootstrap | — |
| glossary 條目驗證 | `ai-skill glossary` | Domain | maintained | knowledge glossary | — |
| Markdown checkbox 掃描 | `ai-skill scan-checkboxes` | Domain | maintained | plan lifecycle | — |
| enforcement registry lint / coverage | `ai-skill enforcement` | Domain | maintained | enforcement | — |
| plan-tree 呈現 | `ai-skill plans` | Domain | maintained | plan lifecycle | — |
| KGE check / validate（Logic 在 Shared `portable/kge`） | `ai-skill kge` | Domain | maintained | enforcement | — |
| workflow route-candidate proposals | `ai-skill router` | Domain | maintained | workflow routing | — |
| image-processing verify / fixtures | `ai-skill image-processing` | Domain | maintained | image-processing workflow | — |
| repo-local release binaries 建置（不經 dispatch） | `go run ./cmd/releasebuild` | Domain | maintained | CLI toolchain | — |
| Evidence Candidate 組裝（不經 dispatch；目的與輸入／輸出見下方註記） | `go run ./cmd/evidencecandidate` | Research | experimental | evidence candidate system | 2026-12-23 |

- `evidencecandidate`：
  - **實驗目的**：在來源 plan [`2026-06-16-1131-evidence-candidate-system.md`](../../plans/active/2026-06-16-1131-evidence-candidate-system.md) 的 observation period，把人工標註的 criteria_hits 組成 candidate。
  - **輸入／輸出**：寫在 [`cmd/evidencecandidate/main.go`](cmd/evidencecandidate/main.go) 檔頭的 contract。輸入是 artifact、criteria_hits、matched_plans；輸出是 candidate，只會寫入 inbox；沒有任何決策權限。
  - **到期處理**：TTL 2026-12-23。來源 plan 的 observation period 若更早結束，就改用那一天。到期時三選一：轉為 maintained、附理由續期，或 deprecated。到期不代表自動刪除。
- `scripts/git-hooks/*` 是 `hooks` 的舊路徑轉發外殼，state = deprecated：只服務仍設定 `core.hooksPath=scripts/git-hooks` 的 clone。移除條件：確認沒有 clone 使用舊路徑；使用情況無法機械確認，所以不做自動清理。

## 開發 gate

`docs/` 的 Phase 0 產物關卡已完成，Phase 1 可開始 Go implementation。後續新增命令時仍必須先對照 [`docs/command-contract.md`](docs/command-contract.md)、[`docs/script-parity-inventory.md`](docs/script-parity-inventory.md)、[`docs/bdd-scenarios.md`](docs/bdd-scenarios.md) 與 [`docs/test-fixture-plan.md`](docs/test-fixture-plan.md)。

## 開發指令

```bash
./bin/ai-skill-darwin-arm64 version --json
./bin/ai-skill-darwin-arm64 doctor --json
./bin/ai-skill-darwin-arm64 doctor --check-runtime --json
./bin/ai-skill-darwin-arm64 runtime validate --repo ../.. --json
./bin/ai-skill-darwin-arm64 glossary validate --repo ../.. --json
```

Phase 1 / Phase 3 採用 [`modernc.org/sqlite`](docs/dependency-policy.md) 作為 pure Go SQLite engine；`doctor --check-runtime` 已覆蓋 in-memory 與 temporary file-backed write / query / integrity proof。Git 維持 desktop external dependency；runtime core path 不依賴 Ruby、Python 或外部 `sqlite3` CLI。

## Repo-local binaries

Committed binaries live in [`bin/`](bin/) so the repo can run the CLI without a local Go install:

- `bin/ai-skill-darwin-arm64`
- `bin/ai-skill-darwin-amd64`
- `bin/ai-skill-linux-amd64`
- `bin/ai-skill-linux-arm64`
- `bin/ai-skill-windows-amd64.exe`
- `bin/BUILDINFO`
- `bin/SHA256SUMS`

Use the binary matching the host OS/architecture. Rebuild these files only after CLI source changes:

```bash
go run ./cmd/releasebuild --stable-names --version "repo-$(git rev-parse --short HEAD)" --commit "$(git rev-parse --short HEAD)" --dist bin
```

`go test ./...` verifies `bin/SHA256SUMS`, checks `bin/BUILDINFO` against the latest CLI source commit, and smoke-tests the current host binary. If CLI source changes, two commits are required: (1) the source change **together with** a full rebuilt `bin/` (pre-commit `validateReleaseBinaryParityStaged` blocks a CLI source commit without all five binaries + `BUILDINFO` + `SHA256SUMS`), then (2) a bin-only rebuild with `--commit <source commit>`, so `BUILDINFO` points at the latest CLI source commit (pre-push test). Use the source commit hash, not `HEAD`, if any other commit sits in between.

Release artifacts：`go run ./cmd/releasebuild` 會輸出 Windows amd64、macOS amd64/arm64、Linux amd64/arm64 binaries 與 `SHA256SUMS`。`ai-skill version` 支援 `-ldflags` 注入 version / commit / date。

GitHub Actions：`.github/workflows/ai-skill-cli.yml` 會在 Windows、macOS、Linux 執行 `go test ./...` 與 `doctor` smoke checks。CI 明訂 Go package 測試上限 30 分鐘、test job 上限 45 分鐘；file-backed SQLite fixtures 的批次建置使用 transaction，減少 Windows runner 的逐筆同步成本。修改 `scripts/ai-skill-cli/**`、`scripts/git-hooks/**` 或 workflow 行為時，必須同步檢查 `.github/workflows/ai-skill-cli.yml` 的 trigger、matrix、test command 與 artifact build steps。

Push 前可手動執行與 hook 相同的 CI preflight：

```bash
./bin/ai-skill-darwin-arm64 hooks run pre-push --repo ../.. --json
```

若上一輪 GitHub Actions 失敗，下一次 push 前至少先跑 `cd scripts/ai-skill-cli && go test ./...`，再跑上面的 pre-push hook runner。`pre-push` 只在本分支相對 upstream 改到 CLI、git hooks 或 CLI workflow 時查詢 GitHub 上最近一次 **completed** 的 `ai-skill CLI` workflow；若上一個 completed run 是紅燈，它會顯示 run URL，但不等待本次 push 的新 run。是否放行仍以本機 `go test ./...` 為準，避免 GitHub runner 太慢造成 workflow。

上游計畫：[`plans/active/2026-05-21-0834-cross-platform-go-script-runtime.md`](../../plans/archived/2026-05-21-0834-cross-platform-go-script-runtime.md)

Unity FX discovery wiring is checked by `runtime refresh` / `runtime validate`, including canonical-source FTS retrieval. See [the procedure](../../analysis/unity/workflows/unity-fx-driver-trace.md) and `internal/app/unity_fx_discovery_test.go` for positive and negative regression cases.

Feedback lesson closure validates structured transfer contracts and intelligence dispositions from index snapshots, including routed extraction targets and Discovery summaries. Run `go test ./portable/kge ./internal/app -run TestFeedback -count=1`; see [the review flow](../../feedback/extraction/lesson-review.md).
