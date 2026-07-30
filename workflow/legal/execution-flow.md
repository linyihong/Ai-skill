# Legal Workflow Execution Flow

本檔是 `workflow/legal/` 的 **canonical lifecycle 與分派表**。只記 stage 順序與 dispatch，
各 slice 正文在對應子檔（見 [`README.md`](README.md) §thin index），避免 dual source-of-truth。

> 本 domain 是 **intake-dispatched workflow**：intake 的答案決定要跑哪些 stage，
> 不是每個任務都跑完 10 個 stage。

> **核心設計**：本流程不是「問問題 → 寫合約」，而是
> **「問問題 → 推理最佳策略 → 使用者決策 → 才決定怎麼寫」**。
> Stage 3 [`strategy/`](strategy/README.md) 是這條流程與模板式產出的分水嶺。

## Lifecycle（stage 順序）

```text
0.  Frame            → 確認這是法律任務，宣告非法律意見邊界
1.  Task Intake      → legal task type + jurisdiction(P0) + 我方角色 + S0 必問
2.  Risk Tier        → Green / Yellow / Red；Red 直接跳 Stage 9
3a. Strategy Pass 1  → 決策點 + 選項 + 利益分析 + 待查證前提 → Decision Register
4.  Dispatch         → 依 Frame + Decision Register 決定下游 stage 組合
5.  Due Diligence    → Identity → Corporate Status → Risk Signals（有對手方時）
6.  Applicable Law   → Jurisdiction → Laws → Recent Updates → Gov Sources → Risk Summary
7.  Reference        → 選定權威來源（範本 + 法條 + Q&A + 函釋），鎖定版本
3b. Strategy Pass 2  → 以已核實前提收斂 → Recommendation + Trade-offs → 使用者確認
8.  Produce          → draft / review / compare / explain / negotiation
9.  Validate         → artifact gates 逐項核對
10. Close            → 產出 + 未決清單 + escalation card + disclaimer
```

Stage 0–4、3b 與 9–10 **每個任務都跑**。Stage 5–8 由 Dispatch Matrix 決定。

**為何 Strategy 分兩趟**：策略推理依賴法律前提（強制法、判決可執行性、範本可否增修）。
Pass 1 在查證前只能給 `provisional` 建議並標記待查證項；Pass 2 在 Research／Reference
完成後才收斂為 `confirmed` 並取得使用者決策。理由見
[`strategy/README.md`](strategy/README.md) §為什麼要兩趟。

## Stage 明細

| Stage | 動作 | Canonical source | Gate |
| --- | --- | --- | --- |
| 0 Frame | 確認任務屬法律 domain；一句宣告產出非法律意見；明顯屬 Red tier 領域先標記 | 本檔 | — |
| 1 Task Intake | 分層問卷 S0→S1→S2；未答完不得進入 Stage 8 的實質產出 | [`intake.md`](intake.md) | `gate.legal.intake_complete` |
| 2 Risk Tier | 依 tier table 判定並在回覆早段明說 tier 與理由 | [`risk-classification.md`](risk-classification.md) | `gate.legal.risk_tier_declared` |
| 3a Strategy Pass 1 | 列出決策點、選項、利益分析、trade-off 與待查證前提 | [`strategy/README.md`](strategy/README.md) + [`strategy/decision-playbooks.md`](strategy/decision-playbooks.md) | `gate.legal.decision_register_present` |
| 4 Dispatch | 依 task type × jurisdiction × 有無對手方 × Decision Register 選 stage 組合 | 本檔 §Dispatch Matrix | — |
| 5 Due Diligence | 三層核實 + risk flag → 對應條款調整建議 | [`due-diligence/README.md`](due-diligence/README.md) | `gate.legal.counterparty_identified` |
| 6 Applicable Law | 五步研究；每條引用帶版本 + 查核日 | [`research/README.md`](research/README.md) | `gate.legal.law_citation_versioned` |
| 7 Reference | 選來源並鎖定版本；政府採購須涵蓋四類來源 | [`reference-sources.md`](reference-sources.md) | `gate.legal.source_version_pinned` |
| 3b Strategy Pass 2 | 收斂建議為 `confirmed`；四欄 Decision Reasoning；請使用者決策 | [`strategy/README.md`](strategy/README.md) | `gate.legal.strategy_reasoned` |
| 8 Produce | draft / review 5 階段 / compare / explain / negotiation | 對應 sub-flow | 見各 sub-flow |
| 9 Validate | artifact gates 逐項核對，未過不得宣稱完成 | [`artifact-gates.md`](artifact-gates.md) | `gate.legal.artifacts_complete` |
| 10 Close | 產出 + Open Questions + Escalation Card + disclaimer | [`artifact-gates.md`](artifact-gates.md) | `gate.legal.escalation_declared` |

## Dispatch Matrix

依 Stage 1 的答案決定 Stage 5–8 的最小必跑集合（可加不可漏）：

| Task type | Stage 5 DD | Stage 6 Law | Stage 7 Ref | Stage 8 Produce |
| --- | --- | --- | --- | --- |
| `draft` | 有對手方且金額／期間非瑣碎時必跑 | 必跑 | 有官方範本可用時必跑 | [`draft/`](draft/README.md) |
| `review` | 對手方不熟悉時必跑 | 必跑 | 對方稿源自官方範本時必跑 | [`review/`](review/README.md) 5 階段 |
| `explain` | 不跑 | 必跑（解讀需綁法域） | 選用 | [`review/`](review/README.md) §Clause Review |
| `compare` | 不跑 | 版本差異涉法規變動時必跑 | 官方範本改版比較時必跑 | [`review/`](review/README.md) §Structural + Clause |
| `research` | 不跑 | **主流程** | 必跑 | [`research/`](research/README.md) |
| `due-diligence` | **主流程** | 制裁／黑名單涉法規時必跑 | 選用 | [`due-diligence/`](due-diligence/README.md) |
| `negotiation-support` | 已有 DD 則複用 | 爭點涉法規時必跑 | 選用 | [`negotiation/`](negotiation/README.md) |
| `strategy` | 對手方信用影響方案時必跑 | 必跑 | 選用 | [`strategy/`](strategy/README.md)（策略本身即產出） |
| `lifecycle` | 必跑 | 必跑 | 必跑 | [`lifecycle/`](lifecycle/README.md) 統籌全段 |

**Strategy 例外**：`explain` 與 `compare` 的 Stage 3a／3b 可簡化為「指出對方安排偏離常見
最佳實務之處」，不必產出完整 Decision Register；其餘 task type 一律必跑。

**Jurisdiction 加成**：jurisdiction 為跨境（≥2 法域）或非 TW 時，Stage 3a 之前必須先讀
[`jurisdiction.md`](jurisdiction.md) 的五變數模型，否則策略推理缺少 Enforcement 這個
反向決定因素。

**Red tier 短路**：Stage 2 判定 Red 時，**跳過 Stage 3a–8**（含策略建議），直接進 Stage 10
產出 Escalation Card。不得因使用者要求完整度而繞過。

## 常見失效模式

| 失效 | 症狀 | 防呆 |
| --- | --- | --- |
| Implementation-first | 拿到「幫我寫一份服務合約」直接輸出條款全文 | Stage 1 `gate.legal.intake_complete` blocking |
| Strategy-less execution | 問完問題直接寫條款，沒有方案比較與建議 | Stage 3a/3b `gate.legal.strategy_reasoned` blocking |
| Ask-only | 把「法院選哪裡」原封丟回使用者，不做分析 | Decision Reasoning 四欄強制 |
| Silent jurisdiction default | 未問即假設台灣法 | Stage 1 jurisdiction 為 P0 必問；UNKNOWN 不可往下 |
| Unsourced statute | 憑記憶給條號或數字上限 | Stage 6 `gate.legal.law_citation_versioned` |
| Stale template | 引用「官方範本」但未鎖版本 | Stage 7 `gate.legal.source_version_pinned` |
| Provisional as confirmed | Pass 1 的建議標成已確認 | Confidence 欄 + Pass 2 來源要求 |
| Diligence as appendix | 背調結果沒有對應到條款調整 | Stage 5 要求每個 risk flag 附條款影響 |
| Disclaimer-only escalation | Red tier 照常出建議，結尾加免責 | Stage 2 短路 + `gate.legal.escalation_declared` |

## 驗證

收尾前確認：task type / jurisdiction / risk tier 都已明說；Decision Register 每個決策點
四欄齊備且 Confidence 正確；Dispatch Matrix 要求的 stage 都跑過或標明不適用；
[`artifact-gates.md`](artifact-gates.md) 的 blocking gates 全部通過。
