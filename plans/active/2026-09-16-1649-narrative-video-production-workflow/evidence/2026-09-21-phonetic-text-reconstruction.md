# Observation — phonetic text reconstruction

**Run ID**：2026-09-21-phonetic-text-reconstruction  
**Kind**：Phase 3 observation（**不是** `contract_gap`；**不**重跑 ASR；**不**擴 workflow）  
**Extends**：[`2026-09-21-spoken-vs-subtitle-reconstruction.md`](2026-09-21-spoken-vs-subtitle-reconstruction.md)、[`2026-09-18-episode-vs-knowledge-accumulation.md`](2026-09-18-episode-vs-knowledge-accumulation.md)

## 三通道都有價值

| 通道 | 價值 | 缺口 |
| --- | --- | --- |
| ASR 近音錯字 | 音接近 spoken | grapheme 可能錯 |
| OCR 改詞 | 字對 | 可能 sanitization |
| 後半句語氣／句式 | constraint | 單獨還原不了被改的詞 |
| 正解 spoken | 目標 | 系統常沒有直接 evidence |

禁止再問「ASR 與 OCR 誰對」。問：constraint 能否共同生成 **phonetic candidate**，再重建 spoken text。

## 第二次 decoding（不是新 ASR）

Audio → ASR decoder → grapheme。再：grapheme → phonetic normalize → near-homophone set → context／OCR／vocabulary ranking。Raw ASR 不動。

`char_count` 是弱證據。優先保存 syllable／phoneme 序列（近音對、字面不同時仍可對齊）。

## LLM 最後一層

順序：alignment → phonetic candidates → OCR candidates → context／pattern（例 vocative `你這個 X`）→ sanitization／vocabulary knowledge → **candidate set** → LLM ranking／reconstruction → independent verify。禁止 ASR → LLM 直接出漂亮句子。

## Sanitization knowledge

`spoken_pattern` → `subtitle_pattern` 是可學習 transformation，走 Learning Inbox：本集 candidate → 多集 evidence_count → `supported` mapping。禁止開局寫死「字幕詞 = 髒話」。也可累積句式 pattern（負面稱呼槽），縮小候選空間。未驗證不得進 Story Truth。
