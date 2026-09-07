> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-07 - Frida attach Android process label, not package name

Status: validated

#### One-line Summary

`frida.get_usb_device().attach('<package>')` 常失敗；`adb pidof <package>` 仍有 PID。用 `enumerate_processes()` 或 **PID** attach，顯示名稱多半是 launcher 標題。

#### Human Explanation

Android 上 Frida 的 process name 是 `ps`/launcher label（例如短英文遊戲名），不是 `applicationId`。Python `attach('com.example.app')` 丟 `ProcessNotFoundError` 不代表 App 沒在跑。也不要為了找 CLI 去掃整個 home 目錄。

#### Trigger

- USB 裝置上 App 已前景，`adb pidof` 有數字。
- `attach(package)` 或 `frida -n <package>` 找不到 process。

#### Evidence

- Tool: `adb pidof` vs `frida.get_usb_device().enumerate_processes()`.
- Sanitized excerpt: PID matched; Frida name was a short store/launcher title, not the dotted package.
- Evidence path: target protocol / Frida notes under `<PROJECT_ROOT>`

#### Generalized Lesson

1. 先 `adb pidof <package>`。
2. `enumerate_processes()` 用 PID 或名稱子字串比對（忽略大小寫的 title）。
3. `attach(pid)`。記錄「Frida name ≠ package」以免下一輪重踩。
4. 本機 Frida：`python3 -c "import frida; print(frida.__version__)"`，不要 `find ~ -name frida`。

#### Agent Action

attach 失敗先列 process 表，不要改 spawn 或重裝 frida-server 當第一反應。

#### Goal / Action / Validation

- Goal: 對上正確 session。
- Action: pidof + enumerate + attach(pid)。
- Validation: `libil2cpp.so` / `libapp.so` 或 Java 執行期可列模組。

#### Applies When

- 授權 USB Frida；Android 前景 App。

#### Does Not Apply When

- spawn `-f <package>` 本來就走 package（那是啟動參數，不是 enumerate 名稱）。

#### Validation

同一 PID 在 adb 與 Frida 一致；script `READY`/`DONE` 能印出。

#### Promotion Target

- `analysis/apk/tools-and-failures.md`

#### Required Linked Updates

- 必須：tools-and-failures 加 Frida attach 一列
- 必須：`feedback/history/apk-analysis/README.md` Recent
- 專案 package／launcher 名留 `<PROJECT_ROOT>`
