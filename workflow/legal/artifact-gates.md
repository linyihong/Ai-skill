# Legal Artifact Gates（產出格式與完成定義）

本檔定義**每個法律任務必須產出什麼、什麼情況下不得宣稱完成**。
Gate 為 blocking 者，未通過即不得輸出實質產出或宣稱完成。

## 必要產出（依 task type）

| 產出 | draft | review | explain | compare | research | due-diligence | strategy | negotiation | lifecycle |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Legal Task Frame | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| Risk Tier 判定 | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| Open Questions（blocking / non-blocking 分開） | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| Disclaimer | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| Decision Register + 四欄 Decision Reasoning | ✅ | ✅ | — | — | ✅ | — | ✅ | ✅ | ✅ |
| Applicable Law Table | ✅ | ✅ | ✅ | 條件 | ✅ | 條件 | ✅ | 條件 | ✅ |
| Reference Source Table | 條件 | 條件 | — | 條件 | ✅ | — | 條件 | — | ✅ |
| Counterparty Diligence Card | 條件 | 條件 | — | — | — | ✅ | 條件 | 複用 | ✅ |
| Jurisdiction Block | 跨境✅ | 跨境✅ | — | — | 跨境✅ | — | 跨境✅ | 跨境✅ | ✅ |
| Draft Handover | ✅ | — | — | — | — | — | — | — | ✅ |
| Review Memo（含風險分級） | — | ✅ | — | ✅ | — | — | — | — | ✅ |
| Issue Ledger + Concession Matrix | — | — | — | — | — | — | — | ✅ | 條件 |
| Escalation Card | Red✅ | Red✅ | Red✅ | Red✅ | Red✅ | Red✅ | Red✅ | Red✅ | Red✅ |

「條件」＝依 [`execution-flow.md`](execution-flow.md) §Dispatch Matrix 判定。

## Blocking gates

| Gate | 條件 | 未通過時 |
| --- | --- | --- |
| `gate.legal.intake_complete` | [`intake.md`](intake.md) §S0 全部有值或明確 `UNKNOWN`／`DECLINED`；jurisdiction 已確認或已列為 blocking 問題 | 只能輸出 Frame + blocking 問題清單 + 法域中立結構檢查；**不得**輸出條款全文或法律結論 |
| `gate.legal.risk_tier_declared` | Risk tier 已依 [`risk-classification.md`](risk-classification.md) 判定，並在回覆早段明說 tier 與理由 | 不得輸出任何實質產出 |
| `gate.legal.decision_register_present` | Strategy Pass 1 已產出 Decision Register（決策點 + 選項 + 待查證前提） | 不得進入 Draft / Negotiate |
| `gate.legal.strategy_reasoned` | 每個決策點四欄齊備（Recommendation / Reason / Alternative / Trade-offs）+ Confidence 正確；Interest Analysis 已分析對手方利益 | 不得把決策落成條款文字 |
| `gate.legal.law_citation_versioned` | 每筆法規／範本引用有「名稱 + 版本或最新修正日 + 查核日」；無來源者標 `unverified` | 該主張不得以確定語氣呈現；引用它的結論降級 |
| `gate.legal.source_version_pinned` | 使用官方範本時版本號已鎖定 | 不得寫「依官方範本」 |
| `gate.legal.counterparty_identified` | 對手方法人全稱 + 登記編號 + 簽署權限已核實，或明確標 `unverified` 並說明影響 | 不得產出以對手方可信度為前提的條款建議 |
| `gate.legal.artifacts_complete` | 上表對應 task type 的必要產出齊備 | 不得宣稱完成 |
| `gate.legal.escalation_declared` | Risk tier 與產出深度一致；Red tier 已輸出 Escalation Card 且無實質建議 | 不得輸出 |
| `gate.legal.no_outbound_send` | 任何對外文件（存證信函、通知、正式回覆、寄給對方的稿件）在發送前已取得使用者明確確認 | 不得發送 |

## Advisory checks（不 blocking，但要回報）

- 條款清單完整性（對照 [`review/README.md`](review/README.md) Stage 4 缺漏表）
- 定義／交叉引用／附件／雙語一致性
- 未決爭點數是否逐輪下降（negotiation）
- Archive 關鍵日期是否登記（lifecycle）

## Disclaimer（每份產出必附）

內容要求（措辭可調整，但四點不可少）：

1. 本產出**不是法律意見**，不構成律師與委任人關係。
2. 法規與官方範本會改版；引用內容以標註的**版本與查核日**為限。
3. 具體個案的結果取決於完整事實與最新法規，**簽署前建議由執業律師覆核**。
4. Yellow tier 時附**具體**需覆核項清單；Red tier 時附 Escalation Card。

**禁止**：把 disclaimer 當作輸出 Red tier 實質建議的通行證（見
[`risk-classification.md`](risk-classification.md)）。

## Escalation Card（Red tier 專用）

格式與禁止事項見 [`risk-classification.md`](risk-classification.md) §Red。
`gate.legal.escalation_declared` 驗證：無存證信函內容、無條號依據、無勝敗或金額評估、
且 Card 含文件準備清單與專業人士類型。

## Confidence 標籤

| 標籤 | 用在哪 | 條件 |
| --- | --- | --- |
| `confirmed` | 法規引用、策略建議、背調結論 | A／B 層來源 + 版本 + 查核日 |
| `probable` | 背調結論 | 多個 C 層來源一致但未回溯 A 層 |
| `provisional` | 策略建議 | Strategy Pass 1，前提未查證 |
| `unverified` | 任何主張 | 無 A／B 層來源支撐 |

不得對 `unverified` 或 `provisional` 的內容使用確定語氣。

## 完成定義（Definition of Done）

宣稱完成前逐項確認：

- [ ] Task type / Jurisdiction / Risk tier 三者都已明說
- [ ] 所有 blocking gate 通過（或明確說明被什麼阻塞）
- [ ] 對應 task type 的必要產出齊備
- [ ] 每個法規／範本引用有版本 + 查核日；`unverified` 已標記
- [ ] Decision Register 每個決策點四欄齊備，Confidence 正確
- [ ] 背調 risk flag 都對應到具體條款影響
- [ ] Open Questions 分 blocking / non-blocking，且 blocking 項已明確請使用者處理
- [ ] Disclaimer 四點齊備；Yellow 有具體覆核清單；Red 有 Escalation Card
- [ ] 未在未取得確認的情況下對外發送任何文件

任一項未過即不得宣稱完成——回報卡在哪一項，而不是宣稱「已完成」。
