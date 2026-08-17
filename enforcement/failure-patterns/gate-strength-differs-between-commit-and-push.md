# Gate Strength Differs Between Commit And Push（commit 與 push 的閘強度不同）

Status: candidate
Class: `process-gap`

## Trigger

專案的驗證分階段：pre-commit 跑一部分檢查，pre-push 跑完整套。而其中某些檢查在兩階段的**嚴格度不同** —— commit 時 `mode=warn`，push 時 `mode=block`。

Agent 連續提交數個 commit，全部順利通過，然後推送。

## Failure Mode

Agent 把「commit 成功」讀成「這次變更是乾淨的」。但 warn 級的檢查不會阻止提交，其輸出混在大量綠色訊息裡，通常不會被注意。問題於是**累積數個 commit**，在推送時一次爆發。

代價不只是多一次來回：

- 每次推送失敗要付一次完整驗證的時間（含 e2e），而失敗訊息只揭露**第一個**擋下的閘 —— 修完它推第二次，可能被第二個閘擋下
- 修正需要新的 commit，於是歷史上留下一串「修上一個 commit 沒過的閘」的提交
- 若 agent 在推送尚未落地時就回報「已完成」，該回報是錯的

觀察到的實例：`check_fe_doc_sync`（FE 行為變更需契約文件）與 `verify-bdd-scenario-coverage`（測試標題須對應 Scenario 或標記豁免）皆為 commit-warn／push-block，於是連續兩次推送各被其中一個擋下。

## Risk

- 反覆盲推，每次付完整驗證成本卻只換到一個失敗訊息
- 錯把「commit 通過」當成品質訊號，在未推送前宣告完成
- 修正 commit 稀釋歷史，且訊息描述的是流程而非變更意圖

## Required Agent Action

**推送前先在本地跑完整的 push 模式驗證，而不是把推送當成第一次完整驗證。**

1. **確認專案是否有分階段閘與強度差異。** 通常在 `verify.sh`／hook 腳本中以 `mode=warn`／`mode=block` 或「deferred until git push」字樣標示 —— 後者本身就是一份清單，說明哪些檢查 commit 時不會擋。
2. **變更觸及以下任一類時，提交前主動跑 push 模式**：測試標題／新增測試檔、FE 可觀察行為、文件契約、產生器輸出、manifest。
3. **不以「commit 通過」作為變更乾淨的證據。** 那只證明通過了 commit 階段的子集。
4. **推送失敗後，先跑完整本地驗證再推第二次。** 直接修掉第一個訊息就重推，只會換到第二個閘的訊息。
5. **推送未以 ref 狀態確認落地前，不宣告完成。** 見 [`proxy-signal-substituted-for-target-state.md`](proxy-signal-substituted-for-target-state.md)。

## Prevention Gate

推送前：

- 這個專案有哪些檢查是 commit 時 warn、push 時 block？我這次的變更碰到它們了嗎？
- 我在本地跑過 push 模式的完整驗證了嗎，還是打算讓推送替我跑？
- 我已經推送失敗過一次了嗎？若是，我修的是全部問題還是只有第一個訊息提到的那個？

## 驗證

1. 推送前有一次本地完整 push 模式驗證的紀錄，且其結論為通過
2. 推送落地以 ref 狀態確認（雙端 rev 相同、ahead/behind 為 0）
3. 若曾推送失敗，修正的範圍以完整本地驗證界定，而非以失敗訊息的第一項界定
4. 採用「提交後先跑本地 push 模式、通過才推」的節奏後，不再出現連續盲推（實測：改用此節奏前連續兩次推送各被不同閘擋下，改用後零次）

## Linked Rules

- [`proxy-signal-substituted-for-target-state.md`](proxy-signal-substituted-for-target-state.md) — push 成敗只以 ref 狀態判定
- [`commit-before-validation-skip.md`](commit-before-validation-skip.md) — 同家族：以「已提交」替代「已驗證」
- [`path-limited-commit-omits-untracked-files.md`](path-limited-commit-omits-untracked-files.md) — 同屬提交階段的 process-gap
- [`../failure-learning-system.md`](../failure-learning-system.md) — `process-gap` 分類
