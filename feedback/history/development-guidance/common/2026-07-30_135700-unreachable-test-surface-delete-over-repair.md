> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md)、[reusable-guidance-boundary](../../../../enforcement/reusable-guidance-boundary.md)、[sanitization](../../../../enforcement/sanitization.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-07-30 — Unreachable test surface: prove rot, then delete over repair

Status: candidate

#### One-line Summary

當某個測試 surface 無法被任何 gate 執行時，先確認 canonical 策略是否早已把測試放在別處；若是，正解是「保留斷言、刪除 surface、加機械禁令」，而不是修好設定讓它跑起來。

#### Human Explanation

「測試存在但不會失敗」比沒有測試更危險：後續作者會在其中加斷言，看不到錯誤，於是相信自己有覆蓋。

收到「這個測試設定壞掉」的回報時，直覺是修設定。但修設定前必須先問一個更前面的問題：**這個 surface 原本就該存在嗎？** 專案常同時存在兩份放置慣例——一份寫在測試策略文件裡（canonical），一份是 scaffolding 工具產生的預設值（殘留）。若 canonical 策略早已把測試放在別處，修好殘留 surface 等於憑空長出第二條 lane，兩條 lane 會各自漂移。

判斷不能靠偏好，要靠證據。決定性的一步是：找出該 surface 中**設定碰巧沒壞**的那一個實例並實際執行它。若它是紅的，就證明這個 surface 已經無人維護、且「修好它」會立刻讓 gate 變紅。

#### Trigger

- 有人回報某類測試「無法執行」「設定指向不存在的檔案」「跑起來就崩」
- 該類測試同時不在任何 verification / pre-commit 指令的涵蓋範圍內
- 專案的測試策略文件已指定另一個 canonical 測試放置位置

#### Evidence

- Tool: monorepo 前端 workspace session（已去敏）
- Sanitized excerpt: 某層級的所有測試 target 因設定 `extends` 指向不存在的檔案而崩潰；同時該層級也不在 verification 指令與 BDD scenario 掃描範圍內。該層級中唯一設定正確的實例實跑後為多數失敗，斷言已與現行契約不符。
- Evidence path: 具體 lib/檔名、失敗數與指令輸出留在 `<PROJECT_ROOT>` 的 plan / commit message

#### Generalized Lesson

```
1. 先查「可達性」再查「正確性」：
   列出所有 gate（CI 指令、pre-commit、lint/coverage scanner）
   實際涵蓋哪些路徑，確認該 surface 是否落在任何一個之內
2. 查 canonical 放置策略：測試策略文件、既有 gate script 的註解與內部邏輯、
   設定檔頂端註解 — 三者一致就是決定性證據
3. 找「碰巧能跑」的那一個實例並實跑，用結果證明 surface 已 rot
4. 選 delete 時，不可靜默丟掉斷言：
   先把仍有效的斷言遷到 gated 位置並看到測試數上升，再刪除舊 surface
5. 同時移除該 surface 的所有復活入口（設定檔、build target、目錄慣例），
   否則 scaffolding 工具會再生一份
```

「修好它讓它能跑」只有在 canonical 策略確實要求該位置存在時才是正解。

#### Agent Action

看到「測試設定壞掉」時：

1. 不要先修設定。先 grep 各 gate script 實際掃描的路徑集合
2. 讀測試策略文件與 gate script 的檔頭註解，判斷 canonical 位置
3. 找出該 surface 中設定未壞的實例並執行，把結果當成 fix-vs-delete 的裁決依據
4. 遷移斷言時，用「測試總數變化」當作沒有靜默掉覆蓋的證據
5. 具體 lib 名稱、失敗數與指令輸出留在專案文件，不寫進 reusable docs

#### Goal / Action / Validation

- Goal: 讓「不會失敗的測試」無法繼續存在，且不因此損失既有覆蓋
- Action: 可達性盤點 → canonical 策略確認 → 實跑證明 rot → 遷移斷言 → 刪除 surface + 機械禁令
- Validation: 遷移後 gated 測試數上升；刪除後既有 gate 全綠；復活測試（重新放回一份）會被機械檢查擋下

#### Applies When

- 專案有多個測試放置慣例，且其中一份來自 scaffolding 預設
- 測試策略文件明確指定 canonical 位置
- 該 surface 可以被完整列舉並一次清乾淨

#### Does Not Apply When

- Canonical 策略確實要求該位置存在（此時修設定並納入 gate 才是正解）
- 該 surface 有大量高價值測試，遷移成本高於維持兩條 lane
- 無法判斷 canonical 位置時 — 應先向維護者確認，不要單方面刪除

#### Validation

- 刪除前後各跑一次 gated 測試套件，數字必須上升或持平，不可下降
- 重新放回一個測試檔，機械檢查必須回非零 exit code

#### Promotion Target

- `intelligence/engineering/` 測試治理指引（若同類情境重複出現）

#### Required Linked Updates

- 關聯 [`2026-07-30_135800-run-new-ban-against-existing-tree`](2026-07-30_135800-run-new-ban-against-existing-tree.md)：本條決定「刪除」，該條處理「加禁令時的必要步驟」
- 依 [`reusable-guidance-boundary.md`](../../../../enforcement/reusable-guidance-boundary.md) 檢查：本檔只保留 generalized lesson，專案 lib 名稱／失敗數／指令輸出留在專案文件
