> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-04 - 歸因給共用基線之前，先排除本地衍生狀態

Status: validated

#### One-line Summary

「原始碼輸入相同」加上「可穩定重現」不足以證明共用基線壞掉；這兩件事在本地衍生狀態過期時同樣成立，而衍生狀態不會出現在原始碼 diff 裡。

#### Human Explanation

要把一個失敗歸因給別人剛合入的變更，直覺的檢查是兩步：確認自己的變更沒碰到相關檔案，再確認失敗可以重現。兩步都通過就很容易下結論說基線壞了。

但這兩步都無法區分「基線真的壞了」與「本地衍生狀態過期」：

- **原始碼輸入相同**——用 diff 比較的是受版本控制的檔案。鎖定檔、解析快取、產生的中繼檔、建置輸出、索引資料庫都不在裡面，而它們會直接影響工具行為。
- **可穩定重現**——過期的衍生狀態是持久的，不是隨機的。它每次都會失敗，看起來就跟真正的缺陷一樣。

這個誤判的代價不對稱：它會把一個十秒就能修好的本地問題，變成一份指名道姓、說某人的 commit 讓主線編譯不過的報告。撤回這種說法比多花兩分鐘查證貴得多。

#### Trigger

準備把某個失敗歸因給他人變更或共用分支狀態，尤其在剛完成 pull / rebase / 切換分支之後。

#### Evidence

- Tool: 版本控制的 diff 與時間戳；建置系統的診斷 log
- Sanitized excerpt: 一次 gate 失敗經過兩項檢查——本輪變更未觸及任何相關原始碼、且單獨重跑可重現——因而被歸因為共用分支上他人的變更。實際原因是一份日期早於該變更的本地解析鎖定檔；重跑解析後同一棵樹通過。原歸因不成立。
- Evidence path: 具體分支、commit 與人員資訊留在 `<PROJECT_ROOT>` 的 incident 紀錄；此處只保留一般化規則。

#### Generalized Lesson

**在把失敗歸因給共用基線之前，明確列出可能影響該工具的衍生狀態，並逐項確認它們與工作樹同步。**

有用的分辨方式：問「這個工具讀了哪些不在版本控制內的東西？」典型有相依解析鎖定檔、套件與編譯快取、產生的程式碼、建置輸出、索引或 attestation 資料庫。任何一項過期都能同時滿足「原始碼相同」與「可重現」。

在真正驗證之前，把歸因表述成假設而不是結論。

#### Agent Action

**應該做：**

1. 歸因給他人變更前，先重新產生相關衍生狀態（重跑解析、清掉快取、重建索引），然後再試一次。這通常比繼續讀原始碼便宜。
2. 檢查衍生狀態產物的時間戳。早於被懷疑變更的時間戳，就是強烈訊號。
3. 若已對外表達過歸因，在推翻後**明確更正**，並指出原歸因錯在哪裡——不要只是換個說法帶過。

**不應該做：**

- 不要把「我的 diff 只有 X 目錄」當成編譯輸入相同的證明；受版本控制的 diff 看不見衍生狀態。
- 不要把「可穩定重現」當成缺陷為真的證明；持久性的過期狀態同樣穩定。
- 不要為了讓症狀消失而修改他人在進行中的工作，直到成因已經確認。

#### Goal / Action / Validation

- Goal: 讓失敗歸因反映實際成因，避免把本地環境問題誤報成他人破壞共用基線。
- Action: 在歸因前重新產生衍生狀態並重試；比對衍生檔的時間戳與被懷疑變更的時間。
- Validation: 重新產生後仍失敗，才支持基線缺陷的結論；若通過，則成因為本地衍生狀態，原歸因作廢並需更正。

#### Applies When

- 工具鏈會讀取不受版本控制的衍生狀態（鎖定檔、快取、產生檔、索引）
- 剛完成 pull / rebase / 分支切換，或長時間未重新產生衍生狀態
- 準備公開陳述某人的變更破壞了共用分支

#### Does Not Apply When

- 失敗發生在乾淨環境（容器／CI 全新 checkout），衍生狀態必然新鮮
- 失敗可在他人未受影響的機器上獨立重現
- 該工具鏈完全不使用受版本控制以外的中間狀態

#### Validation

以本 lesson 的來源事件反向驗證：重新產生衍生狀態後，原本「可重現且原始碼相同」的失敗消失，證明原本兩項檢查不足以支撐該歸因。

#### Promotion Target

- `enforcement/failure-patterns/`（此為 agent 歸因失誤模式，適合提升為 cross-skill failure pattern）
- `enforcement/goal-action-validation.md`（驗證章節：歸因類結論的最低證據標準）

#### Required Linked Updates

- 本 lesson 與 [`2026-09-04_101559-build-gate-must-restore-not-only-build`](2026-09-04_101559-build-gate-must-restore-not-only-build.md) 同源；該條寫的是 gate 設計，本條寫的是歸因紀律，兩條互相連結。
- 依 [`failure-learning-system.md`](../../../../enforcement/failure-learning-system.md) 判斷：本條屬 agent 行為失效模式，已在 Promotion Target 標示 `enforcement/failure-patterns/` 為候選；是否新增該檔待後續確認，本輪只沉澱 lesson。
- 依 [`reusable-guidance-boundary.md`](../../../../enforcement/reusable-guidance-boundary.md) 檢查：人員、分支、commit 等具體證據留在 project docs。
- 已更新 [`feedback/history/development-guidance/README.md`](../README.md) 的 Recent 索引。
