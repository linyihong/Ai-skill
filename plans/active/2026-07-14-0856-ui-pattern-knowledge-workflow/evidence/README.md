# Plan Evidence Index — 2026-07-14-0856-ui-pattern-knowledge-workflow

本目錄存放本 plan 的 **research cycle / Phase 量測 / Readiness gate** 全文。[`_plan.md`](../_plan.md) 只留決策摘要與檔案路徑連結。

> **Canonical 規則**：[`governance/lifecycle/plan-evidence.md`](../../../../governance/lifecycle/plan-evidence.md)（commit-msg `validatePlanEvidenceConvention` 機械強制）

## 引用規則（避免行號漂移）

| 做法 | 說明 |
|---|---|
| **用檔案路徑** | `evidence/<file>.md` 或相對連結；**不要**寫 kit／README 絕對行號 |
| **用標題錨點** | 定位寫 `evidence/foo.md` 內 `##` 標題或表格欄，不用絕對行號 |
| **專案細節** | dogfood / inner commit → consumer `<PROJECT_ROOT>`；本目錄只留 generalized metrics（[`enforcement/sanitization.md`](../../../../enforcement/sanitization.md)） |
| **新 run** | 新增 `evidence/<run-id>-<slug>.md` + 更新本表同一 commit |

## Run 索引

| Run ID | 檔案 | 狀態 | 摘要 |
|---|---|---|---|
| **RC1** | [`research-cycle-1.md`](research-cycle-1.md) | CLOSED | Research Cycle 1 回顧：代表性 / 可推斷 / 可組合假設驗證 |
| **P2-sum** | [`phase2-summary.md`](phase2-summary.md) | 完成 | Phase 2：5 entries + 10 selection scenarios；index / schema 凍結 |
| **P2a** | [`2a-family-inferability-run.md`](2a-family-inferability-run.md) | 完成 | Family inferability rule-trace + blind LLM |
| **sel** | [`selection-scenarios.yaml`](selection-scenarios.yaml) | 完成 | 10 selection scenarios 資料 |
| **P3-start** | [`phase3-start.md`](phase3-start.md) | 完成 | Phase 3 開場與 Composition Closure 目標 |
| **3-metrics** | [`3-metrics.md`](3-metrics.md) | 完成 | Composition metrics（Entry Mods primary） |
| **3H4** | [`3h4-independence-stress.md`](3h4-independence-stress.md) | 完成 | H4 Independence：concurrent overlay 規則；Entry Mods=0 |
| **3H5** | [`3h5-completeness-disposition.md`](3h5-completeness-disposition.md) | 完成 | H5 Completeness：deferred disposition 枚舉 |
| **3H6** | [`3h6-traceability.md`](3h6-traceability.md) | 完成 | H6 Traceability：complete \| waived 終止 |
| **P4-ready** | [`phase4-readiness-gate.md`](phase4-readiness-gate.md) | ✅ Closed | Readiness R1→R2→R3 PASS → RC2 started & closed |
| **R1-dogfood** | [`r1-consumer-dogfood-2026-07-15.md`](r1-consumer-dogfood-2026-07-15.md) | 完成 | `<PROJECT_ROOT>` dogfood：C1 preview gate + C2 payment leave；R1∧R2 PASS → 待 R3 |
| **RC2-P1** | [`rc2-p1-interaction-representability-start.md`](rc2-p1-interaction-representability-start.md) | ✅ Closed | H1–H3 kickoff；`preview_gate_transition` scope lock |
| **RC2-P1-run** | [`rc2-p1-preview-gate-representability-run.md`](rc2-p1-preview-gate-representability-run.md) | ✅ Closed | `preview_gate_transition` H1 PASS；Frozen Layer Mods=0 |
| **RC2-P1-closure** | [`rc2-p1-interaction-representability-closure.md`](rc2-p1-interaction-representability-closure.md) | ✅ Closed | P1 🟢 Stable |
| **RC2-P2** | [`rc2-p2-interaction-inferability-start.md`](rc2-p2-interaction-inferability-start.md) | ✅ Closed | IH1–IH3 PASS；blind cumulative 8/8 |
| **RC2-P2-closure** | [`rc2-p2-interaction-inferability-closure.md`](rc2-p2-interaction-inferability-closure.md) | ✅ Closed | P2 symmetric with RC1-P2 |
| **RC2-P3** | [`rc2-p3-interaction-composition-start.md`](rc2-p3-interaction-composition-start.md) | ✅ Closed | CH1–CH3 PASS |
| **RC2-P3-closure** | [`rc2-p3-interaction-composition-closure.md`](rc2-p3-interaction-composition-closure.md) | ✅ Closed | Interaction Knowledge 🟢 Stable |
| **RC2-P3-CH3** | [`rc2-p3-ch3-traceability.md`](rc2-p3-ch3-traceability.md) | ✅ Complete | CT1–CT3 PASS |
| **RC2-P3-metrics** | [`rc2-p3-composition-metrics.md`](rc2-p3-composition-metrics.md) | ✅ Closed | P3 exit met |
| **RC2-P3-intake** | Consumer `<PROJECT_ROOT>/.ai-skill/project/evidence/rc2-p3-interaction-composition-intake.md` | ✅ Writeback | C1/C2 from I-04 · P3 closure writeback done |
| **RC2-P3-CH1** | [`rc2-p3-ch1-independence-stress.md`](rc2-p3-ch1-independence-stress.md) | ✅ Complete | CH1 PASS · 2 constraints · Entry Mods=0 |
| **RC2-P3-CH2** | [`rc2-p3-ch2-completeness-disposition.md`](rc2-p3-ch2-completeness-disposition.md) | ✅ Complete | CH2 PASS · 8/8 disposition |
| **RC2-P2-intake** | Consumer `<PROJECT_ROOT>/.ai-skill/project/evidence/rc2-p2-interaction-incident-intake.md` | ✅ Writeback | 10 incidents；I-05 觸發第二 entry |
| **RC2** | [`research-cycle-2.md`](research-cycle-2.md) | ✅ CLOSED | RC2 retrospective：Representability → Inferability → Composability |
| **RC2-vocab** | [`rc2-vocabulary-exit-review.md`](rc2-vocabulary-exit-review.md) | ✅ Closed | Post-P3 vocabulary exit：四欄維持；無 schema 擴充 |
| **maint-gov** | [`maintenance-governance.md`](../../../../workflow/software-delivery/maintenance-governance.md) | ▶ Active | Post-research：Stable Maintenance Dogfood；no RC3 |
| **RC-close** | [`stakeholder-research-line-closure-2026-07-15.md`](stakeholder-research-line-closure-2026-07-15.md) | ✅ Closed | Stakeholder：research line ended · maintenance watch |
| **RC2-P2-intake-wb** | [`rc2-p2-interaction-incident-intake-summary.md`](rc2-p2-interaction-incident-intake-summary.md) | ✅ Writeback | Generalized metrics + Table 1 |
| **RC2-P2-sel** | [`interaction-inferability-scenarios.yaml`](interaction-inferability-scenarios.yaml) | 完成 | 8 scenarios：I-01–I-05 + decoys I-06–I-08 |
| **RC2-P2-run** | [`rc2-p2-inferability-run.md`](rc2-p2-inferability-run.md) | ✅ Closed | rule-trace 8/8 · blind cumulative 8/8 · Exit met |
