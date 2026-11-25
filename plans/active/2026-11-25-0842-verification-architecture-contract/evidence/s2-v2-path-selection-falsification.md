# S2 — V2 驗證：路徑式測試選擇會漏掉讀取 repo 資料的測試

- 日期：2026-11-25
- 前提（S1 observed）：pre-push 只在 `scripts/ai-skill-cli/`、`scripts/git-hooks/`、`.githooks/`、CLI workflow 有變更時才跑 `go test ./...`；但 `internal/app` 有測試讀取 repo 其他路徑。
- 判定：**verified** — 本機 pre-push 對兩個探針都放行，且兩個探針都讓測試失敗。

## 方法

1. 從 `main`（`db862d62`）建立本機暫存 worktree 與兩個本機分支：探針分支，以及設為其 upstream 的基準分支。都沒有推送；結束後已刪除，`git ls-remote` 確認遠端沒有對應分支。
2. 每個探針只改一個非 CLI 路徑的檔案，先跑對應測試確認失敗，再提交並執行 `hooks run pre-push`。
3. 在 worktree 中，pre-commit 本身無法執行（見下方 §旁支觀察），所以用 `--no-verify` 提交。pre-push 的 `push_governance_replay` 會對每個待推 commit 重放 pre-commit 與 commit-msg 檢查，因此以它作為提交層的把關。

## 結果

| 探針 | 改動（模擬真實編輯） | 受影響的測試 | 測試結果 | pre-push |
| --- | --- | --- | --- | --- |
| P1 | `metadata/project/example-ai-skill-project.yaml` 的 `project.id` 改為空字串 | `TestLoadProjectMetadata_ExampleFileParsesCleanly` | FAIL（`project.id must be present and non-empty`） | exit 0：replay ok、kge_check ok、`cli_ci_preflight: skipped` |
| P2 | archived plan evidence 的 selector prompt 檔尾多一個換行 | `TestLookupPromptAndPackerPolicyUnchanged` | FAIL（selector prompt drift） | exit 0：同上 |

- 兩個改動在本機都會被推出去，pre-push 不會執行任何 Go 測試。
- GitHub CI 的 `paths` 過濾只含 `scripts/**` 與 CLI workflow 檔，`metadata/` 與 `plans/` 都不在內。依 workflow 定義，CI 也不會觸發。這一點是讀設定得出的，沒有實際推送驗證。
- 下一次有人改 CLI、觸發 `go test` 時，這些失敗才會出現，而且會被歸咎到不相關的 CLI 改動。

## 結論

- 契約 V2「不確定 → 執行」在 Ai-skill 自身**實際被違反**。原因不是快取，而是以路徑決定要不要跑測試，而選擇範圍沒有涵蓋測試實際讀取的輸入。
- 已知的反例：`review_architecture_doc_drift` 有 commit 時的 staged validator，同類依賴已被其他 gate 涵蓋。所以不是每個讀取 repo 的測試都有缺口，但至少 P1、P2 這兩類有。
- 修補方向（見 §裁決與修補）：
  1. **保守全跑**：pre-push 只要有任何變更就跑 `go test ./...`。Go 測試快取命中時約 2 s（S1）；未命中時，正是本來就該跑的時候。
  2. **擴大路徑**：把測試讀取的 repo 路徑加進選擇範圍。需要維護清單，而且新測試讀新路徑時又會漏，違反 V2 的精神。
  3. CI 的 `paths` 過濾要同步調整，否則本機與 CI 的涵蓋範圍不一致。

## 裁決與修補（2026-11-25）

- 使用者選方案 1（保守全跑），CI 一併調整。
- pre-push：`hasCLICIPreflightChange` 改為 `hasPrePushChange`，只要 upstream 以來有任何變更就跑 `go test ./...`；沒有 `scripts/ai-skill-cli/go.mod` 時跳過。
- CI：`.github/workflows/ai-skill-cli.yml` 移除 `push` / `pull_request` 的 `paths` 過濾。
- 代價：每次 push 多跑 `go test`（快取命中約 2 s），以及 `github_workflow_history` 查詢（逾時上限 5 s）。
- 已知同族：[`validation-coverage-gap-executor-placement.md`](../../../../enforcement/failure-patterns/validation-coverage-gap-executor-placement.md)（採樣 #2）。

## 旁支觀察（observed，不在本計畫範圍）

- 在 `git worktree` 中提交時，pre-commit 的 author-email-allowlist 讀取 `<worktree>/.git/info/ai-skill-allowed-author-emails` 失敗（`not a directory`）。worktree 的 `.git` 是檔案，不是目錄，所以 pre-commit 在任何 worktree 都無法執行。尚未查證影響範圍與修法。
