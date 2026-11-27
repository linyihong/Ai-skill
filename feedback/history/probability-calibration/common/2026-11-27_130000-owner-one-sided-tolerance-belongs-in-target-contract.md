> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-11-27 - A decision owner's one-sided tolerance belongs in the target contract

Status: candidate

#### Reuse Evidence

N=1（同一 consumer 的第二輪實戰，第一輪見其他 lesson）；尚無獨立 consumer 重用證據。

#### One-line Summary

決策者說「略低於參照值沒關係」時，容忍帶是單側且有方向；契約只有對稱的絕對／相對容忍度，這個方向性與縮放係數必須寫進 Target，否則驗收時會把「刻意偏低」誤判成缺口或誤判成通過。

#### Human Explanation

實作目標常是「參照觀測值 × 產品縮放係數」，而不是參照值本身；決策者再補一句偏好（某些指標寧可偏低不要偏高）。若契約只寫 `target ± tolerance`，驗收者無法分辨「低於容忍帶是被允許的」與「低於容忍帶是失敗」。方向性影響的是判定規則，不是量測精度，所以屬於 Target 契約，不屬於 Record 的註解。

#### Trigger

- 決策者表達非對稱偏好（寧低勿高、或寧高勿低）
- 實作目標由參照值乘上一個產品決定的縮放係數
- 同一指標在不同層級（hard_constraint／primary_objective）方向不同

#### Evidence

- Tool: 以縮放後的參照值為基準，記錄每個指標的實測值、縮放後目標、偏差方向
- Sanitized excerpt: 一輪調整後，主要指標落在縮放後目標 −4%，整場尾端指標 −20% 至 −35%；決策者先前明說後者偏低可接受，但契約沒有欄位承載，只能寫在對話與文件散文裡
- Evidence path: `<PROJECT_ROOT>/docs/` 的校準紀錄；探針與量測腳本留在 consumer

#### Generalized Lesson

1. 縮放係數是 Target 的一部分：記為 `implementation = reference_observation × scale`，並記 scale 的 `decision_owner` 與 `provisional｜confirmed`。
2. 每個指標宣告容忍方向（兩側／僅下界／僅上界）；方向由 decision owner 決定，不由執行者推斷。
3. 偏差落在「被允許的單側」時，仍要記錄實際偏差量與方向（Gap Ledger 的 `accepted_risk`），不得因為被允許就省略。
4. 偏差落在「不被允許的一側」時照一般缺口處理。

#### Agent Action

- 先做：合約化方向與縮放係數；驗收時分別報告偏差量與方向。
- 不要做：把「偏低沒關係」當作整體放寬、停止回報偏差；把單側許可套到其他指標。

#### Goal / Action / Validation

- Goal: 讓決策者的非對稱偏好可被機械判定。
- Action: Target 記縮放係數與容忍方向；驗收依方向判定；偏差入 Gap Ledger。
- Validation or reference source: 兩個方向各一個反例情境（偏低被接受、偏高被擋）能被 verifier 獨立判出。

#### Applies When

- 目標由參照值換算而來
- 決策者對偏差方向有偏好

#### Does Not Apply When

- 目標為規格值、決策者要求雙側精確
- 沒有決策者（純探索）

#### Validation

用兩個方向相反的合成偏差跑判定，確認結果不同；縮放係數變更後重算目標而不是沿用舊容忍帶。

#### Promotion Target

- `workflow/probability-calibration/contracts/calibration-contract.yaml`（C3b 容忍度：新增方向欄位、C1 縮放係數記錄）
- `workflow/probability-calibration/contracts/gap-ledger.yaml`（`accepted_risk` 附偏差方向與量）

#### Promotion Record

尚未 promotion；交由 plan 擁有者於 Phase 3 靜態走查回寫契約時評估（Phase 0.1 凍結期已過，Phase 3 允許依 dogfood 證據補欄位）。

#### Required Linked Updates

- [x] `feedback/history/probability-calibration/common/README.md`
- [x] `feedback/history/probability-calibration/README.md`（domain 已列）
- [ ] 契約欄位（待 plan 擁有者）

#### Transfer Contract

- Inputs: 參照值、縮放係數、決策者的方向偏好
- Mechanism: 判定規則依賴方向，對稱容忍度表達不出來
- Decision: 方向與係數寫進 Target 契約
- Verification: 兩個方向相反的合成偏差得到不同判定
- Counterexample: 雙側精確的規格型 Target
- Transfer case: 轉換率、掉落率、成本類模型調參中「寧低勿高」的指標

#### Intelligence Review

- Decision: deferred
- Reason: 單一 consumer（N=1）；契約欄位變更由 plan 擁有者決定
- Target: none
- Revisit: plan Phase 3 靜態走查或第二個 consumer

#### Abstraction Check

- Swap-project test: PASS — 任何「參照值×係數＋決策者偏好」的目標都適用
- Project evidence kept in: `<PROJECT_ROOT>/docs/`
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS
