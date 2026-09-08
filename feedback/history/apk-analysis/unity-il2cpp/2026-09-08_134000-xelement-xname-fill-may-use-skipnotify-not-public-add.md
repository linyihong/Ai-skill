> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-08 - XElement(XName) fill may use SkipNotify, not public Add

Status: validated

#### One-line Summary

`XElement(XName)` 看起來沒有內容時，填充常走 `XContainer.AddNodeSkipNotify`／`AppendNodeSkipNotify`／`ValidateNode`，**不是** public `Add`／`set_Value`。

#### Human Explanation

只鉤 `Add`、`set_Value`、`SetAttributeValue` 會得到 `mut=none`，誤判 Parse 參數是空名稱殼。LINQ to XML 的內部路徑用 SkipNotify。子節點類別名與 Load 樹 EQ 分開記；不要 dump `XName` 或 XML。

#### Trigger

- Ctor 是 `XElement(XName)`，Parse 參數 identity 命中 ctor `this`。
- Public `Add` 在 Decrypt 窗口沒有打到該物件。

#### Evidence

- Tool: Frida; ctor `this` pool then fill methods; class names + EQ/NE only.
- Sanitized observation: child class was `XElement`; Load Root/Element identity stayed negative; public `Add` silent.
- Evidence path: `<PROJECT_ROOT>/<App>/docs/` protocol notes

#### Generalized Lesson

1. Public `Add` 靜默 ≠ 沒有填充。
2. 優先鉤 `AddNodeSkipNotify`、`AppendNodeSkipNotify`、`ValidateNode`、`AddAttributeSkipNotify`。
3. 子節點不必 EQ Load 樹。
4. 不要 dump `XName`／XML。

#### Agent Action

`XElement(XName)` + empty public mutators 時改鉤 SkipNotify／ValidateNode，再比 Load 樹 EQ。

#### Goal / Action / Validation

- Goal: 避免把「沒打到 Add」寫成空元素 hop。
- Action: SkipNotify pool on ctor `this` before Parse.
- Validation: fill methods hit; public Add remains silent or unhooked.

#### Applies When

- 授權 IL2CPP；LINQ to XML 在 Load 之後再建 Parse 用的 `XElement`。

#### Does Not Apply When

- Public `Add`／`XElement(XName, object)` 已是 live hop。
- 純靜態。

#### Validation

Same Decrypt window: SkipNotify/ValidateNode on Parse-bound ctor objects; Add silent.

#### Promotion Target

- N/A this round（category index only）

#### Required Linked Updates

- `feedback/history/apk-analysis/unity-il2cpp/README.md`
- `feedback/history/apk-analysis/README.md`
