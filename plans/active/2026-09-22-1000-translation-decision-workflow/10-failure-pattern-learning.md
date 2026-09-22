# Phase 1/3 reinforcement — Translation Failure Pattern learning

Companion to [`_plan.md`](_plan.md)。**不成**新 workflow；**不**把 dogfood 例句永久塞進 Selection prompt。

## Anti-pattern

```text
错误 → 加进 prompt（小姐→Cik, 今晚→malam ini, 八点半→…）
     → if Malay: 30 cases / if Arabic: 50 / if Japanese: 80
```

短期有效，長期膨脹，且與 Translation Decision（Constraint ≠ Selection）衝突。

## Target loop

```text
Dogfood
  → Failure Evidence
  → Failure Classification
  → Candidate Rule / Pattern (AI may propose)
  → Governance Review          ← LLM 不可自提升
  → Registry / Knowledge
  → Constraints + Validation on next Decision
```

對齊 [`enforcement/failure-learning-system.md`](../../../enforcement/failure-learning-system.md)：Capture → Classify → Promote → Strengthen → Validate。本檔只定 **translation domain** 形狀。

## Three durable layers（取代一坨 prompt）

| Layer | 內容 | 例 |
| --- | --- | --- |
| **Registry** | 穩定抽象 expression／failure | `address_title`, `temporal_clock`, `foreign_honorific_leak` |
| **Knowledge** | locale-specific Candidate Space seeds | `ms-MY/address-title`, `ja-JP/name-realization` |
| **Failure Guards** | dogfood → constraint／validation response | I13；見 [`failure-patterns.yaml`](../../../workflow/translation/registry/failure-patterns.yaml) |

Prompt／Selection adapter **只**接收：active constraints + Candidate Space + selection policy — 不是 100 個例句。

## Abstraction examples

| Prompt case（丟棄） | Abstract form（保留） |
| --- | --- |
| 小姐 → Cik/Puan | `address_title` + `native_honorific` + forbid foreign honorific leak |
| 今晚 = malam ini | `temporal_reference` + `preserve_temporal_direction` |
| 八点半 = 8:30 | `temporal_clock` + `preserve_numeric_time`（mechanical） |
| 陈总 = السيد… | `address_title`／name realization + locale knowledge seed |

Disposition on next error：

> 這是新的 failure_pattern，還是已知 pattern 的新 evidence？

## Landing

| 產物 | 路徑 |
| --- | --- |
| Contract（I13） | [`failure-pattern.yaml`](../../../workflow/translation/contracts/failure-pattern.yaml) |
| Registry | [`failure-patterns.yaml`](../../../workflow/translation/registry/failure-patterns.yaml) |
| Knowledge seeds | `ms-MY/address-title`, `temporal-reference` |
| Expression types | `temporal_reference`, `temporal_clock` |

## Explicit non-goals

- 自動從每次翻譯錯誤寫 active rule（無 Governance Review）
- `if japanese: base += …` Selection adapter 膨脹
- 把本 registry 複製成 `enforcement/failure-patterns/`（那是 agent 跨 skill 失效）
