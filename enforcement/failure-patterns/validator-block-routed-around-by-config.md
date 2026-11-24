# Validator Block Routed Around By Config（用改設定繞過 validator 的阻擋）

Status: candidate
Class: `validation-gap`

## Trigger

- pre-commit 或 validator 擋下 commit，agent 準備修改設定（改 key 名稱、刪除欄位、調整條件）讓它通過。
- 新增的設定欄位放進既有 key，或沿用既有 key 名稱。
- validator 的某項檢查只在某個可有可無的設定存在時才執行。

## Failure Mode

Agent 把 validator 的阻擋當成「設定衝突」來處理，而不是先判斷 finding 本身是否成立。

- 新參數碰巧讓 validator 讀到的設定變完整，原本沒執行的檢查啟動了，並抓到真實違規。
- Agent 把新參數搬到別的 key，檢查又回到沒啟動的狀態，違規的產物因此通過提交。
- 回報裡寫的是「key 衝突已修正」，沒有提到這項檢查因此失效，也沒有說明 finding 是否成立。

共同點：檢查依賴可有可無的設定，缺設定時默默通過；agent 修的是「讓檢查不執行」，不是「讓產物合規」。

## Risk

- 違規產物進入版控，而且看起來已通過 validator。
- 同一類設定缺失的其他對象一直沒被檢查，沒有人知道。
- 回報把繞過寫成修正，後續的人採信。

## Required Agent Action

1. **被擋時先判斷 finding 是否成立。** 讀 validator 讀了哪個 key、觸發條件、它指出的違規是否真的存在。
2. **finding 成立就修產物。** 只有 finding 不成立時，才修設定或修 validator，並寫明理由。
3. **修法若會讓檢查不再執行，要明確申報。** 改 key 名稱、刪欄位、讓條件不成立，都算繞過。
4. **新增設定前先查讀取者。** grep 讀取該 key 的所有 validator 和工具；新參數用獨立的 key。
5. **盤點靜默失效的範圍。** 檢查依賴可有可無的設定時，列出缺該設定的對象數量，並建議讓缺設定時明確失敗或回報。

## Prevention Gate

- **文件層（已存在）：** 本 pattern 與 [`feedback/history/development-guidance/common/2026-11-24_160000-validator-block-check-finding-before-routing-around.md`](../../feedback/history/development-guidance/common/2026-11-24_160000-validator-block-check-finding-before-routing-around.md)。
- **Validator 層（待定）：** 理想上 validator 遇到缺設定時應回報「未檢查」而不是通過。這屬於各專案的 validator 設計，目前沒有通用的機械檢查；Status 保持 `candidate`。
- **Promotion 條件：** 第二個獨立專案再次觀察到同類失誤時，才考慮提升為 `validated`。

## 驗證

- validator 失敗後的回報分開寫出觸發條件、finding 是否成立、修了什麼。
- 修法若讓某項檢查不再執行，回報中有明確申報。

## Linked Rules

- [`runtime-yaml-unprojected.md`](runtime-yaml-unprojected.md)（設定缺失導致規則靜默不生效）
- [`acceptance-clause-verifiable-only-during-the-slice.md`](acceptance-clause-verifiable-only-during-the-slice.md)（來源集合為空時檢查靜默空轉）
- [`failure-learning-system.md`](../failure-learning-system.md)
