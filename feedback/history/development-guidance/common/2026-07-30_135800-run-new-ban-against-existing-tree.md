> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md)、[reusable-guidance-boundary](../../../../enforcement/reusable-guidance-boundary.md)、[sanitization](../../../../enforcement/sanitization.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-07-30 — Run a new mechanical gate against the whole tree before wiring it

Status: candidate

#### One-line Summary

新增機械檢查後，先對整棵既有樹跑一次再接進 verification chain：同一 root cause 的兄弟實例會一起浮現，而那會改變這次任務的 scope。

#### Human Explanation

寫機械 gate 時，作者腦中通常只有觸發它的那一個案例。但 gate 的判斷條件通常比那個案例更廣——例如「設定檔的相對 `extends` 必須指向存在的檔案」這個條件，會同時掃到 build 設定、test 設定與任何其他衍生設定。

若直接把 gate 接進 chain 就交付，兩種結果都不好：gate 在別人的 commit 上爆掉（看起來像新 gate 有 bug），或作者為了讓自己的 commit 過關而把 gate 的判斷條件縮小到「剛好只涵蓋我修的那幾個檔案」——後者等於用一個假的綠燈換掉真實訊號。

正確做法是把「首跑輸出」當成 scope 探測工具：它會告訴你這個 root cause 實際影響多廣。發現兄弟實例時，要嘛一起修（並在回報中說明擴大了 scope），要嘛明確記錄為 pre-existing 並縮小 gate 條件——但縮小必須是有意識的決定，不是為了讓紅燈消失。

另一個常見誤判：機械檢查若用 regex 剝除 JSON 註解，會被字串內的萬用字元或路徑片段誤傷。設定檔類的 gate 應該用逐字元 scanner 處理字串邊界，並在首跑時把「parse 失敗」與「規則違反」分開看。

#### Trigger

- 新增任何 repo-wide 機械檢查（禁用某種檔案、驗證設定欄位、強制某種命名）
- 檢查條件比觸發它的原始案例更一般化
- 準備把檢查接進 pre-commit / verification chain

#### Evidence

- Tool: monorepo 前端 workspace session（已去敏）
- Sanitized excerpt: 為了擋住某類測試檔而新增的機械檢查，順帶加了「相對 extends 必須可解析」條件；首跑對既有樹回報數個 **build** 設定檔有相同的壞 extends，該類 build 指令原本就是壞的但沒人發現。首跑同時暴露檢查自身的 JSON 註解剝除以 regex 實作、被字串內的萬用字元誤傷。
- Evidence path: 受影響檔案清單與指令輸出留在 `<PROJECT_ROOT>` 的 commit message

#### Generalized Lesson

```
1. 新 gate 寫完後，第一件事是對整棵既有樹跑一次，不要先接進 chain
2. 把首跑輸出分成三類：
   a. 我這次要修的目標
   b. 同一 root cause 的兄弟實例  → 通常應一起修，並回報 scope 擴大
   c. 檢查自身的 bug（parse 失敗、誤判） → 先修檢查
3. 絕不為了讓紅燈消失而縮小判斷條件；縮小必須有獨立理由
4. 接進 chain 前補一次負向測試：刻意違反一次，確認 exit code 非零
5. 解析設定檔時，註解剝除要處理字串邊界，不要用 regex 掃整份檔案
```

#### Agent Action

新增機械檢查時：

1. 先獨立執行檢查（不接 chain），對整個 repo 跑
2. 逐條分類首跑輸出，不要一律當成「待修」或一律當成「pre-existing」
3. 發現兄弟實例時，在最終回報中明確說明「這是我擴大的 scope 及理由」
4. 修完後做負向測試，證明 gate 真的會擋
5. 明確標示哪些是 pre-existing 且本輪不處理，並說明如何確認它是 pre-existing
   （例如在未修改的樹上重跑一次）

#### Goal / Action / Validation

- Goal: 讓新 gate 上線時是可信的——既不誤擋，也不因遷就現況而失去訊號
- Action: 獨立首跑 → 分類輸出 → 修檢查自身 bug → 決定兄弟實例處置 → 負向測試 → 接 chain
- Validation: 乾淨樹 exit 0；刻意違反 exit 非零；pre-existing 項目已在未修改樹上覆核

#### Applies When

- 新增 repo-wide 的機械 gate、lint rule 或 structural validator
- Gate 的判斷條件可一般化到超出原始觸發案例

#### Does Not Apply When

- 檢查範圍被刻意限定在單一新增目錄，且該目錄本輪才建立
- 純粹的單元測試（非跨 repo 掃描）

#### Validation

- 在未修改的既有樹上跑新 gate，輸出即為 scope 探測結果
- 修復後乾淨樹 exit 0
- 刻意重新引入一次違規，gate 回非零

#### Promotion Target

- `intelligence/engineering/` 機械檢查撰寫指引（若同類情境重複出現）
- `enforcement/`（若「新 gate 必須先對既有樹跑一次」值得升為全庫規則）

#### Required Linked Updates

- 關聯 [`2026-07-30_135700-unreachable-test-surface-delete-over-repair`](2026-07-30_135700-unreachable-test-surface-delete-over-repair.md)：該條決定加禁令，本條規範加禁令的執行方式
- 依 [`reusable-guidance-boundary.md`](../../../../enforcement/reusable-guidance-boundary.md) 檢查：本檔只保留 generalized lesson，受影響檔案清單留在專案文件
