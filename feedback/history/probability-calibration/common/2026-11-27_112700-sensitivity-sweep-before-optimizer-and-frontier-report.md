> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-11-27 - Sweep sensitivity before optimizing; report the trade-off frontier

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate（N=1）。一次調參嘗試沒有找到可採用的候選、也沒有採用任何變更；執行者當時自行寫下 `partially_calibrated`，事後依 workflow 規則更正（見 Generalized Lesson 4）。

#### One-line Summary

多個硬約束加一個新目標時，先做一維掃描與小格點找出「哪一族旋鈕真的能動目標指標」，再決定要不要跑優化器；找不到可行點就回報權衡前緣與缺口，不要硬宣稱不可行，也不要為了收工採用破壞硬約束的點。

#### Human Explanation

隨機局部搜尋在多個強耦合約束下常卡在原地，且高負載機器上每次評估很貴。先用同一個種子對每個旋鈕單獨掃描，可立刻看出哪些旋鈕對目標指標幾乎無效（本輪的跨維耦合旋鈕）、哪些有效但同時推動硬約束（結構旋鈕）。再用小格點把「目標指標 vs 硬約束」畫成前緣。若前緣上沒有同時滿足所有硬約束的點，結論是搜尋範圍內的權衡，不是結構性證明。

#### Trigger

- 新增一個校準目標，既有硬約束都已對齊
- 優化器多輪不動或只在硬約束邊界附近震盪
- 想用一個總分把多個指標合併

#### Evidence

- Tool: 同一種子的單旋鈕掃描 → 小格點 → 貪婪搜尋；每點記錄所有 hard_constraint 與目標指標
- Sanitized excerpt: 3×3 格點中，目標指標改善的點都伴隨硬約束越界（一側總期望值掉到目標以下，另一側命中率上升）；最佳折衷仍離目標約 30–40%
- Evidence path: `<PROJECT_ROOT>/docs/` 的缺口與試驗紀錄；探針與搜尋腳本留在 consumer

#### Generalized Lesson

1. 先做單旋鈕敏感度，再決定搜尋空間；無效旋鈕不要放進優化器。
2. 所有候選用同一種子比較（共同亂數），降低比較雜訊。
3. 輸出前緣表：每點同時列硬約束與目標指標，而不是單一分數（`single_aggregate_score_cannot_decide_success`）。
4. 沒有可行點時，結論限定為「搜尋範圍內」；`structural_blocked` 需要 C7 完整證據；證據不足時執行者只寫缺口與下一步選項（加結構約束或放寬某目標，由 decision owner 決定）。**執行者不寫最終 status**：沒有候選可獨立驗證時，依 `evaluation_order` 第一條落在 `inconclusive`，由 verifier／decider 定案；`partially_calibrated` 需要完成的驗證鏈。
5. 目標改動必須由 decision owner 明確授權；本輪「缺口原本刻意不動」的決定先寫入 ledger，授權後才開工。

#### Agent Action

- 先做：單旋鈕掃描；同種子格點；前緣表；缺口寫入 Gap Ledger，status 留給 verifier／decider。
- 不要做：用總分挑點；放寬硬約束換目標；把搜尋失敗寫成結構不可行；以執行者身分自行寫 `partially_calibrated` 等最終 status。

#### Goal / Action / Validation

- Goal: 在不破壞硬約束下改善目標指標，或誠實記錄權衡。
- Action: 敏感度 → 前緣 → 狀態。
- Validation or reference source: 前緣表；若採用變更需 validation 種子與 fresh-context final holdout（PC4）。

#### Applies When

- 多硬約束 + 單一新目標
- 評估昂貴、旋鈕數有限

#### Does Not Apply When

- 單一指標、無硬約束
- 已知可解析解／封閉式調整

#### Validation

在另一個 validation 種子重算前緣上的關鍵點，確認排序不變；若有採用變更，需獨立驗證者以未用過的 holdout 種子驗收。

#### Promotion Target

- `workflow/probability-calibration/execution-flow.md`（Tune：先敏感度，後優化；結果表達為前緣）
- `intelligence/engineering/probability-calibration/`（plan Phase 2 的 I7 單旋鈕副作用、I9 優化操作衛生 atoms 可吸收）

#### Promotion Record

尚未 promotion；交由 plan 擁有者併入 Phase 2 atoms 評估。

#### Required Linked Updates

- [x] `feedback/history/probability-calibration/common/README.md`
- [x] `feedback/history/probability-calibration/README.md`
- [x] `feedback/history/README.md` domain 表
- [ ] Phase 2 atoms（待 plan 擁有者）

#### Transfer Contract

- Inputs: 一組硬約束、一個新目標、有限旋鈕
- Mechanism: 旋鈕族對指標的敏感度不均；目標與硬約束共享同一組結構旋鈕時形成前緣
- Decision: 先敏感度與前緣，再決定是否優化或換結構
- Verification: 同種子前緣；validation 種子重算
- Counterexample: 無耦合的獨立旋鈕（可逐一調）
- Transfer case: 任何有多個必須同時守住的比率的模型調參（轉換率、掉落率、成本）

#### Intelligence Review

- Decision: deferred
- Reason: 單一 consumer（N=1）；承接位置（workflow 契約或 Phase 2 atoms）由 plan 擁有者決定，本 lesson 不直接改 workflow
- Target: none
- Revisit: plan Phase 2／Phase 3 評估這條或第二個 consumer 出現同型情境時

#### Abstraction Check

- Swap-project test: PASS — 多約束調參都適用
- Project evidence kept in: `<PROJECT_ROOT>/docs/`
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS
