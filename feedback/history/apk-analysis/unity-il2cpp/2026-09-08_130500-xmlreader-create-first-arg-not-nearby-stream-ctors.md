> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-08 - XmlReader.Create first-arg class beats nearby stream ctors

Status: validated

#### One-line Summary

定位「解密後的 bytes 如何變成 XML」時，以 **`XmlReader.Create`／reader ctor 的 runtime 第一參數類型** 為準；同窗口的 `MemoryStream`／`Encoding.GetString` LIVE **不能**證明 `Load(Stream)`。

#### Human Explanation

IL2CPP 遊戲常同時做 UTF-8 decode、暫存 `MemoryStream`、以及真正餵給 `XDocument.Load` 的 reader。Idle session 幾乎一定會打到 `GetString` 與 `MemoryStream..ctor`。若據此寫「XML 來自 Stream」，會跟實際 `Create(TextReader)`（例如 `StringReader`）衝突。正確邊界是 hook `XmlReader.Create`（及實際 Load 用到的 reader 類型），只記錄第一參數的 **class 名**，不要 dump 字串或 stream 內容。

#### Trigger

- 已證明 live XML 走 `XDocument.Load(XmlReader)`，下一步要找 bytes → reader。
- `MemoryStream` 或 `Encoding.GetString` 在同一 attach 視窗大量 LIVE。

#### Evidence

- Tool: Frida unique-type hooks on `XmlReader.Create` / reader ctors / `StringReader` / `MemoryStream` / `Encoding.GetString`.
- Sanitized observation: `Create(TextReader)` fired with a string-backed reader; `Create(Stream)` did not; stream and GetString ctors still fired elsewhere.
- Evidence path: `<PROJECT_ROOT>/<App>/docs/` protocol notes

#### Generalized Lesson

1. 以 **Create／Load 的實際引數 runtime 類型** 當 XML 來源，不以同進程的 stream／encoding 熱點當來源。
2. `MemoryStream` LIVE ≠ `XmlReader.Create(Stream)` LIVE。
3. `Encoding.GetString(Byte[])` LIVE 只證明有 UTF-8（或該 encoding）轉換，不證明那就是 Load 的輸入，除非與 Decrypt／Create 做時間或 call-stack 相關。
4. 不要 dump `String`／stream／XML。

#### Agent Action

先 hook `XmlReader.Create` 各 overload 的第一參數 class；再決定是否需要 `StringReader`／`GetString` 相關性探針。

#### Goal / Action / Validation

- Goal: 避免把忙碌的 stream／encoding API 誤寫成 XML Load 路徑。
- Action: unique-type cap；對照 `Create(Stream)` 是否出現。
- Validation: `Create(TextReader)`（或實際 overload）有 LIVE，且 `Create(Stream)` 在同窗口無 LIVE。

#### Applies When

- 授權 IL2CPP；post-decrypt 或 keepalive 已進入 `System.Xml` Load／Create。

#### Does Not Apply When

- `Create(Stream)` 確實 LIVE（那就記錄 Stream 路徑）。
- 純靜態、無 live Create 證據。

#### Validation

Create first-arg class vs absent Stream overload in the same window.

#### Promotion Target

- N/A this round（category index only）

#### Required Linked Updates

- `feedback/history/apk-analysis/unity-il2cpp/README.md`
- `feedback/history/apk-analysis/README.md`
