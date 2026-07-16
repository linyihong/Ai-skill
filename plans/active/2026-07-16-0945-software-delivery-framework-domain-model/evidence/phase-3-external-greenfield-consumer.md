# Phase 3 External Dogfood — Greenfield Consumer（2026-07-16）

**Run**: phase-3-external-greenfield-consumer  
**Plan**: `2026-07-16-0945-software-delivery-framework-domain-model`  
**Consumer**: `<PROJECT_ROOT>`（sibling greenfield；離線 portable yaml + framework charter）  
**Run boundary（本機，不進版控）**：`<AI_SKILL_REPO>/local/plan-evidence/`（見 evidence README）  
**Method**: ERA 式外部 repo 觀察 + Asset/Policy/Process 分類（比照 [`2026-07-08-0825-delegation-verification-arbitration-loop`](../../2026-07-08-0825-delegation-verification-arbitration-loop/_plan.md) dogfood 證據慣例；**本 run 為 doc-only 分類，未派 E/V**）

---

## 1. 假說

若 N=3（Asset / Policy / Process）為最小完整模型，則**已採 Ai-skill 交付紀律的 greenfield 專案**應能：

1. 用三桶解釋現有目錄，無需新增第一級概念；
2. 專案 overlay（charter + portable YAML）對應 Policy + Automation projection，不與 framework `domain-policies.md` 衝突；
3. 與 ERA 三角色 loop 可並存（Process 子視圖 + 獨立 loop 契約）。

---

## 2. Consumer 三桶對照（generalized）

| 桶 | 典型表面 | 角色 |
| --- | --- | --- |
| **Policy** | `docs/architecture/framework-charter.md`、`docs/workflow/software-delivery.yaml`（`delivery_policy`、`document_priority`、`change_classification`）、`pre-commit-validation.md` | placement、命名、authority、禁止事項、離線流程契約 |
| **Asset** | `docs/domains/<domain>/*`、`plans/active/`、`openapi/`、`docs/Reference Project/analysis/*`（唯讀） | 契約、brief、spec、執行計畫、機器契約 |
| **Process** | yaml `phases:`、`docs/workflow/delegated-execution.md` `delegation_loop` | intake → contracts → implementation → validation；Execute 時 O→E→V |

```text
Intent（使用者 / plan phase）
    → Asset class（domain bundle、plan、openapi）
    → Policy（charter + yaml rules）
    → Process（phases + optional delegation loop）
    → Automation（scripts/verify、git hooks — projection）
```

**結論**：三桶 **95%+ 單一主歸屬**；portable yaml 同時承載 Policy 片段與 Process phase 表 → 印證 **Automation/YAML = projection**，非第四核心。

---

## 3. Artifact 分類矩陣（摘錄）

| Artifact | Asset class | Placement（consumer） | domain-policies §3.2 預設 | Owner | Process stage |
| --- | --- | --- | --- | --- | --- |
| `change-brief.md` | Plan deliverable | `docs/domains/<domain>/` | `docs/planning/` | PM / 發起人 | phase.intake |
| `api-contract.md` | Contract | `docs/domains/<domain>/` | `docs/contracts/` | 架構負責人 | phase.contracts |
| `screen-mapping.md` | Contract (traceability) | `docs/domains/<domain>/` | `docs/contracts/` | 全端 / UX | phase.contracts |
| `domain-invariants.md` | Contract | `docs/domains/<domain>/` | `docs/contracts/` | 域負責人 | phase.contracts |
| domain spec | Plan / Spec | `docs/domains/<domain>/` | `docs/planning/` | 產品 | intake 前 |
| `_plan.md` | Plan deliverable | `plans/active/` | `plans/` ✓ | plan owner | 全 phase |
| `framework-charter.md` | **Policy**（非 Asset） | `docs/architecture/` | overlay（見 §2） | 專案架構 | 實作前必讀 |
| `software-delivery.yaml` | Process + Policy projection | `docs/workflow/` | — | project | 全 lifecycle |
| `delegated-execution.md` | Process component | `docs/workflow/` | Ai-skill 鏡像 | orchestrator 紀律 | Execute |
| OpenAPI artifact | Contract (machine) | `openapi/` | 合理擴充 | API owner | implementation+ |
| Reference analysis | Knowledge (read-only) | `docs/Reference Project/analysis/` | N/A（專案證據邊界） | 分析者 | intake / research |

---

## 4. 與 framework `domain-policies.md` 的落差

| 觀察 | 嚴重度 | 建議 |
| --- | --- | --- |
| **Domain bundle** 把 brief + contracts 共置 `docs/domains/{domain}/` | 預期內 overlay | framework §3.2 已補 domain-bundle |
| 專案稱 `framework-charter` 為「框架憲章」 | 用語 | charter = **Policy 正文**；Principles 在禁止清單背後 |
| `implementation plan` 預設 `docs/plans/` vs `plans/active/_plan.md` | 小 | 語意一致 |
| `docs/evidence/` 尚未使用（pre-code） | — | scaffold 階段應啟用 |
| Authority：`document_priority` 在 YAML 頂層 | 正面 | 與 `domain-policies` §4 相容 |

---

## 5. 與 ERA plan 的接點

Consumer **已採用** ERA 消費者模式：

| ERA 元素 | Consumer 落地 |
| --- | --- |
| 三角色 loop | `docs/workflow/delegated-execution.md` + yaml `delegation_loop` |
| Execute mandatory | overlay README + yaml `mandatory_when` |
| Brief + backfill | active plan 執行段（下一 phase scaffold 待跑） |
| 專案證據邊界 | generalized metrics 留 Ai-skill；class 名 / live 留 `<PROJECT_ROOT>` plan |

**建議下一輪 dogfood（未執行）**：active plan 下一 phase **SPA scaffold** 走完整 O→E→V；Orchestrator 步驟 0 強制 **Asset class 分類**。可登記 ERA **2v** 或本 plan **phase-3-external-greenfield-consumer-execute**。

---

## 6. 回饋摘要（framework ← consumer）

### 6.1 模型驗證（正面）

- N=3 足以解釋 consumer，**零新第一級目錄**假說成立。
- portable YAML 是 Process+Policy 離線投影 — Automation 不升格。
- `framework-charter` = **project Policy canonical**；與 Ai-skill `domain-policies.md` 為 reference/overlay，非 dual source。

### 6.2 Framework 改進

| 項 | 動作 |
| --- | --- |
| domain-bundle overlay | ✅ `domain-policies.md` §3.2 |
| OQ-2 Intent | yaml `phases[].actions` 把 intent 當 ephemeral gate 輸入 |
| OQ-5 cognitive-slice | consumer 用 yaml phases；slice = framework 粒度，phase = project 粒度 |

### 6.3 Consumer 建議（留 `<PROJECT_ROOT>`，不寫入 Ai-skill）

1. charter 開頭加 Primary Model 指針（charter=Policy；domains=Asset bundle；yaml=Process）。
2. Execute 前在 plan 補 delegation brief。
3. 首個 `docs/evidence/` 對齊 ERA verifier finding 欄位。

---

## 7. 量測欄

| 欄位 | 值 |
| --- | --- |
| run_kind | external_observational |
| e_v_loop | 未執行 |
| n_equals_3_fit | pass |
| placement_dual_write | pass |
| era_consumer_adopted | yes |
| framework_patch | domain-policies §3.2 domain-bundle note |
