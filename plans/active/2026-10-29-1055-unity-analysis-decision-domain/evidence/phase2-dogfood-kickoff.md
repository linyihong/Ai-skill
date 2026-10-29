# Phase 2 — 5×2 soft-enter dogfood kickoff

- **Opened**: 2026-10-29（stakeholder：Phase 1 結構性驗收通過 → 正式進入 Phase 2）
- **Architecture**: **不**再調 `analysis/unity` 整體架構；**不**做 Phase 3
- **Pass meaning**: 決策流程找出第一個有證據的必要差異，**或**明確 `blocked`（缺對照／缺證據／下一實驗）——**不是**畫面修復成功；**不是** ADR

## Reviewer pack（GitHub）

| Item | URL |
| --- | --- |
| Repo | https://github.com/linyihong/Ai-skill |
| Phase 1 closeout commit | https://github.com/linyihong/Ai-skill/commit/68d3108b3b68d3e759dee876a130a1ad71712979 |
| Domain + IL2CPP migrate | https://github.com/linyihong/Ai-skill/commit/bb927a7e |
| CLI binary rebuild（discovery path） | https://github.com/linyihong/Ai-skill/commit/4f86bbf4 |
| Plan | https://github.com/linyihong/Ai-skill/blob/main/plans/active/2026-10-29-1055-unity-analysis-decision-domain/_plan.md |
| Phase 1 evidence | https://github.com/linyihong/Ai-skill/blob/main/plans/active/2026-10-29-1055-unity-analysis-decision-domain/evidence/phase1-structural-acceptance.md |
| Phase 0 baseline | https://github.com/linyihong/Ai-skill/blob/main/plans/active/2026-10-29-1055-unity-analysis-decision-domain/evidence/phase0-preflight-baseline.md |
| Decision SoT | https://github.com/linyihong/Ai-skill/blob/main/analysis/unity/decision-workflow.md |
| Case ledger template | https://github.com/linyihong/Ai-skill/blob/main/analysis/unity/case-ledger-template.md |
| Soft procedure | https://github.com/linyihong/Ai-skill/blob/main/analysis/unity/workflows/unity-soft-visual-restore.md |
| Routing handoff test | https://github.com/linyihong/Ai-skill/blob/main/scripts/ai-skill-cli/internal/app/unity_route_handoff_test.go |

## Phase 1 判讀摘要（架構；非審核者獨立重跑）

| 項目 | 判定 |
| --- | --- |
| SoT 唯一 | PASS |
| 舊入口 | PASS |
| 文件連結 | PASS |
| Routing 正向 | PASS |
| Routing 負向 | PASS |
| 工具與 Discovery | PASS |

## Phase 2 執行契約（專案側）

1. 自然路徑 vs soft-enter 對照（可比物件身分／觀測條件；記 Creator、事件序列、GetContent）
2. 找**第一個**未解釋差異（不預設 `_Symbol`／atlas／sprite binding；API 成功 ≠ 完成）
3. 最小區分實驗（兩假設預測、反駁條件、副作用與回復）
4. 裁決與剩餘未知；缺自然對照 → `blocked`，不强行結案

## 本庫產物邊界

- 業務 raw evidence（截圖、位址、fixture）**不**進 Ai-skill
- 本檔只記 kickoff／驗收判定；專案 ledger 複製自 `analysis/unity/case-ledger-template.md`

## Status

| Field | Value |
| --- | --- |
| Phase 2 | **open** — Soft49 ledger v1 started in business project |
| Project path | `<PROJECT_ROOT>` = `Programs/unwrapping/pojiepokerist`（cabinet `docs/slots/el-dorado-dices`） |
| Soft49 ledger | `<PROJECT_ROOT>/docs/slots/el-dorado-dices/lab/phase2-soft49-case-ledger-v1.md` |
| First difference / blocked | Create/5 ABI **ruled-out**；Rect origin vs per-reel **observed**；Creator／物件身分／`_Symbol` **not-tested**；新自然進盤受 Guest＜參賽費 **blocked**（口述） |
| Decision-domain worth keeping | _pending_（ledger 已能標 blocked／not-tested；尚未跑完最小區分實驗） |
| ADR / Phase 3 | **out of scope** this phase |
