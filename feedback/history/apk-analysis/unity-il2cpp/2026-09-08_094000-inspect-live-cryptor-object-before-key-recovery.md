> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-08 - Inspect live cryptor object before key recovery

Status: validated

#### One-line Summary

IL2CPP session 欄位若宣告為 cryptor interface，先讀 live object class 與 wrapper child class，再在加密前觀察 packet shape；通常不需要擷取 key 或明文正文。

#### Human Explanation

Metadata 同時列出多個 cryptor implementation，只能證明候選集合。於 live `EncryptData` 呼叫讀取持有者的 interface field，並以 `il2cpp_object_get_class` 取得實際 class，才能判定這個 session 用哪一個 implementation。若該 implementation 是 wrapper，再以 field offset 讀 child object 的 class 名，可確認 native delegate 類型。

要理解 request 結構，不必先恢復 wire key。可在 `SendPacket` 與 `EncryptData` 之間記錄 packet inheritance、field name/type，以及敏感 string 的 length / first-char class；若 packet request 長度與加密輸入一致，即可建立 pre-encrypt dataflow，而不輸出正文。

#### Trigger

- Metadata 同時出現 AES／XOR／plain 等多個候選 implementation。
- 目標是確認 live protocol shape，而不是建立外部解密器。
- Session object 有 interface-typed cryptor field。

#### Evidence

- Tool: Frida + IL2CPP reflection exports.
- Sanitized observation: interface field resolved to one live wrapper class; its encrypt/decrypt children resolved to the same native delegate class.
- Evidence path: `<PROJECT_ROOT>/<App>/docs/protocol.md`

#### Generalized Lesson

1. 候選 class name 不等於 live implementation。
2. 先解析 interface field 的 runtime class；wrapper 再解析 child runtime class。
3. 先在加密前量 packet field shape，避免不必要的 key extraction。
4. 敏感字串只記 type、length、first-char class；禁止印 body。
5. Runtime class 能證明 implementation family，但不能單獨證明 mode、IV、padding 或 key schedule。**修訂（2026-09-08）：** family 仍不夠；若 live algorithm object 暴露 `CipherMode` / `PaddingMode` backing field，可讀 **enum 成員名**（不要讀 `Byte[]` key/IV）。字串表出現 CBC 等名稱仍不能當 live mode。見 `2026-09-08_112200-live-cipher-mode-from-backing-enum-not-string-table.md`。

#### Agent Action

建立 names-only probe：`owner field declared type → actual class → child actual class`，並將 packet fields 與加密輸入以長度交叉核對。

#### Goal / Action / Validation

- Goal: 判定 live cryptor implementation 並取得安全的 pre-encrypt packet shape。
- Action: hook 持有 cryptor 的 session method 與 packet send boundary；只讀 class/field/type metadata。
- Validation: 同一 session 至少看到一次 encrypt call，且 interface field、wrapper child 與 method receiver class 相互一致。

#### Applies When

- 授權 IL2CPP 分析；cryptor 由 managed session object 持有；可 attach live process。

#### Does Not Apply When

- cryptor 完全位於 native global state，managed object 沒有可反查欄位。
- 目標明確要求可攜式 wire decoder；那需要另外的授權、secret handling 與測試契約。

#### Validation

用低風險 request 重跑一次；確認 runtime class 與 encrypt receiver 一致，且 probe 沒有輸出 key、buffer 或 string body。

#### Promotion Target

- N/A this round（category index only）

#### Required Linked Updates

- `feedback/history/apk-analysis/unity-il2cpp/README.md`
- `feedback/history/apk-analysis/README.md`
