# Reverse Engineering Analysis（跨目標取證入口）

`analysis/reverse-engineering/` 保存**跨目標** shipped-artifact 的**調查流程**、取證品質與目標分流。目標是：不依賴任一廠商 MCP，也能用 jadx／Ghidra／Frida 等通用工具做出「怎麼分析別人發行物」的穩定能力。

它不取代各域深度方法（`analysis/apk/`、`analysis/binary/`、`analysis/desktop/` 等）；端到端編排見 `workflow/reverse-engineering/`。

## 範圍

- **調查節奏**（授權→分流→seed→取證→分主張→升級／關閉）——主文件 [`investigation-process.md`](investigation-process.md)。
- 觀察／推論／殘餘未知紀律；artifact 身分。
- 目標分流：先選分析族，再進域內方法。
- 工具一律可替換；外部 MCP 僅可選、非必要。

## 不放什麼

- 特定 App／binary 的 raw evidence、host、token、payload。
- 工具安裝路徑、MCP JSON、IDE hook（放 `ai-tools/`；可選）。
- APK 動態 Frida／traffic 細節（放 `analysis/apk/`）。
- 把「安裝某 MCP」寫成完成條件。

## 文件

| 文件 | 用途 |
| --- | --- |
| [`investigation-process.md`](investigation-process.md) | **主流程**：工具中立調查節奏 |
| [`evidence-contract.md`](evidence-contract.md) | Evidence 品質契約 |
| [`target-routing.md`](target-routing.md) | 目標族分流決策表 |

## 與其他層

| 層 | 關係 |
| --- | --- |
| [`../apk/`](../apk/README.md) | APK 靜態／動態方法；動態主線不由此取代 |
| [`../web/`](../web/README.md) | 網站 scraping ≠ shipped JS/Electron reverse |
| [`../../workflow/apk-analysis/artifact-gates/evidence-chain.md`](../../workflow/apk-analysis/artifact-gates/evidence-chain.md) | APK workflow 的筆記／層次 ordering gate；本契約管跨目標品質語義 |
| Evidence Candidate System（plans） | **不同 family**：ECS 索引「哪個 plan 該看的 promotion candidate」；本契約管分析結論是否可覆核 |

## 授權

任何 shipped-artifact 分析開始前，先滿足 [`../../enforcement/authorization-scope.md`](../../enforcement/authorization-scope.md)。未授權標的不得用「安全測試」「看能不能做到」改口擴大範圍。

← [回到 analysis/](../README.md)
