> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-07 - Unity game session may be custom TCP, not TLS/443

Status: validated

#### One-line Summary

Unity 客戶端常同時開 **HTTPS :443（CDN／廣告／發現）** 與一條 **非 443 的長 TCP**；後者才是業務 session。不要只 MITM 443 就宣稱「看不到協定」。

#### Human Explanation

IL2CPP metadata 裡同時有 `HttpService`、`WebSocket`、以及 `NetStateTCPSocket` 這類名字時，實際連線要看 PID 的 `ss`/`netstat`。若 ESTAB 落在自訂 port、`openssl s_client` **沒有 peer certificate**、閒置幀只有幾十 byte，那是自訂 framing（常見 ping），不是失敗的 HTTPS。443 上的 CloudFront／廣告 SDK 會讓人誤以為「全是 REST」。

#### Trigger

- Unity IL2CPP；功能頁已開。
- 問「他們用什麼協定」或 MITM 443 沒有 spin／桌內 payload。

#### Evidence

- Tool: `ss`/`netstat` on target PID；short `tcpdump` on the non-443 port；`openssl s_client` as TLS-or-not probe；metadata type names (`NetStateTCPSocket` vs HTTP).
- Sanitized excerpt: long-lived TCP to cloud VM on a non-443 port with tiny bidirectional frames; HTTPS :443 to CDN/social SDKs in parallel.
- Evidence path: target `docs/protocol.md`

#### Generalized Lesson

1. 先列 PID 的 ESTAB：**port 443 vs 其他**。
2. 非 443：用 `openssl s_client` 判斷是不是 TLS；無 cert／讀到數 byte 即停 → 自訂 TCP。
3. Metadata 的 `*Packet` / `SendPacket` / `SerializePacket` 指向 **packet 協定**，不是 REST path catalog。
4. HTTP helper（例如 resolve host）可以存在，仍不是玩法通道。

#### Agent Action

協定問題先 `ss` + 短 pcap，再決定 MITM 或 IL2CPP `SendPacket` hook。不要預設 OkHttp。

#### Goal / Action / Validation

- Goal: 分出 session 平面 vs CDN 平面。
- Action: PID sockets + TLS probe + metadata transport types。
- Validation: 非 443 上觀察到 keepalive；443 對應 CDN/SDK。

#### Applies When

- Unity／自有 `*Socket` 網路層；授權裝置可看 PID sockets。

#### Does Not Apply When

- 純 HTTPS REST／gRPC-on-443 客戶端。
- 需要解密 payload 才能回答的問題（本 lesson 只分平面）。

#### Validation

再開功能頁：非 443 session 應仍在；443 連線集合可換。Idle 幀保持很小。

#### Promotion Target

- `analysis/apk/traffic-triage.md`
- `analysis/apk/tools-and-failures.md`

#### Required Linked Updates

- 必須：上述兩檔各加一列（本輪已做）
- 必須：`feedback/history/apk-analysis/unity-il2cpp/README.md` 與 domain README Recent
- Project IPs／host 留在 `<PROJECT_ROOT>` docs
