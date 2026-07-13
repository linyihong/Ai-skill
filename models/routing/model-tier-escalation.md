# 模型層級升降階梯（Model Tier Escalation Ladder）

本檔定義**依失敗次數觸發的 model 層級升降級規則**。這是 [`multi-model-handoff.md`](multi-model-handoff.md)（管「何時 handoff」）與 [`autonomy-routing.md`](autonomy-routing.md)（管「cognitive state → strategy」）**沒有**涵蓋的一段：當委派出去的 model 產出被驗證為錯，主 session 該讓它重試、還是換更強 / 更弱的 model。

> **與 escalation-policy 的明文互斥分工（重要，別搞混）**：
> - [`enforcement/escalation-policy.md`](../../enforcement/escalation-policy.md) 管的是**同一個 agent 內部**的 evidence-mismatch recovery——實際證據推翻假設時，停止局部 patch、重讀 source、重建 execution graph。**不涉及換 model**。
> - **本檔**管的是**跨 model tier** 的升降級——換一個更強 / 更弱的 model 執行。
> - 兩者可同時發生（同一 agent 先做 evidence recovery，仍失敗才 tier 升級），但**不是同一件事**，規則不可互相套用。

## Model Tier 定義（tool-neutral）

本系統用三個 tier 描述能力等級，不綁特定 model 名稱（名稱會隨版本漂移）：

| Tier | 能力定位 | 對應範例（本環境，2026-07；會漂移見 §參數保鮮） |
|---|---|---|
| **低 tier** | 快、便宜、適合機械批次套用已解出的模式 | Haiku 級 |
| **中 tier** | 日常執行主力，多數實作 / 重構 / 搜尋 | Sonnet 級 |
| **高 tier** | 難判斷、模糊需求、架構決策、對抗性驗證 | Opus 級 |

> 本 repo 的日常運作者是中 / 低 tier；高 tier（如本 session）用來**立制度**，不做日常任務。

## 升級階梯（失敗 → 升 tier）

**觸發**：委派給某 tier 的 model，其產出被 verifier / 測試 / read-back 驗證為錯或不合格。

**動作**（依當前 tier 套用門檻）：

| 當前 tier | 升級門檻 | 升級時必附 |
|---|---|---|
| **低 tier** | **錯 1 次**即升級（不硬撐） | 失敗產出 + 錯誤訊息 |
| **中 tier** | **同一子任務連錯 2 次**即升級 | **完整失敗軌跡**：兩次的 diff / 錯誤 / 已試方法 |
| **高 tier** | 不再自動升級（已是頂）；失敗 → 走 R3 停下問使用者 或 誠實條款 | 選項 + 各 trade-off |

**完成判準（二值）**：失敗次數 ≥ 當前 tier 門檻 → 升級（附上表要求的軌跡）；否則允許再試，但**同一子任務總重試上限 2 輪**（超過即強制升級或換路，見 [`../../governance/judgment-rubrics.md`](../../governance/judgment-rubrics.md) R4）。

- **正例**：Sonnet 執行「重構 auth 模組」，第一次漏改一個 caller 導致測試 fail，第二次改了 caller 但引入型別錯誤 → 連錯 2 次，帶兩次完整 diff + 錯誤訊息升級 Opus。
- **反例**：Sonnet 連錯 2 次，主 session 只丟一句「上次錯了，再做一次」給 Opus，不附失敗軌跡 → Opus 缺 context 可能重犯同錯，違反「升級必附完整失敗軌跡」。

## 降級階梯（模式已解出 → 降 tier 批次套用）

**觸發**：高 / 中 tier 已把某個問題的**做法解出來**（有明確、可重複的步驟），剩下的是把同一模式套到多個檔 / 多個位置。

**動作**：把解出的模式寫成明確步驟（輸入、動作、預期輸出），降回**低 tier** 批次執行；高 tier 不該親自做重複勞動。

**完成判準（二值）**：剩餘工作是否為「同一已解模式的重複套用」且步驟可明確描述？是 → 降級批次套用；否（仍需判斷）→ 留在當前 tier。

- **正例**：Opus 解出「每個 validator 需在 registry 補一行 executor entry」的模式後，把「對這 8 個 validator 各補一行」降給 Haiku 批次做，Opus 只驗收。
- **反例**：Opus 解出模式後，繼續自己逐一改 8 個檔 → 高 tier 做低 tier 的活，浪費成本；或反向錯誤——把「還需判斷哪些算 executor」這種未解判斷降給 Haiku，導致錯誤蔓延。

## 失敗計數的載體（跨 session 傳遞）

「連錯 2 次」的計數與失敗軌跡**必須有落點**，否則跨 session 會遺失（見本 plan Q12）。落點優先序：

1. 若在 delegation loop 內 → 記在被委派 sub-plan 的仲裁紀錄（fix / defer / reject 逐條記錄，見 [`plans/README.md`](../../plans/README.md) §Delegation）。
2. 若是需執行但本輪不修的失敗 → 進 deferred-feedback ledger（本 plan Workstream A，`feedback/pipeline/deferred/`）。
3. **禁止**只存在對話裡：對話會被 compaction，升級後的 model 讀不到 → 重犯同錯。

**完成判準（二值）**：失敗軌跡是否已寫入 1 或 2 的持久載體？否 → 尚未完成升級交接。

## Model / Effort 參數（實查，不憑印象）

**硬規則**：主 session 要顯式指定委派的 model 時，**只能用工具實際支援的參數**；查不到的標 `unverified`，不猜。各工具的可用參數維護在 tool adapter 層（[`../../ai-tools/agent/`](../../ai-tools/agent/)），不在本檔重述——因為它會隨工具版本漂移（見 §參數保鮮 / 本 plan Q11）。

**本環境（Claude Code，2026-07）實際驗證到的**（來源：Agent 工具 schema 本身，非記憶）：

| 參數 | 已驗證值 | 來源 |
|---|---|---|
| Agent 委派 model | `sonnet` / `opus` / `haiku` / `fable`（Agent 工具 `model` enum） | Agent 工具 schema |
| Reasoning effort | 由 agent 定義（`.claude/agents/*.md` frontmatter）決定，**非**每次呼叫傳入 | Agent 工具 schema 說明 |
| Model IDs | Opus 4.8 `claude-opus-4-8`、Sonnet 5 `claude-sonnet-5`、Haiku 4.5 `claude-haiku-4-5-20251001`、Fable 5 `claude-fable-5` | 環境 system context |

> 其他工具（Cursor / Codex / Gemini CLI 等）的 model / effort 控制：**目前標 `unverified`**，需在對應 `ai-tools/agent/<tool>.md` 實查後補；不得憑印象寫。

## 參數保鮮（避免憑印象）

Model 集合會漂移（新 model 上市、舊 model 退役）。維護規則：

- **觸發**：要在委派中指定具體 model 名 / ID，或發現 adapter 的 model 表可能過時。
- **動作**：以「工具當下實際回報 / schema 暴露的值」為準；查不到即標 `unverified`，用 tier 概念（低 / 中 / 高）委派而非硬編 model 名。
- **完成判準（二值）**：委派指令中每個具體 model 名 / ID 是否都有當次可驗證來源？否 → 降回 tier-level 委派。

## 誰會參考這裡（Inbound References）

- [`plans/active/2026-07-10-0955-final-form-feedback-execution-closure.md`](../../plans/active/2026-07-10-0955-final-form-feedback-execution-closure.md) — Workstream F / Phase 5b F2
- [`../../governance/judgment-rubrics.md`](../../governance/judgment-rubrics.md) — R1 / R4 引用本階梯
- [`README.md`](README.md) — routing steps 第 4 步（handoff）前先套用本階梯

## 與既有層的關係

- [`multi-model-handoff.md`](multi-model-handoff.md) — 管「何時 handoff + handoff packet 格式」；本檔管「失敗幾次該升 / 降哪個 tier」。handoff 時用 multi-model-handoff 的 Handoff Packet，升降決策用本檔。
- [`autonomy-routing.md`](autonomy-routing.md) — 管 cognitive state；本檔管 model tier。正交。
- [`enforcement/escalation-policy.md`](../../enforcement/escalation-policy.md) — evidence-mismatch recovery，**不換 model**；見本檔開頭互斥分工。
