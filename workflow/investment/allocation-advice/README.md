# Sub-flow: allocation-advice

在策略＋資產＋**費用**約束下比較可行方案，產出有界「較有利」配置 brief。  
**預設強制 DVA**。

## Blocking intake

| 欄位 | 規則 |
| --- | --- |
| 策略 profile | 缺 → blocking |
| 現有資產 | 缺 → blocking |
| 費用 profile | 缺 → `provisional` 可續跑但 Interest 必標 provisional；不得假設零成本 |

## Interest Analysis（必做）

1. 對使用者較有利的安排（含少交易 vs 一次調整的費用差）  
2. 現實約束（流動性、集中度、禁止清單）  
3. 費用摩擦：手續費／稅／匯費／保管／管理／融資利息 → 再平衡門檻  

## 步驟

1. Intake S0–S3；寫策略／資產／費用摘要（去敏）  
2. Risk tier（通常 Yellow；命中 Red 表則短路）  
3. Pass 1：方案空間 ≥2；待查證市場前提  
4. Research：僅補方案比較所需證據（authority 封頂建議強度）  
5. Pass 2：四欄＋Interest＋uncertainty；**明示 human selection**  
6. **DVA**（除非 `skipped(user)`）：  
   - Orchestrator brief：acceptance 含禁未校準 %、禁「直接執行」、證據帽  
   - Executor：allocation 草稿  
   - Verifier：findings only（對齊 delegated-execution）  
   - 仲裁 fix／defer／reject → 定稿  
7. Gates  

## DVA brief 最小模板

```markdown
## DVA Brief — allocation
- Slice goal: allocation brief under strategy/asset/fee constraints
- Acceptance:
  - [ ] No uncalibrated success % / “high probability optimal”
  - [ ] Recommendation ≠ execute / human selection explicit
  - [ ] Evidence cap: recommendation ≤ ledger authority
  - [ ] Fee friction addressed or marked provisional
  - [ ] ≥2 alternatives with trade-offs
- Verifier_only challenges: planted % ; silent decision closure ; D-only support
- Transport: prefer fresh Task; if same-session role-sim, record limitation
```

## 產出

- Allocation brief（方案表＋費用淨效應）  
- Evidence-ledger＋推算鏈  
- Verifier 摘要＋仲裁  
- Disclaimer  

參考：dogfood `evidence/04`／`05`；DVA 跨域 `evidence/4a-…`。
