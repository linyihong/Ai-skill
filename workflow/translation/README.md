# Translation Decision Workflow

`workflow/translation/` 是 **cross-cutting governed translation decision** capability：  
Meaning → Semantic Roles → Target Syntax／Naturalness → Target Expression，不是「原文 → LLM → 完成」。

> **狀態**：Phase 1–3 + I12 title + I13 failure + **I14–I16 semantic／syntax**（[`11`](../../plans/active/2026-09-22-1000-translation-decision-workflow/11-semantic-syntactic-realization.md)）。  
> Prompt = Selection adapter only。未註冊 route。

## 一句話責任邊界

| 元件 | 問句 |
| --- | --- |
| Context | Where am I translating?（content.type） |
| Analysis | What is this expression + **semantic structure**? |
| Realization／Knowledge | How does it look in the target locale? |
| Registry／Guards | What possibilities／failure patterns exist? |
| Constraints | What must not be wrong?（roles／polarity／numbers／identity） |
| Selection | Among feasible — which surface／order／honorific? |
| Target Realization | Semantic → Syntax → Naturalness under I14–I16 |
| Verifier | semantic≠syntactic≠grammatical≠naturalness |
| Finality／Learning | Closed? New pattern or known evidence? |

## Mechanical invariants

| ID | 規則 |
| --- | --- |
| I1–I13 | Context／Analysis／Candidate／Selection／residue／title／failure |
| I14 | Preserve **semantic relations**, not source surface form |
| I15 | Target reorder／omit／restructure OK **only if** relations preserved |
| I16 | Naturalness MUST NOT alter roles／entities／temporal／polarity／intent |

## Failure taxonomy（摘要）

F1 role loss · F2 grammatical relation · F3 syntax distortion · F4 unnatural · F5 social address · F6 idiom · F7 locale residue · F8 name script · F9 invented — 見 [`registry/failure-patterns.yaml`](registry/failure-patterns.yaml)。

## 何時讀哪個檔

| 認知階段 | 檔案 |
| --- | --- |
| Lifecycle | [`execution-flow.md`](execution-flow.md) |
| Context | [`contracts/translation-context.yaml`](contracts/translation-context.yaml) |
| Analysis | [`contracts/expression-analysis.yaml`](contracts/expression-analysis.yaml) |
| Decision | [`contracts/translation-decision.yaml`](contracts/translation-decision.yaml) |
| Failure Pattern | [`contracts/failure-pattern.yaml`](contracts/failure-pattern.yaml) |
| Validate | [`contracts/validation.yaml`](contracts/validation.yaml) |
| Types／guards | [`registry/`](registry/) |
| Fixtures | [`examples/`](examples/)（含 [`social-address-laozhang`](examples/social-address-laozhang.yaml)） |

## 明確不做

runtime route、完整姓氏庫、LLM 自動 active guards、locale if-case prompt 膨脹、要求 source word order == target。
