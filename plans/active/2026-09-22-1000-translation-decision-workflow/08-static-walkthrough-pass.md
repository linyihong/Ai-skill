# Static walkthrough PASS — 陈小姐 (id-ID + ja-JP)

Date: 2026-09-22  
Method: doc-only contract walk against Phase 1 SoTs + P0 fixtures. **No test runner.**  
Fixtures: [`address-title-chen-xiaojie-id.yaml`](../../workflow/translation/examples/address-title-chen-xiaojie-id.yaml), [`address-title-chen-xiaojie-ja.yaml`](../../workflow/translation/examples/address-title-chen-xiaojie-ja.yaml).

## Walkthrough A — zh → id-ID

| Step | Check | Result |
| --- | --- | --- |
| 1 Context | `target.locale=id-ID` authoritative（not LLM guess） | PASS |
| 2 Analysis | 陈／小姐 split；artifact shape | PASS |
| 3 Candidate Space | title seeds Nona／Miss／Ms. ≠ final | PASS |
| 4 Feasible | `Nona Chen` true；`Miss`/`Ms.` false + reason | PASS |
| 5 Selection | policy + `decision_basis` present；selected ∈ feasible | PASS |
| 6 Regression | `Miss Chen` → `locale_consistency=review` → not accepted w/o waiver | PASS |
| 7 Finality | `Nona Chen` + I9 → accepted | PASS |

## Walkthrough B — zh → ja-JP

| Step | Check | Result |
| --- | --- | --- |
| 1 Context | `ja-JP` + `realization_profile` | PASS |
| 2 Analysis | name_realization vs locale_aware title | PASS |
| 3 Realization space | チェン／陳／Chen × さん／ミス → composed forms | PASS |
| 4 Feasible | all three composed feasible；`Ms. Chen` false | PASS |
| 5 Distinction | `Chenさん` = title OK + **name_realization=review**（≠ locale residue） | PASS |
| 6 Preferred | `チェンさん` → name_realization pass | PASS |
| 7 Brake | non-katakana **not** hard FAIL | PASS |

## Verdict

**A+B PASS** → Phase 1 static gate closed → **Phase 2 subtitle adapter eligible**.

Not proven: live provider quality, full surname lexicon, runtime gates.
