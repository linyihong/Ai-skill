> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-07 - ParseResult arg class can be XDocument when C2S is not XML

Status: validated

#### One-line Summary

IL2CPP 自訂 TCP 上，C2S `EncryptData` 字串的 first-char 分類**不能**推論 S2C 格式；要對 `ParseResult` 的 **arg1 class 名**（常見是 `System.Xml.Linq.XDocument`），且功能手勢不一定會走進該 packet 的 `ParseResult`。

#### Human Explanation

同一條 session 常常不對稱：送出的命令字串看起來像字母 token（不是 `{` / `<`），收回來的 RPC 解析入口卻吃 LINQ-to-XML 文件物件。只 hook `EncryptData` 長度／首字元會把協議標成「非 XML」，接著錯過真正的 S2C schema 入口。

另外，HUD 已更新不代表對應動作型別呼叫了 `ParseResult`。有的動作包沒有 override；界面方法指標（例如 `IPacket.ParsePacket`）hook 了也不會接到實作。要記 **this class + arg1 class + retval class**，不要印 XDocument／string 內容。

#### Trigger

- 已用 first-char 把 C2S 標成 alpha token / not-JSON。
- 問「response 是 JSON 還是 XML」或「結果物件型別」。
- 功能手勢後 HUD 變了，但 type-name log 沒有該動作 packet 的 `ParseResult`。

#### Evidence

- Tool: Frida attach; intercept IL2CPP methods named `ParseResult`; log `il2cpp_object_get_class` for `this`, arg1, retval only.
- Sanitized excerpt: idle/background RPC `ParseResult` arg1 = XML document type; retval = typed arrays or value-type/`nobj`. Feature action that previously showed a distinct C2S packet class did not emit `ParseResult` on that class in the same window.
- Evidence path: `<PROJECT_ROOT>/<App>/docs/protocol.md`

#### Generalized Lesson

1. 分開記 C2S 字串分類與 S2C parse **參數型別**。
2. `ParseResult`：log class 名，不 log XML/JSON 正文。
3. Hook 具體 `Packet\`1.ParseResult`（或同等 base），不要只 hook interface `ParsePacket`。
4. 功能 packet 沒有 `ParseResult` hit 時，標 `parse-path-unknown`，不要把 idle RPC 的 XML 入口當成該動作已解析。
5. retval 若 `nobj`，多半是 valuetype boxing；記「非 reference DTO」，不要當 hook 失敗。

#### Agent Action

在 type-name C2S hook 之後加 names-only `ParseResult`。禁止 dump `XDocument`、明文、金鑰。

#### Goal / Action / Validation

- Goal: 知道 S2C 解析入口的 **CLR 型別名**，而不是猜 JSON。
- Action: 一次 idle 基線 + 一次功能手勢；比對 `ParseResult` this/arg1/ret。
- Validation: 至少一種背景 RPC 的 arg1 class 穩定可再現；功能包若未出現則寫明缺口。

#### Applies When

- Unity IL2CPP；授權 attach；存在 `ParseResult` 方法名。

#### Does Not Apply When

- 純 HTTPS JSON REST（看 converter / path 即可）。
- 需要元素名或欄位值才能回答的問題。

#### Validation

重複 idle 窗口應再現相同 arg1 class。換一個背景 packet 型別，arg1 仍應是同一 XML 文件型別才算「envelope 級」結論。

#### Promotion Target

- N/A this round（category index only）

#### Required Linked Updates

- 必須：`feedback/history/apk-analysis/unity-il2cpp/README.md`、`feedback/history/apk-analysis/README.md`（count + recent row）
- Project 細節留在 `<PROJECT_ROOT>` docs
