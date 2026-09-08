> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-08 - Same-thread dt window after Decrypt, not global LIVE

Status: validated

#### One-line Summary

要證明「Decrypt 的 `Byte[]` 如何變成 XML reader」，用 **同一 thread、Decrypt `onLeave` 後極短時間內** 的 follower 命中；不要用全進程的 `GetString`／`MemoryStream` LIVE。

#### Human Explanation

Encoding 與 stream API 在 idle session 幾乎一直在跑。全域 unique-type LIVE 只能說「有人呼叫」，不能說「接在這次 Decrypt 後面」。在 Decrypt `onLeave` 記下 thread id 與時間，follower 只在同 thread、數十毫秒內記一次 **方法簽名**（不要 dump `Byte[]`／`String`），才能把 hop 寫進協議。

#### Trigger

- 已有 live Decrypt 與 live `XmlReader.Create`／`StringReader`，但兩者可能只是同窗口熱點。
- 需要 bytes → string → reader 的順序，且禁止讀明文。

#### Evidence

- Tool: Frida Interceptor on Decrypt plus followers; unique-tag cap; same-thread short dt.
- Sanitized observation: GetString / string-backed reader / Create(TextReader) fired on the decrypt thread immediately after Decrypt; Stream Create still absent.
- Evidence path: `<PROJECT_ROOT>/<App>/docs/` protocol notes

#### Generalized Lesson

1. Producer `onLeave` 設 thread + timestamp；follower 過濾同 thread 短 dt。
2. 只記錄 follower **方法名與第一參數類型**，不要讀 buffer／string。
3. 全域 LIVE 仍可能是別條路徑；相關性窗口失敗就不要寫成 hop。
4. dt 視窗是啟發式，不是密碼學證明；需要更硬證據時再做 **指標相等**（仍不讀內容）。

#### Agent Action

先相關性窗口；通過後才考慮 retval vs GetString 參數 pointer 比較。

#### Goal / Action / Validation

- Goal: 把 Decrypt 與 XML reader 建成可寫進文件的 hop，且不去敏失敗。
- Action: same-thread dt filter + unique tags。
- Validation: Decrypt 後同 thread 出現 GetString→StringReader→Create(TextReader)；Create(Stream) 仍無。

#### Applies When

- 授權 IL2CPP；session cryptor Decrypt 回 `Byte[]`，XML 走 `System.Xml`。

#### Does Not Apply When

- Decrypt 與 parse 不在同 thread（視窗會假陰性；改 queue／handler hook）。
- 只做靜態 metadata。

#### Validation

至少一次 Decrypt 後同 thread 短 dt 的 follower 序列；對照 Stream Create 未命中。

#### Promotion Target

- N/A this round（category index only）

#### Required Linked Updates

- `feedback/history/apk-analysis/unity-il2cpp/README.md`
- `feedback/history/apk-analysis/README.md`
