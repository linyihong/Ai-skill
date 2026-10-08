# Translation Decision Workflow

`workflow/translation/` 是 **cross-cutting governed translation decision** capability：  
Meaning → Semantic Roles → Target Syntax／Naturalness → Target Expression，不是「原文 → LLM → 完成」。

> **狀態**：Phase 1–3 + I12–I23（ep8 Finality／Reference [`13`](../../plans/active/2026-09-22-1000-translation-decision-workflow/13-ep8-contract-revision.md)）。  
> Prompt = Selection adapter only。未註冊 route。不成固定譯詞庫。

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
| I14b | Target **Lexical Realization** — script-legal ≠ target lexeme（F18） |
| I15 | Target reorder／omit／restructure OK **only if** relations preserved |
| I16 | Naturalness MUST NOT alter roles／entities／temporal／polarity／intent |
| I17 | False cognate／kanji decomposition flagged；no sole zh→ja kanji gloss |
| I18 | semantic_expansion ≠ marketing invented_information |
| I19 | Preserve participants／event／modality／polarity |
| I20 | Meaning-correct + register drift → register fail／review |
| I21 | incomplete_source_must_not_be_completed → **BLOCK** |
| I22 | PASS／REVIEW／BLOCK — uncertainty alone ≠ REVIEW |
| I23 | Kinship／address via reference_resolution — no sole fixed gloss |

跨域 pattern：[`traceable-evidence`](../cross-cutting/traceable-evidence/README.md)（編號用 `TE` 前綴，與上表 I1–I23 無關）。`translation-decision` 的 candidates + infeasible reason 實作 TE1；selection policy／rationale／decision_basis 與 finality `review.reason[]` 實作 TE4；I22 實作 TE5。

## Failure registry（摘要）

| 層 | 路徑 | 角色 |
| --- | --- | --- |
| Binding | [`failure-registry-binding.yaml`](registry/failure-registry-binding.yaml) | Locale Resolution 後載入哪些 registry |
| Core SoT | [`failure-patterns.yaml`](registry/failure-patterns.yaml) | 跨語言 failure concepts |
| Locale SoT | [`locale/<locale>/`](registry/locale/README.md) | manifestations（`manifests: F*`） |

Binding ≠ pattern definition。`pattern_ids` 是 registry references，不是第二套 taxonomy。

## 何時讀哪個檔

| 認知階段 | 檔案 |
| --- | --- |
| Lifecycle | [`execution-flow.md`](execution-flow.md) |
| Context | [`contracts/translation-context.yaml`](contracts/translation-context.yaml) |
| Analysis | [`contracts/expression-analysis.yaml`](contracts/expression-analysis.yaml) |
| Decision | [`contracts/translation-decision.yaml`](contracts/translation-decision.yaml) |
| Failure Pattern | [`contracts/failure-pattern.yaml`](contracts/failure-pattern.yaml) |
| Reference | [`contracts/reference-resolution.yaml`](contracts/reference-resolution.yaml) |
| Finality | [`contracts/finality.yaml`](contracts/finality.yaml) — PASS／REVIEW／BLOCK |
| Validate | [`contracts/validation.yaml`](contracts/validation.yaml) |
| Dogfood acceptance projection | [`adapters/dogfood-acceptance.md`](adapters/dogfood-acceptance.md) — identity／semantic／lexical／multi-dimensional evidence；非新 workflow |
| Types／guards | [`registry/`](registry/)（含 [`failure-registry-binding`](registry/failure-registry-binding.yaml)） |
| Fixtures | [`examples/`](examples/)（含 [`ep8-ja-walkthrough`](examples/ep8-ja-walkthrough.yaml)） |

## 明確不做

runtime route、完整姓氏庫、LLM 自動 active guards、locale if-case prompt 膨脹、要求 source word order == target。
