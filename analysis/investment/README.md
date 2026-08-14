# Investment Analysis Methods（投資研究分析方法）

本層回答「如何取證、拆解主題／標的、分級來源、保留 uncertainty」，**不是**下單流程、也不是 broker 整合。

**上游 dogfood**：[`plans/active/2026-08-14-1101-investment-research-decision-support/`](../../plans/active/2026-08-14-1101-investment-research-decision-support/_plan.md) Phase 1（①–⑥）。  
**執行編排**：[`workflow/investment/`](../../workflow/investment/README.md)（Phase 3 doc-only；`route.workflow.investment` 尚未註冊）。

## 範圍與邊界

### 屬於本層

- 主題／供應鏈拆解方法（choke-point、上下游曝險）
- **Source authority** 分級（A–D）與公開來源類型（無 stale 硬編碼 URL）
- 新聞／趨勢摘要模板
- 個人研究帳／「大神筆記」**對照**方法（引用、不同意點、時效）— **不是** Expert Knowledge
- 證據表欄位與推算鏈形狀（給 workflow evidence-ledger 用）

### 不屬於本層

- Agent 執行順序、intake 問題表、DVA 角色編排 → `workflow/`（Phase 3）
- 使用者真實持倉、券商帳號、付費訂閱憑證 → 專案本地／user settings（**不得**寫進本庫）
- 下單、接券商 API、跨工具 cron → **out of scope**（report producer only）
- 法律合約／投資協議審閱 → `workflow/legal/`（Q8 Case B／C）
- 跨域 generic primitive 定稿 → Decision Support follow-up（dogfood lean-promote 後另開）

## 何時進入此分析領域

1. 需要拆解 **主題／供應鏈**（例：AI／semi／光通訊／DC）再談標的曝險  
2. 需要對單一 **ticker／公司** 做公開資訊 diligence  
3. 需要把 **新聞／趨勢／研究筆記** 編成可覆核摘要，而非聊天結論  
4. 需要對照 **個人研究帳／社群筆記** 與官方／監管來源（H6）  
5. 配置／Interest Analysis 前，需要把市場主張壓成 **證據表＋authority 等級**

若 utterance 的 decision object 是 **合約權利義務**，即使含「投資」字樣 → **不要**進本層當主路由（見 Q8 Case B）。

## 方法地圖

| 檔案 | 用途 |
| --- | --- |
| [`sources-and-tools.md`](sources-and-tools.md) | Source authority A–D、公開來源類型、付費摘錄規則、記錄格式 |
| [`theme-and-supply-chain.md`](theme-and-supply-chain.md) | 主題／供應鏈拆解、choke-point、地圖產出形狀 |
| [`news-trend-summary.md`](news-trend-summary.md) | 新聞與趨勢摘要模板、時效與 delta 規則 |
| [`researcher-note-contrast.md`](researcher-note-contrast.md) | 研究帳／筆記對照（引用、不同意、時效）；禁 Expert Knowledge |

## 取證原則（Phase 1 疤 → 方法）

| 原則 | 來源 |
| --- | --- |
| Recommendation strength ≤ evidence strength | H1／DVA 05 |
| Uncertainty 用 **framing labels**（情境／信心帶），禁止無校準精確勝率 % | H3；01→05 修正 |
| D 級來源不可獨撐高信心建議 | H6；reject「Expert Knowledge」 |
| Paywall／C 級轉述易膨脹 → 標轉述並降級 | 01 scar |
| Recommendation ≠ Human Selection／Decision | DVA／ERA；寫回 DVA evidence 4a |
| 使用者持倉／費用 → user-local；canonical 只用去敏虛構或類型描述 | H7／Q10 still-open |

## 建議讀序

1. 本 README（邊界）  
2. `sources-and-tools.md`（權威分級）  
3. 依任務：`theme-and-supply-chain.md` 或 `news-trend-summary.md`／`researcher-note-contrast.md`  
4. 產出進未來 `workflow/investment/` artifact-gates（尚未建時：對齊 plan evidence-ledger 形狀）

## 與其他層的關係

| 層 | 關係 |
| --- | --- |
| `workflow/cross-cutting/decision-support/` | Investment = case #2 **lab**；本層餵 Pass1→Research→Pass2 的 Research 方法 |
| `workflow/legal/` | 協議／股權文件；Case C = primary+secondary，不單靠本層 |
| `analysis/travel/`、`analysis/web/` | 同源「sources-and-tools」形狀參考 |
| `plans/…/delegation-verification-arbitration-loop/` | 高利害報告消費 O→E→V；本層只定義 Verifier 可檢查的證據形狀 |
| `intelligence/` | 尚無 investment intelligence；有穩定 lesson 再抽，勿空建 |

## Inbound References

- Plan：[`2026-08-14-1101-investment-research-decision-support`](../../plans/active/2026-08-14-1101-investment-research-decision-support/_plan.md)  
- Dogfood：同 plan `evidence/01`–`06`  
- DVA 跨域回饋：[`4a-investment-decision-support-allocation-dva.md`](../../plans/active/2026-07-08-0825-delegation-verification-arbitration-loop/evidence/4a-investment-decision-support-allocation-dva.md)

## 狀態

- Phase 2 方法層：**active（doc-only）**  
- `route.workflow.investment`：**未註冊**（Phase 5）  
- Glossary candidates：**延後註冊**（plan Glossary Impact）
