# 判斷 Rubric 單一入口（Judgment Rubrics）

本檔把散在 `enforcement/rule-weight.md`、`enforcement/escalation-policy.md`、`workflow/software-delivery/closure.md`（DoD）、`enforcement/rollout-boundary` 的判斷規則，收斂成**一個入口**，讓較弱的模型不必自己跨檔推導。每組 rubric 附觸發、動作、二值判準、一正例一反例（遵守 [`weak-model-rule-authoring.md`](weak-model-rule-authoring.md) 四要件）。

> **怎麼用**：遇到「我該升級模型嗎 / 這算完成了嗎 / 我該問使用者嗎 / 我該換方法還是再試一次嗎 / 這品質夠嗎」五類猶豫時，直接查對應 rubric，照二值判準決定。**不要**憑感覺，也不要重新推導 P0–P3 權重。

## Rubric 索引

| # | 猶豫 | 一句話判準 |
|---|---|---|
| R1 | 該升級模型嗎 | 弱模型錯一次、中階同一子任務連錯兩次 → 升級（見 [`models/routing/model-tier-escalation.md`](../models/routing/model-tier-escalation.md)） |
| R2 | 這算真的完成嗎 | 「輸出存在」不是完成；「可被獨立覆核為對」才是完成 |
| R3 | 該停下問使用者嗎 | 涉及不可逆 / scope / 權限 / 互斥結果 → 停下問；有合理預設 → 做並說明 |
| R4 | 換路還是重試 | 同一假設沒有新證據卻反覆失敗 → 換路，不是重試 |
| R5 | 品質夠了嗎 | 有客觀判準的照判準；無客觀判準的走誠實條款，不自我背書 |

---

## R1 — 該升級模型嗎

**觸發**：目前 model 產出被驗證為錯 / 不合格，你在考慮是自己重試還是換更強的 model。

**動作**：查失敗計數與 model tier，套用升級階梯（詳細階梯在 [`models/routing/model-tier-escalation.md`](../models/routing/model-tier-escalation.md)）：弱 tier（Haiku 級）錯 1 次即升級；中 tier（Sonnet 級）同一子任務連錯 2 次即帶完整失敗軌跡升級；同一子任務最多重試 2 輪。

**完成判準（二值）**：失敗次數 ≥ 該 tier 的升級門檻 → 升級；否則允許再試一輪（但總重試不超過 2）。

- **正例**：Haiku 執行「改 3 個檔的 import」，第一次改錯路徑導致 build fail → 立即升級 Sonnet，附失敗 diff。（弱 tier 錯一次就升，不硬撐）
- **反例**：Haiku 錯一次，主 session 讓它「再試試看」，連試 4 輪都失敗才升級 → 浪費 4 輪成本，違反「弱 tier 錯一次即升」。

---

## R2 — 這算真的完成嗎

**觸發**：你準備回報「完成 / 做好了 / 修好了」。

**動作**：對照完成三層，全部成立才可宣稱完成：

1. **輸出存在**：檔案已寫 / 程式碼已改 / commit 已建。
2. **可獨立覆核**：檔案用 read-back 確認內容真的在；程式碼用測試或實跑確認行為對；不能只憑「我剛剛寫了」。
3. **閉環乾淨**：`git status` clean **且** 該 push 的已 push（`git log origin/<branch>..HEAD` 為空）。

> **「pending 狀態」不等於「完成」**：若 commit 了但未獲授權 push，這是 `complete-pending-user-decision`（完成待使用者決定），**不可**回報為「完成」——必須明說「本地已 commit、尚未 push，等你決定」。第 3 層的二值判準是「clean + 已 push」；未 push 一律當作**未通過第 3 層**，只是允許以「pending 待決」形式交回使用者。

**完成判準（二值）**：三層全 yes → 可宣稱完成；任一 no → 未完成（或 pending 待決），說明卡在哪層。

- **正例**：改完 validator，跑 `go test ./...` 全綠，commit + push，`git log origin/main..HEAD` 為空 → 回報「完成，測試綠、已推送」。
- **反例（第 2 層）**：改完 validator，「測試應該會過」就回報完成 → 違反第 2 層（未實跑覆核）；本系統 P1 §全系統審計的 scenario 自申報就是這個病。
- **反例（第 3 層）**：改完 + commit，但沒 push，回報「完成」→ 違反第 3 層；正確回報是「本地已 commit，尚未 push（pending 待你決定）」。

---

## R3 — 該停下問使用者嗎

**觸發**：你要做一個動作，但不確定使用者是否會同意。

**動作**：判斷該動作是否落入「必須停下問」的類別：

- **必停**：不可逆（刪資料 / rewrite git history / 覆蓋既有內容）、改變 scope（做超出被交代範圍的事）、動權限 / 帳號 / 對外發布、兩個合理選項互斥且選錯代價高。
- **可自主做並說明**：有明顯合理預設、可回退、屬於被交代任務的必要步驟。

**完成判準（二值）**：動作 ∈ 必停類別 → 停下問，附選項與風險；否則做並在回報中說明所做選擇。

- **正例**：使用者要「清理 repo」，你發現最有效的是 rewrite git history（不可逆、影響所有 clone）→ 停下，說明風險與替代方案，等使用者拍板。（本 plan Phase 9 就是這樣處理 bin history 清理）
- **反例**：使用者要「加一個 validator」，你順手把三個既有 validator 重構了 → 違反 scope，應該只做被交代的、或先問。

---

## R4 — 換路還是重試

**觸發**：你的方法失敗了，在考慮再試一次還是換做法。

**動作**：檢查失敗是否帶來**新證據**：

- **有新證據**（錯誤訊息指出新的 root cause、讀到之前沒讀的 source）→ 用新證據調整後重試一次是合理的。
- **無新證據**（同樣假設、同樣做法、只是「再跑一次」）→ 換路：重讀 source-of-truth、換方法、或升級 / 問使用者。反覆用沒有新證據的方式 patch 是 autonomy downgrade 訊號（見 [`models/routing/autonomy-routing.md`](../models/routing/autonomy-routing.md) §Downgrade Triggers）。

> **「新證據」的判斷**：`新` 指「這次失敗揭露了上次不知道的資訊」。**第一次就該讀的錯誤訊息，第二次才去讀 → 不算新證據**（那是上次偷懶）。只有「已經看過訊息、據此修正、又出現不同的新錯誤 / 新 root cause」才算新證據。判準：你能不能具體說出「這次比上次多知道了什麼」？說不出 → 無新證據。

**完成判準（二值）**：兩次失敗之間有無新證據（你能具體說出多知道了什麼）？無 → 禁止第三次同法重試，必須換路。同一子任務總重試上限 2 輪。

- **正例**：測試 fail，錯誤訊息顯示是缺一個 registry entry（新證據）→ 補 entry 重試，通過。（本 session 的 orphan_executor 修復就是這樣：錯誤訊息指出確切 symbol，補 allowlist 一次過）
- **反例**：測試 fail，不看錯誤訊息就「再跑一次看看」，連跑三次一樣 fail → 無新證據的重試，應該去讀錯誤訊息 / source。

---

## R5 — 品質夠了嗎

**觸發**：你要判斷產出的品質是否達標。

**動作**：先分類這個品質判斷有沒有客觀判準：

- **有客觀判準**（測試通過率、link check、去敏 grep、格式 validator）→ 跑判準，二值決定，不用主觀感覺。
- **無客觀判準**（文案好不好、設計取捨、需求模糊下的選擇）→ 走**誠實條款**（見下），不自我背書「我覺得夠好」。

**完成判準（二值）**：客觀判準存在且通過 → 達標；客觀判準不存在 → 不宣稱「高品質」，改用誠實條款輸出。

- **正例**：判斷「這份規則文件合格嗎」→ 跑 [`weak-model-rule-authoring.md`](weak-model-rule-authoring.md) 四要件 checklist（客觀）+ fresh 弱 model 終驗 → 二值結論。
- **反例**：判斷「這段架構分析寫得好不好」→ 回報「品質很高」→ 無客觀判準的自我背書，應該走誠實條款：標明「此為判斷性內容，建議第二意見覆核」。

---

## 誠實條款（品味 / 模糊判斷的 escape hatch）

本系統靠拆解、驗證、多樣本評審能補**執行品質**；但**補不了模糊題與品味判斷**。遇到無客觀判準的判斷，弱模型**不得**假裝有判準、不得自我背書「夠好」。改用下列三選一，並在回報中明說採用哪個：

| Escape hatch | 何時用 | 具體動作 |
|---|---|---|
| **升級 model** | 判斷需要更強推理，且成本可接受 | 依 R1 / model-tier-escalation 升級，附完整 context |
| **第二意見 / 多樣本評審** | 有多個候選答案要選優 | 派 fresh-context agent 產生獨立判斷，或生成 N 個答案交叉評審選優 |
| **明說做不到** | 判斷本質需要人類品味 / 授權 | 回報「此項無機械判準，需人工裁量」+ 列出選項與各自 trade-off，不硬給結論 |

**反例（禁止）**：把品味判斷硬編成假 rubric（例如「文案品質 = 字數 / 段落數」）假裝可判 → 誤導弱模型產生 false confidence，比誠實說「做不到」更糟。

## 誰會參考這裡（Inbound References）

- [`plans/active/2026-07-10-0955-final-form-feedback-execution-closure.md`](../plans/active/2026-07-10-0955-final-form-feedback-execution-closure.md) — Workstream F / Phase 5b F3
- [`models/routing/model-tier-escalation.md`](../models/routing/model-tier-escalation.md) — R1 / R4 的 model 升降級細節
- [`weak-model-rule-authoring.md`](weak-model-rule-authoring.md) — R5 的客觀判準來源

## 與既有層的關係（本檔是收斂入口，不取代 source）

- R1/R4 的 model 升降級 canonical → [`models/routing/`](../models/routing/README.md)
- R2 完成閉環 canonical → [`enforcement/dependency-reading.md`](../enforcement/dependency-reading.md) §Ai-skill 回寫完成門檻 + close-loop obligation
- R3 停下問使用者 canonical → [`enforcement/rule-weight.md`](../enforcement/rule-weight.md) §不確定時 + 系統 prompt 的 action categories
- R4 換路訊號 canonical → [`enforcement/escalation-policy.md`](../enforcement/escalation-policy.md) + [`models/routing/autonomy-routing.md`](../models/routing/autonomy-routing.md)

本檔只做**收斂與正反例補充**，衝突時以上述 canonical source 為準。
