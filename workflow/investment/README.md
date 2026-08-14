# Investment Workflow

`workflow/investment/` 負責「投資研究與決策輔助的執行順序」。邊界是 **Decision Support**，
不是交易執行系統：不接券商、不下單、不擁有跨工具 cron。

> **狀態**：Phase 3 **doc-only**。`route.workflow.investment` **尚未註冊**（plan Phase 5）。
> Dogfood 證據見 [`plans/active/2026-08-14-1101-investment-research-decision-support/`](../../plans/active/2026-08-14-1101-investment-research-decision-support/_plan.md)。

> **語意邊界（Q8）**：utterance 含「投資」字樣**不是**本 route 的充分條件。
> Decision object＝證券／主題／配置 → 本 domain；＝合約／股權權利義務文件 → [`../legal/`](../legal/README.md)。
> Mixed（公司分析＋協議）→ framing 後 **primary + secondary**（見 plan evidence 06）；**不**偷升 multi-route runtime（Q12 open）。

> **P0**：無舉證不得建議；禁止保證報酬／「必買」話術；Recommendation ≠ Human Selection。
> 高利害 `allocation-advice` 預設走 DVA（見 [`artifact-gates.md`](artifact-gates.md)）。

## Investment task type（第一級分派）

| Task type | 意義 | 主 sub-flow |
| --- | --- | --- |
| `need-framing` | 還不知道要分析什麼 | [`need-framing/`](need-framing/README.md) |
| `theme-research` | 產業／供應鏈主題 | [`theme-research/`](theme-research/README.md) |
| `name-diligence` | 單一標的 | [`name-diligence/`](name-diligence/README.md) |
| `position-review` | 已持倉 thesis 複核 | [`position-review/`](position-review/README.md) |
| `allocation-advice` | 配置／風險配置（需策略＋資產＋費用） | [`allocation-advice/`](allocation-advice/README.md) |
| `event-check` | 單次新聞／事件 | [`event-check/`](event-check/README.md) |
| `periodic-sweep` | 定期巡檢（report producer；排程在外） | [`periodic-sweep/`](periodic-sweep/README.md) |

Task type 未知時**必須先**跑 `need-framing`，不得從單句需求靜默推定。

## 何時讀哪個檔（thin index）

| 認知階段 | 檔案 | load_when |
| --- | --- | --- |
| Lifecycle 與分派 | [`execution-flow.md`](execution-flow.md) | 任何投資任務入口 |
| 問什麼 | [`intake.md`](intake.md) | Stage 1；策略／資產／費用未確認時 |
| Decision Support | [`strategy/README.md`](strategy/README.md) | Pass 1／Pass 2（除純 Green 名詞整理可簡化） |
| 決策點 playbooks | [`strategy/decision-playbooks.md`](strategy/decision-playbooks.md) | 配置、費用門檻、深度、升 Red |
| 風險分級 | [`risk-classification.md`](risk-classification.md) | **每個任務** |
| 產出 gates | [`artifact-gates.md`](artifact-gates.md) | 產出前與宣稱完成前 |
| 取證方法 | [`../../analysis/investment/README.md`](../../analysis/investment/README.md) | Research stage |

## 核心原則

1. **Intake before advice**：策略／資產／費用缺欄 → blocking 或 `provisional`（見 intake）；配置建議不得假裝零成本。
2. **Two-pass Decision Support**：Pass 1 provisional → Research（analysis 方法）→ Pass 2 收斂；對齊 [`../cross-cutting/decision-support/`](../cross-cutting/decision-support/README.md)。
3. **Evidence-to-decision**：recommendation strength ≤ evidence authority（A–D）；見 analysis `sources-and-tools.md`。
4. **Uncertainty framing**：用情境／信心帶；禁未校準精確勝率 %（除非標 `source-stated-%` 且不採納為校準）。
5. **Source authority**：D（個人研究帳）不可獨撐高信心建議；拒 Expert Knowledge。
6. **Risk tier gates output**：Green／Yellow／Red 見 `risk-classification.md`；Red 停止實質交易指令。
7. **DVA for high-stakes**：`allocation-advice`／含資產之 `position-review` 強制（除非使用者明示跳過並記錄）。契約消費 [`../software-delivery/delegated-execution.md`](../software-delivery/delegated-execution.md) 與 [`../../plans/README.md`](../../plans/README.md) §Delegation loop SOP——**不**平行重寫一份 DVA。
8. **Report producer ≠ scheduler**：可提醒設排程；本 workflow 不執行 cron。
9. **Sanitization**：真實持倉／券商帳號／付費憑證不進 canonical；user-local only。

## 與既有層的關係

- 方法：[`analysis/investment/`](../../analysis/investment/README.md)
- Decision Support generic：[`cross-cutting/decision-support/`](../cross-cutting/decision-support/README.md)（本 domain = **converged case #2**；全庫 stage 仍待 3／3）
- Legal：協議／股權文件仍走 `workflow/legal/`
- DVA loop plan：[`plans/active/2026-07-08-0825-delegation-verification-arbitration-loop/`](../../plans/active/2026-07-08-0825-delegation-verification-arbitration-loop/_plan.md)（跨域證據 4a）

## Executable contracts

- Phase 5 才加 `execution-flow.yaml`／`artifact-gates.yaml` 並 compile。本階段 **markdown-only**。

## 驗證

- 能否說出 task type、risk tier、是否跑 DVA？
- 配置建議是否有策略／資產／費用摘要（或 provisional）？
- Evidence-ledger 是否可獨立覆核？Recommendation 是否留下 human selection？
