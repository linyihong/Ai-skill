# Phase 1/3 revision — ep8 Finality／Reference／Target Realization

Companion to [`_plan.md`](_plan.md)。Evidence：[`evidence/2026-09-29-ep8-ja-walkthrough.md`](evidence/2026-09-29-ep8-ja-walkthrough.md)。  
**不成** Phase 4；**不**把 女强人／绿帽子／姐夫 寫成唯一 mapping。

## Verdict on ep8

Workflow **sees** errors (needs_review／blocked) instead of swallowing them — direction confirmed.  
Next gap：too many REVIEW；must learn **when to PASS**.

| Area | Grade |
| --- | --- |
| Locale／JA-F binding／false-cognate／idiom／expansion／truncated block | ✅ |
| Reference Resolution | 🟡 formalize before Analysis |
| Target Realization（voice／gender／politeness） | 🟡 formalize |
| Finality PASS vs REVIEW vs BLOCK | 🟡 revise |
| Mapping DB of fixed glosses | ❌ do not build |

## Finality ternary（I22）

| Status | Alias | When |
| --- | --- | --- |
| `accepted` | **PASS** | required gates pass；`review.required=false`；no blocking reason |
| `needs_review` | **REVIEW** | explicit `review.reason[]` only — not “any uncertainty” |
| `blocked` | **BLOCK** | incomplete source／missing required context／hard constraint fail |
| `rejected` | — | validation fail without waiver path |

**禁止**：不確定 → 一律 needs_review（會變成人工重翻清單產生器）。

PASS examples from ep8：`糟了→まずい`、`你先回房休息→先に部屋で休んでて`、`一心都在工作上→仕事一筋で`（無特殊 voice／reference 衝突時）。

## Incomplete source（I21）

```text
incomplete_source_must_not_be_completed
Evidence insufficiency ≠ permission to infer missing content.
```

Cross-locale core。ep8 `#11`／`#24` BLOCK 是成功案例。

## Reference before kinship surface（JA-F05 retune）

```text
姐夫 → reference = sister's_husband → kinship realization Candidate Space
     ≠ fixed gloss 義兄さん
```

同構：陈小姐 → name + locale-aware title。

## Target Realization fields

```text
context: speaker / addressee / relationship / scene
target_realization: register / character_voice / politeness / gendered_expression
                  + locale/script/name/kinship realization
```

`我帮あなた` → 僕がやるよ／私がやるよ／手伝うよ 由 Selection 在 Constraint 內選。

## Knowledge discipline

存：`expression_type` + `strategy` + Candidate Space seeds。  
**不**存唯一答案：`女强人 = バリバリのキャリアウーマン`。

## Landing

| 產物 | 路徑 |
| --- | --- |
| Finality I21–I22 | `contracts/finality.yaml` |
| Reference | `contracts/reference-resolution.yaml` |
| Context／realization | `translation-context.yaml`、`translation-decision.yaml` |
| Flow | `execution-flow.md` |
| Ep8 retune | `examples/ep8-ja-walkthrough.yaml` |
