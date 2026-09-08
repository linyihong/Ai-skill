> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-08 - SkipNotify child XElement may be another ctor, not CloneNode or Load

Status: validated

#### One-line Summary

掛到 Parse 用 `XElement` 上的 SkipNotify 子節點，可能是 **另一個 `XElement..ctor`**，不是 `CloneNode` return，也不是 Load 樹節點。

#### Human Explanation

看到 SkipNotify 的 child class 是 `XElement` 時，容易猜 Clone 或從 Load `Element()` 直接掛上去。三池比對（ctor `this`、accessor `x0`、CloneNode `x0`）才能分開。不要 dump `XName` 或 XML。

#### Trigger

- Parse 參數是 ctor wrapper，fill 已是 SkipNotify。
- 子節點與 Load Root/`Element` EQ 失敗。

#### Evidence

- Tool: Frida three pools; class names + EQ/NE only.
- Sanitized observation: child hit ctor pool; Load and CloneNode stayed negative.
- Evidence path: `<PROJECT_ROOT>/<App>/docs/` protocol notes

#### Generalized Lesson

1. SkipNotify child ≠ 自動等於 Load 或 Clone。
2. 用 ctor `this`、CloneNode return register、Load accessor return register 三池。
3. 不要 dump `XName`／XML。

#### Agent Action

SkipNotify 後立刻對 child 做三池 EQ，再寫 hop。

#### Goal / Action / Validation

- Goal: 避免把 nested wrap 寫成 Clone 或 Load graft。
- Action: three-pool EQ on SkipNotify arg1.
- Validation: ctor hit with Load/Clone miss, or the reverse recorded honestly.

#### Applies When

- 授權 IL2CPP；LINQ to XML 在 Load 後再建 Parse 樹。

#### Does Not Apply When

- Child 確實 EQ Load 或 CloneNode。
- 純靜態。

#### Validation

Same Decrypt window: child ctor EQ; Load/Clone NE.

#### Promotion Target

- N/A this round（category index only）

#### Required Linked Updates

- `feedback/history/apk-analysis/unity-il2cpp/README.md`
- `feedback/history/apk-analysis/README.md`
