# Passive runtime 與已保存抓包

跨目標 Capture review 節奏見 [`../reverse-engineering/collection/capture-review.md`](../reverse-engineering/collection/capture-review.md)。本檔補 **desktop／Electron／瀏覽器** 落地步驟。

## Passive browser／Electron

僅觀察使用者已開啟、已授權的目標。不宣稱工具代為 click／navigate／eval／invoke IPC。記錄結構、腳本、網路形狀、可選截圖；刻意排除 credentials、cookie、raw body（依引擎契約）。

靜態 graph 與 runtime 對照時：digest 不一致標 mismatch；路徑對應是 inference。

## HAR／已保存 capture（可執行短法）

1. **授權**：確認允許「讀此檔」；historical capture ≠ live 授權擴大。
2. **身分**：記 format（HAR／mitm dump）、大致時間、digest 或 `digest_unknown`。
3. **Inventory**：列出 method、去敏 host、path、status、mime；寫入專案筆記，不把 raw 入 Ai-skill。
4. **分層**：SNI／CONNECT／時序是傳輸外觀；只有明文才可寫 path／欄位；業務語義需對照 UI／seed。
5. **Coverage**：有 filter、截斷、過大省略時標 `coverage: partial`。
6. **轉檔**：mitmproxy 原生格式若需轉 HAR／JSON，記工具邊界。
7. **Fail-closed**：無明文（僅加密）→ 停；不發明 API；升級權威或 `accepted_gap`。
8. **交接**：需要 live 再抓時帶 unknown 清單 handoff；見 capture-review §7。

主張標籤遵守 [`../reverse-engineering/evidence-contract.md`](../reverse-engineering/evidence-contract.md)。

## 與 APK 即時流量

裝置側 pcap／MITM／proxy 分流仍用 [`../apk/traffic-triage.md`](../apk/traffic-triage.md)。本檔不取代該分流；兩邊都遵守 sanitized inventory 與 fail-closed。

## 失敗

| 現象 | 處理 |
| --- | --- |
| CDP 端點不可達 | 記 environment unknown |
| 截斷／scope filter | coverage: partial |
| 過大 Evidence | 用摘要／分頁投影，保留原 digest |
| 把 SNI 當 API | 改寫為傳輸外觀 observation；補明文或 unknown |

← [desktop/](README.md)
