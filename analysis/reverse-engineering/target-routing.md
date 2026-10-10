# Reverse-Engineering 目標分流

在選工具或開 workflow 之前，先用本表決定**分析族**。選錯族會把 scraping 當 reverse、或把動態 APK 問題丟給只做靜態反編譯的引擎。

## 決策表

| 使用者標的／問題 | 先走 | 主要方法入口 | 常見下一步 |
| --- | --- | --- | --- |
| 授權 APK：class／method／manifest／靜態 refs | **Static APK** | [`../apk/`](../apk/README.md)（靜態 JADX 路徑落地中） | unknowns 涉及 runtime → Dynamic APK |
| 授權 APK：流量、hook、Flutter／Unity、local proxy | **Dynamic APK** | [`../apk/traffic-triage.md`](../apk/traffic-triage.md)、`workflows/` | 靜態只做輔助定位 |
| Native 可執行檔／`.so`／偽碼／xref | **Native binary** | [`../binary/`](../binary/README.md)（落地中） | 需要實際呼叫觀測 → 授權下的 process／debugger 路徑 |
| Linux ELF layout／core crash（非 live） | **Offline binary** | `analysis/binary/` offline 文件（落地中） | 不授權 live attach |
| .NET／managed／NativeAOT | **Managed** | `analysis/binary/managed-code.md`（落地中） | native 橋接用 binary 方法 |
| ASAR／Electron／打包 JS 樹：模組、IPC、source map | **Desktop JS** | [`../desktop/`](../desktop/README.md)（落地中） | 需要頁面行為 → passive runtime 或 web 工具 |
| 網站爬取／anti-bot／selector | **Web scrape** | [`../web/`](../web/README.md) | **不是** desktop reverse |
| HAR／已保存抓包 | **Capture review** | desktop／apk traffic 交叉；通用 HAR 方法落地中 | 不把 capture 當 live 授權擴大 |
| Firmware 映像 | **Firmware** | [`../firmware/`](../firmware/README.md)（stub） | 抽出後可能 handoff Native |
| EVM bytecode（離線 hex／raw） | **EVM** | [`../evm/`](../evm/README.md)（stub） | 不做鏈上查詢除非另授權 |
| 完整原始碼 repo 架構 | **Repo analysis** | [`../repo/`](../repo/README.md) | **跳過** reverse-engineering 入口 |

## 互斥與交接

1. **Static APK 與 Dynamic APK 互補，不互斥。** 靜態引擎若明確不執行 APK、不做 native-lib／runtime capture，不得宣稱已完成動態目標；應留下 residual unknown 並 handoff。
2. **Desktop JS ≠ Web scrape。** 問的是「發行物裡功能怎麼接」，走 desktop；問的是「網站怎麼抓資料」，走 web。
3. **有原始碼且任務是維護／架構** → repo；只有在結論依賴發行物／反編譯時才進本分流。
4. **授權失敗** → 停止；見 [`../../enforcement/authorization-scope.md`](../../enforcement/authorization-scope.md)。

## Evidence

分流結果本身應記一筆 observation：「選定族 X，因為標的型態 Y／問題 Z」。後續 findings 遵守 [`evidence-contract.md`](evidence-contract.md)。

## 可選引擎對照（非唯一路徑）

| 分析族 | 通用工具思路 | 可選本地 MCP／CLI 橋接 |
| --- | --- | --- |
| Static APK | jadx、apktool、aapt | Android inspect／search／method／refs 類工具（若 advertised） |
| Native | Ghidra／IDA／Hopper／radare2 | open_binary／overview／decompile 類工具（若 advertised） |
| Desktop JS | 解 ASAR、靜態 graph、source map | analyze_javascript_application 類工具（若 advertised） |
| Dynamic APK | Frida、pcap、MITM | **通常不走**靜態專用 MCP；用本庫 apk workflows |

工具名以當時環境實際 catalog 為準；本表只鎖能力類別。

← [回到 reverse-engineering/](README.md)
