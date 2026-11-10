---
id: 2026-11-06-1800-unity-bonus-dynamic-dogfood-ida
plan_kind: main
status: completed
owner: unity-analysis-maintainer
created: 2026-11-06
updated: 2026-11-10
parent: null
owner_layer: analysis
---

# Bonus 動態流程 dogfood 與 Ghidra 接入驗證

**狀態 completed，限指定手機動態路徑及 Ghidra 有限接入。** 兩項各自 PASS。[最終驗收與交接](evidence/final-closeout-20261110.md) 為目前裁決，歷史 evidence 保留當時限制。採用 Ghidra 取代本輪 IDA 工具驗證，計畫 ID 沿用。

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

- [x] Phase 0。完成同版 binary／metadata、可回復場景與工具核對，專案保存六欄 case ledger。
- [x] Phase 1。取得同棋子自然移動、到站、獎勵與結算對照，使用獨立畫面和邏輯格讀回。
- [x] Phase 2。四方法有限匯出重跑一致，核對指令與位址，先寫預測再以自然 Runtime 裁決。
- [x] Phase 3。完成指定路徑手機重跑、逐格、逐手、結算及回主盤，另保留隔離中止、重進、立即重用反例。
- [x] Phase 4。分別記錄 A／B 裁決、更新可重用流程、專案待辦與索引，執行 Plan Completion Closure。

正式共用佇列取消與立即重用驗收 deferred。隔離中止的結果限獨立 action。完整的限制、責任及重啟條件見 [結案交接](evidence/final-closeout-20261110.md)。

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

[證據索引](evidence/README.md) · [最終裁決及驗收數據](evidence/final-closeout-20261110.md)
