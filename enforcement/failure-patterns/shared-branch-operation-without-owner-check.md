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

## Risk

- 覆蓋或干擾其他 agent / 使用者未提交的工作
- Push 了未經目標分支最新狀態驗證的內容（本地綠、合併後紅）
- 違反使用者明確的「不要動這個檔案」邊界而不自知
- 在共用 remote 上造成不可逆的 history 變更

## Required Agent Action

在對共用分支執行任何 history operation（merge / rebase / fast-forward / push）**之前**，依序完成：

1. **列出 worktree 與擁有者**：確認目標分支 checked out 在哪裡；若不在目前 worktree，該處可能有他人活動。
2. **量測落差**：分別確認目前分支相對目標分支的 ahead / behind 數量，不要只看其中一邊。
3. **檢查目標分支是否已包含本輪「受限檔案」**：使用者說過不要動的路徑，要在目標分支上實際查一次是否已存在。
4. **檢查目標 worktree 的 dirty 狀態**：若有未提交變更，先確認與本輪變更是否有路徑重疊；有重疊就停止並詢問。
5. **落差存在時，先同步再驗證**：rebase / merge 之後**重跑**驗證，不可沿用同步前的測試結果。
6. **操作後覆核並回報**：確認他人工作完好，並在最終回覆中明確說明曾在共用 worktree 上操作。

若步驟 1–4 任一項顯示有其他活躍擁有者且範圍可能重疊，停止並詢問使用者，不要自行判斷「應該沒差」。

## Prevention Gate

準備執行 merge / push 到共用分支前，agent 必須能回答：

- 目標分支目前 checked out 在哪個 worktree？那裡是否乾淨？
- 我的分支相對目標分支 ahead 幾個、behind 幾個？
- 目標分支是否已經包含本輪被要求「不要修改」的檔案？
- 同步（rebase / merge）之後，我是否**重新**跑過驗證，而不是引用同步前的結果？
- 若我在別人的 worktree 上操作過，我是否已覆核並在回覆中說明？

任一項答不出來，就不要執行 history operation。

## 驗證

1. 操作前的落差量測與 worktree 盤點有實際指令輸出，而非推測
2. 同步後的驗證結果與同步前分開記錄，可看出是重跑而非沿用
3. 操作後對他人未提交工作做過重疊檢查，且結果寫進最終回覆
4. 最終 `git status --short --branch` 於**所有**被觸及的 worktree 皆已回報

## Linked Rules

- [`../conversation-goal-ledger.md`](../conversation-goal-ledger.md) — git history operations 屬 `non-parallelizable`；owner / lock 判斷
- [`../dependency-reading.md`](../dependency-reading.md) — writeback transaction 的 commit / push / readback / clean status 條件
- [`../failure-learning-system.md`](../failure-learning-system.md) — `parallelization-risk` 分類
- [`commit-before-validation-skip.md`](commit-before-validation-skip.md) — 同家族：commit/push 前跳過驗證；本檔補的是「共用分支與他人擁有者」面向
- [`mandatory-step-blocker-bypass.md`](mandatory-step-blocker-bypass.md) — 若同步後驗證被環境阻斷，依該 pattern 停止並通知使用者
