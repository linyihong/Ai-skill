---
id: 2026-07-08-0825-delegation-verification-arbitration-loop
plan_kind: main
status: draft
owner: linyihong
created: 2026-07-08
parent: null
baseline_ref: 2026-06-22-1009-subplan-agent-delegation
---

# Delegation Verification & Arbitration Loop（委派執行→獨立驗證→仲裁閉環）

**Status**: `draft`
**Owner**: linyihong
**建立日期**: 2026-07-08
**Source**: 2026-07-08 對話 — 使用者觀察到外部框架的三角色模式：主 session 只做規劃 / 切分 / 仲裁，執行交給獨立 agent session，驗證再交給另一個獨立 session，最後由主 session 仲裁每條驗證發現（要修 / 超出範圍 / 駁回）。目標：補漏「預計與實現的落差」。主要針對 `workflow/software-delivery` 的交付處理；Ai-skill 自身任務比照辦理，觀察品質是否提升。
**Baseline**: [`03-subplan-agent-delegation`](2026-06-22-1009-plans-system-portability-and-delivery-integration/03-subplan-agent-delegation.md)（completed，2026-07-06）— delegation `brief` schema + 雙路徑 dogfood ★★★★☆。本 plan 是其 loop 延伸（情境 C：sibling main plan + baseline_ref，不重開該 tree）。
**Glossary Impact**: yes — candidate terms：`independent_verification`（fresh-context 驗證 leg，非 executor 自驗、非 orchestrator 自 review）、`arbitration`（orchestrator 對 verifier findings 的處置協議：fix / defer / reject）。graduate 時才註冊到 `knowledge/glossary/ai-skill.md`；未定稿前不註冊。

> **Watch-Out List citation**：對應 [`architecture/ai-native-cognitive-ecosystem-system.md`](../../architecture/ai-native-cognitive-ecosystem-system.md) §Watch-Out List 的「process bloat」「premature abstraction」「over-engineering」防呆：
> - **不建自動 orchestrator** — 03 的 reservation 邊界維持不變；本 plan 是**角色協議**（主 session 人工扮演 orchestrator），不是 automation。
> - **不強制所有任務走三角色 loop** — 只適用於已宣告 `delegation.enabled: true` 的 sub-plan / 委派任務，且為 advisory；小修補直接做。
> - **不先動 schema** — Phase 1 為 doc-only 協議；schema promotion 需 Phase 2 dogfood 證據（falsification ladder，一次一階）。

## Decision Rationale

### Problem & Why Now

03 已證明 delegation `brief` 能形成 capability（fresh executor 只憑 brief + `context.required` 完成任務），但 loop 只有**去程**沒有**回程**：

1. **驗證由誰做沒有契約**：目前 brief 的 `verification` 由 executor 自驗，或 orchestrator 自己 review——前者缺獨立性（executor 對自己的產出有確認偏誤），後者讓 orchestrator 被迫載入執行細節，失去仲裁位置。
2. **findings 沒有處置協議**：驗證發現的問題，「要修 / 超出範圍 / 駁回」的決定散落在對話中，沒有結構化紀錄，無法回答「預計與實現的落差有沒有被系統性收斂」。
3. **角色邊界未文件化**：「主 session 只規劃 / 切分 / 仲裁、不執行」這條邊界目前不存在於任何 workflow 文件；沒有邊界就無法觀察越界（orchestrator 忍不住自己動手改 code = 信號遺失）。

Why now：03 剛 completed、dogfood 方法論（brief independence score、fresh-context 紀律）記憶猶新；review capability（`fault_finding` stance，ADR-013/014）已落地可復用；工具面（Agent + worktree + 獨立 session）已驗證可行。

### Decision

以 **doc-only 協議 + 雙 dogfood** 落地三角色 loop，不動 schema、不接 runtime：

- **Orchestrator（主 session）**：規劃、切分 sub-plan、寫 brief、仲裁 findings。**不執行、不自己驗**。
- **Executor（獨立 session / agent，可選 worktree）**：僅憑 brief + `context.required` 完成，交付 diff / artifact。（= 03 已驗證的去程，不變）
- **Verifier（另一個 fresh-context session / agent）**：輸入 = 同一份 `brief`（`acceptance` + `verification` 為量尺）+ executor 的 diff / artifact；輸出 = 結構化 findings。**只產證據，不做決定**（evidence ≠ decision）。驗證 leg 復用 review capability 的 `fault_finding` stance invoke，不另定 stance。

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

**Role boundary invariants（3 條，行為層）**：
1. Verifier 必須是 fresh context——不是 executor 的 session、也不是 orchestrator 自己。
2. Orchestrator 在 loop 內不產生 implementation diff；發現自己動手 = 越界信號，記入 dogfood evidence。
3. Loop 關閉條件：所有 findings 都有仲裁處置，且 `fix` 項全部重驗通過。

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

## Runtime Execution Path

**doc-only trial 宣告**：本 plan 不接入 runtime——不新增 `route.*`、不新增 commit-msg validator、不動 `runtime.db` generated surfaces、不動 delegation schema / `validatePlanTreeFrontmatter`。協議以文件 + 行為紀律承載；驗證 leg 復用既有 review capability invoke（`ai-skill runtime capability-invoke --capability code-review --stance fault_finding`，既有 warning-only surface，無新 wiring）。

**未來接入條件（graduation）**：Phase 3 證據評估時決策——若 (a) 三角色 loop 在 ≥2 個真實任務有效、且 (b) role boundary invariant 出現行為維持不住的證據（如 verifier fresh-context 被反覆略過），才評估 schema 欄位（`delegation.verification`）或機械檢查；由後續 plan 承載，本 plan 不 carry。**決策 deadline：2026-08-31**（與本 plan closure 同批；未達證據門檻則明確記錄「維持 doc-only」）。

## Open Questions

| ID | Question | Owner / Resolved By | Status | Closed Criteria | Resolution Evidence |
|----|----------|---------------------|--------|-----------------|---------------------|
| Q1 | Verifier 報告最小欄位集（evidence / acceptance_ref / classification / status）是否足以讓 orchestrator 仲裁而不回讀 diff 細節？ | Phase 1 定稿、Phase 2 驗證 | open | 雙 dogfood 中 orchestrator 均未被迫回讀 diff（或缺欄位已補進契約） | <dogfood evidence> |
| Q2 | 協議文件落點：plans/README.md §Delegation 擴充（delegation 擁有 loop）vs `workflow/cross-cutting/review/` consumer doc（review 擁有驗證 leg）？stance 復用不得重定義 | Phase 1 | open | 落點決定 + 文件落地，且未在 consumer 層重定義 stance / requires_context | <commit> |
| Q3 | 品質信號怎麼量：verifier 差集 findings 數 + 仲裁分佈（fix/defer/reject 比例）是否構成「品質提升」的有效指標？null result 如何記錄？ | Phase 2 | open | 雙 dogfood 各留一份差集 + 分佈紀錄；null result 亦為有效關閉 | <dogfood evidence> |
| Q4 | 仲裁紀錄落點：被委派 sub-plan 內 table（傾向）vs 獨立 artifact？ | Phase 1 | open | 落點決定並在 dogfood 實際使用 | <commit + dogfood> |
| Q5 | Schema promotion 門檻：什麼證據才允許動 delegation schema（如 `delegation.verification`）？ | Phase 3 | open | 門檻明文化；未達門檻則明確記錄維持 doc-only | <Phase 3 決策紀錄> |

## 完成條件

- [ ] Phase 1 協議落地（verifier 契約 + 仲裁協議 + 3 條 role boundary invariants，落點依 Q2）
- [ ] Phase 2 雙 dogfood 完成：software-delivery 外部 repo 真實任務 ×1 + Ai-skill 內部任務 ×1，各留差集 + 仲裁分佈 evidence
- [ ] Phase 3 證據評估：Q5 schema promotion 決策（promote 或明確維持 doc-only）+ glossary 註冊決策落實
- [ ] Open Questions 全部 `resolved` / `deferred`（附原因）並回寫
- [ ] 執行 Plan Completion Closure

## Phase 0 — Pre-Build Interrogation

### Phase 0.0 — Open Questions 核對（公版，必填）

逐條核對本 plan §Open Questions，標記處置並回寫：

- [ ] 已讀本 plan §Open Questions 全部條目
- [ ] 對每條標記 `resolved`（附 Phase 0 證據）/ `still-open` / `deferred`（附原因）
- [ ] `resolved` 的條目已同步勾選 / 附註於 §Open Questions
- [ ] 若盤點新發現問題，已加入 §Open Questions

| Open Question | 處置 | 證據 / 原因 |
|---|---|---|
| Q1–Q5 | still-open | 依 Phase 1/2/3 各自解決 |

### Phase 0.1 — 架構相容性 preflight

- [ ] 確認 `plans/README.md` §Delegation 與 `governance/lifecycle/plan-tree-hierarchy.md` §Delegation 現行內容（本 plan 不得與 03 已定稿 schema 衝突；只加 loop、不改 brief 必填集）
- [ ] 確認 review capability 現行 invoke 契約（`knowledge/runtime/capability-registry.yaml` + `workflow/cross-cutting/review/README.md`）：驗證 leg 復用 `fault_finding` stance，consumer 不重定義 stance
- [ ] 確認 03 baseline dogfood 方法論（brief independence score、fresh-context proxy）可平移為 verifier independence 判準
- [ ] 確認 `workflow/software-delivery/README.md` §Review invoke 邊界：本 loop 的 verifier leg 是 capability invoke 消費者，不是新 lifecycle slice

## Phase 1 — 協議定稿（doc-only）

- [ ] 決定 Q2 落點；落地 verifier 報告契約（4 欄位）+ 仲裁協議（fix/defer/reject 表）+ 3 條 role boundary invariants
- [ ] 決定 Q4 仲裁紀錄落點（傾向：被委派 sub-plan 內 table）
- [ ] 更新 `plans/README.md` §Delegation：從「雙路徑派發」擴為「派發 → 獨立驗證 → 仲裁」loop SOP（或依 Q2 決定的落點放置 + 交叉連結）
- [ ] Ai-skill repo 內委派的 bootstrap 注意事項寫入 SOP（brief `context.required` 須含 bootstrap 檔案）

## Phase 2 — 雙 dogfood

**焦點紀律（承 03）**：驗的是**契約自足性**（verifier 報告是否足以仲裁、brief 是否足以執行），不是執行者 / 驗證者聰不聰明。任何失效先反問契約缺漏。

- [ ] **2a — software-delivery 外部 repo 真實任務**：挑一個真實、小、可驗收的交付任務走完整 loop（orchestrator 寫 brief → executor 執行 → verifier 報告 → 仲裁）。記錄：verifier 差集 findings、仲裁分佈、orchestrator 是否越界執行、是否被迫回讀 diff。
- [ ] **2b — Ai-skill 內部任務**：比照 2a，驗證同一協議在治理 repo 內成立（含 bootstrap gate 注意事項實測）。
- [ ] 回饋迴路：任一 dogfood 暴露契約缺漏 → 修 Phase 1 文件（v2），不修執行者；重跑失效 leg。

## Phase 3 — 證據評估與收斂

- [ ] 彙整雙 dogfood evidence，回答 Q3（品質信號成立 / null result）
- [ ] Q5 決策：schema promotion（另立 plan）或明確維持 doc-only（記錄門檻與未達原因）
- [ ] Glossary 註冊決策落實（`independent_verification` / `arbitration` 註冊或明確不註冊）
- [ ] 執行 Plan Completion Closure（含 plans/README.md 狀態表更新、搬移 archived）

## Stakeholder 同意項目

> 描述現行選定策略（治理現況），改方向時直接更新本表。

| 決策面 | Current selected strategy |
|--------|---------------------------|
| Loop 形狀 | 三角色：orchestrator（規劃/切分/仲裁，不執行）/ executor（brief-only）/ verifier（fresh-context，只產證據） |
| 落地方式 | doc-only 協議 + 雙 dogfood；不動 schema、不接 runtime、不建自動 orchestrator |
| 驗證 leg | 復用 review capability `fault_finding` stance invoke，不另定 stance |
| 適用範圍 | advisory；只適用已宣告 delegation 的委派任務；主打 software-delivery，Ai-skill 比照 |
| Schema promotion | gated on Phase 2 證據（Q5），deadline 2026-08-31 |

## 與其他 plans 的關係

- [`2026-06-22-1009-plans-system-portability-and-delivery-integration/03-subplan-agent-delegation.md`](2026-06-22-1009-plans-system-portability-and-delivery-integration/03-subplan-agent-delegation.md) — **baseline**（completed）：delegation brief schema + 雙路徑 dogfood。本 plan 延伸其 loop（去程 → 回程），不重開該 sub-plan；自動 orchestrator reservation 邊界維持。
- [`2026-06-22-1009-plans-system-portability-and-delivery-integration/02-software-delivery-plan-first-ordering.md`](2026-06-22-1009-plans-system-portability-and-delivery-integration/02-software-delivery-plan-first-ordering.md) — plan-first ordering 是本 loop 的前置（orchestrator 產 plan artifact 先於執行）；本 plan 不改其 Q4 關閉條件。
- [`archived/2026-07-06-review-architecture-adr/_plan.md`](archived/2026-07-06-review-architecture-adr/_plan.md) — review = cross-cutting capability invoke（ADR-013 D2）；verifier leg 是其消費者。
- [`active/2026-06-16-1131-evidence-candidate-system.md`](active/2026-06-16-1131-evidence-candidate-system.md) — `defer` 處置的 findings 可轉 evidence candidate（人工 capture，不新增 scanner 職責）。
