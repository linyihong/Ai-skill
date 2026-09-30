# Evidence: OCR semantic_anchor + ASR anomaly → semantic_reconstruction

**狀態**: contract extension (dogfood)
**Plan**: [`43-multimodal-text-evidence-resolution.md`](../43-multimodal-text-evidence-resolution.md)
**Workflow**: [`text-evidence-multimodal-resolution.md`](../../../workflow/narrative-video-production/text-evidence-multimodal-resolution.md)

## 觀察（去敏）

同窗三層 evidence：

| 層 | 例子 | 角色 |
| --- | --- | --- |
| observed ASR | 近音／不自然片語（如「…拍皮」） | acoustic／phonetic |
| observed OCR | 英文領域術語硬字幕（如 `Last Lot`） | lexical → **semantic_anchor** |
| candidate | 重建口播（如「…拍卖品」） | semantic_reconstruction |
| resolved | 被接受的 spoken text | narrative consumer 唯一 SoT |

`Last Lot` 不是逐字翻譯（Last＋Lot），而是拍賣領域固定術語 →
`semantic_anchor.domain=auction` → candidate「最後一件拍賣品」。
ASR「拍皮」語意異常＋音近，共同收斂到同一 candidate。

同一族（皆 **semantic_reconstruction**，不是 typo correction）：

- 不偿不偿我 → 補償補償我（OCR Compensate…）
- …拍皮 → …拍卖品（OCR Last Lot）
- 秦舍 → 禽獸（phonetic＋context）

## 三層命名（強制）

1. **observed** — 設備實際看到／聽到（不可覆寫）
2. **candidate** — 由 evidence 推導（含 provenance）
3. **resolved** — resolution 後供劇情分析使用

禁止把「拍皮→拍卖品」記成 silently corrected ASR；必須能追溯 observed vs resolved。

## Evidence convergence

- OCR anchor 非必要，但有則 strength 更高。
- 僅 ASR＋scene 也可產 candidate，confidence 應低於 ASR＋OCR anchor＋domain。
- 權重看收斂度，禁止全域 OCR>ASR。

## 不改

- 不新增 Phase／Agent；仍在 Multimodal Text Evidence Resolution 內。
- 不覆寫 OCR observed；anchor／candidate 另欄。
