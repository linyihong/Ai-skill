# Native 調查方法

授權 native 可執行檔／library 的靜態（與可選有界動態）取證。Evidence：[`../reverse-engineering/evidence-contract.md`](../reverse-engineering/evidence-contract.md)。

## 步驟

1. 確認授權；記錄 path、digest、架構、格式（Mach-O／ELF／PE）。
2. 開分析資料庫或載入目標（Ghidra／IDA／Hopper／radare2；可選 MCP `open_binary` 類）。
3. Overview：entry、匯出／匯入、字串熱點、節區。
4. 依功能字串／符號搜尋 → 反編譯／組語 → callers／callees／xref。
5. 每條結論標 observation／inference／unknown；偽碼 ≠ 原始碼。
6. Provider 不支援的架構／metadata → unknown，不填 false。

## 可選有界動態

僅在授權且需要「實際是否被呼叫」時：debugger／trace 須有時間與事件上限；hardened runtime／缺 entitlement 標 blocked。不得把「能 attach」當成授權擴大。

## 失敗判讀

| 現象 | 處理 |
| --- | --- |
| Provider 啟動超時 | 加大 timeout 或換 provider；保留 partial |
| 剝殼／加殼 | 分開 original vs derived artifact digest |
| 符號剝除 | 依字串／xref／行為推論，標 inference |

← [binary/](README.md)
