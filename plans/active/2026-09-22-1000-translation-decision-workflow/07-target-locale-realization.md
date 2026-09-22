# Phase 1 reinforcement — Target-Locale Realization

Companion to [`_plan.md`](_plan.md)。**不改 Phase 0 freeze**；屬 Phase 1 contract 補強。

## Dogfood signal（Qwen ep12–14）

| source | en | ja | id |
| --- | --- | --- | --- |
| 陈小姐 | Ms. Chen | **Chenさん** | Nona Chen |

判讀：

- id／en：name Latin + locale title → 結構正確。
- ja：`さん` = address-title realization **正確**；`Chen` = **name realization 未完成** preferred katakana（チェン）。
- 不是籠統「翻譯錯」；也不是只修 prompt。缺口是 contract concept：**Target-Locale Realization**。

## 抽象（釘在 Translation Decision 內，不成新 workflow）

```text
Translation Strategy  → 要怎麼表達（social / semantic move）
Target-Locale Realization → 在目標語言裡實際長什麼樣（script / phonetics / established form）
```

Lifecycle 插入點（仍同一條 decision chain）：

```text
Expression Analysis
  → Identity / component resolution
  → Target-Locale Realization (candidate seeds)
  → Constraints → Feasible → Selection → …
```

## Phase 1 落地

| 產物 | 路徑 |
| --- | --- |
| Realization strategies registry | [`workflow/translation/registry/realization-strategies.yaml`](../../workflow/translation/registry/realization-strategies.yaml) |
| ja-JP name seeds（陈 only） | [`knowledge/translation/locale/ja-JP/name-realization.yaml`](../../knowledge/translation/locale/ja-JP/name-realization.yaml) |
| P0 regression | [`workflow/translation/examples/address-title-chen-xiaojie-ja.yaml`](../../workflow/translation/examples/address-title-chen-xiaojie-ja.yaml) |

## 煞車

- **禁止** `all_foreign_names → katakana` mechanical FAIL。
- `preferred_script: katakana` = Constraint／preference on Candidate Space。
- `Chenさん` → `name_realization: review`（可 waiver），不是自動 incorrect。
- 完整姓氏庫 **不出** Phase 1。

## Cross-locale 價值

同一 Source Expression，不同 Target Locale → 不同 realization（Ms. Chen／チェンさん／Nona Chen）。日文只是第一個把缺口暴露清楚的 locale。
