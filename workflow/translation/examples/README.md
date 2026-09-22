# Translation examples & regression fixtures

Phase 1 = **doc-only**. No test runner.

| File | Role |
| --- | --- |
| [`address-title-chen-xiaojie-id.yaml`](address-title-chen-xiaojie-id.yaml) | **P0 regression** — locale-aware title + `target_locale_residue` |
| [`literal.yaml`](literal.yaml) | literal |
| [`idiom.yaml`](idiom.yaml) | idiom |
| [`slang.yaml`](slang.yaml) | slang |
| [`proverb.yaml`](proverb.yaml) | proverb |
| [`dialect.yaml`](dialect.yaml) | dialect (anti-flattening) |
| [`wordplay.yaml`](wordplay.yaml) | wordplay |

## Static walkthrough checklist（陈小姐 → id-ID）

Use before Phase 2:

1. `translation_context.target.locale` = `id-ID` from consumer（not LLM guess）
2. Expression Analysis splits 陈 / 小姐
3. **Candidate Space** seeds titles（Nona, Miss, Ms.）— not final
4. **Feasible Candidates**：`Nona Chen` feasible；`Miss Chen`／`Ms. Chen` feasible=false
5. **Selection Policy** non-empty + `decision_basis` lists artifacts
6. `Miss Chen` without waiver → `locale_consistency=review` → **not** `accepted`
7. `Nona Chen` → pass → `finality.accepted` only if I9 holds

PASS this walkthrough → eligible for Phase 2 subtitle adapter.
