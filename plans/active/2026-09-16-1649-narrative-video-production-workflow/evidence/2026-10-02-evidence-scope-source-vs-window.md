# Observation — Source vs window evidence scope

**Run ID**：2026-10-02-evidence-scope-source-vs-window
**Kind**：Phase 3 observation → **contract_gap**（workflow 吸收；產品 window-local fallback）
**原則**：Evidence 必須帶 scope；cache 命中 ≠ 當前 window 充分。

## 觀察（sanitized）

Random-clip／reference dub 路徑：

| 訊號 | 觀察 |
| --- | --- |
| Source probe | `hardsub=True`、`speech=True`（全片抽樣） |
| Cache | `dialogue_cues` 命中 |
| Window excerpt | `slice` 後 0 句 → raise「無可用對白」 |
| 人工抽幀 | 同 source 可見硬字幕對白 |

誤判：把 source-level「有字幕」與 window-level「這段有可用 cues」當成同一證據。

## 目標行為

```text
source probe（保留）→ cues cache → slice window
  → cues>0 → continue
  → cues==0 → cue_coverage(empty|…) + assess
       ├─ insufficient → record / skip（勿無條件 OCR）
       └─ suspicious → window_ocr → 補 cues → 再 slice
```

## 分類

| 標籤 | 判定 |
| --- | --- |
| `evidence_scope_source_vs_window_collapsed` | contract_gap → [`46-evidence-acquisition-escalation-loop.md`](../46-evidence-acquisition-escalation-loop.md) / [`text-evidence-acquisition-loop.md`](../../../workflow/narrative-video-production/text-evidence-acquisition-loop.md) |
