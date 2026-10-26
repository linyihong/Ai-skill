> 遵守 [共用規則索引](../../../../enforcement/README.md) 與 [回饋規則](../../../feedback-lessons.md)；專案原始證據留在來源專案。

### 2026-10-23 - 回放資料必須符合消費端的完整結構契約

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。本次文字整理與離線檢查不構成第二個專案的實測。

#### One-line Summary

展示用 fixture 不一定是完整執行資料；寫入前必須對齊消費端的形狀、座標、缺值語意與生命週期。

#### Human Explanation

同一份可見資料可能分別交給序列化／回放入口與另一個 adapter。消費端若還需要隱藏緩衝列或其他狀態，直接拷貝可見區會產生不完整更新；反過來任意補值，也可能破壞合法 sparse 或佔位編碼。

#### Trigger

- fixture 只描述可見區，但消費端接受較大範圍。
- 回放後狀態停滯、部分格位空白，或不同 adapter 行為不一致。

#### Evidence

- Tool: 來源專案既有分析／驗證紀錄；本輪未重跑其執行環境。
- Sanitized excerpt: 非遊戲的虛擬列表、分塊影像或串流緩衝區，也有展示窗口與完整執行結構不同的契約問題。
- Evidence path: `<PROJECT_ROOT>/` 的受控分析與驗證紀錄；不複製原始物件、樣本或部署紀錄。

#### Generalized Lesson

1. 從正常觀測或權威 schema 取得維度、座標、排序、有效範圍與 sentinel 語意；不能由畫面推測。
2. 對照 fixture 與各 adapter 的輸入／輸出，列出缺失位置與型別。
3. 只在契約明訂時展開／補值／調整容量；合法 sparse 與佔位保持原義。缺失值無合法來源時拒絕寫入或標 blocked，不能 edge-clone 猜值。
4. 更新後追完整狀態轉換與下一次操作；單次寫入成功或一張完整畫面不足以證明一致。
5. 僅在事件證實後續覆寫、且重寫符合生命週期契約時修正排序或重寫；不把固定延遲 re-patch 當通用解。

#### Agent Action

先辨識輸入契約與前置條件，再依本條決策；不照搬來源專案的函式名、腳本路徑、固定數字或狀態碼。記錄採用／不適用的理由。

#### Goal / Action / Validation

- Goal: 展示用 fixture 不一定是完整執行資料；寫入前必須對齊消費端的形狀、座標、缺值語意與生命週期。
- Action: 依 Generalized Lesson 記錄輸入、決策與反例。
- Validation or reference source: 離線覆蓋完整資料、可見子集、合法 sparse 與非法缺值；動態驗證狀態可完成並接受下一次授權操作。這是待驗證標準，不宣稱本輪已重跑裝置。

#### Applies When

- 授權 fixture replay 或 runtime apply，展示模型與執行模型可能不同。

#### Does Not Apply When

- 契約確認兩者完全一致。
- 合法 sparse／佔位資料；不能為了填滿畫面而消除它。

#### Validation

離線覆蓋完整資料、可見子集、合法 sparse 與非法缺值；動態驗證狀態可完成並接受下一次授權操作。這是待驗證標準，不宣稱本輪已重跑裝置。

#### Promotion Target

[analysis/apk/workflows/il2cpp-runtime-contract-validation.md](../../../../analysis/apk/workflows/il2cpp-runtime-contract-validation.md)。

#### Promotion Record

2026-10-26：抽象化整理並接到上述方法／入口；這是候選指引的連結，不宣告 validated／promoted，也不新增獨立重用證據。

#### Required Linked Updates

- 本分類 README 與上述方法／入口已檢查並同步。
- 索引與 RAG 經 runtime refresh 重建；原始專案證據不搬入共用庫。

#### Abstraction Check

- Swap-project test: PASS — 非遊戲的虛擬列表、分塊影像或串流緩衝區，也有展示窗口與完整執行結構不同的契約問題。
- Project evidence kept in: `<PROJECT_ROOT>/` 的受控驗證紀錄。
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS

#### 修訂 2026-10-26

依近期回饋審查，把原先依賴專案術語或操作細節的敘述改為輸入、機制、決策與驗證；修訂前內容可由 Git 歷史追溯。Status 保守維持 candidate。
