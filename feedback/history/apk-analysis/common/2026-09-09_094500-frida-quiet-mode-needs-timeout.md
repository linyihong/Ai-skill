> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-09 - Frida CLI quiet mode exits after -l unless -t

Status: validated

#### One-line Summary

Frida 17 的 `-q` 會在載入 `-l` / `-e` 後立刻結束 session；長窗口 spawn/attach 必須加 `-t <seconds>`，不要依賴本機 `timeout(1)`。

#### Human Explanation

分析腳本用 `console.log` 加 `setTimeout` 以為能撐住 REPL。加上 `-q` 後 CLI 在 script load 完成就退出，hook 還沒裝上或只印了一行 `Spawned`。macOS 常沒有 GNU `timeout`，外層包 `timeout 70 frida ...` 會直接 command not found，誤以為 capture 成功（exit 0）。

#### Trigger

- `frida -U -f <package> -l script.js -q` 約一秒結束，log 只有 spawn/resume。
- 腳本裡有數十秒 `setTimeout` / `LOGIN_TRACE done`，但檔案幾乎是空的。

#### Evidence

- Tool: Frida 17 CLI `--help`（`-q` quiet and quit after `-l` and `-e`；`-t` seconds in quiet mode）。
- Sanitized excerpt: spawn+resume then immediate exit; adding `-q -t 70` kept the session until hooks fired.
- Evidence path: target capture notes under `<PROJECT_ROOT>`

#### Generalized Lesson

1. 長 capture 用 `-q` 時**一定**帶 `-t`，秒數大於腳本 `done` timer 與 IL2CPP 延遲。
2. 不要用 GNU `timeout` 當跨平台保活；用 Frida 自己的 `-t`。
3. 不要把 `-q` 的 exit 0 當成 hook 已跑完；要看到腳本自己的 ready/done 事件。

#### Agent Action

寫自動化 spawn 時預設 `frida -U -f <package> -l <script> -q -t <seconds>`。若只要 REPL，不要加 `-q`。

#### Goal / Action / Validation

- Goal: quiet Frida 真的跑完 bounded probe。
- Action: `-q` 搭配 `-t`；檢查 ready 事件。
- Validation: log 含 hook ready，且 process 存活超過 script load。

#### Applies When

- 授權 USB Frida CLI 17；spawn 或 attach 需要腳本自己計時結束。

#### Does Not Apply When

- 互動 REPL（無 `-q`）。
- 一次性 dump 腳本在 load 當下就印完並故意退出。

#### Validation

同一腳本：無 `-t` 立刻結束；有 `-t` 後出現 ready。

#### Promotion Target

- `analysis/apk/tools-and-failures.md` Frida spawn 範例
- `workflow/apk-analysis/execution-flow.md` Frida 部署策略

#### Required Linked Updates

- 必須：tools-and-failures spawn 命令加 `-q -t`
- 必須：apk-analysis feedback README Recent + common 計數
- 專案 package／log 檔名留 `<PROJECT_ROOT>`
