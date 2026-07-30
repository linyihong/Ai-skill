---
id: 2026-07-30-2101-legal-workflow-domain
status: in-progress
owner_layer: workflow
---

# Legal Workflow Domain（`workflow/legal/`）

**Status**: in-progress（implementation closed）

2026-07-30 建立；同日依 stakeholder review 兩次修訂：(1) 從 `legal-contract` 提升為 `legal`；
(2) 新增 `strategy/` slice 與 Stage 3a／3b，把 workflow 從 execution 提升為 legal decision
support；(3) 抽出 `workflow/cross-cutting/decision-support/` cross-domain pattern。

### 為什麼還不 archive（2026-07-30 結案評估）

Phase 1–7 全部完成（實作、runtime 接線、linked updates、commit + push、clean status），
但**不搬移至 `archived/`**，理由三項：

| # | 阻擋項 | 解除條件 |
| --- | --- | --- |
| 1 | §Stakeholder 同意項目有 **4 項 ⏳ 待使用者確認**（Strategy 兩趟、Optimization Suggestion 界線、三案門檻不推廣、`lifecycle/` 不接簽署工具）。archive 會把「待簽核」靜默轉成「已決定」。 | 使用者逐項確認或否決 |
| 2 | §Open Questions **9 條未結**（Q1、Q4–Q11）。其中 Q8／Q9／Q11 的 cross-domain promotion 明確需 **3 個 converged case**（現 1／3：legal）；Q10 需真實多輪對話觀察 Optimization Suggestion 是否退化成 repeated pushback。 | 三案門檻達成 + 真實任務觀察 |
| 3 | §ADR Promotion Criteria **5 條全未達**（需至少 3 個真實法律任務含 1 個非合約任務跑完 lifecycle）。**進度 1／3**（見 §Phase 8 Dogfood #1）。 | 真實任務累積 |

適用 [`plans/README.md`](../README.md) §不搬移的例外情況第 1 條：Decision Support 的三案
promotion 是**持續生效的 watch**，未來會擴充新 Phase（promotion 或降級），沒有明確的
「完成」邊界。留在 `active/` 並在 `plans/README.md` §目前狀態 說明。

**下一次動這個 plan 的時機**：(a) 使用者回覆 4 項 ⏳；或 (b) 第 2／3 個 domain 實例化
Decision Support；或 (c) 第一個真實法律任務跑完 lifecycle 後回填 dogfood evidence。

**Glossary Impact**: yes — 新引入 framework vocabulary：`intake-dispatched workflow`（intake
答案決定下游 slice 組合的 workflow 形狀）、`legal task type`（draft／review／explain／compare／
research／due-diligence／negotiation 的第一級分派維度）、`risk tier gate`（Green／Yellow／Red
三級升級閘門）、`counterparty diligence card`（對手方三層核實產出物）。四者尚未註冊於
[`knowledge/glossary/ai-skill.md`](../../knowledge/glossary/ai-skill.md)，Phase 6 補登。

## Decision Rationale

### Problem & Why Now

使用者要「寫合約」，但實際痛點不是產生條款文字，而是四件事沒有流程：

1. **問題沒有流程化**：合約類型、我方角色、法域三者不同，要問的問題完全不同。沒有分層
   問卷就會「拿到一句需求直接開始寫條款」，草稿缺準據法、缺驗收標準、缺責任上限。
2. **背景調查（背調）沒有分層**：對手公司是否存在、是否正常營業、有無制裁／裁判／政府停權
   —— 這些決定條款怎麼寫（要不要母公司保證、付款節點怎麼切），但沒有來源分層就是臆測。
3. **法律知識會過期**：法規條文與官方契約範本都會改版（政府電子採購網的資訊服務採購契約
   範本在同一頁列出 1131226 / 1121123 / 1120711 等多版本並附修正明細對照）。模型訓練知識
   有 cutoff，**未經核對的法條引用是本 domain 最高風險的輸出**。
4. **法域決定語意**：同一個 `Termination` 條款在台灣／日本／美國／EU 的法律效果不同。
   Jurisdiction 若不是第一個問到的東西，後面所有分析都建立在錯誤前提上。

現行 repo 沒有任何法律 domain。`workflow/software-delivery/contracts.md` 的「契約」指
API / BDD behavior contract，語意完全不同，且 `route.workflow.software-delivery` 的
`user_signals` 已含「契約」——這是必須在本 plan 內裁決的字面衝突。

### Decision

新增 **`workflow/legal/`** domain，形狀是 **intake-dispatched workflow**：

- **Domain 邊界＝法律工作，不是合約文件。**Contract 只是任務之一。第一級分派維度是
  `legal task type`（draft / review / explain / compare / research / due-diligence /
  negotiation-support / lifecycle）。未來新增 NDA、勞動契約、MOU、授權書、法規查詢、GDPR、
  個資法、日本民法、公司法都在同一 domain 內擴充，不新增 workflow domain。
- **Jurisdiction 是 P0 第一題**，與 task type 並列於 intake 最前段；緊接 Governing Law 區塊
  （Contract Language → Governing Law → Dispute Resolution → Court／Arbitration）。
- **Sub-flow 目錄化**：`draft/`、`review/`、`due-diligence/`、`research/`、`negotiation/`、
  `lifecycle/`，各自 README.md 為該 slice canonical。
- **Due diligence 三層**：Counterparty Identity → Corporate Status → Risk Signals（9 類訊號）。
- **Research 正名為 Applicable Law Research**（取代原 `legal-currency` 命名）：
  Jurisdiction → Applicable Laws → Recent Updates → Government Sources → Risk Summary，
  並附各法域官方來源對照（TW / JP / US / EU / Other）。
- **Reference Sources 而非 Reference Templates**：政府採購路徑必須同時涵蓋契約範本、
  政府採購法條文、官方 Q&A 與函釋解釋，不只範本檔。
- **Review 五階段**：Structural → Clause → Risk → Missing Clause → Negotiation Suggestions，
  輸出分級為 Critical / Major / Minor / Missing / Recommended Changes。
- **Escalation 三級 risk tier gate**：Green（一般 NDA／服務／採購）→ Yellow（跨境／高金額／
  智財）→ Red（勞動／股權／投資／M&A／訴訟／政府行政處分）＝直接升級執業律師。
- **Contract Lifecycle 為獨立 sub-flow**：Need → Intake → Research → Draft → Review →
  Negotiate → Revise → Approve → Sign → Archive，預留未來簽署／存檔工具接入點。
- **Legal Strategy Analysis 為所有任務的共同 stage（第二次 review 新增）**：
  `strategy/` 回答「怎麼安排對使用者最有利」，六步管線（Requirements → Legal Analysis →
  Risk Analysis → Interest Analysis → Trade-off Analysis → Recommendation）+ 強制
  **Decision Reasoning 四欄**（Recommendation / Reason / Alternative / Trade-offs）。
  **拆兩趟**：Pass 1（Stage 3a，Research 前）列決策點與待查證前提、建議標 `provisional`；
  Pass 2（Stage 3b，Reference 後）以已核實前提收斂為 `confirmed` 並取得使用者決策。
  理由：策略推理依賴法律前提，單趟會讓建議建立在未查證假設上。九個反覆決策點的推理因素
  在 `strategy/decision-playbooks.md`。

### Alternatives Considered

- **A. `workflow/legal-contract/`（本 plan 初稿）**：reject — 以「合約」為 domain 名，未來必然
  長出 `legal-regulation` / `legal-dispute` / `legal-risk`，domain 會碎掉。Stakeholder review
  明確要求提升為 `legal`。
- **B. 只寫一份 `合約撰寫.md` 放 `workflow/documentation/`**：reject — 這不是文件撰寫流程，
  它有 intake gate、外部來源核實、時效性強制與 escalation gate，屬 workflow 層執行順序。
- **C. 掛在 `software-delivery` 底下當 slice**：reject — source-of-truth duplication risk。
  software-delivery 的 contracts 是 behavior contract，兩者 required_dependencies、
  artifact gates、風險模型無交集。
- **D. 預先鎖定為「台灣政府採購契約」單一路徑**：reject — 使用者要求「依需求調整流程、
  有點類似兩者並行」，鎖定會讓通用商務合約與跨境合約無路可走。
- **E. `workflow/legal/` + intake-dispatched + task-type 第一級分派（accept）**。
- **F. Strategy 單趟（Intake → Strategy → Research → Draft）**：reject — stakeholder 的原始
  排序，但策略需要已核實的法律前提（強制法、可執行性、範本可否增修）。單趟會產出
  「看起來有理由、前提其實沒查」的建議，正是本 domain 最高風險失效模式的變體。
  改為兩趟（accept），保留 stakeholder 要求的「Strategy 在 Draft 之前」核心語意。

### Why Not an ADR Yet

`intake-dispatched workflow` 這個形狀在本 repo 是第一次出現（既有 workflow 都是 lifecycle 固定、
slice 依 task intent lazy-load；本 domain 是 **intake 答案** 決定 slice 組合）。在只有一個 domain
實例、且尚無真實法律任務跑過之前，無法判斷這個形狀是否值得升為 cross-domain 架構決策。
Open Questions Q1／Q5／Q6 未解。

### ADR Promotion Criteria（completed 時驗證）

- [ ] foundational + cross-session + cross-project + expensive-to-reverse + explains-why 全中
- [ ] 至少 3 個真實法律任務（含 1 個非合約任務，如純法規查詢）跑完 lifecycle，dispatch matrix 未被繞過
- [ ] Open Questions 全解
- [ ] 沒有更輕的 promotion target 適用（per ADR-007）
- [ ] 系統真實使用此 contract（detector 命中率 + risk tier gate 觸發紀錄）

### Consequences（預期）

#### 正面
- 法律任務有 blocking intake gate，資訊不足時不會產出草稿。
- Jurisdiction 前置，避免用錯法域的語意分析整份合約。
- 法條／範本引用有版本欄位，過期知識可被機械偵測。
- Risk tier 三級化，讓「需要律師」從一句免責聲明變成可判定的 gate。
- 未來新增法域／法規／任務類型在 domain 內擴充，不新增 workflow domain。

#### 負面
- Domain 檔案數較多（1 entry + 4 core + 6 sub-flow + 2 yaml）。以 README-as-router 與
  slice lazy-load 吸收 token 成本。
- 官方範本版本號寫進 repo 會隨官方改版而 stale；靠 research slice 的「每次使用前核對」
  規則吸收，而非靠 repo 內容永久正確。

#### 風險
- **字面衝突**：「契約」同時命中 software-delivery。detector 無 negative signal，兩 route 同時
  命中會進 conflict（`ActiveRoute=""`，不 block），落到 workflow-routing.md Stage 2 人工裁決。
  緩解：`route.workflow.legal` 的 user_signals 選用**法律專屬**詞，並在 §常見歧義 加裁決列
  + routing 矩陣 scenario 鎖定。
- **越權風險**：agent 產出被當成法律意見。緩解：risk tier gate（Red 直接停止產出實質建議）
  + 每份產出強制 disclaimer，列為 blocking artifact gate 而非建議。
- **法域覆蓋幻覺**：agent 對不熟悉法域（例如日本民法細節）過度自信。緩解：research slice
  要求每個法域主張綁定官方來源 URL 類型；無官方來源即降級為 `unverified` 並升 Yellow。

## Runtime Execution Path

| 環節 | 內容 |
| --- | --- |
| Runtime owner | `knowledge/runtime/routing-registry.yaml` §`route.workflow.legal`（route 定義）、`runtime/cognitive-modes-discovery.yaml`（discovery signal） |
| Event / signal | 使用者訊息含法律工作意圖（合約草擬／合約審閱／保密協議／NDA／盡職調查／背景調查／準據法／管轄／違約金／法規查詢／採購契約 等 substring 命中）；或 open file 命中 `**/contracts/**`、`**/legal/**` context_signals |
| Detector | `scripts/ai-skill-cli/internal/app/detector.go :: DetectWorkflows`（既有 deterministic detector，本 plan 不改 Go 碼） |
| Route query | `ai-skill runtime workflow-context` → `RuntimeContext.ActiveRoute` |
| Loaded contract / source | `primary_source: workflow/legal/execution-flow.md` + `required_dependencies`（intake / jurisdiction / risk-classification / artifact-gates） |
| Runtime action / blocking gate | `gate.workflow.primary_source_read`（既有 PreToolUse validator `validateWorkflowPrimarySourceRead`）：單一 route 鎖定後未讀 primary_source 即擋非 Read 工具 |
| Observable evidence | `ai-skill runtime workflow-context` 對法律任務文字回 `active_route=route.workflow.legal`、`status=detected`；scenario `workflow-detector-legal-v1` + `legal-routing-matrix-v1` |

**Discovery signal（consumer #2）**：`runtime/cognitive-modes-discovery.yaml` 新增
`user_keyword_legal_work`，`context_mode: SOURCE_BACKED` + `governance_mode: STRICT`。
理由：本 domain 的失效模式是「未核對來源就斷言法條」與「資訊不足就出草稿」，
SOURCE_BACKED 強制引用來源、STRICT 強制 gate 不被跳過。此 signal 滿足
`system-upgrade-governance.yaml` §`define_runtime_trigger_flow` 對
`routing_registry_entry_without_discovery_signal_or_commit_validator` 的禁令。

**Deferred Runtime Projection**：不適用。本 plan 新增的兩個 `workflow/legal/*.yaml`
executable contract 皆宣告 `runtime_projection.enabled: true` 並在 Phase 5 project 到
`runtime.db generated_surfaces`，與 travel-planning / software-delivery 既有形狀一致。

### Per-surface consumer 表

| Generated surface key | Named consumer(s) | Consumer 類型 |
| --- | --- | --- |
| `route.workflow.legal`（routing-registry record） | (a) `detector.go :: DetectWorkflows`（auto-detect route_type）；(b) `hooks.go :: workflowPrimarySourceGate` 讀 `primary_source`；(c) `cognitive-modes-discovery.yaml :: user_keyword_legal_work` | discovery signal + Go validator |
| `workflow.legal.execution_flow.contract` | `ai-skill runtime query` routable lookup（同 travel-planning 既有形狀）；`validateRuntimeYamlProjects` commit validator 驗證 projection 存在 | routable lookup + Go validator |
| `workflow.legal.artifact_gates.contract` | 同上 | routable lookup + Go validator |
| `user_keyword_legal_work`（discovery signal） | `runtime/cognitive-modes-discovery.yaml` resolution_rules → cognitive mode 解析；`validateActivationSignals` commit validator | Go validator |

## Open Questions

- [ ] **Q1**：`intake-dispatched workflow` 是否該升為 cross-domain 形狀（travel-planning 的
      intake 其實也是分派式）？—— 需 2 個真實任務後回答。`still-open`
- [x] **Q2**：合約類型與我方角色要不要寫死？—— `resolved`（使用者 2026-07-30 回答：由詢問取得，
      依答案做不同處理 → 設計為 intake dispatch matrix）
- [x] **Q3**：背調深度？—— `resolved`（公開來源核實 + 風險旗標；stakeholder review 進一步要求
      拆為 Identity → Corporate Status → Risk Signals 三層）
- [ ] **Q4**：「契約」字面衝突是否需要 detector 層 negative signal 支援（Go 改動）？本 plan
      先用「專屬詞 + 選路表裁決 + routing 矩陣 scenario」處理，不改 Go。若真實任務出現高頻
      conflict，再開 follow-up plan。`deferred`（理由：detector 的 no-negative-signal 是刻意的
      deterministic 設計，改動需 ADR-006 對齊）
- [ ] **Q5**：官方範本／法規版本號寫進 `reference-sources.md` 後如何防 stale？目前靠 research
      slice 的「每次使用前核對」行為規則，無機械 gate。`still-open`
- [ ] **Q6**：`lifecycle/` 的 Approve → Sign → Archive 三段需要外部工具（電子簽署、文件庫）才有
      實質行為。本輪僅定義流程與交接欄位，**不**接任何外部服務。`deferred`（理由：無授權、
      無 connector；接入前需另開 plan 並取得使用者同意）
- [x] **Q7**：非中文法域（JP / US / EU）的官方來源清單由 agent 依 research slice 現場查證，
      repo 只記「該查哪一類來源」而不硬寫 URL 清單，是否足夠？
      → **`resolved`（Dogfood #2 證實足夠，但需補一條形態說明）**：JP 法域實跑 11 條條文查證
      成功，靠的是「找該法域官方資料庫的**條文級查詢介面**」這個形態知識，而非 URL 清單。
      已把該形態（逐條查證優於整法載入、枝番條文可直接取）與「搜尋摘要不可信須落到條文原文」
      寫入 [`research/README.md`](../../workflow/legal/research/README.md) §法域來源對照。
      URL 仍不硬寫——會 stale，且形態知識可跨法域遷移。
- [ ] **Q8**（stakeholder 提出，P0）：**Workflow 是否應提供 Strategy Recommendation
      （推理與最佳化建議），而不只是依使用者指定內容產生文件？**
      準據法、管轄、仲裁 vs 訴訟、付款條件、驗收方式、違約責任、IP 歸屬都存在多種合法方案。
      → **legal domain 內 `resolved`**：已落地 `workflow/legal/strategy/`（Stage 3a／3b），
      流程為「詢問 → 分析 → 建議 → 說明 trade-off → 使用者決策 → Draft/Review」。
      → **作為 cross-domain capability 仍 `still-open`**：需 3 個 domain converged case
      才升為各 workflow 正式 stage（見
      [`workflow/cross-cutting/decision-support/README.md`](../../workflow/cross-cutting/decision-support/README.md)
      §Promotion 條件，目前 1／3）。
- [ ] **Q9**（stakeholder 提出）：**Legal workflow 允許 AI 做策略推理與最佳化建議到什麼程度？**
      （例如判斷條款對甲方或乙方有利、指出「其實仲裁比東京法院適合」「分三期比一次付安全」）
      → **部分 `resolved`**：允許，但受兩個上限約束——(a) `risk_tier_gate`：Red tier 不出策略
      建議；(b) `Confidence` 欄：前提未查證只能 `provisional`。
      → **`still-open` 的部分**：目前無機械 gate 可驗證「Pass 2 確實在 Research 之後」，
      也無機制防 `playbook-as-answer`（直接套 playbook 常見結論而未綁本案因素）；
      兩者目前靠行為規則 + scenario `forbidden_outputs`。
- [ ] **Q10**（stakeholder 提出）：**Workflow 是否應主動發現對使用者有利的改善點？**
      （使用者說「我要日本法院」，是照做還是分析「改台灣法院可省成本，要考慮嗎？」）
      → **`resolved`（行為已定義）**：主動提出，但**有界**——依
      [`decision-support/README.md`](../../workflow/cross-cutting/decision-support/README.md)
      §Optimization Suggestion：講一次（含量化差異 + 使用者原方案的合理之處）→ 明確詢問 →
      使用者維持原方案就照做並記錄，**不重複勸說、不擅自改動、不製造無意義替代方案**。
      Scenario：`validation/scenarios/cross-domain/decision-support-optimization-suggestion-v1.yaml`。
      → **`still-open`**：此規則在真實多輪對話中是否會退化成 repeated pushback 或
      silent override，需實跑觀察。
- [ ] **Q11**（stakeholder 提出，架構層）：**Decision Support 應抽為全系統共同能力，而非
      legal 專屬。**travel-planning（星期三較便宜／避開塞車）、software-delivery
      （DDD vs Clean vs Modular Monolith 哪個適合）、investment（風險／報酬／稅／匯率）
      都需要同一能力。
      → **已執行**：抽出 `workflow/cross-cutting/decision-support/`（generic contract：
      六步管線、四欄格式、two-pass、Optimization Suggestion、anti-patterns、
      domain instantiation contract 四項）。legal 為 converged case #1。
      → **`still-open`**：是否升為各 workflow 正式 stage 並註冊 route，依 cross-cutting 的
      三案門檻（目前 1／3）。**刻意不現在推廣**——避免在只有一個實例時把形狀寫死
      （見 §Watch-Out List）。

## Phase 0 — Pre-Build Interrogation

### Phase 0.0 — Open Questions 核對（公版，必填）

逐條核對本 plan §Open Questions，標記處置並回寫：

- [x] 已讀本 plan §Open Questions 全部條目
- [x] 對每條標記 `resolved`（附 Phase 0 證據）/ `still-open` / `deferred`（附原因）
- [x] `resolved` 的條目已同步勾選 / 附註於 §Open Questions（Q2/Q3 已勾選）
- [x] 若盤點新發現問題，已加入 §Open Questions（Q4–Q7 為 Phase 0 與第一次 stakeholder
      review 新增；Q8–Q9 為第二次 review（strategy slice）新增）

### Phase 0.1 — Pre-Build Interrogation 六問

| 問題 | 回答 |
| --- | --- |
| `goal_scope_and_non_goals` | Goal：建立法律工作 domain（含合約與非合約任務）。Non-goals：不代替執業律師出具法律意見；Red tier（勞動／股權／投資／M&A／訴訟／行政處分）不產出實質建議只做升級；不整合付費資料庫；不接電子簽署／文件庫；本輪不起草任何具體合約。 |
| `canonical_source_owner` | `workflow/legal/execution-flow.md` 為 lifecycle canonical；`intake.md`／`jurisdiction.md`／`risk-classification.md`／`reference-sources.md`／`artifact-gates.md` 與 6 個 sub-flow README 各為其主題 canonical。execution-flow 只留 stage 順序與 dispatch matrix，不複製 slice 正文。 |
| `projection_boundary` | 兩個 `.yaml` 為 executable contract，project 到 `runtime.db generated_surfaces`；`.md` 為 human canonical。YAML 不重述 md 敘述，只放 activation events / steps / gates（同 travel-planning 既有邊界）。 |
| `source_of_truth_duplication_risk` | 風險點：`workflow/software-delivery/contracts.md`（API contract）語意撞名。緩解：domain 名 `legal`；README 明寫兩者邊界；選路表加裁決列。已 grep 確認 repo 內無其他法律內容（命中僅 software-delivery API contract 與 governance pattern draft）。 |
| `runtime_trigger_flow_or_doc_only_reason` | 非 doc-only。完整 trigger flow 見 §Runtime Execution Path，含 discovery signal + 既有 primary_source gate。 |
| `validation_targets` | (1) `workflow-detector-legal-v1`：法律任務必觸發本 route 且不被 software-delivery 吸走；(2) `legal-routing-matrix-v1`：10 種任務（Draft NDA／Review SaaS／Review Employment／Gov Procurement／Cross-border／Supplier／Distributor／Service／Software License／MOU）各自的 task type + sub-flow + risk tier 分派正確；(3) `legal-intake-gate-blocks-premature-draft-v1`：資訊不足禁止產出草稿；(4) `legal-applicable-law-verification-v1`：法條引用無版本／查核日必須標 unverified；(5) `legal-escalation-red-tier-v1`：Red tier 必須停止實質建議並升級。 |

### Phase 0.2 — Architecture Compatibility Preflight

| # | 檢查項目 | 結果 |
|---|---------|------|
| 1 | Candidate files 存在性 | `knowledge/runtime/routing-registry.yaml` ✓、`runtime/cognitive-modes-discovery.yaml` ✓、`workflow/README.md` ✓、`workflow/workflow-routing.md` ✓、`knowledge/indexes/README.md` ✓、`validation/scenario.schema.json` ✓、`runtime/runtime.db` ✓。`workflow/legal/` 為新建。 |
| 2 | Source-of-truth 一致性 | routing registry 為 route canonical（非 mirror）；runtime.db 由 `ai-skill runtime compile + refresh` 生成，不手改。 |
| 3 | Layer responsibility | 執行順序 → `workflow/`；問卷與來源分層屬同一執行流程的 slice，留在 `workflow/`；尚無足夠 lesson 沉澱，本輪**不**新增 `analysis/legal/` 或 `intelligence/legal/`（避免空殼層，見 §Watch-Out List）。 |
| 4 | Compiler / generated surface | `ai-skill runtime compile` + `refresh` 需重跑；`runtime.db` 須 `git add`（`validateRuntimeIndexFreshness` 會擋 stale checksum）。 |
| 5 | Pre-build interrogation | 已完成（Phase 0.1）。 |
| 6 | Linked updates | 見 §Linked Updates。 |
| 7 | Open Questions 核對 | 見 Phase 0.0。 |
| 8 | Execution decision | 無架構衝突、無未解 blocker question → 繼續執行。 |

## Phase 1 — Test-First Scenarios（先寫，fail by absence）

- [x] `validation/scenarios/runtime/workflow-detector-legal-v1.yaml`（含 counter_case：
      software-delivery 的「契約」不得被吸走）
- [x] `validation/scenarios/legal/routing-matrix-v1.yaml`（10 任務分派矩陣）
- [x] `validation/scenarios/legal/intake-gate-blocks-premature-draft-v1.yaml`
- [x] `validation/scenarios/legal/applicable-law-verification-v1.yaml`
- [x] `validation/scenarios/legal/escalation-red-tier-v1.yaml`
- [x] `validation/scenarios/legal/strategy-decision-reasoning-v1.yaml`（第二次 review 新增：
      四欄推理強制、not-ask-only、not-conclusion-only、two-sided interest、provisional 標記）

完成條件：六個 scenario 檔存在、通過 `validation/scenario.schema.json` 形狀。✅

## Phase 2 — Domain Core 實作

- [x] `workflow/legal/README.md`（domain 入口 + thin index + 與 software-delivery 邊界）
- [x] `workflow/legal/execution-flow.md`（11 stage lifecycle + Strategy 3a／3b + dispatch matrix）
- [x] `workflow/legal/intake.md`（Legal Task Intake：task type + jurisdiction P0 + S0/S1/S2/S3 分層）
- [x] `workflow/legal/strategy/README.md`（Strategy Recommendation Engine：六步管線 + 四欄 Decision Reasoning + 兩趟設計）
- [x] `workflow/legal/strategy/decision-playbooks.md`（九個決策點推理因素 + 反轉條件）
- [x] `workflow/legal/jurisdiction.md`（五變數模型：Language / Governing Law / Dispute Resolution / Court-Seat / Enforcement 反向決定）
- [x] `workflow/legal/risk-classification.md`（Green／Yellow／Red tier gate + Escalation Card）
- [x] `workflow/legal/reference-sources.md`（四類來源 + PCC 路徑 + 來源可信度階層 A–E）
- [x] `workflow/legal/artifact-gates.md`（產出矩陣 + 10 個 blocking gates + disclaimer + DoD）

## Phase 3 — Sub-flow 實作

- [x] `workflow/legal/draft/README.md`（六步 + Protection Pass 11 面向）
- [x] `workflow/legal/review/README.md`（5 階段 + Critical/Major/Minor + 缺漏表 + 退讓順序）
- [x] `workflow/legal/due-diligence/README.md`（三層 + 九類 Risk Signal + 10 個 flag → 條款影響）
- [x] `workflow/legal/research/README.md`（五步 + 13 類適用法 + 法域來源類型表）
- [x] `workflow/legal/negotiation/README.md`（Issue Ledger + Concession Matrix + Version Tracking + Round Brief）
- [x] `workflow/legal/lifecycle/README.md`（11 段 + 交接欄位 + Archive 關鍵日期）

## Phase 4 — Runtime 接線

- [x] `knowledge/runtime/routing-registry.yaml` 新增 `route.workflow.legal`
- [x] `runtime/cognitive-modes-discovery.yaml` 新增 `user_keyword_legal_work`（SOURCE_BACKED + STRICT，priority 19）
- [x] `workflow/legal/execution-flow.yaml` + `workflow/legal/artifact-gates.yaml`（皆 `runtime_projection.enabled: true`）

完成條件（**實跑證據 2026-07-30**）：
`ai-skill runtime workflow-context --transcript <legal task>` →
`status=detected`、`active_route=route.workflow.legal`、`conflict=false`、`substantive=true`。
Counter-case（`SDK behavior 契約 + BDD acceptance criteria`）→ `detected_routes` **不含**
`route.workflow.legal`（回 software-delivery + requirements-cognition 的既有 conflict，
由 workflow-routing.md Stage 2 裁決）。✅

## Phase 5 — Linked Updates + Projection

- [x] `workflow/README.md` §目前入口 + §Inbound References
- [x] `workflow/workflow-routing.md` §選路表 + §常見歧義（新增 §「契約」語意裁決 判準表）
- [x] `knowledge/indexes/README.md` task intent 列
- [x] `knowledge/summaries/legal-workflow.md` + `knowledge/summaries/README.md` 索引列
- [x] `knowledge/glossary/ai-skill.md` 登記 4 詞（`intake_dispatched_workflow`、`legal_task_type`、
      `risk_tier_gate`、`counterparty_diligence_card`），`owner-layer: workflow-orchestration`；
      `ai-skill glossary validate` 新增 entry 0 violations（27 條為 pre-existing，未觸碰）
- [x] `ai-skill runtime compile` PASSED + `refresh` ok（atoms=309 sources=295
      registry_records=60 summaries=26）；`runtime.db` 已更新

## Phase 6 — Decision Support 抽為 cross-domain pattern（第三次 review）

Stakeholder 指出 decision support 不只屬 legal：travel-planning、software-delivery、
investment 都需要同一能力，抽成共同 pattern 比多加 legal sub-flow 有長期價值。

- [x] `workflow/cross-cutting/decision-support/README.md`：generic contract
      （六步管線、四欄 Decision Reasoning、two-pass 規則、**Optimization Suggestion 規則**、
      domain instantiation contract 四項、9 個 anti-pattern、promotion 條件）
- [x] `workflow/cross-cutting/README.md` §Current concerns + §Slice promotion policy
      （三案門檻同樣適用，目前 1／3）
- [x] `workflow/legal/strategy/README.md` 改為 domain instantiation：generic 規則指向
      cross-cutting，本檔只留 legal 專屬（決策點清單、強制法／可行性 vs optimization 的分界、
      與其他 legal slice 分工）
- [x] `validation/scenarios/cross-domain/decision-support-optimization-suggestion-v1.yaml`
      （Q10 行為鎖定：講一次 / 不沉默照做 / 不重複勸說 / 不擅自改動 / 不 optimization theater）

**刻意不做**（避免 over-engineering）：不註冊 `route.*`、不改 travel-planning 或
software-delivery 的 stage 順序、不宣稱為全庫能力。達三案門檻前 legal 是唯一實例。

## Phase 8 — Dogfood #1（真實合約審閱，2026-07-30）

第一個真實任務：一份製造執行系統開發合約的**出稿前自審**（我方為開發／供應方）。
個案 review memo 留在業務專案，不進本庫（見 §Watch-Out List）。

| 項目 | 結果 |
| --- | --- |
| Task type 解析 | `review`（自製稿自審），未誤判為 `draft` |
| Jurisdiction | TW（由管轄條款推得；稿內**無準據法明文**，已列為缺漏） |
| Risk tier | 🟡 Yellow（智財＋金額未定＋對手方為受監理產業），非 Red，未誤升 |
| Intake gate | ✅ **有效**：總金額與交期為 `UNKNOWN`，流程**未產出替代條款文字**，只出風險清單＋7 題 blocking questions |
| Citation gate | ✅ **有效**：Applicable Law Table 5 列全標 `unverified`，全程未編造條號 |
| Diligence | Layer 1／2 完成（C 層來源，標 `probable` 並要求回溯 A 層）；Layer 3 九類逐項有結果含「未查」 |
| Strategy Pass 1 | 7 個決策點，四欄齊備、全標 `provisional`、列出待查證前提 |
| Dispatch Matrix | 實跑符合：`review` → intake + research + review + artifact-gates（＋DD，因對手方不熟悉） |

### 本輪產出的 learning（已回寫）

| # | Learning | Durable target |
| --- | --- | --- |
| 1 | 背調 Layer 2 的**營業項目會反向改變條款建議**（受監理產業 → 資料保存義務＋衍生責任排除擴充）。原 flag 表缺此項。 | [`due-diligence/README.md`](../../workflow/legal/due-diligence/README.md) 新增 `REGULATED_INDUSTRY_COUNTERPARTY` flag ＋ Layer 2 營業項目改為「問兩件事」＋ 明文定義「flag 必須對應條款調整」；lesson 見 [`feedback/history/legal/`](../../feedback/history/legal/README.md) |
| 2 | 缺 converter／office suite 時的 OPC 文字抽取備援（先探測能力再宣告不支援） | [`feedback/history/development-guidance/common/`](../../feedback/history/development-guidance/common/)（工具環境備援，非 workflow 階段） |

### ADR Promotion Criteria 進度

- 真實法律任務：**1／3**（本次為 `review` + 合約類；仍缺 1 個非合約任務如純法規查詢）
- dispatch matrix 未被繞過：✅（本次）
- Decision Support cross-domain converged case：仍 **1／3**（本次未新增 domain）

## Phase 9 — Dogfood #2（已簽署之日本法契約，2026-07-31）

第二個真實任務：一份**已簽署且期間已屆滿**的日本法委外開發契約，我方為受託方（個人）。
個案評估留在業務專案。

| 項目 | 結果 |
| --- | --- |
| 新法域 | **JP**（首次非 TW 法域）。11 條條文全部由該法域官方資料庫逐條查證 |
| Task type | `review`，但**已簽署** → 暴露 Stage 5 的預設假設缺陷（見下） |
| Risk tier | 🟡 Yellow（未進形式爭議）；已標明升 Red 的具體觸發條件 |
| Citation gate | ✅ 有效。並發現一條**若憑記憶會答錯**的改正（某條的「法院不得增減」後段已刪除） |
| 事實輸入 | 使用者中途補充履約與收款狀態 → 風險方向**反轉**（停滯可歸責於對方），評估就地修訂 |
| Red 邊界 | ✅ 守住。未代擬任何對外請求函；稅務項目僅標示不提供意見 |

### 本輪暴露的缺陷與修補

| # | 缺陷 | Durable target |
| --- | --- | --- |
| 1 | **Stage 5 預設「產出會拿去談判」**。合約已簽署時修正條文無收件人，流程跑到第五階段即失效。 | [`review/README.md`](../../workflow/legal/review/README.md) §Stage 5 依簽署狀態分流（5A 未簽／5B 已簽三段產出＋可行動事項風險分級）；[`intake.md`](../../workflow/legal/intake.md) 新增 S0-9 與已簽署追問 |
| 2 | **法域來源表只寫「該查哪一類」，未記查證形態**。 | [`research/README.md`](../../workflow/legal/research/README.md)：補「逐條查證優於整法載入」「枝番條文可直接取」「搜尋摘要不可信須落到條文原文」（實測曾遇摘要掛錯條號） |

Open Questions **Q7 已於本輪 resolved**。

### ADR Promotion Criteria 進度（更新）

- 真實法律任務：**2／3**（`review`×2；仍缺 1 個**非合約**任務，如純法規查詢或純背調）
- 跨法域驗證：✅ TW + JP，dispatch matrix 與 gate 在兩法域均未被繞過
- Decision Support cross-domain converged case：仍 **1／3**

## Phase 7 — Close Loop

- [x] `git status --short --branch` / `git diff` 去敏檢查（無本機絕對路徑、無個資、無 secrets）
- [x] rebase 到 `origin/main`（本地落後 2 週，遠端有 KGE Phase 2a migration）：
      glossary 衝突保留雙方 entry；`runtime.db` / `runtime-index.sqlite` 取 upstream 後用
      rebuild 後的 CLI 重跑 compile + refresh
- [x] `kge check` → `Ready to push (validation ok)`；rebase 後 detector 重驗仍回
      `active_route=route.workflow.legal`、`conflict=false`
- [x] commit + push（`0411c0bd` 實作 + `313d830b` runtime regeneration）+ readback + clean status
- [x] 結案評估：**不 archive**，理由與解除條件見檔頭 §為什麼還不 archive
- [x] 登記於 [`plans/README.md`](../README.md) §目前狀態

## Stakeholder 同意項目

| 項目 | 狀態 |
| --- | --- |
| Domain 名稱為 `workflow/legal/`（非 `legal-contract`） | ✅ stakeholder review 2026-07-30 要求 |
| 第一級分派維度為 legal task type（7+1 種） | ✅ stakeholder review 2026-07-30 要求 |
| Jurisdiction 為 P0 第一題 | ✅ stakeholder review 2026-07-30 要求 |
| Governing Law 區塊（language / law / dispute / court / arbitration） | ✅ stakeholder review 2026-07-30 要求 |
| Due diligence 三層（Identity → Corporate Status → Risk Signals） | ✅ stakeholder review 2026-07-30 要求 |
| Research 正名為 Applicable Law Research，含各法域官方來源 | ✅ stakeholder review 2026-07-30 要求 |
| Reference **Sources**（範本＋法條＋Q&A＋函釋）而非僅 templates | ✅ stakeholder review 2026-07-30 要求 |
| Review 五階段 + 分級輸出 | ✅ stakeholder review 2026-07-30 要求 |
| Escalation 三級 risk tier gate（Green／Yellow／Red） | ✅ stakeholder review 2026-07-30 要求 |
| Contract Lifecycle 獨立 sub-flow（Need → … → Archive） | ✅ stakeholder review 2026-07-30 要求 |
| `strategy/` slice：Legal Strategy Analysis 為所有任務共同 stage | ✅ 第二次 stakeholder review 2026-07-30 要求 |
| Decision Reasoning 四欄（Recommendation / Reason / Alternative / Trade-offs）強制 | ✅ 第二次 stakeholder review 2026-07-30 要求 |
| 九個決策點 playbook（準據法／管轄／付款／違約金／IP／保密／驗收／不可抗力／終止） | ✅ 第二次 stakeholder review 2026-07-30 要求 |
| Strategy 拆成 Pass 1 / Pass 2（Research 前後），非單趟 | ⏳ 待使用者確認（agent 提出的調整；理由見 §Alternatives F） |
| Decision Support 抽為 `workflow/cross-cutting/decision-support/` cross-domain pattern | ✅ 第三次 stakeholder review 2026-07-30 要求 |
| Optimization Suggestion 有界（講一次 → 詢問 → 尊重決定，不重複勸說／不擅自改動） | ⏳ 待使用者確認（agent 對 Q10 的界線設計） |
| Decision Support 達 3 個 domain converged case 前不註冊 route、不改其他 workflow stage 順序 | ⏳ 待使用者確認（依 cross-cutting 既有三案門檻） |
| 我方角色（起草／審閱／談判）由詢問取得後分流 | ✅ 使用者 2026-07-30 確認 |
| PCC 政府採購來源作為 reference set | ✅ 使用者提供來源 |
| `lifecycle/` 不接外部簽署／文件庫工具（Q6 deferred） | ⏳ 待使用者確認 |

## Watch-Out List citation

[`architecture/ai-native-cognitive-ecosystem-system.md`](../../architecture/ai-native-cognitive-ecosystem-system.md)
§Watch-Out List — 本 plan 相關的 wall：

- **不要為未驗證的形狀建新層**：本輪刻意**不**開 `analysis/legal/` 與 `intelligence/legal/`
  空殼目錄，等真實任務產出 lesson 再依 feedback promotion pipeline 升層。
- **不要把 domain 知識當框架能力**：合約條款與法規知識屬 domain content，不得寫進
  `enforcement/` 或 `runtime/` 成為全庫強制規則。
- **不要建空殼 sub-flow**：6 個 sub-flow README 必須各自帶可執行步驟與產出定義；只有標題與
  「待補」的檔案不算完成 Phase 3。
- **不要在只有一個實例時把 cross-domain 形狀寫死**：Decision Support 抽為 cross-cutting
  pattern，但 legal 是唯一 converged case（1／3）。Phase 6 刻意不註冊 route、不改
  travel-planning / software-delivery 的 stage 順序、不宣稱全庫能力。三案門檻是防
  premature abstraction 的 wall，不是行政程序。

## 完成條件

- [x] Phase 1–8 全部勾選
- [x] `ai-skill runtime workflow-context` 實跑證據（active_route 命中 + counter-case 不命中）
- [x] `ai-skill runtime compile` + `refresh` 無錯
- [x] 六個 validation scenario 存在且形狀合格
- [x] linked updates 全部落地（workflow README / 選路表 / knowledge index / summary / glossary）
- [x] `git status` clean、`git log origin/main..HEAD` 為空
- [ ] **archive 前額外條件**（見檔頭 §為什麼還不 archive）：4 項 stakeholder ⏳ 已確認 +
      Open Questions Q1／Q4–Q11 已結或明確 deferred + ADR Promotion Criteria 已評估

## Linked Updates

| 檔案 | 為何連動 |
| --- | --- |
| `knowledge/runtime/routing-registry.yaml` | 新 route 的 canonical 定義 |
| `runtime/cognitive-modes-discovery.yaml` | route 的 discovery signal consumer |
| `workflow/README.md` | §目前入口 是 workflow 層人類索引 |
| `workflow/workflow-routing.md` | §選路表 + 「契約」字面歧義裁決 |
| `knowledge/indexes/README.md` | task intent → primary source 的人類索引 |
| `knowledge/summaries/legal-workflow.md` | route 的 `candidate_sources` summary layer |
| `knowledge/glossary/ai-skill.md` | Glossary Impact: yes 的四個新詞 |
| `runtime/runtime.db` | routing registry + discovery signal 改動需重新 compile/refresh |
| `validation/scenarios/legal/`, `validation/scenarios/runtime/`, `validation/scenarios/cross-domain/` | test-first scenarios |
| `workflow/cross-cutting/README.md` + `decision-support/README.md` | Decision Support 的 generic contract owner；legal 的 `strategy/` 是其 instantiation |

## 與其他 plans 的關係

- [`2026-06-06-1700-workflow-activation-discovery-bridge.md`](2026-06-06-1700-workflow-activation-discovery-bridge.md)：
  本 plan 依賴其已落地的 detector + `gate.workflow.primary_source_read`，但**不修改**該 plan 的
  範圍；其 goal 已於本輪 paused。
- [`archived/2026-05-31-1900-workflow-activation-engine.md`](../archived/2026-05-31-1900-workflow-activation-engine.md)：
  detector 的 deterministic / no-negative-signal 設計來源，本 plan Q4 的 deferred 理由引用它。
