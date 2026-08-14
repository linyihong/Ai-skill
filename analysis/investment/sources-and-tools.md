# Investment Sources and Tools（來源權威與工具邊界）

定義投資研究的 **source authority**、公開來源**類型**、付費摘錄規則與記錄格式。  
**不**硬編碼易 stale 的具體 URL；**不**集成付費資料商 API。

## Source authority（A–D）

用 **authority**，不用「Expert Knowledge」。

| 等級 | 類型 | 可支撐的 recommendation 強度 | 典型例子（類型，非固定 URL） |
| --- | --- | --- | --- |
| **A** | 官方／監管／公司法定揭露 | 可支撐較強主張（仍須 uncertainty framing） | 年報／季報、IR 簡報、8-K／重大訊息、交易所／證監會公告、央行統計 |
| **B** | 一級媒體或研究機構，且可回溯至 A 或可核對原始數據 | 中等；與 A 交叉後可升 | 具署名＋可追溯原始文件的財經媒體深度稿、券商公開報告（若當下可讀） |
| **C** | 二級媒體、彙整、paywall 轉述、無原文的摘要 | **弱**；不得獨撐「最適／必買」 | 彙整站、標題黨、付費牆後他人轉述、未附連結的社群摘要 |
| **D** | 個人研究帳、社群貼文、匿名論壇、口頭「大神」 | **Discovery only**；不可獨撐高信心建議 | X／部落格個人帳、Telegram 轉發、論壇 ID |

### 強制規則（H6）

1. **D 不可獨撐**：若主張僅有 D，recommendation 上限＝「待查／觀察」，不得寫成配置決策。  
2. **C 轉述須標註**：寫 `paraphrase-of-paywall` 或 `secondary-summary`，並試圖找回 A／B。  
3. **A／B 衝突**：並列兩側；不得靜默選一邊。  
4. **時效**：每列證據必有 `as-of` 日期；過期未复核不得當新證據。

## 公開來源類型（預設可用）

依市場（台／美）選類型；實作時用搜尋／官方站當下頁面，不把下列當永久連結清單。

| 類別 | TW 類型 | US 類型 |
| --- | --- | --- |
| 法定揭露 | 公開資訊觀測站／重大訊息、公司 IR | SEC EDGAR、公司 IR |
| 價格／成交（描述用） | 交易所／合法行情頁 | 交易所／合法行情頁 |
| 產業／供應鏈 | 官方產業報告、公司供應商揭露、可信產業媒體（B） | 同上 |
| 宏觀 | 主計／央行／官方統計 | Fed／BLS／官方統計 |
| 新聞 | 可溯源 mainstream（B）；否則 C | 同上 |
| Discovery | 個人研究帳 watchlist（D） | 同上 |

## 付費資料規則（Q7）

| 規則 | 說明 |
| --- | --- |
| 預設 | 只用公開來源 |
| 提醒 | Intake／設定說明可提醒「若有付費資料可提供摘錄」 |
| 寫入條件 | **僅當**使用者已提供、或當下授權可讀到該摘錄 |
| 禁止 | 接付費 API、假設訂閱存在、把「應有 Bloomberg」寫成已查 |

記錄付費摘錄時：來源等級通常 ≤B（視原始文件），並標 `user-supplied-excerpt`。

## 證據列最小格式

每筆寫入 evidence-ledger 至少：

| 欄位 | 必填 |
| --- | --- |
| Claim | 這列要支持或削弱的主張 |
| Source type | 上表類別 |
| Authority | A／B／C／D |
| Locator | URL 或「user-supplied excerpt #n」或文件名（去敏） |
| As-of | ISO 日期 |
| Stance | supports／weakens／context |
| Notes | 轉述？paywall？衝突？ |

## 推算鏈形狀（給報告用）

```text
Premise(s) [authority…]
  → Inference step(s)（可獨立覆核）
  → Uncertainty label（情境／信心帶；禁未校準精確 %）
  → Recommendation（≤ evidence）
  → Human selection（明示：使用者拍板／不代執行）
```

## 工具邊界

| 允許 | 禁止 |
| --- | --- |
| WebSearch／WebFetch 公開頁 | Broker／trading API |
| 讀使用者提供的本地摘錄／設定（專案本地） | 把持倉寫進 canonical `analysis/`／`workflow/` |
| 提醒使用者可設外部排程 | 本庫擁有 cron／跨工具 scheduler |

## 與其他層

- 權威語言與對照法：[`researcher-note-contrast.md`](researcher-note-contrast.md)  
- 主題地圖：[`theme-and-supply-chain.md`](theme-and-supply-chain.md)  
- Workflow gates（機率欄、DVA）：Phase 3 `artifact-gates.md`（尚未建）
