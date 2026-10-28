# 跨目標採集方法（Collection）

本目錄保存 **HOW TO COLLECT** 的跨目標採集紀律與已保存抓包審查法。抽象自授權業務專案的採集 SOP，去掉產品／引擎專名後仍適用 APK、desktop、native、已保存 HAR/pcap。

它不取代各域深度 procedure（`analysis/apk/workflows/`、`analysis/desktop/`）；也不重複 [`../evidence-contract.md`](../evidence-contract.md) 的主張品質語義。

## 何時讀哪份

| 情境 | 讀 |
| --- | --- |
| 要開始 runtime／靜態／資源採集 | [`collection-discipline.md`](collection-discipline.md) |
| 手上已有 HAR／pcap／inventory／錄影 | [`capture-review.md`](capture-review.md) |
| 採集卡住、假缺失、反覆同症狀 | [`collection-pitfalls.md`](collection-pitfalls.md) |

調查節奏仍走 [`../investigation-process.md`](../investigation-process.md)；本目錄補強 §4 Collect。

## 與其他契約的邊界

| 問題 | 文件 |
| --- | --- |
| 這條主張是 observation／inference／unknown？ | [`../evidence-contract.md`](../evidence-contract.md) |
| 權威／新鮮度／可觀測性是否夠寫進報告？ | [`../../../enforcement/evidence-hierarchy.md`](../../../enforcement/evidence-hierarchy.md) |
| APK 動態 pcap→hook 層次筆記 | [`../../../workflow/apk-analysis/artifact-gates/evidence-chain.md`](../../../workflow/apk-analysis/artifact-gates/evidence-chain.md) |
| Frida／Unity／local-proxy 操作步驟 | [`../../apk/workflows/`](../../apk/workflows/README.md) |
| 授權／去敏 | [`../../../enforcement/authorization-scope.md`](../../../enforcement/authorization-scope.md)、[`../../../enforcement/sanitization.md`](../../../enforcement/sanitization.md) |

## 不放什麼

- Raw HAR、pcap、token、host、私人路徑、產品特異 opcode／tag。
- 引擎專屬 hook 配方（Frida IL2CPP、UnityPy 切圖等）。
- 「安裝某 MCP／IDE」當完成條件。

← [回到 reverse-engineering/](../README.md)
