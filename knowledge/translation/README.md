# Translation knowledge（skeleton）

Locale- and culture-specific **Candidate Space seeds** for `workflow/translation/`.  
Not translation truth. Selection still required（I10）。  
Dogfood errors → [`failure-patterns`](../../workflow/translation/registry/failure-patterns.yaml)，不是 prompt 例句堆（I13）。

| Path | Role |
| --- | --- |
| [`locale/title-mapping.yaml`](locale/title-mapping.yaml) | Address-title seeds（含 ms／id／en／ja／ko／vi／th／ar／es／pt／tr） |
| [`locale/ms-MY/address-title.yaml`](locale/ms-MY/address-title.yaml) | Native ms-MY honorifics（Cik／Puan…） |
| [`locale/ja-JP/name-realization.yaml`](locale/ja-JP/name-realization.yaml) | Foreign surname realization（陈 only） |
| [`locale/ja-JP/kinship.yaml`](locale/ja-JP/kinship.yaml) | 姐夫／姐 Candidate Space（JA-F05） |
| [`locale/temporal-reference.yaml`](locale/temporal-reference.yaml) | Temporal families + clock constraint pointers |

Do **not** grow full surname／honorific／gloss dictionaries without dogfood → failure_pattern evidence.
