# Phase 0 — Classification Matrix

**Run**: phase-0-2026-07-16  
**Plan**: `2026-07-16-0945-software-delivery-framework-domain-model`  
**Method**: repo 實例分類（`workflow/software-delivery/` + 直接關聯的 governance / runtime / cross-cutting）

## Primary Model 判定

| 項目 | 結論 |
| --- | --- |
| **N（核心個數）** | **3** — Asset、Policy、Process |
| **分類覆蓋** | 38/40 概念單一主歸屬（**95%**）；2 項標記為 meta/projection（不計入 N） |
| **第四核心？** | **否** — Intent、Automation、Governing Principles 不升格為與三核心並列的 managed-entity 桶 |
| **working hypothesis** | Asset/Policy/Process **成立**；無需為保留「三個」而硬塞 |

### 不計入 N 的層（meta / projection）

| 名稱 | 處置 | 理由 |
| --- | --- | --- |
| **Governing Principles** | Meta-layer（指導 Policy） | 如 One Asset One Owner；不是可操作 policy 正文，也不是 asset 實例 |
| **Automation** | Projection（Policy + Process） | validator、hook、YAML gate；機械化投影，非獨立 managed-entity 類 |
| **Intent** | Process 輸入 | 使用者請求 / 任務意圖； ephemeral，不持久化為 framework asset |
| **Asset class taxonomy** | Asset 的 ontology view | 分類 Knowledge/Contract/Evidence 等是 Asset 細分，非第四核心 |

### 曾考慮的第四核心與否證

| 候選 | 否證 |
| --- | --- |
| **Intent** | 不持久化、無 owner/lifecycle；觸發 Process 的起點，類似 Git 的「commit message」不是與 Tree/Blob 並列的儲存原語 |
| **Automation** | 所有 validator 可回溯到 underlying Policy 或 Process gate；`artifact-gates.yaml` 是 policy 投影 |
| **Principle** | 指導 Policy 制定但不直接管理 asset；文件化時在 constitution view，不擴 N |
| **Capability** | Framework **Asset**（registry 條目）；invoke 時機屬 **Process** |

## Domain 圖（Phase 0 驗收）

```text
         Governing Principles  ← meta（不計入 N）
                 │
                 ▼
              +--------+
              | Policy |  ownership · placement · promotion · authority · lifecycle
              +--------+
                 ▲   │
    Asset class  │   │ applies to
    (ontology)   │   ▼
              +--------+     Intent ──► (triggers, not stored)
              | Asset  |
              +--------+
                 │
                 │ typically produced/consumed in
                 ▼
              +---------+
              | Process |  intake → … → validation → closure
              +---------+
                 │
                 ▼
            Automation  ← projection（不計入 N）
         (validators · hooks · YAML gates)
```

## 推導鏈（domain-native）

```text
Intent → Asset (class) → Policy (which rules apply) → Process (typical stage) → Automation (if projected)
```

## 完整分類矩陣

狀態：`confirmed` = Phase 0 驗證通過；`meta` = 不計入三核心。

| # | 概念 | 主歸屬 | Asset class | Layer | 主要 Policy | 典型 Process stage | 狀態 |
| --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | change brief | Asset | Plan deliverable | Project | placement, ownership | Intake | confirmed |
| 2 | product brief | Asset | Plan deliverable | Project | ownership, authority | Intake | confirmed |
| 3 | product impact alignment | Asset | Plan deliverable | Project | authority | Intake / Requirements | confirmed |
| 4 | implementation plan | Asset | Plan deliverable | Project | placement, ownership | Planning / Implementation | confirmed |
| 5 | parity inventory | Asset | Plan deliverable | Project | ownership, traceability | Intake (replacement) | confirmed |
| 6 | BDD scenario | Asset | Contract / Behavior | Project | authority, traceability | Requirements / Test strategy | confirmed |
| 7 | Journey Specification | Asset | Contract / Behavior | Project | authority, evidence | Test strategy / Validation | confirmed |
| 8 | Domain Model Contract | Asset | Contract | Project | ownership, authority | Contracts | confirmed |
| 9 | API / Interface Contract | Asset | Contract | Project | ownership, authority | Contracts | confirmed |
| 10 | UI / Screen / Consumer Contract | Asset | Contract | Project | ownership, authority | UI contracts | confirmed |
| 11 | Error Handling Contract | Asset | Contract | Project | ownership, authority | Contracts | confirmed |
| 12 | Screen Mapping | Asset | Contract (traceability) | Project | traceability | UI contracts | confirmed |
| 13 | Evidence artifact | Asset | Evidence | Project | placement, authority | Validation | confirmed |
| 14 | UI governance evidence | Asset | Evidence | Project | placement, severity | UI governance / Validation | confirmed |
| 15 | Journey validation result | Asset | Evidence | Project | authority | Validation | confirmed |
| 16 | Review report | Asset | Evidence | Project | placement | Review invoke | confirmed |
| 17 | ADR (`constitution/`) | Asset | Decision | Framework | promotion, authority | Platform governance | confirmed |
| 18 | project decision (`docs/decisions/`) | Asset | Decision | Project | promotion | Closure | confirmed |
| 19 | Template (`templates/`) | Asset | Template | Framework | placement | Any (shape reuse) | confirmed |
| 20 | Pattern entry (`ui-pattern-knowledge/`) | Asset | Knowledge | Framework | promotion, sanitization | UI contracts (when stable) | confirmed |
| 21 | Glossary term | Asset | Knowledge | Framework | evolution | Pre-build interrogation | confirmed |
| 22 | Failure pattern (referenced) | Asset | Knowledge | Framework | promotion | Closure / feedback | confirmed |
| 23 | Feedback lesson | Asset | Knowledge | Framework | promotion, sanitization | Closure | confirmed |
| 24 | `sd-*` cognitive slice | Asset | Process component | Framework | loading budget | Process (load surface) | confirmed |
| 25 | `execution-flow.md` | Process | — | Framework | — | Lifecycle index | confirmed |
| 26 | `intake.md` slice | Process | — | Framework | — | Intake | confirmed |
| 27 | Incident observe→classify→layer | Process | — | Framework | — | Incident path | confirmed |
| 28 | Capability invoke (review) | Process | — | Framework | invoke policy | Post-implementation | confirmed |
| 29 | Delegation loop | Process | — | Framework | role matrix policy | Delegated execution | confirmed |
| 30 | DoR / DoD closure | Process | — | Framework | completeness policy | Closure | confirmed |
| 31 | Contract precedence table | Policy | — | Framework | authority | Contracts (conflict) | confirmed |
| 32 | Placement rules (fragment) | Policy | — | Framework / Project | placement | Any asset create | confirmed |
| 33 | Promotion pipeline rules | Policy | — | Framework | promotion | Retrospective | confirmed |
| 34 | Reusable guidance boundary | Policy | — | Framework | knowledge boundary | Any writeback | confirmed |
| 35 | `artifact-gates` quality criteria | Policy | — | Framework | artifact quality | Same-session closure | confirmed |
| 36 | Recovery policy (`metadata/recovery/`) | Policy | — | Framework | recovery | Mismatch escalation | confirmed |
| 37 | `execution-flow.yaml` / gates | Automation | — | Framework | — | Runtime activation | meta |
| 38 | commit-msg / stop hook validators | Automation | — | Framework | — | Commit / close-out | meta |
| 39 | Governing Principles | Meta | — | Framework | — | — | meta |
| 40 | User / task Intent | Process input | — | — | — | Process entry | meta |

## Asset class 穩定枚舉（ontology — 可演化）

| Class | 範例實例 |
| --- | --- |
| Plan deliverable | change brief, implementation plan, parity inventory |
| Contract / Behavior | BDD, Journey Spec, domain/API/UI contracts |
| Evidence | validation output, screenshots, journey proof |
| Decision | ADR, project decision record |
| Template | `templates/*.md` |
| Knowledge | pattern entry, glossary, failure pattern |
| Process component | `sd-*` slice definition |
| Runtime | `execution-flow.yaml`, `artifact-gates.yaml`（亦為 Automation source） |
| Capability | `capability-registry` entry（cross-cutting） |

**Artifact**（term）：Project 層為完成 delivery intent 而產生的 Asset **實例**（tag），非與 Asset 並列的核心。

## 邊界案例記錄

| 案例 | 判定 | 說明 |
| --- | --- | --- |
| `cognitive_slice` taxonomy (4 types) | Process 載入粒度 | ADR-009；≠ delivery Asset taxonomy |
| `sd-*` slice | Asset (Process component) | slice 是可管理的 framework 定義；載入時機屬 Process |
| Runtime YAML | Asset (Runtime class) + Automation projection | 雙視角合法；核心歸 Asset |
| Review | Process moment + Evidence output | invoke = Process；report = Asset |
| Intelligence atom | Asset (Knowledge) @ `intelligence/` | SD framework 引用，不擁有 |

## Phase 0 否證結論

- **三核心足夠**：95% 單一主歸屬；例外為 meta/projection，有明確處置。
- **不需第四核心**：Intent / Automation / Principles 已解釋為 meta 或 projection。
- **不需減為二核心**：Policy 無法併入 Asset（橫切多 class）或 Process（非時序）。
- **下一步（Phase 1）**：README 增加 Domain 圖 + 推導鏈；**不**新增第一級目錄。

## Dogfood 抽樣（新人 5 分鐘題）

| 問題 | 答案（用三核心推導） |
| --- | --- |
| change brief 是什麼？ | **Asset**（Plan deliverable，Project 層） |
| 誰擁有？ | **Policy**（ownership：發起人 / PM；One Asset One Owner） |
| 通常何時產生？ | **Process**（Intake stage） |
| 放哪？ | **Policy**（placement；專案 overlay 可覆寫）— 現有碎片見 `artifact-gates.md`，Phase 2 才收斂 |
