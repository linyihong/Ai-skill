> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-08 - Client IL2CPP may embed unused server-named factory types

Status: validated

#### One-line Summary

Client 的 IL2CPP metadata 出現 `Server.*` packet factory／header **不代表** live 收包走這條建構路徑；要看到方法被呼叫，或另證 receive／parse 邊界。

#### Human Explanation

共用遊戲堆疊常把 server 端 packet helper 一起編進 client。列出 factory 方法簽名（`Round`／`Join`／`Response`…）有助於理解 **候選** 信封，但 idle session 或實際 S2C apply 可能從不呼叫它們。Live 路徑可能是另一組 `Network.*` packet、`ParseResult`、或 game-action apply。把未觸發的 factory 當成 wire constructor 會誤導後續 hook。

#### Trigger

- 掃 class 名得到 `Server.Core`（或同等）`PacketFactory`，而 receive hook 沒打到這些方法。
- 已有另一條已驗證的 S2C apply（例如 parse XML 後進 process／action）。

#### Evidence

- Tool: Frida method-name + signature dump；Interceptor on factory methods during idle traffic.
- Sanitized observation: factory signatures were present; idle keepalive did not enter those methods; receive／apply evidence lived on a different type path.
- Evidence path: `<PROJECT_ROOT>/<App>/docs/` protocol notes

#### Generalized Lesson

1. Metadata 方法清單 ≠ live call graph。
2. 先 hook factory；若 idle／已知 RPC 都不進，標成 **dead metadata for this window**，不要寫成「S2C 由此建構」。
3. 繼續以已證明的 parse／apply／SendPacket 邊界為主。
4. 不要 dump factory 的 `object[]` payload。

#### Agent Action

Dump signatures for orientation；attach factory methods with a unique-type cap；若無 LIVE 行，文件寫「未觀察到呼叫」。

#### Goal / Action / Validation

- Goal: 避免把 client 內嵌的 server helper 誤當成 live S2C factory。
- Action: signature dump + 有時限的 idle hook。
- Validation: keepalive 或已知 C2S 期間 factory hook 無命中，且既有 parse／apply hook 仍能解釋 UI 更新。

#### Applies When

- 授權 IL2CPP；client image 同時有 `Server.*` 與 `Network.*`（或同等雙命名空間）packet 類型。

#### Does Not Apply When

- Factory 方法在 live receive 上確實被呼叫（那就當 live 路徑記錄）。
- 純靜態、無裝置可驗證 call graph。

#### Validation

Idle 窗口無 factory LIVE；對照既有 apply／parse 證據。

#### Promotion Target

- N/A this round（category index only）

#### Required Linked Updates

- `feedback/history/apk-analysis/unity-il2cpp/README.md`
- `feedback/history/apk-analysis/README.md`
