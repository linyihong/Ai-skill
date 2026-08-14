# Phase 4 — Abstraction review（H1–H8／不實作 generic）

| 欄位 | 值 |
| --- | --- |
| Run ID | 07-phase4-abstraction-review |
| Date | 2026-08-14 |
| Goal | 依 Phase 1 dogfood 決定 lean-promote／defer／reject；**本 phase 不實作 7 個 generic** |
| DS status | investment = Decision Support **converged case #2**（2／3）；全庫 stage 仍待第三案 |

---

## Instantiation 核對（四項）

| # | 項目 | 路徑 | 狀態 |
| --- | --- | --- | --- |
| 1 | Decision point inventory | `workflow/investment/strategy/README.md` | ✅ |
| 2 | Playbooks | `workflow/investment/strategy/decision-playbooks.md` | ✅（配置／費用門檻／主題深度／升 Red／DVA） |
| 3 | Verification source | `analysis/investment/` | ✅ |
| 4 | Risk／depth gate | `workflow/investment/risk-classification.md` | ✅ |

→ 掛 Instantiations 表為 case #2：**成立**。  
→ **不**因此註冊 `route.workflow.investment`（Phase 5）／**不**宣稱 DS 已是全庫必跑 stage。

---

## H1–H8 處置（本 phase 裁決）

| H | Dogfood | Phase 4 處置 | 說明 |
| --- | --- | --- | --- |
| H1 Evidence→Decision | PASS | **follow-up candidate**（不本 plan 實作） | 與 ERA「evidence constrains decision space」同構；升 generic 需另開 plan＋第三域 |
| H2 DS Lifecycle | PASS | **已由 DS Instantiation 覆蓋** | 不再抽第二套 lifecycle primitive |
| H3 Uncertainty Framing | PASS（修疤後） | **follow-up candidate** | Prefer framing labels；**reject** probability-first glossary 綁死 |
| H4 Decision Depth | PASS | **defer** | 需更多 depth 階證據；暫留 investment risk tiers |
| H5 Observation／Reassessment | PASS | **follow-up candidate** | Observation≠scheduler；投資＋travel 可對照後再抽 |
| H6 Source Authority | PASS strong | **follow-up candidate** | **reject** Expert Knowledge；authority 模型可跨域 |
| H7 Knowledge／User-State | PASS（虛構） | **investment-only／defer** | Q10 schema 未定；真實設定檔未 dogfood |
| H8 Semantic routing | PASS A/B；C→Q12 | **follow-up candidate** | 進 Phase 5 `workflow-routing` 歧義列；**不**本 phase 改 detector |

### 明確 reject（本輪）

| 想法 | 理由 |
| --- | --- |
| Expert Knowledge primitive | H6 |
| 立刻開 7 個 cross-cutting 實作 | Stakeholder Q11；falsification ladder |
| 假精確 % 作為預設建議語言 | H3／DVA 05 |
| Keyword precedence「投資」 | H8／Q8 |

---

## Playbooks（Phase 4 確認）

已涵蓋於 `strategy/decision-playbooks.md`：

1. 配置方案選擇  
2. 再平衡／交易成本門檻（Q13）  
3. 主題深挖深度  
4. 何時升 Red  
5. 是否跑／跳過 DVA  

**不**另建平行 playbook 檔。

---

## Scenarios（本 phase 新增）

| ID | 檔 |
| --- | --- |
| investment-intake-gate-v1 | `validation/scenarios/investment/investment-intake-gate-v1.yaml` |
| investment-evidence-required-before-advice-v1 | `…/investment-evidence-required-before-advice-v1.yaml` |
| investment-probability-framing-v1 | `…/investment-probability-framing-v1.yaml` |
| investment-strategy-asset-required-for-allocation-v1 | `…/investment-strategy-asset-required-for-allocation-v1.yaml` |
| investment-dva-required-for-allocation-v1 | `…/investment-dva-required-for-allocation-v1.yaml` |
| investment-fee-interest-analysis-required-v1 | `…/investment-fee-interest-analysis-required-v1.yaml` |

---

## 完成句（Phase 4）

1. Investment **是** Decision Support case #2（converged instantiation；route 未註冊）。  
2. 值得跨域跟進的是 H1／H3／H5／H6／H8（另 plan）；H4／H7 defer；**不**在本 plan 抽 generic。
