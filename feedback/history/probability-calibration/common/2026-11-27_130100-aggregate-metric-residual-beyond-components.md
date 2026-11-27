> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-11-27 - An aggregate metric has a residual its components do not explain

Status: candidate

#### Reuse Evidence

N=1；尚無獨立重用證據。

#### One-line Summary

整體指標（如整場均值）與其成分指標（線獎、附加獎）分開量與分開定義；成分全部對齊後，整體仍可能低於參照，剩餘量是殘差，要記進 Gap Ledger，不能默默調某個成分去吞掉。

#### Human Explanation

參照側的整體指標常含有採集管線才看得到、成分表看不到的部分（例如結算前的調整、被省略的尾端）。對齊成分只能保證成分，不保證整體。若為了讓整體對齊而放大某個成分，其他以該成分為基準的指標會跑掉，而且把量測缺口偽裝成了產品行為。

#### Trigger

- 參照側同時提供整體指標與成分拆分
- 成分總和與整體指標在參照上本身就不相等
- 調參後成分達標、整體仍偏低或偏高

#### Evidence

- Tool: 對參照側計算「整體 − 成分總和」殘差；對產品側同樣計算
- Sanitized excerpt: 參照側每手整體均值約 5.19，成分總和約 4.94，殘差約 0.25（約 5%）；成分對齊後產品整體仍低於縮放後參照約 20%
- Evidence path: `<PROJECT_ROOT>/docs/` 的校準紀錄；探針與量測腳本留在 consumer

#### Generalized Lesson

1. 校準前先算參照側殘差；殘差本身是 Target 的不確定來源（C1 `uncertainty`）。
2. 成分與整體各自是指標，各自有層級與容忍度；整體不從成分推得。
3. 成分達標後的剩餘缺口寫入 Gap Ledger（`definition_mismatch` 或 `known_model_limitation`），由 decision owner 決定接受或補建模。
4. 不得調整單一成分去補殘差，除非 decision owner 明確選擇並記錄副作用。

#### Agent Action

- 先做：量殘差；成分與整體分開驗收；缺口入帳。
- 不要做：為了讓整體數字好看而調成分；把殘差當成產品誤差。

#### Goal / Action / Validation

- Goal: 避免把量測缺口誤當產品偏差並被錯誤修復。
- Action: 計算殘差、分開判定、入帳。
- Validation or reference source: 參照側殘差計算表。

#### Applies When

- 參照側有整體與成分兩種口徑
- 想用成分調整來逼近整體

#### Does Not Apply When

- 整體恰為成分總和（定義上無殘差）
- 只有單一指標

#### Validation

在另一批參照樣本重算殘差，確認殘差穩定（不是樣本雜訊）；穩定才當作結構性缺口。

#### Promotion Target

- `workflow/probability-calibration/execution-flow.md`（指標對照：先算參照殘差）
- `workflow/probability-calibration/contracts/gap-ledger.yaml`（殘差缺口的 kind 與範例）

#### Promotion Record

尚未 promotion。

#### Required Linked Updates

- [x] `feedback/history/probability-calibration/common/README.md`
- [x] `feedback/history/probability-calibration/README.md`（domain 已列）
- [ ] execution-flow 指標對照步驟（待 plan 擁有者）

#### Transfer Contract

- Inputs: 參照整體指標與成分拆分
- Mechanism: 整體含有成分表未涵蓋的部分
- Decision: 殘差入帳，不調成分吞掉
- Verification: 另一批樣本殘差穩定
- Counterexample: 定義上整體等於成分總和
- Transfer case: 營收與各通路營收、錯誤率與各類錯誤率

#### Intelligence Review

- Decision: deferred
- Reason: 單一 consumer（N=1）
- Target: none
- Revisit: plan Phase 3 或第二個 consumer

#### Abstraction Check

- Swap-project test: PASS — 任何整體／成分雙口徑的參照都適用
- Project evidence kept in: `<PROJECT_ROOT>/docs/`
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS
