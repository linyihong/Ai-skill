# Candidate: Story evidence is state change, not dialogue dump

Companion to [`_plan.md`](_plan.md)。**候選原則，不是 workflow schema。** 禁止把整段 ASR 丟給 LLM 當劇情摘要。  
觀察：[`evidence/2026-09-18-story-evidence-vs-dialogue.md`](evidence/2026-09-18-story-evidence-vs-dialogue.md)。候選閘：[`03-architecture-invariants.md`](03-architecture-invariants.md) §Candidate invariant。

核心：**劇情不是對話摘要；劇情是被 Evidence 支持的 Story State Change 與 Narrative Event。** Dialogue 只是 evidence carrier。有對話 ≠ 有劇情資訊。

```text
observable → linking → evidence_unit → event candidates → story state → narrative decision (script / EDR)
```

聚合層見 [`16-evidence-unit.md`](16-evidence-unit.md)：**先停擴 observable**。`evidence_unit` 不是劇情，只是同一時間窗的證據包。

中間必過 **Relevance Assessment**：不問「這是不是劇情」（單句會誤殺伏筆），問「這組 evidence 有沒有造成可觀察的故事狀態變化」。`narrative_role` 允許多類與 `unknown`，禁止每句二元分類。多句 dialogue／action／OCR／face 聚成 **event candidate**，不要四句四個劇情點。

Low relevance **archive，不刪**（後集可能是伏筆）。Medium／unknown → LLM escalation 或等後續集。High → Story Evidence + state change。獨立審查：Story Event 必須回指原始 evidence（Evidence Traceability Gate）。接 [`12-evidence-refinement.md`](12-evidence-refinement.md)。Phase 3 **不**改 workflow。

首份真實 run 證明 `traceable: true` 仍可能錯誤升格：upstream `final: null`、
空 before／after 與 null subject 不能 accepted。完整升格條件與 event basis 見
[`17-story-promotion-gate.md`](17-story-promotion-gate.md)。
