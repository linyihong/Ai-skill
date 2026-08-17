# Deployment-Scoped Mechanism Verified In Wrong Venue（以測試專案代理部署內容做驗證）

Status: candidate
Class: `validation-gap`

## Trigger

某機制的事實來源是**部署單元實際包含什麼**：

- 掃描 `AppContext.BaseDirectory` 的組件／外掛
- 依「哪些套件被引用」決定能力集合
- classpath／載入目錄掃描、DI 組件掃描
- feature flag 由「該模組是否存在」推導

而 agent 要驗證的是**否定命題**：「Host A 看不到 B 專屬的那個 key」「這個部署不含該能力」。

## Failure Mode

Agent 把測試寫在測試專案裡，斷言某項目不存在。測試紅了 —— 但原因不是實作錯，而是**測試專案的輸出目錄不是 Host 的部署目錄**。

測試專案為了建構待測系統，往往引用了比任何單一部署都多的東西（多個 Host 的 factory、共用測試工具、被測模組的兄弟模組）。於是：

- **否定斷言在測試專案裡不可證偽** —— 東西就是在那裡，因為測試專案引用了它
- 更危險的反向：斷言為了通過而被改弱，或實作被改成迎合測試

共同誤解是把「測試專案的輸出」當成「部署內容」的代理。兩者在肯定命題上常一致（部署有的測試專案通常也有），在否定命題上系統性不一致。

## Risk

- 為了讓不可證偽的斷言通過而扭曲實作或弱化斷言，留下比沒測更糟的狀態
- 誤判實作有缺陷，投入時間修一個不存在的問題
- 若斷言被硬改成通過，該機制的真正隔離性從此無人驗證

## Required Agent Action

**驗證要在該機制實際求值的地方進行。**

1. **先問求值發生在哪個目錄／哪個進程。** 部署掃描類機制的求值點是部署輸出，不是測試組件。
2. **否定命題放進「一個部署對應一個測試專案」的套件。** 每個 Host 有自己的測試專案時，該 Host 的否定斷言寫在它自己的套件裡才有意義。
3. **無法在對的場地驗證時，改驗證它的分解命題。** 例如把「catalog 推導」與「服務給定 catalog 的行為」拆開：前者在各 Host 套件驗，後者以明確注入的假 catalog 驗。**並在測試檔案裡寫明這個替代是什麼、為什麼誠實。**
4. **移除不可證偽的斷言時，寫下理由而非讓它悄悄消失。** 移除本身是正確動作，但無紀錄的移除與掩蓋無法區分。

## Prevention Gate

寫下任何「X 不存在／看不到 Y」的斷言前：

- 這個機制求值時讀的是哪個目錄？我的測試進程的那個目錄，內容與它相同嗎？
- 我這個測試專案引用了什麼，是任何單一部署都不會有的？
- 如果實作是對的，這個斷言**有可能**紅嗎？如果不可能紅，它不是測試。
- 如果實作是錯的，這個斷言**一定**會紅嗎？

## 驗證

1. 部署掃描類機制的否定斷言，位於與單一部署一一對應的測試套件中
2. 以假物件替代真實掃描時，測試檔案內有書面說明替代範圍與其正當性
3. 被移除的不可證偽斷言，其移除理由記錄在 slice 證據文件中
4. 真實推導路徑本身仍有覆蓋（不因替代而整段失去驗證）

## Linked Rules

- [`proxy-signal-substituted-for-target-state.md`](proxy-signal-substituted-for-target-state.md) — 上位 pattern：本 pattern 是「場地」維度的代理替換
- [`validation-coverage-gap-executor-placement.md`](validation-coverage-gap-executor-placement.md) — 同型：驗證放錯層級
- [`../failure-learning-system.md`](../failure-learning-system.md) — `validation-gap` 分類
