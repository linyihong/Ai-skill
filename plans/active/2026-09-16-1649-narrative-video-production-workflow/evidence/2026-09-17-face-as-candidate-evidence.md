# Observation — face as candidate evidence (not identity key)

**Run ID**：2026-09-17-face-as-candidate-evidence  
**Kind**：Phase 3 observation（**不是** `contract_gap`；**不**擴 workflow；**不**接辨識模型）  
**原則**：Face is an evidence link / candidate evidence for name resolution, not a direct identity decision source.

ASR「林小姐」與 OCR「林雪」只能證明文字吻合。同一時間窗再掛 `face_track`，才形成完整 candidate set，供之後 Identity Resolution 使用。跨集可先 cluster tracks，等某一集出現名字再解到 `character_id`。

## 掛點（現在）vs 禁止（現在）

| 現在 | 禁止現在做 |
| --- | --- |
| Face Detection / Tracking → `face_track_id` | Face Recognition → Character Identification |
| Evidence link 引用 `face_track` + keyframe | 把 Face 當 ASR／OCR 人名仲裁器的判定來源 |
| `embedding_ref` / `cluster_id` 可 `null` | 用「長得像」覆寫或建立 canonical identity |

```yaml
evidence_link:
  time_range: [123.2, 125.8]
  asr: { ref: asr_segment_182, text: "林小姐，你真的要去？" }
  ocr: { ref: ocr_region_77, text: "林雪" }
  face:
    track_refs: [face_track_023]
    keyframe_refs: [frame_004821]
    embedding_ref: null
    cluster_id: null
  relations:
    - { type: temporal_overlap, refs: [asr_segment_182, ocr_region_77, face_track_023] }
    - { type: candidate_identity, source: face, target: character_lin, status: candidate }
  status: candidate
```

`status: candidate` 不是已解析身份。`character_lin` 只在 Identity Resolution **之後**才成立。

## 與既有候選的關係

- Observable 層：[`07-material-fact-extraction.md`](../07-material-fact-extraction.md) — Face Track 屬機械採集；cluster／embedding 是後續可選。
- 仲裁與政策：[`2026-09-18-evidence-refinement.md`](2026-09-18-evidence-refinement.md) — Face 只進 supporting／escalation，不當 Decision SoT。
- 身份先於名稱：[`08-identity-precedes-naming.md`](../08-identity-precedes-naming.md) — Face／voice 不是 identity key。
- 畫面文字：[`09-visual-text-evidence.md`](../09-visual-text-evidence.md) — OCR 仍是獨立通道，不經 Face 判定。
- 已解析表：[`05-series-cast-canonicalization.md`](../05-series-cast-canonicalization.md) — 不得用 Face 捷徑寫入。

未來共用（升格後）：人名仲裁、跨集追蹤、角色 canonicalization、場景角色集合、speaker resolution、出場證據。

## 真實片子要數

是否留下可引用的 `face_track`；是否出現「OCR+臉=定名」捷徑；跨集 unnamed track 是否等到名字證據才 link。沒被 bible／catalog／matching／EDR 消費的欄位不進 schema。
