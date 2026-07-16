---
id: 2026-07-16-0945-software-delivery-framework-domain-model
plan_kind: main
status: in-progress
owner: linyihong
created: 2026-07-16
last_updated: 2026-07-16
revision:
  - date: 2026-07-16
    note: "Phase 0 完成 — N=3 驗證；evidence matrix；README + execution-flow pointer（零新第一級目錄）"
  - date: 2026-07-16
    note: "Reviewer feedback — Success Criterion 北極星；Phase 0 改為最小完整核心集合（非預設三核心）；Constitution → Principles → Policy 層次修正"
  - date: 2026-07-16
    note: "Phase 3 execute preflight — dual-plan test with greenfield consumer; evidence phase-3-external-greenfield-consumer-execute"
required_for_completion: false
parent: null
---

# Software Delivery Framework Domain Model

**Status**: `in-progress`（Phase 0 ✅ · Phase 1 ✅ · Phase 2 ✅ · Phase 3 external observational ✅）  
**Owner**: framework maintainer (linyihong)  
**建立日期**：2026-07-16  
**Priority**：**P1**（架構收斂；阻擋 premature documentation sprawl）

> ## Success Criterion（北極星）
>
> **本計畫的成功不是建立新的文件結構**，而是找出一個足以解釋現有 Software Delivery Framework 的**最小 Domain Model**。
>
> 若現有模型已足夠，本計畫可以在**不新增任何第一級文件或目錄**的情況下完成——例如僅：
> - `README` 增加一張 Domain 圖
> - `execution-flow` 改幾行 pointer
> - **零** 新資料夾
>
> 成功 **≠** 新增 `constitution.md`。成功 **=** Domain 可解釋、可推導、可維護。

> ## 本 plan 要解什麼
>
> `workflow/software-delivery/` 已從「交付流程」演化成 **Software Delivery Framework**，
> 但缺少收斂的 **Primary Domain Model**。目前 agent / 貢獻者面對的是流程步驟與散落規則，
> 而非少數核心抽象；每新增一種產物就容易重新發明 `docs/` vs `plans/` vs `contracts/`，
> 或在 workflow slice 內重定義 owner / placement / promotion。
>
> **本 plan 先收斂 Domain Model，再決定 Documentation Architecture。**
> 在 Primary Model 驗證通過前，**不**建立 `constitution.md`、`concepts/`、`policies/` 等
> 第一級目錄樹，避免 over-orthogonalization。

---

## 討論脈絡（conversation capture）

以下為 2026-07-16 設計討論的濃縮記錄，作為本 plan 的 empirical trigger。

### 起點：Placement 缺口

- 問題：`software-delivery` 有 workflow（要做什麼），但缺「產物放哪、誰擁有」的統一規範。
- 初步方向：補 **Artifact Placement Charter**（`artifact-placement.md`）。

### 第一輪升級：Framework Constitution

- 共識：缺的不只是 placement，而是 **Framework Constitution**（世界觀 / 原則）。
- 提議六柱：Ownership、Placement、Promotion、Knowledge Boundary、Authority、Evolution。
- 提議四件套：README、constitution、execution-flow、glossary。
- 提議 taxonomy-first：`Purpose → Artifact Class → Owner → Location → Lifecycle → Workflow`。

### 第二輪升級：四層 + Automation

- 修正：Taxonomy **不應**寫死在 Constitution（ontology 會演化；憲法應極穩定）。
- 分層：
  1. **Constitution** — 信念（幾乎不變）
  2. **Concepts / Taxonomy** — ontology（可演化）
  3. **Policies** — 依 taxonomy 的管理規則（placement、ownership、promotion、authority、lifecycle）
  4. **Workflow** — 何時做（消費上層，不擁有分類）
  5. **Automation** — Policy + Process 的 runtime projection
- 用詞升級：**Asset** 取代 Artifact 作為總稱（Workflow、Capability、Template 等不全是「交付 artifact」）。
- 關鍵句：**Workflow stage = 這類 Asset 通常在哪个階段產生**，而非「這個階段可以建什麼」。

### 第三輪煞車：Over-orthogonalization

- 風險：第一級概念膨脹（Constitution、Derivation、Concepts、Policies、Workflow、Automation、Templates、Examples…）→ 新人不知先看哪份。
- 轉向：**Documentation Architecture ≠ Domain Model**。
- 假設 **Primary Model 可能只有三個第一級概念**：

```text
Asset  — 世界裡有什麼（managed entity）
Policy — 這類東西怎麼管（rules derived from asset class）
Process — 什麼時候做（workflow / lifecycle stage consumption）
```

- 其它（Constitution、Taxonomy、Derivation、Automation）可能是 **View** 而非第一級：
  - Constitution → **Governing Principles**（見下方層次）— 不是 Policy 本身
  - Taxonomy → Asset 的 ontology
  - Workflow → Process 的執行面
  - Automation → Policy + Process 的 runtime projection
  - Derivation → `Intent → Asset → Policy → Process` 的自然推導；**可能只需 README 一張圖**，不必獨立 `derivation.md`

### 本 plan 採納的立場

1. **方向正確**：taxonomy-first、workflow 不擁有 asset 定義、Asset > Artifact、promotion / boundary 需正式化。
2. **順序修正**：先 **Domain Model 驗證**，再 Documentation Architecture。
3. **收斂目標**：找到**最小且完整**的核心概念集合以容納 repo 現有核心名詞——**不預設一定是三個**；三核心僅為 working hypothesis，須經 Phase 0 否證或修訂（避免 solution-first / confirmation bias）。

---

## Problem statement

| 現象 | 根因 |
| --- | --- |
| 每種新 artifact 重新討論 `docs/` vs `plans/` | 無穩定 **Asset class → placement** 推導 |
| `intake` / `artifact-gates` / `contracts` 各自定義放置與 owner | **Policy** 未獨立；混在 **Process** 文件裡 |
| `cognitive-slice-taxonomy` 與 delivery artifact 分類並行 | 載入粒度 taxonomy ≠ delivery **Asset** taxonomy |
| 新人先問「先看哪份」 | 第一級文件/目錄過多，缺 **Primary Model** 入口 |
| 討論快進入九個第一級目錄 | **Documentation architecture 領先 Domain model** |

---

## Primary Model 假設（待驗證 — 不預設核心個數）

> **研究紀律**：Phase 0 的目標**不是證明三核心**，而是找到**最小且完整（minimal and complete）**的核心集合。Asset / Policy / Process 為 **working hypothesis**，可因證據修訂為 2、4 或更多——以分類矩陣與否證為準，不以「三個很漂亮」為準。

### Working hypothesis：三核心（candidate）

```text
                    Framework Domain
                 +------------------+
                 |      Asset       |  可管理實體（含 framework 與 project 實例）
                 +------------------+
                  /    |    |    \
           Knowledge Contract Evidence …（subtype / class，非第一級）
                 \    |    |    /
                  +------------------+
                  |      Policy      |  對 Asset class 的管理規則
                  +------------------+
                           |
                  +------------------+
                  |     Process      |  時間序上的工作與階段（workflow 是 Process 的 view）
                  +------------------+
                           |
                    Workflow slices
```

### 推導鏈（domain-native，未必獨立成檔）

```text
Intent
  → Asset (class + instance)
  → Policy (ownership, placement, promotion, authority, lifecycle)
  → Process (which stage typically produces/consumes this asset)
  → Automation (optional gate projection)
```

### Asset 總稱下的 class（ontology — 可演化，不進 constitution）

候選 class（**非**第一級 domain 概念）：

| Class | 簡述 |
| --- | --- |
| Knowledge | 可重複事實 / 方法 |
| Workflow | 可重複執行順序（framework asset） |
| Policy | 管理規則本身也是 asset |
| Contract | 行為 / 介面承諾 |
| Runtime | 機器可讀 contract / projection |
| Evidence | 證明 claim 的產物 |
| Decision | 選擇記錄 |
| Template | artifact 形狀 |
| Automation | gate / validator / hook |
| Capability | invoke 邊界（cross-cutting） |
| Projection | generated surface（runtime.db） |

**Artifact**：降級為 **Project 層、為完成某次 delivery intent 而產生的 Asset 實例**（tag 或 lifecycle 狀態），非與 Asset 並列的第一級。

### Constitution → Principles → Policy（抽象層次 — 非等同）

Constitution **不是** Policy；Policy **也不是** Constitution。Policy 是依 **Governing Principles** 制定的可操作規則。

```text
Constitution（框架憲章 — 極少數、幾乎不變）
        ↓
Governing Principles（治理原則 — 如 One Asset One Owner、Classification before Creation）
        ↓
Policy（依原則與 Asset class 制定的具體規則 — 如 Evidence owner 必須…、Contract owner 必須…）
        ↓
Automation（Policy + Process 的 runtime projection）
```

| 層次 | 範例 | 性質 |
| --- | --- | --- |
| Principle | One Asset One Owner；Project ≠ Framework Knowledge | 抽象、不綁具體 asset class |
| Policy | Evidence 須有單一 owner；Contract 衝突依 precedence 表處置 | 可操作、可檢查、依 class 分化 |

若未來文件化：`constitution.md` 承載 **Principles**（或指向 principles 小節），**不是** operational policy 正文。

### View 對照（documentation 可重組；domain 不應隨目錄改名而變）

| 若採文件化 | Domain 歸屬 | 變動頻率 |
| --- | --- | --- |
| `constitution.md` | Governing Principles（**非** Policy 正文） | 極低 |
| `concepts/asset-taxonomy.md` | Asset ontology | 中 |
| `policies/placement.md` 等 | Policy（依 Principles 制定） | 中 |
| `execution-flow.md` + slices | Process | 中偏高 |
| `artifact-gates.yaml` 等 | Automation（Policy+Process projection） | 隨 policy 演進 |
| `derivation.md` | **可能不需要** — README 圖即可 | — |

---

## 與 repo 現況的關係

| 現有表面 | 本 plan 視角 |
| --- | --- |
| `constitution/ADR-*` | **Ai-skill 平台** foundational Policy / Decision（跨 framework） |
| `workflow/software-delivery/*` | **Software Delivery Framework** 域；待收斂為 Asset/Policy/Process |
| `governance/cognitive-slice-taxonomy.md` | Process **載入粒度**（cognitive_slice），≠ delivery Asset taxonomy |
| `knowledge/glossary/ai-skill.md` | 平台詞彙；SD domain glossary 延後（見 glossary README） |
| `enforcement/` | 平台級 Policy + Automation |
| `artifact-gates.md` | Policy（placement 碎片）+ Process 品質 gate 混雜 |
| `contracts.md` precedence | Policy（authority）碎片 |
| `decision-promotion-pipeline.md` | Policy（promotion）平台通用；SD 應引用特化 |
| `change-retrospective.md` | Process 出口 + Promotion 決策點 |

**邊界**：本 plan 收斂 **Software Delivery Framework** 的 domain model，不取代 `constitution/ADR-*`，但在 ADR Promotion Criteria 通過前不寫新 ADR。

---

## Phase 0 — Primary Model 驗證（**當前唯一授權工作**）

**目標**：找出**最小且完整**的核心概念集合，足以解釋 repo 內 Software Delivery 相關實體與規則——**不預設核心個數為三**。working hypothesis（Asset / Policy / Process）僅為起點，須接受否證與修訂。

### 0.1 概念盤點（inventory）

從 repo 列出待分類名詞（至少涵蓋）：

- [x] change brief / product brief
- [x] BDD scenario / Journey Specification
- [x] ADR / project decision
- [x] Evidence / evidence shape / journey validation
- [x] UI Contract / Screen Mapping / Pattern entry
- [x] Execution slice (`sd-*`) / cognitive slice
- [x] Capability / invoke envelope
- [x] Runtime contract / YAML projection / generated surface
- [x] Automation rule / validator / hook
- [x] Template
- [x] Plan artifact / implementation plan
- [x] Feedback lesson / failure pattern
- [x] Glossary term
- [x] Recovery policy（`metadata/recovery/` — 與 SD Policy 區分）

**產出**：[`evidence/phase-0-classification-matrix.md`](evidence/phase-0-classification-matrix.md)（40 概念、95% 單一主歸屬、N=3 判定）。

### 0.2 分類規則（working）

每個概念必須回答：

1. **它是 Asset、Policy、還是 Process？**（或證明需要第四類）
2. 若是 Asset：**class** 是什麼？project 還是 framework layer？
3. 適用哪些 **Policy**（ownership / placement / promotion / authority / lifecycle）？
4. 哪個 **Process** stage 通常產生或消費它？
5. 是否已有 **Automation**？若無，是否應有？

### 0.3 否證條件（falsification — 避免 confirmation bias）

| 結果 | 行動 |
| --- | --- |
| 存在 **N** 個核心（N 可能為 2、3、4…），使 ≥ 90% 概念有單一主歸屬且推導鏈一致 | Phase 1：凍結 **N**-core Primary Model，設計**最小**文件面（含 N=已有解釋力、零新目錄） |
| working hypothesis（三核心）無法容納穩定例外 | **修訂 N**（增或減），記錄證據；**不得**為保留「三個」而硬塞 |
| 分類混亂、同一概念多歸屬無解 | **停止**文件化；回到 Phase 0 細化定義，不開新目錄 |
| 現有 README + execution-flow pointer 已足夠表達模型 | **直接完成 plan**（見 Success Criterion）；不為完成感而新增目錄 |

### 0.4 視覺驗收

- [x] 能用 **一張圖**（見 [`evidence/phase-0-classification-matrix.md`](evidence/phase-0-classification-matrix.md)）表達 Asset → Policy → Process，無「懸空盒子」
- [x] 圖上每個盒子能對應 ≥ 3 個 repo 實例（見 matrix §Asset class 穩定枚舉）
- [x] 新人 5 分鐘內能回答：「change brief 是什麼？誰擁有？通常何時產生？」（見 matrix §Dogfood 抽樣）

### Phase 0 結論（2026-07-16）

| 項目 | 結果 |
| --- | --- |
| **N** | **3** — Asset、Policy、Process |
| **覆蓋率** | 38/40 confirmed 單一主歸屬；2 meta/projection |
| **第四核心** | 不需要（Intent / Automation / Principles 已解釋） |
| **文件產出** | README §Framework Domain Model + execution-flow pointer；**零** 新第一級目錄 |
| **詳細證據** | [`evidence/phase-0-classification-matrix.md`](evidence/phase-0-classification-matrix.md) |

**Phase 0 完成前禁止**：

- 新增 `workflow/software-delivery/constitution.md`
- 新增 `concepts/`、`policies/` 第一級目錄樹
- 新增 `derivation.md` 作為 canonical source
- 大規模改寫 `intake.md` / `artifact-gates.md`（除加「待 domain model plan 收斂」pointer）

---

## Classification Matrix（Phase 0 完成）

完整 40 列矩陣、否證記錄、Asset class 枚舉見 [`evidence/phase-0-classification-matrix.md`](evidence/phase-0-classification-matrix.md)。

**摘要**：working hypothesis（Asset / Policy / Process）**成立**；Governing Principles 與 Automation 為 meta/projection，不擴 N。

---

## Phase 1 — 凍結 Primary Model（僅在 Phase 0 通過後）

- [x] 在 plan 記錄 **N = 3** 與定義（見 §Phase 0 結論）
- [x] 更新 `workflow/software-delivery/README.md`：§Framework Domain Model（圖 + 推導鏈 + 三核心表）
- [x] `execution-flow.md` 加 Domain Model pointer（Process 層邊界）
- [x] 決定 glossary 策略：`knowledge/glossary/software-delivery.md`（5 個 candidate terms；owner-layer workflow-orchestration）
- [x] **零** 新第一級目錄；未建 `domain-model.md`（README 足夠）

---

## Phase 2 — Documentation Architecture（完成）

依 Phase 0 結論選 **最小** 文件切分：

| 檔 / 目錄 | 狀態 |
| --- | --- |
| `domain-policies.md` | ✅ 單檔 Policy canonical（無 `policies/` 目錄） |
| `knowledge/glossary/software-delivery.md` | ✅ Primary Model 詞彙 |
| `artifact-gates.md` §2 | ✅ pointer；placement 雙寫已移除 |
| `contracts.md` authority | ✅ pointer 至 domain-policies §4 |
| `constitution.md` | ⏸ 未建（Principles 在 domain-policies §1 足夠） |
| `intake.md` | ✅ classify-before-create pointer |

**原則**：文件架構可重組；**Domain 名詞不隨目錄改名**。

### Workflow slice 遷移規則（Phase 2）

- `intake` / `contracts` / `artifact-gates`：**刪除** placement / owner 正文 → pointer 到 Policy
- `execution-flow.md`：只保留 Process 順序 + 「本 stage 常觸發的 Asset class」表
- 新增 asset 時：先更新 taxonomy（Asset class），再更新 policy 表，最後才動 workflow

---

## Phase 3 — External validation + Automation projection（選做）

### 3a External dogfood — greenfield consumer（observational ✅）

比照 [`2026-07-08-0825-delegation-verification-arbitration-loop`](../../2026-07-08-0825-delegation-verification-arbitration-loop/_plan.md) 外部 repo 證據慣例；**本 run 僅分類與回饋，未派 E/V**。Consumer 綁定寫 `local/plan-evidence/`（gitignored），不進版控。

| 項 | 結論 |
| --- | --- |
| N=3 適用 consumer | ✅ pass |
| `framework-charter` = project Policy | ✅ 非 dual source |
| domain-bundle `docs/domains/{domain}/` | ✅ 合法 overlay；已回寫 `domain-policies.md` §3.2 |
| ERA 已落地 | ✅ `delegated-execution.md` + yaml `delegation_loop` |
| 下一輪 | ~~consumer Phase 2 SPA scaffold O→E→V~~ → **preflight ✅**；完整 loop 待使用者 Execute |

全文：[`evidence/phase-3-external-greenfield-consumer.md`](evidence/phase-3-external-greenfield-consumer.md) · Execute：[`evidence/phase-3-external-greenfield-consumer-execute.md`](evidence/phase-3-external-greenfield-consumer-execute.md)

### 3b Automation projection（仍 open）

- [ ] 評估 `artifact-gates.yaml` / intake gate 是否可表達 `Asset class → required policy fields`
- [ ] 與 `enforcement-registry` 對齊：SD framework policy 是否需新 `rule_class`
- [ ] Test-first：validation scenario 先於 mechanical gate（見 `test-first-framework-upgrade`）

---

## Decision Rationale

### 為何不先寫 constitution.md？

- 討論已顯示 **Documentation Architecture 領先 Domain** 的風險。
- 若未來文件化，constitution 承載 **Governing Principles**，**不是** operational Policy 正文；在 Primary Model 邊界未清前，憲法會把暫定 ontology 寫死。
- 對照 [`ADR-012`](../constitution/ADR-012-route-type-activation-behavior-family.md)：ontology 與憲法混綁導致 ontology collapse。

### 為何 working hypothesis 暫用三個（但不預設結論）？

- Git / Kubernetes / DDD 的成功來自 **極少核心名詞**，不是文件多——**目標是「最小且完整」**，不是「剛好三個」。
- 現有討論中 Constitution（Principles）、Taxonomy、Derivation、Automation 均可視為核心的 **view 或投影**——待 Phase 0 驗證。
- 若驗證顯示三個不夠或太多，**依證據修訂 N**；寧可 **N 有證據**，也不要 **九個第一級目錄** 或 **為三而三**。

### 為何 Workflow 不應擁有 Asset？

- 否則每個 slice 成為小型 constitution，造成 dual source-of-truth（已在 `development-process.md` Phase 2 重構中部分修復）。
- Process 只回答 **when**；**what** 與 **how managed** 歸 Asset + Policy。

### Alternatives considered

| 方案 | 不採納原因 |
| --- | --- |
| 立即建六層目錄（Constitution / Concepts / Policies / Workflow / Automation / Derivation） | 第一級過多；新人入口模糊；over-orthogonalization |
| 只補 `artifact-placement.md` | 治標；未解 taxonomy / authority / promotion 同一套模型 |
| 直接寫 ADR | 違反 No-Proposed-ADR；domain 未驗證 |
| 把 taxonomy 寫進 `constitution.md` | ontology 演化會拖垮憲法穩定性 |

---

## Open questions

| ID | 問題 | 關閉條件 |
| --- | --- | --- |
| OQ-1 | 最小完整核心集合的 **N** 是多少？ | **已關閉：N = 3**（見 evidence/phase-0） |
| OQ-2 | `Intent` 是否為第四核心，還是 Process 的輸入？ | 分類時若 Intent 無法歸入三類則升級 |
| OQ-3 | `Automation` 獨立為核心，還是 Policy+Process 投影？ | 與 enforcement-registry 對照後決定 |
| OQ-4 | SD glossary 放 `knowledge/glossary/software-delivery.md` 還是 workflow 內？ | Phase 1；遵循 glossary owner-layer 規則 |
| OQ-5 | 與 `cognitive-slice-taxonomy` 如何共處？ | 寫清：slice = Process 載入粒度；asset class = 交付實體分類 |

---

## Acceptance gate（plan 可 archive 條件）

- [x] Phase 0 分類矩陣完成（95%，N=3）
- [x] Primary Model 圖通過「懸空盒子」檢查
- [x] **Success Criterion 滿足**：README + execution-flow pointer；零新第一級目錄
- [x] Phase 1 README 更新合併
- [x] 至少 1 次 dogfood：用模型回答 3 個新 artifact 問題（placement / owner / stage）無需開新目錄 — 見 [`evidence/phase-2-dogfood.md`](evidence/phase-2-dogfood.md)
- [x] 若 Phase 2 執行：workflow slice 無 placement dual source-of-truth（grep 驗證）— 見 dogfood evidence
- [x] Linked updates：`workflow/software-delivery/README.md`、`knowledge/glossary/software-delivery.md`

---

## 非目標（out of scope）

- 改寫整個 Ai-skill 平台三層架構（ADR-003）— 僅對齊用語
- 合併 `cognitive-slice-taxonomy` 與 delivery taxonomy — 只定義邊界
- 立即機械化所有 placement 規則 — Phase 3 選做
- 業務專案目錄強制統一為 `docs/plans/` — 專案 overlay 仍允許 `.ai-skill/project/rules/`

---

## Related

- [`workflow/software-delivery/README.md`](../workflow/software-delivery/README.md)
- [`workflow/software-delivery/artifact-gates.md`](../workflow/software-delivery/artifact-gates.md)
- [`workflow/software-delivery/contracts.md`](../workflow/software-delivery/contracts.md)
- [`governance/cognitive-slice-taxonomy.md`](../governance/cognitive-slice-taxonomy.md)
- [`governance/lifecycle/decision-promotion-pipeline.md`](../governance/lifecycle/decision-promotion-pipeline.md)
- [`constitution/ADR-012-route-type-activation-behavior-family.md`](../constitution/ADR-012-route-type-activation-behavior-family.md) — ontology vs constitution 分離教訓
- [`plans/active/2026-07-08-0825-delegation-verification-arbitration-loop/_plan.md`](active/2026-07-08-0825-delegation-verification-arbitration-loop/_plan.md) — Decision Semantics vs Workflow
- Discussion origin: 2026-07-16 software-delivery framework / constitution / placement design session
