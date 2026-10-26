> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-26 - 勿 Interceptor.attach 稍後要用 NativeFunction／runtime_invoke 呼叫的同一方法

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。來源：lab 對 Drop／通知 `Show` 同時 attach 與 game-thread `NativeFunction`／`runtime_invoke` 時觸發 Frida「breakpoint」重入，呼叫鏈中斷。改為只 observe 下游事件（如 HandKnock）或僅 invoke、不 attach Show。

#### One-line Summary

若 lab 需要 **主動呼叫** 某 managed 方法（`NativeFunction`／`il2cpp_runtime_invoke`），**不要**對同一 impl 再 `Interceptor.attach`；attach 會在重入時觸發 Frida breakpoint，導致 force Show／Complete 失敗。改 hook 下游 AnimationEvent／Setup，或只 invoke 不 attach。

#### Human Explanation

Frida 在已 patch 的入口上再從腳本同步呼叫同一地址，會撞上自己的 interceptor。腳本載入時就 attach（早於 fixture 旗標）尤其危險——即使後來想「只觀察」，impl 已被 patch。常見於 Drop／Respin `Show`：需要 sync invoke 才能在正確執行緒播完，卻又想 log `Show` 進出。

#### Trigger

- `NativeFunction`／`runtime_invoke` 回傳異常或 log「breakpoint」
- force Show 計數為 0，但 Setup 已成功
- 同一 MethodInfo 既 attach 又被 lab 主動呼叫

#### Evidence

- Tool: Frida apply-shim；game-thread invoke vs Interceptor.attach
- Sanitized excerpt: attach Show → invoke Show 失敗／breakpoint；移除 attach 後 invoke＋下游 HandKnock 成功
- Evidence path: `<PROJECT_ROOT>/docs/slots/*/lab/` hand-knock gates notes

#### Generalized Lesson

1. **互斥**：主動 invoke 的方法 ⊆ 不 attach；要 log 則 hook **下游**或 **Setup**（若 Setup 不被自己呼叫）。
2. **載入期 attach**：腳本 load 時的無條件 attach 不受後續 fixture「不要 force」控制——對可能被 invoke 的 API 預設 false。
3. **IL2CPP 位址**：attach／invoke 都用 `MethodInfo.readPointer()`（或等價 impl），不要把 MethodInfo 指標當 code。
4. **觀察自然 Show**：靠自然圖排程或只 attach 不 invoke；force 路徑則只 invoke。

#### Agent Action

- 設計 force Show 前列出「會被腳本呼叫的方法」→ 禁止 attach。
- 手垂／AnimationEvent：hook `Invoke*Event`，勿 attach 被 invoke 的 `Show`。

#### Goal / Action / Validation

- Goal: lab force 與 observe 不互相踩斷
- Action: invoke 與 attach 分集合；驗證下游事件
- Validation or reference source: 無 breakpoint；force 路徑 frameShow／HandKnock≥1

#### Applies When

- Frida＋IL2CPP lab 需 sync invoke Unity／feature UI 方法

#### Does Not Apply When

- 純觀察、從不從腳本呼叫該方法
- 非 Frida interceptor 模型

#### Validation

- 同裝置：attach+invoke → 失敗；僅 invoke → 成功＋下游事件

#### Promotion Target

- [Frida hook flow](../../../../analysis/apk/workflows/frida-hook-flow.md)（若有 lab invoke 節）；否則 [unity-fx-driver-trace.md](../../../../analysis/apk/workflows/unity-fx-driver-trace.md) §2

#### Promotion Record

尚未 promotion；保留為 candidate history。

#### Required Linked Updates

- category README 追加本條
- `unity-fx-driver-trace.md` §2 補「勿 attach 將 invoke 的方法」一句
- 專案 evidence：`<PROJECT_ROOT>/docs/slots/*/lab/`

#### Transfer Contract

- Inputs: 需 lab invoke 的 MethodInfo、可選 observe
- Mechanism: interceptor 與腳本重入衝突
- Decision: invoke 集合與 attach 集合互斥
- Verification: 無 breakpoint；下游事件到達
- Counterexample: 純觀察路徑
- Transfer case: 任何 Frida sync invoke＋同函數 hook

#### Intelligence Review

- Decision: method-only
- Reason: Frida／lab 操作邊界，併入 FX／hook 流程說明
- Target: analysis/apk/workflows/unity-fx-driver-trace.md
- Revisit: none
- Discovery: not applicable — method-only

#### Abstraction Check

- Swap-project test: PASS — 任何 Frida 對同一 impl 又 attach 又 NativeFunction／runtime_invoke 的情境均成立
- Project evidence kept in: `<PROJECT_ROOT>/docs/slots/*/lab/`
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS
