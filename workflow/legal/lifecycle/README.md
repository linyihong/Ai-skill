# Contract Lifecycle（端到端合約生命週期）

適用 task type `lifecycle`：從「有一個需求」到「簽完存檔並可追蹤」的統籌流程。
其他 task type 是這條流程的**片段**；本 slice 負責串接與交接欄位。

```text
Need → Intake → Strategy → Research → Draft → Review → Negotiate
     → Revise → Approve → Sign → Archive
```

> **Q6 邊界（本輪未接外部工具）**：`Approve` / `Sign` / `Archive` 三段目前只定義流程與
> 交接欄位，**不接**任何電子簽署服務或文件庫。接入前需另開 plan 並取得使用者授權。
> 見 [`../../../plans/active/2026-07-30-2101-legal-workflow-domain.md`](../../../plans/active/2026-07-30-2101-legal-workflow-domain.md) §Open Questions Q6。

## 各段的動作與交接欄位

| # | 段 | 動作 | 進入下一段需具備 | Canonical source |
| --- | --- | --- | --- | --- |
| 1 | **Need** | 記錄商業需求與觸發事件（誰要、為什麼、什麼時候要） | 需求敘述 + 內部需求方 + 硬期限 | 本檔 |
| 2 | **Intake** | 分層問卷；產出 Legal Task Frame | `gate.legal.intake_complete` | [`../intake.md`](../intake.md) |
| 3 | **Strategy** | Decision Register（Pass 1）；決策點與待查證前提 | `gate.legal.decision_register_present` | [`../strategy/README.md`](../strategy/README.md) |
| 4 | **Research** | 適用法與版本核對；對手方背調 | `gate.legal.law_citation_versioned` + `gate.legal.counterparty_identified` | [`../research/README.md`](../research/README.md)、[`../due-diligence/README.md`](../due-diligence/README.md) |
| 5 | **Draft** | Strategy Pass 2 確認後落成條款 | `gate.legal.strategy_reasoned` + Draft Handover | [`../draft/README.md`](../draft/README.md) |
| 6 | **Review** | 內部審閱（或對方稿審閱）五階段 | Review Memo（含風險分級） | [`../review/README.md`](../review/README.md) |
| 7 | **Negotiate** | Issue Ledger + 讓步矩陣 + 版本追蹤 | 未決爭點清單 + 每項的處置 | [`../negotiation/README.md`](../negotiation/README.md) |
| 8 | **Revise** | 依談判結果改稿並回歸一致性檢查 | 定義／交叉引用／附件／雙語一致 + 變更對照表 | [`../draft/README.md`](../draft/README.md) Step 5 |
| 9 | **Approve** | 內部核決 | 核決層級確認 + 未決事項已揭露給核決者 | 本檔 §Approve |
| 10 | **Sign** | 簽署 | 簽署人權限文件 + 附件齊備 + 生效條件確認 | 本檔 §Sign |
| 11 | **Archive** | 存檔與後續追蹤 | 存檔位置 + 關鍵日期已登記 | 本檔 §Archive |

**回圈是正常的**：Negotiate → Revise → Review 會多輪；Revise 若動到 Decision Register
的決策，必須回 Strategy Pass 2 重新確認，不可只改文字。

## Approve（內部核決）

| 檢查 | 內容 |
| --- | --- |
| 核決層級 | 依金額／期間／風險 tier 決定需要誰核准 |
| 揭露完整性 | **未決事項與已知風險必須揭露給核決者**，不可只送乾淨版 |
| 例外核准 | 若有偏離我方標準或 red-line 的讓步，需標明並單獨取得核准 |
| 條件核准 | 記錄「以 X 為條件的核准」，並在簽署前確認條件成立 |

**最常見的失效**：把 Review Memo 的 Critical 風險留在法務端，核決者看到的是「已審閱完成」。

## Sign（簽署）

| 檢查 | 內容 |
| --- | --- |
| 簽署人權限 | 登記代表人或有效授權書（→ [`../due-diligence/README.md`](../due-diligence/README.md) Layer 1） |
| 主體正確 | 簽約主體與背調確認的法人一致（非品牌名、非未登記名義） |
| 附件齊備 | 正文引用的每個附件都存在且為最終版 |
| 版本鎖定 | 簽署版本與最後確認版本一致（比對 hash 或逐頁確認） |
| 形式要求 | 依法域與交易類型：書面／印鑑／電子簽章的效力要求 |
| 份數與保管 | 各方份數、正副本 |
| 生效條件 | 生效日；若附條件（付款、擔保、許可），條件成立前不視為生效 |

## Archive（存檔與追蹤）

存檔不只是放檔案，是**登記需要被提醒的日期與義務**：

| 登記項 | 為什麼 |
| --- | --- |
| 合約期間屆滿日 | 避免被自動續約 |
| **自動續約的拒絕通知截止日** | 最常錯過的日期；通常需提前 N 日通知 |
| 付款節點 | 收付款追蹤 |
| 驗收期限與修正期 | 逾期視為通過的風險 |
| 保固／保證期屆滿 | 主張權利的期限 |
| 保密期間屆滿 | 分層期間各自的屆滿日 |
| 擔保／保證金返還條件 | 避免遺留資金 |
| 存續條款清單 | 終止後仍有效的義務 |
| 變更與補充協議 | 與主約連結，避免只找到主約 |

存檔欄位建議：合約名稱、對方主體全稱與登記號、簽署日、生效日、期間、金額、
準據法與爭議解決、關鍵日期清單、風險 tier、Review Memo 連結、
Decision Register 連結、未決事項。

## 驗證

- 每段進入下一段前的交接欄位是否齊備？
- Revise 若動到 Decision Register 決策，是否回 Strategy Pass 2 重新確認？
- Approve 是否把未決事項與 Critical 風險揭露給核決者？
- Sign 前是否確認簽署人權限、主體、附件、版本、生效條件？
- Archive 是否登記了**自動續約拒絕通知截止日**與存續條款清單？
