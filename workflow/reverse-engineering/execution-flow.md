# Reverse-Engineering Execution Flow

本 workflow 執行 [`../../analysis/reverse-engineering/investigation-process.md`](../../analysis/reverse-engineering/investigation-process.md) 的七步節奏。**不要求**任何外部 MCP；通用反編譯／hook 工具即可。

## 1. 授權（blocking）

確認具名標的與授權範圍（[`../../enforcement/authorization-scope.md`](../../enforcement/authorization-scope.md)）。未授權則停止；不得以「安全測試」改口擴大。

記錄：標的型態、版本／digest（或 unknown）、允許操作（decompile／靜態 only／runtime observe／process）。

## 2. 目標分流

讀 [`../../analysis/reverse-engineering/target-routing.md`](../../analysis/reverse-engineering/target-routing.md)，選定分析族。

| 族 | 下一站 |
| --- | --- |
| Static APK | [`../../analysis/apk/static-jadx-path.md`](../../analysis/apk/static-jadx-path.md)；需要動態時 → [`../apk-analysis/execution-flow.md`](../apk-analysis/execution-flow.md) |
| Dynamic APK | [`../apk-analysis/execution-flow.md`](../apk-analysis/execution-flow.md) |
| Native／offline／managed | [`../../analysis/binary/`](../../analysis/binary/README.md) |
| Desktop JS／Electron | [`../../analysis/desktop/`](../../analysis/desktop/README.md) |
| Firmware／EVM | 對應 stub README |
| Web scrape | [`../../analysis/web/`](../../analysis/web/README.md)（離開本 workflow） |
| Source repo | [`../../analysis/repo/`](../../analysis/repo/README.md)（離開本 workflow） |

## 3. 取證（依調查流程 §3–5）

定 seed → 靜態先、動態後 → 分 observation／inference／unknown。域內方法見 apk／binary／desktop。缺某一工具時換同等權威，標 `engine_unavailable` 若完全無法取證。可選 MCP 橋接僅附錄：[`../../ai-tools/rea-mcp.md`](../../ai-tools/rea-mcp.md)。

## 4. 產出與關閉

通過 [`artifact-gates.md`](artifact-gates.md)。Raw evidence 留業務專案；可重用 lesson 去敏後進 `feedback/history/`。

殘餘 unknown 若需另一族（例如靜態→動態），明確 handoff 並帶 unknown 清單。
