> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-08 - XName.Get string may be String.Substring, not Decrypt GetString

Status: validated

#### One-line Summary

`XName.Get(String)` 的 String 可能是同窗口 **`String.Substring` 回傳物件**，不是 Decrypt `GetString`；其他 Get 可能完全 NE。不要 dump 切片。

#### Human Explanation

名稱 hop 對不上完整 UTF-8 字串時，下一步是池化 `Substring`／`Intern`／`Concat` 的 return register，再比 Get 第一參數。NE 可能是 literal。禁止讀字串內容。

#### Trigger

- `XName.Get` 的 String ≠ Decrypt `GetString`。
- 懷疑從 XML 字串切片出 local name。

#### Evidence

- Tool: Frida; Substring/utf8/LocalName pools vs Get arg0; class names + EQ/NE only.
- Sanitized observation: at least one Get hit Substring; another Get stayed NE.
- Evidence path: `<PROJECT_ROOT>/<App>/docs/` protocol notes

#### Generalized Lesson

1. Get 參數可能是 Substring 物件。
2. 同窗口可同時有 NE（literal／未入池）。
3. 不要 dump String／`XName`。

#### Agent Action

GetString NE 之後加 Substring（及 Intern／Concat）池，再比 Get。

#### Goal / Action / Validation

- Goal: 避免把「不是 GetString」停在死巷，也不用 dump 補洞。
- Action: Substring x0 pool vs Get arg.
- Validation: Substring hit and/or honest NE.

#### Applies When

- 授權 IL2CPP；Decrypt→UTF-8 後 `XName.Get`。

#### Does Not Apply When

- Get 已 EQ GetString。
- 純靜態。

#### Validation

Same Decrypt window: Get arg EQ Substring x0, or NE with nonempty producer pool.

#### Promotion Target

- N/A this round（category index only）

#### Required Linked Updates

- `feedback/history/apk-analysis/unity-il2cpp/README.md`
- `feedback/history/apk-analysis/README.md`
