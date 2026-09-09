> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-09 - Clear-data IL2CPP spawn needs delayed thread_attach

Status: validated

#### One-line Summary

`pm clear` 後立刻 spawn hook 時，`libil2cpp.so` 與 domain pointer 可能已可見，但 `il2cpp_thread_attach` 仍會 abort。等 domain 非空之後再延遲數秒再 attach。

#### Human Explanation

現有 Guest 的 force-stop spawn 用很短 delay 就能 hook。清資料／first-run 冷啟動較慢：過早 `thread_attach` 會 `Error: abort was called`，整段 login 窗口作廢。這不是 PID 錯或 frida-server 掛了。

這是對既有 spawn-vs-attach lesson 的補充：問題不只是 attach 太晚錯過 init，也可以是 spawn 太早撞上未就緒 runtime。

#### Trigger

- 剛 `pm clear` / first-run spawn。
- 腳本在 `il2cpp_thread_attach` abort；force-stop-only 同一腳本卻成功。

#### Evidence

- Tool: Frida spawn + IL2CPP export `il2cpp_thread_attach`。
- Sanitized excerpt: domain pointer passed a minimum-address check; 250 ms later attach aborted; ~2 s delay then hook install succeeded.
- Evidence path: names-only login probe notes under `<PROJECT_ROOT>`

#### Generalized Lesson

1. 等 `il2cpp_domain_get` 回傳看起來已初始化的 domain，**再** `setTimeout` 數秒才 `thread_attach`。
2. attach abort 時重試有上限，不要當成 hook offset 錯誤。
3. 在文件標 reset level：`clear app data` vs `force-stop only` 的 spawn delay 可以不同。

#### Agent Action

清資料 login/bootstrap 窗口預設延遲 thread_attach；不要把 abort 當 Frida 壞掉而重裝 server。

#### Goal / Action / Validation

- Goal: first-run spawn 裝上 names-only hooks。
- Action: domain ready + delayed attach + bounded retry。
- Validation: 腳本印 hookCount / ready，且前景仍是目標 package。

#### Applies When

- 授權 IL2CPP spawn；reset 含 `clear app data` 或 reinstall。

#### Does Not Apply When

- Java-only / 非 IL2CPP。
- 已在 lobby 的 attach 窗口（runtime 已熱）。

#### Validation

同一裝置：clear-data 短 delay abort；加長 delay 後 ready。

#### Promotion Target

- `workflow/apk-analysis/execution-flow.md` spawn 策略
- 既有 lesson：`common/2026-05-14_073700-frida-spawn-vs-attach-init-timing-no-buffer.md`

#### Required Linked Updates

- 必須：unity-il2cpp README 表格列
- 必須：apk-analysis README Recent
- 具體 ctor／packet 名留專案文件
