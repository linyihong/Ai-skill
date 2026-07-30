## workflow.legal

| 欄位 | 值 |
| --- | --- |
| Atom ID | `workflow.legal` |
| Source path | `workflow/legal/execution-flow.md` |
| Lifecycle | `candidate` |
| Summary | 法律工作的 intake-dispatched workflow。Domain 邊界是**法律任務**而非合約文件：第一級分派維度為 legal task type（draft / review / explain / compare / research / due-diligence / strategy / negotiation-support / lifecycle）。Lifecycle：Frame → Task Intake → Risk Tier → **Strategy Pass 1** → Dispatch → Due Diligence → Applicable Law → Reference Sources → **Strategy Pass 2** → Produce → Validate → Close。核心不是「問問題→寫合約」而是「問問題→推理最佳策略→使用者決策→才決定怎麼寫」。Jurisdiction 為 P0；策略建議強制四欄（Recommendation / Reason / Alternative / Trade-offs）；法規與官方範本引用強制版本 + 查核日；背調三層（Identity → Corporate Status → 九類 Risk Signals）且每個 risk flag 對應具體條款調整；風險分 Green / Yellow / Red，Red 短路為 Escalation Card。 |
| When to read | 使用者要起草／審閱／解釋／比較法律合約、查法規或官方契約範本版本、調查對手方公司、決定準據法／管轄／仲裁、設計付款或違約金條件、或需要談判支援時。 |
| Do not use for | ❌ 不是法律意見，不取代執業律師。❌ Red tier 領域（勞動個案、股權、投資、M&A、訴訟、政府行政處分、制裁命中、稅務規劃）不產出實質建議與策略推薦，只做升級。❌ 不可用於 API / BDD behavior contract——那是 `route.workflow.software-delivery`（裁決見 `workflow/workflow-routing.md` §「契約」語意裁決）。❌ 不可憑訓練記憶斷言法條；無官方來源即標 `unverified`。 |
| Context cost | ~380 tokens |
| Estimated full cost | ~3200 tokens（core 5 檔）；含全部 sub-flow 約 ~7500 tokens |
| Validation signal | `ai-skill runtime workflow-context --transcript <legal task>` 回 `active_route=route.workflow.legal`、`conflict=false`；execution-flow / artifact-gates 的 executable contract 已 project 到 `runtime.db generated_surfaces`。 |
| Last checked | 2026-07-30 |
