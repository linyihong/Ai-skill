> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-08 - XName.Get Substring.this may not be Decrypt GetString

Status: validated

#### One-line Summary

`XName.Get` 吃到的 `String.Substring` 回傳值，其 **`this` 不必是 Decrypt `GetString` 物件**；utf8 池非空時 NE 才有意義。不要 dump 切片。

#### Human Explanation

Get←Substring 之後，人會假設切片來源是剛解密的 XML 字串。要對 `Substring` 的 `this`（args[0]）與 GetString return register。NE 表示名稱切片來自別的 String（literal／reader／intern）。禁止讀內容。

#### Trigger

- Get arg 已 EQ Substring retval。
- 想關閉 Decrypt→name 字串 hop。

#### Evidence

- Tool: Frida; utf8 x0 pool vs Substring this vs Get arg; EQ/NE only.
- Sanitized observation: Get hit Substring; Substring.this stayed NE; utf8 pool size 1.
- Evidence path: `<PROJECT_ROOT>/<App>/docs/` protocol notes

#### Generalized Lesson

1. Substring retval hop ≠ Substring.this hop。
2. 先確認 utf8 池 nonempty 再解釋 NE。
3. 不要 dump String。

#### Agent Action

Get←Substring 後立刻比 Substring.this 與 GetString x0。

#### Goal / Action / Validation

- Goal: 避免把切片 hop 寫成「從解密 XML 切 local name」。
- Action: three-pointer EQ (Get arg, Substring ret, Substring this, GetString).
- Validation: thisUtf8 true or honest NE with nonempty utf8 pool.

#### Applies When

- 授權 IL2CPP；Decrypt→UTF-8 且 `XName.Get`←Substring。

#### Does Not Apply When

- Substring.this 確實 EQ GetString。
- 純靜態。

#### Validation

Same Decrypt window: Get←Substring; this≠utf8; nUtf8>0.

#### Promotion Target

- N/A this round（category index only）

#### Required Linked Updates

- `feedback/history/apk-analysis/unity-il2cpp/README.md`
- `feedback/history/apk-analysis/README.md`
