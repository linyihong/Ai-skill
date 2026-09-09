> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-09 - Extract preference key names on-device, never cat values

Status: validated

#### One-line Summary

Unity PlayerPrefs / Android SharedPreferences 的 XML 含 token。只列檔名，並在裝置上抽出 `name=` 鍵名；不要 `cat`／pull 整檔進 transcript。Magisk `su 0 ls -la` 會把 `-la` 當 su 選項。

#### Human Explanation

要證明 session 存在本機，列出 `shared_prefs/*.xml` 檔名通常就夠。把 XML 拉到 host 或 `cat` 進 agent log 等於匯出 token。鍵名可用裝置上的 `grep`/`sed` 只印 `name="..."`。

Magisk：`su 0 ls -la path` 會報 `invalid option -- a`。應 `su -c 'ls -la path'`。`run-as` 在非 debuggable 包會失敗，不代表沒有 root 路徑。

#### Trigger

- 要記錄 token **存放位置**，但禁止讀值。
- `su 0 ls -la` 失敗；`run-as` 說 not debuggable。

#### Evidence

- Tool: Magisk `su -c`；on-device extract of XML `name=` attributes。
- Sanitized excerpt: Unity `*.v2.playerprefs.xml` held login-shaped key names; values were not copied. `su 0 ls -la` parsed flags as su options.
- Evidence path: login bootstrap storage section under `<PROJECT_ROOT>`

#### Generalized Lesson

1. 先 `ls` 檔名；只把 login-shaped **key names** 寫進文件。
2. 抽鍵名的腳本放 `/data/local/tmp`，跑完刪掉；stdout 不得含 element 文字。
3. Root listing 用 `su -c '...'`，不要 `su 0` 加 dash 參數。
4. 鍵名含 install-hash／uid 後綴時，文件裡改成 `<install-hash>`／`<uid-suffix>`。

#### Agent Action

找 session 存放時不要 pull prefs XML。需要鍵名就 on-device 抽 `name=`。

#### Goal / Action / Validation

- Goal: 文件有檔名與鍵名，transcript 無 token。
- Action: `su -c` list + on-device name extract。
- Validation: git diff／chat 不含 prefs 值；只有 key 名。

#### Applies When

- 授權讀 app data；Unity 或 SharedPreferences 可能存 session。

#### Does Not Apply When

- 已有官方 API 說明 token 生命週期、不需要看裝置檔。

#### Validation

文件列出檔名與 redacted key 名；沒有 string 值。

#### Promotion Target

- `workflow/apk-analysis/artifact-gates/self-generation-audits.md` lifecycle 列
- `analysis/apk/tools-and-failures.md` adb/root 注意

#### Required Linked Updates

- 必須：apk-analysis README Recent + common 計數
- 具體 key 字面值（若過於專案）留專案文件並去敏後綴
