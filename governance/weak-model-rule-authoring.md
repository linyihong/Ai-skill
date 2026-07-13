# 弱模型可讀性寫作標準（Weak-Model Rule Authoring Standard）

本檔是本系統所有 canonical 規則文件（enforcement / governance / workflow / models 規則正文）的**寫作驗收標準**。讀者假設是**能力較弱的模型**（Sonnet / Haiku 等級），不是能自行補完抽象意圖的強模型。

> **為什麼需要這條標準**：本系統的長期運作者是較弱的模型。抽象要求（「保持高品質」「謹慎處理」「適當驗證」）對強模型是提示，對弱模型等於沒寫——它無法把抽象詞轉成具體動作。本標準把「一條規則怎樣才算可被弱模型執行」寫成四要件 + 驗收 checklist。本檔自身也遵守此標準（dogfood）。

## 觸發條件（何時套用本標準）

出現下列任一情況時，套用本標準檢查你正在寫的規則：

- 新增或大幅改寫 `enforcement/` / `governance/` / `workflow/` / `models/` 的規則正文。
- 為某個 obligation、gate、checklist、rubric 寫「agent 應該怎麼做」的段落。
- Phase 5b（本 plan）產出 F2–F6 任一檔時，用本標準當量尺。

不套用：純索引 README（只有連結）、純資料表（glossary 詞條）、plan 的敘事段落（Decision Rationale 等分析文字）。

## 四要件（每條可執行規則必備）

一條「可被弱模型執行」的規則，**必須**同時具備下列四項。缺任一項即不合格。

| # | 要件 | 判準（怎樣算有） | 反例（這樣算沒有） |
|---|---|---|---|
| 1 | **觸發條件** | 讀者能機械判斷「現在該不該套用這條」——用可觀察的檔案類型 / 動作 / 狀態描述 | 「在適當時機」「必要時」「處理重要變更時」——無法判斷何時算「適當 / 必要 / 重要」 |
| 2 | **具體動作** | 一個弱模型讀完能立刻知道要**做什麼**（讀哪個檔、跑哪個命令、輸出什麼格式），不需推理意圖 | 「謹慎處理」「確保品質」「適當驗證」——沒說具體做什麼 |
| 3 | **可驗證完成判準** | 有一個**二值**（是 / 否）可檢查的完成條件，不靠自我感覺 | 「直到滿意為止」「確認沒問題」——沒有可覆核的判準 |
| 4 | **≥1 正例 + ≥1 反例** | 至少一個「這樣做對」和一個「這樣做錯」的具體例子，讓弱模型能對照 | 只有規則敘述，沒有例子 |

## 執行動作（寫規則時逐步做）

1. 寫完一條規則後，對照四要件表逐項自檢。
2. 任一要件缺 → 補齊，不是留待日後。
3. 特別檢查要件 2：把每個抽象動詞（「處理」「確保」「驗證」「注意」）替換成具體動作（「讀 X 檔」「跑 `ai-skill Y`」「輸出 Z 格式」）。做不到替換 = 這條規則其實是意圖不是規則，改寫或標為 behavioral-only 判斷（見誠實條款）。
4. 每條規則附至少一組正反例。

## 可驗證完成判準（本標準自身的驗收）

一份規則文件通過本標準，當且僅當：

- [ ] 每條可執行規則都有明確觸發條件（無「適當 / 必要 / 重要」等未定義詞當觸發器）
- [ ] 每條規則的動作是具體的（無單獨出現的「處理 / 確保 / 驗證 / 注意」而不接具體對象）
- [ ] 每條規則有二值完成判準
- [ ] 每條規則附 ≥1 正例 + ≥1 反例
- [ ] **終驗（最硬）**：一個 Sonnet 或更弱等級的 fresh session，只讀這份文件就能照做，且做出的結果可被獨立覆核為「對」

**終驗誰來做（流程所有權）**：作者**不能自簽完成**。流程是「作者自檢四要件 → 自改提交 → 由一個 fresh-context 較弱 model（verifier）執行終驗」。作者自檢通過只代表「準備接受驗收」，不代表「已完成」；終驗由 verifier 產出證據，做不對 = **文件不合格改文件，不是怪 model**。終驗做法見 [`../plans/active/2026-07-10-0955-final-form-feedback-execution-closure.md`](../plans/active/2026-07-10-0955-final-form-feedback-execution-closure.md) Phase 5b 驗收欄。本檔自身即用此流程驗收（Haiku verifier 找出 8 處弱模型可讀性缺陷，逐條修正）。

## 完整正反例（改寫示範）

**反例（不合格規則）**：

> 「Agent 在處理重要變更時應謹慎驗證，確保品質後再提交。」

四要件全缺：觸發器「重要變更」未定義；動作「謹慎驗證」「確保品質」抽象；判準「品質」不可覆核；無例子。弱模型讀完不知道要做什麼。

**正例（合格改寫）**：

> **觸發**：staged diff 含 `enforcement/` 或 `workflow/` 的 `.md` 檔。
> **動作**：提交前跑 `cd scripts/ai-skill-cli && go test ./...`；並用 `git diff --cached` 逐檔確認沒有本機絕對路徑前綴（即 macOS 家目錄那種 `<絕對路徑>` 字面，應改成 `<PROJECT_ROOT>` 占位符）。
> **完成判準**：`go test` 全綠 且 staged diff 經 sanitization scan（`ai-skill hooks run pre-commit`——這是本 repo pre-commit hook，會掃 staged 內容找 forbidden token；forbidden token 定義見 [`../enforcement/sanitization.md`](../enforcement/sanitization.md)，命令契約見 [`../scripts/ai-skill-cli/docs/command-contract.md`](../scripts/ai-skill-cli/docs/command-contract.md)）回報 0 個 forbidden token。
> **正例**：測試綠 + sanitization scan 乾淨 → 提交。
> **反例**：測試綠但 sanitization scan 報 2 個 forbidden path token（diff 夾了兩條本機絕對路徑）→ 不提交，先把路徑換成 `<PROJECT_ROOT>` 占位符。

## 誠實條款（本標準補不了的）

有些判斷**本質上無法降成四要件**——品味判斷（這段文案好不好）、模糊需求的取捨、缺乏客觀判準的設計選擇。遇到這種：

- **不要**硬編一個假 rubric 假裝可判（那會誤導弱模型產生 false confidence）。
- 明文標成 `behavioral-only judgment`，並給 escape hatch：升級到更強 model、要求使用者/第二意見、或明說「此項無機械判準，需人工裁量」。
- 判斷 rubric 的 escape hatch 統一格式見 [`judgment-rubrics.md`](judgment-rubrics.md) §誠實條款。

## 誰會參考這裡（Inbound References）

- [`plans/active/2026-07-10-0955-final-form-feedback-execution-closure.md`](../plans/active/2026-07-10-0955-final-form-feedback-execution-closure.md) — Workstream F / Phase 5b F1
- [`judgment-rubrics.md`](judgment-rubrics.md) — F3，套用本標準的四要件
- [`enforcement/edit-authority-map.md`](../enforcement/edit-authority-map.md) — F5，套用本標準

## 與既有層的關係

- [`document-sizing.md`](document-sizing.md) — 管「文件多大該拆」；本檔管「規則怎麼寫才可執行」。兩者正交。
- [`enforcement/content-layering.md`](../enforcement/content-layering.md) — 管「內容該放哪層」；本檔管「放好之後怎麼寫」。
