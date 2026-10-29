> 遵守 [共用規則索引](../../../../enforcement/README.md) 與 [回饋規則](../../../feedback-lessons.md)；專案原始證據留在來源專案。

### 2026-10-23 - 分析方法與專案報告實作分開維護

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。本次文字整理與離線檢查不構成第二個專案的實測。

#### One-line Summary

共用庫保存證據取得與驗證方法；專案的報告生成器、版型、manifest 與命令契約留在專案。

#### Human Explanation

把某個報告生成器的名稱、外觀或章節號寫成跨專案必跑規則，會讓沒有該工具的專案無法套用方法。工具存在與否也不應改變證據品質要求。

#### Trigger

- 回饋正文要求執行某專案的固定腳本或固定報告版型。
- 只有一個專案的修正成功，卻被標成跨專案 validated。

#### Evidence

- Tool: 來源專案既有分析／驗證紀錄；本輪未重跑其執行環境。
- Sanitized excerpt: 另一個使用不同報告格式或沒有生成器的 Unity 專案，仍能依相同證據門檻工作。
- Evidence path: `<PROJECT_ROOT>/` 的受控分析與驗證紀錄；不複製原始物件、樣本或部署紀錄。

#### Generalized Lesson

1. 先寫方法的輸入、輸出、證據來源與品質門檻。
2. 在來源專案選擇或實作符合當地契約的報告工具；不把具體腳本 API 複製成共用義務。
3. 只有需要跨專案相容性的欄位才另建通用 schema；不能因來源專案存在 manifest 就推定全庫必填。
4. 分開記錄「已抽象／已接入口」與「已獨立重用驗證」，前者不自動提升 lifecycle。

#### Agent Action

先辨識輸入契約與前置條件，再依本條決策；不照搬來源專案的函式名、腳本路徑、固定數字或狀態碼。記錄採用／不適用的理由。

#### Goal / Action / Validation

- Goal: 共用庫保存證據取得與驗證方法；專案的報告生成器、版型、manifest 與命令契約留在專案。
- Action: 依 Generalized Lesson 記錄輸入、決策與反例。
- Validation or reference source: 拿掉所有專案腳本名、章節號與版型描述後，另一個專案仍能說明如何取得與驗證同類證據。

#### Applies When

- 從專案分析／報告工具提取共用方法。

#### Does Not Apply When

- 純專案工具操作手冊；此時直接留在專案文件。

#### Validation

拿掉所有專案腳本名、章節號與版型描述後，另一個專案仍能說明如何取得與驗證同類證據。

#### Promotion Target

[analysis/unity/workflows/unity-asset-parse-flow.md](../../../../analysis/unity/workflows/unity-asset-parse-flow.md)。

#### Promotion Record

2026-10-26：抽象化整理並接到上述方法／入口；這是候選指引的連結，不宣告 validated／promoted，也不新增獨立重用證據。

#### Required Linked Updates

- 本分類 README 與上述方法／入口已檢查並同步。
- 索引與 RAG 經 runtime refresh 重建；原始專案證據不搬入共用庫。

#### Abstraction Check

- Swap-project test: PASS — 另一個使用不同報告格式或沒有生成器的 Unity 專案，仍能依相同證據門檻工作。
- Project evidence kept in: `<PROJECT_ROOT>/` 的受控驗證紀錄。
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS

#### 修訂 2026-10-26

依近期回饋審查，把原先依賴專案術語或操作細節的敘述改為輸入、機制、決策與驗證；修訂前內容可由 Git 歷史追溯。Status 保守維持 candidate。

#### Transfer Contract

- Inputs: 方法、報告實作、可攜性推演與獨立重用證據
- Mechanism: 一次修正成功或路由已接線不會增加實測覆蓋深度
- Decision: 方法與實作分層，狀態依獨立證據決定
- Verification: 拿掉原腳本仍能解釋方法，且不冒充第二次實測
- Counterexample: 純專案工具手冊無需提成通用方法
- Transfer case: 使用不同報告工具的 Unity 專案

#### Intelligence Review

- Decision: existing
- Reason: 方法接線與證據成熟度不相同，可由既有 evidence depth 補足判斷。
- Target: intelligence/engineering/execution/validation-reasoning/evidence-depth.md
