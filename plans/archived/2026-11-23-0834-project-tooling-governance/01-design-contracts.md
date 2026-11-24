# 01 — Design Contracts（設計大綱）

> Companion of [`_plan.md`](_plan.md)。本檔只放六份契約的設計大綱、概念定義與相互關係。
> **執行狀態（Phase、checkbox、Open Questions 處置）只在 `_plan.md` 維護**，本檔不另記狀態。
> Phase 1 會依本大綱在 `workflow/cross-cutting/project-tooling/` 產出正式契約（每份 ≤150 行）。

## 範圍

治理對象是 **AI 在專案中創造、使用、保留與淘汰工具的行為**，不只腳本：一次性資料庫查詢、API 測試、圖片 / OCR / ASR 分析、Unity / Frida 實驗、部署與驗證工具都適用。腳本只是目前最明顯的案例。

設計目標：**新增一個工具是有意識的架構決策，不是 AI 完成任務時的預設副作用**；同時不讓 AI 在寫工具前填一堆表格。

## 1. Role contract（工具內部角色；Q10 已裁決）

**三個獨立維度，不得互相取代或混用**：Internal Role（本節）、Placement（§3）、Lifecycle State（§3）。目錄可以改名搬移；責任邊界與生命週期才是長期穩定的契約。

| Role | 允許 | 不允許 |
| --- | --- | --- |
| Entry | 解析參數、取得必要設定、呼叫 Logic、傳遞結果與結束狀態 | 領域邏輯、重複的 device / network / 檔案 I/O、資料處理 |
| Logic | 實際執行邏輯 | 寫死應由 Parameters 提供的值（目標 ID、座標、序號、門檻） |
| Parameters | 設定、策略選擇、數值 | 大量條件分支、流程描述（避免 configuration sprawl） |

- **Placement 不決定 Role**：Target、Domain、Shared 都可有自己的 Entry / Logic / Parameters，不要求同時具備三種。
- **違規描述格式**：`<placement> 的 <role> 放在 <實際 role / 位置>`。例：「Target 的 Logic 放在 Entry」。此格式不依賴任何語言的目錄結構。

## 2. Discovery contract（先找再建）

建立新工具前，agent 完成：

1. 搜尋專案既有工具與 capability catalog；
2. 搜尋 domain / shared 是否已有相同能力；
3. 確認能否以參數解決；
4. 確認能否擴充既有 harness；
5. 仍需新增時，給出 **creation justification** 與 placement / state（見 §3）。

合理新增的理由只有四種：新的執行責任、新的資料來源或介面、現有抽象無法合理涵蓋、實驗需要隔離。

**輕量原則**：這是執行前的工具發現與選擇行為，不是表單；justification 一兩句，寫在檔頭或 commit message。先作 behavioral workflow invariant，機械化 deferred（Q7）。

## 3. Lifecycle contract（兩個獨立維度）

**Placement 是服務範圍，State 是生命週期狀態，兩者獨立判斷，也都與 §1 Role 獨立。** 不是同一條階梯。

**Dimension A — Placement**（放哪、服務誰）：`Scratch → Research → Target → Domain → Shared`

**Dimension B — State**（目前的維護狀態）：`temporary / experimental / maintained / deprecated`

| 工具例 | Placement | State |
| --- | --- | --- |
| 臨時資料轉換 | Scratch | temporary |
| 尚未確認的實驗 | Research | experimental |
| 正式單目標工具 | Target | maintained |
| 領域共用工具 | Domain | maintained |
| 已淘汰的共用工具 | Shared | deprecated |

推論：

- Research 不一定升級成 Target；實驗失敗可直接 deprecated → 移除。
- Shared 也可以 deprecated。
- 不是每個工具都要升級，但**每個被保留的工具都要有存在理由**。「成功執行過」≠「值得永久保留」。

State 轉移的最小規則：

| State | 存在條件 | 到期 / 退出 |
| --- | --- | --- |
| temporary | 只為當前任務 | 任務結束清理 |
| experimental | 有 TTL、目的、輸入 / 輸出定義 | 到期必須重新評估，去向三選一：轉 maintained / 重設 TTL（附理由）/ deprecated；TTL 不是自動刪除指令 |
| maintained | 有 owner（缺 owner 屬 ownership violation，state 不變） | 無使用者 → 由 Selection 明確決定是否 deprecated |
| deprecated | 禁止新增使用 | 移除 |

Phase 1 已定義的最小責任（不做自動清理工具或 ownership registry；全文見 [`lifecycle-contract.md`](../../../workflow/cross-cutting/project-tooling/lifecycle-contract.md)）：

- **Q8 Scratch**：位置為專案指定的不進版控目錄或 session 暫存區，不得放在正式工具樹內；由建立它的 agent 在任務 close-out 時清理，無法清理須回報殘留路徑、原因與後續處理。
- **Q9 Ownership**：capability catalog 的 `owner` 欄位是唯一記錄點；缺 owner 是治理資訊不完整，不改變 state。
- **機械清理（Q12 G4）**：只限指定 Research 區域、排除使用情況無法可靠確認的工具、先隔離並保留可恢復期間。

## 4. Growth rule + Reuse evidence + Promotion ladder

### Concept check（Selection Responsibility）

查 capability catalog，回答「是否已有同概念的能力？」

| 判定 | 允許的動作 | 禁止 |
| --- | --- | --- |
| 同概念、只差數值 / 目標 | 在既有工具加參數，或加 config / strategy 值 | 複製一個新工具 |
| 同概念、流程有差異 | 先把共同邏輯抽到該分類的共用模組，再在同一功能分類下長 sibling（只放差異） | 另起爐灶；sibling 重寫共用邏輯 |
| 新概念（catalog 無對應） | 開新功能分類 + 新工具，並登記 catalog | 不登記 catalog |

**概念 = 能力或責任，不是相似的程式碼。** 兩支腳本都有迴圈、讀 JSON、輸出報告，不代表同概念；程式碼完全不同但承擔相同業務責任，可能就是同概念。最小準則（Q2a 細化）：「產出物 + 核心流程（責任）」相同即同概念；只差目標 ID、座標、門檻、站點屬參數。

### Config-first ≠ Config-everything

| 差異型態 | 落點 |
| --- | --- |
| 流程相同、資料不同 | Config |
| 流程部分不同、核心契約相同 | Strategy（例：多種 OCR 引擎共用輸入 / 輸出契約，各自實作） |
| 執行責任不同 | 獨立模組或工具（新概念） |
| 實驗尚未穩定 | Research |

### Reuse evidence（Governance Constraint）

計數對象是**同一能力 / 責任的出現次數**，不是相似程式碼。

| 次數 | 必須採取的行動 |
| --- | --- |
| 第 1 次 | 依概念決定工具位置（placement） |
| 第 2 次 | 進行重用與抽象化評估 |
| 第 3 次 | 明確裁決：抽取共用能力，或留下保留重複的理由 |
| 第 2 個 domain 使用 | 評估是否上推 Shared |

分工：concept check 決定「是不是同一件事、該怎麼長」；reuse evidence 確保抽象化決策不能無限期逃避。觸發的是決策，不是強制合併。

### Promotion ladder（抽共用時的落點）

沿 Placement 維度上推，每階可退回；共用邏輯抽到「涵蓋所有使用者的最低一層」即可，不預先上推到 Shared。

## 5. Lock taxonomy

layout / thin-entry / boundary / lifecycle / completeness / drift / parity / safety（八族，Q12 G1 新增 safety）；只用於工具治理鎖，產品不變式鎖不套族形狀但仍適用 Review Rule（Q12 G2）。每族一個「scope + forbidden pattern + allowlist」形狀；allowlist 必須附退場條件。**鎖保護架構不變式，不保護檔案數量**；只擋已明確違反的不變式。

### Governance check architecture（概念，不實作）

single check entry + check registry（跑哪些檢查）+ scope registry（適用範圍）+ 統一結果契約。各 consumer 用自己的語言實作；統一的是治理概念、檢查描述與結果形狀，不是 runner（Q5 deferred）。

## 6. Compensating-lock Review Rule

新增鎖時必答：「這個鎖防的重複，能不能從結構上消除？」能消除就先消除，鎖只守剩下的不變式。

## 責任切分與相互關係

| 責任 | 內容 | 承擔者 | 對應契約 |
| --- | --- | --- | --- |
| Selection | 用既有、擴充或新建；是否抽象化 | agent，留下 justification / 評估紀錄 | §2、§4 concept check |
| Constraint | 角色邊界、放置位置、依賴、生命週期、升級規則 | 專案契約 + 機械鎖 | §1、§3、§4 reuse evidence、§5、§6 |
| Verification | 無重複能力、無繞過共用層、無未處置臨時產物 | 獨立驗證（review / verifier） | 全部 |

**明確不做**：自動 orchestrator（看到新工具就自動搬層）。是否抽象化屬 Selection Responsibility，需保留決策與證據。
