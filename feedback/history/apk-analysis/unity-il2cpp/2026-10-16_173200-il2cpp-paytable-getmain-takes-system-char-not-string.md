> 遵守 [共用規則索引](../../../../enforcement/README.md) 與 [回饋規則](../../../feedback-lessons.md)；專案原始證據留在來源專案。

### 2026-10-16 - IL2CPP 呼叫結果為空時先核對參數型別

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。本次文字整理與離線檢查不構成第二個專案的實測。

#### One-line Summary

IL2CPP 呼叫結果異常時，先對照 managed 簽名、值型別表示與呼叫入口，再判斷資料是否缺失。

#### Human Explanation

顯示上是一個字元，不代表 API 接受 string。若簽名是 System.Char，把 managed string 指標當值型別參數可能得到無效結果；空結果不能直接證明目標沒有資料。

#### Trigger

- 已確認目標物件存在，但呼叫結果持續為空。
- 顯示值與反編譯的參數型別不一致。

#### Evidence

- Tool: 來源專案既有分析／驗證紀錄；本輪未重跑其執行環境。
- Sanitized excerpt: 另一個 IL2CPP 設定查詢或 UI 字元查詢也能使用「核對型別→建立參數→正反對照」的方法。
- Evidence path: `<PROJECT_ROOT>/` 的受控分析與驗證紀錄；不複製原始物件、樣本或部署紀錄。

#### Generalized Lesson

1. 以當前版本 metadata 確認 overload、參數、static／instance 與 by-ref。
2. 區分 runtime_invoke 的參數儲存位址與直接 Native ABI，不能共用同一包裝假設。
3. 若簽名確為 System.Char，按 UTF-16 code unit 準備值；不能把任意 Unicode code point 截斷成 char。
4. 用已知有效輸入與錯型別對照；必要時讀已確認 layout 的集合，但空欄位同樣不代表不存在資料。

#### Agent Action

先辨識輸入契約與前置條件，再依本條決策；不照搬來源專案的函式名、腳本路徑、固定數字或狀態碼。記錄採用／不適用的理由。

#### Goal / Action / Validation

- Goal: IL2CPP 呼叫結果異常時，先對照 managed 簽名、值型別表示與呼叫入口，再判斷資料是否缺失。
- Action: 依 Generalized Lesson 記錄輸入、決策與反例。
- Validation or reference source: 已知有效輸入可取得預期結果，錯型別對照可重現差異；驗證記錄包含簽名與實際呼叫入口。

#### Applies When

- 授權 IL2CPP 分析，且可核對目標方法簽名。

#### Does Not Apply When

- 已證實是 string key 的 API。
- 沒有足夠版本／layout 資訊；此時標 unknown，不直接猜記憶體結構。

#### Validation

已知有效輸入可取得預期結果，錯型別對照可重現差異；驗證記錄包含簽名與實際呼叫入口。

#### Promotion Target

[analysis/apk/workflows/il2cpp-runtime-contract-validation.md](../../../../analysis/apk/workflows/il2cpp-runtime-contract-validation.md)。

#### Promotion Record

2026-10-26：抽象化整理並接到上述方法／入口；這是候選指引的連結，不宣告 validated／promoted，也不新增獨立重用證據。

#### Required Linked Updates

- 本分類 README 與上述方法／入口已檢查並同步。
- 索引與 RAG 經 runtime refresh 重建；原始專案證據不搬入共用庫。

#### Abstraction Check

- Swap-project test: PASS — 另一個 IL2CPP 設定查詢或 UI 字元查詢也能使用「核對型別→建立參數→正反對照」的方法。
- Project evidence kept in: `<PROJECT_ROOT>/` 的受控驗證紀錄。
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS

#### 修訂 2026-10-26

依近期回饋審查，把原先依賴專案術語或操作細節的敘述改為輸入、機制、決策與驗證；修訂前內容可由 Git 歷史追溯。Status 保守維持 candidate。

#### Transfer Contract

- Inputs: 當前方法簽名、已知有效輸入與空回傳觀測
- Mechanism: 型別與表示不一致會掩蓋有效資料
- Decision: 先核對呼叫契約再判定資料缺失
- Verification: 已知有效輸入與離線拒絕錯型別的對照
- Counterexample: 簽名與資料已完整驗證時不需再改參數
- Transfer case: 不同 IL2CPP UI 字元設定查詢

#### Intelligence Review

- Decision: extract
- Reason: 這會改變空回傳時的診斷選路，適合候選判斷智慧。
- Target: intelligence/engineering/analytical-reasoning/heuristics/runtime-evidence-boundaries.md
- Discovery: knowledge/summaries/runtime-evidence-boundaries.md
