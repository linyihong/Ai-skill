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

**Status**: in-progress — Plan Preparation、Phase 0、**Phase 1 完成**（2026-11-23）；Phase 2/3 已改為「consumer 唯讀評估 + Ai-skill 自家套用」（2026-11-23 使用者決定）；**Phase 2 完成**；Q12 已裁決並寫回契約（2026-11-23）；第一次契約一致性審閱（2026-11-23）：Q12 寫回與正式契約 PASS，設計大綱 Lifecycle/Q8/Q9 與 Phase 3 Acceptance 已同步；**Phase 3 已開放**（2026-11-23 使用者決定）；**Phase 3 已關閉**（2026-11-23 使用者裁決）：S1 套用與審查完成，驗證缺口已移交 F1、F2、F6（見 evidence/phase3-application.md §5）。下一步 Phase 4 尚未開放。

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
5. **Lock taxonomy** — 八族工具治理鎖形狀（Q12 加入 safety；產品不變式鎖另歸一類）；鎖保護架構不變式，不保護檔案數量。
6. **Compensating-lock Review Rule** — 新增鎖前必答「能否從結構上消除？」

**明確不做**：共用 runner（Q5 deferred）、runtime route、自動 orchestrator（是否抽象化屬 Selection Responsibility）。

### Alternatives Considered

- A. 維持現狀，各專案各自整理：reject — 成果無法遷移，compensating lock 無框架可指認。
- B. 立即實作 Ai-skill CLI 通用 lint runner（Go）：reject（暫緩）— 證據 n=2、兩個 consumer 語言不同，runner 介面的 invariant 未知（Watch-Out wall 2 / 6）。
- C. 只修 Consumer A 的漏洞：reject — 失去跨專案學習，B4 照樣長。
- D. **Doc-only 契約 + 兩個 consumer 唯讀評估 + Ai-skill 自家套用，再依證據升級**（accept；2026-11-23 修訂：不修改外部 consumer repo）。
- E. 規定 agent「不准新增工具」：reject — 一次性與實驗工具有正當價值；要治理的是生命週期與發現。
- F. 自動 orchestrator 偵測新工具並自動搬層：reject — 越過 Selection Responsibility，抽象化決策失去證據。
- G. 名稱縮小為 Script Governance：reject — 真正的對象是工具資產的生命週期（DB 查詢、API 測試、OCR / ASR、部署驗證工具同樣適用）。

### Why Not an ADR Yet

概念判定準則、discovery 是否需機械化、lifecycle 雙維度是否完備、lock taxonomy 是否完備、共用 runner 是否值得，全部未驗證；且有更輕的 promotion target（cross-cutting pilot contract）。

### ADR Promotion Criteria（completed 時驗證）

- [ ] foundational + cross-session + cross-project + expensive-to-reverse + explains-why 全中
- [ ] 兩個 consumer 唯讀評估與 Ai-skill 自家套用皆證實 contract 可行，且未需修改 invariant 核心
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

**Doc-only trial，不接入 runtime。** 不新增 `route.*`、不 project 到 `runtime.db`、不新增 Ai-skill commit-msg validator。Phase 3 若在 Ai-skill 自己的 hooks 調整既有 validator，屬於套用切片內的改動，不等於本契約接入 runtime；consumer repo 不做任何機械強制。

- 未來接入：Phase 4 決策；entry condition = Phase 2、Phase 3 acceptance 通過，且符合 cross-cutting 三案例門檻（唯讀評估是否計入見 Q11）。
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
- [ ] **Q11 唯讀評估是否計入三案例門檻**（2026-11-23 範圍修訂後新增）：consumer 只做唯讀評估，證明的是「契約能判斷」，不是「照契約改得動」。Phase 4 裁決它們算不算收斂案例；Ai-skill 自家套用算一個實際套用案例。
- [x] **Q12 Phase 2 契約缺口裁決** — **decided (2026-11-23, user)**，已寫回契約：
  - G1 **Accept**：新增 safety 族（[`lock-taxonomy.md`](../../../workflow/cross-cutting/project-tooling/lock-taxonomy.md)）
  - G2 **Accept**：工具治理鎖與產品不變式鎖分開登記與計數；兩者皆適用 Review Rule
  - G3 **Accept with clarification**：forbidden 涵蓋常見寫法；反向樣本雙向——違規變體抓得到、合法相似寫法不誤判
  - G4 **Revise**：「到期 + 未引用 + 未登記」不足以單獨證明可刪；機械清理僅限指定 Research 區域、排除無法可靠確認使用情況者、須先隔離並有可恢復的保留期（[`lifecycle-contract.md`](../../../workflow/cross-cutting/project-tooling/lifecycle-contract.md) §機械清理的邊界）
  - G5 **Accept with clarification**：文件連結須同時說明新增原因與既有能力為何不適用（[`discovery-contract.md`](../../../workflow/cross-cutting/project-tooling/discovery-contract.md)）
  - G6 **Accept**：鎖的 reuse evidence 依同一不變式 / 能力計數，不依鎖檔案數量
  - A-1 **Accept with clarification**：點擊座標與隨目標調整的參數屬 Parameters；演算法幾何可留在具名、集中的 Logic（[`role-contract.md`](../../../workflow/cross-cutting/project-tooling/role-contract.md) §判定細則）
  - A-2 **Accept with clarification**：單純依序呼叫屬 Entry；含條件決策、快取判斷或範圍選擇屬 Logic

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
| Q11（範圍修訂後新增） | still-open | Phase 4 裁決 |
| Q12（Phase 2 新增） | decided（2026-11-23, user） | 已寫回契約 |

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
- [x] `lock-taxonomy.md`（七族鎖形狀，Q12 後為八族、compensating-lock Review Rule、check architecture 概念）
- [x] `workflow/cross-cutting/README.md` 加入連結；linked-updates 檢查（另同步 `workflow/README.md` cross-cutting 列；不新增 route，故 routing registry / 根 README 不適用）

**完成條件**：6 檔各 ≤150 行；Q8、Q9 resolved；Q10 已裁決並寫入 Role contract（含違規判定與跨語言原則）；Q2a–c resolved 或 deferred（附原因）；無 consumer 識別資訊（sanitization gate 通過）；不新增 route、runner 或 validator。

**Phase 1 交付（2026-11-23）**：6 檔已建於 [`workflow/cross-cutting/project-tooling/`](../../../workflow/cross-cutting/project-tooling/README.md)（56–72 行，皆 ≤150）；Q2b、Q2c、Q8、Q9 皆 accepted（Q9 依審閱修訂後通過）；**Phase 1 完成（2026-11-23，使用者最終 sign-off）**；sanitization 通過；無 route / runner / validator。

**Phase 2/3 必須等 Phase 1 完成才開始**（先確認契約正確，再用真實案例驗證）。— Phase 1 已完成。

> **2026-11-23 範圍修訂（使用者決定）**：Ai-skill 治理計畫**不修改外部 consumer repo**。Consumer 只做唯讀評估、當作經驗參考；契約的實際套用在 Ai-skill 自己的工具樹進行。原「在 consumer repo 執行」的 Phase 2/3 定義作廢。

## Phase 2 — 兩個 consumer 的唯讀評估

**唯讀**：不修改任何 consumer repo（tracked 內容、hook、docs 都不動）；只讀取並用契約判斷。具體檔名留各 consumer 的 project-local `.ai-skill/dogfood/` 對照檔。

- [x] **Role**：對兩個 consumer 的入口與共用模組套 [`role-contract.md`](../../../workflow/cross-cutting/project-tooling/role-contract.md)，違規一律用「`<placement>` 的 `<role>` 放在 `<實際位置>`」句型記錄
- [x] **Concept / Discovery**：以 [`growth-rule.md`](../../../workflow/cross-cutting/project-tooling/growth-rule.md) 判定同概念群；檢查既有工具是否有 creation justification、是否有 capability catalog 可供 discovery
- [x] **Placement × State × Ownership**：依 [`lifecycle-contract.md`](../../../workflow/cross-cutting/project-tooling/lifecycle-contract.md) 標示；記錄 ownership violation、逾期 experimental、殘留 scratch
- [x] **Lock review**：對現有鎖依 [`lock-taxonomy.md`](../../../workflow/cross-cutting/project-tooling/lock-taxonomy.md) 分族、填四欄形狀（scope / forbidden / allowlist / intent），逐把回答 compensating-lock Review Rule；記錄 scope 盲區與 pattern 漏洞
- [x] **建議改動清單（不執行）**：每項附對應的違規句型或契約條款
- [x] **契約缺口**：判不出、判斷歧義、或需改契約的地方，各附建議的契約修改

**Acceptance**：(1) consumer repo 沒有任何由本 phase 造成的變更；(2) 每個被評估的工具群都有可判定的結論，或明確記錄為歧義；(3) 契約缺口逐條列出並附建議修改；(4) 證據去敏後寫入 `evidence/`。

**Phase 2 結果（2026-11-23）**：完成，見 [`evidence/phase2-readonly-evaluation.md`](evidence/phase2-readonly-evaluation.md)。(1) consumer repo 由本 phase 造成的變更為 0；(2) Role / Concept / Placement × State / Ownership 皆可判定，僅 A-1、A-2 兩處歧義；(3) 契約缺口 G1–G6 已列並附建議修改，Q12 裁決後已寫回契約（G4 依裁決修訂）；(4) 證據已去敏。無需推翻的 invariant。

## Phase 3 — Ai-skill 自家套用

在 Ai-skill 自己的工具樹（`scripts/`：Go CLI、git-hooks 與其 validator）實際套用契約。Ai-skill 是本契約真正的使用者，改動走 Ai-skill 一般的 writeback 流程。

- [x] **Phase 3.0 範圍盤點（唯讀）**（2026-11-23，見 [`evidence/phase3-0-scope-inventory.md`](evidence/phase3-0-scope-inventory.md)；切片裁決 S1 = hook pipeline，見下方 §套用切片）：列出入口、validator（鎖）、共用模組，決定本 phase 的套用切片；確認 Go-first 政策與 committed `bin/` 重建要求；回寫範圍後才動手
### 套用切片（2026-11-23 使用者裁決：S1 + 五點邊界）

- **Catalog**：涵蓋全部 18 個能力（16 個頂層命令與 2 個獨立 binary）。
- **Role**：只檢查 `.githooks/*`、`scripts/git-hooks/*`、`cmd/ai-skill`、`app.go`、hooks 編排。
- **Lock**：全部的鎖都要分類登記；**完整審查只做工具治理鎖**。產品不變式鎖在 evidence 明列為「未審、後續處理」，不算已完成審查，因為契約仍要求它們有 intent、allowlist 退場條件與 Review Rule。
- **延後**：大型檔案拆分、registry schema 變更（四欄只寫在 evidence）、KGE 呼叫鏈重構。
- **O6**：視為既有架構的盤點結果，不承諾「只補 scope registry 就完整」。

- [x] **Bin 流程預查（第一次修改 Go source 前）**：依現行 hook 實作（`validateReleaseBinaryParityStaged`）與建置文件確認提交要求。若要求同一 commit，就照此規劃，不沿用過去拆 commit 的慣例。
- [x] **Capability catalog**：在 Ai-skill 工具樹 README 建立能力表（capability / entry / placement / state / owner / ttl），共 18 列。
- [x] **Lock 對帳**：建立 hook 執行點 → executor symbol → registry entry 的對照表，寫入 evidence。46 次 validator 執行和 44 個 registry entry 的計數單位不同，數字不一致本身不代表缺漏。以對帳結果確定工具治理鎖的邊界。
- [x] **Role**：薄入口判定與「重複入口」**分開判斷**。兩套外殼可能各自符合薄入口；O2 的一行轉呼叫也不足以直接判定違規。先確認各自的責任與呼叫需求，再決定是否修正；違規依句型記錄並修正（限本切片）。
- [x] **Lock review**：對工具治理鎖填四欄形狀並套用 compensating-lock Review Rule；能從結構消除的列為改動或後續 plan，不在本 phase 擴大實作。
- [x] **O4 契約問題**：記錄各 `[skip-*]` 標記的作用範圍、授權方式與失效時點。單次 skip 的有效期間和 allowlist 的退場條件是不同問題，暫不把 28 種標記判為違規；是否修訂契約留到 Phase 4。
- [x] Phase 2 的契約缺口若已裁決，先回寫契約再套用（Q12 已於 cdd835c4 回寫）

結果：[`evidence/phase3-application.md`](evidence/phase3-application.md)（2026-11-23）。切片內無 Role 違規。7 把工具治理鎖都已審查，四欄與 Review Rule 完整。樣本狀態分三類：T2、T5、T6 雙向通過；T1、T3、T4 發現缺口，其中 T4 的 executor 方向尚未驗證；T7 不適用雙向樣本。**完成審查不等於驗證通過。** **結論（2026-11-23 使用者裁決）：S1 套用與審查完成，驗證缺口已移交**——T4 executor 方向與 L1 → F1，T3 → F2，T1 → F6；本 phase 不補樣本。`evidencecandidate` 的 TTL 已設為 2026-12-23。後續工作 F1–F8 只記錄，F1 優先並另開修復切片。

**Acceptance**：

1. catalog 存在且涵蓋切片內所有 capability；切片內入口無 Role 違規。
2. 每把被審的鎖都有完整四欄形狀（`scope / forbidden / allowlist / intent`），`forbidden` 涵蓋同一違規的常見寫法。
3. Pattern-based 鎖的雙向樣本狀態逐把標示（通過 / 缺口 / 尚未驗證，依實際跑過的 test 判定）；未通過的方向必須列出缺口內容並移交具名的後續工作，本 phase 不要求補齊樣本。非 pattern-based 鎖不強制套用不適合的測試形式。（2026-11-23 使用者裁決修訂；原條件為「有雙向樣本測試」）
4. 每把被審的鎖都有 Compensating-lock Review Rule 結論。
5. 全部的鎖都已分類登記（工具治理鎖 / 產品不變式鎖分開計數），並附 hook → executor → registry 對帳；「被審的鎖」指本切片的工具治理鎖，第 2–4 項只適用它們；產品不變式鎖明列「未審、後續處理」。
6. 切片以外、能從結構消除的問題只記錄為後續工作，不擴大實作。
7. 所有改動通過 Ai-skill hooks 並已推送。

依據：[`lock-taxonomy.md`](../../../workflow/cross-cutting/project-tooling/lock-taxonomy.md)（Q12 G2、G3、G6）。

**Phase 3 驗收結果（2026-11-23）**：1–7 項皆成立（第 3 項依修訂後條件）。T2、T5、T6 雙向通過；T1、T3、T4 的缺口已列並移交 F6、F2、F1；T7 不適用。驗證缺口不算已驗證，Phase 4 引用 Ai-skill 自家套用證據時須一併標示。

## Phase 4 — 證據裁決與升級決策

- [ ] 對每條契約 invariant 標記：兩個 consumer 唯讀評估與 Ai-skill 自家套用是否都成立、是否需修改
- [ ] 裁決 Q11（唯讀評估是否計入三案例門檻）
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
- [x] Phase 2/3 改為 consumer 唯讀評估 + Ai-skill 自家套用（2026-11-23）
- [x] Q12 契約缺口裁決（2026-11-23）
- [x] 契約一致性審閱後開放 Phase 3（2026-11-23）
- [x] Phase 3 套用切片範圍（2026-11-23：S1 + 五點邊界，見 §Phase 3 套用切片）
- [x] Phase 3 結案方式（2026-11-23：缺口移交後續工作並修訂驗收第 3 項，不在本 phase 補樣本）

## 與其他 plans 的關係

- [`2026-06-08-2100-governance-pattern-library-extraction.md`](../2026-06-08-2100-governance-pattern-library-extraction.md)：lock taxonomy 可能成為其 pattern 候選；本 plan 不搶先登記。
- `archived/2026-10-04-0936-evidence-record-contract`：cross-cutting pilot contract 的形狀先例。
- Go-first script policy（Ai-skill 自身腳本）：僅影響 Q5 的 runner 實作語言，不約束 consumer。

## 完成條件

- [ ] Plan Preparation、Phase 0–4 完成條件全部成立
- [ ] 兩個 consumer 唯讀評估 evidence（sanitized）與 Ai-skill 自家套用 evidence 已記錄，project-local 對應檔存在
- [ ] Open Questions 全部 resolved / deferred（附原因）
- [ ] Sanitization gate 通過；`git status` clean 且已 push
- [ ] Plan Completion Closure 執行完畢
