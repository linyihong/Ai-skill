> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-08 - Parse XElement args may be ctor copies, not Load tree identity

Status: validated

#### One-line Summary

`XDocument.Load` 樹上的 `get_Root`／`Element` **不必**等於後續 `Parse(..., XElement, XElement)` 的參數；參數可能是 **新 `XElement` 建構出來的物件**（第二個可為 null）。

#### Human Explanation

相關性窗口裡 Root／Element 查找與 round Parse 都可能出現，但 pointer EQ 會失敗。把 `XElement..ctor` 的 `this`（`onLeave` 的 `args[0]`）放進 pool 再比 Parse 參數，才能證明是 **copy／wrap**，不是同一棵樹節點。不要 dump `XName` 或內容。

#### Trigger

- Parse 的 `XElement` 與 last Root／Element retval 都是 `XElement` 但 EQ 失敗。
- 懷疑 enumerator `Current`，但 `get_Current` 掛不到 generated 類別。

#### Evidence

- Tool: Frida pool of ctor `this` vs Parse args; EQ/NE and class names only.
- Sanitized observation: two-arg Parse matched two ctor instances; one-arg-plus-null matched one ctor. Root/Element identity stayed negative.
- Evidence path: `<PROJECT_ROOT>/<App>/docs/` protocol notes

#### Generalized Lesson

1. Load 樹 identity ≠ Parse 參數 identity。
2. ctor 用 **`this`（args[0]）**，不要用 return register。
3. 第二個 `XElement` 可為 null；那不是 hook 失敗。
4. 不要 dump `XName`／XML。

#### Agent Action

Root/Element EQ 失敗時加 `XElement..ctor` pool，再比 Parse。

#### Goal / Action / Validation

- Goal: 避免把「都是 XElement」寫成同一物件 hop。
- Action: ctor `this` pool + Parse EQ。
- Validation: Parse args hit ctor pool; Root/Element pool misses.

#### Applies When

- 授權 IL2CPP；S2C XML Load 後再進 `Parse(XElement, …)`。

#### Does Not Apply When

- Parse 參數確實 EQ Root／Element（那就記樹節點 hop）。
- 純靜態。

#### Validation

Ctor pool hit vs Root/Element miss in the same Decrypt window.

#### Promotion Target

- N/A this round（category index only）

#### Required Linked Updates

- `feedback/history/apk-analysis/unity-il2cpp/README.md`
- `feedback/history/apk-analysis/README.md`
