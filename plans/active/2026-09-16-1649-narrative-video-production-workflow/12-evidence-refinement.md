# Candidate: Evidence refinement loop

Companion to [`_plan.md`](_plan.md)。**候選獨立契約，不是 ASR／OCR parser，也不是 Phase 2 workflow 檔。** Parser 只採集；Refinement 負責連結、仲裁、驗證、學習。Phase 3 **不**建 `workflow/.../evidence-refinement.md`。  
觀察：[`evidence/2026-09-18-evidence-refinement.md`](evidence/2026-09-18-evidence-refinement.md)。

原則：不是 OCR／ASR 二選一，也不是讓 LLM 一次猜對。多種 evidence 經評估、仲裁、獨立審查與回饋，讓資料室逐步接近正確值。對齊 invariant 5：Need → Constraints → Feasible set → **Selection policy** → Selection。第一版用作品級 policy，**不**寫死全域權重。

```text
layer.observable   video / shot / frame / visual_text+geometry / ASR+time / speaker / face_track / audio
layer.linking      ASR↔OCR↔Face↔Speaker↔Shot；temporal／spatial overlap
layer.canonical    identity / dialogue / scene / place / time / event / naming
layer.narrative    script / template / matching / EDR / 成片
outer loop         observe → resolve → produce → verify → correct → learn policy
```

禁止：

- 只存 OCR `text`、不存 pixel／normalized box 與 persistence（幾何與持續性是一級 metadata）
- parser 宣布「這是浮水印／字幕」（那是 Classification；只允許 `role.candidate`）
- 高 persistence 邊角文字因像人名就進 ASR 人名仲裁
- `OCR weight = 1.0` 當全域真理（字幕延遲、誤讀、LOGO 都可能）
- LLM `confidence` 當 Decision SoT（只當另一條 evidence）
- Script 反過來改寫 evidence（script 是 consumer）
- 把 linking／仲裁塞進 parser
