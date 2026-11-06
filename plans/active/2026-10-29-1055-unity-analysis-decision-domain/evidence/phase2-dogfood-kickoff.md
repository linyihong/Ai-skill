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
| First difference / blocked | Create/5 ABI **ruled-out**；Rect **observed≠causal**；Soft50-A＝執行期 Create 關聯證據（非根因）；空石殼根因／畫面修復 **未確定／未驗收** |
| Soft50 | **A＝observed**。**164752**：deferred locate 窗內成立；`cellId`≠跨 enter id。Same-session 配對 **未完成**：165913／170809＝`data_insufficient`（無新鮮 Create）；172646＝操作者收工中止。**B／Rect write／C＝deferred**；H2b 身分級配對未立。下輪＝hook 後進 Bonus 取 natural Create→soft observe-only。收工筆記 pojie `boardgame-paint-soft50-session-closeout-20261029-1730.md`。不擴架構 |
| Decision-domain worth keeping | _pending_ |
| ADR / Phase 3 | **out of scope** this phase |
| Review pack（專案側 zip；非 Ai-skill） | `<PROJECT_ROOT>/docs/slots/el-dorado-dices/lab/phase2-soft49-review-pack-20261029.zip` |

## 2026-11-06 自然格位證據更新

專案被動觀測已建立十個 controller/model 的 cell ID 與 view 世界座標關聯，固定樣本中 757 次成功內容註冊全部可關聯同一 view/model。自然側格位換算已有獨立位置讀回支持。原始 trace、指標和換算保存在專案側 soft50-20261106/full/。接下來依既有契約取得同場 soft 對照並裁決第一個差異。Phase 2 維持 open，Soft50-B／Rect 寫入／C 的裁決待對照證據。

## 同場回放後的契約修復

同場 soft seating 暴露舊 build 的 hardcoded slots field offsets。改為按執行期欄位名稱與繼承鏈解析後，唯讀 layer/model/board probes 與後續 soft model 存取恢復。來源窗口標記仍需區分持續自然 callback 與 soft 操作。可重用方法為先核對欄位布局與型別，再裁決渲染差異。專案已保存修復前後 trace、回傳診斷和截圖。

剩餘程式缺口為 legacy 五軸雙列重建與自然十軸單列拓撲不符。soft glyph oracle 尚未通過，下一步修復 model/view 格位換算並逐格核對 fixture。Phase 2 維持 open。
