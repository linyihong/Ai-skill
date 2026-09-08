> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-08 - Parse XElement args may navigate via XContainer.Element

Status: validated

#### One-line Summary

`Parse(..., XElement, XElement)` 在參數物件上可能走 **`XContainer.Element` / `Elements` / `GetElements`**，而不是對那些參數呼叫 `XAttribute.get_Value`。

#### Human Explanation

填 wrapper 時可能已呼叫 `Attribute`（甚至與 Load 樹共用 `XAttribute`），但 **Parse 本體**仍可能只做子節點查找。把 getter 掛在「Parse 前的 wrapper pool」會得到 none；應以 **Parse 當下的 args[n]** 當 `this` 再記方法名。不要 dump `XName` 或內容。

#### Trigger

- Parse 參數是新 `XElement` ctor。
- `get_Value` / `get_Name` on wrapper `XAttribute` 在 Parse 進出都是 none。

#### Evidence

- Tool: Frida hooks on `XObject`/`XNode`/`XContainer`/`XElement`/`XAttribute` instance methods; `this` must EQ Parse args; method names + counts only.
- Sanitized observation: one-arg Parse used `XContainer.Element`; two-arg Parse used `Element` plus `Elements`/`GetElements`. Attribute getters on those args were not observed.
- Evidence path: `<PROJECT_ROOT>/<App>/docs/` protocol notes

#### Generalized Lesson

1. Wrapper 填充 API ≠ Parse 讀取 API。
2. 用 Parse args 當 identity，不要用可能被滑出窗口的 ctor pool。
3. 子節點 hop 優先 hook `XContainer.Element`/`Elements` 的回傳，再比 Load 樹。
4. 不要 dump `XName`／XML。

#### Agent Action

Parse 進出記錄參數上的 LINQ 方法名；`get_Value` none 時改看 `Element`/`Elements`。

#### Goal / Action / Validation

- Goal: 找出 Parse 實際導航 API，避免誤記成屬性 getter。
- Action: Parse-arg identity + 方法名計數。
- Validation: 至少出現 `XContainer.Element`；屬性 getter 可為 none。

#### Applies When

- 授權 IL2CPP；XML Load 後 `Parse(XElement, …)`。

#### Does Not Apply When

- Parse 確實呼叫 `get_Value`／`Attribute`（那就記屬性 hop）。
- 純靜態。

#### Validation

Same Decrypt window; Parse onLeave method-name set.

#### Promotion Target

- N/A this round（category index only）

#### Required Linked Updates

- `feedback/history/apk-analysis/unity-il2cpp/README.md`
- `feedback/history/apk-analysis/README.md`
