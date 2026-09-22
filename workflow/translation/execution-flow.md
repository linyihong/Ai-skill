# Translation Decision — Execution Flow

Canonical lifecycle。欄位 SoT 在 [`contracts/`](contracts/)。**不要**在本檔寫 provider／prompt／model 步驟。

## Lifecycle

```text
0. Bind TranslationContext     → consumer / locale pack（authoritative locales）
1. Locale Resolution           → bind only; ≠ Language Detection
2. Expression Analysis         → artifact（any producer）
3. Registry lookup             → Candidate Space seeds
4. Apply Constraints           → Feasible Candidates (candidates[] + feasible)
5. Selection Policy + Actor    → selected + decision_basis（Actor ≠ Constraint owner）
6. Independent Review          → semantic / cultural dimensions
7. Mechanical + Locale Gate    → residue 分欄；locale residue → review 非 auto-fail
8. Finality                    → accepted only if I9 holds
```

## Stage 明細

| Stage | 讀／填 | 推進條件 | 失敗 |
| --- | --- | --- | --- |
| 0 Context | [`translation-context.yaml`](contracts/translation-context.yaml) | context 完整；有 target_locale | blocked |
| 1 Locale | 同上 | target_locale 來自 consumer，非 segment 猜測 | blocked |
| 2 Analysis | [`expression-analysis.yaml`](contracts/expression-analysis.yaml) | required fields 齊；producer 可任意 | needs_review |
| 3 Registry | [`registry/expression-types.yaml`](registry/expression-types.yaml) | type ∈ registry | unresolved |
| 4 Feasible | [`translation-decision.yaml`](contracts/translation-decision.yaml) | Candidate Space ≠ candidates[]；feasible 標齊 | blocked |
| 5 Select | 同上 | `selection.policy` + `selected` + `decision_basis`；selected ∈ feasible | needs_review |
| 6–7 Validate | [`validation.yaml`](contracts/validation.yaml) | 分欄 pass／fail／review；blocking 已處理或 waiver | needs_review |
| 8 Finality | [`finality.yaml`](contracts/finality.yaml) | I9 closure | 不得 accepted |

## 禁止

- `{ src, dst }` 當唯一 decision 輸入（I3）
- 每段 LLM 猜 `target_locale`（I2）
- 把 Expression Analysis 寫成「必須 AI」（I4）
- 從 Candidate Space 直接當 selected，跳過 feasible／policy（I5／I6／I10）
- LLM 改寫 constraints 或自證關閉 locale gate（I7）
- 合併 `source_language_residue` 與 `target_locale_residue`（I8）
- blocking 未清仍 `accepted`（I9）

## Static walkthrough（Phase 1 gate）

完成 contracts 後，用 [`examples/address-title-chen-xiaojie-id.yaml`](examples/address-title-chen-xiaojie-id.yaml) 走讀 Candidate Space → Feasible → Policy → Finality；**PASS 後再**開 Phase 2 subtitle adapter。不需 test runner。
