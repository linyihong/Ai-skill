# 回饋抽象化與 intelligence 承接的機械補強

## Observation

前一輪 [12 條回饋覆核](2026-10-26-recent-feedback-abstraction-audit.md) 修了正文、狀態與方法入口，但沒有把新的實質審查項目接進 executor，亦沒有逐條留下 intelligence 決定。這是流程／執行器缺口，不只是某一條 lesson 的文字問題。

原 `FeedbackLessonClosureRule` 只檢查章節存在、PASS 關鍵字與 Reuse Evidence 非空；commit adapter 用工作目錄內容。因而漏掉：

- 形式 PASS 沒有輸入、原因、反例或可攜情境。
- 回饋只進 analysis／workflow，未判斷 HOW TO THINK 是否應進 intelligence。
- 宣告已抽取，但目標／來源反向連結／父索引／發現入口不存在。
- validated 以「尚未獨立重用」的文字填滿 Evidence，仍被接受。
- 工作目錄修好或新檔 untracked，卻讓 staged 內容通過。

## Rule / Registry / Executor

- Canonical contract：[`enforcement/feedback-lessons.yaml`](../../enforcement/feedback-lessons.yaml)，寫作與流程 companion 為 [`feedback-lessons.md`](../feedback-lessons.md)、[lesson review](../extraction/lesson-review.md)。
- Registry：沿用 `feedback_lesson_closure`，不把語意抽象化的 behavioral class 誤報為完全機械化。
- Portable executor：[`feedback_review.go`](../../scripts/ai-skill-cli/portable/kge/feedback_review.go) 由既有 closure rule 執行。
- Adapter：[`feedback_review_adapter.go`](../../scripts/ai-skill-cli/internal/app/feedback_review_adapter.go) 從 git index 讀 lesson、target、父索引、evidence、summary 與 registry records；KGE CLI 與 commit-msg 共用相同 snapshot 原則。
- Contract 投影由 runtime compile 處理；executor 直接消費 index-attested Context，不另建一份可變的 promotion state store。

## 機械阻擋

1. Transfer Contract 缺具體 inputs／mechanism／decision／verification／counterexample／transfer case。
2. Intelligence Review 沒有 extract／existing／method-only／deferred 決定、理由或有效目標；deferred 缺觸發重審條件。
3. 目標不在正確層、未進 index、缺來源反向連結或父 README 索引；新 atom 卻冒稱 existing。
4. extract 缺 metadata、registry record 引用或指回目標的 Discovery summary；candidate source 卻提高新 atom 信心。
5. validated／promoted 缺 independent context、result、獨立 evidence reference，或正文明確否定獨立重用。
6. staged 壞內容以 unstaged／untracked 修補充數；新 lesson 只 stage 分類 README 卻沒有實際連結。

## 近期 intelligence 承接

[逐條決定表](../extraction/recent-review-map.md) 覆蓋同一批 12 條：

- 5 條抽成 2 個 candidate-intelligence：[局部證據歸因](../../intelligence/engineering/analytical-reasoning/heuristics/runtime-evidence-boundaries.md)、[表示契約對齊](../../intelligence/engineering/analytical-reasoning/heuristics/representation-contract-parity.md)。
- 4 條補入已有 evidence model、evidence depth 與最高槓桿分析路徑，避免重複建 atom。
- 3 條明確 method-only，保留資產取證與圖板裁切程序；不能因沒有新 atom 就認為沒有處理。

來源與新 atom 皆保留 candidate；沒有補造第二個專案的實測。

## Validation

- `TestFeedbackReviewContractMutations`：缺欄、placeholder、重複決定、錯層／越界路徑、無目標／反向來源／父索引／metadata／registry／summary、提高信心均阻擋。
- `TestFeedbackReviewDispositionAndEvidence`：extract／existing／method-only／deferred 正向通過；未解析 evidence、證據與狀態矛盾、新 atom 冒稱 existing 阻擋。
- `TestFeedbackReviewUsesIndexSnapshot`：未 staged 的來源、backlink、registry、summary 修復及 untracked 目標不能充數；完整 index 即使工作目錄檔案移除仍正確讀取。
- `TestFeedbackReviewCanonicalSchemaParity`：YAML 欄位、enum 與 git_index snapshot 不得與 executor 宣告漂移。
- [回歸 scenario](../../validation/scenarios/failure-derived/feedback-extraction-review-v1.yaml) 連回既有 failure pattern，實際結果由本次測試與 commit／close-out 記錄。

## 邊界

機械鎖驗證結構、決定與接線，不能證明作者的可攜性推演正確或證據真的獨立。對新增／修改條目增量生效；未碰觸的歷史不宣稱已回填，需以 `kge validate --all-feedback` 顯式盤點。候選抽取與提升成熟度分開，沒有自動 promotion。
