# UI Pattern Knowledge — reusable rule seeds（Ai-skill）

本目錄存放 **Ai-skill reusable 規則種子**（首波：`overlay` family）。

| 誰擁有 | 內容 |
| --- | --- |
| **Ai-skill** | Core selection / family / neighbors / intent_examples；optional Extended 雛形 |
| **Consumer 專案** | 採用名、`platform_map.project`、composition、project alias；可覆寫 / 延伸 entries |

## 佈局（Phase 1 鎖定 — 不要膨脹）

```text
ui-pattern-knowledge/
  README.md
  pattern-index.yaml
  composition_rules.yaml      # Phase 3+：composition 問題寫這裡（勿改 entries）
  validation/
    entry-schema.yaml
  entries/                    # FROZEN after Phase 2（inferability validated）
    …
  compositions/
    episode_detail.yaml       # Phase 3 Pattern Composition seed
```

- **一 pattern 一 YAML**（現狀）。未來若需敘事長文，可選 `entries/<id>/entry.md` + `metadata.yaml`，**不**改回單一 mega-markdown。
- **不要**新增 `runtime-index.yaml` — 那會與 Runtime Projection 混淆；discovery 只用 `pattern-index.yaml`。

## Core schema（FROZEN v1.0）

見 [`../templates/ui-pattern-knowledge.entry.template.yaml`](../templates/ui-pattern-knowledge.entry.template.yaml) 與 [`validation/entry-schema.yaml`](validation/entry-schema.yaml)。  
Core 欄位凍結；新能力進 Extended，不反复改 Core key 名。

## 使用方式

1. 選型不明或宣告 pattern 對齊時，先讀 [`../ui-contracts.md`](../ui-contracts.md) **Pattern Knowledge Lock**。
2. Discovery：讀 [`pattern-index.yaml`](pattern-index.yaml)。
3. 寫／擴 entry：對照 template + `validation/entry-schema.yaml`（完整 entry = Core 齊）。**Phase 3+ 預設不改 entries**——composition 問題寫 [`composition_rules.yaml`](composition_rules.yaml)。
4. Pattern Composition：見 [`compositions/`](compositions/) + composition template。
5. 專案 overlay 放 consumer repo；不要把專案私有名寫進本目錄 `entries/`。

**禁止**：`platform_map` 使用 Material / iOS / Fluent 等 DS 百科 key；不做獨立 Intent DB；不做全庫 glossary 單檔。

## Phase 狀態

- **Research ladder**：P1 Representability (Entry) → P2 Inferability (Scenario) → P3 Composability (Screen)。
- **Phase 2 Completed**：inferability（非 Coverage）。見 plan `evidence/phase2-summary.md`。
- **Phase 3**：Pattern Composition · Episode only · H4/H5/H6；`composition_rules.yaml` = **Composition Constraints**（非 Rule Library）；**不要**為 Composition 補 Player/App Bar Entry。
