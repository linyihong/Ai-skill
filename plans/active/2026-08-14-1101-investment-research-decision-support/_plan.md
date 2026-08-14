---
id: 2026-08-14-1101-investment-research-decision-support
status: in-progress
owner_layer: workflow
---

# Investment Research & Decision Support（`workflow/investment/`）

**Status**: in-progress

**Phase 0**: ✅ CLOSED / FROZEN（2026-08-14）— 定界完成；**停止繼續改 plan 架構推測**，未知項留 Q10／Q12 由 dogfood 餵。

**Phase 1**: 🟢 IN PROGRESS — 執行順序凍結為 ① theme → ② name → ③ sweep → ④ allocation → ⑤ DVA → ⑥ Q8 A/B/C；evidence 用 PASS/FAIL/MIXED 檢驗假說是否成立，不是打勾證明「有遵守」。

2026-08-14 建立。Stakeholder 定界：形狀對齊 `workflow/legal/`（intake-dispatched +
Decision Support），預設市場 **台股＋美股、主題預設 AI／semi／供應鏈**；語言跟隨本庫
預設（繁體中文，見 `enforcement/neutral-language.md`），使用者可改；允許機率化建議，
但必須綁新聞／趨勢／大神筆記等舉證；另含 **盯盤＋大神筆記巡檢** 產出流程。

**排程邊界（2026-08-14 確認）**：本 skill／workflow **只負責被呼叫時產出報告**；
何時跑由使用者在各 AI 工具的定時任務／提醒自行設定。不接券商、不下單；
agent 最多在 intake／設定說明裡**提醒**可設排程，不擁有 cron 執行權。

**資產／策略／費用＋三角色（2026-08-14 追加）**：Intake 主動詢問（或讀取使用者設定中的）
**投資策略 profile**、**現有資產／持倉**與**手續費／相關費用**；Interest Analysis 須納入費用
摩擦，再推算對使用者較有利的方案（機率化＋trade-off，非保證最優）。高利害報告（尤其
`allocation-advice`／含資產的 `position-review`）
走 [`delegation-verification-arbitration`](2026-07-08-0825-delegation-verification-arbitration-loop/_plan.md)
三角色 loop，並強制 **證據推算表**，讓報告可覆核、可閱讀。

**系統定位（2026-08-14 外部 review 採納）**：Investment **不是**「再多一個投資 skill」的終點；
它是 **Decision Support 的第二個大型 dogfood domain／實驗室**。真正要觀察的是它是否暴露
可升級成 cross-cutting primitive 的不變量（見 §Cross-domain Abstraction Hypotheses）。
**現在不新增 7 個 generic capability / route / glossary 定稿**——先 domain dogfood →
證明 invariant 跨 domain → 再抽象（對齊 DVA／ERA 的 falsification ladder 紀律）。

**Glossary Impact**: yes（**延後註冊**；dogfood 前只用 candidate 名）—
domain 詞：`investment task type`、`strategy-asset profile`、`investment DVA loop`。
abstraction **candidates**（勿過早綁 investment 措辭）：`evidence-to-decision gate`、
`decision-support lifecycle`、`uncertainty framing`（**不是**把 glossary 定死成
`probability-framed recommendation`）、`decision depth gate`、
`periodic observation/reassessment`、`source authority model`、
`knowledge/user-state boundary`、`semantic route disambiguation`（Q8 candidate）。尚未註冊於 `knowledge/glossary/ai-skill.md`。

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
  3. **交易成本／費用 profile**（股票與相關商品；缺則追問或標 provisional）：
     券商手續費／佣金、交易稅（若適用）、匯費／換匯成本、保管／管理費、ETF 費用率、
     借券／融資利息（若使用）、其他平台費用；台股／美股分開記（費率結構不同）。
  4. 市場範圍、是否允許配置建議、大神 watchlist（可從使用者設定讀）。
  使用者已提供策略／資產／費用於設定檔時：intake **確認摘要**即可，不必重問全文；缺欄位才追問。
- **「對使用者最有利」的定義（有界）**：在策略 profile＋現有資產＋**交易成本／費用**約束下，
  比較可行方案的風險／報酬／集中度／流動性／稅與匯率 friction／**手續費與相關費用淨效應**，
  輸出 **機率化情境＋trade-off＋需拍板點**。
  **不是**全域數學最優、**不是**保證報酬。
  **Interest Analysis（必做，對齊 decision-support）**須涵蓋：
  (a) 對使用者較有利的安排（含「少交易／一次建倉 vs 分批」的費用差）；
  (b) 其他方利益與現實約束（流動性、對手方可接受性）；
  (c) **費用摩擦**：來回手續費、稅、匯費、保管／管理費、融資利息等對期望淨報酬與再平衡門檻的影響；
     費用未提供時標 `provisional` 並列待查，**不得**假裝零成本。
- **Decision Support 兩趟**（對齊 cross-cutting contract）：Pass 1 provisional（含策略／資產約束下的
  方案空間）→ Research（新聞／財報／趨勢／大神筆記）→ Pass 2 機率化建議＋需使用者決策點。
- **預設市場**：台股＋美股；**預設主題焦點**：AI／半導體／光通訊／資料中心供應鏈
  （可在 intake 改）。
- **語言**：本庫可重用文件與預設對話輸出＝繁體中文；ticker／專有名詞保留英文；
  使用者明確要求時可改語言。
- **建議規則**：可給方向性／配置建議，但必須 (a) 附來源舉證、(b) 趨勢摘要、
  (c) 大神／公開研究筆記對照（若有 watchlist 命中）、(d) **機率或情境權重**呈現，
  禁止「必然漲／必買」話術。
- **資料來源邊界**：預設公開來源。可提醒使用者「若有付費資料可提供」；**僅當該資料
  當下可查到（已提供／已授權可讀）**才寫進證據表。不接付費資料商 API、不假設訂閱存在。
- **Sweep 流程**：獨立 task type `periodic-sweep`——盯 watchlist 標的＋大神筆記更新＋
  重大新聞，產出 sweep brief（非即時交易訊號）。**觸發在外**（使用者／各 AI 定時任務）；
  本 workflow 是 report producer，不是 scheduler。
- **證據推算報告（Evidence-ledger report）**：凡 Yellow／配置類產出，artifact 必須含可閱讀區塊：
  | 區塊 | 內容 |
  | --- | --- |
  | 策略／資產／費用摘要 | 本次分析所依據的 profile（去敏後；含手續費假設或 provisional） |
  | 證據表 | 主張 → 來源等級 A–D → URL／日期 → 支持／削弱 |
  | 推算鏈 | 前提 → 推論步驟 → 機率／情境權重 → 結論 |
  | 方案比較 | ≥2 方案＋trade-off（**含費用／手續費淨效應**）＋為何較利於**該使用者**策略 |
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
| `allocation-advice` | 配置／風險配置（需策略＋資產＋費用） | Allocation brief（較有利方案＋Interest／費用分析＋證據推算；預設 DVA） |
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
- [x] **Q7** 是否需要付費資料（Bloomberg 等）？→ `resolved`（2026-08-14）：
      **預設公開來源**。Intake／設定說明可**告知使用者**：若持有付費資料／訂閱，可自行提供
      摘錄、匯出或可查連結。Agent **僅在付費資料當下可查到（使用者已提供或已授權可讀）**
      時才納入證據表；不得假設有 Bloomberg 等帳號、不得要求本庫集成付費 API。
      付費證據仍走來源分級與 H6；無公開交叉驗證時不得單獨支撐最高信心建議。
- [x] **Q8** 與 legal「投資／股權」是否衝突？→ `resolved`（2026-08-14 採納外部 review）：
      **要加 routing arbitration，但不是 keyword precedence**（禁止 `investment > legal` 或
      反向固定優先）。裁決依 **decision object + task semantics**：
      - Market／securities／portfolio／allocation／earnings／sector → `investment`
      - Contract／equity rights／obligation／dispute／shareholder rights／投資協議 → `legal`
      - 語意不足 → **先 framing**（investment 的 `need-framing` 或跨域澄清），**不得**僅因
        「投資／股權」字面自動選 `investment`
      Discovery signals（ticker、合約、協議…）只是 detector 提示，**不是最終裁決權**。
      Phase 5 把 invariant 投影到 `workflow-routing.md` 歧義列＋routing signals；
      Phase 1 先用 boundary dogfood 驗證（含 mixed case）。
- [ ] **Q12** Mixed investment＋legal（公司分析＋投資協議）應走 single primary＋secondary、
      還是正式 multi-domain decomposition？`still-open`（Phase 1 Case C dogfood 後再決；
      本 plan 不預設升級 multi-route runtime）
- [x] **Q9** 是否納入策略／資產 intake＋DVA＋證據推算？→ `resolved`（2026-08-14）：是。
      配置建議缺策略或資產摘要 → blocking；高利害報告強制 DVA（可明示跳過並記錄）。
- [ ] **Q10** 策略／資產設定檔的建議 schema 與範例路徑（專案本地）？`still-open`（Phase 2/3 定稿）
- [x] **Q11** 是否把 plan 定位為 DS 實驗室＋H1–H7 觀察、暫不抽 7 個 generic？→ `resolved`
      （2026-08-14 採納外部 review）：是。優先驗證 H1 Evidence-to-Decision Gate、H2 Lifecycle、
      H3 Uncertainty Framing（與 ERA 銜接）。
- [x] **Q13** 股票手續費與相關費用是否納入利益／方案評估？→ `resolved`（2026-08-14）：
      **必須**。Interest Analysis 與 allocation／position／再平衡建議須評估手續費、稅、匯費、
      保管／管理費、融資利息等；費用未知則 provisional＋待查，禁止假設零成本。

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
| 12 | 定位為 DS case #2 實驗室；H1–H7 觀察；暫不抽 7 個 generic | ✅ 採納外部 review（2026-08-14） |
| 13 | 資料：公開預設；付費僅在使用者可提供／可查到時納入 | ✅ 同意（2026-08-14） |
| 14 | Q8：semantic routing（decision object）；禁 keyword precedence；Phase 1 含 A/B/C | ✅ 採納（2026-08-14） |
| 15 | 手續費／相關費用納入 Interest Analysis 與方案比較 | ✅ 同意（2026-08-14） |

## Phase 0 — Pre-Build Interrogation

### Phase 0.0 — Open Questions 核對

- [x] 已讀 §Open Questions
- [x] Q1–Q6 `resolved`（stakeholder 2026-08-14）
- [x] Q7–Q8／Q11 `resolved`；Q10／Q12 `still-open` 不挡 Phase 1（Q12 靠 Case C 餵）

### Phase 0.1 — Pre-Build Interrogation 六問

| 問題 | 回答 |
| --- | --- |
| `goal_scope_and_non_goals` | Goal：建立投資研究／決策輔助 domain（intake、分析分派、舉證建議、sweep 報告）。Non-goals：不下單、不接券商、不實作跨工具 cron、不集成付費資料商 API、不保證報酬、不把個人持倉寫進 canonical docs、不複製大神結論當真理。 |
| `canonical_source_owner` | `workflow/investment/execution-flow.md`＝lifecycle；`intake.md`／`risk-classification.md`／`artifact-gates.md`／sub-flows 各為主題 canonical；`analysis/investment/`＝取證方法。 |
| `projection_boundary` | Phase 5 才加 `.yaml` executable contract 並 project；Phase 1–4 不新增未 project 的 `runtime/*.yaml`。 |
| `source_of_truth_duplication_risk` | (a) legal「投資／股權」語意撞車；(b) engineering economics「investment」。緩解：**Q8 semantic routing**（decision object／task semantics），非 keyword precedence；README＋選路表歧義列。 |
| `runtime_trigger_flow_or_doc_only_reason` | Phase 1–4 doc-only + dogfood；Phase 5 才 runtime integration（route + discovery signal + yaml projection）。 |
| `validation_targets` | intake gate；無舉證不得建議；uncertainty framing；periodic-sweep；**Q8 Case A/B/C**（semantic routing vs legal）；H1–H7 觀察表。 |

### Phase 0.2 — Architecture Compatibility Preflight

| # | 檢查項目 | 結果（draft） |
|---|---------|------|
| 1 | Candidate files | `workflow/`、`analysis/`、`decision-support/`、routing-registry、cognitive-modes-discovery 存在；`workflow/investment/`、`analysis/investment/` 待建 |
| 2 | Source-of-truth | registry／runtime.db 不手改；watchlist 不進 reusable canonical |
| 3 | Layer responsibility | 執行順序→workflow；取證方法→analysis；lesson 夠了再→intelligence（本輪可不建空殼） |
| 4 | Compiler | Phase 5 才 compile/refresh |
| 5 | Linked updates | workflow README、analysis README、decision-support Instantiations、routing（Phase 5）、glossary |
| 6 | Execution decision | Q1–Q9／Q11 已收斂；Q8 已定 invariant；Q10／Q12 靠 dogfood → **可進 Phase 1** |

## Cross-domain Abstraction Hypotheses（實驗室觀察，非現在實作）

> Stakeholder／外部 review 2026-08-14：investment 的架構價值 ≥ domain 本身。
> 本節是 **Phase 1–4 dogfood 的額外觀察契約**；**禁止**據此立刻開 7 個 cross-cutting
> 實作 plan 或塞進 runtime。每一條都要回答：*這個 invariant 成立是因為投資特殊，
> 還是跨 domain 認知／治理規律？*

與既有線對齊（消費、不重造）：
[`decision-support`](../../workflow/cross-cutting/decision-support/README.md)、
[`delegation-verification-arbitration`／ERA](2026-07-08-0825-delegation-verification-arbitration-loop/_plan.md)
（Evidence constrains Decision Space；Evidence Producer ≠ Closure Authority）。

### 優先序與候選抽象名

| 優先 | Investment 表面物 | 候選抽象 | 核心 invariant |
| --- | --- | --- | --- |
| ⭐⭐⭐⭐⭐ | 無舉證不得建議 | **Evidence-to-Decision Gate** | Decision strength cannot exceed evidence strength |
| ⭐⭐⭐⭐⭐ | Pass 1 → Research → Pass 2 | **Decision Support Lifecycle** | Frame → provisional → evidence → reconcile → support → human selection（≠ 系統內部 O→R→E→V） |
| ⭐⭐⭐⭐⭐ | 機率／情境化建議 | **Uncertainty Framing** | Evidence 不足時保留 uncertainty；機率只是一種 representation |
| ⭐⭐⭐⭐ | Green／Yellow／Red | **Decision Depth Gate** | Task depth → required evidence／verification／permitted output |
| ⭐⭐⭐⭐ | periodic-sweep | **Periodic Observation／Reassessment** | Observation producer ≠ scheduler；Previous → New → Delta → Reassess |
| ⭐⭐⭐⭐ | 大神筆記來源分級 | **Source Authority／Evidence Quality** | Authority×recency×directness×corroboration（非「Expert Knowledge」） |
| ⭐⭐⭐⭐ | watchlist／持倉不進 canonical | **Knowledge／User-State Boundary** | Reusable Knowledge ≠ User State ≠ Task／Evidence／Runtime State |
| ⭐⭐⭐⭐ | 「投資」撞 legal／investment | **Semantic Route Disambiguation**（candidate） | Route by decision object／task semantics；ambiguous → framing；禁固定 keyword precedence |

### Q8 Routing invariant（定案）

```text
User request → task framing → domain object + task type → route
```

- **Investment research**（市場／標的／配置／財報／供應鏈）→ `workflow/investment`
- **Legal investment／equity**（協議／股東權利義務／爭議）→ `workflow/legal`
- **Ambiguous** → framing／澄清 decision object；**不得** `contains("投資") → investment`
- Discovery signals 輔助 detector，**不**當最終裁決
- **Mixed**（公司分析＋投資協議）→ Phase 1 Case C 觀察；是否 multi-domain 見 Q12

與系統內部 Observation→Registry→Executor→Validation **不同層**；此為 **task-semantics routing**，
升 cross-cutting 前須 legal＋investment 雙側 dogfood（同 H1–H7 紀律）。

### Responsibility surfaces（銜接 ERA）

```text
Evidence → (constrains) Claim → (supports) Recommendation → (human selection) Decision
```

四個責任面（對齊 ERA：Evidence Producer ≠ Closure Authority）：

1. Evidence Producer  
2. Claim／Analysis Producer  
3. Recommendation Producer  
4. Decision／Selection Authority（預設＝使用者）

### 刻意不做（直到 legal＋investment 雙側 dogfood 支持）

- 不把 Decision Support 立刻拆成 Intake／Evidence／Uncertainty 等獨立 route。
- 不把 glossary 定稿為 `probability-framed recommendation`（用 **uncertainty framing** 當 candidate）。
- 不把「大神」升成 Expert Knowledge primitive。
- 不實作跨工具 cron（維持 Observation producer ≠ scheduler）。

## Phase 1 — Spike / Dogfood（不建 route）

**執行順序（凍結，勿並行亂跑）**：

| # | Run | 主要測 |
| --- | --- | --- |
| ① | `need-framing` → `theme-research` | H1–H4、H6 baseline |
| ② | `name-diligence` | H6（researcher note 是否必要／是否只是 source type） |
| ③ | `periodic-sweep` | H5（delta vs 全量重寫） |
| ④ | `allocation-advice`（虛構策略／資產／費用） | H1–H4、Q13、較有利是否有約束 |
| ⑤ | DVA（接在 allocation 後；爭議結論） | Verifier 挑戰機率／證據強度 |
| ⑥ | Q8 Case A／B／C | H8；Case C 餵 Q12 |

**Evidence 形狀（每假說）**：`PASS`／`FAIL`／`MIXED`／`N/A` + Observation + Evidence + Why matters + Investment-specific or cross-domain? + Legal comparison + Candidate consequence（Promote／defer／investment-only／reject）。**禁止**只打勾證明「有遵守」。

Evidence 目錄：[`evidence/`](evidence/README.md)

- [x] ① `need-framing` → `theme-research` — 見 [`evidence/01-theme-research-cpo-optical.md`](evidence/01-theme-research-cpo-optical.md)
- [x] ② `name-diligence` — 見 [`evidence/02-name-diligence-lite.md`](evidence/02-name-diligence-lite.md)（$LITE；H6 lean-promote Source Authority）
- [x] ③ `periodic-sweep` — 見 [`evidence/03-periodic-sweep-optical-lite.md`](evidence/03-periodic-sweep-optical-lite.md)（H5 lean-promote Observation／Reassessment）
- [x] ④ `allocation-advice` — 見 [`evidence/04-allocation-advice-fictional.md`](evidence/04-allocation-advice-fictional.md)
- [ ] ⑤ DVA（allocation 後）
- [ ] ⑥ Q8 Case A／B／C
- [ ] Phase 1 總結：workflow 能跑？H1–H8 哪些升 cross-domain？Q12？

### Phase 1 跨 domain 觀察表（每 run 必填）

| Observation | Hypothesis | 記什麼 |
| --- | --- | --- |
| 建議能否回溯到 evidence？ | H1 Evidence→Decision | 主張→來源對照是否可獨立覆核 |
| evidence 強度是否限制 recommendation 強度？ | H1 Gate | 弱證據是否仍寫成強建議（失效） |
| uncertainty 是否被保留（非只數字）？ | H3 Uncertainty | 是否被壓成肯定句 |
| 不同 task 是否需要不同 depth／gate？ | H4 Depth | Green 與 allocation 的 gate 差是否合理 |
| sweep 是否能只處理 delta？ | H5 Observation | 全量重寫 vs delta／reassessment |
| 低權威 source 是否被過度放大？ | H6 Source Authority | D 級是否獨撐高信心 |
| user-specific state 是否污染 canonical？ | H7 State Boundary | 持倉／watchlist 是否誤進 reusable path |
| Pass1→Research→Pass2 是否被跳過？ | H2 Lifecycle | 是否退化成 question→answer |
| 「投資」字面是否誤吸錯 route？ | H8／Q8 Semantic routing | Case A/B 是否正確；Case C 如何分解 |
| Mixed 任務是否暴露 single-route 不足？ | Q12 | Case C → single vs multi-domain 證據 |

完成條件：去敏 dogfood notes（`evidence/`），含至少一份 allocation＋DVA 形狀樣本，
**且** H1–H7 觀察表有真實勾選／失敗紀錄（不只「workflow 能跑」）。

Phase 1 結束時要能回答兩句：
1. Investment workflow 能不能跑？
2. 哪些 cognitive／governance primitive **值得**升 cross-cutting 候選（附 legal 對照或「僅投資特殊」）？

## Phase 2 — `analysis/investment/` 方法層

- [ ] `analysis/investment/README.md`
- [ ] 供應鏈／主題拆解方法
- [ ] 來源分級（官方／監管／媒體／個人研究帳）——實作時用 **source authority** 語言，避免「Expert Knowledge」框架
- [ ] 新聞與趨勢摘要模板
- [ ] 大神筆記對照方法（引用、不同意點、时效）
- [ ] `sources-and-tools.md`（公開來源類型為預設；附「使用者可選提供付費摘錄」規則；不硬編碼易 stale URL；不集成付費 API）

## Phase 3 — `workflow/investment/` domain core

- [ ] `README.md`、`execution-flow.md`、`intake.md`
- [ ] `risk-classification.md`（Green／Yellow／Red + 無舉證不得建議）
- [ ] `strategy/` 或等同 Decision Support instantiation（四項：inventory／playbooks／verification／depth gate）
- [ ] `artifact-gates.md`（機率欄位、disclaimer、evidence-ledger、策略／資產摘要 gate、DVA 適用表）
- [ ] Sub-flows：`theme-research/`、`name-diligence/`、`allocation-advice/`、`periodic-sweep/`（其餘可薄 README）
- [ ] `allocation-advice/` 明寫：策略／資產／**費用** blocking-or-provisional intake、Interest Analysis（含手續費）、較有利方案比較、DVA brief 模板
- [ ] 連到 `workflow/software-delivery/delegated-execution.md`／plans README 三角色契約（不重寫一份平行 DVA）

## Phase 4 — Decision Support 掛接

- [ ] 更新 `workflow/cross-cutting/decision-support/README.md` Instantiations：investment = case #2（或 candidate→converged，依 dogfood）
- [ ] **Abstraction review**：依 H1–H7 dogfood 結果，決定哪些升 follow-up plan／哪些標 investment-only（**本 phase 仍不實作 7 個 generic**）
- [ ] Playbooks：配置、**交易成本／再平衡費用門檻**、主題深挖深度、何時升 Red
- [ ] Scenario：`investment-evidence-required-before-advice-v1`、`investment-probability-framing-v1`、`investment-intake-gate-v1`、
      `investment-strategy-asset-required-for-allocation-v1`、`investment-dva-required-for-allocation-v1`、
      `investment-fee-interest-analysis-required-v1`

## Phase 5 — Runtime 接線（graduation）

- [ ] `route.workflow.investment` + discovery signal
- [ ] `execution-flow.yaml` / `artifact-gates.yaml` + compile/refresh
- [ ] `knowledge/summaries/` + glossary 詞條
- [ ] Linked updates：`workflow/README.md`、`analysis/README.md`、`workflow-routing.md` **semantic 歧義列**（Q8 invariant：decision object，非 keyword precedence）
- [ ] Investment／legal discovery signals 分列（ticker／portfolio vs contract／協議）；signals ≠ final arbitration
- [ ] 若 Case C 支持 multi-domain：另開 follow-up plan（不在本 plan 偷升 multi-route runtime）
- [ ] Per-surface consumer 表填實

## Phase 6 — Close-loop

- [ ] Diff review、sanitization、commit（分 owner group）、push（需授權）、readback、clean status

## 完成條件

- [ ] Phase 1–5 完成或明確 deferred 剩餘項
- [x] Stakeholder 同意項目 1–15 已決（2026-08-14）；Q10／Q12 不挡 Phase 1 開工
- [ ] Decision Support 對 investment 的 instantiation 可引用
- [ ] 至少一份真實（去敏）allocation 或 sweep dogfood evidence
- [ ] Phase 1–4 產出 H1–H7 觀察結論（升候選／investment-only／棄）——**不**要求本 plan 內實作 generic primitives
- [ ] Runtime：若宣稱 route 可用，則 trigger flow 與 consumer 表完整；否則維持 doc-only 且不宣稱 integration 完成

## Watch-Out List

- 防把個人持倉／券商帳號寫進 canonical docs（sanitization）。
- 防把單一 X 帳號結論當 source of truth（來源分級 D 不可獨撐）。
- 防未 dogfood 就註冊 route（對齊 legal／decision-support 三案紀律）。
- 防把本庫做成跨工具 scheduler（排程屬使用者／各 AI 定時任務）。
- 防把「最有利」寫成保證最優／保證報酬；必須綁策略約束＋uncertainty framing。
- 防配置／再平衡建議假設零手續費或忽略稅／匯費／保管費（Q13）。
- 防 DVA 形式化：Verifier 只複讀 Executor 結論而無 L2 證據鏈檢查。
- 防 Phase 1 還沒跑就開 7 個 cross-cutting primitive 實作（實驗室 ≠ 立刻建框架）。
- 防把 Decision Support Lifecycle 與系統內部 Observation→Registry→Executor→Validation 混層。
- 防 `contains("投資") → investment` 或固定 `investment > legal` keyword precedence（Q8）。
- 防 Case C 未 dogfood 就宣稱已支援 multi-domain routing。
- 防 glossary 被 `probability-framed recommendation` 綁死（用 uncertainty framing candidate）。
- 防 scope 膨脹成「全市場量化平台」。

## 與其他 plans 的關係

- **依賴／對齊**：[`2026-07-30-2101-legal-workflow-domain.md`](2026-07-30-2101-legal-workflow-domain.md)（intake-dispatched + Decision Support case #1）；`workflow/cross-cutting/decision-support/`（三案門檻）；[`2026-07-08-0825-delegation-verification-arbitration-loop`](2026-07-08-0825-delegation-verification-arbitration-loop/_plan.md)（O→E→V→仲裁＋ERA；本 domain 消費並用 H1 對照 Evidence→Claim→Recommendation→Human Decision）。
- **定位**：Investment = Decision Support instantiation #2 **＋** cross-domain abstraction extraction opportunity（見 §Cross-domain Abstraction Hypotheses）。
- **不取代**：legal Red tier 的「股權／投資契約」法律任務仍走 `route.workflow.legal`。
- **參考外部風格**：公開研究帳如 [Serenity (@aleabitoreddit)](https://x.com/aleabitoreddit) — 學供應鏈／choke-point 敘事與更新節奏，**不**把其結論寫進本庫當事實。
