> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-08 - Ctor XName object may not equal Load Element/get_Name XName

Status: validated

#### One-line Summary

`XElement(XName)` 的 `XName` 引數 **不必**與同窗口 Load 側 `Element(XName)`／樹上 `get_Name` 是同一個物件；不要用字串 dump 去「證明」相等。

#### Human Explanation

LINQ to XML 常 intern `XName`，容易假設 pointer EQ 會過。Load 側已有一池 `XName` 時 ctor 仍可能 NE。跳到比較 local name 字串就違規。只記 EQ/NE 與池是否非空。

#### Trigger

- Parse 參數已是 ctor copy，接下來想對上 Load 樹。
- 想確認是不是共用 interned `XName`。

#### Evidence

- Tool: Frida; Load `Element` arg1 + tree `get_Name` x0 pool vs ctor arg1; class names + EQ/NE + pool size only.
- Sanitized observation: pool nonempty; ctor `XName` stayed NE.
- Evidence path: `<PROJECT_ROOT>/<App>/docs/` protocol notes

#### Generalized Lesson

1. `XName` intern 不能當成 pointer hop。
2. 池大小與 EQ/NE 分開記；池空時 NE 無意義。
3. 不要 dump `XName` 或 XML。

#### Agent Action

先建 Load 側 `XName` 池並確認 nonempty，再比 ctor arg；禁止讀 local name 字串。

#### Goal / Action / Validation

- Goal: 避免把「都是 XName」寫成同一物件 hop，也不要用字串比對。
- Action: pointer EQ + nonempty pool.
- Validation: NE with nonempty pool, or EQ recorded honestly.

#### Applies When

- 授權 IL2CPP；Load 後再 `XElement(XName)` 建 Parse 樹。

#### Does Not Apply When

- ctor `XName` 確實 EQ Load 池。
- 純靜態。

#### Validation

Same Decrypt window: load-name pool > 0; Parse ctor XName NE.

#### Promotion Target

- N/A this round（category index only）

#### Required Linked Updates

- `feedback/history/apk-analysis/unity-il2cpp/README.md`
- `feedback/history/apk-analysis/README.md`
