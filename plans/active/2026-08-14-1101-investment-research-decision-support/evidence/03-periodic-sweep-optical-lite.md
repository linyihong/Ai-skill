# Dogfood 03 — Periodic Sweep（H5：delta vs 全量重寫）

| 欄位 | 值 |
| --- | --- |
| Run ID | 03-periodic-sweep-optical-lite |
| Date | 2026-08-14 |
| Task type | `periodic-sweep` |
| Depth | Yellow（巡檢 brief；**非**交易指令） |
| Watchlist（去敏模擬） | Theme: AI optical／CPO；Name: $LITE |
| Prior state | [`01-theme-research-cpo-optical.md`](01-theme-research-cpo-optical.md) + [`02-name-diligence-lite.md`](02-name-diligence-lite.md) |
| Goal of run | 測 H5：產出是否為 **Delta／Reassessment**，而非重寫整份主題／diligence |

**Scheduler**：本 run 由對話觸發（Q5：workflow = report producer only）。

---

## Previous observation（凍結基線 — 只引用，不重抄全文）

| 基線 ID | 已成立／未決（摘要） |
| --- | --- |
| T-01 | CPO 進入量產／出貨早期敘事；瓶頸含引擎良率／封裝／上游；LPO／NPO／CPO 可能並存 |
| T-01 open | 「sold out／2028」類需求硬度不足（paywall／D） |
| N-02 | NVDA–LITE $2B＋採購／產能敘事有雙 A 級 IR |
| N-02 open | 「sold out through 2028」未升格；完整 8-K／財報細節待補 |

---

## New observation（本 sweep 新抓取）

來源仍分級；sweep **只記錄相對基線的新訊號**。

| ID | Delta claim | Source | Grade | vs 基線 |
| --- | --- | --- | --- | --- |
| D1 | FY26 Q4 營收／獲利加速敘事（媒體彙整：約 $1.01B、YoY 大增等） | [TIKR Q4 optics writeup](https://www.tikr.com/blog/what-lumentum-q4-earnings-reveal-about-the-next-leg-of-the-ai-optics-trade) | C | **新**：02 未鎖財報數字 |
| D2 | 管理層稱 high-power laser 供應落後需求；基板端有 AXT 協議／Greensboro 產能敘事 | 同上＋[druckfin call notes 2026-08-11](https://www.druckfin.com/en/articles/lumentum-says-high-power-laser-shortage-is-worsening-even-faster-than-demand-whi-20260811) | C | **新／強化**：供應缺口具體化（仍非 A 級 IR 原文） |
| D3 | **NPO 被描述為相對 CPO 的 additive TAM**；部分客戶以 NPO 為中間步 | druckfin／[Converge Digest](https://convergedigest.com/lumentum-ai-scale-up-optics-data-center/) | C | **新**：01 有「並存」假設；此為 **LITE 管理層／轉述層的 NPO additive 訊號** |
| D4 | Lead CPO 客戶時程「on track」；ultra-high-power laser **高量出貨敘事 ~2027 H2**，客戶部署語境 **~2028** | Converge Digest／call 轉寫 | C–D | **時間窗細化**（相對 01 粗 calendar）；**非** A 級 IR 原文 |
| D5 | 首次 ELS module PO 敘事（交付語境 2027 H2） | druckfin | C | **新產品形態訊號** |
| D6 | Researcher／大神筆記 | （本 sweep **未**強制重抓 Serenity） | — | **刻意省略**：A/C 已夠標 delta；H6：D 非必要 |

**未做**：重寫 NVDA $2B 協議全文（已在 02，**無變更則不重述**）。

---

## Delta table（核心產出）

| 項目 | 狀態 | 說明 |
| --- | --- | --- |
| NVDA $2B／策略協議事實 | **Unchanged** | 仍依 02 E1–E3；本 sweep 無推翻訊號 |
| CPO 主題仍相關 | **Unchanged** | 需求／供應敘事仍在光互連 |
| 並存架構 | **Updated** | NPO **additive** 敘事變具體（D3）→ 削弱「只押單一 CPO 赢家」 |
| 供應約束 | **Updated** | high-power laser 落後（D2）比 02 更尖 |
| 2027–28 時間窗 | **Updated（低權威）** | D4／D5 細化；升格需公司 IR／SEC |
| 「sold out 2028」開放項 | **Still open** | 仍無 A 級關閉 |

---

## Reassessment（相對 01／02）

| 問題 | Sweep 結論 |
| --- | --- |
| 主題是否失效？ | **否** |
| $LITE diligence 主線是否失效？ | **否**（協議基線未變） |
| 不確定性哪裡變了？ | **NPO additive** 與 **供應缺口／2027 H2–2028 時間窗** 權重上升；單一 CPO 敘事應更謙抑 |
| 需要升級 task？ | 建議後續：用 **公司 IR／8-K／10-Q** 核對 D1–D5（event-check 或 name-diligence refresh）；**不是**本 sweep 給買賣建議 |

### Uncertainty

| 標籤 | 內容 |
| --- | --- |
| Likely | 光互連／雷射供應緊張敘事持續出現在 C 級 earnings 轉述 |
| Plausible | NPO 與 CPO 並行擴大 TAM（依管理層轉述） |
| Unresolved | 精確出貨 digi／毛利／「sold out」硬指引 — 待 A/B |

### Recommendation boundary

- Sweep brief = **狀態變更＋待核對清單**  
- **禁止**：因 Q4 亮眼數字直接輸出買入  
- Human：是否觸發更深 refresh — 使用者／排程決定  

---

## H5 假說檢驗（本 run 主測）

### H5 Periodic Observation／Reassessment

| 欄位 | 內容 |
| --- | --- |
| Result | **PASS** |
| Observation | 產出是 Previous→New→**Delta**→Reassessment；**沒有**重貼 01 供應鏈全文或 02 協議全文 |
| Evidence | 本檔「Previous／Delta／Unchanged」結構；D6 省略 D 級重抓 |
| Counterfactual（FAIL 長什麼樣） | 若把 CPO 定義、NVDA 協議、三情境表再寫一遍當「sweep」→ **H5 FAIL**（全量重寫） |
| Why matters | Observation producer ≠ 每次重建知識庫；也 ≠ scheduler |
| Cross-domain? | **Likely**（法規／CVE／依賴版本巡檢同構） |
| Legal comparison | 「法規是否改版」應報 delta，不是重貼整部法典 |
| Consequence | **lean promote** Periodic Observation／Reassessment |

### 其他 H（簡表）

| H | Result | 一句 |
| --- | --- | --- |
| H1 | PASS | 只允許「更新追蹤／待核對」，未因 C 級財報轉述給交易建議 |
| H2 | PASS | Sweep 仍是短 lifecycle：基線→新證→調和→交人 |
| H3 | PASS | Likely／Plausible／Unresolved；無用假 % |
| H4 | PASS | Sweep ≠ allocation |
| H6 | PASS | 本 sweep **不需要**大神筆記即可標 delta（強化 02） |
| H7 | PASS | 無持倉入庫 |
| H8 | N/A | 未測 A/B/C |

---

## Friction

1. C 級 earnings 部落格／轉寫很多 → 易把 **轉述**當 **A 級**；delta 可記，升格要另跑。  
2. 若 agent「為了完整」重抄 01／02，H5 會假綠流程、真失敗 — 本 run 用 Unchanged 列刻意防呆。

---

## Next

→ Dogfood **04 allocation-advice**（虛構策略／資產／**手續費**；≥2 options＋Interest／費用＋uncertainty＋evidence ledger）。
