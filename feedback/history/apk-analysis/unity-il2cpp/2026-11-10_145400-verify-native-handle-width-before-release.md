### 2026-11-10 — 依同版指令核對原生控制代碼寬度

Status: candidate

#### Reuse Evidence

單一目標的原生呼叫與更正後重跑，待獨立重用。

#### One-line Summary

配置成功、清理才失敗時，比對配置回傳值與釋放參數在同版二進位的完整寬度。

#### Human Explanation

以較窄型別保存不透明控制代碼，可能在配置時沒有錯誤，直到釋放才發生存取例外。第一個生命週期成功也不足以確認後續配置。

#### Trigger

原生操作已完成，清理時才出現低位址存取錯誤。

#### Evidence

- Tool: 同版 ARM64 指令與受控重跑。
- Sanitized excerpt: 原生釋放入口用完整 64 位元暫存器做地址遮罩。橋接錯用 32 位元，更正為指標後同一分段清理和下一段通過。
- Evidence path: <PROJECT_ROOT>/docs/slots/<cabinet>/lab/。

#### Generalized Lesson

不透明控制代碼的寬度依目前安裝版本的原生介面裁決。核對配置與釋放兩端，保留同一完整值，按所有權只釋放一次。保存操作完成與釋放完成兩個讀回，避免清理錯誤遮蔽原始結果。

#### Agent Action

比對同版介面、暫存器寬度與實際配置值，再修正橋接宣告。

#### Goal / Action / Validation

- Goal: 核對原生控制代碼的配置與清理。
- Action: 對照同版原生指令，使用一致寬度和單次所有權移轉。
- Validation or reference source: 重跑原故障分段並驗證清理與下一段。

#### Applies When

可查同版原生入口和清理所有權。

#### Does Not Apply When

原生介面明確宣告並使用狹位元索引，或清理失敗來自重複釋放而未驗證寬度。

#### Validation

記錄配置成功、原生動作完成、釋放完成和下一次配置結果。

#### Promotion Target

- analysis/unity/workflows/il2cpp-runtime-contract-validation.md

#### Promotion Record

維持候選，待獨立重用。

#### Required Linked Updates

已更新既有原生契約流程與分類索引。原始位址與 trace 保留業務專案。

#### Transfer Contract

- Inputs: 同版配置和釋放入口、橋接型別、失敗階段。
- Mechanism: 轉型截斷位址或不透明控制代碼。
- Decision: 依介面寬度與所有權修正配置及釋放。
- Verification: 修正後重跑配置、清理及下一次配置。
- Counterexample: 32 位元 ID 本來就是索引。
- Transfer case: 另一個原生圖形資源配置與清理橋接。

#### Intelligence Review

- Decision: method-only
- Reason: 單一目標，先由既有驗證程序承接。
- Target: analysis/unity/workflows/il2cpp-runtime-contract-validation.md

#### Abstraction Check

- Swap-project test: PASS — 可移用於其他原生配置與清理契約。
- Project evidence kept in: <PROJECT_ROOT>/docs/slots/<cabinet>/lab/
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS
