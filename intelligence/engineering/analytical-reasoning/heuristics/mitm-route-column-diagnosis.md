# MITM Route-Column Diagnosis Heuristic（MITM 三欄分流啟發式）

## 問題

系統代理 + MITM 開啟後，如何避免把「廣告能解密／UI 跳 SSL／看不見業務 API」合成一個錯誤結論（例如全怪 pinning）？

## 原則

- MITM 結果是**多路由**現象，不是單一 yes/no。
- 裝置 UI 的 SSL certificate verification 與 MITM `certificate unknown` 同屬 **trust 失敗**訊號，應留證。
- 主業務 API **零 CONNECT** ≠ pinning；那是 **proxy membership／bypass**，要用 no-proxy pcap／SNI + in-process hook 證明。

## 決策表

| 同窗觀察 | 欄位標籤 | HOW TO THINK | HOW TO DO 下一動 |
| --- | --- | --- | --- |
| Ads／SDK CONNECT 可解密 | A Proxy-aware OK | 導流成立 | 不要宣稱業務主線已覆蓋 |
| 某一第一方 host CONNECT + handshake fail；或 UI SSL | B First-party trust fail | 該 host pinning／custom trust | 記 trust 證據；WebView／H5 可另線處理 |
| 主業務 API 零 CONNECT；關代理後 pcap 仍見業務 IP | C API bypass | client 不理 HTTP proxy | Frida／Cronet／OkHttp／native 高語意 hook |
| 主業務 API CONNECT + handshake fail | C pinning-tier | 與 B 同類，但是業務 host | trust bypass + 繼續用 MITM 或 hook |
| 三欄同時出現 | Mixed | 禁止單標籤 | 專案 dynamic note 分三欄寫；intelligence 用本表裁決 |

## 不建議的做法

- 見 UI SSL → 關掉代理當「修好」，不記 trust 證據。
- 見大量廣告解密 → 以為 MITM 已抓到主線，乾等業務 path。
- 見 MITM 無業務 host → 直接寫 pinning，不做 no-proxy 對照。

## 相關 atoms

- Procedure：[`analysis/apk/traffic-triage.md`](../../../../analysis/apk/traffic-triage.md) §三欄分流
- Failures：[`analysis/apk/tools-and-failures.md`](../../../../analysis/apk/tools-and-failures.md)
- Signal：[`../signals/mitm-route-column-signals.md`](../signals/mitm-route-column-signals.md)
- Lessons：`feedback/history/apk-analysis/local-proxy/2026-07-14_095200-*`、`2026-07-14_095210-*`

## Token 影響

低。僅在 MITM／proxy 動態窗 lazy-load，約 180–220 tokens。
