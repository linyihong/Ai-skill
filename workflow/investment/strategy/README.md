# Investment Strategy（Decision Support instantiation）

本 slice 是投資任務的**共同 Decision Support stage**（Pass 1／Pass 2），對齊
[`../../cross-cutting/decision-support/README.md`](../../cross-cutting/decision-support/README.md)
generic contract。本檔只放 **investment 專屬** inventory／verification／depth 接法。

**狀態**：Decision Support **converged case #2**；`route.workflow.investment` **已註冊**（Phase 5）。  
**Red**：不產出實質交易指令（見 [`../risk-classification.md`](../risk-classification.md)）。

## Instantiation 四項

| # | 項目 | 本域實例 |
| --- | --- | --- |
| 1 | Decision point inventory | 見下＋[`decision-playbooks.md`](decision-playbooks.md) |
| 2 | Playbooks | [`decision-playbooks.md`](decision-playbooks.md) |
| 3 | Verification source | [`../../../analysis/investment/`](../../../analysis/investment/README.md)（公開揭露／新聞／筆記對照） |
| 4 | Risk／depth gate | [`../risk-classification.md`](../risk-classification.md) |

## Decision point inventory（常出現）

1. 任務類型／分析深度（framing vs diligence vs allocation）  
2. 主題邊界與 choke-point 是否成立  
3. 單一標的 thesis：持有／觀察／避開（敘事，非下單）  
4. 配置方案選擇（A／B／…）與再平衡是否划算（**含費用**）  
5. 證據不足時：升級 Research vs 降級建議 vs 停止  
6. 是否觸發／跳過 DVA  
7. Human selection：使用者拍板點清單  

## Two-pass

```text
Intake → Pass 1 (provisional register + 待查證)
      → Research (analysis/investment)
      → Pass 2 (confirmed where sourced + Interest + human selection)
      → Produce (± DVA)
```

Pass 1 建議一律 `provisional`。Pass 2 方可對有 A／B 支撐的主張提高信心；**仍**保留 uncertainty framing。

## Decision Reasoning（強制四欄）

沿用 cross-cutting 格式。Investment 加欄：

- **Evidence cap**: 本建議不得超過的 authority 上限  
- **Fee / friction**: 費用假設或 `provisional`  
- **需使用者決策**: 必須為「是」當涉及配置／買賣方向  

## Interest Analysis（投資專化）

| 面向 | 內容 |
| --- | --- |
| (a) 對使用者較有利 | 在策略＋持倉約束下的方案淨效應 |
| (b) 其他方／現實 | 流動性、價差、公司行為、稅務現實 |
| (c) 費用摩擦 | 來回手續費、稅、匯費、保管／管理、融資利息；再平衡門檻 |

缺費用 → 標 provisional；**不得**假設零成本。

## 與 DVA

Pass 2 草稿若屬強制 DVA 情境 → 進 O→E→V→仲裁後才定稿。Verifier 不改建議。
