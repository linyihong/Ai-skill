> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-11-27 - Reference completeness is an uncertainty the standard error does not cover

Status: candidate

#### Reuse Evidence

N=1：同一 consumer 第一次完整走 workflow（契約先行、獨立 fresh-context 驗收）的一輪；尚無獨立 consumer 重用證據。

#### One-line Summary

參照樣本的記錄完整性（有結算記錄的事件佔全部事件的比例）要和抽樣誤差分開記錄；同一個比率若有兩種可操作定義（以事件觸發的盤面計、以有結算記錄的事件計），把兩者當作區間，不要選一個當點估計。

#### Human Explanation

標準誤只反映樣本內的隨機性，不反映「哪些事件根本沒進樣本」。當事件數可由兩條途徑估計（觸發端與結算端）且兩者不一致，差距就是遺失比例。若遺失與結果大小有關（長場次更容易遺失），參照均值就會有系統偏差，而 SE 看不出來。契約應記錄兩個定義、遺失比例、以及遺失是否可能與結果相關。

#### Trigger

- 同一事件有觸發端與結算端兩種記錄
- 兩端計出的事件率差很多
- 要把事件率或事件均值當參照

#### Evidence

- Tool: 觸發端事件率對結算端事件率
- Sanitized excerpt: 觸發端約 1.45%、結算端約 0.96%，約三分之一觸發沒有結算記錄；結算端的事件均值 SE 為 0.73，但這個 SE 不含遺失造成的偏差
- Evidence path: `<PROJECT_ROOT>/docs/` 的校準契約、紀錄與缺口登錄；採集與探針腳本留在 consumer

#### Generalized Lesson

1. 兩個定義都記進契約，參照事件率用區間，而不是挑一個。
2. 遺失比例與「遺失是否可能與結果相關」寫入 Gap Ledger（`insufficient_evidence`）。
3. 驗收門檻若必須選一個點，說明選擇理由與另一端的差距。
4. 能補強就補強（補採集、用另一來源核對），否則誠實標示參照均值可能偏低或偏高的方向。

#### Agent Action

- 先做：同時算兩個端點的事件率；記錄遺失比例；把遺失與結果的可能關聯寫入缺口。
- 不要做：只用結算端或只用觸發端；把 SE 當作參照的完整不確定性。

#### Goal / Action / Validation

- Goal: 讓參照的不確定性包含「沒進樣本的部分」。
- Action: 兩端對照、區間化、入帳。
- Validation or reference source: 兩端事件數與遺失比例表。

#### Applies When

- 事件有兩種以上的記錄管道
- 事件率或事件均值要當參照

#### Does Not Apply When

- 記錄管道只有一種且保證完整

#### Validation

用另一批時段重算兩端事件率，遺失比例穩定才當作系統性；若遺失比例隨時段漂移，改為記錄漂移。

#### Promotion Target

- `workflow/probability-calibration/contracts/calibration-contract.yaml`（C1 `uncertainty`：包含 coverage／capture loss）
- `workflow/probability-calibration/contracts/gap-ledger.yaml`（`insufficient_evidence` 範例）

#### Promotion Record

尚未 promotion；交由 plan 擁有者評估。

#### Required Linked Updates

- [x] `feedback/history/probability-calibration/common/README.md`
- [x] `feedback/history/probability-calibration/README.md`（domain 已列）
- [ ] 契約 `uncertainty` 說明（待 plan 擁有者）

#### Transfer Contract

- Inputs: 同一事件的觸發端與結算端記錄
- Mechanism: 遺失事件不在樣本內，SE 不涵蓋
- Decision: 區間化事件率、記錄遺失與可能的相關性
- Verification: 另一時段兩端比例穩定
- Counterexample: 單一且完整的記錄管道
- Transfer case: 日誌取樣、轉換漏斗（點擊 vs 完成）、掉落記錄

#### Intelligence Review

- Decision: deferred
- Reason: 單一 consumer（N=1）；承接位置由 plan 擁有者決定
- Target: none
- Revisit: plan Phase 3 或第二個 consumer

#### Abstraction Check

- Swap-project test: PASS — 與領域無關的資料完整性問題
- Project evidence kept in: `<PROJECT_ROOT>/docs/`
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS
