> 遵守 [共用規則索引](../../../../enforcement/README.md)、[sanitization](../../../../enforcement/sanitization.md)、[reusable-guidance-boundary](../../../../enforcement/reusable-guidance-boundary.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson。

### 2026-07-14 - Mixed MITM: ads OK, first-party UI SSL fail, API proxy-bypass

Status: candidate

#### One-line Summary

同一冷啟動窗可同時成立三件事：第三方／廣告 **MITM 可解密**、某一第一方 host **進代理但 UI+TLS 憑證失敗**、**主業務 API 從不進代理**（僅 no-proxy pcap 可見）——必須分三欄報告，不可合成單一「pinning」結論。

#### Human Explanation

延續「代理設定 ≠ 業務路由」與「MITM 有 CDN 無 API → pinning tier」。更新邊界：第一方不一定只有「API pinning」一種失敗。實測混合象限：

| 路由 | MITM | 其他證據 |
| --- | --- | --- |
| Ads / analytics / mediation | CONNECT + decrypt OK | — |
| First-party H5／mobile web host | CONNECT + `certificate unknown`；裝置 UI SSL 提示 | Trust／pinning on that host |
| Primary business API | **Zero CONNECT** | no-proxy tcpdump/SNI 仍見候選 API IP |

錯判模式：

- 見 UI SSL → 關掉代理當修好，漏記 trust 證據。
- 見大量廣告解密 → 以為「MITM 已覆蓋主線」，乾等靜態 catalog 裡的業務 path 字樣。
- 見 MITM 無 API → 直接寫 pinning，其實是 **bypass**（應用 Frida／Cronet／native 路徑，而非再拉長純 MITM）。

#### Trigger

- Cold-start MITM 充滿廣告 SDK，零業務 path。
- 使用者看到 App 內 SSL certificate verification UI。
- Clearing proxy makes App usable again while pcap still shows business IP.

#### Evidence

- Tool: mitmdump + global http_proxy + optional rooted `tcpdump` without proxy.
- Sanitized outcome columns as in the table above.
- Project evidence: `<PROJECT_ROOT>/api/dynamic-*.md`（hosts/paths stay project-local）。

#### Generalized Lesson

Report three independent judgments per window:

1. **Proxy-aware stack health**（ads decrypt?）
2. **Proxy-entered first-party trust**（UI SSL / handshake fail?）
3. **Business API proxy membership**（CONNECT present? if no → bypass confirm via no-proxy SNI）

Only when (3) is CONNECT+handshake-fail may you label that API as pinning-tier for MITM.

#### Agent Action

1. Never wait longer on pure MITM after (1) OK +(3) absent.
2. Log UI SSL as evidence for (2).
3. Prefer Frida/in-process for (3) bypass; keep MITM for (1)/(2) only.

#### Goal / Action / Validation

- Goal: Prevent collapsed “everything is pinning” conclusions.
- Validation: Same timestamped window documents all three columns with distinct evidence sources.

#### Applies When

Android apps mixing WebView/H5, ad SDKs, and a separate primary API client.

#### Does Not Apply When

Single-stack apps where all traffic shares one HTTP client and proxy policy.

#### Promotion Target

`analysis/apk/traffic-triage.md` — candidate；links to existing proxy-config / pinning-tier lessons.
