# Project Tooling Governance（專案工具治理）

AI 在專案裡為了完成當下任務，很容易順手寫出新的腳本、查詢、測試或分析工具，卻不確認同樣的能力是否已經存在、是否該成為共用能力、任務結束後該怎麼處置。本 pattern 治理的是 **AI 在專案中創造、使用、保留與淘汰工具的行為**。

> **狀態**：`pilot`（reference contract）。定義**語意與判定**，不定義目錄名或語言；各專案用自己的結構實作，再連回本檔。
> **不是**必跑 stage、**不**註冊 `route.*`、**不**加 validator、**不**提供共用 runner。
> 來源計畫與證據：[`plans/active/2026-11-23-0834-project-tooling-governance/`](../../../plans/active/2026-11-23-0834-project-tooling-governance/_plan.md)（[`evidence/phase0-baseline.md`](../../../plans/active/2026-11-23-0834-project-tooling-governance/evidence/phase0-baseline.md)）。

## 設計目標

- **新增一個工具是有意識的架構決策，不是完成任務的預設副作用。**
- 不禁止新增工具：一次性與實驗工具有正當價值。要管的是何時可建、放哪、何時重用、何時升級、何時淘汰。
- 不讓 agent 在建工具前填一堆表格：判斷是輕量的，紀錄一兩句即可。
- 機械鎖保護架構不變式，不保護檔案數量；不再隨每個新需求線性增加。

範圍不只腳本：一次性資料庫查詢、API 測試、圖片 / OCR / ASR 分析、逆向實驗探針、部署與驗證工具同樣適用。

## 三個獨立維度

| 維度 | 值 | 回答的問題 | 契約 |
| --- | --- | --- | --- |
| Internal Role | Entry / Logic / Parameters | 這段程式在工具內負責什麼 | [`role-contract.md`](role-contract.md) |
| Placement | Scratch / Research / Target / Domain / Shared | 這個工具服務誰、範圍多大 | [`lifecycle-contract.md`](lifecycle-contract.md) |
| Lifecycle State | temporary / experimental / maintained / deprecated | 這個工具目前的維護狀態 | [`lifecycle-contract.md`](lifecycle-contract.md) |

三者**不得互相取代或混用**。目錄可以改名、搬移；責任邊界與生命週期才是應該長期穩定的契約。

**Ownership**（誰負責維護）不是第四個分類軸，而是記在 capability catalog 的治理資訊；同樣**不得互相推導**——不能用 Placement 推導 Role，也不能用 Ownership 缺失推導 Lifecycle State（見 [`lifecycle-contract.md`](lifecycle-contract.md) §Ownership）。

## 契約索引

| 檔案 | 解決的問題 | 何時載入 |
| --- | --- | --- |
| [`discovery-contract.md`](discovery-contract.md) | agent 為什麼一直重新創造工具 | 準備新增任何工具之前 |
| [`role-contract.md`](role-contract.md) | 入口變胖、邏輯藏進設定 | 寫或審查工具內部結構時 |
| [`lifecycle-contract.md`](lifecycle-contract.md) | 臨時工具變永久、正式工具無人負責 | 決定工具放哪、要不要保留時 |
| [`growth-rule.md`](growth-rule.md) | 同概念重複、過度或不足的抽象化 | 發現相似能力、考慮加參數或抽共用時 |
| [`lock-taxonomy.md`](lock-taxonomy.md) | 機械鎖逐事件蔓延、用鎖補償結構缺陷 | 想新增或修改 pre-commit / lint / gate 時 |

## 責任切分

| 責任 | 內容 | 承擔者 |
| --- | --- | --- |
| Selection | 用既有工具、擴充、或新建；是否抽象化 | agent；留下 justification 或評估紀錄 |
| Constraint | 角色邊界、放置範圍、生命週期、升級規則 | 專案契約 + 機械鎖（只擋已明確違反的不變式） |
| Verification | 無重複能力、無繞過共用層、無未處置的臨時產物 | 獨立驗證（review / verifier） |

**明確不做**：自動 orchestrator（例如偵測到新工具就自動搬層）。是否抽象化屬 Selection，必須保留決策與證據。

## 採用方式（給專案）

1. 在專案的工具樹，每個功能分類 README 放一張 **capability catalog** 表（欄位見 [`lifecycle-contract.md`](lifecycle-contract.md) §Capability catalog）。
2. 依 [`role-contract.md`](role-contract.md) 檢查現有入口；依 [`lock-taxonomy.md`](lock-taxonomy.md) 盤點現有鎖，先套 compensating-lock Review Rule。
3. 之後新增工具前走 [`discovery-contract.md`](discovery-contract.md)。

## 收錄門檻與升級

比照本目錄其他 pattern：需 **≥3 個收斂案例**才考慮升級為 stage、route、validator 或共用 runner。目前：2 個外部 consumer 做**唯讀評估**（不修改其 repo）＋ Ai-skill 自身工具樹做**實際套用**；唯讀評估是否計入門檻由來源計畫 Phase 4 裁決。

← [Back to cross-cutting index](../README.md)
