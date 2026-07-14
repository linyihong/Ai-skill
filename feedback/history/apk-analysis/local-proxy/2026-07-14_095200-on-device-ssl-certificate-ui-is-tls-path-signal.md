> 遵守 [共用規則索引](../../../../enforcement/README.md)、[sanitization](../../../../enforcement/sanitization.md)、[reusable-guidance-boundary](../../../../enforcement/reusable-guidance-boundary.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson。

### 2026-07-14 - On-device SSL certificate UI is a TLS-path signal

Status: candidate

#### One-line Summary

系統代理開啟後若 **App 介面**出現 SSL／憑證驗證錯誤，與 MITM log 的 `certificate unknown` 同屬 **TLS trust 失敗**訊號；不可誤判成「App 壞了」或「沒有業務流量」。

#### Human Explanation

MITM 視窗常見兩種獨立訊號：

1. 代理 log：`Client TLS handshake failed` / `certificate unknown`。
2. **裝置 UI**：App 內 WebView／網路錯誤頁顯示 SSL certificate verification failed（或同等文案）。

兩者可同時成立：client 已導流到代理、但拒收代理憑證。此時 UI 是使用者可見的 trust 失敗證據，應寫入專案 dynamic note，並與「第三方流量已解密成功」對照——證明導流成功、**第一方**卡在 trust／pinning 層。

若同時業務主 API **完全沒有 CONNECT**，而 no-proxy pcap／SNI 仍見業務候選 IP，則另屬 **proxy bypass**（見既有「proxy config vs business route」lesson），與 UI SSL 錯誤是**不同路徑**——可能同一 App 內 H5／次要 host 進代理並彈 UI，主 API 直連繞過。

#### Trigger

- 裝置開著系統／Wi‑Fi HTTP(S) proxy 冷啟動或進某一頁。
- 使用者回報介面出現 SSL／certificate verification。
- MITM 對某一第一方 host 有 CONNECT + handshake fail；或對主 API 零 CONNECT。

#### Evidence

- Tool: PC MITM log + user-visible on-device UI observation + optional no-proxy pcap.
- Sanitized: third-party hosts decrypt OK；one first-party mobile/web host CONNECT + cert reject；primary API host absent from MITM；pcap without proxy still shows API-tier destination IP.
- Evidence path: `<PROJECT_ROOT>/api/dynamic-*.md`、`<PROJECT_ROOT>/capture/`（勿把 raw UI 截圖／host 寫進本 lesson）。

#### Generalized Lesson

```text
UI SSL error while proxy on
  → treat as TLS trust failure evidence (same tier as mitm handshake fail)
  → still split routes:
       host entered proxy + UI/SSL fail  → pinning / custom trust on that host
       primary API never entered proxy  → bypass; confirm with no-proxy pcap/SNI
```

Do not stop the analysis window solely because the App shows a certificate page—that page is often the trust-layer proof you needed.

#### Agent Action

1. Record UI SSL observation in project dynamic notes（not as “install broken”).
2. Correlate with MITM：which host handshake-failed vs which never appeared.
3. If primary API absent from MITM, run/compare no-proxy SNI/pcap before concluding “app offline”.
4. Escalate business API capture to Frida / in-process hooks; keep MITM for proxy-aware stacks only.

#### Goal / Action / Validation

- Goal: Use on-device SSL UI as structured evidence, not noise.
- Action: Pair with proxy-導流 vs TLS 兩層 + business-route bypass lessons.
- Validation: Same window shows ads decryptable + UI SSL on one first-party host + primary API only on no-proxy path.

#### Applies When

- Android APK dynamic analysis with PC MITM and interactive cold start.

#### Does Not Apply When

- Pure offline static analysis with no runtime proxy.
- Non-HTTP features with no network UI.

#### Promotion Target

`analysis/apk/traffic-triage.md` / `tools-and-failures.md` — candidate until second consumer corroborates UI-as-signal wording.
