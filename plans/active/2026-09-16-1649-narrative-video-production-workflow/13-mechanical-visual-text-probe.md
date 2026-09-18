# Candidate: Mechanical visual-text probe

Companion to [`_plan.md`](_plan.md)。**候選，不是 workflow。** 掃區是低成本高召回的搜尋策略，不是 canonical truth。  
觀察：[`evidence/2026-09-18-mechanical-visual-text-probe.md`](evidence/2026-09-18-mechanical-visual-text-probe.md)。證據形狀：[`09-visual-text-evidence.md`](09-visual-text-evidence.md)。歧義升級：[`12-evidence-refinement.md`](12-evidence-refinement.md)。

三層分開：

```text
① Mechanical Probe     掃哪裡（bands → coverage → expand → full-frame）
② Mechanical Features  位置／面積／持續／重複／timestamp → 清楚時的 role.candidate
③ LLM Classification   僅 ambiguous：這些字是什麼角色（不是哪裡有字）
```

第一版：bottom／top／optional center／fallback full-frame。有結果不立刻結束；先 coverage check，不足再擴大。清楚的右上小 box + 高 persistence → `role.candidate: watermark`、`resolver: mechanical`。中央短時「林雪」才 LLM vision。

禁止：LLM 決定 crop；LLM 一張圖改成 `subtitle_y: 0.65`；LLM 直接改全局 probe 規則。LLM 觀察只進 evidence，經 Observation → Accumulation → Policy candidate → Validation → Promotion 才改探針。換更強 Vision LLM 只換 ambiguous resolver，不重寫 pipeline。Phase 3 **不**改 workflow。
