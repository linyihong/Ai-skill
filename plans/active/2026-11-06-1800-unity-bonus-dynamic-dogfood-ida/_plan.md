---
id: 2026-11-06-1800-unity-bonus-dynamic-dogfood-ida
plan_kind: main
status: draft
owner: unity-analysis-maintainer
created: 2026-11-06
updated: 2026-11-06
parent: null
owner_layer: analysis
---

# Bonus 動態流程 dogfood 與 IDA 接入驗證

狀態為 **draft，計畫與 workflow 補強已完成，動態 dogfood／IDA 實測尚未開始**。

## 已接受審查與階段提升條件

採納 2026-11-06 審查，計畫維持 draft。依序檢查工具環境、建立可重現靜態基線、取得不介入的自然動態流程，再做少量預測的交叉驗證。只有環境確認、第一批 IDA 匯出可重現，且至少一項預先記錄的預測經 Runtime 支持後，才考慮提升階段；這不取代兩項最終 oracle。

靜態觀察必須先轉為預測與區分條件，再決定 hook。預測不符時保存差異與原始觀測，不改寫觀測以配合反編譯。環境檢查結果見 [preflight](evidence/environment-preflight-20261106.md)。

## Goal 與延續關係

在已通過靜態畫面回放的案例上，驗證棋子移動、到站事件與下一步流程，並分別評估 Unity 決策流程與 IDA 靜態證據接入的效用。

基線為 [已結案的決策域計畫](../../archived/2026-10-29-1055-unity-analysis-decision-domain/_plan.md) 與 [範圍結案裁決](../../archived/2026-10-29-1055-unity-analysis-decision-domain/evidence/phase2-final-closeout-20261106.md)。此為同目標的第二個案例，不視為跨專案重用。

## Scope／Non-goals

- In: 自然事件與動態位置關聯、授權 lab 回放、事件完成／取消與重進驗證、IDA 方法索引與有限匯出。
- Out: RNG／RTP／伺服器生成公式、餘額或派彩偽造、全遊戲閉鎖、自動 patch、MCP 導入與通用 orchestrator。
- 靜態回放既有 PASS 不代表棋子移動、特殊格或結算流程 PASS。

## Architecture Compatibility 與 SoT

沿用 `analysis/unity/decision-workflow.md`，動態取證見 [動態流程程序](../../../analysis/unity/workflows/unity-dynamic-flow-validation.md)。端到端編排維持 APK workflow。

IDA／IDAPython 為靜態證據來源，Runtime 為實際執行來源；兩者不取代 case ledger 的裁決。業務專案保留 binary、metadata、IDB、原始 trace、fixture、裝置及截圖。本庫只保存方法、計畫與去敏 verdict。

## Evidence Contract

每次記錄 build／binary 與 metadata fingerprint、方法 RVA、物件及 cell ID 的有效窗口、session／phase／時間與觀測範圍。反編譯推論、指令觀察、Runtime 觀察與剩餘未知分欄。靜態 XREF 不證明實際執行順序。

## 兩項獨立驗收

| Track | PASS 條件 | 不足時 |
| --- | --- | --- |
| A 動態 dogfood | 有自然對照，已選路徑從輸入／起點、移動中、終點、到站事件到完成均可關聯；明確 lab 回放符合獨立位置及時序 oracle，延後讀回無回彈或重複完成 | blocked 或 partial，列明缺哪個窗口；只觀察到自然移動不算回放 PASS |
| B IDA 接入 | 同版 ARM64 工具實際可用；至少一條移動或事件路徑的函式／XREF／欄位存取可重跑匯出，形成預先寫下預測的 hook 候選，Runtime 能支持或反駁該預測 | IDA 不可用則 blocked；其他靜態工具結果單列，不冒稱 IDA PASS |

兩項均有 verdict 才做整體結案。blocked 是有效的分析出口，但不等於該 track 的成功。任一項未 PASS 時，整體不得標示雙項成功。

## Phases

### Phase 0 — 基線與工具 preflight

- [x] 寫明基線、兩項 oracle、責任邊界與 deferred 工作。
- [x] 補上動態驗證程序與既有入口連結。
- [ ] 核對目前版本、fixture、乾淨可回復場景及同版 binary／metadata。
- [ ] 確認授權 IDA／ARM64 Hex-Rays、IDAPython 可用，記版本及實際反編譯 smoke；未具備即記 B blocked，不購買或假定工具可用。
- [ ] 在專案建立六欄 case ledger，列主要假說、可區分預測與回復方式。

### Phase 1 — 自然動態對照

- [ ] 進入稀有窗口前掛好輸入事件、移動入口、位置更新、到站與完成／取消觀測；保存整段 phase，避免只取單次 hook。
- [ ] 取得一條普通移動，以及一條特殊格／邊界路徑。記起點、目標／步數、執行中位置、終點、事件與完成順序。
- [ ] 驗證邏輯格與畫面座標映射、物件生命週期、完成次數及獨立讀回。尚未捕獲的分支保持 not-tested。

### Phase 2 — IDA 靜態證據與 Runtime 對照

- [ ] 建立 metadata method → RVA → IDA EA → Runtime module-relative address 索引，核對 executable segment 與入口。
- [ ] 對移動入口、到站 dispatch、完成回呼匯出有限 caller／callee 與欄位讀寫；間接呼叫與反編譯失敗顯示 unknown。
- [ ] 固定 binary／工具／目標清單重跑兩次，正規化後比較；人工抽查一條 XREF 與一處狀態寫入。
- [ ] 依靜態結果選最小被動 hook，先寫兩假說預測，再用 Runtime 裁決；找不到直接路徑也需保存可反駁結果。

### Phase 3 — 授權 lab 動態回放

- [ ] 以明確 fixture／輸入啟動原生移動流程，記修改前狀態；Unity 操作採已驗證主執行緒契約。
- [ ] 用另一個唯讀取證來源與時間序列畫面驗證中途移動、終點、到站事件及恰好一次完成。
- [ ] 普通路徑與所選特殊／邊界路徑分別通過，特殊格的事件語意以自然證據為準。
- [ ] 重複回放、離開重進及取消／失效物件至少各驗證一次；無法安全執行則保留缺口，不擴大 PASS 範圍。

### Phase 4 — 裁決與收尾

- [ ] 分別記 A／B verdict、有效範圍、剩餘未知與反證。
- [ ] 分開評估決策流程效用與 IDA 產出價值，不把安裝成功當分析效用。
- [ ] 新可重用發現經抽象化後回饋，單目標方法仍 candidate。
- [ ] 更新專案待辦、證據索引與 Plan Completion Closure，驗證、提交、推送並讀回。

## Runtime Execution Path

本計畫是既有分析流程的實測，入口為既有 Unity routing → decision-workflow → dynamic-flow procedure → 專案 case ledger。不增加 runtime schema、route 或機械 gate，不將 IDA database 設為 canonical SoT。

## Validation Strategy

文件階段核對連結、plan tree、runtime index freshness。實測階段保存自然與 lab 對照、獨立位置讀回及時間序列畫面；IDA 另驗版本、位址與可重跑匯出。不得用單一 invoke 回傳或同一探針自行寫下的終點作完成證明。

## Open Questions／停止條件

- IDA／Hex-Rays 授權與安裝目前 unknown，Phase 0 先查。
- 特殊格選擇依自然可取得窗口決定，先記選樣理由與未覆蓋分支。
- 連續兩輪無法區分假說時重新設計；缺自然窗口或可信物件關聯則 blocked。
- 不將為取得 Bonus 的等待當成無限自動下注授權，沿用專案現有投注與裝置控制規範。

## Deferred 與責任

負責方為 Unity 分析維護者。唯讀 MCP、跨專案重用與 ADR 仍 deferred；僅於本輪證據揭示明確需求或取得獨立目標重用後重啟。本計畫不重新開啟已結案基線的完成狀態。

## Evidence

[證據索引](evidence/README.md)。目前只有計畫與文件驗證，沒有動態或 IDA 成功證據。
