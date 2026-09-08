> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-08 - XName.Get string may not be the Decrypt GetString object

Status: validated

#### One-line Summary

同窗口的 `XName.Get(String)`，其 String **不必**是 Decrypt 後 UTF-8 `GetString` 的回傳物件；不要 dump 字串來對。

#### Human Explanation

建 `XElement(XName)` 時名稱常走 `XName.Get`。人會猜 Get 吃的是剛解密的 XML 字串。pointer EQ 可能失敗（literal／intern／切片）。禁止讀字串內容。Parse 用 wrapper 也可能完全沒有 `set_Value`／`AddString`。

#### Trigger

- ctor `XName` 與 Load `get_Name` pointer NE。
- 想對上 Decrypt→GetString hop。

#### Evidence

- Tool: Frida; GetString x0 pool vs XName.Get arg0; class names + EQ/NE only.
- Sanitized observation: GetString pool nonempty; Get arg stayed NE; Parse wrappers had no text mutators.
- Evidence path: `<PROJECT_ROOT>/<App>/docs/` protocol notes

#### Generalized Lesson

1. `XName.Get` 的 String ≠ 自動等於 Decrypt `GetString`。
2. 只記 EQ/NE 與池是否非空。
3. 不要 dump String／`XName`／XML。

#### Agent Action

先池化 GetString return register，再比 `XName.Get` 第一個 String；禁止讀內容。

#### Goal / Action / Validation

- Goal: 避免把 Get 寫成 Decrypt 字串 hop，也不要用 dump 補洞。
- Action: pointer EQ + nonempty utf8 pool.
- Validation: NE with nonempty pool, or EQ recorded honestly.

#### Applies When

- 授權 IL2CPP；Decrypt→UTF-8 後再建 LINQ `XName`。

#### Does Not Apply When

- Get 的 String 確實 EQ GetString return。
- 純靜態。

#### Validation

Same Decrypt window: utf8 pool > 0; XName.Get String NE.

#### Promotion Target

- N/A this round（category index only）

#### Required Linked Updates

- `feedback/history/apk-analysis/unity-il2cpp/README.md`
- `feedback/history/apk-analysis/README.md`
