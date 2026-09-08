> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-08 - Elements iterator Current may be nested XElement ctor

Status: validated

#### One-line Summary

`Elements`／`GetElements` 的回傳值不必直接被 `get_Current`；Parse 可能先 `GetEnumerator`，再 `MoveNext`／`get_Current`，而 **Current 可以是巢狀 `XElement` ctor**。

#### Human Explanation

只 pool iterator 物件再掛 `get_Current` 會得到 none，因為 `this` 可能是 `GetEnumerator` 之後的同一編譯器 iterator，或介面方法名不是短名 `get_Current`。應掃描類別名含 iterator 工廠的 `GetEnumerator`／`MoveNext`／`get_Current`（含介面映射名）。`MoveNext` 的回傳可能是值型別，return register 看起來不像 managed object。不要 dump 內容。

#### Trigger

- `Elements`／`GetElements` 回傳 iterator，但 Parse 期間短名 `get_Current` 為 none。
- 需要把 foreach 子節點接回 ctor pool。

#### Evidence

- Tool: Frida on compiler iterator methods during Parse; class names + EQ vs ctor pool.
- Sanitized observation: GetEnumerator returned the iterator type; generic `get_Current` returned `XElement` EQ ctor. MoveNext retval was not a managed object.
- Evidence path: `<PROJECT_ROOT>/<App>/docs/` protocol notes

#### Generalized Lesson

1. Iterator retval ≠ 自動呼叫 `get_Current`。
2. 掛 `GetEnumerator` + 介面映射 `get_Current`，不要只掛短名。
3. Current 常 EQ 巢狀 ctor，與 Load 樹無關。
4. 不要 dump XML／`XName`。

#### Agent Action

Parse 期間對 iterator 類別掛 `GetEnumerator`／`MoveNext`／`get_Current`，Current 比 ctor pool。

#### Goal / Action / Validation

- Goal: 把 foreach 子節點接到 ctor hop。
- Action: 介面方法名 + ctor EQ。
- Validation: 至少一次 `get_Current`:`XElement`{ctor}。

#### Applies When

- 授權 IL2CPP；Parse 走 `XContainer.Elements`／`GetElements`。

#### Does Not Apply When

- Current 確實 EQ Load 樹。
- 沒有 foreach（只有 `Element` 單點查找）。
- 純靜態。

#### Validation

Same Parse window; Current class `XElement` and ctor pool hit.

#### Promotion Target

- N/A this round（category index only）

#### Required Linked Updates

- `feedback/history/apk-analysis/unity-il2cpp/README.md`
- `feedback/history/apk-analysis/README.md`
