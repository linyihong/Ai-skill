# Translation Decision — Execution Flow

Canonical lifecycle。欄位 SoT 在 [`contracts/`](contracts/)。**不要**在本檔寫 provider／prompt／locale if-case 例句堆。

## Lifecycle

```text
0. Bind TranslationContext     → consumer / locale pack（+ realization_profile）
1. Locale Resolution           → bind only; ≠ Language Detection
2. Content-Type Resolution     → content.type + decision_mode
3. Title Structure Analysis    → when content.type=title
4. Expression + Semantic Analysis
   ├── lexical / pragmatic / syntactic / discourse / cultural
   └── semantic_structure + roles + modality（I14／I19）
5. Target-Locale Realization   → name/script／kinship Candidate Space seeds
6. Registry + Locale Taxonomy  → expression + **locale-failure-taxonomy** + patterns
7. Apply Constraints / Guards  → Feasible Candidates（core F* + locale JA-F*／…）
8. Selection + Target Realization
   ├── semantic realize
   ├── syntax realize（reorder OK under I15）
   └── naturalness（I16／I18；≠ grammatical）
9. Independent Review          → semantic／roles／modality／expansion
10. Mechanical + Locale + Grammar + Naturalness + Register
11. Finality                   → accepted only if I9 holds
12. (post) Failure Learning    → evidence → pattern candidate → Governance Review（I13）
```

## Locale Failure Taxonomy（step 6–7）

| 讀 | 路徑 |
| --- | --- |
| Index | [`registry/locale-failure-taxonomy.yaml`](registry/locale-failure-taxonomy.yaml) |
| Pattern bodies | [`registry/failure-patterns.yaml`](registry/failure-patterns.yaml) |

規則：`target_locale` 綁定後載入 `cross_locale_core` + `locale_taxonomies[target_locale]`。  
例：`ja-JP` → JA-F01–F10；`id-ID` → 最小 honorific residue；無證據 locale → core only（禁止發明 prompt 規則）。

## Stage 明細

| Stage | 讀／填 | 推進條件 | 失敗 |
| --- | --- | --- | --- |
| 0–2 Context／Type | [`translation-context.yaml`](contracts/translation-context.yaml) | authoritative locale + content.type | blocked／needs_review |
| 3–4 Analysis | [`expression-analysis.yaml`](contracts/expression-analysis.yaml) | structure＋roles／modality when needed | needs_review |
| 5–7 Realization／Guards | registries + **locale taxonomy** | I5／I11–I20；feasible 標齊 | blocked |
| 8 Select＋Target Realize | [`translation-decision.yaml`](contracts/translation-decision.yaml) | policy + decision_basis | needs_review |
| 9–10 Validate | [`validation.yaml`](contracts/validation.yaml) | 分欄；kinship residue fail | needs_review／fail |
| 11 Finality | [`finality.yaml`](contracts/finality.yaml) | I9 | 不得 accepted |
| 12 Learning | [`failure-pattern.yaml`](contracts/failure-pattern.yaml) | candidate ≠ active without review | — |

## 禁止

- `{ src, dst }` only（I3）
- token mapping without semantic roles（I14）
- require source word order（I15）／naturalness 加語義（I16／I18）
- 姐夫等 kinship 未譯（JA-F05）／緑の帽子 literal（JA-F08）
- LLM 自提升 failure_pattern → active（I13）
- Selection adapter `if locale: += 例句` 膨脹
- 對 truncated source 自行補全（I18）

## Dogfood walkthroughs

- Ep7 JA：[`evidence/2026-09-29-ep7-ja-dogfood`](../../../plans/active/2026-09-22-1000-translation-decision-workflow/evidence/2026-09-29-ep7-ja-dogfood.md)  
- Ep8 JA：**待人審** [`evidence/2026-09-29-ep8-ja-walkthrough`](../../../plans/active/2026-09-22-1000-translation-decision-workflow/evidence/2026-09-29-ep8-ja-walkthrough.md) · [`examples/ep8-ja-walkthrough.yaml`](examples/ep8-ja-walkthrough.yaml)
