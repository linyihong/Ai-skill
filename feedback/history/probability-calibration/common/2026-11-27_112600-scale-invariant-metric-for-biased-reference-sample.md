> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-11-27 - Biased reference sample: compare a scale-invariant ratio, not the absolute count

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate（N=1）。指標改定義後，缺口大小從「多 1.4–1.9 倍」變成可量化、可掃描的單一比值。

#### One-line Summary

參考樣本是偏誤子母體（熱樣本）時，絕對頻率（每 N 局幾次）不是可調目標；改用以**該事件的前置條件為分母**的比值（例如「每個觸發局的長連線數」），才是尺度不變、可跨樣本比較的校準指標。

#### Human Explanation

熱樣本的某些前置事件（觸發局）占比遠高於產品目標，直接比「每 10 萬局的事件數」會把觸發率差異和事件結構差異混在一起，且調到一致會破壞其他硬約束。把分母換成觸發局數後，指標只回答「一旦觸發，結構長怎樣」，與樣本偏誤脫鉤，也能直接拿來做參數掃描。

#### Trigger

- 參考樣本被標記為熱／農場／選擇性樣本
- 缺口是稀有事件的絕對頻率，且調它會撞硬約束
- 事件必然以某前置事件為條件（無前置就不可能發生）

#### Evidence

- Tool: 指標定義 `ratio = count(event) / count(precondition)`；兩側同定義後比較
- Sanitized excerpt: 參考 ~0.107、產品 ~0.171；以此比值掃描參數時可與其他指標並列比較
- Evidence path: `<PROJECT_ROOT>/docs/` 與本地探針（不進本檔）

#### Generalized Lesson

1. 先問事件的必要前置條件，把它放進分母，再比兩側。
2. 新指標先在兩側以同一定義重量（contract C2），再談差距。
3. 該指標的定義版本要升版並記在 Calibration Record，避免與舊絕對頻率混用。
4. 比值對齊後仍要檢查分母本身（觸發率）是否為獨立目標，不可因此遺漏。
5. **比值只消除前置事件率的偏誤，不消除潛在狀態混合的偏誤。** 若參照樣本還混了多個狀態各自使用不同的產生器（例如依帳號狀態切換不同的參數組），而各組的條件結構本來就不同，比值仍會隨混合權重漂移。有分層資料就分層比較；沒有就把「混合偏誤未量化」記進參照的 uncertainty 與 Gap Ledger，不要宣稱參照已不受偏誤影響。

#### Agent Action

- 先做：辨認前置條件；改定義；兩側重算；寫入 `definition_version`。
- 不要做：把熱樣本絕對頻率 1:1 當 implementation target；調參去追絕對頻率。

#### Goal / Action / Validation

- Goal: 缺口以不受樣本偏誤影響的形式表達。
- Action: 條件化分母。
- Validation or reference source: 兩側同定義重算；分母指標另行對照。

#### Applies When

- 事件嚴格依賴某前置事件
- 參考樣本非代表性

#### Does Not Apply When

- 偏誤來自潛在狀態混合且各狀態的條件結構不同、又沒有分層資料：比值仍受混合影響，只能當暫定參照
- 參考樣本被確認具代表性且絕對頻率本身即規格
- 前置事件與目標事件獨立

#### Validation

在第二個參數點重算，確認比值對分母變動不敏感，絕對頻率則隨分母漂移。

#### Promotion Target

- `workflow/probability-calibration/execution-flow.md`（Metric Parity：條件化比值）
- `workflow/probability-calibration/profiles/slot-rtp.md`（可作為 profile 指標詞彙）

#### Promotion Record

尚未 promotion；與 plan 既有規則 `hot_or_farmed_sample_absolute_values_not_targets` 互補（它禁止絕對值當目標，本 lesson 給替代指標形式）。

#### Required Linked Updates

- [x] `feedback/history/probability-calibration/common/README.md`
- [x] `feedback/history/probability-calibration/README.md`
- [x] `feedback/history/README.md` domain 表
- [ ] profile／契約更新（待 plan 擁有者）

#### Transfer Contract

- Inputs: 偏誤參考樣本、帶前置條件的稀有事件
- Mechanism: 絕對頻率 = 前置率 × 條件結構，偏誤只污染前置率
- Decision: 以條件結構比值為校準指標
- Verification: 兩側同定義；分母另列
- Counterexample: 代表性樣本、規格即絕對頻率；各狀態條件結構不同的混合樣本（比值仍隨混合權重漂移）
- Transfer case: 任何「熱樣本觀察 + 目標是條件行為」的機率模型（轉換漏斗、掉落鏈）

#### Intelligence Review

- Decision: deferred
- Reason: 單一 consumer（N=1）；承接位置（workflow 契約或 Phase 2 atoms）由 plan 擁有者決定，本 lesson 不直接改 workflow
- Target: none
- Revisit: plan Phase 2／Phase 3 評估這條或第二個 consumer 出現同型情境時

#### Abstraction Check

- Swap-project test: PASS — 任一條件事件在偏誤樣本中的條件化比值皆適用
- Project evidence kept in: `<PROJECT_ROOT>/docs/`
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS
