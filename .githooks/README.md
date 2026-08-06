# Git hooks（canonical，隨 repo 版本化）

本目錄是 **Ai-skill 本 repo 的 canonical Git hook 來源**。業務邏輯在 Go CLI（`ai-skill hooks run …`），這裡只有薄 adapter。

## 首次 clone 後（必做）

```bash
ai-skill hooks install
```

會把本機 `core.hooksPath` 設為 `.githooks`（相對 repo root）。之後每次 `git commit` 都會跑：

- **pre-commit** — **author email allowlist**（opt-in，見下）、runtime validate、**sanitization_scan**（shared-layer 去敏）、knowledge 檢查
- **commit-msg** — plan-tree / cognitive contract / 其他治理 validators
- **post-commit** — 閉環輔助
- **pre-push** — **push governance replay** + **`kge check`（D9；validation 可 block、advisory 不擋）** + CLI CI preflight

### Author email allowlist（本機、不進 commit）

若要限制「只有某個 email 能 commit」（例如避免 global 誤設成公司信箱），在本機設定白名單即可；**未設定時不限制**：

```bash
# 只允許這個 email（可 --add 多次；換成你的允許地址）
git config --local --add ai-skill.allowedAuthorEmail <USER_EMAIL>

# 或用檔案（同樣不會被 git 追蹤）
printf '%s\n' '<USER_EMAIL>' > .git/info/ai-skill-allowed-author-emails
```

同時請把實際身份設對：

```bash
git config --local user.email <USER_EMAIL>
git config --local user.name <USER>
```

關閉白名單：`git config --local --unset-all ai-skill.allowedAuthorEmail` 並刪除 `.git/info/ai-skill-allowed-author-emails`（若有）。

Adapter 解析順序：`AI_SKILL_CLI` → `bin/ai-skill-<os>-<arch>` → `bin/ai-skill`（後者常為本機 alias，可能過期）。

驗證安裝：

```bash
git config --get core.hooksPath   # 應為 .githooks
ai-skill doctor --plain
```

## 為什麼不用手動複製到 `.git/hooks/`

`core.hooksPath = .githooks` 讓 hook **跟著 git pull 更新**，不需每人維護一份複本。`scripts/git-hooks/` 保留為 legacy 轉發路徑。

## 限制

- `git commit --no-verify` 仍可跳过**本机 commit 阶段** hook；**pre-push 会重跑 sanitization + commit-msg validators**，未通过时需 `git reset --soft @{u}` 后重新 commit（不用 `--no-verify`）再 push。
- `git push --no-verify` 仍可跳过 pre-push；请依赖 CI / branch protection 作为第三道防线。
- 未执行 `hooks install` 的 clone **不会**自动启用 — 这是 Git 安全模型限制，没有 repo 内置「零设定强制 hook」。

← 詳見 [`metadata/project/README.md`](../metadata/project/README.md)、[`enforcement/sanitization-mechanical.md`](../enforcement/sanitization-mechanical.md)
