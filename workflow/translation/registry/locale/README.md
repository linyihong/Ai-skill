# Locale failure-pattern surfaces

After Locale Resolution, Constraint Space =

```text
registry/failure-patterns.yaml          # cross-locale CORE
  ∪
registry/locale/<target_locale>/failure-patterns.yaml   # manifestations
```

| Locale | File | Status |
| --- | --- | --- |
| (core) | [`../failure-patterns.yaml`](../failure-patterns.yaml) | active |
| ja-JP | [`ja-JP/failure-patterns.yaml`](ja-JP/failure-patterns.yaml) | active（JA-F01–F11） |
| id-ID | [`id-ID/failure-patterns.yaml`](id-ID/failure-patterns.yaml) | minimal（ID-F01） |
| en | [`en/failure-patterns.yaml`](en/failure-patterns.yaml) | placeholder |
| ko-KR | [`ko-KR/failure-patterns.yaml`](ko-KR/failure-patterns.yaml) | placeholder |

Index：[`../locale-failure-taxonomy.yaml`](../locale-failure-taxonomy.yaml)。  
**禁止**把 sole surface gloss 寫進 pattern body；表面形在 `knowledge/translation/`。
