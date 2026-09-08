> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-08 - XName.Get Substring.this may not be XmlReader/XName getters

Status: validated

#### One-line Summary

`XName.Get`←`Substring` 時，`Substring.this` **不必**等於同窗口 XmlReader／`XName`／`XElement` 的 name getter 回傳字串。名稱池 nonempty 時 NE 才有意義。不要 dump。

#### Human Explanation

Decrypt GetString 對不上之後，下一步常猜 reader `get_LocalName`／`get_Name`。要池化那些 getter 的 return register 再比 Substring.this。NE 表示切片來源是別條 String（內部 buffer／literal）。禁止讀內容。

#### Trigger

- Get←Substring 且 Substring.this ≠ Decrypt GetString。
- 想對上 XmlReader 名稱 hop。

#### Evidence

- Tool: Frida; reader/`XName`/`XElement` name getter x0 pool vs Substring.this; EQ/NE + pool size.
- Sanitized observation: name pool nonempty; Substring.this stayed NE.
- Evidence path: `<PROJECT_ROOT>/<App>/docs/` protocol notes

#### Generalized Lesson

1. reader name getter ≠ 自動等於 Substring.this。
2. 先確認名稱池 nonempty。
3. 不要 dump String／`XName`。

#### Agent Action

Get←Substring 後池化 XmlReader/`XName` getters，再比 this。

#### Goal / Action / Validation

- Goal: 避免把名稱切片 hop 寫成 reader LocalName，也不用 dump。
- Action: getter x0 pool vs Substring.this.
- Validation: EQ or honest NE with nonempty pool.

#### Applies When

- 授權 IL2CPP；Decrypt 窗口裡 `XName.Get`←Substring。

#### Does Not Apply When

- this 確實 EQ reader getter。
- 純靜態。

#### Validation

Same Decrypt window: nName>0; Get=sub this=NE.

#### Promotion Target

- N/A this round（category index only）

#### Required Linked Updates

- `feedback/history/apk-analysis/unity-il2cpp/README.md`
- `feedback/history/apk-analysis/README.md`
