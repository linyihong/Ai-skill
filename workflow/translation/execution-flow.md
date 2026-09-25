# Translation Decision — Execution Flow

Canonical lifecycle。欄位 SoT 在 [`contracts/`](contracts/)。**不要**在本檔寫 provider／prompt／locale if-case 例句堆。

## Lifecycle

```text
0. Bind TranslationContext     → consumer / locale pack（+ realization_profile）
1. Locale Resolution           → bind only; ≠ Language Detection
2. Content-Type Resolution     → content.type + decision_mode
3. Title Structure Analysis    → when content.type=title
4. Expression + Semantic Analysis
   ├── lexical / pragmatic / syntactic
   └── semantic_structure + roles（I14）
5. Target-Locale Realization   → name/script Candidate Space seeds
6. Registry + Knowledge lookup → expression + failure_patterns + locale seeds
7. Apply Constraints / Guards  → Feasible Candidates
8. Selection + Target Realization
   ├── semantic realize
   ├── syntax realize（reorder OK under I15）
   └── naturalness（I16；≠ grammatical）
9. Independent Review          → semantic／roles／invented_information
10. Mechanical + Locale + Grammar + Naturalness
11. Finality                   → accepted only if I9 holds
12. (post) Failure Learning    → evidence → pattern candidate → Governance Review（I13）
```

## Stage 明細

| Stage | 讀／填 | 推進條件 | 失敗 |
| --- | --- | --- | --- |
| 0–2 Context／Type | [`translation-context.yaml`](contracts/translation-context.yaml) | authoritative locale + content.type | blocked／needs_review |
| 3–4 Analysis | [`expression-analysis.yaml`](contracts/expression-analysis.yaml) | structure＋semantic_structure when multi-arg | needs_review |
| 5–7 Realization／Guards | registries + [`failure-patterns`](registry/failure-patterns.yaml) | I5／I11–I16；feasible 標齊 | blocked |
| 8 Select＋Target Realize | [`translation-decision.yaml`](contracts/translation-decision.yaml) | policy + decision_basis；I14–I16 | needs_review |
| 9–10 Validate | [`validation.yaml`](contracts/validation.yaml) | semantic≠syntactic≠grammatical≠naturalness | needs_review／fail |
| 11 Finality | [`finality.yaml`](contracts/finality.yaml) | I9 | 不得 accepted |
| 12 Learning | [`failure-pattern.yaml`](contracts/failure-pattern.yaml) | candidate ≠ active without review | — |

## 禁止

- `{ src, dst }` only（I3）
- token mapping without semantic roles for multi-arg clauses（I14）
- require source word order in target（I15）
- naturalness that drops／adds roles／entities／temporal／polarity／intent（I16）
- 老／小／阿 → 默認字面「老／年齡」（F5）
- LLM 自提升 failure_pattern → active（I13）
- Selection adapter `if locale: += 例句` 膨脹
- 另開 title-／semantic-translation-workflow

Companion：[`11-semantic-syntactic-realization`](../../../plans/active/2026-09-22-1000-translation-decision-workflow/11-semantic-syntactic-realization.md)。  
Fixture：[`social-address-laozhang.yaml`](examples/social-address-laozhang.yaml)。
