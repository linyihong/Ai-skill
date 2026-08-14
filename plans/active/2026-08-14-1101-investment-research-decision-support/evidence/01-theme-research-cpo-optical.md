# Dogfood 01 — Theme Research：AI DC 光互連／CPO 供應鏈 choke points

| 欄位 | 值 |
| --- | --- |
| Run ID | 01-theme-research-cpo-optical |
| Date | 2026-08-14 |
| Task types | `need-framing` → `theme-research` |
| Depth | Yellow（分析／主題研究；**非**買賣建議） |
| Markets | 台＋美語境；主題 AI／semi／光通訊 |
| Goal of run | 取 H1–H4、H6 baseline；**不是**測投資判斷準不準 |

---

## Need-framing（Pass 0）

| 問題 | 本 run 答案（模擬使用者已確認） |
| --- | --- |
| Decision object | **市場／供應鏈主題**（非合約、非單一 ticker 喊單） |
| 想達成什麼 | 弄清「AI 資料中心互連升級 → CPO／光模組鏈」的瓶頸與催化劑時間窗 |
| Horizon | 12–24 個月觀察窗 |
| 允許輸出 | Theme note＋不確定性；**不**輸出目標價／倉位 |
| 策略／資產／費用 | 本 run **不適用**（非 allocation）；標 N/A |

**Dispatch**：`theme-research`（非 legal；decision object = market/theme）。

---

## Pass 1 — Provisional model（Research 前）

| # | Provisional claim | 待查證 |
| --- | --- | --- |
| P1 | AI 訓練／推論叢集推高機架內外頻寬，傳統 pluggable＋長銅跡線功率／密度吃緊 | 官方／產業對 CPO 動機的公開說明 |
| P2 | CPO（光學引擎與 switch ASIC／加速器同封裝）是下一階段互連路徑之一 | 量產時程是否進入 2026 量產窗 |
| P3 | 瓶頸可能在光學引擎良率、先進封裝產能、InP／SiPh 上游集中 | 是否有可引用的公開研究／新聞（非單一個人帳） |
| P4 | LPO／NPO／CPO 可能並存，而非單一技術立刻取代 | 是否有「並存／分層」敘事 |

Confidence：**provisional**（尚未綁定來源）。

---

## Research — Evidence table

來源等級：A 官方／監管 · B 交易所／公司法定揭露 · C 主流產業媒體／研究機構公開稿 · D 個人研究／部落格／社群

| ID | Claim（支持／削弱） | Source | Grade | Date／取用 | 支持方向 |
| --- | --- | --- | --- | --- | --- |
| E1 | OpenAI Stargate 與 Oracle 等夥伴擴充 GW 級 AI DC 容量（互連需求的**下游需求語境**，非 CPO 本體證明） | [OpenAI — Stargate / Oracle 4.5GW](https://openai.com/index/stargate-advances-with-partnership-with-oracle/) | A | 取用 2026-08-14 | 支持「AI infra 規模持續擴大」；**不**直接證明 CPO 必勝 |
| E2 | OpenAI／Oracle／Vantage 等規劃 Wisconsin 近 GW 級站點（需求語境） | [DCD — Stargate Wisconsin / Vantage](https://www.datacenterdynamics.com/en/news/openai-oracle-and-vantage-plan-stargate-wisconsin-data-center-expected-to-be-close-to-a-gigawatt/) | C | 取用 2026-08-14 | 同 E1；媒體層 |
| E3 | TrendForce：NVIDIA Spectrum-X CPO 開始出貨給選定夥伴；Broadcom Bailly 51.2T CPO 持續有限出貨；量產階段；瓶頸含光學引擎良率與先進封裝（與 AI／HPC 搶產能） | [TrendForce press 2026-07-27](https://www.trendforce.com/presscenter/news/20260727-13151.html) | C | 2026-07-27 | 支持 P2／P3 |
| E4 | TrendForce 研究摘要：2026 為 CPO 量產元年敘事；2027–28 大規模採購窗；材料高度集中 | [TrendForce CPO report page](https://www.trendforce.com/research/download/RP260605GV) | C | 頁面摘要（付費全文**未**取得） | 支持 P2／P3；**全文未讀 → 不得升為高信心** |
| E5 | 產業說明文：CPO 定義、800G→1.6T、NPO／LPO 並存語境、封裝／熱／可維修挑戰 | [HTX Insights optical/CPO guide](https://www.htx.com/news/standing-in-the-light-a-comprehensive-guide-to-the-optical-m-W6PlkLJd/) | C–D | 取用 2026-08-14 | 支持 P1／P4；獨立性弱於 A/B |
| E6 | 個人／獨立供應鏈地圖敘事（90+ 公司、上游集中） | [leoinai Substack SiPh/CPO](https://leoinai.substack.com/p/supply-chain-photonics-and-co-packaged) | D | 取用 2026-08-14 | **僅**作地圖假設；**不得獨撐**高信心建議 |
| E7 | 互動供應鏈地圖（SOI／InP／封裝等 choke 敘事） | [leonardo-boquillon CPO map](https://leonardo-boquillon.com/photonic-cop-supply-chain) | D | 取用 2026-08-14 | 同 E6 |

**未納入**：Bloomberg 等付費終端（本 run 無可查付費摘錄）。依 Q7：僅公開來源。

---

## Pass 2 — Evidence reconciliation + uncertainty-framed note

### Theme thesis（不是投資指令）

AI 叢集規模擴大（E1–E2）提高互連頻寬／功耗壓力；產業公開敘事將 **CPO** 推入 **2026 量產／出貨早期**（E3–E4），同時強調 **光學引擎良率、先進封裝、上游材料集中** 為擴產瓶頸（E3–E4）。**LPO／NPO／CPO 可能分層並存**（E4–E5），不宜把主題收成「只買 CPO 單一贏家」。

### Uncertainty framing（非單一機率點估計）

| 情境 | 粗權重（主觀、可修正） | 含義 |
| --- | --- | --- |
| S1 CPO 沿 2026–28 量產／採購敘事推進，瓶頸決定節奏 | ~50% | 主題研究主線；時間窗與供應商集中度重要 |
| S2 技術／良率／標準化延遲，pluggable／LPO／NPO 更長時間主導增量 | ~30% | 削弱「CPO 立刻取代」；並存更久 |
| S3 需求放緩或 capex 延期，互連升級延後 | ~20% | 下游 AI DC 敘事若降溫則整鏈延後 |

> 這些權重**不是**由單一 A 級來源校準的統計估計；E3–E4 為 C 級機構稿／摘要。若只有 D 級地圖（E6–E7），**不得**把任一情境抬到「高信心」。

### Recommendation boundary（本 run）

- **允許**：持續追蹤光學引擎／封裝／InP–SiPh 上游與 hyperscaler 互連路線圖。  
- **不允許本 run 輸出**：ticker 買賣、倉位、目標價。  
- **Human selection**：是否把本主題列入個人 watchlist／後續 name-diligence — **使用者決定**。

### Catalyst calendar（粗）

| 時點 | 觀察項 | Confidence |
| --- | --- | --- |
| 2026 H2 | CPO switch 出貨／良率／封裝產能公開訊號 | medium（依 E3） |
| 2027–28 | 大規模採購敘事是否兌現 | low–medium（E4 摘要；全文未讀） |

---

## H1–H8 假說檢驗（本 run）

### H1 Evidence → Decision Gate

| 欄位 | 內容 |
| --- | --- |
| Result | **PASS**（在「主題研究」深度） |
| Observation | Pass 2 主張可回溯到 E1–E7；高影響力主張（量產／瓶頸）主要靠 E3–E4，未用 D 級獨撐 |
| Evidence | 見上表；明確拒絕「無來源喊單一赢家」 |
| Why this matters | Decision strength（只出主題追蹤建議）未超過 evidence strength |
| Investment-specific or cross-domain? | **Likely cross-domain**（legal 也有「無來源不得給策略建議」） |
| Legal comparison | Legal：無版本／來源的法條主張須標 unverified — 同構 |
| Candidate consequence | **defer** 升 cross-cutting；再等 allocation／legal 第二輪 |

### H2 Decision Support Lifecycle

| 欄位 | 內容 |
| --- | --- |
| Result | **PASS** |
| Observation | Frame → provisional → research → reconcile → bounded support → human selection；**未**一步 question→answer |
| Evidence | 本檔章節結構 |
| Why this matters | 證明 lifecycle 可在 investment theme 跑通，不依賴下單系統 |
| Investment-specific or cross-domain? | **Likely cross-domain**（對齊 decision-support two-pass） |
| Legal comparison | Legal Stage 3a／3b 同形 |
| Candidate consequence | **defer**（需 allocation 再驗證是否被跳過） |

### H3 Uncertainty Framing

| 欄位 | 內容 |
| --- | --- |
| Result | **MIXED** |
| Observation | 有三情境權重，但權重是主觀粗分，**易被誤讀成校準機率**；若讀者只看「50%」會過度確定 |
| Evidence | Uncertainty 表 + 明確註記「非統計校準」 |
| Why this matters | 證明「保留 uncertainty」有用，也證明 **數字機率可能是壞的 uncertainty UX** |
| Investment-specific or cross-domain? | **Cross-domain candidate**：應用 confidence bands／情景敘事，不必綁 probability % |
| Legal comparison | Legal 多用 provisional／confirmed／unverified，未必用 % |
| Candidate consequence | **Promote uncertainty framing 語言**；**reject** 過早 glossary=`probability-framed recommendation` |

### H4 Decision Depth Gate

| 欄位 | 內容 |
| --- | --- |
| Result | **PASS** |
| Observation | Yellow：允許主題 note；禁止交易指令；與 Green「名詞解釋」可區分 |
| Evidence | Need-framing depth + Recommendation boundary |
| Why this matters | Task depth → permitted output 可操作 |
| Investment-specific or cross-domain? | **Likely cross-domain** |
| Legal comparison | Green／Yellow／Red tier 同構 |
| Candidate consequence | **defer** |

### H5 Periodic Observation／Reassessment

| 欄位 | 內容 |
| --- | --- |
| Result | **N/A** |
| Observation | 本 run 非 sweep |
| Candidate consequence | 留給 run 03 |

### H6 Source Authority／Evidence Quality

| 欄位 | 內容 |
| --- | --- |
| Result | **PASS（with scar）** |
| Observation | A／C／D 分級有用；D 級地圖（E6–E7）易「看起來很完整」但權威低；E4 付費報告**只有摘要頁**卻常被當成全文真理 — 這是真實摩擦 |
| Evidence | E4 明標「全文未讀」；D 不得獨撐 |
| Why this matters | 「來源數量／地圖華麗度」≠ evidence strength；**researcher note 在本 run 不是必要條件**，只是 D 級 source type |
| Investment-specific or cross-domain? | **Likely cross-domain**（法律部落格 ≠ 法條） |
| Legal comparison | 官方條文 > 律師行文章 > 匿名貼文 |
| Candidate consequence | **defer** Source Authority model；**reject**「Expert／大神 Knowledge」抽象 |

### H7 Knowledge／User-State Boundary

| 欄位 | 內容 |
| --- | --- |
| Result | **PASS** |
| Observation | 本 evidence 無持倉／watchlist 個資；主題可進 plan evidence |
| Candidate consequence | **defer**（allocation run 才是硬測） |

### H8／Q8 Semantic Route Disambiguation

| 欄位 | 內容 |
| --- | --- |
| Result | **PASS（本 case）**／完整 A/B/C 仍 **pending** |
| Observation | 「CPO 供應鏈主題研究」decision object = market/theme → investment；無人會誤判成股權協議 |
| Candidate consequence | ⑥ 再測 A/B/C |

---

## Friction／failures worth keeping

1. **付費研究摘要頁（E4）**：容易在 dogfood 裡「看起來有 C 級機構背書」但 agent 沒讀全文 → 必須強制標「摘要 only」。  
2. **主觀機率表**：對 H3 是雙刃；下輪可改情景敘事＋confidence label，少用假精確 %。  
3. **下游 AI DC（E1）與 CPO 本體（E3）**：需求語境 ≠ 技術路線證明；混為一談會膨脹 recommendation strength。

---

## Next

→ Dogfood **02 name-diligence**（單一標的 + 新聞／趨勢 + 至少一則 researcher／公開研究對照；繼續餵 H6）。
