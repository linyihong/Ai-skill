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

**Failure Learning（I13）**：Dogfood failure → classification → pattern **candidate** → Governance Review → Registry／Knowledge／Guards。禁止把案例永久堆進 Selection prompt。見 [`10`](10-failure-pattern-learning.md)。

**Semantic／Syntactic Realization（I14–I16）**：Meaning → Semantic Roles → Target Syntax／Naturalness。Preserve relations, not word order。見 [`11`](11-semantic-syntactic-realization.md)。

## Locale Resolution（P0，在 Analysis 之前）

**Locale Resolution ≠ Language Detection。**

```text
Consumer / job / locale pack
       │
       ▼
TranslationContext (source_locale + target_locale authoritative)
       │
       ▼
Content-Type / Context Resolution   ← content.type (subtitle|title|ui|…)
       │
       ▼
Title Structure Analysis (when content.type=title)
       │
       ▼
Expression Analysis
```

- `target_locale` 由 consumer 傳入，**不是**每段 LLM 猜出來的。
- **Invariant**：`target_locale is authoritative input, not an inferred translation decision.`
- **Target locale = Constraint，不是 Selection。**
- **content.type** 選 Selection Policy 家族；title ≠ ordinary subtitle sentence（I12；[`09`](09-title-content-type.md)）。

若 target 只存在 pipeline 最外層、未進入每段 Decision，稱謂等會被模型自行假設 target（典型 zh→en）。Dogfood：[`04`](04-dogfood-case-address-title-id-ID.md)。Freeze：[`06`](06-phase-0-freeze-invariants.md) I1。

**Invariant**：actor 輸入 = `TranslationContext` + segment + evidence；**禁止**僅 `{ src, dst }`。

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
| 1 | INGEST | **TranslationContext** 自 consumer／locale pack 綁定 |
| 1b | LOCALE RESOLVE | 綁定 authoritative locales（≠ language detection） |
| 2 | UNDERSTAND | **expression_analysis artifact**（producer 可替換） |
| 3 | CLASSIFY | type ∈ registry → **Candidate Space** |
| 4 | DECIDE | Constraints → feasible `candidates[]` → selection + **decision_basis** |
| 5 | VALIDATE | validation.* + mechanical.* + locale_consistency（review ≠ auto-fail） |
| 6 | FINALITY | I9 closure：context + validation + no blocking + selection |

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

必須記錄：`selection.policy`、`selected`、`rationale`、**`decision_basis`**（artifact 依據，非僅「比較自然」）。見 [`06`](06-phase-0-freeze-invariants.md) I8。

## Expression Analysis（artifact，Invariant）

禁止：

```text
source_text → translation
```

要求：

```text
source_text → expression_analysis (artifact) → translation_decision → translation
```

**Expression Analysis ≠ LLM step。** Contract 只要求欄位；producer 可為 registry matcher／dictionary／OCR+LLM／LLM。見 I2。

最小方向：`surface`、`composite_type`、`structure[]`、`pragmatic_meaning`、`register`、`ambiguity[]`（及條件欄位）。

**Structured 範例**：[`05`](05-example-address-title-chen-xiaojie-id.yaml)。

## Candidate Space vs Candidates（兩層）

```text
Registry / knowledge → Candidate Space
  → Constraints → Feasible Candidates (candidates[]: feasible, reason)
  → Selection Policy → Selection Actor → selected
```

見 I3／I7。`title_mapping` 種子 Candidate Space，不是 final answer。

## Translation Decision Record（TDR）

正式形狀見 Phase 1 `contracts/translation-decision.yaml`。Selection 含 `decision_basis`。

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
    status: pass|fail|review   # English title under id-ID → review（I5），非自動 fail
    reason: target_locale_residue
  # source_language_residue 與 target_locale_residue 分欄（I6）
finality:
  status: accepted|needs_review|unresolved|rejected
  blocking_reasons: []
```

**Finality（I9）**：`accepted` ⇔ required context present ∧ required validation pass／waivered ∧ no unresolved blocking ∧ selection exists。

文字正確但文化不確定 → `needs_review`。`id-ID` + `Miss`/`Ms.` → `review` + rationale／waiver 才能 accepted。

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
