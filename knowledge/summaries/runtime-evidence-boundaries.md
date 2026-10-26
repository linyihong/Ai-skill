# runtime-evidence-boundaries

| 欄位 | 值 |
| --- | --- |
| Atom ID | `intelligence.analysis.runtime-evidence-boundaries` |
| Source path | [候選智慧](../../intelligence/engineering/analytical-reasoning/heuristics/runtime-evidence-boundaries.md) |
| Lifecycle | candidate |
| Summary | runtime evidence boundary：hook 命中或空回傳只覆蓋局部邊界，先找首個未驗證契約再判定 driver 或資料缺失。 |
| When to read | IL2CPP、Animator.Play、hook lifecycle、negative evidence 或 local success 與視覺不一致。 |
| Do not use for | 不把 candidate 抽取當作獨立重用實測；不取代 analysis 操作程序。 |
| Validation signal | 判斷有前置條件、可推翻證據與反例；來源 lesson／目標／索引可雙向追溯。 |
| Last checked | 2026-10-26 |
