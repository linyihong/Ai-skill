---
id: 2026-11-23-0834-project-tooling-script-governance
plan_kind: main
status: draft
owner: workflow
owner_layer: workflow
created: 2026-11-23
parent: null
---

# Project Tooling Script Governance（專案腳本收斂框架）

**Status**: draft — 待使用者審閱。尚未進入 Phase 0；不改任何 consumer repo。

**Glossary Impact**: yes（候選，Phase 1 前不登記）— `tooling_layer_contract`、`promotion_ladder`、`compensating_lock`、`registry_driven_entry`。

## 一句話目標

讓專案的分析／驗證腳本從「每做一次分析就多一支腳本」收斂成 **分層 + 參數化 + 機械鎖** 的可演化工具面：新需求優先「改設定 / 加參數」，其次「擴充共用模組」，最後才「新增腳本」，並用可證偽的升級觸發條件與鎖，防止回流。

## Decision Rationale

### Problem & Why Now

兩個性質不同的 consumer 專案都出現同一個根因：**沒有「何時必須抽共用 / 參數化」的規則，也沒有單一 registry 驅動的入口**，於是鎖只能逐事件加，腳本逐分析長。

**Consumer A — 逆向分析專案（Python + Frida，裝置自動化與資料擷取）**

已具備分層雛形與 24 個 pre-commit 機械鎖：

| 層 | 現況 |
| --- | --- |
| L0 薄入口 | `scripts/cli/<domain>/*.py`，約 15 行：`sys.path` + 呼叫 `main()` |
| L1 領域模組 | `lobby/`、`farm/`、`login/`、`slots/{capture,paytable,sprites,registry,lab}` |
| L2 共用基礎 | `shared/{device,ocr,scene,identity}` |
| L3 參數化 | 每目標 `capture-config.json` + `waitHook.strategy` enum（三種策略取代每目標 300+ 行複製檔） |
| 例外區 | `<target>/<id>/`（單目標邏輯）、`frida/research/`（實驗，有 cleanup 鎖） |

鎖可歸為五族：layout、boundary（禁止重造 device I/O）、completeness、drift、lifecycle。

觀察到的缺口（observed，Phase 0 verify）：

| # | 缺口 | 量化 |
| --- | --- | --- |
| A1 | L0 薄入口是慣例不是鎖 | 3 支 `cli/` 入口非薄（31 / 132 / 260 行）；最大者自帶 `tap()`、裝置序號預設值、模組級座標常數 |
| A2 | boundary 鎖 scope 只覆蓋單目標目錄 | 放進 `cli/` 即繞過 device I/O 鎖 |
| A3 | 裝置序號預設值分散 | 7 個檔案各自 `os.environ.get(...)` 帶預設，未統一走 shared resolver |
| A4 | 無升級觸發條件 | README 只寫「禁止」，未寫「第 N 次出現就必須抽」 |
| A5 | 鎖本身逐檔蔓延 | 24 支獨立 check 腳本 + 手寫 pre-commit 清單；新增一鎖 = 新增一檔 + 改 hook |

**Consumer B — 業務交付 monorepo（bash + PowerShell + Node 驗證閘）**

邏輯已抽到 `lib/check_*.mjs`（好的部分），但：

| # | 缺口 | 量化 |
| --- | --- | --- |
| B1 | 入口 wrapper 大量重複 | `scripts/` 根目錄 134 個項目；約 50 組 `verify-*.sh` + `verify-*.ps1`，多數只是 `node lib/check_x.mjs --root` 加 echo（~100 個近乎相同檔） |
| B2 | **compensating lock** | `verify-scripts-parity` 鎖存在的理由是 wrapper 被複製成兩份；鎖在補償結構重複，而非消除它 |
| B3 | `lib/` 平鋪 | 150 檔混放 `check_*`／`test_*`／deploy／scope helper，無子分層 |
| B4 | **一次性分析腳本蔓延** | `performance/` 8 支 per-case 腳本（27–154 行），各自重寫 fixture 斷言、計時、報告；**零個專案內 import**（只 import `node:*`） |

B4 與 Consumer A 的 A1 是同一現象：每做一次分析就複製一支，不抽 harness、不參數化。

**Why now**：兩個 consumer 的鎖數量都在快速增加（A：24；B：~50 gate），每次新增都在擴大維護面；Consumer A 剛完成一輪分層整理，是把「整理後的形狀」抽成可重用契約的最佳時機（形狀新鮮、證據可直接量測）。

### Decision

建立一份 **tool-neutral、language-neutral** 的 Project Tooling Contract，先以 doc-only pilot reference contract 落在 `workflow/cross-cutting/project-tooling/`（比照 `traceable-evidence` 模式：非 stage、不註冊 route），由兩個 consumer 各做一個 dogfood slice 驗證，再決定是否升級。

契約四部分：

1. **Layer contract**：L0 entry（只解析參數 + 呼叫）→ L1 domain → L2 shared → L3 parameters/config（strategy enum、per-target config）；例外區（單目標、research）必須有明確 TTL 或升級出口。
2. **Promotion ladder**（每階可退回，觸發條件可證偽）：
   `research（TTL）→ target/<id>（單目標）→ domain 模組 → shared 模組 → 參數 / 設定`
   - 同一段邏輯出現在第 2 個 target → 必須抽到 domain。
   - 第 2 個 domain 使用 → 必須抽到 shared。
   - 兩個 target 只差數值不差流程 → 必須改成 config / strategy enum，而非新腳本。
3. **Lock taxonomy**：layout / thin-entry / boundary / lifecycle / completeness / drift / parity。每族一個「填 scope + forbidden pattern + allowlist」的形狀；allowlist 必須附退場條件。
4. **Compensating-lock test**：新增鎖前先問「這個鎖防的重複，能不能從結構上消除？」能消除就先消除，鎖只守剩下的不變式（B2 是反例）。

**Registry-driven entry**（單一入口 + 宣告式 check / task 清單）列為 pattern，不列為必須；是否做成 Ai-skill CLI 子指令留到 Phase 4 依證據決定。

### Alternatives Considered

- A. 維持現狀，各專案各自整理：reject — Consumer A 的整理成果無法遷移到 B，B 的 compensating lock 問題沒有框架可指認。
- B. 立即實作 Ai-skill CLI 通用 lint runner（Go）：reject（暫緩）— 證據 n=2、且兩個 consumer 語言不同（Python / Node），runner 介面還不知道哪些欄位是 invariant；過早實作即 Watch-Out wall 2。
- C. 只修 Consumer A 的漏洞：reject — 修得掉 A1/A2，但失去跨專案學習，B4 會照樣長。
- D. **Doc-only 契約 + 兩個 consumer 各一 slice dogfood，再依證據升級**（accept）。

### Why Not an ADR Yet

升級觸發門檻（2 vs 3）、lock taxonomy 是否完備、registry-driven entry 是否值得共用 runtime，全部未驗證；且有更輕的 promotion target（cross-cutting pilot contract）可先承載。

### ADR Promotion Criteria（completed 時驗證）

- [ ] foundational + cross-session + cross-project + expensive-to-reverse + explains-why 全中
- [ ] 兩個 consumer dogfood 皆證實 contract 可行，且未需修改 invariant 核心
- [ ] Open Questions 全解
- [ ] 沒有更輕的 promotion target 適用（per ADR-007）
- [ ] 第三個專案採用時不需修改 contract（真實使用 evidence）

### Consequences（預期）

#### 正面
- 新分析需求有明確的「先參數、再共用、最後新檔」決策順序。
- 鎖從逐事件新增轉為依 taxonomy 套形狀；可辨識 compensating lock。

#### 負面
- 多一份 cross-cutting 契約要維護。
- Consumer B 若選擇收斂 wrapper，會影響 hook / CI 入口（Phase 3 刻意不先做）。

#### 風險
- 契約膨脹成 BPMN 式流程（Watch-Out wall 2）— 緩解：Phase 1 契約限 4 檔、每檔 ≤150 行；只寫 invariant 與觸發條件，不寫語言專屬實作。
- 把 consumer 專案細節帶進共用層 — 緩解：見 §Sanitization。

### Watch-Out List citation

[`architecture/ai-native-cognitive-ecosystem-system.md`](../../architecture/ai-native-cognitive-ecosystem-system.md) §Watch-Out List — **Wall 2 Workflow inflation**（契約只定 invariant，不定逐步流程）；**Wall 6 Optimization hallucination**（不在 n=2 時做通用 runner）。

## Runtime Execution Path

**Doc-only trial，不接入 runtime。** 本 plan 不新增 `route.*`、不 project 到 `runtime.db`、不新增 Ai-skill commit-msg validator。機械強制發生在 **consumer repo 自己的 hook**（Phase 2/3），不屬於 Ai-skill runtime integration，本 plan 不聲稱已完成 runtime integration。

- 未來接入 phase：Phase 4 決策；entry condition = 兩個 consumer dogfood PASS 且 Open Questions Q1–Q4 resolved。
- 候選接入形式：discovery signal（偵測「新增 scripts/ 檔案」時載入本契約）或 `init-project` 範本；需另開 plan 並附 Per-surface consumer 表。
- Graduation deadline：2027-02-28（逾期未接入 → 維持 doc-only pilot 並在 plans/README 說明）。

### Per-surface consumer 表

不適用 — 本 plan 不新增 generated surface、route 或 validator。

## Sanitization

共用層只用 Consumer A / Consumer B 抽象描述（領域 + 語言），不寫 repo 名、裝置序號、座標、本機路徑或 SHA。具體對應（repo、檔案路徑、gate 結果）放各 consumer 的 project-local、非 git 的 `.ai-skill/dogfood/2026-11-23-0834-project-tooling-script-governance.md`。是否把兩個 consumer 名稱加入 `.ai-skill-project.yaml` `private_entities` 由 Q6 決定。

## Open Questions

- [ ] **Q1 放置位置**：`workflow/cross-cutting/project-tooling/` vs `workflow/software-delivery/` 子能力？（傾向 cross-cutting：Consumer A 不是交付專案）
- [ ] **Q2 升級門檻**：「第 2 次出現就抽」是否太激進？分析類一次性腳本是否允許 research 區 TTL 內不抽？
- [ ] **Q3 Compensating-lock test**：寫成規則（新增鎖必答）還是 heuristic（review 提示）？
- [ ] **Q4 Consumer B wrapper 收斂**：保留 sh/ps1 成對 + parity 鎖，或改為單一 Node 入口 + check registry（Node 已是必要依賴）？屬 B 專案決策，本 plan 只提供分析。
- [ ] **Q5 Runner 語言**：若 Phase 4 決定做共用 runner，是 Ai-skill Go CLI（Go-first policy）讀 consumer YAML，還是各 consumer 原生語言？
- [ ] **Q6 Sanitization denylist**：兩個 consumer 名稱是否登記到 `private_entities`？

## Phase 0 — Pre-Build Interrogation + Preflight

### Phase 0.0 — Open Questions 核對（公版，必填）

逐條核對本 plan §Open Questions，標記處置並回寫：

- [ ] 已讀本 plan §Open Questions 全部條目
- [ ] 對每條標記 `resolved`（附 Phase 0 證據）/ `still-open` / `deferred`（附原因）
- [ ] `resolved` 的條目已同步勾選 / 附註於 §Open Questions
- [ ] 若盤點新發現問題，已加入 §Open Questions

| Open Question | 處置 | 證據 / 原因 |
|---|---|---|
| Q1–Q6 | pending | — |

### Phase 0.1 — Baseline 量測（兩個 consumer，可重跑的指令記在 project-local dogfood 檔）

| 指標 | 意義 |
| --- | --- |
| scripts 根目錄項目數 | layout 蔓延程度 |
| 非薄 L0 入口數（> 25 行或含 I/O） | thin-entry 缺口 |
| 近重複 wrapper 數 | 結構重複 / compensating lock 規模 |
| per-case 腳本的專案內 import 比例 | 一次性腳本是否共用 harness |
| 機械鎖數、每鎖 scope | 鎖蔓延與盲區 |

把 A1–A5、B1–B4 從 `observed` 升為 `verified` 或 `refuted`（verify 對 HEAD）。

### Phase 0.2 — Framework discovery（避免雙份 source-of-truth）

- [ ] 讀 [`workflow/cross-cutting/README.md`](../../workflow/cross-cutting/README.md)、[`traceable-evidence/`](../../workflow/cross-cutting/traceable-evidence/README.md)（pilot contract 形狀參考）
- [ ] 讀 [`2026-06-08-2100-governance-pattern-library-extraction.md`](2026-06-08-2100-governance-pattern-library-extraction.md) — 確認 lock taxonomy 不與 pattern library 重疊
- [ ] 讀 software-delivery framework domain model（archived）§Automation deferred — 確認不重複
- [ ] [`workflow/software-delivery/requirements/pre-build-interrogation.md`](../../workflow/software-delivery/requirements/pre-build-interrogation.md) 完成記錄
- [ ] Architecture Compatibility Preflight 最低記錄格式填寫完畢 → Decision: proceed / revise

**Phase 0 完成條件**：baseline 表、缺口 verified/refuted、preflight Decision 已回寫本 plan。

## Phase 1 — 契約草稿（Ai-skill，doc-only）

- [ ] `workflow/cross-cutting/project-tooling/README.md`（索引 + 適用範圍 + 何時載入）
- [ ] `layer-contract.md`（L0–L3 + 例外區 TTL）
- [ ] `promotion-ladder.md`（觸發條件 + 可證偽 predicate）
- [ ] `lock-taxonomy.md`（七族鎖形狀 + compensating-lock test）
- [ ] `workflow/cross-cutting/README.md` 加入連結；linked-updates 檢查

**完成條件**：4 檔各 ≤150 行；無 consumer 識別資訊（sanitization gate 通過）；不新增 route。

## Phase 2 — Consumer A dogfood（在 consumer repo 執行，非本 repo）

Slice：補 thin-entry 漏洞。

- [ ] 新增 thin-entry 鎖：L0 入口只允許 import + `main()`；含 subprocess / 自寫 tap / 座標常數即擋
- [ ] 擴大 device I/O boundary 鎖 scope 至 L0
- [ ] 3 支非薄入口：先放 allowlist（附退場條件）→ 邏輯移到單目標區並改走 shared device helper → 移出 allowlist
- [ ] 裝置序號預設值收斂到 shared resolver

**Acceptance**：baseline「非薄 L0 入口數」= 0；allowlist 為空；鎖有反向測試（放一支違規入口會被擋）。

## Phase 3 — Consumer B dogfood（在 consumer repo 執行）

Slice：一次性分析腳本 → 共用 harness（B4）。wrapper 收斂（B1/B2）不在本 phase，等 Q4。

- [ ] 從 `performance/` 抽出共用 harness（fixture 驗證、計時、報告輸出）
- [ ] 先遷移 2 支 per-case 腳本，per-case 檔只剩 case 參數 / 宣告
- [ ] 評估剩餘腳本是否可改為 config-only case（L3）
- [ ] 視需要加一條 lock：per-case 腳本不得直接重寫 harness 職責

**Acceptance**：遷移後腳本行數下降且輸出報告與遷移前等價（同一 fixture 對照）；專案內 import 比例 > 0。

## Phase 4 — 證據裁決與升級決策

- [ ] 對每條契約 invariant 標記：兩個 consumer 是否都成立、是否需修改
- [ ] 決定 Q5（是否做共用 runner）與 runtime 接入（見 §Runtime Execution Path）
- [ ] 回寫契約；更新 `plans/README.md` 狀態
- [ ] 執行 Plan Completion Closure

## Stakeholder 同意項目

- [ ] 契約放置位置（Q1）
- [ ] 升級門檻（Q2）
- [ ] Phase 2 / Phase 3 dogfood slice 範圍（各在 consumer repo 動到什麼）
- [ ] Consumer B wrapper 收斂是否另案處理（Q4）

## 與其他 plans 的關係

- [`2026-06-08-2100-governance-pattern-library-extraction.md`](2026-06-08-2100-governance-pattern-library-extraction.md)：lock taxonomy 可能成為其 pattern 候選；本 plan 不搶先登記。
- `archived/2026-10-04-0936-evidence-record-contract`：cross-cutting pilot contract 的形狀先例。
- Go-first script policy（Ai-skill 自身腳本）：僅影響 Q5 的 runner 實作語言，不約束 consumer。

## 完成條件

- [ ] Phase 0–4 完成條件全部成立
- [ ] 兩個 consumer dogfood evidence（sanitized）已記錄，project-local 對應檔存在
- [ ] Open Questions 全部 resolved / deferred（附原因）
- [ ] Sanitization gate 通過；`git status` clean 且已 push
- [ ] Plan Completion Closure 執行完畢
