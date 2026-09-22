# SoT contracts & repository layout

Companion to [`_plan.md`](_plan.md)。Phase 1 實作清單。

## 第一版 SoT（優先順序）

0. **Translation Context Contract**（P0）— `contracts/translation-context.yaml` — `translation_context.source`／`target`（language、locale、script、register）；**所有 actor 必須消費**
1. **Translation Decision Contract** — `contracts/translation-decision.yaml`
2. **Expression Type Registry** — `registry/expression-types.yaml` + `registry/translation-strategies.yaml`
3. **Validation Contract** — `contracts/validation.yaml` + `registry/validation-rules.yaml`（含 **locale_consistency**）

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
    literal.yaml
    idiom.yaml
    slang.yaml
    proverb.yaml
    dialect.yaml
    wordplay.yaml
  evidence/                 # Phase 3 dogfood 後
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
| cultural_reference | preserve, adapt, explain |
| proper_noun / terminology | terminology_table, transliterate, preserve |
| proper_name_transliteration | transliterate, preserve_identity（陈→Chen，非語意翻譯） |

完整 enum 在 Phase 1 YAML。

## TDR 欄位方向（illustrative）

Phase 1 contract 為準；此處只作 plan 內導讀。

```yaml
translation_decision:
  id: TD-000123
  source: { text, locale }
  context: { speaker, addressee, scene }
  expression_analysis: { ... }
  constraints: { preserve, prohibit }
  candidates:
    - { text, strategy }
  selection: { policy, selected, rationale }
  review: { semantic, register, cultural }
  finality: { status }
```

## Mechanical validation（translation-core）

硬檢查（LLM 不可自證）：

- source segment coverage、missing／duplicated segments
- placeholders、numbers、dates、URLs、variables
- proper names、terminology、tags、markup
- source-language residue

## Locale validation（translation-core，P0）

與 mechanical 分欄；**不可**由 LLM 自證「英文稱謂也可以」。

- target **language** consistency
- target **locale** consistency（language ≠ locale）
- register consistency
- title／honorific consistency（例：`id-ID` + `Miss`/`Ms.` → `review` + `target_locale_residue`）
- dialect consistency（source／target）
- source-language residue（與 mechanical 可交叉引用）

Regression fixtures：[`05-example-address-title-chen-xiaojie-id.yaml`](05-example-address-title-chen-xiaojie-id.yaml) §`regression_fixtures`。

## Knowledge：`title_mapping`（種子 candidate space）

路徑方向：`knowledge/translation/locale/title-mapping.yaml`。  
`小姐 → Nona`（id）等條目只 **種子 Candidate Space**；final selection 仍靠 policy + context。見 [`04-dogfood-case-address-title-id-ID.md`](04-dogfood-case-address-title-id-ID.md)。

Subtitle **adapter** 追加（消費 NVP，不重定義 content）：

- CPS、duration、line count、line length、break position、speaker attribution、cue overlap、reading speed（對齊 NVP timing／layout 分工）

## 成語／諺語／流行語策略差異（設計備忘）

- **成語**（畫蛇添足）：semantic equivalent 或 target idiom，非字面 snake + feet
- **諺語**（覆水難收）：target proverb 或 semantic equivalent（What's done is done）
- **流行語**：先 source culture 語氣／誰說／對誰／極性，再問 target 是否有同等 social function

## Phase 1 checklist（從 plan 複製追蹤）

- [ ] `schema_version` 與 `artifact-record/v1` 慣例對齊 NVP records
- [ ] `runtime_projection: enabled: false` 直到 Phase 4 條件
- [x] plan-local example：`05-example-address-title-chen-xiaojie-id.yaml`（regression）
- [ ] examples 六類至少各一檔或合併為三檔（slang／proverb／dialect 優先）
- [ ] `translation-context.yaml` + locale_consistency rules
