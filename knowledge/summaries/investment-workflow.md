## workflow.investment

| 欄位 | 值 |
| --- | --- |
| Atom ID | `workflow.investment` |
| Source path | `workflow/investment/execution-flow.md` |
| Lifecycle | `candidate` |
| Summary | 投資研究與決策輔助的 intake-dispatched workflow（Decision Support case #2）。第一級分派＝investment task type（need-framing／theme-research／name-diligence／position-review／allocation-advice／event-check／periodic-sweep）。Lifecycle：Frame → Intake → Risk Tier → Strategy Pass 1 → Dispatch → Research（analysis/investment）→ Pass 2 →（條件）DVA → Produce → Validate → Close。無舉證不得建議；uncertainty framing；D 級研究帳不可獨撐；配置須策略／資產／費用；allocation 預設 DVA。Route by decision object——「投資協議」走 legal，非 keyword「投資」。Report producer only（無券商／無 cron）。 |
| When to read | 主題／供應鏈研究、單一標的 diligence、持倉複核、配置／再平衡建議、事件影響、定期盯盤 sweep。 |
| Do not use for | ❌ 下單／券商 API。❌ 法律投資／股權協議（`route.workflow.legal`）。❌ 把個人研究帳當 Expert Knowledge。❌ 未校準精確勝率當定論。❌ Case C 混合句不得單靠本 route 吞掉協議審閱（Q12 still-open）。 |
| Context cost | ~400 tokens |
| Estimated full cost | ~3500 tokens（core）；含 analysis 方法約 ~6000 |
| Validation signal | `ai-skill runtime workflow-context --transcript <investment task>` → `active_route=route.workflow.investment`；contracts projected in `runtime.db`。 |
| Last checked | 2026-08-14 |
