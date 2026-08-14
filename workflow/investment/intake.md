# Investment Task Intake（分層問卷）

Stage 1 canonical source：**問什麼、順序、何時可停**。

> **Blocking**：配置類（`allocation-advice`、含資產之 `position-review`）在策略／資產摘要
> 未齊前，不得輸出「應執行的配置指令」。可輸出 Frame＋blocking 問題＋標記假設的骨架。

## 問法

一輪 3–6 題；已能從訊息推定的標「推定值」供否認。  
**真實持倉／費用數字**優先讀使用者本地設定；缺欄才追問。**不**把持倉寫進本庫。

## S0 — 任何投資任務

| # | 問題 | 為何 |
| --- | --- | --- |
| S0-1 | **你要做什麼？**（task type 表） | 第一級分派 |
| S0-2 | **Decision object？** 證券／主題／配置 vs 合約／協議 | Q8；協議 → legal |
| S0-3 | **市場範圍？** TW／US／both／其他 | 預設 TW+US |
| S0-4 | **主題焦點？**（可空）預設 AI／semi／光通訊／DC 可改 | 縮小 Research |
| S0-5 | **輸出語言？** 預設繁中；ticker 保留英文 | — |
| S0-6 | **是否允許方向性／配置建議？** | 影響 depth |
| S0-7 | **硬期限／sweep 視窗？** | 決定深度 |

## S1 — 策略 profile（配置／持倉複核／有界「較有利」必問）

| # | 問題 |
| --- | --- |
| S1-1 | 目標：保值／成長／主題曝險／現金管理／其他 |
| S1-2 | Horizon |
| S1-3 | 風險承受與最大可接受回撤（自述即可） |
| S1-4 | 再平衡規則（若有） |
| S1-5 | 禁止標的／槓桿／ deriv 偏好 |
| S1-6 | 稅／匯率約束（若有） |

已有設定檔 → **確認摘要**即可。

## S2 — 現有資產（同上情境）

| # | 問題 |
| --- | --- |
| S2-1 | 現金約略 |
| S2-2 | 持倉清單（標的／約略權重或部位；去敏記錄） |
| S2-3 | 負債／流動性需求（若有） |

缺 S1／S2 → `allocation-advice` **blocking**；其他 task 可標 `UNKNOWN` 繼續研究但不給配置指令。

## S3 — 交易成本／費用 profile（Q13）

| # | 問題 |
| --- | --- |
| S3-1 | 券商手續費／佣金（TW／US 分開） |
| S3-2 | 交易稅（若適用） |
| S3-3 | 匯費／換匯 |
| S3-4 | 保管／管理費、ETF 費用率 |
| S3-5 | 借券／融資利息（若用） |

缺則標 `provisional` 並列待查；**不得**假裝零成本做 Interest Analysis。

## S4 — 依 task type

### need-framing
已知興趣、禁忌、時間盒；「成功長什麼樣」。

### theme-research
主題一句定義＋排除；要不要落到標的候選。

### name-diligence
Ticker／公司名；要對照的研究帳（user-local watchlist）。

### position-review
持倉代號；原 thesis（若有）；關心的失效條件。

### allocation-advice
S1–S3 必齊或 provisional；可比較方案數（建議 ≥2）；是否跳過 DVA（需明示）。

### event-check
事件描述、時點、相關標的。

### periodic-sweep
Watchlist 來源（user-local）；自上次 sweep 的 as-of；只要 delta。

## Investment Task Frame（輸出形狀）

```markdown
## Investment Task Frame
- Task type:
- Decision object: investment | legal | mixed→split
- Markets / theme:
- Risk tier: (Stage 2)
- Strategy / assets / fees: present | provisional | N/A
- DVA: required | skipped(user) | N/A
- Blocking questions:
- Assumptions:
```
