# Execution SoT for narrative-video record fields

本目錄是 Phase 2 **workflow-local** field contract。`runtime_projection.enabled: false`。
Phase 4 才考慮 route；至少一次真實 EDR dogfood 後才有資格當 runtime candidate。

去敏可填示範：[`examples/sanitized-matching-and-edr.yaml`](examples/sanitized-matching-and-edr.yaml)（虛構系列，無真實片名／主機／金鑰）。
對白語義欄位：[`dialogue-semantic-context.yaml`](dialogue-semantic-context.yaml)（Phase 3 起 optional）。
自製口播時軸：[`speech-unit.yaml`](speech-unit.yaml)。Speech Unit 是語音 SoT；caption 是視覺投影，不得改 unit 原文。
字幕行數：cue `layout.lines` vs `max_lines`（見 [`../subtitle-layout.md`](../subtitle-layout.md)）。字級必須同時滿足 absolute、profile、`max_delta`；`text_integrity_gate` 必須 lossless。斷點 schema：[`break-candidate.yaml`](break-candidate.yaml)（clause 先於換行；`natural_boundary_first`；`prefer_at` 只在同一 tier）。`typography_scope: cue_uniform`。跨語系對齊 glyph 視覺高度。
