# Pattern Prompt Expansion（transient checklist）

> **用途**：從 canonical Pattern Knowledge（repo entry / composition）展開成**本次任務**執行 checklist。  
> **禁止**：把某次任務的 checklist **commit 進 canonical Knowledge**（Ai-skill 種子或專案 `ui-pattern-knowledge/` entry）。Checklist 屬 runtime／任務報告；可選在 plan `evidence/` 留**摘要**，不得把展開項整份回寫成 entry。

## Source

- **Entries loaded**: <path or ids>
- **Composition** (if any): <path or `none`>
- **Task claim**: <pattern selection / alignment / implementation>
- **Non-goals**: <e.g. redesign tokens, governance domain rewrite>

## Expansion Shape

Each expanded item MUST carry `evidence_level` so verified ≠ observed ≠ hypothesized 不被當成同一硬度。

| # | Item | Source pattern | Derived from | evidence_level | Hardness |
| --- | --- | --- | --- | --- | --- |
| 1 | <actionable check> | <pattern_id> | selection_rules / recipe / anti_pattern / composition | verified \| observed \| hypothesized | hard \| suggest \| hint |
| 2 | … | … | … | … | … |

Hardness policy（對齊 D7）:

- `verified` → 可當 hard gate
- `observed` → 建議；未證明勿升級為硬約束
- `hypothesized` → 僅提示

## Intent → Selection (task-local)

- **Task intent** (推導，不進 Intent DB): <one line>
- **Selected pattern(s)**: <ids>
- **Rejected neighbors**: <id — why not>
- **Composition pointer**: <screen composition path or TBD>

## Recipe notes (Extended; may be unknown)

- **Recipe status**: unknown \| partial \| complete
- **Required capabilities checked**: <list or N/A>
- **Deferred**: <owner / follow-up>

## Closure

- **Decision**: selected / deferred / needs-human
- **Canonical Knowledge changed?** no（預設）／yes — **only** if entry/composition schema 本身被修正，非本 checklist
- **Where checklist lives**: task report \| plan `evidence/` summary \| discarded
