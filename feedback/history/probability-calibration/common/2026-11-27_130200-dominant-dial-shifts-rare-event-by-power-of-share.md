> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-11-27 - A dominant dial moves rare-event rates by a power of the change

Status: candidate

#### Reuse Evidence

N=1（I7 單旋鈕副作用同 consumer 的第二個實例；不算獨立重用）。

#### One-line Summary

找到能動目標指標的主旋鈕後，動它之前先把「需要 k 個同類才發生」的稀有事件列為診斷指標並預測其變化：事件率約隨旋鈕占比的 k 次方改變，小幅調整就能讓稀有事件消失一個量級。

#### Human Explanation

調參時常只盯著均值類指標；稀有事件（連續 k 個同類、多條件同時成立）的機率對占比極敏感。占比降 40%、k=5，事件率約降到 0.6^5≈8%。若事件是產品特色（大獎），調整會悄悄改變體驗，而契約的指標清單裡沒有它。

#### Trigger

- 主旋鈕是某類元素的占比
- 產品有「需要 k 個同類」的稀有事件或大獎
- 準備調整占比以壓低均值類指標

#### Evidence

- Tool: 動旋鈕前用解析式或大樣本估稀有事件率，調整後再量
- Sanitized excerpt: 占比降 40%（k=5），預測倍率約 0.078；實測稀有事件約從 1/50 萬降到 1/400 萬（倍率約 0.125），方向與量級一致，差異來自非占比來源的組合
- Evidence path: `<PROJECT_ROOT>/docs/` 的校準紀錄；探針與量測腳本留在 consumer

#### Generalized Lesson

1. 旋鈕敏感度掃描要同時掃稀有事件率，不只掃均值類指標。
2. 動旋鈕前先寫預測（k 次方），調後比較；預測與實測差一個量級以上代表有未知機制。
3. 稀有事件若是產品特色，列為 `diagnostic_metric`，或由 decision owner 明確接受其改變（Gap Ledger `accepted_risk`）。
4. 稀有事件樣本以事件數判斷充足（C3），不以手數。

#### Agent Action

- 先做：動旋鈕前預測稀有事件變化；列入診斷指標；把代價交 decision owner。
- 不要做：只回報均值類指標的改善；在不知道稀有事件率改變的情況下交付。

#### Goal / Action / Validation

- Goal: 讓旋鈕變更的稀有事件副作用事先可見。
- Action: 預測→量測→交 owner。
- Validation or reference source: 預測與實測的比對表。

#### Applies When

- 占比類旋鈕
- 產品存在 k 個同類的稀有事件

#### Does Not Apply When

- 旋鈕與稀有事件獨立
- 稀有事件不在產品範圍內

#### Validation

用另一個 seed 與更大樣本重量稀有事件，確認倍率落在預測區間；事件數不足時只記錄不下結論。

#### Promotion Target

- `intelligence/engineering/probability-calibration/single-knob-side-effects.md`（補第二個實例與預測式）
- `workflow/probability-calibration/execution-flow.md`（Tune：動旋鈕前預測稀有事件）

#### Promotion Record

尚未 promotion；同 consumer 實例不計入升級門檻。

#### Required Linked Updates

- [x] `feedback/history/probability-calibration/common/README.md`
- [x] `feedback/history/probability-calibration/README.md`（domain 已列）
- [ ] I7 atom 補充（待 plan 擁有者）

#### Transfer Contract

- Inputs: 占比旋鈕、k 個同類的稀有事件
- Mechanism: 事件率≈占比^k
- Decision: 預測並列入診斷指標或明確接受
- Verification: 預測倍率與實測倍率同量級
- Counterexample: 與占比無關的稀有事件
- Transfer case: 任何「連續 k 次成功」型事件的比率調整（連擊、保底觸發、串聯故障）

#### Intelligence Review

- Decision: deferred
- Reason: 單一 consumer（N=1）
- Target: none
- Revisit: plan Phase 2 atoms 檢視或第二個 consumer

#### Abstraction Check

- Swap-project test: PASS — 「k 次方」推論與領域無關
- Project evidence kept in: `<PROJECT_ROOT>/docs/`
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS
