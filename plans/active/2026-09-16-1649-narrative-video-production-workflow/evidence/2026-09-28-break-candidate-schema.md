# Observation — Phase 1 break-candidate schema

**Run ID**：2026-09-28-break-candidate-schema  
**Kind**：contract（schema freeze；尚無本 repo 程式）  
**Extends**：[`2026-09-28-break-candidate-system.md`](2026-09-28-break-candidate-system.md)

## 決定

`balance_score` 描述左右平衡；`score` 才是機械選擇成本。`hard_violation` 非空就離開可行集，不是扣分後仍可能中選。Phase 1 不設 `source`。`semantic_boundary` 用 `unknown | boundary | non_boundary`，Phase 1 保持 `unknown`。`phrase_integrity` 只是 advisory。`best_cut` 是 fallback selector。

## 尚未做

本 repo 沒有 `BreakPolicy` 程式。Python dataclass 由消費這個 schema 的實作填。Phase 2 才接 LLM selection 與 `break_evidence`。
