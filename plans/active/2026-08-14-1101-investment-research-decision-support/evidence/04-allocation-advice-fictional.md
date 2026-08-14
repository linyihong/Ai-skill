# Dogfood 04 — Allocation Advice（虛構策略／資產／費用）

| 欄位 | 值 |
| --- | --- |
| Run ID | 04-allocation-advice-fictional |
| Date | 2026-08-14 |
| Task type | `allocation-advice` |
| Depth | Yellow（配置決策支援；**非**下單指令） |
| Inputs | **全部虛構／去敏**（不得當真實持倉） |
| Theme evidence | 引用 01–03 為市場語境；**不**把 C 級轉述升成保證報酬 |

---

## Intake — Strategy / Asset / Fee profiles（虛構）

### Strategy profile（虛構）

| 欄位 | 值 |
| --- | --- |
| Goal | 主題成長（AI／光互連曝險）＋保留部分穩健 beta |
| Horizon | 18 個月 |
| Risk | 可承受 −25% 帳面回撤；單一標的硬上限 **25%** |
| Rebalance | 偏離帶 ±5% 才調；避免頻繁交易 |
| Constraints | 不使用融資／選擇權；台＋美皆可 |
| 禁止 | 保證報酬話術 |

### Asset / holdings（虛構紙上組合）

| 標的 | 權重 | 註 |
| --- | --- | --- |
| 現金（USD／TWD 混合，簡化為現金） | 40% | |
| 寬基科技 ETF（標籤 `BETA-ETF`） | 30% | 非單一公司 |
| 光通訊主題 ETF（標籤 `OPTIX-ETF`） | 20% | 虛構 ticker 標籤 |
| `$LITE`（紙上） | 10% | 接續 02／03 追蹤名 |

總權重 100%。**無真實帳戶 ID。**

### Fee profile（虛構假設；provisional 若未校準真實券商）

| 市場 | 假設來回成本（含佣金／雜費簡化） | 匯費／其他 |
| --- | --- | --- |
| 美股 | **0.20%** round-trip | 換匯假設 **0.40%**／次（若動用非美幣現金） |
| 台股 | **0.30%** round-trip（含簡化稅費假設） | — |
| ETF 內含費用 | `BETA-ETF` 0.05%／年；`OPTIX-ETF` 0.45%／年（虛構） | |

> 費用未知時應標 provisional。本 run **明示虛構費率**，僅測 Interest／費用摩擦是否進入比較。

---

## Pass 1 — Provisional options（Research 前）

| Option | 構想 | 待查 |
| --- | --- | --- |
| A | 12 個月內把 `OPTIX-ETF`+`$LITE` 合計從 30% → **40%**，現金降至 30%；分 **4 次**買入 | 光互連敘事是否仍成立；費用摩擦 |
| B | **維持**現權重；只做 sweep 追蹤 | 機會成本 |
| C | `$LITE` 一次加到 **25%** 上限，砍現金 | 集中度＋單名風險＋費用 |

---

## Evidence ledger（約束配置強度）

| ID | 主張 | 來源 | Grade | 可支撐的配置強度 |
| --- | --- | --- | --- | --- |
| E1 | NVDA–LITE 策略投資／採購敘事 | 02 A 級 IR | A | 允許「維持／溫和增加追蹤曝險」討論 |
| E2 | CPO／NPO／供應缺口敘事 | 01／03 | C | **方向性**；不可支撐「高信心重倉」 |
| E3 | 費用假設 | 本 intake 虛構表 | 假設 | 只影響 trade-off 計算，非市場預測 |

**Gate**：E2 為 C → **禁止**把 Option C 寫成「高機率最優」。

---

## Interest Analysis（含費用）

假設調整名義金額 **N = 100**（單位貨幣，虛構）。

| Option | 對使用者利益 | 其他方／現實 | **費用摩擦**（粗算） | 淨效應直覺 |
| --- | --- | --- | --- | --- |
| A | 提高主題曝險但仍分散（ETF+單名）；符 25% 單名上限 | 市場流動性通常足夠 | 分 4 次美股操作：約 4×(0.20%×交易額)；若每次動 ~2.5 單位買主題 → 主題增量 10 單位，來回成本量級 **~0.08–0.15 單位**＋可能匯費 | 費用低於「一次大調」；再平衡門檻合理 |
| B | 零交易費用；保留彈性 | 若主題續強，相對落後 | **~0** | 費用最優；曝險不變 |
| C | 高主題／高單名凸性 | 單名事件風險；違反「分散」精神雖未破 25% | 一次把 LITE +15 單位：成本 ~**0.03 單位**＋匯費；**低費用但高集中風險** | 費用不是主矛盾；**風險／約束**才是 |

**Interest 結論**：對「費用敏感＋單名上限＋18m 主題」的虛構使用者，**A 與 B 較可辯**；C 費用雖低，但把 recommendation strength 推高超過 E2 → **Interest＋Evidence 雙殺**。

---

## Pass 2 — 方案比較（≥2）＋ uncertainty

| Option | Trade-offs | Uncertainty 標籤 | 相對排序（約束下） |
| --- | --- | --- | --- |
| **A 分批加主題** | 費用略增；執行複雜度↑；保留分散 | **Plausible 較優**（在 E1+策略約束下） | **1（建議討論）** |
| **B 維持** | 零費用；可能錯過加碼窗 | **Plausible 合理解** | **2** |
| **C 衝 LITE 上限** | 費用低但單名／敘事依賴 C 級過重 | **Unresolved／不建議作為主方案** | **3（降級）** |

### 「較有利」定義（本 run 驗證）

較有利 = **策略約束 ∩ 證據強度 ∩ 費用摩擦** 下的排序，**不是**模型喊「A 最好／60% 會漲」。

### Recommendation boundary

- **輸出**：請使用者在 A vs B 決策；C 僅作反例。  
- **禁止**：目標價、保證出績、自動下單。  
- **Human selection**：最終選 A／B／混合 — 使用者。

---

## H 檢驗（本 run）

| H | Result | 一句 |
| --- | --- | --- |
| H1 | **PASS** | C 因證據弱被降級；未輸出強買 |
| H2 | **PASS** | Intake→Pass1→ledger→Interest→Pass2→human |
| H3 | **PASS** | Plausible／Unresolved；無假精確總勝率 |
| H4 | **PASS** | Allocation 深度＞ theme；門檻含費用／單名上限 |
| H5 | N/A | |
| H6 | **PASS** | 未用 D 級大神抬 C |
| H7 | **PASS** | 持倉為虛構標籤，未寫入 reusable knowledge 當事實 |
| H8 | N/A | |
| Q13 | **PASS** | 費用進入 Interest 與排序 |

**「較有利」是否有約束？** → **Yes（本虛構 case）**：A/B 可辯、C 被證據＋風險擋下。

---

## Next

→ **05 DVA**：對本卡有爭議的結論做 O→E→V→仲裁（Executor 若寫「Scenario A 60%」應被 Verifier 擋）。
