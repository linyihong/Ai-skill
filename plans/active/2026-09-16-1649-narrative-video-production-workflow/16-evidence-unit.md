# Candidate: Evidence unit before event (stop expanding observable)

Companion to [`_plan.md`](_plan.md)。**候選，不是 workflow。** 真實 evidence chain 已接近既有架構，**不推倒**。observable 層已夠：shot／keyframe／visual text 幾何與 persistence／ASR+speaker+voice／face／links。**停止再加 detector。**  
觀察：[`evidence/2026-09-18-evidence-unit.md`](evidence/2026-09-18-evidence-unit.md)。劇情升格：[`15-story-evidence-vs-dialogue.md`](15-story-evidence-vs-dialogue.md)。

缺口不在缺證據，而在 **跨模態證據 → 劇情語義** 之間沒有聚合層：

```text
observable（凍結擴張）
  → text resolution + linking（誰跟誰有什麼關係；勿讓每個物件互指）
  → evidence_unit（同一時間窗的證據包；不是劇情）
  → narrative_window（跨 unit context；不是固定秒數）
  → event assembly + candidate relations
  → narrative relevance / story event / typed state claim
  → script / EDR
```

`role.candidate` + `resolver: mechanical` + `final: null` 已正確：機械像字幕 ≠ 系統宣告就是字幕。LLM 只做候選解讀；升格靠 traceability + independent verifier。Phase 3 **不**改 schema；先用現有 evidence 跑反例：閒聊、關鍵對白、跨 shot 對白、字幕／ASR 衝突、人物切鏡。

首份真實 run 已產出完整 evidence units，故此層方向 **PASS for continued
dogfood**。新瓶頸是下游 promotion，不是 unit 數量。Event 應補最小
`basis.*_refs` 以區分 dialogue-only 與真正 multimodal 支撐；見
[`17-story-promotion-gate.md`](17-story-promotion-gate.md)。不因此新增 detector。
Window／assembly／relation 的候選邊界見
[`19-text-resolution-and-narrative-assembly.md`](19-text-resolution-and-narrative-assembly.md)。
