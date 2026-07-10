---
id: 2026-07-08-0825-delegation-verification-arbitration-loop
plan_kind: main
status: in-progress
owner: linyihong
created: 2026-07-08
last_updated: 2026-07-09
parent: null
baseline_ref: 2026-06-22-1009-subplan-agent-delegation
revision:
  - date: 2026-07-08
    note: "Verifier 三層驗證契約 + 測試職責分工（防 executor 自寫測試 + verifier 只重跑的自證循環）"
  - date: 2026-07-08
    note: "Dogfood 2d — 外部 monorepo outbound sync Phase 3（4 slices）；consumer overlay slice_kind/backfill；hook allowlist 契約回饋"
  - date: 2026-07-09
<<<<<<< HEAD
    note: "Dogfood 2h — ExternalRepoC common-url Execute 验证不严：RBAC 三连、V5 api-surface、combined 不得 inner-only 关闭"
  - date: 2026-07-09
    note: "第四輪 review — mutation testing 定位為 V3 evidence producer（非 V6）；Behavioral Falsification producer family 立 Q9（forming abstraction）"
=======
    note: "Dogfood 2i — ExternalRepoC user-feedback S0–S4 Execute：Stop/resume、inventory gate、2h 教训迁移、sync_jobs 分表"
>>>>>>> 96460bd (docs(dogfood): add 2i user-feedback pull Execute evidence)
---

# Delegation Verification & Arbitration Loop（委派執行→獨立驗證→仲裁閉環）

**Status**: `in-progress`（Phase 0–2 完成；**2c + 2d** 外部 monorepo dogfood 證據已寫入 kit，2026-07-08；Phase 3 / closure **不收斂**）
**Owner**: linyihong
**建立日期**: 2026-07-08
**Source**: 2026-07-08 對話 — 使用者觀察到外部框架的三角色模式：主 session 只做規劃 / 切分 / 仲裁，執行交給獨立 agent session，驗證再交給另一個獨立 session，最後由主 session 仲裁每條驗證發現（要修 / 超出範圍 / 駁回）。目標：補漏「預計與實現的落差」。主要針對 `workflow/software-delivery` 的交付處理；Ai-skill 自身任務比照辦理，觀察品質是否提升。
**Baseline**: [`03-subplan-agent-delegation`](../../archived/2026-06-22-1009-plans-system-portability-and-delivery-integration/03-subplan-agent-delegation.md)（completed，2026-07-06）— delegation `brief` schema + 雙路徑 dogfood ★★★★☆。本 plan 是其 loop 延伸（情境 C：sibling main plan + baseline_ref，不重開該 tree）。
**Glossary Impact**: yes — candidate terms：`independent_verification`（fresh-context 驗證 leg，非 executor 自驗、非 orchestrator 自 review）、`arbitration`（orchestrator 對 verifier findings 的處置協議：fix / defer / reject）、`evidence_driven_control_loop`（四責任閉環通用化候選，Q6 gated，見 §架構收斂觀察）、`behavioral_falsification`（V3 evidence producer family 候選，Q9 gated，見 §架構收斂觀察第四輪）。graduate 時才註冊到 `knowledge/glossary/ai-skill.md`；未定稿前不註冊。

> **Watch-Out List citation**：對應 [`architecture/ai-native-cognitive-ecosystem-system.md`](../../../architecture/ai-native-cognitive-ecosystem-system.md) §Watch-Out List 的「process bloat」「premature abstraction」「over-engineering」防呆：
> - **不建自動 orchestrator** — 03 的 reservation 邊界維持不變；本 plan 是**角色協議**（主 session 人工扮演 orchestrator），不是 automation。
> - **不強制所有任務走三角色 loop** — 只適用於已宣告 `delegation.enabled: true` 的 sub-plan / 委派任務，且為 advisory；小修補直接做。
> - **不先動 schema** — Phase 1 為 doc-only 協議；schema promotion 需 Phase 2 dogfood 證據（falsification ladder，一次一階）。

## Decision Rationale

### Problem & Why Now

03 已證明 delegation `brief` 能形成 capability（fresh executor 只憑 brief + `context.required` 完成任務），但 loop 只有**去程**沒有**回程**：

1. **驗證由誰做沒有契約**：目前 brief 的 `verification` 由 executor 自驗，或 orchestrator 自己 review——前者缺獨立性（executor 對自己的產出有確認偏誤），後者讓 orchestrator 被迫載入執行細節，失去仲裁位置。
1b. **自證循環（symmetric blind spot）**：若 executor 同時寫實作與測試，而 verifier **只重跑** `brief.verification` 命令，兩者共用同一套測試量尺——測試可能只覆蓋「實作怎麼寫」而非「acceptance 要求什麼」，綠燈仍可能漏掉架構違規與負面 case（2026-07-08 外部 monorepo tiered plan 執行規劃回饋）。
2. **findings 沒有處置協議**：驗證發現的問題，「要修 / 超出範圍 / 駁回」的決定散落在對話中，沒有結構化紀錄，無法回答「預計與實現的落差有沒有被系統性收斂」。
3. **角色邊界未文件化**：「主 session 只規劃 / 切分 / 仲裁、不執行」這條邊界目前不存在於任何 workflow 文件；沒有邊界就無法觀察越界（orchestrator 忍不住自己動手改 code = 信號遺失）。

Why now：03 剛 completed、dogfood 方法論（brief independence score、fresh-context 紀律）記憶猶新；review capability（`fault_finding` stance，ADR-013/014）已落地可復用；工具面（Agent + worktree + 獨立 session）已驗證可行。

### Decision

以 **doc-only 協議 + 雙 dogfood** 落地三角色 loop，不動 schema、不接 runtime：

- **Orchestrator（主 session）**：規劃、切分 sub-plan、寫 brief、仲裁 findings。**不執行、不自己驗**。
- **Executor（獨立 session / agent，可選 worktree）**：僅憑 brief + `context.required` 完成，交付 diff / artifact。（= 03 已驗證的去程，不變）
- **Verifier（另一個 fresh-context session / agent）**：輸入 = 同一份 `brief`（**`acceptance` 為主量尺**；`verification` 為 executor 自驗底線，非 verifier 唯一動作）+ executor 的 diff / artifact；輸出 = 結構化 findings。**只產證據，不做決定**（evidence ≠ decision）。驗證 leg 復用 review capability 的 `fault_finding` stance invoke，不另定 stance。

**Verifier 三層驗證契約**（2026-07-08 補強，doc-only）— 三層皆須執行；**僅第 1 層不足**：

| 層 | 動作 | 誰做 | 目的 |
|---|---|---|---|
| **L1 重跑** | 實際執行 `brief.verification` 列出的命令 | Verifier | 確認可獨立重現、非環境偶發 |
| **L2 讀碼審查** | 讀 diff，對照 `acceptance` 與架構/禁止事項（grep、靜態檢查、契約欄位） | Verifier | **不依賴** executor 自寫測試即可發現違規 |
| **L3 對抗性驗證** | 補寫或執行 **負面 / 邊界** case（見 orchestrator 在 backfill 標 `verifier_only` 的條目） | Verifier（可寫測試或補測） | 抓 happy-path-only 與「測試跟著實作走」 |

**測試職責分工**（orchestrator 在 brief / verification_backfill 設計時必填）：

| 角色 | 測試範圍 |
|---|---|
| **Orchestrator** | 每條 `acceptance` 映射證據；**明確標**哪些 case 為 `executor`（happy path）、哪些為 `verifier_only`（負面 / 架構 / 禁止事項） |
| **Executor** | 實作 + **happy path** 整合測試；**不**壟斷 `verifier_only` 對抗性 case |
| **Verifier** | L1–L3 全做；`verifier_only` 未覆蓋 → `acceptance-violation` finding |

**反模式（禁止視為獨立驗證）**：Verifier 只重跑 executor 自寫測試且不做 L2/L3 → 記入 dogfood 量測欄「verifier 降級為複讀自驗」。

**Verifier 報告最小契約**（Phase 1 定稿）：每條 finding 至少含 —

| 欄位 | 值域 |
|---|---|
| `evidence` | 具體檔案 / 行為觀察，可獨立覆核 |
| `acceptance_ref` | 對應 brief acceptance 條目，或標 `beyond-acceptance` |
| `classification` | `acceptance-violation` / `out-of-scope` / `observation` |
| `status` | `observed` / `verified` / `refuted`（沿用 observation status 紀律） |

**Arbitration 協議**（Phase 1 定稿）：orchestrator 對每條 finding 三選一，逐條記錄於被委派 sub-plan（或委派任務的 plan artifact）：

| 處置 | 語意 | 後續 |
|---|---|---|
| `fix` | 違反 acceptance，需修 | re-delegate 給 executor（brief 修訂或補充指示），修完重驗 |
| `defer` | 真實但超出本次 scope | 轉 observation / 新 plan / evidence candidate，**不**在本輪修 |
| `reject` | 經覆核不成立 | 標 `refuted` + 理由，留證據不刪 |

**Role boundary invariants（4 條，行為層）**：
1. Verifier 必須是 fresh context——不是 executor 的 session、也不是 orchestrator 自己。
2. Orchestrator 在 loop 內不產生 implementation diff；發現自己動手 = 越界信號，記入 dogfood evidence。
3. Loop 關閉條件：所有 findings 都有仲裁處置，且 `fix` 項全部重驗通過。
4. Verifier 不得將「重跑 `brief.verification`」視為充分獨立驗證；**必須**完成 L2 讀碼審查 + L3 對抗性驗證（含 `verifier_only` case）；僅 L1 = 降級，量測欄須記錄。

### Alternatives Considered

- **A. 維持現狀（executor 自驗 + orchestrator 順手 review）**：reject — 缺獨立性；orchestrator 被迫載入執行細節，仲裁與執行混位，正是使用者要補的落差來源。
- **B. 直接擴 delegation schema（加 `verification.mode: independent` / `execution.modes: [reviewer]`）**：reject（本輪）— 零 dogfood 證據就動 schema 違反 falsification ladder；03 的 `execution.modes` 已預留擴充點，證據成立後再升。
- **C. 建自動 orchestrator（自動 spawn executor + verifier + 自動收斂）**：reject — 03 明確 reservation；先證明角色協議本身有效，automation 是之後的事。
- **D. doc-only 協議 + 雙 dogfood，證據決定 schema promotion（accept）**。

### Why Not an ADR Yet

協議形狀（verifier 報告欄位、仲裁紀錄落點）會隨 Phase 0/1 調整；品質是否提升尚無證據（Q3）；可能有更輕 promotion target（plans/README.md §Delegation 擴充即足，未必需要 ADR）。雙 dogfood 完成後再依 decision-promotion-pipeline 評估。

### ADR Promotion Criteria（completed 時驗證）

- [ ] foundational + cross-session + cross-project + expensive-to-reverse + explains-why 全中
- [ ] 雙 dogfood 證實三角色 loop 可行且品質信號為正（或 null-result 明確記錄）
- [ ] Open Questions 全解
- [ ] 沒有更輕的 promotion target 適用（per ADR-007）
- [ ] 至少一個真實 software-delivery 任務走完整 loop 並留下可覆核 evidence

### Consequences（預期）

#### 正面
- 「預計與實現的落差」有結構化收斂路徑：acceptance 為量尺、findings 為證據、仲裁為決定。
- Orchestrator 保持低 context 載量（不讀執行細節，只讀 findings），符合 evidence-decision 分離。
- Verifier 差集（獨立驗證抓到而 executor 自驗沒抓到的）成為可量測的品質信號。

#### 負面
- 每個委派任務多一個 session 成本（verifier）；小任務不划算 → advisory、不強制。
- 仲裁紀錄增加 plan 表面積。

#### 風險
- Verifier 報告若不夠自足，orchestrator 被迫回讀 diff 細節 → 仲裁位置崩壞（= verifier 契約缺漏，比照 brief independence 紀律：修契約，不是怪 verifier）。
- 在 Ai-skill repo 內委派時，executor / verifier session 會撞 bootstrap gate → brief `context.required` 必須含 bootstrap 檔案；外部 repo 無此問題（gate fail-open）。
- 品質提升可能量不到（null result）→ 依 observation-only 紀律，null result 是有效信號，記錄後停在 doc-only，不硬推 schema。

## 架構收斂觀察 — 四責任分離（使用者 review 回寫，2026-07-08）

使用者 review 2b 後提出的抽象收斂，記錄為 **forming abstraction（observe-only，未 graduate）**：三角色是表象，真正的控制流是四責任閉環——

```text
Specification（brief：契約與成功條件）
  ↓
Production（executor：依契約產出；不改 acceptance / scope、不自判 Done）
  ↓
Evidence Collection（verifier：只產證據不做決定；L1 replay / L2 inspection / L3 adversarial）
  ↓
Decision / Arbitration（orchestrator：fix / defer / reject，唯一裁決者）
  ↓
下一輪 Specification（裁決回饋契約，非直接改實作）
```

- **Execution 不直接改變 Plan**：execution → evidence → decision → plan。evidence ≠ decision 原則第一次落到 delegation。
- **Verifier ≠ Reviewer**：傳統 review 把「找問題＋下判決＋開修正」三件事混在一個角色，reviewer 容易變成第二個 executor；本協議拆給三個責任，只留一個裁決者。
- **2b 已實測第四條箭頭**：F2 裁決未直接改實作，而是回饋 Specification（brief v2 追加 acceptance 9）再重跑 Production——「Decision → 下一輪 Specification」是 2b 的實際路徑，非推測。

**跨域實例（candidate analogies，observe-only）**：

| Domain | Production | Evidence | Decision | 證據狀態 |
|---|---|---|---|---|
| Coding | Executor | Verifier | Orchestrator | **已驗證（2b）** |
| Research | Research agent | Fact checker | Planner | analogy，無真實 run |
| Architecture | Designer | Architecture reviewer | Architect | analogy，無真實 run |
| Knowledge | Extractor | Evidence validator | Knowledge maintainer | analogy，無真實 run |

**紀律邊界（依 falsification ladder / governance veto test）**：真實證據目前全在 delivery 域（2b / 2a-external / 2c / **2d**）；「很像 ≠ 同 family」，其餘三域在有真實 run 前維持 analogy 紀錄。通用化定位——graduate 時以「Evidence-driven Closed Control Loop（Specification → Production → Independent Evidence → Arbitration → Specification）」取代「Delegation」——列為 Q6，gated on 至少一個非 delivery 域的真實 run；在此之前 SOP 維持 delegation 措辭，不新增通用 primitive、不改名、不建跨域框架。

**Execution Pattern ≠ Role Topology（使用者 review 第二輪，2026-07-08）**：穩定的候選是**四責任**（Spec → Produce → Evidence → Decision），不是三角色。Role topology 是 domain-variable 實例化——Research 可能是 Planner → Research Agent → Fact Checker → Planner、Knowledge 是 Curator → Extractor → Validator → Curator、Architecture 是 Architect → Designer → Architecture Review → Architect；角色名全換、四責任不變。**Q6 驗的是 pattern（四責任是否自然收斂），不是 topology（角色名是否對得上）**。

**Specification 是可演化 artifact**：2b F2 的真正新發現——契約缺漏回流到 Specification（brief v2）再重跑 Production，而非 verifier → 直接 fix code。Specification 不是一次寫死的輸入，是 loop 中會演化的 artifact。

**Adoption 三階段（使用者裁決，2026-07-08，防「stability 升級成 correctness」）**：

| 階段 | 條件 | 定位 |
|---|---|---|
| **1（現在）** | evidence 全來自 delivery 域 | 維持 **Delegation Loop**；作為 software-delivery 委派任務的 execution pattern 證據已強（**仍 advisory**，不動 SOP 強制度） |
| 2 | Research / Knowledge / Architecture 各一輪真實 run **自然收斂**到四責任閉環（非靠類比解釋） | 才可稱「Evidence-driven Control Loop」是一個 family（Q6 close） |
| 3 | cross-domain + cross-workflow + cross-project evidence 齊備 | 才考慮 execution runtime 全面預設採用 |

現在最多能說「對 software delivery 這是有效模式」，推不出「所有 workflow 都該採用」。

**第三輪 review（使用者，2026-07-08，讀 `sd-delegated-execution` 後）**——三個新命題，各立 open question：

1. **`sd-delegated-execution` 實際是 Software Delivery Execution Model，不是 Delegation SOP**。它定義的不是工具或角色，是 execution 本身：執行前 Specification → Verification Backfill → Deliverables = **Execution Contract**；V1–V4 不是 CI，是 **Evidence Production Pipeline**；整份文件描述的是 Plan → Contract → Production → Evidence → Decision → Plan。sd 域天生符合 Specification → Implementation → Verification → Acceptance，只是把 Verification 拆成 Evidence → Decision——**比一般 CI/CD 更完整，不是多一個流程，是把混在一起的責任拆開**。使用者立場：**sd 域支持全面採用**；系統級仍不預設。
2. **Verification Backfill 是候選 primitive（Evidence-first Execution）**：它回答「acceptance 如何在 execution 前就映射成證據」——從「做完再想怎麼驗」變成「Acceptance → Evidence Mapping → Execution」。可能比 delegation 本身更重要（→ Q7）。
3. **系統級不採用的判準改變**：不是角色問題，是 **Evidence Backfill 是否存在**。Research（Question → Exploration → Evidence → Hypothesis）、Knowledge（Raw → Extraction → Normalization → Validation）、Architecture（Problem → Alternatives → Tradeoff → Decision）的生命週期可能沒有「execution 前的 acceptance→evidence 映射」——這是 Q6/Q7 跨域驗證要直接觀察的點。
4. **最深層命題（→ Q8）**：真正在收斂的可能不是 Delegation、也不是 Evidence-driven Control Loop，而是 **Evidence Responsibility（證據責任）**——整份 sd 文件每節都在回答：誰產生哪種證據（backfill owner / V1–V4）、哪種證據能關哪個狀態（C1–C5）、哪種證據不能單獨 closure（inner-only 禁令）、誰依證據做決策（orchestrator 唯一）。**跨域觀察重點從「四責任閉環是否重現」細化為「證據責任分配結構是否重現」**；若重現，可升格的 primitive 是更底層的 Evidence Responsibility Model。

**第四輪 review（使用者，2026-07-09，Mutation Testing 討論後）**——四個命題，doc-only 回寫：

1. **Mutation testing = V3 的 evidence generator，不是新驗證層（不設 V6）**。架構優勢在於以 evidence 為中心而非以 testing 技術為中心（Executor → Evidence → Verifier → Evidence → Orchestrator → Arbitration）；mutation 只是「一種產生 Evidence 的方法」，架構完全不用改。已落地：[`delegated-execution.md`](../../../workflow/software-delivery/delegated-execution.md) §5 V3 evidence producer + anti-pattern「mutation score KPI」；[`test-strategy.md`](../../../workflow/software-delivery/test-strategy.md) §Mutation 加 verifier-consumer back-pointer。
2. **L3 從 imagination-driven 升級為 mechanical falsification**。現行 L3 是 verifier「猜反例」（human/AI imagination driven）；mutation 是機械枚舉行為區分點（Code → Systematically enumerate behavior changes → 哪些沒死 = verifier 要看的），與本 repo「Mechanical Enforcement > Human Discipline」的一貫立場同構。Verifier 不再需要自己想到 `price==100`。
3. **Survived mutant 只是資訊，finding 才是 evidence**。契約：Mutant → **Semantic Gap** → Verifier Finding（`evidence` = mutation + behavioral implication + 建議 `verifier_only` case；`acceptance_ref`；`classification`）。Orchestrator 不需知道 mutation engine 存在——只讀 finding。AI-native 延伸（observe-only，未落地）：Executor → Mutation → **AI Analysis**（LLM 從 survived mutant 推論 missing behavior、生成 `verifier_only`）→ Verifier 只需 Run → Observe → Evidence。
4. **不做 mutation score KPI**（82%/90%/95% 一律不做）——與 architecture 反 KPI、重 evidence quality 的立場一致；保留 targeted mutation（risk-triggered：boundary / boolean / null / authorization / invariant / guard），不跑「PIT score=92% done」。
5. **更通用的抽象候選（→ Q9）**：「Behavioral Falsification」——mutation 只是 producer 之一，未來可有 fault injection / property-based testing / model-based scenarios，全部產出同一種 evidence：「目前這個行為沒有被驗證區分」。這讓 loop 不綁定任何特定測試技術，維持「以證據為核心、可替換 producer」的設計。**紀律邊界**：forming abstraction（observe-only）；graduate 前不建 producer registry、不改 slice 措辭為通用 family、mutation 以外的 producer 無真實 run 前只是 analogy。

## Runtime Execution Path

**doc-only trial 宣告**：本 plan 不接入 runtime——不新增 `route.*`、不新增 commit-msg validator、不動 `runtime.db` generated surfaces、不動 delegation schema / `validatePlanTreeFrontmatter`。協議以文件 + 行為紀律承載；驗證 leg 復用既有 review capability invoke（`ai-skill runtime capability-invoke --capability code-review --stance fault_finding`，既有 warning-only surface，無新 wiring）。

**未來接入條件（graduation）**：Phase 3 證據評估時決策——若 (a) 三角色 loop 在 ≥2 個真實任務有效、且 (b) role boundary invariant 出現行為維持不住的證據（如 verifier fresh-context 被反覆略過），才評估 schema 欄位（`delegation.verification`）或機械檢查；由後續 plan 承載，本 plan 不 carry。**決策 deadline：2026-08-31**（與本 plan closure 同批；未達證據門檻則明確記錄「維持 doc-only」）。

## Open Questions

| ID | Question | Owner / Resolved By | Status | Closed Criteria | Resolution Evidence |
|----|----------|---------------------|--------|-----------------|---------------------|
| Q1 | Verifier 報告最小欄位集（evidence / acceptance_ref / classification / status）是否足以讓 orchestrator 仲裁而不回讀 diff 細節？ | Phase 1 定稿、Phase 2 驗證 | **resolved（2026-07-08）** | 雙 dogfood 中 loop 後 orchestrator 均未被迫回讀 diff；缺欄位已補進契約 | 2b + **2a-external（外部 repo）**：兩輪 verifier 報告自足、仲裁未回讀 diff；2a-external 有 loop 前 orchestrator 寫 code 越界（commit `<HASH-a>`），屬 role boundary 非報告欄位問題 → mechanical reminders 已補 |
| Q2 | 協議文件落點：plans/README.md §Delegation 擴充（delegation 擁有 loop）vs `workflow/cross-cutting/review/` consumer doc（review 擁有驗證 leg）？stance 復用不得重定義 | Phase 1 | **resolved（2026-07-08）** | 落點決定 + 文件落地，且未在 consumer 層重定義 stance / requires_context ✅ | 落點 = plans/README.md §Delegation loop SOP 子節（delegation 擁有 loop 生命週期；review 只被 invoke 引用）。commit `af26064` + `2d5bc60`；獨立驗證確認未重定義 stance。副作用：F1 措辭 drift 隨 canonical 移轉消解 |
| Q3 | 品質信號怎麼量：verifier 差集 findings 數 + 仲裁分佈（fix/defer/reject 比例）是否構成「品質提升」的有效指標？null result 如何記錄？ | Phase 2 | **resolved（2026-07-08，advisory 指標）** | 雙 dogfood 各留差集 + 分佈；複合指標明文化 | **複合指標**（kit §2a-external 結論表）：(1) acceptance-violation 率（2a-external **0/2 rounds**）；(2) test delta（+6）；(3) pre-merge bug fix 數（2：guard + envelope）；(4) 協調成本（spawn×4、plan commit×6）；(5) orchestrator 越界次數（1）。**結論**：品質↑有量化證據；orchestrator 寫 code↓、協調↑；verifier 邊際 catch 本任務為中等（強制 IT/結構化 defer 價值 > acceptance 差集）。null result 未出現 |
| Q4 | 仲裁紀錄落點：被委派 sub-plan 內 table（傾向）vs 獨立 artifact？ | Phase 1 | **resolved（2026-07-08）** | 落點決定並在 dogfood 實際使用 ✅ | 落點 = 被委派任務的 plan artifact 內 table（SOP 已載明）；dogfood 期記於 kit §Dogfood 紀錄（2b 仲裁表實際使用） |
| Q5 | Schema promotion 門檻：什麼證據才允許動 delegation schema（如 `delegation.verification`）？ | Phase 3 | open | 門檻明文化；未達門檻則明確記錄維持 doc-only | kit §2c + **§2d** 增強信號；consumer 機械 gate（2c/2d）後 orchestrator 零 manageCode diff — **尚不足以** close Phase 3 / schema 決策 |
| Q6 | 通用化定位：graduate 時是否以「Evidence-driven Closed Control Loop」（四責任分離：Specification → Production → Independent Evidence → Arbitration → Specification）取代「Delegation」定位？（使用者 review 2026-07-08 提出，見 §架構收斂觀察） | Phase 3（adoption stage 2 gate） | open | 至少一個**非 delivery 域**（Research / Knowledge / Architecture）真實 run **自然收斂**到四責任閉環（非類比解釋）；**驗 pattern 不驗 topology**（角色名可全換）；stage 3（runtime 全面預設）另需 cross-domain + cross-workflow + cross-project evidence，不在本 plan scope | <跨域 run evidence> |
| Q7 | **Verification Backfill 是否為獨立 primitive（Evidence-first Execution）**：「acceptance 在 execution 前映射成證據」是否比 delegation 本身更根本？（第三輪 review 命題 2） | Phase 3 / stage 2 觀察 | open | (a) sd 域內：backfill 在 ≥2 個真實委派任務穩定使用且能擋「做完再想怎麼驗」；(b) 跨域：至少一個非 delivery 域**自然出現或明確缺席**「execution 前的 acceptance→evidence 映射」——缺席也是有效答案（支持「backfill 是 sd-specific，不是 primitive」） | **2d 正向（sd 域第 2 個外部 run）**：`verification_backfill` + `deliverables[]` + `slice_kind` + V4 產出物核對；L1–L3 外層鏈為 user-visible slice 關閉條件 — kit §2d；跨域觀察仍 open |
| Q8 | **Evidence Responsibility Model 是否為更底層共同骨架**：跨域是否自然收斂出相同的「證據責任分配」結構——誰產生哪種證據 / 哪種證據能關哪個狀態 / 哪種不能單獨 closure / 誰依證據決策？（第三輪 review 命題 4；**細化 Q6 的觀察鏡頭**） | Phase 3 / stage 2 觀察 | open | 非 delivery 域真實 run 記錄其證據責任分配（不預設 sd 詞彙）；若 ≥1 域重現同構的四類責任回答 → Q8 升格候選 = Evidence Responsibility Model（而非 Delegated Execution / Control Loop）；若各域結構不同構 → 記錄差異、維持 domain-local | <跨域證據責任紀錄> |
| Q9 | **Behavioral Falsification 是否為 V3 evidence producer family**：mutation / fault injection / property-based / model-based 是否收斂為可替換 producer（皆產出「此行為未被驗證區分」型 evidence）？（第四輪 review 命題 5，見 §架構收斂觀察） | Phase 3 / 後續 delivery dogfood | open | (a) targeted mutation 作為 V3 producer 在 ≥1 個真實委派 run 實際使用，且 survived mutant → semantic-gap finding 契約成立（orchestrator 未被迫理解 mutation engine）；(b) 至少第二種 producer（fault injection / property-based / model-based）自然出現於真實 run——缺席亦為有效答案（family 維持 mutation-only，不建抽象）；graduate 前不建 producer registry / 通用 taxonomy | <V3 producer run evidence> |

## 完成條件

- [x] Phase 1 協議落地（verifier 契約 + 仲裁協議 + 3 條 role boundary invariants，落點依 Q2 = plans/README.md §Delegation loop SOP）— 2026-07-08，經 2b 委派 loop 產出
- [x] Phase 2 雙 dogfood 完成：software-delivery 外部 repo 真實任務 ×1 + Ai-skill 內部任務 ×1，各留差集 + 仲裁分佈 evidence — 2026-07-08（2a demo SD read-only + **2a-external 外部 sync-adapter Step 6** 實作 + 2b SOP 擴充；evidence → kit §Dogfood 紀錄）
- [ ] Phase 3 證據評估：Q5 schema promotion 決策（promote 或明確維持 doc-only）+ glossary 註冊決策落實
- [ ] Open Questions 全部 `resolved` / `deferred`（附原因）並回寫
- [ ] 執行 Plan Completion Closure

## Phase 0 — Pre-Build Interrogation

### Phase 0.0 — Open Questions 核對（公版，必填）

逐條核對本 plan §Open Questions，標記處置並回寫：

- [x] 已讀本 plan §Open Questions 全部條目
- [x] 對每條標記 `resolved`（附 Phase 0 證據）/ `still-open` / `deferred`（附原因）
- [x] `resolved` 的條目已同步勾選 / 附註於 §Open Questions（本輪無 resolved）
- [x] 若盤點新發現問題，已加入 §Open Questions（無新問題；tool transport 選擇記入 §Stakeholder 表，非 open question）

| Open Question | 處置 | 證據 / 原因 |
|---|---|---|
| Q1 verifier 報告自足性 | resolved | 2b + 2a-external 量測欄：loop 後仲裁未回讀 diff |
| Q3 品質信號 | resolved（advisory） | 複合指標見 kit §2a-external 結論表 + 2b 量測欄 |
| Q4 仲裁紀錄落點 | still-open（dogfood 期 interim） | dogfood 期記於 kit §Dogfood 紀錄；真實委派任務落點 Phase 1 決 |
| Q5 schema promotion 門檻 | still-open | Phase 3 |

### Phase 0.1 — 架構相容性 preflight ✅（2026-07-08）

- [x] 確認 `plans/README.md` §Delegation 與 `governance/lifecycle/plan-tree-hierarchy.md` §Delegation 現行內容：4 必填集不動，本 plan 只加 loop；無 schema 衝突
- [x] 確認 review capability 現行 invoke 契約：`capability-registry.yaml` 有 `code-review` + `requires_context.stance: [fault_finding]`（status: active）；`stance_enum.reserved_policy` 禁止新增 stance 值 → verifier leg 復用 `fault_finding`，consumer 不重定義（符合 review/README.md governance invariant）
- [x] 確認 03 baseline dogfood 方法論可平移：brief independence score 的「讀檔差集」判準平移為 executor / verifier 的「讀檔紀錄」欄（kit 模板 A/B）
- [x] 確認 `workflow/software-delivery/README.md` §Review invoke 邊界：Review 不是 workflow phase / lifecycle slice，是 capability invoke → verifier leg 定位為 invoke 消費者，成立
- [x] （新增）確認 Layer 3 tool adapter 存在：`ai-tools/agent/cursor.md`（Cursor transport 細節歸該層 + kit §Cursor 傳輸備註，模板本體 tool-neutral）

## Phase 1 — 協議定稿（doc-only）✅（2026-07-08，經 2b 委派落地）

- [x] 決定 Q2 落點 = `plans/README.md` §Delegation「派發 → 獨立驗證 → 仲裁（loop SOP）」子節（canonical）；verifier 4 欄位契約 + 仲裁三處置表 + 3 條 invariants 已落地（commit `af26064` + fix `2d5bc60`，**由 2b 委派 loop 產出，orchestrator 未寫實作**）
- [x] 決定 Q4 仲裁紀錄落點 = 被委派任務的 plan artifact（sub-plan table）；dogfood 期記於 kit §Dogfood 紀錄（2b 已實際使用）
- [x] 更新 `plans/README.md` §Delegation loop SOP（同上，經獨立驗證 acceptance 8/8 + fix 重驗 pass）
- [x] Ai-skill repo 內委派的 bootstrap 注意事項寫入 SOP（tool-neutral gate id `gate.bootstrap.receipt_present`；executor / verifier 實測 Bootstrap Receipt 通過）
- [x] **Verifier 三層驗證契約** + 測試職責分工（`executor` / `verifier_only`）補強 — 2026-07-08，外部 monorepo tiered plan 執行規劃回饋；見 §Decision Rationale、plans/README.md、kit 模板 B
- [x] **software-delivery 泛化 slice**（`sd-delegated-execution`，maturity: candidate）— 2026-07-08，從外部 consumer overlay 泛化：角色×證據責任矩陣、verification backfill（tier+owner）、`deliverables[]`、slice 關閉狀態（`implementation_done` ≠ `slice_compliant_closed`，C1–C5）、Verifier V1–V4（V4 產出物/流程核對）、orchestrator 錨定紀律、anti-patterns。落點 [`workflow/software-delivery/delegated-execution.md`](../../../workflow/software-delivery/delegated-execution.md)；tier 語意引用既有 test-strategy / validation slices 不另建 taxonomy；`deliverable-omission` / `process-omission` 為 delivery 域 candidate 分類，第二 consumer 證據前不進 canonical enum。**第二批泛化（orchestrator 脊椎，同日）**：觸發條件（含 loop-未關閉持續 + transport-adaptation 例外須 plan 註明）、執行順序 6 步不可跳過、主 session 禁止 4 條 + 被擋時的正確反應、機械 gate 事件生命週期模式（Layer 3 optional，arm → deny → executor grant → clear + bypass 書面記錄）——泛化自 consumer 的 orchestrator 機械提醒檔
- [x] **V3 evidence producer 定位**（第四輪 review 回寫，2026-07-09）：targeted mutation 作為 V3 的 evidence producer 寫入 [`delegated-execution.md`](../../../workflow/software-delivery/delegated-execution.md) §5（含 survived mutant → semantic-gap finding 契約、anti-pattern「mutation score KPI」）；[`test-strategy.md`](../../../workflow/software-delivery/test-strategy.md) §Mutation 加 back-pointer。doc-only、不設 V6、不建 producer registry；family 通用化 gated on Q9

## Phase 2 — 雙 dogfood

**焦點紀律（承 03）**：驗的是**契約自足性**（verifier 報告是否足以仲裁、brief 是否足以執行），不是執行者 / 驗證者聰不聰明。任何失效先反問契約缺漏。

**Transport（tool-portability 驗證，2026-07-08 調整）**：雙 transport 分工——**2a 用 Cursor**（human 路徑：使用者把模板貼進獨立 fresh chat）、**2b 用 Claude Code Agent 工具**（agent 路徑：orchestrator session 直接 spawn executor / verifier 獨立 agent，使用者授權全程自駕 + 每階段 commit/push）。同一套模板跑兩種 transport = 「模板 tool-neutral、工具細節只在傳輸備註」分層成立的更強證據。模板見 [`01-dogfood-prompt-kit.md`](01-dogfood-prompt-kit.md)。

- [x] Dogfood prompt kit 建立（模板 A executor / B verifier / C 仲裁+量測欄，Cursor-first；kit 為 rendered transport artifact，canonical 在本檔 §Decision Rationale）— 2026-07-08
- [x] **2a — software-delivery 外部 repo 真實任務**：挑一個真實、小、可驗收的交付任務走完整 loop（orchestrator 寫 brief → executor 執行 → verifier 報告 → 仲裁）。記錄：verifier 差集 findings、仲裁分佈、orchestrator 是否越界執行、是否被迫回讀 diff。 — 2026-07-08：**2a-external** 外部 monorepo sync-adapter Step 6（kit §2a-external）；另 2a demo SD read-only（kit §2a）
- [x] **2b — Ai-skill 內部任務** ✅（2026-07-08，agent transport）：委派任務 = Phase 1 SOP 擴充本身。完整 loop 走完：brief → executor（worktree，自驗 8/8，讀檔零差集）→ verifier（fresh，findings ×3 全 observation 級，0 violation）→ 仲裁（fix 1 / defer 1 / reject 1）→ fix 回派（brief v2）→ delta 重驗 pass → merge。**Q1 正向證據：orchestrator 全程未回讀 diff，僅憑 verifier 報告仲裁**。完整 evidence 見 [`01-dogfood-prompt-kit.md`](01-dogfood-prompt-kit.md) §Dogfood 紀錄 2b。
- [x] **2a-external — 外部 monorepo sync-adapter Step 6** ✅（2026-07-08，Cursor Task transport）— kit §2a-external
- [x] **2c — 外部 monorepo tiered archive 全線（8 slices，Phase A–D）** — 2026-07-08 **證據 only**（kit §2c）；強化 Q3 品質信號，**不**視為 Phase 3 closure
- [x] **2d — 外部 monorepo outbound sync Phase 3（4 slices）** — 2026-07-08 **證據 only**（kit §2d）；強化 backfill / consumer gate 信號；**不**視為 Phase 3 closure
- [x] **2h — ExternalRepoC common-url Execute 验证缺口** — 2026-07-09 **證據 only**（[`evidence/2h-externalrepoc-common-url-verification-gaps.md`](evidence/2h-externalrepoc-common-url-verification-gaps.md)）；RBAC 三连、V5 全 API 面、combined 不得 defer L1–L3；**不**視為 Phase 3 closure
- [ ] **2e — 跨域 run（Research/Audit 域）：grandfather sunset audit** — 2026-07-08 啟動（kit §2e）；Q6/Q7(b)/Q8 的 stage-2 裁決 run；任務 = `pre_2026_05_28_doc_only_completion` sunset（2026-08-31）的 4-plan wiring 調查與處置建議（真實 deadline 待辦，非 manufactured）；brief 以 Research 域原生語彙撰寫，觀察四責任 / backfill / 證據責任結構是否**自然**出現
- [x] 回饋迴路（2b 觸發 ×1）：F2 暴露 brief v1 缺「reusable doc 目標須含 tool-neutral 措辭條款」→ brief v2 追加 acceptance 9、kit 使用說明補教訓；修契約未修執行者；fix leg 重跑通過。2a 若再暴露缺漏比照處理。

## Phase 3 — 證據評估與收斂

- [x] 彙整 dogfood evidence，回答 Q3（品質信號成立 / null result）— 2026-07-08：成立（advisory 複合指標）；null result 未出現；**2c 補強**（8-slice、acceptance-violation 2/8、slice 紀律後 orchestrator 零實作 diff）；**2d 補強**（4-slice、outer L1–L3 關閉紀律、gate 後零 manageCode diff）— 見 kit §2c / §2d
- [ ] Q5 決策：schema promotion（另立 plan）或明確維持 doc-only（記錄門檻與未達原因）
- [ ] Q6 決策：通用化定位（Evidence-driven Closed Control Loop vs Delegation）——依非 coding / delivery 域真實 run 證據裁決；無證據則維持 delegation 定位、四責任觀察留檔（§架構收斂觀察）
- [ ] Q7 決策：Verification Backfill 是否升格為獨立 primitive（Evidence-first Execution）——依 sd 域 ≥2 次真實使用 + 跨域出現/缺席觀察裁決；缺席亦為有效關閉（sd-specific）
- [ ] Q8 決策：Evidence Responsibility Model 是否為更底層骨架——跨域 run 以「證據責任分配結構是否同構」為觀察鏡頭（取代單看角色/loop 形狀）；同構 → 升格候選改為 Evidence Responsibility Model；不同構 → 記錄差異維持 domain-local
- [ ] Q9 決策：Behavioral Falsification producer family——依 (a) mutation 作為 V3 producer 的真實 run 使用證據 + (b) 第二種 producer 出現/缺席裁決；缺席 → family 維持 mutation-only、不建抽象（有效關閉）
- [ ] sd 域定位落地評估：依使用者第三輪 review「sd 支持全面採用」，評估將 `sd-delegated-execution` 從「advisory + delegation 宣告任務」重定位為 **Software Delivery Execution Model**（含 slice 更名/正文重框、execution-flow 導航同步、advisory→default 的升級條件明文化）——獨立 linked-update 批次，不與 Q7/Q8 混批
- [ ] Glossary 註冊決策落實（`independent_verification` / `arbitration` / `evidence_driven_control_loop` 註冊或明確不註冊）
- [ ] 執行 Plan Completion Closure（含 plans/README.md 狀態表更新、搬移 archived）

## Stakeholder 同意項目

> 描述現行選定策略（治理現況），改方向時直接更新本表。

| 決策面 | Current selected strategy |
|--------|---------------------------|
| Loop 形狀 | 三角色：orchestrator（規劃/切分/仲裁，不執行）/ executor（brief-only，happy path 測試）/ verifier（fresh-context，L1–L3 驗證，可補 `verifier_only` 測試） |
| 落地方式 | doc-only 協議 + 雙 dogfood；不動 schema、不接 runtime、不建自動 orchestrator |
| 通用化 adoption | **三階段**（§架構收斂觀察）：stage 1 現況 — **sd 域使用者支持全面採用**（第三輪 review 2026-07-08：定位為 Software Delivery Execution Model；落地重框列 Phase 3 checkbox）、系統級不預設；stage 2 gated on 跨域自然收斂（Q6/Q7/Q8）；stage 3（runtime 全面預設）需 cross-domain + cross-workflow + cross-project evidence，超出本 plan scope |
| Stage 2 觀察鏡頭 | **證據責任分配結構**（Q8）：跨域 run 記錄「誰產生哪種證據 / 哪種證據關哪個狀態 / 哪種不能單獨 closure / 誰依證據決策」，不預設 sd 詞彙、不驗角色名 |
| 驗證 leg | 復用 review capability `fault_finding` stance invoke，不另定 stance |
| V3 evidence producer | targeted mutation（risk-triggered：boundary / boolean / null / authorization / invariant / guard）；survived mutant 須轉 semantic-gap finding，**不做 mutation score KPI**；producer family 通用化（Behavioral Falsification）gated on Q9，graduate 前 mutation-only（2026-07-09 第四輪 review 裁決） |
| 適用範圍 | advisory；只適用已宣告 delegation 的委派任務；主打 software-delivery，Ai-skill 比照 |
| Schema promotion | gated on Phase 2 證據（Q5），deadline 2026-08-31 |
| Dogfood transport | 雙 transport：2a Cursor（human 路徑，使用者操作）/ 2b Claude Code Agent（agent 路徑，orchestrator 自駕，2026-07-08 使用者授權）；模板 tool-neutral，工具細節只在 kit 傳輸備註 + `ai-tools/agent/`（Layer 3） |
| 2b 委派任務 | Phase 1 的 plans/README.md §Delegation loop SOP 擴充本身（真實待辦、可驗收、orchestrator 全程不碰實作） |

## 與其他 plans 的關係

- [`2026-06-22-1009-plans-system-portability-and-delivery-integration/03-subplan-agent-delegation.md`](../../archived/2026-06-22-1009-plans-system-portability-and-delivery-integration/03-subplan-agent-delegation.md) — **baseline**（completed）：delegation brief schema + 雙路徑 dogfood。本 plan 延伸其 loop（去程 → 回程），不重開該 sub-plan；自動 orchestrator reservation 邊界維持。
- [`2026-06-22-1009-plans-system-portability-and-delivery-integration/02-software-delivery-plan-first-ordering.md`](../../archived/2026-06-22-1009-plans-system-portability-and-delivery-integration/02-software-delivery-plan-first-ordering.md) — plan-first ordering 是本 loop 的前置（orchestrator 產 plan artifact 先於執行）；本 plan 不改其 Q4 關閉條件。
- [`archived/2026-07-06-review-architecture-adr/_plan.md`](../../archived/2026-07-06-review-architecture-adr/_plan.md) — review = cross-cutting capability invoke（ADR-013 D2）；verifier leg 是其消費者。
- [`active/2026-06-16-1131-evidence-candidate-system.md`](../2026-06-16-1131-evidence-candidate-system.md) — `defer` 處置的 findings 可轉 evidence candidate（人工 capture，不新增 scanner 職責）。
