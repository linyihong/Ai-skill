# 11 — Phase 3 Dogfood + Failure Learning

← [`_plan.md`](_plan.md)

## 目標

在可重現合成資料上跑通：**多類負向挡住 + baseline candidate（rembg）IDR + 量測證據**。
不把單次 rembg 成功升級為永久 guard。

## 指令（歷史 Python gate 已刪；現行組裝 Go methods）

```text
# 對五類負向各跑：
ai-skill image-processing verify …/neg.<name>.png --profile neg.<name> --expect-block
# rembg dogfood 屬 consumer 本機 adapter；證據已落在 fixtures/dogfood/
```

見 [`tools/README.md`](../../../workflow/image-processing/tools/README.md)。

## 產物

| 產物 | 路徑 |
| --- | --- |
| Evidence | [`evidence/2026-10-14-phase3-synthetic-dogfood.md`](evidence/2026-10-14-phase3-synthetic-dogfood.md) |
| Gate JSON | `workflow/image-processing/fixtures/reports/phase3-gate-latest.json` |
| IDR | `workflow/image-processing/fixtures/dogfood/reports/*.idr.json` |
| Category sources | `fixtures/dogfood/sources/dogfood.*` |

## Acceptance

- [x] 五類負向 fixture 可執行且挡住
- [x] ≥1 rembg dogfood IDR（合成源）
- [x] 品質／耗時證據入 `evidence/`
- [x] 明確禁止無覆核升 guard
