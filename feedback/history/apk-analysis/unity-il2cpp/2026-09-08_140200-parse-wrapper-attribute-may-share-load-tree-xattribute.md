> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-08 - Parse wrapper Attribute may share Load-tree XAttribute

Status: validated

#### One-line Summary

`Parse(..., XElement, XElement)` 的參數可以是新 `XElement` 建構物件，但其 `Attribute` 回傳值仍可能與同窗口 `XDocument.Load` 樹節點上的 `Attribute` **同一 `XAttribute` 物件**。

#### Human Explanation

Load 樹的 Root／Element identity 失敗時，不要立刻斷定 apply 路徑與 Load 樹完全無關。LINQ-to-XML 複本常複製元素節點、但 **屬性物件可被共用**。第二個 Parse `XElement` 可為 null；有物件時其 `Attribute` 也不必 EQ 樹。不要 dump `XName` 或屬性值。

#### Trigger

- Parse 參數 EQ `XElement..ctor`，NE Load Root／Element。
- 需要把 apply 路徑接回 Load 樹，又不想讀字串內容。

#### Evidence

- Tool: Frida pools of Load-tree `XElement.Attribute` return-register objects vs Attribute called on Parse-bound ctor wrappers; EQ/NE and class names only.
- Sanitized observation: first Parse wrapper Attribute hit `XAttribute` objects already seen on the Load tree; second wrapper was null or Attribute NE tree; null Attribute returns are not objects.
- Evidence path: `<PROJECT_ROOT>/<App>/docs/` protocol notes

#### Generalized Lesson

1. Element identity 失敗 ≠ 屬性 identity 也失敗。
2. 分開 pool：樹上 `Attribute` 的 `XAttribute` vs Parse wrapper 上 `Attribute` 的回傳。
3. 第二個 `XElement` 可為 null；不要把 null 當 hook 失敗。
4. 不要 dump `XName`／屬性值。

#### Agent Action

Element NE 之後加 `Attribute` 物件 pool，再比 Parse 前 wrapper 上的 `Attribute` 回傳。

#### Goal / Action / Validation

- Goal: 在不 dump XML 的前提下，把 ctor 複本接回 Load 樹。
- Action: 同 Decrypt 窗口比 `XAttribute` pointer EQ。
- Validation: 至少一個 Parse wrapper 的 `Attribute` EQ 樹 pool；另一個可為 null 或 NE。

#### Applies When

- 授權 IL2CPP；S2C XML Load 後再進 `Parse(XElement, …)`。

#### Does Not Apply When

- Parse 參數已經 EQ Root／Element（直接記樹 hop）。
- Attribute 全部 NE 且沒有 null（那就要找別的 fill／copy API）。
- 純靜態。

#### Validation

Same-window Load-tree Attribute pool vs Parse-wrapper Attribute returns; class names + EQ/NE only.

#### Promotion Target

- N/A this round（category index only）

#### Required Linked Updates

- `feedback/history/apk-analysis/unity-il2cpp/README.md`
- `feedback/history/apk-analysis/README.md`
