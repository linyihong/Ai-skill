# Phase 1 reinforcement — Semantic & Syntactic Realization

Companion to [`_plan.md`](_plan.md)。**不成**新 workflow；補 Source Understanding → Target Expression 中間層。

## Gap

已有：locale／idiom／title／name／failure_pattern。  
仍缺：語意角色、格關係、target syntax 重構、naturalness 分欄、social address（老張 ≠ オヤジ）。

## Core invariants（I14–I16）

1. **I14** — Translation preserves **semantic relations**, not source surface form.
2. **I15** — Target syntax may reorder／omit／restructure **only when** corresponding semantic relations remain preserved.
3. **I16** — Naturalness optimization MUST NOT introduce／remove／alter source-supported roles、entities、temporal、polarity、or intent.

## Layers

```text
Expression Analysis
├── lexical
├── semantic   → semantic_structure (predicate / args / temporal / …)
├── pragmatic
└── syntactic  → relations feeding Target Syntax Realization

Source Semantic Structure
  → Target Semantic Structure
  → Target Syntax Planning
  → Target Surface Sentence
  → Naturalness (Selection Responsibility; constrained by I16)
```

**Preserve semantic relations, not source word order.**  
例：`残業して12時まで` → `12時まで残業して` 可為合法 Target Syntax Realization。

## Semantic Role Preservation

`semantic_roles` 對照 source↔target；verifier 查 missing／mismatch／polarity／temporal。  
跨語言：Semantic Role → Target Grammatical Realization（が／に／を／省略）— **不是**逐詞翻介詞。

## Responsibility split

| Responsibility | Owns | Does not own |
| --- | --- | --- |
| Constraint | role completeness、polarity、numbers、identity、no invented info | which synonym sounds best |
| Selection | natural order、honorific choice、idiom equivalent、sentence reshape | rewriting constraints |
| Target Realization | semantic／syntax／naturalness **under** constraints | self-closing validation |

## Validation columns（分欄）

`semantic` ≠ `syntactic` ≠ `grammatical` ≠ `naturalness` ≠ `semantic_roles`

## Social address

`老張`：`老` = `social_address_marker`／familiarity — **≠** age=old → 禁止默認 `オヤジ張さん`。

## Failure taxonomy seeds（F1–F5）

| ID | Pattern |
| --- | --- |
| F1 | `semantic_role_loss` |
| F2 | `grammatical_relation_error` |
| F3 | `target_syntax_distortion`（死守 source word order） |
| F4 | `unnatural_surface_realization` |
| F5 | `social_address_misinterpretation` |

既有：idiom_literalization、target_locale_residue、name_script_mismatch、invented_information（+ I13 learning）。

## Landing

| 產物 | 路徑 |
| --- | --- |
| Analysis | [`expression-analysis.yaml`](../../../workflow/translation/contracts/expression-analysis.yaml) |
| Decision I14–I16 | [`translation-decision.yaml`](../../../workflow/translation/contracts/translation-decision.yaml) |
| Validation dims | [`validation.yaml`](../../../workflow/translation/contracts/validation.yaml) |
| Types／patterns | expression-types + failure-patterns |
| Fixture | [`social-address-laozhang.yaml`](../../../workflow/translation/examples/social-address-laozhang.yaml) |
