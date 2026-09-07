> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-07 - IL2CPP SendPacket arg class is the opcode

Status: validated

#### One-line Summary

自訂 TCP session 要對 UI 操作時，hook `SendPacket` 並只記 **arg 的 IL2CPP class 名**；不要印 buffer、`EncryptData` 字串或金鑰。

#### Human Explanation

線上 C2S 常是密文，tcpdump 只有長度。真正的「這次送了哪種封包」在序列化物件的型別上。`SendPacket(this, packet)` 的 `this` 往往是 socket；`packet` 才是 `Network.Packets.*`。`EncryptData` 的參數可能已經是 `System.String`，印出來就等於 dump 明文序列化結果。

心跳路徑可能是 `SendPingPacket` 再送一個時間類封包，不要把每次 `SendPacket` 都當成功能操作。

#### Trigger

- Unity IL2CPP；已確認非 443 自訂 TCP session。
- 問「旋轉／進桌對應哪個 packet」或「有哪些屬性」。

#### Evidence

- Tool: Frida attach + `il2cpp_*` exports on `libil2cpp.so`; intercept method pointers named `SendPacket` / `EncryptData`; log `il2cpp_object_get_class` names only.
- Sanitized excerpt: idle = ping helper + time packet; enter-feature = play-now / observe / join-shaped types; action = bet-shaped type then next-shaped type. Info DTO field names like amount/stars may appear on a sibling `*Info` class, not on the packet type.
- Evidence path: `<PROJECT_ROOT>/<App>/docs/protocol.md`

#### Generalized Lesson

1. 對 `SendPacket` 記 **arg1 class**，不是 socket class。
2. 不要 log `EncryptData` 的 string／byte 內容。
3. 先看 idle 基線，再點一次功能，diff 新增的型別名。
4. `il2cpp_class_get_fields` 只當欄位**名**清單；空清單不代表沒有序列化內容（可能在 base／properties／DTO）。

#### Agent Action

先 type-name hook，再考慮 field **types**。禁止 `GetCryptorKeyBuffer`、明文 dump、重放。

#### Goal / Action / Validation

- Goal: UI 步驟對到 packet **類別名**。
- Action: Frida names-only；一次明確 gesture。
- Validation: 功能步驟出現與 idle 不同的 `Network.Packets.*`（或同等 namespace）名稱。

#### Applies When

- IL2CPP 仍 export `il2cpp_class_get_name` / `il2cpp_object_get_class`；授權裝置可 attach。

#### Does Not Apply When

- 純 HTTPS REST（看 path 即可）。
- 需要欄位值或解密才能回答的問題。

#### Validation

同一手勢重複一次，應再現相同 class 名序列（允許穿插 time／buddy 類背景包）。

#### Promotion Target

- N/A this round（category index only）

#### Required Linked Updates

- 必須：`feedback/history/apk-analysis/unity-il2cpp/README.md`
- Project 細節留在 `<PROJECT_ROOT>` docs
