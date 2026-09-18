# Observation — episode evidence vs knowledge accumulation

**Run ID**：2026-09-18-episode-vs-knowledge-accumulation  
**Kind**：Phase 3 observation（**不是** `contract_gap`；**不**擴 workflow；**不**建完整 Knowledge DB）  
**原則**：Episode analysis may propose learning candidates. It may not write knowledge.

首份真實 run 把 vocative 暱稱直接升成 canonical identity，正是把本集觀察與知識累積混在同一欄。正確路徑：

```yaml
episode_evidence:
  identity_observation:
    observed_id: entity_003
    speaker_id: speaker_02
    candidate_names:
      - { name: vocative_nickname, status: candidate }

learning_candidate:
  domain: identity
  type: identity_link
  status: pending_verification
  requested_action: extend_identity_evidence
```

Verifier 至少看：face／voice continuity、OCR name（role 已解析）、explicit self-reference、cross-episode evidence、relationship evidence。不足 → keep candidate。不得因單集稱呼寫 `canonical_name: resolved`。

Knowledge 不是 AI 記憶。每筆必須有 source、evidence、provenance、status、validity。先分 Identity／Narrative／Retrieval 三個知識面，不要一個大庫全塞。

機械累積（例：反覆出現的 watermark 正規化框）走 Mechanical Registry；語意累積（稱呼 → 關係候選）走 Knowledge Store。兩者都要 Observation → Repeated Evidence → Candidate → Validation → Promotion。

下一集 identity dogfood 只驗：vocative 停在 episode evidence；跨集才產生 learning candidate；無 verifier 不得改 knowledge／mechanical registry。
