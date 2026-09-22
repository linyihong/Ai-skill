# Translation Decision — Execution Flow

Canonical lifecycle。欄位 SoT 在 [`contracts/`](contracts/)。**不要**在本檔寫 provider／prompt／model 步驟。

## Lifecycle

```text
0. Bind TranslationContext     → consumer / locale pack（+ realization_profile）
1. Locale Resolution           → bind only; ≠ Language Detection
2. Content-Type Resolution     → content.type (subtitle|title|ui|…) + decision_mode
3. Title Structure Analysis    → when content.type=title（truncated? idiom? part?）
4. Expression Analysis         → artifact（any producer）
5. Target-Locale Realization   → name/script Candidate Space seeds
6. Registry lookup             → expression + realization + title strategies
7. Apply Constraints           → Feasible Candidates (+ decision_class)
8. Selection Policy + Actor    → selected + decision_basis
9. Independent Review          → semantic / cultural / invented_information
10. Mechanical + Locale + Name → residue 分欄；I11／I12
11. Finality                   → accepted only if I9 holds
```

## Stage 明細

| Stage | 讀／填 | 推進條件 | 失敗 |
| --- | --- | --- | --- |
| 0–1 Context／Locale | [`translation-context.yaml`](contracts/translation-context.yaml) | authoritative target_locale | blocked |
| 2 Content-Type | 同上 `content.type` | title ≠ ordinary subtitle sentence | needs_review |
| 3 Title Structure | [`expression-analysis.yaml`](contracts/expression-analysis.yaml) | title_structure when title | needs_review |
| 4 Analysis | 同上 | structure＋ambiguity on idioms | needs_review |
| 5–7 Realization／Feasible | decision + registries | I5／I11／I12；feasible 標齊 | blocked |
| 8 Select | [`translation-decision.yaml`](contracts/translation-decision.yaml) | policy + decision_basis | needs_review |
| 9–10 Validate | [`validation.yaml`](contracts/validation.yaml) | invented_information 分欄 | needs_review／fail |
| 11 Finality | [`finality.yaml`](contracts/finality.yaml) | I9 | 不得 accepted |

## 禁止

- `{ src, dst }` only（I3）
- 每段猜 `target_locale`（I2）
- Analysis＝必須 AI（I4）
- Candidate Space 直接當 selected（I5／I6／I10）
- LLM 改 constraints（I7）
- 合併 residue 欄（I8）
- `Chenさん` 當完整 ja name realization 卻無 review／waiver（I11）
- 片名／標題自行加 セクシー 等 marketing（I12）
- **所有**外國名非片假 → mechanical FAIL
- 另開 title-translation-workflow

Static walkthrough A+B **PASS** — [`08`](../../../plans/active/2026-09-22-1000-translation-decision-workflow/08-static-walkthrough-pass.md)。  
Title fixture — [`examples/title-kongjie-yiriqianli.yaml`](examples/title-kongjie-yiriqianli.yaml)。  
Subtitle adapter — [`adapters/subtitle.yaml`](adapters/subtitle.yaml)。
