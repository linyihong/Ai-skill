# Candidate: Story promotion gate

Companion to [`_plan.md`](_plan.md)。**候選 gate，不是 Phase 2 workflow schema。**
首份真實 source-analysis run 證明 observable／evidence unit 可用，但
candidate → accepted 過鬆。去敏 run：
[`evidence/2026-09-18-real-run-promotion-gaps.md`](evidence/2026-09-18-real-run-promotion-gaps.md)。

## Story Evidence 升格

`traceable: true` 只證明「找得到來源」，不證明事件語意或 state claim 成立。
`story_evidence.status: accepted` 必須同時滿足：

1. 所有 blocking upstream decisions 已 resolved；任何 `final: null` → stop。
2. Narrative consumer 使用 resolved text；raw ASR／未解析 subtitle candidate
   不得成為 final event semantics。
3. Event semantics 已 resolved；`dialogue_cluster`／raw summary 只算 evidence
   grouping，不是已解析 event。
4. Evidence traceability 通過。
5. 若宣稱 state change，必須有非空 `subject`、`change_type`，以及可驗證的
   `before`／`after` 或明示「新增已知事實」；空物件與 null 不得 accepted。
6. Independent verifier 通過。

`possible_state_delta` 只能停在 `state_change_candidate`。生命週期：

```text
candidate → resolved → accepted
```

禁止 candidate 直接跳 accepted。

## Semantic authority

Mechanical 可以建立 event window、算 overlap、提出 `role_candidate`，但不得
自行宣告 `narrative_relevance: high`。除非有已登記、可重跑、已驗證的
deterministic policy；否則寫：

```yaml
narrative_relevance:
  status: candidate
  level: unresolved
  proposed_by: mechanical
  escalate_to: llm_narrative
```

LLM proposal 同樣不是 final。

Text resolution、window、event assembly 與 candidate relations 見
[`19-text-resolution-and-narrative-assembly.md`](19-text-resolution-and-narrative-assembly.md)。
Relation 的 temporal proximity 只能當 evidence，不能自行宣告因果。

## Event basis

每個 event candidate 應能回答「由哪些 modality 支持」，但不複製全部 evidence：

```yaml
basis:
  dialogue_refs: [...]
  action_refs: [...]
  visual_change_refs: [...]
  location_change_refs: [...]
  participant_change_refs: [...]
  visual_text_refs: [...]
```

空 refs 合法；它讓 verifier 看出「只有 dialogue」而不是假裝 multimodal event。

## Identity／name promotion

- `speaker_id`／voice cluster ≠ character。
- Vocative 是 **稱呼關係 evidence**；先解析 addressee，不得把稱呼貼到 uttering
  speaker，更不得直接 `canonical_name: resolved`。單集 vocative 只進 Episode
  Evidence；跨集累積走 Learning Inbox，見
  [`18-episode-vs-knowledge-accumulation.md`](18-episode-vs-knowledge-accumulation.md)。
- OCR `name_mention` ≠ entity。Visual-text role 未解析、watermark variants、
  普通片語、情緒詞與稱謂都不得建立 resolved entity。
- 原始 ASR 永不覆寫；校正寫 `text_resolution: raw | candidate_corrected |
  resolved` + alternatives。

## Archive reactivation

Low relevance 保留 `archive_reason`，並允許後集 evidence 重新啟用：

```yaml
reactivation:
  allowed: true
  triggers: [identity_link, relationship_change, later_episode_reference]
```

本候選至少再跑一集，驗證 dialogue-only event、low relevance reactivation、
跨集 identity resolution 後，才決定是否升 workflow contract。
