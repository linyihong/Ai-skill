# Phase 0 — Baseline、Concept 試判與 Preflight（2026-11-23）

> 唯讀量測，對兩個 consumer 的 HEAD 執行；consumer repo 的 tracked 內容未修改。
> 可重跑指令、具體檔名、HEAD SHA 記在各 consumer 的 project-local dogfood 檔（見 [`README.md`](README.md) 引用規則）。

## Baseline 量測

| 指標 | Consumer A（Python + Frida） | Consumer B（bash + PowerShell + Node） |
| --- | --- | --- |
| scripts 根目錄項目數 | 7（只放分類目錄） | 132（53 組 `verify-*` 成對腳本平放於根） |
| 工具檔數 | 180 `.py`、82 `.js` | 67 `.sh`、67 `.ps1`、158 `.mjs` |
| L0 入口數 | 49 | 53 組入口 wrapper（各 sh + ps1） |
| 非薄 L0 入口（> 25 行或含 I/O） | **3**（31 / 132 / 260 行；後兩者各含 10+ 處 device I/O） | 14 組 > 60 行的 orchestrator（屬合理入口職責，非違規） |
| 近重複 wrapper | 低（L0 為 import + `main()` 樣板） | **26 組 ≤ 20 行**、只呼叫單一 lib check 的 wrapper（×2 語言 = 52 檔） |
| per-case 工具的專案內 import 比例 | 單目標 spin-hunt 類 8／10 支引用共用 helper | `performance/` **0／8** |
| 機械鎖數 | 24 個 check（pre-commit 逐行列出 24 個呼叫）；8 個鎖有測試 | ~53 個 verify 入口 + 82 個 lib check |

## 缺口驗證（對 HEAD）

| # | 原觀察 | 結果 | 證據摘要 |
| --- | --- | --- | --- |
| A1 | L0 薄入口是慣例不是鎖 | **verified** | 3 支非薄入口；其中 2 支無任何共用 helper 引用，自帶 tap、座標常數、序號預設值 |
| A2 | boundary 鎖 scope 只覆蓋單目標目錄 | **verified** | device I/O 鎖只掃單目標目錄樹；L0 不在範圍內 |
| A3 | 裝置序號預設值分散（原記 7 檔） | **verified（數量修正）** | 帶**寫死預設值**的是 **4 檔**：3 支在 L0、**1 支在 L2 shared 本身**；另 3 檔只讀環境變數、不帶預設（合規）。原「7 檔」把合規讀取也算進去 |
| A4 | 無升級觸發條件 | **verified** | README 只有放置與禁止規則；「promote」字樣只出現在資產（圖檔 / protocol ID）流程，與工具抽象化無關 |
| A5 | 鎖本身逐檔蔓延 | **verified** | 24 個獨立 check 檔，pre-commit 逐行呼叫 24 次；新增一鎖 = 新檔 + 改 hook |
| B1 | 入口 wrapper 大量重複（原記「多數只是呼叫 lib」） | **verified（比例修正）** | 53 組中 **26 組**是 ≤ 20 行的單一 lib 呼叫；14 組是 > 60 行的合理 orchestrator；其餘介於兩者 |
| B2 | compensating lock | **verified** | parity 鎖唯一職責是檢查 sh / ps1 成對存在 |
| B3 | `lib/` 平鋪 | **verified** | 150 檔：82 check、46 test、22 其他（deploy / scope / attestation helper）同層 |
| B4 | 一次性分析腳本蔓延 | **verified** | 8 支 per-case 腳本 0 個專案內 import；計時 7／8、fixture 環境斷言 6／8、報告寫檔 7／8、hash 5／8 各自重寫 |

## Concept 試判（以能力 / 責任分組，不以程式碼相似度）

| 概念（能力 / 責任） | Consumer | 工具數 | 判定 | 現況 |
| --- | --- | --- | --- | --- |
| 「坐檯旋轉直到特定事件，並擷取證據」 | A | 10 | 同概念、流程有差異（每目標的事件條件不同） | 8 支已走「共用 helper + 單目標 sibling」；**2 支誤置於 L0 且完全不用共用 helper**（= A1 的實體） |
| 大廳 tile 處理（擷取 / dump / 匯出 / 拉取 / 去 overlay / 同步） | A | 6 | **不同概念**：同一產出物，但各自是 pipeline 的不同責任 | 合理分立 |
| 「對隔離 fixture 跑一條效能 lane 並出報告」 | B | 8 | 同概念、流程有差異（每案的請求序列不同） | 每支自帶 harness = 應抽共用 harness + per-case sibling（Phase 3 slice） |
| 「呼叫第 X 個 lib check」入口 | B | 26 組 | 同概念、只差參數（check 名稱） | 應為單一參數化入口；屬 Q4 另案 |
| 各 lib check（82 個不變式） | B | 82 | 不同概念（各守不同不變式） | 合理分立；適合 check registry 概念（Q5 deferred） |

**Q2a 證據**：「產出物 + 核心流程（責任）相同 = 同概念」在兩個 consumer 都判得出來，且兩條都需要——大廳 tile 案例只有產出物相同、責任不同，正確判為不同概念。沒有遇到判不出來的案例。建議措辭改為「產出物 + 核心責任」。

## Placement × State 試標

| 工具群 | Placement | State | 是否被雙維度描述得清楚 |
| --- | --- | --- | --- |
| Consumer A 實驗用 Frida 探針 | Research | experimental（有 cleanup 鎖） | 是 |
| Consumer A 單目標 hunt 工具 | Target | maintained（**無 owner 紀錄**） | 是，但缺 owner → 支持 Q9 |
| Consumer A 2 支誤置 L0 的 hunt 工具 | 應為 Target，實際在 L0 | maintained？（無 justification、無 owner） | **否**：L0 不是 placement 值 → 見下方 Q10 |
| Consumer A capture / paytable / lobby 模組 | Domain | maintained | 是 |
| Consumer A shared/device | Shared | maintained | 是（但其中 1 檔仍帶寫死序號，A3） |
| Consumer B `performance/` per-case | Target | maintained | 是 |
| Consumer B 根目錄單一 SQL 修復腳本 | Scratch？ | temporary 卻長期留存 | 是，且正好是「成功執行過 ≠ 值得保留」的候選 |
| Consumer B lib check | Domain | maintained | 是 |

## 新發現

1. **Q10 — Layer contract 與 Placement 維度部分重疊。** Layer contract 用 L0 entry / L1 domain / L2 shared / L3 config；Placement 用 Scratch / Research / Target / Domain / Shared。Domain、Shared 同時出現在兩套分類，而 L0 / L3 不是 placement。這是本契約內部的雙份分類（duplication risk）。
   **提議解法（safe_assumption，Phase 1 可逆）**：Layer contract 改為描述**工具內部的角色**——entry（只解析參數 + 呼叫）/ logic / parameters（config、strategy）；**範圍歸 Placement 單一維度管**。A1 的誤置因此表述為「Target 層的 logic 放在 entry 角色裡」。
2. **Cross-cutting 收錄門檻是三個收斂案例。** `workflow/cross-cutting/README.md` 對 pattern 升為 slice / stage 要求 ≥3 個收斂案例；本 plan 目前 n=2，維持 pilot 與 doc-only 一致。Phase 4 的升級判斷應沿用三案例門檻。
3. **鎖 = consumer 的 Automation 層。** software-delivery framework domain model 把 gate / validator / hook 定位為 Policy + Process 的 runtime projection（通用 Automation 實作 deferred）。compensating lock 可表述為「Automation 在補 Process / 結構的缺口」；本 plan 不實作通用 Automation，無重複。
4. **Governance pattern library 無重疊。** 該 plan 治理的是 Ai-skill 自身治理子系統的形狀（Observation → … → Validation），不是 consumer 工具鎖；lock taxonomy 不搶先登記為其 pattern。

## Framework discovery 與 Preflight

| 欄位 | 內容 |
| --- | --- |
| Trigger | 開始執行本 plan Phase 0 |
| Checked sources | `workflow/cross-cutting/README.md`、`traceable-evidence/README.md`、governance pattern library plan、software-delivery framework domain model（archived）、`pre-build-interrogation.md`、`plans/README.md` preflight 清單、`governance/lifecycle/plan-evidence.md` |
| Conflicts | 無外部 source-of-truth 衝突；內部有 Q10（Layer vs Placement 重疊） |
| Open Questions 核對 | Q2a → resolved（見 Concept 試判）；Q2b、Q2c → 有證據但留 Phase 1 定案；Q8、Q9 → Phase 1 必交（Q9 有新證據）；新增 Q10 |
| Decision | **proceed**（Q10 以 safe_assumption 進 Phase 1，Phase 1 產出後由使用者確認） |
| Validation | 量測以 shell 對 consumer HEAD 重跑；本檔與 project-local 檔 sanitization scan 通過 |

### Pre-build Interrogation

- Goal：產出 doc-only 的 Project Tooling Contract（6 份），讓 consumer 能據以收斂工具。
- Scope：`workflow/cross-cutting/project-tooling/` 6 檔 + cross-cutting README 連結。
- Non-goals：runner、runtime route、validator、consumer wrapper 收斂（Q4）、自動 orchestrator。
- Acceptance / validation target：Phase 1 完成條件（6 檔各 ≤150 行、Q8 / Q9 / Q10 resolved、sanitization 通過、無 route）。
- Framework discovery：同上表 Checked sources。
- Duplication risk：契約內部 Layer vs Placement（Q10）；對外無。
- Open questions：Q2b、Q2c、Q8、Q9、Q10。
- Assumptions：Q10 提議解法（可逆）。
- Decision：proceed。
