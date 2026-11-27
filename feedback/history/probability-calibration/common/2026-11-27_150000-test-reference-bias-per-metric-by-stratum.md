> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-11-27 - Test a reference sample's bias per metric and stratum before excluding the whole sample

Status: candidate

#### Reuse Evidence

N=1：同一 consumer 第一次完整走 workflow（契約先行、獨立 fresh-context 驗收）的一輪；尚無獨立 consumer 重用證據。

#### One-line Summary

參照樣本被標為「偏誤樣本」時，不要整份排除：依已知偏誤因子（如帳號年資）對每個指標分層檢查；有指標隨因子漂移、有指標平坦，平坦的指標仍可當參照。

#### Human Explanation

熱樣本規則（絕對值不當目標）原本以整份樣本為單位。實際上偏誤常只影響部分指標。對某一指標，把樣本依偏誤因子切成幾層，各層的均值若沒有單調趨勢，且層間差異在抽樣誤差內，就可判定該指標未受該偏誤影響；漂移的指標仍只取尺度不變結構。這讓目標能由實際資料決定，而不是因為樣本整體可疑就丟掉可用資料。

#### Trigger

- 參照樣本有已知偏誤因子（新舊帳號、時段、來源）
- 同一樣本含多類指標（單手回報與整場結算）
- 想引用其中一類指標當目標

#### Evidence

- Tool: 以偏誤因子分 7 層，逐層算指標均值與分位數
- Sanitized excerpt: 單手回報在最年輕一層約為穩態的 2 倍（明顯偏誤）；整場結算在 7 層的均值落在 28.9 到 33.7，無單調趨勢，樣本數每層 39 到 684
- Evidence path: `<PROJECT_ROOT>/docs/` 的校準契約、紀錄與缺口登錄；採集與探針腳本留在 consumer

#### Generalized Lesson

1. 偏誤判斷單位是「指標 × 偏誤因子」，不是整份樣本。
2. 分層檢查要預先宣告因子與層數，並記錄每層樣本數；樣本太少的層只當觀察。
3. 指標平坦只證明「對該因子不敏感」，不證明樣本完整（見另一條關於參照完整性的 lesson）。
4. 仍受偏誤的指標維持不當目標，並在 Gap Ledger 記錄原因。

#### Agent Action

- 先做：分層檢查每個候選參照指標；通過者當 reference_observation，未通過者只取尺度不變結構。
- 不要做：因為樣本整體標示為偏誤就丟掉所有指標；或沒有檢查就當作可用。

#### Goal / Action / Validation

- Goal: 在不誤用偏誤資料的前提下，用得上的參照資料。
- Action: 指標×因子分層檢查。
- Validation or reference source: 分層表（每層 n、均值、分位數）。

#### Applies When

- 參照樣本有已知偏誤因子
- 指標類別多於一種

#### Does Not Apply When

- 偏誤因子未知或無法分層
- 樣本太小無法分層

#### Validation

換另一個偏誤因子再分層，結論不翻轉；層間差異以誤差帶而不是肉眼判斷。

#### Promotion Target

- `workflow/probability-calibration/contracts/calibration-contract.yaml`（hot sample 規則改為「指標×偏誤因子」層級；診斷順序 b 步驟寫入分層檢查）
- `workflow/probability-calibration/profiles/slot-rtp.md`（單手與整場結算的偏誤差異）

#### Promotion Record

尚未 promotion；交由 plan 擁有者評估。

#### Required Linked Updates

- [x] `feedback/history/probability-calibration/common/README.md`
- [x] `feedback/history/probability-calibration/README.md`（domain 已列）
- [ ] 契約／診斷 b 步驟（待 plan 擁有者）

#### Transfer Contract

- Inputs: 帶偏誤因子的參照樣本、多類指標
- Mechanism: 偏誤只作用於部分指標
- Decision: 指標×因子分層檢查後決定可用性
- Verification: 另一偏誤因子結論不翻轉
- Counterexample: 偏誤因子未知
- Transfer case: 灰度用戶樣本、早期用戶樣本、特定時段樣本的行為指標

#### Intelligence Review

- Decision: deferred
- Reason: 單一 consumer（N=1）；承接位置由 plan 擁有者決定
- Target: none
- Revisit: plan Phase 3 或第二個 consumer

#### Abstraction Check

- Swap-project test: PASS — 與領域無關的取樣偏誤檢查
- Project evidence kept in: `<PROJECT_ROOT>/docs/`
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS
