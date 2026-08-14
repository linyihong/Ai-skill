---
id: 2026-08-14-1101-investment-research-decision-support
status: draft
owner_layer: workflow
---

# Investment Research & Decision Support（`workflow/investment/`）

**Status**: draft（Phase 0 定界已收斂；可進 Phase 1 dogfood）

2026-08-14 建立。Stakeholder 定界：形狀對齊 `workflow/legal/`（intake-dispatched +
Decision Support），預設市場 **台股＋美股、主題預設 AI／semi／供應鏈**；語言跟隨本庫
預設（繁體中文，見 `enforcement/neutral-language.md`），使用者可改；允許機率化建議，
但必須綁新聞／趨勢／大神筆記等舉證；另含 **盯盤＋大神筆記巡檢** 產出流程。

**排程邊界（2026-08-14 確認）**：本 skill／workflow **只負責被呼叫時產出報告**；
何時跑由使用者在各 AI 工具的定時任務／提醒自行設定。不接券商、不下單；
agent 最多在 intake／設定說明裡**提醒**可設排程，不擁有 cron 執行權。

**資產／策略＋三角色（2026-08-14 追加）**：Intake 主動詢問（或讀取使用者設定中的）
**投資策略 profile** 與 **現有資產／持倉**；在約束下推算對使用者較有利的方案（機率化＋
trade-off，非保證最優）。高利害報告（尤其 `allocation-advice`／含資產的 `position-review`）
走 [`delegation-verification-arbitration`](2026-07-08-0825-delegation-verification-arbitration-loop/_plan.md)
三角色 loop，並強制 **證據推算表**，讓報告可覆核、可閱讀。

**Glossary Impact**: yes — 預期引入（Phase 實作時註冊）：`investment task type`、
`probability-framed recommendation`、`expert-note watchlist`、`periodic market sweep`、
`strategy-asset profile`、`evidence-ledger report`、`investment DVA loop`。
尚未註冊於 `knowledge/glossary/ai-skill.md`。

## Decision Rationale

### Problem & Why Now

使用者要「投資相關幫忙」，但痛點不是直接喊單，而是：

1. **需求沒被問清楚**：類似法律，使用者常只丟一句「看一下 XX」或「幫我看看配置」，
   沒有 horizon、風險承受、現金用途、持倉現況，分析會偏題。
2. **分析路徑要被設計**：要先決定是主題研究、單一標的 diligence、持倉複核，還是配置建議。
3. **建議需要舉證與機率**：新聞、趨勢、公開研究／大神筆記都要進 artifact；結論用機率／
   情境呈現，避免說死。
4. **需要定期 sweep**：不是單次對話才有用——要有可重複的「盯盤＋大神筆記巡檢」流程。
5. **策略與現有資產未進場**：沒有策略／持倉約束的「最有利方案」只是空泛建議；必須先問清或讀設定。
6. **單人產出可信度不足**：配置類報告需要獨立驗證與仲裁，並把證據推算寫進可讀 artifact。

Repo 現況：`workflow/cross-cutting/decision-support/` 已把 **investment（未來）** 列為
候選；尚無 `workflow/investment/` / `analysis/investment/`。Legal 已是 Decision Support
converged case #1；本 domain 目標成為 **case #2**。

### Decision

新增 **`workflow/investment/`**（執行順序）＋ **`analysis/investment/`**（取證方法），
形狀為 **intake-dispatched workflow**：

- **Domain 邊界＝投資研究與決策輔助**，不是交易執行系統。不接券商 API、不下單。
- **第一級分派＝investment task type**（見下表）。
- **Intake P0 問題**（缺則 blocking，不得進入配置建議）：
  1. **策略 profile**：目標（保值／成長／主題曝險等）、horizon、風險承受、再平衡規則、
     禁止標的／槓桿偏好、稅／匯率約束（若有）。
  2. **現有資產**：現金、持倉清單（標的／成本／權重或約略部位）、負債／流動性需求（若有）。
  3. 市場範圍、是否允許配置建議、大神 watchlist（可從使用者設定讀）。
  使用者已提供策略／資產於設定檔時：intake **確認摘要**即可，不必重問全文；缺欄位才追問。
- **「對使用者最有利」的定義（有界）**：在策略 profile＋現有資產約束下，比較可行方案的
  風險／報酬／集中度／流動性／稅與匯率 friction，輸出 **機率化情境＋trade-off＋需拍板點**。
  **不是**全域數學最優、**不是**保證報酬。Interest Analysis 必須同時看「市場對手／流動性現實」
  （對齊 decision-support 不得 one-sided interest）。
- **Decision Support 兩趟**（對齊 cross-cutting contract）：Pass 1 provisional（含策略／資產約束下的
  方案空間）→ Research（新聞／財報／趨勢／大神筆記）→ Pass 2 機率化建議＋需使用者決策點。
- **預設市場**：台股＋美股；**預設主題焦點**：AI／半導體／光通訊／資料中心供應鏈
  （可在 intake 改）。
- **語言**：本庫可重用文件與預設對話輸出＝繁體中文；ticker／專有名詞保留英文；
  使用者明確要求時可改語言。
- **建議規則**：可給方向性／配置建議，但必須 (a) 附來源舉證、(b) 趨勢摘要、
  (c) 大神／公開研究筆記對照（若有 watchlist 命中）、(d) **機率或情境權重**呈現，
  禁止「必然漲／必買」話術。
- **Sweep 流程**：獨立 task type `periodic-sweep`——盯 watchlist 標的＋大神筆記更新＋
  重大新聞，產出 sweep brief（非即時交易訊號）。**觸發在外**（使用者／各 AI 定時任務）；
  本 workflow 是 report producer，不是 scheduler。
- **證據推算報告（Evidence-ledger report）**：凡 Yellow／配置類產出，artifact 必須含可閱讀區塊：
  | 區塊 | 內容 |
  | --- | --- |
  | 策略／資產摘要 | 本次分析所依據的 profile（去敏後） |
  | 證據表 | 主張 → 來源等級 A–D → URL／日期 → 支持／削弱 |
  | 推算鏈 | 前提 → 推論步驟 → 機率／情境權重 → 結論 |
  | 方案比較 | ≥2 方案＋trade-off＋為何較利於**該使用者**策略 |
  | Verifier 摘要 | 若跑 DVA：findings 與仲裁結果 |
- **三角色 DVA（delegation → verification → arbitration）**：
  對齊 [`2026-07-08-0825-delegation-verification-arbitration-loop`](2026-07-08-0825-delegation-verification-arbitration-loop/_plan.md)。
  | 情境 | 是否強制 DVA |
  | --- | --- |
  | Green 名詞／地圖整理 | 否（可單 session） |
  | Yellow `theme-research`／`name-diligence`／`event-check`／`periodic-sweep` | **建議**；使用者要求「可看／可驗證報告」時強制 |
  | Yellow/Red 邊界之 `allocation-advice`、含資產之 `position-review` | **強制**（除非使用者明示跳過並記錄） |
  角色：Orchestrator（問需求、寫 brief、仲裁）／Executor（取證＋草稿報告）／Verifier
  （獨立覆核證據鏈、機率是否說死、策略約束是否被違反）。Verifier **只產 evidence findings**，
  不改建議；Orchestrator 仲裁 `fix`／`defer`／`reject` 後才定稿。

#### Investment task types（intake 第一級）

| Task type | 何時 | 主要產出 |
| --- | --- | --- |
| `need-framing` | 還不知道要分析什麼 | 澄清後的分析題＋建議路徑 |
| `theme-research` | 產業／供應鏈主題 | Thesis note＋供應鏈圖＋catalyst |
| `name-diligence` | 單一標的 | Diligence card＋證據表 |
| `position-review` | 已持倉 | Thesis still-valid?＋失效條件 |
| `allocation-advice` | 配置／風險配置（需策略＋資產） | Allocation brief（對使用者較有利方案＋證據推算；預設 DVA） |
| `event-check` | 單次新聞／事件 | Event impact card |
| `periodic-sweep` | 定期巡檢 | Sweep brief（標的＋大神筆記＋新聞） |

#### Risk / depth gate（草案）

| Tier | 例子 | Agent 行為 |
| --- | --- | --- |
| Green | 名詞、產業地圖、公開來源整理 | 可產出 research note |
| Yellow | 單一公司／配置情境／機率化建議 | 必須舉證＋標 confidence；建議用機率／情境 |
| Red | 保證報酬、無來源喊單、 lev/期權複雜結構未核實 | **停止實質交易指令**；改為 Escalation／待查清單 |

> Stakeholder 2026-08-14：**允許建議**，但必須新聞等舉證、趨勢分析、大神筆記對照，
> 且以機率呈現。因此本 domain **不採**「永不給方向建議」；改採「無舉證不得建議」＋
> 「禁止說死」。

### Alternatives Considered

- **A. 只做 theme-research（Serenity 筆記克隆）**：reject — 使用者明確要 intake＋配置建議＋定期盯盤。
- **B. 純 chat prompt、不建 workflow**：reject — 無法 gate、無法與 Decision Support 三案門檻對齊。
- **C. 接券商／自動下單**：reject — 超出授權與安全邊界；本 plan 明確 non-goal。
- **D. intake-dispatched + analysis 取證 + periodic-sweep（partial）** — 保留，但不足以涵蓋策略／資產與報告可信度。
- **E. D + strategy/asset intake + 有界「較有利方案」+ DVA + evidence-ledger（accept）**。

### Why Not an ADR Yet

Intake-dispatched 與 Decision Support 已有 legal 先例與 cross-cutting pilot；investment
是否成為穩定 case #2 需 dogfood。未驗證前不升 ADR。

### ADR Promotion Criteria（completed 時驗證）

- [ ] 至少 3 個真實任務（含 1 個含策略／資產之 `allocation-advice`、1 個 `periodic-sweep`）跑完 lifecycle
- [ ] 至少 1 次完整 O→E→V→仲裁（investment DVA）且 evidence-ledger 齊備
- [ ] Decision Support instantiation 四項齊備，且 cross-cutting 計數變 2／3
- [ ] Open Questions 收斂或明確 deferred
- [ ] 機率化建議＋舉證 gate＋「無策略／資產不得給配置建議」scenario 覆蓋

### Consequences

#### 正面
- 投資對話先問清楚再分析，減少偏題；策略／資產進場後「較有利方案」才有意義。
- 建議可執行但仍可追溯來源與機率；evidence-ledger 提高可讀性與可覆核性。
- DVA 降低單人幻覺與「說死」風險。
- 定期 sweep 讓「大神筆記＋新聞」變成可重複流程，而非一次性搜尋。
- 推進 Decision Support 三案門檻（legal → investment → ?）。

#### 負負面／風險
- 投資內容易被當成投顧意見 → 強制 disclaimer + 機率框架 + 無舉證不得建議。
- 大神筆記來源品質參差 → 來源分級（A 官方／B 監管揭露／C 主流媒體／D 個人研究帳號），
  D 級不得單獨支撐高信心建議。
- 定期 sweep 若寫進本庫 watchlist 會混入個人持倉 → **watchlist／持倉留專案本地或
  `.agent-goals/`／使用者指定路徑**，不進 canonical reusable docs（sanitization）。

## Runtime Execution Path

| 環節 | 內容 |
| --- | --- |
| Runtime owner | 預定：`knowledge/runtime/routing-registry.yaml` §`route.workflow.investment`（**Phase 5 才註冊**；Phase 1–4 dogfood 前不註冊） |
| Event / signal | 投資研究／配置／盯盤／大神筆記等（Phase 5 定義 user_signals） |
| Detector | 既有 `DetectWorkflows`；本 plan 不改 Go（除非 dogfood 證明字面衝突） |
| Loaded contract | `workflow/investment/execution-flow.md` + required_dependencies |
| Discovery signal | Phase 5：`cognitive-modes-discovery.yaml` 新增投資關鍵字 → 傾向 `SOURCE_BACKED` |
| Doc-only trial 宣告 | Phase 1–4 為 **doc + dogfood trial**，**不宣稱已完成 runtime integration**；graduation 條件見 Phase 5 |

### Per-surface consumer 表（Phase 5 填實；draft 預留）

| Generated surface key | Named consumer(s) | Consumer 類型 |
| --- | --- | --- |
| `route.workflow.investment`（預定） | `DetectWorkflows` + primary_source gate + discovery signal | discovery + Go validator |
| `workflow.investment.execution_flow.contract`（預定） | routable lookup + `validateRuntimeYamlProjects` | routable lookup + Go validator |
| `workflow.investment.artifact_gates.contract`（預定） | 同上 | routable lookup + Go validator |

## Open Questions

- [x] **Q1** 市場範圍？→ `resolved`：台股＋美股；主題預設 AI。
- [x] **Q2** 語言？→ `resolved`：本庫預設繁體中文；使用者可改。
- [x] **Q3** 可否給買賣／配置建議？→ `resolved`：可給建議，須舉證＋趨勢＋大神筆記對照，機率呈現、避免說死。
- [x] **Q4** 大神／研究帳 watchlist 初值？→ `resolved`（2026-08-14）：可納入；初值含
      [@aleabitoreddit](https://x.com/aleabitoreddit)；其餘帳號由使用者設定擴充。來源分級仍為
      D（個人研究帳），不得單獨支撐高信心建議。更新頻率＝sweep 被呼叫時重抓，不由本庫 cron。
- [x] **Q5** `periodic-sweep` 觸發方式？→ `resolved`（2026-08-14）：**排程在使用者／各 AI
      工具的定時任務**；本 skill／workflow 只在被呼叫時產出報告。Intake／設定說明可**提醒**
      使用者自行設排程；本庫不實作跨工具 scheduler。
- [x] **Q6** 個人持倉與 watchlist 存放？→ `resolved`（2026-08-14）：**使用者設定**（專案本地
      或工具側設定）；skill 執行時讀取該設定。不寫入 Ai-skill canonical reusable docs。
- [ ] **Q7** 是否需要付費資料（Bloomberg 等）？預設只用公開來源？`deferred`（預設公開來源 only；
      付費源另開 plan）
- [ ] **Q8** 與 legal Red tier「投資／股權」字面是否衝突？選路表是否要加裁決？`still-open`
      （Phase 5 routing 時處理；不挡 Phase 1）
- [x] **Q9** 是否納入策略／資產 intake＋DVA＋證據推算？→ `resolved`（2026-08-14）：是。
      配置建議缺策略或資產摘要 → blocking；高利害報告強制 DVA（可明示跳過並記錄）。
- [ ] **Q10** 策略／資產設定檔的建議 schema 與範例路徑（專案本地）？`still-open`（Phase 2/3 定稿）

## Stakeholder 同意項目

| # | 項目 | 狀態 |
| --- | --- | --- |
| 1 | 形狀＝legal 式 intake → Decision Support → Research → 建議 | ✅ 同意（2026-08-14） |
| 2 | 含配置建議（allocation-advice） | ✅ 同意 |
| 3 | 含盯盤＋大神筆記巡檢（periodic-sweep 產出） | ✅ 同意 |
| 4 | 市場預設：台＋美、主題 AI | ✅ 同意 |
| 5 | 語言：本庫預設繁中，可改 | ✅ 同意 |
| 6 | 建議須舉證＋趨勢＋大神筆記；機率呈現 | ✅ 同意 |
| 7 | 不接下單／券商 API；僅提醒可設排程 | ✅ 同意（2026-08-14） |
| 8 | Q4 大神 watchlist 初值含 @aleabitoreddit，可擴充 | ✅ 同意 |
| 9 | Q5／Q6：觸發＝各 AI 定時；設定＝使用者側；workflow＝報告產出 | ✅ 同意 |
| 10 | Intake 詢問／讀取策略 profile＋現有資產；據此推算較有利方案（機率化） | ✅ 同意（2026-08-14） |
| 11 | 高利害報告走 DVA 三角色＋證據推算表 | ✅ 同意（2026-08-14） |

## Phase 0 — Pre-Build Interrogation

### Phase 0.0 — Open Questions 核對

- [x] 已讀 §Open Questions
- [x] Q1–Q6 `resolved`（stakeholder 2026-08-14）
- [x] Q7 `deferred`（公開來源 only）；Q8 `still-open` 不挡 Phase 1

### Phase 0.1 — Pre-Build Interrogation 六問

| 問題 | 回答 |
| --- | --- |
| `goal_scope_and_non_goals` | Goal：建立投資研究／決策輔助 domain（intake、分析分派、舉證建議、sweep 報告）。Non-goals：不下單、不接券商、不實作跨工具 cron、不保證報酬、不把個人持倉寫進 canonical docs、不複製大神結論當真理。 |
| `canonical_source_owner` | `workflow/investment/execution-flow.md`＝lifecycle；`intake.md`／`risk-classification.md`／`artifact-gates.md`／sub-flows 各為主題 canonical；`analysis/investment/`＝取證方法。 |
| `projection_boundary` | Phase 5 才加 `.yaml` executable contract 並 project；Phase 1–4 不新增未 project 的 `runtime/*.yaml`。 |
| `source_of_truth_duplication_risk` | (a) legal Red「投資」字面；(b) intelligence/engineering/economics「investment」工程投資語意。緩解：route signals 用股市／配置／ticker 等專屬詞；README 明寫邊界。 |
| `runtime_trigger_flow_or_doc_only_reason` | Phase 1–4 doc-only + dogfood；Phase 5 才 runtime integration（route + discovery signal + yaml projection）。 |
| `validation_targets` | intake gate；無舉證不得建議；機率框架強制；periodic-sweep artifact；routing 不與 legal／software-delivery 誤吸。 |

### Phase 0.2 — Architecture Compatibility Preflight

| # | 檢查項目 | 結果（draft） |
|---|---------|------|
| 1 | Candidate files | `workflow/`、`analysis/`、`decision-support/`、routing-registry、cognitive-modes-discovery 存在；`workflow/investment/`、`analysis/investment/` 待建 |
| 2 | Source-of-truth | registry／runtime.db 不手改；watchlist 不進 reusable canonical |
| 3 | Layer responsibility | 執行順序→workflow；取證方法→analysis；lesson 夠了再→intelligence（本輪可不建空殼） |
| 4 | Compiler | Phase 5 才 compile/refresh |
| 5 | Linked updates | workflow README、analysis README、decision-support Instantiations、routing（Phase 5）、glossary |
| 6 | Execution decision | Q1–Q6 已收斂、Q7 deferred、Q8 不挡 → **可進 Phase 1 dogfood** |

## Phase 1 — Spike / Dogfood（不建 route）

- [ ] 選 1 個 AI／semi 主題跑 `need-framing` → `theme-research` 全流程（對話即可）
- [ ] 選 1 檔標的跑 `name-diligence`（新聞＋趨勢＋至少 1 則大神／公開研究對照）
- [ ] 模擬一次 `periodic-sweep` 產出形狀
- [ ] 用**去敏虛構**策略 profile＋資產表跑一次 `allocation-advice` 草稿（含方案比較＋機率）
- [ ] 模擬 DVA：寫 brief → Executor 產 evidence-ledger 報告 → Verifier findings 表（可同人分角色或 Task）
- [ ] 記錄 failure／摩擦 → 回寫本 plan Open Questions

完成條件：去敏 dogfood notes（可放本 plan `evidence/`），含至少一份 allocation＋DVA 形狀樣本。

## Phase 2 — `analysis/investment/` 方法層

- [ ] `analysis/investment/README.md`
- [ ] 供應鏈／主題拆解方法
- [ ] 來源分級（官方／監管／媒體／個人研究帳）
- [ ] 新聞與趨勢摘要模板
- [ ] 大神筆記對照方法（引用、不同意點、时效）
- [ ] `sources-and-tools.md`（公開來源類型；不硬編碼易 stale 的 URL 清單為真理）

## Phase 3 — `workflow/investment/` domain core

- [ ] `README.md`、`execution-flow.md`、`intake.md`
- [ ] `risk-classification.md`（Green／Yellow／Red + 無舉證不得建議）
- [ ] `strategy/` 或等同 Decision Support instantiation（四項：inventory／playbooks／verification／depth gate）
- [ ] `artifact-gates.md`（機率欄位、disclaimer、evidence-ledger、策略／資產摘要 gate、DVA 適用表）
- [ ] Sub-flows：`theme-research/`、`name-diligence/`、`allocation-advice/`、`periodic-sweep/`（其餘可薄 README）
- [ ] `allocation-advice/` 明寫：策略／資產 blocking intake、較有利方案比較、DVA brief 模板
- [ ] 連到 `workflow/software-delivery/delegated-execution.md`／plans README 三角色契約（不重寫一份平行 DVA）

## Phase 4 — Decision Support 掛接

- [ ] 更新 `workflow/cross-cutting/decision-support/README.md` Instantiations：investment = case #2（或 candidate→converged，依 dogfood）
- [ ] Playbooks：配置、主題深挖深度、何時升 Red
- [ ] Scenario：`investment-evidence-required-before-advice-v1`、`investment-probability-framing-v1`、`investment-intake-gate-v1`、
      `investment-strategy-asset-required-for-allocation-v1`、`investment-dva-required-for-allocation-v1`

## Phase 5 — Runtime 接線（graduation）

- [ ] `route.workflow.investment` + discovery signal
- [ ] `execution-flow.yaml` / `artifact-gates.yaml` + compile/refresh
- [ ] `knowledge/summaries/` + glossary 詞條
- [ ] Linked updates：`workflow/README.md`、`analysis/README.md`、`workflow-routing.md` 歧義列（vs legal「投資」）
- [ ] Per-surface consumer 表填實

## Phase 6 — Close-loop

- [ ] Diff review、sanitization、commit（分 owner group）、push（需授權）、readback、clean status

## 完成條件

- [ ] Phase 1–5 完成或明確 deferred 剩餘項
- [x] Stakeholder 同意項目 1–11 已決（2026-08-14）；Q7 deferred／Q8／Q10 不挡 Phase 1
- [ ] Decision Support 對 investment 的 instantiation 可引用
- [ ] 至少一份真實（去敏）allocation 或 sweep dogfood evidence
- [ ] Runtime：若宣稱 route 可用，則 trigger flow 與 consumer 表完整；否則維持 doc-only 且不宣稱 integration 完成

## Watch-Out List

- 防把個人持倉／券商帳號寫進 canonical docs（sanitization）。
- 防把單一 X 帳號結論當 source of truth（來源分級 D 不可獨撐）。
- 防未 dogfood 就註冊 route（對齊 legal／decision-support 三案紀律）。
- 防把本庫做成跨工具 scheduler（排程屬使用者／各 AI 定時任務）。
- 防把「最有利」寫成保證最優／保證報酬；必須綁策略約束＋機率／情境。
- 防 DVA 形式化：Verifier 只複讀 Executor 結論而無 L2 證據鏈檢查。
- 防 scope 膨脹成「全市場量化平台」。

## 與其他 plans 的關係

- **依賴／對齊**：[`2026-07-30-2101-legal-workflow-domain.md`](2026-07-30-2101-legal-workflow-domain.md)（intake-dispatched + Decision Support case #1）；`workflow/cross-cutting/decision-support/`（三案門檻）；[`2026-07-08-0825-delegation-verification-arbitration-loop`](2026-07-08-0825-delegation-verification-arbitration-loop/_plan.md)（O→E→V→仲裁契約，本 domain 消費不另造輪子）。
- **不取代**：legal Red tier 的「股權／投資契約」法律任務仍走 `route.workflow.legal`。
- **參考外部風格**：公開研究帳如 [Serenity (@aleabitoreddit)](https://x.com/aleabitoreddit) — 學供應鏈／choke-point 敘事與更新節奏，**不**把其結論寫進本庫當事實。
