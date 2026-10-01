# Candidate: Evidence Acquisition & Escalation Loop

Companion to [`45-ocr-discovery-layout-probe.md`](45-ocr-discovery-layout-probe.md)、
[`12-evidence-refinement.md`](12-evidence-refinement.md)、
[`43-multimodal-text-evidence-resolution.md`](43-multimodal-text-evidence-resolution.md)。
**不是新 Phase、不是新 Agent、不推翻 Phase 1/2、不接 runtime。**
Workflow：[`text-evidence-acquisition-loop.md`](../../../workflow/narrative-video-production/text-evidence-acquisition-loop.md)。
Evidence：[`evidence/2026-10-01-evidence-acquisition-loop.md`](evidence/2026-10-01-evidence-acquisition-loop.md)。

## 問題

45 已把「probe miss ≠ 無字幕」寫進 OCR discovery。Dogfood 仍見：

1. Probe／discovery 已 `dialogue_candidates`／`has_hardsub=True`，但成片中後段時間窗對白漏錄。
2. OCR 與 ASR 密度不合理時，系統仍把第一次採集當最終 acquisition。
3. Story Event 過少往往在下游才被發現，acquisition 階段沒有「證據可疑 → 升級採集」閘。

根因層級比單一 OCR region 更高：**缺少 Evidence Monitor + Escalation Policy**。

## 核心 refinement

```text
Probe = discovery（非 exclusion）
Monitor = confirmed | inconclusive | suspicious（依 expected + cross-modal）
suspicious → Escalation Level 0–4（機械先於 LLM）
only then → Text Resolution / Story
```

OCR×ASR 互為 recovery 觸發：ASR 密／OCR 疏 → OCR escalate；反向 → ASR escalate。

## 與既有 companion 的邊界

| 檔 | 管什麼 | 本檔補強 |
| --- | --- | --- |
| 45／ocr-discovery | 單模態 region miss／layout | 跨模態 Monitor；時間窗 coverage gap |
| 43／multimodal resolution | observed→resolved 語義 | 先確保 acquisition 充分，再決議 |
| 12／refinement | linking／story 外環 | 本檔是 acquisition 閘，不取代 12 |

## Phase 3 分類

| 標籤 | 判定 |
| --- | --- |
| `evidence_acquisition_monitor_missing` | **contract_gap** → workflow 吸收；產品落地 Monitor＋至少 `ocr_asr_coverage_gap` escalation |

## 產品 adapter（本輪）

- 新增 mechanical Evidence Monitor（窗密度＋signals）。
- OCR 路徑在 fuse 後若 `ocr_asr_coverage_gap` → Level 1／2 重採（denser／wider／window）並記 before／after。
- 不把 Level-4 LLM 當第一刀；不因單次 miss 寫無字幕。
