# Translation Decision Workflow

`workflow/translation/` 是 **cross-cutting governed translation decision** capability：  
Meaning → Intent → Register → Cultural Expression → Target Expression，不是「原文 → LLM → 完成」。

> **狀態**：Phase 1–3 + **content_type／title**（I12）補強。Phase 2 subtitle adapter。  
> **沒有** `route.workflow.translation`。YAML **不**投影。不接 provider／prompt。  
> Plan：[`2026-09-22-1000-translation-decision-workflow`](../../plans/active/2026-09-22-1000-translation-decision-workflow/_plan.md)。  
> Walkthrough PASS：[`08`](../../plans/active/2026-09-22-1000-translation-decision-workflow/08-static-walkthrough-pass.md)。  
> Realization：[`07`](../../plans/active/2026-09-22-1000-translation-decision-workflow/07-target-locale-realization.md)。  
> Title／content_type：[`09`](../../plans/active/2026-09-22-1000-translation-decision-workflow/09-title-content-type.md)。

## 一句話責任邊界

| 元件 | 問句 |
| --- | --- |
| Context | Where am I translating?（含 **content.type**） |
| Title Structure | Is this a title／truncated／idiom hook／part marker? |
| Analysis | What is this expression? |
| Realization | How does it look in the target locale? |
| Registry | What possibilities exist? |
| Constraints | What is not allowed to be wrong? |
| Candidates | What is feasible?（含 decision_class） |
| Policy | What should we optimize? |
| Selection Actor | Select among **feasible** candidates |
| Verifier | Is the decision defensible? |
| Finality | Can this be closed? |

**Translation Strategy ≠ Target-Locale Realization**（I11）。  
**translation ≠ adaptation ≠ marketing_generation**（I12）。

## Mechanical invariants

| ID | 規則 |
| --- | --- |
| I1–I11 | 見既有表（Context／Analysis／Candidate／Selection／residue／realization） |
| I12 | `content.type` 選 policy；title 先結構分析；**禁止** unsupported marketing／invented hashtags |

## 何時讀哪個檔

| 認知階段 | 檔案 |
| --- | --- |
| Lifecycle | [`execution-flow.md`](execution-flow.md) |
| Context | [`contracts/translation-context.yaml`](contracts/translation-context.yaml) |
| Analysis | [`contracts/expression-analysis.yaml`](contracts/expression-analysis.yaml) |
| Decision | [`contracts/translation-decision.yaml`](contracts/translation-decision.yaml) |
| Validate／Close | [`contracts/validation.yaml`](contracts/validation.yaml)、[`contracts/finality.yaml`](contracts/finality.yaml) |
| Types／strategies／realization | [`registry/`](registry/) |
| Subtitle → NVP | [`adapters/subtitle.yaml`](adapters/subtitle.yaml) |
| Title P0 | [`examples/title-kongjie-yiriqianli.yaml`](examples/title-kongjie-yiriqianli.yaml) |
| Name seeds | [`knowledge/translation/`](../../knowledge/translation/README.md) |

## 核心原則

1. Locale Resolution ≠ Language Detection。
2. Content-Type Resolution 在 Analysis 之前（title ≠ subtitle）。
3. Constraints 定義可行集；Selection 是明示 policy。
4. `preferred_script` 約束 Candidate Space，禁止非片假一律 FAIL。
5. invented_information 與 idiom semantic uncertainty **分欄**。
6. NVP 字幕 content vs timing／layout 分界不變。

## 明確不做

runtime route、provider／prompt、完整姓氏庫、title-translation-workflow、test runner。
