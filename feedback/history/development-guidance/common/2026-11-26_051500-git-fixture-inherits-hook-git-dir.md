> 遵守 [共用規則索引](../../../../enforcement/README.md) 與 [回饋規則](../../../feedback-lessons.md)；原始證據留在專案 repo 的 commit 與 session transcript。

### 2026-11-26 - 在 git hook 裡跑的 git 固定樣本，會把 `git config` 寫進真正的 repo

Status: candidate

#### Reuse Evidence

單一專案、兩支測試、同一個機制；另一台機器上已實際造成 5 個 commit 以錯誤身分推上共用主線。

#### One-line Summary

hook 執行時 git 會匯出 `GIT_DIR`／`GIT_WORK_TREE`／`GIT_INDEX_FILE`（以及 `git commit` 時的 `GIT_AUTHOR_*`／`GIT_COMMITTER_*`）。測試在暫存資料夾 `git init`、`git config user.name Fixture`，若子程序繼承這些變數，指令會作用在**被提交的那個真 repo**：本機設定被改成測試身分，之後的真實 commit 都掛著測試作者。

#### Human Explanation

1. 測試用 `spawnSync('git', …, { cwd: tmp })` 或 `git -C tmp`，以為已經隔離。但 `GIT_DIR` 環境變數優先於 cwd 與 `-C`。
2. 平常在終端機跑測試一切正常；只有從 pre-commit／pre-push hook 裡跑（且 git 有匯出 `GIT_DIR`，例如 linked worktree 或呼叫端已設定）才會中。所以很久都沒被發現。
3. 症狀不在測試本身：測試照樣通過，受害的是「之後」的 commit 作者欄位，而共用主線上的 commit 不能改寫。
4. 同一類還有作者變數：`git commit` 會把 `GIT_AUTHOR_NAME` 等匯出給 hook，讓固定樣本 repo 的作者斷言在 hook 裡失敗、在終端機通過。

#### Trigger

- 測試或工具在 hook 執行路徑中會呼叫 `git init`／`git config`／`git commit`。
- 共用主線出現 `fixture@example.invalid` 之類的測試身分作者。
- 同一支 git 固定樣本測試在終端機過、在 hook 裡失敗（或反過來）。

#### Evidence

- Tool: `git log --author`、對照兩支測試的 git helper；以「誘餌 repo」重現：`GIT_DIR=<decoy>/.git GIT_WORK_TREE=<decoy> node <舊版測試>` → 誘餌 repo 的 `user.name` 被寫成測試身分；修正版不寫。
- Evidence path: 專案 repo 的修正 commit（說明與重現方式在 commit message）。

#### Generalized Lesson

1. **任何會建立或設定 git repo 的固定樣本，第一件事是清掉 git 的 repo-local 環境變數**（`git rev-parse --local-env-vars` 列出的那組）與 `GIT_AUTHOR_*`／`GIT_COMMITTER_*`，而不是依賴 cwd 或 `-C`。最好集中在一個共用 helper，並讓靜態檢查找出「spawn git 卻沒用 helper」的測試。
2. **驗證隔離要用誘餌，不是看測試通過。** 設一個假的 `GIT_DIR` 指向誘餌 repo 再跑測試，事後檢查誘餌有沒有被動到；「測試綠」證明不了沒有副作用。
3. **身分類副作用要追到下游。** 測試身分出現在共用分支時，先找出是哪個測試的 helper 外洩，再通知受影響的機器檢查並清除本機 `.git/config` 的 `user.name`／`user.email`。
