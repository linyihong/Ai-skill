# APK 操作方法摘要

| 欄位 | 值 |
| --- | --- |
| Atom ID | `apk.workflows` |
| Source path | [IL2CPP 執行契約驗證](../../analysis/apk/workflows/il2cpp-runtime-contract-validation.md) |
| Lifecycle | candidate |
| Summary | APK 分析依證據選 Frida、proxy 或資產方法。IL2CPP 空回傳先核對 ABI／參數型別；共用 hook 建立目標×階段矩陣；fixture 的展示窗口與完整 runtime 形狀分開，合法 sparse 不任意補值。 |
| When to read | IL2CPP ABI、System.Char、runtime_invoke、hook lifecycle、fixture shape、visible subset、sparse 或 runtime contract 不一致時。 |
| Do not use for | 不複製專案類名、腳本、固定窗口或延遲 re-patch；candidate 方法不代表已跨專案驗證。 |
| Validation signal | 參數正反對照、階段矩陣、完整／子集／sparse 反例有明確結論；未知契約保持 unknown。 |
| Last checked | 2026-10-26 |
