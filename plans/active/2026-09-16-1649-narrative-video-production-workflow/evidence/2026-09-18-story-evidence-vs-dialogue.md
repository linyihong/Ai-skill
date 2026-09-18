# Observation — story evidence vs dialogue dump

**Run ID**：2026-09-18-story-evidence-vs-dialogue  
**Kind**：Phase 3 observation（**不是** `contract_gap`；**不**擴 workflow）  
**原則**：Plot is evidence-backed story-state change and narrative events, not an ASR summary.

閒聊有語意 ≠ 應進 Script／Template／EDR。不要追求「現在就得到正確劇情」。路徑：Evidence → candidate interpretation → story state → new evidence → revision → verified story evidence。

## Relevance 與 narrative_role

問「對故事狀態有沒有可觀察影響」，不要問「這是不是劇情」。`unknown` 合法。

例類：`social`／`filler`、`characterization`、`relationship`、`world_building`、`setup`、`plot_event`、`conflict`、`decision`、`revelation`、`foreshadowing`、`transition`、`unknown`。

Low → archive（保留 raw evidence）。High → Story Event。Medium／unknown → 更深 LLM 或等後集。咖啡閒聊可 `changed: false` 仍留在資料室。

## Event 聚合，不是逐句 summary

00:12 A「你真的要去？」／00:14 B「我要去。」／00:18 拿鑰匙／00:20 離開 → **一個** `event`（type departure），掛多條 evidence，不是四個劇情點。

## Story state change

重要性看狀態差，不看句子漂不漂亮。例：before `destination: unknown` → after `tokyo` = `information_reveal`。單句「你明天真的要去東京嗎？」沒有 state 時可保持 unknown。

## Traceability

獨立 LLM 只檢查：Story Event 能否回指 ASR／OCR／face／action／voice。能 → 保留。不能 → 降級／重分析。不得因「這句很有意思」升格。

## 真實片子要數

ASR dump→摘要 是否污染 matching／EDR；low 列是否被刪；event 是否逐句膨脹。沒被消費的 role enum 不進 schema。
