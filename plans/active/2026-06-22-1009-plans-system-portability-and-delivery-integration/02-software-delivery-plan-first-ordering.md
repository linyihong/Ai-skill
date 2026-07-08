---
id: 2026-06-22-1009-software-delivery-plan-first-ordering
plan_kind: sub
status: completed
owner: linyihong
created: 2026-06-22
parent: 2026-06-22-1009-plans-system-portability-and-delivery-integration
required_for_completion: true
sub_plan_reason: >
  plan-first ordering 只動 software-delivery workflow 文件（intake 段），與 01
  的 CLI / validator 工作完全不同 owner / 不同 review 焦點，可獨立 sign-off 並
  與 01 並行。獨立成 sub-plan 避免把 workflow doc 改動混進 01 的 Go / CLI commit。
---

# Software-delivery Plan-First Ordering（sub-plan）

**Status**: `completed`（2026-07-08 Plan Completion Closure）
**Owner**: linyihong
**Parent**: [`_plan.md`](_plan.md)

## Source Request
讓 software-delivery 接入 plans 系統：「以後開發需要先寫 plans 系統把所有規劃好」。使用者選 **workflow 層 ordering（advisory）**，不做機械 block。

## Scope
- **In**：在 `workflow/software-delivery/` intake 段加入 plan-first ordering，明文接在 pre-build-interrogation / Test-First Ordering 之後；advisory + review checklist；可選 validation scenario。
- **Out**：commit-msg 機械 block（無 active plan 不准 commit code）— 保留為後續 maturity ladder 升級候選。
- **Affected**：`workflow/software-delivery/intake.md`、`workflow/software-delivery/execution-flow.md`（導航）、可能 `workflow/software-delivery/test-strategy.md`（Test-First Ordering 接點）、可能新增 validation scenario。

## Decision Rationale（sub 層）
現有 intake 已有 pre-build-interrogation（goal/scope/non-goals/acceptance/source-of-truth/duplication risk）與 Test-First Ordering（framework/runtime/governance 升級強制順序）。plan-first **不是新 gate，而是把「實作前先有 plan artifact」明文化為 intake 的一環**，並連到 plans 系統（`plans/active/` + plan-tree）。

**關鍵修正（回應 review #4）：plan 是 artifact，不是一個一次性 stage。** 線性「interrogation → plan → preflight → implement」是錯的，因為 preflight 會**回改 plan**（架構相容性檢查發現衝突就更新 plan）。正確模型是 plan 在 preflight 間反覆收斂：

```
Discover → Interrogate → Draft Plan ⟲ Preflight → Execute
                              └──────────┘
                         preflight 可回改 plan
                         （plan 非一次生成）
```

不重複 pre-build-interrogation（Q4）的分工：pre-build-interrogation = 需求拷問（產出 plan 的輸入）；plan-first = 拷問結果落成可收斂的 plan artifact；Architecture Compatibility Preflight = 對 draft plan 做相容檢查並回饋修正。三者不是序列三段，而是 interrogation 餵入、plan 為中心 artifact、preflight 反覆驗證。

### Alternatives
- A. 硬機械 gate：reject（本輪）— 易誤擋小修補，使用者已選 advisory。
- B. 完全不接、只靠既有 preflight：reject — 缺「先有 plan artifact」的明文順序，plan 與 delivery 仍脫節。
- C. advisory ordering + review checklist（accept）。

## Open Questions（本 sub）
- Q4（與 pre-build-interrogation 不重複）— 見 main §Open Questions。
- 是否所有 software-delivery 任務都要 plan，還是依規模分級（小修補豁免）？

## Phase 0 — Pre-Build Interrogation

### Phase 0.0 — Open Questions 核對（公版，必填）
- [x] 已讀 main + 本 sub §Open Questions 全部條目
- [x] 對每條標記 `resolved` / `still-open` / `deferred`
- [x] `resolved` 條目回寫（main Q4 → resolving；規模分級 → resolved-by-doc）
- [x] 新問題已加入 §Open Questions（無新問題）

| Open Question | 處置 | 證據 / 原因 |
|---|---|---|
| Q4 不重複 pre-build-interrogation | resolving（Phase 1 落地後 close） | Phase 0.1 已界定分工：interrogation = 需求拷問（產出 plan 的輸入）；plan-first = 拷問結果落成可收斂的 plan artifact；Architecture Compatibility Preflight（`plans/README.md` L122-161）= 對 draft plan 做相容檢查並回饋修正。三者非序列三段。intake §Plan-First Ordering 落地 + 一次真實 intake 含 preflight 回改實例後由 02 owner 於 main §Open Questions 標 resolved |
| 規模分級豁免 | resolved | 接 `plans/README.md` plan-tree「何時不開 sub-plan」（L41：`sub_plan_reason` 非空為唯一強制；單一 phase step 用 checkbox、< 1 session inline、純文件補強直接 commit，不開 sub-plan）。intake advisory 直接引用此既有規則作豁免條件，不新增平行規則 |

### Phase 0.1 — 架構盤點
- [x] 讀 `workflow/software-delivery/intake.md` 現行 intake 順序與 pre-build-interrogation 內容。（Change Intake → Pre-build Interrogation Gate → Requirements Cognition Checkpoint → Parity Gate；Plan-First Ordering 接在 Pre-build Interrogation Gate 之後）
- [x] 讀 `workflow/software-delivery/test-strategy.md` Test-First Ordering，確認接點。（§5 Test-First Ordering 為 framework/runtime/governance 升級**強制**順序；plan-first 為 advisory，與其正交、互不覆蓋，於 intake 交叉引用）
- [x] 讀 `plans/README.md` Architecture Compatibility Preflight（已要求實作前 preflight），界定 plan-first 與 preflight 的關係（避免三重 gate）。（`plans/README.md` L126 已排序 pre-build-interrogation → preflight，且 preflight 可回改 plan；plan-first **只引用不重寫**該排序，避免 dual source-of-truth / triple gate）

### Phase 0.2 — Architecture Compatibility Preflight 記錄

| 欄位 | 內容 |
| --- | --- |
| Trigger | 開始執行 sub-plan 02 Phase 1（software-delivery intake plan-first ordering doc） |
| Checked sources | `intake.md`、`test-strategy.md` §5、`execution-flow.md`、`plans/README.md`（Preflight + plan-tree）、`requirements/pre-build-interrogation.md`、`closure.md`、`review-checklist.md`、software-delivery `README.md` |
| Conflicts | 無架構衝突。關鍵邊界：`plans/README.md` L126 已序列化 pre-build-interrogation → Architecture Compatibility Preflight（preflight 可回改 plan）；Test-First Ordering（test-strategy §5）為 framework/runtime/governance 變更的**強制**順序。plan-first 須**引用**上述而非重寫，避免 triple gate / dual source-of-truth |
| Interrogation | Goal：intake 明文化「先有 plan artifact 再實作」的 advisory ordering + loop 模型 + 規模分級豁免 + review checklist。Scope in：intake 小節、execution-flow 導航、review-checklist 一條；out：commit-msg 機械 block、runtime gate。Acceptance/validation target：intake 小節落地 + loop 模型 + Q4 分工清楚 + 規模分級 + checklist + doc-only 宣告；完整 acceptance（真實 intake 走過 loop + 一次 preflight 回改）為 completion #6，需真實使用實例。Framework discovery：canonical source = workflow layer（intake.md）；plans 系統 canonical = `plans/README.md`；不接 runtime.db。Duplication risk：僅引用既有 preflight/interrogation 排序，不重寫。 |
| Open Questions 核對 | Q4（main）→ resolving（Phase 1 落地 + 真實 intake 實例後 close）；規模分級豁免 → resolved（引用 plan-tree 既有規則） |
| Decision | proceed |
| Validation | diff review + markdown link 一致性 + plan checkbox sync；真實 intake loop 實例列為 needs-validation（completion #6） |

**Glossary discovery**：plan-first 復用既有 plans 系統與 software-delivery 詞彙（`plan_kind` / `plan_tree` / `sub_plan_reason` 已在 `knowledge/glossary/ai-skill.md`），no new framework vocabulary introduced。

## Phase 1 — Plan-first ordering 文件化
- [x] 在 intake 段新增「Plan-First Ordering」小節（`intake.md` §Plan-First Ordering，接在 Pre-build Interrogation Gate 之後）：明文「會導向 code/workflow/governance/runtime/schema/generated artifact/tool adapter 改動的任務，實作前應有對應 `plans/active/` plan（inline 小 plan 或 plan-tree）」。
- [x] **明文一句防 loop 被忘**：小節開頭 blockquote「Plan 是 artifact，不是一個一次性 stage」+「plan 在 preflight 間反覆收斂」——避免誤讀成 `Interrogate → Plan → Preflight` 線性。
- [x] **用 loop 模型描述（非線性三段）**：`Discover → Interrogate → Draft Plan ⟲ Preflight → Execute` code block，明寫 preflight 可回改 plan、plan 非一次生成。
- [x] 分工說明：三欄表（interrogation 餵入 / plan 為中心 artifact / preflight 反覆驗證回饋），各指向 canonical source，不重寫（Q4）。
- [x] 規模分級：直接引用 `plans/README.md` plan-tree「何時不開 sub-plan」（< 1 session inline / 純文件補強直接 commit / surgical 小修補走 surgical-changes），不新增平行規則。
- [x] advisory 語氣（blockquote 明寫「非 commit-time 機械 block」，用「應 / 建議」）。

## Phase 2 — Review checklist + 可選 scenario
- [x] 在 `review-checklist.md` §Change Intake 加一條 Plan-First Ordering（advisory）檢查項，指向 intake 規模分級豁免。
- [x] （可選 scenario 決策）**不新增 validation scenario；本輪 doc-only**。理由：plan-first 為 advisory workflow ordering，不接 runtime gate、無 commit-msg validator、無可機械觀察的 blocking 判定；新增 scenario 會暗示 mechanical enforcement 意圖，與 advisory 定位矛盾。**未來升級條件**：若累積證據顯示 advisory 被忽略造成 plan↔delivery 脫節（maturity ladder：ordering 觀察 → evidence → 再評估 commit-msg validator / runtime gate），屆時才補 scenario-first + validator（沿 `test-strategy.md` §5 Test-First Ordering 強制流程）。
- [x] 更新 `execution-flow.md` Cognitive Slice 導航 Intake 列，加入 Plan-First Ordering 指向新小節。

## 完成條件
- [x] intake plan-first ordering 小節落地，用 loop 模型、與 pre-build-interrogation / preflight 分工清楚（Q4 分工已界定；待真實 intake 實例後由 02 owner close main Q4）
- [x] 規模分級豁免條件落地（引用 plan-tree 既有規則）
- [x] review checklist 更新
- [x] doc-only 宣告明確（Phase 2 可選 scenario 決策 + main §Runtime Execution Path「02 若不接 runtime，須明寫 doc-only」）
- **Acceptance evidence（回應 review #6）拆為兩條不同成熟度的子命題（2026-07-06，使用者定調）**：ordering（plan-first 是否發生）與 feedback-loop（preflight 是否實質改 plan）是兩個獨立能力，不混為一談、不因前者成立就宣告後者完成。

    | 子命題 | 狀態 | 證據 |
    |---|---|---|
    | Plan-first ordering | ✅ Verified | Brower `verification-code-center` 真實演化、Git 時序可驗、非 backfill |
    | Preflight feedback loop | ✅ Verified | Brower [`preflight-feedback-log §2026-07-07 tiered worker`](../../../Brower/docs/plans/preflight-feedback-log.md) — clean T1（改 plan 前獨立段落）+ T0<T1<T2；material ∈ {scope, sequencing, dependency, acceptance}；**closes 02-B / main Q4**（2026-07-08） |

  - [x] **A. Ordering Evidence — ✅ Verified（Brower，2026-07-06）**：真實 software-delivery 專案 Browser Manage 的 [`docs/plans/active/2026-07-03-verification-code-center.md`]——真功能（非 toy）、`T0`=2026-07-03 起草於 implementation 前、2026-07-06 多次**有序** git commit 演化、git timestamp 可驗、README/`intake.md` §Plan-First Ordering 對齊、**非 backfill**。→ 證明 plan-first ordering 可操作、可持續、**非事後敘事**。**不再要求更多同型證據**（accumulating，非 hunting）。
  - [x] **B. Feedback Loop Evidence — ✅ Verified（Brower，2026-07-08）**：[`preflight-feedback-log §2026-07-07 tiered-data-archive-platform — worker 模組分離`](../../../Brower/docs/plans/preflight-feedback-log.md) — `T0`（2026-07-06 初稿：retention job 可放 admin API 進程）→ `T1`（2026-07-07 本 log 獨立段落：batch 不應與 API 同進程、需 worker + ShedLock）→ `T2`（同日 plan 修訂 D7–D10 + Phase A acceptance）。Material：scope（worker 模組）、sequencing（API 禁止跑 job）、dependency（verification Phase 2）、acceptance（§7 worker 驗收）。**不 hunt 歷史湊例** — 自然累積於 tiered platform preflight。
  - **B 的 falsifiable rigor（沿用）— planning artifact 的「演化」而非某檔案的演化**：planning artifact = 當下真正承載 planning state 的產物（`plan.md` / analysis doc / ADR / design doc），(a) 不綁單一檔名；(b) 不鼓勵為驗證刻意改 plan 檔；(c) artifact-centered。
  - **B 的 falsifiable rigor（沿用）— 三個 ordered timestamp 必留**：`T0` 可辨識 planning artifact / `T1` **獨立** evidence 顯示假設被推翻（非終稿內事後補述）/ `T2` 後續 planning artifact 明顯不同且**可追溯到 T1**。缺 T0<T1<T2 獨立時序痕跡 → 「preflight 真的改變計畫」與「只是發現問題 / 事後重建」無法區分 → 不得 close。
  - **B 的候選評估紀錄（不代表已 close）**：Browser Manage 另一候選為 evidence-driven 修正（初版假設被 live 抓包推翻 → 全面回改 SDK/docs），時序獨立可查、behaviorally satisfies loop，但 planning state 落在 analysis/SDK 而非集中 planning artifact（artifact locality 偏弱）→ **strong supporting**，未達 B。consolidation/backfill 候選（execute→plan）→ **不計入**。
- [x] linked-updates 檢查（execution-flow 導航 / review-checklist / intake 三處同步；plans/README 只引用不改）
- [x] 執行 Plan Completion Closure — 2026-07-08（sub-plan 留 tree folder；`status: completed`）

## Plan Completion Closure（2026-07-08）

| # | 檢查 | 結果 |
|---|---|---|
| 1 | 完成條件 A+B | Ordering ✅ + Feedback loop ✅（tiered worker preflight） |
| 2 | runtime refresh | N/A — doc-only workflow intake |
| 3 | linked-updates | intake / execution-flow / review-checklist 已同步 |
| 4 | main Q4 | 由 02 owner 標 **resolved**（見 parent `_plan.md`） |
| 5 | 搬移 | **不搬** — 與 01 相同，sub-plan 留 portability tree folder；lifecycle `completed` 與 storage 分離 |

## Glossary Impact
Glossary Impact: no — plan-first ordering 復用既有 plans 系統與 software-delivery 詞彙，no new framework vocabulary introduced。

## 與其他 plans 的關係
- 接入 [`workflow/software-delivery/intake.md`](../../../workflow/software-delivery/intake.md) 與 [`workflow/software-delivery/test-strategy.md`](../../../workflow/software-delivery/test-strategy.md)。
- 復用 [`plans/README.md`](../../README.md) Architecture Compatibility Preflight 與 plan-tree「何時不開 sub-plan」規則。
