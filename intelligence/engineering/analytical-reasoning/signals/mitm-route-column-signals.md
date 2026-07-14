# MITM Route-Column Signals（MITM 三欄訊號）

## 問題

動態 MITM 窗要用哪些觀測訊號，才能分開「導流成功／第一方 trust 失敗／業務 API bypass」？

## 判斷信號

### 欄 A — Proxy-aware stack

| 信號 | 檢查方式 | 可信度 |
| --- | --- | --- |
| 廣告／mediation／analytics CONNECT + HTTP 200 解密 | mitmdump／Proxyman flow list | 高 |
| `settings get global http_proxy` 非 `:0` | adb | 高（設定層） |

### 欄 B — Proxy-entered first-party trust fail

| 信號 | 檢查方式 | 可信度 |
| --- | --- | --- |
| MITM：`Client TLS handshake failed` / `certificate unknown` 且 host 為第一方 | mitmdump log | 高 |
| **裝置 UI：SSL certificate verification／憑證錯誤頁**（代理開啟時） | 使用者／截圖觀察 | 高（可見 trust 失敗） |
| 同一 host 關代理後可開、開代理即 SSL UI | A/B 短窗 | 高 |

### 欄 C — Business API proxy membership

| 信號 | 檢查方式 | 可信度 |
| --- | --- | --- |
| 靜態／已知業務 host 在 MITM **零 CONNECT** | flow／log 對照 | 高 |
| 關代理後 short-window `tcpdump`／pcap 仍見業務候選 dst IP:443 | adb + dig 對照 | 高（bypass 或直連仍活） |
| 業務 host CONNECT + handshake fail | mitmdump | 高（改標 pinning-tier，不是 bypass） |
| Frida：OkHttp／Retrofit 打到業務 URL，且 Cronet 業務計數≈0 | in-process URL hook | 高（棧＝OkHttp） |
| Frida／靜態：`OkHttpClient.Builder` 上 **no-proxy／NO_PROXY／\*ProxyConfig\*apply*(…, true)** | builder hook | 高（bypass 機制＝顯式 no-proxy） |

### 排除／防呆

| 誤讀 | 正確 |
| --- | --- |
| UI SSL → App 壞掉／不要開代理 | UI SSL = 欄 B 證據 |
| 有廣告解密 → 主線 MITM 完成 | 只證明欄 A |
| 無業務 CONNECT → pinning | 先欄 C；零 CONNECT = bypass 候選 |
| 零 CONNECT → 一定是 Cronet | Frida 若證 OkHttp，查 **顯式 no-proxy** |

## 判斷流程

```text
MITM 冷啟動窗
  → 有任何可解密流量？否 → 查 proxy 設定／Wi‑Fi
  → 是 → 記欄 A
  → 第一方 handshake fail 或 UI SSL？是 → 記欄 B（分 host）
  → 主業務 host 有 CONNECT？
        是 + handshake fail → 欄 C pinning-tier
        是 + 解密成功 → MITM 可繼續業務
        否 → no-proxy pcap？
              有業務 IP → 欄 C bypass → Frida／in-process
                    → OkHttp？查 builder no-proxy → 標 okhttp-no-proxy
                    → Cronet／native？走對應 hook
              無 → 查流程未觸發／DNS／權限
```

## 相關 atoms

- [`../heuristics/mitm-route-column-diagnosis.md`](../heuristics/mitm-route-column-diagnosis.md)
- [`local-proxy-detection.md`](local-proxy-detection.md)（loopback／TUN 另線；與本三欄可並存）
- [`analysis/apk/traffic-triage.md`](../../../../analysis/apk/traffic-triage.md)
- [`feedback/history/apk-analysis/local-proxy/2026-07-14_105100-column-c-bypass-may-be-okhttp-no-proxy-not-cronet.md`](../../../../feedback/history/apk-analysis/local-proxy/2026-07-14_105100-column-c-bypass-may-be-okhttp-no-proxy-not-cronet.md)
