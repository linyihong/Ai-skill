# Legal Feedback History

對應 domain：[`workflow/legal/`](../../../workflow/legal/README.md)（法律工作 workflow —— 合約起草／
審閱／解釋／比較、適用法研究、對手方背景調查、法律策略分析、談判支援、合約生命週期）。

## 分類

| 分類 | 數量 | 說明 |
|------|------|------|
| [`common/`](common/) | 2 | 跨任務類型的通用 lesson（背調、策略推理、驗收與付款結構、風險分級、交付物權利鏈） |

## 邊界

- 本 domain 只放**可重用的法律工作方法**：背調欄位如何對應條款、策略推理的失效模式、
  常見結構缺口的辨識方式、風險分級的判斷依據。
- **不放**：具體合約的當事人、統一編號、地址、代表人、金額、條號爭點或任何可識別
  單一專案之資訊。個案 review memo 留在業務專案。
- **不放**法規適用結論。法規內容與版本由 [`workflow/legal/research/README.md`](../../../workflow/legal/research/README.md)
  現場查證；lesson 只能主張「該問什麼問題」，不能主張「法律規定是什麼」。
- Red tier 個案（勞動、股權、M&A、訴訟、行政處分）不在本 domain 沉澱實質建議，
  只沉澱「如何辨識並升級」的方法。

## Recent (2026-07-30)

| Slug | 摘要 |
|------|------|
| `common/2026-07-30_231500-verify-upstream-ip-chain-before-promising-source-delivery` | 承諾交付原始碼前，反向確認我方是否擁有全部交付物的權利。「出資聘請他人完成」在多數法域的預設是權利歸受聘人，不是出資的我方——此風險不在對方稿件裡，逐條審閱找不到。 |
| `common/2026-07-30_221637-counterparty-business-scope-drives-clause-terms` | 背調 Layer 2 的營業項目要問兩件事：締約能力涵蓋範圍，以及對手方是否受法規監理（後者會反向要求新增資料保存義務並擴充衍生責任排除）。Risk flag 的價值在於它改了哪一條條款。 |

## 首次 dogfood

本 domain 於 2026-07-30 隨 [`plans/active/2026-07-30-2101-legal-workflow-domain.md`](../../../plans/active/2026-07-30-2101-legal-workflow-domain.md)
建立。第一個真實任務為一份製造執行系統開發合約的出稿前自審（`review` task type、
TW jurisdiction、Yellow tier），驗證了 intake gate 確實在金額與交期未知時阻擋條款產出。
