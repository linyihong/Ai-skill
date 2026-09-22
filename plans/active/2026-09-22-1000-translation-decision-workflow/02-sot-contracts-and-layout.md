# SoT contracts & repository layout

Companion to [`_plan.md`](_plan.md)。Phase 1 實作清單。

## 第一版 SoT（優先順序）

0. **Translation Context Contract**（P0）— `contracts/translation-context.yaml` — `translation_context.source`／`target`（language、locale、script、register）；**所有 actor 必須消費**
1. **Translation Decision Contract** — `contracts/translation-decision.yaml`
2. **Expression Type Registry** — `registry/expression-types.yaml` + `registry/translation-strategies.yaml`
3. **Validation Contract** — `contracts/validation.yaml` + `registry/validation-rules.yaml`（含 **locale_consistency**）
4. **Failure Pattern Contract**（I13）— `contracts/failure-pattern.yaml` + `registry/failure-patterns.yaml` — dogfood → abstract guards；**不成** Selection prompt 例句庫

定穩後，LLM／翻譯模型／供應商 = Selection Actor 替換，不重做 workflow。

## 建議 `workflow/translation/` 布局（Phase 1）

```text
workflow/translation/
  README.md
  execution-flow.md
  contracts/
    translation-context.yaml
    source.yaml
    expression-analysis.yaml
    translation-decision.yaml
    validation.yaml
    finality.yaml
  registry/
    expression-types.yaml
    translation-strategies.yaml
    validation-rules.yaml
  adapters/
    document.yaml
    subtitle.yaml
    ui.yaml
  examples/
    address-title-chen-xiaojie-id.yaml   # copy from plan 05-example-*
    address-title-chen-xiaojie-ja.yaml
    title-kongjie-yiriqianli.yaml        # I12 title + invented_information
    literal.yaml
    idiom.yaml
    slang.yaml
    proverb.yaml
    dialect.yaml
    wordplay.yaml
  evidence/                 # Phase 3 dogfood 後（plan-local）
```

## 建議 `knowledge/translation/` 布局（Phase 3 骨架）

```text
knowledge/translation/
  README.md
  terminology/
  locale/
  dialect/
  cultural-expression/
  character-voice/
```

Registry 不寫死在 prompt；新增日本語関西弁／若者言葉等 = 擴 registry + knowledge，不改整條 workflow 形狀。

## Expression type → strategy（v0 草案）

| type | 允許 strategies（摘要） |
| --- | --- |
| literal | semantic_translation |
| idiom | semantic_equivalent, cultural_equivalent, literal_plus_explanation |
| proverb | target_proverb, semantic_equivalent |
| slang | target_slang, register_preserving_adaptation |
| meme | cultural_adaptation, preserve_reference |
| dialect | dialect_equivalent, regional_register, preserve_source_flavor |
| wordplay | recreate_wordplay, semantic_substitution |
| honorific | relationship_preserving |
| address_title | locale_aware, target_locale_equivalent |
| name_with_address_title | transliteration + locale_aware title; components name + title |
| title | title_faithful, title_semantic_adaptation（需 content.type=title） |
| title_idiom | semantic_equivalent, cultural_equivalent, title_semantic_adaptation |
| title_hook / title_wordplay | title_semantic_adaptation, preserve_drama_hook |
| part_marker | part_normalization, preserve |
| cultural_reference | preserve, adapt, explain |
| proper_noun / terminology | terminology_table, transliterate, preserve |
| proper_name_transliteration | transliterate, preserve_identity（陈→Chen，非語意翻譯） |

完整 enum 在 Phase 1 YAML。Title／I12：[`09`](09-title-content-type.md)。

## TDR 欄位方向（illustrative）

Phase 1 contract 為準；此處只作 plan 內導讀。

```yaml
translation_decision:
  id: TD-000123
  translation_context: { source, target }   # authoritative locales
  expression_analysis: { ... }              # artifact; producer opaque
  candidate_space_refs: [...]               # registry / title_mapping seeds
  constraints: { preserve, prohibit }
  candidates:
    - text: "Nona Chen"
      strategy: locale_aware
      feasible: true
    - text: "Miss Chen"
      strategy: literal_equivalent
      feasible: false
      reason: target_locale_mismatch
  selection:
    policy: [...]
    selected: "Nona Chen"
    rationale: "..."
    decision_basis:
      - target_locale
      - expression_analysis
      - scene_context
      - candidate_registry
  review: { semantic, register, cultural }
  finality: { status }
```

## Mechanical validation（translation-core）

硬檢查（LLM 不可自證）：

- source segment coverage、missing／duplicated segments
- placeholders、numbers、dates、URLs、variables
- proper names、terminology、tags、markup
- **source_language_residue**（與 target_locale_residue **分欄**，I6）

## Locale validation（translation-core，P0）

與 mechanical 分欄；**不可**由 LLM 自證關閉。

- target language／locale／register／title-honorific／dialect
- **target_locale_residue** → 預設 **`review`**（I5），非自動 incorrect／fail
- `accepted` 需 explicit rationale 或 waiver

**P0 regression fixture**：[`05-example-address-title-chen-xiaojie-id.yaml`](05-example-address-title-chen-xiaojie-id.yaml)。

## Knowledge：`title_mapping`（種子 Candidate Space）

路徑：`knowledge/translation/locale/title-mapping.yaml`（Phase 3）。  
**Registry invariant（I7）**：Candidate Space ≠ Final Answer。禁止把 mapping 當 translation truth。

Subtitle **adapter**（Phase 2）追加 NVP timing／layout 消費，不重定義 content。

## Phase 1 checklist

- [ ] `schema_version` 對齊 NVP records 慣例
- [ ] `runtime_projection: enabled: false`
- [x] P0 regression：`05-example-*`（I4）
- [ ] contracts 寫清 I1／I2／I3（見 [`06`](06-phase-0-freeze-invariants.md)）
- [ ] `decision_basis` + Finality I9
- [ ] slang／proverb／dialect examples
- [ ] **不加** I10 禁止清單（memory／prompt／score／route…）
