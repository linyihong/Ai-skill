> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-07-30 - 對手方營業項目會反向改變條款建議

Status: candidate

#### One-line Summary

背調 Layer 2 的「營業項目」不只是身分核實欄位——對手方所處產業若受法規監理，會反向要求合約新增資料保存／匯出義務並收緊衍生責任排除。

#### Human Explanation

背景調查（due diligence）常被當成「確認對方是不是真公司」的一次性檢查，做完就附在附錄。實務上
Layer 2 的欄位裡，**營業項目**是唯一會直接改寫條款內容的一項：它決定對手方自身受哪些產業法規
拘束，而**受監理產業的客戶會把供應商系統產生的資料納入自己的法規遵循體系**。

常見誤判是把營業項目只用來檢查「締約能力是否涵蓋本交易」（scope mismatch），檢查完就跳過。
但當對手方屬於受監理產業（例如需要品質管理系統、批次可追溯性或稽核紀錄保存的製造業、
金融、醫療、食品、運輸），供應商交付的系統若產生追溯／履歷／過站／批號類資料，該資料
很可能被客戶用於因應主管機關查核。這會產生兩個原稿通常沒有的需求：

1. **資料保存期限與匯出義務**——客戶日後可能主張系統須保存 N 年並支援稽核匯出。若合約未約定，
   這個義務會在保固／維護期間被當成「本來就該有的功能」要求。
2. **衍生責任的排除範圍**——若客戶因系統缺陷而未通過法規查核，其主張的損失屬間接／衍生損失。
   標準的「排除停工／生產／營業／預期利益損失」清單未必涵蓋「法規查核未通過」與
   「主管機關處分」，需要明文補上。

#### Trigger

在合約審閱或起草流程中完成對手方背調，Layer 2 顯示其營業項目屬受法規監理產業，而交付標的
會產生追溯／履歷／稽核類資料。

#### Evidence

- Tool: 公開企業登記資訊查詢（C 層彙整來源，未回溯官方登記）
- Sanitized excerpt: 一份製造執行系統開發合約的委託方，其登記營業項目為受法規監理之製造業；
  合約標的包含批號與過站履歷追蹤，但原稿既無資料保存期限條款，責任限制條款的排除清單
  亦未涵蓋「客戶法規查核未通過」類衍生損失。
- Evidence path: `<PROJECT_ROOT>/review-memo-<date>.md`（專案 review memo 留在業務專案，不進本庫）

#### Generalized Lesson

背調 Layer 2 的營業項目要問兩個問題，不是一個：

1. （既有）營業項目是否涵蓋本交易標的？→ 締約能力 / scope mismatch
2. （本 lesson）**對手方是否處於受法規監理產業，且交付標的會產生其法規遵循所需的資料？**
   → 若是，觸發 `REGULATED_INDUSTRY_COUNTERPARTY` 風險旗標，對應兩項條款調整：
   資料保存期限與匯出義務、衍生責任排除範圍擴充。

這條的普適形式是：**risk flag 的價值在於它改變了哪一條條款，而不在於它描述了對手方什麼特徵。**
任何背調欄位若無法對應到條款調整，就只是資料，不是 flag。

#### Agent Action

下次在背調流程中：

- 讀到營業項目時，**除了** scope mismatch 檢查，**另外**判斷該產業是否受法規監理。
- 若是，在 Diligence Card 的「Risk Flags → 條款影響」表加一列，並在 Decision Register
  對應決策點（資料保存、責任限制排除項）中反映。
- **不要**把背調結果只寫成公司資料附錄；每一列都要能回答「這改了哪一條」。
- **不要**憑推測斷言該產業的具體法規義務內容——標 `unverified` 並列為律師覆核項；
  本 lesson 只主張「要問這個問題」，不主張任何法規適用結論。

#### Goal / Action / Validation

- Goal: 讓背調的 Layer 2 產出能改變條款，而不只是核實身分。
- Action: 在 `workflow/legal/due-diligence/README.md` 的 risk flag 表新增
  `REGULATED_INDUSTRY_COUNTERPARTY`，並在 Layer 2 檢查表的營業項目列加上第二個問題。
- Validation or reference source: `workflow/legal/due-diligence/README.md` §Risk Flags → 條款影響；
  首次實跑於一份製造執行系統合約審閱，該旗標產出兩項原稿缺漏的條款建議。

#### Applies When

- 有對手方且需核實其身分／營運狀態的合約任務。
- 交付標的會產生追溯、履歷、過站、批號、稽核紀錄類資料。
- 對手方營業項目落在需要品質管理系統、批次追溯或稽核保存義務的受監理產業。

#### Does Not Apply When

- 無對手方的任務（純法規查詢、內部評估）。
- 交付標的不產生對手方法規遵循所需之資料（例如純介面美化、內部工具）。
- 對手方營業項目為一般商業服務且無產業特別法拘束。
- Red tier 任務——該情形下不產出實質條款建議，僅升級。

#### Validation

- 對照 `workflow/legal/due-diligence/README.md` 的 flag 表，確認新增列具備「觸發條件」與
  「建議的條款調整」兩欄（缺條款調整欄即不符合本庫對 risk flag 的定義）。
- 在後續合約任務中檢查：受監理產業對手方的 Diligence Card 是否確實產出資料保存與
  責任排除兩項建議；若沒有，表示 flag 未被實際消費。

#### Promotion Target

- `workflow/legal/due-diligence/README.md`（本輪已落地：新增 risk flag 與 Layer 2 第二問）
- 未來若在多個 domain 重複出現「背調欄位須對應下游決策」的形態，可考慮提升為
  `workflow/cross-cutting/decision-support/` 的 instantiation contract 補充條件。

#### Required Linked Updates

- `workflow/legal/due-diligence/README.md` — 已更新（risk flag 表 + Layer 2 營業項目列）
- `feedback/history/legal/README.md` — 已建立 domain 索引
- `feedback/history/README.md` — 已在 domains 表新增 `legal`
- 依 [`reusable-guidance-boundary.md`](../../../../enforcement/reusable-guidance-boundary.md) 檢查：
  本檔不含當事人名稱、統一編號、地址、代表人姓名、金額或任何可識別單一專案之資訊；
  具體 review memo 與背調資料留在業務專案。
