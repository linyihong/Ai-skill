# Observation — Subtitle Candidate Detector（precision on escalation）

**Run ID**：2026-10-01-subtitle-candidate-detector
**Kind**：Phase 3 observation → **contract_gap**（workflow 已吸收；產品 classifier 第一刀已落地；完整 pack dogfood 未完成）
**原則**：Loop＝recall；Candidate Detector＝precision。「畫面有文字」≠字幕 existence。

## 觀察（sanitized）

硬字幕短劇 reference pack（campus-memoir 類）成片：

| 訊號 | 觀察 |
| --- | --- |
| Escalation loop（46）後 | OCR recall 上升，但鐘錶時間／郵件畫面／UI 字進入字幕／破題語料 |
| 下游 | 成片破題語音與硬字幕對齊變差；雜訊被當 dialogue |
| 根因 | coverage／escalation 把「OCR 有字」當「字幕存在」 |

## 目標行為

```text
OCR text → Subtitle Candidate Detector
  → subtitle_like | non_subtitle_like | uncertain
Evidence Monitor densities / dialogue corpus：僅 subtitle_like
ASR dense ∧ subtitle_like=0 → may escalate
OCR text 全是 clock/mail/UI → 不因「有字」當 recovered
probe_decision.reason 記錄 escalate / no_escalation
```

## 分類

| 標籤 | 判定 |
| --- | --- |
| `subtitle_candidate_precision_missing` | contract_gap → [`47-subtitle-candidate-detector.md`](../47-subtitle-candidate-detector.md) → workflow [`text-evidence-subtitle-candidate.md`](../../../workflow/narrative-video-production/text-evidence-subtitle-candidate.md) |

## 產品 adapter（已落地／未完成）

| 項 | 狀態 |
| --- | --- |
| mechanical `subtitle_candidate`（morphology／geometry／zone） | landed（Windows SoT pipeline） |
| `dialogue_candidate_count`／Evidence Monitor 只計 subtitle_like | landed |
| `probe_decision`＋`ocr_scene_text_only` signal | landed |
| fuse 前過濾 non_subtitle | landed（indent 修復後） |
| unit：clock／email → non；CJK dialogue → subtitle；ASR-dense+clock → escalate | passed |
| 完整 pack dogfood（成片語音／字幕對齊） | **open**（host 進程在模型載入中途中斷；非 classifier 邏輯失敗） |

## Validation

- [x] Workflow + README／acquisition-loop／execution-flow／invariant 7f 已連結
- [x] Plan companion 47 + 本 evidence + `evidence/README.md` 索引
- [x] Product classifier 第一刀 + unit smoke

完整 episode pack dogfood（雜訊↓、對白不失真）仍 open：列於 companion 47 Acceptance 未勾項；host 進程在模型載入中途中斷，非本輪 classifier 邏輯失敗。
