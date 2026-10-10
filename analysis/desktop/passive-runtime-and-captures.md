# Passive runtime 與已保存抓包

## Passive browser／Electron

僅觀察使用者已開啟、已授權的目標。不宣稱工具代為 click／navigate／eval／invoke IPC。記錄結構、腳本、網路形狀、可選截圖；刻意排除 credentials、cookie、raw body（依引擎契約）。

靜態 graph 與 runtime 對照時：digest 不一致標 mismatch；路徑對應是 inference。

## HAR／已保存 capture

檢視 requests／responses／exposed payloads／source locations。不把 historical capture 當成對 live 系統的新授權。mitmproxy 原生格式若需轉檔，記工具邊界。

## 失敗

| 現象 | 處理 |
| --- | --- |
| CDP 端點不可達 | 記 environment unknown |
| 截斷／scope filter | coverage: partial |
| 過大 Evidence | 用摘要／分頁投影，保留原 digest |

← [desktop/](README.md)
