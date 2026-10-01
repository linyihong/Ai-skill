# Evidence — source vs publish timebase

Status: contract refinement（observation → contract candidate）
Date: 2026-10-01
Plan: Narrative Video Production（Phase 3 dogfood；**不**開新大 phase）

## Observation（sanitized）

同一敘事內容產出兩條成片：一條近 1× 參考時長，一條經 constant-speed publish（例 rate≈1.5、時長縮短）。人工用**相同 wall-clock 秒數**對畫面時，發布片較早秒數看到不同鏡頭——這在 transform 下是預期行為，不是字幕／ASR／matching 必然錯位。

語音＋字幕若互鎖、僅口型偏快：屬 **presentation_comfort**，不是 **temporal_integrity** failure。

## Contract implication

寫入 workflow：[`source-publish-timebase.md`](../../../../workflow/narrative-video-production/source-publish-timebase.md)

Invariant（候選，未凍結進十條）：

> All source understanding and evidence extraction MUST use the canonical source timebase. Publish-time speed transforms MUST NOT be treated as new evidence.

## Product adapter note

Product 側既有「源片分析 + 末端 playback speed」傾向與此一致；缺口是 EDR／Timeline IR **明示** `timeline_transform` 與比較協議，以及 locale CPS 必須標 `timebase: publish`。具體 host／job／檔名留在 `<PROJECT_ROOT>`。

## Validation checklist

- [ ] 契約檔可從 NVP README／captions／assemble 連到
- [ ] EDR optional 欄位含 `timeline_transform`（或等價）
- [ ] 比較協議：`publish_t ↔ source_t = publish_t * rate`
- [ ] 禁止把 sped media 當 OCR／ASR evidence source（文件＋adapter 註記）
