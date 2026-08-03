# Shared-branch Operation Without Owner Check（未確認擁有者就在共用分支／他人 worktree 上操作 git history）

Status: candidate
Class: `parallelization-risk` / `validation-gap`

## Trigger

當使用者要求「把改動放到主線」「commit 並 push」「合併過去」，而 agent 目前在 feature branch、worktree 或 detached checkout 時，使用此 pattern。

具體觸發信號：

- 目標分支（main / trunk / 主線）checked out 在**另一個** worktree
- 目前分支與目標分支有落差（ahead / behind / diverged）
- 專案同時有其他 session、agent 或使用者在活動（dev server 在跑、背景任務、近期 commit）
- 使用者曾說某個檔案「不要動」，但該檔案可能已在目標分支上存在

## Failure Mode

Agent 把「push 到主線」當成單純的 `commit` + `push` 兩步，直接在目標分支所在的 worktree 執行 merge / fast-forward / push，而沒有先建立操作前的狀態快照。

Git history operations 在 [`conversation-goal-ledger.md`](../conversation-goal-ledger.md) 中被明列為 `non-parallelizable`，但該規則的檢查點是 `.agent-goals/` lock；當另一個 session 沒有建立 lock（或使用不同工具）時，agent 若不主動探測，就會在毫無警覺的情況下改動別人正在使用的工作區。

三個常見的具體漏檢：

1. **落差未檢**：假設目前分支是最新的，實際上落後目標分支數個 commit；直接 push 會失敗，或在 force 之下覆蓋他人工作。
2. **「不要動的檔案」已存在於目標分支**：使用者在任務開始時說某檔案不要碰，agent 在舊 base 上看不到它，於是自行建立了同名檔案；rebase 後才發現衝突或重複。
3. **目標分支的 worktree 有他人未提交工作**：agent 在該 worktree 執行 checkout / merge，可能觸及他人 staged / unstaged 變更。

即使結果碰巧無害（無路徑重疊、fast-forward 乾淨），**檢查發生在操作之後就不算 prevention**。

### 變體：同一 working tree 的共存 session（preflight 盲點）

上述檢查全部是**快照**，且都假設「其他擁有者在別的 worktree」。當兩個 agent session 共用**同一個** working tree 時，這些檢查會全部回報乾淨，卻完全偵測不到共存者：

- `git worktree list` 只有一筆 → 看起來沒有其他 worktree
- `.agent-goals/locks/` 空 → 對方若用不同工具或未建 lock 就查不到
- `ahead/behind` 為 0 → 只反映查詢當下那一刻

具體後果：agent 在 `git add` 之後、`git commit` 之前，共存 session 執行了自己的 `git add -A` + `commit`，於是 agent 的 staged 檔案被**捲進對方的 commit**，掛在一個與內容無關的 message 下；agent 自己的 `git commit` 回報 "nothing to commit"。之後對方 push，agent 的 push 因 ref lock 失敗（remote 已前進）。內容不會遺失，但 authorship 與 commit message 已錯置。

追加要求：

- **不要只在開始時檢查一次**。`git add` 與 `git commit` 之間若夾雜其他工具呼叫，commit 前重新確認 HEAD 未移動（先記錄 HEAD SHA，commit 前比對）。
- **主動偵測共存訊號**：working tree 出現不是自己改的 modified 檔案、背景 dev server／任務在跑、`git log` 出現分鐘級的新 commit。任一出現就假設有共存者，並在回報中說明。
- **commit 後立即覆核**：確認 `git log -1` 的 message 與 file list 確實是自己的。若被併入他人 commit，**停止並回報**，不要 amend 或 reset 他人剛產生的 commit。
- **push 被 ref lock 拒絕時先 fetch 再判斷**：對方可能已把你的內容一併推上去，此時不需要也不應該再 push；盲目 rebase／force 會製造重複或覆蓋。

### 變體：目標分支本身已領先 remote（push 夾帶他人 commit）

前述檢查都在問「我的分支 vs 目標分支」。還有一個方向沒被覆蓋：**目標分支 vs 它的 remote**。

當目標分支在 agent 動手前就已經 ahead（既有未推送 commit），`push <remote> <target>` 推的不只是本輪的 commit，而是**連同那些既有 commit 一起發佈**。Git 沒有「只推最上面那一個」的選項，這是 push 語意的必然結果，不是可以繞過的細節。

風險在於這些既有 commit 不是本輪產生的，agent 對它們一無所知：可能是他人或前次 session 刻意留在本地、尚未定案、等待 review，或含有還不打算公開的內容。發佈到共用 remote 之後，回退需要 force push，成本遠高於事前確認。

要求：

- **push 前分別回報兩個數字**：本輪新增幾個 commit、既有未推送幾個。只看「ahead 幾個」會把兩者混為一談。
- **若既有未推送 commit 存在，在 push 之前**列出它們（`git log <remote>/<target>..<target> --oneline`）並向使用者說明「這次 push 會一併發佈這些」，取得確認後再執行。
- **「使用者授權 push」不等於「使用者知道會夾帶什麼」**。授權涵蓋動作，不涵蓋 agent 事後才發現的範圍擴大；範圍與授權不一致時要重新確認，而不是自行判斷「反正也是他自己的 commit」。
- 事後才在最終回覆中補充說明，**不算 prevention**——與本檔開頭「檢查發生在操作之後就不算 prevention」同一原則。

## Risk

- 覆蓋或干擾其他 agent / 使用者未提交的工作
- 把使用者尚未打算公開的既有 commit 一併推上共用 remote（回退需 force push）
- Push 了未經目標分支最新狀態驗證的內容（本地綠、合併後紅）
- 違反使用者明確的「不要動這個檔案」邊界而不自知
- 在共用 remote 上造成不可逆的 history 變更

## Required Agent Action

在對共用分支執行任何 history operation（merge / rebase / fast-forward / push）**之前**，依序完成：

1. **列出 worktree 與擁有者**：確認目標分支 checked out 在哪裡；若不在目前 worktree，該處可能有他人活動。
2. **量測落差（兩個方向都要）**：確認目前分支相對目標分支的 ahead / behind，**以及目標分支相對其 remote 的 ahead / behind**。後者決定 push 會不會夾帶既有未推送 commit；若非零，先列出並向使用者說明後才能 push。
3. **檢查目標分支是否已包含本輪「受限檔案」**：使用者說過不要動的路徑，要在目標分支上實際查一次是否已存在。
4. **檢查目標 worktree 的 dirty 狀態**：若有未提交變更，先確認與本輪變更是否有路徑重疊；有重疊就停止並詢問。
5. **落差存在時，先同步再驗證**：rebase / merge 之後**重跑**驗證，不可沿用同步前的測試結果。
6. **操作後覆核並回報**：確認他人工作完好，並在最終回覆中明確說明曾在共用 worktree 上操作。

若步驟 1–4 任一項顯示有其他活躍擁有者且範圍可能重疊，停止並詢問使用者，不要自行判斷「應該沒差」。

## Prevention Gate

準備執行 merge / push 到共用分支前，agent 必須能回答：

- 目標分支目前 checked out 在哪個 worktree？那裡是否乾淨？
- 我的分支相對目標分支 ahead 幾個、behind 幾個？
- 目標分支相對它的 remote ahead 幾個？其中哪些**不是**本輪產生的？使用者知道這次 push 會一併發佈它們嗎？
- 目標分支是否已經包含本輪被要求「不要修改」的檔案？
- 同步（rebase / merge）之後，我是否**重新**跑過驗證，而不是引用同步前的結果？
- 若我在別人的 worktree 上操作過，我是否已覆核並在回覆中說明？

任一項答不出來，就不要執行 history operation。

## 驗證

1. 操作前的落差量測與 worktree 盤點有實際指令輸出，而非推測；落差量測涵蓋「目前分支 vs 目標分支」與「目標分支 vs remote」兩個方向
2. 若 push 會夾帶非本輪的既有 commit，確認的時間點在 push **之前**（可從對話順序反查），而非最終回覆才補述
3. 同步後的驗證結果與同步前分開記錄，可看出是重跑而非沿用
4. 操作後對他人未提交工作做過重疊檢查，且結果寫進最終回覆
5. 最終 `git status --short --branch` 於**所有**被觸及的 worktree 皆已回報

## Linked Rules

- [`../conversation-goal-ledger.md`](../conversation-goal-ledger.md) — git history operations 屬 `non-parallelizable`；owner / lock 判斷
- [`../dependency-reading.md`](../dependency-reading.md) — writeback transaction 的 commit / push / readback / clean status 條件
- [`../failure-learning-system.md`](../failure-learning-system.md) — `parallelization-risk` 分類
- [`commit-before-validation-skip.md`](commit-before-validation-skip.md) — 同家族：commit/push 前跳過驗證；本檔補的是「共用分支與他人擁有者」面向
- [`mandatory-step-blocker-bypass.md`](mandatory-step-blocker-bypass.md) — 若同步後驗證被環境阻斷，依該 pattern 停止並通知使用者
