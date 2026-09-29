# Phase 1 refinement — Locale-layered Failure Registry

Companion to [`_plan.md`](_plan.md)。**不成**新 workflow；只拆 registry 分層。

## Shape

```text
Translation Workflow Core
        │
        ▼
Constraint Space =
  registry/failure-patterns.yaml                 # cross-locale concepts
∪ registry/locale/<target_locale>/failure-patterns.yaml  # manifestations
```

| Layer | Owns |
| --- | --- |
| Core | modality_loss、semantic_expansion、residue、incomplete_source… |
| Locale | JA-F*／ID-F* as `manifests: <core_id>` + applies_when／constraints |
| Knowledge | 義兄さん／チャオ・ウェイ／Nona — Candidate Space only |

## Moved

- JA-F01–F11 → [`registry/locale/ja-JP/failure-patterns.yaml`](../../workflow/translation/registry/locale/ja-JP/failure-patterns.yaml)
- ID-F01 → [`registry/locale/id-ID/failure-patterns.yaml`](../../workflow/translation/registry/locale/id-ID/failure-patterns.yaml)
- Core cleaned → [`registry/failure-patterns.yaml`](../../workflow/translation/registry/failure-patterns.yaml)

Adding ko／th／vi = new locale file；**不**改 execution-flow 主鏈。
