# Git hooks（canonical，隨 repo 版本化）

本目錄是 **Ai-skill 本 repo 的 canonical Git hook 來源**。業務邏輯在 Go CLI（`ai-skill hooks run …`），這裡只有薄 adapter。

## 首次 clone 後（必做）

```bash
ai-skill hooks install
```

會把本機 `core.hooksPath` 設為 `.githooks`（相對 repo root）。之後每次 `git commit` 都會跑：

- **pre-commit** — runtime validate、**sanitization_scan**（shared-layer 去敏）、knowledge 檢查
- **commit-msg** — plan-tree / cognitive contract / 其他治理 validators
- **post-commit** / **pre-push** — 閉環與 CI preflight

驗證安裝：

```bash
git config --get core.hooksPath   # 應為 .githooks
ai-skill doctor --plain
```

## 為什麼不用手動複製到 `.git/hooks/`

`core.hooksPath = .githooks` 讓 hook **跟著 git pull 更新**，不需每人維護一份複本。`scripts/git-hooks/` 保留為 legacy 轉發路徑。

## 限制

- `git commit --no-verify` 仍可跳過本機 hook；請依賴 CI / review 作為第二道防線。
- 未執行 `hooks install` 的 clone **不會**自動啟用 — 這是 Git 安全模型限制，沒有 repo 內建「零設定強制 hook」。

← 詳見 [`metadata/project/README.md`](../metadata/project/README.md)、[`enforcement/sanitization-mechanical.md`](../enforcement/sanitization-mechanical.md)
