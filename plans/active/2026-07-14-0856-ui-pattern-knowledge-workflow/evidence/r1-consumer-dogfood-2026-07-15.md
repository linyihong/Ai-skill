# R1 — Consumer dogfood evidence（Phase 4 Readiness）

**Plan**: [`../_plan.md`](../_plan.md)  
**Gate**: [`phase4-readiness-gate.md`](phase4-readiness-gate.md)  
**Date**: 2026-07-15  
**Role**: R1 搜證者（read-only in consumer；evidence writeback only）  
**Consumer**: `<PROJECT_ROOT>`（`<AI_SKILL_DOGFOOD_EVIDENCE>` pilot）

---

## Summary

| # | Candidate | R1 | R2 | R3 | Disposition |
| --- | --- | --- | --- | --- | --- |
| **C1** | Player preview gate projection break | **PASS** | **PASS** | — | 進 R3 假說起草；**不**開 Interaction schema |
| **C2** | Membership payment leave confirm（C.5） | **PASS** | **PASS** | — | 同上 |

**Rejected near-miss（wrong layer）**

| Candidate | R1 | 應留層 |
| --- | --- | --- |
| TAB-005 / G4 player-return | FAIL | Continuation（capture→restore；非 Pattern Flow） |
| `player-episode-sheet-hit-target` | FAIL | Composition（TabBar / empty overlay 空間互斥） |
| Scroll load-more bottom gate stall | FAIL | Pagination runtime（非 `episode_detail` Pattern Tree） |

---

## C1 — Player preview gate projection break

### Consumer anchors

| Kind | Path（under `<PROJECT_ROOT>`） |
| --- | --- |
| Screen mapping | `docs/frontend-contracts/screen-mapping/player-preview-gate.md` §Known projection break |
| Overlay rule | `.ai-skill/project/rules/player-client-patterns.md` |
| Experience runtime | `<AI_SKILL_REPO>/workflow/cross-cutting/experience-runtime/player.yaml` |
| Failure catalog | `<AI_SKILL_REPO>/workflow/software-delivery/validation/failure-evolution-catalog.md` → `player_preview_gate_projection_break` |
| Integration | `tests/integration/player-preview-gate.integration.mjs` |

### Pattern attestation ✅

| Check | Verdict | Authority |
| --- | --- | --- |
| Gate overlay 選型 | **modal_dialog** 族（blocking / short_form_gate） | `entries/modal_dialog.yaml#selection_rules.when` |
| 非 bottom_sheet | 訂閱阻斷非「底部多選」 | `entries/bottom_sheet.yaml#selection_rules.not_when` |
| 非 toast | 需決策 + 阻斷 | `entries/toast.yaml` near_neighbor 分界 |

### Composition attestation ✅

| Check | Verdict | Authority |
| --- | --- | --- |
| Pattern Tree | `episode_detail`: player (waived) + modal_dialog + scrim | `compositions/episode_detail.yaml` |
| Constraints | `overlay.dialog_requires_scrim` 滿足；無 sheet+dialog 雙開 | `composition_rules.yaml` |
| 空間語意 | mask + guide 與主 player stage 並存 — 樹合法 | integration envelope + player-spec |

### Flow attestation ❌

```text
Pattern layer     ✅  modal_dialog gate 選型正確
Composition layer ✅  episode_detail 樹 + scrim constraint 正確
Flow layer        ❌  preview → gated on preview_limit_reached 未在主舞台發火
```

| Dimension | Expected | Observed (incident) |
| --- | --- | --- |
| State machine | `preview` → `gated` on `preview_limit_reached` | poll / `timeupdate` 綁在 adjacent preload video |
| Temporal | bounded time mask + guide；preview 禁 auto-next | BDD source 綠；browser 無 mask |
| Evidence gap | `evidence:temporal_behavior` | 修復前僅 behavior-layer source assert → projection break |

**Remediation status**: integration 現綠（`player-preview-gate.integration.mjs`）；歷史 incident 仍證成 R1。

### R2 probe

> 加一條 `composition_rules` 能否修？

**不能 → R2 PASS**

| Hypothetical constraint | Why insufficient |
| --- | --- |
| `player.cannot_coexist_with: adjacent_preload` | 垂直 snap 需要 preload；非 composition 修補 |
| `overlay.dialog_requires_scrim`（已有） | 不解決 listener 綁錯 video 的時序 |
| `overlay.no_concurrent_temporary_overlays`（已有） | 與本 incident 無關 |

需要：狀態擁有者（`boundVideo` after mount）、轉移觸發（main-stage `preview_limit_reached`）、preview 態 `onEnded` guard。

---

## C2 — Membership payment leave confirm（C.5）

### Consumer anchors

| Kind | Path（under `<PROJECT_ROOT>`） |
| --- | --- |
| C.5 trial | `docs/plans/c5-trials/2026-07-13-payment-leave-confirm-dialog.yaml` |
| Screen mapping | `docs/frontend-contracts/screen-mapping/membership-payment-sync-trust-journey.md` |
| Interaction overlay | `.ai-skill/project/rules/interaction-hazard-review.md` |
| Capability integration | `tests/integration/payment-leave-confirm.browser.integration.mjs` |
| BDD | `tests/bdd/payment-leave-confirm.test.mjs` |

### Pattern attestation ✅

| Check | Verdict | Authority |
| --- | --- | --- |
| Leave confirm | **modal_dialog**（destructive_confirm / blocking_choice） | `entries/modal_dialog.yaml#selection_rules.when` |
| 非 bottom_sheet | Stay/Leave 單次決策 | `entries/bottom_sheet.yaml#selection_rules.not_when` |

### Composition attestation ✅

| Check | Verdict | Authority |
| --- | --- | --- |
| Dialog + scrim | Leave confirm 為 modal 族 | `overlay.dialog_requires_scrim` |
| 與 pending panel | Accordion 內 QR + hoisted confirm — 空間合法 | C.5 `design_delta` |

### Flow attestation ❌（counterfactual — pre-C.5）

```text
Pattern layer     ✅  modal_dialog 離開確認選型正確
Composition layer ✅  modal+scrim+accordion 空間組合正確
Flow layer        ❌  pending → accordion unmount → 未經 confirm 的狀態失效
```

| State | Owner | Invalidation (without flow fix) | Recovery boundary |
| --- | --- | --- | --- |
| `pendingPaymentVisible` | Subscribe form + pending panel | Accordion unmount → poll 停、QR 消失 | sync 或 explicit leave confirm |
| `leaveConfirmIntent` | Leave confirm dialog | backdrop dismiss / accordion race | `disablePointerDismissal`；Stay 保持 pending |
| `orderStatusTrusted` | Form view model | pending 中改 plan / 重送 | pending 中 lock controls |

**Ship status**: C.5 `incident.did_happen: false`（預防成功）；capability integration pass on consumer test deploy。Flow 失敗為 **validated counterfactual**。

### R2 probe

**不能 → R2 PASS** — `accordion.cannot_collapse_while: pending` 等價編碼狀態轉移，非 spatial edge；需 invalidation_event + recovery_boundary 鏈。

---

## R3 draft hint（本輪不定稿）

兩案共同形狀：

> Interaction Knowledge 描述 **在 Pattern 選型與空間組合已正確時，跨時間的合法狀態擁有、失效事件與恢復邊界** — 而非再選 overlay 或加 spatial constraint。

---

## Explicit non-actions（本 run）

- [x] 未改 `ui-pattern-knowledge/entries/*.yaml`
- [x] 未起草 Phase 4 / Interaction Knowledge schema
- [x] 未將 Continuation / Navigation 詞彙替代 Pattern Knowledge attestation
