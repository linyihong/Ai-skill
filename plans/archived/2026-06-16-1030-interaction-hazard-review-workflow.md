---
id: 2026-06-16-1030-interaction-hazard-review-workflow
plan_kind: main
status: completed
owner: linyihong
created: 2026-06-16
priority: P1
required_for_completion: false
---

# State Trust Transition — Promotion Discipline (Decision Framework)

**Status**: `completed`
Owner: framework maintainer (linyihong)
**建立日期**：2026-06-16
**最後修訂**：2026-07-13（補 ECS Phase 1 冷啟動貢獻互指；Plan Completion Closure 已完成）
**Priority**：**P1**
**Plan archive**：✅ archived 2026-07-13
**Downstream pilot**（canonical evidence consumer）：<AI_SKILL_DOGFOOD_EVIDENCE> `docs/plans/archived/2026-06-16-state-trust-transition-pilot.md`（**completed** 2026-07-13）；C.5/ADR#4 @ `6a7cc1c` + `docs/plans/c5-trials/2026-07-13-payment-leave-confirm-dialog.yaml`

## Executive summary

本 plan 已从「设计 workflow」转为 **升级判准（promotion discipline）**：

```text
primitive → evidence → consumer → 是否值得升 workflow
```

**不是**：问题 → 建 workflow。

**Primitive / consumer 分离**（本版最成熟处）：

- **Primitive**：State Trust Transition table（Ownership Map + Invalidation + Recovery）
- **Evidence**：`temporal_behavior` 子形状（event_trace / dom_presence / ownership）
- **Consumer**：software-delivery、UI overlay、integration claims — **仅为消费者**

**真正在验的题**（不是 O2/O3 标签本身）：

```text
同一 trust transition（invalidate → recover）能否形成最小闭环？
```

若 Dialog、rollback、websocket 都能用 **同一四栏、不改字** 表达 trust lost → trust restored ⇒ lean **O3**（generic model）。否则 lean **O2**（UI conditional gate）。

**Blocking decision**（O1 = future promotion only）：

| Path | 选 when | Pilot status (2026-06-16) |
|---|---|---|
| **O2 — Conditional gate** | B + B.5 失败：栏位不能跨域复用 | not selected |
| **O3 — Generic trust transition model** | B + B.5 通过：四栏不改字可套 rollback + websocket | **final** — downstream C.5/ADR#4 pass 2026-07-13 |

---

## Stage Decision（O3 — **final** 2026-07-13）

Downstream pilot closed Criterion 4 with a deliberate C.5 Predictive Interception Trial (PaymentLeaveConfirmDialog). This plan now records **final O3**.

| Layer | Status |
|---|---|
| Primitive | Mature / Stable |
| Evidence model | Validated |
| Promotion discipline | Established |
| O2 vs O3 | **Final O3** |
| Workflow / runtime surface graduation | **Still deferred** — no new lifecycle phase; evidence accumulation, not framework expansion |

Accordingly:

- **O2** is not selected.
- **O3** is the **final** direction for the State Trust Transition primitive.
- Criterion 4 evidence: <AI_SKILL_DOGFOOD_EVIDENCE> `docs/plans/c5-trials/2026-07-13-payment-leave-confirm-dialog.yaml` + L1 @ test `/h5` `6a7cc1c` (`adr4_verdict: pass`).
- **Do not** manufacture additional “graduation features”; further work is consumer evidence and optional naming ADR — not new framework design unless evidence invalidates the primitive.

**Governance invariant:** No Experience Runtime lifecycle row, no `ownership_map` rename, no Interaction Evidence Hierarchy extraction from this decision alone. Future work focuses on **evidence accumulation, not framework expansion**.

---

## Downstream pilot gate

Phase 勾选 **不以 Ai-skill 自证**；以 downstream validation pilot 产物 + commit 为准（[`reusable-guidance-boundary.md`](../../enforcement/reusable-guidance-boundary.md)）。

| Phase | Gate pass when | Downstream evidence | Status |
|---|---|---|---|
| A0 | 四栏 + Recovery Boundary 定稿并 sync 到 consumer workflow | <AI_SKILL_DOGFOOD_EVIDENCE> `bcce737` — `framework-development-workflow.yaml`, `interaction-hazard-review.md` | **pass** |
| A1 | Coupon 四栏 trust table | <AI_SKILL_DOGFOOD_EVIDENCE> `6665b77` — `screen-mapping/episode-coupon-redeem-journey.md` | **pass** |
| B | 非 player invalidate↔recover 闭环 | <AI_SKILL_DOGFOOD_EVIDENCE> `6665b77` — `screen-mapping/membership-payment-sync-trust-journey.md` | **pass** |
| B.5 | 四栏不改字压力测试 | <AI_SKILL_DOGFOOD_EVIDENCE> `6665b77` — `screen-mapping/websocket-subscription-trust-journey.md` | **pass** |
| C | field survival scenario + predictive prevention | <AI_SKILL_DOGFOOD_EVIDENCE> BDD + C.5 PaymentLeaveConfirmDialog | **pass** |
| D | 全部 ADR criteria + O2/O3 final | <AI_SKILL_DOGFOOD_EVIDENCE> archived pilot §D final **O3** | **pass** (runtime surface still deferred) |

**Promotion gate summary**（ADR criteria 1–6）：

| # | Criterion | Status |
|---|---|---|
| 1 | ≥2 cases, four-column table | **pass** (coupon + payment sync) |
| 2 | ≥1 validation scenario consumes trust evidence | **pass** (BDD + coupon journey + payment-leave L1 browser) |
| 3 | O2 or O3 resolved | **final O3** |
| 4 | ≥1 previously unknown prevention | **pass** — <AI_SKILL_DOGFOOD_EVIDENCE> C.5 PaymentLeaveConfirmDialog / test `/h5` `6a7cc1c` |
| 5 | template field survives renaming | **pass** (B.5 websocket, same headers) |
| 6 | no rubber-stamp | **pass** — refused continuation dogfood as ADR #4 substitute |

---

## Core primitive — State Trust Transition

### Ownership Map（四栏 — A0 定稿）

| State | Owner | Invalidation Event | Recovery Boundary |
|---|---|---|---|

**Invalidation Event** — any event after which the state must no longer be trusted.

**Recovery Boundary** — what evidence makes this state trustworthy again.

```text
trust lost  →  trust restored
```

| State | Owner | Invalidation Event | Recovery Boundary |
|---|---|---|---|
| `playbackAllowed` | Entitlement | refresh | grant readback |
| `optimisticBalance` | QueryCache | rollback | server sync |
| `dialogOpen` | CouponPanel | unmount | explicit reopen |
| `websocketReady` | Connection | reconnect | handshake complete |

**Why Recovery Boundary matters**：invalidated ≠ consumer must immediately stop — e.g. refresh 后是否可暂时信、等 readback？无 Recovery 栏，Phase B 只能描述问题，不能描述 **结束**。

Hazard class：`owner-invalidation-before-complete` when recovery boundary crossed while async work still in flight.

Ownership 是 trust transition 的 **一个来源**，不是全部 — 模型名倾向 **State Trust Transition**，Ownership Map 为表格名。

---

## Promotion discipline（ADR criteria）

Phase D graduation **全部**满足：

1. ≥ **2** independent cases，四栏 finalized template 填完
2. ≥ **1** validation scenario 机械消费 interception / trust evidence
3. **O2 or O3** resolved（B + B.5 证据）
4. ≥ **1 previously unknown prevention** — not post-hoc explanation of shipped bugs
5. ≥ **1 template field survives renaming pressure** — 若 Dialog → rollback → websocket 后 **Invalidation Event** 与 **Recovery Boundary** 列名仍存在、不需改栏位 ⇒ primitive 有生命力；若每案改栏 ⇒ abstraction noise，**勿 O3**
6. No sustained rubber-stamp on empty sections

```text
Post-hoc explanatory power ≠ predictive interception
Abstraction that renames every case ≠ primitive
```

---

## Evidence（consumer 层 — 不膨胀 taxonomy）

```yaml
temporal_behavior:
  event_trace:
  dom_presence:
  ownership:    # trust boundary / no-invalidate-before-complete
```

**新增 consumer ≠ 新增 taxonomy.**

Governance invariant（consumer）：`observable_outcome_must_survive_owner_refresh` — outcome trusted across refresh window until recovery boundary evidence arrives.

---

## Evidence Rule

> Machine-readable evidence-rule（schema `evidence-rule-v1`），索引於
> [`governance/evidence-candidates/evidence-rules/interaction-hazard.pointer.yaml`](../../governance/evidence-candidates/evidence-rules/interaction-hazard.pointer.yaml)。
> **Phase 1A Step 2（consumer attach）**：本 section 成立 = consumer hook 建立；criterion 內容是
> **Step 3（criteria authoring）**，下方刻意留 placeholder。rule 定義 owner = 本 plan。acceptance-gate
> 形狀候選 `pilot_complete + criteria_pass >= 6`；證據可跨 repo（下游 <AI_SKILL_DOGFOOD_EVIDENCE> commit）。notify
> 屬 acceptance-gate（gate projection），不在 evidence_rule。設計來源見
> [`evidence-candidate-system`](../active/2026-06-16-1131-evidence-candidate-system.md)。

```yaml
evidence_rule:
  collect: true
  match:
    artifact_types: [commit, screen-mapping-doc, test, bdd, adr]   # commit 含下游 repo（跨 repo 證據）
    criteria:
      - id: trust_transition_case
        description: 新的 invalidate↔recover 案例可用四欄 Ownership Map（State|Owner|Invalidation|Recovery）不改字表達
      - id: field_survival
        description: 四欄在 rename pressure（換域、不換列名）下仍成立 —— primitive 有生命力（criterion 5）
      - id: predictive_prevention
        description: ≥1 previously-unknown prevention，非已上線 bug 的事後解釋（criterion 4）
      - id: downstream_pilot_evidence
        description: 下游 consuming 專案（如 <AI_SKILL_DOGFOOD_EVIDENCE>）commit/diff 驗證 trust transition —— 跨 repo 證據
  exclusions:
    - post-hoc 解釋已 ship 的 bug（非 predictive interception）
    - 每案都要改栏位的 abstraction（abstraction noise，非 primitive）
```

## Roadmap

本 plan 已不是「做 workflow」，而是在管理**何时一个 primitive 有资格升级成 workflow**。这两个生命周期现在可以清楚分离：

```text
Lifecycle 1 — Primitive Maturation (completed)
  Primitive → Consumer validation → Promotion discipline
  ──────────────────────────────────────────────────────
  A0  — Template + Recovery Boundary（四栏定稿 + downstream sync）
  A1  — Coupon trust transition table
  B   — Optimistic rollback（invalidate ↔ recover 闭环）
  B.5 — Rename pressure test（websocket 或第三域；不必新 incident）

Lifecycle 2 — Graduation (discipline complete; surface promotion deferred)
  Predictive evidence accumulation → O3 decision（Phase D）
  ──────────────────────────────────────────────────────
  C   — Scenario spike（field survival ✅；predictive ✅ C.5 2026-07-13）
  D   — ADR criteria → **final O3**（runtime surface / slice **not** registered）
  E   — Project overlay advisory（out of scope for this plan）
  F   — Mechanical promotion（deferred）
```

### A0 — Template + Recovery Boundary

- [x] 四栏定义写入本 plan
- [x] Downstream project overlay / screen-mapping sync Recovery Boundary — <AI_SKILL_DOGFOOD_EVIDENCE> `bcce737`
- [x] Side Effect Chain：`invalidation_events` + **`recovery_evidence`** per step — downstream workflow yaml + screen mappings

### A1 — Coupon

- [x] 四栏填 coupon — <AI_SKILL_DOGFOOD_EVIDENCE> `episode-coupon-redeem-journey.md` (`6665b77`); aligns Appendix A
- [x] Counterfactual documented — coupon unmount hazard + recovery (post-ship fix recorded; Criterion 4 later satisfied via C.5, not this case)

### B — Optimistic rollback / payment sync trust

- [x] Primary：invalidate + server sync recovery 落入四栏 — membership payment sync mapping (`6665b77`)
- [x] 验 **trust transition 闭环** — pending UI vs sync + `router.refresh` recovery boundary

### B.5 — Rename pressure test

- [x] **不用同名案例** — websocket subscription sketch (`6665b77`)
- [x] 四栏 **不改字** 套用 — BDD `template field survival` test passes
- [x] **B.5 pass ⇒ tentative O3**（非最终 Phase D graduation）

### C — Scenario spike

- [x] ≥1 previously unknown prevention — <AI_SKILL_DOGFOOD_EVIDENCE> C.5 PaymentLeaveConfirmDialog (2026-07-13)
- [x] ≥1 scenario asserts template field survival — <AI_SKILL_DOGFOOD_EVIDENCE> `state-trust-transition-pilot.test.mjs` (`6665b77`)
- [~] Draft ids promoted to Ai-skill `validation/scenarios/` — **not done**；本 plan 接受 deferred（downstream BDD equivalent green；無 Ai-skill scenario 檔）

### D — Graduation

- [x] O2 / O3 **final** written decision — **final O3** (downstream pilot §D, 2026-07-13)
- [x] Criteria 4 + integration journey evidence met for promotion *discipline*
- [x] **Do not** register runtime surface / lifecycle phase — decision **held**（evidence accumulation ≠ framework expansion；repo 內無 ownership_map / experience-runtime lifecycle 升格）

---

## Watch-out: execution primitive drift

Many primitives are **not designed** — they are **pulled out by three or four consumers**.

**Do not optimize for becoming a primitive.**

- Phase D 前：不称 slice / lifecycle phase
- 观察最小闭环：**Ownership + Invalidation + Recovery** 能否跨 consumer 成立
- 若能：它不是 UI workflow — 单独 ADR 再议 naming / owner layer
- 若不能：stay O2 conditional gate

---

## Stakeholder review log

| Review | Key outcome |
|---|---|
| #1 | scenario→slice；evidence 不膨胀 |
| #2 | O2 vs O3；Invalidation Event；predictive ADR |
| #3 | Recovery Boundary；trust transition 验题；B.5 rename pressure；field survival ADR；anti optimize-for-primitive |
| #4 | Downstream pilot gate 回写；A0–B.5 pass；C partial；tentative O3 |
| #5 | Formalized tentative O3 as the current stage decision；separated framework maturity from workflow graduation；predictive evidence 须自然产生而非为过 gate 制造 |
| #6 | Downstream evidence sync；ADR #4 operationalized；Interaction Evidence Hierarchy watch |
| #7 | Downstream C.5 PaymentLeaveConfirmDialog + L1 @ test `/h5` `6a7cc1c` → Criterion 4 pass；**final O3**；runtime surface仍 deferred |
| #8 | Checkbox audit + **Plan Completion Closure**：Ai-skill scenarios `[~]` deferred；runtime surface decision held；`status: completed` → `plans/archived/` |
| #9 | 回寫 ECS 關係：本 plan 是 Evidence Candidate System **Phase 1 冷啟動消費者**（criteria + gate 形狀 + expire fixture），非 observation 期真實 accept 樣本 — 見 [`2026-06-16-1131-evidence-candidate-system`](../active/2026-06-16-1131-evidence-candidate-system.md) §與其他 plans 的關係 |

---

## Related — Evidence Candidate System（把 ECS 建起來）

本 plan **completed** 後不升 Ai-skill workflow；對
[`2026-06-16-1131-evidence-candidate-system`](../active/2026-06-16-1131-evidence-candidate-system.md)
的貢獻停在 **Phase 1 observation infrastructure**：

| 貢獻 | 內容 |
|---|---|
| Consumer attach | `## Evidence Rule` + pointer `interaction-hazard.pointer.yaml` |
| Criteria | `trust_transition_case` / `field_survival` / `predictive_prevention` / `downstream_pilot_evidence` |
| Gate 形狀 | `pilot_complete + criteria_pass>=6`（第三種 maturity 形狀） |
| Plumbing fixture | Phase 1B `C-0003` → **expire**（ratio 不失真） |

**Non-claim：** final O3 / C.5 證據**未**經 ECS candidate→accept 寫回；不計入 ECS `phase2_gate` 真實樣本。

---

## 完成条件

- [x] A0 四栏 synced（downstream `bcce737`）
- [x] A1 + B + B.5 填表（downstream `6665b77`）
- [x] C scenario complete — field survival **pass**; predictive prevention **pass** (C.5 2026-07-13)
- [x] O2/O3 **final** 书面决策 — **final O3**
- [x] Phase D promotion **discipline** — ADR criteria 1–6 pass
- [x] Runtime surface / workflow slice registration — **deferred by design**（explicit non-goal；決策已落地＝不註冊，非漏做）

---

## Appendix A — Coupon

| State | Owner | Invalidation Event | Recovery Boundary |
|---|---|---|---|
| `previewLimitReached` | ImmersivePlayerFrame | early redeem success clears mask | refresh + entitlement readback |
| `dialogOpen` | Hoisted Frame (was CouponPanel) | panel unmount | explicit reopen after terminal redeem state |
| `pendingRedeem` | Redeem mutation | cancel / onOpenChange(false) | idle / new user action |
| `playbackAllowed` | Entitlement gate | refresh before grant readback | grant readback + poll window clear |

## Appendix B — Optimistic rollback

| State | Owner | Invalidation Event | Recovery Boundary |
|---|---|---|---|
| `optimisticBalance` | QueryCache | rollback | server sync / query refetch |
| `optimisticUISuccess` | View layer | invalidate (server truth differs) | reconciled view model |
| `pendingMutation` | Mutation hook | supersede | new mutation settled |

## Appendix C — B.5 websocket sketch（rename pressure）

| State | Owner | Invalidation Event | Recovery Boundary |
|---|---|---|---|
| `websocketReady` | Connection manager | reconnect / drop | handshake complete |
| `subscriptionActive` | Client session | connection invalidate | resubscribe ack |

**B.5 pass** = 上表无需改列名即可填写。
