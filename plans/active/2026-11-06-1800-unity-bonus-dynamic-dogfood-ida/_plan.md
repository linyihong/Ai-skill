---
id: 2026-11-06-1800-unity-bonus-dynamic-dogfood-ida
plan_kind: main
status: draft
owner: unity-analysis-maintainer
created: 2026-11-06
updated: 2026-11-09
parent: null
owner_layer: analysis
---

# Bonus 動態流程 dogfood 與 Ghidra 接入驗證

狀態為 **draft，Ghidra 靜態基線已可重跑，第一個 Runtime 預測獲自然觀測支持，相鄰格移動與到站回放通過，特殊格獎勵鏈與取消待驗**。

## 已接受審查與階段提升條件

採納 2026-11-06 審查，計畫維持 draft。依序檢查工具環境、建立可重現靜態基線、取得不介入的自然動態流程，再做少量預測的交叉驗證。只有環境確認、第一批 Ghidra 匯出可重現，且至少一項預先記錄的預測經 Runtime 支持後，才考慮提升階段；這不取代兩項最終 oracle。

靜態觀察必須先轉為預測與區分條件，再決定 hook。預測不符時保存差異與原始觀測，不改寫觀測以配合反編譯。環境檢查結果見 [preflight](evidence/environment-preflight-20261106.md)。

2026-11-09 決定改用 Ghidra 進行本機 ARM64 靜態分析，搭配 Il2CppDumper、LLVM 與 Frida。計畫路徑沿用原識別碼，兩項驗收仍分開裁決。

2026-11-09 已取得 [自然 Runtime 對照](evidence/natural-runtime-20261109.md)。B 的首案預測驗收通過，A 已取得自然樣本，並完成相鄰格回放的獨立位置讀回、重複與同場景重進。普通格與特殊格的到站回放亦已通過，特殊格獎勵鏈及取消仍待驗。保留 draft，待整體階段審查。

## Goal 與延續關係

在已通過靜態畫面回放的案例上，驗證棋子移動、到站事件與下一步流程，並分別評估 Unity 決策流程與 Ghidra 靜態證據接入的效用。

基線為 [已結案的決策域計畫](../../archived/2026-10-29-1055-unity-analysis-decision-domain/_plan.md) 與 [範圍結案裁決](../../archived/2026-10-29-1055-unity-analysis-decision-domain/evidence/phase2-final-closeout-20261106.md)。此為同目標的第二個案例，不視為跨專案重用。

## Scope／Non-goals

- In: 自然事件與動態位置關聯、授權 lab 回放、事件完成／取消與重進驗證、Ghidra 方法索引與有限匯出。
- Out: RNG／RTP／伺服器生成公式、餘額或派彩偽造、全遊戲閉鎖、自動 patch、MCP 導入與通用 orchestrator。
- 靜態回放既有 PASS 不代表棋子移動、特殊格或結算流程 PASS。

## Architecture Compatibility 與 SoT

沿用 `analysis/unity/decision-workflow.md`，動態取證見 [動態流程程序](../../../analysis/unity/workflows/unity-dynamic-flow-validation.md)。端到端編排維持 APK workflow。

Ghidra／GhidraScript 為靜態證據來源，Runtime 為實際執行來源；兩者不取代 case ledger 的裁決。業務專案保留 binary、metadata、Ghidra 專案、原始 trace、fixture、裝置及截圖。本庫只保存方法、計畫與去敏 verdict。

## Evidence Contract

每次記錄 build／binary 與 metadata fingerprint、方法 RVA、物件及 cell ID 的有效窗口、session／phase／時間與觀測範圍。反編譯推論、指令觀察、Runtime 觀察與剩餘未知分欄。靜態 XREF 不證明實際執行順序。

## 兩項獨立驗收

| Track | PASS 條件 | 不足時 |
| --- | --- | --- |
| A 動態 dogfood | 有自然對照，已選路徑從輸入／起點、移動中、終點、到站事件到完成均可關聯；明確 lab 回放符合獨立位置及時序 oracle，延後讀回無回彈或重複完成 | blocked 或 partial，列明缺哪個窗口；只觀察到自然移動不算回放 PASS |
| B Ghidra 接入 | 同版 ARM64 工具實際可用；至少一條移動或事件路徑的函式／XREF／欄位存取可重跑匯出，形成預先寫下預測的 hook 候選，Runtime 能支持或反駁該預測 | 工具不可用或缺 Runtime 對照，記錄該項缺口與來源 |

兩項均有 verdict 才做整體結案。blocked 是有效的分析出口，但不等於該 track 的成功。任一項未 PASS 時，整體不得標示雙項成功。

## Phases

### Phase 0 — 基線與工具 preflight

- [x] 寫明基線、兩項 oracle、責任邊界與 deferred 工作。
- [x] 補上動態驗證程序與既有入口連結。
- [ ] 核對目前版本、fixture、乾淨可回復場景及同版 binary／metadata。
- [x] 確認 Ghidra、JDK 與原生 ARM64 反編譯器可用，見 [靜態基線](evidence/ghidra-static-baseline-20261109.md)。
- [ ] 在專案建立六欄 case ledger，列主要假說、可區分預測與回復方式。

### Phase 1 — 自然動態對照

- [x] 進場前掛好移動、停止、到站與位置收尾觀測，保存自然進盤至結算。取消分支另列待測。
- [x] 取得普通移動與 +1 特殊格樣本，記錄目標身分、事件順序與獨立畫面。完整逐段座標與步數對齊仍待核對。
- [ ] 驗證邏輯格與畫面座標映射、物件生命週期、完成次數及獨立讀回。尚未捕獲的分支保持 not-tested。

### Phase 2 — Ghidra 靜態證據與 Runtime 對照

- [ ] 建立 metadata method → RVA → Ghidra address → Runtime module-relative address 索引，核對 executable segment 與入口。
- [ ] 對移動入口、到站 dispatch、完成回呼匯出有限 caller／callee 與欄位讀寫；間接呼叫與反編譯失敗顯示 unknown。
- [ ] 固定 binary／工具／目標清單重跑兩次，正規化後比較；人工抽查一條 XREF 與一處狀態寫入。
- [x] 以預先寫下的兩個假說做被動 Runtime 對照，裁決樣本內停止與中途移動的事件差異。

### Phase 3 — 授權 lab 動態回放

[相鄰格回放證據](evidence/dynamic-replay-20261109.md) 覆蓋同一行程的四次中途移動。此樣本涵蓋中途移動。另見 [到站回放證據](evidence/stop-replay-20261109.md)，覆蓋普通格及特殊格的原生 Stop 到站。

- [ ] 以明確 fixture／輸入啟動原生移動流程，記修改前狀態；Unity 操作採已驗證主執行緒契約。
- [ ] 用另一個唯讀取證來源與時間序列畫面驗證中途移動、終點、到站事件及恰好一次完成。
- [ ] 普通路徑與所選特殊／邊界路徑分別通過，特殊格的事件語意以自然證據為準。
- [ ] 重複回放、離開重進及取消／失效物件至少各驗證一次；無法安全執行則保留缺口，不擴大 PASS 範圍。

### Phase 4 — 裁決與收尾

- [ ] 分別記 A／B verdict、有效範圍、剩餘未知與反證。
- [ ] 分開評估決策流程效用與 Ghidra 產出價值，不把安裝成功當分析效用。
- [ ] 新可重用發現經抽象化後回饋，單目標方法仍 candidate。
- [ ] 更新專案待辦、證據索引與 Plan Completion Closure，驗證、提交、推送並讀回。

## Runtime Execution Path

本計畫是既有分析流程的實測，入口為既有 Unity routing → decision-workflow → dynamic-flow procedure → 專案 case ledger。不增加 runtime schema、route 或機械 gate，不將 Ghidra 專案 設為 canonical SoT。

## Validation Strategy

文件階段核對連結、plan tree、runtime index freshness。實測階段保存自然與 lab 對照、獨立位置讀回及時間序列畫面；Ghidra 另驗版本、位址與可重跑匯出。不得用單一 invoke 回傳或同一探針自行寫下的終點作完成證明。

## Open Questions／停止條件

- IDA 接入另行保留，待具備適用 ARM64 授權再評估。
- 特殊格選擇依自然可取得窗口決定，先記選樣理由與未覆蓋分支。
- 連續兩輪無法區分假說時重新設計；缺自然窗口或可信物件關聯則 blocked。
- 不將為取得 Bonus 的等待當成無限自動下注授權，沿用專案現有投注與裝置控制規範。

## Deferred 與責任

負責方為 Unity 分析維護者。唯讀 MCP、跨專案重用與 ADR 仍 deferred；僅於本輪證據揭示明確需求或取得獨立目標重用後重啟。本計畫不重新開啟已結案基線的完成狀態。

## Evidence

[證據索引](evidence/README.md)。靜態工具已通過樣本 smoke、重跑比較及自然 Runtime 預測。動態 dogfood 的中途移動回放已通過。普通格及特殊格的到站事件已通過。特殊格獎勵鏈、取消與跨行程重進仍待測。
