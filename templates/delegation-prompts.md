# 委派派工模板（Delegation Prompt Templates）

本檔提供**五種常見任務型態**的委派 prompt 填空模板，給主 session（orchestrator）直接套用。每份都內建**派工三件套**（目標與動機 / 驗收條件 / 回報格式），與 delegation `brief` schema 對齊（`brief.goal` / `brief.acceptance` / `brief.verification`）。

> **與既有 kit 的關係**：三角色 loop 的 canonical 契約在 [`../plans/README.md`](../plans/README.md) §Delegation；executor/verifier/仲裁的傳輸模板在 [`../plans/active/2026-07-08-0825-delegation-verification-arbitration-loop/01-dogfood-prompt-kit.md`](../plans/active/2026-07-08-0825-delegation-verification-arbitration-loop/01-dogfood-prompt-kit.md)。**本檔不重複那些**，只補「按任務型態分化的填空模板」——搜尋 / 實作 / 重構 / 研究 / 審查各一，方便主 session 選型直接貼。
>
> **模板刻意 self-contained**（貼進 fresh session 的人不讀本 repo）——這是 transport 本質，不是 dual source。

## 共用規則（五型都適用）

1. **回報合約**：subagent 只回**結論 + 檔案:行號**；長產物（完整清單 / 大 diff / 報告全文）**落檔**再傳路徑，不要整段貼回主對話（省 token、可覆核）。
2. **model tier**：依 [`../models/routing/model-tier-escalation.md`](../models/routing/model-tier-escalation.md) 選 tier；搜尋 / 批次套用用低 tier，實作 / 重構用中 tier，研究 / 審查 / 對抗性驗證用中～高 tier。
3. **驗證不自驗**：executor 交付後，驗收由 fresh-context agent 做（見 [`../plans/README.md`](../plans/README.md) §Delegation Verifier）。
4. **Ai-skill repo 內委派**：`context.required` 須含 `CORE_BOOTSTRAP.md` + `runtime/core-bootstrap.yaml`，executor 首則回覆須輸出 Bootstrap Receipt（格式：`Bootstrap: rules=✓ phase=<phase-id> obligations=<n> gates=<n>` 加一行 `Active per-turn obligations: <ids>`；值用 `ai-skill runtime receipt` 取得，完整定義見 [`../CORE_BOOTSTRAP.md`](../CORE_BOOTSTRAP.md)）。外部 repo 委派無此需求。

---

## T1 — 搜尋 / 盤點（Search / Inventory）

**何時用**：要在多檔 / 多目錄掃描、找命名慣例、盤點某類現象，只需要結論不需要過程。派低 tier（讀取密集、判斷輕）。

```text
你是搜尋執行者。目標是盤點，不是修改——不要改任何檔案。

## 目標與動機
盤點 <範圍，例如 "所有 enforcement/*.md"> 中 <要找什麼，例如 "使用未定義觸發詞（適當/必要/重要）的規則">。
動機：<為什麼要找，例如 "F1 弱模型可讀性標準要清查違規清單">。

## 驗收條件
1. 涵蓋 <範圍> 全部檔案（不遺漏；列出實際掃描的檔數）。
2. 每個命中附「檔案:行號 + 命中的具體字串」，可獨立覆核。
3. 零命中也是有效結論（明確回報「無命中」，不要編造）。

## 回報格式
- 主對話只回：命中總數 + 前 3 個範例（檔案:行號）。
- 完整清單落檔到 <路徑，例如 scratchpad/inventory.md>，回傳路徑。
```

- **正例**：回「12 命中，範例：`enforcement/foo.md:23`（"適當時"）…完整清單見 `scratchpad/inventory.md`」。
- **反例**：把 12 條命中全文貼回主對話（洗版 + 難覆核），或宣稱「大致掃過沒問題」（未列掃描檔數，不可覆核）。

---

## T2 — 實作（Implementation）

**何時用**：明確 contract 下新增功能 / 檔案。派中 tier。

```text
你是實作執行者。只憑這份 brief 完成，不需找其他規劃文件。

## 目標與動機
實作 <要做什麼>。動機：<為什麼現在做>。

## 必讀 context（context.required）
- <路徑1>（<為什麼要讀>）
- <路徑2>

## 驗收條件（acceptance — 做到什麼算完成）
1. <可觀察、可驗證的條目>
2. <...>
（每條標 executor=happy-path 自驗 / verifier_only=負面·架構，供後續獨立驗證）

## 自驗底線（verification — 你交付前至少要跑）
- <命令，例如 `cd scripts/ai-skill-cli && go test ./...`>，貼結果。

## 回報格式
- 主對話：完成的 acceptance 條目對照 + 自驗結果（通過/失敗）+ branch/diff 位置。
- 不要宣稱「應該會過」——未實跑的不算完成（見 judgment-rubrics R2）。
```

- **正例**：回「acceptance 1–3 完成，`go test` 全綠（貼輸出），diff 在 branch X」。
- **反例**：回「實作好了，測試應該沒問題」→ 違反 R2 第 2 層（未實跑覆核）。

---

## T3 — 重構（Refactor）

**何時用**：改既有程式碼 / 文件結構但**不改外部行為**。派中 tier。重點是控制 diff 純度。

```text
你是重構執行者。行為必須不變，只改結構。

## 目標與動機
重構 <目標>，使 <改善點>。**外部行為/介面不變**。

## 必讀 context
- <被重構的檔 + 其測試 + 呼叫方>

## 驗收條件
1. 重構前後，<既有測試套件> 全綠且測試未被修改（行為不變的證據）。
2. diff 只含結構改動，無夾帶新功能（surgical：見 workflow surgical-changes）。
3. <改善點可量測，例如 "單檔行數 < X" / "重複邏輯消除">。

## 自驗底線
- 跑既有測試貼結果；`git diff` 自檢無「順便」改動。

## 回報格式
- 主對話：測試前後對照 + diff 統計（檔數/行數）+ 有無夾帶（明確回報）。
```

- **正例**：回「既有 42 測試全綠且未動，diff 僅 3 檔結構調整，無新功能」。
- **反例**：重構時「順便」加了個小功能 / 改了個 unrelated bug → 違反 surgical diff 純度，驗收該退回。

---

## T4 — 研究 / 調查（Research）

**何時用**：開放問題、需要收集證據形成判斷（非直接改檔）。派中～高 tier。**產出是證據，不是決定**。

```text
你是研究執行者。產出證據與選項，不做最終決定（決定由主 session 做）。

## 目標與動機
調查 <問題>。動機：<這個判斷會影響什麼決定>。

## 邊界
- 只收集證據 + 列選項 + 各 trade-off；不要替使用者拍板。
- 每個結論附來源（檔案:行號 / URL / 實測輸出），不可憑印象。

## 驗收條件
1. 問題被拆成可回答的子問題，各有證據。
2. 至少 2 個選項，各附 trade-off 與適用條件。
3. 明確標「已驗證」vs「假設 / 待驗證」，不混為一談。

## 回報格式
- 主對話：一句話結論傾向 + 選項對照表。
- 完整證據落檔到 <路徑>，回傳路徑。
- 查不到的明確標「未找到」，不編造（誠實條款）。
```

- **正例**：回「傾向選項 B；A/B/C 對照見 `scratchpad/research.md`；C 的成本數據未找到，已標明」。
- **反例**：直接回「我建議用 X」而不附證據與替代方案 → 越界替使用者決定，且不可覆核。

---

## T5 — 審查 / 驗證（Review / Verification）

**何時用**：對 executor 產出或既有變更做獨立審查。派中～高 tier。**fresh context，不可沿用 executor 的 session**。

```text
你是獨立驗證者（fresh context）。你沒有參與實作，只審查。

## 輸入
- brief 的 acceptance（主量尺）：<貼 acceptance>
- 被審對象：<branch / diff / 檔案位置>

## 三層驗證（都要做，只做 L1 不足）
- L1 重跑：實際執行 <verification 命令>，貼結果。
- L2 讀碼：讀 diff，對照 acceptance 與架構禁止事項（grep / 靜態檢查）。
- L3 對抗性：執行或補寫負面 / 邊界 case（brief 標 verifier_only 的條目）。

## 驗收條件（你的報告要滿足）
每條 finding 附四欄：evidence（可覆核）/ acceptance_ref（對應條目或 beyond-acceptance）/ classification（acceptance-violation / out-of-scope / observation）/ status（observed / verified / refuted）。

## 回報格式
- 主對話：findings 四欄表 + 一句話總評（通過 / 有 violation）。
- 只產證據，不做「要不要修」的決定（那是 orchestrator 仲裁）。
```

- **正例**：回四欄 findings 表，L1 重跑貼輸出，L3 補了 2 個負面 case。
- **反例**：只重跑 executor 自寫測試（L1）就回「看起來沒問題」→ 降級為複讀自驗，違反三層驗證。

## 誰會參考這裡（Inbound References）

- [`../plans/active/2026-07-10-0955-final-form-feedback-execution-closure.md`](../plans/active/2026-07-10-0955-final-form-feedback-execution-closure.md) — Workstream F / Phase 5b F4
- [`../plans/README.md`](../plans/README.md) — §Delegation canonical 契約
- [`../models/routing/model-tier-escalation.md`](../models/routing/model-tier-escalation.md) — 每型的 tier 選擇

## 與既有層的關係

- delegation dogfood kit（plan-local transport）管 executor/verifier/仲裁三角色 loop；本檔管**按任務型態分化的填空**。兩者互補，不重複契約。
