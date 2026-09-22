# Dogfood case — 陈小姐 → id-ID（稱謂／語系判定混線）

Companion to [`_plan.md`](_plan.md)。**P0 regression fixture**（不只是一般 example）：[`05-example-address-title-chen-xiaojie-id.yaml`](05-example-address-title-chen-xiaojie-id.yaml)。Freeze：[`06`](06-phase-0-freeze-invariants.md) I4–I6。

## 觀察到的輸入／輸出

```json
{
  "src": "陈小姐",
  "dst": "Nona Chen",
  "episode": 14
}
```

目標 **id-ID**。`Nona Chen` 可為正確結果。常見偏離路徑：未綁定 target locale，走「中文→英文」稱謂，產出 `Miss Chen`／`Ms. Chen` → **`locale_consistency: review`**（`target_locale_residue`），**不是**自動判定「翻譯錯誤／fail」（borrowed／international／character voice／mixed language 可能合法；`accepted` 需 rationale／waiver）。

## 暴露的三責任混線

| 責任 | 應做 | 混線時 |
| --- | --- | --- |
| **Locale Resolution**（≠ Language Detection） | consumer／locale pack → `TranslationContext`；target = Constraint | 每段 LLM 猜 target／最外層 locale 未進 decision |
| **Name resolution** | `陈` → transliteration `Chen` | 整串當普通句子 |
| **Address title** | `小姐` → locale-aware **Candidate Space**（含 Nona） | 當普通詞或英文預設 Miss／Ms. |

根因通常不是「LLM 不會翻」，而是 **target_locale 未當 authoritative Constraint 進入 decision**。

## 正確決策鏈（目標 id-ID）

```text
NVP locale pack / job → TranslationContext (id-ID)
  ↓
Expression Analysis (artifact)
  陈 → proper_name / transliteration
  小姐 → address_title / locale_equivalent
  ↓
Candidate Space (title_mapping seeds: Nona, Miss, Ms., …)
  ↓
Constraints → Feasible Candidates
  Nona Chen feasible:true
  Miss Chen feasible:false reason:target_locale_mismatch
  ↓
Selection + decision_basis → Nona Chen
  ↓
Locale Validation → pass（無未 waiver 的 en title residue）
```

## Regression（P0）

`target.locale = id-ID` 且 dst 含 `Miss`／`Ms.`／`Mr.`…：

```yaml
validation:
  locale_consistency:
    status: review
    reason: target_locale_residue
```

不得以 LLM「也可以吧」關閉；需 `needs_review` 或 explicit rationale + waiver。

**勿與 `source_language_residue` 合併**（I6）：ja→id 殘日文 ≠ zh→id 的 English title。

## Knowledge：`title_mapping`

只種子 Candidate Space；**≠ Final Answer**（I7）。
