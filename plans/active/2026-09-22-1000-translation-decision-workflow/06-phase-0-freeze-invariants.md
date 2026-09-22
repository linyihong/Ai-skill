# Phase 0 freeze — boundary invariants（2026-09-22 review）

Companion to [`_plan.md`](_plan.md)。Stakeholder review：**Phase 0 freeze，進 Phase 1**。  
本檔只釘 **contract 邊界**；不重寫 plan、不擴 Phase 1 scope。

## Verdict

| 項 | 裁決 |
| --- | --- |
| Phase 0 | **freeze** |
| Phase 1 | **開跑**（doc-only contracts + registry；不增平台能力） |
| Plan 重做 | **不需要** |

## Phase 1 shape（凍結）

```text
TranslationContext
       │
       ▼
┌──────────────────┐
│ Locale Resolution│  ← bind authoritative locales; ≠ Language Detection
└────────┬─────────┘
         ▼
Expression Analysis   ← artifact（producer 可替換）
         ▼
Expression Registry
         ▼
Candidate Space       ← registry / knowledge 種子
         ▼
Constraints
         ▼
Feasible Candidates   ← candidates[] with feasible / reason
         ▼
Selection Policy
         ▼
Selection Actor       ← LLM 等；只在可行集內選
         ▼
Translation Decision
    ┌────┴────┐
    ▼         ▼
Independent  Mechanical
Review       Gate
    └────┬────┘
         ▼
      Finality
```

## Responsibility boundary（一句話）

| 元件 | 問句 |
| --- | --- |
| Context | Where am I translating? |
| Analysis | What is this expression? |
| Registry | What possibilities exist? |
| Constraints | What is not allowed to be wrong? |
| Candidates | What is feasible? |
| Policy | What should we optimize? |
| LLM / Selection Actor | Select among feasible candidates |
| Verifier | Is the decision defensible? |
| Runtime | Did the contract actually hold? |
| Finality | Can this be closed? |

## Frozen invariants（Phase 1 contracts 必須寫清）

### I1 — Locale Resolution ≠ Language Detection

- `source_locale` / `target_locale` 由 **consumer／job／locale pack** 傳入 `TranslationContext`。
- **禁止**每段 LLM 猜 target language／locale。
- **Invariant**：`target_locale is authoritative input, not an inferred translation decision.`
- **Target locale = Constraint，不是 Selection。**

NVP 路徑：`locale pack → TranslationContext → Expression Analysis`。

### I2 — Expression Analysis is an artifact, not an LLM step

Contract 只要求欄位形狀（`composite_type`、`structure`、`pragmatic_meaning`、`register`、`ambiguity`…）。  
Producer 可為：registry matcher、dictionary、OCR+LLM、純 LLM——**不得**把「Expression Analysis = AI analysis」寫進 SoT。

### I3 — Candidate Space ≠ Feasible Candidates ≠ Selected

```text
Knowledge / Registry → Candidate Space
  → Constraints → Feasible Candidates (candidates[] + feasible + reason)
  → Selection Policy → Selection Actor → selected
```

`title_mapping` 等只種子 Candidate Space；**Candidate Space ≠ Final Answer**（registry invariant）。

### I4 — 陈小姐 id-ID = P0 regression fixture

[`05-example-address-title-chen-xiaojie-id.yaml`](05-example-address-title-chen-xiaojie-id.yaml) 定位為 **workflow invariant regression fixture**（locale_aware title + English title residue）。  
Phase 1：doc-only `regression_fixtures` 即可；**不**建 test runner。未來可路徑：`tests/regression/locale/`（不擋 Phase 1）。

### I5 — target_locale_residue → `review`，不是自動 `fail` / incorrect

`id-ID` + English title token（`Miss`／`Ms.`）→ `locale_consistency: review` + `target_locale_residue`。  
可能合法：borrowed expression、international title、character voice、intentionally mixed language。  
`accepted` 需 **explicit rationale 或 waiver**；LLM「也可以吧」不可關閉。

### I6 — `source_language_residue` ≠ `target_locale_residue`

| 例 | 分類 |
| --- | --- |
| ja→id，dst 含「それは…」 | `source_language_residue` |
| zh→id，dst = `Miss Chen` | `target_locale_residue` / locale_consistency |

兩欄分欄，**禁止合併**。

### I7 — `title_mapping` 不是翻譯字典真理

Phase 3 種子表不得被當 translation truth；否則退回 dictionary pipeline。  
Invariant：**Candidate Space ≠ Final Answer**。

### I8 — Selection 必填 `decision_basis`

除 `policy` / `selected` / `rationale` 外，TDR 加：

```yaml
selection:
  decision_basis:
    - target_locale
    - expression_analysis
    - scene_context
    - candidate_registry   # illustrative
```

`rationale` 可含自然語；`decision_basis` 必須指向 **artifact／契約依據**（Evidence Chain）。

### I9 — Finality mechanical closure

`accepted` 僅當：

1. required context present  
2. required validation pass（或已 waivered per policy）  
3. no unresolved blocking reason  
4. selection exists  

缺 context、或任何 **blocking** validation 未 resolved → 不得 `accepted`。

### I10 — Phase 1 停在資料契約；不加平台

Phase 1 **只做**：TranslationContext、Expression Analysis、TDR、Expression／Strategy Registry、Validation（含 locale）、examples、P0 regression fixture、README、execution-flow。

**明確不加**：translation memory、glossary engine、automatic terminology extraction、LLM prompt、model routing、translation／confidence score、runtime route、automatic correction。

## Phase 1 SoT 寫作檢查（三處特別小心）

寫 contracts 時若偏離下列任一條 → 先停並改契約，勿堆 AI-specific 欄位：

1. Locale Resolution 被寫成 per-segment language detection  
2. Expression Analysis 被寫成「必須 LLM」  
3. Candidate Space 與 `candidates[]` 混成一層、或 `title_mapping` 當 final answer  
