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
| **P4-ready** | [`phase4-readiness-gate.md`](phase4-readiness-gate.md) | ✅ Closed | Readiness R1∧R2∧R3 PASS；RC2 啟動（非 Phase 4） |
| **R1-dogfood** | [`r1-consumer-dogfood-2026-07-15.md`](r1-consumer-dogfood-2026-07-15.md) | 完成 | C1+C2；R1∧R2∧R3 PASS → Interaction 🟡 Research Justified |
| **RC2-P1** | [`rc2-p1-interaction-representability-start.md`](rc2-p1-interaction-representability-start.md) | ✅ Closed | Kickoff：Vocabulary Freeze、metrics、scope lock |
| **RC2-P1-run** | [`rc2-p1-preview-gate-representability-run.md`](rc2-p1-preview-gate-representability-run.md) | ✅ Closed | `preview_gate_transition` H1 PASS；Frozen Layer Mods=0 |
| **RC2-P1-close** | [`rc2-p1-interaction-representability-closure.md`](rc2-p1-interaction-representability-closure.md) | ✅ Closed | Stakeholder closure：四假說 PASS；P1 🟢 Stable |
| **RC2-P2** | [`rc2-p2-interaction-inferability-start.md`](rc2-p2-interaction-inferability-start.md) | ▶ Active | IH1–IH3；layer-first；案例=`payment_leave_transition` |
