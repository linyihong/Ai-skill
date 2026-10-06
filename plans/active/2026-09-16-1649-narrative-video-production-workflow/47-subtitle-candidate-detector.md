# Candidate: Subtitle Candidate Detector（precision on escalation）

Companion to [`46-evidence-acquisition-escalation-loop.md`](46-evidence-acquisition-escalation-loop.md)、
[`45-ocr-discovery-layout-probe.md`](45-ocr-discovery-layout-probe.md)。
Workflow：[`text-evidence-subtitle-candidate.md`](../../../workflow/narrative-video-production/text-evidence-subtitle-candidate.md)。

**不是新 Phase、不是新 Agent、不拿掉 escalation loop。**

## 問題

46 的 Evidence Monitor 提高了 OCR recall（ASR 密 → 加密／擴大 OCR），dogfood 卻把：

- 手錶時間（`12:35`）
- 郵件／信件畫面文字
- 手機 UI／浮水印／招牌

當成字幕候選，污染破題語音與字幕對齊。

根因：**「OCR 有字」被當成「字幕存在」**；缺 Subtitle Candidate Detector。

## 核心 refinement

```text
Loop = recall
Subtitle Candidate Detector = precision
OCR text → classify (geometry + zone + morphology + persistence + ASR support)
  → subtitle_like | non_subtitle_like | uncertain
Escalation only when subtitle-existence evidence is missing/suspicious
  — not when only non_subtitle scene text is present
```

## Phase 3 分類

| 標籤 | 判定 |
| --- | --- |
| `subtitle_candidate_precision_missing` | **contract_gap** → workflow 吸收；產品落地 mechanical classifier + escalation gate |

## 產品 adapter（本輪）

- `subtitle_candidate` mechanical module（morphology／geometry／zone evidence）。
- `dialogue_candidate_count`／Evidence Monitor densities 只計 `subtitle_like`（或 uncertain 可選；**排除** clock／document／ui）。
- `probe_decision` 記錄 escalate／no_escalation reason。
- Dogfood：`校园回忆录` 類片驗證雜訊下降且對白不失真。

## Evidence

[`evidence/2026-10-01-subtitle-candidate-detector.md`](evidence/2026-10-01-subtitle-candidate-detector.md)

## Acceptance

- [x] Workflow + README／acquisition-loop 連結
- [x] Product classifier + monitor／probe 接入（第一刀；見 evidence）
- [x] Escalation 不把 non_subtitle 當 recovered dialogue（unit `probe_decision`）
- [x] Sanitized evidence note under plan `evidence/`（indexed）
- [ ] 完整 pack dogfood（成片語音／字幕對齊、雜訊↓）— 見 evidence Validation deferred

此 checkbox 保持 open：獨立 frame negative pair 發現附近 clothing text
可借用 dialogue band 被升格。新的 failing regression 應保留，同時保護
鄰近真字幕；不可用 raw retention 或 process completion 宣稱 precision
通過。採集完整度與 timing readback 亦分開驗證，見
[`evidence/2026-10-06-attribution-precision-and-acquisition-integrity.md`](evidence/2026-10-06-attribution-precision-and-acquisition-integrity.md)。
