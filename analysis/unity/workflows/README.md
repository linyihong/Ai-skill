# Unity／IL2CPP procedures（HOW TO COLLECT）

本目錄是 analysis-local procedure。先讀 [`../decision-workflow.md`](../decision-workflow.md) 再進單一流程。

| Workflow | 說明 |
| --- | --- |
| [`unity-asset-parse-flow.md`](unity-asset-parse-flow.md) | UnityCache／AssetBundle；UnityPy／AssetRipper／AssetStudio／Il2CppDumper 分工 |
| [`unity-fx-driver-trace.md`](unity-fx-driver-trace.md) | Animator 命中但視覺缺口；feature FX → Timeline／粒子／Update／Native |
| [`unity-soft-visual-restore.md`](unity-soft-visual-restore.md) | soft／paint API ok 但 GetContent null 或畫面彈回；資料→停驅動→paint |
| [`il2cpp-runtime-contract-validation.md`](il2cpp-runtime-contract-validation.md) | IL2CPP 參數／hook 階段／fixture 形狀契約 |

編排與完成 gate：[`workflow/apk-analysis/execution-flow.md`](../../../workflow/apk-analysis/execution-flow.md)。
工具選擇總表（Unity 段）：[`../../apk/tools-and-failures.md`](../../apk/tools-and-failures.md)。
