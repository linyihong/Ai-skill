> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-08 - Same packet class EncryptData length can vary

Status: validated

#### One-line Summary

同一 `SendPacket` 型別名的 `EncryptData` 字串長度可以隨桌／下注變，不能把某次量到的 len 當成 opcode 常數；分類仍看 first-char，不要印 body。

#### Human Explanation

命令字串常含桌或下注相關 token，長度會 +1／+N。若文件寫死「下注封包永遠 52」，下一台櫃或另一檔總下注就會看起來像新 opcode。正確錨點是 packet **class 名** + first-char 類別（alpha / json / xml），len 只當觀察值。

**2026-09-08 revision:** 同一長連線稍後再量，**連 ping（`GetTime`）與進場／離桌 opcode 都可能整批 +1**。不要把「只有旋轉變長」當成規則；first-char 仍穩。

**Same-day revision 2:** 同一 class 也可以差超過 1：大廳 **featured vs grid** 的 category 封包、以及進場 `PlayNow*` 都量到更短／更長的觀察值。仍用 class 名，不要把區間上下界當 opcode。

另外：掃 IL2CPP 方法時 `Interceptor.attach` 碰到不可 hook 的 pointer 會整支 script abort，其餘 C2S hook 也一起沒了。應用 try/catch 跳過該 pointer。

#### Trigger

- 同一 packet class 在不同機台或不同總下注出現不同 EncryptData len。
- Frida script 在掃 class 中途 `unable to intercept function` 後完全靜音。

#### Evidence

- Tool: Frida names-only EncryptData length + first UTF-16 code unit.
- Sanitized excerpt: bet/next types unchanged; len +1 vs earlier cabinets. Same session later: ping/enter/leave also +1 vs 2026-09-07 constants.
- Evidence path: `<PROJECT_ROOT>/<App>/docs/protocol.md`

#### Generalized Lesson

1. 記錄 `class + len + first-char class`，不要把單次 len 當常數。
2. 比較機台時，len 漂移 ≠ 新協議。
3. 批量 `Interceptor.attach` 必須 try/catch；失敗跳過該 method ptr。
4. 禁止印 command 正文。

#### Agent Action

短窗、names-only。Len 只寫數字與 kind。

#### Goal / Action / Validation

- Goal: 避免把字串長度誤當成 opcode。
- Action: 至少兩台櫃或兩檔下注各一次旋轉。
- Validation: 型別名相同、first-char 相同、len 允許不同。

#### Applies When

- IL2CPP；命令是 `EncryptData(System.String)`；授權 attach。

#### Does Not Apply When

- 固定長度 binary opcode（不是字串命令）。
- HTTP path 本身就是 opcode。

#### Validation

再換一檔總下注或另一台櫃，應仍是同一 pair of packet classes。

#### Promotion Target

- N/A this round（category index only）

#### Required Linked Updates

- `feedback/history/apk-analysis/unity-il2cpp/README.md`
- `feedback/history/apk-analysis/README.md`
- Project 細節留 `<PROJECT_ROOT>` docs
