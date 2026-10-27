> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-27 - SetActive 只用 NativeFunction(uint8) 打目標 GO；breakpoint→NF 須滿參數個數

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。來源：lab 用 `il2cpp_runtime_invoke` 傳布林／沿 parent 鏈 `SetActive` 導致黑屏／重連；改為對**單一** GameObject `NativeFunction(SetActive, void, [pointer, uint8])`。另：`runtime_invoke` 觸發 Frida breakpoint 後，1 參數啟發式 NativeFunction fallback 在 2 參數 `CreateContent(cell, rect)` 上 AV；改依 `method_get_param_count` 組齊所有 pointer 參數後恢復多數建立。

#### One-line Summary

IL2CPP＋Frida 啟用物件時，**不要**用 `runtime_invoke` 亂塞 bool、更不要沿 parent 逐級 SetActive；只對目標 GO 用 **NativeFunction + uint8**。Breakpoint 後的 NativeFunction fallback **必須**按 MethodInfo 參數個數傳齊（含 valuetype 的指向緩衝區），不可假設「instance + 一個 arg」。

#### Human Explanation

`GameObject.SetActive(bool)` 經 `il2cpp_runtime_invoke` 時，布林裝箱／參數槽錯誤容易寫到錯誤物件或把整棵 UI 樹關掉，表現為黑屏、重連。沿 transform parent 向上「幫你打開」更危險。安全預設：解析目標 GO 指標，直接 `new NativeFunction(impl, 'void', ['pointer', 'uint8'])(go, 1)`，必要時先開根節點再開目標，**禁止** parent-walk invoke。

另一條常見路徑：腳本裡的 `invokeMi` 在 `runtime_invoke` 丟「breakpoint triggered」時退回 NativeFunction。若 fallback 只讀 `args[0]`，對 2＋參數方法（例如 CreateContent(cell, Rect*)）會用錯 ABI → access violation。正確做法：`param_count = method_get_param_count(mi)`，`this` + 每個 param 都以 pointer 從 args 陣列取出再 `NativeFunction`。這與「勿 attach 又 invoke 同一方法」互補：即使不 attach，重入／延遲呼叫仍可能 breakpoint，fallback 仍要 ABI 正確。

#### Trigger

- SetActive／顯示後黑屏、HUD 全沒、需重連
- log 含 `breakpoint triggered`，接著 `access violation` 在 Create／Show
- 同方法有時 createOk、有時 AV，且失敗格剛好是多參數 Creator

#### Evidence

- Tool: Frida；GameObject.SetActive；Creator CreateContent
- Sanitized excerpt: parent-walk invoke SetActive → 黑屏；NF uint8 僅目標 → 恢復。CreateContent breakpoint + 1-arg NF → AV；full-arity NF → 多數格 createOk
- Evidence path: `<PROJECT_ROOT>/docs/slots/*/lab/` enter-board／apply-shim notes

#### Generalized Lesson

1. **SetActive**：僅目標 GO；`NativeFunction(..., uint8)`；禁止 runtime_invoke 布林裝箱與 parent-walk。
2. **Breakpoint fallback**：依 `method_get_param_count` 組 `(this?, a0…aN-1)` 全 pointer（valuetype 傳指向已填緩衝的 pointer）；記錄 `via: native-pcN` 便於對照。
3. **與 attach 互斥**：仍遵守「會被腳本 invoke 的方法不要 Interceptor.attach」。
4. **失敗隔離**：單格 Create AV 應 catch 後 continue，勿讓整次 enter UI 中斷。

#### Agent Action

- 寫 activate helper：只收 GO 指標 + uint8；單元測試／煙測故意錯用 invoke 應能察覺黑屏風險並拒絕合併。
- 維護通用 `invokeMi`：breakpoint／AV 時 full-arity NF；多參數方法煙測至少一條。

#### Goal / Action / Validation

- Goal: 啟用與 Creator 呼叫不毀 UI、不因錯誤 ABI 全滅
- Action: NF uint8 SetActive；full-arity invoke fallback
- Validation or reference source: 無黑屏；CreateContent 多數 ok；via 日誌含 pc≥2

#### Applies When

- Frida＋IL2CPP 操作 Unity GameObject／Creator／多參數 managed 方法

#### Does Not Apply When

- 純 Java／ART hook、無 il2cpp_runtime_invoke
- 已確認不會走 breakpoint fallback 的單參數 void 方法（仍建議統一 helper）

#### Validation

- 錯用 parent-walk → 可復現黑屏；改 NF 目標 → 畫面保留
- 故意 1-arg fallback vs full-arity：後者 Create 成功率更高

#### Promotion Target

- analysis/apk/workflows/frida-hook-flow.md
- analysis/apk/workflows/unity-fx-driver-trace.md

#### Promotion Record

尚未 promotion；保留為 candidate history。

#### Required Linked Updates

- category README 追加本條
- `unity-fx-driver-trace.md` §2.1／§3 補 SetActive uint8 與 full-arity fallback 一句＋連結
- `frida-hook-flow.md` 若有 invoke helper 節則補連結；否則本輪以 fx-driver 為主承接
- 專案 evidence：`<PROJECT_ROOT>/docs/slots/*/lab/`

#### Transfer Contract

- Inputs: GO 指標、MethodInfo、args 陣列、Frida breakpoint／AV
- Mechanism: bool／parent SetActive 毀樹；1-arg NF 與多參 ABI 不符
- Decision: 目標-only uint8 NF；fallback 用 param_count 組參
- Verification: 無黑屏；多參 Create 成功或可隔離失敗
- Counterexample: 非 IL2CPP／從不 SetActive／單參且無 breakpoint
- Transfer case: 任何 Frida 對 Unity GO／Creator 的 sync invoke

#### Intelligence Review

- Decision: method-only
- Reason: Frida／IL2CPP 呼叫慣例增量；承接既有 hook／FX 流程
- Target: analysis/apk/workflows/unity-fx-driver-trace.md
- Revisit: none
- Discovery: not applicable — method-only

#### Abstraction Check

- Swap-project test: PASS — 任何 Frida＋IL2CPP 對 GameObject.SetActive 或多參數 runtime_invoke→NF fallback 的情境均成立
- Project evidence kept in: `<PROJECT_ROOT>/docs/slots/*/lab/`
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS
