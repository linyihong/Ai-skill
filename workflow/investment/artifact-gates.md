# Investment Artifact Gates

定義**必須產出什麼、何時不得宣稱完成**。

## 必要產出（依 task type）

| 產出 | framing | theme | name | position | allocation | event | sweep |
| --- | --- | --- | --- | --- | --- | --- | --- |
| Investment Task Frame | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| Risk tier 宣告 | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| Disclaimer | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| Open Questions | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| Decision Register（Pass1／2） | 簡 | ✅ | ✅ | ✅ | ✅ | 條件 | 條件 |
| Evidence-ledger | — | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| Strategy／資產／費用摘要 | — | — | — | 條件 | ✅ | — | — |
| Interest Analysis（含費用） | — | — | — | 條件 | ✅ | — | — |
| ≥2 方案比較 | — | — | — | — | ✅ | — | — |
| Uncertainty framing | — | Yellow+ | Yellow+ | Yellow+ | ✅ | Yellow+ | Yellow+ |
| Human selection 明示 | — | 有建議時 | 有建議時 | ✅ | ✅ | — | — |
| Verifier 摘要 | — | DVA 時 | DVA 時 | DVA 時 | DVA 時 | — | DVA 時 |
| Escalation Card | Red | Red | Red | Red | Red | Red | Red |

## Blocking gates

| Gate | 條件 | 未通過 |
| --- | --- | --- |
| `gate.investment.intake_complete` | S0 有值或 UNKNOWN；配置類 S1／S2 齊或明確 blocking | 不得輸出配置執行指令 |
| `gate.investment.risk_tier_declared` | tier＋理由在回覆早段 | 不得實質建議 |
| `gate.investment.decision_register_present` | Pass 1 決策點＋待查證（簡化情形除外） | 不得進 Pass 2 定稿 |
| `gate.investment.strategy_reasoned` | 四欄＋Interest；Confidence 正確 | 不得把建議寫成已決策 |
| `gate.investment.strategy_asset_fee_present` | 配置類有摘要或 `provisional`＋待查 | 不得假裝完整 Interest |
| `gate.investment.evidence_ledger_present` | Yellow+ 主張可回溯 authority／locator／as-of | 該主張降級或刪建議 |
| `gate.investment.uncertainty_framed` | 無未校準精確勝率；有 framing labels | 改寫或 Verifier fix |
| `gate.investment.no_uncalibrated_probability` | 禁裸「60%／高機率最優」 | blocking |
| `gate.investment.recommendation_ne_decision` | 明示 human selection／不代執行 | blocking |
| `gate.investment.dva_complete` | 強制 DVA 時有 O／E／V／仲裁紀錄；或 `skipped(user)` | 不得宣稱 allocation 定稿 |
| `gate.investment.artifacts_complete` | 上表必要產出齊 | 不得宣稱完成 |
| `gate.investment.disclaimer_present` | disclaimer 要點齊 | 不得宣稱完成 |

## DVA 適用表

| 情境 | DVA |
| --- | --- |
| Green 名詞／地圖 | 否 |
| Yellow theme／name／event／sweep | 建議；「可驗證報告」時強制 |
| `allocation-advice` | **強制** |
| 含資產之 `position-review` | **強制** |
| 使用者明示跳過 | 允許；Frame 記 `DVA: skipped(user)` |

契約：[`../software-delivery/delegated-execution.md`](../software-delivery/delegated-execution.md)＋[`../../plans/README.md`](../../plans/README.md) §Delegation loop SOP。  
Acceptance 必含（investment）：`verifier_only` 挑戰未校準 %；挑戰 Recommendation→Decision；挑戰 D 獨撐。

## Evidence-ledger 最小區塊

| 區塊 | 內容 |
| --- | --- |
| 策略／資產／費用摘要 | 去敏；provisional 標註 |
| 證據表 | Claim → A–D → locator → as-of → stance |
| 推算鏈 | Premise → inference → uncertainty → recommendation |
| 方案比較 | ≥2；含費用淨效應（配置類） |
| Verifier 摘要 | findings＋仲裁（若有） |

## Disclaimer（要點）

1. 本產出**不是**受託投資建議執行、不構成下單授權。  
2. 市場與公司資訊會變；以標註 as-of 為限。  
3. 建議強度受證據限制；**決策與執行由使用者負責**。  
4. Yellow：具體待覆核項；Red：Escalation Card。

## Definition of Done

- [ ] Task type／tier／DVA 狀態已明說  
- [ ] Blocking gates 通過或標阻塞  
- [ ] 必要產出齊  
- [ ] 無保證報酬／代客下單步驟  
- [ ] 持倉未寫入 canonical  
