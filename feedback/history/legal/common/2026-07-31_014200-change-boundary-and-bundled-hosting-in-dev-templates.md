> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-07-31 - 委外開發範本的兩類高頻缺口：變更邊界與主機捆綁

Status: candidate

#### One-line Summary

供應方軟體委外範本常有完整的變更五步與默示驗收（優點應吸收），但「範圍內免加價」若無可判定邊界、以及「開發費含一年主機／維護」若無 SLA／續約／遷移，仍是 Major／Critical 級缺漏——結構看起來齊全不代表 Stage 4 可過。

#### Human Explanation

公開的供應方軟體委託開發範本，常把變更流程寫得很細（申請→評估→確認→重訂時程→實施與紀錄），並搭配默示驗收、保固與範圍外收費、依已收款階段的終止退款／交付。這些是**值得寫進 draft Protection Pass 與 playbook 的優點**。

同一類範本也常同時出現兩個高頻缺口：

1. **「未超出附件範圍不得要求額外費用」卻沒有可判定的範圍邊界。**  
   變更程序再完整，若 SOW／附件一本身模糊，「範圍內」變成供應方無法拒絕的無償工單，或變成雙方對「算不算超出」各說各話。五步流程不能替代邊界定義。

2. **固定總價寫「含一年基本款主機與維護」，卻沒有主機規格、可用性、維護範圍、續約與結束遷移。**  
   等於把未定範圍的運維塞進已定開發價金。審閱時若只看「有沒有維護條款」會誤判為齊全；正確問法是「維運範圍是否可執行、期滿怎麼結束」。

另外常見內在張力：同時寫「直至甲方確認合格為止」與「N 日未回覆視為合格」卻未寫優先序——Stage 2 應標出自動效果條款的衝突。

#### Trigger

審閱或起草軟體／委外開發合約（尤其供應方標準範本），且稿內出現「範圍內變更免加價」或「開發費含主機／維護」任一字樣。

#### Evidence

- Tool: 公開供應方軟體委託開發合約範本的 unsigned `review` dogfood（TW、Yellow：智財＋金額空白）
- Sanitized excerpt: 範本含完整變更五步與默示驗收；同時以「未超出附件範圍不得加價」與「含一年基本款主機與維護」定價，附件與 SLA 均未在正文可判定；另見「直至確認合格」與「逾期視為合格」並存。
- Evidence path: 個案 Review Memo 留對話／業務專案，不進本庫（本 lesson 只沉澱方法）

#### Generalized Lesson

1. **吸收優點**：變更五步、階段確認天數、默示驗收方向、保固 vs 新功能、依已收款階段的終止義務——寫進
   [`workflow/legal/draft/README.md`](../../../../workflow/legal/draft/README.md) Protection Pass 與
   [`decision-playbooks.md`](../../../../workflow/legal/strategy/decision-playbooks.md)。
2. **同時擋缺口**：Stage 4 Missing Clause 對軟體委外固定檢查  
   - 範圍內免加價的**判定邊界**是否存在；  
   - 主機／維運捆綁是否拆出 SLA／續約／遷移；  
   - 默示驗收與「直至確認合格」的**優先序**。
3. **結構齊全 ≠ 缺漏已清**：Stage 1 骨架可過關的範本，仍可能在 Stage 4 因邊界與捆綁範圍不合格。

#### Agent Action

下次遇到軟體委外範本：

- Protection Pass／缺漏表：變更五步與邊界、保固分界、主機捆綁、終止依收款階段——逐項勾。
- 看到「範圍內免加價」→ 立刻問：附件範圍如何客觀判定？評估工時公式？誰有最終認定權？
- 看到「含主機／維護」→ 立刻問：規格、SLA、期滿、遷移是否可執行？不可則標 Major／Missing。
- 同時存在默示驗收與無限修正至合格 → 標內在張力，要求寫優先序。
- **不要**因變更流程寫得細就跳過 Stage 4；**不要**把對方公司名或具體金額寫進 reusable lesson。

#### Goal / Action / Validation

- Goal: 讓軟體委外範本的優點進 workflow，同時把「看起來完整」的高頻缺口變成可檢查項。
- Action: 更新 `draft/README.md`、`strategy/decision-playbooks.md`、`intake.md`、`review/README.md` Stage 4；本 lesson 索引於 `feedback/history/legal/README.md`。
- Validation or reference source: 上述四檔含對應檢查項；本輪 dogfood Review Memo 的缺漏表命中變更邊界、主機捆綁、驗收張力三項。

#### Applies When

- Task type 為 `draft` 或 `review`，交易類型為軟體／委外開發或含客製交付的服務合約。
- 稿內出現變更免加價、保固、主機／維運捆綁、默示驗收等供應方範本常見條款。

#### Does Not Apply When

- 純法規查詢、純背調、Red tier（本 lesson 不產出實質法律意見）。
- 政府採購走官方範本路徑時——變更／驗收依招標文件與機關增修，不套用民間供應方範本假設。

#### Related

- [`workflow/legal/draft/README.md`](../../../../workflow/legal/draft/README.md) §變更控制五步、§主機或維運捆進開發費
- [`workflow/legal/strategy/decision-playbooks.md`](../../../../workflow/legal/strategy/decision-playbooks.md) §付款、§驗收、§終止
- [`workflow/legal/review/README.md`](../../../../workflow/legal/review/README.md) Stage 4
