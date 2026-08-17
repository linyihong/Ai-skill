# Cross-Boundary Agreement Never Mechanically Checked（跨邊界一致性從未被機械驗證）

Status: candidate
Class: `validation-gap`

## Trigger

同一項事實被兩處各自宣告，而兩處分屬不同語言、不同專案、不同團隊節奏：

- 後端的權限／角色矩陣 ↔ 前端的同一份矩陣
- API 的錯誤碼 ↔ 前端的錯誤訊息目錄
- 契約文件的表格 ↔ 實作的常數表
- 產生器的輸入清單 ↔ 消費端的可選項

兩邊**各自都是正確的**。沒有任何編譯器、型別系統或測試看得見它們之間的關係。

## Failure Mode

兩處靠「意圖」一致，而意圖不會在漂移時報錯。於是漂移不是被發現的，是被**遇到**的 —— 通常在某個機制轉換、某個環境切換，或某位使用者剛好走到那條路徑時。

實際觀察到的形態：

| 形態 | 症狀 |
| --- | --- |
| 前端矩陣少一個 key | 「套用建議」給出的組合與產品文件不符，無人察覺 |
| 前端有後端不存在的 key | 儲存被**整筆**拒絕，連同一次勾選的其他項目一起失效 |
| 兩邊角色層級不同 | 該擋沒擋，或該給沒給；資料庫層看不出異常 |
| 後端可授予但前端無勾選框 | API 能給一個管理員永遠找不到的權限 |
| 錯誤碼被收斂成單一碼 | 具體性遺失；若該碼在訊息目錄中不存在，畫面直接印出 i18n key |

共同點：**沒有任何一側「壞了」**，壞的是兩側之間那條沒人檢查的關係。

## Risk

- 漂移長期存在且累積，發現時已無法判斷哪一側才是原意
- 症狀出現的位置遠離原因（前端畫面 vs 後端表），因此被歸類為 UI 缺陷或環境問題
- 若該關係是產品功能的核心（例如權限設定），功能會在無人宣告它壞掉的情況下失去意義

## Required Agent Action

**跨邊界的一致性必須有一個會失敗的檢查，而不是一段說明它們應該一致的文字。**

1. **先問「誰是權威」，並寫進檢查裡。** 兩邊都可能先開發完成；權威方由決策指定（常見是後端／契約），檢查的錯誤訊息應直接說出哪一側該改。
2. **雙向都要檢查。** 單向檢查會製造已驗證的錯覺：「後端每個 key 都有前端對應」通過，不代表前端沒有多出後端不認的 key —— 而後者往往是更嚴重的那個方向。
3. **檢查放在能同時看到兩側的層級。** 跨語言時單一語言的單元測試看不到對面，落點通常是 repo 層的 gate（`scripts/` 類）而非測試專案。
4. **把一次性比對固化為 gate。** 為釐清問題寫的臨時腳本，其價值在於之後每次變更都跑；留在對話裡等於沒有。
5. **平行實作要逐一檢查。** 多站台／多 Host／多 app 各有一份時，通過其中一個不代表其餘通過（見 [`proxy-signal-substituted-for-target-state.md`](proxy-signal-substituted-for-target-state.md) 呼叫形式規則第 4 條）。
6. **新 gate 首次執行的失敗清單，先懷疑 gate 自己。** 大量失敗更可能是比對邏輯的缺陷（例如比較了集合的順序而非成員），而非受檢對象全錯。確認 gate 正確後再採信其輸出。

## Prevention Gate

發現兩處宣告同一件事時：

- 目前有任何東西會在它們不一致時失敗嗎？若沒有，這個一致性只是巧合。
- 我要加的檢查是雙向的嗎？哪一個方向的漂移後果更嚴重？
- 權威是哪一側？檢查的訊息說得出來嗎？
- 這份比對我是打算跑一次，還是讓它每次都跑？

## 驗證

1. 存在一個會因不一致而失敗的機械檢查，且已接上 pre-commit／pre-push 或 CI
2. 該檢查對每一類不一致都經 mutation 驗證（各自製造一次，確認會紅且訊息指名對象）
3. 檢查涵蓋雙向，且權威方在訊息中明示
4. 刻意允許的例外（單側獨有、機器面專用）在檢查中**具名記錄**並附理由，而非以靜默排除實現

## Linked Rules

- [`proxy-signal-substituted-for-target-state.md`](proxy-signal-substituted-for-target-state.md) — 上位家族：以間接訊號替代目標狀態
- [`decision-revised-without-contract-authority-update.md`](decision-revised-without-contract-authority-update.md) — 同型：單一事實多落點，讀者不同
- [`deployment-scoped-mechanism-verified-in-wrong-venue.md`](deployment-scoped-mechanism-verified-in-wrong-venue.md) — 檢查放錯層級的具體形態
- [`../failure-learning-system.md`](../failure-learning-system.md) — `validation-gap` 分類
