> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-08 - Event wrapper extras may be only ctor/Invoke/Clone, not nested ParseAll

Status: validated

#### One-line Summary

Round/result wrapper 在 `Parse` 之外可能只有 **`.ctor` / `Invoke` / `Clone`**；巢狀 `ParseAll` 不必由 wrapper 的 Invoke/Clone 呼叫。

#### Human Explanation

nested `ParseAll` 出現在 wrapper Parse 前／後時，下一步掛 wrapper 其餘非 accessor 方法。若 live 只有 IN `.ctor`，代表 nested fill 的 caller 在別的型別。不要 dump Invoke 參數。

#### Trigger

- Nested `Parse*` 已在 PRE/POST。
- 懷疑 wrapper 還有 Fill/ParseSpin 之類 helper。

#### Evidence

- Tool: Frida enumerate non-get/set methods on wrapper types except `Parse`; PRE/IN/POST/APPLY names only.
- Sanitized observation: extras were `.ctor`, `Invoke`, `Clone`; live IN was both `.ctor` only; PRE/POST/APPLY none.
- Evidence path: `<PROJECT_ROOT>/<App>/docs/` protocol notes

#### Generalized Lesson

1. Wrapper 不一定有第二層 Parse helper。
2. `Invoke`/`Clone` 存在 ≠ Parse 窗口會跑。
3. nested `ParseAll` 的 caller 要到別的 class 找。

#### Agent Action

nested ParseAll 後列 wrapper 非 accessor；IN 只有 ctor 就不要再追 wrapper。

#### Goal / Action / Validation

- Goal: 確認 nested fill 是否由 wrapper helper 呼叫。
- Action: wrapper 方法名 + 時間桶。
- Validation: LIST 含 Invoke/Clone 或沒有；live 可為僅 ctor。

#### Applies When

- 授權 IL2CPP；靜態 wrapper `Parse` + nested DTO `ParseAll`。

#### Does Not Apply When

- wrapper live 確實呼叫 Fill/Parse* helper（照實記錄）。
- 純靜態。

#### Validation

One round; method names only.

#### Promotion Target

- N/A this round（category index only）

#### Required Linked Updates

- `feedback/history/apk-analysis/unity-il2cpp/README.md`
- `feedback/history/apk-analysis/README.md`
