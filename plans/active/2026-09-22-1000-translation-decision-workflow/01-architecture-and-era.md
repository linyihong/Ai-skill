# Architecture — Translation ERA & six layers

Companion to [`_plan.md`](_plan.md)。

## 定位

**Translation Decision Workflow**，不是 **Translation Pipeline**。核心邊界：

```text
             ┌── Mechanical validation
             │
Source ──────┼── Constraints (Constraint Responsibility)
             │
             └── Evidence
                     ↓
              Feasible candidate space
                     ↓
               LLM Selection Actor (Selection Responsibility)
                     ↓
              Independent review (semantic / cultural)
                     ↓
              Mechanical gate
                     ↓
                 Finality
```

LLM 負責理解與選擇；Runtime／workflow 負責限制、驗證與收斂。

## Translation ERA（對齊 ERA v2）

```text
Evidence
  → Expression Space
  → Constraints
  → Feasible Candidates
  → Selection Policy
  → Translation Decision
  → Validation
  → Finality
```

Evidence **不直接決定譯文**；Evidence 約束 candidate space，再由 Selection Policy 選定。

## Locale Resolution（P0，在 Analysis 之前）

```text
Source segment
  ↓
Language / Locale Resolution   ← mandatory; before Expression Analysis
  ↓
TranslationContext (source + target locale, script, register hints)
  ↓
Expression Analysis
  ↓
…
```

若 `target locale` 只存在 pipeline 最外層、未進入每段 Translation Decision，稱謂／成語／人名／方言段會被模型 **自行假設 target language**（典型：zh→en 稱謂）。Dogfood：[`04-dogfood-case-address-title-id-ID.md`](04-dogfood-case-address-title-id-ID.md)。

**Invariant**：translation actor 輸入 = `TranslationContext` + segment + evidence；**禁止** decision 路徑僅 `{ src, dst }`。

## 七步決策鏈（邏輯模型）

與六層執行模型對應；細節在 `execution-flow`（Phase 1）。Step 0 = Locale Resolution。

1. **Source Understanding** — language／locale、speaker／addressee、meaning、intent、register、cultural markers、ambiguity
2. **Expression Classification** — registry lookup（literal、idiom、proverb、slang、meme、dialect、wordplay、honorific、cultural_reference、proper_noun／terminology）
3. **Translation Strategy Selection** — literal、semantic_equivalent、cultural_equivalent、adaptation、explanation、preserve + annotation
4. **Candidate Translation** — 多候選，每個標 strategy
5. **Linguistic / Cultural Review** — meaning、intent、tone、culture、character voice
6. **Mechanical + Locale Validation** — placeholders、names、numbers、tags、punctuation、length、source-language residue；**locale_consistency**（target language／locale、register、title-honorific、dialect）；subtitle adapter 另接 timing／layout（NVP）
7. **Finality** — `accepted` | `needs_review` | `unresolved` | `rejected`

## 六層執行責任

| Layer | 名稱 | 主要產出 |
| --- | --- | --- |
| 1 | INGEST | **TranslationContext**（source／target locale）+ segment + evidence refs |
| 1b | LOCALE RESOLVE | 固定 target locale；禁止 implicit zh→en path |
| 2 | UNDERSTAND | structured **expression_analysis**（必填才進 DECIDE） |
| 3 | CLASSIFY | expression.type ∈ registry |
| 4 | DECIDE | candidates + selection.policy + selected |
| 5 | VALIDATE | validation.* + mechanical.* |
| 6 | FINALITY | status + blocking_reasons |

## Constraint vs Selection

### Constraint Responsibility（可機械化為主）

```yaml
constraints:
  preserve:
    - speaker_intent
    - factual_meaning
    - proper_names
    - numbers
    - temporal_reference
  prohibit:
    - source_language_residue
    - invented_information
    - wrong_person
    - wrong_polarity
```

### Selection Responsibility（不可當 governance）

在符合 constraints 的候選中選表達。取決於：上下文、speaker、relationship、場景、時代、locale、角色個性、詞彙極性（例：ヤバい 正／負）、target audience。

必須記錄：`selection.policy`、`selected`、`rationale`（與 NVP invariant 5 同形）。

## Expression Analysis（中間層，Invariant）

禁止：

```text
source_text → translation
```

要求：

```text
source_text → expression_analysis → translation_decision → translation
```

最小方向：`surface`、`composite_type`、`structure[]`（span、type、role、handling）、`literal_meaning`、`pragmatic_meaning`、`register`、`speaker_intent`、`ambiguity[]`。

**Structured 範例**（陈小姐 → id-ID）：[`05-example-address-title-chen-xiaojie-id.yaml`](05-example-address-title-chen-xiaojie-id.yaml) — `陈` = proper_name／transliteration；`小姐` = address_title／locale_equivalent。

## Translation Decision Record（TDR）

正式 artifact 形狀見 Phase 1 `contracts/translation-decision.yaml`。用途：爭議時追溯整條 decision chain，對齊 Evidence Chain／EDR 思維。

## Validation（禁止單一分數）

```yaml
validation:
  semantic: pass|fail|review
  intent: pass|fail|review
  register: pass|fail|review
  terminology: pass|fail|review
  cultural: pass|fail|review
  format: pass|fail|review
  completeness: pass|fail|review
  locale_consistency:
    status: pass|fail|review
    reason: target_locale_residue   # e.g. id-ID dst contains Miss/Ms.
finality:
  status: accepted|needs_review|unresolved|rejected
  blocking_reasons: []
```

文字完全正確但文化等價不確定 → `needs_review` + `cultural_equivalence_uncertain` 類 blocking reason。

## Locale（反 dialect flattening）

```yaml
source_locale:
  language: ja
  region: JP
  dialect: kansai          # optional
  register: colloquial
target_locale:
  language: zh
  region: TW
  variant: traditional
  register: colloquial
```

不合理路徑：Kansai → Standard Japanese → Chinese（dialect flattening）。

## Evidence 類型（非 exhaustive）

- corpus、dictionary、contextual_usage（episode／scene）、llm_analysis（須標 model，不可當 closure authority）

Decision 可要求 `evidence_required` 依 expression type（Phase 1 registry）。
