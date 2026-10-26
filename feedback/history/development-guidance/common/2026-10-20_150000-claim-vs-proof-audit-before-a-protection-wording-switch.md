> 遵守 [共用規則索引](../../../../enforcement/README.md) 與 [回饋規則](../../../feedback-lessons.md)；專案原始證據留在來源專案。

### 2026-10-20 - 保護性宣稱需逐句對應證據與反例

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。本次文字整理與離線檢查不構成第二個專案的實測。

#### One-line Summary

對外文件使用「無法／拒絕／防止」等保證前，逐句核對執行器、角色能力、驗證範圍與反例。

#### Human Explanation

通過一組測試不能支持超出其角色與操作範圍的保證。只有偵測能力的控制，也不能在缺乏強制排序時被寫成防止能力。

#### Trigger

- 契約或文件準備升級保護性措辭。
- 驗證報告使用總數，卻沒有角色與操作明細。

#### Evidence

- Tool: 來源專案既有分析／驗證紀錄；本輪未重跑其執行環境。
- Sanitized excerpt: 不同儲存、權限或交付系統都能逐句建立宣稱與證據對照，不需沿用原系統資料模型。
- Evidence path: `<PROJECT_ROOT>/` 的受控分析與驗證紀錄；不複製原始物件、樣本或部署紀錄。

#### Generalized Lesson

1. 建立宣稱→執行控制→證據→限制對照。
2. 依威脅模型分別測試各角色可執行的操作；合併統計前去除重複身分與重複情境。
3. 區分偵測、阻擋與復原；缺失證據的宣稱降為設計意圖或明示限制。
4. 在宣稱發布前完成獨立核對，保留可重現反例或尚未覆蓋範圍。

#### Agent Action

先辨識輸入契約與前置條件，再依本條決策；不照搬來源專案的函式名、腳本路徑、固定數字或狀態碼。記錄採用／不適用的理由。

#### Goal / Action / Validation

- Goal: 對外文件使用「無法／拒絕／防止」等保證前，逐句核對執行器、角色能力、驗證範圍與反例。
- Action: 依 Generalized Lesson 記錄輸入、決策與反例。
- Validation or reference source: 每個保護性宣稱都可反查證據與角色範圍；至少包含符合授權的負向或反例情境。

#### Applies When

- 對外或契約文字聲稱某角色不能完成某操作。

#### Does Not Apply When

- 明示為待驗證設計意圖的內部筆記；仍不可冒充已驗證保證。

#### Validation

每個保護性宣稱都可反查證據與角色範圍；至少包含符合授權的負向或反例情境。

#### Promotion Target

[workflow/software-delivery/release/README.md](../../../../workflow/software-delivery/release/README.md)。

#### Promotion Record

2026-10-26：抽象化整理並接到上述方法／入口；這是候選指引的連結，不宣告 validated／promoted，也不新增獨立重用證據。

#### Required Linked Updates

- 本分類 README 與上述方法／入口已檢查並同步。
- 索引與 RAG 經 runtime refresh 重建；原始專案證據不搬入共用庫。

#### Abstraction Check

- Swap-project test: PASS — 不同儲存、權限或交付系統都能逐句建立宣稱與證據對照，不需沿用原系統資料模型。
- Project evidence kept in: `<PROJECT_ROOT>/` 的受控驗證紀錄。
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS

#### 修訂 2026-10-26

依近期回饋審查，把原先依賴專案術語或操作細節的敘述改為輸入、機制、決策與驗證；修訂前內容可由 Git 歷史追溯。Status 保守維持 candidate。

#### Transfer Contract

- Inputs: 準備發布的保護性宣稱、控制與角色能力
- Mechanism: 低層或部分角色測試不能支持更廣的保證
- Decision: 逐句限制宣稱到已覆蓋的操作與角色
- Verification: 宣稱對應證據，且未測範圍被明示
- Counterexample: 僅描述待驗證設計意圖時不宣稱實證保護
- Transfer case: 不同資料保存或權限系統的對外保證

#### Intelligence Review

- Decision: existing
- Reason: 既有 evidence model 已處理 claim 與證據範圍，補充角色反例而不另建重複 atom。
- Target: intelligence/engineering/execution/validation-reasoning/evidence-model.md
