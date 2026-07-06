# ADR-013: Cognitive Role Primitive Gate (Review as First Consumer)

## Status

**Proposed**

## Framework Generation

- **世代分類**：Gen 3 — Cognitive Execution System 子系統邊界擴充候選
- **當前世代文件**：[`architecture/ai-native-cognitive-execution-system.md`](../architecture/ai-native-cognitive-execution-system.md)；[`constitution/ADR-008-runtime-cognitive-modes.md`](ADR-008-runtime-cognitive-modes.md)（既有 4 維 `cognitive_mode` primitive）
- **適用狀態**：本 ADR **不**立即引入 runtime primitive 或 workflow 目錄變更。Accept 前僅鎖定決策框架與 promotion gate。Software-delivery review gap 為觸發證據，**不是**自動批准 D1 的理由。

## Date

2026-07-06

## Source Plan

- [`plans/active/2026-07-06-review-architecture-adr/_plan.md`](../plans/active/2026-07-06-review-architecture-adr/_plan.md)
- [`plans/active/2026-07-06-review-architecture-adr/01-phase0b-perspective-generalization-evidence.md`](../plans/active/2026-07-06-review-architecture-adr/01-phase0b-perspective-generalization-evidence.md) — Phase 0b evidence
- [`plans/active/2026-07-06-review-architecture-adr/02-phase0b5-perspective-taxonomy-validation.md`](../plans/active/2026-07-06-review-architecture-adr/02-phase0b5-perspective-taxonomy-validation.md) — Phase 0b.5 perspective taxonomy (in progress)

## Context

### Trigger: software-delivery 三層契約不一致

| 層 | 現況 |
|---|---|
| **Governance** | [`software-delivery-governance.md`](../governance/ai-runtime-governance/software-delivery-governance.md) — Artifact completeness 含 review report |
| **Execution** | [`execution-flow.yaml`](../workflow/software-delivery/execution-flow.yaml) — 無 review loading surface；步驟在 implementation 後直接 `validate_and_close` |
| **README** | [`workflow/software-delivery/README.md`](../workflow/software-delivery/README.md) — review-checklist 為「需要時參考」 |
| **Routing** | [`routing-registry.yaml`](../knowledge/runtime/routing-registry.yaml) — `code-review` / `design-review` triggers 存在，required_dependencies **不含** review surface |

### Deeper gap: AI 無 Reviewer 執行模型

Agent 在 software-delivery 任務中常見路徑：

```text
Implement → Run tests → Finish
```

缺少可發現的 **persona 切換**（停止實作、findings-only、產 review report）。此缺口觸發本 ADR，但 **解法是否為新 Runtime Primitive 尚未證明**。

### Review 的 lifecycle 分布（反對單一 Workflow Phase）

| Review 類型 | 典型時機 | Caller 上下文 |
|---|---|---|
| Contract / API design review | 平行實作前 | `sd-contracts` / `sd-ui-contracts` |
| Architecture review | 架構決策後 | `architecture/` |
| Security review | API/Auth 設計後 | contracts / implementation |
| Code review | 實作完成後 | `sd-implementation` |
| Performance review | PR / perf-sensitive | `sd-test-strategy` / perf-risk-gate |
| Release review | 發布/合併前 | `sd-validation` / `sd-closure` |

**共同點不是同一時間點，而是同一類行為：以 Reviewer 視角執行檢查並產可追溯 artifact。**

### 與 ADR-008 的關係

[`cognitive_mode`](../models/cognitive-modes/README.md)（ADR-008）描述 **如何**執行（FAST/DEEP、SUMMARY_FIRST、STRICT …）。  
本 ADR 評估的是 **是否**需要獨立的 **誰在執行** primitive（`cognitive_role`），或是否可用 **capability invoke + context** 表達。

兩者應保持 **正交**（若 D1 成立）：

```text
cognitive_role: Reviewer
cognitive_mode: DEEP + SOURCE_BACKED + STANDARD + NONE
capability: code-review
```

---

## Primary Question

> **Does the Cognitive System need a new Runtime Primitive: Cognitive Role?**

| 答案 | 含義 |
|---|---|
| **Yes (D1)** | `cognitive_role` 與 `cognitive_mode` 並列；Review 為 **第一個 consumer**；Planner / Debugger 等可共享 |
| **No (D2 / B)** | Persona 以 **capability context** 表達；不擴張 routing/runtime/graphs/overlays 的 primitive 表面 |

「Review 是什麼？」是 **derivative**。「Role 要不要成為 primitive？」是 **primary**。

---

## Decision Tree

```text
Q1: Is Review a fixed Workflow Phase (single lifecycle slice)?
    Yes → Option A (sd-review) — evaluate
    No  → continue

Q2: Does persona separation require a new Runtime Primitive?
    Yes AND generalization test passes → D1 (cognitive_role primitive)
    No OR review-only special case      → D2 (capability + context) → B fallback
```

```mermaid
flowchart TD
  q1{Q1_Review_fixed_Workflow_Phase}
  q2{Q2_Persona_needs_Runtime_Primitive}
  A[Option_A_reject_expected]
  D1[D1_Role_Primitive]
  D2[D2_Capability_Context]

  q1 -->|Yes| A
  q1 -->|No| q2
  q2 -->|Yes_generalizes| D1
  q2 -->|No| D2
```

---

## Five Required Answers

| # | Question | Draft answer | Evidence / notes |
|---|---|---|---|
| 1 | Review 是否跨 Workflow？ | **Yes** | 表見 Context §lifecycle 分布 |
| 2 | Review 是否需要 Persona 切換？ | **Yes** | 停止 feature coding；findings-only；report artifact |
| 3 | Persona 是否需要 **Runtime Primitive**？ | **No (leaning)** | Phase 0b — `context.perspective` sufficient; see Generalization Test |
| 4 | Capability 是否足以表達 Reviewer Context？ | **Yes** | D2 invoke envelope |
| 5 | 新 Primitive 是否泛化到其他 Domain？ | **No** | Phase 0b — slice/mode/objective cover most activities |

---

## Generalization Test

> **Phase 0b complete** — full evidence: [`plans/active/2026-07-06-review-architecture-adr/01-phase0b-perspective-generalization-evidence.md`](../plans/active/2026-07-06-review-architecture-adr/01-phase0b-perspective-generalization-evidence.md)

### Phase 0b method

- Test **non-Review** activities and **counter-examples** (Refactoring, Documentation, Test authoring)
- Success = define **Role boundary**, not only pro-Role cases
- Separate: **Review ≠ Workflow Phase** (proven) vs **Role ≠ Runtime Primitive** (tested here)

### Activity × placement matrix (summary)

| Activity | Workflow phase? | Capability enough? | Role more natural? | Conclusion |
|---|---|---|---|---|
| Review | No | △ needs `perspective` | Yes | **Perspective candidate** |
| Planning | No (slice owns) | Yes | △ overlaps intake | **Capability / slice** |
| Debugging | No | △ + FORENSIC mode | △ **not falsified** | **Perspective candidate — Phase 0b.5** |
| Architecture fit | No (slice owns) | Yes | No | **Slice** |
| Validation | No (slice owns) | Yes | No | **Slice** |
| Refactoring | No | Yes (`execution_mode`) | No | **Counter-example** |
| Documentation | No | Yes | No | **Counter-example** |
| Test authoring | No | Yes | △ weak | **Counter-example** |

### Promotion rule outcome (Phase 0b)

- **≥3 activities share stable role primitive semantics?** → **No**
- **Only Review (+ Debugger under test) show perspective switch?** → **Under Phase 0b.5**
- **Strong counter-examples?** → Refactoring, Documentation, Planning (slice), Architecture (slice)

**Phase 0b updates Q3 / Q5:**

| # | Question | Phase 0b answer |
|---|---|---|
| 3 | Persona needs Runtime Primitive? | **No (leaning)** — `context.perspective` sufficient |
| 5 | Generalizes to other domains? | **No** — slices / modes / objectives already cover most |

### Role taxonomy explosion risk

若 Role 成 primitive 而無 bounded catalog，預期膨脹：

`Tester → Refactorer → Optimizer → Writer → Documenter → …`

D1 **必須**附帶 **bounded role catalog policy** 與新增 role 的 ADR/plan gate，否則 reject。

---

## Evaluation Criteria

| Criteria | 說明 |
|---|---|
| Lifecycle Independence | 跨 Requirements / Architecture / Implementation / Release 的 review |
| Cognitive Separation | Reviewer vs Implementer 真正分離 |
| Discoverability | AI 知何時 invoke review |
| Extensibility | 新增 UX/Compliance review 成本 |
| Contract Clarity | workflow contract 保持 thin |
| Runtime Cost | routing/loading 維護 |
| **Primitive Justification** | 獨立語意 + 跨 domain 泛化 |

### Scoring matrix (Phase 0 draft — not final)

Legend: `S` strong · `A` adequate · `W` weak · `F` fails · `O` open

| Criteria | A | B | C | D1 | D2 |
|---|---|---|---|---|---|
| Lifecycle Independence | W | S | W | S | S |
| Cognitive Separation | A | W | A | S | A |
| Discoverability | W | A | W | S | A |
| Extensibility | W | S | W | S | S |
| Contract Clarity | W | S | F | S | S |
| Runtime Cost | A | S | W | W | S |
| Primitive Justification | F | A | F | **F** | S |
| **Overall (Phase 0b)** | reject | fallback | reject | **reject** | **recommend** |

> Phase 0b evidence: [`01-phase0b-perspective-generalization-evidence.md`](../plans/active/2026-07-06-review-architecture-adr/01-phase0b-perspective-generalization-evidence.md)

---

## Options

### Option A — Workflow Phase (`sd-review`)

- Review = 單一 lifecycle slice + post-implementation Review Gate
- **預期 reject**：pre-impl review 錯置；lifecycle 綁定錯誤

### Option B — Cross-cutting Capability

- `workflow/cross-cutting/review/` + per-slice `review_invocation`
- 無 persona primitive
- **Fallback** 若 D1 未通過

### Option C — Hybrid (slice hook + capability body)

- **退化路徑**：Hook → Capability → Review Workflow → Gate → 回到 Option A
- **預期 reject** 除非 anti-degeneration 邊界可證明

### Option D1 — Role as Runtime Primitive

```text
Workflow → cognitive_role (primitive) → review capability → artifact
```

- 新增 `governance/cognitive-role.md`（living spec，accept 後）
- routing / runtime / graphs / overlays **須**理解 role — **高承諾**
- **條件接受**：Q2=Yes 且 Generalization Test 通過且 bounded catalog

### Option D2 — Role as Capability Context

```text
Workflow → review capability → context.perspective → artifact
```

```yaml
# Illustrative invoke envelope — not canonical until accept
invoke:
  capability: code-review
  context:
    perspective: reviewer
    caller_slice: sd-implementation
```

- **Role 不是第一級概念**
- 與 `cognitive_mode` 正交：`perspective` = 以誰的視角執行 capability
- **Leading fallback** in Phase 0 draft

---

## Primitive Promotion Cost (D1 only)

若 `cognitive_role` 晉升為 runtime primitive，下列表面 **必須**同步設計（不可 silent 擴張）：

- `knowledge/runtime/routing-registry.yaml` — activation / loading
- `runtime/runtime.db` — 若未來投影（需另 plan）
- `knowledge/graphs/*` — workflow edges
- project overlays — Brower 等 consumer
- glossary — `cognitive_role` vs `cognitive_mode` 不混淆

**Primitive 一旦建立很難刪。** 本 ADR 在 accept D1 前不寫入上述任何機械表面。

---

## Review Type × Invocation Point (shared by D1/D2)

| Review type | Caller workflow | D1 role switch | D2 invoke |
|---|---|---|---|
| Contract review | `sd-contracts` | Designer → Reviewer | `contract-review` + `perspective: reviewer` |
| Architecture review | `architecture/` | Architect → Reviewer | `architecture-review` + `perspective: reviewer` |
| Security review | contracts / impl | Implementer → Reviewer | `security-review` + `perspective: reviewer` |
| Code review | `sd-implementation` | Implementer → Reviewer | `code-review` + `perspective: reviewer` |
| Performance review | test-strategy / perf-risk-gate | Implementer → Reviewer | `performance-review` + `perspective: reviewer` |
| UI / compliance review | `sd-ui-governance` | Implementer → Reviewer | `ui-review` + `perspective: reviewer` |
| Release review | `sd-validation` / `sd-closure` | Validator → Reviewer | `release-review` + `perspective: reviewer` |

---

## Recommendation (Phase 0c stakeholder input — not final D2 acceptance)

> **Review is not a Workflow Phase** — proven (Phase 0a).
>
> **Reject D1 (Accepted):** Cognitive Role as Runtime Primitive is **not justified** by Phase 0b generalization evidence. The rejection is because **Role does not generalize as a runtime primitive**, not because Review is special.
>
> **D2 as working model (Accepted provisionally):** cross-cutting review capabilities with bounded `context.perspective`. Closes the software-delivery review gap **without** promoting `cognitive_role` to a runtime primitive. Phase 1 **may plan** against this envelope; implementation treats `perspective` enum as **draft** until Phase 0b.5 + 0d.
>
> **D2 final acceptance (Deferred):** `context.perspective: reviewer` is **not** finalized. Phase 0b.5 tests whether Review, Debugger, Incident Analysis, and Security Audit share a higher-order perspective (e.g. `fault_finding` / negative evidence seeking). See [`02-phase0b5-perspective-taxonomy-validation.md`](../plans/active/2026-07-06-review-architecture-adr/02-phase0b5-perspective-taxonomy-validation.md).
>
> **Research priority:** Perspective taxonomy — not Role — is the open design space. A future ADR (e.g. ADR-014) may own perspective boundaries if validation confirms a stable abstraction.

Phase 0a draft leaned D2; **Phase 0b strengthens reject-D1 / provisional-D2**; **Phase 0c accepts reject-D1 + provisional-D2, defers final D2 accept**.

---

## Decision

**Partial (Phase 0c)** — Status remains **Proposed** until Phase 0b.5 + 0d close.

| Outcome | Phase 0c status |
|---|---|
| **Reject D1** | **Accepted** |
| **D2 working model** | **Accepted provisionally** |
| **D2 final acceptance** | **Deferred** — pending perspective taxonomy validation |

Acceptance outcomes (after Phase 0b.5 + 0d):

| Outcome | Action |
|---|---|
| **Accept D1** | Author `governance/cognitive-role.md`; plan runtime/routing extension; Phase 1 implementation — **requires new counter-evidence** |
| **Accept D2 (final)** | Author `governance/review-capability.md` with validated `context.perspective`; Phase 1 without new primitive |
| **Accept B only** | Capability catalog without formal context envelope (weaker persona separation) |
| **Reject A, C** | Document as rejected paths |

---

## Consequences

### If D1 accepted

- Review = first `cognitive_role` consumer
- `workflow/cross-cutting/review/` — capability bodies
- Per-slice `role_switch` + `review_invocation` hooks (thin)
- Glossary: `cognitive_role` (distinct from `cognitive_mode`)
- **No** `sd-review` lifecycle slice

### If D2 accepted (leading fallback)

- `workflow/cross-cutting/review/` — capability bodies
- Per-slice `review_invocation` with `context.perspective`
- Glossary: `review_capability`, `capability_context` (or ADR-finalized term)
- **No** routing/runtime primitive expansion
- **No** `sd-review` lifecycle slice

### Either path — shared fixes

- Migrate [`review-checklist.md`](../workflow/software-delivery/review-checklist.md) under cross-cutting (ADR accept 後)
- Fix [`validation.md`](../workflow/software-delivery/validation.md) misplaced review-report-template reference
- Align governance Artifact completeness language
- Validation scenarios: pre-impl architecture review · post-impl code review · pre-release release review

### Rejected (Phase 0)

- Option A — single post-impl `sd-review` slice as primary model
- Option C — hybrid without anti-degeneration proof
- Review-only primitive without generalization test pass

---

## Related

- [`constitution/ADR-008-runtime-cognitive-modes.md`](ADR-008-runtime-cognitive-modes.md)
- [`constitution/ADR-009-cognitive-slice-taxonomy.md`](ADR-009-cognitive-slice-taxonomy.md)
- [`workflow/cross-cutting/README.md`](../workflow/cross-cutting/README.md)
- [`governance/cognitive-slice-taxonomy.md`](../governance/cognitive-slice-taxonomy.md)
- [`knowledge/glossary/ai-skill.md`](../knowledge/glossary/ai-skill.md) — `validation_capability`（cross-cutting capability 先例）
- [`workflow/software-delivery/review-checklist.md`](../workflow/software-delivery/review-checklist.md)

## Open Questions (Phase 0b.5 — perspective taxonomy)

1. Do Review, Debugger, Security Audit, and Incident Analysis share one perspective (`fault_finding`)?
2. If yes — is `reviewer | debugger | auditor` the wrong enum layer (consumer labels vs perspective)?
3. D2 vs B — is formal context envelope required, or capability-only sufficient?
4. Peer review vs self-review — same capability + different context flag?
5. Mechanical enforcement — advisory vs transition-block (deferred to Phase 1 plan)
6. ~~Generalization Test / Reject D1~~ — **closed in Phase 0b** (accepted Phase 0c)
