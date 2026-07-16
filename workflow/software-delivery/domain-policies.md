# Software Delivery Framework — Domain Policies

Canonical **Policy** layer for [`README.md`](README.md) §Framework Domain Model（N = 3：Asset、Policy、Process）。

> **邊界**：本檔定義可操作規則。Governing Principles（信念）見 §1；Process 順序見 [`execution-flow.md`](execution-flow.md)；Automation 投影見 `artifact-gates.yaml` / `execution-flow.yaml`。  
> **驗證**：[`plans/active/2026-07-16-0945-software-delivery-framework-domain-model/evidence/phase-0-classification-matrix.md`](../../plans/active/2026-07-16-0945-software-delivery-framework-domain-model/evidence/phase-0-classification-matrix.md)

建立任何 asset 前：`Intent → Asset (class) → **本檔 Policy** → Process stage`。

---

## 1. Governing Principles（指導 Policy，非 Policy 正文）

| Principle | 含義 |
| --- | --- |
| Classification before Creation | 先定 Asset class，再建檔案 |
| One Asset, One Owner | 每個 asset 實例僅一個 owner |
| Project ≠ Framework Knowledge | 專案證據不進可重用 framework source |
| Contract over Implementation | 衝突時 contract / behavior 高於 code |
| Promotion before Reuse | 可重用規則須經 promotion，不直接生在 project |
| Workflow consumes; does not define Assets | Process 不定義 placement / owner |

---

## 2. Ownership Policy

| Asset class | 預設 owner | 備註 |
| --- | --- | --- |
| Plan deliverable（brief、implementation plan、parity inventory） | 發起人 / PM（專案） | change brief 由 intake 產生 |
| Contract / Behavior（BDD、domain/API/UI contract） | 架構或全端負責人（專案） | consumer contract 與 API 同步審查 |
| Evidence（validation output、review report） | 驗證或 review 執行者（專案） | evidence 不得單獨 override contract |
| Decision（ADR、project decision） | 決策者（framework 或 project） | 平台 ADR → `constitution/`；專案 → `docs/decisions/` |
| Template、Pattern entry、Knowledge | framework maintainer（Ai-skill） | promotion 見 §5 |
| Process component（`sd-*` slice） | software-delivery workflow owner | 見 `governance/cognitive-slice-taxonomy.md` |

專案覆寫：`<PROJECT_ROOT>/.ai-skill/project/rules/`（不得改寫 framework term 定義）。

---

## 3. Placement Policy

### 3.1 Framework knowledge（Ai-skill repo）

| 內容 | 放置位置 |
| --- | --- |
| 跨平台安全原則 | `controls/` |
| 平台 / 嵌入式實作細節 | `platforms/` |
| 語言陷阱 | `languages/` |
| 可建置模式 | `implementation/` |
| 審查流程 | `checklists/` |
| 浮現中教訓 | `feedback/history/development-guidance/` |
| APK 分析 | `analysis/apk/`、`workflow/apk-analysis/` |
| UI pattern seeds | [`ui-pattern-knowledge/`](ui-pattern-knowledge/README.md) |
| 交付流程與 Policy | `workflow/software-delivery/` |
| 共享清理 / 連動規則 | `enforcement/` |

完整表歷史正文曾於 [`artifact-gates.md`](artifact-gates.md) §2；**canonical 移本檔**，`artifact-gates` 保留品質門檻與 evidence shape。

### 3.2 Project delivery assets（預設 — 專案可 overlay）

| Asset class | 預設路徑 | Template |
| --- | --- | --- |
| Plan deliverable | `docs/planning/` 或單一 `docs/planning/delivery-pack.md`（小專案） | [`templates/change-brief-template.md`](templates/change-brief-template.md) |
| Contract | `docs/contracts/` | [`templates/contract-template.md`](templates/contract-template.md) |
| BDD / behavior | `tests/bdd/` 或 `docs/behavior/` | [`templates/bdd-scenario-template.md`](templates/bdd-scenario-template.md) |
| Implementation plan | `docs/plans/` | [`templates/implementation-plan-template.md`](templates/implementation-plan-template.md) |
| Evidence | `docs/evidence/` | [`templates/ui-governance-evidence-template.md`](templates/ui-governance-evidence-template.md) 等 |
| Review report | `docs/reviews/` | [`templates/review-report-template.md`](templates/review-report-template.md) |
| Project decision | `docs/decisions/` | ADR 風格即可 |
| 憑證 / 樣本 ID | gitignored `.local/` 或專案宣告的 ignored 路徑 | 不進版控 |

多倉 workspace：外層 orchestrator vs sibling 分工見 [`implementation/multi-repository-workspace-mode.md`](implementation/multi-repository-workspace-mode.md)。

---

## 4. Authority Policy（文件衝突優先序）

當專案內多份文件不一致時，**除非專案有更強本地規則**，採下列優先序（canonical 原出處 [`contracts.md`](contracts.md) §Contract Governance Gate）：

1. Governance / framework contract（repo invariants、命名、build/run）
2. Product plan / accepted brief
3. BDD behavior（可觀察行為與 acceptance）
4. Domain、architecture、API、consumer/UI、error、hardware contracts
5. 實作與 generated clients
6. Tests、fixtures、examples

較低層發現較高層錯誤時：**不得默默改 code**；依 [`contracts.md`](contracts.md) 衝突表分類處理。

Review / invoke **不能** override acceptance gate 或 contract precedence。

---

## 5. Promotion Policy

可重用內容不得直接生在 project。平台通用管道：[`governance/lifecycle/decision-promotion-pipeline.md`](../../governance/lifecycle/decision-promotion-pipeline.md)。

| 來源 | 典型去向 |
| --- | --- |
| Session / project lesson | `feedback/history/` → `intelligence/` 或 `workflow/` |
| Ship 後 UI incident 學習 | [`change-retrospective.md`](change-retrospective.md) 三選一 |
| 架構級不可逆決策 | `constitution/ADR-*`（須滿足 ADR Promotion Criteria） |

---

## 6. Knowledge Boundary Policy

套用 [`enforcement/reusable-guidance-boundary.md`](../../enforcement/reusable-guidance-boundary.md)：

- **Framework**：通用規則、故障模式、決策規則、驗證方法（已清理）
- **Project**：主機、token、樣本 ID、類別名、一次性結論

---

## 7. Lifecycle Policy（Project asset 實例）

| 狀態 | 適用 |
| --- | --- |
| draft | 討論中；不可作為 implementation 唯一依據 |
| accepted | 可作為 contract / plan 輸入 |
| superseded / deprecated | 保留連結；新工作不得引用為 canonical |

Framework Knowledge asset 生命週期見各 layer README（candidate → validated → promoted）。
