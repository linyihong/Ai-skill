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

**Glossary Impact**: yes（候選，Phase 1 前不登記）— `tooling_layer_contract`、`tool_discovery`、`creation_justification`、`tooling_lifecycle`、`concept_check`、`capability_catalog`、`reuse_evidence`、`promotion_ladder`、`compensating_lock`。

## 一句話目標

讓專案的分析／驗證腳本從「每做一次分析就多一支腳本」收斂成 **分層 + 參數化 + 機械鎖** 的可演化工具面。**新腳本可以長，但只在確認沒有同概念功能時才長**：同概念 → 加參數，或抽出共用邏輯後在同一功能分類下長一個 sibling；新概念 → 才開新分類、新腳本。並用鎖防止回流。

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

**更深一層的根因（agent 工作習慣）**：AI 執行能力越強，越容易為了完成當下任務創造新工具，卻不判斷工具是否已存在、是否該成為共用能力、任務結束後如何處置。累積路徑：

1. 遇到問題，寫一支一次性腳本；
2. 下次遇到相似問題，沒找到或沒辨識出舊腳本，再寫一支；
3. 新腳本變成半永久工具，沒有正式入口、文件或維護責任；
4. 為防再犯，加一個 lint / pre-commit；
5. 腳本、wrapper、validator、allowlist、hook 一起增長。

第 4 步是關鍵：**正在用更多機械鎖，管理原本可以靠更好的工具架構消除的重複**（compensating lock）。

**Why now**：兩個 consumer 的鎖數量都在快速增加（A：24；B：~50 gate），每次新增都在擴大維護面；Consumer A 剛完成一輪分層整理，是把「整理後的形狀」抽成可重用契約的最佳時機（形狀新鮮、證據可直接量測）。

### Decision

建立一份 **tool-neutral、language-neutral** 的 Project Tooling Contract，先以 doc-only pilot reference contract 落在 `workflow/cross-cutting/project-tooling/`（比照 `traceable-evidence` 模式：非 stage、不註冊 route），由兩個 consumer 各做一個 dogfood slice 驗證，再決定是否升級。

**治理對象是腳本的生命週期，不是禁止新增腳本。** 一次性腳本在逆向、效能、資料擷取中有價值；要管的是：何時可建、放哪、何時重用、何時升級、何時淘汰。

契約六部分（每份文件 ≤150 行，不合併成大文件）：

1. **Layer contract**：L0 entry（只解析參數 + 呼叫）→ L1 domain → L2 shared → L3 parameters/config（strategy enum、per-target config）；例外區（單目標、research）必須有明確 TTL 或升級出口。
2. **Discovery contract（先找再建）** — 解決「AI 為什麼一直重新創造腳本」：建立新腳本前，agent 必須完成
   1. 搜尋專案既有 scripts 與 capability catalog；
   2. 搜尋 domain / shared 是否已有相同能力；
   3. 確認能否以參數解決；
   4. 確認能否擴充既有 harness；
   5. 仍需新增時，說明 **creation justification** 與生命週期狀態。

   合理新增的理由只有四種：新的執行責任、新的資料來源或介面、現有抽象無法合理涵蓋、實驗需要隔離。
   **輕量原則**：這是執行前的工具發現與選擇行為，不是表單；justification 一兩句即可，寫在檔頭或 commit message。候選升為 workflow invariant（behavioral，見 Q7）。
3. **Lifecycle contract（保留與淘汰）**：

   | 層 | 狀態 | 存在條件 / 處置 |
   | --- | --- | --- |
   | Scratch | `temporary` | 只為當前任務；固定暫存位置，任務結束清理 |
   | Research | `experimental` | 有 TTL、目的、輸入 / 輸出定義；到期重新評估（升級 / 延期 / 刪除） |
   | Target | `maintained` | 確認要保留、服務單一目標；有 owner |
   | Domain | `maintained` | 同領域可重用；共用模組 + 參數化入口 |
   | Shared | `maintained` | 跨領域基礎能力 |
   | 任一層 | `deprecated` | 禁止新增使用，等待移除；已無使用者即移除 |

   原則：**不是每支腳本都要升級，但每支被保留的腳本都要有存在理由。**「成功執行過」≠「值得永久保留」。
4. **Growth rule + Promotion ladder（重複與抽象）** — concept-first（使用者 2026-11-23 決定）：

   新增前的 **concept check**：查 capability catalog（每個功能分類的 README 能力表），回答「是否已有同概念的功能？」

   | 判定 | 允許的動作 | 禁止 |
   | --- | --- | --- |
   | **同概念、只差數值 / 目標** | 在既有腳本 **加參數** 或加 config / strategy enum 值 | 複製一支新腳本 |
   | **同概念、流程有差異** | 先把共同邏輯 **抽到該分類的共用模組**，再在 **同一功能分類下** 長一個 sibling（只放差異） | 在別的分類或根目錄另起爐灶；sibling 重寫共用邏輯 |
   | **新概念**（catalog 無對應） | 開新功能分類 + 新腳本，並登記到 catalog | 不登記 catalog |

   概念判定的最小準則（Phase 0/1 細化，見 Q2a）：**輸入對象、產出物、核心流程** 三者中若「產出物 + 核心流程」相同，即視為同概念；只差目標 ID、座標、門檻、站點等，屬參數。

   **Config-first ≠ Config-everything**（防止 script sprawl 變成 configuration sprawl）：

   | 差異型態 | 落點 |
   | --- | --- |
   | 流程相同、資料不同 | Config |
   | 流程部分不同、核心契約相同 | Strategy（例：多種 OCR 引擎共用輸入 / 輸出契約，各自實作） |
   | 執行責任不同 | 獨立模組或腳本（屬新概念） |
   | 實驗尚未穩定 | Research |

   **Reuse evidence**（決定「抽象化決策何時到期」，不是機械合併）：

   | 觸發事件 | 必須採取的行動 |
   | --- | --- |
   | 同概念第 2 次出現 | 進行抽象化評估（走上方 concept check 表） |
   | 同概念第 3 次出現 | 抽取共用能力，或寫下「保留重複」的理由 |
   | 第 2 個 domain 使用相同能力 | 評估是否上推 Shared |
   | 僅參數不同 | 評估 Config / Strategy |

   兩者的關係：concept check 決定「是不是同一件事、該怎麼長」；reuse evidence 決定「這個決策最晚何時必須做」。表面相同不代表未來變化方向相同，所以第 2 次觸發的是決策，不是強制合併。

   **Promotion ladder** 作為抽共用時的落點（每階可退回）：
   `scratch → research（TTL）→ target/<id> → domain 模組 → shared 模組`，數值差異另落 L3 config / strategy。
   共用邏輯抽到「涵蓋所有使用者的最低一層」即可，不預先上推到 shared。
5. **Lock taxonomy**：layout / thin-entry / boundary / lifecycle / completeness / drift / parity。每族一個「填 scope + forbidden pattern + allowlist」的形狀；allowlist 必須附退場條件。**鎖保護架構不變式，不保護檔案數量。**
6. **Compensating-lock test**（Review Rule，新增鎖時必答，見 Q3）：「這個鎖防的重複，能不能從結構上消除？」能消除就先消除，鎖只守剩下的不變式（B2 是反例）。

**Governance check architecture（概念，不實作）**：single check entry + check registry（要跑哪些檢查）+ scope registry（適用範圍）+ 統一的結果契約。各 consumer 用自己的語言與執行環境實作；統一的是治理概念、檢查描述與結果形狀，不是 runner。是否做成 Ai-skill CLI 子指令 deferred（Q5）。

**責任切分**（對齊 Ai-skill 既有的 selection / constraint / verification 分工）：

| 責任 | 內容 | 由誰承擔 |
| --- | --- | --- |
| Selection | 用既有工具、擴充、或新建；是否抽象化 | agent，留下 justification / 評估紀錄 |
| Constraint | 放置位置、依賴、生命週期、升級規則 | 專案契約 + 機械鎖（只擋已明確違反的不變式） |
| Verification | 無重複能力、無繞過共用層、無未處置的臨時產物 | 獨立驗證（review / verifier） |

**明確不做**：自動 orchestrator（例如看到新腳本就自動搬到 shared）。是否抽象化屬 selection responsibility，需保留決策與證據。

**設計目標**：新增一支腳本是有意識的架構決策，不是 AI 完成任務時的預設副作用；同時不讓 AI 在寫腳本前填一堆表格。

### Alternatives Considered

- A. 維持現狀，各專案各自整理：reject — Consumer A 的整理成果無法遷移到 B，B 的 compensating lock 問題沒有框架可指認。
- B. 立即實作 Ai-skill CLI 通用 lint runner（Go）：reject（暫緩）— 證據 n=2、且兩個 consumer 語言不同（Python / Node），runner 介面還不知道哪些欄位是 invariant；過早實作即 Watch-Out wall 2。
- C. 只修 Consumer A 的漏洞：reject — 修得掉 A1/A2，但失去跨專案學習，B4 會照樣長。
- D. **Doc-only 契約 + 兩個 consumer 各一 slice dogfood，再依證據升級**（accept）。
- E. 規定 agent「不准新增腳本」：reject — 一次性與實驗腳本有正當價值；要治理的是生命週期與發現，不是數量。
- F. 自動 orchestrator 偵測新腳本並自動搬層：reject — 越過 selection responsibility，抽象化決策失去證據。

### Why Not an ADR Yet

概念判定準則、discovery 是否需機械化、lifecycle 狀態是否完備、lock taxonomy 是否完備、registry-driven entry 是否值得共用 runtime，全部未驗證；且有更輕的 promotion target（cross-cutting pilot contract）可先承載。

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
- 契約膨脹成 BPMN 式流程（Watch-Out wall 2）— 緩解：Phase 1 + 1.5 契約限 6 檔、每檔 ≤150 行；只寫 invariant 與觸發條件，不寫語言專屬實作。
- 把 consumer 專案細節帶進共用層 — 緩解：見 §Sanitization。

### Watch-Out List citation

[`architecture/ai-native-cognitive-ecosystem-system.md`](../../architecture/ai-native-cognitive-ecosystem-system.md) §Watch-Out List — **Wall 2 Workflow inflation**（契約只定 invariant，不定逐步流程）；**Wall 6 Optimization hallucination**（不在 n=2 時做通用 runner）。

## Runtime Execution Path

**Doc-only trial，不接入 runtime。** 本 plan 不新增 `route.*`、不 project 到 `runtime.db`、不新增 Ai-skill commit-msg validator。機械強制發生在 **consumer repo 自己的 hook**（Phase 2/3），不屬於 Ai-skill runtime integration，本 plan 不聲稱已完成 runtime integration。

- 未來接入 phase：Phase 4 決策；entry condition = 兩個 consumer dogfood PASS 且 Q2a–c、Q7–Q9 resolved。
- 候選接入形式：discovery signal（偵測「新增 scripts/ 檔案」時載入本契約）或 `init-project` 範本；需另開 plan 並附 Per-surface consumer 表。
- Graduation deadline：2027-02-28（逾期未接入 → 維持 doc-only pilot 並在 plans/README 說明）。

### Per-surface consumer 表

不適用 — 本 plan 不新增 generated surface、route 或 validator。

## Sanitization

共用層只用 Consumer A / Consumer B 抽象描述（領域 + 語言），不寫 repo 名、裝置序號、座標、本機路徑或 SHA。具體對應（repo、檔案路徑、gate 結果）放各 consumer 的 project-local、非 git 的 `.ai-skill/dogfood/2026-11-23-0834-project-tooling-script-governance.md`。是否把兩個 consumer 名稱加入 `.ai-skill-project.yaml` `private_entities` 由 Q6 決定。

## Open Questions

- [x] **Q1 放置位置** — **decided (2026-11-23, user review)**：`workflow/cross-cutting/project-tooling/`。腳本治理跨逆向分析與業務交付，不綁 software-delivery。
- [x] **Q2 生長 / 升級門檻** — **decided (2026-11-23, user)**：concept-first（無同概念才長新腳本；同概念加參數或「抽共用 + 同分類 sibling」）。同日 user review 補充 reuse evidence：第 2 次 → 抽象化評估、第 3 次 → 抽共用或寫保留重複的理由。兩者合併見 §Decision 第 4 部分。
- [ ] **Q2a 概念判定準則**：「產出物 + 核心流程相同 = 同概念」是否足夠？邊界案例（同產出物、不同資料來源；同契約、不同實作 → Strategy）怎麼判？Phase 0 用兩個 consumer 的實際腳本逐支試判。
- [ ] **Q2b Capability catalog 形狀**：每個功能分類 README 的能力表即可，還是需要機器可讀清單（讓鎖可檢查「新檔是否已登記」）？
- [ ] **Q2c Research 區**：實驗腳本在 TTL 內是否豁免 concept check（只要求到期前登記或刪除）？
- [x] **Q3 Compensating-lock test** — **decided (2026-11-23, user review)**：新增鎖時必須回答的 Review Rule，不是可選 heuristic。
- [x] **Q4 Consumer B wrapper 收斂** — **decided (2026-11-23, user review)**：獨立裁決，不放進本 plan 的 dogfood，避免 wrapper 整理干擾一次性腳本治理的證據。
- [x] **Q5 Runner 語言** — **deferred (2026-11-23, user review)**：目前沒有證據證明需要共用 runner；重新評估時機＝Phase 4。
- [x] **Q6 Sanitization denylist** — **decided (2026-11-23, user review)**：登記必要的私有識別資訊；在 Phase 0 以獨立 commit 修改 `.ai-skill-project.yaml` `private_entities`（本輪只修 plan）。
- [ ] **Q7 Discovery 強制程度**：Discovery contract 先作 behavioral workflow invariant；是否、何時需要機械化（例：新增 scripts 檔時檢查是否有 justification / catalog 登記）？前提是不變成表單負擔。
- [ ] **Q8 Scratch 位置與清理**：temporary 腳本的固定暫存位置（專案內 ignored 目錄 vs session scratch）與清理責任由誰確認？
- [ ] **Q9 Ownership 表達**：`maintained` 工具的 owner 記在哪（catalog 欄位 / 檔頭 / CODEOWNERS）？

## Phase 0 — Pre-Build Interrogation + Preflight

### Phase 0.0 — Open Questions 核對（公版，必填）

逐條核對本 plan §Open Questions，標記處置並回寫：

- [ ] 已讀本 plan §Open Questions 全部條目
- [ ] 對每條標記 `resolved`（附 Phase 0 證據）/ `still-open` / `deferred`（附原因）
- [ ] `resolved` 的條目已同步勾選 / 附註於 §Open Questions
- [ ] 若盤點新發現問題，已加入 §Open Questions

| Open Question | 處置 | 證據 / 原因 |
|---|---|---|
| Q1、Q2、Q3、Q4、Q6 | decided（2026-11-23） | 使用者決定；Q6 的 denylist 修改在 Phase 0 執行 |
| Q5 | deferred | 重新評估時機＝Phase 4 |
| Q2a–c、Q7–Q9 | pending | — |

### Phase 0.1 — Baseline 量測（兩個 consumer，可重跑的指令記在 project-local dogfood 檔）

| 指標 | 意義 |
| --- | --- |
| scripts 根目錄項目數 | layout 蔓延程度 |
| 非薄 L0 入口數（> 25 行或含 I/O） | thin-entry 缺口 |
| 近重複 wrapper 數 | 結構重複 / compensating lock 規模 |
| per-case 腳本的專案內 import 比例 | 一次性腳本是否共用 harness |
| 機械鎖數、每鎖 scope | 鎖蔓延與盲區 |

把 A1–A5、B1–B4 從 `observed` 升為 `verified` 或 `refuted`（verify 對 HEAD）。

**Concept 試判**：對兩個 consumer 的既有腳本逐支套 concept check，產出「概念 → 腳本清單」表；同概念多支者即為 Phase 2/3 收斂候選，並用來回答 Q2a（判不出來的邊界案例記錄下來）。

### Phase 0.2 — Framework discovery（避免雙份 source-of-truth）

- [ ] 讀 [`workflow/cross-cutting/README.md`](../../workflow/cross-cutting/README.md)、[`traceable-evidence/`](../../workflow/cross-cutting/traceable-evidence/README.md)（pilot contract 形狀參考）
- [ ] 讀 [`2026-06-08-2100-governance-pattern-library-extraction.md`](2026-06-08-2100-governance-pattern-library-extraction.md) — 確認 lock taxonomy 不與 pattern library 重疊
- [ ] 讀 software-delivery framework domain model（archived）§Automation deferred — 確認不重複
- [ ] [`workflow/software-delivery/requirements/pre-build-interrogation.md`](../../workflow/software-delivery/requirements/pre-build-interrogation.md) 完成記錄
- [ ] Architecture Compatibility Preflight 最低記錄格式填寫完畢 → Decision: proceed / revise

- [ ] Q6：`.ai-skill-project.yaml` `private_entities` 登記兩個 consumer 的識別資訊（獨立 commit）

Phase 0 只讀、只量測，不修改任何 consumer repo。

**Phase 0 完成條件**：baseline 表、缺口 verified/refuted、concept 試判表、preflight Decision 已回寫本 plan。

## Phase 1 — 結構契約草稿（Ai-skill，doc-only）

- [ ] `workflow/cross-cutting/project-tooling/README.md`（索引 + 適用範圍 + 何時載入 + 設計目標：腳本是有意識的架構決策，不是預設副作用）
- [ ] `layer-contract.md`（L0–L3 + 例外區出口）
- [ ] `growth-rule.md`（concept check + config-first ≠ config-everything + reuse evidence + promotion ladder 落點）
- [ ] `lock-taxonomy.md`（七族鎖形狀 + compensating-lock Review Rule + governance check architecture 概念）
- [ ] `workflow/cross-cutting/README.md` 加入連結；linked-updates 檢查

## Phase 1.5 — AI Tool Creation Governance（Ai-skill，doc-only）

Phase 2/3 dogfood **必須等本 phase 完成**才開始：先確認契約正確，再用真實案例驗證。

- [ ] `discovery-contract.md`（先找再建 5 步 + 4 種合理新增理由 + creation justification 的輕量寫法 + selection / constraint / verification 責任切分 + 不做自動 orchestrator）
- [ ] `lifecycle-contract.md`（Scratch / Research / Target / Domain / Shared × temporary / experimental / maintained / deprecated；TTL 到期處置；deprecation 與移除條件；ownership）
- [ ] README 索引補上兩檔；確認 Discovery 寫成 behavioral invariant，不是填表流程

**Phase 1 + 1.5 完成條件**：6 檔各 ≤150 行；無 consumer 識別資訊（sanitization gate 通過）；不新增 route、runner 或 validator；Q2a–c、Q7–Q9 各有處置（resolved / deferred 附原因）。

## Phase 2 — Consumer A dogfood（在 consumer repo 執行，非本 repo）

Slice：補 thin-entry 漏洞。

- [ ] 新增 thin-entry 鎖：L0 入口只允許 import + `main()`；含 subprocess / 自寫 tap / 座標常數即擋
- [ ] 擴大 device I/O boundary 鎖 scope 至 L0
- [ ] 3 支非薄入口：先放 allowlist（附退場條件）→ 邏輯移到單目標區並改走 shared device helper → 移出 allowlist
- [ ] 裝置序號預設值收斂到 shared resolver

**Acceptance**：baseline「非薄 L0 入口數」= 0；allowlist 為空；鎖有反向測試（放一支違規入口會被擋）。

## Phase 3 — Consumer B dogfood（在 consumer repo 執行）

Slice：一次性分析腳本 → 共用 harness（B4）。wrapper 收斂（B1/B2）依 Q4 另案裁決，不在本 plan。

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

- [x] 契約放置位置：cross-cutting（Q1）
- [x] 生長規則採 concept-first + reuse evidence（Q2）
- [x] Compensating-lock 為 Review Rule（Q3）；wrapper 收斂另案（Q4）；runner deferred（Q5）；登記 denylist（Q6）
- [ ] 概念判定準則與 capability catalog 形狀（Q2a / Q2b）
- [ ] Discovery 強制程度、scratch 位置、ownership 表達（Q7–Q9）
- [ ] Phase 2 / Phase 3 dogfood slice 範圍（各在 consumer repo 動到什麼）

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
