# Evidence: ASR lexical anomaly + OCR semantic reconstruction

**狀態**: contract_gap candidate
**Plan**: [`43-multimodal-text-evidence-resolution.md`](../43-multimodal-text-evidence-resolution.md)
**相關**: [`21-spoken-vs-subtitle-reconstruction.md`](../21-spoken-vs-subtitle-reconstruction.md)、[`22-phonetic-text-reconstruction.md`](../22-phonetic-text-reconstruction.md)、[`40-language-role-before-text-resolution.md`](../40-language-role-before-text-resolution.md)

## 觀察（去敏）

同窗硬字幕／口播對齊時：

| 通道 | Observed | 角色 |
| --- | --- | --- |
| ASR | 近音錯字片語（字面不自然的中文） | phonetic／acoustic evidence |
| OCR | 英文 dialogue subtitle（語義清晰的英文短語） | lexical evidence |
| Semantic | OCR→目標語 translation candidate | semantic evidence |
| Context | 衝突／索賠類場景約束 | semantic constraint |

合理 spoken reconstruction 同時需要：**音近、字幕語義、場景契合**。不是「OCR 比 ASR 高」單鍵覆蓋。

同一更高層問題亦見近音錯字＋字幕改寫（phonetic candidate vs sanitization candidate）—
見既有 phonetic／spoken-vs-subtitle evidence。

## 缺口

1. ASR 多只存字面，未強制拆 `phonetic`／`lexical_confidence`。
2. OCR 跨語言時缺少正式 `semantic_candidate`（translation 有 provenance）。
3. Resolution 缺 `resolution_reason`／三角 evidence refs。
4. 語意異常可機械偵測，但產品易靜默採用 ASR 字面或固定 OCR 勝出。

## 不改

- 不新增 Phase／Agent；不凍結為單一 composite「subtitle PASS」。
- 不把翻譯字幕直接當 spoken SoT。
- Narrative Video Workflow 主架構不因本證據改動。

## 下一步

Workflow：[`text-evidence-multimodal-resolution.md`](../../../workflow/narrative-video-production/text-evidence-multimodal-resolution.md)。
產品：在 Language Relation Gate 之後產出 candidates＋reason；dogfood 驗一例。
