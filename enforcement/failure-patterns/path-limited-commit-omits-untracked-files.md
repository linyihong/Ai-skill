# Path-Limited Commit Omits Untracked Files（路徑限定提交漏掉新增檔案）

Status: candidate
Class: `process-gap`

## Trigger

Agent 在共用 worktree（多 session／多 agent 同時工作）中提交，因此**正確地**避開 `git add -A`，改用路徑限定：

```
git commit -- <dir>
git commit -m "..." -- src/Foo tests/Foo
```

而本次變更**產生了新檔案** —— 特別是產生器輸出：OpenAPI client、codegen、snapshot、migration。

## Failure Mode

`git commit -- <path>` 只提交該路徑下**已追蹤檔案的修改**。未追蹤的新檔案不在其中，且**不會有任何警告**。

這個陷阱之所以難察覺，是因為 agent 為了遵守「共用 worktree 禁用 `git add -A`」而採用路徑限定 —— 一個正確的決定 —— 然後把它誤當成 `git add -A` 的完整等價物。提交成功、gate 全綠（gate 讀的是工作樹，新檔案在那裡）、`git status` 只剩幾個未追蹤項而容易被當成雜訊。

失效在別人 clone 或 CI 乾淨建置時才顯現：產生的 client 檔不存在，建置紅。

## Risk

- 產生器輸出缺漏，本機全綠而 CI／他人環境紅，且錯誤訊息指向缺檔而非缺提交
- 後續 slice 重新產生時把「新增」誤判為「本次變更」，污染下一個 commit 的範圍
- 若該 slice 已宣告關閉，缺漏會被繼承進交接文件的「已完成」

## Required Agent Action

**新檔案必須明確 `git add`，路徑限定只覆蓋修改。**

1. **提交前先看未追蹤清單。** `git status --porcelain` 中的 `??` 項目逐一判斷歸屬，不當成雜訊略過。
2. **產生器跑過就假設有新檔案。** codegen／OpenAPI／snapshot／migration 的正常產出就是新增檔案。
3. **明確 `git add <具體路徑>` 再提交** —— 這仍然符合共用 worktree 的約束，因為指名了路徑；被禁止的是 `-A` 的無差別暫存，不是 `add` 本身。
4. **提交後驗證目標狀態**：`git status --porcelain` 對相關路徑應為空。不看回執，看狀態。

## Prevention Gate

路徑限定提交前：

- 本次是否執行過任何產生器？若有，新檔案在哪裡？
- `git status --porcelain` 有 `??` 嗎？每一項我都判斷過歸屬了嗎？
- 我是不是把「避開 `git add -A`」誤解成「不要用 `git add`」？

## 驗證

1. 提交後該 slice 相關路徑的 `git status --porcelain` 為空
2. 產生器輸出目錄在 commit 的 `--stat` 中出現（若本次執行過產生器）
3. 若曾發生漏提交，補提交獨立成 commit 並在訊息中說明，而非混入下一個功能 commit

## Linked Rules

- [`shared-branch-operation-without-owner-check.md`](shared-branch-operation-without-owner-check.md) — 共用 worktree 的操作約束來源
- [`proxy-signal-substituted-for-target-state.md`](proxy-signal-substituted-for-target-state.md) — 同家族：以「commit 指令成功」替代「檔案已入庫」
- [`../failure-learning-system.md`](../failure-learning-system.md) — `process-gap` 分類
