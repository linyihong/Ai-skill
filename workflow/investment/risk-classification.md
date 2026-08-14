# Risk Classification（Green / Yellow / Red）

Stage 2 canonical source。Tier **決定產出深度**，不是文末免責。

> Stakeholder：允許舉證後的方向性建議；**無舉證不得建議**；禁止說死。  
> 與 legal Red「股權／投資契約」不同——那是法律文件任務；本檔管**市場研究／配置建議**風險。

## 判定規則

取**最高**命中 tier。

### Green — 完整研究整理

| 條件 | 例 |
| --- | --- |
| 名詞、產業地圖、公開來源彙整 | 供應鏈節點表、詞彙整理 |
| 無方向性買賣／配置指令 | 「這是 CPO 節點」 |
| 證據以 A／B 為主或明確標 C／D 僅 context | |

**產出**：research note／地圖；可附 open questions。

### Yellow — 可建議，必須舉證＋uncertainty

| 條件 | 為什麼 |
| --- | --- |
| 單一公司 diligence、事件影響 | 錯誤成本高 |
| 主題 thesis 含標的候選 | 易被讀成喊單 |
| 配置／再平衡／「較有利方案」 | 影響真實資金 |
| 使用 C／D 來源支撐敘事 | 權威不足 |
| 機率化或情境權重建議 | 易假精確 |

**產出必須**：
1. Evidence-ledger（主張→authority→locator→as-of）  
2. Uncertainty labels（禁未校準精確勝率 %）  
3. Recommendation strength ≤ evidence  
4. 明示 **需使用者決策／不代執行**

### Red — 停止實質交易指令

| 條件 | 行為 |
| --- | --- |
| 要求保證報酬／「必漲／必買／必賣」且拒絕改寫 | 停止該指令；改 Escalation／可做的研究範圍 |
| 無來源喊單、拒絕舉證 | 停止建議 |
| 複雜槓桿／期權／衍生結構且關鍵條款未核實 | 停止結構性交易指令；列待查 |
| 要求代接下單／接券商／給出可直接送單的指令序 | **停止**；本 workflow 無執行權 |
| 法律投資協議／股權文件審閱（誤入） | 轉 `workflow/legal/`；本域不產法律意見 |

**Red 仍可**：整理公開事實、列文件／數據待查、產出 Escalation Card。

```markdown
## Escalation Card（investment）
- Risk tier: Red
- Trigger:
- Why no trade instruction:
- What research is still allowed:
- Docs / data to prepare:
- Suggested human next step: (e.g. licensed advisor / self decision with ledger)
```

**禁止**：用 disclaimer 當通行證輸出保證報酬或代客下單步驟。

## 升降

| 情境 | 處置 |
| --- | --- |
| 中途發現無證據卻要強建議 | 升 Yellow／Red |
| 使用者堅持保證口吻 | 說明一次界線；不降 tier |
| 「只是參考不會下單」 | tier 不因這句話降級 |

## 與 DVA

Yellow／Red 邊界之 `allocation-advice`、含資產 `position-review` → **強制 DVA**（除非明示跳過並記錄）。Verifier 必挑戰：未校準 %、Recommendation→Decision 僭越、D 獨撐。
