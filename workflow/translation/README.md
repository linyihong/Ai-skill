# Translation Decision Workflow

`workflow/translation/` 是 **cross-cutting governed translation decision** capability：  
Meaning → Intent → Register → Cultural Expression → Target Expression，不是「原文 → LLM → 完成」。

> **狀態**：Phase 1 contracts PASS（靜態走讀 A+B）+ **Phase 2 subtitle adapter**（doc-only）。  
> **沒有** `route.workflow.translation`。YAML **不**投影。不接 provider／prompt。  
> Plan：[`2026-09-22-1000-translation-decision-workflow`](../../plans/active/2026-09-22-1000-translation-decision-workflow/_plan.md)。  
> Walkthrough PASS：[`08`](../../plans/active/2026-09-22-1000-translation-decision-workflow/08-static-walkthrough-pass.md)。  
> Phase 0 freeze：[`06`](../../plans/active/2026-09-22-1000-translation-decision-workflow/06-phase-0-freeze-invariants.md)。  
> Realization：[`07`](../../plans/active/2026-09-22-1000-translation-decision-workflow/07-target-locale-realization.md)。

## 一句話責任邊界

| 元件 | 問句 |
| --- | --- |
| Context | Where am I translating? |
| Analysis | What is this expression? |
| Realization | How does it look in the target locale (script／phonetics／form)? |
| Registry | What possibilities exist? |
| Constraints | What is not allowed to be wrong? |
| Candidates | What is feasible? |
| Policy | What should we optimize? |
| Selection Actor（含 LLM） | Select among **feasible** candidates |
| Verifier | Is the decision defensible? |
| Finality | Can this be closed? |

**Translation Strategy ≠ Target-Locale Realization**（I11）。例：`Chenさん` = title OK，name 未完成 katakana realization。

## Mechanical invariants（Phase 1）

| ID | 規則 |
| --- | --- |
| I1 | `TranslationContext` 必須存在 |
| I2 | `target_locale` 是 **authoritative input**（Constraint） |
| I3 | **禁止** src／dst-only path |
| I4 | Expression Analysis 是 **artifact** |
| I5 | **Candidate Space ≠ Feasible Candidates** |
| I6 | Selection 必須有 **explicit `selection.policy`** |
| I7 | LLM 只當 Selection Actor |
| I8 | `source_language_residue` ≠ `target_locale_residue` |
| I9 | `finality.accepted` ⇔ context + validation／waiver + selection + no blocking |
| I10 | mapping／realization seeds **只**種子 Candidate Space |
| I11 | Realization ≠ translation strategy；incomplete name realization → **review**（非硬 FAIL 全部非片假） |

## 何時讀哪個檔

| 認知階段 | 檔案 |
| --- | --- |
| Lifecycle | [`execution-flow.md`](execution-flow.md) |
| Context | [`contracts/translation-context.yaml`](contracts/translation-context.yaml) |
| Analysis | [`contracts/expression-analysis.yaml`](contracts/expression-analysis.yaml) |
| Decision | [`contracts/translation-decision.yaml`](contracts/translation-decision.yaml) |
| Validate／Close | [`contracts/validation.yaml`](contracts/validation.yaml)、[`contracts/finality.yaml`](contracts/finality.yaml) |
| Types／strategies／realization | [`registry/`](registry/)（含 [`realization-strategies.yaml`](registry/realization-strategies.yaml)） |
| Subtitle → NVP | [`adapters/subtitle.yaml`](adapters/subtitle.yaml) — **content_gate only** |
| Name seeds | [`knowledge/translation/locale/ja-JP/name-realization.yaml`](../../knowledge/translation/locale/ja-JP/name-realization.yaml) |
| Walkthrough | [`examples/`](examples/) — id + ja P0；PASS 見 plan [`08`](../../plans/active/2026-09-22-1000-translation-decision-workflow/08-static-walkthrough-pass.md) |

## 核心原則

1. Locale Resolution ≠ Language Detection。
2. Constraints 定義可行集；Selection 是明示 policy。
3. `preferred_script` 約束 Candidate Space，**禁止**「非片假＝FAIL」。
4. 換模型不應改變契約形狀。
5. NVP 字幕 content vs timing／layout 分界不變。

## 明確不做（Phase 1）

runtime route、provider／prompt、完整姓氏庫、glossary engine、test runner。
