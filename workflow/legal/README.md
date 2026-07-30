# Legal Workflow

`workflow/legal/` 負責「法律工作的執行順序」。本 domain 的邊界是**法律任務**，不是合約文件——
Contract 只是任務之一。新增法域、法規、契約類型或任務型態時在本 domain 內擴充，
**不新增 workflow domain**。

> **邊界（必讀）**：本 domain 的「合約 / 契約」指**法律合約**。
> [`../software-delivery/contracts.md`](../software-delivery/contracts.md) 的「契約」指
> **API / BDD behavior contract**，兩者語意無關。選路裁決見
> [`../workflow-routing.md`](../workflow-routing.md) §常見歧義。

> **P0 邊界**：本 workflow 產出**不是法律意見**。Red tier 領域（勞動個案、股權、投資、
> M&A、訴訟、政府行政處分）只做升級，不產出實質建議。見
> [`risk-classification.md`](risk-classification.md)。

## Legal task type（第一級分派維度）

Intake 第一題問「你要做什麼」，答案決定下游要跑哪些 sub-flow：

| Task type | 意義 | 主 sub-flow |
| --- | --- | --- |
| `draft` | 我方出稿 | [`draft/`](draft/README.md) |
| `review` | 審閱對方稿 | [`review/`](review/README.md) |
| `explain` | 解釋既有條款／法律概念 | [`review/`](review/README.md) §Clause Review（只做解讀，不改稿） |
| `compare` | 比較兩個版本或兩份合約 | [`review/`](review/README.md) §Structural + Clause Review |
| `research` | 法規查詢／適用法研究 | [`research/`](research/README.md) |
| `due-diligence` | 對手方背景調查 | [`due-diligence/`](due-diligence/README.md) |
| `strategy` | 法律策略分析（怎麼安排最有利） | [`strategy/`](strategy/README.md) |
| `negotiation-support` | 談判支援（讓步矩陣、爭點收斂） | [`negotiation/`](negotiation/README.md) |
| `lifecycle` | 端到端合約生命週期管理 | [`lifecycle/`](lifecycle/README.md) |

Task type 未知時**必須先問**，不得從單句需求推定。

> [`strategy/`](strategy/README.md) 同時是**所有任務的共同 stage**（Stage 3a／3b），
> 不只是一種 task type。本 workflow 的核心不是「問問題 → 寫合約」，而是
> **「問問題 → 推理最佳策略 → 使用者決策 → 才決定怎麼寫」**。

## 何時讀哪個檔（thin index）

| 認知階段 | 檔案 | load_when |
| --- | --- | --- |
| Lifecycle 順序與分派 | [`execution-flow.md`](execution-flow.md) | 任何法律任務的入口（本 domain 的 `primary_source`） |
| 問什麼（分層問卷） | [`intake.md`](intake.md) | Stage 1；task type / jurisdiction / 我方角色未確認時 |
| **推理最佳策略** | [`strategy/README.md`](strategy/README.md) | **Stage 3a／3b，除 explain / compare 外一律必跑**；需要方案比較、建議與 trade-off 時 |
| 九個決策點的推理因素 | [`strategy/decision-playbooks.md`](strategy/decision-playbooks.md) | 決策點涉準據法／管轄／付款／違約金／IP／保密／驗收／不可抗力／終止 |
| 法域與準據法模型 | [`jurisdiction.md`](jurisdiction.md) | 跨境、或需決定 governing law / dispute resolution / court vs arbitration |
| 風險分級與升級 | [`risk-classification.md`](risk-classification.md) | **每個任務都要跑**；決定 Green / Yellow / Red 與是否停止產出 |
| 權威來源分類 | [`reference-sources.md`](reference-sources.md) | 需要官方範本、法條、Q&A、函釋；政府採購路徑必讀 |
| 產出格式與 gates | [`artifact-gates.md`](artifact-gates.md) | 產出前與宣稱完成前 |
| 起草 | [`draft/README.md`](draft/README.md) | task type = draft |
| 審閱／解釋／比較 | [`review/README.md`](review/README.md) | task type = review / explain / compare |
| 背景調查 | [`due-diligence/README.md`](due-diligence/README.md) | 有對手方且需核實其身分、營運狀態或風險訊號 |
| 適用法研究 | [`research/README.md`](research/README.md) | 需引用任何法規、範本或官方解釋 |
| 談判支援 | [`negotiation/README.md`](negotiation/README.md) | 多輪來回、需讓步矩陣或爭點收斂 |
| 生命週期管理 | [`lifecycle/README.md`](lifecycle/README.md) | 從需求到簽署存檔的端到端管理 |

**Suppression**：evidence-only 或純閒聊不應載入本 domain。只需一個 slice 時（例如純法規查詢）
載入 `execution-flow.md` + `intake.md` + `research/README.md` + `risk-classification.md` 即可。

## 核心原則

1. **Strategy before drafting**：先推理「怎麼安排對使用者最有利」，再決定條款怎麼寫。
   任何建議都要有 Recommendation / Reason / Alternative / Trade-offs 四欄，不得只給結論，
   也不得把決策原封丟回使用者。
2. **Jurisdiction first**：法域是 P0。同一個 `Termination` 條款在台灣／日本／美國／EU 的法律
   效果不同；法域未定即不做實質條款分析。
3. **Source-backed law**：任何法規／範本引用必須帶「名稱 + 版本或最新修正日 + 查核日」。
   無來源即標 `unverified`，不以確定語氣陳述。
4. **Intake before draft**：S0 必問未答完不得產出條款全文；只能輸出標記假設的骨架 +
   待確認清單。
5. **Two-sided interest analysis**：策略只算我方風險會得出對方不可能接受的方案。
   建議必須同時分析對手方的真實利益與可交換籌碼。
6. **Risk tier gates output**：Green 可完整產出；Yellow 產出但標明需覆核項；Red 停止實質
   建議與策略推薦、只做升級。
7. **Diligence changes clauses**：背調的 risk flag 必須對應到具體條款調整（母公司保證、
   付款節點、擔保、仲裁地），而不是附錄一份公司資料。
8. **No silent defaults**：法域、準據法、我方角色、金額量級都不得靜默預設。

## 與既有層的關係

- `workflow/legal/` 是法律工作執行流程的唯一入口。
- 本輪**不**建立 `analysis/legal/` 與 `intelligence/legal/`：尚無真實任務沉澱的 lesson，
  避免空殼層。等 lesson 出現後依 feedback promotion pipeline 升層。
- `enforcement/` 提供 dependency reading、linked updates、evidence hierarchy；
  法律 domain 知識**不**寫進 `enforcement/` 或 `runtime/` 成為全庫規則。
- Routing：`route.workflow.legal`（見 [`../../knowledge/runtime/routing-registry.yaml`](../../knowledge/runtime/routing-registry.yaml)）。

## Executable contracts

- [`execution-flow.yaml`](execution-flow.yaml) — lifecycle 的 executable contract
- [`artifact-gates.yaml`](artifact-gates.yaml) — 產出 gates 的 executable contract

## 驗證

- Agent 能否說出本次任務的 legal task type、jurisdiction 與 risk tier？
- 是否在產出任何條款文字前跑完 `intake.md` §S0？
- 是否在產出條款前跑過 Strategy（Decision Register + 四欄 Decision Reasoning）？
- 每個法規／範本引用是否有版本與查核日？
- Red tier 是否停止實質建議？
