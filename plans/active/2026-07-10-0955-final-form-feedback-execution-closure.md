---
id: 2026-07-10-0955-final-form-feedback-execution-closure
plan_kind: main
status: draft
owner: linyihong
created: 2026-07-10
last_updated: 2026-07-10
parent: null
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

## Decision Rationale

### Problem & Why Now

**斷鏈診斷（2026-07-10 盤點）**：

1. **Feedback 只有申報、沒有執行鏈**。[`obligation.feedback.learning_report`](../../runtime/core-bootstrap.yaml) 已上線（2026-06-08 plan completed），但 stop hook 只驗**格式**（`scripts/ai-skill-cli/internal/app/hooks.go` 只檢查 enum 值合法）。`Writeback: DEFERRED` 與 `UNAVAILABLE` 是合法**終態**：沒有 queue、沒有下次 session resurface、沒有 closure 指標。「看到 feedback」與「執行 feedback」之間沒有任何結構連接——這正是使用者痛點的機械成因。
2. **Cross-repo feedback 死在對話裡**。在 consumer repo 使用本系統時，`RepoContext: NON_LOCAL` + `Writeback: UNAVAILABLE` 合法地結束一輪；learning 內容留在該對話 transcript，永不回流 Ai-skill。
3. **思維載體部分是 tool-private**。相當一部分操作紀律（falsification ladder、evidence-decision separation、observation status discipline、rollout boundary…）目前活在單一 agent 工具的 private memory；其他 model / 工具 bootstrap 後**讀不到**。這直接違反「不管其他模型都能按照同一思維」的理想。
4. **Enforcement ladder 未完成**。目前 41 rule classes：mechanical=21、behavioral=13、not_mech=5、research=2。behavioral 規則對不同 model 的遵循度天然不穩定，是輸出不穩定的最大殘餘來源。

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

## Phase 7 — Plan Completion Closure

- [ ] 執行 [`plans/README.md`](../README.md) §Plan 完成閉環 checklist（validator / linked updates / 狀態表 / archive 評估——本 plan 的 Phase 6 為 ongoing，評估是否適用「持續生效基礎建設」例外留 active）
- [ ] ADR Promotion Criteria 核對

## 完成條件

- [ ] Phase 1–5 全部完成條件達成；Phase 6 至少一輪 sweep
- [ ] **北極星判準可量測**：closure ratio 有 ≥30 天資料；parity 分類表 100% 處置；靜默流失在機械 gate 下結構性不可能（DEFERRED 必有 pointer）
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

## 與其他 plans 的關係

- [`archived/2026-06-08-1047-feedback-learning-report-obligation.md`](../archived/2026-06-08-1047-feedback-learning-report-obligation.md) — 本 plan 的直接前置：report 申報層已落地，本 plan 補執行層；其 deferred ADR promotion 與本 plan Phase 4 合併評估。
- [`2026-06-16-1131-evidence-candidate-system.md`](2026-06-16-1131-evidence-candidate-system.md) — observation infra 模式來源（lifecycle / README index / 反自我繁殖 invariant）；ledger 與 candidate inbox 分工：執行債 vs 未接受觀察。
- [`2026-07-08-0825-delegation-verification-arbitration-loop`](2026-07-08-0825-delegation-verification-arbitration-loop/_plan.md) — verifier findings 的 `defer` 處置是 ledger 的 feed 來源之一；仲裁協議（fix/defer/reject）與 ledger lifecycle 語義對齊。
- §未來 plans 排程表內各 plan — sequencing 關係如表；不 re-parent、不改內容。
- [`governance/lifecycle/system-upgrade-governance.md`](../../governance/lifecycle/system-upgrade-governance.md) — Phase 4 runtime 接入走其 §define_runtime_trigger_flow 規則。
