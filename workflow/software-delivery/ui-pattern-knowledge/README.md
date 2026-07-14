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
  pattern-index.yaml          # Discovery index（≠ Runtime Projection）
  validation/
    entry-schema.yaml         # 結構完整性（Verifier / AI 寫 entry 時對照）
  entries/
    bottom_sheet.yaml         # 一 pattern 一檔（勿做成單一超大 glossary）
    modal_dialog.yaml
    drawer.yaml
    toast.yaml
    scrim.yaml
```

- **一 pattern 一 YAML**（現狀）。未來若需敘事長文，可選 `entries/<id>/entry.md` + `metadata.yaml`，**不**改回單一 mega-markdown。
- **不要**新增 `runtime-index.yaml` — 那會與 Runtime Projection 混淆；discovery 只用 `pattern-index.yaml`。

## Core schema（FROZEN v1.0）

見 [`../templates/ui-pattern-knowledge.entry.template.yaml`](../templates/ui-pattern-knowledge.entry.template.yaml) 與 [`validation/entry-schema.yaml`](validation/entry-schema.yaml)。  
Core 欄位凍結；新能力進 Extended，不反复改 Core key 名。

## 使用方式

1. 選型不明或宣告 pattern 對齊時，先讀 [`../ui-contracts.md`](../ui-contracts.md) **Pattern Knowledge Lock**。
2. Discovery：讀 [`pattern-index.yaml`](pattern-index.yaml)。
3. 寫／擴 entry：對照 template + `validation/entry-schema.yaml`（完整 entry = Core 齊）。
4. 專案 overlay 放 consumer repo；不要把專案私有名寫進本目錄 `entries/`。

**禁止**：`platform_map` 使用 Material / iOS / Fluent 等 DS 百科 key；不做獨立 Intent DB；不做全庫 glossary 單檔。

## Phase 2 dogfood（可推理性，不堆名詞）

目標：**五 Entry + 十 Selection Scenario 全部通過**（見 plan evidence）。  
順序：`scrim` → `modal_dialog` → `bottom_sheet` → `drawer` → `toast`（`feedback`）。  
**不要**急著加 popover / tooltip / fab / command_palette。
