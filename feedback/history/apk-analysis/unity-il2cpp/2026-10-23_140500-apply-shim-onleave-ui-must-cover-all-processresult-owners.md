> 遵守 [共用規則索引](../../../../enforcement/README.md) 與 [回饋規則](../../../feedback-lessons.md)；專案原始證據留在來源專案。

### 2026-10-23 - 共用 hook 的前後階段需明確覆蓋各目標

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。本次文字整理與離線檢查不構成第二個專案的實測。

#### One-line Summary

多目標共用 hook 必須分別驗證進入、返回與後續處理的適用範圍，避免只有資料更新、沒有預期後續行為。

#### Human Explanation

擴充共用 hook 時，入口分派可能已涵蓋新目標，但返回後的 UI 或清理路徑仍只允許舊目標。這是生命週期分派缺口，不能只依入口命中或資料寫入成功判斷完成。

#### Trigger

- 新目標有入口／寫入事件，卻沒有預期的返回後事件。
- 共用 hook 的不同階段分別使用硬編碼 owner 條件。

#### Evidence

- Tool: 來源專案既有分析／驗證紀錄；本輪未重跑其執行環境。
- Sanitized excerpt: 多個 UI component 或非遊戲資料處理器共用 interceptor 時，也能以目標×階段矩陣找出分派缺口。
- Evidence path: `<PROJECT_ROOT>/` 的受控分析與驗證紀錄；不複製原始物件、樣本或部署紀錄。

#### Generalized Lesson

1. 建立目標×階段矩陣：resolved、enter、leave、後續行為、清理。
2. 各階段以 capability／契約分派，明示哪些目標應執行、哪些刻意排除。
3. 集合不必永遠相等；但每個差異須有理由，不能因複製舊目標條件而靜默漏掉。
4. 先定位首個缺失階段，再追資料或 renderer；不要把缺失 callback 誤判成資料模型錯誤。

#### Agent Action

先辨識輸入契約與前置條件，再依本條決策；不照搬來源專案的函式名、腳本路徑、固定數字或狀態碼。記錄採用／不適用的理由。

#### Goal / Action / Validation

- Goal: 多目標共用 hook 必須分別驗證進入、返回與後續處理的適用範圍，避免只有資料更新、沒有預期後續行為。
- Action: 依 Generalized Lesson 記錄輸入、決策與反例。
- Validation or reference source: 至少對兩種契約情境驗證完整階段與明確排除；有失敗時要回報缺失階段，而非只有總命中數。

#### Applies When

- 多個型別／物件共享 hook，後續動作依目標能力有所不同。

#### Does Not Apply When

- 單一專用目標，或經契約確認某目標不需要返回後動作。

#### Validation

至少對兩種契約情境驗證完整階段與明確排除；有失敗時要回報缺失階段，而非只有總命中數。

#### Promotion Target

[analysis/apk/workflows/il2cpp-runtime-contract-validation.md](../../../../analysis/apk/workflows/il2cpp-runtime-contract-validation.md)。

#### Promotion Record

2026-10-26：抽象化整理並接到上述方法／入口；這是候選指引的連結，不宣告 validated／promoted，也不新增獨立重用證據。

#### Required Linked Updates

- 本分類 README 與上述方法／入口已檢查並同步。
- 索引與 RAG 經 runtime refresh 重建；原始專案證據不搬入共用庫。

#### Abstraction Check

- Swap-project test: PASS — 多個 UI component 或非遊戲資料處理器共用 interceptor 時，也能以目標×階段矩陣找出分派缺口。
- Project evidence kept in: `<PROJECT_ROOT>/` 的受控驗證紀錄。
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS

#### 修訂 2026-10-26

依近期回饋審查，把原先依賴專案術語或操作細節的敘述改為輸入、機制、決策與驗證；修訂前內容可由 Git 歷史追溯。Status 保守維持 candidate。
