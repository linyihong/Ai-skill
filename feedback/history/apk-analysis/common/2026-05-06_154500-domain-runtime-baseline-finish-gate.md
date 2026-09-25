Status: candidate

# Domain／執行環境基線 — 取數門檻
# Extracted — See [`workflow/apk-analysis/artifact-gates.md`](../../../../workflow/apk-analysis/artifact-gates.md) (Domain/Runtime Baseline)

**Date:** 2026-05-06  
**Category:** common  
**Status:** validated (method) — apply per project  

## Observation

Teams often stop after per-API request/response documentation and schema catalogs. Implementers then cannot reliably attach transports, derive opaque parameters (`l`-like session scalars), or choose pagination semantics without guessing.

## Lesson

Maintain a **project-level domain / runtime baseline** (separate from entity-level Domain Concepts in feature handoff) that records: environment/host family placeholders, TLS/proxy path, login/device dependency for list calls, lineage of opaque query fields, signing/gateway prerequisites (no secrets), pagination ground truth, rate limits. Cross-link rows to API Catalog entries and UI operation ids.

**Finish gate:** if the outcome includes SDK/client/replay/integration, baseline must exist or be an explicit skeleton with tracked open questions in the same work unit.

**Development gate:** if development is about to start for a live-facing SDK/client/app tool, the baseline must be more than a skeleton. It must answer the minimum runnable factors (endpoint/path family, route/service mapping or adapter strategy, session/bootstrap dependency, opaque parameter source/lifetime, signing/gateway prerequisites, response decrypt/unwrap boundary, pagination truth, error/session recovery, replay checklist). Missing factors block live-facing implementation unless explicitly scoped out. Skeleton baselines may only support offline parser, fixture, mock transport, or documentation work.

## Validation

Concrete baseline shape and checklist: [`../../DOCUMENTATION.md`](../../DOCUMENTATION.md) § *Domain／執行環境基線*. Skill entry updated in [`../../SKILL.md`](../../SKILL.md) Quick Start §7 and Default Workflow handoff.

## Applicability

Any authorized APK traffic analysis whose downstream consumes real HTTP or decrypted JSON outside the APK.

#### One-line Summary

既有 lesson 的結論維持於本檔原始內容；此 closure 補記不新增專案事實。

#### Evidence

既有工具輸出、觀察與專案證據已記於本 lesson 的原始段落；未新增或推論額外證據。

#### Generalized Lesson

將本條的具體情境視為候選通則；未在獨立情境重複驗證前，維持 candidate。

#### Agent Action

重用前先核對本條既有前提、限制與驗證方法；前提不符時重新取證。

#### Goal / Action / Validation

- Goal: 保留既有 lesson 的可驗證結論。
- Action: 依原始 Evidence 與 Trigger 重做相關檢查。
- Validation: 結果與原始結論一致才可重用。

#### Applies / Does Not Apply

- Applies: 本條既有 Trigger、前提與 Evidence 相符時。
- Does not apply: 跨專案、前提不明或缺少原始證據時；先重新驗證。

#### Promotion Target

- 尚未 promotion；保留為 candidate history，需有獨立 reuse evidence 才可提升。

#### Required Linked Updates

- Not applicable: 這是歷史 closure 修復；未新增可安全提升的 canonical guidance，category index 已存在。

#### One-line Summary

既有 lesson 的結論維持於本檔原始內容；此 closure 補記不新增專案事實。

#### Evidence

既有工具輸出、觀察與專案證據已記於本 lesson 的原始段落；未新增或推論額外證據。

#### Generalized Lesson

將本條的具體情境視為候選通則；未在獨立情境重複驗證前，維持 candidate。

#### Agent Action

重用前先核對本條既有前提、限制與驗證方法；前提不符時重新取證。

#### Goal / Action / Validation

- Goal: 保留既有 lesson 的可驗證結論。
- Action: 依原始 Evidence 與 Trigger 重做相關檢查。
- Validation: 結果與原始結論一致才可重用。

#### Applies / Does Not Apply

- Applies: 本條既有 Trigger、前提與 Evidence 相符時。
- Does not apply: 跨專案、前提不明或缺少原始證據時；先重新驗證。

#### Promotion Target

- 尚未 promotion；保留為 candidate history，需有獨立 reuse evidence 才可提升。

#### Required Linked Updates

- Not applicable: 這是歷史 closure 修復；未新增可安全提升的 canonical guidance，category index 已存在。
