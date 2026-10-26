# representation-contract-parity

| 欄位 | 值 |
| --- | --- |
| Atom ID | `intelligence.analysis.representation-contract-parity` |
| Source path | [候選智慧](../../intelligence/engineering/analytical-reasoning/heuristics/representation-contract-parity.md) |
| Lifecycle | candidate |
| Summary | representation contract parity：fixture 展示子集與完整 runtime 結構依契約轉換，合法 sparse 不補滿，未知缺值拒絕寫入。 |
| When to read | fixture shape、sparse、visible subset、padding、runtime representation mismatch。 |
| Do not use for | 不把 candidate 抽取當作獨立重用實測；不取代 analysis 操作程序。 |
| Validation signal | 判斷有前置條件、可推翻證據與反例；來源 lesson／目標／索引可雙向追溯。 |
| Last checked | 2026-10-26 |
