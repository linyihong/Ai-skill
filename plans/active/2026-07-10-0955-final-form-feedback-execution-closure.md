---
id: 2026-07-10-0955-final-form-feedback-execution-closure
plan_kind: main
status: draft
owner: linyihong
created: 2026-07-10
last_updated: 2026-07-10
parent: null
revision:
  - date: 2026-07-10
    note: "全系統架構審計入 plan：層級地圖 + P1–P12 問題清單（314 orphans、scenario 無 runner、.git 6.9GB、lesson 75% 滯留、behavioral 計數器缺失等）；新增 Workstream E + Phase 7–9；closure 改 Phase 10；Q7–Q10"
---

# Final Form — Feedback Execution Closure & Model-Neutral Stability（最終形態：回饋執行閉環與模型中立穩定輸出）

**Status**: `draft`
**Owner**: linyihong
**建立日期**: 2026-07-10
**Source**: 2026-07-10 對話 — 使用者要求「系統最終形態」計畫。核心理想：**不是系統不能犯錯，而是任何 model（不限單一 agent）bootstrap 後都繼承同一套思維，持續進化與提煉**。核心痛點：**每次明明看到可以 feedback 的點，卻都沒有執行**——不管是在修改本 repo，還是在其他 repo 使用本系統時。
**Glossary Impact**: yes — candidate terms：`feedback_execution_closure`（feedback 從「被看到」到「被執行或被顯式仲裁」的閉環性質）、`deferred_feedback_ledger`（DEFERRED / UNAVAILABLE writeback 的 durable queue）、`model_neutral_cognition_parity`（思維載體全部落在 repo、任何 agent 工具讀得到的性質）。graduate 時才註冊到 `knowledge/glossary/ai-skill.md`；未定稿前不註冊。

> **Watch-Out List citation**：對應 [`architecture/ai-native-cognitive-ecosystem-system.md`](../../architecture/ai-native-cognitive-ecosystem-system.md) §Watch-Out List 的「process bloat」「premature abstraction」「autonomous optimizer」防呆：
> - **不建 feedback telemetry DB、不建自動 promotion engine** — ledger 是人工 accept/close 的索引層，復用 evidence-candidate 已驗證的 observation-infra 模式，不長第二條 authority。
> - **機械化 gate 一次一階** — 先 doc-only ledger + 手動 round-trip，falsification 證據成立才升 stop-hook 機械 gate（governance promotion falsification ladder 紀律）。
> - **不因「更好的想法」改契約** — 本 plan 的 working model 凍結於 §Decision；只有 dogfood 證據能改它。

## 北極星：最終形態的定義（What "final form" means）

系統的最終形態**不是**「不會犯錯的系統」，而是滿足以下三條性質的認知作業系統：

| # | 性質 | 可觀察判準 |
|---|---|---|
| 1 | **Feedback Execution Closure** | 任何一次「值得沉澱的 learning」要嘛當輪 writeback COMPLETED，要嘛進入 durable ledger 並在後續 session 被 resurface，直到 closed 或被顯式仲裁（won't-fix / refuted）。**靜默流失率 = 0** 是結構性質，不是自律性質。 |
| 2 | **Model-Neutral Cognition Parity** | 思維（evidence≠decision、falsification ladder、observation-first、minimal governance surface、surgical discipline）的 canonical 載體 100% 在 repo（governance / enforcement / workflow / runtime.db），**0% 依賴單一工具的 private memory**。換 model、換工具，bootstrap 後行為分佈不變。 |
| 3 | **Mechanical-First Stability** | 高頻失效模式的 enforcement 由 behavioral 逐階升 mechanical；「超級穩定輸出」= 輸出品質由契約與 gate 決定，不由模型個性決定。promotion 順序由 ledger 的失效頻率資料驅動，不由直覺驅動。 |

三條性質互相餵養：(1) 產生失效頻率資料 → (3) 用資料決定 promotion 優先序 → (2) 保證任何 model 都被同一套 gate 約束 → 回到 (1) 讓任何 model 的 feedback 都進同一個 ledger。這就是「按照同一套思維持續進化與提煉」的機械化表述。

## 全系統架構審計（2026-07-10 盤點）

> 使用者要求：不只解使用者講的痛點，把整個系統攤開解析、找出存在的問題點。以下每條都附可覆核證據；量測值以 2026-07-10 為準。

### 層級地圖（現況）

| 群組 | 層 | 規模 | 健康度一句話 |
|---|---|---|---|
| 入口/契約 | `CLAUDE.md` / `AGENTS.md` / `GEMINI.md` / `.cursor` → `CORE_BOOTSTRAP.md` ↔ `runtime/core-bootstrap.yaml` → `runtime.db` | hooks 全鏈機械強制 | 最成熟的一段；但每 session 固定注入 ~40KB（見 P8） |
| 思維/policy | `enforcement/`(63md) `governance/`(48) `constitution/`(15 ADR) `models/`(31) `anti-patterns/`(6) | 41 rule classes | mechanical 51%；behavioral 13 條的升級條件沒人在數（P6） |
| 執行 | `workflow/`(86md) `plans/`(52 archived + 9 active) | delivery domain 獨大 | plan 治理最強（19+ validators）；並行回寫曾真實相撞（P11） |
| 知識 | `knowledge/`(37md, glossary 1355 行) `intelligence/`(226md) `analysis/`(27) `memory/`(18) `metadata/`(19) | 962 md 全庫 | 寫入多、讀出少：routes 59 條中 43 orphan（P2） |
| 回饋 | `feedback/`(199 lessons) `validation/`(225 scenarios) `evaluations/`(1md) `traces/`(1md) | — | **全系統最弱的一環**：lesson 75% 滯留（P5）、scenario 無 runner（P1）、evaluations/traces 是空殼層（P12） |
| 工具 | `scripts/ai-skill-cli/`（Go；hooks.go 4408 行）+ 6 平台 committed bin（各 ~17MB） | CI 1 條 | bin 歷史 179 次 commit → `.git` 6.9GB（P3）；validator 單檔巨石（P10） |

### 問題點清單（P1–P12，依「對穩定輸出的威脅」排序）

| # | 問題 | 證據 | 為什麼威脅「超級穩定輸出」 |
|---|---|---|---|
| **P1** | **驗證層是自申報的**：225 個 validation scenario YAML **沒有 runner**；CLI 只有 filename heuristic（`enforcement.go` walk），無任何路徑執行 scenario 斷言。plan 裡的 `scenario X PASS` 全是 agent 自我宣稱 | `ai-skill` 無 scenario 子命令；grep 僅 filename 引用 | 系統穩定性敘事建立在 gates 上，但 gate 的 scenario 層是 behavioral——換一個不自律的 model，PASS 可以是幻覺。這是 inflated-reporting 的結構性溫床 |
| **P2** | **314 orphan 宣告面**：routes 43/59、surfaces 66/77、scenarios 205/220 無 consumer | `ai-skill runtime audit` | 宣告遠超消費 = 假治理感 + context 稅 + drift 溫床。Gen3 audit 只保護**新增** wiring，存量沒有清償計畫 |
| **P3** | **repo 健康：`.git` 6.9GB**：6 平台 bin（各 16–18MB）已 commit 179 次 | `du .git`；`git log --follow bin/` | clone/CI/工具操作全面變慢；Go-first policy 的未付帳單。任何 history 清理都是 destructive，拖越久越貴 |
| **P4** | Learning report 只有申報層無執行層（使用者原始痛點） | `hooks.go` 只驗 enum | 已由本 plan Workstream A 承接 |
| **P5** | **feedback/history 199 條 lesson，僅 ~49 條有 promotion 痕跡（75% 滯留）** | grep promotion feedback/history | promotion pipeline 文件齊全但無 pull 機制——與 P4 同構：回饋層「進得來、出不去」 |
| **P6** | **behavioral 升級條件沒有計數器**：13 條 behavioral 的 sunset review 條件全是「≥N incidents accumulate」，但系統沒有任何機制在數 incident | `ai-skill enforcement coverage` 輸出 | 升級條件永遠不會被觸發 = ladder 永遠停在 behavioral。Workstream A 的 ledger 應兼任此 counter（synergy） |
| **P7** | **grandfather sunset 2026-08-31 逼近**：3-4 個 doc-only completed plans 需在 deadline 前升 auto-detected 或降 orphan，目前無排程 | `plans/README.md` ⚠️ 標記 | deadline 違約會讓「completed」語義失真 |
| **P8** | **Bootstrap 稅無分級**：每 session 固定注入 ~40KB + 必讀 3 檔 + receipt 查詢，trivial/read-only 任務同價 | SessionStart hook 輸出 40.7KB | 成本高 → 使用者/agent 有繞過誘因 → 繞過即失去全部 gate 保護。分級（light bootstrap）不存在 |
| **P9** | **知識讀出率低**：intelligence/ 226 md 但對應 routes 多 orphan；`knowledge/summaries/` 僅 26 份，覆蓋率低 | 層級地圖 + audit | 知識寫入不被消費 = 進化的「提煉」半途而廢；也加重 P2 |
| **P10** | **hooks.go 4408 行單檔巨石**：28 個 commit validators 硬編碼 dispatch；dispatcher refactor 自 Phase 7 起持續 deferred | `wc -l`；`per_commit_dispatcher_status.go_dispatch_refactor: deferred` | 所有機械強制的單點；新 validator 邊際成本遞增，review 難度上升 |
| **P11** | **並行回寫無鎖已真實相撞**：2026-07-10 delegation plan 兩個併發 session 回寫，舊底稿覆蓋掉 5 輪內容，靠 git 考古重建 | delegation plan frontmatter revision note | `.agent-goals/` lock 只管 project-local，Ai-skill plan 檔案本身無 single-writer 保護；多 agent 化（delegation loop）會放大此風險 |
| **P12** | **空殼層 / layer sprawl**：21 個頂層目錄中 `traces/` `evaluations/` `templates/` `tools/` 各只有 1–4 md；`evaluations/`（scenario 結果）從未運轉 | 層級地圖 | 宣告了層但沒長肉 = 導航成本 + 「系統很完整」的錯覺；應併層、填實或明文 reservation |

### 審計結論（一段話）

系統的**去程**（契約 → gate → 執行）已經相當成熟且機械化；**回程**（執行 → 證據 → 回饋 → 提煉 → 升級）整段都是 behavioral 或斷裂的：scenario 自申報（P1）、lesson 滯留（P5）、incident 沒人數（P6）、learning report 無執行鏈（P4）。使用者感受到的「看到 feedback 但沒執行」是整個回程斷裂的其中一個症狀。因此本 plan 的最終形態工作 = **把回程接通並機械化**，加上清償三筆已量化的結構債（orphan 存量 P2、git 肥大 P3、驗證自申報 P1）。

## Decision Rationale

### Problem & Why Now

**斷鏈診斷（2026-07-10 盤點）**：

1. **Feedback 只有申報、沒有執行鏈**。[`obligation.feedback.learning_report`](../../runtime/core-bootstrap.yaml) 已上線（2026-06-08 plan completed），但 stop hook 只驗**格式**（`scripts/ai-skill-cli/internal/app/hooks.go` 只檢查 enum 值合法）。`Writeback: DEFERRED` 與 `UNAVAILABLE` 是合法**終態**：沒有 queue、沒有下次 session resurface、沒有 closure 指標。「看到 feedback」與「執行 feedback」之間沒有任何結構連接——這正是使用者痛點的機械成因。
2. **Cross-repo feedback 死在對話裡**。在 consumer repo 使用本系統時，`RepoContext: NON_LOCAL` + `Writeback: UNAVAILABLE` 合法地結束一輪；learning 內容留在該對話 transcript，永不回流 Ai-skill。
3. **思維載體部分是 tool-private**。相當一部分操作紀律（falsification ladder、evidence-decision separation、observation status discipline、rollout boundary…）目前活在單一 agent 工具的 private memory；其他 model / 工具 bootstrap 後**讀不到**。這直接違反「不管其他模型都能按照同一思維」的理想。
4. **Enforcement ladder 未完成**。目前 41 rule classes：mechanical=21、behavioral=13、not_mech=5、research=2。behavioral 規則對不同 model 的遵循度天然不穩定，是輸出不穩定的最大殘餘來源。

以上 1–4 是使用者痛點的直接成因。**全系統審計（見 §全系統架構審計）進一步發現 8 個使用者未提到的結構問題（P1–P3、P5–P9 尤其），共同模式：系統的「回程」（執行→證據→回饋→提煉→升級）整段 behavioral 或斷裂**。

Why now：learning report obligation、evidence-candidate observation infra（Phase 1 completed）、delegation loop（三角色閉環 dogfood 中）三塊拼圖都已就位——feedback 執行閉環可以**復用**它們，不需要新機制；再晚，deferred feedback 的流失持續累積且不可回溯。

### Decision

四個 workstream，漸進落地，全部復用既有基礎設施：

**A. Feedback Execution Closure（核心，Phase 1–4）**
建 `feedback/pipeline/deferred/` deferred-feedback ledger（committed、去敏後入庫）：

- 行為規則：`FeedbackDecision: NEEDED` 且 `Writeback: DEFERRED|UNAVAILABLE` 時，**必須**附 durable pointer（ledger entry 路徑）；沒有 pointer 的 DEFERRED 視為未完成 close-out。
- Lifecycle 沿用 evidence-candidate 拍板慣例：`create → close(applied) | reject(refuted) | expire`（expire ≠ reject，防 closure-rate 失真）；entry 不可指向 entry。
- Resurface：`ai-skill runtime receipt` 輸出加一行 open-deferred 計數（不加 bootstrap 檔案、不擴 SessionStart 注入量，避開 bootstrap-entry-bloat failure pattern）。
- Cross-repo：consumer repo session 結束時，NEEDED 的 learning 以結構化 block 寫進 final response + 寫入 `<AI_SKILL_REPO>` ledger（機器上路徑可達時）；不可達時 fallback = final response block，由使用者/下個 Ai-skill session 搬運（此 fallback 的實際流失率是 Phase 3 量測欄）。
- 機械化 gate **gated on 證據**（Phase 4 entry condition，見下）。

**B. Model-Neutral Cognition Parity（Phase 5）**
盤點 tool-private memory ↔ repo 的 parity gap：每條 memory 分類為 `user-specific`（留 memory，不進 repo）/ `discipline`（可般化 → 移植成 tool-neutral governance / enforcement / intelligence 文件，memory 降級為 pointer）/ `already-covered`（repo 已有 canonical，刪重複）。之後用**非 Claude 的 model** 跑同一組任務 dogfood，量測 obligation 遵循度與思維紀律遵循度（量測欄見 Phase 5）。

**C. Mechanical Enforcement Ladder Completion（Phase 6，ongoing）**
以 ledger 的失效頻率為排序輸入，對 13 條 behavioral rule classes 逐條評估 behavioral → mechanical promotion；一次一階、每階可回退（enforcement-registry Status Transition Matrix 既有紀律，不新增流程）。

**D. Roadmap Consolidation（§未來 plans 排程，非 phase）**
既有 active plans 的排序與 entry conditions 收斂成一張表（reference-link，不重寫、不 re-parent——各 plan 維持獨立 main plan）。

**E. Verification & Structural Debt（Phase 7–9，源自 §全系統架構審計）**
清償三筆已量化的結構債 + 排程既有 deadline：

- **E1 Scenario Runner（P1）**：讓 `scenario X PASS` 從 agent 自申報變成機器可執行——先盤 225 個 scenario 的 assertion 形態，選出機器可跑子集（例如 file-exists / grep-able / CLI-invocable 斷言），落 `ai-skill scenarios run`；不可機器化的 scenario 明文標 `agent-judged` 並留 review 抽查。
- **E2 Orphan 存量清償（P2）**：314 orphans 逐批處置——wire consumer / 標 `manual_activation`（附 reason）/ deprecate 下架；每批走既有 audit 工具驗收，orphan 數只降不升（ratchet）。
- **E3 Repo 健康（P3）**：committed bin 策略改造（Git LFS / release artifacts / 單平台 bin + CI 分發擇一），並向使用者提案 history 清理（**destructive，僅提案不執行**）。
- **E4 既有 deadline 排程（P7）**：grandfather doc-only plans 在 2026-08-31 前逐個升 auto-detected 或降 orphan。
- **E5 併行回寫保護（P11）**：plan 檔案 single-writer 慣例（延伸 `.agent-goals/` lock 語義到 Ai-skill plan 回寫，或 commit-time 衝突偵測）——先 doc-only 慣例 + 觀察，機械化 gated。

P8（bootstrap 分級）、P9（知識讀出率）、P10（hooks.go 巨石）、P12（空殼層）**本 plan 只記錄不執行**：P8/P9 進 Open Questions（Q9/Q10）等證據；P10 已有 deferred dispatcher refactor 決議（`per_commit_dispatcher_status`）不重複開工；P12 留待 orphan 清償（E2）時一併盤層。

### Alternatives Considered

- **A. 維持現狀（learning report 只報不追）**：reject — 正是本 plan 要解的 pain point；report 已證明「申報格式」擋不住流失。
- **B. 立即機械化 hard block（DEFERRED 一律擋 stop）**：reject — 未先驗證 ledger 可用性就上 hard gate，會誘發假 `COMPLETED`（inflated-reporting failure pattern 已記錄此風險）；先 doc-only + 手動 round-trip。
- **C. 建全新 feedback telemetry / 自動 promotion engine**：reject — Gen4 watch-out 明確反對 autonomous optimizer；evidence-candidate 模式已證明「被動 detection + 人工 capture」夠用且不長第二條 authority。
- **D. 把思維移植做成一次性大遷移**：reject — memory 條目需逐條分類（user-specific vs discipline），bulk-apply 違反 rollout boundary 紀律；逐條、可回退。
- **E. 漸進：ledger + resurface + 量測 → 證據成立才機械化（accept）**。

### Why Not an ADR Yet

未驗證：ledger 的 closure ratio、cross-repo fallback 流失率、非 Claude model 的 parity 量測都還沒有第一筆證據；scope 會依 Phase 1–3 dogfood 調整；Q1–Q5 未解。等 Phase 4 機械化 gate graduate 且 Phase 5 parity dogfood 有結果，才評估是否有 ADR-worthy 的 durable decision（候選：「feedback execution closure 是 per-turn obligation 的必要延伸」）。

### ADR Promotion Criteria（completed 時驗證）

- [ ] foundational + cross-session + cross-project + expensive-to-reverse + explains-why 全中
- [ ] Plan 結果證實 decision 可行（closure ratio 與 parity 量測達 §完成條件）
- [ ] Open Questions 全解
- [ ] 沒有更輕的 promotion target 適用（per ADR-007）
- [ ] 系統真實使用此 contract（≥ 30 天 ledger 運轉證據 + ≥ 1 次非 Claude model dogfood）

### Consequences（預期）

#### 正面
- Feedback 靜默流失從「常態」變成「結構上不可能不留痕」。
- 思維載體 repo 化後，model 更換 / 多工具並用不再稀釋系統行為。
- Enforcement promotion 有了資料驅動的優先序來源。

#### 負面
- 每輪 close-out 多一個 ledger 寫入動作（NEEDED+DEFERRED 時）；token 與操作成本上升。
- ledger 需要去敏紀律（cross-repo feedback 易夾帶專案細節）。

#### 風險
- **假 COMPLETED 規避**：機械化後 agent 可能謊報 COMPLETED 繞過 ledger。緩解：Phase 4 抽查 COMPLETED 的 commit 證據（writeback 必留 git 痕跡，可機械 cross-check）。
- **ledger 變垃圾場**：低價值 entry 淹沒高價值。緩解：expire lifecycle + Phase 2 量測欄（accepted_ratio）；evidence-candidate 已驗證此配方。
- **parity 移植過度**：把 user-specific 偏好誤植為系統紀律。緩解：Phase 5 逐條分類需 stakeholder sign-off。

## Runtime Execution Path

**First landing 是 doc-only + CLI consumer，不接 runtime**（Deferred Runtime Projection 宣告）：

- **不 project 的 reason**：ledger schema 未經 dogfood，closure lifecycle 的 enum 可能依 Phase 1–3 證據調整；先外放 `feedback/pipeline/deferred/`（markdown + README index），不建 `runtime/*.yaml`、不 project `runtime.db`。
- **預定 project 的 phase / 條件**：Phase 4 機械化時，若 stop-hook gate 需要 machine-readable contract，才建 `runtime/feedback-deferred-ledger.yaml`（`runtime_projection.enabled: true`）並同批 wire named consumer。**graduation deadline：2026-09-30**；逾期未 graduate 則本 plan 的機械化 phase 降級為 explicit-defer 並記錄原因。
- **Trigger flow（Phase 4 目標形態，先寫明避免「routing 會處理」空話）**：event = stop-hook 收到 final response → detector = learning report parser（現有 `hooks.go` feedback 驗證函式擴充）→ query = `FeedbackDecision==NEEDED && Writeback∈{DEFERRED,UNAVAILABLE}` → loaded contract = ledger 路徑 convention → runtime action = 無 durable pointer 時 block + 提示 ledger 寫入指令 → evidence = hook 輸出 + ledger entry。
- Phase 1–3 期間唯一的 runtime 接觸點：`ai-skill runtime receipt` 加一行 open-deferred 計數（讀 `feedback/pipeline/deferred/` 檔案系統，不進 DB；Go-first，改 `scripts/ai-skill-cli/`，rebuild committed bin）。

### Per-surface consumer 表（Phase 4 生效時填實；現為預告）

| Generated surface key | Named consumer(s) | Consumer 類型 |
|---|---|---|
| `gate.feedback.deferred_pointer_required`（Phase 4 才建） | stop-hook learning report validator（`hooks.go`） | Go validator |
| `feedback/pipeline/deferred/` README index（Phase 1） | `ai-skill runtime receipt` open-deferred 行 + 人工 review | CLI / manual_activation（reason: observation-period infra，等 Phase 4 wire gate） |

## Open Questions

- [ ] **Q1 — ledger 位置與提交策略**：committed `feedback/pipeline/deferred/`（跨 session / 跨機器可見，需去敏 gate）vs gitignored inbox（evidence-candidate 慣例，但 deferred feedback 的核心需求正是跨 session 存活）。傾向 committed；Phase 0 確認與 evidence-candidate 的 inbox-gitignored 拍板不衝突（兩者用途不同：candidate 是「未接受的觀察」，deferred entry 是「已判定 NEEDED 的執行債」）。
- [ ] **Q2 — cross-repo 寫入可達性**：consumer repo session 對 `<AI_SKILL_REPO>` 的寫入在各工具（Claude Code / Cursor / 其他）的權限模型下是否穩定可達？不可達時 fallback block 的實際搬運率多少？（Phase 3 量測）
- [ ] **Q3 — memory→repo parity 邊界**：哪些 tool-private memory 是 user-specific（不進 repo）？分類判準草案：涉及「使用者授權/偏好」→ user-specific；涉及「證據與決策怎麼處理」→ discipline。需 Phase 5 逐條核對 + sign-off。
- [ ] **Q4 — 非 Claude model dogfood 的量測欄**：obligation 遵循度（receipt / mode report / learning report 出現率）之外，思維紀律遵循度怎麼量測？候選：violation-per-session 計數（surgical violation、evidence-decision leakage、premature promotion）。
- [ ] **Q5 — resurface 注入點的噪音成本**：receipt 加一行是否足夠讓 deferred 被撿起？若 30 天 closure ratio 過低，是否需要升級為 SessionStart 注入（有 bootstrap-entry-bloat 風險）？
- [ ] **Q6 — 假 COMPLETED 的機械 cross-check**：`Writeback: COMPLETED` 宣告與當輪 git 痕跡（feedback-history / enforcement / workflow 路徑的 diff）能否機械對帳？
- [ ] **Q7 — scenario 機器化比例**：225 個 scenario 中多少比例的斷言可機器執行？低於多少比例時 E1 的 runner 投資不划算（改用抽查制）？（Phase 7 盤點回答）
- [ ] **Q8 — bin 儲存策略**：LFS（需 remote 支援）vs release artifacts（CI 改造）vs 單平台 bin（跨平台 policy 弱化）？history 清理（filter-repo）使用者是否授權？
- [ ] **Q9 — bootstrap 分級（P8）**：trivial/read-only 任務的 light bootstrap 是否值得？風險 = 分級判定本身可被濫用為繞過。需要先有「bootstrap 稅 vs 繞過率」的觀察資料，不先動 P0 gate。
- [ ] **Q10 — 知識讀出率量測（P9）**：intelligence/ 的實際被讀率怎麼量測（route hit? summary 載入率?）？低讀出的 atom 應 deprecate 還是改 routing？

## 未來 plans 排程（Roadmap Consolidation — reference-link，不 re-parent）

既有 active plans 與本 plan 的關係與建議順序。各 plan 維持獨立 main plan；本表只收斂 sequencing 與 entry conditions，不改各 plan 內容：

| 順位 | Plan | 現狀 | Entry condition / 與本 plan 的關係 |
|---|---|---|---|
| 1 | 本 plan Phase 1–3（ledger + resurface + cross-repo） | draft | 立即可做；不依賴其他 plan |
| 2 | [`2026-05-28-1830-plan-archival-audit-validator`](2026-05-28-1830-plan-archival-audit-validator.md) | draft | 獨立 quick win；建議先 graduate 保護後續 archive commits |
| 3 | [`2026-07-08-0825-delegation-verification-arbitration-loop`](2026-07-08-0825-delegation-verification-arbitration-loop/_plan.md) | in-progress | 續跑 dogfood；其 verifier findings 是本 plan ledger 的天然 feed 來源之一 |
| 4 | [`2026-06-16-1131-evidence-candidate-system`](2026-06-16-1131-evidence-candidate-system.md) Phase 2 | in-progress（observation period） | 等 `phase2_gate`（count≥20 / reviewed≥80% / accepted>50% / age_p95<30d）；本 plan ledger 與 candidate inbox 保持分工（執行債 vs 觀察） |
| 5 | 本 plan Phase 4（機械化 gate）+ Phase 6（enforcement ladder sweep） | — | entry: ledger 運轉 ≥2 週、≥10 entries、closure 流程走通 |
| 6 | [`2026-05-27-1557-tool-runtime-signal-economics-integration`](2026-05-27-1557-tool-runtime-signal-economics-integration.md) | draft | 原 sequencing 條件已滿足（validateRuntimeTriggerWiring active）；排在機械化 gate 之後避免同期改 hooks.go 疊風險 |
| 7 | 本 plan Phase 5（model-neutral parity + 非 Claude dogfood） | — | 可與 6 並行；需要使用者提供第二 model 環境 |
| 8 | [`2026-06-06-1700-workflow-activation-discovery-bridge`](2026-06-06-1700-workflow-activation-discovery-bridge.md) / [`2026-06-08-2100-governance-pattern-library-extraction`](2026-06-08-2100-governance-pattern-library-extraction.md) / [`2026-06-29-1430-preparatory-refactoring-workflow`](2026-06-29-1430-preparatory-refactoring-workflow/_plan.md) | draft/in-progress | 按各自 plan 的 gate；governance-pattern extraction 遵守 N≥5 紀律 |
| 9 | [`2026-05-28-1636-gen4-fitness-optimization-memory-interface-reservation`](2026-05-28-1636-gen4-fitness-optimization-memory-interface-reservation.md) | draft | 最後：其 fitness 輸入正是本 plan ledger + economics plan 的 telemetry primitives |
| 10 | [`2026-06-16-1030-interaction-hazard-review-workflow`](2026-06-16-1030-interaction-hazard-review-workflow.md) | draft | 按其 A0→D roadmap 自走；與本 plan 無依賴 |
| — | 本 plan Phase 8（grandfather 結案部分） | — | **hard deadline 2026-08-31**：不論其他順位，需插隊在 deadline 前完成 |
| — | 本 plan Phase 7 / 9（scenario runner、repo 健康） | — | 可與順位 3–6 並行；Phase 9 的 history 清理需使用者授權（Q8） |

## Phase 0 — 盤點與 Preflight

### Phase 0.0 — Open Questions 核對（公版，必填）

逐條核對本 plan §Open Questions，標記處置並回寫：

- [ ] 已讀本 plan §Open Questions 全部條目
- [ ] 對每條標記 `resolved`（附 Phase 0 證據）/ `still-open` / `deferred`（附原因）
- [ ] `resolved` 的條目已同步勾選 / 附註於 §Open Questions
- [ ] 若盤點新發現問題，已加入 §Open Questions

| Open Question | 處置 | 證據 / 原因 |
|---|---|---|
| Q1 ledger 位置 | pending | |
| Q2 cross-repo 可達性 | pending | |
| Q3 parity 邊界 | pending | |
| Q4 dogfood 量測欄 | pending | |
| Q5 resurface 噪音 | pending | |
| Q6 假 COMPLETED cross-check | pending | |

### Phase 0.1 — Architecture Compatibility Preflight + Pre-build Interrogation

- [ ] 讀 `feedback/pipeline/` 現有 YAML（lifecycle-automation / promotion-engine / promotion-workflow），確認 deferred ledger 與 promotion pipeline 的分工邊界（ledger 管「執行債被撿起」，promotion 管「撿起後往哪放」），不建重複 surface
- [ ] 讀 `governance/evidence-candidates/` schema 與拍板紀錄，確認可復用的 lifecycle / README-index 慣例與**不可**復用的部分（inbox gitignored）
- [ ] 讀 `hooks.go` learning report 驗證函式現況，確認 Phase 4 擴充點
- [ ] 確認 `archived/2026-06-08-1047-feedback-learning-report-obligation.md` 的 deferred ADR promotion 條件（post-use spam check）與本 plan 是否互鎖
- [ ] 完成最低記錄格式（Trigger / Checked sources / Conflicts / Interrogation / Open Questions 核對 / Decision / Validation）

## Phase 1 — Deferred Feedback Ledger（doc-only）

- [ ] `feedback/pipeline/deferred/README.md`：entry schema（id / created / source_repo_context / target enum 沿用 learning report 的 `feedback-history|intelligence|workflow|enforcement|project-docs` / status `open|closed|refuted|expired` / closure evidence 欄）+ 索引表 + 去敏規則（引用 `enforcement/sanitization.md`）
- [ ] 兩條 invariant 寫入 README：entry 不可指向 entry；ledger 是索引層不是 authority（closure 的 authority 在被 writeback 的目標層）
- [ ] 手動 round-trip ×1：拿一條真實 deferred feedback（本 plan 撰寫過程即產生候選）走 create → writeback → close，證明鏈路可走通
- [ ] 完成條件：README + schema + 1 筆 closed entry 入庫

## Phase 2 — Resurface + CLI

- [ ] `ai-skill feedback` subcommand（Go-first：`scripts/ai-skill-cli/` 新 cmd，rebuild committed bin，遵守 pre-push binary guard）：`list --open` / `defer <summary>` / `close <id> --evidence <path>`
- [ ] `ai-skill runtime receipt` 輸出加 open-deferred 一行（filesystem read，不進 runtime.db）
- [ ] 量測欄啟動：entries created / closed / expired、closure ratio、age p95（人工月結即可，不建 telemetry）
- [ ] 完成條件：CLI 三動作可用 + receipt 顯示計數 + 量測欄第一筆

## Phase 3 — Cross-repo Path（dogfood）

- [ ] 在 ≥1 個真實 consumer repo session 走完整流程：NON_LOCAL 發現 learning → 寫入 Ai-skill ledger（或 fallback block）→ 回 Ai-skill session close
- [ ] 量測：寫入可達率、fallback 搬運率、去敏 gate 是否擋住專案細節
- [ ] Q2 回寫
- [ ] 完成條件：≥1 筆 cross-repo entry 完成 create→close 全循環

## Phase 4 — 機械化 Gate（entry condition gated）

**Entry condition**：Phase 1–3 完成 + ledger 運轉 ≥2 週 + ≥10 entries + closure ratio 有第一筆量測。

- [ ] `hooks.go` learning report 驗證擴充：`NEEDED + DEFERRED|UNAVAILABLE` 且無 ledger pointer → block（含修復提示）
- [ ] 假 COMPLETED cross-check（Q6）：可行則同批落地，不可行則記 explicit-defer
- [ ] 若需要 machine-readable contract：建 `runtime/feedback-deferred-ledger.yaml`（`runtime_projection.enabled: true`）+ 填實 §Per-surface consumer 表 + scenario tests
- [ ] enforcement-registry 登記 coverage 變化（behavioral → mechanical；走 Status Transition Matrix）
- [ ] 完成條件：gate live + scenario tests PASS + registry 同步

## Phase 5 — Model-Neutral Cognition Parity

- [ ] 盤點 tool-private memory 全清單 → 逐條分類 `user-specific` / `discipline` / `already-covered`（分類表入本 plan evidence 或 companion）
- [ ] `discipline` 條目移植為 tool-neutral repo 文件（governance / enforcement / intelligence 依內容歸層；每條獨立 commit，可回退）；memory 端降級為 pointer
- [ ] 非 Claude model dogfood ×1：同一任務、fresh session、只憑 repo bootstrap；量測 obligation 遵循度 + Q4 定案的紀律量測欄
- [ ] gap 回寫：dogfood 發現的遵循缺口 → 進 ledger（吃自己的狗糧）
- [ ] 完成條件：分類表 sign-off + ≥1 次非 Claude dogfood 證據入庫

## Phase 6 — Enforcement Ladder Sweep（ongoing）

- [ ] 以 ledger 失效頻率排序 13 條 behavioral rule classes，選 top 1–2 做 behavioral → mechanical promotion 評估（一次一階）
- [ ] 每次 promotion 走 enforcement-registry Status Transition Matrix + scenario tests；不可機械化者明文記 not_mech 理由
- [ ] 完成條件：≥1 條 promotion 完成或明文 not_mech 化；sweep 節奏（月度）寫入 registry companion

## Phase 7 — Scenario Runner（E1，驗證債）

- [ ] 盤點 225 個 scenario 的 assertion 形態分類（機器可執行 / agent-judged），回答 Q7
- [ ] `ai-skill scenarios run` 落地機器可執行子集（Go-first；rebuild bin）；輸出進 `evaluations/`（讓空殼層長肉，部分回應 P12）
- [ ] 不可機器化 scenario 標 `agent-judged` + 抽查制規則寫入 `validation/README.md`
- [ ] plan acceptance 慣例更新：新 plan 的 `scenario X PASS` 需附 runner 輸出或 `agent-judged` 標記
- [ ] 完成條件：runner 可跑 + 首次全量執行報告入 `evaluations/` + Q7 回寫

## Phase 8 — Orphan 存量清償 + Deadline 排程（E2 + E4）

- [ ] 314 orphans 分批處置（每批 ≤ 20，wire / manual_activation+reason / deprecate），`ai-skill runtime audit` 驗收，orphan 數 ratchet 只降不升
- [ ] grandfather doc-only plans（3–4 個）在 2026-08-31 前逐個結案（升 auto-detected 或降 orphan）
- [ ] 空殼層盤點（P12）：traces / evaluations / templates / tools 逐層判定 併層 / 填實 / 明文 reservation
- [ ] 完成條件：orphan 總數 < 100 或全數帶顯式處置標記；grandfather 清零

## Phase 9 — Repo 健康 + 併行回寫保護（E3 + E5）

- [ ] bin 儲存策略提案（Q8 三選一 + trade-off 表）→ **使用者拍板後**才執行；history 清理僅提案（destructive，P0 授權邊界）
- [ ] plan 檔案 single-writer 慣例 doc-only 落地（回寫前宣告 owner / 檢查併發 session；延伸 `.agent-goals/` lock 語義）+ writeback-collision failure pattern 正式化（`enforcement/failure-patterns/`）
- [ ] 完成條件：策略拍板 + 慣例文件入庫 + failure pattern 註冊

## Phase 10 — Plan Completion Closure

- [ ] 執行 [`plans/README.md`](../README.md) §Plan 完成閉環 checklist（validator / linked updates / 狀態表 / archive 評估——本 plan 的 Phase 6 為 ongoing，評估是否適用「持續生效基礎建設」例外留 active）
- [ ] ADR Promotion Criteria 核對

## 完成條件

- [ ] Phase 1–5、7–9 全部完成條件達成；Phase 6 至少一輪 sweep
- [ ] **北極星判準可量測**：closure ratio 有 ≥30 天資料；parity 分類表 100% 處置；靜默流失在機械 gate 下結構性不可能（DEFERRED 必有 pointer）
- [ ] **審計債判準**：scenario PASS 有 runner 輸出或顯式 `agent-judged` 標記；orphan 數 < 100 或全數帶顯式處置；grandfather 清零；bin 策略拍板
- [ ] Open Questions 全解或明文 deferred
- [ ] `git status` clean、全部 push、plans/README.md 狀態表同步

## Stakeholder 同意項目

| # | 項目 | 需要同意的原因 |
|---|---|---|
| 1 | ledger 採 **committed**（非 gitignored）+ 去敏 gate | 改變「feedback 內容入 repo」的邊界（Q1） |
| 2 | Phase 4 機械化 gate 的 block 行為 | 增加每輪 close-out 摩擦；需確認 friction 可接受 |
| 3 | Phase 5 memory 分類表（哪些進 repo、哪些留 user-specific） | 涉及使用者偏好與系統紀律的邊界（Q3）；bulk 移植前需逐條 sign-off |
| 4 | 非 Claude model dogfood 的環境與 model 選擇 | 需要使用者提供 / 授權第二 model 環境 |
| 5 | §未來 plans 排程表的順位 | 改變既有 draft plans 的啟動順序 |
| 6 | bin 儲存策略（Q8）與任何 git history 清理 | **destructive**（P0 授權邊界）：history rewrite 影響所有 clone；未經明確同意絕不執行 |
| 7 | orphan 清償的 deprecate 批次 | 下架 route / surface / scenario 屬於能力移除，需可回退且逐批確認 |

## 與其他 plans 的關係

- [`archived/2026-06-08-1047-feedback-learning-report-obligation.md`](../archived/2026-06-08-1047-feedback-learning-report-obligation.md) — 本 plan 的直接前置：report 申報層已落地，本 plan 補執行層；其 deferred ADR promotion 與本 plan Phase 4 合併評估。
- [`2026-06-16-1131-evidence-candidate-system.md`](2026-06-16-1131-evidence-candidate-system.md) — observation infra 模式來源（lifecycle / README index / 反自我繁殖 invariant）；ledger 與 candidate inbox 分工：執行債 vs 未接受觀察。
- [`2026-07-08-0825-delegation-verification-arbitration-loop`](2026-07-08-0825-delegation-verification-arbitration-loop/_plan.md) — verifier findings 的 `defer` 處置是 ledger 的 feed 來源之一；仲裁協議（fix/defer/reject）與 ledger lifecycle 語義對齊。
- §未來 plans 排程表內各 plan — sequencing 關係如表；不 re-parent、不改內容。
- [`governance/lifecycle/system-upgrade-governance.md`](../../governance/lifecycle/system-upgrade-governance.md) — Phase 4 runtime 接入走其 §define_runtime_trigger_flow 規則。
