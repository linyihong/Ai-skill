# Translation knowledge（skeleton）

Locale- and culture-specific **Candidate Space seeds** for `workflow/translation/`.  
Not translation truth. Selection still required（I10）。  
Dogfood errors → [`failure-patterns`](../../workflow/translation/registry/failure-patterns.yaml)，不是 prompt 例句堆（I13）。

| Path | Role |
| --- | --- |
| [`locale/title-mapping.yaml`](locale/title-mapping.yaml) | Address-title seeds（含 ms／id／en／ja） |
| [`locale/ms-MY/address-title.yaml`](locale/ms-MY/address-title.yaml) | Native ms-MY honorifics（Cik／Puan…） |
| [`locale/ja-JP/name-realization.yaml`](locale/ja-JP/name-realization.yaml) | Foreign surname realization（陈 only） |
| [`locale/temporal-reference.yaml`](locale/temporal-reference.yaml) | Temporal families + clock constraint pointers |

Do **not** grow full surname／honorific／gloss dictionaries without dogfood → failure_pattern evidence.
