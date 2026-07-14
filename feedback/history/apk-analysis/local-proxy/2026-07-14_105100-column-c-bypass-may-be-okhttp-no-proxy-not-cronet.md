> 遵守 [共用規則索引](../../../../enforcement/README.md)、[sanitization](../../../../enforcement/sanitization.md)、[reusable-guidance-boundary](../../../../enforcement/reusable-guidance-boundary.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson。

### 2026-07-14 - Column-C bypass may be explicit OkHttp no-proxy (not Cronet)

Status: validated

#### One-line Summary

MITM 對主業務 API **零 CONNECT**、但 Frida 已確認仍是 **OkHttp／Retrofit** 時，優先在 client builder 上找 **顯式 no-proxy／Proxy.NO_PROXY／disable system proxy** 設定；不要預設結論為 Cronet／QUIC。

#### Human Explanation

三欄分流的欄 C（business API proxy membership）過去常把「零 CONNECT」歸因於 Cronet、native，或「自訂 client」。實測可出現第三種已驗證型態：

1. no-proxy pcap／SNI 仍見業務 API；
2. in-process hook 打到 `RealInterceptorChain`／Retrofit service；
3. App 在建立 `OkHttpClient.Builder` 時呼叫類似 `applyReleaseNoProxy`／`applyNoProxyIfNeeded(builder, true)` 的路徑，使 **release 業務 client 不使用系統 HTTP 代理**。

此時 MITM 拉再久也看不到業務 path；正確升級是 Frida／in-process，並把 bypass **機制**記成「OkHttp no-proxy config」，而非「換了傳輸棧」。

#### Trigger

- 欄 C：業務 API 零 CONNECT + no-proxy 仍有 443。
- Frida：OkHttp／Retrofit 有 first-party URL；Cronet `newUrlRequestBuilder` 計數為 0（或僅 ads）。
- DEX／runtime 出現 `*ProxyConfig*`、`NO_PROXY`、`noProxy`、`setProxy(Proxy.NO_PROXY)`、`proxySelector` 覆寫。

#### Evidence

- Tool: Frida hook on app networking helper that configures `OkHttpClient.Builder`；paired with prior MITM／pcap triage.
- Sanitized: methods resembling `applyReleaseNoProxy` + `applyNoProxyIfNeeded(builder, true)` fired at client build.
- Target hosts／package names stay in `<PROJECT_ROOT>/…/api/dynamic-*.md`.

#### Generalized Lesson

After confirming column-C bypass:

1. Identify the **business HTTP stack** in-process（OkHttp vs Cronet vs other）.
2. If OkHttp：grep／hook **proxy configuration** on the client builder **before** assuming Cronet.
3. Record bypass class as one of: `okhttp-no-proxy`｜`cronet-quic`｜`native`｜`unknown`——分項書寫，避免混成「自訂 client」單字。
4. Ads 可 MITM 只證明欄 A，不推翻欄 C 的 no-proxy client。

#### Agent Action

1. MITM 欄 C = 零 CONNECT → Frida URL hook。
2. 見 OkHttp first-party → hook／靜態讀 `*Proxy*`／`OkHttpClient.Builder` 組裝點。
3. 若確認 no-proxy：停止延長純 MITM；文件寫 `bypass=okhttp-no-proxy`。
4. 繼續章節／sign／decrypt 用 in-process；不要為了「讓流量進 Charles」硬改未授權的 proxy 設定除非使用者要求。

#### Promotion Target

- Layer: analysis + intelligence
- Status: promoted
- Required Linked Updates:
  - `analysis/apk/traffic-triage.md` — 欄 C 下一步補 OkHttp no-proxy 分支
  - `intelligence/engineering/analytical-reasoning/signals/mitm-route-column-signals.md` — 欄 C 增加 OkHttp no-proxy 訊號
  - N/A: workflow／enforcement（本條不改）

#### Verification

Validated when Frida shows OkHttp business URLs under column-C zero-CONNECT, and builder hooks prove an explicit no-proxy path；Cronet business counters stay inactive in the same window.
