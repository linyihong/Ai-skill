# Observation — Evidence Acquisition Loop（Monitor before Story）

**Run ID**：2026-10-01-evidence-acquisition-loop
**Kind**：Phase 3 observation → **contract_gap**（workflow 吸收；產品 adapter 同輪落地 OCR×ASR gap）
**原則**：Probe 發現證據；Monitor 判可疑；Escalation 加大採集；不得等 Story 過少才回頭修 OCR。

## 觀察（sanitized）

硬字幕短劇 random-clip 成片：

| 訊號 | 觀察 |
| --- | --- |
| 早期 mechanical probe | `subtitle_like=0` 但 `dialogue_candidates>0` → hardsub 路由成立（45 已修 exclusion） |
| 成片中後段時間窗 | 畫面仍有對白硬字幕／口播，但節錄 cues 明顯偏稀 |
| 下游 | 若無 Monitor，會把「已選 OCR」當成 acquisition 完成 |

結論：45 解決「有沒有字幕」的 exclusion 塌縮；**時間窗／跨模態 coverage gap** 需要 Acquisition Loop。

## 目標行為

```text
ASR window density high ∧ OCR window density low
  → signal: ocr_asr_coverage_gap (suspicious)
  → escalate OCR (widen / denser / targeted window)
  → record before/after coverage
  → only then fuse / resolve / story
```

## 分類

| 標籤 | 判定 |
| --- | --- |
| `evidence_acquisition_monitor_missing` | contract_gap → [`46-evidence-acquisition-escalation-loop.md`](../46-evidence-acquisition-escalation-loop.md) |
