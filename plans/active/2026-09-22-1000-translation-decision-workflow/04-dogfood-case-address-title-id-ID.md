# Dogfood case — 陈小姐 → id-ID（稱謂／語系判定混線）

Companion to [`_plan.md`](_plan.md)。**第一個 real-data failure case**；機械 regression 見 [`05-example-address-title-chen-xiaojie-id.yaml`](05-example-address-title-chen-xiaojie-id.yaml)。

## 觀察到的輸入／輸出

```json
{
  "src": "陈小姐",
  "dst": "Nona Chen",
  "episode": 14
}
```

目標語為 **印尼語**（`id` / `id-ID`）。`Nona Chen` 本身可以是正確結果（`Nona` ≈ Miss／小姐），但常見錯誤路徑是 workflow 把片段當「中文 → 英文」處理，產出 `Miss Chen`、`Ms. Chen` 等 **target locale 不一致** 的稱謂。

## 暴露的三責任混線

| 責任 | 應做 | 混線時 |
| --- | --- | --- |
| **Language / Locale Resolution** | 在 Translation Decision **之前** 固定 `source` + `target` locale；所有 actor 只吃 `TranslationContext` | 目標語只存在 pipeline 最外層，未進入每段 decision |
| **Name resolution** | `陈` = proper name（family_name）→ transliteration `Chen`，非語意翻譯 | 整串 `陈小姐` 當一般句子丟給模型 |
| **Address title translation** | `小姐` = address_title → **locale-aware** 候選（id：`Nona` 等） | `小姐` 當普通詞彙或英文預設 `Miss`/`Ms.` |

根因通常 **不是**「LLM 不會翻」，而是 **target locale 未進入 expression decision**。

## 正確決策鏈（目標 id-ID）

```text
陈小姐
  ↓
[Locale Resolution]  target = id-ID（script: Latn）
  ↓
[Expression Analysis]
  陈 → proper_name / family_name / transliteration
  小姐 → address_title / honorific / locale_equivalent
  ↓
[Constraints]
  preserve person identity; title must match target locale; no semantic translation of name
  ↓
[Candidate Space]
  Nona Chen | Miss Chen | Ms. Chen | Chen
  ↓
[Selection]  policy: target_locale_naturalness + preserve_address_title + subtitle_register
  ↓
Nona Chen
  ↓
[Locale Validation Gate]  PASS（無 en title residue）
```

## 錯誤路徑（regression 要擋）

```text
陈小姐 → LLM → English-like title → Miss Chen / Ms. Chen
```

當 `target.locale = id-ID` 且 `dst` 含英文稱謂 token（`Miss`、`Ms.`、`Mr.`…）時：

```yaml
validation:
  locale_consistency:
    status: review   # 或 warn，依 profile
    reason: target_locale_residue
```

LLM 不得用「Ms. 也可以吧」關閉此項；需 `needs_review` 或 explicit selection rationale + waiver policy。

## P0 contract 增量（相對初版 plan）

1. **Translation Context Contract**（Phase 1 SoT #0 或併入 `source.yaml`）— 所有 translation actor **禁止** 只吃 `{ src, dst }`。
2. **Locale Resolution** — 在 Expression Analysis **之前** 的 mandatory step（INGEST 子步或 Layer 1.5）。
3. **Structured Expression Analysis** — `surface` + `structure[]`（span、type、role、handling）。
4. **Registry：`name_with_address_title` / `address_title`** — `translation_mode: locale_aware`；knowledge `title_mapping` 提供 **candidate space**，非 final answer。
5. **Locale Validation**（併入 validation contract）— target language／locale／register／title-honorific／dialect／source residue 分欄。

`target language ≠ target locale`：例如 `zh-CN → id-ID` vs `zh-TW → id-ID` vs `zh-HK → id-ID` 可能影響 source 表達與 title 候選，不可只寫 `Chinese → Indonesian`。

## Knowledge：`title_mapping`（示意，非 exhaustive）

Registry 不應只是一張死表；表只 **縮小／種子** candidate space，Selection 仍看 scene、relationship、formal／informal、subtitle style。

```yaml
# knowledge/translation/locale/title-mapping.yaml (Phase 3 方向)
title_mapping:
  zh:
    小姐:
      id: [Nona]           # 候選，非唯一答案
      en: [Ms., Miss]
      ja: [さん]
```

同一 `小姐` 在 id-ID 下仍可能選 `Nona Chen` vs `Miss Chen`（后者應被 locale gate 標 review）。

## Phase 對照

| Phase | 本案例觸發物 |
| --- | --- |
| 0 | 本 companion + example YAML |
| 1 | `contracts/translation-context.yaml`、`expression-analysis` structure、`validation` locale_consistency |
| 3 | dogfood run 引用 episode 14；`locale_consistency` regression fixture |

## 與 NVP

若 `dst` 進 caption pack 且 `target_locale: id-ID`，`content_gate` 應能引用 TDR + locale validation；timing／layout 不變。
