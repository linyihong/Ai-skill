# Candidate: Visual text evidence

Companion to [`_plan.md`](_plan.md)。**候選，不是 workflow contract。** Locale 已有 `text_origin: ocr`；源片分析尚未要求一級 `visual_text_evidence`。  
觀察：[`evidence/2026-09-17-visual-text-evidence.md`](evidence/2026-09-17-visual-text-evidence.md)。

資料室保存的不是「OCR 結果」，而是：**某時間、某空間位置出現了某段文字。** OCR、字幕檔、scene-text、UI 抽字、人工標註都是 `acquisition.method`。

機械層一級 metadata：pixel `box` **與** `normalized_box`（跨解析度）、polygon、`frame_ref`。可從 box **算出** `spatial_features`／`temporal_features`。另可機械量測 `visual_style`（色盤／對比／描邊／陰影）——是 evidence feature，不是字幕判決：[`20-ocr-visual-style.md`](20-ocr-visual-style.md)。

`role.candidate` + `resolver: mechanical` + `final: null` 是正確形狀：機械「像字幕」≠ 系統宣告就是字幕。空間／時間清楚時不必 LLM vision。掃區與 fallback：[`13-mechanical-visual-text-probe.md`](13-mechanical-visual-text-probe.md)。右上角高覆蓋率的「林雪」不得因像人名就進 ASR 人名仲裁。仲裁見 [`12-evidence-refinement.md`](12-evidence-refinement.md)。下一層聚合：[`16-evidence-unit.md`](16-evidence-unit.md)。Phase 3 **不**改 workflow。
