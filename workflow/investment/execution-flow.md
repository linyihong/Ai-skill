# Investment Workflow Execution Flow

本檔是 `workflow/investment/` 的 **canonical lifecycle 與分派表**。各 slice 正文在子檔；
取證方法在 `analysis/investment/`，避免 dual source-of-truth。

> **intake-dispatched**：不是每個任務都跑完全部 stage。  
> **核心**：問清楚 → Decision Support Pass 1 → Research → Pass 2 →（必要時 DVA）→ 產出 → gates。

## Lifecycle

```text
0.  Frame            → 確認投資／研究決策輔助；宣告非投顧執行／不下單
1.  Task Intake      → task type + 市場範圍 +（配置類）策略／資產／費用
2.  Risk Tier        → Green / Yellow / Red
3a. Strategy Pass 1  → Decision Register（provisional）+ 待查證前提
4.  Dispatch         → 選 Research／Produce sub-flow；決定是否 DVA
5.  Research         → analysis/investment 方法（主題／新聞／筆記對照／公開揭露）
3b. Strategy Pass 2  → 收斂 recommendation + Interest Analysis + human selection 點
6.  DVA（條件）      → O brief → E 草稿 → V findings → 仲裁（消費 delegated-execution 契約）
7.  Produce          → 對應 task type artifact
8.  Validate         → artifact-gates
9.  Close            → 產出 + open questions + disclaimer（+ Verifier 摘要若有）
```

Stage 0–4、3b、8–9 **多數任務都跑**。Stage 5 深度依 task type。Stage 6 依 DVA 表。

## Stage 明細

| Stage | 動作 | Canonical source | Gate |
| --- | --- | --- | --- |
| 0 Frame | 確認 decision object；排除「投資協議」誤吸；一句非執行／非保證 | 本檔 + README Q8 | — |
| 1 Intake | 分層問卷；配置類缺策略／資產 → blocking | [`intake.md`](intake.md) | `gate.investment.intake_complete` |
| 2 Risk Tier | 明說 tier 與理由 | [`risk-classification.md`](risk-classification.md) | `gate.investment.risk_tier_declared` |
| 3a Pass 1 | 決策點＋選項＋待查證；Confidence=`provisional` | [`strategy/`](strategy/README.md) | `gate.investment.decision_register_present` |
| 4 Dispatch | 依 task type × tier 選 sub-flow 與 DVA | 本檔 §Dispatch | — |
| 5 Research | 取證；填 evidence 表 | [`analysis/investment/`](../../analysis/investment/README.md) | `gate.investment.evidence_ledger_present`（Yellow+） |
| 3b Pass 2 | 四欄 Decision Reasoning；Interest（含費用）；human selection | [`strategy/`](strategy/README.md) | `gate.investment.strategy_reasoned` |
| 6 DVA | 三角色；Verifier 只產 findings | [`allocation-advice/`](allocation-advice/README.md) §DVA + [`../software-delivery/delegated-execution.md`](../software-delivery/delegated-execution.md) | `gate.investment.dva_complete`（強制情境） |
| 7 Produce | 各 sub-flow artifact | 對應 README | — |
| 8 Validate | gates 逐項 | [`artifact-gates.md`](artifact-gates.md) | `gate.investment.artifacts_complete` |
| 9 Close | disclaimer + open questions | [`artifact-gates.md`](artifact-gates.md) | `gate.investment.disclaimer_present` |

## Dispatch Matrix

| Task type | Research 重點 | Produce | DVA |
| --- | --- | --- | --- |
| `need-framing` | 輕 | Framing note＋建議路徑 | 否 |
| `theme-research` | 主題／供應鏈 | Thesis＋地圖 | 建議（可驗證報告時強制） |
| `name-diligence` | 單一標的揭露＋新聞＋筆記 | Diligence card | 建議 |
| `position-review` | Thesis 對照新證據 | Still-valid card | **含資產時強制** |
| `allocation-advice` | 方案比較所需市場證據 | Allocation brief | **強制**（除非明示跳過並記錄） |
| `event-check` | 事件來源 | Event impact card | 否（升 Yellow 建議時可加） |
| `periodic-sweep` | Delta only | Sweep brief | 建議 |

**Red 短路**：Stage 2＝Red → 跳過實質交易指令與「保證報酬」類建議；可產出 Escalation／待查清單與允許的研究整理（見 risk-classification）。

**Pass 簡化**：純 Green 名詞／地圖整理可合併 3a／3b 為短 Decision Register；一有方向性建議即恢復完整 two-pass。

## 常見失效模式

| 失效 | 防呆 |
| --- | --- |
| Keyword→investment | Stage 0 Q8 decision object 檢查 |
| 無策略給配置 | `gate.investment.strategy_asset_fee_present` |
| 假精確 % | Verifier `verifier_only`＋artifact uncertainty 規則 |
| D 獨撐建議 | evidence authority 檢查 |
| Recommendation＝Decision | 強制 human selection 欄 |
| 形式 DVA（Verifier 複讀） | findings 須挑戰 acceptance；見 dogfood 05／4a |
| 同 session 當 fresh Task | 記錄 transport；正式強制時應分離 context |

## 驗證

收尾前：task type／tier／DVA 適用已明說；blocking gates 通過或標阻塞；evidence 可覆核；未代執行下單。
