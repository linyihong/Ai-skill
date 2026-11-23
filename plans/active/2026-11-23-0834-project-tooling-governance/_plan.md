---
id: 2026-11-23-0834-project-tooling-governance
plan_kind: main
status: in-progress
owner: workflow
owner_layer: workflow
created: 2026-11-23
parent: null
---

# Project Tooling Governance（專案工具治理）

**Status**: in-progress — Plan Preparation、Phase 0、**Phase 1 完成**（2026-11-23，契約通過使用者最終 sign-off）；Phase 2 / 3 未開始（consumer dogfood，需另行啟動）。不改任何 consumer repo。

**本 plan 是唯一的執行狀態來源。** 契約設計大綱見 [`01-design-contracts.md`](01-design-contracts.md)（不記狀態）。

**Glossary Impact**: yes（候選，Phase 1 前不登記）— `tooling_role_contract`、`tool_discovery`、`creation_justification`、`tool_placement`、`lifecycle_state`、`concept_check`、`capability_catalog`、`reuse_evidence`、`promotion_ladder`、`compensating_lock`。

## 一句話目標

治理 AI 在專案中**創造、使用、保留與淘汰工具**的行為，讓分析／驗證工具從「每做一次分析就多一支」收斂成分層、參數化、有生命週期的工具面。新工具可以長，但只在確認沒有同概念能力時才長；新增一個工具應該是有意識的架構決策，不是完成任務的預設副作用。腳本是目前最明顯的案例，不是唯一範圍。

## Decision Rationale

### Problem & Why Now

兩個性質不同的 consumer 專案都出現同一個根因：**沒有「何時必須抽共用 / 參數化」的規則，也沒有單一 registry 驅動的入口**，於是鎖只能逐事件加，腳本逐分析長。

**Consumer A — 逆向分析專案（Python + Frida，裝置自動化與資料擷取）**：已有分層雛形（L0 薄入口 `scripts/cli/<domain>/` → L1 領域模組 → L2 `shared/` → L3 per-target config + strategy enum；例外區 `<target>/<id>/`、`frida/research/`）與 24 個 pre-commit 機械鎖（layout、boundary、completeness、drift、lifecycle 五族）。

| # | 缺口（Phase 0：全部 verified，見 [`evidence/phase0-baseline.md`](evidence/phase0-baseline.md)） | 量化 |
| --- | --- | --- |
| A1 | L0 薄入口是慣例不是鎖 | 3 支 `cli/` 入口非薄（31 / 132 / 260 行）；最大者自帶 `tap()`、裝置序號預設值、模組級座標常數 |
| A2 | boundary 鎖 scope 只覆蓋單目標目錄 | 放進 `cli/` 即繞過 device I/O 鎖 |
| A3 | 裝置序號預設值分散 | 4 個檔案帶寫死預設（3 支在 L0、1 支在 L2 shared 本身），未統一走 shared resolver（Phase 0 修正原記 7 檔） |
| A4 | 無升級觸發條件 | README 只寫「禁止」，未寫何時必須抽 |
| A5 | 鎖本身逐檔蔓延 | 24 支獨立 check 腳本 + 手寫 pre-commit 清單 |

**Consumer B — 業務交付 monorepo（bash + PowerShell + Node 驗證閘）**：邏輯已抽到 `lib/check_*.mjs`，但：

| # | 缺口（Phase 0：全部 verified，見 [`evidence/phase0-baseline.md`](evidence/phase0-baseline.md)） | 量化 |
| --- | --- | --- |
| B1 | 入口 wrapper 大量重複 | `scripts/` 根目錄 132 項；53 組 `verify-*.sh` + `.ps1`，其中 26 組是 ≤20 行的單一 lib 呼叫（Phase 0 修正原記「多數」） |
| B2 | **compensating lock** | `verify-scripts-parity` 存在的理由是 wrapper 被複製成兩份 |
| B3 | `lib/` 平鋪 | 150 檔混放 check / test / deploy / scope helper |
| B4 | **一次性分析腳本蔓延** | `performance/` 8 支 per-case 腳本（27–154 行）各自重寫 fixture 斷言、計時、報告；零個專案內 import |

**更深一層的根因（agent 工作習慣）**：AI 執行能力越強，越容易為完成當下任務創造新工具，卻不判斷工具是否已存在、是否該共用、任務結束後如何處置。累積路徑：一次性腳本 → 沒找到舊的再寫一支 → 變成無入口無責任的半永久工具 → 為防再犯加 lint / pre-commit → 腳本、wrapper、validator、allowlist、hook 一起增長。第 4 步是關鍵：**用更多機械鎖，管理原本可以靠更好的工具架構消除的重複**。

**Why now**：兩個 consumer 的鎖都在快速增加（A：24；B：~50 gate）；Consumer A 剛完成一輪分層整理，形狀新鮮、證據可直接量測。

### Decision

建立 **tool-neutral、language-neutral** 的 Project Tooling Contract，以 doc-only pilot reference contract 落在 `workflow/cross-cutting/project-tooling/`（比照 `traceable-evidence`：非 stage、不註冊 route），由兩個 consumer 各做一個 dogfood slice 驗證，再依證據決定是否升級。

**治理對象是工具的生命週期，不是禁止新增工具。** 契約六部分（詳見 [`01-design-contracts.md`](01-design-contracts.md)）：

1. **Role contract**（原 Layer contract，Q10）— 工具內部角色 Entry / Logic / Parameters；違規以「某 placement 的某角色放錯位置」描述，跨語言適用。
2. **Discovery contract** — 先找再建；四種合理新增理由；輕量 justification，不是表單。
3. **Lifecycle contract** — **Placement**（Scratch → Research → Target → Domain → Shared）與 **State**（temporary / experimental / maintained / deprecated）；與 Role 共三個獨立維度（Q10）。
4. **Growth rule** — concept-first（Selection）+ reuse evidence（Constraint：第 1 次定位置、第 2 次評估、第 3 次明確裁決；計數對象是同一能力 / 責任，不是相似程式碼）+ config-first ≠ config-everything + promotion ladder。
5. **Lock taxonomy** — 七族鎖形狀；鎖保護架構不變式，不保護檔案數量。
6. **Compensating-lock Review Rule** — 新增鎖前必答「能否從結構上消除？」

**明確不做**：共用 runner（Q5 deferred）、runtime route、自動 orchestrator（是否抽象化屬 Selection Responsibility）。

### Alternatives Considered

- A. 維持現狀，各專案各自整理：reject — 成果無法遷移，compensating lock 無框架可指認。
- B. 立即實作 Ai-skill CLI 通用 lint runner（Go）：reject（暫緩）— 證據 n=2、兩個 consumer 語言不同，runner 介面的 invariant 未知（Watch-Out wall 2 / 6）。
- C. 只修 Consumer A 的漏洞：reject — 失去跨專案學習，B4 照樣長。
- D. **Doc-only 契約 + 兩個 consumer 各一 slice dogfood，再依證據升級**（accept）。
- E. 規定 agent「不准新增工具」：reject — 一次性與實驗工具有正當價值；要治理的是生命週期與發現。
- F. 自動 orchestrator 偵測新工具並自動搬層：reject — 越過 Selection Responsibility，抽象化決策失去證據。
- G. 名稱縮小為 Script Governance：reject — 真正的對象是工具資產的生命週期（DB 查詢、API 測試、OCR / ASR、部署驗證工具同樣適用）。

### Why Not an ADR Yet

概念判定準則、discovery 是否需機械化、lifecycle 雙維度是否完備、lock taxonomy 是否完備、共用 runner 是否值得，全部未驗證；且有更輕的 promotion target（cross-cutting pilot contract）。

### ADR Promotion Criteria（completed 時驗證）

- [ ] foundational + cross-session + cross-project + expensive-to-reverse + explains-why 全中
- [ ] 兩個 consumer dogfood 皆證實 contract 可行，且未需修改 invariant 核心
- [ ] Open Questions 全解
- [ ] 沒有更輕的 promotion target 適用（per ADR-007）
- [ ] 第三個專案採用時不需修改 contract（真實使用 evidence）

### Consequences（預期）

- **正面**：新需求有明確的「先找、先參數、再共用、最後新建」順序；鎖依 taxonomy 套形狀並可辨識 compensating lock。
- **負面**：多一份 cross-cutting 契約要維護。
- **風險**：契約膨脹成 BPMN 式流程（Wall 2）— 緩解：6 檔各 ≤150 行，只寫 invariant 與觸發條件；consumer 細節外洩 — 緩解：§Sanitization。

### Watch-Out List citation

[`architecture/ai-native-cognitive-ecosystem-system.md`](../../../architecture/ai-native-cognitive-ecosystem-system.md) §Watch-Out List — **Wall 2 Workflow inflation**（契約只定 invariant，不定逐步流程）；**Wall 6 Optimization hallucination**（不在 n=2 時做通用 runner、不做自動 orchestrator）。

## Runtime Execution Path

**Doc-only trial，不接入 runtime。** 不新增 `route.*`、不 project 到 `runtime.db`、不新增 Ai-skill commit-msg validator。機械強制發生在 consumer repo 自己的 hook（Phase 2/3），不屬於 Ai-skill runtime integration。

- 未來接入：Phase 4 決策；entry condition = 兩個 consumer dogfood PASS，且符合 cross-cutting 三案例門檻（目前 n=2）。
- 候選形式：discovery signal（新增工具檔時載入本契約）或 `init-project` 範本；需另開 plan 並附 Per-surface consumer 表。
- Graduation deadline：2027-02-28（逾期未接入 → 維持 doc-only pilot 並在 plans/README 說明）。

### Per-surface consumer 表

不適用 — 本 plan 不新增 generated surface、route 或 validator。

## Sanitization

共用層只用 Consumer A / Consumer B 抽象描述（領域 + 語言），不寫 repo 名、裝置序號、座標、本機路徑或 SHA。具體對應放各 consumer 的 project-local、非 git `.ai-skill/dogfood/2026-11-23-0834-project-tooling-governance.md`。私有識別資訊登記見 Plan Preparation。

## Open Questions

- [x] **Q1 放置位置** — decided (2026-11-23, user review)：`workflow/cross-cutting/project-tooling/`。
- [x] **Q2 生長 / 升級門檻** — decided (2026-11-23, user)：concept-first + reuse evidence（第 1 次定位置、第 2 次評估、第 3 次明確裁決）；計數對象是同一能力 / 責任，不是相似程式碼。
- [x] **Q2a 概念判定準則** — resolved (2026-11-23, Phase 0)：「產出物 + 核心責任」兩者都相同才算同概念；兩個 consumer 共 5 個概念群都判得出來，大廳 tile 群（同產出物、不同責任）正確判為不同概念；無無法判定的案例。見 [`evidence/phase0-baseline.md`](evidence/phase0-baseline.md) §Concept 試判。
- [x] **Q2b Capability catalog 形狀** — **accepted (2026-11-23, user)**：功能分類 README 能力表（capability / entry / placement / state / owner / ttl）；capability 是主要識別單位（可多個 entry、只登記一列）；多分類時工具樹根 README 須有 catalog 索引入口。機器可讀版本 deferred，與 Q7 一起評估。見 [`lifecycle-contract.md`](../../../workflow/cross-cutting/project-tooling/lifecycle-contract.md) §Capability catalog。
- [x] **Q2c Research 區** — **accepted (2026-11-23, user)**：Research = Placement、experimental = State、TTL = 期限而非自動刪除；TTL 內免 concept check 但須記錄目的與 TTL；到期必須重評 concept + reuse evidence；逾期未評估不再視為有效 experimental；去向三選一（maintained / 重設 TTL 附理由 / deprecated）。見 [`lifecycle-contract.md`](../../../workflow/cross-cutting/project-tooling/lifecycle-contract.md) §Research 與 concept check。
- [x] **Q3 Compensating-lock test** — decided (2026-11-23, user review)：新增鎖時必答的 Review Rule。
- [x] **Q4 Consumer B wrapper 收斂** — decided (2026-11-23, user review)：獨立裁決，不進本 plan dogfood。
- [x] **Q5 Runner 語言** — deferred (2026-11-23, user review)：無證據需要共用 runner；Phase 4 再評估。
- [x] **Q6 Sanitization denylist** — decided (2026-11-23, user review)：登記必要私有識別資訊；於 Plan Preparation 以獨立 commit 執行（Phase 0 維持唯讀）。
- [x] **Q7 Discovery 機械化** — deferred (2026-11-23, user review)：先以 workflow behavioral 契約在 dogfood 驗證；Phase 4 再評估。
- [x] **Q8 Scratch 位置與清理責任** — **accepted (2026-11-23, user)**：不進版控固定目錄或 session 暫存區、不得在正式工具樹；建立者於 close-out 清理；**無法清理時回報殘留路徑、原因、後續處理**；有保留價值轉 Research，否則刪除。見 [`lifecycle-contract.md`](../../../workflow/cross-cutting/project-tooling/lifecycle-contract.md) §Scratch 位置與清理責任。
- [x] **Q9 Ownership** — **accepted (2026-11-23, user)**（依審閱修訂後通過）：單一記錄點 = catalog `owner` 欄、可為團隊 / 角色、不建 registry（保留）；**刪除「無 owner 視為 experimental 或 deprecated」**——owner 缺失是治理資訊不完整，不改變 state：maintained + 無 owner = ownership violation 須補 owner；experimental + 無 owner 可接受但須 TTL 與目的；deprecated 不要求維護 owner 但須退役處理責任方；改 deprecated 須由 Selection 明確決定。見 [`lifecycle-contract.md`](../../../workflow/cross-cutting/project-tooling/lifecycle-contract.md) §Ownership。
- [x] **Q10 Layer / Placement 分離** — **decided (2026-11-23, user)**：三個維度正式分離、不得互相取代或混用——(1) **Internal Role**：Entry / Logic / Parameters；(2) **Placement**：Scratch / Research / Target / Domain / Shared；(3) **Lifecycle State**：temporary / experimental / maintained / deprecated。附帶規則：Placement 不決定 Role（任一 placement 都可有三種角色，但不要求同時具備）；Entry 維持薄入口（只解析參數、取得必要設定、呼叫 Logic、傳遞結果與結束狀態）；Parameters 不得成為邏輯藏身處。Phase 1 正式定義違規判定與跨語言適用原則。

## Plan Preparation（Phase 0 之前，獨立 commit）

- [x] Q6：在 `.ai-skill-project.yaml` `private_entities` 登記兩個 consumer（Consumer A = ExternalRepoD、Consumer B = ExternalRepoE）；`ai-skill runtime compile` 後 `derived_match_tokens` 含其 case variants，暫存 probe 檔驗證 pre-commit 會擋（2026-11-23）
- [x] 前置清掃（使用者 2026-11-23 選擇「登記＋清掃既有」）：共用層 16 檔中的兩個名稱改為 alias，5 個 archived delegation-loop evidence 檔改中性檔名並修連結；獨立 commit，先於登記
- [x] 本步驟是唯一會修改 Ai-skill 設定的前置工作；完成後才進 Phase 0

## Phase 0 — Read-only Discovery + Preflight

**唯讀**：只讀、只量測，不修改 Ai-skill 設定或任何 consumer repo。

### Phase 0.0 — Open Questions 核對（公版，必填）

- [x] 已讀本 plan §Open Questions 全部條目
- [x] 對每條標記 `resolved`（附 Phase 0 證據）/ `still-open` / `deferred`（附原因）
- [x] `resolved` 的條目已同步勾選 / 附註於 §Open Questions
- [x] 若盤點新發現問題，已加入 §Open Questions（Q10）

| Open Question | 處置 | 證據 / 原因 |
|---|---|---|
| Q1、Q2、Q3、Q4、Q6 | decided（2026-11-23） | 使用者決定；Q6 於 Plan Preparation 執行 |
| Q5、Q7 | deferred | Phase 4 再評估 |
| Q2a | resolved | Phase 0 concept 試判（evidence/phase0-baseline.md） |
| Q2b、Q2c、Q8 | accepted（2026-11-23, user） | 補充約束已寫入 lifecycle-contract |
| Q9 | accepted（2026-11-23, user） | owner 缺失不推導 state |
| Q10（新增） | decided（2026-11-23, user） | 三維度分離，見 §Open Questions |

### Phase 0.1 — Baseline 量測（可重跑指令記在 project-local dogfood 檔）

| 指標 | 意義 |
| --- | --- |
| scripts 根目錄項目數 | layout 蔓延程度 |
| 非薄 L0 入口數（> 25 行或含 I/O） | thin-entry 缺口 |
| 近重複 wrapper 數 | 結構重複 / compensating lock 規模 |
| per-case 工具的專案內 import 比例 | 一次性工具是否共用 harness |
| 機械鎖數、每鎖 scope | 鎖蔓延與盲區 |

- [x] A1–A5、B1–B4 從 `observed` 升為 `verified` 或 `refuted`（verify 對 HEAD）— 全部 verified，A3、B1 數量修正
- [x] **Concept 試判**：逐支套 concept check，產出「能力 / 責任 → 工具清單」表（以責任分組，不以程式碼相似度）；同概念多支者為 Phase 2/3 候選；邊界案例記錄給 Q2a
- [x] 每支工具試標 placement × state，檢驗雙維度是否足以描述現況 — 足以描述，唯一例外引出 Q10

### Phase 0.2 — Framework discovery（避免雙份 source-of-truth）

- [x] 讀 [`workflow/cross-cutting/README.md`](../../../workflow/cross-cutting/README.md)、[`traceable-evidence/`](../../../workflow/cross-cutting/traceable-evidence/README.md)
- [x] 讀 [`2026-06-08-2100-governance-pattern-library-extraction.md`](../2026-06-08-2100-governance-pattern-library-extraction.md) — 確認 lock taxonomy 不重疊（無重疊）
- [x] 讀 software-delivery framework domain model（archived）§Automation deferred — 確認不重複（鎖 = consumer 的 Automation 層；本 plan 不實作通用 Automation）
- [x] [`pre-build-interrogation.md`](../../../workflow/software-delivery/requirements/pre-build-interrogation.md) 完成記錄
- [x] Architecture Compatibility Preflight 最低記錄格式 → Decision: **proceed**（記錄於 [`evidence/phase0-baseline.md`](evidence/phase0-baseline.md) §Framework discovery 與 Preflight）

**Phase 0 完成條件**：baseline 表、缺口 verified / refuted、concept 試判表、placement × state 試標、preflight Decision 已回寫本 plan。— **完成（2026-11-23）**。Project-local 對應檔已建於兩個 consumer 的 `.ai-skill/dogfood/`（非 git；Consumer A 新增 `.git/info/exclude` 條目，tracked 內容未改）。Phase 4 升級判斷沿用 cross-cutting 三案例門檻。

## Phase 1 — 契約交付（Ai-skill，doc-only）

依 [`01-design-contracts.md`](01-design-contracts.md) 在 `workflow/cross-cutting/project-tooling/` 產出：

- [x] `README.md`（索引、適用範圍、何時載入、設計目標）
- [x] `role-contract.md`（Q10：Entry / Logic / Parameters、違規判定、跨語言原則）
- [x] `discovery-contract.md`（behavioral invariant，非填表）
- [x] `lifecycle-contract.md`（placement × state 雙維度；**含 Q8 scratch 位置與清理責任、Q9 ownership 最小契約**）
- [x] `growth-rule.md`（concept check、reuse evidence、config-first ≠ config-everything、promotion ladder）
- [x] `lock-taxonomy.md`（七族鎖形狀、compensating-lock Review Rule、check architecture 概念）
- [x] `workflow/cross-cutting/README.md` 加入連結；linked-updates 檢查（另同步 `workflow/README.md` cross-cutting 列；不新增 route，故 routing registry / 根 README 不適用）

**完成條件**：6 檔各 ≤150 行；Q8、Q9 resolved；Q10 已裁決並寫入 Role contract（含違規判定與跨語言原則）；Q2a–c resolved 或 deferred（附原因）；無 consumer 識別資訊（sanitization gate 通過）；不新增 route、runner 或 validator。

**Phase 1 交付（2026-11-23）**：6 檔已建於 [`workflow/cross-cutting/project-tooling/`](../../../workflow/cross-cutting/project-tooling/README.md)（56–72 行，皆 ≤150）；Q2b、Q2c、Q8、Q9 皆 accepted（Q9 依審閱修訂後通過）；**Phase 1 完成（2026-11-23，使用者最終 sign-off）**；sanitization 通過；無 route / runner / validator。

**Phase 2/3 必須等 Phase 1 完成才開始**（先確認契約正確，再用真實案例驗證）。

## Phase 2 — Consumer A dogfood（在 consumer repo 執行）

Slice：補 thin-entry 漏洞，並以 discovery / lifecycle 契約處置 3 支非薄入口。

- [ ] 新增 thin-entry 鎖：L0 入口只允許 import + `main()`；含 subprocess / 自寫 tap / 座標常數即擋
- [ ] 擴大 device I/O boundary 鎖 scope 至 L0（先套 compensating-lock Review Rule）
- [ ] 3 支非薄入口：concept check → 標 placement × state → allowlist（附退場條件）→ 邏輯移到正確 placement 並改走 shared device helper → 移出 allowlist
- [ ] 裝置序號預設值收斂到 shared resolver

**Acceptance**：非薄 L0 入口數 = 0；allowlist 為空；鎖有反向測試；每支被移動的工具有 justification 與 placement × state。

## Phase 3 — Consumer B dogfood（在 consumer repo 執行）

Slice：一次性分析工具 → 共用 harness（B4）。wrapper 收斂依 Q4 另案。

- [ ] concept check：確認 per-case 工具屬同一能力 / 責任
- [ ] 抽出共用 harness（fixture 驗證、計時、報告輸出）
- [ ] 先遷移 2 支，per-case 檔只剩 case 參數 / 宣告
- [ ] 評估其餘是否可改為 config-only case（並檢查未變成 config-everything）
- [ ] 視需要加一條 lock（先套 compensating-lock Review Rule）

**Acceptance**：遷移後行數下降且輸出報告與遷移前等價（同一 fixture 對照）；專案內 import 比例 > 0。

## Phase 4 — 證據裁決與升級決策

- [ ] 對每條契約 invariant 標記：兩個 consumer 是否都成立、是否需修改
- [ ] 重新評估 Q5（共用 runner）、Q7（discovery 機械化）與 runtime 接入
- [ ] 回寫契約；更新 `plans/README.md` 狀態
- [ ] 執行 Plan Completion Closure

## Stakeholder 同意項目

- [x] 名稱維持 Project Tooling Governance（不縮小為 Script Governance）
- [x] 放置位置 cross-cutting（Q1）；生長規則 concept-first + reuse evidence（Q2）
- [x] Compensating-lock 為 Review Rule（Q3）；wrapper 收斂另案（Q4）；runner、discovery 機械化 deferred（Q5、Q7）
- [x] Q6 移到 Plan Preparation，Phase 0 維持唯讀
- [x] Q10：Role / Placement / Lifecycle State 三維度分離（2026-11-23）
- [x] 概念判定準則與 capability catalog 形狀（Q2a Phase 0 證據；Q2b accepted）
- [x] Research 免 concept check 的範圍（Q2c accepted）
- [x] Scratch 最小契約（Q8 accepted）
- [x] Ownership 最小契約（Q9 accepted）
- [x] Phase 1 契約最終 sign-off（2026-11-23）
- [ ] Phase 2 / Phase 3 dogfood slice 範圍

## 與其他 plans 的關係

- [`2026-06-08-2100-governance-pattern-library-extraction.md`](../2026-06-08-2100-governance-pattern-library-extraction.md)：lock taxonomy 可能成為其 pattern 候選；本 plan 不搶先登記。
- `archived/2026-10-04-0936-evidence-record-contract`：cross-cutting pilot contract 的形狀先例。
- Go-first script policy（Ai-skill 自身腳本）：僅影響 Q5 的 runner 實作語言，不約束 consumer。

## 完成條件

- [ ] Plan Preparation、Phase 0–4 完成條件全部成立
- [ ] 兩個 consumer dogfood evidence（sanitized）已記錄，project-local 對應檔存在
- [ ] Open Questions 全部 resolved / deferred（附原因）
- [ ] Sanitization gate 通過；`git status` clean 且已 push
- [ ] Plan Completion Closure 執行完畢
