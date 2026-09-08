# Named Work Unit Executed Without Claim（有名稱的工作單元未宣告擁有者就被第二條線重做）

Status: candidate
Class: `parallelization-risk` / `handoff-gap`

## Trigger

使用者要求執行一份計畫裡**已命名**的工作單元（slice、Phase、numbered walkthrough），而該 repo 同時有其他 session／agent 在同一條共用 branch 上活動。

具體觸發信號：

- 計畫檔列出編號切片，且尚未全部關閉
- `git log` 或 remote 上已有標題／路徑命中同一 slice id 的近期 commit
- 工作樹對該 slice 的候選路徑已 dirty，但本 session 還沒宣告 owner
- `.agent-goals/` 沒有對應該 slice id 的 `single-owner` 列與 lock

## Failure Mode

Agent 把「清單上還沒打勾」讀成「還沒人做」。它沒有先搜同一 slice id 的既有實作，就再開一條獨立 Executor 把同一批 acceptance 再做一遍。

檔案尚未重疊時，`.agent-goals/` 的 path-overlap 檢查與 git worktree 快照都會回報乾淨——兩個 session 各自建了平行的 `<pkg>/api`、平行的測試改寫、平行的 evidence 檔。衝突要到其中一條線 push、另一條 merge 時才出現，那時已經是語意衝突（add/add、同一測試兩種實作），不是文字衝突。

這與 [`shared-branch-operation-without-owner-check.md`](shared-branch-operation-without-owner-check.md) 不同：那條管的是 **git history 操作前**有沒有盤點 branch／worktree 擁有者；本條管的是 **開始實作一個有名稱的工作單元前**有沒有宣告「這格已有人」。Goal ledger 的 lock 綁的是 conversation goal，不是 plan 上的 slice id；只檢查 lock 目錄仍會漏。

## Risk

- 同一 acceptance 被清兩次，合併時必須在兩套實作之間取捨，驗證成本落在後到的那條線
- 較晚的那條線可能覆寫較完整的入口（或相反），而棘輪／閘只看見「列已刪」
- 計畫上的「未關閉」被當成待辦，實際上遠端已經 `slice_compliant_closed`

## Required Agent Action

**有名稱的工作單元在派發或動手之前必須先被認領。認領失敗就停，不要開第二條實作。**

1. **用 slice id 搜，不要只用「清單還沒勾」。** 查 `git log <remote>/<branch>..HEAD`、遠端最近 commit、工作樹 dirty 路徑、`evidence/` 檔名、以及 `.agent-goals/` 的 Planning 連結是否已指向同一 `slice_id`。
2. **若已有實作（已 commit、已 push、或工作樹已在改那些檔）：停止。** 向使用者列出那條線的 SHA／路徑／關閉狀態，問要接手、合併還是放棄本線。不要默默重做。
3. **若確定無人認領：先寫 goal ledger。** `parallelization: single-owner`（同一 slice 的實作預設不是 parallelizable），Planning 連結寫 plan 路徑 + `slice_id`，再建立 lock，然後才派 Executor 或改產品檔。
4. **遠端在執行期間前進時，把「同 slice 的新 commit」當成衝突訊號**，即使目前檔案還沒 overlap。先 fetch 再決定，不要把第二份實作做到 merge 才發現。

## Prevention Gate

動手或派發 Executor 前：

- 這個 slice id 在本地、遠端、工作樹、goal ledger 四處搜過了嗎？
- 若已有證據檔或刪債 commit，為什麼還要再開一條實作？
- goal file 的 Planning 連結有沒有寫 `slice_id`，lock 有沒有建？
- 只看到計畫列未勾，有沒有可能是另一條線已經做完、只是還沒改那一格？

## 驗證

1. 開始實作前的紀錄裡，有對該 `slice_id` 的搜尋結果（無命中或已停止詢問）
2. `.agent-goals/` 有對應該 slice 的 `single-owner` 列與 lock，或明確標 not-applicable（單人、單 session、計畫無編號切片）
3. 沒有在已知已關閉或進行中的 slice 上再開一條平行實作
4. 若發生過重工，最終回覆說明取捨依據，而不是默默覆蓋

## Linked Rules

- [`shared-branch-operation-without-owner-check.md`](shared-branch-operation-without-owner-check.md) — 共用 branch 的 git 操作前檢查；本條是其「開始實作前」的姊妹項
- [`../conversation-goal-ledger.md`](../conversation-goal-ledger.md) — lock 與 `single-owner`；slice id 必須寫進 Planning 連結，否則 lock 綁不到工作單元
- [`deferred-item-not-backfilled-to-target-phase.md`](deferred-item-not-backfilled-to-target-phase.md) — 同家族：清單上的狀態與實際工作狀態脫鉤
- [`../failure-learning-system.md`](../failure-learning-system.md) — `parallelization-risk` / `handoff-gap`
