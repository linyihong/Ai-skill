> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-08-03 - Windows 上的 .cmd shim 要靠 shell 啟動，不是補 .cmd 副檔名

Status: validated

#### One-line Summary

Node 在 Windows spawn 套件管理器的 wrapper 指令時，`shell: false` 找不到 shim（ENOENT），改指名 `<cmd>.cmd` 也會被拒（EINVAL）；唯一可行是走 shell，且要傳單一字串以免觸發 args-array 的 deprecation。

#### Human Explanation

跨平台建置腳本常用 `spawn`／`spawnSync` 呼叫套件管理器的執行器（如 Node 生態的 `npx`、`npm`、`yarn`）。在 POSIX 上這些是實體可執行檔，在 Windows 上卻是 `.cmd` batch shim。

多數人記得的修法是「依 `process.platform` 補上 `.cmd` 副檔名」。這個修法**在現代 Node 上已經失效**：自 CVE-2024-27980 的緩解措施（Node 18.20.2 / 20.12.2 / 21.7.3 起）之後，Node 拒絕在沒有 shell 的情況下直接 spawn `.bat` / `.cmd`，會回 `EINVAL`。因此舊部落格與舊 issue 上的 `.cmd` 解法會讓人以為修好了，實際只是把錯誤碼從 ENOENT 換成 EINVAL。

真正的修法是在 Windows 上走 shell。但走 shell 又有第二個坑：同時傳 `shell: true` 與 args array 會觸發 DEP0190 deprecation warning（因為參數只是被串接、不會被跳脫）。若腳本在迴圈裡跑 N 次，這個 warning 就印 N 次，把真正的輸出淹掉。傳入**單一預先引號化的命令字串**可同時避開 warning 與跳脫語意不明的問題。

第三個坑是診斷。spawn 失敗時 `status` 是 `null` 而非數字，`process.exit(status ?? 1)` 這種寫法會讓腳本**安靜地**以 1 結束、完全不印任何東西，看起來像卡住或無故失敗。真正的原因藏在沒人讀的 `result.error` 裡。

#### Trigger

- 跨平台建置／codegen 腳本在 Windows 上以 exit 1 結束，但沒有任何輸出
- 同一腳本在 POSIX 正常，使用者被迫手動逐一執行底層指令當作 workaround
- `spawnSync` 的結果 `status` 為 `null`
- 修成 `<cmd>.cmd` 之後仍然失敗，錯誤碼從 ENOENT 變成 EINVAL

#### Evidence

- Tool: Node `child_process.spawnSync` 探測腳本（Windows，Node 24.x）
- Sanitized excerpt:
  - `shell: false` + wrapper 名稱 → `status: null`、`error: ENOENT`
  - `shell: false` + `<wrapper>.cmd` → `status: null`、`error: EINVAL`
  - `shell: true` + 單一命令字串 → `status: 0`，且無 deprecation warning
  - `shell: true` + args array → `status: 0`，但每次呼叫印一則 DEP0190
- Evidence path: 具體專案的腳本路徑、設定檔清單與 live run 輸出留在該專案 repo 與其 commit message，不複製到本檔。

#### Generalized Lesson

1. **平台差異在啟動機制，不在檔名。** Windows 的 batch shim 需要 command interpreter 才能執行；補副檔名不會讓它變成可直接 spawn 的映像檔。判斷式應是「這個目標是不是 shim」，而不是「要不要加 `.cmd`」。
2. **shell 只開在需要的那一邊。** 只有 shim 需要 shell；同一腳本裡呼叫真實可執行檔（語言 runtime、版本控制工具）的地方應維持 `shell: false`——走 shell 反而會讓含空白的路徑參數被錯誤切分。不要為了一致性把整份腳本都改成 shell。
3. **走 shell 時傳單一字串並自行引號化**，不要傳 args array，避免 deprecation warning 與未定義的跳脫行為。
4. **spawn 失敗與非零結束是兩種不同狀況，要分開處理。** `status === null` 或 `error` 存在代表根本沒啟動成功；把 `error.message`（或 signal）印出來。`status ?? <fallback>` 這種寫法會把啟動失敗偽裝成一般的非零結束，是「安靜失敗」的來源。
5. **迴圈執行多個設定時，非零結束要指名是哪一個**失敗，否則使用者要自己比對輸出順序。

#### Agent Action

看到「腳本無輸出就 exit 1」時：

- 先確認 `status` 是不是 `null`，並印出 `error`，不要先猜是下游工具的問題。
- 不要採用「依平台補 `.cmd`」這個網路上最常見的答案，先用最小腳本實測三種組合（bare name / `.cmd` / shell）再決定；Node 的行為已因安全性修補改變過。
- 修完後同時補上失敗診斷，不要只讓 happy path 通過——原始 bug 的實際傷害是**沒有訊息**，不只是沒有執行。
- 驗證要包含刻意製造的失敗路徑（例如把目標指令移出 PATH），確認新的診斷真的會印出來。

#### Goal / Action / Validation

- Goal: 讓跨平台建置腳本在 Windows 可執行，且失敗時可診斷。
- Action: 平台分支只包住 shim 呼叫；shell 分支傳單一引號化字串；補 spawn-failure 與非零結束的訊息。
- Validation: 在 Windows 實跑成功路徑（exit 0）並比對產物與手動 workaround 一致；另以移除 PATH 目標的方式跑失敗路徑，確認兩種失敗都印出指令與原因。

#### Applies When

- 以 Node `child_process` 呼叫套件管理器 wrapper、CLI shim 或任何 batch/`.cmd` 進入點
- 腳本需同時支援 Windows 與 POSIX
- Node 版本 ≥ 18.20.2 / 20.12.2 / 21.7.3（含之後所有版本）

#### Does Not Apply When

- 目標是真實可執行映像檔（runtime、編譯器、版本控制工具）——這些不需要 shell，加了反而增加引號化風險
- 僅在 POSIX 執行的腳本
- 執行環境的 Node 早於上述修補版本（此時 `.cmd` 寫法仍可運作，但不應作為新程式碼的預設）

#### Validation

- 三種 spawn 組合的實測結果可重跑重現（bare / `.cmd` / shell）
- 成功路徑的產物與既有已提交結果零差異，證明修正沒有改變輸出
- 失敗路徑實測會印出指令與原因，而非空白 exit 1

#### Promotion Target

- `workflow/software-delivery/execution-flow.md`（跨平台建置腳本的驗收項目：失敗路徑必須有診斷）
- `intelligence/`（若日後累積更多「安全性修補改變 runtime 行為，舊解法靜默失效」的案例，可抽為更高階 atom）

#### Required Linked Updates

- 已依 [`linked-updates.md`](../../../enforcement/linked-updates.md) 更新 [`feedback/history/development-guidance/README.md`](../README.md) 的分類數量與 Recent 表格。
- 本條為單一 lesson，尚未成熟到改寫主流程文件，故未修改 `workflow/software-delivery/` 正文；Promotion Target 已標示未來位置。
- 本條源自 project incident，已依 [`reusable-guidance-boundary.md`](../../../enforcement/reusable-guidance-boundary.md) 檢查：正文只保留泛化的平台行為、錯誤碼語意與驗證方法；專案名稱、分支、腳本路徑、設定檔清單與 live run 輸出留在該專案 repo。
