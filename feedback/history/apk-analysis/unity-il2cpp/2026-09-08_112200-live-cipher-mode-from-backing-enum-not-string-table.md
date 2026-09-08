> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-08 - Live cipher mode from backing enum, not string table

Status: validated

#### One-line Summary

證明 live AES **mode / padding** 時，讀 managed algorithm 物件上 `CipherMode` / `PaddingMode` 的 **enum 成員名**；不要用 metadata 字串表，也不要 dump key／IV buffer。

#### Human Explanation

同一 IL2CPP image 常同時出現多種 mode／padding 字串與多個 AES 實作 class。那只證明編譯進 binary，不證明這個 session 的 live 物件用哪一個。若 cryptor wrapper 持有 `System.Security.Cryptography.SymmetricAlgorithm`／`Aes` 子類，backing field（常見名 `ModeValue`、`PaddingValue`）就是 enum。對 live object 讀該欄位並 map 成 **成員名**，即可記錄 mode／padding，而不讀 `KeyValue`／`IVValue`。

#### Trigger

- 已確認 live cryptor family 是 managed AES wrapper，但仍寫「mode 未證明」。
- Metadata／string table 出現 CBC、ECB、PKCS7 等多個候選。

#### Evidence

- Tool: Frida + IL2CPP field type／offset + enum static names.
- Sanitized observation: live algorithm class was a concrete `Aes*` provider; backing mode／padding fields mapped to enum member names; key／IV byte arrays were not read.
- Evidence path: `<PROJECT_ROOT>/<App>/docs/protocol.md`

#### Generalized Lesson

1. 字串表或未使用的 `RijndaelManaged`／`AesManaged` class **不是** live mode。
2. 在 live encrypt／decrypt 的 wrapper 上走到 algorithm 物件，讀 `CipherMode`／`PaddingMode` **名稱**。
3. 禁止讀取或記錄 `Byte[]` key／IV、key getter、native key-buffer API。
4. 這補足「只知道 implementation family」；仍不能從 mode 名推出 session key。

#### Agent Action

Hook 已證明的 encrypt／decrypt 入口；從 wrapper 讀 algorithm 指標；只 log 型別名與 mode／padding enum 名。

#### Goal / Action / Validation

- Goal: 在不擷取 key 的前提下記錄 live cipher mode／padding 名稱。
- Action: names-only probe；跳過所有 `Byte[]` 與 key 相關欄位。
- Validation: 同一 session 至少一次 encrypt 與 decrypt 得到相同 mode／padding 名；probe 輸出不含 buffer。

#### Applies When

- 授權 IL2CPP；live cryptor 持有 managed `SymmetricAlgorithm`／`Aes`。

#### Does Not Apply When

- 加密完全在 native、沒有 managed mode 欄位。
- 目標是可攜式 decoder／重放（另需授權與 secret handling）。

#### Validation

對 keepalive 或低風險 RPC 重跑；mode／padding 名穩定；git diff 無 key 素材。

#### Promotion Target

- N/A this round（category index only）

#### Required Linked Updates

- `feedback/history/apk-analysis/unity-il2cpp/README.md`
- `feedback/history/apk-analysis/README.md`
- `feedback/history/apk-analysis/unity-il2cpp/2026-09-08_094000-inspect-live-cryptor-object-before-key-recovery.md`（修訂第 5 點）
