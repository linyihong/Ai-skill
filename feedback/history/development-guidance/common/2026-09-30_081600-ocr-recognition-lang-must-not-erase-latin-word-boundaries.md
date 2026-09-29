> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-30 - OCR recognition language must not erase Latin word boundaries

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。來自硬字幕英文 dogfood：畫面有詞間空白，OCR 快取成無空格長串（單框）。

#### One-line Summary

`recognition_language`（引擎模型，如 `ch`）≠ `observed_script`；Latin 觀測不得因 ch 路徑刪空格；單框無空格長串要標 boundary suspicious 並做幾何 recovery，derived 不可覆蓋 raw。

#### Human Explanation

短劇硬字幕常用中文 OCR 引擎讀畫面，但字幕可能是英文。若管線用 `ocr_lang=ch` → `latin=False` → 刪光空白，或引擎輸出單框黏字串，後面 Text Resolution／雙語 region 都無法還原詞界。正確做法：依觀測 script 正規化；保留 raw；幾何／ink-gap 產 derived candidate；詞典分詞只作 candidate。

#### Trigger

- OCR／dialogue cue 出現長 Latin 無空格串，而畫面可見空格
- `ocr_lang=ch`（或非拉丁 recognition）卻對英文 hardsub 去空白
- 單 `bbox`、`space_count=0`、長串仍被當最終 SoT

#### Evidence

- Tool: product dogfood；OCR cache 單框黏字串 → dialogue subtitle.observed
- Sanitized：recognition≠observed script；boundary loss at acquisition
- Episode titles／paths：`<PROJECT_ROOT>` only

#### Generalized Lesson

1. **recognition_language ≠ observed_script**：空白政策跟 script，不跟引擎語言。
2. **raw ≠ derived**：recovery 必須可重跑；禁止改寫／刪除 raw。
3. **Geometry before lexicon**：單框黏字串先 ink／char gap；詞典只 candidate。
4. **不開 LLM 斷詞 stage**：屬 OCR evidence acquisition contract gap，不是新 Phase／Agent。

#### Agent Action

改 OCR／硬字幕管線時：先查 normalize／join 是否 script-aware；Latin 單框無空格是否標 boundary 並嘗試 recovery；確認 raw 仍在快取。

#### Goal / Action / Validation

- Goal: Latin hardsub 詞界在 acquisition 層可追溯（raw＋boundary／derived）。
- Action: script-aware normalize／join；suspicious → geometry recovery；lexical 僅 candidate。
- Validation or reference source: 重跑後 sticky Latin 有空白 derived 或 boundary 標記；raw 仍為黏字串原文。

#### Applies When

- 硬字幕／OCR 管線可能讀到 Latin（含用 ch 模型）
- 下游消費 OCR text 當 subtitle／alignment 證據

#### Does Not Apply When

- 純 CJK OCR 且無需詞間空格
- 已有多 box 且空格正確的 Latin 行（只需 script-aware join 即可）

#### Validation

Dogfood：Latin 黏字串案例有 `boundary`／derived candidate；`raw_text` 未被覆寫；ch 路徑不再刪 Latin 空格。

#### Promotion Target

- `workflow/narrative-video-production/text-evidence-ocr-boundary.md`
- NVP plan companion `42-ocr-boundary-and-script-aware-normalization.md`

#### Promotion Record

本輪同步寫入 workflow／plan companion／evidence；產品 adapter 另改。

#### Required Linked Updates

- `feedback/history/development-guidance/common/README.md` index
- `feedback/history/development-guidance/README.md` recent／count
- workflow／plan：本 commit 一併更新

#### Linked Plans

- Narrative Video Production：OCR evidence acquisition／boundary
- Translation Decision：locale content 勿把黏字串當高信心源（間接）
