# Translation Decision — Execution Flow

Canonical lifecycle。欄位 SoT 在 [`contracts/`](contracts/)。**不要**在本檔寫 provider／prompt／model 步驟。

## Lifecycle

```text
0. Bind TranslationContext     → consumer / locale pack（+ optional realization_profile）
1. Locale Resolution           → bind only; ≠ Language Detection
2. Expression Analysis         → artifact（any producer）
3. Target-Locale Realization   → name/title script／phonetic／form Candidate Space seeds
4. Registry lookup             → expression + realization strategy catalogs
5. Apply Constraints           → Feasible Candidates (candidates[] + feasible)
6. Selection Policy + Actor    → selected + decision_basis（Actor ≠ Constraint owner）
7. Independent Review          → semantic / cultural dimensions
8. Mechanical + Locale + Name  → residue 分欄；realization incomplete → review
9. Finality                    → accepted only if I9 holds
```

## Stage 明細

| Stage | 讀／填 | 推進條件 | 失敗 |
| --- | --- | --- | --- |
| 0–1 Context／Locale | [`translation-context.yaml`](contracts/translation-context.yaml) | authoritative target_locale | blocked |
| 2 Analysis | [`expression-analysis.yaml`](contracts/expression-analysis.yaml) | structure＋optional realization strategies | needs_review |
| 3 Realization | [`realization-strategies.yaml`](registry/realization-strategies.yaml) + knowledge seeds | Candidate Space 含 name／title realization | needs_review |
| 4–5 Registry／Feasible | [`translation-decision.yaml`](contracts/translation-decision.yaml) | I5／I11；feasible 標齊 | blocked |
| 6 Select | 同上 | policy + selected + decision_basis | needs_review |
| 7–8 Validate | [`validation.yaml`](contracts/validation.yaml) | locale ≠ name_realization 分欄 | needs_review |
| 9 Finality | [`finality.yaml`](contracts/finality.yaml) | I9 | 不得 accepted |

## 禁止

- `{ src, dst }` only（I3）
- 每段猜 `target_locale`（I2）
- Analysis＝必須 AI（I4）
- Candidate Space 直接當 selected（I5／I6／I10）
- LLM 改 constraints（I7）
- 合併 residue 欄（I8）
- `Chenさん` 當完整 ja name realization 卻無 review／waiver（I11）
- **所有**外國名非片假 → mechanical FAIL

## Static walkthrough（Phase 1 gate）

1. id-ID：[`examples/address-title-chen-xiaojie-id.yaml`](examples/address-title-chen-xiaojie-id.yaml)  
2. ja-JP：[`examples/address-title-chen-xiaojie-ja.yaml`](examples/address-title-chen-xiaojie-ja.yaml) — `Chenさん`→review；`チェンさん`→preferred pass  

兩案靜態走讀 PASS 後再開 Phase 2。不需 test runner。
