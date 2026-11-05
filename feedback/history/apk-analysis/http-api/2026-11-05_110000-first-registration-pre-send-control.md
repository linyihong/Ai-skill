# 首次裝置建帳的送出前對照

Status: candidate

#### One-line Summary

在乾淨裝置環境攔住官方首次送出的登入，讓獨立客戶端用同一份測試身分輸入先完成建帳，再以受保護的唯讀查詢驗證新會話。

#### Generalized Lesson

重開、清除 App 資料及新的作業系統識別可能對應不同狀態。應分別標記恢復、首跑及識別替換。即使 HTTP 和業務碼都成功，回應可能只有會員資料外框。先確認服務端新建標記、憑證存在，以及後續唯讀查詢成功，才能判定建帳完成。

#### Evidence

同一官方傳輸以合成識別重現缺少會話欄位的回應。使用另一個全新測試環境，攔下官方首包後，以獨立客戶端重算簽章並取得新會話及唯讀查詢結果。原始證據留在 `<PROJECT_ROOT>`。

#### Applies When

已授權分析裝置登入，而且可在隔離環境攔截首次請求並保留乾淨身分。

#### Does Not Apply When

只需驗證舊會話續期，或登入依賴不可轉移的硬體證明、互動授權或受限制的裝置憑證。

#### Validation

記錄 App 請求未送出、兩種簽章比對、服務端新建證據、憑證存在與同會員唯讀查詢。另記錄身分輸入由哪個環境核發。

#### Promotion Target

[首次建帳對照方法](../../../../analysis/apk/workflows/first-registration-pre-send-control.md)。以候選方法接線，待獨立情境重用。

#### Abstraction Check

- Swap-project test: PASS
- Project evidence kept in: `<PROJECT_ROOT>` 的登入分析紀錄
- Product and sample identifiers absent from Generalized Lesson: PASS

#### Transfer Contract

- Inputs: 隔離的首次執行環境、可驗證的簽章實作、可阻止送出的 App 首次請求
- Mechanism: 首次請求一旦抵達服務端可能消耗新建條件，舊會話與合成識別也可能走不同回應路徑
- Decision: 先保留乾淨身分，確認簽章等價，再讓獨立傳輸成為首個送出者
- Verification: 新建標記、憑證與同會員唯讀查詢均成立，且可證明官方請求尚未送出
- Counterexample: 裝置憑證只能在原程序使用時，應保留裝置端驗證邊界
- Transfer case: 另一個使用裝置識別建立訪客工作階段的閱讀 App

#### Intelligence Review

- Decision: method-only
- Reason: 增量是隔離首跑、攔截與對照的可重現操作順序
- Target: analysis/apk/workflows/first-registration-pre-send-control.md

#### Agent Action

按方法文件分開記錄首次建帳與恢復，先比對簽章，再驗證新會話及唯讀查詢。

#### Goal / Action / Validation

- Goal: 可重現的首次建帳證據
- Action: 隔離環境與送出前對照
- Validation: 新建標記、憑證存在與同會員查詢

#### Required Linked Updates

方法文件、父索引及 APK execution-flow 入口已在同一變更接線。
