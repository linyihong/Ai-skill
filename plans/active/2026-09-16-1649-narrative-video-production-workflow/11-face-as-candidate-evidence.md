# Candidate: Face as identity evidence link

Companion to [`_plan.md`](_plan.md)。**候選原則，不是 workflow schema。** Phase 3 **只留掛點**：`face_track` 可被 evidence link 引用。不實作 Face Recognition，也不改 `narrative-video-production`。  
觀察：[`evidence/2026-09-17-face-as-candidate-evidence.md`](evidence/2026-09-17-face-as-candidate-evidence.md)。

原則：Face 是 ASR／OCR 人名仲裁的 **candidate evidence**，不是 Identity 的直接判定器。

四層與 refinement loop：[`12-evidence-refinement.md`](12-evidence-refinement.md)。Face 停在 observable；linking 只掛 candidate，不裁定身份。

禁止：`OCR「林雪」+ Face 長得像 = 這人就是林雪`。同一演員可演不同角色；Face 不是 canonical identity key。

現在合法產物：`face_track_id`、keyframes；`embedding_ref`／`cluster_id` 可空。跨集 cluster 與人物圖像判別延後到 evidence 足夠。  
ASR／OCR parser 不吸收 Face；採集與 refinement 分開——不要做「ASR-OCR-Face 仲裁器」塞進 parser。
