# Translation Decision Workflow

`workflow/translation/` 是 **cross-cutting governed translation decision** capability：  
Meaning → Intent → Register → Cultural Expression → Target Expression，不是「原文 → LLM → 完成」。

> **狀態**：Phase 1 contract／registry（doc-only）。**沒有** `route.workflow.translation`。  
> YAML **不**投影（`runtime_projection.enabled: false`）。不接 provider／prompt／model routing。  
> Plan：[`2026-09-22-1000-translation-decision-workflow`](../../plans/active/2026-09-22-1000-translation-decision-workflow/_plan.md)。  
> Phase 0 freeze：[`06-phase-0-freeze-invariants.md`](../../plans/active/2026-09-22-1000-translation-decision-workflow/06-phase-0-freeze-invariants.md)。

## 一句話責任邊界

| 元件 | 問句 |
| --- | --- |
| Context | Where am I translating? |
| Analysis | What is this expression? |
| Registry | What possibilities exist? |
| Constraints | What is not allowed to be wrong? |
| Candidates | What is feasible? |
| Policy | What should we optimize? |
| Selection Actor（含 LLM） | Select among **feasible** candidates |
| Verifier | Is the decision defensible? |
| Finality | Can this be closed? |

## Mechanical invariants（Phase 1）

| ID | 規則 |
| --- | --- |
| I1 | `TranslationContext` 必須存在 |
| I2 | `target_locale` 是 **authoritative input**（Constraint），不是推論出的 decision |
| I3 | **禁止** src／dst-only translation decision path |
| I4 | Expression Analysis 是 **artifact**；不限定 producer |
| I5 | **Candidate Space ≠ Feasible Candidates** |
| I6 | Selection 必須有 **explicit `selection.policy`** |
| I7 | LLM／模型只當 Selection Actor；**不**擁有 Constraint Responsibility |
| I8 | `source_language_residue` ≠ `target_locale_residue`（分欄） |
| I9 | `finality.accepted` ⇔ context + validation（或 waiver）+ selection + no unresolved blocking |
| I10 | `title_mapping`／registry seeds **只**種子 Candidate Space；**不得**當 final answer |

## 何時讀哪個檔

| 認知階段 | 檔案 |
| --- | --- |
| Lifecycle | [`execution-flow.md`](execution-flow.md) |
| Context | [`contracts/translation-context.yaml`](contracts/translation-context.yaml)、[`contracts/source.yaml`](contracts/source.yaml) |
| Analysis | [`contracts/expression-analysis.yaml`](contracts/expression-analysis.yaml) |
| Decision | [`contracts/translation-decision.yaml`](contracts/translation-decision.yaml) |
| Validate／Close | [`contracts/validation.yaml`](contracts/validation.yaml)、[`contracts/finality.yaml`](contracts/finality.yaml) |
| Types／strategies | [`registry/`](registry/) |
| Walkthrough | [`examples/`](examples/) — P0：[`address-title-chen-xiaojie-id.yaml`](examples/address-title-chen-xiaojie-id.yaml) |

## 核心原則

1. Locale Resolution ≠ Language Detection；locales 由 consumer／job／locale pack 傳入。
2. Constraints 定義可行集；Selection 是明示 policy。
3. 禁止單一 translation quality／confidence score。
4. 換模型／API 不應改變本目錄契約形狀。
5. NVP 字幕：本 workflow 管 **content** decision；timing／layout 仍屬 [`narrative-video-production`](../narrative-video-production/captions-and-locales.md)（Phase 2 adapter）。

## 明確不做（Phase 1）

runtime route、provider／prompt、translation memory、glossary engine、auto terminology、auto-correct、test runner。
