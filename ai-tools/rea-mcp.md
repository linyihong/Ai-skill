# REA MCP／CLI（可選工具 adapter）

**預設不需要本檔。** Canonical 調查流程在 [`../analysis/reverse-engineering/investigation-process.md`](../analysis/reverse-engineering/investigation-process.md)，用 jadx／Ghidra／Frida 等通用工具即可。

本檔只在使用者**明確要**接 [REA](https://github.com/morluto/rea)（`rea-agents`）MCP 時使用。不 vendor REA skill 全文；不把安裝 REA 當分析能力完成條件。

## 何時需要（少數情況）

- 已決定用該 MCP 當 agent 呼叫介面，且本機有 JADX／Ghidra／Hopper 等。
- **不需要**：建立／執行 Ai-skill reverse-engineering 能力；已有 CLI 工作流。

## 前置（REA 不安裝這些）

| 用途 | 前置 |
| --- | --- |
| 執行 | Node.js 22.19+／24.11+／26+ 與 npm（以上游 README 為準） |
| Android 靜態 | 完整 JDK 17+（含 compiler）、headless JADX MCP JAR；環境變數如 `REA_JADX_MCP_JAR`、`JAVA_HOME` |
| Native 深層 | 既有 Hopper／Ghidra／IDA 之一 |
| 靜態 JS | 通常只需 Node |

## 建議操作

1. 只讀診斷：`npx -y rea-agents@latest doctor --client cursor --json`（或其他 client id）。
2. Dry-run setup：`npx -y rea-agents@latest setup --client cursor --dry-run --json`——**審查**將改寫的路徑與備份。
3. 使用者明確同意後：`setup --client cursor --yes`（Hopper 安裝需另旗標與另同意）。
4. 重啟／重連 agent；以 live `tools/list`／`binary_session` 為準，勿假設舊 npm 與 main 同能力。
5. 更新：`rea update` 或 `npx rea-agents@latest setup`；再審 diff。

## 與 Ai-skill 的邊界

| 可以 | 不可以 |
| --- | --- |
| 當 `analysis_engine` 可選實作 | 把 REA skill 鏡像進 `analysis/`／`workflow/` |
| 失敗時標 `engine_unavailable` | 因未安裝 REA 而跳過 authorization |
| 結果遵守 `analysis/reverse-engineering/evidence-contract.md` | 把 raw token／私人路徑寫進本庫 |

## Cursor

短 pointer 見 [`agent/cursor.md`](agent/cursor.md)。專案 MCP 設定屬本機／使用者設定，不進 Ai-skill canonical（除非團隊另有決策）。

← [ai-tools/](README.md)
