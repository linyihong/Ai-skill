# Translation Decision Workflow

`workflow/translation/` 是 **cross-cutting governed translation decision** capability：  
Meaning → Intent → Register → Cultural Expression → Target Expression，不是「原文 → LLM → 完成」。

> **狀態**：Phase 1–3 + title／I12 + **Failure Pattern learning（I13）**。  
> **沒有** `route.workflow.translation`。YAML **不**投影。Prompt = Selection adapter only。  
> Plan：[`2026-09-22-1000-translation-decision-workflow`](../../plans/active/2026-09-22-1000-translation-decision-workflow/_plan.md)。  
> Title／content_type：[`09`](../../plans/active/2026-09-22-1000-translation-decision-workflow/09-title-content-type.md)。  
> Failure learning：[`10`](../../plans/active/2026-09-22-1000-translation-decision-workflow/10-failure-pattern-learning.md)。

## 一句話責任邊界

| 元件 | 問句 |
| --- | --- |
| Context | Where am I translating?（含 **content.type**） |
| Title Structure | Is this a title／truncated／idiom hook／part marker? |
| Analysis | What is this expression? |
| Realization／Knowledge | How does it look in the target locale? |
| Registry | What possibilities／failure guards exist? |
| Constraints／Guards | What is not allowed to be wrong? |
| Candidates | What is feasible?（含 decision_class） |
| Policy | What should we optimize? |
| Selection Actor | Select among **feasible** candidates（prompt ≠ rule store） |
| Verifier | Is the decision defensible? |
| Finality | Can this be closed? |
| Failure Learning | New case, or evidence for a known pattern?（I13） |

**Translation Strategy ≠ Target-Locale Realization**（I11）。  
**translation ≠ adaptation ≠ marketing_generation**（I12）。  
**prompt case list ≠ failure_pattern**（I13）。

## Mechanical invariants

| ID | 規則 |
| --- | --- |
| I1–I11 | Context／Analysis／Candidate／Selection／residue／realization |
| I12 | `content.type`；title 先結構分析；禁止 unsupported marketing |
| I13 | Dogfood → pattern **candidate** → Governance Review → active guard；LLM 不可自提升 |

## 三層（取代 prompt if-blocks）

| Layer | 位置 |
| --- | --- |
| Registry（抽象） | [`registry/expression-types.yaml`](registry/expression-types.yaml)、[`registry/failure-patterns.yaml`](registry/failure-patterns.yaml) |
| Knowledge（locale） | [`knowledge/translation/`](../../knowledge/translation/README.md) |
| Guards（validation） | [`registry/validation-rules.yaml`](registry/validation-rules.yaml) |

## 何時讀哪個檔

| 認知階段 | 檔案 |
| --- | --- |
| Lifecycle | [`execution-flow.md`](execution-flow.md) |
| Context | [`contracts/translation-context.yaml`](contracts/translation-context.yaml) |
| Analysis | [`contracts/expression-analysis.yaml`](contracts/expression-analysis.yaml) |
| Decision | [`contracts/translation-decision.yaml`](contracts/translation-decision.yaml) |
| Failure Pattern | [`contracts/failure-pattern.yaml`](contracts/failure-pattern.yaml) |
| Validate／Close | [`contracts/validation.yaml`](contracts/validation.yaml)、[`contracts/finality.yaml`](contracts/finality.yaml) |
| Types／strategies／guards | [`registry/`](registry/) |
| Subtitle → NVP | [`adapters/subtitle.yaml`](adapters/subtitle.yaml) |
| Fixtures | [`examples/`](examples/) |

## 核心原則

1. Locale Resolution ≠ Language Detection。
2. Content-Type Resolution 在 Analysis 之前。
3. Constraints／failure guards 定義可行集；Selection 是明示 policy。
4. Dogfood 錯誤抽成 pattern，不堆進 Selection prompt。
5. invented_information／idiom／clock／residue **分欄**。
6. NVP content vs timing／layout 分界不變。

## 明確不做

runtime route、完整姓氏庫、title-translation-workflow、LLM 自動寫 active guards、locale if-case prompt 膨脹。
