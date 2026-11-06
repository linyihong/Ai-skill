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
| Phase 2 | **project-side PASS，framework review pending** — 單案差異裁決與獨立驗收已記錄，尚未關閉整份計畫 |
| Project path | `<PROJECT_ROOT>` = `Programs/unwrapping/pojiepokerist`（cabinet `docs/slots/el-dorado-dices`） |
| Soft49 ledger | `<PROJECT_ROOT>/docs/slots/el-dorado-dices/lab/phase2-soft49-case-ledger-v1.md` |
| First difference / blocked | 早期 Create ABI 假說排除；同版布局、model/view 拓撲、內容快取、座標單位及實際 pivot 差異已有區分實驗與專案證據。具體修法與結果以專案驗收記錄為準 |
| Soft50 | **scoped PASS** — 目前版本與固定 fixture 的回放已由 model、獨立內容讀回和畫面三層驗收。其他 fixture／build／重進窗口覆蓋仍開放 |
| Decision-domain worth keeping | 單案試跑支持繼續使用，跨案例與架構升級仍待審查 |
| ADR / Phase 3 | **out of scope** this phase |
| Review pack（專案側 zip；非 Ai-skill） | `<PROJECT_ROOT>/docs/slots/el-dorado-dices/lab/phase2-soft49-review-pack-20261029.zip` |

## 2026-11-06 自然格位證據更新

專案被動觀測已建立十個 controller/model 的 cell ID 與 view 世界座標關聯，固定樣本中 757 次成功內容註冊全部可關聯同一 view/model。自然側格位換算已有獨立位置讀回支持。原始 trace、指標和換算保存在專案側 soft50-20261106/full/。接下來依既有契約取得同場 soft 對照並裁決第一個差異。Phase 2 維持 open，Soft50-B／Rect 寫入／C 的裁決待對照證據。

## 同場回放後的契約修復

同場 soft seating 暴露舊 build 的 hardcoded slots field offsets。改為按執行期欄位名稱與繼承鏈解析後，唯讀 layer/model/board probes 與後續 soft model 存取恢復。來源窗口標記仍需區分持續自然 callback 與 soft 操作。可重用方法為先核對欄位布局與型別，再裁決渲染差異。專案已保存修復前後 trace、回傳診斷和截圖。

剩餘程式缺口為 legacy 五軸雙列重建與自然十軸單列拓撲不符。soft glyph oracle 尚未通過，下一步修復 model/view 格位換算並逐格核對 fixture。Phase 2 維持 open。

## 單格座標修復與投影校準

專案單格實驗在主執行緒修正 model/view 座標單位後，目標骰子恢復可見。可見標記也推翻先前僅依 world Y 排序的上下排推估，已修正程式與 regression。十格回放的 Symbol／型別讀回通過，整體畫面仍有 Item 尺寸／位置差異，保持未驗收。下一步核對自然 creator/factory 的 Rect 傳遞及相同內容物件的 renderer/transform。Phase 2 仍 open。

## 乾淨場景的範圍驗收

專案已保存乾淨自然窗口、明確回放操作、另一支唯讀探針及延後畫面。驗收記錄為 `<PROJECT_ROOT>/docs/slots/el-dorado-dices/lab/boardgame-paint-soft50-repair-acceptance-20261106.md`，程式與證據提交 `3817d75`。本次具體結果留在專案，不作跨目標有效性宣稱。

前述未驗收描述為歷史階段。現已完成單案差異裁決與限定版本／fixture 的回放驗收。Phase 2 的 framework review、跨案例重用及整份 plan closure 未自動通過，Phase 3 與 ADR 仍未啟動。可重用候選補充設定 pivot 與實際 pivot 分離、探索後前置狀態檢查，以及模型／獨立內容／畫面的分層驗證。
