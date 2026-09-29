# Phase 1 refinement — Locale-layered Failure Registry

Companion to [`_plan.md`](_plan.md)。**不成**新 workflow；只拆 registry 分層。

## Shape

```text
Locale Resolution
        ▼
failure-registry-binding.yaml          # binding / index（非 pattern SoT）
     /                    \
    ▼                      ▼
failure-patterns.yaml    locale/<locale>/failure-patterns.yaml
(core SoT)               (manifestations; manifests: F*)
        \                    /
         ▼                  ▼
           Constraint Space
```

| Layer | Owns |
| --- | --- |
| Binding | which registries load（`pattern_ids`） |
| Core | modality_loss、semantic_expansion、residue… |
| Locale | JA-F*／ID-F* as `manifests: <core_id>` |
| Knowledge | surface Candidate Space |

Former name `locale-failure-taxonomy.yaml` kept as deprecated pointer → [`failure-registry-binding.yaml`](../../workflow/translation/registry/failure-registry-binding.yaml).

## Moved

- JA-F01–F11 → [`registry/locale/ja-JP/failure-patterns.yaml`](../../workflow/translation/registry/locale/ja-JP/failure-patterns.yaml)
- ID-F01 → [`registry/locale/id-ID/failure-patterns.yaml`](../../workflow/translation/registry/locale/id-ID/failure-patterns.yaml)
- Core cleaned → [`registry/failure-patterns.yaml`](../../workflow/translation/registry/failure-patterns.yaml)

Adding ko／th／vi = new locale file + binding entry；**不**改 execution-flow 主鏈。
