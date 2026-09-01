# Artifact Gates

Eligibility **只讀 record 欄位**。對照表在
[`records/artifact-gates.yaml`](records/artifact-gates.yaml)（與 Phase 2 同形）。

不要在本檔列出 mutation class。下游改動後：

```text
mutation_event → identity-acceptance.yaml → validity → 本閘讀 validity
```

## Blocking（欄位條件）

| Gate | 讀取 | 未過 |
| --- | --- | --- |
| lock complete | spec + reference lock | 不得比較候選 |
| identity eligible | `decision==accepted` ∧ blocking 空 ∧ `validity==current` | 不得 mesh／rig |
| mesh geometry eligible | `mesh_qa.decision==pass` ∧ identity accepted/current | 不得 rig |
| deform eligible | body `deformation.decision==pass` ∧ identity accepted/current | 不得 face |
| face eligible | `facial_expression.decision==pass` ∧ UV/material pass ∧ identity accepted/current | 不得 outfit／animation |
| surface defer | UV/material fail 或 not-evaluated + blocking defect defer | 只允許 diagnostic；maturity 封頂 prototype |
| promotion | candidate `decision`+`reason`；provenance 最小或 unavailable/partial | 不得當正式候選 |
| completion | pack 適用欄位完整且 pass；identity accepted/current；無 blocking defect；`fresh_reviewer` | 不得 runtime-ready |

作者自驗可推進 stage；**completion 需要 fresh reviewer**。
