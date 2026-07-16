---
id: 2026-07-08-0825-delegation-verification-arbitration-loop
plan_kind: main
status: in-progress
owner: linyihong
created: 2026-07-08
last_updated: 2026-07-14
parent: null
baseline_ref: 2026-06-22-1009-subplan-agent-delegation
revision:
  - date: 2026-07-08
    note: "Verifier 三層驗證契約 + 測試職責分工（防 executor 自寫測試 + verifier 只重跑的自證循環）"
  - date: 2026-07-08
    note: "Dogfood 2d — 外部 monorepo outbound sync Phase 3（4 slices）；consumer overlay slice_kind/backfill；hook allowlist 契約回饋"
  - date: 2026-07-09
    note: "Dogfood 2d′ — ExternalRepoC 9j2 module alignment follow-on；integration gate、remote_absent_delete、live teardown、release-time gate；證據全文遷至 evidence/"
  - date: 2026-07-09
    note: "Dogfood 2h — ExternalRepoC common-url Execute 验证不严：RBAC 三连、V5 api-surface、combined 不得 inner-only 关闭"
  - date: 2026-07-09
    note: "Mutation review — mutation testing 定位為 V3 evidence producer（非 V6）；Behavioral Falsification producer family 立 Q9（forming abstraction）"
  - date: 2026-07-09
    note: "Dogfood 2i — ExternalRepoC user-feedback S0–S4 Execute：Stop/resume、inventory gate、2h 教训迁移、sync_jobs 分表"
  - date: 2026-07-10
    note: "Dogfood 2j — ExternalRepoC push Execute **负向证据**：单 Task 跳过 Verifier、delegation.enabled:false 误当豁免；consumer verifier-after-executor gate 回饋"
  - date: 2026-07-10
    note: "Dogfood 2k — ExternalRepoC push **纠偏后 post-close**：用户手验 runtime/UI 缺口、Worker 拓扑、post-close surgical debt；V5-W/U 契约候选"
  - date: 2026-07-10
    note: "Dogfood 2l — ExternalRepoC common-url S2′ mirror **负向证据**：0 Executor/Verifier、surgical bypass 滥用、Shell gate 洞；retroactive R1 Verifier 契约"
  - date: 2026-07-10
    note: "還原 #1：併發 plan 回寫（b6481e5 / 0958a38）自陳舊底稿覆蓋第五~十輪 + ERA + 2e/2f/2g + Q9；自 e2d5091 / bfb2704 / 66f58ed 重建"
  - date: 2026-07-10
    note: "Dogfood 2m — ExternalRepoC Phase G-mirror **批量 retrofit**：V-m1–V-m5 模板 + 登记总表；02/01 合规 loop 对照 2l；stale JVM V5-A 复发；phase vs slice close_kind"
  - date: 2026-07-10
    note: "Dogfood 2n — ExternalRepoC 07 push DEL-S1–S6 **正向闭环**：6/6 E+V、sub-plan completed；2e 勾选完成；ADR SD 完整 loop 证据；Phase 3 Q7 sd 域信号增强（仍 open Q5）"
  - date: 2026-07-10
    note: "Dogfood 2o — <PROJECT_ROOT> tab-scroll **单 session vs 三角色** 对照；partial authority / deploy smoke≠L3；Q8 ERA 信号"
  - date: 2026-07-10
    note: "還原 #2：2de1686 三度自陳舊底稿覆蓋第五~十二輪；自 dcd6f9e 重建並疊回 2k–2o 新增。**Collision N=3 → 依 failure-to-validator-closure，機械 validator 升為 due**（獨立 task）"
  - date: 2026-07-13
    note: "Dogfood 2p — ExternalRepoC 09 Integration 默认切流 INT-D0–D5：6/6 E+V、same-branch、一口气未跳 V；live defer"
  - date: 2026-07-13
    note: "2p 契约回写：多 todo=多轮 E→V；brief 累积表；Verifier 四栏强制 → plans/README + kit + delegated-execution + consumer overlay"
  - date: 2026-07-13
    note: "Dogfood 2q — transport cutover Verifier inner-only 假绿：loop≠路径通；features+L3+V5-A；consumer gate.plan_transport_runtime_evidence"
  - date: 2026-07-13
    note: "Dogfood 2r — <PROJECT_ROOT> player overlay Mode A：soft-nav 绿→cold URL 死；entry-path / elementFromPoint 回馈"
  - date: 2026-07-13
    note: "Sync 跨域表：Research 改標 2e 已驗證；Architecture / Knowledge 仍 analogy；紀律邊界與 adoption stage 2 進度同步"
  - date: 2026-07-14
    note: "Dogfood 2s — Architecture 域：UI Pattern Knowledge plan review（R1+R2 Verifier）；跨域表 Architecture→已驗證；stage 2 = 2/3（Knowledge 仍缺）；evidence/2s-…"
  - date: 2026-07-14
    note: "Domain Boundary — APK Analysis ↔ Software Delivery Capability Handoff（不對稱成熟度）；companion 04；預註冊 dogfood 2t（2t-A Discovery / 2t-B Handoff→SD Intake）；不套三角色於 apk-analysis"
  - date: 2026-07-16
    note: "Dogfood 2v — greenfield consumer Phase 2 preflight（brief+backfill；0 E+V；配對 domain-model execute）"
---

# Delegation Verification & Arbitration Loop（委派執行→獨立驗證→仲裁閉環）

**Status**: `in-progress`（Phase 0–2 完成；外部 monorepo dogfood **2a–2q** + consumer **2o/2r**；跨域 **2e Research** + **2s Architecture**；**2n/2p 正向闭环**；**2t 預註冊** APK↔SD Capability Handoff（見 [`04-apk-capability-handoff-boundary.md`](04-apk-capability-handoff-boundary.md)，未跑）；Phase 3 / closure **仍不收斂** — Q5 schema promotion open，deadline 2026-08-31；stage 2 **2/3** Knowledge 仍缺）
**Owner**: linyihong
**建立日期**: 2026-07-08
**Source**: 2026-07-08 對話 — 使用者觀察到外部框架的三角色模式：主 session 只做規劃 / 切分 / 仲裁，執行交給獨立 agent session，驗證再交給另一個獨立 session，最後由主 session 仲裁每條驗證發現（要修 / 超出範圍 / 駁回）。目標：補漏「預計與實現的落差」。主要針對 `workflow/software-delivery` 的交付處理；Ai-skill 自身任務比照辦理，觀察品質是否提升。**2026-07-14 延伸**：真實 APK 分析專案進入視野 → 落實 Domain vs Workflow 邊界（APK Discovery candidate ≠ SD Delegated Execution validated）；Capability Handoff 為 Domain Boundary，非把三角色套到 apk-analysis。
**Baseline**: [`03-subplan-agent-delegation`](../../archived/2026-06-22-1009-plans-system-portability-and-delivery-integration/03-subplan-agent-delegation.md)（completed，2026-07-06）— delegation `brief` schema + 雙路徑 dogfood ★★★★☆。本 plan 是其 loop 延伸（情境 C：sibling main plan + baseline_ref，不重開該 tree）。
**Glossary Impact**: yes — candidate terms：`independent_verification`（fresh-context 驗證 leg，非 executor 自驗、非 orchestrator 自 review）、`arbitration`（orchestrator 對 verifier findings 的處置協議：fix / defer / reject）、`evidence_driven_control_loop`（四責任閉環通用化候選，Q6 gated）、`evidence_responsibility_architecture`（ERA，Q8 假說，第四輪 review 命名）、`evidence_first_acceptance`（Q7 結論的 universal 候選，parent of backfill / embedded evidence rules）、`behavioral_falsification`（V3 evidence producer family 候選，Q9 gated）、`capability_handoff` / `deliverable_capability` / `discovery_evidence`（2026-07-14 Domain Boundary；見 [`04`](04-apk-capability-handoff-boundary.md)）。見 §架構收斂觀察。graduate 時才註冊到 `knowledge/glossary/ai-skill.md`；未定稿前不註冊。

> **Watch-Out List citation**：對應 [`architecture/ai-native-cognitive-ecosystem-system.md`](../../../architecture/ai-native-cognitive-ecosystem-system.md) §Watch-Out List 的「process bloat」「premature abstraction」「over-engineering」防呆：
> - **不建自動 orchestrator** — 03 的 reservation 邊界維持不變；本 plan 是**角色協議**（主 session 人工扮演 orchestrator），不是 automation。
> - **不強制所有任務走三角色 loop** — advisory 僅限非 Execute 情境（純問答 / 只讀 / surgical 小修直接做）。**2026-07-10 更新（2j F2 裁決）**：使用者 Execute 意圖 = mandatory loop，`delegation.enabled: false` 不是豁免（見 §Stakeholder 適用範圍）。
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
- [x] 至少一個真實 software-delivery 任務走完整 loop 並留下可覆核 evidence — **2n**（ExternalRepoC 07 push DEL-S1–S6，6 slice E+V，[`evidence/2n-externalrepoc-push-delivery-s1-s6-compliant-loop.md`](evidence/2n-externalrepoc-push-delivery-s1-s6-compliant-loop.md)）

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

**跨域實例（candidate analogies → 真實 run 進度）**：

| Domain | Production | Evidence | Decision | 證據狀態 |
|---|---|---|---|---|
| Coding | Executor | Verifier | Orchestrator | **已驗證（2b）** |
| Research | Research agent | Fact checker | Planner | **已驗證（2e）** — grandfather sunset audit |
| Architecture | Designer | Architecture reviewer | Architect | **已驗證（2s）** — UI Pattern Knowledge plan review |
| Knowledge | Extractor | Evidence validator | Knowledge maintainer | analogy，無真實 run |
| **APK Analysis** | Discovery / RE producer |（尚無獨立 Verifier 契約）| Capability Assessment（*Can we explain?*） | **Domain Knowledge 成熟**；Discovery workflow = **candidate**；Delegated Execution = **未宣稱**（見 [`04`](04-apk-capability-handoff-boundary.md)；dogfood **2t** 預註冊） |

**紀律邊界（依 falsification ladder / governance veto test）**：真實證據目前在 delivery 域（2b / 2a-external / 2c / **2d**…）+ **Research（2e）** + **Architecture（2s）**；Knowledge 在有真實 run 前維持 analogy。「很像 ≠ 同 family」。**APK Analysis 不得因為與 SD 共用 ERA Decision Semantics，就被標成 Delegated Execution 已驗證**——成熟度不對稱是刻意的（2026-07-14 stakeholder）。Q6 最低門檻（≥1 非 delivery 域）已由 2e 滿足；adoption stage 2「三域各一輪」= Research ✅ / Architecture ✅ / **Knowledge 仍缺（2/3）**；**2t 不填 Knowledge、不偷換 stage 2**。通用化定位——graduate 時以「Evidence-driven Closed Control Loop（Specification → Production → Independent Evidence → Arbitration → Specification）」取代「Delegation」——列為 Q6，改名裁決留 Phase 3；在此之前 SOP 維持 delegation 措辭，不新增通用 primitive、不改名、不建跨域框架。

**第十三輪 review（使用者，2026-07-14）——Domain vs Workflow 邊界落實（APK ↔ SD）**：

1. **Handoff 觸發 ≠ 分析完成**，= **形成 deliverable capability**（SDK / Client / Contract / OpenAPI / BDD / tests / library）。API catalog、protocol、crypto/UI/event model 仍是 Discovery Evidence。
2. **Capability Handoff = Domain Boundary 契約**，不是 apk-analysis workflow 多一步；SD Intake 只吃 Capability Proposal，不吃 Frida/mitm。
3. **不對稱寫法強制**：apk-analysis = Current Workflow（candidate）；software-delivery = Delegated Execution（validated）。禁止直接複製三角色到 APK。
4. **既有** [`feature-handoff`](../../../workflow/apk-analysis/artifact-gates/feature-handoff.md) **保留為 artifact gate**；Capability Proposal 是其上界（「能重建」≠「能交付」）。全文與 dogfood 預註冊見 [`04-apk-capability-handoff-boundary.md`](04-apk-capability-handoff-boundary.md)。

**Execution Pattern ≠ Role Topology（使用者 review 第二輪，2026-07-08）**：穩定的候選是**四責任**（Spec → Produce → Evidence → Decision），不是三角色。Role topology 是 domain-variable 實例化——Research 可能是 Planner → Research Agent → Fact Checker → Planner、Knowledge 是 Curator → Extractor → Validator → Curator、Architecture 是 Architect → Designer → Architecture Review → Architect；角色名全換、四責任不變。**Q6 驗的是 pattern（四責任是否自然收斂），不是 topology（角色名是否對得上）**。

**Specification 是可演化 artifact**：2b F2 的真正新發現——契約缺漏回流到 Specification（brief v2）再重跑 Production，而非 verifier → 直接 fix code。Specification 不是一次寫死的輸入，是 loop 中會演化的 artifact。

**Adoption 三階段（使用者裁決，2026-07-08，防「stability 升級成 correctness」）**：

| 階段 | 條件 | 定位 |
|---|---|---|
| **1（已過）** | evidence 來自 delivery 域 | 維持 **Delegation Loop**；作為 software-delivery 委派任務的 execution pattern 證據已強（**仍 advisory**，不動 SOP 強制度） |
| **2（進行中，2/3）** | Research / Knowledge / Architecture 各一輪真實 run **自然收斂**到四責任閉環（非靠類比解釋） | Research ✅（2e）；Architecture ✅（2s）；**Knowledge 仍缺**。Q6 最低門檻（≥1 非 delivery）已滿足；family 宣稱仍等 3/3 |
| 3 | cross-domain + cross-workflow + cross-project evidence 齊備 | 才考慮 execution runtime 全面預設採用 |

現在最多能說「對 software delivery + Research/Audit + Architecture（plan review）這是有效模式」，推不出「所有 workflow 都該採用」；Knowledge 尚未有真實 run。

**第三輪 review（使用者，2026-07-08，讀 `sd-delegated-execution` 後）**——三個新命題，各立 open question：

1. **`sd-delegated-execution` 實際是 Software Delivery Execution Model，不是 Delegation SOP**。它定義的不是工具或角色，是 execution 本身：執行前 Specification → Verification Backfill → Deliverables = **Execution Contract**；V1–V4 不是 CI，是 **Evidence Production Pipeline**；整份文件描述的是 Plan → Contract → Production → Evidence → Decision → Plan。sd 域天生符合 Specification → Implementation → Verification → Acceptance，只是把 Verification 拆成 Evidence → Decision——**比一般 CI/CD 更完整，不是多一個流程，是把混在一起的責任拆開**。使用者立場：**sd 域支持全面採用**；系統級仍不預設。
2. **Verification Backfill 是候選 primitive（Evidence-first Execution）**：它回答「acceptance 如何在 execution 前就映射成證據」——從「做完再想怎麼驗」變成「Acceptance → Evidence Mapping → Execution」。可能比 delegation 本身更重要（→ Q7）。
3. **系統級不採用的判準改變**：不是角色問題，是 **Evidence Backfill 是否存在**。Research（Question → Exploration → Evidence → Hypothesis）、Knowledge（Raw → Extraction → Normalization → Validation）、Architecture（Problem → Alternatives → Tradeoff → Decision）的生命週期可能沒有「execution 前的 acceptance→evidence 映射」——這是 Q6/Q7 跨域驗證要直接觀察的點。
4. **最深層命題（→ Q8）**：真正在收斂的可能不是 Delegation、也不是 Evidence-driven Control Loop，而是 **Evidence Responsibility（證據責任）**——整份 sd 文件每節都在回答：誰產生哪種證據（backfill owner / V1–V4）、哪種證據能關哪個狀態（C1–C5）、哪種證據不能單獨 closure（inner-only 禁令）、誰依證據做決策（orchestrator 唯一）。**跨域觀察重點從「四責任閉環是否重現」細化為「證據責任分配結構是否重現」**；若重現，可升格的 primitive 是更底層的 Evidence Responsibility Model。

**第四輪 review（使用者，2026-07-09，讀 2e 裁決結果後）**——假說升級與方法論轉向：

1. **「Topology 變、責任不變」= 抽象層抓對的證據**：2e 的價值不在「Research 也能用」，而在角色全換（Executor→Investigator、Verifier→Fact-checker、決策層多出 Maintainer 一級）但 Specification → Production → Independent Evidence → Decision 責任結構不變。**若兩域長得一模一樣，反而可能只是 copy SOP；topology 不同而責任一致，才代表抽象抓對了。**
2. **Evidence-first Acceptance 抽象樹**（Q7 結果的正確畫法——抽象往上一層、具體往下一層）：
   ```text
   Evidence-first Acceptance（invariant：acceptance 在開始前就必須定義證據）
   ├── Delivery  → Backfill Table（tier + owner，結構化強形式）
   ├── Research  → Embedded Evidence Rules（acceptance 內嵌證據標準，弱形式）
   └── 其他 domain → <該域自然長出的形式，待觀察>
   ```
3. **Invariant candidate 正式命名（Q8 核心）**：**Evidence Producer must not be Closure Authority（自產證據不能自我關閉）**——sd（inner test 不能關 user-visible slice）與 research（調查者自查 necessary but insufficient）**自然形成、非 SOP 規定**，兩域同構。
4. **Primitive 候選排序（使用者裁定，2026-07-09）**：① Evidence Responsibility（誰可以生產什麼證據）② Evidence-first Acceptance（acceptance 如何先定義證據）③ Independent Verification（誰不能驗證自己）④ Four Responsibility Loop（工作如何流動）。**Delegation 大幅降級——它是一種 deployment，不是底層不變量。**
5. **核心假說改名**：Evidence-driven Control Loop 是**表面結構**；真正浮現的核心 = **Evidence Responsibility Architecture（ERA）**——跨域最終都回答同四問（誰產哪類證據 / 哪些證據足以支持哪類結論 / 誰不能靠自產證據關閉 / 裁決權屬誰）。已有跨域實證（成功域 **N=3**：sd + Research + Architecture；2s 帶 orchestrator 越界疤仍 pattern-held），仍需 Knowledge 與**反例**。
6. **方法論轉向（下一步）**：**falsification-first**——不急著找第三個成功域，**刻意選預期會失敗的域**（Brainstorming / Creative Writing / Open-ended Design / Ideation）做下一個 run：若它們也自然長出 evidence-first acceptance + independent evidence + closure authority → 假說極強；若沒有 → 明確畫出適用邊界。**兩種結果都是有效裁決素材，比繼續累積成功案例更有研究價值。**仍須真實任務（自然出現的 ideation 需求），不 manufacture。

**第五輪 review（使用者，2026-07-09）——研究階段轉換：accumulation → boundary discovery**：

1. **預註冊的反漂移理由明文化**：無預註冊會落入「無限弱形式回歸」——沒有 backfill → 說有弱形式；沒有弱形式 → 說有更弱形式；最後任何東西都算同一 pattern，**假說永遠不可被推翻**。預註冊 = 先定義 failure 長什麼樣。
2. **Falsification 拆成兩個獨立觀察（2f 預註冊判準，run 前定死）**：
   - **F1**：acceptance 是否**自然形成 evidence requirement**（不誘導）。
   - **F2**：closure 是否**真的依賴 independent evidence**——證據可以存在但最終以偏好/品味/創意/領導裁決關閉（「我喜歡第 7 個」）＝ F2 fail。**Evidence 有、closure 不靠它，也是 failure。**
   - 判讀：F1✓F2✓ = ERA 成立於該任務；F1✓F2✗ = 證據裝飾性，邊界訊號；F1✗ = 邊界訊號。
3. **ERA 最終假說收斂為單一問句**：「**Decision 是否必須依賴 Independent Evidence？**」Yes → ERA 成立；No → 邊界。（四步 loop 是這個問句的展開，不是假說本體。）
4. **邊界維度猜想（待 2f 驗證）**：ERA 可能不是「所有 AI 工作」的 primitive，而是 **High-integrity Work**（sd / research / audit / compliance / architecture governance / knowledge management——共同點：decision 必須可 justify）的 primitive。**邊界不是 domain 名，是工作性質：Justification Required vs Preference Allowed**——同一 domain 可同時含兩種工作（architecture 安全審查 vs architecture 概念發想）。此分類比 domain 更穩定。
5. **成熟標誌重定義**：好的架構模式成熟的標誌不是「可以用在所有地方」，而是「**知道它在哪裡不該用**」——最期待的不是 2f 成功，是 2f 依預註冊標準**真實失敗**且失敗原因正是「decision 只需偏好裁決」。

**第六輪 review（使用者，2026-07-09，讀 2f 期中資料後）——ERA 假說重構：從「決定」到「約束」**：

1. **三層結構取代二分法**（2f 期中資料直接支撐）：(i) Producer 無 contract 也**自發**產生 evidence awareness（撞名註記、截點聲明）——**Evidence Requirement 不是 evidence awareness 的唯一來源**，此為獨立記錄的觀察；(ii) Reviewer 的功能是 **Filter 不是 Selector**（全部輸出是「這不能選、這有事實錯誤、這類別有盲點」，零「我喜歡」）；(iii) 真正的 Selection 在使用者。
2. **ERA v2**：不是 `Evidence → Decision`，是 **`Evidence → Decision Space → Preference → Selection`**——**Evidence 約束 Decision，不必然決定 Decision**。此抽象更強，因為它把各工作型態放上同一光譜：
   | 工作型態 | Evidence 約束強度 | Selection 空間 |
   |---|---|---|
   | Software Delivery | 完全約束 | 幾乎不存在（唯一合法答案） |
   | Research | 高度約束 | 剩一個最可信答案 |
   | Naming | 部分約束 | 排除危險候選後由偏好選 |
   | Brainstorm / Creative | 弱約束 | 避開明顯失敗後由品味形成 |
3. **邊界重定義**：ERA 的 boundary 不是「有/沒有 evidence」，是「**Evidence Constraint 有多強**」——continuum 取代 binary。
4. **F2 三模式量測尺度**（本輪提出時 F2 尚未量測，時序合法；**記為補充量測，原始二元判準仍為本 run 正式判定基礎**）：Evidence-determined（證據幾乎唯一決定）/ **Evidence-constrained**（證據縮小範圍、偏好完成選擇）/ Preference-determined（證據幾乎無影響）。
5. **反事實鑑別問題**（F2 量測的關鍵細化）：若最終選了某候選，須區分「reviewer 排除其它後它成為最佳剩餘」vs「即使 review 不存在本來就會選它」——**同一個選擇、對 ERA 意義完全不同**。

**第七輪 review（使用者，2026-07-09）——Feasible Set 形式化 + v3 候選假說**：

1. **v2 真正強的原因**：不只解釋更多 domain，是解釋更多 **Decision 類型**——Decision 不是單一步驟。sd 實際鏈是 `Evidence → 哪些不能 Close → fix/defer/reject → Closure`；Evidence 沒有直接輸出 Decision，是先定義 **Decision 的合法範圍**。
2. **形式化（feasible set）**：`Decision Space → Evidence Constraints → Feasible Set → Decision`（All Possibilities → Feasible → Chosen）。Decision 不從零開始，Evidence 持續縮小可行集。
3. **C1b 重讀為 constraint 的典型範例**：evidence 不說「要選 integration」，只說「**不能只選 inner**」——排除而非指定。
4. **ERA 名稱自洽**：Responsibility = **誰有權縮小 Decision Space**。Executor 增加 implementation evidence、Verifier 增加 independent evidence、Orchestrator 做最終 Selection——每個角色都是 Decision Space 的操作者。
5. **第二維度猜想（未驗證，記錄待觀察）**：Constraint Strength 不是唯一維度，還有 **Constraint Type**——sd 的 evidence 主要「排除非法」（Illegal）、naming 主要「降低風險」（Risk）、research 主要「增加可信度」（Confidence）、creative 剩「Preference」。**呼應（orchestrator 補充）**：各域自然長出的 finding 分類恰是 constraint-type 標記——sd `acceptance-violation`＝Illegal 型、naming review 的撞名/文化坑＝Risk 型、research 查核的引文覆核＝Confidence 型；分類 enum 未經協調卻對上猜想分型，是此維度真實存在的間接信號。
6. **v3 候選假說（明標：不入正式定義，下一階段待驗證）**：`v1: Evidence → Decision`／`v2: Evidence constrains Decision`／**`v3(候選): Evidence progressively shapes the feasible decision space`**。開放問題：不同類型的證據是以不同機制（排除非法／降風險／升可信度）塑造可行集，還是存在更統一的機制？——ERA 下一階段的核心待驗證問題。

**第八輪 review（使用者，2026-07-09）——理論收斂判定 + 抽象凍結（stakeholder 決策）**：

1. **理論收斂特徵判定：開始有了**。依據不是假說變漂亮，而是每輪修正共有的特性：**新抽象不推翻舊資料，而是讓舊資料成為新模型的特例**（delivery 的 `Evidence → Decision` = feasible set 只剩一個元素的極限情況）；每輪產生新的可反證預測；假說不靠改寫歷史成立。
2. **修正第七輪：Preference 不是 Constraint Type，是 Selection Policy**——Illegal / Risk / Confidence 都在「限制可行集」，Preference 是「從合法集合裡挑一個」，屬不同機制、不同層。修訂模型：
   ```text
   Decision Space
     ↓ Evidence Constraints（Illegal / Risk / Confidence / …）
   Feasible Set
     ↓ Selection Policy（Preference / Optimization / Random / User Choice）
   Decision
   ```
3. **Responsibility 二分（讓 ERA 的 Responsibility 第一次完整）**：**Constraint Responsibility**（誰能改變 Feasible Set——producer 加 evidence、verifier 加 independent constraints）vs **Selection Responsibility**（誰能從可行集合做最終選擇——orchestrator / decision holder 的 selection policy）。
4. **不急 enum 化** constraint types（Illegal/Risk/Confidence 目前是觀察分型，非 schema）。
5. **抽象凍結（stakeholder 決策，2026-07-09，效期至 Phase 3 / 2026-08-31）**：**在 Phase 3 前刻意不再提升抽象層級**。v2（Evidence constrains Decision）= 正式工作模型；v3 = 維持候選；後續只做證據收集——特別觀察 **Constraint Responsibility 與 Selection Responsibility 是否在新案例中自然分離**。v3 升格條件：3–4 個不同領域重複出現此分離。理由：目前「每輪能回頭解釋舊資料 + 產生可反證預測」的優勢，會被急著升格破壞。

**Constraint-type 猜想的第一個 production incident 實證（consumer bookmark 案，2026-07-09；外部 review 回饋入檔）**：

- **事實**：mock IT ✅ + build ✅ + grep ✅ → UI 開啟即 missing column（dev DB 未跑 migration）。錯誤關閉的根因**不是 verifier 少做一步，是缺一種 constraint type**——既有證據全屬 implementation constraint，只能排除 implementation failure；**沒有任何證據有能力約束 runtime decision space**，Feasible Set 沒有真正縮小，Close 在過大的可行集裡做出。
- **對 ERA 的意義**：(a) constraint-type 維度（第七輪猜想）從觀察分型升為**有 incident 後果的實證**——缺 type 會直接導致錯誤 closure；(b) Responsibility 補洞案例：runtime constraint 原本**無人負責** → verifier V5 補位（constraint responsibility 分配的自然演化實例）；(c) v3 的 feasible-set 語言直接解釋了 incident（凍結期內又一次「新模型解釋新資料」）。
- **凍結合規處理**：V5 / `runtime` tier / `runtime-omission` 以 **mechanism** 入 `sd-delegated-execution`（不動抽象）；**constraint-family 重構**（以 Spec / Implementation / Runtime / Delivery / Safety constraint 族取代 V-編號，V1–V5 降為 mechanism——外部 review 建議，防 V6/V7/V8 無限疊 checklist）記為 **Phase 3 解凍後的重組候選**，屆時與 Q8/v3 裁決同批評估。

**第九輪 review（使用者，2026-07-09）——研究線收官凍結（line-level freeze）**：

1. **凍結判定**：研究線暫時凍結——不是沒東西可研究，是它已具備成熟理論雛形的特徵。收斂節奏本身即成果：`Observation → Hypothesis → Cross-domain → Falsification → Boundary → Freeze`——這是 research methodology 的節奏，不是「抽象再抽象」。
2. **每次修正保留舊模型（正式確認）**：v1 = constraint strength ≈ 100% 的極限情況；v2 = 把 strength 顯式化；v3 = 把 Decision 拆成 Constraint + Selection——**沒有任何一層需要否定前面**。
3. **九輪最大成果 = Responsibility 重定義**（比 ERA 本身更重要，比 RACI 更底層）：不是「誰寫 code / 誰 review / 誰 approve」，是「**誰可以改變決策空間**」——Constraint Responsibility（誰有權縮小 Feasible Set：executor 的 implementation evidence、verifier 的 independent constraints、runtime smoke 的 runtime constraints）+ Selection Responsibility（誰有權從 Feasible Set 做最終選擇：PO / maintainer / user / architecture board）。
4. **隱含成果正式記錄：固定住的是 Decision Semantics，不是 Workflow**——sd 的 Close、research 的 Disposition、naming 的 Selection 表面完全不同，但全部化約到 `Decision Space → Evidence Constraints → Feasible Set → Selection Policy → Decision`。**Workflow 可以不同、Decision Semantics 不變——這比 Workflow 更接近 Primitive。**
5. **Phase 3 重定位（stakeholder 裁定）**：不再研究抽象，改為**驗證穩定性**——問題從「還有沒有更高一層」改為「**未來三個月，新案例是否不用修改模型就自然落進目前模型**」。連續多案例成立 = 模型具 **Predictive Power**（而非僅 Retrospective Power）——兩者差異極大。
6. **Post-Phase-3 候選研究主題（登記不展開）**：**Evidence Lifecycle**（`Produced → Challenged → Confirmed → Deprecated → Archived`）——Evidence 目前在模型中是靜態節點，但它怎麼累積/失效是隱含的下一問。此線不建立在新抽象上，而建立在既有研究紀律（temporal evidence / provenance / pre-registration / falsification）之上，很可能自然長出來。**現在不開**：ERA 剛進入穩定收斂期。

**第十輪 review（使用者，2026-07-09）——研究治理層的收官觀察（meta-level，本身依「登記不升格」處理）**：

1. **Freeze-at-peak 是治理模式，不是研究步驟**：「在最想繼續的時候停下來收證據」與一般「沒想法才停」本質不同——前者是**證據紀律**，後者是資源不足。核心句：**「Freeze 不是結束，是把模型固定，讓未來有機會失敗」**——模型一直改就永遠不知道它有沒有預測能力；凍結後 Case D/E/F 全用同一模型，成功與失敗才都有研究價值。
2. **兩種 Freeze 的區分**：**Model Freeze**（v2 固定為正式工作模型）vs **Hypothesis Freeze**（v3 / Evidence Lifecycle / constraint-family——`Interesting → Register → Do Nothing`，等證據）。後者罕見且珍貴：多數研究的最大問題是看到新想法立即展開。
3. **Working Model 的修改閘門 = 證據，不是更好的想法**：討論可以一直發生、hypothesis 可以一直累積，但**正式工作模型只有新的、足以改變模型的證據才能改**。這回答「什麼東西可以改變正式模型」——比「不要再抽象」更根本。
4. **Research Governance Primitive（候選節奏，登記不展開）**：`Exploration → Freeze → Prediction → Accumulation → Revision`——九輪實際走的節奏。
5. **三個最有長期價值的可移植成果（使用者總結）**——即使 ERA 被修正、v3 不成立、Decision Semantics 被取代，這三者仍保留並可用於下一個問題：
   1. **Working Model Freeze**：證據不足時正式模型保持穩定；新想法先登記不升格。
   2. **Pre-registered Falsification**：先定義成功/失敗判準再收資料，防事後合理化。
   3. **Predictive Validation**：凍結後不追更高抽象，觀察新案例是否自然落位。
   → **處置**：三者標記為 **本 plan closure 時的 intelligence atom 提煉候選**（plans/README 原則 5 的既有管道；不新建 governance 文件、不即時升格——與其自述的紀律自我一致）。**九輪最大產出不只是候選理論，是一套能持續產生可靠理論的研究流程。**

**Mutation review（使用者，2026-07-09，Mutation Testing 討論；獨立對話線，還原自 commit 28692fd）**——四個命題，doc-only 回寫：

1. **Mutation testing = V3 的 evidence generator，不是新驗證層（不設 V6）**。架構優勢在於以 evidence 為中心而非以 testing 技術為中心；mutation 只是「一種產生 Evidence 的方法」，架構完全不用改。已落地：[`delegated-execution.md`](../../../workflow/software-delivery/delegated-execution.md) §5 V3 evidence producer；[`test-strategy.md`](../../../workflow/software-delivery/test-strategy.md) §Mutation 加 verifier-consumer back-pointer。
2. **L3 從 imagination-driven 升級為 mechanical falsification**——mutation 機械枚舉行為區分點，與「Mechanical Enforcement > Human Discipline」同構；verifier 不再需要自己想到 `price==100`。
3. **Survived mutant 只是資訊，finding 才是 evidence**：Mutant → Semantic Gap → Verifier Finding；orchestrator 不需知道 mutation engine 存在。
4. **不做 mutation score KPI**；保留 targeted mutation（risk-triggered：boundary / boolean / null / authorization / invariant / guard）。
5. **通用抽象候選（→ Q9）**：「Behavioral Falsification」producer family（mutation / fault injection / property-based / model-based，皆產出「此行為未被驗證區分」型 evidence）。**紀律邊界**：forming abstraction（observe-only）；graduate 前不建 producer registry、mutation 以外 producer 無真實 run 前只是 analogy。

**Writeback-collision 事件（2026-07-10，本 plan 自身的 ERA 實例——負向）**：

- **事實**：兩個併發 evidence session 從陳舊底稿回寫 `_plan.md` / kit——`b6481e5` 覆蓋掉第五~十輪 review、ERA v2/v3、兩層凍結、2e/2f/2g run 紀錄與 glossary terms；`bfb2704` 帶未解決 conflict markers 入 commit；`0958a38` 解衝突時再丟失 Q9。**內容全數自 git 歷史還原**（e2d5091 / bfb2704 / 66f58ed），run 紀錄依新 evidence/ 慣例落檔。
- **對 ERA 的意義（自我指涉實例）**：plan writeback 本身就是一種 Production——這兩次 push 是「**自產證據自我關閉**」的實例（無獨立 verifier 覆核 diff 即 push），與 2j 的「單 Task 跳過 Verifier」同構、發生在治理 repo 自身。**模型再次正確診斷斷裂形狀**（Production 與 Evidence 合併）。
- **契約回饋（登記，mechanism 候選）**：(a) evidence-only session 應**只 append `evidence/` 檔 + 一行 plan checkbox**，不得整檔重寫 `_plan.md`（93bde60 的 evidence/ 分離方向正確，需成為紀律）；(b) push 前 `git pull --rebase` + 衝突不得以「取我版」解掉共享敘事檔；(c) **conflict-marker pre-commit scan** 為可機械化候選（Go-first，涉 CLI——登記不即時實作）。

**第十一輪 review（使用者，2026-07-10，讀 collision 還原後）——治理自我適用的高價值驗證（非新假說，凍結合規）**：

1. **Collision 是 Governance 失敗，不是 Git 問題**：merge conflict / 資料遺失只是工程表象；真正的 failure 結構 = Evidence Session 修改 canonical plan、無 Independent Review、直接 Push——**Production + Evidence + Closure 由同一 actor 完成**，與 Evidence Producer ≠ Closure Authority 完全同構。Remediation（scan / rebase）只是補救，不是解釋。
2. **Structural Predictive Power（結構預測力）**：Predictive Power 的第二形態——模型不預測事件（merge / overwrite / 錯誤 approve / collision 都只是 manifestation），預測**失敗結構**：「Production 與 Evidence 不分離 → 治理失敗」。與「新案例自然落位」並列為 Phase 3 穩定性量尺的兩翼。
3. **三條契約回饋的正確定性**：它們不是 patch，是 **Constraint Responsibility 的恢復**——append-only = 阻止 Evidence Session 修改 Canonical Decision；rebase = 重新接受新的 Constraint；conflict scan = 阻止未完成的 Decision 進入 Canonical。三條都是 ERA 在治理層的具體化。
4. **Q5 的機械化邊界明晰**：「擋寫不擋不驗」揭示——**Mechanization 只能約束 Execution，不能產生 Evidence / Judgement**。工具可以禁止直接寫，不能生成 Independent Evidence → Constraint / Selection Responsibility **不可完全自動化**（Q5 Phase 3 門檻明文化的定性基礎）。
5. **關鍵升層句（使用者，值得長期保留）**：「**Canonical Writeback 本身就是一種 Selection**」——寫回 canonical 不是 IO，是 `Candidate Knowledge → Canonical Knowledge` 的 Selection Policy 行為。Evidence Session 直接寫 canonical 不是「多做一步」，是**跨越 Responsibility Boundary**。適用於 plan、Knowledge Base、Glossary、ADR、Pattern Library——所有 canonical 面。
6. **研究線最新評價**：前九輪建立 Decision Semantics；本次首次出現「**Governance 本身也服從同一套 Decision Semantics**」——模型開始能約束自己的演化。**當一個模型能解釋自己的失敗模式而不需發明新解釋框架，是理論成熟度提升的重要訊號**。定性：目前工作模型（v2 + v3 候選）的高價值驗證，非新假說——凍結不動。

**第十二輪 review（使用者，2026-07-10，證據充足性對帳後）——Phase 3 決策框架定稿**：

1. **Open questions 三分類（Phase 3 開批照此裁，勿混）**：**Evidence Complete**（Q5 / Q6 / sd 重定位——可裁決只等批次；Q5 已無研究價值：新 delivery case 不可能推翻「gate 能限制寫入、不能替代 verification」，2j + collision 兩域負向已足；Q6 從 "Can we rename" 變成 "**Do we rename now**"——governance judgment 非 research）／**Evidence Pending**（Q8-F2 / Q9——缺**不可替代**資料：F2 只有使用者能產生，Selection Responsibility 在 user，任何 agent 不能代替；Q9 mutation producer 零真實 run）／**Time Validation**（v3 / predictive power——**缺的是時間不是案例**：domain coverage 已近完成，若 8 月中出現需改 feasible-set 模型的案例 predictive power 即降，這不是 domain 能補的）。
2. **研究治理原則（登記）**：**Evidence Saturation ≠ Observation Completion**——「飽和的閘門放著不會壞，時間窗補不回來」；多數研究證據夠了就急著 publish，本線要的是 **Prediction Window**。
3. **Time 是 Observation Constraint（非 Evidence Constraint）**：Working Model 受兩種 constraint——**Evidence Constraints 決定今天**（哪些結論成立）、**Observation Constraints 決定是否可升格**（freeze → 未來案例 → 仍支持，才解鎖修改權）。今天所有案例都支持 ≠ 模型穩。
4. **反過早 Closure 裁定（stakeholder，2026-07-10）**：Q5/Q6 雖可裁，**維持 8/31 一起裁**——非行政理由：讓 Working Model 有完整的**無修改觀察期**；「模型足以應對新案例而無需修正」的信心**靠增加案例數量無法替代**。

**Collision 復發（N=3，2026-07-10）——機械化升為 due**：`2de1686`（Cursor session）三度自陳舊底稿覆蓋 `_plan.md`（-146 行：第五~十二輪 + Q7/Q8/Q9 回退），已自 `dcd6f9e` 重建。行為層契約回饋（append-only / rebase-before-push）已證明**擋不住跨工具 session**——依 [`failure-to-validator-closure`](../../../enforcement/failure-patterns/failure-to-validator-closure.md)，N=3 達機械化門檻：**canonical-narrative shrink guard**（commit 對 `plans/active/*/_plan.md` 刪除行數超閾值時 block，除非 `[plan-restructure]` opt-in）+ conflict-marker scan，開獨立 task 實作（Go-first，涉 CLI）。

**外部相鄰模式對照（2026-07-10，使用者提問觸發；observe-only，Q6 定位參照）**：

- **對象**：[claude-skills-llm-council](https://github.com/aiwithremy/claude-skills-llm-council)（Karpathy LLM Council 方法論的 Claude skill 版；全 repo = README + SKILL.md 單 prompt 檔）——一問題 → 5 個思考角度顧問（Contrarian / First Principles / Expansionist / Outsider / Executor）獨立分析 → 匿名互評 → 主席綜合裁決；明文限定判斷型決策、排除創作與處理任務。
- **表面相似**：多 agent 角色分離 + 互評環節 + 最終裁決者；Outsider 零 context ≈ fresh-context 的弱形式。
- **本質差異（ERA 語言）**：Council 流通的是**意見**（opinions 互評 opinions，無事實查核、無 acceptance 量尺、無可覆核 evidence）；closure 為 **preference-based**（主席綜合即結束）；不存在 Evidence Producer ≠ Closure Authority。定位：**preference-allowed 決策空間的 Selection 輔助工具**（活在約束光譜弱端），與本 loop（justification-required 執行的證據治理，強端）是**同剪影、不同家族**——governance veto test 教科書案例：去掉 domain 內容後，Council 剩「意見聚合」、本 loop 剩「證據責任分配」。
- **「基底」判定**：對其自身目標成熟可用；對本 plan 目標**不構成基底**（無契約 / 證據 / 仲裁 / 關閉 / gate 任一層）。**吸收方向（登記不展開）**：Council 可作為本框架的 technique——Specification 階段的決策空間發散（orchestrator 拆 slice 前的方案比較）、或 preference-allowed 任務的 Selection 輔助（接 2f filter/selector 發現）；依 register-don't-promote，待真實使用需求出現再落。

**doc-only trial 宣告**：本 plan 不接入 runtime——不新增 `route.*`、不新增 commit-msg validator、不動 `runtime.db` generated surfaces、不動 delegation schema / `validatePlanTreeFrontmatter`。協議以文件 + 行為紀律承載；驗證 leg 復用既有 review capability invoke（`ai-skill runtime capability-invoke --capability code-review --stance fault_finding`，既有 warning-only surface，無新 wiring）。

**未來接入條件（graduation）**：Phase 3 證據評估時決策——若 (a) 三角色 loop 在 ≥2 個真實任務有效、且 (b) role boundary invariant 出現行為維持不住的證據（如 verifier fresh-context 被反覆略過），才評估 schema 欄位（`delegation.verification`）或機械檢查；由後續 plan 承載，本 plan 不 carry。**決策 deadline：2026-08-31**（與本 plan closure 同批；未達證據門檻則明確記錄「維持 doc-only」）。

## Open Questions

| ID | Question | Owner / Resolved By | Status | Closed Criteria | Resolution Evidence |
|----|----------|---------------------|--------|-----------------|---------------------|
| Q1 | Verifier 報告最小欄位集（evidence / acceptance_ref / classification / status）是否足以讓 orchestrator 仲裁而不回讀 diff 細節？ | Phase 1 定稿、Phase 2 驗證 | **resolved（2026-07-08）** | 雙 dogfood 中 loop 後 orchestrator 均未被迫回讀 diff；缺欄位已補進契約 | 2b + **2a-external（外部 repo）**：兩輪 verifier 報告自足、仲裁未回讀 diff；2a-external 有 loop 前 orchestrator 寫 code 越界（commit `<HASH-a>`），屬 role boundary 非報告欄位問題 → mechanical reminders 已補 |
| Q2 | 協議文件落點：plans/README.md §Delegation 擴充（delegation 擁有 loop）vs `workflow/cross-cutting/review/` consumer doc（review 擁有驗證 leg）？stance 復用不得重定義 | Phase 1 | **resolved（2026-07-08）** | 落點決定 + 文件落地，且未在 consumer 層重定義 stance / requires_context ✅ | 落點 = plans/README.md §Delegation loop SOP 子節（delegation 擁有 loop 生命週期；review 只被 invoke 引用）。commit `af26064` + `2d5bc60`；獨立驗證確認未重定義 stance。副作用：F1 措辭 drift 隨 canonical 移轉消解 |
| Q3 | 品質信號怎麼量：verifier 差集 findings 數 + 仲裁分佈（fix/defer/reject 比例）是否構成「品質提升」的有效指標？null result 如何記錄？ | Phase 2 | **resolved（2026-07-08，advisory 指標）** | 雙 dogfood 各留差集 + 分佈；複合指標明文化 | **複合指標**（kit §2a-external 結論表）：(1) acceptance-violation 率（2a-external **0/2 rounds**）；(2) test delta（+6）；(3) pre-merge bug fix 數（2：guard + envelope）；(4) 協調成本（spawn×4、plan commit×6）；(5) orchestrator 越界次數（1）。**結論**：品質↑有量化證據；orchestrator 寫 code↓、協調↑；verifier 邊際 catch 本任務為中等（強制 IT/結構化 defer 價值 > acceptance 差集）。null result 未出現 |
| Q4 | 仲裁紀錄落點：被委派 sub-plan 內 table（傾向）vs 獨立 artifact？ | Phase 1 | **resolved（2026-07-08）** | 落點決定並在 dogfood 實際使用 ✅ | 落點 = 被委派任務的 plan artifact 內 table（SOP 已載明）；dogfood 期記於 kit §Dogfood 紀錄（2b 仲裁表實際使用） |
| Q5 | Schema promotion 門檻：什麼證據才允許動 delegation schema（如 `delegation.verification`）？ | Phase 3 | open | 門檻明文化；未達門檻則明確記錄維持 doc-only | kit §2c + **§2d** 增強信號；consumer 機械 gate（2c/2d）後 orchestrator 零 manageCode diff；**§2g 第二個獨立 consumer**（ExternalRepoA overlay + backfill + 自建 gate，BDD 7/7）——consumer-layer 自理模式 ×2 成立，slice 各項 promotion eligibility 達標、評估歸 Phase 3 批次 — schema 本身**仍維持 doc-only**。**2j 負向補充（2026-07-10）**：機械 gate「擋寫不擋不驗」缺口實證——role boundary 需 **Verifier-spawn tracking** 類 gate（consumer `verifier-after-executor` 回饋），Phase 3 門檻明文化的直接素材 |
| Q6 | 通用化定位：graduate 時是否以「Evidence-driven Closed Control Loop」（四責任分離：Specification → Production → Independent Evidence → Arbitration → Specification）取代「Delegation」定位？（使用者 review 2026-07-08 提出，見 §架構收斂觀察） | Phase 3（adoption stage 2 gate） | open（**stage-2 進行中 2/3**，2026-07-14；改名裁決留 Phase 3 / maintainer） | Research / Knowledge / Architecture 各一輪真實 run；**驗 pattern 不驗 topology**；stage 3 另需 cross-domain + cross-workflow + cross-project，不在本 plan scope | **2e Research** ✅；**2s Architecture** ✅（UI Pattern Knowledge plan review；Task Verifier；orchestrator 越界疤；**不**填 Knowledge）——pattern held, topology differed；**Knowledge 仍缺** |
| Q7 | **Verification Backfill 是否為獨立 primitive（Evidence-first Execution）**：「acceptance 在 execution 前映射成證據」是否比 delegation 本身更根本？（第三輪 review 命題 2） | Phase 3 / stage 2 觀察 | open | (a) sd 域內：backfill 在 ≥2 個真實委派任務穩定使用且能擋「做完再想怎麼驗」；(b) 跨域：至少一個非 delivery 域**自然出現或明確缺席**「execution 前的 acceptance→evidence 映射」——缺席也是有效答案（支持「backfill 是 sd-specific，不是 primitive」） | **resolved（2026-07-09）**：(a) sd 域 ≥2 真實使用成立（2c/2d）；(b) 2e Research 域觀察——**結構化 backfill（tier+owner 表）明確缺席且不需要**，出現的是弱形式 **evidence-first acceptance**（acceptance 內嵌證據標準）。**結論：Verification Backfill 是 sd 域強形式 primitive，非 universal；universal 候選改為 evidence-first acceptance**。Reopen 條件：任一非 delivery 域自然長出結構化 tier+owner 映射需求 |
| Q8 | **Evidence Responsibility Architecture（ERA）是否為更底層共同骨架**：跨域是否自然收斂出相同的「證據責任分配」結構——誰產生哪種證據 / 哪種證據足以支持哪類結論 / 誰不能靠自產證據關閉 / 裁決權屬誰？（第三輪 review 命題 4 立案；**第四輪 review 命名 ERA + invariant candidate「Evidence Producer ≠ Closure Authority」**） | Phase 3 / stage 2 觀察 | open（成功域 **N=3**：sd + Research + Architecture；**2f falsification** 預註冊，[`evidence/2f-falsification-naming-run.md`](evidence/2f-falsification-naming-run.md)；Knowledge 未跑） | **雙向裁決**：(a) 成功域同構累積（不預設 sd 詞彙）；(b) **反例探測（2f，判準已預註冊）**——F1 + F2；目標選擇依**工作性質**（preference-allowed）非 domain 名。邊界維度猜想：Justification Required vs Preference Allowed | **2e** 第一個跨域同構；**2s Architecture** 四問重現（見 [`evidence/2s-architecture-ui-pattern-knowledge-plan-review.md`](evidence/2s-architecture-ui-pattern-knowledge-plan-review.md)）。Primitive 候選排序：ERA > Evidence-first Acceptance > Independent Verification > Four-Responsibility Loop；Delegation 降級為 deployment。負向同構 ×3（2j/writeback-collision/2l）仍有效 |
| Q9 | **Behavioral Falsification 是否為 V3 evidence producer family**：mutation / fault injection / property-based / model-based 是否收斂為可替換 producer（皆產出「此行為未被驗證區分」型 evidence）？（Mutation review 命題 5，見 §架構收斂觀察） | Phase 3 / 後續 delivery dogfood | open | (a) targeted mutation 作為 V3 producer 在 ≥1 個真實委派 run 實際使用，且 survived mutant → semantic-gap finding 契約成立（orchestrator 未被迫理解 mutation engine）；(b) 至少第二種 producer 自然出現於真實 run——缺席亦為有效答案（family 維持 mutation-only，不建抽象）；graduate 前不建 producer registry / 通用 taxonomy | <V3 producer run evidence> |

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

## Phase 2 — 雙 dogfood

**焦點紀律（承 03）**：驗的是**契約自足性**（verifier 報告是否足以仲裁、brief 是否足以執行），不是執行者 / 驗證者聰不聰明。任何失效先反問契約缺漏。

**Transport（tool-portability 驗證，2026-07-08 調整）**：雙 transport 分工——**2a 用 Cursor**（human 路徑：使用者把模板貼進獨立 fresh chat）、**2b 用 Claude Code Agent 工具**（agent 路徑：orchestrator session 直接 spawn executor / verifier 獨立 agent，使用者授權全程自駕 + 每階段 commit/push）。同一套模板跑兩種 transport = 「模板 tool-neutral、工具細節只在傳輸備註」分層成立的更強證據。模板見 [`01-dogfood-prompt-kit.md`](01-dogfood-prompt-kit.md)。

- [x] Dogfood prompt kit 建立（模板 A executor / B verifier / C 仲裁+量測欄，Cursor-first；kit 為 rendered transport artifact，canonical 在本檔 §Decision Rationale）— 2026-07-08
- [x] **2a — software-delivery 外部 repo 真實任務**：挑一個真實、小、可驗收的交付任務走完整 loop（orchestrator 寫 brief → executor 執行 → verifier 報告 → 仲裁）。記錄：verifier 差集 findings、仲裁分佈、orchestrator 是否越界執行、是否被迫回讀 diff。 — 2026-07-08：**2a-external** 外部 monorepo sync-adapter Step 6（kit §2a-external）；另 2a demo SD read-only（kit §2a）
- [x] **2b — Ai-skill 內部任務** ✅（2026-07-08，agent transport）：委派任務 = Phase 1 SOP 擴充本身。完整 loop 走完：brief → executor（worktree，自驗 8/8，讀檔零差集）→ verifier（fresh，findings ×3 全 observation 級，0 violation）→ 仲裁（fix 1 / defer 1 / reject 1）→ fix 回派（brief v2）→ delta 重驗 pass → merge。**Q1 正向證據：orchestrator 全程未回讀 diff，僅憑 verifier 報告仲裁**。完整 evidence 見 [`01-dogfood-prompt-kit.md`](01-dogfood-prompt-kit.md) §Dogfood 紀錄 2b。
- [x] **2a-external — 外部 monorepo sync-adapter Step 6** ✅（2026-07-08，Cursor Task transport）— kit §2a-external
- [x] **2c — 外部 monorepo tiered archive 全線（8 slices，Phase A–D）** — 2026-07-08 **證據 only**（kit §2c）；強化 Q3 品質信號，**不**視為 Phase 3 closure
- [x] **2d — 外部 monorepo outbound sync Phase 3（4 slices）** — 2026-07-08 **證據 only**（kit §2d）；強化 backfill / consumer gate 信號；**不**視為 Phase 3 closure
- [x] **2g — 第二個外部 consumer（ExternalRepoA）：server_doc test placement + delegation overlay** — 2026-07-09 **證據 only**（kit §2g）。雙重意義：(a) **Phase 3 穩定性視角第一筆資料**——新案例未修改模型自然落位（同 overlay + backfill 模式、ERA 分工自然出現、consumer 機械 gate 自理）；(b) sd-delegated-execution §Provenance 的升級條件「第二個獨立 consumer 真實使用」**已滿足** → 分類 enum / backfill 模板化 / 機械 gate 泛化的 promotion **eligibility 成立**，依凍結紀律評估延至 Phase 3 批次，一次一階
- [x] **2d′ — ExternalRepoC 9j2 module alignment follow-on** — 2026-07-09 **證據 only**（[`evidence/2d-prime-externalrepoc-module-alignment.md`](evidence/2d-prime-externalrepoc-module-alignment.md)）。§2d 同一 consumer 延續：integration gate（平行 branch UX fail ×1）、`remote_absent_delete` fix、live 雙邊 teardown、pre-push build + inner src/test block；模型自然落位 **是**；**不**視為 Phase 3 closure
- [x] **2h — ExternalRepoC common-url Execute 驗證缺口** — 2026-07-09 **證據 only**（[`evidence/2h-externalrepoc-common-url-verification-gaps.md`](evidence/2h-externalrepoc-common-url-verification-gaps.md)）；RBAC 三連、V5 全 API 面、combined 不得 defer L1–L3；**不**視為 Phase 3 closure
- [x] **2i — ExternalRepoC user-feedback S0–S4 Execute** — 2026-07-09 **證據 only**（[`evidence/2i-externalrepoc-user-feedback-pull-execute.md`](evidence/2i-externalrepoc-user-feedback-pull-execute.md)）；Stop/resume、inventory gate、2h 教訓遷移驗證；**不**視為 Phase 3 closure
- [x] **2j — ExternalRepoC push Execute 跳過 Verifier loop** — 2026-07-10 **負向證據 only**（[`evidence/2j-externalrepoc-push-execute-skip-verifier-loop.md`](evidence/2j-externalrepoc-push-execute-skip-verifier-loop.md)）；Execute 意圖 > `delegation.enabled:false`（F2 stakeholder 裁決）；單 Task 包辦 = Production/Evidence 合併；consumer `verifier-after-executor` gate 回饋；**不**視為 Phase 3 closure
- [x] **2k — ExternalRepoC push 纠偏后 post-close runtime 缺口** — 2026-07-10 **證據 only**（[`evidence/2k-externalrepoc-push-post-close-runtime-gaps.md`](evidence/2k-externalrepoc-push-post-close-runtime-gaps.md)）；用户手验暴露 V5 未覆蓋 create 表单 / Worker 拓扑；`post-close-surgical-debt`；**不**視為 Phase 3 closure
- [x] **2l — ExternalRepoC common-url S2′ mirror 再跳过三角色 loop** — 2026-07-10 **负向证据 only**（[`evidence/2l-externalrepoc-common-url-s2-mirror-skip-loop.md`](evidence/2l-externalrepoc-common-url-s2-mirror-skip-loop.md)）；0 Executor/Verifier、surgical bypass 滥用、Shell 绕过 preToolUse；`retroactive-r1-verifier`；**不**視為 Phase 3 closure
- [x] **2m — ExternalRepoC Phase G-mirror 批量 retrofit** — 2026-07-10 **正负对照**（[`evidence/2m-externalrepoc-phase-g-mirror-batch-retrofit.md`](evidence/2m-externalrepoc-phase-g-mirror-batch-retrofit.md)）；V-m1–V-m5 + 登记总表；02/01 合规 loop vs 03/2l；`retrofit-v-m-template` / `stale-jvm-v5-a-checklist`；**不**視為 Phase 3 closure
- [x] **2n — ExternalRepoC 07 push delivery DEL-S1–S6 合规 loop** — 2026-07-10 **正向证据**（[`evidence/2n-externalrepoc-push-delivery-s1-s6-compliant-loop.md`](evidence/2n-externalrepoc-push-delivery-s1-s6-compliant-loop.md)）；6/6 slice E+V、sub-plan `completed`、零 post-close bypass；对照 2j/2k/2l；**不**單獨視為 Phase 3 closure（Q5 仍 open）
- [x] **2o — consumer tab-scroll：单 session vs 三角色 对照** — 2026-07-10 **證據 only**（[`evidence/2o-consumer-tab-scroll-single-vs-delegation.md`](evidence/2o-consumer-tab-scroll-single-vs-delegation.md)）；partial authority / deploy smoke ≠ L3；Q8 ERA 信号；**不**視為 Phase 3 closure
- [x] **2p — ExternalRepoC Integration 默认切流 INT-D0–D5** — 2026-07-13 **正向证据**（[`evidence/2p-externalrepoc-integration-default-cutover-d0-d5.md`](evidence/2p-externalrepoc-integration-default-cutover-d0-d5.md)）；6/6 E+V、same-branch 连续、一口气压力未跳 Verifier；live defer；**不**單獨視為 Phase 3 closure（Q5 仍 open）
- [x] **2q — ExternalRepoC transport inner-only runtime gap** — 2026-07-13 **负向/纠偏**（[`evidence/2q-externalrepoc-transport-inner-only-runtime-gap.md`](evidence/2q-externalrepoc-transport-inner-only-runtime-gap.md)）；loop 绿≠路径通；features+L3+V5；**不**视为 Phase 3 closure
- [x] **2r — <PROJECT_ROOT> player overlay Mode A hit-trap** — 2026-07-13 **负向证据**（[`evidence/2r-consumer-player-overlay-mode-a-hit-trap.md`](evidence/2r-consumer-player-overlay-mode-a-hit-trap.md)）；soft-nav 绿→cold URL 全死；programmatic click / 单入口假绿；consumer entry-path 矩阵回馈；**不**视为 Phase 3 closure
- [x] **2s — 跨域 run（Architecture 域）：UI Pattern Knowledge plan review** — 2026-07-14（[`evidence/2s-architecture-ui-pattern-knowledge-plan-review.md`](evidence/2s-architecture-ui-pattern-knowledge-plan-review.md)）；R1+R2 Task Verifier；stakeholder 仲裁；**不**填 Knowledge 格；**不**視為 UI Pattern Knowledge Phase 1 完成；**不**視為 Phase 3 closure
- [x] **2v — greenfield consumer Phase 2 preflight** — 2026-07-16（[`evidence/2v-external-greenfield-consumer-phase2-preflight.md`](evidence/2v-external-greenfield-consumer-phase2-preflight.md)）；brief+backfill 寫入 `<PROJECT_ROOT>` plan evidence；0 E+V；配對 domain-model execute；consumer verify OK；**不**視為 Phase 3 closure
- [ ] **2t — APK Analysis ↔ Software Delivery Capability Handoff（預註冊，2026-07-14）** — 契約與雙軌設計見 [`04-apk-capability-handoff-boundary.md`](04-apk-capability-handoff-boundary.md)。**2t-A**：真實 APK Discovery（candidate workflow；Decision = *Can we explain?*；**不**強制三角色）。**2t-B**（僅 Capability Assessment = Yes）：Capability Proposal → SD Intake → **既有** Delegated Execution。預註冊 F1–F4。**不**填 Knowledge；**不**宣稱 APK Delegated Execution 已驗證；**不**視為 Phase 3 closure。啟動：指定 consumer `<PROJECT_ROOT>` 後開 `evidence/2t-…`
- [x] **2e — 跨域 run（Research/Audit 域）：grandfather sunset audit** ✅ — 2026-07-08–09（[`evidence/2e-grandfather-sunset-audit.md`](evidence/2e-grandfather-sunset-audit.md)）；Q6/Q7(b)/Q8 的 stage-2 裁決 run。完整 loop：調查者（worktree，252 行報告 `c8ff035`，中斷後 resume 完成）→ 事實查核者（fresh，引文逐條命中、5 surfaces 獨立重跑一致、findings ×2 全 observation）→ 仲裁（defer×2，無 fix）。**實質產出**：5/5 surfaces 已 wired、flag 條款過時、延展不觸發、sunset 只剩行政收尾（處置決定保留 maintainer，見 `02-grandfather-sunset-audit.md`）。**跨域觀察**：四責任自然成立（topology 不同：+maintainer 第二層 decision）；backfill 結構化形式明確缺席、弱形式（evidence-first acceptance）出現；證據責任四問同構重現（含「自產證據不能自我關閉」跨域不變式）——詳 kit §2e Q6/Q7/Q8 觀察表
- [x] 回饋迴路（2b 觸發 ×1）：F2 暴露 brief v1 缺「reusable doc 目標須含 tool-neutral 措辭條款」→ brief v2 追加 acceptance 9、kit 使用說明補教訓；修契約未修執行者；fix leg 重跑通過。2a 若再暴露缺漏比照處理。

## Phase 3 — 證據評估與收斂

> **定位改寫（第九輪 review，2026-07-09）**：Phase 3 不研究抽象，**驗證穩定性**——每個新案例記錄「是否不修改模型即自然落位」；連續成立 = predictive power 證據。裁決素材照舊（Q5/Q6/Q8 + 下列 checkbox），但評估視角以穩定性為主。
>
> **決策框架（第十二輪 review，2026-07-10）**：開批時按三分類裁決——**Evidence Complete**（Q5 / Q6 / sd 重定位：素材已飽和，照裁）、**Evidence Pending**（Q8-F2 / Q9：開批時若仍缺不可替代資料 → 明確 deferred 附 reopen 條件，不硬裁）、**Time Validation**（v3 / predictive power：以無修改觀察期的落位紀錄裁）。**反過早 Closure**：8/31 前不提前裁 Evidence Complete 項——無修改觀察期本身是 predictive validation 素材。

- [x] 彙整 dogfood evidence，回答 Q3（品質信號成立 / null result）— 2026-07-08：成立（advisory 複合指標）；null result 未出現；**2c 補強**（8-slice、acceptance-violation 2/8、slice 紀律後 orchestrator 零實作 diff）；**2d 補強**（4-slice、outer L1–L3 關閉紀律、gate 後零 manageCode diff）— 見 kit §2c / §2d
- [ ] Q5 決策：schema promotion（另立 plan）或明確維持 doc-only（記錄門檻與未達原因）
- [ ] Q6 決策：通用化定位（Evidence-driven Closed Control Loop vs Delegation）——依非 coding / delivery 域真實 run 證據裁決；無證據則維持 delegation 定位、四責任觀察留檔（§架構收斂觀察）
- [ ] Q7 決策：Verification Backfill 是否升格為獨立 primitive（Evidence-first Execution）——依 sd 域 ≥2 次真實使用 + 跨域出現/缺席觀察裁決；缺席亦為有效關閉（sd-specific）
- [ ] Q8 決策：Evidence Responsibility Architecture（ERA）是否為更底層骨架——**裁決素材 = 成功域同構（已 N=2）+ falsification run（預期失敗域的真實任務，畫適用邊界）**；升格候選順位（使用者 2026-07-09）：ERA > Evidence-first Acceptance > Independent Verification > Four-Responsibility Loop，Delegation 降級為 deployment
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
| Stage 2 觀察鏡頭 | **證據責任分配結構**（Q8 / ERA）：跨域 run 記錄「誰產生哪種證據 / 哪種證據足以支持哪類結論 / 誰不能靠自產證據關閉 / 裁決權屬誰」，不預設 sd 詞彙、不驗角色名 |
| 下一裁決 run | **Falsification-first**（第四輪 review，2026-07-09）：刻意選預期失敗的域（brainstorming / creative writing / open-ended design / ideation）的**真實任務**跑 loop——成功則假說極強、失敗則畫出適用邊界，皆有效；不 manufacture、不再累積成功域 |
| 抽象凍結 | **Phase 3 前不再提升抽象層級**（第八輪 review，2026-07-09）：v2 = 正式工作模型、v3 = 候選；後續只收證據，觀察焦點 = Constraint / Selection Responsibility 是否自然分離；v3 升格需 3–4 域重複出現分離 |
| 研究線凍結 + Phase 3 定位 | **Line-level freeze**（第九輪 review，2026-07-09）：研究線收官；**Phase 3 = 驗證穩定性非研究抽象**——量尺為「新案例是否不改模型自然落位」（predictive vs retrospective power）。Post-Phase-3 候選主題：**Evidence Lifecycle**（登記不展開） |
| 驗證 leg | 復用 review capability `fault_finding` stance invoke，不另定 stance |
| 適用範圍 | **Execute 意圖（「開始執行 plan / sub-plan / slice」）= mandatory 三角色 loop**（2j F2 stakeholder 裁決，2026-07-10）——`delegation.enabled: false` **不是豁免**：Execute 前 orchestrator 須翻 `true` + 補 backfill，或 plan 明記 transport adaptation；advisory 僅限非 Execute 情境（純問答 / 只讀審計 / surgical 小修）。主打 software-delivery，Ai-skill 比照。**APK Analysis Discovery 不在此 mandatory 範圍**（無三角色證據；見 Domain Boundary 列） |
| Domain Boundary（APK ↔ SD） | **Capability Handoff**（2026-07-14）：觸發 = deliverable capability，非 analysis-complete；輸入 = Capability Proposal；SD 不消費 RE 機械細節。apk-analysis = candidate Discovery；software-delivery = validated Delegated Execution。細節 [`04`](04-apk-capability-handoff-boundary.md)；dogfood **2t** |
| V3 evidence producer | targeted mutation（risk-triggered：boundary / boolean / null / authorization / invariant / guard）；survived mutant 須轉 semantic-gap finding，**不做 mutation score KPI**；producer family 通用化（Behavioral Falsification）gated on Q9，graduate 前 mutation-only（Mutation review 裁決，2026-07-09） |
| Schema promotion | gated on Phase 2 證據（Q5），deadline 2026-08-31 |
| Dogfood transport | 雙 transport：2a Cursor（human 路徑，使用者操作）/ 2b Claude Code Agent（agent 路徑，orchestrator 自駕，2026-07-08 使用者授權）；模板 tool-neutral，工具細節只在 kit 傳輸備註 + `ai-tools/agent/`（Layer 3） |
| 2b 委派任務 | Phase 1 的 plans/README.md §Delegation loop SOP 擴充本身（真實待辦、可驗收、orchestrator 全程不碰實作） |

## 與其他 plans 的關係

- [`2026-06-22-1009-plans-system-portability-and-delivery-integration/03-subplan-agent-delegation.md`](../../archived/2026-06-22-1009-plans-system-portability-and-delivery-integration/03-subplan-agent-delegation.md) — **baseline**（completed）：delegation brief schema + 雙路徑 dogfood。本 plan 延伸其 loop（去程 → 回程），不重開該 sub-plan；自動 orchestrator reservation 邊界維持。
- [`2026-06-22-1009-plans-system-portability-and-delivery-integration/02-software-delivery-plan-first-ordering.md`](../../archived/2026-06-22-1009-plans-system-portability-and-delivery-integration/02-software-delivery-plan-first-ordering.md) — plan-first ordering 是本 loop 的前置（orchestrator 產 plan artifact 先於執行）；本 plan 不改其 Q4 關閉條件。
- [`archived/2026-07-06-review-architecture-adr/_plan.md`](../../archived/2026-07-06-review-architecture-adr/_plan.md) — review = cross-cutting capability invoke（ADR-013 D2）；verifier leg 是其消費者。
- [`active/2026-06-16-1131-evidence-candidate-system.md`](../2026-06-16-1131-evidence-candidate-system.md) — `defer` 處置的 findings 可轉 evidence candidate（人工 capture，不新增 scanner 職責）。
- [`04-apk-capability-handoff-boundary.md`](04-apk-capability-handoff-boundary.md) — **本 plan companion**（2026-07-14）：APK ↔ SD Domain Boundary + Capability Handoff 契約候選 + dogfood **2t** 預註冊；不替代 `workflow/apk-analysis` 正文。
- [`archived/2026-05-11-1129-apk-analysis-pilot-migration.md`](../../archived/2026-05-11-1129-apk-analysis-pilot-migration.md) — APK workflow/analysis/intelligence 分層 pilot（completed）；本輪 Boundary 建立在該分層之上，不重開 migration。
