# Local Proxy — Feedback Lessons

| 檔名 | Status | 標題 | 一句話摘要 |
|------|--------|------|-----------|
| `2026-10-23_134500-wire-result-glass-must-match-processresult-strip.md` | candidate | Wire RESULT GLASS must match ProcessResult strip | session-proxy rewrite 的 GLASS cell 數須對齊柜机 strip 契約；thin visible-only 常卡停輪 |
| `2026-07-14_105100-column-c-bypass-may-be-okhttp-no-proxy-not-cronet.md` | candidate | Column-C bypass may be OkHttp no-proxy | 部分「繞過」是 OkHttp 不走系統代理，未必是 Cronet |
| `2026-07-14_095210-mixed-mitm-ads-ok-firstparty-ssl-api-bypass.md` | candidate | Mixed MITM: ads OK, first-party SSL bypass | 廣告路徑可見但主 API SSL 失敗時分層判斷 |
| `2026-07-14_095200-on-device-ssl-certificate-ui-is-tls-path-signal.md` | candidate | On-device SSL certificate UI is TLS path signal | 裝置憑證 UI 可當 TLS 路徑信號 |
| `2026-06-22_120100-mitm-cdn-visible-primary-api-invisible-pinning-tier.md` | candidate | MITM CDN visible, primary API invisible | CDN 可見而主 API 不可見時多半是 pinning／分層 |
| `2026-05-13_095200-multi-layer-proxy-architecture-capture-strategy.md` | candidate | Multi-layer proxy capture strategy | 多層代理需分層 hook，單一策略不夠 |
| `2026-05-13_095000-non-standard-tls-bypasses-ssl-hooks.md` | candidate | Non-standard TLS bypasses SSL hooks | 非標準 TLS 使 SSL_read／write hook 失效 |
| `2026-05-01_132400-netty-aggregated-request-tostring-headers.md` | candidate | Netty aggregated request toString headers | 聚合請求 toString 可讀 headers |
| `2026-05-01_131000-cast-netty-request-for-handler-route.md` | candidate | Cast Netty request for handler route | handler 路由前正確 cast Netty request |
| `2026-05-01_114300-local-proxy-handler-uri-hook.md` | candidate | Local proxy handler URI hook | 本機 proxy handler URI hook 要點 |
| `2026-04-30_120010-本機-loopback-proxyserver-轉發會讓-wi-fi-http-mitm-看不到業務-connect.md` | candidate | Loopback ProxyServer hides Wi-Fi MITM CONNECT | loopback 轉發會讓 Wi-Fi HTTP MITM 看不到業務 CONNECT |
| `2026-04-30_120009-內建-sing-box-tun-類通道可能繞過-wi-fi-系統代理.md` | candidate | Built-in TUN may bypass Wi-Fi system proxy | TUN 類通道可能繞過系統代理 |
