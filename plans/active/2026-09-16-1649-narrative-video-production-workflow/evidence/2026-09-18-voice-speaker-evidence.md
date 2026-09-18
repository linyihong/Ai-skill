# Observation — voice / speaker evidence

**Run ID**：2026-09-18-voice-speaker-evidence  
**Kind**：Phase 3 observation（**不是** `contract_gap`；**不**擴 workflow；**不**接聲紋產品）  
**原則**：ASR supplies speech text. Voice supplies speaker evidence. Identity resolution combines them later.

即使兩句都沒有人名，不同 `speaker_id` 仍表示很可能不同說話者。跨集：ep01 `entity` + `face_cluster` + `voice_cluster`、`name: null`；ep02 OCR／ASR「林雪」+ 同一 cluster → `identity_link`，不覆寫 ep01 列。

## Canonical hang（採集層）

```yaml
asr_segment:
  id: asr_00182
  timing: { start: 123.42, end: 125.81 }
  transcript: { text: "你真的要去？" }
  speaker:
    speaker_id: speaker_04
    diarization_ref: diarization_04
    voice_evidence_ref: voice_track_04
  evidence_links:
    visual_text_refs: [ocr_00182]
    face_track_refs: [face_track_023]

voice_evidence:
  voice_track_id: voice_track_04
  features: { embedding_ref: null }
  source_segments: [asr_00182, asr_00183]
  identity: { candidate_refs: [] }   # 解析後才填；不是 ASR 輸出
```

禁止在 ASR 階段寫 `speaker_04 = face_track_023` 或 `speaker_04 = character_lin`。`evidence_links` = 同時段共現。

## 與既有候選

Face：[`11-face-as-candidate-evidence.md`](../11-face-as-candidate-evidence.md)。身份先於名稱：[`08-identity-precedes-naming.md`](../08-identity-precedes-naming.md)。Visual text：[`09-visual-text-evidence.md`](../09-visual-text-evidence.md)。Linking：[`12-evidence-refinement.md`](../12-evidence-refinement.md)。

## 真實片子要數

無對白人名時 speaker 切分是否被消費；是否出現 speaker=character 捷徑。沒被 bible／catalog／matching／EDR 消費的欄位不進 schema。
