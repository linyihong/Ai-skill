# Shipped-Artifact 調查流程（工具中立）

本文件是跨目標 reverse 的**怎麼想、怎麼走**，不綁定任何 MCP／廠商 CLI。通用工具（jadx、Ghidra、Frida、解 ASAR…）只是實作手段；缺某一工具時換同等權威，流程不變。

靈感來自「先分流目標 → 帶證據追功能 → 分開觀察／推論／未知 → 需要時再升級權威」這類 agent 調查型流程；**Ai-skill 的 canonical 步驟以本檔與本庫 workflow 為準**。

## 與既有能力的關係

| 我們已強 | 本流程補上 |
| --- | --- |
| APK 動態：traffic triage、Frida、Flutter／Unity | 跨目標靜態起步、Evidence／unknown 紀律、功能追蹤閉環 |
| APK workflow 端到端 | 非 APK（native／Electron…）同一套「調查節奏」 |
| 授權／去敏 | 每步仍強制；流程不授權就停 |

## 七步節奏

```text
1 Authorize → 2 Route family → 3 Seed feature
    → 4 Collect evidence → 5 Separate claims
    → 6 Escalate authority OR close
    → 7 Handoff / reconstruct (optional)
```

### 1. Authorize（blocking）

具名標的 + 允許操作（只靜態／可觀察 runtime／可 hook）。未授權停止。見 [`../../enforcement/authorization-scope.md`](../../enforcement/authorization-scope.md)。

### 2. Route family

用 [`target-routing.md`](target-routing.md) 選分析族。
有完整原始碼且任務是維護／架構 → **離開**本流程，走 [`../repo/`](../repo/README.md)。

### 3. Seed the feature

把使用者問題收成可追的 seed（擇一或多）：

| Seed 類型 | 例子 |
| --- | --- |
| UI／產品語意 | 「搜尋」「剪貼簿」「冷啟動進某頁」 |
| 字串／符號 | 錯誤訊息、IPC channel 字面量、export 名 |
| 結構入口 | Activity／主視窗／已知 DLL export |
| 對照物 | 兩版發行物、前後 HAR |

沒有 seed 就先做 inventory（manifest／模組表／overview），再問使用者收斂——不要廣鉤一切。

### 4. Collect evidence（由近到遠）

預設**靜態先、動態後**（動態成本與授權更高）：

1. Inventory／manifest／package graph
2. 定位候選（搜 class／符號／模組）
3. 讀 body／偽碼／IPC 邊界
4. 靜態 refs／callers（標明非 runtime）
5. 僅在 unknown **blocking** 且已授權時：pcap／hook／passive observe／有界 process

工具表見各域 `tools-and-failures`；換工具不換步驟順序。

### 5. Separate claims

每條結論標 [`evidence-contract.md`](evidence-contract.md) 三類：

- observation
- inference
- residual unknown（寫 `needed_authority`）

禁止：偽碼當原始碼；靜態 xref 當已執行呼叫；「反編譯完整」當語意完整。

### 6. Escalate or close

| 情況 | 動作 |
| --- | --- |
| Unknown 不擋 goal | 標 `accepted_gap` 或 open，可關閉本輪 |
| Unknown 擋 goal且權威未用 | 升級權威（例如 Static APK → Dynamic APK） |
| 缺工具 | `engine_unavailable` → 換同等 CLI／手動；不假裝完成 |
| 要宣稱重建／對等 | 只對**有限可驗證主張**關閉；見下節 |

### 7. Optional：compare / reconstruct

- **版本對照**：兩邊覆蓋完整才可稱 added／removed；否則 unknown。
- **重建**：只驗證有限義務（行為／fixture／對照），不宣稱全域等價。
- APK 動態重建交接仍走 [`../../workflow/apk-analysis/`](../../workflow/apk-analysis/README.md) artifact gates。

## 與 REA 類 prompt 意圖的對照（非依賴）

| 調查意圖 | 本庫落點 |
| --- | --- |
| 追一個功能到證據 | 本檔 §3–5 + 域方法 |
| 兩版對照 | §7 + desktop／binary compare 思路 |
| 清點殘餘未知 | evidence-contract residual unknown |
| 有界 process／runtime | desktop passive／binary 有界動態；apk Frida workflows |
| 一般 source repo | **不做**；repo analysis |

不要求安裝任何外部 MCP。若環境已有橋接工具，僅作 §4 的可選實作，見 [`../../ai-tools/rea-mcp.md`](../../ai-tools/rea-mcp.md)。

## 最小完成檢查

- [ ] 授權已記錄
- [ ] 分析族已選
- [ ] seed 明確
- [ ] findings 已分 observation／inference／unknown
- [ ] blocking unknown 已升級或明示 accepted_gap
- [ ] 未把靜態結果宣稱成動態覆蓋

編排入口：[`../../workflow/reverse-engineering/execution-flow.md`](../../workflow/reverse-engineering/execution-flow.md)。

← [reverse-engineering/](README.md)
